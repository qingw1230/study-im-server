package group

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
	pbGroup "github.com/qingw1230/study-im-server/pkg/proto/group"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
)

func CreateGroup(c *gin.Context) {
	log.Info("call api CreateGroup")
	params := base_info.CreateGroupReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("CreateGroup BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbGroup.CreateGroupReq{GroupInfo: &pbPublic.GroupInfo{}}
	copier.Copy(req.GroupInfo, &params)
	req.GroupInfo.CreateUserId = opUserId
	req.OpUserId = opUserId
	req.OwnerUserId = params.OwnerUserId
	log.Info("CreateGroup args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.GroupName)
	client := pbGroup.NewGroupClient(conn)
	reply, err := client.CreateGroup(context.Background(), req)
	if err != nil {
		log.Error("call rpc CreateGroup failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.CreateGroupResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
	log.Info("api CreateGroup return")
}

func DeleteGroup(c *gin.Context) {
	log.Info("call api DeleteGroup")
	params := base_info.DeleteGroupReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("DeleteGroup BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbGroup.DeleteGroupReq{
		GroupId:  params.GroupId,
		OpUserId: opUserId,
	}
	log.Info("DeleteGroup args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.GroupName)
	client := pbGroup.NewGroupClient(conn)
	reply, err := client.DeleteGroup(context.Background(), req)
	if err != nil {
		log.Error("DeleteGroup failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.DeleteFriendResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
	log.Info("api DeletGroup return")
}

func QuitGroup(c *gin.Context) {
	log.Info("call api QuitGroup")
	params := base_info.QuitGroupReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("QuitGroup BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbGroup.QuitGroupReq{
		GroupId:  params.GroupId,
		OpUserId: opUserId,
	}
	log.Info("QuitGroup args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.GroupName)
	client := pbGroup.NewGroupClient(conn)
	reply, err := client.QuitGroup(context.Background(), req)
	if err != nil {
		log.Error("QuitGroup failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.QuitGroupResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
	log.Info("api QuitGroup return")
}

func GetJoinedGroupList(c *gin.Context) {
	params := base_info.GetJoinedGroupListReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("GetJoinedGroupList BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbGroup.GetJoinedGroupListReq{
		FromUserId: params.FromUserId,
		RoleLevel:  params.RoleLevel,
		OpUserId:   opUserId,
	}
	log.Info("GetJoinedGroupList args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.GroupName)
	client := pbGroup.NewGroupClient(conn)
	reply, err := client.GetJoinedGroupList(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.GetJoinedGroupListResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	resp.CommonResp.Data = reply.GroupInfoList
	c.JSON(http.StatusOK, resp)
}

func GetGroupInfo(c *gin.Context) {
	params := base_info.GetGroupInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("GetGroupInfo BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbGroup.GetGroupInfoReq{
		GroupId:  params.GroupId,
		OpUserId: opUserId,
	}
	log.Info("GetGroupInfo args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.GroupName)
	client := pbGroup.NewGroupClient(conn)
	reply, err := client.GetGroupInfo(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.GetGroupInfoResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	resp.Data = reply.GroupInfo
	c.JSON(http.StatusOK, resp)
}

func SetGroupInfo(c *gin.Context) {
	log.Info("call api SetGroupInfo")
	params := base_info.SetGroupInfoReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("SetGroupInfo BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbGroup.SetGroupInfoReq{GroupInfo: &pbPublic.GroupInfo{}}
	copier.Copy(req.GroupInfo, &params)
	req.OpUserId = opUserId
	log.Info("SetGroupInfo args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.GroupName)
	client := pbGroup.NewGroupClient(conn)
	reply, err := client.SetGroupInfo(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.SetGroupInfoResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
	log.Info("api SetGroupInfo return")
}
