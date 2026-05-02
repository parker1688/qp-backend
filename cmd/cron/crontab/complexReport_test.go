package crontab

import (
	"bootpkg/common"
	"testing"
)

func init() {
	common.Initialization("../../../conf.yaml")
}

func TestTodayComplexReport(t *testing.T) {
	//paymentOut.OrderWithdrawAnotherPayFailRemark("81296806419365889", "代付失败", 0)
	YesterdayComplexReport()
}
