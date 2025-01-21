package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

// InsertIntoBlacklist 向黑名单中插入数据
func InsertIntoBlacklist(black *db.Black) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	black.CreateTime = time.Now()
	return dbConn.Table(constant.DBTableBlack).Create(black).Error
}
