// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

const venuePriorityOrderExpr = "CASE UPPER(venue_code) " +
	"WHEN 'TH' THEN 1 " +
	"WHEN 'PG' THEN 2 " +
	"WHEN 'FG' THEN 3 " +
	"WHEN 'JDB' THEN 4 " +
	"WHEN 'CQ9' THEN 5 " +
	"WHEN 'MGPLUS' THEN 6 " +
	"WHEN 'BBINNEW' THEN 7 " +
	"WHEN 'KYNEW' THEN 8 " +
	"WHEN 'LEYOUNEW' THEN 9 " +
	"WHEN 'LGD' THEN 10 " +
	"WHEN 'BG' THEN 11 " +
	"WHEN 'MW' THEN 12 " +
	"WHEN 'MT' THEN 13 " +
	"WHEN 'WALI' THEN 14 " +
	"WHEN 'JILI' THEN 15 " +
	"WHEN 'KX' THEN 16 " +
	"WHEN 'TT' THEN 17 " +
	"ELSE 999 END"

func SaveFcVenue(vo *dos.FcVenue) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVenue(page, pageSize int, vo *dos.FcVenue) (ret []*dos.FcVenue, total int64) {
	query := global.G_DB.Model(&dos.FcVenue{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.VenueType > 0 {
		query = query.Where("venue_type = ?", vo.VenueType)
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

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVenue
	query.Order(venuePriorityOrderExpr).Order("venue_name asc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVenue(vo *dos.FcVenue) []*dos.FcVenue {
	var data []*dos.FcVenue
	query := global.G_DB.Model(&dos.FcVenue{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.VenueType > 0 {
		query = query.Where("venue_type = ?", vo.VenueType)
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

	query.Find(&data)
	return data
}

func FindByKeyFcVenueFirst(vo *dos.FcVenue) *dos.FcVenue {
	var data *dos.FcVenue
	query := global.G_DB.Model(&dos.FcVenue{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.Language) > 0 {
		query = query.Where("language = ?", vo.Language)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if vo.VenueType > 0 {
		query = query.Where("venue_type = ?", vo.VenueType)
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

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcVenue(vo *dos.FcVenue) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"venue_name": vo.VenueName,
		"venue_code": vo.VenueCode,
		"language":   vo.Language,
		"status":     vo.Status,
		"venue_type": vo.VenueType,
		"currency":   vo.Currency,
		"create_by":  vo.CreateBy,
		"update_by":  vo.UpdateBy,
	}).Error == nil
}

func DeleteFcVenue(vo *dos.FcVenue) bool {
	return global.G_DB.Model(&dos.FcVenue{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func GetFcVenueNameMap(venueCodes []string) map[string]string {
	result := map[string]string{}
	data := []dos.FcVenue{}
	query := global.G_DB.Model(&dos.FcVenue{}).Select("venue_name", "venue_code")
	if len(venueCodes) > 0 {
		query = query.Where("venue_code in ?", venueCodes)
	}
	err := query.Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcVenueNameMap] Find venue names failed: %v", err.Error())
		return result
	}

	for _, v := range data {
		result[v.VenueCode] = v.VenueName
	}

	return result
}

func GetFcVenueList() []dos.FcVenue {
	data := []dos.FcVenue{}
	query := global.G_DB.Model(&dos.FcVenue{}).Select("venue_name", "venue_code")
	err := query.Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcVenueList] Find venue list failed: %v", err.Error())
		return data
	}

	return data
}
