// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcNotifyTemplate(vo *dos.FcNotifyTemplate) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcNotifyTemplate(page, pageSize int, vo *dos.FcNotifyTemplate, pageTimeQuery *response.PageTimeQuery) (ret []*dos.FcNotifyTemplate, total int64) {
	query := global.G_DB.Model(&dos.FcNotifyTemplate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.TemplateContent) > 0 {
		query = query.Where("template_content = ?", vo.TemplateContent)
	}

	if len(vo.NotifyFlag) > 0 {
		query = query.Where("notify_flag = ?", vo.NotifyFlag)
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

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if pageTimeQuery.StartAt != "" {
		query = query.Where("create_time >= ?", vo.CreateTime)
	}
	if pageTimeQuery.EndAt != "" {
		query = query.Where("create_time <= ?", vo.CreateTime)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcNotifyTemplate
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcNotifyTemplate(vo *dos.FcNotifyTemplate) []*dos.FcNotifyTemplate {
	var data []*dos.FcNotifyTemplate
	query := global.G_DB.Model(&dos.FcNotifyTemplate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.TemplateContent) > 0 {
		query = query.Where("template_content = ?", vo.TemplateContent)
	}

	if len(vo.NotifyFlag) > 0 {
		query = query.Where("notify_flag = ?", vo.NotifyFlag)
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

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcNotifyTemplateFirst(vo *dos.FcNotifyTemplate) *dos.FcNotifyTemplate {
	var data *dos.FcNotifyTemplate
	query := global.G_DB.Model(&dos.FcNotifyTemplate{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.TemplateContent) > 0 {
		query = query.Where("template_content = ?", vo.TemplateContent)
	}

	if len(vo.NotifyFlag) > 0 {
		query = query.Where("notify_flag = ?", vo.NotifyFlag)
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

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

//根据主键Update
func UpdateFcNotifyTemplate(vo *dos.FcNotifyTemplate) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"title":            vo.Title,
		"template_content": vo.TemplateContent,
		"notify_flag":      vo.NotifyFlag,
		"create_by":        vo.CreateBy,
		"update_by":        vo.UpdateBy,
		"merchant_code":    vo.MerchantCode,
	}).Error == nil
}

func DeleteFcNotifyTemplate(vo *dos.FcNotifyTemplate) bool {
	return global.G_DB.Model(&dos.FcNotifyTemplate{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
