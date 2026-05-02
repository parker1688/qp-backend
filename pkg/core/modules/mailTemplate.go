// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"time"
)

func SaveMailTemplate(vo *dos.MailTemplate) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageMailTemplate(page, pageSize int, vo *dos.MailTemplate) (ret []*dos.MailTemplate, total int64) {
	query := global.G_DB.Model(&dos.MailTemplate{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	var count int64
	query.Count(&count)
	dataSlice := []*dos.MailTemplate{}
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyMailTemplate(vo *dos.MailTemplate) []*dos.MailTemplate {
	var data []*dos.MailTemplate
	query := global.G_DB.Model(&dos.MailTemplate{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	query.Find(&data)
	return data
}

func FindByKeyMailTemplateFirst(vo *dos.MailTemplate) *dos.MailTemplate {
	var data *dos.MailTemplate
	query := global.G_DB.Model(&dos.MailTemplate{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	query.Take(&data)
	return data
}

func UpdateMailTemplate(vo *dos.MailTemplate) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"title":       vo.Title,
		"content":     vo.Content,
		"update_time": automaticType.Time(time.Now()),
		"update_by":   vo.UpdateBy,
	}).Error == nil
}

func DeleteMailTemplate(vo *dos.MailTemplate) bool {
	return global.G_DB.Model(&dos.MailTemplate{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
