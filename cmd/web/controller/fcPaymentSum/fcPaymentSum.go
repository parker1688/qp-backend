// The build tag makes sure the stub is not built in the final build.

package fcPaymentSum

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcPaymentSum/save
func SaveFcPaymentSumControl(c *gin.Context) {
	var jsonp dos.FcPaymentSum
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
		jsonp.UpdateBy = jsonp.CreateBy
	}
	// 查询渠道code是否存在
	if jsonp.ChannelCode == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "channel_code is required")
	}
	channelSum := modules.FindByKeyFcPayChannelSumFirst(&dos.FcPayChannelSum{
		ChannelCode: jsonp.ChannelCode,
	})
	if channelSum.ChannelCode == "" {
		global.G_LOG.Errorf("channel_code: %s not exist", jsonp.ChannelCode)
		response.FailErrJSON(c, response.ERROR_PARAMETER, "channel_code not exist")
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.CreateTime.Timer().IsZero() {
		jsonp.CreateTime = automaticType.Now()
	}

	data, err := modules.SaveFcPaymentSum(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	merchantCodeArr := []*dos.FcMerchant{}
	// 查询所有的商户，给商户添加通道
	err = global.G_DB.Model(&dos.FcMerchant{}).Where("status = 1").Find(&merchantCodeArr).Error
	if err != nil {
		tmpStr := fmt.Sprintf("query all merchant err: %v", err)
		global.G_LOG.Error(tmpStr)
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	for _, v := range merchantCodeArr {
		tmp := dos.FcPayment{}
		tool.JsonMapper(&jsonp, &tmp)
		tmp.MerchantCode = v.MerchantCode
		tmp.Id = ""

		_, err = modules.SaveFcPayment(&tmp)
		if err != nil {
			tmpStr := fmt.Sprintf("SaveFcPayment channelCode: %v channelName: %v payCode: %v payName: %v merchantCode: %v err: %v",
				tmp.ChannelCode, tmp.ChannelName, tmp.PaymentCode, tmp.PaymentName, tmp.MerchantCode, err)
			global.G_LOG.Error(tmpStr)
			continue
		}
	}

	response.SuccessJSON(c, data)
}

// api: api/fcPaymentSum/findPage
func FindPageFcPaymentSumControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcPaymentSum
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.PaymentName = c.DefaultQuery("payment_name", "")
	jsonp.PaymentCode = c.DefaultQuery("payment_code", "")
	jsonp.PayId = c.DefaultQuery("pay_id", "")
	jsonp.ChannelName = c.DefaultQuery("channel_name", "")
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))

	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.AmountRange = c.DefaultQuery("amount_range", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcPaymentSum(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPaymentSum)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcPaymentSum/findByKey
func FindByKeyFcPaymentSumControl(c *gin.Context) {
	var jsonp dos.FcPaymentSum
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
	data := modules.FindByKeyFcPaymentSum(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPaymentSum/update
func UpdateFcPaymentSumControl(c *gin.Context) {
	var jsonp dos.FcPaymentSum
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

	// 查询渠道code是否存在
	if jsonp.ChannelCode == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "channel_code is required")
	}
	channelSum := modules.FindByKeyFcPayChannelSumFirst(&dos.FcPayChannelSum{
		ChannelCode: jsonp.ChannelCode,
	})
	if channelSum.ChannelCode == "" {
		global.G_LOG.Errorf("channel_code: %s not exist", jsonp.ChannelCode)
		response.FailErrJSON(c, response.ERROR_PARAMETER, "channel_code not exist")
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}

	data := modules.UpdateFcPaymentSum(&jsonp)
	if data {
		//同步通道子级的状态
		err = global.G_DB.Model(&dos.FcPayment{}).Where("payment_code = ? AND pay_id = ?", jsonp.PaymentCode, jsonp.PayId).Updates(map[string]interface{}{
			"status": jsonp.Status,
		}).Error
		if err != nil {
			global.G_LOG.Error("update FcPayment paymentCode: %v err: %v", jsonp.PaymentCode, err)
		}
	}
	response.SuccessJSON(c, data)
}

// api: api/fcPaymentSum/delete
func DeleteFcPaymentSumControl(c *gin.Context) {
	var jsonp dos.FcPaymentSum
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

	row := modules.FindByKeyFcPaymentSumFirst(&jsonp)
	if row.Id == "" {
		response.SuccessJSON(c, struct{}{})
		return
	}

	data := modules.DeleteFcPaymentSum(&jsonp)

	// 删除所有其他商户的通道 code的通道id
	err = global.G_DB.Where("payment_code = ? and pay_id = ?", row.PaymentCode, row.PayId).Delete(&dos.FcPayment{}).Error
	if err != nil {
		global.G_LOG.Error("delete FcPayment paymentCode: %v err: %v", row.PaymentCode, err)
	}

	response.SuccessJSON(c, data)
}
