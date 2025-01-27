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
	STR_TOKEN = "token"
	SINGLE_   = "single_"
	GROUP_    = "group_"
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

	// MsgReceiveOpt
	ReceiveMessage          = 1
	NotReceiveMessage       = 2
	ReceiveNotNotifyMessage = 3
)
