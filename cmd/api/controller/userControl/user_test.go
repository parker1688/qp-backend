package userControl

import (
	"bootpkg/common"
	"bootpkg/common/global"
	"bootpkg/common/tool"
	"bootpkg/pkg/core/modules/enmus"
	vos "bootpkg/pkg/core/modules/vo"
	"context"
	"fmt"
	"testing"
	"time"
)

func init() {
	common.Initialization("../../../../conf.yaml")
}

func TestRechargeSuccessInfoVO(t *testing.T) {
	global.G_REDIS.Set(context.Background(), fmt.Sprintf(enmus.RECHARGESUCCESSINFO, "62748047137832961"), tool.String(&vos.RechargeSuccessInfoVO{
		Success:  true,
		Amount:   tool.String(100),
		Currency: "THB",
	}), 30*time.Hour)
}
