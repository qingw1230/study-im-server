package main

import (
	"flag"

	"github.com/qingw1230/study-im-server/internal/rpc/group"
)

func main() {
	rpcPort := flag.Int("port", 10500, "rpc default port")
	flag.Parse()
	rpcServer := group.NewGroupServer(*rpcPort)
	rpcServer.Run()
}
