package controller

import (
	"bootpkg/cmd/api/controller/userControl"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/common/tool/plugins/kafkaMQ"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/channelData"
	"bootpkg/pkg/service/vip"
	"context"
	"errors"
	"fmt"
	"math"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func init() {
	for i := 0; i < 6; i++ {
		kafkaInitFun = append(kafkaInitFun, BetRecordDataConsumer)
	}
}

type UserFlow struct {
	UserId         string  `json:"user_id"`         // 用户ID
	UserName       string  `json:"user_name"`       // 用户账号
	VenueCode      string  `json:"venue_code"`      // 场馆 code
	GameType       string  `json:"game_type"`       // 游戏类型
	ValidBetamount float64 `json:"valid_betamount"` // 有效投注额
}

// UserInviteConsumer
//
//	@Description: 会员邀请消费者
func BetRecordDataConsumer(ctx context.Context) {
	conf := global.CONFIG.Mq.Kafka
	kafkaOpt.CreateConsumerGroup(strings.Split(conf.Addr, ","), &kafkaMQ.Kafka{
		Version: conf.Version,
		GroupId: "BetRecordData",
		Topic:   []string{channelData.Kakfa_Topic_Bet_Record_Data},
		Handler: func(message *sarama.ConsumerMessage) error {
			err := BetRecordDataRecord(message.Value)
			if err != nil {
				global.G_LOG.Errorf("BetRecordDataConsumer message: %s  err: %v", message.Value, err)
			}
			return nil
		},
	}, ctx)
}

func FilterRecord(data []*dos.FcBetRecord) ([]*dos.FcBetRecord, []*dos.FcBetRecordUnsettled) {
	data1 := []*dos.FcBetRecord{}
	data2 := []*dos.FcBetRecordUnsettled{}
	for _, v := range data {
		if v.IsSettled == 0 {
			data1 = append(data1, v)
		} else {
			v2 := dos.FcBetRecordUnsettled{}
			tool.JsonMapper(v, &v2)
			data2 = append(data2, &v2)
		}
	}
	return data1, data2
}

func BetRecordUnsettled(data []*dos.FcBetRecordUnsettled) {
	for _, betRecord := range data {
		if !TryLockBetOrderUnsettled(betRecord.OrderSn, betRecord.VenueCode) {
			continue
		}
		if IsExistBetOrderUnsettled(betRecord.OrderSn, betRecord.VenueCode) {
			UnlockBetOrderUnsettled(betRecord.OrderSn, betRecord.VenueCode)
			continue
		}
		err2, _ := CreateOrUpdateUnsettled(betRecord) //入库
		if err2 != nil {
			if isDuplicateKeyError(err2) {
				SetOrderSnUnsettled(betRecord.OrderSn, betRecord.VenueCode)
				UnlockBetOrderUnsettled(betRecord.OrderSn, betRecord.VenueCode)
				continue
			}
			global.G_LOG.Error(err2.Error())
			UnlockBetOrderUnsettled(betRecord.OrderSn, betRecord.VenueCode)
			continue
		}
		SetOrderSnUnsettled(betRecord.OrderSn, betRecord.VenueCode)
		UnlockBetOrderUnsettled(betRecord.OrderSn, betRecord.VenueCode)
	}
}

func DeleteUnsettledRecord(orderSn, venueCode string) {
	if venueCode == "FBTY" || venueCode == "HGTY" || venueCode == "VRCP" || venueCode == "DBDJ" || venueCode == "WALI" ||
		venueCode == "BBIN" || venueCode == "SBTY" || venueCode == "PPDZ" { //这几个才有未结算
		key := fmt.Sprintf("unsettle_record_del:%s:%s", orderSn, venueCode)
		cacheOrder := global.G_REDIS.Get(context.Background(), key).Val()
		if cacheOrder == "" {
			global.G_DB.Where("order_sn = ? and venue_code =?", orderSn, venueCode).Delete(&dos.FcBetRecordUnsettled{})
			global.G_REDIS.Set(context.Background(), key, 1, 3*time.Minute)
		}
	}
}

func BetRecordDataRecord(s []byte) error {
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			global.G_LOG.Error(err, "panic", "stack", "...\n"+string(buf))
		}
	}()

	global.G_LOG.Info("BetRecordDataRecord 消费到了数据")
	var data []*dos.FcBetRecord
	err := tool.JsonUnmarshal(s, &data)
	if err != nil {
		global.G_LOG.Error(err.Error())
		return nil
	}

	data1, data2 := FilterRecord(data)
	go BetRecordUnsettled(data2)

	userBetAmount := map[string]float64{}
	userFlowData := map[string]map[string]*UserFlow{}     // 第一个 key 为 user_id, 第二 key 为 场馆 code:gameType
	userGameTypeData := map[string]*vo.UserGameTypeData{} // 用户游戏类型数据, key 为 user_id:gameType
	//洗码有效转换比例：90%
	costRate := GetCostRate()

	// ============= 游戏任务相关以下
	userTaskParamMap := map[string][]modules.TaskActionParam{} // 用户任务参数表
	// ============= 游戏任务相关以上

	for _, betRecord := range data1 {
		if !TryLockBetOrder(betRecord.OrderSn, betRecord.VenueCode) {
			continue
		}
		//global.G_LOG.Infof("BetRecord -------------------------------1:%v, %v", betRecord.BetAmount, betRecord)
		//去重
		if IsExistBetOrder(betRecord.OrderSn, betRecord.VenueCode) {
			//global.G_LOG.Info("BetRecord -------------------------------2:%v, %v", betRecord.OrderSn, betRecord.VenueCode)
			UnlockBetOrder(betRecord.OrderSn, betRecord.VenueCode)
			continue
		}
		//global.G_LOG.Infof("BetRecord -------------------------------3:%v, %v", betRecord.BetAmount, betRecord)

		err2, record := CreateOrUpdate(betRecord) //入库
		//global.G_LOG.Infof("BetRecord -------------------------------4:%v, %v", betRecord.BetAmount, betRecord)

		if err2 != nil {
			if isDuplicateKeyError(err2) {
				SetOrderSn(betRecord.OrderSn, betRecord.VenueCode)
				DeleteUnsettledRecord(betRecord.OrderSn, betRecord.VenueCode)
				UnlockBetOrder(betRecord.OrderSn, betRecord.VenueCode)
				continue
			}
			//global.G_LOG.Infof("BetRecord -------------------------------5:%v, %v", betRecord.BetAmount, betRecord)
			global.G_LOG.Error(err2.Error())
			UnlockBetOrder(betRecord.OrderSn, betRecord.VenueCode)
			continue
		}
		SetOrderSn(betRecord.OrderSn, betRecord.VenueCode)
		UnlockBetOrder(betRecord.OrderSn, betRecord.VenueCode)
		DeleteUnsettledRecord(betRecord.OrderSn, betRecord.VenueCode)

		if betAmount, ok := userBetAmount[record.UserId]; ok && len(record.UserId) > 0 {
			userBetAmount[record.UserId] = decimal.NewFromFloat(record.BetAmount * (1 - costRate)).Add(decimal.NewFromFloat(math.Abs(betAmount))).Truncate(2).InexactFloat64()
		} else {
			if len(record.UserId) > 0 {
				userBetAmount[record.UserId] = record.BetAmount * (1 - costRate)
			}
		}

		// 打码
		userId := record.UserId

		record.ValidBetamount = record.ValidBetamount * (1 - costRate) //扣掉一个比例
		flowKey := record.VenueCode + ":" + record.GameType
		flowMap, ok := userFlowData[userId]
		if !ok {
			tmpFlow := &UserFlow{}
			tmpFlow.UserId = userId
			tmpFlow.UserName = record.UserName
			tmpFlow.VenueCode = record.VenueCode
			tmpFlow.GameType = record.GameType
			tmpFlow.ValidBetamount = record.ValidBetamount

			flowMap = map[string]*UserFlow{}
			flowMap[flowKey] = tmpFlow
			userFlowData[userId] = flowMap
		} else {
			tmpFlow, ok2 := flowMap[flowKey]
			if !ok2 {
				tmpFlow = &UserFlow{}
				tmpFlow.UserId = userId
				tmpFlow.UserName = record.UserName
				tmpFlow.VenueCode = record.VenueCode
				tmpFlow.GameType = record.GameType
				tmpFlow.ValidBetamount = record.ValidBetamount

				flowMap[flowKey] = tmpFlow
			} else {
				tmpFlow.ValidBetamount += record.ValidBetamount
				flowMap[flowKey] = tmpFlow
			}
		}
		userFlowData[userId] = flowMap

		// 用户游戏分类数据
		gameTypeKey := userId + ":" + record.GameType
		tmpGameTypeData, ok := userGameTypeData[gameTypeKey]
		if !ok {
			tmpGameTypeData = &vo.UserGameTypeData{}
			tmpGameTypeData.UserId = userId
			tmpGameTypeData.UserName = record.UserName
			tmpGameTypeData.GameType = record.GameType
			tmpGameTypeData.ValidBetamount = record.ValidBetamount
			tmpGameTypeData.BetAmount = record.BetAmount
			tmpGameTypeData.NetAmount = record.NetAmount
			tmpGameTypeData.MerchantCode = record.MerchantCode
			userGameTypeData[gameTypeKey] = tmpGameTypeData
		} else {
			tmpGameTypeData.ValidBetamount += record.ValidBetamount
			tmpGameTypeData.BetAmount += record.BetAmount
			tmpGameTypeData.NetAmount += record.NetAmount
			userGameTypeData[gameTypeKey] = tmpGameTypeData
		}

		// ============= 游戏任务相关以下
		userTaskParamMap = modules.ToUserTaskParamsByBetRecord(userTaskParamMap, record)
		// ============= 游戏任务相关以上
	}

	// 用户打码
	UpdateUserFlow(userFlowData, userGameTypeData)

	for userId, betAmount := range userBetAmount {
		//存储用户流水
		global.G_REDIS.IncrByFloat(context.Background(), fmt.Sprintf(enmus.UserTotalBetAmountKey, userId), betAmount)

		isVipUpgradation := false
		var needBetAmount float64
		nextLevelAmountDataExists := global.G_REDIS.Exists(context.Background(), fmt.Sprintf(enmus.UserVipNextLevelAmountKey, userId)).Val()

		//下一级所需流水缓存不存在
		if nextLevelAmountDataExists == 0 {
			var minLevel int
			global.G_DB.Model(&dos.FcVip{}).Select("min(level) as minLevel").Scan(&minLevel)
			//下一级所需流水
			nexVip := modules.FindByKeyFcVipFirst(&dos.FcVip{Level: minLevel + 1})
			if betAmount-nexVip.MinBetAmount > 0 {
				isVipUpgradation = true
			}

		} else {
			needBetAmount, err = global.G_REDIS.Get(context.Background(), fmt.Sprintf(enmus.UserVipNextLevelAmountKey, userId)).Float64()
			if err != nil {
				global.G_LOG.Error(err.Error())
			}

			if betAmount-needBetAmount >= 0 {
				isVipUpgradation = true
			} else {
				global.G_REDIS.Set(context.Background(),
					fmt.Sprintf(enmus.UserVipNextLevelAmountKey,
						userId),
					fmt.Sprintf("%0.2f", decimal.NewFromFloat(needBetAmount).Sub(decimal.NewFromFloat(betAmount)).Truncate(2).InexactFloat64()), -1)
			}
		}

		//触发升级
		if isVipUpgradation {
			for {
				user := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{UserId: userId})
				totalBetAmount, err := global.G_REDIS.Get(context.Background(),
					fmt.Sprintf(enmus.UserTotalBetAmountKey, user.UserId)).Float64()
				if err != nil {
					global.G_LOG.Error()
					return nil
				}
				global.G_LOG.Infof("用户 %s 触发了升级  本次流水 %0.2f  升级所需流水 %0.2f", user.UserName, betAmount, needBetAmount)

				isok, serVipNextLevelAmount := vip.UserVipUpgradation(user, 0, totalBetAmount)
				if !isok {
					break
				}
				if serVipNextLevelAmount != 0 {
					break
				}
				global.G_LOG.Infof("用户 %s 循环升级  本次流水 %0.2f  升级所需流水 %0.2f", user.UserName, totalBetAmount, serVipNextLevelAmount)
			}
		}
	}

	// ============= 游戏任务相关以下
	for userId, taskParams := range userTaskParamMap {
		go modules.DoUserTaskAction(userId, taskParams, true)
	}
	// ============= 游戏任务相关以上

	return nil
}

