package main

import (
	"flag"

	"github.com/qingw1230/study-im-server/internal/rpc/account"
)

func main() {
	rpcPort := flag.Int("port", 10100, "rpc default port")
	flag.Parse()
	rpcServer := account.NewRpcAccountServer(*rpcPort)
	rpcServer.Run()
}
