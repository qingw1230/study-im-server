package main

import (
	"sync"

	"github.com/qingw1230/study-im-server/internal/msg_transfer"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	msg_transfer.Init()
	msg_transfer.Run()
	wg.Wait()
}
