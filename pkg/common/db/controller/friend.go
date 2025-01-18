package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

// InsertToFriend 向 Friend 表中插入单向好友关系
func InsertToFriend(toInsertFollow *db.Friend) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	toInsertFollow.CreateTime = time.Now()
	return dbConn.Table(constant.DBTableFriend).Create(toInsertFollow).Error
}

// GetFriendListByUserID 根据 UserID 获取好友列表
func GetFriendListByUserID(ownerUserID string) ([]db.Friend, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var friends []db.Friend
	err = dbConn.Table(constant.DBTableFriend).Where("owner_user_id = ?", ownerUserID).Find(&friends).Error
	if err != nil {
		return nil, err
	}
	return friends, nil
}

// GetFriendRelationFromFriend 从 Friend 表中获取好友关系
func GetFriendRelationFromFriend(ownerUserID, friendUserID string) (*db.Friend, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var friend db.Friend
	err = dbConn.Table(constant.DBTableFriend).Where("owner_user_id = ? AND friend_user_id = ?", ownerUserID, friendUserID).First(&friend).Error
	if err != nil {
		return nil, err
	}
	return &friend, nil
}
