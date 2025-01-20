package constant

// 数据库表名
const (
	DBTableUser            = "users"
	DBTableUserLuckyNumber = "user_lucky_numbers"
	DBTableFriend          = "friends"
	DBTableFriendRequest   = "friend_requests"
	DBTableGroup           = "groups"
	DBTableGroupMember     = "group_members"
)

// 数字常量
const (
	LENGTH_10 = 10
	LENGTH_11 = 11
	LENGTH_20 = 20
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
