package constant

// 数据库表名
const (
	DBTableUserInfo       = "user_info"
	DBTableUserInfoBeauty = "user_info_beauty"
	DBTableFrindRequest   = "friend_request"
)

// 数字常量
const (
	LENGTH_10 = 10
	LENGTH_11 = 11
	LENGTH_20 = 20
)

// token
const (
	NormalToken = iota
	InvalidToken
	KickedToken
	ExpiredToken
)
