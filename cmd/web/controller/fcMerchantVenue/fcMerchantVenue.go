// The build tag makes sure the stub is not built in the final build.

package fcMerchantVenue

import (
	"bootpkg/cmd/web/model/vo"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcMerchantVenue/save
func SaveFcMerchantVenueControl(c *gin.Context) {
	var jsonp vo.MerchantVenueCreateRequest
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

	var merchantVenue dos.FcMerchantVenue
	var venueImgs = make([]*dos.FcVenueImg, len(jsonp.GameTypeImg))
	tool.JsonMapper(jsonp, &merchantVenue)
	tool.JsonMapper(jsonp.GameTypeImg, &venueImgs)
	userInfo, ok := c.Get("UserInfo")
	var gameTypes []string
	if ok {
		merchantVenue.CreateBy = userInfo.(*dos.AdminUser).UserName
		global.G_DB.
			Where("venue_code=? and merchant_code=?", merchantVenue.VenueCode, merchantVenue.MerchantCode).Delete(&dos.FcVenueImg{})
		for _, v := range venueImgs {
			v.CreateBy = userInfo.(*dos.AdminUser).UserName
			v.VenueCode = merchantVenue.VenueCode
			v.VenueName = merchantVenue.VenueName
			v.MerchantCode = merchantVenue.MerchantCode
			_, _ = modules.SaveFcVenueImg(v)
			gameTypes = append(gameTypes, v.GameType)
		}
	}
	merchantVenue.GameType = strings.Join(gameTypes, ",")
	data, err := modules.SaveFcMerchantVenue(&merchantVenue)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcMerchantVenue/findPage
func FindPageFcMerchantVenueControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcMerchantVenue
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.VenueId = c.DefaultQuery("venue_id", "")
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.MaintainStartTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("maintain_start_time", "")))
	jsonp.MaintainEndTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("maintain_end_time", "")))
	jsonp.ConfigId = c.DefaultQuery("config_id", "")
	jsonp.VenueName = c.DefaultQuery("venue_name", "")

	jsonp.ConfigAlias = c.DefaultQuery("config_alias", "")
	jsonp.Currency = c.DefaultQuery("currency", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcMerchantVenue(jsonp.PageNo, jsonp.PageSize, &jsonp.FcMerchantVenue)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcMerchantVenue/findByKey
func FindByKeyFcMerchantVenueControl(c *gin.Context) {
	var jsonp dos.FcMerchantVenue
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
	data := modules.FindByKeyFcMerchantVenue(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcMerchantVenue/update
func UpdateFcMerchantVenueControl(c *gin.Context) {
	var jsonp vo.MerchantVenueUpdateRequest
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

	// 判断商户是否被关闭
	merchantVenueData := modules.FindByKeyFcMerchantVenueFirst(&dos.FcMerchantVenue{
		BaseDos: dos.BaseDos{Id: jsonp.Id},
	})
	if merchantVenueData.Status == 2 {
		jsonp.Status = merchantVenueData.Status
	}

	var merchantVenue dos.FcMerchantVenue
	var venueImgs = make([]*dos.FcVenueImg, len(jsonp.GameTypeImg))
	tool.JsonMapper(jsonp, &merchantVenue)
	tool.JsonMapper(jsonp.GameTypeImg, &venueImgs)
	var gameTypes []string
	userInfo, ok := c.Get("UserInfo")
	if ok {
		merchantVenue.UpdateBy = userInfo.(*dos.AdminUser).UserName
		err := global.G_DB.
			Where("(venue_code=? and merchant_code=?) or (venue_code=? and merchant_code='')",
				merchantVenue.VenueCode, merchantVenue.MerchantCode, merchantVenue.VenueCode).Delete(&dos.FcVenueImg{}).Error
		if err == nil {
			for _, v := range venueImgs {
				if v.MerchantCode == merchantVenue.MerchantCode || v.MerchantCode == "" {
					v.CreateBy = userInfo.(*dos.AdminUser).UserName
					v.VenueCode = merchantVenue.VenueCode
					v.VenueName = merchantVenue.VenueName
					v.MerchantCode = merchantVenue.MerchantCode
					_, _ = modules.SaveFcVenueImg(v)
					gameTypes = append(gameTypes, v.GameType)
				}
			}
		} else {
			global.G_LOG.Errorf("[UpdateFcMerchantVenueControl] Delete old venue(venue_code=%s, merchant_code=%s) faild: %v",
				merchantVenue.VenueCode, merchantVenue.MerchantCode,
				err.Error())
		}

	}

	data := modules.UpdateFcMerchantVenue(&merchantVenue)
	response.SuccessJSON(c, data)
}

// api: api/fcMerchantVenue/delete
func DeleteFcMerchantVenueControl(c *gin.Context) {
	var jsonp dos.FcMerchantVenue
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
	merchantVenue := modules.FindByKeyFcMerchantVenueFirst(&jsonp)

	global.G_DB.
		Where("venue_code=? and merchant_code=?", merchantVenue.VenueCode, merchantVenue.MerchantCode).Delete(&dos.FcVenueImg{})
	data := modules.DeleteFcMerchantVenue(&jsonp)
	response.SuccessJSON(c, data)
}
