package fcOrderPromotion

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/userTransfer"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func UpdatePromotionNo(c *gin.Context) {
	var jsonp dos.FcOrderPromotion
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
	affect := global.G_DB.Model(&dos.FcOrderPromotion{}).Where("id = ? and status = ?", jsonp.Id, enmus.ORDER_PENDING_STATUS).Update("status", enmus.ORDER_NO_STATUS)
	if affect.RowsAffected > 0 {
		response.SuccessJSON(c, true)
		return
	}
	response.FailErrJSON(c, response.ERROR_PARAMETER, "请刷新后重试")
	return
}

func UpdatePromotionOK(c *gin.Context) {
	var jsonp dos.FcOrderPromotion
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
	affect, err := userTransfer.PromotionOk(&jsonp)
	if !affect {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试：%v", err))
		return
	}
	response.SuccessJSON(c, affect)
}
