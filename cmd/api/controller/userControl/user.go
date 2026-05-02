package userControl

import (
	"bootpkg/cmd/api/controller/versionControl"
	"bootpkg/cmd/api/model/response"
	"bootpkg/cmd/api/model/srv"
	"bootpkg/cmd/api/model/vo"
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/langs"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/kirinlabs/utils"
	"github.com/kirinlabs/utils/convert"
)

const (
	inviteCodeTest = "8888888888"
)

// 获取落地页商户
func FindMerchantCode(c *gin.Context) {
	inviteCode := c.DefaultQuery("invite_code", "")
	url := c.DefaultQuery("url", "")

	var respData struct {
		MerchantCode string `json:"merchant_code"`
		CustomerLink string `json:"customer_link"`
		InviteCode   int    `json:"invite_code"`
	}

	if len(inviteCode) > 0 {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		// 优先根据推广ID查找商户code
		agentDomain := dos.FcAgentDomain{}
		err := global.G_DB.WithContext(ctx).Model(&dos.FcAgentDomain{}).Select("merchant_code", "customer_link", "invite_code").
			Where("invite_code = ?", inviteCode).First(&agentDomain).Error
		if err != nil {
			global.G_LOG.Errorf("[FindMerchantCode] Can not find merchant code by invite code: inviteCode=%s, err=%v",
				inviteCode, err.Error())
			response.FailErrJSON(c, response.ERROR_PARAMETER, "获取商户失败")
			return
		}

		respData.MerchantCode = agentDomain.MerchantCode
		respData.CustomerLink = agentDomain.CustomerLink
		respData.InviteCode = agentDomain.InviteCode
		response.SuccessJSON(c, respData)
		return
	} else if len(url) > 0 {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()
		// 其次该用域名地址查找商户code
		agentDomain := dos.FcAgentDomain{}
		err := global.G_DB.WithContext(ctx).Model(&dos.FcAgentDomain{}).Select("merchant_code", "customer_link", "invite_code").
			Where("jump_link like ?", "%"+strings.Trim(url, "/")+"%").First(&agentDomain).Error
		if err != nil {
			global.G_LOG.Errorf("[FindMerchantCode] Can not find merchant code by invite code: inviteCode=%s, err=%v",
				inviteCode, err.Error())
			response.FailErrJSON(c, response.ERROR_PARAMETER, "获取商户失败")
			return
		}

		respData.MerchantCode = agentDomain.MerchantCode
		respData.CustomerLink = agentDomain.CustomerLink
		respData.InviteCode = agentDomain.InviteCode
		response.SuccessJSON(c, respData)
		return
	}

	response.FailErrJSON(c, response.ERROR_PARAMETER, "请求参数错误")
}

