package msg

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"github.com/qingw1230/study-im-server/pkg/utils"
)

func (s *msgServer) encapsulateMsgData(data *pbPublic.MsgData) {
	data.ServerMsgId = GetMsgId(data.SendId)
	data.SendTime = utils.GetCurrentTimestampByMill()
}

func (s *msgServer) SendMsg(_ context.Context, req *pbMsg.SendMsgReq) (*pbMsg.SendMsgResp, error) {
	log.Info("call rpc SendMsg")
	s.encapsulateMsgData(req.MsgData)
	msgToMq := pbMsg.MsgDataToMq{
		Token: req.Token,
	}
	resp := &pbMsg.SendMsgResp{CommonResp: &pbPublic.CommonResp{}}

	switch req.MsgData.SessionType {
	case constant.SingleChatType:
		msgToMq.MsgData = req.MsgData
		err1 := s.sendMsgToKafka(&msgToMq, msgToMq.MsgData.RecvId)
		if err1 != nil {
			log.Error("sendMsgToKafka error", msgToMq.MsgData.RecvId, msgToMq.String())
			return returnMsg(resp, req, constant.Fail, constant.MsgKafkaSendError, constant.MsgKafkaSendErrorInfo, "", 0)
		}

		// 给自己发消息，kafka 中只需要存一份
		if msgToMq.MsgData.SendId != msgToMq.MsgData.RecvId {
			err2 := s.sendMsgToKafka(&msgToMq, msgToMq.MsgData.SendId)
			if err2 != nil {
				log.Error("sendMsgToKafka error", msgToMq.MsgData.SendId, msgToMq.String())
				return returnMsg(resp, req, constant.Fail, constant.MsgKafkaSendError, constant.MsgKafkaSendErrorInfo, "", 0)
			}
		}
		log.Info("rpc SendMsg success")
		return returnMsg(resp, req, constant.Success, constant.NoError, constant.SuccessInfo, msgToMq.MsgData.ServerMsgId, msgToMq.MsgData.SendTime)
	default:
		return returnMsg(resp, req, constant.Fail, constant.MsgUnknownError, constant.MsgUnknownErrorInfo, msgToMq.MsgData.ServerMsgId, msgToMq.MsgData.SendTime)
	}
}

// sendMsgToKafka 将数据发送到 kafka，key 为 userId
func (s *msgServer) sendMsgToKafka(m *pbMsg.MsgDataToMq, key string) error {
	partition, offset, err := s.producer.SendMessage(m, key)
	if err != nil {
		log.Error("kafka SendMessage failed", m.String(), partition, offset, err.Error(), key)
	}
	return err
}

// GetMsgId 使用时间戳、sendId 和随机值生成
func GetMsgId(sendId string) string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return t + "-" + sendId + "-" + strconv.Itoa(rand.Int())
}

func returnMsg(resp *pbMsg.SendMsgResp, req *pbMsg.SendMsgReq, status string, code int32, info, serverMsgId string, sendTime int64) (*pbMsg.SendMsgResp, error) {
	resp.CommonResp.Status = status
	resp.CommonResp.Code = code
	resp.CommonResp.Info = info
	resp.ServerMsgId = serverMsgId
	resp.ClientMsgId = req.MsgData.ClientMsgId
	resp.SendTime = sendTime
	return resp, nil
}

func modifyMessageByUserMessageReceiveOpt(userId, sourceId string, sessionType int, req *pbMsg.SendMsgReq) bool {
	conversationId := utils.GetConversationIdBySessionType(sourceId, sessionType)
	opt, err := db.DB.GetSingleConversationMsgOpt(userId, conversationId)
	if err != nil && err != redis.ErrNil {
		log.Error("GetSingleConversationMsgOpt from redis error", conversationId, req.String(), err.Error())
		return true
	}

	switch opt {
	case constant.ReceiveMessage:
		return true
	case constant.NotReceiveMessage:
		return false
	}
	return true
}
