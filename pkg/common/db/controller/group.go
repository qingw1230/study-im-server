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
