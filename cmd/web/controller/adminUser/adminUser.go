// The build tag makes sure the stub is not built in the final build.

package adminUser

import (
	"bootpkg/cmd/web/handler"
	vo2 "bootpkg/cmd/web/model/vo"
	"bootpkg/cmd/web/srv"
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/core/modules/vo"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/kirinlabs/utils"
	"github.com/kirinlabs/utils/convert"
	"github.com/kirinlabs/utils/str"

	//"github.com/wenlng/go-captcha-assets/resources/images"
	//"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/captcha"
	"github.com/wenlng/go-captcha/v2/slide"

	//"log"
	"strconv"
	"strings"
	"time"
)

var slideBasicCapt slide.Captcha

//func init() {
//	slideBasicCapt = slide.New(
//		//slide.WithGenGraphNumber(2),
//		slide.WithEnableGraphVerticalRandom(true),
//	)
//
//	// background images
//	imgs, err := images.GetImages()
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	graphs, err := tiles.GetTiles()
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
//	for i := 0; i < len(graphs); i++ {
//		graph := graphs[i]
//		newGraphs = append(newGraphs, &slide.GraphImage{
//			OverlayImage: graph.OverlayImage,
//			MaskImage:    graph.MaskImage,
//			ShadowImage:  graph.ShadowImage,
//		})
//	}
//
//	// set resources
//	slideBasicCapt.SetResources(
//		slide.WithGraphImages(newGraphs),
//		slide.WithBackgrounds(imgs),
//	)
//}

// api: api/adminUser/save
func SaveAdminUserControl(c *gin.Context) {
	var jsonp vo2.AdminUserAddReq
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
	if !tool.IsValidUsername2(jsonp.UserName) {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "用户名非法，不能包含汉字或特殊字符")
		return
	}

	if jsonp.LimitPertimeAmount > jsonp.TotalAmount {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "单词限制不能比总限制额度大")
		return
	}
	var adminUser dos.AdminUser

	tool.JsonMapper(jsonp, &adminUser)

	userInfo, ok := c.Get("UserInfo")
	userInfoM := &dos.AdminUser{}
	if ok {
		userInfoM = userInfo.(*dos.AdminUser)
		adminUser.CreateBy = userInfoM.UserName

		// 普通账号不能创建账号
		if userInfoM.AccountType == 3 {
			response.FailErrJSON(c, ecode.Unauthorized, "无权限")
			return
		}
	}
	// 只有超管账号才能授权给超管
	if adminUser.AccountType == 1 && userInfoM.AccountType != 1 {
		response.FailErrJSON(c, ecode.Unauthorized, "无权限")
		return
	}
	// 普通账号不能赋予管理员权限
	if adminUser.AccountType == 2 && (userInfoM.AccountType != 1 && userInfoM.AccountType != 2) {
		response.FailErrJSON(c, ecode.Unauthorized, "无权限")
		return
	}

	existUser := modules.FindByKeyAdminUser(&dos.AdminUser{UserName: jsonp.UserName})
	if len(existUser) > 0 {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "用户名已经存在")
		return
	}

	adminUser.CreateTime = automaticType.Time(time.Now())

	adminUser.MerchantCodes = strings.Trim(adminUser.MerchantCodes, ",")

	hasher := tool.GetGlobalPasswordHasher()
	hashPwd, hashErr := hasher.HashPassword(jsonp.Pwd)
	if hashErr != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, "密码加密失败")
		return
	}
	adminUser.Pwd = hashPwd
	adminUser.EnforcePwd = 1 //强制修改密码
	data, _ := modules.SaveAdminUser(&adminUser)
	response.SuccessJSON(c, data)
}

// api: api/adminUser/findPage
func FindPageAdminUserControl(c *gin.Context) {
	jsonp := struct {
		response.PageQuery
		dos.AdminUser
	}{}
	err := c.ShouldBindQuery(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if jsonp.PageNo <= 0 {
		jsonp.PageNo = 1
	}
	if jsonp.PageSize <= 0 {
		jsonp.PageSize = 10
	}

	data, total := modules.FindPageAdminUser(jsonp.PageNo, jsonp.PageSize, &jsonp.AdminUser, c)

	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, data)
}

