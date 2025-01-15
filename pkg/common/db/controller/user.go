package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

func CreateUser(user db.UserInfo) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	return dbConn.Table(constant.DBTableUserInfo).Create(&user).Error
}

func IsUserExist(email string) bool {
	_, err := FindUserByEmail(email)
	return err != gorm.ErrRecordNotFound
}

func FindUserByID(id string) (*db.UserInfo, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}

	var user db.UserInfo
	err = dbConn.Table(constant.DBTableUserInfo).Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(email string) (*db.UserInfo, error) {
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
