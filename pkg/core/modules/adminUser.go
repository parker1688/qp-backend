// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/utils/encrypt"
	"gorm.io/gorm"
)

func SaveAdminUser(vo *dos.AdminUser) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageAdminUser(page, pageSize int, vo *dos.AdminUser, c *gin.Context) (ret []*dos.AdminUser, total int64) {
	query := global.G_DB.Model(&dos.AdminUser{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("username = ?", vo.UserName)
	}

	if len(vo.UserNick) > 0 {
		query = query.Where("user_nick = ?", vo.UserNick)
	}

	if vo.AccountType > 0 {
		query = query.Where("account = ?", vo.AccountType)
	}

	if len(vo.Mobile) > 0 {
		query = query.Where("mobile = ?", vo.Mobile)
	}

	if len(vo.DepartmentId) > 0 {
		query = query.Where("department_id = ?", vo.DepartmentId)
	}

	if len(vo.RoleIds) > 0 {
		query = query.Where("role_ids = ?", vo.RoleIds)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserListWithPerms(c, query, vo); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.AdminUser
	query.Offset((page - 1) * pageSize).Limit(pageSize).Order("create_time desc").Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyAdminUser(vo *dos.AdminUser) []*dos.AdminUser {
	var data []*dos.AdminUser
	query := global.G_DB.Model(&dos.AdminUser{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}
	if len(vo.UserName) > 0 {
		query = query.Where("username = ?", vo.UserName)
	}
	query.Take(&data)
	return data
}

func FindByKeyAdminUserWithPerms(vo *dos.AdminUser, c *gin.Context) []*dos.AdminUser {
	var data []*dos.AdminUser
	query := global.G_DB.Model(&dos.AdminUser{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}
	if len(vo.UserName) > 0 {
		query = query.Where("username = ?", vo.UserName)
	}

	merchantLis := GetAdminUserMerchantList(c)
	adminUser := GetTokenAdminUser(c)
	if adminUser != nil {
		query = query.Where("account_type >= ?", adminUser.AccountType)
		query = query.Where("merchant_codes in ?", merchantLis)
	} else {
		return data
	}

	query.Take(&data)
	return data
}

func FirstByKeyAdminUser(vo *dos.AdminUser) *dos.AdminUser {
	var data *dos.AdminUser
	query := global.G_DB.Model(&dos.AdminUser{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}
	if len(vo.UserName) > 0 {
		query = query.Where("username = ?", vo.UserName)
	}
	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateAdminUser(vo *dos.AdminUser) bool {
	updates := map[string]interface{}{
		"username":            vo.UserName,
		"user_nick":            vo.UserNick,
		"mobile":               vo.Mobile,
		"account_type":         vo.AccountType,
		"status":               vo.Status,
		"merchant_codes":       vo.MerchantCodes,
		"total_amount":         vo.TotalAmount,
		"limit_pertime_amount": vo.LimitPertimeAmount,
		//"create_by":      vo.CreateBy,
		"update_by": vo.UpdateBy,
	}
	if len(vo.Pwd) > 0 {
		updates["pwd"] = encrypt.Sha256(vo.Pwd + global.CONFIG.SHA256Salt)
		updates["enforce_pwd"] = 1 //强制修改密码
	}

	return global.G_DB.Model(vo).Where("id=?", vo.Id).Updates(updates).Error == nil
}

func DeleteAdminUser(vo *dos.AdminUser) bool {
	return global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", vo.Id).Delete(vo).Error != nil
}

func UpdateAdminUserByDepartmentId(vo *dos.AdminUser) bool {
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"department_id": vo.DepartmentId,
	}).Error == nil
}

func UpdateAdminUserByRoleIds(vo *dos.AdminUser) bool {
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"role_ids": vo.RoleIds,
	}).Error == nil
}

func FindUserNameByKeyAdminUser(vo *dos.AdminUser, c *gin.Context) []*dos.AdminUser {
	var data []*dos.AdminUser
	query := global.G_DB.Model(&dos.AdminUser{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserListWithPerms(c, query, vo); !ok {
			return data
		}
	}

	query.Select("id", "username").Find(&data)
	return data
}

func FindByKeyAdmindUserFirst(vo *dos.AdminUser) *dos.AdminUser {
	query := global.G_DB.Model(&dos.AdminUser{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}
	adminUser := &dos.AdminUser{}
	err := query.First(&adminUser).Error
	if err != nil {
		global.G_LOG.Errorf("[FindByKeyAdmindUserFirst] find admin user failed err: %v", err.Error())
	}

	return adminUser
}

func UpdateAdminUserPassword(vo *dos.AdminUser) bool {
	updates := map[string]interface{}{
		"update_by": vo.UpdateBy,
	}
	updates["pwd"] = encrypt.Sha256(vo.Pwd + global.CONFIG.SHA256Salt)
	updates["enforce_pwd"] = 0 //已修改密码
	return global.G_DB.Model(vo).Where("id=?", vo.Id).Updates(updates).Error == nil
}

func UpdateAdminUserMfa(vo *dos.AdminUser) bool {
	updates := map[string]interface{}{
		"update_by": vo.UpdateBy,
	}
	updates["mfa"] = vo.Mfa
	updates["mfa_hour"] = vo.MfaHour
	return global.G_DB.Model(vo).Where("id=?", vo.Id).Updates(updates).Error == nil
}

func ClearAdminUserMfa(vo *dos.AdminUser) bool {
	updates := map[string]interface{}{
		"update_by": vo.UpdateBy,
	}
	updates["mfa"] = ""
	return global.G_DB.Model(vo).Where("id=?", vo.Id).Updates(updates).Error == nil
}

func UpdateAdminUserByMerchantCodesId(vo *dos.AdminUser) bool {
	return global.G_DB.Model(vo).Where("id=?", vo.Id).Updates(map[string]interface{}{
		"merchant_codes": vo.MerchantCodes,
	}).Error == nil
}

func UpdateAdminUserByStatusId(vo *dos.AdminUser) bool {
	return global.G_DB.Model(vo).Where("id=?", vo.Id).Updates(map[string]interface{}{
		"status": vo.Status,
	}).Error == nil
}

func QueryAdminUserMerchantCodes(c *gin.Context, query *gorm.DB, queryMerchantCode string) (*gorm.DB, bool) {
	userInfo, ok := c.Get("UserInfo")
	if ok {
		adminUser := userInfo.(*dos.AdminUser)
		if adminUser.AccountType == 1 {
			// 超级管理员
			if len(queryMerchantCode) > 0 {
				query = query.Where("merchant_code like ?", "%"+queryMerchantCode+"%")
			}
			return query, true
		}

		merchantCodes := userInfo.(*dos.AdminUser).MerchantCodes
		if len(merchantCodes) > 0 {
			merchantCodeLis := strings.Split(merchantCodes, ",")
			if len(queryMerchantCode) > 0 && slices.Contains(merchantCodeLis, queryMerchantCode) {
				// 是查询商户并且需要在自己商户权限中
				return query.Where("merchant_code = ?", queryMerchantCode), true
			} else if len(queryMerchantCode) == 0 {
				// 非查询商户
				return query.Where("merchant_code in ?", merchantCodeLis), true
			}
		}
	}

	return query, false
}

func CheckAdminUserMerchantPerms(c *gin.Context, merchantCode string) bool {
	userInfo, ok := c.Get("UserInfo")
	if ok {
		adminUser := userInfo.(*dos.AdminUser)
		if adminUser.AccountType == 1 {
			// 是超级管理员不判断商户
			return true
		}

		merchantCodeLis := strings.Split(adminUser.MerchantCodes, ",")
		return slices.Contains(merchantCodeLis, merchantCode)
	}

	return false
}

func GetAdminUserMerchantList(c *gin.Context) []string {
	userInfo, ok := c.Get("UserInfo")
	if ok {
		if userInfo.(*dos.AdminUser).AccountType == 1 {
			// 是超级管理员获取全部商户
			merchants := []dos.FcMerchant{}
			err := global.G_DB.Model(&dos.FcMerchant{}).Select("merchant_code").Find(&merchants).Error
			if err != nil {
				global.G_LOG.Errorf("[GetAdminUserMerchantList] Find merchants failed err: %v", err.Error())
				return []string{}
			}
			merchantLis := []string{}
			for _, v := range merchants {
				merchantLis = append(merchantLis, v.MerchantCode)
			}

			return merchantLis
		}

		merchantCodes := userInfo.(*dos.AdminUser).MerchantCodes
		return strings.Split(merchantCodes, ",")
	}

	return []string{}
}

func QueryAdminUserListWithPerms(c *gin.Context, query *gorm.DB, vo *dos.AdminUser) (*gorm.DB, bool) {
	isQuery := len(vo.Id) > 0 || len(vo.UserName) > 0 || len(vo.UserNick) > 0 || vo.AccountType > 0 ||
		len(vo.Mobile) > 0 || len(vo.DepartmentId) > 0 || len(vo.RoleIds) > 0 || vo.Status > 0 ||
		len(vo.Remarks) > 0 || len(vo.CreateBy) > 0 || len(vo.UpdateBy) > 0

	userInfo, ok := c.Get("UserInfo")
	if ok {
		adminUser := userInfo.(*dos.AdminUser)
		if adminUser.AccountType == 1 {
			// 超级管理员则全部显示
			return query, true
		} else if adminUser.AccountType == 2 {
			// 商户管理员只显示自己、子账号、关联商户的子账号
			merchantLis := strings.Split(adminUser.MerchantCodes, ",")
			if !isQuery {
				query = query.Or("username = ?", adminUser.UserName)
				query = query.Or("create_by = ?", adminUser.UserName)
				//query = query.Or("account_type > ? and merchant_codes in ?", adminUser.AccountType, merchantLis)
				//len := len(merchantLis)
				//if len == 1 {
				//	query = query.Or("account_type > ? and merchant_codes like ?", adminUser.AccountType, "%"+merchantLis[0]+"%")
				//} else {
				for _, code := range merchantLis {
					//if i == 0 {
					//	query = query.Where("(account_type > ? and merchant_codes like ?", adminUser.AccountType, "%"+code+"%")
					//} else if i == len-1 {
					//	query = query.Or("account_type > ? and merchant_codes like ?)", adminUser.AccountType, "%"+code+"%")
					//} else {
					query = query.Or("account_type > ? and merchant_codes like ?", adminUser.AccountType, "%"+code+"%")
					//}
				}
				//}
			} else {
				if vo.UserName == adminUser.UserName {
					query = query.Where("merchant_codes = ?", adminUser.MerchantCodes)
				} else {
					len := len(merchantLis)
					if len == 1 {
						query = query.Where("account_type > ? and merchant_codes like ?", adminUser.AccountType, "%"+merchantLis[0]+"%")
					} else {
						for i, code := range merchantLis {
							if i == 0 {
								query = query.Where("(account_type > ? and merchant_codes like ?", adminUser.AccountType, "%"+code+"%")
							} else if i == len-1 {
								query = query.Or("account_type > ? and merchant_codes like ?)", adminUser.AccountType, "%"+code+"%")
							} else {
								query = query.Or("account_type > ? and merchant_codes like ?", adminUser.AccountType, "%"+code+"%")
							}
						}
					}
				}

			}

			return query, true
		} else if adminUser.AccountType == 3 {
			// 普通账号仅显示自己
			return query.Where("username = ?", adminUser.UserName), true
		}

		return query, false
	}

	return query, false
}

func CheckAdminUserSubAccount(c *gin.Context, id string) bool {
	userInfo, ok := c.Get("UserInfo")
	if ok {
		adminUser := userInfo.(*dos.AdminUser)
		if adminUser.AccountType == 1 {
			return true
		}

		/*adminUsers := []dos.AdminUser{}
		err := global.G_DB.Model(&dos.AdminUser{}).Select("username").Where("create_by = ?",
			userInfo.(*dos.AdminUser).UserName).Find(&adminUsers).Error
		if err != nil {
			global.G_LOG.Errorf("[CheckAdminUserSubAccount] Find merchants failed err: %v", err.Error())
			return false
		}

		for _, v := range adminUsers {
			if v.UserName == userName {
				return true
			}
		}*/
		user := dos.AdminUser{}
		err := global.G_DB.Model(&dos.AdminUser{}).Select("account_type", "merchant_codes").Where("id = ?", id).Take(&user).Error
		if err != nil {
			global.G_LOG.Errorf("[CheckAdminUserSubAccount] Find admin user failed: %v", err.Error())
			return false
		}

		if user.AccountType != 3 {
			return false // 商户只能修改子账户
		}

		myMerchantCodeLis := strings.Split(adminUser.MerchantCodes, ",")
		merchantCodeLis := strings.Split(user.MerchantCodes, ",")
		for _, v := range merchantCodeLis {
			if v != "" && slices.Contains(myMerchantCodeLis, v) {
				return true
			}
		}
	}

	return false
}

func GetTokenAdminUser(c *gin.Context) *dos.AdminUser {
	userInfo, ok := c.Get("UserInfo")
	if ok {
		return userInfo.(*dos.AdminUser)
	}

	return nil
}

// UseGoogleMfa - 是否使用谷歌mfa
// @returns bool
func UseGoogleMfa() bool {
	googleMfa := FindByKeyDictsDetailFirst(&dos.DictsDetail{
		DictsTypeCode: "Cilent_System_Settings",
		DictsTag:      "GoogleMfaFlag",
	})

	return googleMfa.DictsValue == "1"
}

func IsWhiteIP(ip string) bool {
	if ip == "127.0.0.1" || ip == "localhost" {
		return true
	}
	//whiteSettingValue, exits := _cache.Get("WhiteIP_Status")
	//if !exits {
	whiteSetting := FindByKeyDictsDetailFirst(&dos.DictsDetail{
		DictsTag: "WhiteIP",
	})
	//_cache.SetDefault("WhiteIP_Status", whiteSetting.DictsValue)
	whiteSettingValue := whiteSetting.DictsValue
	//}
	if whiteSettingValue == "1" {
		//ipValue, ipExits := _cache.Get(ip)
		//if !ipExits {
		data := FindByKeyWhiteIpFirst(&dos.WhiteIp{
			IpAddr: ip,
		})
		ipValue := len(data.Id) == 0
		//_cache.SetDefault(ip, ipValue)
		//}
		if ipValue {
			return false
		}
	}
	return true
}
