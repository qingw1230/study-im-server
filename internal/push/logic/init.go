package logic

import (
	"github.com/qingw1230/study-im-server/pkg/common/log"
)

var (
	rpcServer pushServer
)

func Init(rpcPort int) {
	log.NewPrivateLog("push")
	rpcServer.onInit(rpcPort)
}

func Run() {
	go rpcServer.run()
}