const (
	GameInfoRedisKey         = "PullOrderGameInfoRedisKey::%s::%s"
	VenueAccountUserRedisKey = "PullOrderVenueAccountUserRedisKey::%s::%s"
	//VenuePGDZversionRedisKey = "PullOrderVenuePGDZversionRedisKey"
)

func CreateOrUpdate(betRecord *dos.FcBetRecord) (error, *dos.FcBetRecord) {
	betRecord.Id = tool.SnowflakeId()
	//补充用户信息
	venueUser := AccountInfo(betRecord.VenueCode, betRecord.Account)
	if venueUser.UserId == "" {
		global.G_LOG.Error(errors.New(betRecord.VenueCode + "场馆的" + betRecord.Account + "用户不存在"))
		return errors.New("用户不存在"), nil
	}

	betRecord.UserId = venueUser.UserId
	betRecord.UserName = venueUser.UserName
	betRecord.MerchantCode = venueUser.MerchantCode
	betRecord.Account = venueUser.Account

	//补充游戏名称内容
	gameinfo := GameInfo(betRecord.VenueCode, betRecord.GameCode)
	//global.G_LOG.Infof("BetRecord -------------------------------3-8:%v, %v", betRecord.BetAmount, betRecord)

	if betRecord.GameName == "" {
		betRecord.GameName = gameinfo.GameName
	}

	if betRecord.GameType == "" {
		betRecord.GameType = gameinfo.GameType
	}

	betRecord.CreateTime = automaticType.Time(time.Now())
	betRecord.UpdateTime = automaticType.Time(time.Now())
	//global.G_LOG.Info("BetRecord -------------------------------3-9:%v, %v", betRecord.BetAmount, betRecord)
	err := global.G_DB.Model(&dos.FcBetRecord{}).Create(betRecord).Error
	if err != nil {
		global.G_LOG.Error(err.Error())
		return err, nil
	}

	return nil, betRecord
}

