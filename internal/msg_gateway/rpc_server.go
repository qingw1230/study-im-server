package msg_gateway

import (
	"context"
	"encoding/json"
	"net"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbReply "github.com/qingw1230/study-im-server/pkg/proto/reply"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *gatewayServer) OnlinePushMsg(_ context.Context, req *pbReply.OnlinePushMsgReq) (*pbReply.OnlinePushMsgResp, error) {
	log.Info("call rpc OnlinePushMsg")
	msgBytes, _ := json.Marshal(req.MsgData)
	mReply := Resp{
		ReqIdentifier: constant.WSPushMsg,
		Data:          msgBytes,
	}
	replyBytes, err := json.Marshal(mReply)
	if err != nil {
		log.Error("data encode failed", err.Error())
		return nil, err
	}

	var resp []*pbReply.SingleMsgToUser
	if conn := ws.getUserConn(req.PushToUserId); conn != nil {
		log.Info("user online push msg", req.PushToUserId)
		tmp := &pbReply.SingleMsgToUser{
			ResultCode: constant.WSNoError,
			RecvId:     req.PushToUserId,
		}
		err := ws.writeMsg(conn, websocket.TextMessage, replyBytes)
		if err != nil {
			tmp.ResultCode = constant.WSWriteError
		}
		resp = append(resp, tmp)
	} else {
		tmp := &pbReply.SingleMsgToUser{
			ResultCode: constant.WSNoUserConn,
			RecvId:     req.PushToUserId,
		}
		resp = append(resp, tmp)
	}

	return &pbReply.OnlinePushMsgResp{
		Resp: resp,
	}, nil
}

type gatewayServer struct {
	pbReply.UnimplementedOnlineMessageRelayServiceServer
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func (s *gatewayServer) onInit(rpcPort int) {
	s.rpcPort = rpcPort
	s.rpcRegisterName = config.Config.RpcRegisterName.OnlineMessageRelayName
	s.etcdSchema = config.Config.Etcd.EtcdSchema
	s.etcdAddr = config.Config.Etcd.EtcdAddr
}

func (s *gatewayServer) run() {
	log.Info("rpc gateway start...")
	address := utils.ServerIP + ":" + strconv.Itoa(s.rpcPort)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("listen network failed", err.Error(), address)
		return
	}
	defer ln.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	pbReply.RegisterOnlineMessageRelayServiceServer(server, s)
	err = etcdv3.RegisterEtcd(s.etcdSchema, s.etcdAddr, utils.ServerIP, s.rpcPort, s.rpcRegisterName, etcdv3.TIME_TO_LIVE)
	if err != nil {
		log.Error("msg_gateway RegisterEtcd failed", err.Error())
		return
	}
	log.Info("rpc msg_gateway register success")
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed", err.Error())
		return
	}
}
