package msg

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbGroup "github.com/qingw1230/study-im-server/pkg/proto/group"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
	"github.com/qingw1230/study-im-server/pkg/utils"
)

func (s *msgServer) SendMsg(_ context.Context, req *pbMsg.SendMsgReq) (*pbMsg.SendMsgResp, error) {
	log.Info("call rpc SendMsg args:", req.String())
	req.MsgData.ServerMsgId = GetMsgId(req.MsgData.SendId)
	req.MsgData.SendTime = utils.GetCurrentTimestampByMill()
	msgToMq := pbMsg.MsgDataToMq{
		Token: req.Token,
	}
	if req.MsgData.ContentType == 102 || req.MsgData.ContentType == 104 || req.MsgData.ContentType == 105 {
		req.MsgData.Status = -1
	}
	resp := &pbMsg.SendMsgResp{CommonResp: &sdkws.CommonResp{}}

	switch req.MsgData.SessionType {
	case constant.SingleChatType:
		msgToMq.MsgData = req.MsgData
		seq, err := db.DB.IncrSeq(msgToMq.MsgData.RecvId)
		if err != nil {
			log.Error("redis IncrUserSeq failed", err.Error(), msgToMq.MsgData.RecvId)
			return returnMsg(resp, req, constant.Fail, constant.MsgUnknownError, constant.MsgUnknownErrorInfo, msgToMq.MsgData.ServerMsgId, 0, msgToMq.MsgData.SendTime)
		}
		msgToMq.MsgData.Seq = seq
		err1 := s.sendMsgToKafka(&msgToMq, msgToMq.MsgData.RecvId)
		if err1 != nil {
			log.Error("sendMsgToKafka error", msgToMq.MsgData.RecvId, msgToMq.String())
			return returnMsg(resp, req, constant.Fail, constant.MsgKafkaSendError, constant.MsgKafkaSendErrorInfo, "", 0, 0)
		}

		// 给自己发消息，kafka 中只需要存一份
		if msgToMq.MsgData.SendId != msgToMq.MsgData.RecvId {
			seq, err = db.DB.IncrSeq(msgToMq.MsgData.SendId)
			if err != nil {
				log.Error("redis IncrUserSeq failed", err.Error(), msgToMq.MsgData.SendId)
				return returnMsg(resp, req, constant.Fail, constant.MsgUnknownError, constant.MsgUnknownErrorInfo, msgToMq.MsgData.ServerMsgId, 0, msgToMq.MsgData.SendTime)
			}
			msgToMq.MsgData.Seq = seq
			err2 := s.sendMsgToKafka(&msgToMq, msgToMq.MsgData.SendId)
			if err2 != nil {
				log.Error("sendMsgToKafka error", msgToMq.MsgData.SendId, msgToMq.String())
				return returnMsg(resp, req, constant.Fail, constant.MsgKafkaSendError, constant.MsgKafkaSendErrorInfo, "", 0, 0)
			}
		}
		return returnMsg(resp, req, constant.Success, constant.NoError, constant.SuccessInfo, msgToMq.MsgData.ServerMsgId, seq, msgToMq.MsgData.SendTime)
	case constant.GroupChatType:
		conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.GroupName)
		client := pbGroup.NewGroupClient(conn)
		inReq := &pbGroup.GetGroupAllMemberReq{GroupId: req.MsgData.GroupId}
		reply, err := client.GetGroupAllMember(context.Background(), inReq)
		if err != nil {
			log.Error("rpc GetGroupAllMember failed", err.Error())
			return returnMsg(resp, req, constant.Fail, constant.MsgUnknownError, constant.MsgUnknownErrorInfo, msgToMq.MsgData.ServerMsgId, 0, msgToMq.MsgData.SendTime)
		}
		if reply.CommonResp.Code != constant.NoError {
			log.Error("rpc send_msg GetGroupAllMember failed", err.Error())
			return returnMsg(resp, req, constant.Fail, reply.CommonResp.Code, reply.CommonResp.Info, msgToMq.MsgData.ServerMsgId, 0, msgToMq.MsgData.SendTime)
		}

		seq, err := db.DB.IncrSeq(req.MsgData.GroupId)
		if err != nil {
			log.Error("redis IncrUserSeq failed", err.Error(), msgToMq.MsgData.SendId)
			return returnMsg(resp, req, constant.Fail, constant.MsgUnknownError, constant.MsgUnknownErrorInfo, msgToMq.MsgData.ServerMsgId, 0, msgToMq.MsgData.SendTime)
		}
		for _, v := range reply.MemberList {
			req.MsgData.RecvId = v.UserId
			msgToMq.MsgData = req.MsgData
			err := s.sendMsgToKafka(&msgToMq, v.UserId)
			if err != nil {
				log.Error("sendMsgToKafka error", msgToMq.MsgData.SendId, msgToMq.String())
				return returnMsg(resp, req, constant.Fail, constant.MsgKafkaSendError, constant.MsgKafkaSendErrorInfo, "", 0, 0)
			}
		}

		return returnMsg(resp, req, constant.Success, constant.NoError, constant.SuccessInfo, msgToMq.MsgData.ServerMsgId, seq, msgToMq.MsgData.SendTime)
	default:
		return returnMsg(resp, req, constant.Fail, constant.MsgUnknownError, constant.MsgUnknownErrorInfo, msgToMq.MsgData.ServerMsgId, 0, msgToMq.MsgData.SendTime)
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
	t := time.Now().UnixMilli()
	strTimestamp := strconv.FormatInt(t, 10)
	return strTimestamp + "-" + sendId + "-" + strconv.Itoa(rand.Int())
}

func returnMsg(resp *pbMsg.SendMsgResp, req *pbMsg.SendMsgReq, status string, code int32, info, serverMsgId string, seq uint64, sendTime int64) (*pbMsg.SendMsgResp, error) {
	resp.CommonResp.Status = status
	resp.CommonResp.Code = code
	resp.CommonResp.Info = info
	resp.Seq = seq
	resp.ServerMsgId = serverMsgId
	resp.SendTime = sendTime
	return resp, nil
}

type NotificationMsg struct {
	SendId      string
	RecvId      string
	MsgFrom     int32
	SessionType int32
	ContentType int32
	Content     []byte
}

func Notification(n *NotificationMsg) {
	msg := &sdkws.MsgData{
		SendId:      n.SendId,
		RecvId:      n.RecvId,
		MsgFrom:     n.MsgFrom,
		SessionType: n.SessionType,
		ContentType: n.ContentType,
		Content:     string(n.Content),
		CreateTime:  time.Now().UnixMilli(),
	}
	switch n.SessionType {
	case constant.GroupChatType:
		msg.RecvId = ""
		msg.GroupId = n.RecvId
	}
	sender, err := controller.GetUserById(msg.SendId)
	if err != nil {
		log.Error("GetUserById failed", err.Error(), msg.SendId)
	}
	if sender != nil {
		msg.SenderNickName = sender.NickName
		msg.SenderFaceUrl = sender.FaceUrl
	}

	req := &pbMsg.SendMsgReq{MsgData: msg}
	rpcConn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.OfflineMessageName)
	client := pbMsg.NewMsgClient(rpcConn)
	reply, err := client.SendMsg(context.Background(), req)
	if err != nil || reply.CommonResp.Code != constant.NoError {
		log.Error("rpc SendMsg failed", err.Error())
	}
}
