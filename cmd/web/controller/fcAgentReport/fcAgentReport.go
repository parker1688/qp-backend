package fcAgentReport

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func FindAgentReport(c *gin.Context) {
	jsonp := struct {
		dos.FcAgentReport
	}{}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}

	userInfo, ok := c.Get("UserInfo")
	merchantCodes := ""
	if ok {
		merchantCodes = userInfo.(*dos.AdminUser).MerchantCodes
	}
	//global.G_LOG.Infof("fc user report -------------------------------1:%v", merchantCodes)

	groupRows := []dos.FcAgentGroup{}
	query := global.G_DB.Model(&dos.FcAgentGroup{})
	dayNum := jsonp.Days
	if jsonp.StartAt != "" {
		query = query.Where("create_time >= ?", jsonp.StartAt)
	}
	if jsonp.EndAt != "" {
		query = query.Where("create_time <= ?", jsonp.EndAt)
	}
	merchantCodes2 := tool.StrToListForSql(merchantCodes)
	query.Where("merchant_code in (?)", merchantCodes2)

	query.Find(&groupRows)
	//global.G_LOG.Info("fc agent report -------------------------------2:%v", groupRows)

	rsdata := []dos.FcAgentReport{}
	for _, groupInfo := range groupRows {
		//global.G_LOG.Info("fc agent report -------------------------------2:%v", groupInfo)
		//queryGroup := global.G_DB.Model(&dos.FcAgentGroup{})
		//queryGroup = queryGroup.Where("id = ?", groupid)
		groupName := groupInfo.GroupName
		code := groupInfo.InviteCode
		time1 := groupInfo.CreateTime.String()
		//queryGroup.Select("group_name").Scan(&groupName)
		//queryGroup.Select("code").Scan(&code)
		//queryGroup.Select("create_time").Scan(&time1)
		//global.G_LOG.Info("fc agent report -------------------------------3:%v, %v, %v", groupName, code, time1)

		timeStamp := tool.DateToTimeStamp(tool.TimeLayout, time1) + int64(dayNum*86400)
		time2 := tool.TimeStampToDate(tool.TimeLayout, timeStamp) // 开启时间+N天
		//global.G_LOG.Info("fc agent report -------------------------------3-2:%v, %v, %v", code, time1, time2)

		registCount := saveRigster(code, time1, time2)
		//global.G_LOG.Info("fc agent report -------------------------------4:%v", registCount)

		//queryUser := global.G_DB.Model(&dos.FcAgent{})
		//queryUser = queryUser.Where("invite_code = ?", code)
		//merchantCode := ""
		//queryUser.Select("merchant_code").Scan(&merchantCode)
		//global.G_LOG.Info("fc agent report -------------------------------5:%v", merchantCode)

		firstPayPlayerNum := saveFirstPayPlayerNum(code, time1, time2)
		//global.G_LOG.Info("fc agent report -------------------------------6:%v", firstPayPlayerNum)

		payPlayerNum := savePayPlayerNum(code, time1, time2)
		//global.G_LOG.Info("fc agent report -------------------------------7:%v", payPlayerNum)

		firstPayMoney := saveFirstPayMoney(code, time1, time2)
		//global.G_LOG.Info("fc agent report -------------------------------8:%v", firstPayMoney)

		payMoney := savePayMoney(code, time1, time2)
		//global.G_LOG.Info("fc agent report -------------------------------9:%v", payMoney)

		payCount := savePayCount(code, time1, time2)
		//global.G_LOG.Info("fc agent report -------------------------------10:%v", payCount)

		rc := dos.FcAgentReport{InviteCode: code, GroupName: groupName, RegistCount: registCount, FirstPayPlayerNum: firstPayPlayerNum, //MerchantCode:merchantCode,
			PayPlayerNum: payPlayerNum, FirstPayMoney: firstPayMoney, PayMoney: payMoney, PayCount: payCount}
		rsdata = append(rsdata, rc)
	}
	response.SuccessJSON(c, rsdata)
}

