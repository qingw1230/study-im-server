package msg_gateway

import (
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
)

func (ws *WsServer) msgParse(conn *UserConn, binaryMsg []byte) {
	m := Req{}
	err := json.Unmarshal(binaryMsg, &m)
	log.Info(m)
	if err != nil {

	}

	switch m.ReqIdentifier {
	case constant.WSHeartBeat:
		ws.handleHeartBeat(conn, &m)
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
	ws.writeMsg(conn, websocket.TextMessage, reply)
}

// func (ws *WsServer) sendMsgReq(conn *UserConn, m *Req) {
// 	resp := pbMsg.SendMsgResp{CommonResp: &pbPublic.CommonResp{}}
// 	isPass, code, info, pbMsgData := ws.argsValidate(m, constant.WSSendMsg)
// 	if !isPass {
// 		resp.CommonResp.Status = constant.Fail
// 		resp.CommonResp.Code = code
// 		resp.CommonResp.Info = info
// 		// TODO(qingw1230): 发送错误信息
// 	}

// 	data := pbMsgData.(*pbPublic.MsgData)
// 	pbSendMsgReq := pbMsg.SendMsgReq{
// 		Token:   m.Token,
// 		MsgData: data,
// 	}

// 	// TODO(qingw1230): 使用服务发现建立连接
// 	rpcConn, err := grpc.NewClient("127.0.0.1:10300", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Error("grpc.NewClient failed", err.Error())
// 		resp.CommonResp.Status = constant.Fail
// 		resp.CommonResp.Code = constant.WSCommonError
// 		resp.CommonResp.Info = constant.WSCommonErrorInfo
// 		// TODO(qingw1230): 发送错误消息
// 		return
// 	}
// 	client := pbMsg.NewMsgClient(rpcConn)
// 	_, err = client.SendMsg(context.Background(), &pbSendMsgReq)
// 	if err != nil {
// 		resp.CommonResp.Status = constant.Fail
// 		resp.CommonResp.Code = constant.WSCommonError
// 		resp.CommonResp.Info = constant.WSCommonErrorInfo
// 		// TODO(qingw1230): 发送错误消息
// 		return
// 	}
// }
