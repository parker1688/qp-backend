// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

const venueGamePriorityOrderExpr = "CASE UPPER(venue_code) " +
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

func SaveFcVenueGame(vo *dos.FcVenueGame) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcVenueGame(page, pageSize int, vo *dos.FcVenueGame) (ret []*dos.FcVenueGame, total int64) {
	query := global.G_DB.Model(&dos.FcVenueGame{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.VenueType) > 0 {
		query = query.Where("venue_type = ?", vo.VenueType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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

	if len(vo.GameCode) > 0 {
		query = query.Where("game_code = ?", vo.GameCode)
	}

	if vo.Hot > 0 {
		query = query.Where("hot = ?", vo.Hot)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.GameName) > 0 {
		query = query.Where("game_name like ?", "%"+vo.GameName)
	}

	if len(vo.GameNamePinyin) > 0 {
		query = query.Where("game_name_pinyin like ?", "%"+vo.GameNamePinyin)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVenueGame
	query.Order(venueGamePriorityOrderExpr).Order("sort asc").Order("game_name asc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVenueGame(vo *dos.FcVenueGame) []*dos.FcVenueGame {
	var data []*dos.FcVenueGame
	query := global.G_DB.Model(&dos.FcVenueGame{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.VenueType) > 0 {
		query = query.Where("venue_type = ?", vo.VenueType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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

	if len(vo.GameCode) > 0 {
		query = query.Where("game_code = ?", vo.GameCode)
	}

	if vo.Hot > 0 {
		query = query.Where("hot = ?", vo.Hot)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.GameName) > 0 {
		query = query.Where("game_name = ?", vo.GameName)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	query.Order("hot,sort desc").Find(&data)
	return data
}

func FindByKeyFcVenueGameFirst(vo *dos.FcVenueGame) *dos.FcVenueGame {
	var data *dos.FcVenueGame
	query := global.G_DB.Model(&dos.FcVenueGame{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
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

	if len(vo.GameCode) > 0 {
		query = query.Where("game_code = ?", vo.GameCode)
	}

	if vo.Hot > 0 {
		query = query.Where("hot = ?", vo.Hot)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.GameName) > 0 {
		query = query.Where("game_name = ?", vo.GameName)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	query.First(&data)
	return data
}

// 根据主键Update
func UpdateFcVenueGame(vo *dos.FcVenueGame) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"venue_name": vo.VenueName,
		"venue_code": vo.VenueCode,
		"venue_type": vo.VenueType,
		"status":     vo.Status,
		"create_by":  vo.CreateBy,
		"update_by":  vo.UpdateBy,
		"game_code":  vo.GameCode,
		"hot":        vo.Hot,
		"sort":       vo.Sort,
		"img_icon":   vo.ImgIcon,
		"game_name":  vo.GameName,
		"language":   vo.Language,
		"game_type":  vo.GameType,
		"gtype":      vo.Gtype,
	}).Error == nil
}

func DeleteFcVenueGame(vo *dos.FcVenueGame) bool {
	return global.G_DB.Model(&dos.FcVenueGame{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