func Register(c *gin.Context) {
	var jsonp vo.RegisterReq
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

	// 判断黑名单
	if modules.IsBlacklisted(tool.ClientIP(c), jsonp.VisitorId) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "已被加入黑名单")
		return
	}

	if !tool.IsValidUsername(jsonp.UserName) {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "用户名格式不正确")
		return
	}

	userNameFlag, err := tool.IsNumStr(jsonp.UserName)
	if err != nil || !userNameFlag {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "用户名不符合要求")
		return
	}

	if jsonp.RealName == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "用户姓名不能为空")
		return
	}

	if jsonp.Phone == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "手机号码不能为空")
		return
	}
	_, err = strconv.ParseInt(jsonp.Phone, 10, 64)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "手机号码错误")
		return
	}
	if jsonp.Password != jsonp.ConfirmPassword {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "密码不一致")
		return
	}
	pwdFlag, err := tool.IsNumStr(jsonp.Password)
	if err != nil || !pwdFlag {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "密码不合规：请输入7-14位数字和字母的组合!")
		return
	}

	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	if len(merchantCode) == 0 {
		err1 := global.G_DB.Model(&dos.FcAgentDomain{}).Select("merchant_code").
			Where("invite_code = ?", jsonp.InviteCode).Scan(&merchantCode).Error
		if err1 != nil {
			global.G_LOG.Errorf("[Register] Find agent domain data by invite code failed: inviteCode=%d, userName=%s, visitorId=%s, err=%s",
				jsonp.InviteCode, jsonp.UserName, jsonp.VisitorId, err1.Error())
		}
	}

	if len(merchantCode) == 0 {
		global.G_LOG.Errorf("[Register] Get user merchant code by invite code failed: inviteCode=%d, userName=%s, visitorId=%s",
			jsonp.InviteCode, jsonp.UserName, jsonp.VisitorId)
	}

	userNameM := fmt.Sprintf("%s:%s", merchantCode, jsonp.UserName)

	language := c.GetHeader(enmus.LANGUAGE_HEADER)
	if len(language) == 0 {
		language = "zh-CN"
	}

	tmpUser := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserName:     jsonp.UserName,
		MerchantCode: merchantCode,
	})
	if len(tmpUser.UserId) > 0 {
		response.FailErrJSON(c, response.Member_Have_Exsit, "用户名已存在")
		return
	}

	// 查询商户信息
	merchant := modules.FindByKeyFcMerchantFirst(&dos.FcMerchant{
		MerchantCode: merchantCode,
		Status:       1,
	})

	global.G_LOG.Infof("[Register] Looking for merchant with merchantCode: %s", merchantCode)

	if len(merchant.Id) == 0 {
		response.FailErrJSON(c, response.Merchant_Not_Exsit, "商户不存在")
		return
	}
	if merchant.Status != 1 {
		response.FailErrJSON(c, response.Merchant_Not_Exsit, "商户不可用")
		return
	}

	ipKey := fmt.Sprintf("%v:%v", merchantCode, tool.ClientIP(c))
	global.G_REDIS.Expire(context.Background(), ipKey, tool.TimeTomorrowTime())

	newUserId, err := modules.GetNextIdGeneral(modules.GKEY_USER_ID_INCR)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}
	userIdStr := strconv.FormatInt(newUserId, 10)
	hasher := tool.GetGlobalPasswordHasher()
	password, err := hasher.HashPassword(jsonp.Password)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, "密码加密失败")
		return
	}
	userLogin := &dos.FcUserLogin{
		BaseDos: dos.BaseDos{
			Id: userIdStr,
		},
		UserName:     jsonp.UserName,
		Password:     password,
		MerchantCode: merchantCode,
	}
	saveOk, _ := modules.SaveFcUserLogin(userLogin)
	if !saveOk {
		response.FailErrJSON(c, response.ERROR_SERVER, "")
		return
	}

	agentInviteCode := merchant.AgentInviteCode
	IsOfficialAgent := 1
	if jsonp.InviteCode != 0 {
		agentInviteCode = jsonp.InviteCode
		IsOfficialAgent = 2
	}

	agentId := ""
	if jsonp.InviteCode > 0 {
		agent := modules.FindByKeyFcAgentFirst(&dos.FcAgent{
			InviteCode: jsonp.InviteCode,
		})
		if len(agent.Id) > 0 {
			agentId = agent.Id
		}
	}

	var vip *dos.FcVip
	global.G_DB.Model(&dos.FcVip{}).Order("level asc").First(&vip)

	material := &dos.FcUserMaterial{
		UserId:          userIdStr,
		UserName:        jsonp.UserName,
		Sex:             1,
		Tel:             jsonp.Phone,
		RealName:        jsonp.RealName,
		NickName:        jsonp.UserName,
		MerchantCode:    merchantCode,
		Vip:             vip.VipName,
		Level:           1,
		IsWithdraw:      1,
		IsBonus:         1,
		AgentId:         agentId,
		Language:        language,
		RegisterIp:      tool.ClientIP(c),
		LastLoginIp:     tool.ClientIP(c),
		LastLoginTime:   automaticType.Time(time.Now()),
		InviteCode:      userIdStr,
		LoginStatus:     0,
		LastLoginCount:  0,
		CreateTime:      automaticType.Time(time.Now()),
		Website:         c.GetHeader(enmus.CLIENT_TYPE_HEADER),
		RegistVisitorId: jsonp.VisitorId,
		VisitorId:       jsonp.VisitorId,
		AgentInviteCode: agentInviteCode,
		IsOfficialAgent: IsOfficialAgent,
	}
	material.Encrypt()

	rFlag, _ := modules.SaveFcUserMaterial(material)
	if !rFlag {
		response.FailErrJSON(c, response.ERROR_SERVER, "注册失败")
		return
	}

	clientType := c.GetHeader(vo.Client_Type)
	modules.SaveFcLoginLog(&dos.FcLoginLog{
		UserId:       material.UserId,
		UserName:     material.UserName,
		Ip:           tool.ClientIP(c),
		ClientType:   clientType,
		MerchantCode: material.MerchantCode,
		Citys:        "",
	})

	jwtToken := srv.GenTokenJWTToken(userNameM)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	global.G_REDIS.Set(context.Background(), memberRedisKey, utils.Json(material), 16*24*time.Hour)

	modules.RecordAbnormalLogonData(tool.ClientIP(c), jsonp.VisitorId)

	var result struct {
		vo.RegisterResp
		MerchantCode string `json:"merchant_code"`
		InviteCode   int    `json:"invite_code"`
		Website      string `json:"website"`
	}

	result.RegisterResp.UserName = jsonp.UserName
	result.RegisterResp.Token = jwtToken
	result.RegisterResp.Level = material.Level
	result.MerchantCode = merchantCode
	result.InviteCode = jsonp.InviteCode
	result.Website = c.GetHeader(enmus.CLIENT_TYPE_HEADER)

	go modules.DoUserSystemEmail(material, enmus.MailType_FirstLogin)

	response.SuccessMsgJSON(c, result, "Success")
}

