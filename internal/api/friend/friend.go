package friend

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	pbFriend "github.com/qingw1230/study-im-server/pkg/proto/friend"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AddFriend(c *gin.Context) {
	params := base_info.AddFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("AddFriend BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbFriend.AddFriendReq{CommonId: &pbFriend.CommonId{}}
	copier.Copy(req.CommonId, &params)
	req.ReqMsg = params.ReqMsg
	req.CommonId.OpUserId = opUserId
	log.Info("AddFriend args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.AddFriend(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.AddFriendResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
}

func AddFriendResponse(c *gin.Context) {
	params := base_info.AddFriendResponseReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("AddFriendResponse BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbFriend.AddFriendResponseReq{}
	copier.Copy(req, &params)
	req.CommonId.OpUserId = opUserId
	log.Info("AddFriendResponse args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.AddFriendResponse(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.AddFriendResponseResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
}

func DeleteFriend(c *gin.Context) {
	params := base_info.DeleteFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("DeleteFriend BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbFriend.DeleteFriendReq{CommonId: &pbFriend.CommonId{}}
	copier.Copy(req.CommonId, &params)
	req.CommonId.OpUserId = opUserId
	log.Info("DeleteFriend args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.DeleteFriend(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := &base_info.DeleteFriendResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(resp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
}

func GetFriendList(c *gin.Context) {
	params := base_info.GetFriendListReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("GetFriendList BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbFriend.GetFriendListReq{CommonId: &pbFriend.CommonId{}}
	req.CommonId.FromUserId = params.FromUserId
	req.CommonId.OpUserId = opUserId
	log.Info("GetFriendList args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed ", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.GetFriendList(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.GetFriendListResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	resp.CommonResp.Data = reply.FriendInfoList
	c.JSON(http.StatusOK, resp)
}