// 统计注册人数
func saveRigster(code int, time, time2 string) int {
	query := global.G_DB.Model(&dos.FcUserMaterial{})
	query = query.Where("agent_invite_code = ?", code)
	query = query.Where("create_time >= ?", time)
	query = query.Where("create_time <= ?", time2)
	registCount := 0
	query.Select("count(1) as registCount").Scan(&registCount)
	return registCount
}

// 统计首冲人数
func saveFirstPayPlayerNum(code int, time1, time2 string) int {
	query := global.G_DB.Model(&dos.FcFirstOrderDeposit{})
	query = query.Where("invite_code = ?", code)
	query = query.Where("create_time >= ?", time1)
	query = query.Where("create_time <= ?", time2)

	count := 0
	query.Select("count(user_id)").Scan(&count)
	return count
}

// 统计充值人数
func savePayPlayerNum(code int, time1, time2 string) int {
	query := global.G_DB.Model(&dos.FcOrderDeposit{})
	query = query.Where("invite_code = ? and status = 3 ", code)
	query = query.Where("create_time >= ?", time1)
	query = query.Where("create_time <= ?", time2)

	count := 0
	userList := []string{}
	query.Select("user_id").Group("user_id").Find(&userList)
	count = len(userList)
	return count
}

// 统计首冲额度
func saveFirstPayMoney(code int, time1, time2 string) float64 {
	query := global.G_DB.Model(&dos.FcFirstOrderDeposit{})
	query = query.Where("invite_code = ?", code)
	query = query.Where("create_time >= ?", time1)
	query = query.Where("create_time <= ?", time2)

	count := 0.0
	query.Select("sum(amount)").Scan(&count)
	return count
}

// 统计充值额度
func savePayMoney(code int, time1, time2 string) float64 {
	query := global.G_DB.Model(&dos.FcOrderDeposit{})
	query = query.Where("invite_code = ? and status = 3 ", code)
	query = query.Where("create_time >= ?", time1)
	query = query.Where("create_time <= ?", time2)

	count := 0.0
	query.Select("sum(amount)").Scan(&count)
	return count
}

// 统计充值额度
func savePayCount(code int, time1, time2 string) int {
	query := global.G_DB.Model(&dos.FcOrderDeposit{})
	query = query.Where("invite_code = ? and status = 3 ", code)
	query = query.Where("create_time >= ?", time1)
	query = query.Where("create_time <= ?", time2)

	count := 0
	query.Select("count(amount)").Scan(&count)
	return count
}

//
//func saveVistPage(uid int, code int) {
//
//}
//
//func saveDownApp(uid int, code int) {
//
//}

