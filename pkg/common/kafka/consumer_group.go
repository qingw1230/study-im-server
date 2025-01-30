package kafka

import (
	"context"

	"github.com/Shopify/sarama"
)

type MyConsumerGroup struct {
	sarama.ConsumerGroup
	groupId string
	topics  []string
}

type MyConsumerGroupConfig struct {
	KafkaVersion   sarama.KafkaVersion
	OffsetsInitial int64
	IsReturnErr    bool
}

func NewConsumerGroup(consumerConfig *MyConsumerGroupConfig, topics, addr []string, groupId string) *MyConsumerGroup {
	config := sarama.NewConfig()
	config.Version = consumerConfig.KafkaVersion
	config.Consumer.Offsets.Initial = consumerConfig.OffsetsInitial
	config.Consumer.Return.Errors = consumerConfig.IsReturnErr
	client, err := sarama.NewClient(addr, config)
	if err != nil {
		panic(err.Error())
	}
	consumerGroup, err := sarama.NewConsumerGroupFromClient(groupId, client)
	if err != nil {
		panic(err.Error())
	}
	return &MyConsumerGroup{
		ConsumerGroup: consumerGroup,
		groupId:       groupId,
		topics:        topics,
	}
}

func (mc *MyConsumerGroup) RegisterHandleAndConsumer(handler sarama.ConsumerGroupHandler) {
	ctx := context.Background()
	for {
		err := mc.ConsumerGroup.Consume(ctx, mc.topics, handler)
		if err != nil {
			panic(err.Error())
		}
	}
}
