// The build tag makes sure the stub is not built in the final build.

package fcVenueGame

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcVenueGame/save
func SaveFcVenueGameControl(c *gin.Context) {
	var jsonp dos.FcVenueGame
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

	// 同场馆下游戏编码唯一，重复时给出明确提示
	existList := modules.FindByKeyFcVenueGame(&dos.FcVenueGame{
		VenueCode: jsonp.VenueCode,
		GameCode:  jsonp.GameCode,
	})
	if len(existList) > 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "该场馆下游戏编码已存在")
		return
	}

	data, code := modules.SaveFcVenueGame(&jsonp)
	if code == ecode.FAIL {
		response.FailErrJSON(c, response.ERROR_SERVER, "保存游戏失败")
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcVenueGame/findPage
func FindPageFcVenueGameControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcVenueGame
	}{}
	jsonp.PageQuery.PageNo = tool.Atoi(c.DefaultQuery("current", "1"))
	jsonp.PageQuery.PageSize = tool.Atoi(c.DefaultQuery("pageSize", "10"))
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.VenueName = c.DefaultQuery("venue_name", "")
	jsonp.VenueCode = c.DefaultQuery("venue_code", "")
	jsonp.VenueType = c.DefaultQuery("venue_type", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")
	jsonp.GameCode = c.DefaultQuery("game_code", "")
	jsonp.Hot = tool.Atoi(c.DefaultQuery("hot", ""))
	jsonp.Sort = tool.Atoi(c.DefaultQuery("sort", ""))
	jsonp.ImgIcon = c.DefaultQuery("img_icon", "")
	jsonp.GameName = c.DefaultQuery("game_name", "")
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.GameType = c.DefaultQuery("game_type", "")
	jsonp.GameNamePinyin = c.DefaultQuery("game_name_pinyin", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcVenueGame(jsonp.PageNo, jsonp.PageSize, &jsonp.FcVenueGame)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcVenueGame/findByKey
func FindByKeyFcVenueGameControl(c *gin.Context) {
	var jsonp dos.FcVenueGame
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
	data := modules.FindByKeyFcVenueGame(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueGame/update
func UpdateFcVenueGameControl(c *gin.Context) {
	var jsonp dos.FcVenueGame
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

	data := modules.UpdateFcVenueGame(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueGame/delete
func DeleteFcVenueGameControl(c *gin.Context) {
	var jsonp dos.FcVenueGame
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
	data := modules.DeleteFcVenueGame(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcVenueGame/find/venue
func FindFcVenueControl(c *gin.Context) {
	data := modules.FindByKeyFcVenue(&dos.FcVenue{})
	response.SuccessJSON(c, data)
}
