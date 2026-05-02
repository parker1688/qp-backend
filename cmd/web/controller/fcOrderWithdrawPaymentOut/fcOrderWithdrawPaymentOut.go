// The build tag makes sure the stub is not built in the final build.

package fcOrderWithdrawPaymentOut

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/userTransfer"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

func failWithdrawStatsConflict(c *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, gorm.ErrRecordNotFound) ||
		strings.Contains(err.Error(), "update withdraw status fail") ||
		strings.Contains(err.Error(), "update withdraw payment out status fail") ||
		strings.Contains(err.Error(), "update order withdraw another status fail") ||
		strings.Contains(err.Error(), "update withdraw another pay status fail") {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请刷新后重试")
		return true
	}
	return false
}

// api: api/fcOrderWithdrawPaymentOut/save
func SaveFcOrderWithdrawPaymentOutControl(c *gin.Context) {
	var jsonp dos.FcOrderWithdrawPaymentOut
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

	data, _ := modules.SaveFcOrderWithdrawPaymentOut(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderWithdrawPaymentOut/findPage
func FindPageFcOrderWithdrawPaymentOutControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcOrderWithdrawPaymentOut
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.PageTimeQuery.StartAt = c.DefaultQuery("startAt", "")
	jsonp.PageTimeQuery.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.OrderSn = c.DefaultQuery("order_sn", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")

	jsonp.Status = tool.Atoi(c.DefaultQuery("status", "-1"))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.OrderType = tool.Atoi(c.DefaultQuery("order_type", ""))

	jsonp.ChannelId = c.DefaultQuery("channel_id", "")
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")
	jsonp.PaymentId = c.DefaultQuery("payment_id", "")
	jsonp.PaymentCode = c.DefaultQuery("payment_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcOrderWithdrawPaymentOut(jsonp.PageTimeQuery, &jsonp.FcOrderWithdrawPaymentOut, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcOrderWithdrawPaymentOut/findByKey
func FindByKeyFcOrderWithdrawPaymentOutControl(c *gin.Context) {
	var jsonp dos.FcOrderWithdrawPaymentOut
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
	data := modules.FindByKeyFcOrderWithdrawPaymentOut(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderWithdrawPaymentOut/update
func UpdateFcOrderWithdrawPaymentOutControl(c *gin.Context) {
	var jsonp dos.FcOrderWithdrawPaymentOut
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

	orderWithdrawPaymentOut := modules.FindByKeyFcOrderWithdrawPaymentOutFirst(&jsonp)
	if orderWithdrawPaymentOut == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, orderWithdrawPaymentOut.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcOrderWithdrawPaymentOut(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderWithdrawPaymentOut/delete
func DeleteFcOrderWithdrawPaymentOutControl(c *gin.Context) {
	var jsonp dos.FcOrderWithdrawPaymentOut
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

	orderWithdrawPaymentOut := modules.FindByKeyFcOrderWithdrawPaymentOutFirst(&jsonp)
	if orderWithdrawPaymentOut == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, orderWithdrawPaymentOut.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcOrderWithdrawPaymentOut(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcOrderWithdrawPaymentOut/upWithdrawStats
func UpWithdrawStatsFcOrderWithdrawPaymentOutControl(c *gin.Context) {
	var jsonp struct {
		Id             string `json:"id"`
		WithdrawStatus int    `json:"withdraw_status"`
		Remark         string `json:"remark"`
	}

	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	param := dos.FcOrderWithdrawPaymentOut{
		BaseDos: dos.BaseDos{Id: jsonp.Id},
	}

	orderWithdrawPaymentOut := modules.FindByKeyFcOrderWithdrawPaymentOutFirst(&param)
	if orderWithdrawPaymentOut == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在")
		return
	}

	if orderWithdrawPaymentOut.Status == enmus.OrderWithdrawPaymentOutStats_Failed ||
		orderWithdrawPaymentOut.Status == enmus.OrderWithdrawPaymentOutStats_Success {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "订单已结束")
		return
	}

	userInfoF, _ := c.Get("UserInfo")
	userInfo := userInfoF.(*dos.AdminUser)

	var withdrawStatus int

	switch jsonp.WithdrawStatus {
	case enmus.OrderWithdrawStats_Yes: // 人工已打款
		err = global.G_DB.Transaction(func(tx *gorm.DB) error {
			eRow := tx.Model(&dos.FcOrderWithdrawPaymentOut{}).
				Where("id = ? AND status in ? AND withdraw_status = ?", jsonp.Id, []int{
					enmus.OrderWithdrawPaymentOutStats_Prepare,
					enmus.OrderWithdrawPaymentOutStats_Progress,
				}, enmus.OrderWithdrawStats_No).
				Updates(map[string]interface{}{
					"withdraw_status": enmus.OrderWithdrawStats_ManualYes,
					"remark":          jsonp.Remark,
					"update_by":       userInfo.UserName,
					"update_time":     automaticType.Now(),
				})
			if eRow.Error != nil {
				return eRow.Error
			}
			if eRow.RowsAffected != 1 {
				return gorm.ErrRecordNotFound
			}

			// 结束提款管理订单
			eRow = tx.Model(&dos.FcOrderWithdraw{}).Where("order_sn = ? AND another_pay_status = ?",
				orderWithdrawPaymentOut.OrderSn, enmus.OrderWithdrawAnotherPayStats_Progress).
				Update("another_pay_status", enmus.OrderWithdrawAnotherPayStats_Success)
			if eRow.Error != nil {
				return eRow.Error
			}
			if eRow.RowsAffected != 1 {
				return gorm.ErrRecordNotFound
			}

			err1 := tx.Model(&dos.FcTranscation{}).Where("related_id = ? AND funding_type = ?",
				orderWithdrawPaymentOut.OrderSn, int(userTransfer.TranWithdraw)).
				Update("status", 1).Error
			if err1 != nil {
				return err1
			}
			return nil
		})

		if err != nil {
			if failWithdrawStatsConflict(c, err) {
				return
			}
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}
		//提款成功系统邮件提示
		order := orderWithdrawPaymentOut
		go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawSuccess, order.Amount)
		withdrawStatus = enmus.OrderWithdrawStats_ManualYes
	case enmus.OrderWithdrawStats_No: // 人工未打款
		param.OrderSn = orderWithdrawPaymentOut.OrderSn
		param.WithdrawStatus = enmus.OrderWithdrawStats_ManualNo
		param.Remark = jsonp.Remark
		param.UpdateBy = userInfo.UserName
		err := userTransfer.UserWithdrawPaymentOutNo(&param)
		if err != nil {
			if failWithdrawStatsConflict(c, err) {
				return
			}
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}
		//提款成功系统邮件提示
		order := orderWithdrawPaymentOut
		go modules.SendSystemMail(order.UserId, order.MerchantCode, enmus.MailType_WithdrawFail, order.Amount)
		withdrawStatus = enmus.OrderWithdrawStats_ManualNo
	default:
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	response.SuccessJSON(c, struct {
		WithdrawStatus int `json:"withdraw_status"`
	}{
		WithdrawStatus: withdrawStatus,
	})
}
