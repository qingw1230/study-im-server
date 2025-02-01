package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

func InsertIntoConversation(toInsertRecord *db.Conversation) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	toInsertRecord.LastMessageTime = time.Now()
	return dbConn.Table(constant.DBTableConversation).Create(toInsertRecord).Error
}
