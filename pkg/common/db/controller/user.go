package controller

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/utils"
)

func CreateUser(email, nickName, pwd string) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}

	now := time.Now()
	user := db.UserInfo{
		UserID:        utils.GenerateRandomID(constant.LENGTH_11),
		Password:      utils.MakePassword(pwd),
		Email:         email,
		NickName:      nickName,
		JoinType:      constant.UserInfoJoinTypeAPPLY,
		Status:        constant.UserInfoStatusENABLE,
		CreateTime:    now,
		LastLoginTime: now,
		LastOffTime:   int(now.UnixNano()),
	}
	return dbConn.Table(constant.DBTableUserInfo).Create(&user).Error
}

func GetUserByEmail(email string) (*db.UserInfo, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}

	var user db.UserInfo
	err = dbConn.Table(constant.DBTableUserInfo).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
