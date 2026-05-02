package userControl

import (
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/srv"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/sms"
	"bootpkg/common/tool"
	"bootpkg/langs"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	vos "bootpkg/pkg/core/modules/vo"
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/kirinlabs/utils"
	"github.com/kirinlabs/utils/encrypt"
	"github.com/shopspring/decimal"
	"github.com/tidwall/sjson"
	"gorm.io/gorm"
)

const (
	RESDIS_KEY_MATERIAL = "RESDIS_KEY_MATERIAL:%v"
)

func Material(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	rechargeValue := global.G_REDIS.Get(context.Background(), fmt.Sprintf(enmus.RECHARGESUCCESSINFO, userInfo.UserId)).Val()
	cacheValue := global.G_REDIS.Get(context.Background(), fmt.Sprintf(RESDIS_KEY_MATERIAL, userInfo.UserId)).Val()
	if len(cacheValue) > 0 {
		if len(rechargeValue) > 2 {
			if global.CONFIG.General.ENV != enmus.Debug {
				global.G_REDIS.Del(context.Background(), fmt.Sprintf(enmus.RECHARGESUCCESSINFO, userInfo.UserId))
			}
			cacheValue, _ = sjson.Set(cacheValue, "recharge_info", rechargeValue)
		}
		response.SuccessMsgJSONData(c, cacheValue, "")
		return
	}
	userInfoData := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: userInfo.UserId,
		//MerchantCode: userInfo.MerchantCode,
	})

	if len(userInfoData.UserId) == 0 {
		response.FailErrJSON(c, ecode.Unauthorized, "not find user info")
		return
	}
	if userInfoData.Level != userInfo.Level || userInfoData.UpdateTime.String() != userInfo.UpdateTime.String() {
		srv.SetUserMaterial(userInfoData)
	}

	userInfoData.Decrypt()
	var material *vo.MaterialResponse
	tool.JsonMapper(userInfoData, &material)
	registerTime := userInfoData.CreateTime.Timer()

	sub := time.Now().AddDate(0, 0, 1).Sub(registerTime)
	//material.Hide()
	material.RegisterDay = int(math.Floor(sub.Hours() / 24))
	if material.WalletPassword != "" {
		material.WalletPassword = "1"
	} else {
		material.WalletPassword = "0"
	}
	global.G_REDIS.Set(context.Background(), fmt.Sprintf(RESDIS_KEY_MATERIAL, userInfo.UserId), tool.String(material), 5*time.Second)
	if len(rechargeValue) > 2 {
		if global.CONFIG.General.ENV != enmus.Debug {
			global.G_REDIS.Del(context.Background(), fmt.Sprintf(enmus.RECHARGESUCCESSINFO, userInfo.UserId))
		}
		var rsi vos.RechargeSuccessInfoVO
		err := tool.JsonUnmarshalFromString(rechargeValue, &rsi)
		if err == nil {
			material.RechargeInfo = rsi
		}
	}
	response.SuccessMsgJSON(c, material, langs.GetWithLocaleGin(c, "message_16"))
}

// 用户退出登录
func Logout(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	userNameM := userInfo.MerchantCode + ":" + userInfo.UserName
	tokenKeyStr := fmt.Sprintf(enmus.REDIS_MEMBER_LOGIN_TOKEN, userNameM)
	//global.G_LOG.Infof("usrNameM: %s tokenKeyStr: %s", userNameM, tokenKeyStr)

	global.G_REDIS.Del(context.Background(), tokenKeyStr)
	response.SuccessMsgJSON(c, "Success", "退出登录成功")
}

