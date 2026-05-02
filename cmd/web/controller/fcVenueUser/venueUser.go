package fcVenueUser

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/service/venues"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/shopspring/decimal"
)

func TransferIn(c *gin.Context) {
	jsonp := struct {
		dos.FcVenueUser
		Amount float64 `json:"amount"` // 转入金额
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
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	m := &dos.FcVenueUser{}
	m.Id = jsonp.Id
	data := modules.FindByKeyFcVenueUserFirst(m)
	venueDepositRequest := &venues.VenueDepositRequest{
		UserName:     data.UserName,
		UserId:       data.UserId,
		VenueCode:    data.VenueCode,
		MerchantCode: data.MerchantCode,
		IP:           "8.8.8.8",
		OptBy:        jsonp.UpdateBy,
		Amount:       jsonp.Amount,
		Currency:     data.Currency,
	}
	venueDepositResp := venues.VenueDeposit(venueDepositRequest)
	if venueDepositResp.Code != 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, venueDepositResp.Msg)
		return
	}
	response.SuccessJSON(c, true)
}

func TransferOut(c *gin.Context) {
	jsonp := struct {
		dos.FcVenueUser
		Amount float64 `json:"amount"` // 转出金额
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
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	m := &dos.FcVenueUser{}
	m.Id = jsonp.Id
	data := modules.FindByKeyFcVenueUserFirst(m)

	venueWithdrawRequest := &venues.VenueWithdrawRequest{
		UserName:     data.UserName,
		UserId:       data.UserId,
		VenueCode:    data.VenueCode,
		MerchantCode: data.MerchantCode,
		IP:           "8.8.8.8",
		OptBy:        jsonp.UpdateBy,
		Amount:       jsonp.Amount,
		Currency:     data.Currency,
	}
	venueWithdrawResp := venues.VenueWithdraw(venueWithdrawRequest)
	if venueWithdrawResp.Code != 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, venueWithdrawResp.Msg)
		return
	}
	response.SuccessJSON(c, true)
}

// api: api/fcVenueUser/findPage
func FindPageFcVenueUserByUserIdControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVenueUser
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.UserId = c.DefaultQuery("user_id", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyFcVenueUser(&dos.FcVenueUser{
		UserId: jsonp.UserId,
	})
	var balanceRequest []*venues.VenueBalanceReq
	for _, v := range data {
		balanceRequest = append(balanceRequest, &venues.VenueBalanceReq{
			UserName:     v.UserName,
			UserId:       v.UserId,
			VenueCode:    v.VenueCode,
			Currency:     v.Currency,
			MerchantCode: v.MerchantCode,
		})
	}

	venueBalanceResp := venues.VenueBalancesTimeoutAsync(balanceRequest)
	venuesBalanceMap := map[string]float64{}
	for _, v := range venueBalanceResp {
		venuesBalanceMap[v.VenueCode] = decimal.NewFromFloat(v.Amount).Truncate(2).InexactFloat64()
	}
	pageData, total := modules.FindPageFcVenueUser(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVenueUser)
	var newData []*vo.FcVenueUserByUserIdResp
	tool.JsonMapper(pageData, &newData)
	for _, v := range newData {
		balance, ok := venuesBalanceMap[v.VenueCode]
		if ok {
			v.Amount = tool.String(balance)
		}
	}
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, newData)
}
