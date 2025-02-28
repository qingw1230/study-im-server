package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/utils"
)

// CreateGroup 向 Group 表插入群信息
func CreateGroup(groupInfo *db.Group) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	groupInfo.CreateTime = time.Now()
	groupInfo.NotificationUpdateTime = utils.UnixSecondToTime(0)
	return dbConn.Table(constant.DBTableGroup).Create(groupInfo).Error
}

// DeleteGroup 删除指定群的所有信息
func DeleteGroup(groupId string) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	if err := dbConn.Table(constant.DBTableGroup).Where("group_id = ?", groupId).Delete(&db.Group{}).Error; err != nil {
		return err
	}
	if err := dbConn.Table(constant.DBTableGroupMember).Where("group_id = ?", groupId).Delete(&db.GroupMember{}).Error; err != nil {
		return err
	}
	// TODO(qingw1230): group_requests 表要怎么处理
	return nil
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
