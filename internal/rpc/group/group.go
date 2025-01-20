package group

import (
	"context"
	"net"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	pbGroup "github.com/qingw1230/study-im-server/pkg/proto/group"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *groupServer) CreateGroup(_ context.Context, req *pbGroup.CreateGroupReq) (*pbGroup.CreateGroupResp, error) {
	log.Info("call CreateGroup args: ", req.String())
	if !token_verify.CheckAccess(req.OpUserId, req.GroupInfo.CreateUserId) {
		log.Error("CheckAccess false ", req.OpUserId, req.GroupInfo.CreateUserId)
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	// 确保群主存在
	groupOwnerInfo, err := controller.FindUserById(req.OwnerUserId)
	if err != nil {
		log.Error("FindUserById failed ", err.Error(), req.OwnerUserId)
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	// 使用 timestamp 和 md5 随机生成 groupId
	groupId := utils.GenerateGroupId()
	groupInfo := &db.Group{}
	copier.Copy(groupInfo, req.GroupInfo)
	groupInfo.GroupId = groupId
	// 创建群信息
	err = controller.InsertIntoGroup(groupInfo)
	if err != nil {
		log.Error("InsertToGroup failed ", err.Error(), groupInfo)
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	groupMember := &db.GroupMember{
		GroupId:   groupId,
		RoleLevel: constant.GroupOwner,
	}
	copier.Copy(groupMember, groupOwnerInfo)
	err = controller.InsertIntoGroupMember(groupMember)
	if err != nil {
		log.Error("InsertIntoGroupMember failed ", err.Error(), groupInfo)
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	log.Info("rpc CreateGroup return ", req.String())
	return &pbGroup.CreateGroupResp{CommonResp: &constant.PBCommonSuccessResp}, nil
}

type groupServer struct {
	pbGroup.UnimplementedGroupServer
	rpcPort         int
	rpcRegisterName string
	zkSchema        string
	zkAddr          []string
}

func NewGroupServer(port int) *groupServer {
	log.NewPrivateLog("group")
	return &groupServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.GroupName,
		zkSchema:        config.Config.Zookeeper.ZKSchema,
		zkAddr:          config.Config.Zookeeper.ZKAddr,
	}
}

func (s *groupServer) Run() {
	log.Info("rpc friend start...")
	address := utils.ServerIP + ":" + strconv.Itoa(s.rpcPort)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("listen network failed ", err.Error(), address)
		return
	}
	defer ln.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	pbGroup.RegisterGroupServer(server, s)
	// TODO(qingw1230): 将 rpc 服务注册进 zk
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed ", err.Error())
		return
	}
}