func Login(c *gin.Context) {
	var jsonp vo.LoginReq
	err := c.ShouldBind(&jsonp)

	// 判断黑名单
	if modules.IsBlacklisted(tool.ClientIP(c), jsonp.VisitorId) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "已被加入黑名单")
		return
	}

	versionRes := versionControl.LoginCheck(jsonp.UserName)
	if !versionRes {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "服务器维护中")
		return
	}

	err = global.VALIDATE.Struct(jsonp)
	if err != nil {
		var errstr []string
		for _, err := range err.(validator.ValidationErrors) {
			errstr = append(errstr, err.Translate(global.LANG))
		}
		response.FailErrJSON(c, response.ERROR_PARAMETER, strings.Join(errstr, "\n"))
		return
	}
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)
	userNameM := fmt.Sprintf("%s:%s", merchantCode, jsonp.UserName)

	val := global.G_REDIS.Get(context.Background(), enmus.MEMBER_REDIS_LOGIN_ERR_COUNT+userNameM).Val()
	count := convert.Atoi(val) + 1
	if count > 5 {
		response.FailErrJSON(c, ecode.PASSWORD_ERR_NUM, langs.GetWithLocaleGin(c, "message_13"))
		return
	}

	user := modules.FindByKeyFcUserLoginFirst(&dos.FcUserLogin{
		UserName:     jsonp.UserName,
		MerchantCode: merchantCode,
	})

	if user.Id == "" {
		response.FailErrJSON(c, ecode.ACCOUNT_NOT_REGISTER, langs.GetWithLocaleGin(c, "message_11"))
		return
	}

	hasher := tool.GetGlobalPasswordHasher()
	match, needUpgrade := hasher.VerifyPassword(jsonp.Password, user.Password)
	if !match {
		global.G_REDIS.Set(context.Background(), enmus.MEMBER_REDIS_LOGIN_ERR_COUNT+userNameM, count, 30*time.Minute)
		response.FailErrJSON(c, ecode.PASSWORD_FAIL, langs.GetWithLocaleGin(c, "message_20"))
		return
	}

	if needUpgrade {
		if newHash, upErr := hasher.UpgradePasswordToBcrypt(user.Password, jsonp.Password); upErr == nil && newHash != "" {
			user.Password = newHash
			if dbErr := global.G_DB.Model(&dos.FcUserLogin{}).Where("id = ?", user.Id).Update("password", newHash).Error; dbErr != nil {
				global.G_LOG.Warnf("[Login] upgrade password failed for userId=%s: %v", user.Id, dbErr)
			}
		}
	}

	material := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId:       user.Id,
		MerchantCode: merchantCode,
	})

	if material.LoginStatus == 1 {
		response.FailErrJSON(c, ecode.ACCOUNT_DISABLE, "账户已禁用，请联系客服")
		return
	}

	// 判断异常登录（本地开发临时注释）
	if false {
		response.FailErrDataJSON(c, ecode.LOGON_ABNORMAL, "登录异常", struct {
			LoginIp        string `json:"login_ip"`
			LoginVisitorId string `json:"login_visitor_id"`
		}{
			LoginIp:        tool.ClientIP(c),
			LoginVisitorId: jsonp.VisitorId,
		})
		return
	}

	go modules.DoUserMailAction(material)

	modules.UpdateFcUserMaterialLogin(&dos.FcUserMaterial{
		UserId:         user.Id,
		LastLoginIp:    tool.ClientIP(c),
		LastLoginTime:  automaticType.Time(time.Now()),
		LastLoginCount: material.LastLoginCount + 1,
	})

	clientType := c.GetHeader(vo.Client_Type)
	modules.SaveFcLoginLog(&dos.FcLoginLog{
		UserId:       user.Id,
		UserName:     user.UserName,
		Ip:           tool.ClientIP(c),
		ClientType:   clientType,
		MerchantCode: user.MerchantCode,
		VisitorId:    material.VisitorId,
		CreateTime:   automaticType.Time(time.Now()),
	})

	jwtToken := srv.GenTokenJWTToken(userNameM)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	global.G_REDIS.Set(context.Background(), memberRedisKey, utils.Json(material), 16*24*time.Hour)

	global.G_REDIS.Del(context.Background(), enmus.MEMBER_REDIS_LOGIN_ERR_COUNT+userNameM)

	// 返回登录成功信息
	response.SuccessJSON(c, map[string]interface{}{
		"username": jsonp.UserName,
		"token":    jwtToken,
		"level":    material.Level,
	})
}

