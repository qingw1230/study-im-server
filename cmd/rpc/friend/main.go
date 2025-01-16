package main

import (
	"flag"

	"github.com/qingw1230/study-im-server/internal/rpc/friend"
)

func main() {
	rpcPort := flag.Int("port", 10200, "rpc default port")
	flag.Parse()
	rpcServer := friend.NewFriendServer(*rpcPort)
	rpcServer.Run()
}