// api: api/adminUser/findByKey
func FindByKeyAdminUserControl(c *gin.Context) {
	var jsonp dos.AdminUser
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindByKeyAdminUserWithPerms(&jsonp, c)
	if len(data) > 0 {
		data[0].Pwd = ""
	}
	response.SuccessJSON(c, data)
}

// api: api/adminUser/update
func UpdateAdminUserControl(c *gin.Context) {
	var jsonp dos.AdminUser
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
	if !tool.IsValidUsername2(jsonp.UserName) {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "用户名非法，不能包含汉字或特殊字符")
		return
	}
	if jsonp.LimitPertimeAmount > jsonp.TotalAmount {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "单词限制不能比总限制额度大")
		return
	}
	jsonp.MerchantCodes = strings.Trim(jsonp.MerchantCodes, ",")

	userInfo, ok := c.Get("UserInfo")
	userInfoM := &dos.AdminUser{}
	if ok {
		userInfoM = userInfo.(*dos.AdminUser)
		jsonp.UpdateBy = userInfoM.UserName

		if userInfoM.AccountType != 1 && userInfoM.UserName == jsonp.UserName {
			response.FailErrJSON(c, ecode.Unauthorized, "请联系上级管理员")
			return
		}
	}
	// 只有超管账号才能授权给超管
	if jsonp.AccountType == 1 && userInfoM.AccountType != 1 {
		response.FailErrJSON(c, ecode.Unauthorized, "无权限")
		return
	}
	// 普通账号不能赋予管理员权限
	if jsonp.AccountType == 2 && (userInfoM.AccountType != 1 && userInfoM.AccountType != 2) {
		response.FailErrJSON(c, ecode.Unauthorized, "无权限")
		return
	}

	// 判断是否为子账户
	if !modules.CheckAdminUserSubAccount(c, jsonp.Id) {
		response.FailErrJSON(c, ecode.Unauthorized, "无权限，请联系管理员")
		return
	}

	data := modules.UpdateAdminUser(&jsonp)
	response.SuccessJSON(c, data)
}

// api: api/adminUser/delete
func DeleteAdminUserControl(c *gin.Context) {
	var jsonp dos.AdminUser
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	// 判断是否为子账户
	if !modules.CheckAdminUserSubAccount(c, jsonp.Id) {
		response.FailErrJSON(c, ecode.Unauthorized, "无权限，请联系管理员")
		return
	}

	data := modules.DeleteAdminUser(&jsonp)
	response.SuccessJSON(c, data)
}

