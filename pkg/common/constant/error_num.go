package constant

const (
	NoError     = 200
	CommonError = 10000

	// 请求参数相关的
	RequestParamsError      = 10001
	RequestBindJSONError    = 10002
	RequestCheckCodeError   = 10003
	RequestTokenAccessError = 10004

	// token 相关的
	TokenError       = 10100
	TokenExpired     = 10101
	TokenInvalid     = 10102
	TokenMalformed   = 10103
	TokenNotValidYet = 10104
	TokenUnknown     = 10105

	// MySQL 数据库相关的
	RecordMySQLCommonError  = 11000
	RecordAlreadyExists     = 11001
	RecordNotExists         = 11002
	RecordAccountORPwdError = 11003

	// msg 相关的
	MsgCommonError    = 12000
	MsgUnknownError   = 12001
	MsgKafkaSendError = 12002

	// ws 相关的
	WSCommonError    = 13000
	WSUnmarshalError = 13003
	WSValidateError  = 13004
)
