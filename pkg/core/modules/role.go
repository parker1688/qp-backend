// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kirinlabs/utils"
	"gorm.io/gorm"
)

func SaveRole(vo *dos.Role) (bool, ecode.Code) {
	global.G_REDIS.Del(context.Background(), enmus.REDIS_TABLE_ROLE)
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageRole(page, pageSize int, vo *dos.Role, c *gin.Context) (ret []*dos.Role, total int64) {
	query := global.G_DB.Model(&dos.Role{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.RoleName) > 0 {
		query = query.Where("name = ?", vo.RoleName)
	}

	if len(vo.MeusIds) > 0 {
		query = query.Where("meus_ids = ?", vo.MeusIds)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remark = ?", vo.Remarks)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	ok := true
	query, ok = doAdminUserRolePermsQuery(c, query)
	if !ok {
		return ret, total
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.Role
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyRole(vo *dos.Role) []*dos.Role {
	var data []*dos.Role
	query := global.G_DB.Model(&dos.Role{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.RoleName) > 0 {
		query = query.Where("name = ?", vo.RoleName)
	}

	if len(vo.MeusIds) > 0 {
		query = query.Where("meus_ids = ?", vo.MeusIds)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remark = ?", vo.Remarks)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	query.Find(&data)
	return data
}

// 根据主键Update
func UpdateRole(vo *dos.Role) bool {
	global.G_REDIS.Del(context.Background(), enmus.REDIS_TABLE_ROLE)
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"name":      vo.RoleName,
		"meus_ids":  vo.MeusIds,
		"status":    vo.Status,
		"remark":    vo.Remarks,
		"create_by": vo.CreateBy,
		"update_by": vo.UpdateBy,
	}).Error == nil
}

func DeleteRole(vo *dos.Role) bool {
	global.G_REDIS.Del(context.Background(), enmus.REDIS_TABLE_ROLE)
	return global.G_DB.Delete(vo).Error != nil
}

func FindRoleAll() []*dos.Role {
	var dataSlice []*dos.Role
	val := global.G_REDIS.Get(context.Background(), enmus.REDIS_TABLE_ROLE).Val()
	if len(val) > 5 && global.JSON.UnmarshalFromString(val, &dataSlice) == nil {
		return dataSlice
	}
	global.G_DB.Model(&dos.Role{}).Find(&dataSlice)
	global.G_REDIS.Set(context.Background(), enmus.REDIS_TABLE_ROLE, utils.Json(dataSlice), 2*time.Hour).Val()
	return dataSlice
}

func UpdateRoleMenus(vo *dos.Role) bool {
	global.G_REDIS.Del(context.Background(), enmus.REDIS_TABLE_ROLE)
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"meus_ids":   vo.MeusIds,
		"perms_list": vo.PermsList,
	}).Error == nil
}

func doAdminUserRolePermsQuery(c *gin.Context, query *gorm.DB) (*gorm.DB, bool) {
	userInfo, ok := c.Get("UserInfo")
	if ok {
		adminUser := userInfo.(*dos.AdminUser)
		if adminUser.AccountType == 1 {
			return query, true
		} else if adminUser.AccountType == 2 {
			return query, true
		}
	}

	return query, false
}

// 判断角色名是否重复
func CheckRoleNameRepeat(roleName string) bool {
	data := dos.Role{}
	err := global.G_DB.Model(&dos.Role{}).Select("id").Where("role_name = ?", roleName).Take(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[CheckRoleNameRepeat] Can't query role data err: %v", err.Error())
		return false
	}

	return len(data.Id) > 0
}
