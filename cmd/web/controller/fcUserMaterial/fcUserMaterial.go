// The build tag makes sure the stub is not built in the final build.

package fcUserMaterial

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

// api: api/fcUserMaterial/save
func SaveFcUserMaterialControl(c *gin.Context) {
	var jsonp dos.FcUserMaterial
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
		jsonp.CreateBy = userInfo.(*dos.AdminUser).UserName
	}

	data, _ := modules.SaveFcUserMaterial(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserMaterial/findPage
func FindPageFcUserMaterialControl(c *gin.Context) {
	jsonp := struct {
		dos.FcUserMaterial
		response.PageTimeQuery
	}{}

	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.NickName = c.DefaultQuery("nick_name", "")
	jsonp.RealName = c.DefaultQuery("real_name", "")
	jsonp.ParentId = c.DefaultQuery("parent_id", "")
	jsonp.AgentId = c.DefaultQuery("agent_id", "")
	jsonp.AgentName = c.DefaultQuery("agent_name", "")
	jsonp.Sex = tool.Atoi(c.DefaultQuery("sex", ""))
	jsonp.Tel = c.DefaultQuery("tel", "")
	jsonp.Email = c.DefaultQuery("email", "")
	jsonp.Qq = c.DefaultQuery("qq", "")
	jsonp.Wx = c.DefaultQuery("wx", "")
	jsonp.Address = c.DefaultQuery("address", "")
	jsonp.Birthday = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("birthday", "")))
	jsonp.Level = tool.Atoi(c.DefaultQuery("level", ""))
	jsonp.AgentInviteCode = tool.Atoi(c.DefaultQuery("agent_invite_code", ""))
	jsonp.Vip = c.DefaultQuery("vip", "")
	jsonp.RegisterIp = c.DefaultQuery("register_ip", "")
	jsonp.LastLoginIp = c.DefaultQuery("last_login_ip", "")
	jsonp.RegistVisitorId = c.DefaultQuery("regist_visitor_id", "")

	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	jsonp.Nation = c.DefaultQuery("nation", "")
	jsonp.Language = c.DefaultQuery("language", "")
	jsonp.Avatar = c.DefaultQuery("avatar", "")
	jsonp.AgentSubId = c.DefaultQuery("agent_sub_id", "")
	jsonp.AgentSubName = c.DefaultQuery("agent_sub_name", "")
	jsonp.CreateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("create_time", "")))
	jsonp.CreateBy = c.DefaultQuery("create_by", "")
	jsonp.UpdateTime = automaticType.Time(tool.StrToTimeZero(c.DefaultQuery("update_time", "")))
	jsonp.UpdateBy = c.DefaultQuery("update_by", "")

	jsonp.StartAt = c.DefaultQuery("startAt", "")
	jsonp.EndAt = c.DefaultQuery("endAt", "")
	jsonp.LastStartAt = c.DefaultQuery("last_startAt", "")
	jsonp.LastEndAt = c.DefaultQuery("last_endAt", "")
	jsonp.Website = c.DefaultQuery("website", "")
	jsonp.PageTimeQuery.IsFree = c.DefaultQuery("is_free", "")
	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 对电话号码,真实姓名等进行加密处理
	jsonp.Encrypt()
	//global.G_LOG.Infof("username: %s orgTel: %s encrypt tel: %s", jsonp.UserName, tel, jsonp.Tel)

	data, total := modules.FindPageFcUserMaterial(jsonp.PageNo, jsonp.PageSize,
		&jsonp.FcUserMaterial, jsonp.PageTimeQuery, false, c)

	// 处理用户数据
	_, userInfoArr := HandleUserInfo(data, true, true)

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, userInfoArr)
}

