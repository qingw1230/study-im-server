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
	rpc "github.com/qingw1230/study-im-server/pkg/proto/friend"
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
		log.Error("token_verify.GetUserIDFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.TokenUnknown,
			Info:   constant.TokenUnknownMsg.Error(),
		})
		return
	}
	req := &rpc.AddFriendReq{}
	copier.Copy(req, &params)
	req.OpUserID = opUserID
	log.Info("client.AddFriend args: ", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("grpc.NewClient failed ", err.Error())
		c.JSON(http.StatusInternalServerError, constant.CommonFailResp)
		return
	}
	client := rpc.NewFriendClient(conn)
	reply, err := client.AddFriend(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, constant.CommonFailResp)
		return
	}

	resp := base_info.CommonResp{}
	copier.Copy(&resp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
}
