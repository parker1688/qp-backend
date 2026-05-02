// // The build tag makes sure the stub is not built in the final build.
package fcTranscation

//
//import (
//	"bootpkg/common/expands/automaticType"
//	"bootpkg/common/response"
//	"bootpkg/common/tool"
//	"bootpkg/pkg/core/modules"
//	"bootpkg/pkg/core/modules/dos"
//	"github.com/gin-gonic/gin"
//	"github.com/go-playground/validator"
//)
//
//// api: api/fcTranscation/save
//func SaveFcTranscationControl(c *gin.Context) {
//	var jsonp dos.FcTranscation
//	err := c.ShouldBind(&jsonp)
//	err1 := validator.New().Struct(jsonp)
//	if err != nil || err1 != nil {
//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
//		return
//	}
//
//	userInfo, ok := c.Get("UserInfo")
//	if ok {
//		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
//	}
//
//	data, _ := modules.SaveFcTranscation(&jsonp)
//	response.SuccessJSON(c, data)
//}
//
//// api: api/fcTranscation/findPage
//func FindPageFcTranscationControl(c *gin.Context) {
//	jsonp := struct {
//		response.PageTimeQuery
//		dos.FcTranscation
//	}{}
//	jsonp.PageTimeQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
//	jsonp.PageTimeQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
//	jsonp.StartAt = c.DefaultQuery("startAt", "")
//	jsonp.EndAt = c.DefaultQuery("endAt", "")
//	jsonp.Id = c.DefaultQuery("id", "")
//	jsonp.UserId = c.DefaultQuery("user_id", "")
//	jsonp.UserName = c.DefaultQuery("user_name", "")
//	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
//
//	jsonp.Remark = c.DefaultQuery("remark", "")
//	jsonp.FundingType = tool.Atoi(c.DefaultQuery("funding_type", "-1"))
//	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
//	jsonp.CreateBy = c.DefaultQuery("create_by", "")
//	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
//	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
//
//	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
//
//	err := validator.New().Struct(jsonp)
//	if err != nil {
//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
//		return
//	}
//	data, total := modules.FindPageFcTranscation(jsonp.PageNo, jsonp.PageSize, &jsonp.FcTranscation, jsonp.PageTimeQuery)
//	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
//}
//
//// api: api/fcTranscation/findByKey
//func FindByKeyFcTranscationControl(c *gin.Context) {
//	var jsonp dos.FcTranscation
//	err := c.ShouldBind(&jsonp)
//	err1 := validator.New().Struct(jsonp)
//	if err != nil || err1 != nil {
//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
//		return
//	}
//	data := modules.FindByKeyFcTranscation(&jsonp)
//	response.SuccessJSON(c, data)
//}
//
//// api: api/fcTranscation/update
//func UpdateFcTranscationControl(c *gin.Context) {
//	var jsonp dos.FcTranscation
//	err := c.ShouldBind(&jsonp)
//	err1 := validator.New().Struct(jsonp)
//	if err != nil || err1 != nil {
//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
//		return
//	}
//
//	userInfo, ok := c.Get("UserInfo")
//	if ok {
//		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
//	}
//
//	data := modules.UpdateFcTranscation(&jsonp)
//	response.SuccessJSON(c, data)
//}
//
//// api: api/fcTranscation/delete
//func DeleteFcTranscationControl(c *gin.Context) {
//	var jsonp dos.FcTranscation
//	err := c.ShouldBind(&jsonp)
//	err1 := validator.New().Struct(jsonp)
//	if err != nil || err1 != nil {
//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
//		return
//	}
//	data := modules.DeleteFcTranscation(&jsonp)
//	response.SuccessJSON(c, data)
//}
