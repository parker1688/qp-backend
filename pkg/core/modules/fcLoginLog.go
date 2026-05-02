// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcLoginLog(vo *dos.FcLoginLog) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcLoginLog(page, pageSize int, vo *dos.FcLoginLog, pageTimeQuery *response.PageTimeQuery) (ret []*dos.FcLoginLog, total int64) {
	query := global.G_DB.Model(&dos.FcLoginLog{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if len(vo.ClientType) > 0 {
		query = query.Where("client_type = ?", vo.ClientType)
	}

	if len(vo.Version) > 0 {
		query = query.Where("version = ?", vo.Version)
	}

	if len(vo.VisitorId) > 0 {
		query = query.Where("visitor_id = ?", vo.VisitorId)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(pageTimeQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", pageTimeQuery.StartAt)
	}

	if len(pageTimeQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", pageTimeQuery.EndAt)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcLoginLog
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcLoginLog(vo *dos.FcLoginLog) []*dos.FcLoginLog {
	var data []*dos.FcLoginLog
	query := global.G_DB.Model(&dos.FcLoginLog{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if len(vo.ClientType) > 0 {
		query = query.Where("client_type = ?", vo.ClientType)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

func FindByKeyFcLoginLogFirst(vo *dos.FcLoginLog) *dos.FcLoginLog {
	var data *dos.FcLoginLog
	query := global.G_DB.Model(&dos.FcLoginLog{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Ip) > 0 {
		query = query.Where("ip = ?", vo.Ip)
	}

	if len(vo.ClientType) > 0 {
		query = query.Where("client_type = ?", vo.ClientType)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.First(&data)
	return data
}

func FindByUserIdsFcLoginLog(userIdArr []string) []*dos.FcLoginLog {
	var data []*dos.FcLoginLog
	query := global.G_DB.Model(&dos.FcLoginLog{})

	if len(userIdArr) > 0 {
		query = query.Group("user_id").Having("user_id = ?", userIdArr).Order("create_time desc")
	}

	query.Find(&data)
	return data
}

// 根据主键Update
func UpdateFcLoginLog(vo *dos.FcLoginLog) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":     vo.UserId,
		"user_name":   vo.UserName,
		"ip":          vo.Ip,
		"client_type": vo.ClientType,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcLoginLog(vo *dos.FcLoginLog) bool {
	return global.G_DB.Model(&dos.FcLoginLog{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
