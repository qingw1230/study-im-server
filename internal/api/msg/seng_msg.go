package msg

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
)

type paramsUserSendMsg struct {
	SendId         string `json:"sendId" binding:"required"`
	SenderNickName string `json:"senderNickName"`
	SenderFaceUrl  string `json:"senderFaceUrl"`
	Data           struct {
		RecvId      string `json:"recvId" `
		GroupId     string `json:"groupId" `
		SessionType int32  `json:"sessionType" binding:"required"`
		MsgFrom     int32  `json:"msgFrom" binding:"required"`
		ContentType int32  `json:"contentType" binding:"required"`
		Content     string `json:"content" binding:"required"`
		CreateTime  int64  `json:"createTime" binding:"required"`
	}
}

func newUserSendMsgReq(token string, params *paramsUserSendMsg) *pbMsg.SendMsgReq {
	pbData := pbMsg.SendMsgReq{
		Token: token,
		MsgData: &pbPublic.MsgData{
			SendId:         params.SendId,
			RecvId:         params.Data.RecvId,
			GroupId:        params.Data.GroupId,
			SenderNickName: params.SenderNickName,
			SenderFaceUrl:  params.SenderFaceUrl,
			SessionType:    params.Data.SessionType,
			MsgFrom:        params.Data.MsgFrom,
			ContentType:    params.Data.ContentType,
			Content:        params.Data.Content,
			CreateTime:     params.Data.CreateTime,
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
	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.OfflineMessageName)
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
		"sendTime":    reply.SendTime,
	}
	c.JSON(http.StatusOK, resp)
	log.Info("api SendMsg return")
}
