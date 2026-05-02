// The build tag makes sure the stub is not built in the final build.

package fcPayChannelSum

import (
	"bootpkg/common/ecode"
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

// api: api/fcPayChannelSum/save
func SaveFcPayChannelSumControl(c *gin.Context) {
	var jsonp dos.FcPayChannelSum
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
	if jsonp.Currency == "" {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.UpdateBy = jsonp.CreateBy
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	if jsonp.CreateTime.Timer().IsZero() {
		jsonp.CreateTime = automaticType.Now()
	}

	data, err := modules.SaveFcPayChannelSum(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	merchantCodeArr := []*dos.FcMerchant{}
	// 查询所有的商户，给商户添加渠道
	err = global.G_DB.Model(&dos.FcMerchant{}).Where("status = 1").Find(&merchantCodeArr).Error
	if err != nil {
		tmpStr := fmt.Sprintf("query all merchant err: %v", err)
		global.G_LOG.Error(tmpStr)
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	for _, v := range merchantCodeArr {
		tmp := dos.FcPayChannel{}
		tool.JsonMapper(&jsonp, &tmp)
		tmp.MerchantCode = v.MerchantCode
		tmp.Id = ""

		_, saveCode := modules.SaveFcPayChannel(&tmp)
		if saveCode != ecode.OK {
			tmpStr := fmt.Sprintf("SaveFcPayChannel channelCode: %v channelName: %v channelType: %v merchantCode: %v err: %v",
				tmp.ChannelCode, tmp.ChannelName, tmp.ChannelType, tmp.MerchantCode, saveCode)
			global.G_LOG.Error(tmpStr)
			continue
		}
	}

	response.SuccessJSON(c, data)
}

// api: api/fcPayChannelSum/findPage
func FindPageFcPayChannelSumControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcPayChannelSum
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.ChannelName = c.DefaultQuery("channel_name", "")
	jsonp.ChannelCode = c.DefaultQuery("channel_code", "")
	jsonp.Icon = c.DefaultQuery("icon", "")
	jsonp.ChannelType = tool.Atoi(c.DefaultQuery("channel_type", ""))
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.Currency = c.DefaultQuery("currency", "")
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
	data, total := modules.FindPageFcPayChannelSum(jsonp.PageNo, jsonp.PageSize, &jsonp.FcPayChannelSum)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcPayChannelSum/findByKey
func FindByKeyFcPayChannelSumControl(c *gin.Context) {
	var jsonp dos.FcPayChannelSum
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
	data := modules.FindByKeyFcPayChannelSum(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcPayChannelSum/update
func UpdateFcPayChannelSumControl(c *gin.Context) {
	var jsonp dos.FcPayChannelSum
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
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}

	data := modules.UpdateFcPayChannelSum(&jsonp)
	if data {
		//同步渠道子级的状态
		err = global.G_DB.Model(&dos.FcPayChannel{}).Where("channel_code = ?", jsonp.ChannelCode).Updates(map[string]interface{}{
			"icon":   jsonp.Icon,
			"status": jsonp.Status,
		}).Error
		if err != nil {
			global.G_LOG.Error("update FcPayChannel channelCode: %v err: %v", jsonp.ChannelCode, err)
		}
	}
	response.SuccessJSON(c, data)
}

// api: api/fcPayChannelSum/delete
func DeleteFcPayChannelSumControl(c *gin.Context) {
	var jsonp dos.FcPayChannelSum
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

	row := modules.FindByKeyFcPayChannelSumFirst(&jsonp)
	if row.Id == "" {
		response.SuccessJSON(c, struct{}{})
		return
	}

	data := modules.DeleteFcPayChannelSum(&jsonp)

	// 删除所有其他商户的渠道 code
	err = global.G_DB.Where("channel_code = ?", row.ChannelCode).Delete(&dos.FcPayChannel{}).Error
	if err != nil {
		global.G_LOG.Error("delete FcPayChannel channelCode: %v err: %v", row.ChannelCode, err)
	}

	response.SuccessJSON(c, data)
}
