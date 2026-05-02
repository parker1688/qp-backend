// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcPromotionInfo(vo *dos.FcPromotionInfo) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcPromotionInfo(page, pageSize int, vo *dos.FcPromotionInfo, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcPromotionInfo, total int64) {
	query := global.G_DB.Model(&dos.FcPromotionInfo{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.PromotionType > 0 {
		query = query.Where("promotion_type = ?", vo.PromotionType)
	} else {
		query = query.Where("promotion_type != 10") //DV要求排除掉
	}

	if vo.Title != "" {
		query = query.Where("title like ?", "%"+vo.Title+"%")
	}

	if vo.GameType > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.PromotionImg) > 0 {
		query = query.Where("promotion_img = ?", vo.PromotionImg)
	}

	if !vo.StartTime.Timer().IsZero() {
		query = query.Where("start_time = ?", vo.StartTime)
	}

	if !vo.EndTime.Timer().IsZero() {
		query = query.Where("end_time = ?", vo.EndTime)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
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

	if pageTimeQuery.StartAt != "" {
		query = query.Where("start_time >= ?", pageTimeQuery.StartAt)
	}
	if pageTimeQuery.EndAt != "" {
		query = query.Where("start_time <= ?", pageTimeQuery.EndAt)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.ClientType) > 0 {
		query = query.Where("client_type = ?", vo.ClientType)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
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
	var dataSlice []*dos.FcPromotionInfo
	query.Order("sort asc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcPromotionInfo(vo *dos.FcPromotionInfo, c *gin.Context) []*dos.FcPromotionInfo {
	var data []*dos.FcPromotionInfo
	query := global.G_DB.Model(&dos.FcPromotionInfo{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.PromotionType > 0 {
		query = query.Where("promotion_type = ?", vo.PromotionType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.GameType > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.PromotionImg) > 0 {
		query = query.Where("promotion_img = ?", vo.PromotionImg)
	}

	if !vo.StartTime.Timer().IsZero() {
		query = query.Where("start_time = ?", vo.StartTime)
	}

	if !vo.EndTime.Timer().IsZero() {
		query = query.Where("end_time = ?", vo.EndTime)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
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

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.ClientType) > 0 {
		query = query.Where("client_type = ?", vo.ClientType)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
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

func FindByKeyFcPromotionInfoFirst(vo *dos.FcPromotionInfo) *dos.FcPromotionInfo {
	var data *dos.FcPromotionInfo
	query := global.G_DB.Model(&dos.FcPromotionInfo{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if vo.PromotionType > 0 {
		query = query.Where("promotion_type = ?", vo.PromotionType)
	}

	if vo.GameType > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.PromotionImg) > 0 {
		query = query.Where("promotion_img = ?", vo.PromotionImg)
	}

	if !vo.StartTime.Timer().IsZero() {
		query = query.Where("start_time = ?", vo.StartTime)
	}

	if !vo.EndTime.Timer().IsZero() {
		query = query.Where("end_time = ?", vo.EndTime)
	}

	if len(vo.Link) > 0 {
		query = query.Where("link = ?", vo.Link)
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

	if len(vo.Content) > 0 {
		query = query.Where("content = ?", vo.Content)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.ClientType) > 0 {
		query = query.Where("client_type = ?", vo.ClientType)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcPromotionInfo(vo *dos.FcPromotionInfo) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"promotion_type":         vo.PromotionType,
		"game_type":              vo.GameType,
		"status":                 vo.Status,
		"promotion_img":          vo.PromotionImg,
		"h5_img":                 vo.H5Img,
		"start_time":             vo.StartTime,
		"end_time":               vo.EndTime,
		"link":                   vo.Link,
		"h5_link":                vo.H5Link,
		"update_by":              vo.UpdateBy,
		"update_time":            vo.UpdateTime,
		"sort":                   vo.Sort,
		"content":                vo.Content,
		"client_type":            vo.ClientType,
		"language":               vo.Language,
		"merchant_code":          vo.MerchantCode,
		"title":                  vo.Title,
		"h5_content":             vo.H5Content,
		"stage_content":          vo.StageContent,
		"gift_style":             vo.GiftStyle,
		"recharge_balance_ratio": vo.RechargeBalanceRatio,
		"balance":                vo.Balance,
		"first_recharge_amount":  vo.FirstRechargeAmount,
		"bonus_amount":           vo.BonusAmount,
		"reg_start_time":         vo.RegStartTime,
		"reg_end_time":           vo.RegEndTime,
		"cycle":                  vo.Cycle,
		"date_range_data":        vo.DateRangeData,
	}).Error == nil
}

func DeleteFcPromotionInfo(vo *dos.FcPromotionInfo) bool {
	return global.G_DB.Model(&dos.FcPromotionInfo{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func UpdateFcPromotionInfoContent(vo *dos.FcPromotionInfo) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"content":   vo.Content,
		"update_by": vo.UpdateBy,
	}).Error == nil
}
