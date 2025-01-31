package msg_transfer

import "github.com/qingw1230/study-im-server/pkg/common/log"

var (
	persistentCH PersistentConsumerHandler
	historyCH    HistoryConsumerHandler
)

func Init() {
	log.NewPrivateLog("msg_transfer")
	persistentCH.Init()
	historyCH.Init()
}

func Run() {
	go persistentCH.persistentConsumerGroup.RegisterHandleAndConsumer(&persistentCH)
	go historyCH.historyConsumerGroup.RegisterHandleAndConsumer(&historyCH)
}
