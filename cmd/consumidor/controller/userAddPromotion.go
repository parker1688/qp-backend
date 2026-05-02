package controller

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/common/tool/plugins/kafkaMQ"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/service/channelData"
	"context"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/IBM/sarama"
	"gorm.io/gorm"
)

func init() {
	kafkaInitFun = append(kafkaInitFun, PromotionConsumer)
}

// UserInviteConsumer
//
//	@Description: 充值消费
func PromotionConsumer(ctx context.Context) {
	conf := global.CONFIG.Mq.Kafka
	kafkaOpt.CreateConsumerGroup(strings.Split(conf.Addr, ","), &kafkaMQ.Kafka{
		Version: conf.Version,
		GroupId: "userPromotionConsumer",
		Topic:   []string{channelData.Kakfa_Topic_User_Promotion},
		Handler: func(message *sarama.ConsumerMessage) error {
			err := PromotionAccumulation(message.Value)
			if err != nil {
				global.G_LOG.Errorf("UserInviteConsumer message: %s  err: %v", message.Value, err)
			}
			return nil
		},
	}, ctx)
}

// 优惠累计报表
func PromotionAccumulation(s []byte) error {
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

	global.G_LOG.Infof("consumer:%s", s)
	var m channelData.UserPromotionMessage
	err := tool.JsonUnmarshal(s, &m)
	if err != nil {
		global.G_LOG.Infof("消费提款 累计总优惠金额 解码json err:%s", err.Error())
		return err
	}

	//存储用户流水
	//err = global.G_REDIS.IncrByFloat(context.Background(), fmt.Sprintf(enmus.UserTotalRechargeAmountKey, m.UserId), m.DepositAmount).Err()
	//if err != nil {
	//	global.G_LOG.Infof("消费存款 累计总存款 err:%s", err.Error())
	//	return err
	//}

	key := fmt.Sprintf("PromotionAccumulation::%s", m.OrderSn)
	//加分布式锁
	if global.G_REDIS.SetNX(context.Background(), key, "true", time.Duration(30)*time.Minute).Val() {
		report := modules.FindByKeyFcUserReportFirst(&dos.FcUserReport{UserId: m.UserId})

		if len(report.Id) == 0 {
			row := dos.FcUserReport{}
			row.UserId = m.UserId
			row.UserName = m.UserName
			if m.PromotionType == 2 {
				row.RebateAmount = m.PromotionAmount
			} else {
				row.PromotionAmount = m.PromotionAmount
			}

			modules.SaveFcUserReport(&row)
		} else {
			if m.PromotionType == 2 {
				global.G_DB.Model(&dos.FcUserReport{}).Where("user_id=?", m.UserId).Updates(map[string]interface{}{
					"rebate_amount": gorm.Expr("rebate_amount + ?", m.PromotionAmount),
				})
			} else {
				global.G_DB.Model(&dos.FcUserReport{}).Where("user_id=?", m.UserId).Updates(map[string]interface{}{
					"promotion_amount": gorm.Expr("promotion_amount + ?", m.PromotionAmount),
				})
			}
		}
	}

	return nil

}
