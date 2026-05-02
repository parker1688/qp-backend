// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcTranscation(vo *dos.FcTranscation) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcTranscation(page, pageSize int, vo *dos.FcTranscation, pageQuery response.PageTimeQuery) (ret []*dos.FcTranscation, total int64) {
	page, pageSize = response.NormalizePage(page, pageSize)
	query := global.G_DB.Model(&dos.FcTranscation{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if vo.FundingType > -1 {
		query = query.Where("funding_type = ?", vo.FundingType)
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

	if len(pageQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", pageQuery.StartAt)
	}

	if len(pageQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", pageQuery.EndAt)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcTranscation
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcTranscation(vo *dos.FcTranscation) []*dos.FcTranscation {
	var data []*dos.FcTranscation
	query := global.G_DB.Model(&dos.FcTranscation{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if vo.FundingType > 0 {
		query = query.Where("funding_type = ?", vo.FundingType)
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
func UpdateFcTranscation(vo *dos.FcTranscation) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":            vo.UserId,
		"user_name":          vo.UserName,
		"status":             vo.Status,
		"amount":             vo.Amount,
		"transcation_before": vo.TranscationBefore,
		"transcation_after":  vo.TranscationAfter,
		"remark":             vo.Remark,
		"funding_type":       vo.FundingType,
		"create_by":          vo.CreateBy,
		"update_by":          vo.UpdateBy,
		"merchant_code":      vo.MerchantCode,
	}).Error == nil
}

func DeleteFcTranscation(vo *dos.FcTranscation) bool {
	return global.G_DB.Model(&dos.FcTranscation{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