func MaterialUpdate(c *gin.Context) {
	var jsonp dos.FcUserMaterial
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	jsonp.Encrypt()

	if jsonp.Sex != 1 && jsonp.Sex != 2 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "性别参数错误")
		return
	}
	if jsonp.NickName == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "昵称不能为空")
		return
	}
	//if jsonp.Avatar == "" {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, "头像不能为空")
	//	return
	//}
	if jsonp.Birthday.Timer().IsZero() {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "生日不能为空")
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	data := map[string]interface{}{}

	// 查询用户信息, 缓存里面时间类型为空的数据不能用，因为自定义 time json 的时候会默认时间为 1949 导致 time.IsZero 的时候判断错误
	uRow := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: userInfo.UserId,
	})
	if uRow.UserId == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "用户不存在")
		return
	}

	//if jsonp.RealName != "" {
	//	data["real_name"] = jsonp.RealName
	//	userInfo.RealName = jsonp.RealName
	//}
	if jsonp.Sex != 0 {
		data["sex"] = jsonp.Sex
		userInfo.Sex = jsonp.Sex
	}
	if !jsonp.Birthday.Timer().IsZero() {
		if uRow.Birthday.Timer().IsZero() { // 生日为空的时候才能更改
			data["birthday"] = jsonp.Birthday
			userInfo.Birthday = jsonp.Birthday
		} else {
			if jsonp.Birthday != uRow.Birthday {
				response.FailErrJSON(c, response.ERROR_SERVER, "已绑定生日，生日更改请联系客服")
				return
			}
		}
	}
	//if jsonp.Email != "" {
	//	data["email"] = jsonp.Email
	//	userInfo.Email = jsonp.Email
	//}

	if jsonp.NickName != "" {
		data["nick_name"] = jsonp.NickName
		userInfo.NickName = jsonp.NickName
	}

	//if jsonp.Tel != "" {
	//	data["tel"] = jsonp.Tel
	//	userInfo.Tel = jsonp.Tel
	//}

	if jsonp.Avatar != "" {
		data["avatar"] = jsonp.Avatar
		userInfo.Avatar = jsonp.Avatar
	}

	//if jsonp.IsVerification {
	//	data["is_verification"] = jsonp.IsVerification
	//	userInfo.IsVerification = jsonp.IsVerification
	//}

	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id=?", userInfo.UserId).Updates(data).Error
	//res := modules.UpdateFcUserMaterial(userInfo)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	srv.SetUserMaterial(userInfo)
	response.SuccessMsgJSON(c, "success", "success")
	//response.SuccessMsgJSON(c, "Success", langs.GetWithLocaleGin(c, "message_16"))
}

// 更新用户头像
func MaterialUpdateAvatar(c *gin.Context) {
	var jsonp dos.FcUserMaterial
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	data := map[string]interface{}{}

	if jsonp.Avatar != "" {
		data["avatar"] = jsonp.Avatar
		userInfo.Avatar = jsonp.Avatar
	}

	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id=?", userInfo.UserId).Updates(data).Error
	//res := modules.UpdateFcUserMaterial(userInfo)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	srv.SetUserMaterial(userInfo)
	response.SuccessMsgJSON(c, "success", "success")
}

// 更新用户性别
func MaterialUpdateSex(c *gin.Context) {
	var jsonp dos.FcUserMaterial
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	data := map[string]interface{}{}

	if jsonp.Sex != 1 && jsonp.Sex != 2 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "性别参数错误")
		return
	}

	data["sex"] = jsonp.Sex
	userInfo.Sex = jsonp.Sex

	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id=?", userInfo.UserId).Updates(data).Error
	//res := modules.UpdateFcUserMaterial(userInfo)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	srv.SetUserMaterial(userInfo)
	response.SuccessMsgJSON(c, "success", "success")
}

// 更新用户昵称
func MaterialUpdateNickName(c *gin.Context) {
	var jsonp dos.FcUserMaterial
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	data := map[string]interface{}{}

	if jsonp.NickName == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "昵称不能为空")
		return
	}

	data["nick_name"] = jsonp.NickName
	userInfo.NickName = jsonp.NickName

	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id=?", userInfo.UserId).Updates(data).Error
	//res := modules.UpdateFcUserMaterial(userInfo)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	srv.SetUserMaterial(userInfo)
	response.SuccessMsgJSON(c, "success", "success")
}

