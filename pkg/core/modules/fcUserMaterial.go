// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcUserMaterial(vo *dos.FcUserMaterial) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserMaterial(page, pageSize int, vo *dos.FcUserMaterial, pageQuery response.PageTimeQuery, queryIp bool, c *gin.Context) (ret []*dos.FcUserMaterial, total int64) {
	page, pageSize = response.NormalizePage(page, pageSize)
	query := global.G_DB.Model(&dos.FcUserMaterial{})

	if len(vo.UserId) > 0 {
		query = query.Where("user_id like ?", "%"+vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name like ?", "%"+vo.UserName+"%")
	}

	if len(vo.NickName) > 0 {
		query = query.Where("nick_name like ?", "%"+vo.NickName+"%")
	}

	if len(vo.RealName) > 0 {
		query = query.Where("real_name = ?", vo.RealName)
	}

	if len(vo.ParentId) > 0 {
		query = query.Where("parent_id = ?", vo.ParentId)
	}

	if len(vo.AgentId) > 0 {
		query = query.Where("agent_id = ?", vo.AgentId)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.Sex > 0 {
		query = query.Where("sex = ?", vo.Sex)
	}

	if len(vo.Tel) > 0 {
		query = query.Where("tel = ?", vo.Tel)
	}

	if len(vo.Email) > 0 {
		query = query.Where("email = ?", vo.Email)
	}

	if len(vo.Qq) > 0 {
		query = query.Where("qq = ?", vo.Qq)
	}

	if len(vo.Wx) > 0 {
		query = query.Where("wx = ?", vo.Wx)
	}

	if len(vo.Address) > 0 {
		query = query.Where("address = ?", vo.Address)
	}

	if !vo.Birthday.Timer().IsZero() {
		query = query.Where("birthday = ?", vo.Birthday)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.Vip) > 0 {
		query = query.Where("vip = ?", vo.Vip)
	}

	if vo.AgentInviteCode > 0 {
		query = query.Where("agent_invite_code = ?", vo.AgentInviteCode)
	}

	/*if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/

	if len(vo.Nation) > 0 {
		query = query.Where("nation = ?", vo.Nation)
	}

	if queryIp {
		// 注册ip和登录ip一起查询
		if len(vo.RegisterIp) > 0 {
			query = query.Where("register_ip = ?", vo.RegisterIp)

			query = query.Or("last_login_ip = ?", vo.RegisterIp)
		}
	} else {
		if len(vo.RegisterIp) > 0 {
			query = query.Where("register_ip = ?", vo.RegisterIp)
		}

		if len(vo.LastLoginIp) > 0 {
			query = query.Where("last_login_ip = ?", vo.LastLoginIp)
		}
	}

	if len(vo.RegistVisitorId) > 0 {
		query = query.Where("regist_visitor_id = ?", vo.RegistVisitorId)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.Avatar) > 0 {
		query = query.Where("avatar = ?", vo.Avatar)
	}

	if len(vo.AgentSubId) > 0 {
		query = query.Where("agent_sub_id = ?", vo.AgentSubId)
	}

	if len(vo.AgentSubName) > 0 {
		query = query.Where("agent_sub_name = ?", vo.AgentSubName)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}
	if len(pageQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", pageQuery.StartAt)
	}
	if len(pageQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", pageQuery.EndAt)
	}
	if len(pageQuery.LastStartAt) > 0 {
		query = query.Where("last_login_time >= ?", pageQuery.LastStartAt)
	}
	if len(pageQuery.LastEndAt) > 0 {
		query = query.Where("last_login_time <= ?", pageQuery.LastEndAt)
	}
	if len(pageQuery.IsFree) > 0 {
		query = query.Where("is_free = ?", pageQuery.IsFree == "true")
	}

	if vo.IsVerification {
		query = query.Where("is_verification = ?", vo.IsVerification)
	}
	if c == nil && len(vo.LastLoginIp) > 0 {
		query = query.Where("last_login_ip = ?", vo.LastLoginIp)
	}

	if len(vo.VisitorId) > 0 {
		query = query.Where("visitor_id = ?", vo.VisitorId)
	}
	if len(vo.Website) > 0 {
		query = query.Where("website = ?", vo.Website)
	}
	if len(vo.RankFlag) > 0 {
		query = query.Where("rank_flag = ?", vo.RankFlag)
	}
	if len(vo.Rank) > 0 {
		query = query.Where("rank = ?", vo.Rank)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserMaterial
	query.Offset((page - 1) * pageSize).Limit(pageSize).Order("create_time desc").Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserMaterial(vo *dos.FcUserMaterial) []*dos.FcUserMaterial {
	var data []*dos.FcUserMaterial
	query := global.G_DB.Model(&dos.FcUserMaterial{})

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.RealName) > 0 {
		query = query.Where("real_name = ?", vo.RealName)
	}

	if len(vo.ParentId) > 0 {
		query = query.Where("parent_id = ?", vo.ParentId)
	}

	if len(vo.AgentId) > 0 {
		query = query.Where("agent_id = ?", vo.AgentId)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.Sex > 0 {
		query = query.Where("sex = ?", vo.Sex)
	}

	if len(vo.Tel) > 0 {
		query = query.Where("tel = ?", vo.Tel)
	}

	if len(vo.RegisterIp) > 0 {
		query = query.Where("register_ip = ?", vo.RegisterIp)
	}

	if len(vo.LastLoginIp) > 0 {
		query = query.Where("last_login_ip = ?", vo.LastLoginIp)
	}

	if len(vo.RegistVisitorId) > 0 {
		query = query.Where("regist_visitor_id = ?", vo.RegistVisitorId)
	}

	if len(vo.Email) > 0 {
		query = query.Where("email = ?", vo.Email)
	}

	if len(vo.Qq) > 0 {
		query = query.Where("qq = ?", vo.Qq)
	}

	if len(vo.Wx) > 0 {
		query = query.Where("wx = ?", vo.Wx)
	}

	if len(vo.Address) > 0 {
		query = query.Where("address = ?", vo.Address)
	}

	if !vo.Birthday.Timer().IsZero() {
		query = query.Where("birthday = ?", vo.Birthday)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.Vip) > 0 {
		query = query.Where("vip = ?", vo.Vip)
	}

	if vo.AgentInviteCode > 0 {
		query = query.Where("agent_invite_code = ?", vo.AgentInviteCode)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.Nation) > 0 {
		query = query.Where("nation = ?", vo.Nation)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.Avatar) > 0 {
		query = query.Where("avatar = ?", vo.Avatar)
	}

	if len(vo.AgentSubId) > 0 {
		query = query.Where("agent_sub_id = ?", vo.AgentSubId)
	}

	if len(vo.AgentSubName) > 0 {
		query = query.Where("agent_sub_name = ?", vo.AgentSubName)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if vo.IsVerification {
		query = query.Where("is_verification = ?", vo.IsVerification)
	}
	if len(vo.VisitorId) > 0 {
		query = query.Where("visitor_id = ?", vo.VisitorId)
	}
	if len(vo.Website) > 0 {
		query = query.Where("website = ?", vo.Website)
	}
	if len(vo.RankFlag) > 0 {
		query = query.Where("rank_flag = ?", vo.RankFlag)
	}
	if len(vo.Rank) > 0 {
		query = query.Where("rank = ?", vo.Rank)
	}
	query.Order("create_time desc").Find(&data)
	return data
}

func FindByKeyFcUserMaterialFirst(vo *dos.FcUserMaterial) *dos.FcUserMaterial {
	var data *dos.FcUserMaterial
	query := global.G_DB.Model(&dos.FcUserMaterial{})

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.RealName) > 0 {
		query = query.Where("real_name = ?", vo.RealName)
	}

	if len(vo.ParentId) > 0 {
		query = query.Where("parent_id = ?", vo.ParentId)
	}

	if len(vo.AgentId) > 0 {
		query = query.Where("agent_id = ?", vo.AgentId)
	}

	if len(vo.AgentName) > 0 {
		query = query.Where("agent_name = ?", vo.AgentName)
	}

	if vo.Sex > 0 {
		query = query.Where("sex = ?", vo.Sex)
	}

	if len(vo.Tel) > 0 {
		query = query.Where("tel = ?", vo.Tel)
	}

	if len(vo.Email) > 0 {
		query = query.Where("email = ?", vo.Email)
	}

	if len(vo.Qq) > 0 {
		query = query.Where("qq = ?", vo.Qq)
	}

	if len(vo.Wx) > 0 {
		query = query.Where("wx = ?", vo.Wx)
	}

	if len(vo.Address) > 0 {
		query = query.Where("address = ?", vo.Address)
	}

	if !vo.Birthday.Timer().IsZero() {
		query = query.Where("birthday = ?", vo.Birthday)
	}

	if vo.Level > 0 {
		query = query.Where("level = ?", vo.Level)
	}

	if len(vo.Vip) > 0 {
		query = query.Where("vip = ?", vo.Vip)
	}

	if vo.AgentInviteCode > 0 {
		query = query.Where("agent_invite_code = ?", vo.AgentInviteCode)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.Nation) > 0 {
		query = query.Where("nation = ?", vo.Nation)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.Avatar) > 0 {
		query = query.Where("avatar = ?", vo.Avatar)
	}

	if len(vo.AgentSubId) > 0 {
		query = query.Where("agent_sub_id = ?", vo.AgentSubId)
	}

	if len(vo.AgentSubName) > 0 {
		query = query.Where("agent_sub_name = ?", vo.AgentSubName)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if !vo.UpdateTime.Timer().IsZero() {
		query = query.Where("update_time = ?", vo.UpdateTime)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if vo.IsVerification {
		query = query.Where("is_verification = ?", vo.IsVerification)
	}
	if len(vo.VisitorId) > 0 {
		query = query.Where("visitor_id = ?", vo.VisitorId)
	}
	if len(vo.Website) > 0 {
		query = query.Where("website = ?", vo.Website)
	}
	if len(vo.RankFlag) > 0 {
		query = query.Where("rank_flag = ?", vo.RankFlag)
	}
	if len(vo.Rank) > 0 {
		query = query.Where("rank = ?", vo.Rank)
	}
	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcUserMaterial(vo *dos.FcUserMaterial) bool {
	data := map[string]interface{}{

		"sex":               vo.Sex,
		"qq":                vo.Qq,
		"wx":                vo.Wx,
		"address":           vo.Address,
		"level":             vo.Level,
		"nation":            vo.Nation,
		"language":          vo.Language,
		"avatar":            vo.Avatar,
		"update_by":         vo.UpdateBy,
		"is_verification":   vo.IsVerification,
		"website":           vo.Website,
		"rank":              vo.Rank,
		"rank_flag":         vo.RankFlag,
		"is_withdraw":       vo.IsWithdraw,
		"is_bonus":          vo.IsBonus,
		"agent_invite_code": vo.AgentInviteCode,

		//"wallet_password": vo.WalletPassword,
		//"alipay":          vo.Alipay,
		//"alipay_realname": vo.AlipayRealname,
		//"email":           vo.Email,
		//"birthday":  vo.Birthday,
		//"tel":             vo.Tel,
		//"real_name":       vo.RealName,

		//"visitor_id":      vo.VisitorId,
		//"register_ip":     vo.RegisterIp,
		//"last_login_ip":   vo.LastLoginIp,
		//"last_login_time": vo.LastLoginTime,
	}

	if vo.Birthday.String() != "0001-01-01 00:00:00" && vo.Birthday.String() != "" {
		data["birthday"] = vo.Birthday
	}
	if len(vo.Avatar) > 0 {
		data["avatar"] = vo.Avatar
	}
	return global.G_DB.Model(vo).Where(`user_id = ?`, vo.UserId).Updates(data).Error == nil
}

func DeleteFcUserMaterial(vo *dos.FcUserMaterial) bool {
	return global.G_DB.Model(&dos.FcUserMaterial{}).Where("user_id = ?", vo.UserId).Delete(vo).Error == nil
}

// 根据主键Update
func UpdateFcUserMaterialLogin(vo *dos.FcUserMaterial) bool {
	return global.G_DB.Model(vo).Where(`user_id = ?`, vo.UserId).Updates(map[string]interface{}{
		"last_login_ip":    vo.LastLoginIp,
		"last_login_time":  vo.LastLoginTime,
		"last_login_count": vo.LastLoginCount,
	}).Error == nil
}

func SetAbnormalLogonRecord(key string, value string) {
	if len(key) == 0 {
		return
	}

	if !CheckAbnormalLogonRecord(key) {
		global.G_DB.Model(&dos.FcAbnormalLogonRecord{}).Create(&dos.FcAbnormalLogonRecord{
			Key:   key,
			Value: "1",
		})
	}
}

func CheckAbnormalLogonRecord(key string) bool {
	data := dos.FcAbnormalLogonRecord{}
	global.G_DB.Model(&dos.FcAbnormalLogonRecord{}).Select("value").Where("`key` = ?", key).Take(&data)
	return data.Value == "1"
}

func CheckAbnormalLogonPass(ip string, visitorId string) bool {
	return CheckAbnormalLogonRecord(visitorId)
}

func RecordAbnormalLogonData(ip string, visitorId string) {
	SetAbnormalLogonRecord(ip, "1")
	SetAbnormalLogonRecord(visitorId, "1")
}

// GetFcUserMaterialRegTime - 获取用户注册时间
// @param {string} userId
// @returns string
func GetFcUserMaterialRegTime(userId string) string {
	data := dos.FcUserMaterial{}
	err := global.G_DB.Model(&dos.FcUserMaterial{}).
		Select("create_time").
		Where("user_id = ?", userId).
		Take(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcUserMaterialRegTime] Find user material register time failed: userId=%s, err=%s",
			userId, err.Error())
		return ""
	}

	return data.CreateTime.String()
}
