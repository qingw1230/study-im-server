package kafka

import (
	"context"
	"fmt"
	"testing"

	"github.com/IBM/sarama"
	"github.com/qingw1230/study-im-server/pkg/common/config"
)

type exampleConsumerGroupHandler struct{}

func (exampleConsumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (exampleConsumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h exampleConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d\n", msg.Topic, msg.Partition, msg.Offset)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func TestConsumeMessage(t *testing.T) {
	topics := []string{topic}
	group := NewMyConsumerGroup(
		&MyConsumerGroupConfig{
			KafkaVersion:   sarama.V3_5_1_0,
			OffsetsInitial: sarama.OffsetNewest,
			IsReturnErr:    false,
		},
		topics, config.Config.Kafka.OfflineMsgToMongoMysql.Addr, "test",
	)
	defer func() {
		_ = group.Close()
	}()

	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	ctx := context.Background()
	for {
		handler := exampleConsumerGroupHandler{}
		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}
