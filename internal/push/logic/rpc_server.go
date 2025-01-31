package logic

import (
	"context"
	"net"
	"strconv"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbPush "github.com/qingw1230/study-im-server/pkg/proto/push"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *pushServer) PushMsg(_ context.Context, req *pbPush.PushMsgReq) (*pbPush.PushMsgResp, error) {
	log.Info("call rpc PushMsg")
	MsgToUser(req)
	return &pbPush.PushMsgResp{
		ResultCode: constant.WSNoError,
	}, nil
}

type pushServer struct {
	pbPush.UnimplementedPushMsgServiceServer
	rpcPort         int
	rpcRegisterName string
	zkSchema        string
	zkAddr          []string
}

func (s *pushServer) onInit(rpcPort int) {
	s.rpcPort = rpcPort
	s.rpcRegisterName = config.Config.RpcRegisterName.PushName
	s.zkSchema = config.Config.Zookeeper.ZKSchema
	s.zkAddr = config.Config.Zookeeper.ZKAddr
}

func (s *pushServer) run() {
	log.Info("rpc push start...")
	address := utils.ServerIP + ":" + strconv.Itoa(s.rpcPort)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("listen network failed", err.Error(), address)
		return
	}
	defer ln.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	pbPush.RegisterPushMsgServiceServer(server, s)
	// TODO(qingw1230): 将 rpc 服务注册进 zk
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed", err.Error())
		return
	}
}
