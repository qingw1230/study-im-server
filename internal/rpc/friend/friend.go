package friend

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	cp "github.com/qingw1230/study-im-server/pkg/common/utils"
	pbFriend "github.com/qingw1230/study-im-server/pkg/proto/friend"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *friendServer) AddFriend(_ context.Context, req *pbFriend.AddFriendReq) (*pbFriend.AddFriendResp, error) {
	log.Info("call AddFriend args: ", req.String())
	// 确保有权限
	if !token_verify.CheckAccess(req.CommonId.OpUserId, req.CommonId.FromUserId) {
		log.Error("CheckAccess false ", req.CommonId.OpUserId, req.CommonId.FromUserId)
		return &pbFriend.AddFriendResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	// 保证要添加的好友存在
	if _, err := controller.FindUserById(req.CommonId.ToUserId); err != nil {
		return &pbFriend.AddFriendResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordNotExists,
				Info:   constant.RecordAccountNotExistsInfo,
			},
		}, nil
	}

	friendRequest := db.FriendRequest{HandleResult: 0}
	copier.Copy(&friendRequest, req.CommonId)
	friendRequest.ReqMsg = req.ReqMsg
	err := controller.InsertFriendApplication(&friendRequest)
	if err != nil {
		log.Error("InsertFriendApplication failed ", err.Error(), friendRequest)
		return &pbFriend.AddFriendResp{CommonResp: &constant.PBMySQLCommonFailResp}, nil
	}

	// TODO(qingw1230): 给被添加方发送通知
	return &pbFriend.AddFriendResp{CommonResp: &constant.PBCommonSuccessResp}, nil
}

func (s *friendServer) AddFriendResponse(_ context.Context, req *pbFriend.AddFriendResponseReq) (*pbFriend.AddFriendResponseResp, error) {
	log.Info("call AddFriendResponse args: ", req.String())
	// 确保有权限
	if !token_verify.CheckAccess(req.CommonId.OpUserId, req.CommonId.FromUserId) {
		log.Error("CheckAccess false ", req.CommonId.OpUserId, req.CommonId.FromUserId)
		return &pbFriend.AddFriendResponseResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	// 在同意或拒绝好友申请之前，先检查记录是否存在
	friendRequest, err := controller.GetFriendApplicationByBothUserId(req.CommonId.FromUserId, req.CommonId.ToUserId)
	if err != nil {
		return &pbFriend.AddFriendResponseResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordNotExists,
				Info:   constant.RecordNotExistsInfo,
			},
		}, nil
	}
	friendRequest.HandleResult = req.HandleResult
	friendRequest.HandleTime = time.Now()
	friendRequest.HandleUserId = req.CommonId.OpUserId
	friendRequest.HandleMsg = req.HandleMsg
	err = controller.UpdateFriendApplication(friendRequest)
	if err != nil {
		resp := &pbFriend.AddFriendResponseResp{CommonResp: &pbPublic.CommonResp{}}
		copier.Copy(resp.CommonResp, constant.MySQLCommonFailResp)
		return resp, nil
	}
	log.Info("UpdateFriendApplication success ", friendRequest)

	if friendRequest.HandleResult == constant.FriendResponseAgree {
		// 插入两条单向好友关系
		friend, err := controller.GetFriendRelationFromFriend(req.CommonId.FromUserId, req.CommonId.ToUserId)
		log.Debug("GetFriendRelationFromFriend return ", friend, err)
		if err == nil {
			log.Warn("GetFriendRelationFromFriend exist ", req.CommonId.FromUserId, req.CommonId.ToUserId)
		} else if err == gorm.ErrRecordNotFound {
			toInsertFollow := db.Friend{OwnerUserId: req.CommonId.FromUserId, FriendUserId: req.CommonId.ToUserId, OpUserId: req.CommonId.OpUserId}
			err = controller.InsertIntoFriend(&toInsertFollow)
			if err != nil {
				log.Error("InsertToFriend failed ", err.Error(), toInsertFollow)
				resp := &pbFriend.AddFriendResponseResp{CommonResp: &pbPublic.CommonResp{}}
				copier.Copy(resp.CommonResp, constant.MySQLCommonFailResp)
				return resp, nil
			}
		} else {
			log.Error("GetFriendRelationFromFriend failed ", err.Error())
			resp := &pbFriend.AddFriendResponseResp{CommonResp: &pbPublic.CommonResp{}}
			copier.Copy(resp.CommonResp, constant.MySQLCommonFailResp)
			return resp, nil
		}

		_, err = controller.GetFriendRelationFromFriend(req.CommonId.ToUserId, req.CommonId.FromUserId)
		if err == nil {
			log.Warn("GetFriendRelationFromFriend exist ", req.CommonId.ToUserId, req.CommonId.FromUserId)
		} else if err == gorm.ErrRecordNotFound {
			toInsertFollow := db.Friend{OwnerUserId: req.CommonId.ToUserId, FriendUserId: req.CommonId.FromUserId, OpUserId: req.CommonId.OpUserId}
			err = controller.InsertIntoFriend(&toInsertFollow)
			if err != nil {
				log.Error("InsertToFriend failed ", err.Error(), toInsertFollow)
				resp := &pbFriend.AddFriendResponseResp{CommonResp: &pbPublic.CommonResp{}}
				copier.Copy(resp.CommonResp, constant.MySQLCommonFailResp)
				return resp, nil
			}
			// TODO(qingw1230): 通知另一方
		} else {
			log.Error("GetFriendRelationFromFriend failed ", err.Error())
			resp := &pbFriend.AddFriendResponseResp{CommonResp: &pbPublic.CommonResp{}}
			copier.Copy(resp.CommonResp, constant.MySQLCommonFailResp)
			return resp, nil
		}
	}

	// TODO(qingw1230): 通知或拒绝的通知
	return &pbFriend.AddFriendResponseResp{CommonResp: &constant.PBCommonSuccessResp}, nil
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
	log.Info("call GetFriendList args: ", req.String())
	// 确保有权限
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
	return &pbFriend.GetFriendListResp{
		CommonResp:     &constant.PBCommonSuccessResp,
		FriendInfoList: userInfoList,
	}, nil
}

type friendServer struct {
	pbFriend.UnimplementedFriendServer
	rpcPort         int
	rpcRegisterName string
	zkSchema        string
	zkAddr          []string
}

func NewFriendServer(port int) *friendServer {
	log.NewPrivateLog("friend")
	return &friendServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.FriendName,
		zkSchema:        config.Config.Zookeeper.ZKSchema,
		zkAddr:          config.Config.Zookeeper.ZKAddr,
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
	// TODO(qingw1230): 将 rpc 服务注册进 zk
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed ", err.Error())
		return
	}
}
