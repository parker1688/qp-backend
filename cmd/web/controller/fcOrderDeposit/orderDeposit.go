package fcOrderDeposit

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/channelData"
	"bootpkg/pkg/service/userTransfer"
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func OrderDepositOk(c *gin.Context) {
	var jsonp *dos.FcOrderDeposit
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	orderDeposit := dos.FcOrderDeposit{}
	err = global.G_DB.Model(&dos.FcOrderDeposit{}).Select("merchant_code").
		Where("id=?", jsonp.Id).First(&orderDeposit).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, orderDeposit.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	global.G_LOG.Infof("OrderDepositOk %v", tool.String(jsonp))
	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
		jsonp.AuthBy = jsonp.UpdateBy
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}
	jsonp.Status = enmus.ORDER_YES_STATUS
	affect, err := userTransfer.UserDepositSuccess(jsonp)
	if !affect {
		response.FailErrJSON(c, response.ERROR_PARAMETER, fmt.Sprintf("操作失败,请刷新后重试：%v", err))
		return
	}
	response.SuccessJSON(c, affect)
}

func OrderDepositNo(c *gin.Context) {
	var jsonp *dos.FcOrderDeposit
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	orderDeposit := dos.FcOrderDeposit{}
	err = global.G_DB.Model(&dos.FcOrderDeposit{}).Select("merchant_code").
		Where("id=?", jsonp.Id).First(&orderDeposit).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, orderDeposit.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	if jsonp.UpdateTime.Timer().IsZero() {
		jsonp.UpdateTime = automaticType.Now()
	}

	nowTimeStr := automaticType.Now().String()
	updates := map[string]interface{}{}
	updates["status"] = enmus.ORDER_NO_STATUS
	updates["remark"] = jsonp.Remark
	updates["auth_by"] = jsonp.UpdateBy
	updates["auth_time"] = nowTimeStr
	updates["update_by"] = jsonp.UpdateBy
	updates["update_time"] = nowTimeStr

	eRow := global.G_DB.Model(&dos.FcOrderDeposit{}).
		Where("id = ? AND status in ?", jsonp.Id, []int{enmus.Order_STATUS_PENDING_PAY, enmus.ORDER_STATUS_WAIT}).
		Updates(updates)
	err = eRow.Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	if eRow.RowsAffected != 1 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请刷新后重试")
		return
	}

	response.SuccessJSON(c, struct{}{})
}

func OrderDepositPush(c *gin.Context) {
	var jsonp *dos.FcOrderDeposit
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if global.CONFIG.Mq.IsInit {
		var vo *dos.FcOrderDeposit
		global.G_DB.Model(&dos.FcOrderDeposit{}).Where("id = ?", jsonp.Id).Take(&vo)
		if len(vo.Id) == 0 || vo.Status != 3 {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "订单不存在或订单状态不对")
			return
		}
		if !modules.CheckAdminUserMerchantPerms(c, vo.MerchantCode) {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
			return
		}
		key := fmt.Sprintf("UserInviteRecord:%s", vo.UserId) + "2024"
		global.G_REDIS.Del(context.Background(), key)
		msg := &channelData.UserRechargeMessage{
			UserId:        vo.UserId,
			UserName:      vo.UserName,
			OrderSn:       vo.OrderSn,
			DepositTime:   vo.CreateTime.String(),
			DepositAmount: vo.Amount,
			T:             time.Now().UnixMicro(),
			ForceStatus:   1,
		}
		errR := channelData.SendUserRecharge(msg)
		if errR != nil {
			global.G_LOG.Errorf(" send kafka err: %v , msg: %s", errR, tool.String(msg))
			response.FailErrJSON(c, response.ERROR_PARAMETER, "推送失败")
			return
		}
	}
	response.SuccessJSON(c, true)
}
