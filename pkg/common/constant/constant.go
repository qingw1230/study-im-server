package constant

// 数据库表名
const (
	DBTableUser            = "users"
	DBTableUserLuckyNumber = "user_lucky_numbers"
	DBTableFriend          = "friends"
	DBTableFriendRequest   = "friend_requests"
	DBTableGroup           = "groups"
	DBTableGroupMember     = "group_members"
	DBTableGroupRequest    = "group_requests"
	DBTableBlack           = "blacks"
	DBTableChatLog         = "chat_logs"
	DBTableConversation    = "conversations"
)

// 数字常量
const (
	LENGTH_10 = 10
	LENGTH_11 = 11
	LENGTH_12 = 12
	LENGTH_20 = 20
	LENGTH_32 = 32
)

// 字符串常量
const (
	STR_TOKEN              = "token"
	STR_SEND_ID            = "sendId"
	STR_SINGLE_            = "single_"
	STR_GROUP_             = "group_"
	ADD_FRIEND_DEFAULT_MSG = "我们已成功添加为好友，现在可以开始聊天啦～"
)

// token
const (
	NormalToken = iota + 1
	InvalidToken
	KickedToken
	ExpiredToken
)

const (
	// MsgFrom
	UserMsgType = 100
	SysMsgType  = 200

	// SessionType
	SingleChatType = 1
	GroupChatType  = 2

	// Websocket Protocol
	WSHeartBeat            = 1000
	WSGetNewestSeq         = 1001
	WSPullMsgBySeqList     = 1002
	WSSendMsg              = 1003
	WSPullConversationList = 1004
	WSPushMsg              = 2001
	WSKickOnlineMsg        = 2002
	WSDataError            = 3001

	// ContentType
	Text    = 101
	Picture = 102

	FriendRequestNotification         = 1201 // add_friend
	FriendRequestApprovedNotification = 1202 // add_friend_response
	FriendRequestRejectedNotification = 1203 // add_friend_response

	WSNoError    = 0
	WSNoUserConn = -1
	WSWriteError = -2
)