// 更新用户生日
func MaterialUpdateBirthday(c *gin.Context) {
	var jsonp dos.FcUserMaterial
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	data := map[string]interface{}{}

	if jsonp.Birthday.Timer().IsZero() {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "生日不能为空")
		return
	}

	birthDay := time.Time(jsonp.Birthday)
	nowTime := time.Now()
	if birthDay.After(nowTime) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "生日错误")
		return
	}

	data["birthday"] = jsonp.Birthday
	userInfo.Birthday = jsonp.Birthday

	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id=?", userInfo.UserId).Updates(data).Error
	//res := modules.UpdateFcUserMaterial(userInfo)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	srv.SetUserMaterial(userInfo)
	response.SuccessMsgJSON(c, "success", "success")
}

// 更新用户邮箱
func MaterialUpdateEmail(c *gin.Context) {
	var jsonp dos.FcUserMaterial
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	data := map[string]interface{}{}

	if jsonp.Email == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "邮箱不能为空")
		return
	}

	emailFlag, err := tool.IsEmail(jsonp.Email)
	if err != nil || !emailFlag {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "邮箱格式错误")
		return
	}

	data["email"] = jsonp.Email
	userInfo.Email = jsonp.Email

	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id=?", userInfo.UserId).Updates(data).Error
	//res := modules.UpdateFcUserMaterial(userInfo)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	srv.SetUserMaterial(userInfo)
	response.SuccessMsgJSON(c, "success", "success")
}

func MaterialUpdatePhone(c *gin.Context) {
	var jsonp vo.UpdatePhoneReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	if jsonp.VeryCode == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "very code can not empty")
		return
	}

	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	data := map[string]interface{}{}

	if !isCloseSmsVerification() {
		key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", jsonp.Phone, "PhoneBind")
		code := global.G_REDIS.Get(context.Background(), key).Val()

		if code != jsonp.VeryCode {
			response.FailErrJSON(c, ecode.VERFIYCODE_ERROR, langs.GetWithLocaleGin(c, "message_10"))
			return
		}
	}

	if jsonp.Phone == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "手机不能为空")
		return
	}

	userInfo.Tel = jsonp.Phone
	userInfo.Encrypt()
	data["tel"] = userInfo.Tel
	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id=?", userInfo.UserId).Updates(data).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	srv.SetUserMaterial(userInfo)
	response.SuccessMsgJSON(c, nil, "保存成功")
}

func UserReport(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	resp := vo.UserReportResp{}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	// 获取用户账目报表信息
	uReport := dos.FcUserReport{}
	err := global.G_DB.WithContext(ctx).Model(&dos.FcUserReport{}).Where("user_id=?", userInfo.UserId).First(&uReport).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果不存在，不处理

		} else {
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}
	} else {
		tool.JsonMapper(&uReport, &resp)
	}

	// 获取用户游戏报表
	uGameDataArr := []*vo.UserGameDataResp{}
	err = global.G_DB.WithContext(ctx).Model(&dos.FcUserGameReport{}).Where("user_id=?", userInfo.UserId).Scan(&uGameDataArr).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果不存在，不处理

		} else {
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}
	}

	// 查询用户余额
	amount := 0.00
	err = global.G_DB.WithContext(ctx).Model(&dos.FcUserWallet{}).Where("user_id=?", userInfo.UserId).Pluck("ava_amount", &amount).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

		} else {
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}
	}

	resp.UserName = userInfo.UserName
	resp.UserId = userInfo.UserId
	resp.MerchantCode = userInfo.MerchantCode
	resp.Vip = userInfo.Vip
	resp.Level = userInfo.Level
	resp.Amount = amount
	resp.GameList = uGameDataArr

	response.SuccessMsgJSON(c, &resp, "success")
}

