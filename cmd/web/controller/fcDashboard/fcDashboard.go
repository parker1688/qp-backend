package fcDashboard

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcDashboard/pics
func DashboardPics(c *gin.Context) {
	var jsonp vo.DashboardPicReq
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

	//
	query1 := global.G_DB.Model(&dos.FcUserMaterial{})
	query2 := global.G_DB.Model(&dos.FcUserMaterial{})

	query3 := global.G_DB.Model(&dos.FcLoginLog{})
	query4 := global.G_DB.Model(&dos.FcLoginLog{})

	query5 := global.G_DB.Model(&dos.FcOrderDeposit{})
	query6 := global.G_DB.Model(&dos.FcOrderDeposit{})

	query7 := global.G_DB.Model(&dos.FcOrderWithdraw{})
	query8 := global.G_DB.Model(&dos.FcOrderWithdraw{})

	query9 := global.G_DB.Model(&dos.FcFirstOrderDeposit{})
	query10 := global.G_DB.Model(&dos.FcFirstOrderDeposit{})

	query11 := global.G_DB.Model(&dos.FcBetRecord{})
	query12 := global.G_DB.Model(&dos.FcBetRecord{})

	//global.G_LOG.Infof("Dashboard ---------------------1:%v, %v, %v", jsonp.MerchantCode, jsonp.StartTime, jsonp.EndTime)
	//global.G_LOG.Infof("Dashboard ---------------------2:%v, %v, %v", jsonp.MerchantCode, jsonp.PreStartTime, jsonp.PreEndTime)

	//
	if !jsonp.StartTime.Timer().IsZero() && !jsonp.EndTime.Timer().IsZero() {
		query1 = query1.Where("create_time >= ? and create_time <= ?", jsonp.StartTime, jsonp.EndTime)
		query3 = query3.Where("create_time >= ? and create_time <= ?", jsonp.StartTime, jsonp.EndTime)
		query5 = query5.Where("create_time >= ? and create_time <=  ?", jsonp.StartTime, jsonp.EndTime)
		query7 = query7.Where("create_time >= ? and create_time <= ?", jsonp.StartTime, jsonp.EndTime)
		query9 = query9.Where("create_time >= ? and create_time <= ?", jsonp.StartTime, jsonp.EndTime)
		query11 = query11.Where("create_time >= ? and create_time <= ?", jsonp.StartTime, jsonp.EndTime)

	}
	if !jsonp.PreStartTime.Timer().IsZero() && !jsonp.PreEndTime.Timer().IsZero() {
		query2 = query2.Where("create_time >= ? and create_time <= ?", jsonp.PreStartTime, jsonp.PreEndTime)
		query4 = query4.Where("create_time >= ? and create_time <= ?", jsonp.PreStartTime, jsonp.PreEndTime)
		query6 = query6.Where("create_time >= ? and create_time <= ?", jsonp.PreStartTime, jsonp.PreEndTime)
		query8 = query8.Where("create_time >= ? and create_time <= ?", jsonp.PreStartTime, jsonp.PreEndTime)
		query10 = query10.Where("create_time >= ? and create_time <= ?", jsonp.PreStartTime, jsonp.PreEndTime)
		query12 = query12.Where("create_time >= ? and create_time <= ?", jsonp.PreStartTime, jsonp.PreEndTime)
	}

	if jsonp.MerchantCode != "" {
		query1 = query1.Where("merchant_code = ?", jsonp.MerchantCode)
		query2 = query2.Where("merchant_code = ?", jsonp.MerchantCode)
		query3 = query3.Where("merchant_code = ?", jsonp.MerchantCode)
		query4 = query4.Where("merchant_code = ?", jsonp.MerchantCode)
		query5 = query5.Where("merchant_code = ?", jsonp.MerchantCode)
		query6 = query6.Where("merchant_code = ?", jsonp.MerchantCode)
		query7 = query7.Where("merchant_code = ?", jsonp.MerchantCode)
		query8 = query8.Where("merchant_code = ?", jsonp.MerchantCode)
		query9 = query9.Where("merchant_code = ?", jsonp.MerchantCode)
		query10 = query10.Where("merchant_code = ?", jsonp.MerchantCode)
		query11 = query11.Where("merchant_code = ?", jsonp.MerchantCode)
		query12 = query12.Where("merchant_code = ?", jsonp.MerchantCode)
	}

	addUserCount := 0
	preAddUserCount := 0
	loginUserCount := 0
	preLoginUserCount := 0
	rechargeAmount := 0.0
	preRechargeAmount := 0.0
	withDrawAmount := 0.0
	preWithDrawAmount := 0.0
	firstRechargeAmount := 0.0
	preFirstRechargeAmount := 0.0
	netAmount := 0.0
	preNetAmount := 0.0

	query1.Select("count(user_id) as addUserCount").Scan(&addUserCount)
	query2.Select("count(user_id) as preAddUserCount").Scan(&preAddUserCount)

	query3.Select("count(DISTINCT user_id) as loginUserCount").Scan(&loginUserCount)
	query4.Select("count(DISTINCT user_id) as preLoginUserCount").Scan(&preLoginUserCount)

	query5.Select("sum(amount) as rechargeAmount").Where("status=3").Scan(&rechargeAmount)
	query6.Select("sum(amount) as rechargeAmount").Where("status=3").Scan(&preRechargeAmount)

	query7.Select("sum(pre_amount) as withDrawaAmount").Where("status=3").Scan(&withDrawAmount)
	query8.Select("sum(pre_amount) as withDrawaAmount").Where("status=3").Scan(&preWithDrawAmount)

	query9.Select("sum(amount) as firstRechargeAmount").Scan(&firstRechargeAmount)
	query10.Select("sum(amount) as firstRechargeAmount").Scan(&preFirstRechargeAmount)

	query11.Select("sum(net_amount) as netAmount").Scan(&netAmount)
	query12.Select("sum(net_amount) as netAmount").Scan(&preNetAmount)

	//global.G_LOG.Infof("sql1:%v, %v", addUserCount, preAddUserCount)
	//global.G_LOG.Infof("sql2:%v, %v", loginUserCount, preLoginUserCount)
	//global.G_LOG.Infof("sql3:%v, %v", rechargeAmount, preRechargeAmount)
	//global.G_LOG.Infof("sql4:%v, %v", withDrawAmount, preWithDrawAmount)
	//global.G_LOG.Infof("sql5:%v, %v", firstRechargeAmount, preFirstRechargeAmount)
	//global.G_LOG.Infof("sql6:%v, %v", netAmount, preNetAmount)

	data := vo.DashboadPicResp{AddUserNum: addUserCount, PreAddUserNum: preAddUserCount, LoginUserCount: loginUserCount, PreLoginUserCount: preLoginUserCount,
		RechargeAmount: rechargeAmount, PreRechargeAmount: preRechargeAmount, WithdrawAmount: withDrawAmount, PreWithdrawAmount: preWithDrawAmount,
		FirstRechargeAmount: firstRechargeAmount, PreFirstRechargeAmount: preFirstRechargeAmount, NetAmount: netAmount,
		PreNetAmount: preNetAmount}

	response.SuccessJSON(c, data)
}

