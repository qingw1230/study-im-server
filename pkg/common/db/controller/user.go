package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

func CreateUser(user db.User) error {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return err
	}
	return dbConn.Table(constant.DBTableUser).Create(&user).Error
}

func IsUserExist(email string) bool {
	_, err := FindUserByEmail(email)
	return err != gorm.ErrRecordNotFound
}

// FindUserById 用 UserId 查找用户，找到了 err 为 nil，找不到不为 nil
func FindUserById(id string) (*db.User, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}

	var user db.User
	err = dbConn.Table(constant.DBTableUser).Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindUserByEmail(email string) (*db.User, error) {
	dbConn, err := db.DB.MySQLDB.DefaultGormDB()
	if err != nil {
		return nil, err
	}

	var user db.User
	err = dbConn.Table(constant.DBTableUser).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
