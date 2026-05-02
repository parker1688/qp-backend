// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveSplashScreen(vo *dos.SplashScreen) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageSplashScreen(page, pageSize int, vo *dos.SplashScreen, c *gin.Context) (ret []*dos.SplashScreenEx, total int64) {
	query := global.G_DB.Model(&dos.SplashScreen{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
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
	dataSlice := []*dos.SplashScreenEx{}
	query.Offset((page - 1) * pageSize).Limit(pageSize).Preload("Merchant").Find(&dataSlice)
	return dataSlice, count
}

func FindByKeySplashScreen(vo *dos.SplashScreen, c *gin.Context) []*dos.SplashScreenEx {
	var data []*dos.SplashScreenEx
	query := global.G_DB.Model(&dos.SplashScreen{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
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

	query.Preload("Merchant").Find(&data)
	return data
}

func FindByKeySplashScreenFirst(vo *dos.SplashScreen) *dos.SplashScreen {
	var data *dos.SplashScreen
	query := global.G_DB.Model(&dos.SplashScreen{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

func UpdateSplashScreen(vo *dos.SplashScreen) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"logo_img":    vo.LogoImg,
		"banner_img":  vo.BannerImg,
		"screen_img":  vo.ScreenImg,
		"update_time": automaticType.Time(time.Now()),
		"update_by":   vo.UpdateBy,
	}).Error == nil
}

func DeleteSplashScreen(vo *dos.SplashScreen) bool {
	return global.G_DB.Model(&dos.SplashScreen{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

// CheckSplashScreenMerchant - 判断是否已存在商户
// @param {string} merchantCode 商户码
// @param {string} excludeId 排除id
// @returns bool
func CheckSplashScreenMerchant(merchantCode string, excludeId string) bool {
	data := dos.SplashScreen{}
	err := global.G_DB.Model(&dos.SplashScreen{}).Select("id").Where("merchant_code = ?", merchantCode).First(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[CheckSplashScreenMerchant] Query splash screen data failed: %v", err.Error())
		return false
	}

	if len(excludeId) > 0 {
		return data.Id != excludeId
	}

	return len(data.Id) > 0
}

// GetSplashScreenByMerchant - 根据商户码获取开屏数据
// @param {string} merchantCode
// @returns dos.SplashScreen, error
func GetSplashScreenByMerchant(merchantCode string) (dos.SplashScreen, error) {
	data := dos.SplashScreen{}
	err := global.G_DB.Model(&dos.SplashScreen{}).
		Where("merchant_code = ?", merchantCode).Take(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
