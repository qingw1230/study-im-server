package friend

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbFriend "github.com/qingw1230/study-im-server/pkg/proto/friend"
)

func AddFriend(c *gin.Context) {
	log.Info("call api AddFriend")
	params := base_info.AddFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("AddFriend BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbFriend.AddFriendReq{CommonId: &pbFriend.CommonId{
		OpUserId:   opUserId,
		FromUserId: params.FromUserId,
		ToUserId:   params.ToUserId,
	}}
	req.ReqMsg = params.ReqMsg
	log.Info("AddFriend args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.FriendName)
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.AddFriend(context.Background(), req)
	if err != nil {
		log.Error("call rpc AddFriend failed", err.Error())
	}
	resp := base_info.AddFriendResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
	log.Info("api AddFriend return")
}

func AddFriendResponse(c *gin.Context) {
	log.Info("call api AddFriendResponse")
	params := base_info.AddFriendResponseReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("AddFriendResponse BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbFriend.AddFriendResponseReq{}
	copier.Copy(req, &params)
	req.CommonId.OpUserId = opUserId
	log.Info("AddFriendResponse args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.FriendName)
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.AddFriendResponse(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.AddFriendResponseResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
	log.Info("api AddFriendResponse return")
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

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.FriendName)
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
	log.Info("call api GetFriendList")
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

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.FriendName)
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
	log.Info("api GetFriendList return")
}

func GetFriendApplyList(c *gin.Context) {
	log.Info("call api GetFriendApplyList")
	params := base_info.GetFriendApplyListReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("GetFriendApplyList BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbFriend.GetFriendApplyListReq{CommonId: &pbFriend.CommonId{
		FromUserId: params.FromUserID,
		OpUserId:   opUserId,
	}}
	log.Info("GetFriendApplyList args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.FriendName)
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.GetFriendApplyList(context.Background(), req)
	if err != nil {
		log.Error("AddBlacklist failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := &base_info.GetFriendApplyListResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	resp.CommonResp.Data = reply.FriendRequestList
	c.JSON(http.StatusOK, resp)
	log.Info("api GetFriendApplyList return")
}

func AddBlacklist(c *gin.Context) {
	params := base_info.AddBlacklistReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("AddBlack BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbFriend.AddBlacklistReq{CommonId: &pbFriend.CommonId{}}
	copier.Copy(req.CommonId, &params)
	req.CommonId.OpUserId = opUserId
	log.Info("AddBlack args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.FriendName)
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.AddBlacklist(context.Background(), req)
	if err != nil {
		log.Error("AddBlacklist failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.AddBlacklistResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
}
