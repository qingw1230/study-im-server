package constant

const (
	NoError     = 200
	CommonError = 10000

	// token 相关的
	TokenError       = 700
	TokenExpired     = 701
	TokenInvalid     = 702
	TokenMalformed   = 703
	TokenNotValidYet = 704
	TokenUnknown     = 705

	// 请求参数相关的
	RequestParamsError      = 10001
	RequestBindJSONError    = 10002
	RequestCheckCodeError   = 10003
	RequestTokenAccessError = 10004

	// MySQL 数据库相关的
	MySQLCommonError         = 11000
	MySQLRecordAlreadyExists = 11001
	MySQLRecordNotExists     = 11002
	MySQLAccountORPwdError   = 11003

	// msg 相关的
	MsgCommonError    = 12000
	MsgUnknownError   = 12001
	MsgKafkaSendError = 12002

	// ws 相关的
	WSCommonError    = 13000
	WSUnmarshalError = 13003
	WSValidateError  = 13004

	// Mongo 数据库相关的
	MongoCommonError = 14000

	// Redis 数据库相关的
	RedisCommonError = 15000
)