// WalletPasswordUpdate
//
//	@Description: 设置支付密码
//	@param c
func WalletPasswordUpdate(c *gin.Context) {

	var jsonp vo.WalletPasswordUpdateReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	if userInfo.Tel == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "phone not allow empty")
		return
	}

	//如果关闭了短信验证模块，则取消改逻辑
	if !isCloseSmsVerification() {
		key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", userInfo.Tel, "WalletPasswordUpdate")
		code := global.G_REDIS.Get(context.Background(), key).Val()

		if code != jsonp.VeryCode {
			response.FailErrJSON(c, ecode.INPUT_VERFIYCODE, langs.GetWithLocaleGin(c, "message_10"))
			return
		}
	}

	NewPassword := encrypt.Sha256(jsonp.WalletPassword + global.CONFIG.General.ApiSHA256Salt)
	userInfo.WalletPassword = NewPassword
	res := modules.UpdateFcUserMaterial(userInfo)

	if res {
		srv.SetUserMaterial(userInfo)
		response.SuccessJSON(c, langs.GetWithLocaleGin(c, "message_16"))
		return
	}

	response.FailErrJSON(c, response.ERROR_SERVER, langs.GetWithLocaleGin(c, "message_17"))
}

func PasswordUpdate(c *gin.Context) {

	var jsonp vo.PasswordUpdateReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	userLogin := modules.FindByKeyFcUserLoginFirst(&dos.FcUserLogin{BaseDos: dos.BaseDos{Id: userInfo.UserId}})
	OldPassword := encrypt.Sha256(jsonp.OldPassword + global.CONFIG.General.ApiSHA256Salt)

	if userLogin.Password != OldPassword {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "Password verification error")
		return
	}

	if jsonp.NewPassword != jsonp.ConfirmPassword {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "Password inconsistency")
		return
	}

	NewPassword := encrypt.Sha256(jsonp.NewPassword + global.CONFIG.General.ApiSHA256Salt)
	userLogin.Password = NewPassword
	res := modules.UpdateFcUserLogin(userLogin)
	if res {
		srv.SetUserMaterial(userInfo)
		response.SuccessJSON(c, langs.GetWithLocaleGin(c, "message_16"))
		return
	}

	response.FailErrJSON(c, response.ERROR_SERVER, langs.GetWithLocaleGin(c, "message_17"))
}

func PasswordReset(c *gin.Context) {

	var jsonp vo.PasswordResetReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	//userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	//userInfoCache := userInfoF.(*dos.FcUserMaterial) //用户信息
	userInfo := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{UserName: jsonp.UserName})

	if len(userInfo.UserId) == 0 {
		response.FailErrJSON(c, ecode.INPUT_VERFIYCODE, langs.GetWithLocaleGin(c, "message_8"))
		return
	}
	userInfo.Decrypt()
	//if userInfo.Tel != jsonp.Phone {
	//	response.FailErrJSON(c, ecode.INPUT_VERFIYCODE, langs.GetWithLocaleGin(c, "message_8"))
	//	return
	//}
	//key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", userInfo.Tel, "PasswordReset")
	//code := global.G_REDIS.Get(context.Background(), key).Val()
	//
	//if code != jsonp.VeryCode {
	//	response.FailErrJSON(c, ecode.INPUT_VERFIYCODE, langs.GetWithLocaleGin(c, "message_10"))
	//	return
	//}

	userLogin := modules.FindByKeyFcUserLoginFirst(&dos.FcUserLogin{BaseDos: dos.BaseDos{Id: userInfo.UserId}})
	OldPassword := encrypt.Sha256(jsonp.OldPassword + global.CONFIG.General.ApiSHA256Salt)

	if userLogin.Password != OldPassword {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "密码错误")
		return
	}

	if jsonp.NewPassword != jsonp.ConfirmPassword {
		response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_18"))
		return
	}

	NewPassword := encrypt.Sha256(jsonp.NewPassword + global.CONFIG.General.ApiSHA256Salt)
	userLogin.Password = NewPassword
	res := modules.UpdateFcUserLogin(userLogin)
	if res {
		srv.SetUserMaterial(userInfo)

		//删除密码错误次数过多
		userNameM := fmt.Sprintf("%s:%s", userInfo.MerchantCode, userInfo.UserName)
		global.G_REDIS.Del(context.Background(), enmus.MEMBER_REDIS_LOGIN_ERR_COUNT+userNameM).Val()
		response.SuccessMsgJSON(c, "Success", langs.GetWithLocaleGin(c, "message_16"))
		return
	}

	response.FailErrJSON(c, response.ERROR_SERVER, langs.GetWithLocaleGin(c, "message_17"))
}

