package controller

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/common/tool/plugins/kafkaMQ"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	vos "bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/channelData"
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/IBM/sarama"
)

func init() {
	for i := 0; i < 6; i++ {
		kafkaInitFun = append(kafkaInitFun, MessageConsumer)
	}
}

type UserIdName struct {
	UserId   string `gorm:"column:user_id;primary_key" json:"user_id" form:"user_id" uri:"user_id" ` // 用户Id
	UserName string `gorm:"column:user_name" json:"user_name" form:"user_name" uri:"user_name" `     // 用户账号
}

// UserInviteConsumer
//
//	@Description: 站内信
func MessageConsumer(ctx context.Context) {
	conf := global.CONFIG.Mq.Kafka
	kafkaOpt.CreateConsumerGroup(strings.Split(conf.Addr, ","), &kafkaMQ.Kafka{
		Version: conf.Version,
		GroupId: "SiteMessageData",
		Topic:   []string{channelData.Kakfa_Topic_User_Site_Msg_Data},
		Handler: func(message *sarama.ConsumerMessage) error {
			err := SiteMsgHandler(message.Value)
			if err != nil {
				global.G_LOG.Errorf("SiteMsgHandler message: %s  err: %v", message.Value, err)
			}
			return nil
		},
	}, ctx)
}

func SiteMsgHandler(s []byte) error {
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

	global.G_LOG.Info("SiteMsgHandler consumers")
	global.G_LOG.Infof("SiteMsgHandler: %v", string(s))

	siteMsg := vos.SiteMsgVO{}
	err := tool.JsonUnmarshal(s, &siteMsg)
	if err != nil {
		global.G_LOG.Info("SiteMsgHandler JsonUnmarshal s: %v err: %v", string(s), err)
		return err
	}

	// 如果是要发送所有商户信息
	if siteMsg.NotifyType == 1 && siteMsg.MerchantCode != "" {
		err = AllUserSendMsg(&siteMsg)
		if err != nil {
			return err
		}
	} else if siteMsg.NotifyType == 2 {
		err = PartUserSendMsg(&siteMsg)
		if err != nil {
			return err
		}
	}

	return nil
}

func AllUserSendMsg(siteMsg *vos.SiteMsgVO) error {
	if siteMsg.MsgId == "" {
		tmpStr := "send all merchantCode siteMsg.MsgId is empty"
		global.G_LOG.Errorf(tmpStr)
		return errors.New(tmpStr)
	}

	// 根据msgId查询信息
	queryParam := dos.FcSiteMessage{}
	queryParam.Id = siteMsg.MsgId
	row := modules.FindByKeyFcSiteMessageFirst(&queryParam)
	if row.Id == "" {
		tmpStr := fmt.Sprintf("query msgId: %v not exist", siteMsg.MsgId)
		global.G_LOG.Errorf(tmpStr)
		return errors.New(tmpStr)
	}

	// 查询今天登录的用户进行发送站内信, 每次查询 100 条数据
	todayZero, _ := tool.TodayStartEndDate()
	dbQuery := global.G_DB.Model(&dos.FcUserMaterial{}).Select("user_id", "user_name").Where("login_status = 0 AND merchant_code = ? AND update_time >= ?", siteMsg.MerchantCode, todayZero)
	var count int64
	err := dbQuery.Count(&count).Error
	if err != nil {
		tmpStr := fmt.Sprintf("query merchantCode: %v login_status=0 count err: %v", siteMsg.MerchantCode, err)
		global.G_LOG.Errorf(tmpStr)
		return errors.New(tmpStr)
	}
	if count < 1 {
		tmpStr := fmt.Sprintf("query merchantCode: %v login_status=0 count is 0", siteMsg.MerchantCode)
		global.G_LOG.Errorf(tmpStr)
		return nil
	}

	page := int64(0)
	pageSize := int64(100)
	loop := count / pageSize
	if count%pageSize != 0 {
		loop = loop + 1
	}

	for page = int64(0); page < loop; page++ {
		dataSlice := []*UserIdName{}
		err = dbQuery.Offset((int(page)) * int(pageSize)).Limit(int(pageSize)).Scan(&dataSlice).Error
		if err != nil {
			tmpStr := fmt.Sprintf("query userIdName merchantCode: %v page: %v pageSize: %v err: %v", siteMsg.MerchantCode, page, pageSize, err)
			global.G_LOG.Errorf(tmpStr)
		}

		nowTime := automaticType.Now()
		// 插入站内信
		for _, v := range dataSlice {
			userSiteMsg := &dos.FcUserSiteMessage{}
			userSiteMsg.UserId = v.UserId
			userSiteMsg.UserName = v.UserName

			userSiteMsg.MsgId = row.Id
			userSiteMsg.Title = row.Title
			userSiteMsg.Content = row.Content
			userSiteMsg.MsgIdType = row.MsgIdType
			userSiteMsg.MsgType = row.MsgType
			userSiteMsg.NotifyType = row.NotifyType
			userSiteMsg.Language = row.Language
			userSiteMsg.MerchantCode = siteMsg.MerchantCode
			userSiteMsg.CreateBy = row.CreateBy
			userSiteMsg.UpdateBy = row.UpdateBy

			userSiteMsg.DelStatus = 1  // 未删除
			userSiteMsg.ReadStatus = 1 // 未读
			userSiteMsg.UpdateTime = nowTime
			userSiteMsg.CreateTime = nowTime

			err = global.G_DB.Create(userSiteMsg).Error
			if err != nil {
				global.G_LOG.Errorf("AllUserSendMsg Insert userSiteMsg err: %v", err)
				continue
			}

			// 首先判断该用户的站内信消息是否存在
			userMsgKey := fmt.Sprintf(enmus.UserSiteMessageKey, row.Id, v.UserId)
			err = global.G_REDIS.SetEx(context.Background(), userMsgKey, 1, 8*24*time.Hour).Err()
			if err != nil {
				global.G_LOG.Errorf("AllUserSendMsg cash userSiteMsg err: %v", err)
			}
		}
	}

	return nil
}

// 部分用户发送站内信
func PartUserSendMsg(siteMsg *vos.SiteMsgVO) error {
	userSiteMsg := &dos.FcUserSiteMessage{}
	userSiteMsg.UserId = siteMsg.UserId
	userSiteMsg.UserName = siteMsg.UserName

	nowTime := automaticType.Now()

	userSiteMsg.MsgId = siteMsg.MsgId
	userSiteMsg.MsgIdType = siteMsg.MsgIdType
	userSiteMsg.MsgType = siteMsg.MsgType
	userSiteMsg.Title = siteMsg.Title
	userSiteMsg.Content = siteMsg.Content
	userSiteMsg.NotifyType = siteMsg.NotifyType
	userSiteMsg.Language = siteMsg.Language
	userSiteMsg.MerchantCode = siteMsg.MerchantCode
	userSiteMsg.CreateBy = siteMsg.CreateBy
	userSiteMsg.UpdateBy = siteMsg.UpdateBy

	userSiteMsg.DelStatus = 1  // 未删除
	userSiteMsg.ReadStatus = 1 // 未读
	userSiteMsg.UpdateTime = nowTime
	userSiteMsg.CreateTime = nowTime

	err := global.G_DB.Create(userSiteMsg).Error
	if err != nil {
		global.G_LOG.Errorf("PartUserSendMsg Insert userSiteMsg err: %v", err)
		return err
	}

	return nil
}
