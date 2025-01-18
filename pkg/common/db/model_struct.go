package db

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
)

// UserInfo 用户信息表
type UserInfo struct {
	// 新版本添加主键使用: primaryKey
	// 新版本添加唯一索引使用: uniqueIndex
	UserID            string `gorm:"column:user_id;type:varchar(12);not null;primary_key"`
	Email             string `gorm:"column:email;type:varchar(50);not null;unique_index:idx_key_email"`
	Password          string `gorm:"column:password;type:varchar(32);not null"`
	Salt              string `gorm:"column:salt;type:varchar(10);not null"`
	NickName          string `gorm:"column:nick_name;type:varchar(20)"`
	PersonalSignature string `gorm:"column:personal_signature;type:varchar(50)"`
	Sex               int    `gorm:"column:sex;type:tinyint(1)"`
	JoinType          int    `gorm:"column:join_type;type:tinyint(1)"`
	Status            int    `gorm:"column:status;type:tinyint(1)"`
	AreaName          string `gorm:"column:area_name;type:varchar(50)"`
	AreaCode          string `gorm:"column:area_code;type:varchar(50)"`
	CreateTime        int    `gorm:"column:create_time;type:int(11)"`
	LastLoginTime     int    `gorm:"column:last_login_time;type:int(11)"`
	LastOffTime       int    `gorm:"column:last_off_time;type:int(11)"`
}

func (UserInfo) TableName() string {
	return constant.DBTableUserInfo
}

// UserInfoBeauty 靓号表
type UserInfoBeauty struct {
	ID     int    `gorm:"column:id;type:int(11) auto_increment;not null;primary_key"`
	Email  string `gorm:"column:email;type:varchar(50);not null;unique_index:idx_key_email"`
	UserID string `gorm:"column:user_id;type:varchar(12);not null;unique_index:idx_key_user_id"`
	Status int    `gorm:"column:status;type:tinyint(1)"`
}

func (UserInfoBeauty) TableName() string {
	return "user_info_beauty"
}

// GroupInfo 群信息表
type GroupInfo struct {
	GroupID      string `gorm:"column:group_id;type:varchar(12);not null;primary_key"`
	GroupName    string `gorm:"column:group_name;type:varchar(20)"`
	GroupOwnerID string `gorm:"column:group_owner_id;type:varchar(12)"`
	CreateTime   int    `gorm:"column:create_time;type:int(11)"`
	GroupNotice  string `gorm:"column:group_notice;type:varchar(500)"`
	JoinType     int    `gorm:"column:join_type;type:tinyint(1)"`
	Status       int    `gorm:"column:status;type:tinyint(1)"`
}

func (GroupInfo) TableName() string {
	return "group_info"
}

type Friend struct {
	ID           int       `gorm:"column:id;type:int(11) auto_increment;not null;primary_key"`
	OwnerUserID  string    `gorm:"column:owner_user_id;type:varchar(12);not null;unique_index:idx_key"`
	FriendUserID string    `gorm:"column:friend_user_id;type:varchar(12);not null;unique_index:idx_key"`
	Remark       string    `gorm:"column:remark;type:varchar(255)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	AddSource    int32     `gorm:"column:add_source;type:int(11)"`
	OpUserID     string    `gorm:"column:op_user_id;type:varchar(12)"`
	Ex           string    `gorm:"column:ex;type:varchar(1024)"`
}

func (Friend) TableName() string {
	return constant.DBTableFriend
}

// FriendRequest 好友申请表
type FriendRequest struct {
	RequestID    int       `gorm:"column:request_id;type:int(11) auto_increment;not null;primary_key"`
	FromUserID   string    `gorm:"column:from_user_id;type:varchar(12);not null;unique_index:idx_key"`
	ToUserID     string    `gorm:"column:to_user_id;type:varchar(12);not null;unique_index:idx_key"`
	HandleResult int8      `gorm:"column:handle_result;type:tinyint(1)"`
	ReqMsg       string    `gorm:"column:req_msg;type:varchar(255)"`
	HandleUserID string    `gorm:"column:handle_user_id;type:varchar(12)"`
	HandleMsg    string    `gorm:"column:handle_msg;type:varchar(255)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	HandleTime   time.Time `gorm:"column:handle_time;typd:datetime"`
	Ex           string    `gorm:"column:ex;type:varchar(1024)"`
}

func (FriendRequest) TableName() string {
	return constant.DBTableFriendRequest
}