// 忘记密码
func PasswordForget(c *gin.Context) {
	var jsonp vo.PasswordForgotReq
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 验证用户
	encryptTel := (&dos.FcUserMaterial{}).EncryptData(jsonp.Phone)
	userInfo := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserName: jsonp.UserName,
		Tel:      encryptTel, // 查询需要加密后的电话
	})
	if len(userInfo.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "该账号对应的手机号码错误")
		return
	}
	userInfo.Decrypt()

	// 验证验证码
	if !isCloseSmsVerification() {
		key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", userInfo.Tel, "ForgotPass")
		code := global.G_REDIS.Get(context.Background(), key).Val()

		if code != jsonp.VeryCode {
			response.FailErrJSON(c, ecode.INPUT_VERFIYCODE, langs.GetWithLocaleGin(c, "message_10"))
			return
		}
	}

	// 设置忘记密码更新密码的验证key
	global.G_REDIS.Set(context.Background(),
		fmt.Sprintf(enmus.MEMBER_REDIS_PASS_ForgotToken, userInfo.UserName),
		encryptTel,
		time.Duration(5)*time.Minute)

	response.SuccessMsgJSON(c, "Success", langs.GetWithLocaleGin(c, "message_16"))
}

// 忘记密码更新新密码
func PasswordForgotUpdate(c *gin.Context) {
	var jsonp vo.PasswordForgotUpdateReq
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 获取验证key
	authKey := fmt.Sprintf(enmus.MEMBER_REDIS_PASS_ForgotToken, jsonp.UserName)
	encryptTel := global.G_REDIS.Get(context.Background(), authKey).Val()
	if len(encryptTel) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "令牌已过期")
		return
	}

	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)

	// 验证用户
	userInfo := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserName:     jsonp.UserName,
		MerchantCode: merchantCode,
		Tel:          encryptTel,
	})
	if len(userInfo.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "该账号对应的手机号码错误")
		return
	}
	userInfo.Decrypt()

	// 判断验证密码
	if jsonp.NewPassword != jsonp.ConfirmPassword {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "密码不一致")
		return
	}

	// 更新用户密码
	userLogin := modules.FindByKeyFcUserLoginFirst(&dos.FcUserLogin{BaseDos: dos.BaseDos{Id: userInfo.UserId}})
	userLogin.Password = encrypt.Sha256(jsonp.NewPassword + global.CONFIG.General.ApiSHA256Salt)
	ret := modules.UpdateFcUserLogin(userLogin)
	if !ret {
		response.FailErrJSON(c, response.ERROR_SERVER, "更新密码失败")
		return
	}

	global.G_REDIS.Del(context.Background(), authKey)

	response.SuccessMsgJSON(c, nil, "修改成功")
}

