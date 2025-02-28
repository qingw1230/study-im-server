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
	cp "github.com/qingw1230/study-im-server/pkg/common/utils"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbGroup "github.com/qingw1230/study-im-server/pkg/proto/group"
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *groupServer) CreateGroup(_ context.Context, req *pbGroup.CreateGroupReq) (*pbGroup.CreateGroupResp, error) {
	log.Info("call CreateGroup args:", req.String())
	if !token_verify.CheckAccess(req.OpUserId, req.OwnerUserId) {
		log.Error("CheckAccess failed", req.OpUserId, req.OwnerUserId)
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	// 确保群主存在
	groupOwnerInfo, err := controller.GetUserById(req.OwnerUserId)
	if err != nil {
		log.Error("GetUserById failed", err.Error(), req.OwnerUserId)
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	// 使用 timestamp 和 md5 随机生成 groupId
	groupId := utils.GenerateGroupId()
	groupInfo := &db.Group{}
	copier.Copy(groupInfo, req.GroupInfo)
	groupInfo.CreateUserId = req.OpUserId
	groupInfo.GroupId = groupId
	// 创建群信息
	err = controller.CreateGroup(groupInfo)
	if err != nil {
		log.Error("InsertIntoGroup failed", err.Error(), groupInfo)
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	// 添加群成员
	groupMember := &db.GroupMember{
		GroupId:        groupId,
		RoleLevel:      constant.GroupOwner,
		InviterUserId:  req.OpUserId,
		OperatorUserId: req.OpUserId,
	}
	copier.Copy(groupMember, groupOwnerInfo)
	err = controller.AddGroupMember(groupMember)
	if err != nil {
		log.Error("AddGroupMember failed", err.Error(), groupInfo)
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	conversationRecord := &db.Conversation{
		OwnerUserId:      req.OwnerUserId,
		ConversationId:   groupId,
		ConversationType: constant.GroupChatType,
		ConversationName: groupInfo.GroupName,
		GroupId:          groupId,
		MemberCount:      1,
	}
	err = controller.InsertIntoConversation(conversationRecord)
	if err != nil {
		log.Error("InsertIntoConversation failed", err.Error())
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	err = db.DB.SetSeq(groupId, 0)
	if err != nil {
		log.Error("SetSeq failed", err.Error())
		return &pbGroup.CreateGroupResp{CommonResp: &constant.PBRedisCommonFailResp}, nil
	}

	var okUserIdList []string
	for _, user := range req.InitMemberList {
		us, err := controller.GetUserById(user.UserId)
		if err != nil {
			log.Error("GetUserById failed", err.Error(), user.UserId)
			continue
		}
		if user.RoleLevel == constant.GroupOwner {
			log.Error("only one owner", user)
			continue
		}
		groupMember.RoleLevel = user.RoleLevel
		copier.Copy(groupMember, us)
		err = controller.AddGroupMember(groupMember)
		if err != nil {
			log.Error("AddGroupMember failed", err.Error(), groupMember)
			continue
		}

		okUserIdList = append(okUserIdList, user.UserId)
	}

	resp := &pbGroup.CreateGroupResp{CommonResp: &constant.PBCommonSuccessResp, GroupInfo: &sdkws.GroupInfo{}}
	group, err := controller.GetGroupInfoByGroupId(groupId)
	if err != nil {
		log.Error("GetGroupInfoByGroupId failed", err.Error(), groupId)
		resp.CommonResp = &constant.PBMySQLCommonFailResp
		return resp, nil
	}
	copier.Copy(resp.GroupInfo, group)
	resp.GroupInfo.MemberCount, err = controller.GetGroupMemberNumByGroupId(groupId)
	if err != nil {
		log.Error("GetGroupMemberNumByGroupId failed", err.Error(), groupId)
		resp.CommonResp = &constant.PBMySQLCommonFailResp
		return resp, nil
	}
	resp.GroupInfo.OwnerUserId = req.OwnerUserId

	log.Info("rpc CreateGroup return", req.String())
	// TODO(qingw1230): 通知进群的人
	return resp, nil
}

func (s *groupServer) DeleteGroup(_ context.Context, req *pbGroup.DeleteGroupReq) (*pbGroup.DeleteGroupResp, error) {
	log.Info("call rpc DeleteGroup args:", req.String())
	// 只能群主或 APP Admin 才能解散群
	if !hasOwnerAccess(req.GroupId, req.OpUserId) {
		log.Error("hasAccess failed", req.GroupId, req.OpUserId)
		return &pbGroup.DeleteGroupResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	err := controller.DeleteGroup(req.GroupId)
	if err != nil {
		log.Error("DeleteGroup failed", err.Error(), req.GroupId)
		return &pbGroup.DeleteGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	log.Info("rpc DeleteGroup return")
	return &pbGroup.DeleteGroupResp{CommonResp: &constant.PBCommonSuccessResp}, nil
}

func (s *groupServer) QuitGroup(_ context.Context, req *pbGroup.QuitGroupReq) (*pbGroup.QuitGroupResp, error) {
	log.Info("call rpc QuitGroup args:", req.String())
	_, err := controller.GetGroupMemberInfoByGroupIdAndUserId(req.GroupId, req.OpUserId)
	if err != nil {
		log.Error("GetGroupMemberInfoByGroupIdAndUserId failed", req.GroupId, req.OpUserId)
		return &pbGroup.QuitGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	err = controller.DeleteGroupMemberByGroupIdAndUserId(req.GroupId, req.OpUserId)
	if err != nil {
		log.Error("DeleteGroupMemberByGroupIdAndUserId", err.Error(), req.GroupId, req.OpUserId)
		return &pbGroup.QuitGroupResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	log.Info("rpc QuitGroup return")
	// TODO(qingw1230): 给群主和管理员发送通知信息
	return &pbGroup.QuitGroupResp{CommonResp: &constant.PBCommonSuccessResp}, nil
}

func (s *groupServer) GetJoinedGroupList(_ context.Context, req *pbGroup.GetJoinedGroupListReq) (*pbGroup.GetJoinedGroupListResp, error) {
	log.Info("GetJoinedGroupList args:", req.String())
	if !token_verify.CheckAccess(req.OpUserId, req.FromUserId) {
		log.Error("CheckAccess failed", req.OpUserId, req.FromUserId)
		return &pbGroup.GetJoinedGroupListResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	var (
		joinedGroupIdList []string
		err               error
	)
	// 根据传入的值获取创建的或加入的
	if req.RoleLevel == constant.GroupOwner {
		joinedGroupIdList, err = controller.GetOwnedGroupIdListByUserId(req.FromUserId)
	} else {
		joinedGroupIdList, err = controller.GetJoinedGroupIdListByUserId(req.FromUserId)
	}
	if err != nil {
		log.Error("GetJoinedGroupIdListByUserId failed", err.Error(), req.FromUserId)
		return &pbGroup.GetJoinedGroupListResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	resp := &pbGroup.GetJoinedGroupListResp{CommonResp: &sdkws.CommonResp{}}
	for _, groupId := range joinedGroupIdList {
		var groupInfo sdkws.GroupInfo
		num, _ := controller.GetGroupMemberNumByGroupId(groupId)
		owner, err1 := controller.GetGroupOwnerInfoByGroupId(groupId)
		group, err2 := controller.GetGroupInfoByGroupId(groupId)
		if num > 0 && owner != nil && err1 == nil && group != nil && err2 == nil {
			copier.Copy(&groupInfo, group)
			groupInfo.MemberCount = num
			groupInfo.OwnerUserId = owner.UserId
			resp.GroupInfoList = append(resp.GroupInfoList, &groupInfo)
		} else {
			log.Error("check nil", num, owner, group)
			continue
		}
	}

	copier.Copy(resp.CommonResp, &constant.PBCommonSuccessResp)
	log.Info("rpc GetJoinedGroupList return", resp.String())
	return resp, nil
}

func (s *groupServer) GetGroupAllMember(_ context.Context, req *pbGroup.GetGroupAllMemberReq) (*pbGroup.GetGroupAllMemberResp, error) {
	log.Info("call rpc GetGroupAllMember args:", req.String())
	var resp pbGroup.GetGroupAllMemberResp
	memberList, err := controller.GetGroupMemberListByGroupId(req.GroupId)
	if err != nil {
		log.Error("GetGroupMemberListByGroupId failed", err.Error())
		resp.CommonResp = &constant.PBMySQLCommonFailResp
		return &resp, nil
	}

	for _, v := range memberList {
		var node sdkws.GroupMemberInfo
		copier.Copy(&node, &v)
		resp.MemberList = append(resp.MemberList, &node)
	}
	log.Info("rpc GetGroupAllMember return")
	resp.CommonResp = &constant.PBCommonSuccessResp
	return &resp, nil
}

func (s *groupServer) GetGroupInfo(_ context.Context, req *pbGroup.GetGroupInfoReq) (*pbGroup.GetGroupInfoResp, error) {
	log.Info("call rpc GetGroupInfo args:", req.String())
	group, err := controller.GetGroupInfoByGroupId(req.GroupId)
	if err != nil {
		log.Error("GetGroupInfoByGroupId failed", req.GroupId)
		return &pbGroup.GetGroupInfoResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}
	groupInfo := &sdkws.GroupInfo{}
	cp.GroupDBCopyIM(groupInfo, group)
	resp := &pbGroup.GetGroupInfoResp{
		CommonResp: &constant.PBCommonSuccessResp,
		GroupInfo:  groupInfo,
	}
	log.Info("rpc GetGroupInfo return")
	return resp, nil
}

func (s *groupServer) SetGroupInfo(_ context.Context, req *pbGroup.SetGroupInfoReq) (*pbGroup.SetGroupInfoResp, error) {
	log.Info("call rpc SetGroupInfo args:", req.String())
	if !hasOwnerOrAdminAccess(req.GroupInfo.GroupId, req.OpUserId) {
		log.Error("hasAccess failed", req.OpUserId)
		return &pbGroup.SetGroupInfoResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	var groupInfo db.Group
	copier.Copy(&groupInfo, req.GroupInfo)
	err := controller.SetGroupInfo(&groupInfo)
	if err != nil {
		log.Error("SetGroupInfo failed", err.Error())
		return &pbGroup.SetGroupInfoResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}
	log.Info("rpc SetGroupInfo return")
	return &pbGroup.SetGroupInfoResp{CommonResp: &constant.PBCommonSuccessResp}, nil
}

// hasOwnerAccess 检查是否是群拥有者
func hasOwnerAccess(groupId, opUserId string) bool {
	if utils.IsContainString(opUserId, config.Config.Admin.UserIds) {
		return true
	}
	groupUserInfo, err := controller.GetGroupMemberInfoByGroupIdAndUserId(groupId, opUserId)
	if err != nil {
		log.Error("GetGroupMemberInfoByGroupIdAndUserId failed", err.Error(), groupId, opUserId)
		return false
	}
	if groupUserInfo.RoleLevel == constant.GroupOwner || groupUserInfo.RoleLevel == constant.GroupAdmin {
		return true
	}
	return false
}

// hasOwnerOrAdminAccess 检查是否是群拥有者或管理员
func hasOwnerOrAdminAccess(groupId, opUserId string) bool {
	if utils.IsContainString(opUserId, config.Config.Admin.UserIds) {
		return true
	}
	groupUserInfo, err := controller.GetGroupMemberInfoByGroupIdAndUserId(groupId, opUserId)
	if err != nil {
		log.Error("GetGroupMemberInfoByGroupIdAndUserId failed", err.Error(), groupId, opUserId)
		return false
	}
	if groupUserInfo.RoleLevel == constant.GroupOwner || groupUserInfo.RoleLevel == constant.GroupAdmin {
		return true
	}
	return false
}

type groupServer struct {
	pbGroup.UnimplementedGroupServer
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewGroupServer(port int) *groupServer {
	log.NewPrivateLog("group")
	return &groupServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.GroupName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (s *groupServer) Run() {
	log.Info("rpc group start...")
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
	err = etcdv3.RegisterEtcd(s.etcdSchema, s.etcdAddr, utils.ServerIP, s.rpcPort, s.rpcRegisterName, etcdv3.TIME_TO_LIVE)
	if err != nil {
		log.Error("group RegisterEtcd failed", err.Error())
		return
	}
	log.Info("rpc group register success")
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed ", err.Error())
		return
	}
}
