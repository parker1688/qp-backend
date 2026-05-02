// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcSiteMessage(vo *dos.FcSiteMessage) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcSiteMessage(page, pageSize int, vo *dos.FcSiteMessage, pageTimeQuery *response.PageTimeQuery) (ret []*dos.FcSiteMessage, total int64) {
	query := global.G_DB.Model(&dos.FcSiteMessage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.MsgType > 0 {
		query = query.Where("msg_type = ?", vo.MsgType)
	}

	if vo.NotifyType > 0 {
		query = query.Where("notify_type = ?", vo.NotifyType)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	if pageTimeQuery.StartAt != "" {
		query = query.Where("create_time >= ?", vo.CreateTime)
	}
	if pageTimeQuery.EndAt != "" {
		query = query.Where("create_time <= ?", vo.CreateTime)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcSiteMessage
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcSiteMessage(vo *dos.FcSiteMessage) []*dos.FcSiteMessage {
	var data []*dos.FcSiteMessage
	query := global.G_DB.Model(&dos.FcSiteMessage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.MsgType > 0 {
		query = query.Where("msg_type = ?", vo.MsgType)
	}

	if vo.NotifyType > 0 {
		query = query.Where("notify_type = ?", vo.NotifyType)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	query.Find(&data)
	return data
}

func FindByKeyFcSiteMessageFirst(vo *dos.FcSiteMessage) *dos.FcSiteMessage {
	var data *dos.FcSiteMessage
	query := global.G_DB.Model(&dos.FcSiteMessage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.MsgType > 0 {
		query = query.Where("msg_type = ?", vo.MsgType)
	}

	if vo.NotifyType > 0 {
		query = query.Where("notify_type = ?", vo.NotifyType)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	query.Take(&data)
	return data
}

//根据主键Update
func UpdateFcSiteMessage(vo *dos.FcSiteMessage) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"title":         vo.Title,
		"content":       vo.Content,
		"msg_type":      vo.MsgType,
		"notify_type":   vo.NotifyType,
		"language":      vo.Language,
		"merchant_code": vo.MerchantCode,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
	}).Error == nil
}

func DeleteFcSiteMessage(vo *dos.FcSiteMessage) bool {
	return global.G_DB.Model(&dos.FcSiteMessage{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