// api: /api/base/nologin/login
func AdminUserLogin(c *gin.Context) {
	var jsonp vo.UserLoginVO
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	jsonBytes, _ := tool.JsonMarshal(&jsonp)
	if err != nil {
		global.G_LOG.Errorf("accept param=%v err: %v", string(jsonBytes), err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	//global.G_LOG.Infof("accept name=%v param=%v success", jsonp.UserName, string(jsonBytes))

	if !modules.IsWhiteIP(tool.ClientIP(c)) {
		response.FailErrJSON(c, ecode.NotWhiteIp, "非白名单IP")
		return
	}

	// 验证码参数校验, 是否开启谷歌验证
	/*codeNum := 0
	if modules.UseGoogleMfa() && len(jsonp.Key) > 0 {
		codeNum, err = strconv.Atoi(jsonp.Key)
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}
	}*/

	//验证码验证
	//cacheValue := global.G_REDIS.Get(context.Background(), enmus.LOGIN_CAPTCHA+jsonp.Key).Val()
	//if cacheValue != "true" {
	//	response.FailErrJSON(c, response.ERROR_PARAMETER, "人机验证已失效,请重新验证")
	//	return
	//}
	//defer global.G_REDIS.Del(context.Background(), enmus.LOGIN_CAPTCHA+jsonp.Key)

	// 获取用户登录的错误次数
	userLoginCountErrKey := enmus.REDIS_LOGIN_ERR_COUNT + jsonp.UserName
	var count int
	//先隐藏这部分逻辑
	/*
		val := global.G_REDIS.Get(context.Background(), userLoginCountErrKey).Val()
		count := convert.Atoi(val) + 1
		if count > 5 {
			response.FailJSON(c, ecode.PASSWORD_ERR_NUM)
			return
		}
	*/
	data := modules.FindByKeyAdminUser(&dos.AdminUser{UserName: jsonp.UserName})
	if len(data) == 0 {
		global.G_LOG.Errorf("query user=%s from db err: %v", jsonp.UserName, err)
		global.G_REDIS.Set(context.Background(), userLoginCountErrKey, count, 30*time.Minute)
		response.FailJSON(c, ecode.PASSWORD_FAIL)
		return
	}
	if data[0].Status == enmus.STATUS_DISABLE {
		response.FailJSON(c, ecode.ACCOUNT_IS_DISABLED)
		return
	}
	if len(data[0].RoleIds) == 0 {
		response.FailErrJSON(c, ecode.ServerErr, "未分配角色权限")
		return
	}
	hasher := tool.GetGlobalPasswordHasher()
	match, needUpgrade := hasher.VerifyPassword(jsonp.PassWord, data[0].Pwd)
	if !match {
		global.G_LOG.Errorf("user=%s pwd error", jsonp.UserName)
		global.G_REDIS.Set(context.Background(), enmus.REDIS_LOGIN_ERR_COUNT+jsonp.UserName, count, 30*time.Minute)
		response.FailJSON(c, ecode.PASSWORD_FAIL)
		return
	}

	if needUpgrade {
		if newHash, upErr := hasher.UpgradePasswordToBcrypt(data[0].Pwd, jsonp.PassWord); upErr == nil && newHash != "" {
			if dbErr := global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", data[0].Id).Update("pwd", newHash).Error; dbErr != nil {
				global.G_LOG.Warnf("[AdminUserLogin] upgrade password failed for user=%s: %v", jsonp.UserName, dbErr)
			} else {
				data[0].Pwd = newHash
			}
		}
	}

	mfaErrCount := global.G_REDIS.Get(context.Background(), enmus.REDIS_LOGIN_MFA+jsonp.UserName).Val()
	if tool.Int(mfaErrCount) >= 7 {
		f := global.G_REDIS.TTL(context.Background(), enmus.REDIS_LOGIN_MFA+jsonp.UserName).Val().Minutes()
		response.FailErrJSON(c, response.ERROR_PARAMETER, "虚拟6位验证码错误次数过多，账号将被临时锁定 "+tool.String(int(f))+" 分钟")
		return
	}

	// 判断用户的 google 验证码是否正确
	errCode := ecode.OK
	if modules.UseGoogleMfa() {
		if data[0].Mfa == "" {
			errCode = ecode.GoogleMFABind
		} else {
			errCode = ecode.GoogleMFAVerify
		}

		/*mfa, err := GetGoogleMfa(data[0].Mfa)
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}

		flag := googleAuthenticator.VerifyCode(mfa, int32(codeNum))
		if !flag {
			global.G_REDIS.Set(context.Background(), enmus.REDIS_LOGIN_ERR_COUNT+jsonp.UserName, count, 30*time.Minute)
			response.FailErrJSON(c, response.ERROR_PARAMETER, "谷歌验证码验证失败")
			return
		}*/
	}

	//单点
	var token = strings.Replace(uuid.NewString(), "-", "", -1)
	ttl := 2 * time.Hour
	if jsonp.AutoLogin {
		ttl = 8 * time.Hour
	}
	ip := c.ClientIP() + token
	token += "." + tool.MD5([]byte(ip))
	data[0].Token = token

	sessionId, err := c.Cookie(enmus.LOGIN_COOKIE)
	if err != nil {
		// Cookie 不存在，创建一个新的
		sessionId = handler.SetSessionCookie(c)
		global.G_LOG.Infof("user=%v create new cookie: %v", jsonp.UserName, sessionId)
	}

	// 同步商户保持最新
	data[0].MerchantCodes = modules.SyncMerchantCodes(jsonp.UserName, data[0].MerchantCodes)

	// 存储用户 token 等信息
	global.G_REDIS.Set(context.Background(), token, "1", ttl)
	global.G_REDIS.Set(context.Background(), enmus.REDIS_LOGIN_TOKEN+sessionId, jsonp.UserName, ttl)
	global.G_REDIS.Set(context.Background(), enmus.REDIS_LOGIN_USERINFO+jsonp.UserName, utils.Json(data[0]), ttl)
	modules.SaveLoginLog(&dos.LoginLog{
		UserName:   jsonp.UserName,
		Ip:         c.ClientIP(),
		CreateTime: automaticType.Time(time.Now()),
	})

	// 查询角色权限集合
	permsList := ""
	global.G_LOG.Infof("[AdminUserLogin] roleIds=%s, permsList=%s", data[0].RoleIds, permsList)
	if len(data[0].RoleIds) > 0 {
		roleIds := strings.Split(data[0].RoleIds, ",")
		if len(roleIds) > 0 {
			var role dos.Role
			result := global.G_DB.Model(&dos.Role{}).Select("perms_list").Where("id = ?", roleIds[0]).Scan(&role)
			global.G_LOG.Infof("[AdminUserLogin] SQL affected: %d, role.PermsList from DB: %s", result.RowsAffected, role.PermsList)
			permsList = role.PermsList
			global.G_LOG.Infof("[AdminUserLogin] after assignment, permsList=%s", permsList)
		}
	}

	merchantNames := modules.GetFcMerchantNamesStringByCodes(data[0].MerchantCodes)

	result := struct {
		Token         string
		Nickname      string
		MerchantCodes string
		MerchantNames string
		Id            string
		RoleIds       string
		PermsList     string
		AccountType   int    `json:"account_type"`
		Department    string `json:"department"`
		LoginIp       string `json:"login_ip"`
		Mobile        string `json:"mobile"`
	}{
		Token:         token,
		Nickname:      data[0].UserNick,
		MerchantCodes: data[0].MerchantCodes,
		MerchantNames: merchantNames,
		Id:            data[0].Id,
		RoleIds:       data[0].RoleIds,
		PermsList:     permsList,
		AccountType:   data[0].AccountType,
		Department:    modules.GetDepartmentName(data[0].DepartmentId),
		LoginIp:       tool.ClientIP(c),
		Mobile:        data[0].Mobile,
	}
	response.SuccessCodeJSON(c, errCode, result)
}

// api: /api/adminUser/isLogin
func AdminUserIsLogin(c *gin.Context) {
	response.SuccessJSON(c, nil)
}

func UpdateAdminUserByDepartmentIdControl(c *gin.Context) {
	var jsonp dos.AdminUser
	err := c.ShouldBindJSON(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	data := modules.UpdateAdminUserByDepartmentId(&jsonp)
	response.SuccessJSON(c, data)
}

func UpdateAdminUserByRoleIdsControl(c *gin.Context) {
	var jsonp dos.AdminUser
	err := c.ShouldBindJSON(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	data := modules.UpdateAdminUserByRoleIds(&jsonp)
	response.SuccessJSON(c, data)
}

func UpdateAdminUserPwd(c *gin.Context) {
	var jsonp dos.AdminUser
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
	pwd := jsonp.Pwd
	if pwd == "" {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "password is empty")
		return
	}
	hasher := tool.GetGlobalPasswordHasher()
	password, hashErr := hasher.HashPassword(pwd)
	if hashErr != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, "密码加密失败")
		return
	}

	token := c.Request.Header.Get(enmus.LOGIN_TOKEN)
	if len(token) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "token is empty")
		return
	}
	sessionId, err := c.Cookie(enmus.LOGIN_COOKIE)
	if err != nil {
		global.G_LOG.Errorf("user=%v get cookie=%v err: %v", jsonp.UserName, enmus.LOGIN_COOKIE, err)
		response.FailErrJSON(c, response.ERROR_PARAMETER, "无效的访问xxx")
		return
	}

	updateMap := make(map[string]interface{})
	updateMap["pwd"] = password
	updateMap["enforce_pwd"] = 0

	eRow := global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", jsonp.Id).Updates(updateMap)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}

	// 删除 token, 删除用户信息, 删除用户登录错误次数
	tokenKey := enmus.REDIS_LOGIN_TOKEN + sessionId
	userInfoKey := enmus.REDIS_LOGIN_USERINFO + jsonp.UserName
	userLoginCountErrKey := enmus.REDIS_LOGIN_ERR_COUNT + jsonp.UserName
	global.G_REDIS.Del(context.Background(), tokenKey).Val()
	global.G_REDIS.Del(context.Background(), userInfoKey).Val()
	global.G_REDIS.Del(context.Background(), userLoginCountErrKey).Val()

	response.SuccessJSON(c, true)
}

