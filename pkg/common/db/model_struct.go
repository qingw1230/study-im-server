package db

import (
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
)

// User 用户信息表
type User struct {
	// 新版本添加主键使用: primaryKey
	// 新版本添加唯一索引使用: uniqueIndex
	UserId            string    `gorm:"column:user_id;type:varchar(12);not null;primary_key"`
	Email             string    `gorm:"column:email;type:varchar(50);not null;unique_index:idx_key_email"`
	Password          string    `gorm:"column:password;type:varchar(32);not null"`
	Salt              string    `gorm:"column:salt;type:varchar(10);not null"`
	NickName          string    `gorm:"column:nick_name;type:varchar(20)"`
	FaceUrl           string    `gorm:"column:face_url;type:varchar(255)"`
	PersonalSignature string    `gorm:"column:personal_signature;type:varchar(50)"`
	Sex               int       `gorm:"column:sex;type:tinyint(1)"`
	JoinType          int       `gorm:"column:join_type;type:tinyint(1)"`
	AreaName          string    `gorm:"column:area_name;type:varchar(50)"`
	AreaCode          string    `gorm:"column:area_code;type:varchar(50)"`
	CreateTime        time.Time `gorm:"column:create_time;type:datetime"`
	Ex                string    `gorm:"column:ex;type:varchar(1024)"`
}

func (User) TableName() string {
	return constant.DBTableUser
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
	Id           int    `gorm:"column:id;type:int(11) auto_increment;not null;primary_key"`
	FromUserId   string `gorm:"column:from_user_id;type:varchar(12);not null;unique_index:idx_unique_request;index:idx_from"`
	ToUserId     string `gorm:"column:to_user_id;type:varchar(12);not null;unique_index:idx_unique_request;index:idx_to"`
	ReqMsg       string `gorm:"column:req_msg;type:varchar(255)"`
	HandleResult int32  `gorm:"column:handle_result;type:tinyint(1)"`
	CreateTime   int64  `gorm:"column:create_time;type:bigint"`
	HandleTime   int64  `gorm:"column:handle_time;type:bigint"`
}

func (FriendRequest) TableName() string {
	return constant.DBTableFriendRequest
}

type Group struct {
	GroupId                string    `gorm:"column:group_id;type:varchar(12);not null;primary_key"`
	GroupName              string    `gorm:"column:group_name;type:varchar(32)"`
	Notification           string    `gorm:"column:notification;type:varchar(255)"`
	Introduction           string    `gorm:"column:introduction;type:varchar(255)"`
	FaceUrl                string    `gorm:"column:face_url;type:varchar(255)"`
	CreateTime             time.Time `gorm:"column:create_time;type:datetime"`
	Status                 int32     `gorm:"column:status;type:tinyint"`
	CreateUserId           string    `gorm:"column:create_user_id;type:varchar(12)"`
	GroupType              int32     `gorm:"column:group_type;type:tinyint"`
	NeedVerification       int32     `gorm:"column:need_verification;type:tinyint"`
	NotificationUpdateTime time.Time `gorm:"column:notification_update_time;type:datetime"`
	NotificationUserId     string    `gorm:"column:notification_user_id;type:varchar(12)"`
}

func (Group) TableName() string {
	return constant.DBTableGroup
}

type GroupMember struct {
	GroupId        string    `gorm:"column:group_id;type:varchar(12);not null;primary_key"`
	UserId         string    `gorm:"column:user_id;type:varchar(12);not null;primary_key"`
	NickName       string    `gorm:"column:nick_name;type:varchar(20)"`
	FaceUrl        string    `gorm:"column:face_url;type:varchar(255)"`
	RoleLevel      int32     `gorm:"column:role_level;type:tinyint"`
	JoinTime       time.Time `gorm:"column:join_time;type:datetime"`
	JoinSource     int32     `gorm:"column:join_source;type:tinyint"`
	InviterUserId  string    `gorm:"column:inviter_user_id;type:varchar(12)"`
	OperatorUserId string    `gorm:"column:operator_user_id;type:varchar(12)"`
	// MuteEndTime    time.Time
}

func (GroupMember) TableName() string {
	return constant.DBTableGroupMember
}