func FindDetailAgentReport(c *gin.Context) {
	jsonp := struct {
		dos.FcAgentReportDetailRep
	}{}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}

	userInfo, ok := c.Get("UserInfo")
	merchantCodes := ""
	if ok {
		merchantCodes = userInfo.(*dos.AdminUser).MerchantCodes
	}
	//global.G_LOG.Infof("fc user report -------------------------------1:%v, %v", merchantCodes, userInfo)

	// 有效投注额
	sumValidBetAmount := 0.00
	// 用户输赢
	sumWinAmount := 0.00
	// 福利累计
	sumPromotionAmount := 0.00
	// 返水累计
	sumRebateAmount := 0.00
	// 充值累计
	sumRechargeAmount := 0.00
	// 提现累计
	sumWithdrawalAmount := 0.00

	rsdata := []dos.FcAgentReportDetail{}

	userList := []string{}
	query := global.G_DB.Model(&dos.FcUserMaterial{}) //充值
	merchantCodes2 := tool.StrToListForSql(merchantCodes)
	query.Where("merchant_code in (?)", merchantCodes2)
	query.Select("user_id as userList").Where("agent_invite_code=?", jsonp.InviteCode).Scan(&userList)
	//global.G_LOG.Infof("fc user report -------------------------------2:%v", userList)

	for _, uid := range userList {
		q1 := global.G_DB.Model(&dos.FcOrderDeposit{})            //充值
		q2 := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}) //提现
		q3 := global.G_DB.Model(&dos.FcBetRecord{})               //有效下注，利润
		q5 := global.G_DB.Model(&dos.FcOrderPromotion{})          //福利总表
		q6 := global.G_DB.Model(&dos.FcUserRebateRecords{})       //反水

		if len(jsonp.Id) > 0 {
			q1 = q1.Where("id = ?", jsonp.Id)
			q2 = q2.Where("id = ?", jsonp.Id)
			q3 = q3.Where("id = ?", jsonp.Id)
			q5 = q5.Where("id = ?", jsonp.Id)
			q6 = q6.Where("id = ?", jsonp.Id)
		}

		if jsonp.StartAt != "" {
			q1 = q1.Where("create_time >= ?", jsonp.StartAt)
			q2 = q2.Where("create_time >= ?", jsonp.StartAt)
			q3 = q3.Where("create_time >= ?", jsonp.StartAt)
			q5 = q5.Where("create_time >= ?", jsonp.StartAt)
			q6 = q6.Where("create_time >= ?", jsonp.StartAt)
		}
		if jsonp.EndAt != "" {
			q1 = q1.Where("create_time <= ?", jsonp.EndAt)
			q2 = q2.Where("create_time <= ?", jsonp.EndAt)
			q3 = q3.Where("create_time <= ?", jsonp.EndAt)
			q5 = q5.Where("create_time <= ?", jsonp.EndAt)
			q6 = q6.Where("create_time <= ?", jsonp.EndAt)
		}
		rechargeAmount := 0.0
		//global.G_LOG.Info("fc user report -------------------------------2-0:%v, %v, %v", uid, jsonp.StartAt, jsonp.EndAt)
		q1.Select("sum(amount) as rechargeAmount").Where("user_id = ? and status=3", uid).Scan(&rechargeAmount)
		//global.G_LOG.Info("fc user report -------------------------------2-1:%v, %v, %v", rechargeAmount, q1)
		withDrawaAmount := 0.0
		q2.Select("sum(amount) as withDrawaAmount").Where("user_id = ? and status=3", uid).Scan(&withDrawaAmount)
		//global.G_LOG.Info("fc user report -------------------------------2-2:%v, %v, %v", withDrawaAmount, withDrawaCount, q2)
		validBetAmount := 0.0
		q3.Select("sum(valid_betamount) as validBetAmount").Where("user_id = ?", uid).Scan(&validBetAmount)
		netAmount := 0.0
		q3.Select("sum(net_amount) as netAmount").Where("user_id = ?", uid).Scan(&netAmount)
		//global.G_LOG.Info("fc user report -------------------------------2-4:%v, %v", q3, netAmount)
		promotionAmount := 0.0
		q5.Select("sum(amount) as promotionAmount").Where("user_id = ?", uid).Scan(&promotionAmount)
		rebateAmount := 0.0
		q6.Select("sum(bonus_amount) as rebateAmount").Where("user_id = ?", uid).Scan(&rebateAmount)
		subRechargeWithdraw := rechargeAmount - withDrawaAmount
		rc := dos.FcAgentReportDetail{UserId: uid, InviteCode: jsonp.InviteCode,
			RechargeAmount: rechargeAmount, WithdrawalAmount: withDrawaAmount, ValidBetamount: validBetAmount,
			WinAmount: netAmount, RebateAmount: rebateAmount, PromotionAmount: promotionAmount, SubRechargeWithDraw: subRechargeWithdraw}

		//global.G_LOG.Info("fc user report -------------------------------2:%v", rc)
		rsdata = append(rsdata, rc)

		sumValidBetAmount += validBetAmount
		sumWinAmount += netAmount
		sumPromotionAmount += promotionAmount
		sumRebateAmount += rebateAmount
		sumRechargeAmount += rechargeAmount
		sumWithdrawalAmount += withDrawaAmount
	}

	respMap := map[string]interface{}{}
	respMap["sumValidBetAmount"] = sumValidBetAmount
	respMap["sumWinAmount"] = sumWinAmount
	respMap["sumPromotionAmount"] = sumPromotionAmount
	respMap["sumRebateAmount"] = sumRebateAmount
	respMap["sumRechargeAmount"] = sumRechargeAmount
	respMap["sumWithdrawalAmount"] = sumWithdrawalAmount
	respMap["list"] = rsdata

	response.SuccessJSON(c, respMap)
}