// 异常登录恢复
func LoginAbnormalRecover(c *gin.Context) {
	var jsonp vo.LoginAbnormalRecoverReq
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)

	user := modules.FindByKeyFcUserLoginFirst(&dos.FcUserLogin{
		UserName:     jsonp.UserName,
		MerchantCode: merchantCode,
	})

	if user.Id == "" {
		response.FailErrJSON(c, ecode.ACCOUNT_NOT_REGISTER, langs.GetWithLocaleGin(c, "message_11"))
		return
	}

	hasher := tool.GetGlobalPasswordHasher()
	match, needUpgrade := hasher.VerifyPassword(jsonp.Password, user.Password)
	if !match {
		response.FailErrJSON(c, ecode.PASSWORD_FAIL, langs.GetWithLocaleGin(c, "message_20"))
		return
	}

	if needUpgrade {
		if newHash, upErr := hasher.UpgradePasswordToBcrypt(user.Password, jsonp.Password); upErr == nil && newHash != "" {
			user.Password = newHash
			if dbErr := global.G_DB.Model(&dos.FcUserLogin{}).Where("id = ?", user.Id).Update("password", newHash).Error; dbErr != nil {
				global.G_LOG.Warnf("[LoginAbnormalRecover] upgrade password failed for userId=%s: %v", user.Id, dbErr)
			}
		}
	}

	material := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId:       user.Id,
		MerchantCode: merchantCode,
	})

	if material.LoginStatus == 1 {
		response.FailErrJSON(c, ecode.ACCOUNT_DISABLE, langs.GetWithLocaleGin(c, "message_21"))
		return
	}

	// 更新最后登录IP和设备码
	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", material.UserId).Updates(map[string]interface{}{
		"last_login_ip": tool.ClientIP(c),
		"visitor_id":    jsonp.VisitorId,
	}).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, "更新用户信息失败")
	}

	// 执行正常登录流程
	userNameM := fmt.Sprintf("%s:%s", merchantCode, material.UserName)

	go modules.DoUserMailAction(material)

	modules.UpdateFcUserMaterialLogin(&dos.FcUserMaterial{
		UserId:         user.Id,
		LastLoginIp:    tool.ClientIP(c),
		LastLoginTime:  automaticType.Time(time.Now()),
		LastLoginCount: material.LastLoginCount + 1,
	})

	clientType := c.GetHeader(vo.Client_Type)
	modules.SaveFcLoginLog(&dos.FcLoginLog{
		UserId:       user.Id,
		UserName:     user.UserName,
		Ip:           tool.ClientIP(c),
		ClientType:   clientType,
		MerchantCode: user.MerchantCode,
		VisitorId:    material.VisitorId,
	})

	jwtToken := srv.GenTokenJWTToken(userNameM)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	global.G_REDIS.Set(context.Background(), memberRedisKey, utils.Json(material), 16*24*time.Hour)

	modules.RecordAbnormalLogonData(tool.ClientIP(c), jsonp.VisitorId)

	// 返回登录成功信息
	response.SuccessJSON(c, map[string]interface{}{
		"username": material.UserName,
		"token":    jwtToken,
		"level":    material.Level,
	})
}

