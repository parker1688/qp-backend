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
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/kirinlabs/utils"
	"github.com/kirinlabs/utils/encrypt"
)

// api: api/fcUserMaterial/update
func UpdateFcUserMaterialAgent(c *gin.Context) {
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

	data := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(map[string]interface{}{
		"agent_id":  jsonp.AgentId,
		"parent_id": jsonp.ParentId,
	}).Error
	response.SuccessJSON(c, data == nil)
}

// api: api/fcUserMaterial/update
func UpdateFcUserMaterialIsFree(c *gin.Context) {
	var jsonp struct {
		UserId         string `json:"user_id"`
		IsFree         bool   `json:"is_free"`
		LoginStatus    int    `json:"login_status"`
		IsVerification bool   `json:"is_verification"`
		Remark         string `json:"remark"`
	}
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
	eRow := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(map[string]interface{}{
		"is_free":         jsonp.IsFree,
		"login_status":    jsonp.LoginStatus,
		"is_verification": jsonp.IsVerification,
		"remark":          jsonp.Remark,
	})
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}

	mm := &dos.FcUserLogin{}
	mm.Id = jsonp.UserId
	m := modules.FindByKeyFcUserLoginFirst(mm)
	userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	err = global.G_REDIS.Del(context.Background(), memberRedisKey).Err()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	response.SuccessJSON(c, true)
}

// api: api/fcUserMaterial/update
func UpdateFcUserMaterialIsWithdraw(c *gin.Context) {
	var jsonp struct {
		UserId     string `json:"user_id"`
		IsWithdraw int    `json:"is_withdraw"`
	}
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

	mm := &dos.FcUserLogin{}
	mm.Id = jsonp.UserId
	m := modules.FindByKeyFcUserLoginFirst(mm)
	if !modules.CheckAdminUserMerchantPerms(c, m.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	updateBy := ""
	userInfo, ok := c.Get("UserInfo")
	if ok {
		updateBy = userInfo.(*dos.AdminUser).UserName
	}
	updateMap := map[string]interface{}{}
	updateMap["is_withdraw"] = jsonp.IsWithdraw
	if updateBy != "" {
		updateMap["update_by"] = updateBy
	}

	eRow := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(updateMap)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}

	userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	err = global.G_REDIS.Del(context.Background(), memberRedisKey).Err()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	response.SuccessJSON(c, true)
}

// api: api/fcUserMaterial/update
func UpdateFcUserMaterialIsBonus(c *gin.Context) {
	var jsonp struct {
		UserId  string `json:"user_id"`
		IsBonus int    `json:"is_bonus"`
	}
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

	mm := &dos.FcUserLogin{}
	mm.Id = jsonp.UserId
	m := modules.FindByKeyFcUserLoginFirst(mm)
	if !modules.CheckAdminUserMerchantPerms(c, m.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	updateBy := ""
	userInfo, ok := c.Get("UserInfo")
	if ok {
		updateBy = userInfo.(*dos.AdminUser).UserName
	}
	updateMap := map[string]interface{}{}
	updateMap["is_bonus"] = jsonp.IsBonus
	if updateBy != "" {
		updateMap["update_by"] = updateBy
	}

	eRow := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(updateMap)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}

	userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	err = global.G_REDIS.Del(context.Background(), memberRedisKey).Err()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	response.SuccessJSON(c, true)
}

// api: api/fcUserMaterial/update
func UpdateFcUserMaterialRealName(c *gin.Context) {
	var jsonp struct {
		UserId   string `json:"user_id"`
		RealName string `json:"real_name"`
	}
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

	mm := &dos.FcUserLogin{}
	mm.Id = jsonp.UserId
	m := modules.FindByKeyFcUserLoginFirst(mm)
	if !modules.CheckAdminUserMerchantPerms(c, m.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	updateBy := ""
	userInfo, ok := c.Get("UserInfo")
	if ok {
		updateBy = userInfo.(*dos.AdminUser).UserName
	}

	u := dos.FcUserMaterial{}
	u.RealName = jsonp.RealName
	u.Encrypt()

	updateMap := map[string]interface{}{}
	updateMap["real_name"] = u.RealName
	if updateBy != "" {
		updateMap["update_by"] = updateBy
	}

	eRow := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(updateMap)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	/*err = global.G_REDIS.Del(context.Background(), memberRedisKey).Err()*/
	// 直接更新缓存用户可以不下线
	newUserInfo := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: jsonp.UserId,
	})
	if len(newUserInfo.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	err = global.G_REDIS.Set(context.Background(), memberRedisKey, utils.Json(newUserInfo), 16*24*time.Hour).Err()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	response.SuccessJSON(c, true)
}

