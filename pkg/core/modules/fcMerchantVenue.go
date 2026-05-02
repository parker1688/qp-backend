// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"context"
)

func SaveFcMerchantVenue(vo *dos.FcMerchantVenue) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcMerchantVenue(page, pageSize int, vo *dos.FcMerchantVenue) (ret []*dos.FcMerchantVenue, total int64) {
	query := global.G_DB.Model(&dos.FcMerchantVenue{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.VenueId) > 0 {
		query = query.Where("venue_id = ?", vo.VenueId)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if !vo.MaintainStartTime.Timer().IsZero() {
		query = query.Where("maintain_start_time = ?", vo.MaintainStartTime)
	}

	if !vo.MaintainEndTime.Timer().IsZero() {
		query = query.Where("maintain_end_time = ?", vo.MaintainEndTime)
	}

	if len(vo.ConfigId) > 0 {
		query = query.Where("config_id = ?", vo.ConfigId)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.ConfigAlias) > 0 {
		query = query.Where("config_alias = ?", vo.ConfigAlias)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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

	if vo.InVenueCode > 0 {
		query = query.Where("in_venue_code = ?", vo.InVenueCode)
	}

	if vo.OutVenueCode > 0 {
		query = query.Where("out_venue_code = ?", vo.OutVenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.Describe) > 0 {
		query = query.Where("describe = ?", vo.Describe)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcMerchantVenue
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcMerchantVenue(vo *dos.FcMerchantVenue) []*dos.FcMerchantVenue {
	var data []*dos.FcMerchantVenue
	query := global.G_DB.Model(&dos.FcMerchantVenue{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.VenueId) > 0 {
		query = query.Where("venue_id = ?", vo.VenueId)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if !vo.MaintainStartTime.Timer().IsZero() {
		query = query.Where("maintain_start_time = ?", vo.MaintainStartTime)
	}

	if !vo.MaintainEndTime.Timer().IsZero() {
		query = query.Where("maintain_end_time = ?", vo.MaintainEndTime)
	}

	if len(vo.ConfigId) > 0 {
		query = query.Where("config_id = ?", vo.ConfigId)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.ConfigAlias) > 0 {
		query = query.Where("config_alias = ?", vo.ConfigAlias)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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

	if vo.InVenueCode > 0 {
		query = query.Where("in_venue_code = ?", vo.InVenueCode)
	}

	if vo.OutVenueCode > 0 {
		query = query.Where("out_venue_code = ?", vo.OutVenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.Describe) > 0 {
		query = query.Where("describe = ?", vo.Describe)
	}

	query.Find(&data)
	return data
}

func FindByKeyFcMerchantVenueFirst(vo *dos.FcMerchantVenue) *dos.FcMerchantVenue {
	var data *dos.FcMerchantVenue
	query := global.G_DB.Model(&dos.FcMerchantVenue{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.VenueId) > 0 {
		query = query.Where("venue_id = ?", vo.VenueId)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if !vo.MaintainStartTime.Timer().IsZero() {
		query = query.Where("maintain_start_time = ?", vo.MaintainStartTime)
	}

	if !vo.MaintainEndTime.Timer().IsZero() {
		query = query.Where("maintain_end_time = ?", vo.MaintainEndTime)
	}

	if len(vo.ConfigId) > 0 {
		query = query.Where("config_id = ?", vo.ConfigId)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.ConfigAlias) > 0 {
		query = query.Where("config_alias = ?", vo.ConfigAlias)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
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

	if vo.InVenueCode > 0 {
		query = query.Where("in_venue_code = ?", vo.InVenueCode)
	}

	if vo.OutVenueCode > 0 {
		query = query.Where("out_venue_code = ?", vo.OutVenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.Describe) > 0 {
		query = query.Where("describe = ?", vo.Describe)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcMerchantVenue(vo *dos.FcMerchantVenue) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"merchant_code":       vo.MerchantCode,
		"venue_id":            vo.VenueId,
		"venue_code":          vo.VenueCode,
		"status":              vo.Status,
		"maintain_start_time": vo.MaintainStartTime,
		"maintain_end_time":   vo.MaintainEndTime,
		"config_id":           vo.ConfigId,
		"venue_name":          vo.VenueName,
		"config_alias":        vo.ConfigAlias,
		"currency":            vo.Currency,
		"create_by":           vo.CreateBy,
		"update_by":           vo.UpdateBy,
		"venue_fee_rate":      vo.VenueFeeRate,
		"in_venue_code":       vo.InVenueCode,
		"out_venue_code":      vo.OutVenueCode,
		"game_type":           vo.GameType,
		"img_icon":            vo.ImgIcon,
		"img_bar":             vo.ImgBar,
		"describe":            vo.Describe,
	}).Error == nil
}

func DeleteFcMerchantVenue(vo *dos.FcMerchantVenue) bool {
	return global.G_DB.Model(&dos.FcMerchantVenue{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func SyncMerchantVenueByStatus(venueCode string, status int) error {
	err := global.G_DB.Model(&dos.FcMerchantVenue{}).
		Where("venue_code = ?", venueCode).
		Update("status", status).Error
	if err != nil {
		global.G_LOG.Errorf("[SyncMerchantVenueByStatus] Update merchant venue status failed: %v", err.Error())
	} else {
		global.G_REDIS.Set(context.Background(), "VenuesSyncFlag", "1", -1) // 用于全局刷新场馆数据
	}

	return err
}

func SyncVenueGameByStatus(venueCode string, status int) error {
	err := global.G_DB.Model(&dos.FcVenueGame{}).
		Where("venue_code = ?", venueCode).
		Update("status", status).Error
	if err != nil {
		global.G_LOG.Errorf("[SyncVenueGameByStatus] Update venue game status failed: %v", err.Error())
	}

	return err
}