// 用户登录更新站内信信息
func UserLoginUpdateSiteMsg(u *dos.FcUserMaterial) error {
	// 获取近 7 天内的全局站内信消息
	siteMsgArr := []*dos.FcSiteMessage{}
	startAt := time.Now().Add(-7 * 24 * time.Hour)
	registerAt := time.Time(u.CreateTime)
	if registerAt.After(startAt) {
		startAt = registerAt
	}
	startAtStr := startAt.Format("2006-01-02 15:04:05")

	err := global.G_DB.Model(&dos.FcSiteMessage{}).Where("notify_type = 1 AND create_time >= ?", startAtStr).Find(&siteMsgArr).Error
	if err != nil {
		global.G_LOG.Errorf("query FcSiteMessage by create_time >= %s err: %v", startAt, err)
		return err
	}

	nowTime := automaticType.Now()
	for _, v := range siteMsgArr {
		userMsgKey := fmt.Sprintf(enmus.UserSiteMessageKey, v.Id, u.UserId)
		tmpResult := global.G_REDIS.Get(context.Background(), userMsgKey).Val()
		if tmpResult != "" {
			continue
		}

		userIdName := vo.UserIdNameVO{}
		global.G_DB_SHARDING.Model(&dos.FcUserSiteMessage{}).Where("user_id = ? AND msg_id = ?", u.UserId, v.Id).Scan(&userIdName)

		if userIdName.UserId != "" {
			continue
		}

		userSiteMsg := &dos.FcUserSiteMessage{}
		userSiteMsg.UserId = u.UserId
		userSiteMsg.UserName = u.UserName
		userSiteMsg.MerchantCode = u.MerchantCode

		userSiteMsg.MsgId = v.Id
		userSiteMsg.Title = v.Title
		userSiteMsg.Content = v.Content
		userSiteMsg.MsgIdType = v.MsgIdType
		userSiteMsg.MsgType = v.MsgType
		userSiteMsg.NotifyType = v.NotifyType
		userSiteMsg.Language = v.Language
		userSiteMsg.CreateBy = v.CreateBy
		userSiteMsg.UpdateBy = v.UpdateBy

		userSiteMsg.DelStatus = 1
		userSiteMsg.ReadStatus = 1
		userSiteMsg.UpdateTime = nowTime
		userSiteMsg.CreateTime = v.CreateTime

		err = global.G_DB.Create(userSiteMsg).Error
		if err != nil {
			global.G_LOG.Errorf("AllUserSendMsg Insert userSiteMsg err: %v", err)
			continue
		}

		global.G_REDIS.SetEx(context.Background(), userMsgKey, 1, 8*24*time.Hour)
	}

	return nil
}

// 推广页统计
func InviteSave(c *gin.Context) {
	var jsonp vo.FcUserInviteReq
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	visitData := &dos.FcVisit{}
	visitData.UserId = jsonp.UserId
	visitData.InviteCode = jsonp.InviteCode
	err = global.G_DB.Create(visitData).Error
	if err != nil {
		global.G_LOG.Errorf("User Invite Insert err: %v", err)
		response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
		return
	}

	response.SuccessMsgJSON(c, nil, "success")
}

