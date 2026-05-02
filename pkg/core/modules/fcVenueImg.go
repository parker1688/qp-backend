// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcVenueImg(vo *dos.FcVenueImg) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcVenueImg(page, pageSize int, vo *dos.FcVenueImg) (ret []*dos.FcVenueImg, total int64) {
	query := global.G_DB.Model(&dos.FcVenueImg{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("venue_type like %?%", vo.GameType)
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

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.ImgBar) > 0 {
		query = query.Where("img_bar = ?", vo.ImgBar)
	}

	if len(vo.LinkIcon) > 0 {
		query = query.Where("link_icon = ?", vo.LinkIcon)
	}

	if len(vo.LinkBar) > 0 {
		query = query.Where("link_bar = ?", vo.LinkBar)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcVenueImg
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcVenueImg(vo *dos.FcVenueImg, c *gin.Context) []*dos.FcVenueImg {
	var data []*dos.FcVenueImg
	query := global.G_DB.Model(&dos.FcVenueImg{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("venue_type like %?%", vo.GameType)
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

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.ImgBar) > 0 {
		query = query.Where("img_bar = ?", vo.ImgBar)
	}

	if len(vo.LinkIcon) > 0 {
		query = query.Where("link_icon = ?", vo.LinkIcon)
	}

	if len(vo.LinkBar) > 0 {
		query = query.Where("link_bar = ?", vo.LinkBar)
	}

	if c == nil {
		if len(vo.MerchantCode) > 0 {
			query = query.Where("merchant_code = ?", vo.MerchantCode)
		} else {
			query = query.Where("merchant_code = ?", vo.MerchantCode)
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcVenueImgFirst(vo *dos.FcVenueImg) *dos.FcVenueImg {
	var data *dos.FcVenueImg
	query := global.G_DB.Model(&dos.FcVenueImg{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.VenueName) > 0 {
		query = query.Where("venue_name = ?", vo.VenueName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("venue_type like %?%", vo.GameType)
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

	if len(vo.ImgIcon) > 0 {
		query = query.Where("img_icon = ?", vo.ImgIcon)
	}

	if len(vo.ImgBar) > 0 {
		query = query.Where("img_bar = ?", vo.ImgBar)
	}

	if len(vo.LinkIcon) > 0 {
		query = query.Where("link_icon = ?", vo.LinkIcon)
	}

	if len(vo.LinkBar) > 0 {
		query = query.Where("link_bar = ?", vo.LinkBar)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcVenueImg(vo *dos.FcVenueImg) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"venue_name":    vo.VenueName,
		"venue_code":    vo.VenueCode,
		"game_type":     vo.GameType,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"img_icon":      vo.ImgIcon,
		"img_bar":       vo.ImgBar,
		"link_icon":     vo.LinkIcon,
		"link_bar":      vo.LinkBar,
		"merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcVenueImg(vo *dos.FcVenueImg) bool {
	return global.G_DB.Model(&dos.FcVenueImg{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
