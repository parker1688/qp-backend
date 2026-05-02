package walletControl

import (
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/langs"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func BindBank(c *gin.Context) {
	var jsonp vo.FcUserWithdrawBankBind
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if len(jsonp.Currency) == 0 {
		jsonp.Currency = global.CONFIG.General.DefaultCurrency
	}

	jsonp.BankAddress = strings.Replace(jsonp.BankAddress, ".", "", -1)
	jsonp.BankAddress = strings.Replace(jsonp.BankAddress, "-", "", -1)

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	userInfo.Decrypt()
	if !isCloseSmsVerification() {
		key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", userInfo.Tel, "BindBank")
		code := global.G_REDIS.Get(context.Background(), key).Val()

		if code != jsonp.VeryCode {
			response.FailErrJSON(c, ecode.VERFIYCODE_ERROR, langs.GetWithLocaleGin(c, "message_10"))
			return
		}
	}

	numberHash := tool.HMACSignatureSha256("sha256", jsonp.AccountNumber)
	var count, total int64
	global.G_DB.Model(&dos.FcUserWithdrawBankBind{}).Where("number_hash=?", numberHash).Count(&count)
	if count > 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_38"))
		return
	}

	global.G_DB.Model(&dos.FcUserWithdrawBankBind{}).Where("user_id=?", userInfo.UserId).Count(&total)
	if total >= 5 {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "最多绑定5个银行卡")
		return
	}

	bankType := modules.FindByKeyFcBindBankTypeFirst(&dos.FcBindBankType{
		BankCode: jsonp.AccountBankType,
		//MerchantCode: merchantCode,
	})

	if len(bankType.Id) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "不支持该银行卡")
		return
	}
	if jsonp.IsDefault == 0 {
		jsonp.IsDefault = 1
	}
	if total == 0 {
		jsonp.IsDefault = 2
	}
	var withdraw = &dos.FcUserWithdrawBankBind{
		UserId:          userInfo.UserId,
		UserName:        userInfo.UserName,
		BankAddress:     jsonp.BankAddress,
		AccountNumber:   jsonp.AccountNumber,
		AccountHolder:   jsonp.AccountHolder,
		MerchantCode:    merchantCode,
		Currency:        jsonp.Currency,
		AccountBankType: bankType.BankName,
		AccountBankCode: bankType.BankCode,
		NumberHash:      numberHash,
		IsDefault:       jsonp.IsDefault,
	}
	withdraw.Encrypt()

	b, _ := modules.SaveFcUserWithdrawBankBind(withdraw)
	response.SuccessMsgJSON(c, b, "添加成功")
}

func BindBankType(c *gin.Context) {
	var jsonp dos.FcBindBankType
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyFcBindBankType(&dos.FcBindBankType{
		Currency: jsonp.Currency,
	}, nil)
	response.SuccessJSON(c, data)
}

func ChannelBankImg(c *gin.Context) {
	var jsonp dos.FcChannelBankImg
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyFcChannelBankImg(&dos.FcChannelBankImg{
		Status:      1,
		ChannelCode: jsonp.ChannelCode,
		PaymentCode: jsonp.PaymentCode,
	}, nil)
	response.SuccessJSON(c, data)
}

func GetBindBank(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	data := modules.FindByKeyFcUserWithdrawBankBind(&dos.FcUserWithdrawBankBind{
		UserId: userInfo.UserId,
	})

	//var BindBankResp []vo.GetBindBankResp
	//for _, v := range data {
	//
	//	BindBankResp = append(BindBankResp, vo.GetBindBankResp{
	//		Id:              v.Id,
	//		BankAddress:     v.BankAddress,
	//		AccountNumber:   v.AccountNumber,
	//		AccountHolder:   v.AccountHolder,
	//		AccountBankType: v.AccountBankType,
	//		cc:              v.Currency,
	//		IsDefault:       v.IsDefault,
	//	})
	//}
	for _, v := range data {
		v.Decrypt()
	}
	toData := make([]*vo.GetBindBankResp, 0, len(data))
	tool.MapLoose(data, &toData)
	for _, v := range toData {
		v.Hide()
		img := modules.FindByKeyFcChannelBankImgFirst(&dos.FcChannelBankImg{
			Status:      1,
			PaymentCode: v.AccountBankCode,
		})
		v.Img = img.Img
		v.ImgPath = img.ImgPath
		v.Icon = img.Icon
		v.IconPath = img.IconPath

	}
	response.SuccessJSON(c, toData)
}