// api: api/fcUserMaterial/update
func ClearFcUserMaterialTel(c *gin.Context) {
	var jsonp struct {
		UserId       string `json:"user_id"`
		Tel          string `json:"tel"`
		MerchantCode string `json:"merchant_code"`
		Remark       string `json:"remark"`
	}
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

	mm := &dos.FcUserLogin{}
	mm.Id = jsonp.UserId
	m := modules.FindByKeyFcUserLoginFirst(mm)
	if !modules.CheckAdminUserMerchantPerms(c, m.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	//u := dos.FcUserMaterial{}
	//u.Tel = jsonp.Tel
	//u.Remark = jsonp.Remark
	//u.Encrypt()

	updateBy := ""
	userInfo, ok := c.Get("UserInfo")
	if ok {
		updateBy = userInfo.(*dos.AdminUser).UserName
	}

	//// 判断手机号码是否存在，同一个包网商户只允许存在一个
	//if jsonp.Tel != "" {
	//	var telExistCount int64
	//	err = global.G_DB.Model(&dos.FcUserMaterial{}).Where("tel = ? AND merchant_code = ?", u.Tel, jsonp.MerchantCode).Count(&telExistCount).Error
	//	if err != nil {
	//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
	//		return
	//	}
	//
	//	if telExistCount > 0 {
	//		global.G_LOG.Infof("UpdateFcUserMaterialTel tel: %v already exist", jsonp.Tel)
	//		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
	//		return
	//	}
	//}

	updateMap := map[string]interface{}{}
	updateMap["tel"] = ""
	if updateBy != "" {
		updateMap["update_by"] = updateBy
	}

	eRow := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(updateMap)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}

	userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	//err = global.G_REDIS.Del(context.Background(), memberRedisKey).Err()
	// 直接更新缓存用户可以不下线
	newUserInfo := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: jsonp.UserId,
	})
	if len(newUserInfo.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	err = global.G_REDIS.Set(context.Background(), memberRedisKey, utils.Json(newUserInfo), 16*24*time.Hour).Err()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	response.SuccessJSON(c, true)
}

// api: api/fcUserMaterial/update
func UpdateFcUserMaterialLoginStatus(c *gin.Context) {
	var jsonp struct {
		UserId      string `json:"user_id"`
		LoginStatus int    `json:"login_status"`
	}
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

	mm := &dos.FcUserLogin{}
	mm.Id = jsonp.UserId
	m := modules.FindByKeyFcUserLoginFirst(mm)
	if !modules.CheckAdminUserMerchantPerms(c, m.MerchantCode) {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "没有对该商户的处理权限")
		return
	}

	u := dos.FcUserMaterial{}
	u.LoginStatus = jsonp.LoginStatus

	updateBy := ""
	userInfo, ok := c.Get("UserInfo")
	if ok {
		updateBy = userInfo.(*dos.AdminUser).UserName
	}

	updateMap := map[string]interface{}{}
	updateMap["login_status"] = u.LoginStatus
	if updateBy != "" {
		updateMap["update_by"] = updateBy
	}

	eRow := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(updateMap)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}

	userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	err = global.G_REDIS.Del(context.Background(), memberRedisKey).Err()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	response.SuccessJSON(c, true)
}

// api: api/fcUserMaterial/update
func ClearFcUserMaterialWalletPwd(c *gin.Context) {
	var jsonp struct {
		UserId    string `json:"user_id"`
		WalletPwd string `json:"wallet_password"`
	}
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

	updateBy := ""
	userInfo, ok := c.Get("UserInfo")
	if ok {
		updateBy = userInfo.(*dos.AdminUser).UserName
	}

	updateMap := map[string]interface{}{}
	updateMap["wallet_password"] = ""
	if updateBy != "" {
		updateMap["update_by"] = updateBy
	}

	eRow := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(updateMap)
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}

	mm := &dos.FcUserLogin{}
	mm.Id = jsonp.UserId
	m := modules.FindByKeyFcUserLoginFirst(mm)
	userNameM := fmt.Sprintf("%s:%s", m.MerchantCode, m.UserName)
	memberRedisKey := fmt.Sprintf(enmus.REDIS_LOGIN_MEMBERINFO, userNameM)
	err = global.G_REDIS.Del(context.Background(), memberRedisKey).Err()
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "操作失败,请重试")
		return
	}
	response.SuccessJSON(c, true)
}