// api: api/fcDashboard/hour
func DashboardHour(c *gin.Context) {
	var jsonp vo.DashboardHourReq
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

	startTime := tool.HourTimeStamp()
	endTime := time.Now().Unix()

	//global.G_LOG.Infof("register----------merchant_code:%v", jsonp.MerchantCode)
	registerList := []int{}
	for i := 0; i < 24; i++ {
		st := tool.TimeStampToDate(tool.TimeLayout, startTime)
		et := tool.TimeStampToDate(tool.TimeLayout, endTime)
		query := global.G_DB.Model(&dos.FcUserMaterial{})
		if jsonp.MerchantCode != "" {
			query = query.Where("merchant_code = ?", jsonp.MerchantCode)
		}
		query = query.Where("create_time >= ? and create_time < ?", st, et)
		num1 := 0
		query.Select("count(user_id) as num1").Scan(&num1)
		//global.G_LOG.Infof("register----------info:%v, %v, %v", st, et, num1)

		registerList = append(registerList, num1)
		endTime = startTime
		startTime -= int64(3600)
	}

	//global.G_LOG.Infof("register---------list:%v", registerList)
	data := vo.DashboadHourResp{RegisterNum: registerList}
	response.SuccessJSON(c, data)
}

// api: api/fcDashboard/User
func DashboardUser(c *gin.Context) {
	var jsonp vo.DashboardUserReq
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

	query := global.G_DB.Model(&dos.FcUserMaterial{})
	query1 := global.G_DB.Model(&dos.FcLoginLog{})
	query2 := global.G_DB.Model(&dos.FcLoginLog{})
	query3 := global.G_DB.Model(&dos.FcLoginLog{})
	query4 := global.G_DB.Model(&dos.FcUserMaterial{})
	if jsonp.MerchantCode != "" {
		query = query.Where("merchant_code = ?", jsonp.MerchantCode)
		query1 = query1.Where("merchant_code = ?", jsonp.MerchantCode)
		query2 = query2.Where("merchant_code = ?", jsonp.MerchantCode)
		query3 = query3.Where("merchant_code = ?", jsonp.MerchantCode)
		query4 = query4.Where("merchant_code = ?", jsonp.MerchantCode)
	}

	now := time.Now()
	zeroTime, _ := tool.DayStartEndDate(now)

	now2 := time.Now().Unix()
	serverDayTime := now2 - int64(7*86400)
	serverDayTime2 := tool.TimeStampToDate(tool.TimeLayout, serverDayTime)
	monDayTime := now2 - int64(30*86400)
	monDayTime2 := tool.TimeStampToDate(tool.TimeLayout, monDayTime)
	num1 := 0 //当天
	query.Select("count(user_id) as num1").Where("create_time > ?", zeroTime).Scan(&num1)
	num2 := 0 //7日内
	query1.Select("count(DISTINCT user_id) as num2").Where("create_time > ?", serverDayTime2).Scan(&num2)
	num3 := 0 //7日外
	query2.Select("count(DISTINCT user_id) as num3").Where("create_time < ?", serverDayTime2).Scan(&num3)
	num4 := 0
	query3.Select("count(DISTINCT user_id) as num4").Where("create_time < ?", monDayTime2).Scan(&num4)
	totalnum := 0
	query4.Select("count(user_id) as totalnum").Scan(&totalnum)
	data := vo.DashboadUserResp{Num1: num1, Num2: num2, Num3: num3, Num4: num4, TotalNum: totalnum}
	response.SuccessJSON(c, data)
}
