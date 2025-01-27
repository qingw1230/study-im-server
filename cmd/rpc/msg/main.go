package main

import (
	"flag"

	"github.com/qingw1230/study-im-server/internal/rpc/msg"
)

func main() {
	rpcPort := flag.Int("port", 10300, "rpc default port")
	flag.Parse()
	rpcServer := msg.NewMsgServer(*rpcPort)
	rpcServer.Run()
}
