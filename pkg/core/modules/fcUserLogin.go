// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"errors"

	"gorm.io/gorm"
)

var (
	GKEY_USER_ID_INCR     = "USER_ID_INCR"
	GKEY_INVITE_CODE_INCR = "INVITE_CODE_INCR"
)

func GetNextIdGeneral(key string) (int64, error) {
	newUserId := 0
	err := global.G_DB.Transaction(func(tx *gorm.DB) error {
		data := dos.FcGlobal{}
		err := tx.Model(&dos.FcGlobal{}).Select("value").Where("`key` = ?", key).First(&data).Error
		if err != nil {
			global.G_LOG.Errorf("[GetNextIdGeneral] Find %s sequeue failed: %v", key, err.Error())
			return errors.New("无法获取ID")
		}

		newUserId = tool.Atoi(data.Value) + 1

		err = tx.Model(&dos.FcGlobal{}).Where("`key` = ?", key).Update("value", newUserId).Error
		if err != nil {
			global.G_LOG.Errorf("[GetNextIdGeneral] Update %s sequeue failed: err=%v", key, err.Error())
			return errors.New("无法获取ID")
		}

		return nil
	})

	if newUserId == 0 {
		return 0, errors.New("无法获取ID")
	}

	if err != nil {
		global.G_LOG.Errorf("[GetNextIdGeneral] Update user id sequeue failed: newUserId=%d %v",
			newUserId, err.Error())
		return 0, err
	}

	return int64(newUserId), nil
}

func SaveFcUserLogin(vo *dos.FcUserLogin) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcUserLogin(page, pageSize int, vo *dos.FcUserLogin, pageTimeQuery *response.PageTimeQuery) (ret []*dos.FcUserLogin, total int64) {
	query := global.G_DB.Model(&dos.FcUserLogin{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Password) > 0 {
		query = query.Where("password = ?", vo.Password)
	}

	if !vo.CreateTime.Timer().IsZero() {
		query = query.Where("create_time = ?", vo.CreateTime)
	}

	if len(vo.CreateBy) > 0 {
		query = query.Where("create_by = ?", vo.CreateBy)
	}

	if len(pageTimeQuery.StartAt) > 0 {
		query = query.Where("create_time >= ?", pageTimeQuery.StartAt)
	}

	if len(pageTimeQuery.EndAt) > 0 {
		query = query.Where("create_time <= ?", pageTimeQuery.EndAt)
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

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserLogin
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserLogin(vo *dos.FcUserLogin) []*dos.FcUserLogin {
	var data []*dos.FcUserLogin
	query := global.G_DB.Model(&dos.FcUserLogin{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Password) > 0 {
		query = query.Where("password = ?", vo.Password)
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

	query.Find(&data)
	return data
}

func FindByKeyFcUserLoginFirst(vo *dos.FcUserLogin) *dos.FcUserLogin {
	var data *dos.FcUserLogin
	query := global.G_DB.Model(&dos.FcUserLogin{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Password) > 0 {
		query = query.Where("password = ?", vo.Password)
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
func UpdateFcUserLogin(vo *dos.FcUserLogin) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_name": vo.UserName,
		"password":  vo.Password,
		"create_by": vo.CreateBy,
		"update_by": vo.UpdateBy,
		// "merchant_code": vo.MerchantCode,
	}).Error == nil
}

func DeleteFcUserLogin(vo *dos.FcUserLogin) bool {
	return global.G_DB.Model(&dos.FcUserLogin{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
