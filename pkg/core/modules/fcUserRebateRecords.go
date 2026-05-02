// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/dos"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveFcUserRebateRecords(vo *dos.FcUserRebateRecords) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcUserRebateRecords(page, pageSize int, vo *dos.FcUserRebateRecords, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcUserRebateRecords, total int64) {
	query := global.G_DB.Model(&dos.FcUserRebateRecords{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if vo.RebateType > 0 {
		query = query.Where("rebate_type = ?", vo.RebateType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remarks = ?", vo.Remarks)
	}

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	if pageTimeQuery.StartAt != "" {
		query = query.Where("create_time >= ?", pageTimeQuery.StartAt)
	}
	if pageTimeQuery.EndAt != "" {
		query = query.Where("create_time <= ?", pageTimeQuery.EndAt)
	}

	if len(vo.UpdateBy) > 0 {
		query = query.Where("update_by = ?", vo.UpdateBy)
	}

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcUserRebateRecords
	query.Order("create_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcUserRebateRecords(vo *dos.FcUserRebateRecords, c *gin.Context) []*dos.FcUserRebateRecords {
	var data []*dos.FcUserRebateRecords
	query := global.G_DB.Model(&dos.FcUserRebateRecords{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if vo.RebateType > 0 {
		query = query.Where("rebate_type = ?", vo.RebateType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remarks = ?", vo.Remarks)
	}

	/*if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}*/

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

func FindByKeyFcUserRebateRecordsFirst(vo *dos.FcUserRebateRecords) *dos.FcUserRebateRecords {
	var data *dos.FcUserRebateRecords
	query := global.G_DB.Model(&dos.FcUserRebateRecords{})
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if vo.RebateType > 0 {
		query = query.Where("rebate_type = ?", vo.RebateType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
	}

	if len(vo.Remarks) > 0 {
		query = query.Where("remarks = ?", vo.Remarks)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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
func UpdateFcUserRebateRecords(vo *dos.FcUserRebateRecords) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"user_id":       vo.UserId,
		"user_name":     vo.UserName,
		"venue_code":    vo.VenueCode,
		"game_type":     vo.GameType,
		"bet_amount":    vo.BetAmount,
		"rebate_type":   vo.RebateType,
		"status":        vo.Status,
		"bonus_amount":  vo.BonusAmount,
		"bonus_rate":    vo.BonusRate,
		"remarks":       vo.Remarks,
		"merchant_code": vo.MerchantCode,
		"create_by":     vo.CreateBy,
		"update_by":     vo.UpdateBy,
		"update_time":   vo.UpdateTime,
	}).Error == nil
}

func DeleteFcUserRebateRecords(vo *dos.FcUserRebateRecords) bool {
	return global.G_DB.Model(&dos.FcUserRebateRecords{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

func FcUserRebateRecordsList(vo *dos.FcUserRebateRecords, pageTimeQuery *response.PageTimeQuery) (ret []*dos.FcUserRebateRecordsListRow, total int64, totalBonusAmount float64) {
	query := global.G_DB.Model(&dos.FcUserRebateRecords{})
	totalQuery := global.G_DB.Model(&dos.FcUserRebateRecords{})

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
		totalQuery = totalQuery.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
		totalQuery = totalQuery.Where("user_name = ?", vo.UserName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
		totalQuery = totalQuery.Where("venue_code = ?", vo.VenueCode)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
		totalQuery = totalQuery.Where("game_type = ?", vo.GameType)
	}

	if vo.RebateType > 0 {
		query = query.Where("rebate_type = ?", vo.RebateType)
		totalQuery = totalQuery.Where("rebate_type = ?", vo.RebateType)
	}

	if vo.Status > 0 {
		query = query.Where("status = ?", vo.Status)
		totalQuery = totalQuery.Where("status = ?", vo.Status)
	}

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
		totalQuery = totalQuery.Where("merchant_code = ?", vo.MerchantCode)
	}

	if pageTimeQuery.TimeType == -1 {
		if pageTimeQuery.StartAt != "" {
			query = query.Where("create_time >= ?", pageTimeQuery.StartAt)
			totalQuery = totalQuery.Where("create_time >= ?", pageTimeQuery.StartAt)
		}
		if pageTimeQuery.EndAt != "" {
			query = query.Where("create_time <= ?", pageTimeQuery.EndAt)
			totalQuery = totalQuery.Where("create_time <= ?", pageTimeQuery.EndAt)
		}
	} else {
		sTime, eTime := tool.GetDayRange(time.Now(), pageTimeQuery.TimeType)
		query.Where("create_time BETWEEN ? AND ?", sTime, eTime)
		totalQuery.Where("create_time BETWEEN ? AND ?", sTime, eTime)
	}

	totalQuery.Select("sum(bonus_amount)").Scan(&totalBonusAmount)

	var count int64
	query.Count(&count)
	dataSlice := []*dos.FcUserRebateRecordsListRow{}
	query.Select("create_time", "sum(bonus_amount) as bonus_amount").
		Group("create_time").
		Order("create_time desc").
		Offset((pageTimeQuery.PageNo - 1) * pageTimeQuery.PageSize).Limit(pageTimeQuery.PageSize).
		Find(&dataSlice)

	if dataSlice != nil {
		dateVenueMp := getUserRebateRecordDateVenueMap(vo.UserId, dataSlice)
		for i, v := range dataSlice {
			if v, ok := dateVenueMp[v.CreateTime.String()]; ok {
				// 将时间对应最新一条数据的场馆码和游戏类型赋值
				dataSlice[i].VenueCode = v.VenueCode
				dataSlice[i].GameType = v.GameType
			}
		}
	}

	return dataSlice, count, totalBonusAmount
}

func getUserRebateRecordDateVenueMap(userId string, data []*dos.FcUserRebateRecordsListRow) map[string]dos.FcUserRebateRecords {
	result := map[string]dos.FcUserRebateRecords{}

	if data == nil {
		return result
	}

	dateLis := []string{}
	for _, v := range data {
		dateLis = append(dateLis, v.CreateTime.String())
	}

	lis := []dos.FcUserRebateRecords{}

	err := global.G_DB.Model(&dos.FcUserRebateRecords{}).
		Select("create_time", "venue_code", "game_type").
		Where("user_id = ? and create_time in ?", userId, dateLis).
		Order("create_time desc").
		Find(&lis).Error
	if err != nil {
		global.G_LOG.Errorf("[getUserRebateRecordDateVenueMap] Can't find user rebate records by dates: userId=%s, dates=%v, err=%v", userId, dateLis, err.Error())
		return result
	}

	for _, v := range lis {
		if _, ok := result[v.CreateTime.String()]; !ok {
			result[v.CreateTime.String()] = dos.FcUserRebateRecords{
				VenueCode: v.VenueCode,
				GameType:  v.GameType,
			}
		}
	}

	return result
}

// GetFcUserRebateDetailRecordList - 获取用户返水详情记录列表
// @param {string} string
// @param {string} date
// / @returns ]dos.FcUserRebateRecordsDetailListResp
func GetFcUserRebateDetailRecordList(userId string, date string) []dos.FcUserRebateRecordsDetailListResp {
	dataSlice := []dos.FcUserRebateRecordsDetailListResp{}
	global.G_DB.Model(&dos.FcUserRebateRecords{}).
		Where("user_id = ? AND create_time = ?", userId, date).
		Order("create_time desc").
		Find(&dataSlice)

	return dataSlice
}
