package main

import (
	"flag"

	"github.com/qingw1230/study-im-server/internal/rpc/conversation"
)

func main() {
	rpcPort := flag.Int("port", 11100, "rpc default port")
	flag.Parse()
	rpcServer := conversation.NewConversationServer(*rpcPort)
	rpcServer.Run()
}