func CreateOrUpdateUnsettled(betRecord *dos.FcBetRecordUnsettled) (error, *dos.FcBetRecordUnsettled) {
	betRecord.Id = tool.SnowflakeId()
	//补充用户信息
	venueUser := AccountInfo(betRecord.VenueCode, betRecord.Account)
	if venueUser.UserId == "" {
		global.G_LOG.Error(errors.New(betRecord.VenueCode + "场馆的" + betRecord.Account + "用户不存在"))
		return errors.New("用户不存在"), nil
	}

	betRecord.UserId = venueUser.UserId
	betRecord.UserName = venueUser.UserName
	betRecord.MerchantCode = venueUser.MerchantCode
	betRecord.Account = venueUser.Account

	//补充游戏名称内容
	gameinfo := GameInfo(betRecord.VenueCode, betRecord.GameCode)
	//global.G_LOG.Infof("BetRecord -------------------------------3-8:%v, %v", betRecord.BetAmount, betRecord)

	if betRecord.GameName == "" {
		betRecord.GameName = gameinfo.GameName
	}

	if betRecord.GameType == "" {
		betRecord.GameType = gameinfo.GameType
	}

	betRecord.CreateTime = automaticType.Time(time.Now())
	betRecord.UpdateTime = automaticType.Time(time.Now())
	//global.G_LOG.Info("BetRecord -------------------------------3-9:%v, %v", betRecord.BetAmount, betRecord)
	err := global.G_DB.Model(&dos.FcBetRecordUnsettled{}).Create(betRecord).Error
	if err != nil {
		global.G_LOG.Error(err.Error())
		return err, nil
	}

	return nil, betRecord
}

