package msg_transfer

import (
	"github.com/Shopify/sarama"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/common/kafka"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"google.golang.org/protobuf/proto"
)

type fcb func(msg []byte, msgKey string)

type PersistentConsumerHandler struct {
	msgHandle               map[string]fcb
	persistentConsumerGroup *kafka.MyConsumerGroup
}

func (pc *PersistentConsumerHandler) Init() {
	pc.msgHandle = make(map[string]fcb)
	pc.msgHandle[config.Config.Kafka.Ws2mschat.Topic] = pc.handleChatWs2Mysql
	pc.persistentConsumerGroup = kafka.NewMyConsumerGroup(
		&kafka.MyConsumerGroupConfig{
			KafkaVersion:   sarama.V0_10_2_0,
			OffsetsInitial: sarama.OffsetNewest,
			IsReturnErr:    false,
		},
		[]string{config.Config.Kafka.Ws2mschat.Topic},
		config.Config.Kafka.Ws2mschat.Addr,
		config.Config.Kafka.ConsumerGroupId.MsgToMySql,
	)
}

func (pc *PersistentConsumerHandler) handleChatWs2Mysql(msg []byte, msgKey string) {
	msgFromMq := pbMsg.MsgDataToMq{}
	err := proto.Unmarshal(msg, &msgFromMq)
	if err != nil {
		log.Error("msg_transfer Unmarshal failed", err.Error())
		return
	}

	if msgKey == msgFromMq.MsgData.RecvId && msgFromMq.MsgData.SessionType == constant.SingleChatType {
		log.Info("msg_transfer msg persisting")
		if err = controller.InsertMessageToChatLog(&msgFromMq); err != nil {
			log.Error("InsertMessageToChatLog failed", err.Error())
		}
	}
}

func (PersistentConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (PersistentConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (pc *PersistentConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		pc.msgHandle[msg.Topic](msg.Value, string(msg.Key))
		sess.MarkMessage(msg, "")
	}
	return nil
}
