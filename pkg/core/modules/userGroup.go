// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"errors"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveUserGroup(vo *dos.UserGroup) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageUserGroup(page, pageSize int, vo *dos.UserGroup, c *gin.Context) (ret []*dos.UserGroup, total int64) {
	query := global.G_DB.Model(&dos.UserGroup{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.GroupName) > 0 {
		query = query.Where("group_name like ?", "%"+vo.GroupName+"%")
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.UserGroup
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyUserGroup(vo *dos.UserGroup, c *gin.Context) []*dos.UserGroup {
	var data []*dos.UserGroup
	query := global.G_DB.Model(&dos.UserGroup{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.GroupName) > 0 {
		query = query.Where("group_name like ?", "%"+vo.GroupName+"%")
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyUserGroupFirst(vo *dos.UserGroup) *dos.UserGroup {
	var data *dos.UserGroup
	query := global.G_DB.Model(&dos.UserGroup{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.GroupName) > 0 {
		query = query.Where("group_name like ?", "%"+vo.GroupName+"%")
	}

	query.Take(&data)
	return data
}

func UpdateUserGroup(vo *dos.UserGroup) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"group_name":  vo.GroupName,
		"data":        vo.Data,
		"update_time": automaticType.Time(time.Now()),
		"update_by":   vo.UpdateBy,
	}).Error == nil
}

func DeleteUserGroup(vo *dos.UserGroup) bool {
	return global.G_DB.Model(&dos.UserGroup{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// CheckUserGroupName - 判断是否存在组名称
// @param {string} groupName 组名称
// @param {string} excludeId 排除id
// @returns bool
func CheckUserGroupName(groupName string, excludeId string) bool {
	data := dos.UserGroup{}
	err := global.G_DB.Model(&dos.UserGroup{}).Select("id").Where("group_name = ?", groupName).First(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[CheckUserGroupName] Query user group count failed: %v", err.Error())
		return false
	}

	if len(excludeId) > 0 {
		return data.Id != excludeId
	}

	return len(data.Id) > 0
}

// GetUserGroupIdsByGroupName - 根据组名称获取组用户id列表
// @param {string} groupName 组名称
// @returns []string
func GetUserGroupIdsByGroupName(groupName string) []string {
	data := dos.UserGroup{}
	err := global.G_DB.Model(&dos.UserGroup{}).Select("data").Where("group_name = ?", groupName).First(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[CheckUserGroupName] Query user group data failed: %v", err.Error())
		return []string{}
	}

	lis := []string{}
	tool.JsonUnmarshal([]byte(data.Data), &lis)

	return lis
}

// CheckUserGroupIdsByMerchantCode - 判断有无不同商户ID
// @param {string} merchantCode 商户code
// @param {string} data 用户id组
// @returns error
func CheckUserGroupIdsByMerchantCode(merchantCode string, data string) error {
	userIds := []string{}
	err := tool.JsonUnmarshal([]byte(data), &userIds)
	if err != nil {
		global.G_LOG.Errorf("[CheckUserGroupIdsByMerchantCode] Json unmarshal failed: data=%s, err=%v", data, err.Error())
		return errors.New("数据格式存在问题: data=" + data + ", err=" + err.Error())
	}

	users := []dos.FcUserMaterial{}
	err = global.G_DB.Model(&dos.FcUserMaterial{}).Select("user_id", "merchant_code").Where("user_id in ?", userIds).Find(&users).Error
	if err != nil {
		global.G_LOG.Errorf("[CheckUserGroupIdsByMerchantCode] Query user material failed: %v", err.Error())
		return err
	}

	if len(userIds) != len(users) {
		nonexistUserIds := []string{}
		userLis := []string{}
		for _, v := range users {
			userLis = append(userLis, v.UserId)
		}

		for _, id := range userIds {
			if !slices.Contains(userLis, id) {
				nonexistUserIds = append(nonexistUserIds, id)
			}
		}

		if len(nonexistUserIds) > 0 {
			return errors.New("有不存在的玩家ID(" + strings.Join(nonexistUserIds, ",") + ")")
		}
	}

	wrongUserIds := []string{}
	for _, v := range users {
		if v.MerchantCode != merchantCode {
			wrongUserIds = append(wrongUserIds, v.UserId)
		}
	}

	if len(wrongUserIds) > 0 {
		return errors.New("存在不同商户玩家ID(" + strings.Join(wrongUserIds, ",") + ")")
	}

	return nil
}