func AdminUserGetMenusByRole(c *gin.Context) {
	userInfo, ok := c.Get("UserInfo")
	if !ok {
		response.FailJSON(c, response.ERROR_PARAMETER)
		return
	}
	u := userInfo.(*dos.AdminUser)

	// 获取所有菜单
	menus := modules.FindMenusAll()

	// 超级管理员直接返回所有菜单
	if u.AccountType == 1 {
		menusList := make([]*dos.MenusResp, 0)
		var parentIds []string
		for _, v := range menus {
			if len(v.ParentId) > 0 {
				parentIds = append(parentIds, v.ParentId)
			}
		}

		queryParentMenus := []*dos.Menus{}
		global.G_DB.Model(&dos.Menus{}).Select("id", "menu_name", "icon", "role_flag").Where("id IN ?", parentIds).Find(&queryParentMenus)
		parentMenusMp := map[string]*dos.Menus{}
		for _, v := range queryParentMenus {
			parentMenusMp[v.Id] = v
		}

		for _, v := range menus {
			v.CreateBy = ""
			v.UpdateBy = ""
			v.ApiRegular = ""
			if v.MenuName == "菜单管理" {
				v.ShowStatus = 1
			}

			respMenu := dos.MenusResp{}
			tool.JsonMapper(v, &respMenu)
			if val, ok := parentMenusMp[v.ParentId]; ok {
				respMenu.ParentMenuName = val.MenuName
				respMenu.ParentIcon = val.Icon
				respMenu.ParentRoleFlag = val.RoleFlag
			}
			menusList = append(menusList, &respMenu)
		}
		response.SuccessJSON(c, menusList)
		return
	}

	// 非超级管理员：根据角色权限返回菜单
	roles := modules.FindRoleAll()
	var userMenusIdsBuild strings.Builder
	userRolesId := "," + u.RoleIds + ","
	for i := 0; i < len(roles); i++ {
		if str.Contians(userRolesId, ","+convert.String(roles[i].Id)+",") {
			if i > 0 {
				userMenusIdsBuild.WriteString(",")
			}
			userMenusIdsBuild.WriteString(roles[i].MeusIds)
		}
	}
	if userMenusIdsBuild.Len() == 0 {
		response.FailErrJSON(c, ecode.AccessDenied, "权限不足")
		return
	}

	menusList := make([]*dos.MenusResp, 0)
	var parentIds []string
	for _, v := range menus {
		if len(v.ParentId) > 0 {
			parentIds = append(parentIds, v.ParentId)
		}
	}

	queryParentMenus := []*dos.Menus{}
	global.G_DB.Model(&dos.Menus{}).Select("id", "menu_name", "icon", "role_flag").Where("id IN ?", parentIds).Find(&queryParentMenus)
	parentMenusMp := map[string]*dos.Menus{}
	for _, v := range queryParentMenus {
		parentMenusMp[v.Id] = v
	}

	// 把父级加入到userMenusIdsBuild中
	roleMenusParentMp := modules.GetMenusMapByRoleMenusIds(strings.Split(userMenusIdsBuild.String(), ","))
	for k := range roleMenusParentMp {
		userMenusIdsBuild.WriteString("," + k)
	}

	userMenusIds := "," + userMenusIdsBuild.String() + ","
	for _, v := range menus {
		v.CreateBy = ""
		v.UpdateBy = ""
		v.ApiRegular = ""
		if str.Contians(userMenusIds, ","+convert.String(v.Id)+",") {
			respMenu := dos.MenusResp{}
			tool.JsonMapper(v, &respMenu)
			if val, ok := parentMenusMp[v.ParentId]; ok {
				respMenu.ParentMenuName = val.MenuName
				respMenu.ParentIcon = val.Icon
				respMenu.ParentRoleFlag = val.RoleFlag
			}
			menusList = append(menusList, &respMenu)
		}
	}
	if len(menusList) == 0 {
		response.FailErrJSON(c, ecode.AccessDenied, "权限不足")
		return
	}
	response.SuccessJSON(c, menusList)
}