// 推广页跳转链接
func InviteLink(c *gin.Context) {
	var jsonp vo.FcInviteDomainReq
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	var data dos.FcAgentDomain
	err = global.G_DB.Model(&dos.FcAgentDomain{}).Select("invite_code", "domain", "jump_link",
		"merchant_code", "type", "customer_link", "ios_link", "android_link", "banner_img", "logo_img").
		Where("invite_code = ?", jsonp.InviteCode).Scan(&data).Error
	if len(data.CustomerLink) == 0 {
		data.CustomerLink = modules.SyncFcAgentDomainCustomerLink(data.MerchantCode, true)
	}
	merchantLink := modules.GetFcMerchantLinkData(data.MerchantCode)
	data.BannerImg = merchantLink.BannerImg
	var result vo.FcInviteDomainResp
	tool.JsonMapper(data, &result)
	if err != nil || result.InviteCode == 0 {
		global.G_LOG.Error(err)
		response.FailErrJSON(c, response.ERROR_SERVER, "不存在的推广码")
		return
	}

	agentDomain := modules.FindByKeyFcAgentDomainFirst(&dos.FcAgentDomain{
		MerchantCode: data.MerchantCode,
		Type:         enmus.AgentDomainType_Agent,
	})
	result.IosLink = agentDomain.IosLink
	result.IosLink2 = agentDomain.IosLink2
	result.AndroidLink = agentDomain.AndroidLink
	result.AndroidLink2 = agentDomain.AndroidLink2

	response.SuccessMsgJSON(c, result, "success")
}

// 客服链接
func CustomerLink(c *gin.Context) {
	merchantCode := c.GetHeader(vo.MerchantCode_KEY_G)

	// 若 merchantCode 为空，尝试从 MerchantUrl 或 MerchantId 反查
	if len(merchantCode) == 0 {
		merchantUrl := c.GetHeader(vo.MerchantUrl_KEY_G)
		merchantId := c.GetHeader(vo.MerchantID_KEY_G)
		if len(merchantUrl) > 0 {
			// 去掉协议头，只保留 host:port
			cleanUrl := merchantUrl
			if idx := strings.Index(cleanUrl, "://"); idx >= 0 {
				cleanUrl = cleanUrl[idx+3:]
			}
			var domainRow dos.FcAgentDomain
			if e := global.G_DB.Model(&dos.FcAgentDomain{}).Select("merchant_code").
				Where("domain LIKE ?", "%"+strings.Trim(cleanUrl, "/")+"%").
				Order("update_time DESC").Limit(1).Find(&domainRow).Error; e == nil && len(domainRow.MerchantCode) > 0 {
				merchantCode = domainRow.MerchantCode
			}
		}
		if len(merchantCode) == 0 && len(merchantId) > 0 {
			var domainRow dos.FcAgentDomain
			if e := global.G_DB.Model(&dos.FcAgentDomain{}).Select("merchant_code").
				Where("invite_code = ?", merchantId).
				Limit(1).Find(&domainRow).Error; e == nil && len(domainRow.MerchantCode) > 0 {
				merchantCode = domainRow.MerchantCode
			}
		}
	}

	data := dos.FcCustomerLink{}
	err := global.G_DB.Model(&dos.FcCustomerLink{}).Select("link").
		Where("merchant_code = ?", merchantCode).
		First(&data).Error
	if err != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, "无法获取客服链接")
		return
	}

	merchantLink := modules.GetFcMerchantLinkData(merchantCode)
	agentDomain := &dos.FcAgentDomain{}
	// Prefer the latest enabled official-domain row so UI/sign models are deterministic.
	err = global.G_DB.Model(&dos.FcAgentDomain{}).
		Where("merchant_code = ? AND `type` = ? AND status = ?", merchantCode, enmus.AgentDomainType_Agent, 1).
		Order("update_time DESC, id DESC").
		Take(agentDomain).Error
	if err != nil || len(agentDomain.MerchantCode) == 0 {
		// Fallback for historical data without enabled rows.
		err = global.G_DB.Model(&dos.FcAgentDomain{}).
			Where("merchant_code = ? AND `type` = ?", merchantCode, enmus.AgentDomainType_Agent).
			Order("update_time DESC, id DESC").
			Take(agentDomain).Error
	}

	// 准备 UI 资产配置
	var uiAssetsJson interface{} = nil
	if len(merchantLink.HomeEntryIcons) > 0 {
		// 直接使用存储的 JSON 字符串
		uiAssetsJson = merchantLink.HomeEntryIcons
	} else {
		// 如果没有配置，返回空对象
		uiAssetsJson = "{}"
	}

	response.SuccessJSON(c, struct {
		MerchantCode string `json:"merchant_code"`
		MerchantName string `json:"merchant_name"`
		Link         string `json:"link"`
		ShortLink    string `json:"short_link"`
		LogoImg      string `json:"logo_img"`
		BannerImg    string `json:"banner_img"`
		UiAssets     string `json:"ui_assets"`
	}{
		MerchantCode: merchantCode,
		MerchantName: modules.GetMerchantName(merchantCode),
		Link:         data.Link,
		ShortLink:    agentDomain.ShortLink,
		LogoImg:      merchantLink.LogoImg,
		BannerImg:    merchantLink.BannerImg,
		UiAssets:     uiAssetsJson.(string),
	})
}

