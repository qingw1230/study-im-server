package db

import "time"

// UserInfo 用户信息表
type UserInfo struct {
	// 新版本添加主键使用: primaryKey
	// 新版本添加唯一索引使用: uniqueIndex
	UserID            string    `gorm:"column:user_id;type:varchar(12);not null;primary_key"`
	Email             string    `gorm:"column:email;type:varchar(50);not null;unique_index:idx_key_email"`
	Password          string    `gorm:"column:password;type:varchar(32);not null"`
	NickName          string    `gorm:"column:nick_name;type:varchar(20)"`
	PersonalSignature string    `gorm:"column:personal_signature;type:varchar(50)"`
	Sex               int       `gorm:"column:sex;type:tinyint(1)"`
	JoinType          int       `gorm:"column:join_type;type:tinyint(1)"`
	Status            int       `gorm:"column:status;type:tinyint(1)"`
	AreaName          string    `gorm:"column:area_name;type:varchar(50)"`
	AreaCode          string    `gorm:"column:area_code;type:varchar(50)"`
	CreateTime        time.Time `gorm:"column:create_time;type:datetime"`
	LastLoginTime     time.Time `gorm:"column:last_login_time;type:datetime"`
	LastOffTime       int       `gorm:"column:last_off_time;type:bigint(13)"`
}

func (UserInfo) TableName() string {
	return "user_info"
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
	GroupID      string    `gorm:"column:group_id;type:varchar(12);not null;primary_key"`
	GroupName    string    `gorm:"column:group_name;type:varchar(20)"`
	GroupOwnerID string    `gorm:"column:group_owner_id;type:varchar(12)"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
	GroupNotice  string    `gorm:"column:group_notice;type:varchar(500)"`
	JoinType     int       `gorm:"column:join_type;type:tinyint(1)"`
	Status       int       `gorm:"column:status;type:tinyint(1)"`
}

func (GroupInfo) TableName() string {
	return "group_info"
}

// UserContact 用户联系人表
type UserContact struct {
	UserID         string    `gorm:"column:user_id;type:varchar(12);not null;primary_key"`
	ContactID      string    `gorm:"column:contact_id;type:varchar(12);not null;primary_key;index:idx_contact_id"`
	ContactType    int       `gorm:"column:cotact_type;type:tinyint(1)"`
	CreateTime     time.Time `gorm:"column:create_time;type:datetime"`
	Status         int       `gorm:"column:status;type:tinyint(1)"`
	LastUpdateTime int       `gorm:"column:last_update_time;type:timestamp;autoUpdateTime"`
}

func (UserContact) TableName() string {
	return "user_contact"
}

// UserContactApply 联系人申请表
type UserContactApply struct {
	ApplyID       int    `gorm:"column:apply_id;type:int(11) auto_increment;not null;primary_key"`
	ApplyUserID   string `gorm:"column:apply_user_id;type:varchar(12);not null;unique_index:idx_key"`
	ReceiveUserID string `gorm:"column:receive_user_id;type:varchar(12);not null;unique_index:idx_key"`
	ContactType   int    `gorm:"column:contact_type;type:tinyint(1);not null"`
	ContactID     string `gorm:"column:contact_id;type:varchar(12);unique_index:idx_key"`
	LastApplyTime int    `gorm:"column:last_apply_time;type:bigint(20);index:idx_last_apply_time"`
	Status        int    `gorm:"column:status;type:tinyint(1)"`
	ApplyInfo     string `gorm:"column:apply_info;type:varchar(100)"`
}

func (UserContactApply) TableName() string {
	return "user_contact_apply"
}