func FindAdminUserAllControl(c *gin.Context) {
	var jsonp dos.AdminUser
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	data := modules.FindUserNameByKeyAdminUser(&jsonp, c)
	response.SuccessJSON(c, data)
}

func LoginOut(c *gin.Context) {
	if !srv.ExitLogin(c) {
	}
	response.SuccessJSON(c, true)
}

func UpdateAdminUserPassWord(c *gin.Context) {
	var jsonp vo.AdminUserPasswordVO
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if jsonp.ConfirmPwd != jsonp.NewPwd || len(jsonp.ConfirmPwd) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "两次密码输入不一致")
		return
	}
	if jsonp.OldPwd == jsonp.ConfirmPwd {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "新密码不能和原密码一样")
		return
	}

	userInfoGIN, _ := c.Get("UserInfo")
	userInfo := userInfoGIN.(*dos.AdminUser)
	m := &dos.AdminUser{}
	m.Id = userInfo.Id
	data := modules.FirstByKeyAdminUser(m)
	if data == nil || len(data.Id) == 0 {
		response.FailJSON(c, ecode.FAIL)
		return
	}
	hasher := tool.GetGlobalPasswordHasher()
	match, _ := hasher.VerifyPassword(jsonp.OldPwd, data.Pwd)
	if !match {
		response.FailJSON(c, ecode.PASSWORD_FAIL)
		return
	}

	newPwdHash, hashErr := hasher.HashPassword(jsonp.ConfirmPwd)
	if hashErr != nil {
		response.FailErrJSON(c, response.ERROR_SERVER, "密码加密失败")
		return
	}
	u := &dos.AdminUser{
		Pwd:      newPwdHash,
		UpdateBy: userInfo.UserName,
	}
	u.Id = userInfo.Id
	d := modules.UpdateAdminUserPassword(u)
	if !srv.ExitLogin(c) {
	}
	response.SuccessJSON(c, d)
}

