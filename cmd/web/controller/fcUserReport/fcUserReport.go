// The build tag makes sure the stub is not built in the final build.

package fcUserReport

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"fmt"
	"slices"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

var userReportCache = cache.New(5*time.Minute, 5*time.Minute)

// api: api/fcUserReport/save
func SaveFcUserReportControl(c *gin.Context) {
	var jsonp dos.FcUserReport
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
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, err := modules.SaveFcUserReport(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

//// api: api/fcUserReport/findPage
//func FindPageFcUserReportControl(c *gin.Context) {
//	jsonp := struct {
//		response.PageTimeQuery
//		dos.FcUserReport
//	}{}
//	jsonp.PageTimeQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
//	jsonp.PageTimeQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
//	jsonp.StartAt = c.DefaultQuery("startAt", "")
//	jsonp.EndAt = c.DefaultQuery("endAt", "")
//	jsonp.Id = c.DefaultQuery("id", "")
//	jsonp.UserId = c.DefaultQuery("user_id", "")
//	jsonp.UserName = c.DefaultQuery("user_name", "")
//
//	jsonp.RechargeCount = tool.Atoi(c.DefaultQuery("recharge_count", ""))
//
//	jsonp.BetCount = tool.Atoi(c.DefaultQuery("bet_count", ""))
//
//	jsonp.WithdrawalCount = tool.Atoi(c.DefaultQuery("withdrawal_count", ""))
//
//	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
//	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
//	jsonp.CreateBy = c.DefaultQuery("create_by", "")
//	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
//	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
//
//	err := validator.New().Struct(jsonp)
//	if err != nil {
//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
//		return
//	}
//	data, total := modules.FindPageFcUserReport(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserReport, &jsonp.PageTimeQuery)
//
//	query := global.G_DB.Model(&dos.FcUserReport{})
//	if len(jsonp.Id) > 0 {
//		query = query.Where("id = ?", jsonp.Id)
//	}
//
//	if len(jsonp.UserId) > 0 {
//		query = query.Where("user_id = ?", jsonp.UserId)
//	}
//
//	if len(jsonp.UserName) > 0 {
//		query = query.Where("user_name = ?", jsonp.UserName)
//	}
//
//	if len(jsonp.MerchantCode) > 0 {
//		query = query.Where("merchant_code = ?", jsonp.MerchantCode)
//	}
//
//	if jsonp.StartAt != "" {
//		query = query.Where("create_time >= ?", jsonp.CreateTime)
//	}
//	if jsonp.EndAt != "" {
//		query = query.Where("create_time <= ?", jsonp.CreateTime)
//	}
//
//	//sumBetAmount := 0.00
//	//query.Select("sum(bet_amount) as sumBetAmount").Scan(&sumBetAmount)
//
//	// 有效投注额
//	sumValidBetAmount := 0.00
//	query.Select("sum(valid_betamount) as sumValidBetAmount").Scan(&sumValidBetAmount)
//
//	// 用户输赢
//	sumWinAmount := 0.00
//	query.Select("sum(win_amount) as sumWinAmount").Scan(&sumWinAmount)
//
//	// 福利累计
//	sumPromotionAmount := 0.00
//	query.Select("sum(promotion_amount) as sumPromotionAmount").Scan(&sumPromotionAmount)
//
//	// 返水累计
//	sumRebateAmount := 0.00
//	query.Select("sum(rebate_amount) as sumRebateAmount").Scan(&sumRebateAmount)
//
//	// 充值累计
//	sumRechargeAmount := 0.00
//	//query.Select("sum(recharge_amount) as sumRechargeAmount").Scan(&sumRechargeAmount)
//	global.G_DB.Model(&dos.FcOrderDeposit{}).Select("sum(amount)").Where("user_id = ? and status=3", jsonp.UserId).Scan(&sumRechargeAmount)
//
//	// 提现累计
//	sumWithdrawalAmount := 0.00
//	//query.Select("sum(withdrawal_amount) as sumWithdrawalAmount").Scan(&sumWithdrawalAmount)
//	global.G_DB.Model(&dos.FcOrderWithdraw{}).Select("sum(pre_amount)").Where("user_id = ? and status=3", jsonp.UserId).Scan(&sumWithdrawalAmount)
//
//	newData := make([]*dos.FcUserReportListResp, len(data))
//	tool.JsonMapper(data, &newData)
//	for i, v := range newData {
//		subRechargeWithDraw := tool.DecimalFSubTruncate(v.RechargeAmount, v.WithdrawalAmount, 2)
//		fdOdds := 0.00
//		if v.RechargeAmount > 0.00 && subRechargeWithDraw > 0.00 {
//			fdOdds = tool.DecimalFDivTruncate(subRechargeWithDraw, v.RechargeAmount, 4)
//		}
//		betMultiple := 0.00
//		if v.ValidBetamount > 0.00 && v.RechargeAmount > 0.00 {
//			betMultiple = tool.DecimalFDivTruncate(v.ValidBetamount, v.RechargeAmount, 4)
//		}
//
//		newData[i].SubRechargeWithDraw = subRechargeWithDraw
//		newData[i].FdOdds = fdOdds
//		newData[i].BetMultiple = betMultiple
//	}
//
//	respMap := map[string]interface{}{}
//	respMap["sumValidBetAmount"] = sumValidBetAmount
//	respMap["sumWinAmount"] = sumWinAmount
//	respMap["sumPromotionAmount"] = sumPromotionAmount
//	respMap["sumRebateAmount"] = sumRebateAmount
//	respMap["sumRechargeAmount"] = sumRechargeAmount
//	respMap["sumWithdrawalAmount"] = sumWithdrawalAmount
//	respMap["list"] = newData
//
//	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, respMap)
//}

func FindPageFcUserReportControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcUserReport
		GroupName string `json:"group_name"`
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.GroupName = c.DefaultQuery("group_name", "")

	jsonp.RechargeCount = tool.Atoi(c.DefaultQuery("recharge_count", ""))

	jsonp.BetCount = tool.Atoi(c.DefaultQuery("bet_count", ""))

	jsonp.WithdrawalCount = tool.Atoi(c.DefaultQuery("withdrawal_count", ""))

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

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
	//global.G_LOG.Info("fc user report -------------------------------0-0:%v, %v, %v, %v, %v", jsonp.Id, jsonp.UserName, jsonp.MerchantCode, jsonp.StartAt, jsonp.EndAt)

	rsdata := []*dos.FcUserReport{}
	total := int64(0)

	time1 := time.Now().Unix()
	cacheKey := fmt.Sprintf("userReport:%s:%s:%s:%s", jsonp.MerchantCode, jsonp.StartAt, jsonp.EndAt, jsonp.UserId)
	reportData, ok := userReportCache.Get(cacheKey)

	if ok {
		rsdata = reportData.([]*dos.FcUserReport)
		for _, v := range rsdata {
			sumValidBetAmount += v.ValidBetamount
			sumWinAmount += v.WinAmount
			sumPromotionAmount += v.PromotionAmount
			sumRebateAmount += v.RebateAmount
			sumRechargeAmount += v.RechargeAmount
			sumWithdrawalAmount += v.WithdrawalAmount
		}
	} else {
		if len(jsonp.UserId) > 0 { //有用户id只选一个人的

			query1 := global.G_DB.Model(&dos.FcOrderDeposit{})            //充值
			query2 := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}) //提现
			query3 := global.G_DB.Model(&dos.FcBetRecord{})               //有效下注，利润
			query5 := global.G_DB.Model(&dos.FcOrderPromotion{})          //福利总表
			query6 := global.G_DB.Model(&dos.FcUserRebateRecords{})       //反水

			if len(jsonp.MerchantCode) > 0 {
				query1 = query1.Where("merchant_code = ?", jsonp.MerchantCode)
				query2 = query2.Where("merchant_code = ?", jsonp.MerchantCode)
				query3 = query3.Where("merchant_code = ?", jsonp.MerchantCode)
				query5 = query5.Where("merchant_code = ?", jsonp.MerchantCode)
				query6 = query6.Where("merchant_code = ?", jsonp.MerchantCode)
			}

			if jsonp.StartAt != "" {
				query1 = query1.Where("create_time >= ?", jsonp.StartAt)
				query2 = query2.Where("create_time >= ?", jsonp.StartAt)
				query3 = query3.Where("create_time >= ?", jsonp.StartAt)
				query5 = query5.Where("create_time >= ?", jsonp.StartAt)
				query6 = query6.Where("create_time >= ?", jsonp.StartAt)
			}
			if jsonp.EndAt != "" {
				query1 = query1.Where("create_time <= ?", jsonp.EndAt)
				query2 = query2.Where("create_time <= ?", jsonp.EndAt)
				query3 = query3.Where("create_time <= ?", jsonp.EndAt)
				query5 = query5.Where("create_time <= ?", jsonp.EndAt)
				query6 = query6.Where("create_time <= ?", jsonp.EndAt)
			}

			uid := jsonp.UserId

			if len(jsonp.GroupName) > 0 {
				userIds := modules.GetUserGroupIdsByGroupName(jsonp.GroupName)
				if !slices.Contains(userIds, uid) {
					uid = "0" // 没有在分组中不能查询
				}
			}

			rechargeAmount := 0.0
			query1.Select("sum(amount) as rechargeAmount").Where("user_id = ? and status=3", uid).Scan(&rechargeAmount)
			rechargeCount := 0
			query1.Select("count(1) as rechargeCount").Scan(&rechargeCount)
			withDrawaAmount := 0.0
			query2.Select("sum(amount) as withDrawaAmount").Where("user_id = ? and status=3", uid).Scan(&withDrawaAmount)
			withDrawaCount := 0
			query2.Select("count(1) as withDrawaCount").Scan(&withDrawaCount)
			validBetAmount := 0.0
			query3.Select("sum(valid_betamount) as validBetAmount").Where("user_id = ?", uid).Scan(&validBetAmount)
			betAmount := 0.0
			query3.Select("sum(bet_amount) as betAmount").Scan(&betAmount)
			betCount := 0
			query3.Select("count(1) as betCount").Scan(&betCount)

			netAmount := 0.0
			query3.Select("sum(net_amount) as netAmount").Scan(&netAmount)
			promotionAmount := 0.0
			query5.Select("sum(amount) as promotionAmount").Where("user_id = ?", uid).Scan(&promotionAmount)
			promotionAmount += modules.GetUserReportTotalPromotionExVal(uid)
			rebateAmount := 0.0
			query6.Select("sum(bonus_amount) as rebateAmount").Where("user_id = ?", uid).Scan(&rebateAmount)

			userName := jsonp.UserName

			userMerchantCode := ""
			global.G_DB.Model(&dos.FcUserMaterial{}).Select("merchant_code as userMerchantCode").Where("user_id=?", uid).Scan(&userMerchantCode)

			rc := dos.FcUserReport{UserId: jsonp.UserId, UserName: userName, MerchantCode: userMerchantCode,
				RechargeAmount: rechargeAmount, RechargeCount: rechargeCount, WithdrawalAmount: withDrawaAmount,
				WithdrawalCount: withDrawaCount, BetAmount: betAmount, ValidBetamount: validBetAmount,
				BetCount: betCount, WinAmount: netAmount, RebateAmount: rebateAmount, PromotionAmount: promotionAmount}
			rsdata = append(rsdata, &rc)

			sumValidBetAmount += validBetAmount
			sumWinAmount += netAmount
			sumPromotionAmount += promotionAmount
			sumRebateAmount += rebateAmount
			sumRechargeAmount += rechargeAmount
			sumWithdrawalAmount += withDrawaAmount
		} else {
			userList := []string{}
			query := global.G_DB.Model(&dos.FcUserMaterial{}) //充值
			if len(jsonp.GroupName) > 0 {
				userList = modules.GetUserGroupIdsByGroupName(jsonp.GroupName)
				global.G_LOG.Infof("userReport userList-1:%v", len(userList))
			} else {
				query.Select("user_id as userList").Where("merchant_code=?", jsonp.MerchantCode).Scan(&userList)
				global.G_LOG.Infof("userReport userList-2:%v", len(userList))
			}

			for _, uid := range userList {
				q1 := global.G_DB.Model(&dos.FcOrderDeposit{})            //充值
				q2 := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}) //提现
				q3 := global.G_DB.Model(&dos.FcBetRecord{})               //有效下注，利润
				q5 := global.G_DB.Model(&dos.FcOrderPromotion{})          //福利总表
				q6 := global.G_DB.Model(&dos.FcUserRebateRecords{})       //反水

				if len(jsonp.MerchantCode) > 0 {
					q1 = q1.Where("merchant_code = ?", jsonp.MerchantCode)
					q2 = q2.Where("merchant_code = ?", jsonp.MerchantCode)
					q3 = q3.Where("merchant_code = ?", jsonp.MerchantCode)
					q5 = q5.Where("merchant_code = ?", jsonp.MerchantCode)
					q6 = q6.Where("merchant_code = ?", jsonp.MerchantCode)
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
				q1.Select("sum(amount) as rechargeAmount").Where("user_id = ? and status=3", uid).Scan(&rechargeAmount)
				rechargeCount := 0
				q1.Select("count(1) as rechargeCount").Scan(&rechargeCount)
				withDrawaAmount := 0.0
				q2.Select("sum(amount) as withDrawaAmount").Where("user_id = ? and status=3", uid).Scan(&withDrawaAmount)
				withDrawaCount := 0
				q2.Select("count(1) as withDrawaCount").Scan(&withDrawaCount)
				validBetAmount := 0.0
				q3.Select("sum(valid_betamount) as validBetAmount").Where("user_id = ?", uid).Scan(&validBetAmount)
				betAmount := 0.0
				q3.Select("sum(bet_amount) as betAmount").Scan(&betAmount)
				betCount := 0
				q3.Select("count(1) as betCount").Scan(&betCount)
				netAmount := 0.0
				q3.Select("sum(net_amount) as netAmount").Scan(&netAmount)
				promotionAmount := 0.0
				q5.Select("sum(amount) as promotionAmount").Where("user_id = ?", uid).Scan(&promotionAmount)
				promotionAmount += modules.GetUserReportTotalPromotionExVal(uid)
				rebateAmount := 0.0
				q6.Select("sum(bonus_amount) as rebateAmount").Where("user_id = ?", uid).Scan(&rebateAmount)
				userName := ""
				query.Select("user_name as userName").Where("user_id=?", uid).Scan(&userName)

				userMerchantCode := ""
				global.G_DB.Model(&dos.FcUserMaterial{}).Select("merchant_code as userMerchantCode").Where("user_id=?", uid).Scan(&userMerchantCode)

				rc := dos.FcUserReport{UserId: uid, UserName: userName, MerchantCode: userMerchantCode,
					RechargeAmount: rechargeAmount, RechargeCount: rechargeCount, WithdrawalAmount: withDrawaAmount,
					WithdrawalCount: withDrawaCount, BetAmount: betAmount, ValidBetamount: validBetAmount,
					BetCount: betCount, WinAmount: netAmount, RebateAmount: rebateAmount, PromotionAmount: promotionAmount}

				if rechargeAmount > 0 || rechargeCount > 0 || withDrawaAmount > 0 || withDrawaCount > 0 || betAmount > 0 ||
					validBetAmount > 0 || betCount > 0 || netAmount > 0 || rebateAmount > 0 || promotionAmount > 0 {
					rsdata = append(rsdata, &rc)

					sumValidBetAmount += validBetAmount
					sumWinAmount += netAmount
					sumPromotionAmount += promotionAmount
					sumRebateAmount += rebateAmount
					sumRechargeAmount += rechargeAmount
					sumWithdrawalAmount += withDrawaAmount
				}
			}
		}
		userReportCache.Set(cacheKey, rsdata, 1*time.Minute)
	}

	newData := make([]*dos.FcUserReportListResp, len(rsdata))
	tool.JsonMapper(rsdata, &newData)
	for i, v := range newData {
		subRechargeWithDraw := tool.DecimalFSubTruncate(v.RechargeAmount, v.WithdrawalAmount, 2)
		fdOdds := 0.00
		if v.RechargeAmount > 0.00 && subRechargeWithDraw > 0.00 {
			fdOdds = tool.DecimalFDivTruncate(subRechargeWithDraw, v.RechargeAmount, 4)
		}
		betMultiple := 0.00
		if v.ValidBetamount > 0.00 && v.RechargeAmount > 0.00 {
			betMultiple = tool.DecimalFDivTruncate(v.ValidBetamount, v.RechargeAmount, 4)
		}

		newData[i].SubRechargeWithDraw = subRechargeWithDraw
		newData[i].FdOdds = fdOdds
		newData[i].BetMultiple = betMultiple
	}

	respMap := map[string]interface{}{}
	respMap["sumValidBetAmount"] = tool.TruncateFloat(sumValidBetAmount, 2)
	respMap["sumWinAmount"] = tool.TruncateFloat(sumWinAmount, 2)
	respMap["sumPromotionAmount"] = tool.TruncateFloat(sumPromotionAmount, 2)
	respMap["sumRebateAmount"] = tool.TruncateFloat(sumRebateAmount, 2)
	respMap["sumRechargeAmount"] = tool.TruncateFloat(sumRechargeAmount, 2)
	respMap["sumWithdrawalAmount"] = tool.TruncateFloat(sumWithdrawalAmount, 2)
	respMap["list"] = newData
	time2 := time.Now().Unix()
	global.G_LOG.Infof("userReport cost time :%v, %v, %v", time2, time1, time2-time1)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, respMap)
}

// api: api/fcUserReport/findByKey
func FindByKeyFcUserReportControl(c *gin.Context) {
	var jsonp dos.FcUserReport
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
	data := modules.FindByKeyFcUserReport(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserReport/update
func UpdateFcUserReportControl(c *gin.Context) {
	var jsonp dos.FcUserReport
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
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcUserReport(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserReport/delete
func DeleteFcUserReportControl(c *gin.Context) {
	var jsonp dos.FcUserReport
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
	data := modules.DeleteFcUserReport(&jsonp)
	response.SuccessJSON(c, data)
}
