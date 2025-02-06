package kafka

import (
	"github.com/IBM/sarama"
)

type Consumer struct {
	addr          []string
	Topic         string
	PartitionList []int32
	Consumer      sarama.Consumer
}

func NewKafkaConsumer(addr []string, topic string) *Consumer {
	c := Consumer{
		addr:  addr,
		Topic: topic,
	}
	consumer, err := sarama.NewConsumer(c.addr, nil)
	if err != nil {
		panic(err.Error())
	}
	c.Consumer = consumer
	partitionList, err := consumer.Partitions(c.Topic)
	if err != nil {
		panic(err.Error())
	}
	c.PartitionList = partitionList
	return &c
}
