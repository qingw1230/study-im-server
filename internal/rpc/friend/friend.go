package friend

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
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

func (s *friendServer) AddFriend(ctx context.Context, req *pbFriend.AddFriendReq) (*pbFriend.AddFriendResp, error) {
	log.Info("call AddFriend, args: ", req.String())
	// 确保有权限
	ok := token_verify.CheckAccess(req.OpUserID, req.FromUserID)
	if !ok {
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

	friendRequest := db.FriendRequest{
		HandleResult: 0,
		CreateTime:   time.Now(),
	}
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
