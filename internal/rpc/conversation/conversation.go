package conversation

import (
	"context"
	"net"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbConversation "github.com/qingw1230/study-im-server/pkg/proto/conversation"
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *conversationServer) GetConversationList(_ context.Context, req *pbConversation.GetConversationListReq) (*pbConversation.GetConversationListResp, error) {
	log.Info("call rpc GetConversationList")
	if !token_verify.CheckAccess(req.OpUserId, req.FromUserId) {
		log.Error("CheckAccess failed", req.OpUserId, req.FromUserId)
		return &pbConversation.GetConversationListResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	conversationList, err := controller.GetConversationList(req.FromUserId)
	if err != nil {
		log.Error("GetConversationList failed", err.Error())
		return &pbConversation.GetConversationListResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	resp := &pbConversation.GetConversationListResp{CommonResp: &constant.PBCommonSuccessResp}
	for _, c := range conversationList {
		tmp := &sdkws.ConversationInfo{}
		copier.Copy(tmp, &c)
		resp.ConversationList = append(resp.ConversationList, tmp)
	}
	return resp, nil
}

type conversationServer struct {
	pbConversation.UnimplementedConversationServer
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewConversationServer(port int) *conversationServer {
	log.NewPrivateLog("conversation")
	return &conversationServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.ConversationName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (s *conversationServer) Run() {
	log.Info("rpc conversation start...")
	address := utils.ServerIP + ":" + strconv.Itoa(s.rpcPort)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("listen network failed", err.Error(), address)
		return
	}
	defer ln.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	pbConversation.RegisterConversationServer(server, s)
	err = etcdv3.RegisterEtcd(s.etcdSchema, s.etcdAddr, utils.ServerIP, s.rpcPort, s.rpcRegisterName, etcdv3.TIME_TO_LIVE)
	if err != nil {
		log.Error("conversation RegisterEtcd failed", err.Error())
		return
	}
	log.Info("rpc conversatin success")
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed", err.Error())
		return
	}
}
