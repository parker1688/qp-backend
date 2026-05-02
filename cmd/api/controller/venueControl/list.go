package venueControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/venues"
	"context"
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/shopspring/decimal"
)

const (
	REDIS_KEY_BALANCE_F5      = "REDIS_KEY_BALANCE_F5_MERCHANT_%v_USERNAME_%v"
	REDIS_KEY_BALANCE_Recover = "REDIS_KEY_BALANCE_Recover_MERCHANT_%v_USERNAME_%v"
)

var venuesCache = cache.New(3*time.Minute, 3*time.Minute)

func GetVenusList(c *gin.Context) {
	gameType := c.DefaultQuery("game_type", "")

	merchantCode := modules.GetAgentDomainMerchantCodeByHeader(c)

	key := fmt.Sprintf("venue_list:%s:%s", merchantCode, gameType)
	cacheData, ok := venuesCache.Get(key)
	if !ok {
		data := venues.GetVenuesLineWithGameType(merchantCode, gameType)
		var newData []*vo.GetVenusListResp
		for _, v := range data {
			m := &vo.GetVenusListResp{
				VenueCode:  v.VenueCode,
				VenueName:  v.VenueName,
				IsMaintain: v.Status != 1, //1 正常
				ImgIcon:    v.ImgIcon,
				ImgBar:     v.ImgBar,
				Describe:   v.Describe,
				GameType:   v.GameType,
			}
			m.GameTypeImg = modules.FindByKeyFcVenueImg(&dos.FcVenueImg{VenueCode: v.VenueCode, MerchantCode: v.MerchantCode}, nil)
			newData = append(newData, m)
		}
		venuesCache.Set(key, newData, 10*time.Minute)
		response.SuccessJSON(c, newData)
	} else {
		data := cacheData.([]*vo.GetVenusListResp)
		response.SuccessJSON(c, data)
	}

}

func GetVenusBalances(c *gin.Context) {
	var jsonp vo.VenueBalanceReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	balanceF5Key := fmt.Sprintf(REDIS_KEY_BALANCE_F5, userInfo.MerchantCode, userInfo.UserName)
	isWait := global.G_REDIS.SetNX(context.Background(), balanceF5Key, "1", 3*time.Second).Val()
	venuesBalanceMap := map[string]float64{}
	data := venues.GetVenueCodeUniqueVenueLine(userInfo.MerchantCode)
	if isWait { //操作频繁
		var balanceRequest []*venues.VenueBalanceReq
		for _, v := range data {
			if v.InVenueCode != 0 && v.OutVenueCode != 0 { //禁止转入转出,不能获取余额
				continue
			}
			balanceRequest = append(balanceRequest, &venues.VenueBalanceReq{
				UserName:     userInfo.UserName,
				UserId:       userInfo.UserId,
				VenueCode:    v.VenueCode,
				Currency:     jsonp.Currency,
				MerchantCode: userInfo.MerchantCode,
			})
		}
		venueBalanceResp := venues.VenueBalancesTimeoutAsync(balanceRequest)

		for _, v := range venueBalanceResp {
			venuesBalanceMap[v.VenueCode] = decimal.NewFromFloat(v.Amount).Truncate(2).InexactFloat64()
		}
	}
	newData := make([]*vo.GetVenusBalancesResp, 0, len(data))
	for _, v := range data {
		balance, ok := venuesBalanceMap[v.VenueCode]
		m := &vo.GetVenusBalancesResp{
			VenueCode: v.VenueCode,
			Amount:    "", //超时未获取成功
		}
		if ok {
			m.Amount = tool.String(balance)
		}
		newData = append(newData, m)
	}

	response.SuccessJSON(c, newData)
}

