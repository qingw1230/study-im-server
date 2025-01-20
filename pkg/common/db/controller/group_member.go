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
