// The build tag makes sure the stub is not built in the final build.

package fcVenue

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVenue/save
func SaveFcVenueControl(c *gin.Context) {
	var jsonp vo.VenueCreateRequest
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

	var venue dos.FcVenue
	var venueImgs = make([]*dos.FcVenueImg, len(jsonp.GameTypeImg))
	tool.JsonMapper(jsonp, &venue)
	tool.JsonMapper(jsonp.GameTypeImg, &venueImgs)
	userInfo, ok := c.Get("UserInfo")
	if ok {
		venue.CreateBy = userInfo.(*dos.AdminUser).UserName
		for _, v := range venueImgs {
			v.CreateBy = userInfo.(*dos.AdminUser).UserName
			v.VenueCode = venue.VenueCode
			v.VenueName = venue.VenueName

			_, _ = modules.SaveFcVenueImg(v)
		}
	}

	data, code := modules.SaveFcVenue(&venue)
	if code != ecode.OK {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "保存场馆失败")
		return
	}

	response.SuccessJSON(c, data)
}

// api: api/fcVenue/findPage
func FindPageFcVenueControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVenue
	}{}
	//global.G_LOG.Infof("fcVenue ------------------------------0:%v", jsonp)
	jsonp.PageQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
	jsonp.PageQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))

	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.VenueName = c.DefaultQuery("venue_name", "")
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.VenueType = tool.Atoi(c.DefaultQuery("venue_type", ""))
	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	//global.G_LOG.Infof("fcVenue ------------------------------1:%v", jsonp)
	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	//global.G_LOG.Infof("fcVenue ------------------------------2:%v", jsonp)
	data, total := modules.FindPageFcVenue(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVenue)
	//global.G_LOG.Infof("fcVenue ------------------------------3:%v, %v", data, total)

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVenue/findByKey
func FindByKeyFcVenueControl(c *gin.Context) {
	var jsonp dos.FcVenue
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
	data := modules.FindByKeyFcVenue(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenue/update
func UpdateFcVenueControl(c *gin.Context) {
	var jsonp vo.VenueUpdateRequest
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

	var venue dos.FcVenue
	var venueImgs = make([]*dos.FcVenueImg, len(jsonp.GameTypeImg))
	tool.JsonMapper(jsonp, &venue)
	tool.JsonMapper(jsonp.GameTypeImg, &venueImgs)

	userInfo, ok := c.Get("UserInfo")
	if ok {
		venue.UpdateBy = userInfo.(*dos.AdminUser).UserName
		global.G_DB.
			Where("venue_code=? and merchant_code=?", venue.VenueCode, "").Delete(&dos.FcVenueImg{})
		for _, v := range venueImgs {
			v.CreateBy = userInfo.(*dos.AdminUser).UserName
			v.UpdateBy = userInfo.(*dos.AdminUser).UserName
			v.VenueCode = venue.VenueCode
			v.VenueName = venue.VenueName
			_, _ = modules.SaveFcVenueImg(v)
		}
	}

	data := modules.UpdateFcVenue(&venue)
	if data {
		// 同步状态给场馆游戏列表
		modules.SyncMerchantVenueByStatus(venue.VenueCode, venue.Status)
		modules.SyncVenueGameByStatus(venue.VenueCode, venue.Status)
	}

	response.SuccessJSON(c, data)
}

// api: api/fcVenue/delete
func DeleteFcVenueControl(c *gin.Context) {
	var jsonp dos.FcVenue
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
	venue := modules.FindByKeyFcVenueFirst(&jsonp)
	modules.DeleteFcVenue(&jsonp)
	data := modules.DeleteFcVenueImg(&dos.FcVenueImg{
		VenueCode: venue.VenueCode,
	})
	global.G_DB.
		Where("venue_code=? and merchant_code=?", venue.VenueCode, "").Delete(&dos.FcVenueImg{})
	response.SuccessJSON(c, data)
}