func BindBlockchain(c *gin.Context) {
	var jsonp vo.FcUserWithdrawBlockchainBind
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	userInfo.Decrypt()
	if !isCloseSmsVerification() {
		key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", userInfo.Tel, "Blockchain")
		code := global.G_REDIS.Get(context.Background(), key).Val()

		if code != jsonp.VeryCode {
			response.FailErrJSON(c, ecode.VERFIYCODE_ERROR, langs.GetWithLocaleGin(c, "message_10"))
			return
		}
	}
	data := modules.FindByKeyFcUserWithdrawBlockchainBind(&dos.FcUserWithdrawBlockchainBind{
		PaymentCode: jsonp.PaymentCode,
		UserId:      userInfo.UserId,
	})
	if len(data) >= 5 {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "最多绑定5个地址")
		return
	}

	for _, v := range data {
		v.Decrypt()
		if v.BlockchainAddress == jsonp.BlockchainAddress {
			response.FailErrJSON(c, response.ERROR_DEFAULT, "地址重复")
			return
		}
	}
	if jsonp.IsDefault == 0 {
		jsonp.IsDefault = 1
	}
	if len(data) == 0 {
		jsonp.IsDefault = 2
	}
	w := &dos.FcUserWithdrawBlockchainBind{
		UserId:            userInfo.UserId,
		UserName:          userInfo.UserName,
		Blockchain:        jsonp.PaymentName,
		BlockchainAddress: jsonp.BlockchainAddress,
		ContractType:      jsonp.PaymentCode,
		MerchantCode:      merchantCode,
		PaymentCode:       jsonp.PaymentCode,
		PaymentName:       jsonp.PaymentName,
		IsDefault:         jsonp.IsDefault,
		RealName:          jsonp.RealName,
	}
	w.Encrypt()

	b, err := modules.SaveFcUserWithdrawBlockchainBind(w)
	response.SuccessMsgJSON(c, b, "添加成功")
}

func BindBlockchainType(c *gin.Context) {
	data := modules.FindByKeyFcBindBlockchainType(&dos.FcBindBlockchainType{})
	response.SuccessJSON(c, data)
}

func GetBindBlockchain(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	data := modules.FindByKeyFcUserWithdrawBlockchainBind(&dos.FcUserWithdrawBlockchainBind{
		UserId: userInfo.UserId,
	})
	for _, v := range data {
		v.Decrypt()
	}
	toData := make([]*vo.GetBindBlockchainResp, 0, len(data))
	tool.JsonMapper(data, &toData)

	for _, v := range toData {
		img := modules.FindByKeyFcChannelBankImgFirst(&dos.FcChannelBankImg{
			Status:      1,
			PaymentCode: v.PaymentCode,
		})
		v.Img = img.Img
		v.ImgPath = img.ImgPath
		v.Icon = img.Icon
		v.IconPath = img.IconPath
	}
	response.SuccessJSON(c, toData)
}

func GetBindOnlie(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	data := modules.FindByKeyFcUserWithdrawOnlineBind(&dos.FcUserWithdrawOnlineBind{
		UserId: userInfo.UserId,
	})
	for _, v := range data {
		v.Decrypt()
	}
	toData := make([]*vo.GetBindOnlieResp, 0, len(data))
	tool.JsonMapper(data, &toData)
	for _, v := range toData {
		img := modules.FindByKeyFcChannelBankImgFirst(&dos.FcChannelBankImg{
			Status:      1,
			ChannelCode: v.ChannelCode,
		})
		v.Img = img.Img
		v.ImgPath = img.ImgPath
		v.Icon = img.Icon
		v.IconPath = img.IconPath
	}
	response.SuccessJSON(c, toData)
}

