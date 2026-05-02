package crontab

import (
	"bootpkg/common"
	"testing"
)

func init() {
	common.Initialization("../../../conf.yaml")
}

func TestPayOutFailRemark(t *testing.T) {
	//paymentOut.OrderWithdrawAnotherPayFailRemark("81296806419365889", "代付失败", 0)
	t.Log("ok")
}

func TestWithdrawPaymentOut(t *testing.T) {
	WithdrawPaymentOut()
	t.Log("ok")
}

func TestWithdrawPaymentOut2(t *testing.T) {

}