func SimulateUser(c *gin.Context) {
	var jsonp struct {
		Num          int64  `json:"num"`
		Password     string `json:"password"`
		MerchantCode string `json:"merchant_code"` // 商户code
	}
	err := c.ShouldBind(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	if jsonp.Num > 50 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "最大生成50个")
		return
	}

	password := encrypt.Sha256(jsonp.Password + global.CONFIG.General.ApiSHA256Salt)
	var count int64
	global.G_DB.Model(&dos.FcUserLogin{}).Count(&count)
	count = count + 1
	for i := count; i < count+jsonp.Num; i++ {
		username := fmt.Sprintf("9%010s", tool.String(i))
		userLogin := &dos.FcUserLogin{
			UserName:     username,
			Password:     password,
			MerchantCode: jsonp.MerchantCode,
		}
		saveOk, _ := modules.SaveFcUserLogin(userLogin)
		if !saveOk {
			response.FailErrJSON(c, response.ERROR_SERVER, "")
			return
		}
		material := &dos.FcUserMaterial{
			UserId:        userLogin.Id,
			UserName:      userLogin.UserName,
			Tel:           "",
			MerchantCode:  jsonp.MerchantCode,
			Vip:           "VIP0",
			Level:         0,
			ParentId:      "", //邀请好友
			AgentId:       "", //代理邀请
			Language:      "",
			RegisterIp:    "8.8.8.8",
			LastLoginIp:   "8.8.8.8",
			IsFree:        true,
			LastLoginTime: automaticType.Time(time.Now()),
			InviteCode:    tool.HashEncodeInt64([]int64{tool.Int(userLogin.Id)}),
		}
		saveOk, _ = modules.SaveFcUserMaterial(material)
		if !saveOk {
			response.FailErrJSON(c, response.ERROR_SERVER, "")
			return
		}
	}
	response.SuccessJSON(c, true)
}

func UpdateFcUserMaterialRemark(c *gin.Context) {
	var jsonp struct {
		UserId string `json:"user_id"`
		Remark string `json:"remark"`
	}
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
	eRow := global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", jsonp.UserId).Updates(map[string]interface{}{
		"remark": jsonp.Remark,
	})
	if eRow.Error != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, eRow.Error.Error())
		return
	}

	response.SuccessJSON(c, true)
}

// api: api/fcUserMaterial/findPage
func FindPageFcUserMaterialSameControl(c *gin.Context) {
	jsonp := struct {
		dos.FcUserMaterial
		response.PageTimeQuery
	}{}
	response.NormalizePageTimeQuery(&jsonp.PageTimeQuery)
	jsonp.UserId = c.DefaultQuery("user_id", "")
	jsonp.MerchantCode = c.DefaultQuery("merchant_code", "")
	err := validator.New().Struct(jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}
	d := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: jsonp.UserId,
	})
	//查詢相同IP
	jsonp.UserId = ""
	jsonp.LastLoginIp = d.LastLoginIp
	if len(jsonp.LastLoginIp) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "无相同IP")
		return
	}
	data, total := modules.FindPageFcUserMaterial(jsonp.PageNo, jsonp.PageSize, &jsonp.FcUserMaterial, jsonp.PageTimeQuery, false, c)

	vof := make([]*vo.FcUserMaterialVO, 0, len(data))
	tool.JsonMapper(data, &vof)
	for k, v := range vof {
		if v.Email != "" {
			vof[k].Email = v.Email[0:3] + "*******"
		}
		if v.WalletPassword != "" {
			vof[k].WalletPassword = "已设置"
		}

		if v.RealName != "" {
			realNameLen := len(v.RealName)
			if realNameLen > 2 {
				vof[k].RealName = v.RealName[0:1] + "**"
			} else {
				vof[k].RealName = v.RealName[0:1] + "*"
			}
		}
	}
	response.SuccessPageJSON(c, jsonp.PageNo, jsonp.PageSize, total, vof)
}

// api: /api/fcUserMaterial/merchant
func GetFcUserMaterialMerchantControl(c *gin.Context) {
	var jsonp struct {
		UserId string `json:"user_id" form:"user_id"`
	}

	err := c.ShouldBindQuery(&jsonp)
	if err != nil {
		response.FailErrJSON(c, response.ERROR_PARAMETER, err.Error())
		return
	}

	if len(jsonp.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "用户ID为空")
		return
	}

	data := modules.FindByKeyFcUserMaterialFirst(&dos.FcUserMaterial{
		UserId: jsonp.UserId,
	})

	if data == nil || len(data.UserId) == 0 {
		response.FailErrJSON(c, response.ERROR_PARAMETER, "无法找到用户对应商户")
		return
	}

	response.SuccessJSON(c, struct {
		MerchantCode string `json:"merchant_code"`
		MerchantName string `json:"merchant_name"`
	}{
		MerchantCode: data.MerchantCode,
		MerchantName: modules.GetMerchantName(data.MerchantCode),
	})
}
