package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

// GetReceivedFriendRequestList 获取想要添加我的好友请求
func GetReceivedFriendRequestList(toUserId string, offset, limit int) ([]db.FriendRequest, int, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, 0, err
	}

	var friendRequests []db.FriendRequest

	// 查询总记录数
	var total int
	err = dbConn.Table(constant.DBTableFriendRequest).
		Where("to_user_id = ?", toUserId).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = dbConn.Table(constant.DBTableFriendRequest).
		Where("to_user_id = ?", toUserId).
		Offset(offset).
		Limit(limit).
		Find(&friendRequests).Error
	if err != nil {
		return nil, 0, err
	}

	return friendRequests, total, nil
}

// InsertFriendRequest 插入一条好友申请记录
func InsertFriendRequest(friendRequest *db.FriendRequest) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	return dbConn.Table(constant.DBTableFriendRequest).Create(friendRequest).Error
}

func UpdateFriendRequest(friendRequest *db.FriendRequest) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	friendRequest.CreateTime = time.Now().UnixMilli()
	return dbConn.Table(constant.DBTableFriendRequest).Where("from_user_id = ? AND to_user_id = ?",
		friendRequest.FromUserId, friendRequest.ToUserId).Update(friendRequest).Error
}

// GetFriendRequestByBothUserId 根据添加方和被添加方的 UserId 获取申请记录
// 若找到了返回相应的记录，否则返回错误
func GetFriendRequestByBothUserId(fromUserId, toUserId string) (*db.FriendRequest, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var friendRequest db.FriendRequest
	err = dbConn.Table(constant.DBTableFriendRequest).Where("from_user_id = ? AND to_user_id = ?", fromUserId, toUserId).First(&friendRequest).Error
	if err != nil {
		return nil, err
	}
	return &friendRequest, nil
}
