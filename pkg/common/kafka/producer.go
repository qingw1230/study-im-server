package kafka

import (
	"github.com/IBM/sarama"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"google.golang.org/protobuf/proto"
)

type Producer struct {
	topic    string
	addr     []string
	config   *sarama.Config
	producer sarama.SyncProducer
}

func NewKafkaProducer(addr []string, topic string) *Producer {
	log.Debug("call NewKafkaProducer")
	p := &Producer{
		topic: topic,
		addr:  addr,
	}
	p.config = sarama.NewConfig()
	p.config.Producer.Return.Errors = true
	p.config.Producer.Return.Successes = true
	p.config.Producer.RequiredAcks = sarama.WaitForAll
	p.config.Producer.Partitioner = sarama.NewHashPartitioner

	producer, err := sarama.NewSyncProducer(p.addr, p.config)
	if err != nil {
		log.Error("NewSyncProducer failed", p.addr, p.topic)
		panic(err.Error())
	}
	p.producer = producer
	log.Debug("NewKafkaProducer return")
	return p
}

func (p *Producer) SendMessage(m proto.Message, key ...string) (int32, int64, error) {
	kMsg := &sarama.ProducerMessage{}
	kMsg.Topic = p.topic
	if len(key) == 1 {
		kMsg.Key = sarama.StringEncoder(key[0])
	}
	bMsg, err := proto.Marshal(m)
	if err != nil {
		log.Error("proto Marshal err:", err.Error())
		return -1, -1, err
	}
	kMsg.Value = sarama.ByteEncoder(bMsg)
	return p.producer.SendMessage(kMsg)
}