func PhoneVeryCode(c *gin.Context) {
	var jsonp vo.PhoneVeryCodeReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	/*userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息*/

	//if userInfo.Email == jsonp.Email {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, "邮箱已绑定")
	//	return
	//}

	//else{
	//if jsonp.Phone != userInfo.DecryptData(userInfo.Tel) {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, langs.GetWithLocaleGin(c, "message_32"))
	//	return
	//}
	//}
	//if jsonp.Tag == "WalletPasswordUpdate" {
	//	res := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{Tel: jsonp.Phone})
	//	if len(res.UserId) > 0 && res.UserId != userInfo.UserId {
	//		response.FailErrJSON(c, response.ERROR_PARAMETER, "The phone number has already been bound")
	//		return
	//	}
	//}
	/*userInfo.Decrypt()
	if jsonp.Tag == "PhoneBind" {
		userMaterial := &dos.FcUserMaterial{Tel: jsonp.Phone}
		userMaterial.Encrypt()
		//res := modules.FindByKeyFcUserMaterialFirst(userMaterial)
		//if len(res.UserId) > 0 {
		//	response.FailErrJSON(c, ecode.EAMIL_HAS_EXIST, "手机号码已经存在")
		//	return
		//}
		userInfo.Tel = jsonp.Phone
	}*/

	/*var code int
	if global.CONFIG.General.ENV == "Debug" {
		code = 123456
		global.G_REDIS.Set(context.Background(), key, code, time.Duration(5)*time.Minute)
		response.SuccessMsgJSON(c, langs.GetWithLocaleGin(c, "message_16"), "Success")
		return
	}*/

	if !isCloseSmsVerification() { // 状态为关闭则不需要发短信
		key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", jsonp.Phone, jsonp.Tag)
		lockKey := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode_Lock, "0", jsonp.Phone, jsonp.Tag)

		count := global.G_REDIS.Exists(context.Background(), lockKey).Val()
		if count > 0 {
			response.FailErrJSON(c, response.ERROR_SERVER, "Can only be sent once in 2 minutes")
			return
		}

		code := tool.RandInt(100000, 999999)

		err = sms.Handle(jsonp.Phone, tool.String(code))
		if err != nil {
			global.G_LOG.Errorf("[PhoneVeryCode] captcha send failed: %v", err.Error())
			response.FailErrJSON(c, response.ERROR_SERVER, "验证码发送失败")
			return
		}

		global.G_REDIS.Set(context.Background(), key, code, time.Duration(5)*time.Minute)
		global.G_REDIS.Set(context.Background(), lockKey, code, time.Duration(2)*time.Minute)
	}

	response.SuccessMsgJSON(c, langs.GetWithLocaleGin(c, "message_16"), "Success")
}

func Verification(c *gin.Context) {
	var jsonp vo.PhoneVerificationReq
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)

	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息
	if !isCloseSmsVerification() {
		key := fmt.Sprintf(enmus.MEMBER_REDIS_PHONE_VeryCode, "0", userInfo.Tel, "Verification")
		code := global.G_REDIS.Get(context.Background(), key).Val()

		if code != jsonp.Code {
			response.FailErrJSON(c, ecode.VERFIYCODE_ERROR, langs.GetWithLocaleGin(c, "message_10"))
			return
		}
	}

	global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id=?", userInfo.UserId).Update("is_verification", true)
	response.SuccessMsgJSON(c, langs.GetWithLocaleGin(c, "message_15"), "Success")
}

