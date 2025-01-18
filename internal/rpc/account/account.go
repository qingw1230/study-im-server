package account

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
	pbAccount "github.com/qingw1230/study-im-server/pkg/proto/account"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
)

func (s *accountServer) Register(_ context.Context, req *pbAccount.RegisterReq) (*pbAccount.RegisterResp, error) {
	log.Info("call Register args: ", req.String())

	// 确保用户不存在
	if controller.IsUserExist(req.Email) {
		log.Error("IsUserExist failed ", req.Email)
		resp := &pbAccount.RegisterResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordAlreadyExists,
				Info:   constant.RecordEmailAlreadyRegisterErrorInfo,
			},
		}
		return resp, nil
	}

	var user db.UserInfo
	copier.Copy(&user, req)
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

	log.Info("call Register success")
	resp := &pbAccount.RegisterResp{CommonResp: &pbPublic.CommonResp{}}
	copier.Copy(resp.CommonResp, constant.CommonSuccessResp)
	return resp, nil
}

func (s *accountServer) Login(_ context.Context, req *pbAccount.LoginReq) (*pbAccount.LoginResp, error) {
	log.Info("call Login args: ", req.String())

	// 确保用户存在
	user, err := controller.FindUserByEmail(req.UserLoginInfo.Email)
	if err == gorm.ErrRecordNotFound {
		resp := &pbAccount.LoginResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordNotExists,
				Info:   constant.RecordAccountORPwdErrorInfo,
			},
		}
		return resp, nil
	}
	if err != nil {
		log.Error("FindUserByEmail failed ", err.Error())
		return nil, err
	}

	if !utils.ValidPassword(req.UserLoginInfo.Password, user.Salt, user.Password) {
		resp := &pbAccount.LoginResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.RecordAccountORPwdError,
				Info:   constant.RecordAccountORPwdErrorInfo,
			},
		}
		return resp, nil
	}
	log.Debug("user passed ValidPassword ", user.Email)

	// TODO(qingw1230): 多设备登录检测
	token, _, err := token_verify.CreateToken(user.UserID)
	if err != nil {
		resp := &pbAccount.LoginResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Fail,
				Code:   constant.TokenError,
				Info:   constant.CreateTokenMsg.Error(),
			},
		}
		return resp, nil
	}

	resp := &pbAccount.LoginResp{
		CommonResp:           &pbPublic.CommonResp{},
		UserLoginSuccessInfo: &pbPublic.UserLoginSuccessInfo{},
	}
	copier.Copy(resp.CommonResp, &constant.CommonSuccessResp)
	copier.Copy(resp.UserLoginSuccessInfo, user)
	resp.UserLoginSuccessInfo.Token = token
	resp.UserLoginSuccessInfo.Admin = utils.IsContain(user.UserID, config.Config.Admin.UserIDs)
	return resp, nil
}

func (s *accountServer) GetUserInfo(_ context.Context, req *pbAccount.GetUserInfoReq) (*pbAccount.GetUserInfoResp, error) {
	log.Info("call GetUserInfo args: ", req.String())

	user, err := controller.FindUserByID(req.UserID)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		log.Error("FindUserByID failed ", err.Error())
		return nil, err
	}

	resp := &pbAccount.GetUserInfoResp{
		CommonResp:     &pbPublic.CommonResp{},
		PublicUserInfo: &pbPublic.PublicUserInfo{},
	}
	copier.Copy(resp.PublicUserInfo, user)
	return resp, nil
}

type accountServer struct {
	pbAccount.UnimplementedAccountServer
	rpcPort         int
	rpcRegisterName string
	zkSchema        string
	zkAddr          []string
}

func NewRpcAccountServer(port int) *accountServer {
	log.NewPrivateLog("account")
	return &accountServer{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.AccountName,
		zkSchema:        config.Config.Zookeeper.ZKSchema,
		zkAddr:          config.Config.Zookeeper.ZKAddr,
	}
}

func (s *accountServer) Run() {
	log.Info("rpc account start...")
	address := utils.ServerIP + ":" + strconv.Itoa(s.rpcPort)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Error("listen network failed ", err.Error(), address)
		return
	}
	defer ln.Close()

	server := grpc.NewServer()
	defer server.GracefulStop()

	pbAccount.RegisterAccountServer(server, s)
	// TODO(qingw1230): 将 rpc 服务注册进 zk
	err = server.Serve(ln)
	if err != nil {
		log.Error("Server failed ", err.Error())
		return
	}
}
