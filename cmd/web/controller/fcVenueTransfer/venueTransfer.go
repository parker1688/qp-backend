package fcVenueTransfer

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/service/userTransfer"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// DepositOk
//
//	@Description: 场馆转入成功
//	@param c
func DepositOk(c *gin.Context) {
	var jsonp dos.FcVenueTransfer
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
	affect, err := userTransfer.VenueDepositSuccess(&jsonp)
	if !affect {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试：%v", err))
		return
	}
	response.SuccessJSON(c, affect)
}

// DepositNo
//
//	@Description: 场馆转入失败(解除冻结金额)
//	@param c
func DepositNo(c *gin.Context) {
	var jsonp dos.FcVenueTransfer
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
	affect, err := userTransfer.VenueDepositFail(&jsonp)
	if !affect {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试：%v", err))
		return
	}
	response.SuccessJSON(c, affect)
}

// WithdrawNo
//
//	@WithdrawNo: 场馆转出失败
//	@param c
func WithdrawNo(c *gin.Context) {
	var jsonp dos.FcVenueTransfer
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
	affect := global.G_DB.Model(&dos.FcVenueTransfer{}).Where("id = ? and status = ?", jsonp.Id, 0).Updates(map[string]interface{}{
		"update_by": jsonp.UpdateBy,
		"status":    2,
	})
	if affect.RowsAffected > 0 {
		response.SuccessJSON(c, true)
		return
	}
	response.FailErrJSON(c, response.ERROR_PARAMETER, "请刷新后重试")
}

// WithdrawOk
//
//	@Description: 场馆转出成功
//	@param c
func WithdrawOk(c *gin.Context) {
	var jsonp dos.FcVenueTransfer
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
	affect, err := userTransfer.VenueWithdrawSuccess(&jsonp)
	if !affect {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试：%v", err))
		return
	}
	response.SuccessJSON(c, affect)
}