type VenueGameInfo struct {
	GameName string `json:"gameName"`
	GameType string `json:"gameType"`
}

// 后期从缓存里面取
func GameInfo(venueCode, gameCode string) VenueGameInfo {
	var venueGameInfo VenueGameInfo
	//key := fmt.Sprintf(GameInfoRedisKey, venueCode, gameCode)
	//gameInfo := global.G_REDIS.Get(context.Background(), key).Val()
	//
	//if gameInfo != "" {
	//	tool.JsonUnmarshal([]byte(gameInfo), &venueGameInfo)
	//	if venueGameInfo.GameType != "" {
	//		return venueGameInfo
	//	}
	//}

	if gameCode != "" {
		var venueGame dos.FcVenueGame
		global.G_DB_SHARDING.Model(&dos.FcVenueGame{}).Where("venue_code=? AND game_code=?", venueCode, gameCode).Take(&venueGame)
		// 有可能没有配置，则用场馆类型
		if venueGame.Id == "" {
			merchantVenue := modules.FindByKeyFcMerchantVenueFirst(&dos.FcMerchantVenue{
				VenueCode: venueCode,
			})
			// 场馆 gameType 可能存在多种类型，取第一个
			gameTypeArr := strings.Split(merchantVenue.GameType, ",")
			if len(gameTypeArr) > 1 {
				venueGameInfo.GameType = gameTypeArr[0]
			} else {
				venueGameInfo.GameType = merchantVenue.GameType
			}
		} else {
			venueGameInfo.GameName = venueGame.GameName
			venueGameInfo.GameType = venueGame.GameType
		}
	} else {
		merchantVenue := modules.FindByKeyFcMerchantVenueFirst(&dos.FcMerchantVenue{
			VenueCode: venueCode,
		})
		// 场馆 gameType 可能存在多种类型，取第一个
		gameTypeArr := strings.Split(merchantVenue.GameType, ",")
		if len(gameTypeArr) > 1 {
			venueGameInfo.GameType = gameTypeArr[0]
		} else {
			venueGameInfo.GameType = merchantVenue.GameType
		}
	}
	return venueGameInfo
}

