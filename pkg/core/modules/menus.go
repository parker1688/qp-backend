// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"bootpkg/pkg/core/modules/enmus"
	"context"
	"time"

	"github.com/kirinlabs/utils"
)

func SaveMenus(vo *dos.Menus) (bool, ecode.Code) {
	global.G_REDIS.Del(context.Background(), enmus.REDIS_TABLE_MENUS).Val()
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageMenus(page, pageSize int, vo *dos.Menus) (ret []*dos.Menus, total int64) {
	query := global.G_DB.Model(&dos.Menus{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MenuName) > 0 {
		query = query.Where("menu_name = ?", vo.MenuName)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.RoleFlag) > 0 {
		query = query.Where("role_flag = ?", vo.RoleFlag)
	}

	if len(vo.Address) > 0 {
		query = query.Where("address = ?", vo.Address)
	}

	if len(vo.ParentId) > 0 {
		query = query.Where("parentId = ?", vo.ParentId)
	}

	if vo.Type > 0 {
		query = query.Where("type = ?", vo.Type)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.Menus
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyMenus(vo *dos.Menus) []*dos.Menus {
	var data []*dos.Menus
	query := global.G_DB.Model(&dos.Menus{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MenuName) > 0 {
		query = query.Where("menu_name = ?", vo.MenuName)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.RoleFlag) > 0 {
		query = query.Where("role_flag = ?", vo.RoleFlag)
	}

	if len(vo.Address) > 0 {
		query = query.Where("address = ?", vo.Address)
	}

	if len(vo.ParentId) > 0 {
		query = query.Where("parentId = ?", vo.ParentId)
	}

	if vo.Type > 0 {
		query = query.Where("type = ?", vo.Type)
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
func UpdateMenus(vo *dos.Menus) bool {
	global.G_REDIS.Del(context.Background(), enmus.REDIS_TABLE_MENUS).Val()
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"menu_name":   vo.MenuName,
		"icon":        vo.Icon,
		"sort":        vo.Sort,
		"role_flag":   vo.RoleFlag,
		"address":     vo.Address,
		"parent_id":   vo.ParentId,
		"type":        vo.Type,
		"locales":     vo.Locales,
		"update_by":   vo.UpdateBy,
		"show_status": vo.ShowStatus,
		"open_cache":  vo.OpenCache,
		"perms":       vo.Perms,
		"api_regular": vo.ApiRegular,
	}).Error == nil
}

func DeleteMenus(vo *dos.Menus) bool {
	global.G_REDIS.Del(context.Background(), enmus.REDIS_TABLE_MENUS).Val()
	return global.G_DB.Model(&dos.Menus{}).Where("id = ?", vo.Id).Delete(vo).Error != nil
}

func FindMenusAll() []*dos.Menus {
	var dataSlice []*dos.Menus

	global.G_LOG.Infof("[FindMenusAll] === START === Checking global.G_DB is nil? %v", global.G_DB == nil)

	val := global.G_REDIS.Get(context.Background(), enmus.REDIS_TABLE_MENUS).Val()
	if len(val) > 5 && tool.JsonUnmarshalFromString(val, &dataSlice) == nil {
		global.G_LOG.Infof("[FindMenusAll] from Redis, count=%d", len(dataSlice))
		return dataSlice
	}

	// 直接使用原生 SQL 测试
	var count int64
	err := global.G_DB.Raw("SELECT COUNT(*) FROM menu").Scan(&count).Error
	global.G_LOG.Infof("[FindMenusAll] Direct SQL COUNT: %d, error: %v", count, err)

	// GORM 查询
	err = global.G_DB.Model(&dos.Menus{}).Order("sort desc").Find(&dataSlice).Error
	global.G_LOG.Infof("[FindMenusAll] GORM error: %v, count: %d", err, len(dataSlice))

	if len(dataSlice) > 0 {
		global.G_LOG.Infof("[FindMenusAll] First menu: ID=%s, Name=%s", dataSlice[0].Id, dataSlice[0].Name)
	}

	global.G_REDIS.Set(context.Background(), enmus.REDIS_TABLE_MENUS, utils.Json(dataSlice), 2*time.Hour).Val()
	return dataSlice
}

func GetMenusMapByRoleMenusIds(menusIds []string) map[string]*dos.Menus {
	data := []*dos.Menus{}
	global.G_DB.Model(&dos.Menus{}).Select("id", "parent_id").Where("id IN ?", menusIds).Find(&data)
	var parentIds []string
	for _, v := range data {
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

	return parentMenusMp
}
