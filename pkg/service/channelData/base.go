package channelData

import (
	"bootpkg/common/global"
	"bootpkg/common/tool/plugins/kafkaMQ"
	"sync"
)

var producerOtp *kafkaMQ.KafkaOpts
var once sync.Once

func InitProducer() {
	once.Do(func() {
		var err error
		producerOtp, err = kafkaMQ.NewKafkaProducer(global.CONFIG)
		if err != nil {
			panic(err)
		}
	})
}