func BindBankDefault(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBankBind
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	global.G_DB.Model(&dos.FcUserWithdrawBankBind{}).Where("user_id = ?", userInfo.UserId).Update("is_default", 1)
	global.G_DB.Model(&dos.FcUserWithdrawBankBind{}).Where("id = ? and user_id = ?", jsonp.Id, userInfo.UserId).Update("is_default", 2)
	response.SuccessJSON(c, true)
}

func BindBlockchainDefault(c *gin.Context) {
	var jsonp dos.FcUserWithdrawBlockchainBind
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	global.G_DB.Model(&dos.FcUserWithdrawBlockchainBind{}).Where("user_id = ?", userInfo.UserId).Update("is_default", 1)
	global.G_DB.Model(&dos.FcUserWithdrawBlockchainBind{}).Where("id = ? and user_id = ?", jsonp.Id, userInfo.UserId).Update("is_default", 2)
	response.SuccessJSON(c, true)
}

func BindOnlineDefault(c *gin.Context) {
	var jsonp dos.FcUserWithdrawOnlineBind
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	global.G_DB.Model(&dos.FcUserWithdrawOnlineBind{}).Where("user_id = ?", userInfo.UserId).Update("is_default", 1)
	global.G_DB.Model(&dos.FcUserWithdrawOnlineBind{}).Where("id = ? and user_id = ?", jsonp.Id, userInfo.UserId).Update("is_default", 2)
}
func BindOnline(c *gin.Context) {
	var jsonp vo.FcUserWithdrawOnlineBind
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial)
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	userInfo.Decrypt()
	if !isCloseSmsVerification() {
		key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", userInfo.Tel, "BindOnline")
		code := global.G_REDIS.Get(context.Background(), key).Val()

		if code != jsonp.VeryCode {
			response.FailErrJSON(c, ecode.VERFIYCODE_ERROR, langs.GetWithLocaleGin(c, "message_10"))
			return
		}
	}
	data := modules.FindByKeyFcUserWithdrawOnlineBind(&dos.FcUserWithdrawOnlineBind{
		UserId:      userInfo.UserId,
		ChannelCode: jsonp.ChannelCode,
	})

	if len(data) >= 5 {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "最多绑定5个账号")
		return
	}
	if jsonp.IsDefault == 0 {
		jsonp.IsDefault = 1
	}
	if len(data) == 0 {
		jsonp.IsDefault = 2
	}
	for _, v := range data {
		v.Decrypt()
		if v.AccountNumber == jsonp.AccountNumber {
			response.FailErrJSON(c, response.ERROR_DEFAULT, "绑定账号重复")
			return
		}
	}
	w := &dos.FcUserWithdrawOnlineBind{
		UserId:        userInfo.UserId,
		UserName:      userInfo.UserName,
		MerchantCode:  merchantCode,
		AccountHolder: jsonp.AccountHolder,
		AccountNumber: jsonp.AccountNumber,
		ChannelCode:   jsonp.ChannelCode,
		ChannelName:   jsonp.ChannelName,
		IsDefault:     jsonp.IsDefault,
	}
	w.Encrypt()
	b, err := modules.SaveFcUserWithdrawOnlineBind(w)
	response.SuccessMsgJSON(c, b, "添加成功")
}

func isCloseSmsVerification() bool {
	smsVerification := modules.FindByKeyDictsDetailFirst(&dos.DictsDetail{
		DictsTypeCode: "Cilent_System_Settings",
		DictsTag:      "SmsVerification",
	})

	return smsVerification.DictsValue == "1"
}
