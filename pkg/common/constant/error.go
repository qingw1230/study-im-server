package constant

const (
	NoError = 200

	// 请求参数相关的
	RequestParamsError    = 10001
	RequestCheckCodeError = 10002

	// 数据库记录相关的
	RecordAlreadyExists     = 11001
	RecordNotExists         = 11002
	RecordAccountORPwdError = 11003
)
