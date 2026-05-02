// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcWelfareManage(vo *dos.FcWelfareManage) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcWelfareManage(page, pageSize int, vo *dos.FcWelfareManage, c *gin.Context) (ret []*dos.FcWelfareManage, total int64) {
	query := global.G_DB.Model(&dos.FcWelfareManage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.FlowMultiple > 0 {
		query = query.Where("flow_multiple = ?", vo.FlowMultiple)
	}

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	/*if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}*/

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcWelfareManage
	query.Offset((page - 1) * pageSize).Order("sort desc").Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcWelfareManage(vo *dos.FcWelfareManage, c *gin.Context) []*dos.FcWelfareManage {
	var data []*dos.FcWelfareManage
	query := global.G_DB.Model(&dos.FcWelfareManage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.FlowMultiple > 0 {
		query = query.Where("flow_multiple = ?", vo.FlowMultiple)
	}

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	/*if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}*/

	query.Order("sort desc").Find(&data)
	return data
}

func FindByKeyFcWelfareManageFirst(vo *dos.FcWelfareManage) *dos.FcWelfareManage {
	var data *dos.FcWelfareManage
	query := global.G_DB.Model(&dos.FcWelfareManage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.FlowMultiple > 0 {
		query = query.Where("flow_multiple = ?", vo.FlowMultiple)
	}

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcWelfareManage(vo *dos.FcWelfareManage) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"flow_multiple": vo.FlowMultiple,
		"bonus_type":    vo.BonusType,
		"title":         vo.Title,
		"merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcWelfareManage(vo *dos.FcWelfareManage) bool {
	return global.G_DB.Model(&dos.FcWelfareManage{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// GetFcWelfareTitleByBonusType - 根据交易工单类型获取福利标题
// @param {int} bonusType 交易工单类型
func GetFcWelfareTitleByBonusType(bonusType int) string {
	data := dos.FcWelfareManage{}
	err := global.G_DB.Model(&dos.FcWelfareManage{}).Select("title").
		Where("bonus_type = ?", bonusType).First(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcWelfareTitleByBonusType] Find welfare manage data failed: bonusType=%d, err=%v",
			bonusType, err.Error())
		return ""
	}

	return data.Title
}

func GetFcWelfareTitles(kw string) []string {
	result := []string{}

	data := []dos.FcWelfareManage{}
	err := global.G_DB.Model(&dos.FcWelfareManage{}).Select("title").
		Where("title like ?", "%"+kw+"%").Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcWelfareTitleByBonusType] Find welfare manage data failed: kw=%s, err=%v",
			kw, err.Error())
		return result
	}

	for _, v := range data {
		result = append(result, v.Title)
	}

	return result
}
