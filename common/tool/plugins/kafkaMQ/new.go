package kafkaMQ

import (
	"bootpkg/common/conf"
	"strings"
)

func NewKafkaProducer(config *conf.Config) (*KafkaOpts, error) {
	opt := &KafkaOpts{}
	opt.IsAsync = config.Mq.Kafka.IsAsync
	err := opt.CreateProducer(strings.Split(config.Mq.Kafka.Addr, ","), config.Mq.Kafka.Version)
	return opt, err
}
