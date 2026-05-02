package controller

import (
	"bootpkg/common"
	"bootpkg/common/tool"
	"bootpkg/pkg/service/channelData"
	"os"
	"testing"
	"time"
)

func init() {
	common.Initialization("../../../conf.yaml")
}

func TestUserInviteConsumer(t *testing.T) {
	if os.Getenv("QP_ENABLE_INTEGRATION_TESTS") != "1" {
		t.Skip("skip integration test: set QP_ENABLE_INTEGRATION_TESTS=1 to enable")
	}

	channelData.SendUserRecharge(&channelData.UserRechargeMessage{
		UserId:        "11",
		OrderSn:       "22",
		DepositTime:   "33",
		DepositAmount: 0,
		T:             time.Now().UnixMicro(),
	})
	t.Log(tool.String(&channelData.UserRechargeMessage{
		UserId:        "11",
		OrderSn:       "22",
		DepositTime:   "33",
		DepositAmount: 0,
		T:             time.Now().UnixMicro(),
	}))
}
