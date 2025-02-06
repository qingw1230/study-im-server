package kafka

import (
	"testing"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/proto/public"
)

const topic = "test-kafka"

func TestSendMessage(t *testing.T) {
	p := NewKafkaProducer(config.Config.Kafka.OfflineMsgToMongoMysql.Addr, topic)
	msg := &public.CommonResp{
		Status: "success",
		Info:   "成功",
	}
	_, _, err := p.SendMessage(msg)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
