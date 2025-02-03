package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

func GetConversationList(ownerUserId string) ([]db.Conversation, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	conversationList := make([]db.Conversation, 0)
	err = dbConn.Table(constant.DBTableConversation).Where("owner_user_id = ?", ownerUserId).Find(&conversationList).Error
	if err != nil {
		return nil, err
	}
	return conversationList, nil
}

func InsertIntoConversation(toInsertRecord *db.Conversation) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	toInsertRecord.LastMessageTime = time.Now().Unix()
	return dbConn.Table(constant.DBTableConversation).Create(toInsertRecord).Error
}