func VenueRecover(c *gin.Context) {
	var jsonp vo.VenueBalanceReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	venueCode := global.G_REDIS.Get(context.Background(), fmt.Sprintf("venue-status:%s", userInfo.UserId)).Val()
	if venueCode != "" {
		var withdrawRequest *venues.VenueWithdrawRequest
		withdrawRequest = &venues.VenueWithdrawRequest{
			UserName:     userInfo.UserName,
			UserId:       userInfo.UserId,
			VenueCode:    venueCode,
			MerchantCode: userInfo.MerchantCode,
			IP:           c.ClientIP(),
			OptBy:        "",
			Amount:       0,
			Currency:     jsonp.Currency,
			IsAllAmount:  true, //金额全部回收
		}
		venues.VenueWithdraw(withdrawRequest)
	}
	response.SuccessMsgJSON(c, nil, "操作成功")
}

func VenueRecoverAll(c *gin.Context) {
	var jsonp vo.VenueBalanceReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}
	data := venues.GetVenuesLine(userInfo.MerchantCode)

	for _, v := range data {
		if v.OutVenueCode != 0 { //禁止提款
			continue
		}
		withdrawRequest := &venues.VenueWithdrawRequest{
			UserName:     userInfo.UserName,
			UserId:       userInfo.UserId,
			VenueCode:    v.VenueCode,
			MerchantCode: userInfo.MerchantCode,
			IP:           c.ClientIP(),
			OptBy:        "",
			Amount:       0,
			Currency:     jsonp.Currency,
			IsAllAmount:  true, //金额全部回收
		}
		venues.VenueWithdraw(withdrawRequest)
	}
	response.SuccessMsgJSON(c, nil, "操作成功")
}
func GameRecord(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		vo.BetRecordReq
	}{}

	jsonp.PageQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
	jsonp.PageQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.DateEnd = c.DefaultQuery("date_end", "")
	jsonp.DateStart = c.DefaultQuery("date_start", "")
	jsonp.GameType = c.DefaultQuery("game_type", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "0"))
	timeType := tool.Atoi(c.DefaultQuery("time_type", "-1"))

	if jsonp.PageNo == 0 {
		jsonp.PageNo = 1
	}
	if jsonp.PageSize == 0 {
		jsonp.PageSize = 10
	}
	if jsonp.PageNo < 1 {
		jsonp.PageNo = 1
	}
	if jsonp.PageSize < 1 {
		jsonp.PageSize = 10
	}
	if jsonp.PageSize > 100 {
		jsonp.PageSize = 100
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	k1 := 0
	st := ""
	et := ""
	if timeType == -1 {
		if jsonp.DateStart != "" {
			if len(jsonp.DateStart) < 11 {
				jsonp.DateStart += " 00:00:00"
			}
		}
		if jsonp.DateEnd != "" {
			if len(jsonp.DateEnd) < 11 {
				jsonp.DateEnd += " 23:59:59"
			}
		}
	} else {
		k1 = 1
		sTime1, eTime1 := tool.GetDayRange(time.Now(), timeType)
		st = sTime1.Format(tool.TimeLayout)
		et = eTime1.Format(tool.TimeLayout)
	}

	keyData := fmt.Sprintf("venue_record:%s:%s:%s:%v:%s:%s:%v:%v:%v", userInfo.UserId, jsonp.VenueCode, jsonp.GameType, timeType, jsonp.DateStart, jsonp.DateEnd, jsonp.PageNo, jsonp.PageSize, jsonp.Status)
	if k1 == 1 {
		keyData = fmt.Sprintf("venue_record:%s:%s:%s:%v:%s:%s:%v:%v:%v", userInfo.UserId, jsonp.VenueCode, jsonp.GameType, timeType, st, et, jsonp.PageNo, jsonp.PageSize, jsonp.Status)
	}
	keyTotal := keyData + ":total"
	keyBetAmount := keyData + ":betamount"
	keyNetAmount := keyData + ":netamount"
	resData, ok := venuesCache.Get(keyData)

	//global.G_LOG.Infof("venue betrecord------------------------1:%v, %v, %v", keyData, ok, resData)
	//global.G_LOG.Infof("venue betrecord------------------------2:%v, %v", jsonp.PageSize, jsonp.PageNo)
	if !ok {
		clientType := c.GetHeader(enmus.CLIENT_TYPE_HEADER)
		if len(clientType) == 0 {
			clientType = enmus.H5
		}
		language := c.GetHeader(enmus.LANGUAGE_HEADER)
		if len(language) == 0 {
			language = "zh-CN"
		}

		err := validator.New().Struct(jsonp)
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}

		var data []*dos.FcBetRecord
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		query := global.G_DB.WithContext(ctx).Model(&dos.FcBetRecord{}).
			Where("user_id = ? ", userInfo.UserId).
			Order("bet_time desc")
		query1 := global.G_DB.WithContext(ctx).Model(&dos.FcBetRecord{}).
			Where("user_id = ? ", userInfo.UserId)
		//global.G_LOG.Infof("venue betrecord------------------------3:%v", userInfo.UserId)

		if timeType == -1 {
			if jsonp.DateStart != "" {
				query.Where("bet_time >=?", jsonp.DateStart)
				query1.Where("bet_time >=?", jsonp.DateStart)
			}

			if jsonp.DateEnd != "" {
				query.Where("bet_time <=?", jsonp.DateEnd)
				query1.Where("bet_time <=?", jsonp.DateEnd)
			}
		} else {
			query.Where("bet_time BETWEEN ? AND ?", st, et)
			query1.Where("bet_time BETWEEN ? AND ?", st, et)
		}

		if jsonp.VenueCode != "" {
			query.Where("venue_code =?", jsonp.VenueCode)
			query1.Where("venue_code =?", jsonp.VenueCode)
		}

		if jsonp.GameType != "" {
			query.Where("game_type =?", jsonp.GameType)
			query1.Where("game_type =?", jsonp.GameType)
		}

		if jsonp.Status == 1 {
			query.Where("net_amount > ?", 0)
			query1.Where("net_amount > ?", 0)
		}

		if jsonp.Status == 2 {
			query.Where("net_amount < ?", 0)
			query1.Where("net_amount < ?", 0)
		}

		var totalBetAmount, totalNetAmount float64
		var total int64
		query.Count(&total)

		query1.Select("sum(bet_amount) as totalBetAmount").Scan(&totalBetAmount)
		query1.Select("sum(net_amount) as totalNetAmount").Scan(&totalNetAmount)
		//global.G_LOG.Infof("venue betrecord------------------------4:%v, %v", totalBetAmount, totalNetAmount)

		query.Offset((jsonp.PageNo - 1) * jsonp.PageSize).Limit(jsonp.PageSize).
			Scan(&data)

		newData := make([]*vo.BetRecordResp, 0, len(data))
		tool.JsonMapper(data, &newData)
		res := map[string]interface{}{
			"list":           newData,
			"totalBetAmount": totalBetAmount,
			"totalNetAmount": totalNetAmount,
		}
		if newData == nil {
			res["list"] = []string{}
		}
		venuesCache.Set(keyData, newData, 1*time.Minute)
		venuesCache.Set(keyTotal, total, 1*time.Minute)
		venuesCache.Set(keyBetAmount, totalBetAmount, 1*time.Minute)
		venuesCache.Set(keyNetAmount, totalNetAmount, 1*time.Minute)
		response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, res)
	} else {
		total, ok1 := venuesCache.Get(keyTotal)
		betAmount, ok2 := venuesCache.Get(keyBetAmount)
		netAmount, ok3 := venuesCache.Get(keyNetAmount)
		//global.G_LOG.Infof("venue betrecord------------------------2-3:%v, %v, %v", total, betAmount, netAmount)
		//global.G_LOG.Infof("venue betrecord------------------------2-4:%v, %v, %v", ok1, ok2, ok3)
		//global.G_LOG.Infof("venue betrecord------------------------2-5:%v, %v, %v", resData, jsonp.PageNo, jsonp.PageSize)
		if !ok1 || !ok2 || !ok3 {
			total = 0
			betAmount = 0
			netAmount = 0
		}
		total1 := total.(int64)
		netAmount = netAmount.(float64)
		betAmount = betAmount.(float64)
		newData := resData.([]*vo.BetRecordResp)
		res := map[string]interface{}{
			"list":           newData,
			"totalBetAmount": betAmount,
			"totalNetAmount": netAmount,
		}
		response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total1, res)
	}
}