func AccountInfo(venueCode, account string) dos.FcVenueUser {
	//global.G_LOG.Info("AccountInfo -------------------------------1:%v, %v", venueCode, account)

	if venueCode == enmus.LYQP || venueCode == enmus.KYQP {
		arr := strings.Split(account, "_")
		//global.G_LOG.Info("AccountInfo -------------------------------1-1:%v", arr, arr[1])
		if len(arr) > 1 {
			account = arr[1]
		}
	}
	//global.G_LOG.Info("AccountInfo -------------------------------1-2:%v", account)

	venueAccountUserRedisKey := fmt.Sprintf(VenueAccountUserRedisKey, account, venueCode)
	accountInfo := global.G_REDIS.Get(context.Background(), venueAccountUserRedisKey).Val()
	//global.G_LOG.Info("AccountInfo -------------------------------2:%v", accountInfo)

	var venueUser dos.FcVenueUser
	if account != "" && accountInfo != "" {
		err := tool.JsonUnmarshal([]byte(accountInfo), &venueUser)
		if err != nil {
			global.G_LOG.Error(err.Error())
		}

		// tyqp 的账号信息是 dvtony001@1752
		if venueCode == enmus.TYQP {
			arr := strings.Split(venueUser.Account, "@")
			if len(arr) > 1 {
				venueUser.Account = arr[0]
			}
		}

		//global.G_LOG.Info("AccountInfo -------------------------------3:%v", venueUser)
		return venueUser
	}
	// tyqp 的账号信息是 dvtony001@1752
	if venueCode == enmus.TYQP {
		arr := strings.Split(venueUser.Account, "@")
		if len(arr) > 1 {
			account = arr[0]
			venueUser.Account = arr[0]
		}
	}
	//global.G_LOG.Info("AccountInfo -------------------------------4:%v", venueUser)
	global.G_DB.Model(&dos.FcVenueUser{}).Where("venue_code=? AND account=?", venueCode, account).Take(&venueUser)
	venueUserJson, err := tool.MarshalToString(venueUser)
	if err != nil {
		global.G_LOG.Error(err.Error())
	}
	global.G_REDIS.Set(context.Background(), venueAccountUserRedisKey, venueUserJson, time.Duration(24)*time.Hour)
	return venueUser
}