func GetCaptchaData(c *gin.Context) {
	capt := captcha.GetCaptcha()
	chars := []string{"AF", "BF", "CV", "DO", "EG", "FR", "GB", "HM", "JM", "KZ", "LU", "MC", "MC", "MO", "NR", "PM", "TB", "PM"}
	_ = capt.SetRangChars(chars)
	capt.SetRangCheckTextLen(captcha.RangeVal{
		Min: 2,
		Max: 2,
	})
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "请重新获取验证码")
		return
	}
	global.G_REDIS.Set(context.Background(), enmus.LOGIN_CAPTCHA+key, tool.String(dots), 5*time.Minute)
	bt := map[string]interface{}{
		"image_base64": b64,
		"thumb_base64": tb64,
		"captcha_key":  key,
	}
	response.SuccessJSON(c, bt)
}

func CheckCaptchaData(c *gin.Context) {
	var jsonp vo.CaptchaDataVO
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "参数提交错误")
		return
	}
	cacheData := global.G_REDIS.Get(context.Background(), enmus.LOGIN_CAPTCHA+jsonp.Key).Val()
	if len(cacheData) == 0 {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "验证失败")
		return
	}
	var dct map[int]captcha.CharDot
	if err := json.Unmarshal([]byte(cacheData), &dct); err != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "验证失败")
		return
	}

	src := jsonp.Dots
	chkRet := false
	if (len(dct) * 2) == len(src) {
		for i, dot := range dct {
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 76)
			sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 76)
			chkRet = captcha.CheckPointDistWithPadding(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height), 10)
			if !chkRet {
				break
			}
		}
	}
	if chkRet {
		global.G_REDIS.Set(context.Background(), enmus.LOGIN_CAPTCHA+jsonp.Key, "true", 5*time.Minute)
	} else {
		global.G_REDIS.Del(context.Background(), enmus.LOGIN_CAPTCHA+jsonp.Key)
	}
	response.SuccessJSON(c, chkRet)
}

