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
)

// token
const (
	NormalToken = iota
	InvalidToken
	KickedToken
	ExpiredToken
)
