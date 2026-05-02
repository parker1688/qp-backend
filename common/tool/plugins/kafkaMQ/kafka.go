package kafkaMQ

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"time"
)

type KafkaSarama interface {
	CreateProducer(addr []string, version string) error
	Send(topic, value string) error
	AsyncResult(sucFunc func(*sarama.ProducerMessage), failFunc func(*sarama.ProducerError))
	CreateConsumerGroup(addr []string, kafka *Kafka, ctx context.Context) error
}

type KafkaOpts struct {
	IsAsync bool
	sarama.SyncProducer
	sarama.AsyncProducer
}

// todo: 创建生产者
func (k *KafkaOpts) CreateProducer(addr []string, version string) error {
	config := sarama.NewConfig()
	ver, err := sarama.ParseKafkaVersion(version)
	if err != nil {
		return err
	}
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = ver

	if k.IsAsync {
		k.AsyncProducer, err = sarama.NewAsyncProducer(addr, config)
		return err
	}
	k.SyncProducer, err = sarama.NewSyncProducer(addr, config)
	return err
}

// todo: 发送消息同步
func (k *KafkaOpts) Send(topic, value string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
	}
	msg.Value = sarama.ByteEncoder(value)
	if k.IsAsync {
		k.Input() <- msg
		return nil
	}
	_, _, err := k.SendMessage(msg)
	return err
}

// todo: 异步消息监控
func (k *KafkaOpts) AsyncResult(sucFunc func(*sarama.ProducerMessage), failFunc func(*sarama.ProducerError)) {
	go func() {
		for {
			select {
			case suc := <-k.Successes():
				sucFunc(suc)
			case fail := <-k.Errors():
				failFunc(fail)
			}
		}
	}()
}

// todo: group消费
func (k *KafkaOpts) CreateConsumerGroup(addr []string, kafka *Kafka, ctx context.Context) error {
	if kafka.Offset == 0 {
		kafka.Offset = sarama.OffsetOldest
	}
	config := sarama.NewConfig()
	ver, err := sarama.ParseKafkaVersion(kafka.Version)
	if err != nil {
		return err
	}
	config.Version = ver
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRange()}
	config.Consumer.Group.Heartbeat.Interval = 6 * time.Second
	config.Consumer.Group.Session.Timeout = 30 * time.Second
	config.Consumer.Offsets.Initial = kafka.Offset
	client, err := sarama.NewConsumerGroup(addr, kafka.GroupId, config)
	if err != nil {
		return err
	}
	go func() {
		for {
			if err := client.Consume(ctx, kafka.Topic, kafka); err != nil {
				fmt.Printf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				fmt.Println(ctx.Err())
				break
			}
		}
	}()
	return nil
}

type Kafka struct {
	Version string
	GroupId string
	Topic   []string
	Offset  int64
	Handler func(*sarama.ConsumerMessage) error
}

func (p *Kafka) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (p *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (p *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		err := p.Handler(message)
		if err == nil {
			// 更新位移
			session.MarkMessage(message, "")
		}
	}
	return nil

	//for {
	//	select {
	//	case message, ok := <-claim.Messages():
	//		if !ok {
	//			fmt.Println("message channel was closed")
	//		} else {
	//			//fmt.Printf("Message claimed: value = %s, timestamp = %v, topic = %s\n", string(message.Value), message.Timestamp, message.Topic)
	//			// 处理数据
	//
	//			session.MarkMessage(message, "")
	//		}
	//	case <-session.Context().Done():
	//	}
	//}
}
