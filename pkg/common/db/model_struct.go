package db

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
)

// User 用户信息表
type User struct {
	// 新版本添加主键使用: primaryKey
	// 新版本添加唯一索引使用: uniqueIndex
	UserId            string `gorm:"column:user_id;type:varchar(12);not null;primary_key"`
	Email             string `gorm:"column:email;type:varchar(50);not null;unique_index:idx_key_email"`
	Password          string `gorm:"column:password;type:varchar(32);not null"`
	Salt              string `gorm:"column:salt;type:varchar(10);not null"`
	NickName          string `gorm:"column:nick_name;type:varchar(20)"`
	FaceUrl           string `gorm:"column:face_url;type:varchar(255)"`
	PersonalSignature string `gorm:"column:personal_signature;type:varchar(50)"`
	Sex               int    `gorm:"column:sex;type:tinyint(1)"`
	JoinType          int    `gorm:"column:join_type;type:tinyint(1)"`
	AreaName          string `gorm:"column:area_name;type:varchar(50)"`
	AreaCode          string `gorm:"column:area_code;type:varchar(50)"`
	CreateTime        int    `gorm:"column:create_time;type:int(11)"`
}

func (User) TableName() string {
	return constant.DBTableUser
}

// UserLuckyNumber 靓号表
type UserLuckyNumber struct {
	Id     int    `gorm:"column:id;type:int(11) auto_increment;not null;primary_key"`
	Email  string `gorm:"column:email;type:varchar(50);not null;unique_index:idx_key_email"`
	UserId string `gorm:"column:user_id;type:varchar(12);not null;unique_index:idx_key_user_id"`
	Status int    `gorm:"column:status;type:tinyint(1)"`
}

func (UserLuckyNumber) TableName() string {
	return constant.DBTableUserLuckyNumber
}

type Friend struct {
	Id           int       `gorm:"column:id;type:int(11) auto_increment;not null;primary_key"`
	OwnerUserId  string    `gorm:"column:owner_user_id;type:varchar(12);not null;unique_index:idx_key"`
	FriendUserId string    `gorm:"column:friend_user_id;type:varchar(12);not null;unique_index:idx_key"`
	Remark       string    `gorm:"column:remark;type:varchar(255)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	AddSource    int32     `gorm:"column:add_source;type:int(11)"`
	OpUserId     string    `gorm:"column:op_user_id;type:varchar(12)"`
	Ex           string    `gorm:"column:ex;type:varchar(1024)"`
}

func (Friend) TableName() string {
	return constant.DBTableFriend
}

// FriendRequest 好友申请表
type FriendRequest struct {
	RequestId    int       `gorm:"column:request_id;type:int(11) auto_increment;not null;primary_key"`
	FromUserId   string    `gorm:"column:from_user_id;type:varchar(12);not null;unique_index:idx_key"`
	ToUserId     string    `gorm:"column:to_user_id;type:varchar(12);not null;unique_index:idx_key"`
	HandleResult int8      `gorm:"column:handle_result;type:tinyint(1)"`
	ReqMsg       string    `gorm:"column:req_msg;type:varchar(255)"`
	HandleUserId string    `gorm:"column:handle_user_id;type:varchar(12)"`
	HandleMsg    string    `gorm:"column:handle_msg;type:varchar(255)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	HandleTime   time.Time `gorm:"column:handle_time;typd:datetime"`
	Ex           string    `gorm:"column:ex;type:varchar(1024)"`
}

func (FriendRequest) TableName() string {
	return constant.DBTableFriendRequest
}
