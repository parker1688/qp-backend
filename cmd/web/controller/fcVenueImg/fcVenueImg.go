// The build tag makes sure the stub is not built in the final build.

package fcVenueImg

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVenueImg/save
func SaveFcVenueImgControl(c *gin.Context) {
	var jsonp dos.FcVenueImg
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
	}

	data, err := modules.SaveFcVenueImg(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcVenueImg/findPage
func FindPageFcVenueImgControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVenueImg
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.VenueName = c.DefaultQuery("venue_name", "")
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.GameType = c.DefaultQuery("game_type", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.ImgIcon = c.DefaultQuery("img_icon", "")
	jsonp.ImgBar = c.DefaultQuery("img_bar", "")
	jsonp.LinkIcon = c.DefaultQuery("link_icon", "")
	jsonp.LinkBar = c.DefaultQuery("link_bar", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcVenueImg(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVenueImg)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVenueImg/findByKey
func FindByKeyFcVenueImgControl(c *gin.Context) {
	var jsonp dos.FcVenueImg
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
	data := modules.FindByKeyFcVenueImg(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueImg/update
func UpdateFcVenueImgControl(c *gin.Context) {
	var jsonp dos.FcVenueImg
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

	data := modules.UpdateFcVenueImg(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueImg/delete
func DeleteFcVenueImgControl(c *gin.Context) {
	var jsonp dos.FcVenueImg
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
	data := modules.DeleteFcVenueImg(&jsonp)
	response.SuccessJSON(c, data)
}
