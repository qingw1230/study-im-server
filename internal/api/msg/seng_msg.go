package msg

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type paramsUserSendMsg struct {
	SendId         string `json:"sendId" binding:"required"`
	SenderNickName string `json:"senderNickName"`
	SenderFaceUrl  string `json:"senderFaceUrl"`
	Data           struct {
		RecvId      string                    `json:"recvId" `
		GroupId     string                    `json:"groupId" `
		ClientMsgId string                    `json:"clientMsgId" binding:"required"`
		SessionType int32                     `json:"sessionType" binding:"required"`
		MsgFrom     int32                     `json:"msgFrom" binding:"required"`
		ContentType int32                     `json:"contentType" binding:"required"`
		Content     []byte                    `json:"content" binding:"required"`
		ForceList   []string                  `json:"forceList"`
		Options     map[string]bool           `json:"options" `
		CreateTime  int64                     `json:"createTime" binding:"required"`
		OfflineInfo *pbPublic.OfflinePushInfo `json:"offlineInfo"`
	}
}

func newUserSendMsgReq(token string, params *paramsUserSendMsg) *pbMsg.SendMsgReq {
	pbData := pbMsg.SendMsgReq{
		Token: token,
		MsgData: &pbPublic.MsgData{
			SendId:          params.SendId,
			RecvId:          params.Data.RecvId,
			GroupId:         params.Data.GroupId,
			ClientMsgId:     params.Data.ClientMsgId,
			SenderNickName:  params.SenderNickName,
			SenderFaceUrl:   params.SenderFaceUrl,
			SessionType:     params.Data.SessionType,
			MsgFrom:         params.Data.MsgFrom,
			ContentType:     params.Data.ContentType,
			Content:         params.Data.Content,
			CreateTime:      params.Data.CreateTime,
			Options:         params.Data.Options,
			OfflinePushInfo: params.Data.OfflineInfo,
		},
	}
	return &pbData
}

func SendMsg(c *gin.Context) {
	log.Info("call api SendMsg")
	params := paramsUserSendMsg{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("SendMsg BindJSON success")

	token := c.Request.Header.Get(constant.STR_TOKEN)
	req := newUserSendMsgReq(token, &params)
	log.Info("SendMsg args:", req.String())
	// TODO(qingw1230): 使用服务发现建立连接
	conn, err := grpc.NewClient("127.0.0.1:10300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}
	client := pbMsg.NewMsgClient(conn)
	reply, err := client.SendMsg(context.Background(), req)
	if err != nil {
		log.Error("SendMsg failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.CommonResp{}
	copier.Copy(&resp, reply.CommonResp)
	resp.Data = gin.H{
		"serverMsgID": reply.ServerMsgId,
		"clientMsgID": reply.ClientMsgId,
		"sendTime":    reply.SendTime,
	}
	c.JSON(http.StatusOK, resp)
	log.Info("api SendMsg return")
}
