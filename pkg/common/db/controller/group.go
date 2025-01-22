package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

// InsertIntoGroup 向 Group 表插入群信息
func InsertIntoGroup(groupInfo *db.Group) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	groupInfo.CreateTime = time.Now()
	return dbConn.Table(constant.DBTableGroup).Create(groupInfo).Error
}

// GetGroupInfoByGroupId 根据 groupId 获取群聊信息
func GetGroupInfoByGroupId(groupId string) (*db.Group, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var groupInfo db.Group
	err = dbConn.Table(constant.DBTableGroup).Where("group_id = ?", groupId).First(&groupInfo).Error
	if err != nil {
		return nil, err
	}
	return &groupInfo, nil
}

func SetGroupInfo(groupInfo *db.Group) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	return dbConn.Table(constant.DBTableGroup).Where("group_id = ?", groupInfo.GroupId).Update(groupInfo).Error
}
