package modules

import (
	"bootpkg/common/global"
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"

	"github.com/gin-gonic/gin"
)

func SaveFcBetRecordUnsettled(vo *dos.FcBetRecordUnsettled) (bool, error) {
	rRow := global.G_DB.Create(vo)
	if rRow.Error != nil {
		return false, rRow.Error
	}
	return true, nil
}

func FindPageFcBetRecordUnsettled(page, pageSize int, vo *dos.FcBetRecordUnsettled, pageTimeQuery *response.PageTimeQuery, c *gin.Context) (ret []*dos.FcBetRecordUnsettledExt, total int64) {
	query := global.G_DB.Model(&dos.FcBetRecordUnsettled{})
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

	if c == nil && len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
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

	if c != nil {
		ok := true
		if query, ok = QueryAdminUserMerchantCodes(c, query, vo.MerchantCode); !ok {
			return ret, total
		}
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcBetRecordUnsettledExt
	query.Order("bet_time desc, settlement_time").Offset((page - 1) * pageSize).Limit(pageSize).Preload("Merchant").Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcBetRecordUnsettled(vo *dos.FcBetRecordUnsettled, c *gin.Context) []*dos.FcBetRecordUnsettled {
	var data []*dos.FcBetRecordUnsettled
	query := global.G_DB.Model(&dos.FcBetRecordUnsettled{})
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

func FindByKeyFcBetRecordUnsettledFirst(vo *dos.FcBetRecordUnsettled) *dos.FcBetRecordUnsettled {
	var data *dos.FcBetRecordUnsettled
	query := global.G_DB.Model(&dos.FcBetRecordUnsettled{})
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

	if len(vo.MerchantCode) > 0 {
		query = query.Where("merchant_code = ?", vo.MerchantCode)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcBetRecordUnsettled(vo *dos.FcBetRecordUnsettled) bool {
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

func DeleteFcBetRecordUnsettled(vo *dos.FcBetRecordUnsettled) bool {
	return global.G_DB.Model(&dos.FcBetRecordUnsettled{}).Delete(vo).Error == nil
}
