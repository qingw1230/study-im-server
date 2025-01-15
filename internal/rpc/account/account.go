package account

import (
	"context"
	"net"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"github.com/qingw1230/study-im-server/pkg/base_info"
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

func (rpc *rpcAccount) Register(_ context.Context, req *pbAccount.RegisterReq) (*pbAccount.RegisterResp, error) {
	log.Info("call rpcAccount.Register, args: ", req.String())

	// 确保用户不存在
	if controller.IsUserExist(req.UserRegisterInfo.Email) {
		log.Error("controller.IsUserExist failed ", req.UserRegisterInfo.Email)
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

func (rpc *rpcAccount) Login(_ context.Context, req *pbAccount.LoginReq) (*pbAccount.LoginResp, error) {
	log.Info("call rpcAccount.Login, args: ", req.String())

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
		log.Error("controller.FindUserByEmail failed ", err.Error())
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
	log.Debug("user passed utils.ValidPassword ", user.Email)

	flag := false
	for _, str := range config.Config.Admin.Emails {
		if str == req.UserLoginInfo.Email {
			flag = true
			break
		}
	}
	// TODO(qingw1230): 改用 JWT
	token := utils.Md5Encode(user.UserID + utils.GenerateRandomStr(constant.LENGTH_20))
	// TODO(qingw1230): 多设备登录检测
	tokenStruct := base_info.TokenToRedis{
		Token:    token,
		UserID:   user.UserID,
		NickName: user.NickName,
		Admin:    flag,
	}
	db.DB.SetUserToken(tokenStruct)

	resp := &pbAccount.LoginResp{
		CommonResp:           &pbPublic.CommonResp{},
		UserLoginSuccessInfo: &pbPublic.UserLoginSuccessInfo{},
	}
	copier.Copy(resp.CommonResp, &constant.CommonSuccessResp)
	copier.Copy(resp.UserLoginSuccessInfo, user)
	resp.UserLoginSuccessInfo.Token = token
	resp.UserLoginSuccessInfo.Admin = flag
	return resp, nil
}

func (rpc *rpcAccount) GetUserInfo(_ context.Context, req *pbAccount.GetUserInfoReq) (*pbAccount.GetUserInfoResp, error) {
	log.Info("call rpcAccount.GetUserInfo, args: ", req.String())

	user, err := controller.FindUserByID(req.UserID)
	if err == gorm.ErrRecordNotFound {
		resp := &pbAccount.GetUserInfoResp{
			CommonResp: &pbPublic.CommonResp{
				Status: constant.Success,
				Code:   constant.RecordNotExists,
				Info:   constant.RecordAccountNotExistsInfo,
			},
		}
		return resp, nil
	}
	if err != nil {
		log.Error("controller.FindUserByID failed ", err.Error())
		return nil, err
	}

	resp := &pbAccount.GetUserInfoResp{
		CommonResp:     &pbPublic.CommonResp{},
		PublicUserInfo: &pbPublic.PublicUserInfo{},
	}
	copier.Copy(resp.CommonResp, &constant.CommonSuccessResp)
	copier.Copy(resp.PublicUserInfo, user)
	return resp, nil
}

type rpcAccount struct {
	pbAccount.UnimplementedAccountServer
	rpcPort         int
	rpcRegisterName string
	zkSchema        string
	zkAddr          []string
}

func NewRpcAccountServer(port int) *rpcAccount {
	log.NewPrivateLog("account")
	return &rpcAccount{
		rpcPort:         port,
		rpcRegisterName: config.Config.RpcRegisterName.AccountName,
		zkSchema:        config.Config.Zookeeper.ZKSchema,
		zkAddr:          config.Config.Zookeeper.ZKAddr,
	}
}

func (rpc *rpcAccount) Run() {
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
