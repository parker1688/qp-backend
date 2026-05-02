// The build tag makes sure the stub is not built in the final build.

package fcCustomerOrder

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"bootpkg/pkg/service/channelData"
	"bootpkg/pkg/service/userTransfer"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

// api: api/fcCustomerOrder/save
func SaveFcCustomerOrderControl(c *gin.Context) {
	var jsonp dos.FcCustomerOrder
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
	if jsonp.UserId == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "userId is empty")
		return
	}
	mem := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: jsonp.UserId,
	})
	if mem.UserName == "" {
		global.G_LOG.Errorf("userId=%s query userMaterial err: %v", jsonp.UserId, err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, "userId query is empty")
		return
	}
	jsonp.UserName = mem.UserName
	jsonp.Status = 1 // 默认为处理中

	// 判断是否改了该玩家的商户码
	if jsonp.MerchantCode != mem.MerchantCode {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "不能更改用户商户")
		return
	}

	if !modules.CheckAdminUserMerchantPerms(c, mem.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
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

	jsonp.FlowAmount = jsonp.Amount * float64(jsonp.FlowMultiple)

	data, err := modules.SaveFcCustomerOrder(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	response.SuccessJSON(c, data)
}

// api: api/fcCustomerOrder/findPage
func FindPageFcCustomerOrderControl(c *gin.Context) {
	jsonp := struct {
		response.PageTimeQuery
		dos.FcCustomerOrder
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.LastStartAt = c.DefaultQuery("last_startAt", "")
	jsonp.LastEndAt = c.DefaultQuery("last_endAt", "")

	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))

	jsonp.BonusType = tool.Atoi(c.DefaultQuery("bonus_type", ""))
	jsonp.Title = c.DefaultQuery("title", "")
	jsonp.Remark = c.DefaultQuery("remark", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.SolveRemark = c.DefaultQuery("solve_remark", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcCustomerOrder(jsonp.PageNo, jsonp.PageSize, &jsonp.FcCustomerOrder, &jsonp.PageTimeQuery, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcCustomerOrder/findByKey
func FindByKeyFcCustomerOrderControl(c *gin.Context) {
	var jsonp dos.FcCustomerOrder
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
	data := modules.FindByKeyFcCustomerOrder(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcCustomerOrder/update
func UpdateFcCustomerOrderControl(c *gin.Context) {
	var jsonp dos.FcCustomerOrder
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
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	jsonp.UpdateTime = automaticType.Now()

	data := modules.UpdateFcCustomerOrder(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcCustomerOrder/update
func UpdateFcCustomerOrderStatusControl(c *gin.Context) {
	var jsonp vo.CustomerOrderUpdateStatusVO
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
	if !ok {
		response.FailErrJSON(c, response.ERROR_SERVER, "UserInfo error")
		return
	}
	userInfoM, ok := userInfo.(*dos.AdminUser)
	if !ok {
		response.FailErrJSON(c, response.ERROR_SERVER, "UserInfo error2")
		return
	}

	row := dos.FcCustomerOrder{}
	err = global.G_DB.Model(dos.FcCustomerOrder{}).Where(`id = ?`, jsonp.Id).First(&row).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "not found")
		return
	}

	if row.Status == jsonp.Status {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "状态一样，不需要更改")
		return
	}

	if !modules.CheckAdminUserMerchantPerms(c, row.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	err = global.G_DB.Model(dos.FcCustomerOrder{}).Where(`id = ?`, jsonp.Id).Updates(map[string]interface{}{
		"status":       jsonp.Status,
		"solve_remark": jsonp.SolveRemark,
		"update_time":  tool.TimeNowString(),
		"update_by":    userInfoM.UserName,
	}).Error
	if err != nil {
		global.G_LOG.Errorf("update customer id=%s order status=%v error:%v", jsonp.Id, jsonp.Status, err)
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	if jsonp.Status != 3 {
		response.SuccessJSON(c, struct{}{})
		return
	}

	if jsonp.Currency == "" {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	// 如果同意，则需要给用户加钱，写红利记录
	orderSn := tool.SnowflakeId()
	fop := dos.FcOrderPromotion{
		ApplyAmount:  row.Amount,
		OrderSn:      orderSn,
		AppleRate:    1,
		ApplyType:    row.BonusType,
		Amount:       row.Amount,
		Status:       enmus.ORDER_YES_STATUS,
		UserName:     row.UserName,
		UserId:       row.UserId,
		TurnOver:     row.FlowMultiple,
		MerchantCode: row.MerchantCode,
		Remake:       jsonp.SolveRemark,
		Currency:     jsonp.Currency,
		CreateBy:     userInfoM.UserName,
		UpdateBy:     userInfoM.UserName,
	}

	amount := row.Amount
	trsAmountType := userTransfer.TranDiscount
	global.G_DB.Transaction(func(tx *gorm.DB) error {

		err = userTransfer.UserAmountChange(tx, amount, trsAmountType, jsonp.Currency, jsonp.SolveRemark, row.UserId, userInfoM.UserName, "",
			modules.GetFcWelfareTitleByBonusType(row.BonusType))
		if err != nil {
			return err
		}

		// 写红利表
		err = tx.Create(&fop).Error
		if err != nil {
			return err
		}

		return err
	})
	if err != nil {
		global.G_LOG.Error("WalletAmountOpt UserAmountChange username: %v err: %v", row.UserName, err)
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	if global.CONFIG.Mq.IsInit {
		promoData := channelData.UserPromotionMessage{}
		promoData.UserId = row.UserId
		promoData.OrderSn = fop.OrderSn
		promoData.ForceStatus = 1
		promoData.T = time.Now().UnixMicro()
		promoData.PromotionTime = fop.CreateTime.String()
		promoData.PromotionAmount = amount
		promoData.PromotionType = 1

		// 发送红利消息给消息队列
		err = channelData.SendUserPromotion(&promoData)
		if err != nil {
			global.G_LOG.Errorf("UpdateFcCustomerOrderStatusControl SendUserPromotion kafka err: %v , msg: %s", err, tool.String(promoData))
		}
	}

	response.SuccessJSON(c, struct{}{})
}

// api: api/fcCustomerOrder/delete
func DeleteFcCustomerOrderControl(c *gin.Context) {
	var jsonp dos.FcCustomerOrder
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

	customerOrder := modules.FindByKeyFcCustomerOrderFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, customerOrder.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcCustomerOrder(&jsonp)
	response.SuccessJSON(c, data)
}
