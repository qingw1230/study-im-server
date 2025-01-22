package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

// InsertIntoGroupMember 插入群成员
func InsertIntoGroupMember(toInsertInfo *db.GroupMember) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	toInsertInfo.JoinTime = time.Now()
	if toInsertInfo.RoleLevel == 0 {
		toInsertInfo.RoleLevel = constant.GroupOrdinaryUser
	}
	return dbConn.Table(constant.DBTableGroupMember).Create(&toInsertInfo).Error
}

// GetGroupMemberListByUserId 根据 userId 获取所有群聊的 GroupMember
func GetGroupMemberListByUserId(userId string) ([]db.GroupMember, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var groupMemberList []db.GroupMember
	err = dbConn.Table(constant.DBTableGroupMember).Where("user_id = ?", userId).Find(&groupMemberList).Error
	if err != nil {
		return nil, err
	}
	return groupMemberList, nil
}

// GetGroupMemberInfoByGroupIdAndUserId 根据 groupId 和 userId 获取群成员信息
func GetGroupMemberInfoByGroupIdAndUserId(groupId, userId string) (*db.GroupMember, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var groupMember db.GroupMember
	err = dbConn.Table(constant.DBTableGroupMember).Where("group_id = ? AND user_id = ?", groupId, userId).First(&groupMember).Error
	if err != nil {
		return nil, err
	}
	return &groupMember, nil
}

// DeleteGroupMemberByGroupIdAndUserId 根据 groupId 和 userId 删除群成员信息
func DeleteGroupMemberByGroupIdAndUserId(groupId, userId string) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	return dbConn.Table(constant.DBTableGroupMember).Where("group_id = ? AND user_id = ?", groupId, userId).Delete(&db.GroupMember{}).Error
}

// GetGroupMemberNumByGroupId 根据 groupId 获取群成员数量
func GetGroupMemberNumByGroupId(groupId string) (uint32, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return 0, err
	}
	var num uint32
	err = dbConn.Table(constant.DBTableGroupMember).Where("group_id = ?", groupId).Count(&num).Error
	if err != nil {
		return 0, err
	}
	return num, err
}

func GetGroupOwnerManagerInfoByGroupId(groupId string) ([]db.GroupMember, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}
	var groupMemberList []db.GroupMember
	err = dbConn.Table(constant.DBTableGroupMember).Where("group_id = ? AND role_level > ?", groupId, constant.GroupOrdinaryUser).Find(&groupMemberList).Error
	if err != nil {
		return nil, err
	}
	return groupMemberList, nil
}

func GetGroupOwnerInfoByGroupId(groupId string) (*db.GroupMember, error) {
	omList, err := GetGroupOwnerManagerInfoByGroupId(groupId)
	if err != nil {
		return nil, err
	}
	for _, v := range omList {
		if v.RoleLevel == constant.GroupOwner {
			return &v, nil
		}
	}
	return nil, nil
}

// GetOwnedGroupIdListByUserId 根据 userId 获取创建的所有群聊的 GroupId
func GetOwnedGroupIdListByUserId(userId string) ([]string, error) {
	groupMemberList, err := GetGroupMemberListByUserId(userId)
	if err != nil {
		return nil, err
	}
	var groupIdList []string
	for _, v := range groupMemberList {
		if v.RoleLevel == constant.GroupOwner {
			groupIdList = append(groupIdList, v.GroupId)
		}
	}
	return groupIdList, nil
}

// GetJoinedGroupIdListByUserId 根据 userId 获取添加的所有群聊的 GroupId
func GetJoinedGroupIdListByUserId(userId string) ([]string, error) {
	groupMemberList, err := GetGroupMemberListByUserId(userId)
	if err != nil {
		return nil, err
	}
	var groupIdList []string
	for _, v := range groupMemberList {
		if v.RoleLevel != constant.GroupOwner {
			groupIdList = append(groupIdList, v.GroupId)
		}
	}
	return groupIdList, nil
}
