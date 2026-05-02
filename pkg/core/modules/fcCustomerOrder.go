// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcCustomerOrder(vo *dos.FcCustomerOrder) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcCustomerOrder(page, pageSize int, vo *dos.FcCustomerOrder, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcCustomerOrder, total int64) {
	query := global.G_DB.Model(&dos.FcCustomerOrder{})
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

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.SolveRemark) > 0 {
		query = query.Where("solve_remark = ?", vo.SolveRemark)
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

	if pageTimeQuery.StartAt != "" {
		query = query.Where("create_time >= ?", pageTimeQuery.StartAt)
	}
	if pageTimeQuery.EndAt != "" {
		query = query.Where("create_time <= ?", pageTimeQuery.EndAt)
	}
	if pageTimeQuery.LastStartAt != "" {
		query = query.Where("update_time >= ?", pageTimeQuery.LastStartAt)
	}
	if pageTimeQuery.LastEndAt != "" {
		query = query.Where("update_time <= ?", pageTimeQuery.LastEndAt)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcCustomerOrder
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcCustomerOrder(vo *dos.FcCustomerOrder, c *gin.Context) []*dos.FcCustomerOrder {
	var data []*dos.FcCustomerOrder
	query := global.G_DB.Model(&dos.FcCustomerOrder{})
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

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark = ?", vo.Remark)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.SolveRemark) > 0 {
		query = query.Where("solve_remark = ?", vo.SolveRemark)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Find(&data)
	return data
}

func FindByKeyFcCustomerOrderFirst(vo *dos.FcCustomerOrder) *dos.FcCustomerOrder {
	var data *dos.FcCustomerOrder
	query := global.G_DB.Model(&dos.FcCustomerOrder{})
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

	if vo.BonusType > 0 {
		query = query.Where("bonus_type = ?", vo.BonusType)
	}

	if len(vo.Title) > 0 {
		query = query.Where("title = ?", vo.Title)
	}

	if len(vo.Remark) > 0 {
		query = query.Where("remark like ?", "%"+vo.Remark+"%")
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	if len(vo.SolveRemark) > 0 {
		query = query.Where("solve_remark = ?", vo.SolveRemark)
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
func UpdateFcCustomerOrder(vo *dos.FcCustomerOrder) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":       vo.UserId,
		"user_name":     vo.UserName,
		"status":        vo.Status,
		"amount":        vo.Amount,
		"bonus_type":    vo.BonusType,
		"title":         vo.Title,
		"remark":        vo.Remark,
		"merchant_code": vo.MerchantCode,
		"solve_remark":  vo.SolveRemark,
		"update_time":   vo.UpdateTime.String(),
		"update_by":     vo.UpdateBy,
	}).Error == nil
}

func DeleteFcCustomerOrder(vo *dos.FcCustomerOrder) bool {
	return global.G_DB.Model(&dos.FcCustomerOrder{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}
