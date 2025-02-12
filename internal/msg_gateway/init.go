package msg_gateway

import (
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/qingw1230/study-im-server/pkg/common/log"
)

var (
	rw        *sync.RWMutex
	validate  *validator.Validate
	ws        WsServer
	rpcServer gatewayServer
)

func Init(rpcPort, wsPort int) {
	log.NewPrivateLog("msg_gateway")
	rw = &sync.RWMutex{}
	validate = validator.New()
	ws.onInit(wsPort)
	rpcServer.onInit(rpcPort)
}

func Run() {
	go ws.run()
	go rpcServer.run()
}
