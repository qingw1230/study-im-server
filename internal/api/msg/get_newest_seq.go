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
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
)

func GetNewestSeq(c *gin.Context) {
	log.Info("call api GetNewestSeq")
	params := base_info.UserGetNewestSeqReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed", err.Error())
		c.JSON(http.StatusOK, constant.NewBindJSONErrorRespWithInfo(err.Error()))
		return
	}
	log.Info("GetNewestSeq BindJSON success")

	ok, opUserId := token_verify.GetUserIdFromToken(c.Request.Header.Get(constant.STR_TOKEN))
	if !ok {
		log.Error("GetUserIdFromToken failed", c.Request.Header.Get(constant.STR_TOKEN))
		c.JSON(http.StatusOK, constant.NewRespNoData(constant.Fail, constant.TokenUnknown, constant.TokenUnknownMsg.Error()))
		return
	}
	req := &pbMsg.GetNewestSeqReq{
		UserId:   params.UserId,
		OpUserId: opUserId,
	}
	log.Info("GetNewestSeq args:", req.String())

	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.OfflineMessageName)
	client := pbMsg.NewMsgClient(conn)
	reply, err := client.GetNewestSeq(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusOK, constant.CommonFailResp)
		return
	}

	resp := base_info.UserGetNewestSeqResp{CommonResp: base_info.CommonResp{}}
	copier.Copy(&resp.CommonResp, reply.CommonResp)
	resp.CommonResp.Data = reply.NewestSeq
	resp.ReqIdentifier = params.ReqIdentifier
	c.JSON(http.StatusOK, resp)
	log.Info("api GetNewestSeq return")
}
