package account

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/captcha"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	rpc "github.com/qingw1230/study-im-server/pkg/proto/account"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Register(c *gin.Context) {
	params := base_info.RegisterReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("Register BindJSON success")

	// 校验验证码
	if !captcha.Captcha.Verify(params.CheckCodeId, params.CheckCode) {
		log.Error("Captcha.Verify failed ", params.Email)
		c.JSON(http.StatusOK, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestCheckCodeError,
			Info:   constant.RequestCheckCodeErrorInfo,
		})
		return
	}
	log.Info("Register Captcha.Verify success")

	req := &rpc.RegisterReq{}
	copier.Copy(req, &params)
	log.Info("Register rpc client.Register args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("grpc.NewClient failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := rpc.NewAccountClient(conn)
	reply, err := client.Register(context.Background(), req)
	if err != nil {
		log.Error("client.Register internal failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	if reply.CommonResp.Status == constant.Fail {
		log.Error("client.Register failed ", reply.CommonResp.Code)
		resp := base_info.CommonResp{}
		copier.Copy(&resp, reply.CommonResp)
		c.JSON(http.StatusOK, resp)
		return
	}
	log.Info("Register rpc client.Register call success")

	c.JSON(http.StatusOK, constant.CommonSuccessResp)
}

func Login(c *gin.Context) {
	params := base_info.LoginReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("Login BindJSON success")

	// 校验验证码
	if !captcha.Captcha.Verify(params.CheckCodeId, params.CheckCode) {
		log.Error("Captcha.Verify failed ", params.Email)
		c.JSON(http.StatusOK, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestCheckCodeError,
			Info:   constant.RequestCheckCodeErrorInfo,
		})
		return
	}
	log.Info("Login Captcha.Verify success")

	req := &rpc.LoginReq{}
	copier.Copy(&req, &params)
	log.Info("Login rpc client.Login args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("grpc.NewClient failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := rpc.NewAccountClient(conn)
	reply, err := client.Login(context.Background(), req)
	if err != nil {
		log.Error("client.Login internal failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	if reply.CommonResp.Status == constant.Fail {
		log.Error("client.Login failed ", reply.CommonResp.Code)
		resp := base_info.CommonResp{}
		copier.Copy(&resp, reply.CommonResp)
		c.JSON(http.StatusOK, resp)
		return
	}
	log.Info("Login rpc client.Login call success")

	c.JSON(http.StatusOK, constant.NewSuccessRespWithData(reply.UserInfo))
}

type Data struct {
	Id   string `json:"id"`
	B64S string `json:"b64s"`
}

func GetCheckCode(c *gin.Context) {
	id, b64s := captcha.Captcha.Generate()

	resp := base_info.CommonResp{
		Status: constant.Success,
		Code:   constant.NoError,
		Info:   constant.SuccessInfo,
		Data: Data{
			Id:   id,
			B64S: b64s,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func GetUserInfo(c *gin.Context) {
	params := base_info.GetUserInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("Login BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}

	req := &rpc.GetUserInfoReq{
		OpUserId: opUserId,
		UserId:   params.UserId,
	}
	log.Info("GetUserInfo rpc client.GetUserInfo args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10100", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("grpc.NewClient failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := rpc.NewAccountClient(conn)
	reply, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Error("client.GetUserInfo internal failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	log.Info("Login rpc client.Login call success")

	// reply 和 err 同时为 nil，说明用户不存在
	if reply == nil {
		c.JSON(http.StatusOK, constant.CommonSuccessResp)
		return
	}
	c.JSON(http.StatusOK, constant.NewSuccessRespWithData(reply.PublicUserInfo))
}
