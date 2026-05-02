// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"github.com/gin-gonic/gin"
)

func SaveOpRecord(vo *dos.OpRecord) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageOpRecord(page, pageSize int, vo *dos.OpRecord, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.OpRecord, total int64) {
	query := global.G_DB.Model(&dos.FcBetRecord{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Menu1) > 0 {
		query = query.Where("menu1 = ?", vo.Menu1)
	}

	if len(vo.Menu2) > 0 {
		query = query.Where("menu2 = ?", vo.Menu2)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(pageTimeQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", pageTimeQuery.StartAt)
	}
	if len(pageTimeQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", pageTimeQuery.EndAt)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.OpRecord
	query.Order("bet_time desc, settlement_time").Offset((page - 1) * pageSize).Limit(pageSize).Preload("Merchant").Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyOpRecord(vo *dos.OpRecord, c *gin.Context) []*dos.OpRecord {
	var data []*dos.OpRecord
	query := global.G_DB.Model(&dos.OpRecord{})
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

	query.Order("bet_time desc, settlement_time").Find(&data)
	return data
}

// 根据主键Update
func UpdateOpRecord(vo *dos.OpRecord) bool {
	return global.G_DB.Model(vo).Updates(map[string]interface{}{}).Error == nil
}

func DeleteopRecord(vo *dos.OpRecord) bool {
	return global.G_DB.Model(&dos.OpRecord{}).Delete(vo).Error == nil
}

func AddOpRecord(userId, optor, merchantCode, menu1, menu2, ip, opAction, result string) (bool, error) {
	rec := dos.OpRecord{
		UserId:       userId,
		UserName:     optor,
		MerchantCode: merchantCode,
		Menu1:        menu1,
		Menu2:        menu2,
		IP:           ip,
		Op:           opAction,
		Result:       result,
		CreateTime:   automaticType.Now(),
	}
	return SaveOpRecord(&rec)
}
