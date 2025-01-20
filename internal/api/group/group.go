package group

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	pbGroup "github.com/qingw1230/study-im-server/pkg/proto/group"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateGroup(c *gin.Context) {
	params := base_info.CreateGroupReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("CreateGroup BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed ", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbGroup.CreateGroupReq{GroupInfo: &pbPublic.GroupInfo{}}
	copier.Copy(req.GroupInfo, &params)
	req.GroupInfo.CreateUserId = opUserId
	req.OpUserId = opUserId
	req.OwnerUserId = params.OwnerUserId
	log.Info("CreateGroup args:", req.String())

	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10500", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := pbGroup.NewGroupClient(conn)
	reply, err := client.CreateGroup(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.CreateGroupResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	c.JSON(http.StatusOK, resp)
}