// 文本导航信息
func GuideInfo(c *gin.Context) {
	var jsonp struct {
		Id string `form:"id"`
	}

	if err := c.ShouldBindQuery(&jsonp); err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	guides := modules.GetGuideInfo(jsonp.Id)
	response.SuccessJSON(c, guides)
}

// 广告栏信息
func AdsCarouselInfo(c *gin.Context) {
	var jsonp struct {
		Id string `form:"id"`
	}

	if err := c.ShouldBindQuery(&jsonp); err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	data := modules.GetAdsCarouselInfo(jsonp.Id)
	response.SuccessJSON(c, data)
}

// 获取货币汇率
func GetCurrencyFx(c *gin.Context) {
	var result struct {
		Message   string  `json:"message"`
		RatePrice float64 `json:"ratePrice"`
		Status    string  `json:"status"`
	}

	_, err := resty.New().R().
		SetHeader("Accept", "application/json").
		SetResult(&result).
		Get("http://pay2.fulian.co/api/getRate?rateType=CNY")
	if err != nil {
		global.G_LOG.Errorf("[GetCurrencyFx] get currency fx failed: err=%s", err.Error())
		response.FailErrJSON(c, response.ERROR_SERVER, "获取汇率失败稍后再试")
		return
	}

	response.SuccessJSON(c, struct {
		RatePrice float64 `json:"ratePrice"`
	}{
		RatePrice: result.RatePrice,
	})
}

func TempClientLogs(c *gin.Context) {
	var jsonp struct {
		ID   string `json:"ID" form:"ID" uri:"ID"`
		Data string `json:"data" form:"data" uri:"data"`
	}
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
		return
	}

	global.G_LOG.Infof("[ClientLogs] Id=%s, data=%v", jsonp.ID, jsonp.Data)

	response.SuccessJSON(c, true)
}

func GetClientIpInfo(c *gin.Context) {
	var result struct {
		Status      string  `json:"status"`
		Country     string  `json:"country"`
		CountryCode string  `json:"countryCode"`
		Region      string  `json:"region"`
		RegionName  string  `json:"regionName"`
		City        string  `json:"city"`
		Zip         string  `json:"zip"`
		Lat         float64 `json:"lat"`
		Lon         float64 `json:"lon"`
		Timezone    string  `json:"timezone"`
		Isp         string  `json:"isp"`
		Org         string  `json:"org"`
		As          string  `json:"as"`
		Query       string  `json:"query"`
	}

	clientIp := tool.ClientIP(c)

	if clientIp == "127.0.0.1" || clientIp == "localhost" {
		response.FailErrJSON(c, response.ERROR_SERVER, "获取IP信息失败稍后再试")
		return
	}

	_, err := resty.New().R().
		SetHeader("Accept", "application/json").
		SetResult(&result).
		Get(fmt.Sprintf("http://ip-api.com/json/%s?lang=zh-CN", clientIp))
	if err != nil {
		global.G_LOG.Errorf("[GetClientIpInfo] get ip information failed: ip=%s, err=%s",
			tool.ClientIP(c), err.Error())
		response.FailErrJSON(c, response.ERROR_SERVER, "获取IP信息失败稍后再试")
		return
	}

	if result.Status != "success" {
		global.G_LOG.Errorf("[GetClientIpInfo] get ip information failed: ip=%s, result=%+v",
			tool.ClientIP(c), result)
		response.FailErrJSON(c, response.ERROR_SERVER, "获取IP信息失败稍后再试")
		return
	}

	response.SuccessJSON(c, result)
}
