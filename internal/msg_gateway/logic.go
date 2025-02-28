package msg_gateway

import (
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbConversation "github.com/qingw1230/study-im-server/pkg/proto/conversation"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
)

func (ws *WsServer) msgParse(conn *UserConn, binaryMsg []byte) {
	m := Req{}
	err := json.Unmarshal(binaryMsg, &m)
	log.Info("msgParse 收到消息：", m)
	if err != nil {
		log.Error("Unmarshal failed", err.Error())
		return
	}

	switch m.ReqIdentifier {
	case constant.WSHeartBeat:
		ws.handleHeartBeat(conn, &m)
	case constant.WSGetNewestSeq:
		ws.getNewestSeqReq(conn, &m)
	case constant.WSPullMsgBySeqList:
		ws.pullMsgBySeqListReq(conn, &m)
	case constant.WSSendMsg:
		ws.sendMsgReq(conn, &m)
	case constant.WSPullConversationList:
		ws.pullConversationListReq(conn, &m)
	}
}

// handleHeartBeat 处理 ws 收到的心跳消息
func (ws *WsServer) handleHeartBeat(conn *UserConn, m *Req) {
	log.Info("ws call handleHeartBeat")
	resp := Resp{
		ReqIdentifier: constant.WSHeartBeat,
		Code:          constant.NoError,
		Info:          constant.SuccessInfo,
	}
	reply, _ := json.Marshal(resp)
	// TODO(qingw1230): 为每个用户维护定时器，心跳超时后断开 ws 连接
	ws.writeMsg(conn, websocket.TextMessage, reply)
}

func (ws *WsServer) getNewestSeqReq(conn *UserConn, m *Req) {
	log.Info("ws call getNewestSeqReq")
	nReply := pbMsg.GetNewestSeqResp{CommonResp: &sdkws.CommonResp{}}
	pbData := pbMsg.GetNewestSeqReq{
		UserId:   m.SendId,
		OpUserId: m.SendId,
	}

	rpcConn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.OfflineMessageName)
	client := pbMsg.NewMsgClient(rpcConn)
	reply, err := client.GetNewestSeq(context.Background(), &pbData)
	if err != nil {
		log.Error("ws call GetNewestSeq failed", err.Error())
		copier.Copy(nReply.CommonResp, reply.CommonResp)
		ws.getNewestSeqResp(conn, m, &nReply)
		return
	}

	ws.getNewestSeqResp(conn, m, reply)
}

func (ws *WsServer) getNewestSeqResp(conn *UserConn, m *Req, pb *pbMsg.GetNewestSeqResp) {
	b, _ := json.Marshal(pb)
	mReply := Resp{
		ReqIdentifier: m.ReqIdentifier,
		Code:          pb.CommonResp.Code,
		Info:          pb.CommonResp.Info,
		Data:          b,
	}
	ws.sendMsg(conn, mReply)
}

func (ws *WsServer) pullMsgBySeqListReq(conn *UserConn, m *Req) {
	log.Info("ws call pullMsgBySeqListReq")
	nReply := pbMsg.PullMessageBySeqListResp{CommonResp: &sdkws.CommonResp{}}
	isPass, code, info, pData := ws.argsValidate(m, constant.WSPullMsgBySeqList)
	if !isPass {
		nReply.CommonResp.Status = constant.Fail
		nReply.CommonResp.Code = code
		nReply.CommonResp.Info = info
		ws.pullMsgBySeqListResp(conn, m, &nReply)
		return
	}

	pbData := pbMsg.PullMessageBySeqListReq{
		UserId:   m.SendId,
		OpUserId: m.SendId,
		SeqList:  pData.(*pbMsg.PullMessageBySeqListReq).SeqList,
	}

	rpcConn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.OfflineMessageName)
	client := pbMsg.NewMsgClient(rpcConn)
	reply, err := client.PullMessageBySeqList(context.Background(), &pbData)
	if err != nil {
		log.Error("ws call PullMessageBySeqList failed", err.Error())
		copier.Copy(nReply.CommonResp, reply.CommonResp)
		ws.pullMsgBySeqListResp(conn, m, &nReply)
		return
	}

	ws.pullMsgBySeqListResp(conn, m, reply)
}

