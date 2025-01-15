package account

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
	pbAccount "github.com/qingw1230/study-im-server/pkg/proto/account"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (rpc *rpcAccout) Register(_ context.Context, req *pbAccount.RegisterReq) (*pbAccount.RegisterResp, error) {
	log.Info("call rpcAccount.Register, args: ", req.String())

	// 确保用户不存在
	if controller.IsUserExist(req.UserRegisterInfo.Email) {
		log.Error("controller.IsUserExist failed ", req.UserRegisterInfo.Email)
		resp := &pbAccount.RegisterResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordAlreadyExists,
				Info:   constant.EmailAlreadyRegisterInfo,
			},
		}
		return resp, nil
	}

	var user db.UserInfo
	copier.Copy(&user, req.UserRegisterInfo)
	user.UserID = utils.GenerateRandomID(constant.LENGTH_11)
	user.Salt = utils.GenerateRandomStr(constant.LENGTH_10)
	user.Password = utils.MakePassword(user.Password, user.Salt)
	user.JoinType = constant.UserInfoJoinTypeAPPLY
	user.Status = constant.UserInfoStatusENABLE
	now := int(time.Now().Unix())
	user.CreateTime = now
	user.LastLoginTime = now
	user.LastOffTime = now

	err := controller.CreateUser(user)
	if err != nil {
		return nil, err
	}

	log.Info("call rpcAccount.Register success")
	resp := &pbAccount.RegisterResp{CommonResp: &pbPublic.CommonResp{}}
	copier.Copy(resp.CommonResp, constant.CommonSuccessResp)
	return resp, nil
}

type rpcAccout struct {
	pbAccount.UnimplementedAccountServer
	rpcPort         int
	rpcRegisterName string
	zkSchema        string
	zkAddr          []string
}

func NewRpcAccountServer(port int) *rpcAccout {
	log.NewPrivateLog("account")
	return &rpcAccout{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.AccountName,
		zkSchema:        config.Config.Zookeeper.ZKSchema,
		zkAddr:          config.Config.Zookeeper.ZKAddr,
	}
}

func (rpc *rpcAccout) Run() {
	log.Info("rpc account start...")
	address := utils.ServerIP + ":" + strconv.Itoa(rpc.rpcPort)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("listen network failed ", err.Error(), address)
	}

	server := grpc.NewServer()
	defer server.GracefulStop()

	pbAccount.RegisterAccountServer(server, rpc)
	// TODO(qingw1230): 将 rpc 服务注册进 zk
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed ", err.Error())
		return
	}
}
