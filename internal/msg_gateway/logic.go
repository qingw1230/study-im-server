package msg_gateway

import (
	"context"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	case constant.WSSendMsg:
		ws.sendMsgReq(conn, &m)
	}
}

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

func (ws *WsServer) sendMsgReq(conn *UserConn, m *Req) {
	log.Info("ws call sendMsgReq")
	nReply := pbMsg.SendMsgResp{CommonResp: &pbPublic.CommonResp{}}
	isPass, code, info, pData := ws.argsValidate(m, constant.WSSendMsg)
	if !isPass {
		nReply.CommonResp.Status = constant.Fail
		nReply.CommonResp.Code = code
		nReply.CommonResp.Info = info
		ws.sendMsgResp(conn, m, &nReply)
		return
	}

	data := pData.(*pbPublic.MsgData)
	pbData := pbMsg.SendMsgReq{
		Token:   m.Token,
		MsgData: data,
	}

	// TODO(qingw1230): 使用服务发现建立连接
	rpcConn, err := grpc.NewClient("127.0.0.1:10300", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed", err.Error())
		return
	}
	client := pbMsg.NewMsgClient(rpcConn)
	reply, err := client.SendMsg(context.Background(), &pbData)
	if err != nil {
		log.Error("SendMsg failed", err.Error())
		copier.Copy(nReply.CommonResp, reply.CommonResp)
		ws.sendMsgResp(conn, m, &nReply)
	}
	ws.sendMsgResp(conn, m, reply)
}

func (ws *WsServer) sendMsgResp(conn *UserConn, m *Req, pb *pbMsg.SendMsgResp) {
	mReplyData := pbPublic.UserSendMsgResp{
		ServerMsgId: pb.GetServerMsgId(),
		ClientMsgId: pb.GetClientMsgId(),
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

// sendMsg 将 mReply 使用 JSON 序列化后通过 ws 发送
func (ws *WsServer) sendMsg(conn *UserConn, mReply interface{}) {
	b, err := json.Marshal(mReply)
	if err != nil {
		log.Error("JSON Marshal failed", err.Error())
		return
	}
	err = ws.writeMsg(conn, websocket.BinaryMessage, b)
	if err != nil {
		log.Error("ws writeMsg failed", err.Error())
	}
}