func (ws *WsServer) pullMsgBySeqListResp(conn *UserConn, m *Req, pb *pbMsg.PullMessageBySeqListResp) {
	b, _ := json.Marshal(pb)
	mReply := Resp{
		ReqIdentifier: m.ReqIdentifier,
		Code:          pb.CommonResp.Code,
		Info:          pb.CommonResp.Info,
		Data:          b,
	}
	ws.sendMsg(conn, mReply)
}

// sendMsgReq 处理 ws 收到的发送消息请求
func (ws *WsServer) sendMsgReq(conn *UserConn, m *Req) {
	log.Info("ws call sendMsgReq")
	nReply := pbMsg.SendMsgResp{CommonResp: &sdkws.CommonResp{}}
	isPass, code, info, pData := ws.argsValidate(m, constant.WSSendMsg)
	if !isPass {
		nReply.CommonResp.Status = constant.Fail
		nReply.CommonResp.Code = code
		nReply.CommonResp.Info = info
		ws.sendMsgResp(conn, m, &nReply)
		return
	}

	data := pData.(*sdkws.MsgData)
	pbData := pbMsg.SendMsgReq{
		Token:   m.Token,
		MsgData: data,
	}

	rpcConn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.OfflineMessageName)
	client := pbMsg.NewMsgClient(rpcConn)
	reply, err := client.SendMsg(context.Background(), &pbData)
	if err != nil {
		log.Error("ws call SendMsg failed", err.Error())
		copier.Copy(nReply.CommonResp, reply.CommonResp)
		ws.sendMsgResp(conn, m, &nReply)
		return
	}

	ws.sendMsgResp(conn, m, reply)
}

func (ws *WsServer) sendMsgResp(conn *UserConn, m *Req, pb *pbMsg.SendMsgResp) {
	mReplyData := sdkws.UserSendMsgResp{
		ServerMsgId: pb.GetServerMsgId(),
		SendTime:    pb.GetSendTime(),
	}
	b, _ := json.Marshal(&mReplyData)
	mReply := Resp{
		ReqIdentifier: m.ReqIdentifier,
		Code:          pb.CommonResp.Code,
		Info:          pb.CommonResp.Info,
		Data:          b,
	}
	ws.sendMsg(conn, mReply)
}

func (ws *WsServer) pullConversationListReq(conn *UserConn, m *Req) {
	log.Info("ws call pullConversationListReq")
	nReply := pbConversation.GetConversationListResp{CommonResp: &sdkws.CommonResp{}}
	isPass, code, info, pData := ws.argsValidate(m, constant.WSPullConversationList)
	if !isPass {
		nReply.CommonResp.Status = constant.Fail
		nReply.CommonResp.Code = code
		nReply.CommonResp.Info = info
		ws.pullConversationListResp(conn, m, &nReply)
		return
	}

	pbData := pData.(*pbConversation.GetConversationListReq)
	pbData.OpUserId = pbData.FromUserId

	rpcConn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.ConversationName)
	client := pbConversation.NewConversationClient(rpcConn)
	reply, err := client.GetConversationList(context.Background(), pbData)
	if err != nil {
		log.Error("ws call GetConversationList failed", err.Error())
		copier.Copy(nReply.CommonResp, reply.CommonResp)
		ws.pullConversationListResp(conn, m, &nReply)
		return
	}

	ws.pullConversationListResp(conn, m, reply)
}

func (ws *WsServer) pullConversationListResp(conn *UserConn, m *Req, pb *pbConversation.GetConversationListResp) {
	b, _ := json.Marshal(pb)
	mReply := Resp{
		ReqIdentifier: m.ReqIdentifier,
		Code:          pb.CommonResp.Code,
		Info:          pb.CommonResp.Info,
		Data:          b,
	}
	ws.sendMsg(conn, mReply)
}

// sendMsg 将 mReply 使用 JSON 序列化后通过 ws 发送
func (ws *WsServer) sendMsg(conn *UserConn, mReply interface{}) {
	b, err := json.Marshal(mReply)
	if err != nil {
		log.Error("JSON Marshal failed", err.Error())
		return
	}
	err = ws.writeMsg(conn, websocket.TextMessage, b)
	if err != nil {
		log.Error("ws writeMsg failed", err.Error())
	}
}
