package main

import (
	"flag"
	"sync"

	"github.com/qingw1230/study-im-server/internal/msg_gateway"
)

func main() {
	wsPort := flag.Int("ws_port", 17778, "ws default port")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	msg_gateway.Init(0, *wsPort)
	msg_gateway.Run()
	wg.Wait()
}
