package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/utils"
)

// InsertFriendApplication 插入一条好友申请记录
func InsertFriendApplication(friendRequest *db.FriendRequest) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	friendRequest.CreateTime = time.Now()
	friendRequest.HandleTime = utils.UnixSecondToTime(0)
	return dbConn.Table(constant.DBTableFriendRequest).Create(friendRequest).Error
}

func UpdateFriendApplication(friendRequest *db.FriendRequest) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	friendRequest.CreateTime = time.Now()
	return dbConn.Table(constant.DBTableFriendRequest).Where("from_user_id = ? AND to_user_id = ?", friendRequest.FromUserID, friendRequest.ToUserID).Update(friendRequest).Error
}

// GetFriendApplicationByBothUserID 根据添加方和被添加方的 UserID 获取申请记录
// 若找到了返回相应的记录，否则返回错误
func GetFriendApplicationByBothUserID(fromUserID, toUserID string) (*db.FriendRequest, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var friendRequest db.FriendRequest
	err = dbConn.Table(constant.DBTableFriendRequest).Where("from_user_id = ? AND to_user_id = ?", fromUserID, toUserID).First(&friendRequest).Error
	if err != nil {
		return nil, err
	}
	return &friendRequest, nil
}
