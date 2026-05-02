package report

import (
	"bootpkg/common/expands/automaticType"
	"bootpkg/pkg/core/modules/dos"
	"strings"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newComplexReportDryRunDB(t *testing.T) *gorm.DB {
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

func buildBetRecordAggregateSQL(t *testing.T, report ComplexReport, selectExpr string) string {
	t.Helper()

	db := newComplexReportDryRunDB(t)
	tx := buildComplexReportBetRecordQuery(db.Model(&dos.FcBetRecord{}), report).Select(selectExpr)
	tx = tx.Scan(&struct{}{})
	return tx.Statement.SQL.String()
}

func TestBuildComplexReportBetRecordQueryAppliesFilters(t *testing.T) {
	report := ComplexReport{
		UserId:       "u-1",
		MerchantCode: "M001",
		StartAt:      automaticType.Time(time.Date(2026, 4, 1, 0, 0, 0, 0, time.Local)),
		EndAt:        automaticType.Time(time.Date(2026, 4, 2, 0, 0, 0, 0, time.Local)),
	}

	sql := buildBetRecordAggregateSQL(t, report, "sum(net_amount) as betWin")

	checks := []string{
		"FROM `fc_bet_record`",
		"user_id=?",
		"bet_time>=?",
		"bet_time<?",
		"merchant_code =?",
		"sum(net_amount) as betWin",
	}

	for _, want := range checks {
		if !strings.Contains(sql, want) {
			t.Fatalf("expected SQL to contain %q, got %s", want, sql)
		}
	}
}

func TestBuildComplexReportBetRecordQueryOmitsUnsetFilters(t *testing.T) {
	sql := buildBetRecordAggregateSQL(t, ComplexReport{}, "sum(valid_betamount) as betAmount")

	unwanted := []string{"user_id=?", "merchant_code =?", "bet_time>=?", "bet_time<?"}
	for _, fragment := range unwanted {
		if strings.Contains(sql, fragment) {
			t.Fatalf("expected SQL to omit %q, got %s", fragment, sql)
		}
	}

	if !strings.Contains(sql, "FROM `fc_bet_record`") {
		t.Fatalf("expected SQL to target fc_bet_record, got %s", sql)
	}
}