// VipProgress
//
//	@Description: VIP进度条
//	@param c
func VipProgress(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	user := userInfoF.(*dos.FcUserMaterial) //用户信息

	userInfo := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: user.UserId,
	})

	nextVip := modules.FindByKeyFcVipFirst(&dos.FcVip{Level: userInfo.Level + 1})
	nowVip := modules.FindByKeyFcVipFirst(&dos.FcVip{Level: userInfo.Level})
	totalBetAmount, err := global.G_REDIS.Get(context.Background(),
		fmt.Sprintf(enmus.UserTotalBetAmountKey, user.UserId)).Float64()
	if err != nil {
		global.G_LOG.Error(err.Error())
	}

	var nextHaveBetAmount, betProgress, nextNeedBet float64

	if totalBetAmount > 0 {
		//下一级已投注流水
		nextHaveBetAmount = decimal.NewFromFloat(totalBetAmount).
			Sub(decimal.NewFromFloat(nowVip.MinBetAmount)).Truncate(2).InexactFloat64()
	}

	//下一级投注共需流水
	nextBetAmount := decimal.NewFromFloat(nextVip.MinBetAmount).
		Sub(decimal.NewFromFloat(nowVip.MinBetAmount)).Truncate(2).InexactFloat64()

	if nextVip.MinBetAmount > 0 && nextBetAmount > 0 {
		nextNeedBet = decimal.NewFromFloat(nextBetAmount).Sub(decimal.NewFromFloat(nextHaveBetAmount)).Truncate(2).InexactFloat64()
		betProgress = decimal.NewFromFloat(nextHaveBetAmount).Div(decimal.NewFromFloat(nextBetAmount)).Truncate(2).InexactFloat64()
	}

	//totalRechargeAmount, err := global.G_REDIS.Get(context.Background(),
	//	fmt.Sprintf(enmus.UserTotalRechargeAmountKey, user.UserId)).Float64()
	//
	//if err != nil {
	//	global.G_LOG.Error(err.Error())
	//	totalRechargeAmount = 0
	//}

	vipProgressResp := vo.VipProgressResp{
		NowVip:        userInfo.Vip,
		NextVip:       nextVip.VipName,
		NextNeedBet:   nextNeedBet,
		NowBetAmount:  nextHaveBetAmount,
		NextBetAmount: nextBetAmount,
		//NextRechargeAmount: nexVip.MinRecharegeAmount,
		Progress: betProgress,
		//BetProgress:        betProgress,
		//MinWithdrawAmount:  nowVip.MinWithdrawAmount,
		//MinRecharegeAmount: nowVip.MinRecharegeAmount,
		//WithdrawalFee: fmt.Sprintf("%0.f", nowVip.WithdrawalFee) + "%",
		NowLevel: userInfo.Level,
		//TotalBet:      totalBetAmount,
		//TotalRecharge: totalRechargeAmount,
	}

	response.SuccessMsgJSON(c, vipProgressResp, langs.GetWithLocaleGin(c, "message_16"))
}
func getUserSumDepositAmount(userId string) float64 {
	var amount float64
	global.G_DB.Model(&dos.FcOrderDeposit{}).Select("sum(amount)").Where("user_id = ? and status=3", userId).Scan(&amount)
	return amount
}
func ManualTransferWallet(c *gin.Context) {
	var jsonp struct {
		Manual bool `form:"manual" json:"manual"`
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	userInfo.ManualTransferWallet = automaticType.BitBool(jsonp.Manual)
	res := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", userInfo.UserId).Update("manual_transfer_wallet", jsonp.Manual).Error
	if res != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, langs.GetWithLocaleGin(c, "message_17"))
		return
	}
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userInfo.UserName)
	global.G_REDIS.Set(context.Background(), memberRedisKey, utils.Json(userInfo), 30*24*time.Hour)
	response.SuccessJSON(c, langs.GetWithLocaleGin(c, "message_16"))
}

func DayReport(c *gin.Context) {
	userInfoF, _ := c.Get(vo.USER_NAME_INFO_G)
	userInfo := userInfoF.(*dos.FcUserMaterial) //用户信息

	var userReport *dos.FcUserReport
	global.G_DB.Model(&dos.FcUserReport{}).Where("user_id=?", userInfo.UserId).Take(&userReport)
	response.SuccessMsgJSON(c, userReport, langs.GetWithLocaleGin(c, "message_16"))
}

func isCloseSmsVerification() bool {
	smsVerification := modules.FindByKeyDictsDetailFirst(&dos.DictsDetail{
		DictsTypeCode: "Cilent_System_Settings",
		DictsTag:      "SmsVerification",
	})

	return smsVerification.DictsValue == "1"
}
