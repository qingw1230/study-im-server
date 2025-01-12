package db

import "time"

type UserInfo struct {
	// 新版本添加主键使用: primaryKey
	// 新版本添加唯一索引使用: uniqueIndex
	UserID            string    `gorm:"column:user_id;type:varchar(12);not null;primary_key"`
	Password          string    `gorm:"column:password;type:varchar(32);not null"`
	Email             string    `gorm:"column:email;type:varchar(50);not null;unique_index:idx_key_email"`
	NickName          string    `gorm:"column:nick_name;type:varchar(20)"`
	PersonalSignature string    `gorm:"column:personal_signature;type:varchar(50)"`
	JoinType          int       `gorm:"column:join_type;type:tinyint(1)"`
	Sex               int       `gorm:"column:sex;type:tinyint(1)"`
	Status            int       `gorm:"column:status;type:tinyint(1)"`
	AreaName          string    `gorm:"column:area_name;type:varchar(50)"`
	AreaCode          string    `gorm:"columb:area_code;type:varchar(50)"`
	CreateTime        time.Time `gorm:"column:create_time;type:datetime"`
	LastLoginTime     time.Time `gorm:"column:last_login_time;type:datetime"`
	LastOffTime       int       `gorm:"column:last_off_time;type:bigint(13)"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type UserInfoBeauty struct {
	ID     int    `gorm:"column:id;type:int(11);not null;primary_key"`
	Email  string `gorm:"column:email;type:varchar(50);not null;unique_index:idx_key_email"`
	UserID string `gorm:"column:user_id;type:varchar(12);not null;unique_index:idx_key_user_id"`
	Status int    `gorm:"column:status;type:tinyint(1)"`
}

func (UserInfoBeauty) TableName() string {
	return "user_info_beauty"
}
