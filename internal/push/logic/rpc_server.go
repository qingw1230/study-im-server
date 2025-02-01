package logic

import (
	"context"
	"net"
	"strconv"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
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
	etcdSchema      string
	etcdAddr        []string
}

func (s *pushServer) onInit(rpcPort int) {
	s.rpcPort = rpcPort
	s.rpcRegisterName = config.Config.RpcRegisterName.PushName
	s.etcdSchema = config.Config.Etcd.EtcdSchema
	s.etcdAddr = config.Config.Etcd.EtcdAddr
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
	err = etcdv3.RegisterEtcd(s.etcdSchema, s.etcdAddr, utils.ServerIP, s.rpcPort, s.rpcRegisterName, etcdv3.TIME_TO_LIVE)
	if err != nil {
		log.Error("push RegisterEtcd failed", err.Error())
		return
	}
	log.Info("rpc push register success")
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed", err.Error())
		return
	}
}
