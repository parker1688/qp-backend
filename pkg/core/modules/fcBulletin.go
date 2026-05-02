// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcBulletin(vo *dos.FcBulletin) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcBulletin(page, pageSize int, vo *dos.FcBulletin, c *gin.Context) (ret []*dos.FcBulletin, total int64) {
	query := global.G_DB.Model(&dos.FcBulletin{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.IsDisplay > 0 {
		query = query.Where("is_display = ?", vo.IsDisplay)
	}

	if vo.BulletinType > 0 {
		query = query.Where("bulletin_type = ?", vo.BulletinType)
	}

	if vo.ContentType > 0 {
		query = query.Where("content_type = ?", vo.ContentType)
	}

	if len(vo.BulletinImg) > 0 {
		query = query.Where("bulletin_img = ?", vo.BulletinImg)
	}

	if !vo.StartTime.Timer().IsZero() {
		query = query.Where("start_time = ?", vo.StartTime)
	}

	if !vo.EndTime.Timer().IsZero() {
		query = query.Where("end_time = ?", vo.EndTime)
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
	var dataSlice []*dos.FcBulletin
	query.Order("sort asc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcBulletin(vo *dos.FcBulletin) []*dos.FcBulletin {
	var data []*dos.FcBulletin
	query := global.G_DB.Model(&dos.FcBulletin{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.IsDisplay > 0 {
		query = query.Where("is_display = ?", vo.IsDisplay)
	}

	if vo.BulletinType > 0 {
		query = query.Where("bulletin_type = ?", vo.BulletinType)
	}

	if vo.ContentType > 0 {
		query = query.Where("content_type = ?", vo.ContentType)
	}

	if len(vo.BulletinImg) > 0 {
		query = query.Where("bulletin_img = ?", vo.BulletinImg)
	}

	if !vo.StartTime.Timer().IsZero() {
		query = query.Where("start_time = ?", vo.StartTime)
	}

	if !vo.EndTime.Timer().IsZero() {
		query = query.Where("end_time = ?", vo.EndTime)
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

	query.Order("sort desc").Find(&data)
	return data
}

func FindByKeyFcBulletin2(vo *dos.FcBulletin, c *gin.Context) []*dos.FcBulletin {
	var data []*dos.FcBulletin
	query := global.G_DB.Model(&dos.FcBulletin{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.IsDisplay > 0 {
		query = query.Where("is_display = ?", vo.IsDisplay)
	}

	if vo.BulletinType > 0 {
		query = query.Where("bulletin_type = ?", vo.BulletinType)
	}

	if vo.ContentType > 0 {
		query = query.Where("content_type = ?", vo.ContentType)
	}

	if len(vo.BulletinImg) > 0 {
		query = query.Where("bulletin_img = ?", vo.BulletinImg)
	}

	if !vo.StartTime.Timer().IsZero() {
		query = query.Where("start_time = ?", vo.StartTime)
	}

	if !vo.EndTime.Timer().IsZero() {
		query = query.Where("end_time = ?", vo.EndTime)
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

	query.Order("sort desc").Find(&data)
	return data
}

func FindByKeyFcBulletinFirst(vo *dos.FcBulletin) *dos.FcBulletin {
	var data *dos.FcBulletin
	query := global.G_DB.Model(&dos.FcBulletin{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if vo.IsDisplay > 0 {
		query = query.Where("is_display = ?", vo.IsDisplay)
	}

	if vo.BulletinType > 0 {
		query = query.Where("bulletin_type = ?", vo.BulletinType)
	}

	if vo.ContentType > 0 {
		query = query.Where("content_type = ?", vo.ContentType)
	}

	if len(vo.BulletinImg) > 0 {
		query = query.Where("bulletin_img = ?", vo.BulletinImg)
	}

	if !vo.StartTime.Timer().IsZero() {
		query = query.Where("start_time = ?", vo.StartTime)
	}

	if !vo.EndTime.Timer().IsZero() {
		query = query.Where("end_time = ?", vo.EndTime)
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

// 根据主键Update
func UpdateFcBulletin(vo *dos.FcBulletin) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"title":         vo.Title,
		"content":       vo.Content,
		"sort":          vo.Sort,
		"is_display":    vo.IsDisplay,
		"bulletin_type": vo.BulletinType,
		"content_type":  vo.ContentType,
		"bulletin_img":  vo.BulletinImg,
		//"start_time":    vo.StartTime,
		//"end_time":      vo.EndTime,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"update_time":   vo.UpdateTime,
		"merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcBulletin(vo *dos.FcBulletin) bool {
	return global.G_DB.Model(&dos.FcBulletin{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
