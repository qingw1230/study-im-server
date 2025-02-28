package kafka

import (
	"testing"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
)

const topic = "test-kafka"

func TestSendMessage(t *testing.T) {
	p := NewKafkaProducer(config.Config.Kafka.OfflineMsgToMongoMysql.Addr, topic)
	msg := &sdkws.CommonResp{
		Status: "success",
		Info:   "成功",
	}
	_, _, err := p.SendMessage(msg)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