/*
func GetCaptchaDataSlide(c *gin.Context) {
	captData, err := slideBasicCapt.Generate()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "Generate fail")
		return
	}
	blockData := captData.GetData()
	var masterImageBase64, tileImageBase64 string
	masterImageBase64 = captData.GetMasterImage().ToBase64()
	tileImageBase64 = captData.GetTileImage().ToBase64()
	dotsByte := tool.String(blockData)
	key := tool.MD5([]byte(dotsByte))
	global.G_REDIS.Set(context.Background(), enmus.LOGIN_CAPTCHA+key, dotsByte, 5*time.Minute)
	bt := map[string]interface{}{
		"code":         0,
		"captcha_key":  key,
		"image_base64": masterImageBase64,
		"tile_base64":  tileImageBase64,
		"tile_width":   blockData.Width,
		"tile_height":  blockData.Height,
		"tile_x":       blockData.TileX,
		"tile_y":       blockData.TileY,
	}
	response.SuccessJSON(c, bt)
}

func GetCaptchaDataSlideCheck(c *gin.Context) {
	var jsonp vo.CaptchaDataVO
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "参数提交错误")
		return
	}
	cacheData := global.G_REDIS.Get(context.Background(), enmus.LOGIN_CAPTCHA+jsonp.Key).Val()
	if len(cacheData) == 0 {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "验证失败")
		return
	}
	var dct *slide.Block
	err = tool.JsonUnmarshalFromString(cacheData, &dct)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "请重选验证码")
		return
	}
	src := jsonp.Dots

	chkRet := false
	if 2 == len(src) {
		sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[0]), 64)
		sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[1]), 64)
		chkRet = slide.CheckPoint(int64(sx), int64(sy), int64(dct.X), int64(dct.Y), 6)
	}
	if chkRet {
		global.G_REDIS.Set(context.Background(), enmus.LOGIN_CAPTCHA+jsonp.Key, "true", 5*time.Minute)
	} else {
		global.G_REDIS.Del(context.Background(), enmus.LOGIN_CAPTCHA+jsonp.Key)
	}
	response.SuccessJSON(c, chkRet)
}

*/

func ClearMAF(c *gin.Context) {
	var jsonp dos.AdminUser
	err := c.ShouldBind(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	if !modules.CheckAdminUserSubAccount(c, jsonp.Id) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "无权限，请联系管理员")
		return
	}

	data := modules.ClearAdminUserMfa(&jsonp)
	response.SuccessJSON(c, data)
}

func GetIp(c *gin.Context) {
	ip := c.GetHeader("cf-connecting-ip")
	if len(ip) == 0 {
		ip = c.ClientIP()
	}
	response.SuccessJSON(c, ip)
}

// UpdateMerchantCodes
//
//	@Description: 绑定商户Code集合
//	@param c
func UpdateMerchantCodes(c *gin.Context) {
	var jsonp dos.AdminUser
	err := c.ShouldBindJSON(&jsonp)
	err = validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}
	data := modules.UpdateAdminUserByMerchantCodesId(&jsonp)
	user := modules.FirstByKeyAdminUser(&dos.AdminUser{BaseDos: dos.BaseDos{Id: jsonp.Id}})
	//踢他下线
	global.G_REDIS.Del(context.Background(), enmus.REDIS_LOGIN_USERINFO+user.UserName)
	response.SuccessJSON(c, data)
}

func UpdateStatusBan(c *gin.Context) {
	var jsonp dos.AdminUser
	err := c.ShouldBindJSON(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	err1 := validator.New().Struct(jsonp)
	if err1 != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err1.Error())
		return
	}

	if jsonp.Id == "" {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "Id 不能为空")
		return
	}

	// 校验 status 只能是 0 或 1
	if jsonp.Status != dos.USER_STATUS_DISABLED && jsonp.Status != dos.USER_STATUS_ENABLED {
		response.FailErrJSON(c, response.ERROR_DEFAULT, "Status 只能是 0（禁用）或 1（启用）")
		return
	}
	userInfo, ok := c.Get("UserInfo")
	if ok {
		jsonp.UpdateBy = userInfo.(*dos.AdminUser).UserName
	}

	data := modules.UpdateAdminUserByStatusId(&jsonp)
	user := modules.FirstByKeyAdminUser(&dos.AdminUser{BaseDos: dos.BaseDos{Id: jsonp.Id}})

	if jsonp.Status == dos.USER_STATUS_DISABLED { //如果是禁用就t下线
		global.G_REDIS.Del(context.Background(), enmus.REDIS_LOGIN_USERINFO+user.UserName)
	}

	response.SuccessJSON(c, data)
}

