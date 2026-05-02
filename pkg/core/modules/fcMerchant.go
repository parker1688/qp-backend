package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
	"strings"

	"github.com/gin-gonic/gin"
)

func SaveFcMerchant(vo *dos.FcMerchant) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcMerchant(page, pageSize int, vo *dos.FcMerchant, c *gin.Context) (ret []*dos.FcMerchant, total int64) {
	query := global.G_DB.Model(&dos.FcMerchant{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.Logo) > 0 {
		query = query.Where("logo = ?", vo.Logo)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if len(vo.Prefix) > 0 {
		query = query.Where("prefix = ?", vo.Prefix)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcMerchant
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcMerchant(vo *dos.FcMerchant, c *gin.Context) []*dos.FcMerchant {
	var data []*dos.FcMerchant
	query := global.G_DB.Model(&dos.FcMerchant{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.Logo) > 0 {
		query = query.Where("logo = ?", vo.Logo)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if len(vo.Prefix) > 0 {
		query = query.Where("prefix = ?", vo.Prefix)
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

func FindByKeyFcMerchantFirst(vo *dos.FcMerchant) *dos.FcMerchant {
	var data dos.FcMerchant

	// 先直接使用原生 SQL 查询测试
	var count int64
	global.G_DB.Raw("SELECT COUNT(*) FROM fc_merchant WHERE merchant_code = ? AND status = ?", vo.MerchantCode, vo.Status).Scan(&count)
	global.G_LOG.Infof("[FindByKeyFcMerchantFirst] Direct SQL COUNT: %d for merchant_code=%s, status=%d", count, vo.MerchantCode, vo.Status)

	// 使用 GORM 查询
	query := global.G_DB.Model(&dos.FcMerchant{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.MerchantName) > 0 {
		query = query.Where("merchant_name = ?", vo.MerchantName)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.Logo) > 0 {
		query = query.Where("logo = ?", vo.Logo)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Currency) > 0 {
		query = query.Where("currency = ?", vo.Currency)
	}

	if len(vo.Prefix) > 0 {
		query = query.Where("prefix = ?", vo.Prefix)
	}

	global.G_LOG.Infof("[FindByKeyFcMerchantFirst] Querying with MerchantCode: %s, Status: %d", vo.MerchantCode, vo.Status)

	err := query.Take(&data).Error
	if err != nil {
		global.G_LOG.Infof("[FindByKeyFcMerchantFirst] GORM error: %v", err)
		return nil
	}

	global.G_LOG.Infof("[FindByKeyFcMerchantFirst] Found data: Id=%s, MerchantCode=%s", data.Id, data.MerchantCode)

	return &data
}

// 根据主键Update
func UpdateFcMerchant(vo *dos.FcMerchant) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"merchant_name": vo.MerchantName,
		//// "merchant_code": vo.MerchantCode,
		"logo":     vo.Logo,
		"status":   vo.Status,
		"currency": vo.Currency,
		//"prefix":        vo.Prefix,
	}).Error == nil
}

func DeleteFcMerchant(vo *dos.FcMerchant) bool {
	return global.G_DB.Model(&dos.FcMerchant{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func GetFcMerchantNamesStringByCodes(codes string) string {
	if len(codes) == 0 {
		return ""
	}

	codeLis := strings.Split(codes, ",")

	merchants := []dos.FcMerchant{}
	err := global.G_DB.Model(&dos.FcMerchant{}).Select("merchant_name", "merchant_code").Where("merchant_code in ?", codeLis).Find(&merchants).Error
	if err != nil {
		global.G_LOG.Errorf("[GetFcMerchantNamesByCodes] query fc_merchant err: %v", err.Error())
		return ""
	}

	merchantsMp := map[string]string{}
	for _, v := range merchants {
		merchantsMp[v.MerchantCode] = v.MerchantName
	}

	str := ""
	len := len(codeLis)
	for i, v := range codeLis {
		if val, ok := merchantsMp[v]; ok {
			str = str + val
			if len-1 > i {
				str = str + ","
			}
		} else {
			if len-1 > i {
				str = str + ","
			}
		}
	}

	return str
}

func SyncMerchantCodes(userName string, codes string) string {
	merchants := []dos.FcMerchant{}
	err := global.G_DB.Model(&dos.FcMerchant{}).Select("merchant_code").Find(&merchants).Error
	if err != nil {
		global.G_LOG.Errorf("[SyncMerchantCodes] find merchants failed err: %v", err.Error())
		return codes
	}

	// 是超级管理员直接返回全部
	loginUser := dos.AdminUser{}
	err = global.G_DB.Model(&dos.AdminUser{}).Where("username = ?", userName).Take(&loginUser).Error
	if err != nil {
		global.G_LOG.Errorf("[SyncMerchantCodes] find admin user failed err: %v", err.Error())
		return codes
	}

	merchatLis := []string{}
	merchantMap := map[string]int{}
	for _, v := range merchants {
		merchantMap[v.MerchantCode] = 1
		merchatLis = append(merchatLis, v.MerchantCode)
	}

	if loginUser.AccountType == 1 {
		// 是超级管理员返回全部商户
		err = global.G_DB.Model(&dos.AdminUser{}).Where("id = ?", loginUser.Id).
			Update("merchant_codes", strings.Join(merchatLis, ",")).Error
		if err != nil {
			global.G_LOG.Errorf("[SyncMerchantCodes] Update admin user data failed: %v", err.Error())
		}
		return strings.Join(merchatLis, ",")
	}

	ret := true
	newCodes := []string{}
	oldCodes := strings.Split(codes, ",")
	for _, merchantCode := range oldCodes {
		if _, ok := merchantMap[merchantCode]; ok {
			newCodes = append(newCodes, merchantCode)
		} else {
			// 已不存在该商户
			ret = false
		}
	}

	newCodesStr := strings.Join(newCodes, ",")

	if !ret {
		err = global.G_DB.Model(&dos.AdminUser{}).Where("username = ?", userName).Updates(map[string]any{
			"merchant_codes": newCodesStr,
		}).Error
		if err != nil {
			global.G_LOG.Errorf("[SyncMerchantCodes] update admin user merchant codes failed err: %v", err.Error())
		}
	}

	return newCodesStr
}

// GetMerchantName - 获取商户名
// @param {string} merchantCode
// @returns string
func GetMerchantName(merchantCode string) string {
	data := dos.FcMerchant{}
	err := global.G_DB.Model(&dos.FcMerchant{}).Select("merchant_name").
		Where("merchant_code = ?", merchantCode).First(&data).Error
	if err != nil {
		global.G_LOG.Errorf("[GetMerchantName] Find merchant failed: merchantCode=%s, err=%v",
			merchantCode, err.Error())
		return ""
	}

	return data.MerchantName
}
