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

func newPbSendMsgReq(token string, params *base_info.SendMsgReq) *pbMsg.SendMsgReq {
	pbData := pbMsg.SendMsgReq{
		Token: token,
		MsgData: &pbPublic.MsgData{
			SendId:         params.SendId,
			RecvId:         params.RecvId,
			GroupId:        params.GroupId,
			SenderNickName: params.SenderNickName,
			SenderFaceUrl:  params.SenderFaceUrl,
			SessionType:    params.SessionType,
			MsgFrom:        params.MsgFrom,
			ContentType:    params.ContentType,
			Content:        params.Content,
			CreateTime:     params.CreateTime,
		},
	}
	return &pbData
}

func SendMsg(c *gin.Context) {
	log.Info("call api SendMsg")
	params := base_info.SendMsgReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("SendMsg BindJSON success")

	token := c.Request.Header.Get(constant.STR_TOKEN)
	req := newPbSendMsgReq(token, &params)
	log.Info("SendMsg args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.OfflineMessageName)
	client := pbMsg.NewMsgClient(conn)
	reply, err := client.SendMsg(context.Background(), req)
	if err != nil {
		log.Error("SendMsg failed", err.Error())
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.SendMsgResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	copier.Copy(&resp, reply)
	c.JSON(http.StatusOK, resp)
	log.Info("api SendMsg return")
}