// 字符串
var nonAlphaRegex = regexp.MustCompile(`[^a-zA-Z ]+`)

// 获取 google 密钥
func GetGoogleMfa(str string) (string, error) {
	// 去掉不是字母的字符
	tmpStr := nonAlphaRegex.ReplaceAllString(str, "")
	tmpStrUpper := strings.ToUpper(tmpStr)

	strLen := len(tmpStrUpper)
	if strLen < 1 {
		tmpStr := fmt.Sprintf("secret=%s too short, less than 16 characters", tmpStrUpper)
		return "", errors.New(tmpStr)
	}

	secretStr := ""
	if strLen < 8 {
		secretStr = tmpStrUpper
		for i := strLen; i < 8; i++ {
			secretStr = secretStr + "X"
		}
	} else {
		secretStr = tmpStrUpper[:8] + tmpStrUpper[strLen-8:strLen]
	}

	return secretStr, nil
}

// 账户安全
func AdminUserSecurity(c *gin.Context) {
	var jsonp vo.AdminUserSecurityReq
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

	currUser := modules.GetTokenAdminUser(c)
	if currUser == nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "未登录账号")
		return
	}

	if len(jsonp.Mobile) > 0 {
		// 修改联系方式
		/*if jsonp.Mobile != jsonp.ConfirmMobile {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "确认手机号错误")
			return
		}*/

		err := global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", currUser.Id).Updates(map[string]interface{}{
			"mobile": jsonp.Mobile,
		}).Error
		if err != nil {
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}

		response.SuccessMsgJSON(c, nil, "更新联系方式成功")
		return
	} else if len(jsonp.NewPwd) > 0 {
		token := c.Request.Header.Get(enmus.LOGIN_TOKEN)
		if len(token) == 0 {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "账号未登录")
			return
		}

		// 判断确认密码
		if jsonp.NewPwd != jsonp.ConfirmPwd {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "确认密码错误")
			return
		}

		// 判断密码格式
		if !tool.PwdFormatValid(jsonp.NewPwd) {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "密码格式不正确")
			return
		}

		hasher := tool.GetGlobalPasswordHasher()

		// 判断是否为旧密码
		authUser := dos.AdminUser{}
		err := global.G_DB.Model(&dos.AdminUser{}).Select("pwd").Where("id = ?", currUser.Id).First(&authUser).Error
		if err != nil {
			response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
			return
		}
		if samePwd, _ := hasher.VerifyPassword(jsonp.NewPwd, authUser.Pwd); samePwd {
			response.FailErrJSON(c, response.ERROR_PARAMETER, "新密码不能跟旧密码一样")
			return
		}

		password, hashErr := hasher.HashPassword(jsonp.NewPwd)
		if hashErr != nil {
			response.FailErrJSON(c, response.ERROR_SERVER, "密码加密失败")
			return
		}

		err = global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", currUser.Id).Updates(map[string]interface{}{
			"pwd":         password,
			"enforce_pwd": 0,
		}).Error
		if err != nil {
			response.FailErrJSON(c, response.ERROR_SERVER, err.Error())
			return
		}

		// 删除 token, 删除用户信息, 删除用户登录错误次数
		sessionId, _ := c.Cookie(enmus.LOGIN_COOKIE)
		tokenKey := enmus.REDIS_LOGIN_TOKEN + sessionId
		userInfoKey := enmus.REDIS_LOGIN_USERINFO + currUser.UserName
		userLoginCountErrKey := enmus.REDIS_LOGIN_ERR_COUNT + currUser.UserName
		global.G_REDIS.Del(context.Background(), tokenKey)
		global.G_REDIS.Del(context.Background(), userInfoKey)
		global.G_REDIS.Del(context.Background(), userLoginCountErrKey)
		response.SuccessMsgJSON(c, nil, "更新密码成功")
		return
	}

	response.FailErrJSON(c, response.ERROR_PARAMETER, "参数错误")
}
