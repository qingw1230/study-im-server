package main

import (
	"flag"
	"sync"

	"github.com/qingw1230/study-im-server/internal/push/logic"
)

func main() {
	rpcPort := flag.Int("port", 10700, "rpc default port")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	logic.Init(*rpcPort)
	logic.Run()
	wg.Wait()
}
