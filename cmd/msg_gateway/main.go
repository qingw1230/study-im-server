package main

import (
	"flag"
	"sync"

	"github.com/qingw1230/study-im-server/internal/msg_gateway"
)

func main() {
	rpcPort := flag.Int("rpc_port", 10400, "rpc default port")
	wsPort := flag.Int("ws_port", 17778, "ws default port")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	msg_gateway.Init(*rpcPort, *wsPort)
	msg_gateway.Run()
	wg.Wait()
}
