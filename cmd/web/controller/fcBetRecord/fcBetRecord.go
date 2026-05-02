package fcBetRecord

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/venues"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcBetRecord/save
func SaveFcBetRecordControl(c *gin.Context) {
	var jsonp dos.FcBetRecord
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, err := modules.SaveFcBetRecord(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcBetRecord/findPage
func FindPageFcBetRecordControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcBetRecord
	}{}

	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "") // 投注时间
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.LastStartAt = c.DefaultQuery("last_startAt", "") // 结算时间
	jsonp.LastEndAt = c.DefaultQuery("last_endAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Account = c.DefaultQuery("account", "")
	jsonp.PlayerName = c.DefaultQuery("player_name", "")
	jsonp.OrderSn = c.DefaultQuery("order_sn", "")
	jsonp.GameCode = c.DefaultQuery("game_code", "")

	jsonp.BetTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("bet_time", "")))
	jsonp.SettlementTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("settlement_time", "")))
	jsonp.ThirdBettime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("third_bettime", "")))
	jsonp.ThirdSettlementtime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("third_settlementtime", "")))
	jsonp.GameType = c.DefaultQuery("game_type", "")
	jsonp.GameName = c.DefaultQuery("game_name", "")
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")

	jsonp.OddsType = tool.Atoi(c.DefaultQuery("odds_type", ""))
	jsonp.Version = tool.Int(c.DefaultQuery("version", ""))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	//global.G_LOG.Infof("fcBetRecord ------------1:%v", jsonp.FcBetRecord)
	data, total, betTime, betAmount, netAmount, validAmount := modules.FindPageFcBetRecord(jsonp.PageNo, jsonp.PageSize, &jsonp.FcBetRecord, &jsonp.PageTimeQuery, c)
	//global.G_LOG.Infof("fcBetRecord ------------2:%v", data)
	//global.G_LOG.Infof("fcBetRecord ------------3:%v", total)
	list := []*dos.FcBetRecordResp{}
	for _, v := range data {
		betRecordList := dos.FcBetRecordResp{}
		tool.JsonMapper(v, &betRecordList)
		betRecordList.MerchantName = v.Merchant.MerchantName
		list = append(list, &betRecordList)
	}

	res := dos.FcBetRecordListResp{RecordList: list, TotalBetTime: betTime, TotalBetAmount: betAmount, TotalNetAmount: netAmount, TotalValidBetAmount: validAmount}

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, res)
}

// api: api/fcBetRecord/findByKey
func FindByKeyFcBetRecordControl(c *gin.Context) {
	var jsonp dos.FcBetRecord
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
	data := modules.FindByKeyFcBetRecord(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcBetRecord/update
func UpdateFcBetRecordControl(c *gin.Context) {
	var jsonp dos.FcBetRecord
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

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcBetRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBetRecord/delete
func DeleteFcBetRecordControl(c *gin.Context) {
	var jsonp dos.FcBetRecord
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

	betRecord := modules.FindByKeyFcBetRecordFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, betRecord.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcBetRecord(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcBetRecord/delete
func VenuePlaybackControl(c *gin.Context) {
	var jsonp vo.VenuePlaybackReq
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

	playBackResp, err := venues.GetVenuePlayback(&jsonp)
	if err != nil {
		global.G_LOG.Errorf("GetVenuePlayback venueCode: %s err: %v", jsonp.VenueCode, err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	playbackUrl := playBackResp.Data.PlaybackUrl
	response.SuccessJSON(c, playbackUrl)
}
