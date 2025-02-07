package friend

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/qingw1230/study-im-server/internal/rpc/msg"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	cp "github.com/qingw1230/study-im-server/pkg/common/utils"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbFriend "github.com/qingw1230/study-im-server/pkg/proto/friend"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *friendServer) AddFriend(_ context.Context, req *pbFriend.AddFriendReq) (*pbFriend.AddFriendResp, error) {
	log.Info("call rpc AddFriend args:", req.String())
	if !token_verify.CheckAccess(req.CommonId.OpUserId, req.CommonId.FromUserId) {
		log.Error("CheckAccess failed", req.CommonId.OpUserId, req.CommonId.FromUserId)
		return &pbFriend.AddFriendResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	if _, err := controller.GetUserById(req.CommonId.ToUserId); err != nil {
		log.Error("GetUserById failed", err.Error())
		return &pbFriend.AddFriendResp{CommonResp: &pbPublic.CommonResp{Status: constant.Fail, Code: constant.MySQLRecordNotExists, Info: constant.MySQLAccountNotExistsInfo}}, nil
	}

	// TODO(qingw1230): 检查是否已是好友
	// TODO(qingw1230): 检查对方有没有添加自己的记录

	friendRequest := &db.FriendRequest{
		FromUserId:   req.CommonId.FromUserId,
		ToUserId:     req.CommonId.ToUserId,
		ReqMsg:       req.ReqMsg,
		HandleResult: constant.FriendRequestDefault,
		CreateTime:   time.Now().UnixMilli(),
	}
	_, err := controller.GetFriendRequestByBothUserId(friendRequest.FromUserId, friendRequest.ToUserId)
	if err == nil {
		err = controller.UpdateFriendRequest(friendRequest)
	} else {
		err = controller.InsertFriendRequest(friendRequest)
	}
	if err != nil {
		log.Error("InsertFriendRequest failed", err.Error(), friendRequest)
		return &pbFriend.AddFriendResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	msg.FriendRequestNotification(req, friendRequest)
	log.Info("rpc AddFriend return")
	return &pbFriend.AddFriendResp{CommonResp: &constant.PBCommonSuccessResp}, nil
}

func (s *friendServer) AddFriendResponse(_ context.Context, req *pbFriend.AddFriendResponseReq) (*pbFriend.AddFriendResponseResp, error) {
	log.Info("call rpc AddFriendResponse args:", req.String())
	if !token_verify.CheckAccess(req.CommonId.OpUserId, req.CommonId.ToUserId) {
		log.Error("CheckAccess false ", req.CommonId.OpUserId, req.CommonId.ToUserId)
		return &pbFriend.AddFriendResponseResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	// 在同意或拒绝好友申请之前，先检查记录是否存在
	friendRequest, err := controller.GetFriendRequestByBothUserId(req.CommonId.FromUserId, req.CommonId.ToUserId)
	if err != nil {
		return &pbFriend.AddFriendResponseResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.MySQLRecordNotExists,
				Info:   constant.MySQLRecordNotExistsInfo,
			},
		}, nil
	}
	friendRequest.HandleResult = req.HandleResult
	friendRequest.HandleTime = time.Now().UnixMilli()
	resp := &pbFriend.AddFriendResponseResp{CommonResp: &constant.PBMySQLCommonFailResp}
	err = controller.UpdateFriendRequest(friendRequest)
	if err != nil {
		return resp, nil
	}
	log.Info("UpdateFriendApplication success ", friendRequest)

	// 同意好友请求，向好友表和会话表插入记录
	if friendRequest.HandleResult == constant.FriendResponseAgree {
		// 插入两条单向好友关系
		_, err1 := controller.GetFriendRelationFromFriend(req.CommonId.FromUserId, req.CommonId.ToUserId)
		if err1 == nil {
			log.Warn("GetFriendRelationFromFriend exist", req.CommonId.FromUserId, req.CommonId.ToUserId)
		} else if err1 == gorm.ErrRecordNotFound {
			toInsertFollow := db.Friend{OwnerUserId: req.CommonId.FromUserId, FriendUserId: req.CommonId.ToUserId, OpUserId: req.CommonId.OpUserId}
			err := controller.InsertIntoFriend(&toInsertFollow)
			if err != nil {
				log.Error("InsertIntoFriend failed", err.Error(), toInsertFollow)
				return resp, nil
			}
		} else {
			log.Error("GetFriendRelationFromFriend failed", err1.Error())
			return resp, nil
		}

		_, err2 := controller.GetFriendRelationFromFriend(req.CommonId.ToUserId, req.CommonId.FromUserId)
		if err2 == nil {
			log.Warn("GetFriendRelationFromFriend exist", req.CommonId.ToUserId, req.CommonId.FromUserId)
		} else if err2 == gorm.ErrRecordNotFound {
			toInsertFollow := db.Friend{OwnerUserId: req.CommonId.ToUserId, FriendUserId: req.CommonId.FromUserId, OpUserId: req.CommonId.OpUserId}
			err := controller.InsertIntoFriend(&toInsertFollow)
			if err != nil {
				log.Error("InsertIntoFriend failed", err.Error(), toInsertFollow)
				return resp, nil
			}
			// TODO(qingw1230): 通知另一方
		} else {
			log.Error("GetFriendRelationFromFriend failed", err2.Error())
			return resp, nil
		}

		if err1 == gorm.ErrRecordNotFound && err2 == gorm.ErrRecordNotFound {
			user1, _ := controller.GetUserById(req.CommonId.ToUserId)
			conversationRecord1 := &db.Conversation{
				OwnerUserId:      req.CommonId.FromUserId,
				ConversationId:   req.CommonId.FromUserId + req.CommonId.ToUserId,
				ConversationType: constant.SingleChatType,
				ConversationName: user1.NickName,
				UserId:           req.CommonId.ToUserId,
			}
			controller.InsertIntoConversation(conversationRecord1)

			user2, _ := controller.GetUserById(req.CommonId.FromUserId)
			conversationRecord2 := &db.Conversation{
				OwnerUserId:      req.CommonId.ToUserId,
				ConversationId:   req.CommonId.ToUserId + req.CommonId.FromUserId,
				ConversationType: constant.SingleChatType,
				ConversationName: user2.NickName,
				UserId:           req.CommonId.FromUserId,
			}
			controller.InsertIntoConversation(conversationRecord2)
			// TODO(qingw1230): 向会话添加 ReqMsg
		}
	}

	// TODO(qingw1230): 通知或拒绝的通知
	copier.Copy(resp.CommonResp, &constant.PBCommonSuccessResp)
	return resp, nil
}