func IsExistBetOrder(orderSn, venueCode string) bool {
	if GetOrderSn(orderSn, venueCode) == "1" {
		return true
	}
	var record *dos.FcBetRecord
	global.G_DB.Model(&dos.FcBetRecord{}).Where("order_sn = ? and venue_code = ?", orderSn, venueCode).Take(&record)
	if len(record.Id) > 0 {
		SetOrderSn(orderSn, venueCode)
		return true
	}
	return false
}

func SetOrderSn(orderSn, venueCode string) {
	key := fmt.Sprintf("BetRecord:%s:%s", orderSn, venueCode)
	global.G_REDIS.Set(context.Background(), key, 1, 10*time.Minute)
}

func TryLockBetOrder(orderSn, venueCode string) bool {
	key := fmt.Sprintf("BetRecordLock:%s:%s", orderSn, venueCode)
	return global.G_REDIS.SetNX(context.Background(), key, 1, 30*time.Second).Val()
}

func UnlockBetOrder(orderSn, venueCode string) {
	key := fmt.Sprintf("BetRecordLock:%s:%s", orderSn, venueCode)
	global.G_REDIS.Del(context.Background(), key)
}

func GetOrderSn(orderSn, venueCode string) string {
	key := fmt.Sprintf("BetRecord:%s:%s", orderSn, venueCode)
	res := global.G_REDIS.Get(context.Background(), key).Val()
	return res
}

func IsExistBetOrderUnsettled(orderSn, venueCode string) bool {
	if GetOrderSnUnsettled(orderSn, venueCode) == "1" {
		return true
	}
	var record *dos.FcBetRecordUnsettled
	global.G_DB.Model(&dos.FcBetRecordUnsettled{}).Where("order_sn = ? and venue_code = ?", orderSn, venueCode).Take(&record)
	if len(record.Id) > 0 {
		SetOrderSnUnsettled(orderSn, venueCode)
		return true
	}
	return false
}

func SetOrderSnUnsettled(orderSn, venueCode string) {
	key := fmt.Sprintf("BetRecordUnsettled:%s:%s", orderSn, venueCode)
	global.G_REDIS.Set(context.Background(), key, 1, 10*time.Minute)
}

func TryLockBetOrderUnsettled(orderSn, venueCode string) bool {
	key := fmt.Sprintf("BetRecordUnsettledLock:%s:%s", orderSn, venueCode)
	return global.G_REDIS.SetNX(context.Background(), key, 1, 30*time.Second).Val()
}

func UnlockBetOrderUnsettled(orderSn, venueCode string) {
	key := fmt.Sprintf("BetRecordUnsettledLock:%s:%s", orderSn, venueCode)
	global.G_REDIS.Del(context.Background(), key)
}

func GetOrderSnUnsettled(orderSn, venueCode string) string {
	key := fmt.Sprintf("BetRecordUnsettled:%s:%s", orderSn, venueCode)
	res := global.G_REDIS.Get(context.Background(), key).Val()
	return res
}

func isDuplicateKeyError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "duplicate entry") || strings.Contains(errMsg, "error 1062")
}

// @Description: 重置版本号
// @param recordVersion
func ResetVersion(venueCode string, recordVersion int64) {
	key := fmt.Sprintf(enmus.VenueVersionRedisKey, venueCode)
	vesion := global.G_REDIS.Get(context.Background(), key).Val()
	var newVersion int64

	if vesion != "" {
		vesionInt, err := strconv.ParseInt(vesion, 10, 64)
		if err != nil {
			global.G_LOG.Error(err.Error())
		}
		if recordVersion > vesionInt {
			newVersion = recordVersion
		}
	} else {
		newVersion = recordVersion
	}
	if newVersion != 0 {
		err := global.G_REDIS.Set(context.Background(), key, fmt.Sprintf("%d", newVersion), time.Duration(48)*time.Hour).Err()
		if err != nil {
			global.G_LOG.Errorf("redis set %s %s err:%s", key, fmt.Sprintf("%d", newVersion), err.Error())
		}
	}
}

//func ResetLastBetId(venueCode string, betId string) {
//	key := fmt.Sprintf(enmus.VenuLastUidRedisKey, venueCode)
//
//	err := global.G_REDIS.Set(context.Background(), key, betId, time.Duration(48)*time.Hour).Err()
//	if err != nil {
//		global.G_LOG.Errorf("redis set %s %s err:%s", key, betId, err.Error())
//	}
//
//}

