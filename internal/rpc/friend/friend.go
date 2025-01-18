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
	pbFriend "github.com/qingw1230/study-im-server/pkg/proto/friend"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *friendServer) AddFriend(_ context.Context, req *pbFriend.AddFriendReq) (*pbFriend.AddFriendResp, error) {
	log.Info("call AddFriend args: ", req.String())
	// 确保有权限
	if !token_verify.CheckAccess(req.OpUserID, req.FromUserID) {
		log.Error("CheckAccess false ", req.OpUserID, req.FromUserID)
		return &pbFriend.AddFriendResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RequestTokenAccessError,
				Info:   constant.RequestTokenAccessErrorInfo,
			},
		}, nil
	}

	// 保证要添加的好友存在
	if _, err := controller.FindUserByID(req.ToUserID); err != nil {
		return &pbFriend.AddFriendResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordNotExists,
				Info:   constant.RecordAccountNotExistsInfo,
			},
		}, nil
	}

	friendRequest := db.FriendRequest{HandleResult: 0}
	copier.Copy(&friendRequest, req)
	err := controller.InsertFriendApplication(&friendRequest)
	if err != nil {
		log.Error("InsertFriendApplication failed ", err.Error(), friendRequest)
		return &pbFriend.AddFriendResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordMySQLCommonError,
				Info:   constant.RecordMySQLCommonErrorInfo,
			},
		}, nil
	}

	// TODO(qingw1230): 给被添加方发送通知
	resp := &pbFriend.AddFriendResp{CommonResp: &pbPublic.CommonResp{}}
	copier.Copy(resp.CommonResp, constant.CommonSuccessResp)
	return resp, nil
}

func (s *friendServer) AddFriendResponse(_ context.Context, req *pbFriend.AddFriendResponseReq) (*pbFriend.AddFriendResponseResp, error) {
	log.Info("call AddFriendResponse args: ", req.String())
	// 确保有权限
	if !token_verify.CheckAccess(req.OpUserID, req.FromUserID) {
		log.Error("CheckAccess false ", req.OpUserID, req.FromUserID)
		return &pbFriend.AddFriendResponseResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RequestTokenAccessError,
				Info:   constant.RequestTokenAccessErrorInfo,
			},
		}, nil
	}

	// 在同意或拒绝好友申请之前，先检查记录是否存在
	friendRequest, err := controller.GetFriendApplicationByBothUserID(req.FromUserID, req.ToUserID)
	if err != nil {
		return &pbFriend.AddFriendResponseResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordNotExists,
				Info:   constant.RecordNotExistsInfo,
			},
		}, nil
	}
	friendRequest.HandleResult = int8(req.HandleResult)
	friendRequest.HandleTime = time.Now()
	friendRequest.HandleUserID = req.OpUserID
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
		friend, err := controller.GetFriendRelationFromFriend(req.FromUserID, req.ToUserID)
		log.Debug("GetFriendRelationFromFriend return ", friend, err)
		if err == nil {
			log.Warn("GetFriendRelationFromFriend exist ", req.FromUserID, req.ToUserID)
		} else if err == gorm.ErrRecordNotFound {
			toInsertFollow := db.Friend{OwnerUserID: req.FromUserID, FriendUserID: req.ToUserID, OpUserID: req.OpUserID}
			err = controller.InsertToFriend(&toInsertFollow)
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

		_, err = controller.GetFriendRelationFromFriend(req.ToUserID, req.FromUserID)
		if err == nil {
			log.Warn("GetFriendRelationFromFriend exist ", req.ToUserID, req.FromUserID)
		} else if err == gorm.ErrRecordNotFound {
			toInsertFollow := db.Friend{OwnerUserID: req.ToUserID, FriendUserID: req.FromUserID, OpUserID: req.OpUserID}
			err = controller.InsertToFriend(&toInsertFollow)
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
	resp := &pbFriend.AddFriendResponseResp{CommonResp: &pbPublic.CommonResp{}}
	copier.Copy(resp.CommonResp, constant.CommonSuccessResp)
	return resp, nil
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
