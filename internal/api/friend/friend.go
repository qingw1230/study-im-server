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
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestBindJSONError,
			Info:   err.Error(),
		})
		return
	}
	log.Info("AddFriend BindJSON success")

	ok, opUserID := token_verify.GetUserIDFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIDFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusInternalServerError, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.TokenUnknown,
			Info:   constant.TokenUnknownMsg.Error(),
		})
		return
	}
	req := &pbFriend.AddFriendReq{CommonID: &pbFriend.CommonID{}}
	copier.Copy(req.CommonID, &params)
	req.ReqMsg = params.ReqMsg
	req.CommonID.OpUserID = opUserID
	log.Info("AddFriend args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed ", err.Error())
		c.JSON(http.StatusInternalServerError, constant.CommonFailResp)
		return
	}
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.AddFriend(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constant.CommonFailResp)
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
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestBindJSONError,
			Info:   err.Error(),
		})
		return
	}
	log.Info("AddFriendResponse BindJSON success")

	ok, opUserID := token_verify.GetUserIDFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIDFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusInternalServerError, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.TokenUnknown,
			Info:   constant.TokenUnknownMsg.Error(),
		})
		return
	}
	req := &pbFriend.AddFriendResponseReq{}
	copier.Copy(req, &params)
	req.CommonID.OpUserID = opUserID
	log.Info("AddFriendResponse args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed ", err.Error())
		c.JSON(http.StatusInternalServerError, constant.CommonFailResp)
		return
	}
	client := pbFriend.NewFriendClient(conn)
	reply, err := client.AddFriendResponse(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constant.CommonFailResp)
		return
	}

	resp := base_info.AddFriendResponseResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
}