// 更新用户打码
func UpdateUserFlow(userFlowData map[string]map[string]*UserFlow, userGameTypeData map[string]*vo.UserGameTypeData) {
	// 更新用户的游戏类型数据
	for _, v := range userGameTypeData {
		if v.UserId == "" { // 有时候共用一个环境 user_id 为空
			continue
		}

		tmpRow := dos.FcUserGameReport{}
		nowTime := automaticType.Now()
		gameCount := 1 // 游戏局数

		err := global.G_DB.Model(&dos.FcUserGameReport{}).Where("user_id = ? AND game_type = ?", v.UserId, v.GameType).First(&tmpRow).Error
		if err != nil {
			// 如果不存在就插入
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tmpRow.UserId = v.UserId
				tmpRow.UserName = v.UserName
				tmpRow.GameType = v.GameType
				tmpRow.GameCount = gameCount
				tmpRow.BetAmount = v.BetAmount
				tmpRow.ValidBetamount = v.ValidBetamount
				tmpRow.NetAmount = v.NetAmount
				tmpRow.MerchantCode = v.MerchantCode
				tmpRow.CreateBy = "system"
				tmpRow.UpdateBy = "system"
				tmpRow.CreateTime = nowTime
				tmpRow.UpdateTime = nowTime

				err = global.G_DB.Save(&tmpRow).Error
				if err != nil {
					global.G_LOG.Errorf("UpdateuserGameTypeData username: %s gameType: %s insert err: %s", v.UserName, v.GameType, err.Error())
					continue
				}
				continue
			}
			continue
		}

		rowMap := map[string]interface{}{
			"bet_amount":      gorm.Expr("bet_amount + ?", v.BetAmount),
			"valid_betamount": gorm.Expr("valid_betamount + ?", v.ValidBetamount),
			"net_amount":      gorm.Expr("net_amount + ?", v.NetAmount),
			"game_count":      gorm.Expr("game_count + ?", gameCount),
			"update_time":     nowTime,
		}

		// 存在则更新数据
		err = global.G_DB.Model(&dos.FcUserGameReport{}).Where("user_id = ? AND game_type = ?", v.UserId, v.GameType).Updates(rowMap).Error
		if err != nil {
			global.G_LOG.Errorf("UpdateuserGameTypeData username: %s gameType: %s update err: %s", v.UserName, v.GameType, err.Error())
			continue
		}
	}

	// userFlowData 第一个 key 为 user_id, 第二 key 为 场馆 code:gameType
	for userId, flowMap := range userFlowData {
		if userId == "" { // 有时候共用一个环境 user_id 为空
			continue
		}

		flowKey := fmt.Sprintf(enmus.UserRebateFlow, userId)
		for key, value := range flowMap {
			_, err := global.G_REDIS.HIncrByFloat(context.Background(), flowKey, key, value.ValidBetamount).Result()
			if err != nil {
				global.G_LOG.Errorf("redis UpdateUserFlow username: %s venueCode: %s  gameType: %s HIncrByFloat err: %s",
					value.UserName, value.VenueCode, value.GameType, err.Error())
				continue
			}
		}

		// 设置 key 的过期时间为 3 个月
		err := global.G_REDIS.Expire(context.Background(), flowKey, 24*time.Hour*90).Err()
		if err != nil {
			global.G_LOG.Errorf("redis UpdateUserFlow setKey: %s expire err: %s", flowKey, err.Error())
		}
	}
}

func GetCostRate() float64 {
	rate := 1.0
	rateStr := global.G_REDIS.Get(context.Background(), "venue_cost_rate").Val()
	if rateStr != "" {
		f, err := strconv.ParseFloat(rateStr, 64)
		if err != nil {
			global.G_LOG.Errorf("转换错误:%v", err)
			return rate
		}
		rate = f
	} else {
		rate = userControl.GetDictRebateCostRateValue()
		global.G_REDIS.Set(context.Background(), "venue_cost_rate", rate, 10*time.Minute)
	}
	return rate
}
