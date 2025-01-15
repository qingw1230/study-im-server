package account

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/captcha"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	rpc "github.com/qingw1230/study-im-server/pkg/proto/account"
	"github.com/qingw1230/study-im-server/pkg/proto/public"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Register(c *gin.Context) {
	params := base_info.RegisterReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestBindJSONError,
			Info:   err.Error(),
		})
		return
	}
	log.Info("Register BindJSON success")

	// 校验验证码
	if !captcha.Captcha.Verify(params.CheckCodeID, params.CheckCode) {
		log.Error("Captcha.Verify failed ", params.Email)
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestCheckCodeError,
			Info:   constant.CheckCodeInfo,
		})
		return
	}
	log.Info("Register Captcha.Verify success")

	req := &rpc.RegisterReq{UserRegisterInfo: &public.UserRegisterInfo{}}
	copier.Copy(req.UserRegisterInfo, &params)
	log.Info("Register rpc client.Register args", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("grpc.NewClient failed ", err.Error())
		c.JSON(http.StatusInternalServerError, constant.CommonFailResp)
		return
	}
	client := rpc.NewAccountClient(conn)
	reply, err := client.Register(context.Background(), req)
	if err != nil {
		log.Error("client.Register internal failed ", err.Error())
		c.JSON(http.StatusInternalServerError, constant.CommonFailResp)
		return
	}
	if reply.CommonResp.Status == constant.Fail {
		log.Error("client.Register failed ", reply.CommonResp.Code)
		resp := base_info.CommonResp{}
		copier.Copy(&resp, reply.CommonResp)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	log.Info("Register rpc client.Register call success")

	c.JSON(http.StatusOK, constant.CommonSuccessResp)
}

func Login(c *gin.Context) {
	params := base_info.LoginReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestParamsError,
			Info:   constant.FailInfo,
		})
		return
	}

	if !captcha.Captcha.Verify(params.CheckCodeID, params.CheckCode) {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestCheckCodeError,
			Info:   constant.FailInfo,
		})
		return
	}

	user, err := controller.FindUserByEmail(params.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RecordNotExists,
			Info:   constant.FailInfo,
		})
		return
	}

	if !utils.ValidPassword(params.Password, user.Salt, user.Password) {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RecordAccountORPwdError,
			Info:   constant.FailInfo,
		})
		return
	}

	flag := false
	for _, str := range config.Config.Admin.Emails {
		if str == params.Email {
			flag = true
			break
		}
	}

	token := utils.Md5Encode(user.UserID + utils.GenerateRandomStr(constant.LENGTH_20))

	// TODO(qingw1230): 多设备登录检测
	tokenStruct := base_info.TokenToRedis{
		Token:    token,
		UserID:   user.UserID,
		NickName: user.NickName,
		Admin:    flag,
	}
	db.DB.SetUserToken(tokenStruct)

	resp := base_info.LoginResp{
		Token:             token,
		Admin:             flag,
		UserID:            user.UserID,
		NickName:          user.NickName,
		PersonalSignature: user.PersonalSignature,
		JoinType:          user.JoinType,
		Sex:               user.Sex,
		AreaName:          user.AreaName,
		AreaCode:          user.AreaCode,
	}
	c.JSON(http.StatusOK, base_info.CommonResp{
		Status: constant.Success,
		Code:   constant.NoError,
		Info:   constant.SuccessInfo,
		Data:   resp,
	})
}

type Data struct {
	ID   string `json:"id"`
	B64S string `json:"b64s"`
}

func GetCheckCode(c *gin.Context) {
	id, b64s := captcha.Captcha.Generate()

	resp := base_info.CommonResp{
		Status: constant.Success,
		Code:   constant.NoError,
		Info:   constant.SuccessInfo,
		Data: Data{
			ID:   id,
			B64S: b64s,
		},
	}
	c.JSON(http.StatusOK, resp)
}
