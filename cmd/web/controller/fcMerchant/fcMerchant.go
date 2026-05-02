// The build tag makes sure the stub is not built in the final build.

package fcMerchant

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// api: api/fcMerchant/save
func SaveFcMerchantControl(c *gin.Context) {
	var jsonp dos.FcMerchant
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
	jsonp.MerchantCode = strings.ToUpper(jsonp.MerchantCode)
	m := modules.FindByKeyFcMerchantFirst(&dos.FcMerchant{
		MerchantCode: jsonp.MerchantCode,
	})
	if m != nil && len(m.Id) > 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户Code已经存在")
		return
	}
	if jsonp.Prefix == "" {
		jsonp.Prefix = jsonp.MerchantCode
	}

	if !modules.CheckAdminUserMerchantPerms(c, jsonp.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	//mP := modules.FindByKeyFcMerchantFirst(&dos.FcMerchant{
	//	Prefix: jsonp.Prefix,
	//})
	//if len(mP.Id) > 0 {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, "游戏账号前缀已经存在")
	//	return
	//}

	// 获取包网商户的代理 ID
	//agentInviteCode := global.G_REDIS.Incr(context.Background(), enmus.Merchant_Agent_Invite_Code_Auto).Val()
	agentInviteCode, err := modules.GetNextIdGeneral(modules.GKEY_INVITE_CODE_INCR)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, "获取商户代理ID错误")
		return
	}
	jsonp.AgentInviteCode = int(agentInviteCode)

	data, code := modules.SaveFcMerchant(&jsonp)
	if code == ecode.FAIL {
		response.FailErrJSON(c, response.ERROR_SERVER, "保存商户失败")
		return
	}

	// 给商户添加渠道和通道关联（容错执行，不影响商户主数据保存）
	func() {
		defer func() {
			if r := recover(); r != nil {
				global.G_LOG.Errorf("[SaveFcMerchantControl] bind pay relation panic merchantCode=%s panic=%v", jsonp.MerchantCode, r)
			}
		}()

		// 获取全部渠道并关联
		payChannels := []dos.FcPayChannelSum{}
		err = global.G_DB.Model(&dos.FcPayChannelSum{}).Where("status = 1").Find(&payChannels).Error
		if err != nil {
			global.G_LOG.Errorf("query all pay_channel_sum err: %v", err)
			return
		}

		for _, v := range payChannels {
			tmp := dos.FcPayChannel{}
			tool.JsonMapper(v, &tmp)
			tmp.MerchantCode = jsonp.MerchantCode
			tmp.Currency = global.CONFIG.General.DefaultCurrency
			tmp.Id = ""

			_, err = modules.SaveFcPayChannel(&tmp)
			if err == ecode.FAIL {
				global.G_LOG.Errorf("SaveFcPayChannel channelCode: %v channelName: %v channelType: %v merchantCode: %v err: %v",
					tmp.ChannelCode, tmp.ChannelName, tmp.ChannelType, tmp.MerchantCode, err)
				continue
			}
		}

		// 获取全部通道并关联
		payments := []dos.FcPaymentSum{}
		err = global.G_DB.Model(&dos.FcPaymentSum{}).Where("status = 1").Find(&payments).Error
		if err != nil {
			global.G_LOG.Errorf("query all payment_sum err: %v", err)
			return
		}

		for _, v := range payments {
			tmp := dos.FcPayment{}
			tool.JsonMapper(v, &tmp)
			tmp.MerchantCode = jsonp.MerchantCode
			tmp.Id = ""

			_, err = modules.SaveFcPayment(&tmp)
			if err != nil {
				global.G_LOG.Errorf("SaveFcPayment channelCode: %v channelName: %v payCode: %v payName: %v merchantCode: %v err: %v",
					tmp.ChannelCode, tmp.ChannelName, tmp.PaymentCode, tmp.PaymentName, tmp.MerchantCode, err)
				continue
			}
		}
	}()

	response.SuccessJSON(c, data)
}

// api: api/fcMerchant/findPage
func FindPageFcMerchantControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.FcMerchant
	}{}
	response.NormalizePageQuery(&jsonp.PageQuery)
	jsonp.Id = c.DefaultQuery("id", "")
	jsonp.MerchantName = c.DefaultQuery("merchant_name", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Logo = c.DefaultQuery("logo", "")
	jsonp.Status = tool.Atoi(c.DefaultQuery("status", ""))

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data, total := modules.FindPageFcMerchant(jsonp.PageNo, jsonp.PageSize, &jsonp.FcMerchant, c)
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/fcMerchant/findByKey
func FindByKeyFcMerchantControl(c *gin.Context) {
	var jsonp dos.FcMerchant
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
	data := modules.FindByKeyFcMerchant(&jsonp, c)
	response.SuccessJSON(c, data)
}

// api: api/fcMerchant/update
func UpdateFcMerchantControl(c *gin.Context) {
	var jsonp dos.FcMerchant
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

	merchant := modules.FindByKeyFcMerchantFirst(&dos.FcMerchant{BaseDos: dos.BaseDos{Id: jsonp.Id}})
	if merchant == nil || len(merchant.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.UpdateFcMerchant(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcMerchant/delete
func DeleteFcMerchantControl(c *gin.Context) {
	var jsonp dos.FcMerchant
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

	merchant := modules.FindByKeyFcMerchantFirst(&dos.FcMerchant{BaseDos: dos.BaseDos{Id: jsonp.Id}})
	if merchant == nil || len(merchant.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "商户不存在")
		return
	}
	if !modules.CheckAdminUserMerchantPerms(c, merchant.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcMerchant(&jsonp)
	response.SuccessJSON(c, data)
}
