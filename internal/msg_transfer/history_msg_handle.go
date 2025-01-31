package msg_transfer

import (
	"github.com/Shopify/sarama"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/kafka"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"google.golang.org/protobuf/proto"
)

type HistoryConsumerHandler struct {
	msgHandle            map[string]fcb
	historyConsumerGroup *kafka.MyConsumerGroup
}

func (hc *HistoryConsumerHandler) Init() {
	hc.msgHandle = make(map[string]fcb)
	hc.msgHandle[config.Config.Kafka.Ws2mschat.Topic] = hc.handleChatWs2Mongo
	hc.historyConsumerGroup = kafka.NewMyConsumerGroup(
		&kafka.MyConsumerGroupConfig{
			KafkaVersion:   sarama.V0_10_2_0,
			OffsetsInitial: sarama.OffsetNewest,
			IsReturnErr:    false,
		},
		[]string{config.Config.Kafka.Ws2mschat.Topic},
		config.Config.Kafka.Ws2mschat.Addr,
		config.Config.Kafka.ConsumerGroupId.MsgToMongo,
	)
}

func (mc *HistoryConsumerHandler) handleChatWs2Mongo(msg []byte, msgKey string) {
	log.Info("call handleChatWs2Mongo")
	msgFromMq := pbMsg.MsgDataToMq{}
	err := proto.Unmarshal(msg, &msgFromMq)
	if err != nil {
		log.Error("msg_transfer Unmarshal failed", err.Error())
		return
	}

	switch msgFromMq.MsgData.SessionType {
	case constant.SingleChatType:
		err := saveUserChat(msgKey, &msgFromMq)
		if err != nil {
			log.Error("single data insert to mongo failed", err.Error())
			return
		}
	}

	log.Info("handleChatWs2Mongo return")
}

func (HistoryConsumerHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (HistoryConsumerHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (mc *HistoryConsumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		mc.msgHandle[msg.Topic](msg.Value, string(msg.Key))
		sess.MarkMessage(msg, "")
	}
	return nil
}
