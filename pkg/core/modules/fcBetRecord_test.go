package modules

import (
	"bootpkg/common/response"
	"bootpkg/pkg/core/modules/dos"
	"strings"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newModulesDryRunDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:root@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local",
		SkipInitializeWithVersion: true,
	}), &gorm.Config{DryRun: true, DisableAutomaticPing: true})
	if err != nil {
		t.Fatalf("open dry run db: %v", err)
	}

	return db
}

func TestBuildFcBetRecordPageTotalsQueryUsesSingleAggregateSelect(t *testing.T) {
	db := newModulesDryRunDB(t)
	baseQuery := db.Model(&dos.FcBetRecord{})
	baseQuery = applyFcBetRecordFilters(baseQuery, &dos.FcBetRecord{MerchantCode: "M001", UserId: "u-1"}, true)
	baseQuery = applyFcBetRecordPageTimeFilters(baseQuery, &response.PageTimeQuery{StartAt: "2026-04-01 00:00:00", EndAt: "2026-04-02 00:00:00"})

	tx := buildFcBetRecordPageTotalsQuery(baseQuery).Scan(&fcBetRecordPageTotals{})
	sql := tx.Statement.SQL.String()

	checks := []string{
		"count(1) as total_bet_time",
		"coalesce(sum(bet_amount), 0) as total_bet_amount",
		"coalesce(sum(net_amount), 0) as total_net_amount",
		"coalesce(sum(valid_betamount), 0) as total_valid_bet_amount",
		"FROM `fc_bet_record`",
		"merchant_code = ?",
		"user_id = ?",
		"bet_time >= ?",
		"bet_time <= ?",
	}

	for _, want := range checks {
		if !strings.Contains(sql, want) {
			t.Fatalf("expected SQL to contain %q, got %s", want, sql)
		}
	}
}