// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveAdsCarousel(vo *dos.AdsCarousel) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageAdsCarousel(page, pageSize int, vo *dos.AdsCarousel, c *gin.Context) (ret []*dos.AdsCarouselResp, total int64) {
	query := global.G_DB.Model(&dos.AdsCarousel{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Key) > 0 {
		query = query.Where("`key` like ?", "%"+vo.Name+"%")
	}

	if len(vo.Name) > 0 {
		query = query.Where("name like ?", "%"+vo.Name+"%")
	}

	if vo.Status > -1 {
		query = query.Where("status = ?", vo.Status)
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
	dataSlice := []*dos.AdsCarouselResp{}
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)

	for i := range dataSlice {
		dataSlice[i].SourceNum = strings.Count(dataSlice[i].Sources, ",") + 1
	}

	return dataSlice, count
}

func FindByKeyAdsCarousel(vo *dos.AdsCarousel, c *gin.Context) []*dos.AdsCarouselResp {
	var data []*dos.AdsCarouselResp
	query := global.G_DB.Model(&dos.AdsCarousel{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Key) > 0 {
		query = query.Where("`key` like ?", "%"+vo.Name+"%")
	}

	if len(vo.Name) > 0 {
		query = query.Where("name like ?", "%"+vo.Name+"%")
	}

	if vo.Status > -1 {
		query = query.Where("status = ?", vo.Status)
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

	for i := range data {
		data[i].SourceNum = strings.Count(data[i].Sources, ",") + 1
	}

	return data
}

func FindByKeyAdsCarouselFirst(vo *dos.AdsCarousel) *dos.AdsCarousel {
	var data *dos.AdsCarousel
	query := global.G_DB.Model(&dos.AdsCarousel{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Key) > 0 {
		query = query.Where("`key` like ?", "%"+vo.Name+"%")
	}

	if len(vo.Name) > 0 {
		query = query.Where("name like ?", "%"+vo.Name+"%")
	}

	query.Take(&data)
	return data
}

func UpdateAdsCarousel(vo *dos.AdsCarousel) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"merchant_code": vo.MerchantCode,
		"key":           vo.Key,
		"name":          vo.Name,
		"sort":          vo.Sort,
		//"is_carousel":   vo.IsCarousel,
		"sources":     vo.Sources,
		"update_time": automaticType.Time(time.Now()),
		"update_by":   vo.UpdateBy,
		"status":      vo.Status,
		"jumpto":      vo.Jumpto,
	}).Error == nil
}

func DeleteAdsCarousel(vo *dos.AdsCarousel) bool {
	return global.G_DB.Model(&dos.AdsCarousel{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// GetAdsCarouselInfo - 获取广告栏信息
// @param {string} id
// @returns []dos.GuideInfoResp
func GetAdsCarouselInfo(id string) []dos.AdsCarouselRes {
	data := []dos.AdsCarousel{}
	query := global.G_DB.Model(&dos.AdsCarousel{})
	if len(id) > 0 {
		query = query.Where("id = ?", id)
	}
	err := query.Where("status = 1").Order("sort").Find(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetAdsCarouselInfo] Find ads carousel data failed: %v", err.Error())
		return []dos.AdsCarouselRes{}
	}

	result := []dos.AdsCarouselRes{}
	if len(data) > 0 {
		tool.JsonMapper(data, &result)
	}

	return result
}
