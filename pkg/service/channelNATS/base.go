package channelNATS

import (
	"bootpkg/common/global"
	"github.com/nats-io/nats.go"
)

var _natPS *nats.Conn

func InitNATS() {
	nc, err := nats.Connect(global.CONFIG.NATS.URL)
	if err != nil {
		global.G_LOG.Errorf("nats connection fail:%v", err)
		panic(err)
	}
	_natPS = nc
}
