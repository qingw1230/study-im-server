package constant

import (
	"errors"
)

const (
	Success     = "success"
	SuccessInfo = "请求成功"
	Fail        = "fail"
	FailInfo    = "请求失败"
)

const (
	RequestCheckCodeErrorInfo   = "验证码错误"
	RequestTokenAccessErrorInfo = "没有访问权限"

	MySQLCommonErrorInfo               = "MySQL 数据库通用错误"
	MySQLRecordNotExistsInfo           = "MySQL 相应记录不存在"
	MySQLEmailAlreadyRegisterErrorInfo = "邮箱已注册"
	MySQLAccountORPwdErrorInfo         = "账号或密码错误"
	MySQLAccountNotExistsInfo          = "账号不存在"

	MsgCommonErrorInfo    = "msg 通用错误"
	MsgUnknownErrorInfo   = "未知的 sessionTye"
	MsgKafkaSendErrorInfo = "向 kafka 发送失败"

	WSCommonErrorInfo    = "ws 通用错误"
	WSUnmarshalErrorInfo = "反序列化失败"
	WSValidateErrorInfo  = "字段验证失败"

	MongoCommonErrorInfo = "Mongo 数据库通用错误"
)

var (
	CreateTokenMsg      = errors.New("create token failed")
	ParseTokenMsg       = errors.New("parse token failed")
	TokenExpiredMsg     = errors.New("token is timed out, please login again")
	TokenInvalidMsg     = errors.New("token has been invalidated")
	TokenMalformedMsg   = errors.New("that's not even a token")
	TokenNotValidYetMsg = errors.New("token not active yet")
	TokenUnknownMsg     = errors.New("couldn't handle this token")
)
