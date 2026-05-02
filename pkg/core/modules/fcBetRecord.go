// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SaveFcBetRecord(vo *dos.FcBetRecord) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func applyFcBetRecordFilters(query *gorm.DB, vo *dos.FcBetRecord, includeMerchantCode bool) *gorm.DB {
	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.UserId) > 0 {
		query = query.Where("user_id = ?", vo.UserId)
	}

	if len(vo.UserName) > 0 {
		query = query.Where("user_name = ?", vo.UserName)
	}

	if len(vo.Account) > 0 {
		query = query.Where("account = ?", vo.Account)
	}

	if len(vo.PlayerName) > 0 {
		query = query.Where("player_name = ?", vo.PlayerName)
	}

	if len(vo.OrderSn) > 0 {
		query = query.Where("order_sn = ?", vo.OrderSn)
	}

	if len(vo.GameCode) > 0 {
		query = query.Where("game_code = ?", vo.GameCode)
	}

	if !vo.BetTime.Timer().IsZero() {
		query = query.Where("bet_time = ?", vo.BetTime)
	}

	if !vo.SettlementTime.Timer().IsZero() {
		query = query.Where("settlement_time = ?", vo.SettlementTime)
	}

	if !vo.ThirdBettime.Timer().IsZero() {
		query = query.Where("third_bettime = ?", vo.ThirdBettime)
	}

	if !vo.ThirdSettlementtime.Timer().IsZero() {
		query = query.Where("third_settlementtime = ?", vo.ThirdSettlementtime)
	}

	if len(vo.GameType) > 0 {
		query = query.Where("game_type = ?", vo.GameType)
	}

	if len(vo.GameName) > 0 {
		query = query.Where("game_name = ?", vo.GameName)
	}

	if len(vo.VenueCode) > 0 {
		query = query.Where("venue_code = ?", vo.VenueCode)
	}

	if vo.OddsType > 0 {
		query = query.Where("odds_type = ?", vo.OddsType)
	}

	if vo.Version > 0 {
		query = query.Where("version = ?", vo.Version)
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

	if includeMerchantCode && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	return query
}

func applyFcBetRecordPageTimeFilters(query *gorm.DB, pageTimeQuery *response.PageTimeQuery) *gorm.DB {
	if pageTimeQuery == nil {
		return query
	}

	if len(pageTimeQuery.StartAt) > 0 {
		query = query.Where("bet_time >= ?", pageTimeQuery.StartAt)
	}
	if len(pageTimeQuery.EndAt) > 0 {
		query = query.Where("bet_time <= ?", pageTimeQuery.EndAt)
	}
	if len(pageTimeQuery.LastStartAt) > 0 {
		query = query.Where("settlement_time >= ?", pageTimeQuery.LastStartAt)
	}
	if len(pageTimeQuery.LastEndAt) > 0 {
		query = query.Where("settlement_time <= ?", pageTimeQuery.LastEndAt)
	}

	return query
}

type fcBetRecordPageTotals struct {
	TotalBetTime        int64   `gorm:"column:total_bet_time"`
	TotalBetAmount      float64 `gorm:"column:total_bet_amount"`
	TotalNetAmount      float64 `gorm:"column:total_net_amount"`
	TotalValidBetAmount float64 `gorm:"column:total_valid_bet_amount"`
}

func buildFcBetRecordPageTotalsQuery(query *gorm.DB) *gorm.DB {
	return query.Select(
		"count(1) as total_bet_time, " +
			"coalesce(sum(bet_amount), 0) as total_bet_amount, " +
			"coalesce(sum(net_amount), 0) as total_net_amount, " +
			"coalesce(sum(valid_betamount), 0) as total_valid_bet_amount",
	)
}

func findFcBetRecordPageTotals(query *gorm.DB) fcBetRecordPageTotals {
	totals := fcBetRecordPageTotals{}
	buildFcBetRecordPageTotalsQuery(query).Scan(&totals)
	return totals
}

func FindPageFcBetRecord(page, pageSize int, vo *dos.FcBetRecord, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcBetRecordExt, total int64, bettime int64, betAmount float64, netAmount float64, validBetAmount float64) {
	query := global.G_DB.Model(&dos.FcBetRecord{})
	query = applyFcBetRecordFilters(query, vo, c == nil)
	query = applyFcBetRecordPageTimeFilters(query, pageTimeQuery)

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total, 0, 0.0, 0.0, 0.0
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcBetRecordExt
	query.Order("bet_time desc, settlement_time").Offset((page - 1) * pageSize).Limit(pageSize).Preload("Merchant").Find(&dataSlice)
	totals := findFcBetRecordPageTotals(query)

	return dataSlice, count, totals.TotalBetTime, totals.TotalBetAmount, totals.TotalNetAmount, totals.TotalValidBetAmount
}

func FindByKeyFcBetRecord(vo *dos.FcBetRecord, c *gin.Context) []*dos.FcBetRecord {
	var data []*dos.FcBetRecord
	query := global.G_DB.Model(&dos.FcBetRecord{})
	query = applyFcBetRecordFilters(query, vo, c == nil)

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return data
		}
	}

	query.Order("bet_time desc, settlement_time").Find(&data)
	return data
}

func FindByKeyFcBetRecordFirst(vo *dos.FcBetRecord) *dos.FcBetRecord {
	var data *dos.FcBetRecord
	query := global.G_DB.Model(&dos.FcBetRecord{})
	query = applyFcBetRecordFilters(query, vo, true)

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcBetRecord(vo *dos.FcBetRecord) bool {
	return global.G_DB.Model(vo).Updates(map[string]interface{}{
		"id":                   vo.Id,
		"user_id":              vo.UserId,
		"user_name":            vo.UserName,
		"account":              vo.Account,
		"player_name":          vo.PlayerName,
		"order_sn":             vo.OrderSn,
		"game_code":            vo.GameCode,
		"bet_amount":           vo.BetAmount,
		"net_amount":           vo.NetAmount,
		"valid_betamount":      vo.ValidBetamount,
		"bet_time":             vo.BetTime,
		"settlement_time":      vo.SettlementTime,
		"third_bettime":        vo.ThirdBettime,
		"third_settlementtime": vo.ThirdSettlementtime,
		"game_type":            vo.GameType,
		"game_name":            vo.GameName,
		"venue_code":           vo.VenueCode,
		"odds":                 vo.Odds,
		"odds_type":            vo.OddsType,
		"version":              vo.Version,
		"create_by":            vo.CreateBy,
		"update_by":            vo.UpdateBy,
		"merchant_code":        vo.MerchantCode,
		"merchant_net_amount":  vo.MerchantNetAmount,
		"after_balance":        vo.AfterBalance,
	}).Error == nil
}

func DeleteFcBetRecord(vo *dos.FcBetRecord) bool {
	return global.G_DB.Model(&dos.FcBetRecord{}).Delete(vo).Error == nil
}
