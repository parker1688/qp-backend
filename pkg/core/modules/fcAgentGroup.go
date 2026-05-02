// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcAgentGroup(vo *dos.FcAgentGroup) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcAgentGroup(page, pageSize int, vo *dos.FcAgentGroup, merchantCodes string) (ret []*dos.FcAgentGroup, total int64) {
	query := global.G_DB.Model(&dos.FcAgentGroup{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.InviteCode > 0 {
		query = query.Where("invite_code = ?", vo.InviteCode)
	}
	merchantCodes2 := tool.StrToListForSql(merchantCodes)
	query = query.Where("merchant_code in (?)", merchantCodes2)
	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcAgentGroup
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcAgentGroup(vo *dos.FcAgentGroup, merchantCodes string) []*dos.FcAgentGroup {
	var data []*dos.FcAgentGroup
	query := global.G_DB.Model(&dos.FcAgentGroup{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.InviteCode > 0 {
		query = query.Where("invite_code = ?", vo.InviteCode)
	}

	if len(vo.GroupName) > 0 {
		query = query.Where("group_name = ?", vo.GroupName)
	}
	merchantCodes2 := tool.StrToListForSql(merchantCodes)
	query = query.Where("merchant_code in (?)", merchantCodes2)

	query.Find(&data)
	return data
}

// 根据主键Update
func UpdateFcAgentGroup(vo *dos.FcAgentGroup, merchantCodes string) bool {
	merchantCodes2 := tool.StrToListForSql(merchantCodes)
	return global.G_DB.Model(vo).Where(`id = ? and merchant_code in (?)`, vo.Id, merchantCodes2).Updates(map[string]interface{}{
		"invite_code": vo.InviteCode,
		"group_name":  vo.GroupName,
	}).Error == nil
}

func DeleteFcAgentGroup(vo *dos.FcAgentGroup, merchantCodes string) bool {
	merchantCodes2 := tool.StrToListForSql(merchantCodes)
	return global.G_DB.Model(&dos.FcAgentGroup{}).Where("id = ? and merchant_codes in (?)", vo.Id, merchantCodes2).Delete(vo).Error == nil
}
