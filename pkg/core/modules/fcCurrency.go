// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcCurrency(vo *dos.FcCurrency) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcCurrency(page, pageSize int, vo *dos.FcCurrency) (ret []*dos.FcCurrency, total int64) {
	query := global.G_DB.Model(&dos.FcCurrency{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Name) > 0 {
		query = query.Where("name = ?", vo.Name)
	}

	if len(vo.Code) > 0 {
		query = query.Where("code = ?", vo.Code)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcCurrency
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcCurrency(vo *dos.FcCurrency) []*dos.FcCurrency {
	var data []*dos.FcCurrency
	query := global.G_DB.Model(&dos.FcCurrency{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Name) > 0 {
		query = query.Where("name = ?", vo.Name)
	}

	if len(vo.Code) > 0 {
		query = query.Where("code = ?", vo.Code)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcCurrencyFirst(vo *dos.FcCurrency) *dos.FcCurrency {
	var data *dos.FcCurrency
	query := global.G_DB.Model(&dos.FcCurrency{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Name) > 0 {
		query = query.Where("name = ?", vo.Name)
	}

	if len(vo.Code) > 0 {
		query = query.Where("code = ?", vo.Code)
	}

	if len(vo.Icon) > 0 {
		query = query.Where("icon = ?", vo.Icon)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcCurrency(vo *dos.FcCurrency) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"name": vo.Name,
		"code": vo.Code,
		"rate": vo.Rate,
		"icon": vo.Icon,
	}).Error == nil
}

func DeleteFcCurrency(vo *dos.FcCurrency) bool {
	return global.G_DB.Model(&dos.FcCurrency{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
