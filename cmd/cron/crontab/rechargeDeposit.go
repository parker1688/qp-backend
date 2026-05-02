package crontab

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"time"
)

func init() {
	cronFunc = append(cronFunc, cronTabEvery{
		spec: "@every 1m", //5分钟一次
		cmd:  RechargeDeposit,
	})
}

// RechargeDeposit
//
//	@Description: 充值订单，定期将未支付的订单设置为失败
func RechargeDeposit() {
	nowTime := time.Now()
	nowTimeStr := nowTime.Format("2006-01-02 15:04:05")
	//global.G_LOG.Infof("RechargeDeposit cron startAt: %v", nowTimeStr)
	t := nowTime.Add(-1 * time.Hour).Format(tool.TimeLayout)
	updateMap := map[string]interface{}{}
	updateMap["status"] = enmus.ORDER_NO_STATUS
	updateMap["auth_by"] = "system"
	updateMap["auth_time"] = nowTimeStr

	for {

		eRows := global.G_DB.Model(&dos.FcOrderDeposit{}).Where("status = ? and create_time <= ? limit 20", enmus.Order_STATUS_PENDING_PAY, t).Updates(updateMap)
		if eRows.Error != nil {
			global.G_LOG.Errorf("update FcOrderDeposit status fail, %s", eRows.Error.Error())
			break
		}

		// 如果没有则直接跳出
		effected := eRows.RowsAffected
		if effected < 1 {
			break
		}
		//global.G_LOG.Infof("RechargeDeposit cron effected: %v", effected)

		time.Sleep(100 * time.Millisecond)
	}
}
