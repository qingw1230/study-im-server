package constant

// 数据库表名
const (
	DBTableUserInfo       = "user_info"
	DBTableUserInfoBeauty = "user_info_beauty"
	DBTableFriend         = "friend"
	DBTableFriendRequest  = "friend_request"
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
