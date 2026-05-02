package fcOrderWithdrawPaymentOut

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/paymentOut"

	"github.com/gin-gonic/gin"
)

func PayOutOk(c *gin.Context) {
	jsonp := struct {
		Ids         []string `json:"ids"`          //ID集合
		PaymentCode string   `json:"payment_code"` //选择支付通道
	}{}

	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)
	affected := 0
	for _, v := range jsonp.Ids {
		if len(v) == 0 {
			continue
		}
		order := modules.FindByKeyFcOrderWithdrawPaymentOutFirst(&dos.FcOrderWithdrawPaymentOut{BaseDos: dos.BaseDos{Id: v}})
		if order == nil {
			continue
		}

		if ok, _ := paymentOut.OrderWithdrawAnotherPaySuccess(order.Id, userInfo.UserName); ok {
			affected++
		}

	}
	if affected == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请刷新后重试")
		return
	}
	response.SuccessJSON(c, true)
}

func PayOutFail(c *gin.Context) {
	jsonp := struct {
		Ids            []string `json:"ids"`             //ID集合
		CallbackRemark string   `json:"callback_remark"` //拒绝理由
	}{}

	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)
	affected := 0
	for _, v := range jsonp.Ids {
		if len(v) == 0 {
			continue
		}
		order := modules.FindByKeyFcOrderWithdrawPaymentOutFirst(&dos.FcOrderWithdrawPaymentOut{BaseDos: dos.BaseDos{Id: v}})
		if order == nil {
			continue
		}

		if paymentOut.OrderWithdrawAnotherPayFailRemark(order.Id, jsonp.CallbackRemark, enmus.OrderWithdrawPaymentOutStats_Progress, userInfo.UserName) { //已经提交的状态
			affected++
		}

	}
	if affected == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请刷新后重试")
		return
	}
	response.SuccessJSON(c, true)
}
