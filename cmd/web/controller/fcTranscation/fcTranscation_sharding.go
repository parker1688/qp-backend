// The build tag makes sure the stub is not built in the final build.

package fcTranscation

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcTranscation/save
func SaveFcTranscationControl(c *gin.Context) {
	var jsonp dos.FcTranscation
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

	jsonp.CreateTime = automaticType.Time(time.Now())
	jsonp.UpdateTime = jsonp.CreateTime
	data, _ := modules.SaveFcTranscationSharding(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcTranscation/findPage
func FindPageFcTranscationControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcTranscationSharding
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))

	jsonp.Remark = c.DefaultQuery("remark", "")
	jsonp.FundingType = tool.Atoi(c.DefaultQuery("funding_type", "-1"))
	//jsonp.CreateTime = c.DefaultQuery("create_time", "")
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcTranscationSharding(jsonp.PageNo, jsonp.PageSize, &jsonp.FcTranscationSharding, jsonp.PageTimeQuery, c)
	//排序
	//sort.Slice(data, func(i, j int) bool {
	//	return data[i].CreateTime.Timer().Unix() > data[j].CreateTime.Timer().Unix()
	//})
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcTranscation/findByKey
func FindByKeyFcTranscationControl(c *gin.Context) {
	var jsonp dos.FcTranscation
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
	data := modules.FindByKeyFcTranscationSharding(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcTranscation/update
func UpdateFcTranscationControl(c *gin.Context) {
	var jsonp dos.FcTranscation
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

	data := modules.UpdateFcTranscationSharding(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcTranscation/delete
func DeleteFcTranscationControl(c *gin.Context) {
	var jsonp dos.FcTranscation
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

	fcTranscation := modules.FindByKeyFcTranscationFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, fcTranscation.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcTranscationSharding(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcTranscation/cashflow
func CashFlowFcTranscationControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcTranscationSharding
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.FundingType = tool.Atoi(c.DefaultQuery("funding_type", "-1"))
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcTranscationSharding(jsonp.PageNo, jsonp.PageSize, &jsonp.FcTranscationSharding, jsonp.PageTimeQuery, c)

	resp := map[string]interface{}{}
	resp["list"] = data
	resp["quotaCovertStatis"] = modules.AcumulateQuotaCovertStatisData(&jsonp.FcTranscationSharding, jsonp.PageTimeQuery, c)
	resp["onlineDepositStatis"] = modules.AcumulateOnlineDepositStatisData(&jsonp.FcTranscationSharding, jsonp.PageTimeQuery, c)
	resp["manualDepositStatis"] = modules.AcumulateManualDepositStatisData(&jsonp.FcTranscationSharding, jsonp.PageTimeQuery, c)
	resp["wthdrawStatis"] = modules.AcumulateWithdrawStatisData(&jsonp.FcTranscationSharding, jsonp.PageTimeQuery, c)
	resp["promotionStatis"] = modules.AcumulatePromotionStatisData(&jsonp.FcTranscationSharding, jsonp.PageTimeQuery, c)
	resp["rebateStatis"] = modules.AcumulateRebateStatisData(&jsonp.FcTranscationSharding, jsonp.PageTimeQuery, c)

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, resp)
}
