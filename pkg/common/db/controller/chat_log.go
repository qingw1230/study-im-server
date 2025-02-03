package controller

import (
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
)

func InsertMessageToChatLog(msg *pbMsg.MsgDataToMq) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}

	chatLog := &db.ChatLog{}
	copier.Copy(chatLog, msg.MsgData)
	switch msg.MsgData.SessionType {
	case constant.SingleChatType:
		chatLog.RecvId = msg.MsgData.RecvId
	}
	chatLog.Content = msg.MsgData.Content
	chatLog.CreateTime = msg.MsgData.CreateTime
	chatLog.SendTime = msg.MsgData.SendTime
	return dbConn.Table(constant.DBTableChatLog).Create(chatLog).Error
}