// type GroupRequest struct {
// 	RequestId    int       `gorm:"column:request_id;type:int(11) auto_increment;not null;primary_key"`
// 	UserId       string    `gorm:"column:user_id;type:varchar(12);not null;unique_index:idx_key"`
// 	GroupId      string    `gorm:"column:group_id;type:varchar(12);not null;unique_index:idx_key"`
// 	HandleResult int32     `gorm:"column:handle_result;type:tinyint(1)"`
// 	ReqMsg       string    `gorm:"column:req_msg;type:varchar(255)"`
// 	CreateTime   time.Time `gorm:"column:create_time;type:datetime"`
// 	HandleUserId string    `gorm:"column:handle_user_id;type:varchar(12)"`
// 	HandleMsg    string    `gorm:"column:handle_msg;type:varchar(255)"`
// 	HandleTime   time.Time `gorm:"column:handle_time;typd:datetime"`
// 	Ex           string    `gorm:"column:ex;type:varchar(1024)"`
// }

// func (GroupRequest) TableName() string {
// 	return constant.DBTableGroupRequest
// }

type Black struct {
	OwnerUserId string    `gorm:"column:owner_user_id;type:varchar(12);not null;primary_key"`
	BlockUserId string    `gorm:"column:block_user_id;type:varchar(12);not null;primary_key"`
	CreateTime  time.Time `gorm:"column:create_time;type:datetime"`
	AddSource   int32     `gorm:"column:add_source;type:tinyint(1)"`
	OpUserId    string    `gorm:"column:op_user_id;type:varchar(12)"`
	Ex          string    `gorm:"column:ex;type:varchar(1024)"`
}

func (Black) TableName() string {
	return constant.DBTableBlack
}

type ChatLog struct {
	ServerMsgId    string `gorm:"column:server_msg_id;type:char(64);not null;primary_key"`
	SendId         string `gorm:"column:send_id;type:varchar(12)"`
	RecvId         string `gorm:"column:recv_id;type:varchar(12)"`
	SenderNickName string `gorm:"column:sender_nick_name;type:varchar(20)"`
	SenderFaceUrl  string `gorm:"column:sender_face_url;type:varchar(255)"`
	SessionType    int32  `gorm:"column:session_type"`
	MsgFrom        int32  `gorm:"column:msg_from"`
	ContentType    int32  `gorm:"column:content_type"`
	Content        string `gorm:"column:content;type:varchar(3000)"`
	FileSize       int64  `gorm:"column:file_size;type:bigint"`
	FileName       string `gorm:"column:file_name;type:varchar(255)"`
	FilePath       string `gorm:"column:file_path;type:varchar(255)"`
	FileType       int32  `gorm:"column:file_type;type:tinyint"`
	Status         int32  `gorm:"column:status;type:tinyint"`
	SendTime       int64  `gorm:"column:send_time;type:bigint"`
	CreateTime     int64  `gorm:"column:create_time;type:bigint"`
}

func (ChatLog) TableName() string {
	return constant.DBTableChatLog
}

type Conversation struct {
	OwnerUserId      string    `gorm:"column:owner_user_id;type:varchar(12);not null;primary_key"`
	ConversationId   string    `gorm:"column:conversation_id;type:varchar(24);not null;primary_key"`
	ConversationType int32     `gorm:"column:conversation_type;type:tinyint;not null"`
	ConversationName string    `gorm:"column:conversation_name;type:varchar(50)"`
	UserId           string    `gorm:"column:user_id;type:varchar(12)"`
	GroupId          string    `gorm:"column:group_id;type:varchar(12)"`
	MemberCount      int32     `gorm:"column:member_count;type:int"`
	NoReadCount      int32     `gorm:"column:no_read_count;type:int"`
	TopType          int32     `gorm:"column:toy_type;type:tinyint"`
	LastMessage      string    `gorm:"column:last_message;type:varchar(1024)"`
	LastMessageTime  time.Time `gorm:"column:last_message_time;type:datetime"`
	Status           int32     `gorm:"column:status;type:tinyint(1);default:1"`
	// RecvMsgOpt       int32     `gorm:"column:recv_msg_opt;type:tinyint"`
}

func (Conversation) TableName() string {
	return constant.DBTableConversation
}
