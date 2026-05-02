package controller

import (
	"bootpkg/common/tool/plugins/kafkaMQ"
	"context"
)

var (
	kafkaOpt *kafkaMQ.KafkaOpts
)

var kafkaInitFun []func(ctx context.Context)

func NewConsumers(ctx context.Context) {
	kafkaOpt = &kafkaMQ.KafkaOpts{}
	//自动执行方法Map
	if len(kafkaInitFun) > 0 {
		for _, f := range kafkaInitFun {
			f(ctx)
		}
	}
	kafkaInitFun = nil
}
