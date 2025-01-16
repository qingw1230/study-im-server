package controller

import (
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

// InsertFriendApplication 插入一条好友申请记录
func InsertFriendApplication(friendRequest *db.FriendRequest) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	return dbConn.Table(constant.DBTableFrindRequest).Create(friendRequest).Error
}