func (s *friendServer) DeleteFriend(_ context.Context, req *pbFriend.DeleteFriendReq) (*pbFriend.DeleteFriendResp, error) {
	log.Info("call DeleteFriend args: ", req.String())
	// 确保有权限
	if !token_verify.CheckAccess(req.CommonId.OpUserId, req.CommonId.FromUserId) {
		log.Error("CheckAccess false ", req.CommonId.OpUserId, req.CommonId.FromUserId)
		return &pbFriend.DeleteFriendResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	err := controller.DeleteSingleFriendRelation(req.CommonId.FromUserId, req.CommonId.ToUserId)
	if err != nil {
		log.Error("DeleteSingleFriendRelation failed ", err.Error())
		return &pbFriend.DeleteFriendResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	return &pbFriend.DeleteFriendResp{CommonResp: &constant.PBCommonSuccessResp}, nil
}

func (s *friendServer) GetFriendList(_ context.Context, req *pbFriend.GetFriendListReq) (*pbFriend.GetFriendListResp, error) {
	log.Info("call rpc GetFriendList args: ", req.String())
	if !token_verify.CheckAccess(req.CommonId.OpUserId, req.CommonId.FromUserId) {
		log.Error("CheckAccess false ", req.CommonId.OpUserId, req.CommonId.FromUserId)
		return &pbFriend.GetFriendListResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	friends, err := controller.GetFriendListByUserId(req.CommonId.FromUserId)
	if err != nil {
		log.Error("GetFriendListByUserId failed ", err.Error(), req.CommonId.FromUserId)
		return &pbFriend.GetFriendListResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	var userInfoList []*pbPublic.FriendInfo
	for _, user := range friends {
		friendUserInfo := &pbPublic.FriendInfo{FriendInfo: &pbPublic.PublicUserInfo{}}
		cp.FriendDBCopyIM(friendUserInfo, &user)
		userInfoList = append(userInfoList, friendUserInfo)
	}
	log.Info("rpc GetFriendList return")
	return &pbFriend.GetFriendListResp{
		CommonResp:     &constant.PBCommonSuccessResp,
		FriendInfoList: userInfoList,
	}, nil
}

func (s friendServer) GetFriendApplyList(_ context.Context, req *pbFriend.GetFriendApplyListReq) (*pbFriend.GetFriendApplyListResp, error) {
	log.Info("call rpc GetFriendApplyList args:", req.String())
	if !token_verify.CheckAccess(req.CommonId.OpUserId, req.CommonId.FromUserId) {
		log.Error("CheckAccess failed", req.CommonId.OpUserId, req.CommonId.FromUserId)
		return &pbFriend.GetFriendApplyListResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	offset := (req.PageNumber - 1) * req.ShowNumber
	friendRequests, total, err := controller.GetReceivedFriendRequestList(req.CommonId.FromUserId, int(offset), int(req.ShowNumber))
	if err != nil {
		log.Error("GetReceivedFriendApplicationListByUserId failed", err.Error(), req.CommonId.FromUserId)
		return &pbFriend.GetFriendApplyListResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	var applyUserList []*pbFriend.FriendRequest
	for _, fr := range friendRequests {
		var pbFr pbFriend.FriendRequest
		cp.FriendRequestDBCopyIM(&pbFr, &fr)
		applyUserList = append(applyUserList, &pbFr)
	}

	log.Info("rpc GetFriendApplyListResp return")
	return &pbFriend.GetFriendApplyListResp{
		CommonResp:        &constant.PBCommonSuccessResp,
		FriendRequestList: applyUserList,
		Total:             int64(total),
	}, nil
}

func (s *friendServer) AddBlacklist(_ context.Context, req *pbFriend.AddBlacklistReq) (*pbFriend.AddBlacklistResp, error) {
	log.Info("call rpc AddBlack args:", req.String())
	if !token_verify.CheckAccess(req.CommonId.OpUserId, req.CommonId.FromUserId) {
		log.Error("CheckAccess failed", req.CommonId.OpUserId, req.CommonId.FromUserId)
		return &pbFriend.AddBlacklistResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	black := &db.Black{
		OwnerUserId: req.CommonId.FromUserId,
		BlockUserId: req.CommonId.ToUserId,
		OpUserId:    req.CommonId.OpUserId,
	}
	err := controller.InsertIntoBlacklist(black)
	if err != nil {
		log.Error("InsertIntoBlacklist failed", err.Error())
		return &pbFriend.AddBlacklistResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	log.Info("rpc AddBlack return")
	return &pbFriend.AddBlacklistResp{CommonResp: &constant.PBCommonSuccessResp}, nil
}

type friendServer struct {
	pbFriend.UnimplementedFriendServer
	rpcPort         int
	rpcRegisterName string
	etcdSchema      string
	etcdAddr        []string
}

func NewFriendServer(port int) *friendServer {
	log.NewPrivateLog("friend")
	return &friendServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.FriendName,
		etcdSchema:      config.Config.Etcd.EtcdSchema,
		etcdAddr:        config.Config.Etcd.EtcdAddr,
	}
}

func (s *friendServer) Run() {
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

	pbFriend.RegisterFriendServer(server, s)
	err = etcdv3.RegisterEtcd(s.etcdSchema, s.etcdAddr, utils.ServerIP, s.rpcPort, s.rpcRegisterName, etcdv3.TIME_TO_LIVE)
	if err != nil {
		log.Error("friend RegisterEtcd failed", err.Error())
		return
	}
	log.Info("rpc friend register success")
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed ", err.Error())
		return
	}
}
