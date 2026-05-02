// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcUserSiteMessage(vo *dos.FcUserSiteMessage) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserSiteMessage(page, pageSize int, vo *dos.FcUserSiteMessage, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcUserSiteMessage, total int64) {
	query := global.G_DB.Model(&dos.FcUserSiteMessage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
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

	if vo.MsgIdType > 0 {
		query = query.Where("msg_id_type = ?", vo.MsgIdType)
	}

	if vo.NotifyType > 0 {
		query = query.Where("notify_type = ?", vo.NotifyType)
	}

	if vo.ReadStatus > 0 {
		query = query.Where("read_status = ?", vo.ReadStatus)
	}

	if vo.DelStatus > 0 {
		query = query.Where("del_status = ?", vo.DelStatus)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserSiteMessage
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserSiteMessage(vo *dos.FcUserSiteMessage, c *gin.Context) []*dos.FcUserSiteMessage {
	var data []*dos.FcUserSiteMessage
	query := global.G_DB.Model(&dos.FcUserSiteMessage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
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

	if vo.ReadStatus > 0 {
		query = query.Where("read_status = ?", vo.ReadStatus)
	}

	if vo.DelStatus > 0 {
		query = query.Where("del_status = ?", vo.DelStatus)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcUserSiteMessageFirst(vo *dos.FcUserSiteMessage) *dos.FcUserSiteMessage {
	var data *dos.FcUserSiteMessage
	query := global.G_DB.Model(&dos.FcUserSiteMessage{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
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

	if vo.ReadStatus > 0 {
		query = query.Where("read_status = ?", vo.ReadStatus)
	}

	if vo.DelStatus > 0 {
		query = query.Where("del_status = ?", vo.DelStatus)
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

// 根据主键Update
func UpdateFcUserSiteMessage(vo *dos.FcUserSiteMessage) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":       vo.UserId,
		"user_name":     vo.UserName,
		"title":         vo.Title,
		"content":       vo.Content,
		"msg_type":      vo.MsgType,
		"notify_type":   vo.NotifyType,
		"read_status":   vo.ReadStatus,
		"del_status":    vo.DelStatus,
		"language":      vo.Language,
		"merchant_code": vo.MerchantCode,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
	}).Error == nil
}

func DeleteFcUserSiteMessage(vo *dos.FcUserSiteMessage) bool {
	return global.G_DB.Model(&dos.FcUserSiteMessage{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