// api: api/fcUserMaterial/findByKey
func FindByKeyFcUserMaterialControl(c *gin.Context) {
	var jsonp dos.FcUserMaterial
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
	data := modules.FindByKeyFcUserMaterial(&jsonp)

	// 处理用户数据
	_, userInfoArr := HandleUserInfo(data, true, false)

	var userDetailLis []*vo.MaterialDetail
	for _, v := range userInfoArr {
		uDetail := vo.MaterialDetail{}
		tool.JsonMapper(v, &uDetail)
		q1 := global.G_DB.Model(&dos.FcOrderDeposit{})            //充值
		q2 := global.G_DB.Model(&dos.FcOrderWithdrawPaymentOut{}) //提现
		q3 := global.G_DB.Model(&dos.FcBetRecord{})               //有效下注，利润
		q5 := global.G_DB.Model(&dos.FcOrderPromotion{})          //福利总表
		q6 := global.G_DB.Model(&dos.FcUserRebateRecords{})       //反水

		merchantLis := modules.GetAdminUserMerchantList(c)

		q1 = q1.Where("merchant_code in ?", merchantLis)
		q2 = q2.Where("merchant_code in ?", merchantLis)
		q3 = q3.Where("merchant_code in ?", merchantLis)
		q5 = q5.Where("merchant_code in ?", merchantLis)
		q6 = q6.Where("merchant_code in ?", merchantLis)

		uid := v.UserId
		rechargeAmount := 0.0
		q1.Select("sum(amount) as rechargeAmount").Where("user_id = ? and status=3", uid).Scan(&rechargeAmount)
		withDrawaAmount := 0.0
		q2.Select("sum(amount) as withDrawaAmount").Where("user_id = ? and status=3", uid).Scan(&withDrawaAmount)
		validBetAmount := 0.0
		q3.Select("sum(valid_betamount) as validBetAmount").Where("user_id = ?", uid).Scan(&validBetAmount)
		betAmount := 0.0
		q3.Select("sum(bet_amount) as betAmount").Where("user_id = ?", uid).Scan(&betAmount)
		netAmount := 0.0
		q3.Select("sum(net_amount) as netAmount").Where("user_id = ?", uid).Scan(&netAmount)
		promotionAmount := 0.0
		q5.Select("sum(amount) as promotionAmount").Where("user_id = ?", uid).Scan(&promotionAmount)
		rebateAmount := 0.0
		q6.Select("sum(bonus_amount) as rebateAmount").Where("user_id = ?", uid).Scan(&rebateAmount)

		uDetail.PromotionAmount = promotionAmount
		uDetail.WithdrawalAmount = withDrawaAmount
		uDetail.RebateAmount = rebateAmount
		uDetail.BetAmount = betAmount
		uDetail.ValidBetamount = validBetAmount
		uDetail.WinAmount = netAmount
		uDetail.RechargeAmount = rechargeAmount

		userDetailLis = append(userDetailLis, &uDetail)
	}

	response.SuccessJSON(c, userDetailLis)
}

// api: api/fcUserMaterial/detail
func UserMaterialDetail(c *gin.Context) {
	var jsonp dos.FcUserMaterial
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
	data := modules.FindByKeyFcUserMaterialFirst(&jsonp)

	if !modules.CheckAdminUserMerchantPerms(c, data.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	// 处理用户数据
	u := HandleUserInfoRow(data, true, true)
	uDetail := vo.MaterialDetail{}
	tool.JsonMapper(u, &uDetail)

	uDetail.IsLogin = 1 // 默认在线
	userNameM := u.MerchantCode + ":" + u.UserName
	tokenKeyStr := fmt.Sprintf(enmus.REDIS_MEMBER_LOGIN_TOKEN, userNameM)
	uExist := global.G_REDIS.Exists(context.Background(), tokenKeyStr).Val()
	if uExist == 0 {
		uDetail.IsLogin = 2
	}

	uReport := dos.FcUserReport{}
	err = global.G_DB.Model(&dos.FcUserReport{}).Where("user_id = ?", u.UserId).First(&uReport).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.SuccessJSON(c, uDetail)
			return
		}
		global.G_LOG.Errorf("FindByKeyFcUserMaterialControl query FcUserReport userId: %v err:%v", u.UserId, err)
		response.SuccessJSON(c, uDetail)
		return
	}

	uDetail.PromotionAmount = uReport.PromotionAmount
	uDetail.RebateAmount = uReport.RebateAmount
	uDetail.BetAmount = uReport.BetAmount
	uDetail.ValidBetamount = uReport.ValidBetamount
	uDetail.WinAmount = uReport.WinAmount
	uDetail.RechargeAmount = uReport.RechargeAmount

	response.SuccessJSON(c, uDetail)
}

