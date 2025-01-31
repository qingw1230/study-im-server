package msg_gateway

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	"github.com/qingw1230/study-im-server/pkg/utils"
)

type UserConn struct {
	*websocket.Conn
	mu *sync.Mutex
}

type WsServer struct {
	wsAddr       string
	wsMaxConnNum int
	wsUpGrader   *websocket.Upgrader
	wsConnToUser map[*UserConn]string
	wsUserToConn map[string]*UserConn
}

// onInit 初始化 WsServer 结构体
func (ws *WsServer) onInit(port int) {
	ws.wsAddr = ":" + utils.IntToString(port)
	ws.wsMaxConnNum = config.Config.LongConnSvr.WebsocketMaxConnNum
	ws.wsConnToUser = make(map[*UserConn]string)
	ws.wsUserToConn = make(map[string]*UserConn)
	ws.wsUpGrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(config.Config.LongConnSvr.WebsocketTimeout),
		ReadBufferSize:   config.Config.LongConnSvr.WebsocketMaxMsgLen,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (ws *WsServer) run() {
	http.HandleFunc("/", ws.wsHandler)
	err := http.ListenAndServe(ws.wsAddr, nil)
	if err != nil {
		log.Error("ws ListenAndServer failed", err.Error())
	}
}

// wsHandler 将 HTTP 连接升级成 ws 连接
func (ws *WsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	if ws.headerCheck(w, r) {
		conn, err := ws.wsUpGrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error("Upgrade failed", err.Error())
			return
		}
		log.Info("Upgrade success")
		sendId := r.Header.Get(constant.STR_SEND_ID)
		newConn := &UserConn{Conn: conn, mu: &sync.Mutex{}}
		ws.addUserConn(sendId, newConn)
		go ws.readMsg(newConn)
	}
}

// readMsg 不断读取 ws 中的消息并处理
func (ws *WsServer) readMsg(conn *UserConn) {
	for {
		messageType, msg, err := conn.ReadMessage()
		if messageType == websocket.PingMessage {
			log.Info("receive a PingMessage")
		}
		if err != nil {
			log.Error("ws ReadMessage failed", "userId:", ws.wsConnToUser[conn], err.Error())
			ws.delUserConn(conn)
			return
		}
		log.Info("ReadMessage success, type", messageType)
		ws.msgParse(conn, msg)
	}
}

func (ws *WsServer) writeMsg(conn *UserConn, a int, msg []byte) error {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	return conn.WriteMessage(a, msg)
}

// addUserConn 添加 userId 与 ws 连接的映射
func (ws *WsServer) addUserConn(userId string, conn *UserConn) {
	log.Info("call addUserConn")
	rw.Lock()
	defer rw.Unlock()
	ws.wsUserToConn[userId] = conn
	ws.wsConnToUser[conn] = userId
	log.Info("addUserConn return")
}

// delUserConn 删除 userId 与 ws 连接的映射，并关闭 ws 连接
func (ws *WsServer) delUserConn(conn *UserConn) {
	log.Info("call delUserConn")
	rw.Lock()
	defer rw.Unlock()
	if userId, ok := ws.wsConnToUser[conn]; ok {
		delete(ws.wsUserToConn, userId)
		delete(ws.wsConnToUser, conn)
	}
	err := conn.Close()
	if err != nil {
		log.Error("close ws failed", err.Error())
	}
	log.Info("delUserConn return")
}

// headerCheck 检查 URL 参数部分的 token 和 sendId
func (ws *WsServer) headerCheck(w http.ResponseWriter, r *http.Request) bool {
	log.Info("call headerCheck")
	query := r.URL.Query()
	sendId := query[constant.STR_SEND_ID][0]
	token := query[constant.STR_TOKEN][0]
	if ok, err := token_verify.VerifyToken(token, sendId); !ok {
		log.Error("VerifyToken failed", err.Error(), sendId, token)
		e := err.(*constant.ErrInfo)
		http.Error(w, e.ErrMsg, e.ErrCode)
		return false
	}
	log.Info("headerCheck return")
	return true
}