// api: api/fcUserMaterial/update
func UpdateFcUserMaterialControl(c *gin.Context) {
	var jsonp dos.FcUserMaterial
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

	user := modules.FindByKeyFcUserMaterialFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, user.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	if jsonp.Birthday.Timer().Year() < 1970 {
		jsonp.Birthday = automaticType.Time(tool.StrToTimeZero("1970-01-01 00:00:00"))
	}
	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateFcUserMaterial(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/fcUserMaterial/delete
func DeleteFcUserMaterialControl(c *gin.Context) {
	var jsonp dos.FcUserMaterial
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

	user := modules.FindByKeyFcUserMaterialFirst(&jsonp)
	if !modules.CheckAdminUserMerchantPerms(c, user.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	data := modules.DeleteFcUserMaterial(&jsonp)
	response.SuccessJSON(c, data)
}

// 处理一些用户数据
func HandleUserInfo(data []*dos.FcUserMaterial, moneyFlag bool, isPrivate bool) ([]string, []*vo.FcUserMaterialVO) {
	dataLen := len(data)
	userIdArr := make([]string, dataLen)
	userInfoArr := make([]*vo.FcUserMaterialVO, dataLen)
	userInfoMap := make(map[string]*vo.FcUserMaterialVO, dataLen)

	for k, v := range data {
		userIdArr[k] = v.UserId

		tmpU := HandleUserInfoRow(data[k], false, isPrivate)

		userInfoArr[k] = tmpU
		userInfoMap[v.UserId] = tmpU
	}

	// 是否需要获取用户余额
	if moneyFlag {
		userWallets := modules.FindByUserIdsFcUserWallet(userIdArr)
		// 有一些用户没有创建钱包
		for _, v := range userWallets {
			userInfo, ok := userInfoMap[v.UserId]
			if ok {
				userInfo.Currency = v.Currency
				userInfo.TotalAmount = v.TotalAmount
				userInfo.AvaAmount = v.AvaAmount
				userInfo.FronzenAmount = v.FronzenAmount
				userInfo.IsLock = v.IsLock
			}
		}
	}

	return userIdArr, userInfoArr
}

// 处理一些用户数据
func HandleUserInfoRow(data *dos.FcUserMaterial, moneyFlag bool, isPrivate bool) *vo.FcUserMaterialVO {
	startAt := time.Now().Add(-7 * 24 * time.Hour)
	u := vo.FcUserMaterialVO{}
	data.Decrypt()

	tool.JsonMapper(data, &u)

	p := &vo.FcPrivate{}
	p.Email = data.Email
	p.Alipay = data.Alipay
	p.Tel = data.Tel
	p.AlipayRealname = data.AlipayRealname
	p.RealName = data.RealName
	p.WalletPassword = data.WalletPassword
	if isPrivate {
		PrivateDataHandler(p)
	}
	u.Email = p.Email
	u.Tel = p.Tel
	u.RealName = p.RealName

	u.IsActive = 2 // 默认为不活跃
	if !data.LastLoginTime.Timer().IsZero() {
		lastLoginTime := time.Time(data.LastLoginTime)
		if lastLoginTime.After(startAt) {
			u.IsActive = 1
		}
	}

	/*registIpInfo, err := tool.IPCityInfo(data.RegisterIp)
	if err != nil {
		global.G_LOG.Errorf("username=%s registIP=%s IPCityInfo err: %v", data.UserName, data.RegisterIp, err)
	} else {
		registIpCity := registIpInfo.CountryName
		if registIpInfo.RegionName != "" {
			registIpCity += "|" + registIpInfo.RegionName
		}
		if registIpInfo.CityName != "" {
			registIpCity += "|" + registIpInfo.RegionName
		}
		u.RegisterIpCity = registIpCity
	}

	lastIpInfo, err := tool.IPCityInfo(data.LastLoginIp)
	if err != nil {
		global.G_LOG.Errorf("username=%s lastIP=%s IPCityInfo err: %v", data.UserName, data.LastLoginIp, err)
	} else {
		lastIpCity := lastIpInfo.CountryName
		if lastIpInfo.RegionName != "" {
			lastIpCity += "|" + lastIpInfo.RegionName
		}
		if registIpInfo.CityName != "" {
			lastIpCity += "|" + lastIpInfo.RegionName
		}
		u.LastLoginIpCity = lastIpCity
	}*/

	// 是否需要获取用户余额
	if moneyFlag {
		userWallets := modules.FindByKeyFcUserWalletFirst(&dos.FcUserWallet{
			UserId: data.UserId,
		})
		if userWallets.Id != "" {
			u.Currency = userWallets.Currency
			u.TotalAmount = userWallets.TotalAmount
			u.AvaAmount = userWallets.AvaAmount
			u.FronzenAmount = userWallets.FronzenAmount
			u.IsLock = userWallets.IsLock
		}
	}

	return &u
}

// 敏感数据处理
func PrivateDataHandler(v *vo.FcPrivate) {
	emailLen := len(v.Email)
	if emailLen > 4 {
		v.Email = v.Email[0:3] + "***" + v.Email[emailLen-4:]
	}

	telLen := len(v.Tel)
	if telLen > 4 {
		v.Tel = v.Tel[0:3] + "****" + v.Tel[telLen-4:]
	}

	alipayLen := len(v.Alipay)
	if alipayLen > 4 {
		v.Alipay = v.Alipay[0:3] + "****" + v.Alipay[alipayLen-4:]
	}

	aliPayRealNameRune := []rune(v.AlipayRealname)
	alipayRealNameLen := len(aliPayRealNameRune)
	if alipayRealNameLen > 2 {
		v.AlipayRealname = string(aliPayRealNameRune[0:1]) + "**"
	} else if alipayRealNameLen > 1 {
		v.AlipayRealname = string(aliPayRealNameRune[0:1]) + "*"
	}

	// 中文字符需要特殊处理
	realNameRune := []rune(v.RealName)
	realNameLen := len(realNameRune)
	if realNameLen > 2 {
		v.RealName = string(realNameRune[0:1]) + "**"
	} else if realNameLen > 1 {
		v.RealName = string(realNameRune[0:1]) + "*"
	}

	if v.WalletPassword != "" {
		v.WalletPassword = "已设置"
	}

	// 银行卡号
	accountNumLen := len(v.AccountNumber)
	if accountNumLen > 4 {
		v.AccountNumber = v.AccountNumber[0:3] + "*****" + v.AccountNumber[accountNumLen-4:]
	}

	// 银行卡号持有人
	accountHolderLen := len(v.AccountHolder)
	if accountHolderLen > 1 {
		tmpStr := ""
		for i := 0; i < accountHolderLen-1; i++ {
			tmpStr += "*"
		}
		v.AccountHolder = v.AccountHolder[0:1] + tmpStr
	}
}

// 用户查重
func FindPageUserRepeatControl(c *gin.Context) {
	jsonp := struct {
		dos.FcUserMaterial
		response.PageTimeQuery
	}{}

	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.UserName = c.DefaultQuery("user_name", "")
	jsonp.RealName = c.DefaultQuery("real_name", "")
	jsonp.Tel = c.DefaultQuery("tel", "")
	jsonp.Email = c.DefaultQuery("email", "")
	jsonp.Qq = c.DefaultQuery("qq", "")
	jsonp.Wx = c.DefaultQuery("wx", "")
	jsonp.Address = c.DefaultQuery("address", "")
	jsonp.RegisterIp = c.DefaultQuery("register_ip", "")
	jsonp.LastLoginIp = c.DefaultQuery("last_login_ip", "")
	jsonp.RegistVisitorId = c.DefaultQuery("regist_visitor_id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")

	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 对电话号码,真实姓名等进行加密处理
	jsonp.Encrypt()
	//global.G_LOG.Infof("username: %s orgTel: %s encrypt tel: %s", jsonp.UserName, tel, jsonp.Tel)

	data, total := modules.FindPageFcUserMaterial(jsonp.PageNo,
		jsonp.PageSize, &jsonp.FcUserMaterial, jsonp.PageTimeQuery, true, c)

	// 处理用户数据
	_, userInfoArr := HandleUserInfo(data, false, true)

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, userInfoArr)
}
