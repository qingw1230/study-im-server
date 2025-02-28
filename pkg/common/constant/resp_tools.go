package constant

import (
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
)

var (
	CommonSuccessResp = base_info.CommonResp{
		Status: Success,
		Code:   NoError,
		Info:   SuccessInfo,
	}

	CommonFailResp = base_info.CommonResp{
		Status: Fail,
		Code:   CommonError,
		Info:   FailInfo,
	}

	MySQLCommonFailResp = base_info.CommonResp{
		Status: Fail,
		Code:   MySQLCommonError,
		Info:   MySQLCommonErrorInfo,
	}

	MongoCommonFailResp = base_info.CommonResp{
		Status: Fail,
		Code:   MongoCommonError,
		Info:   MongoCommonErrorInfo,
	}
)

var (
	PBCommonSuccessResp = sdkws.CommonResp{
		Status: Success,
		Code:   NoError,
		Info:   SuccessInfo,
	}

	PBCommonFailResp = sdkws.CommonResp{
		Status: Fail,
		Code:   CommonError,
		Info:   FailInfo,
	}

	PBTokenAccessErrorResp = sdkws.CommonResp{
		Status: Fail,
		Code:   RequestTokenAccessError,
		Info:   RequestTokenAccessErrorInfo,
	}

	PBMySQLCommonFailResp = sdkws.CommonResp{
		Status: Fail,
		Code:   MySQLCommonError,
		Info:   MySQLCommonErrorInfo,
	}

	PBMongoCommonFailResp = sdkws.CommonResp{
		Status: Fail,
		Code:   MongoCommonError,
		Info:   MongoCommonErrorInfo,
	}

	PBRedisCommonFailResp = sdkws.CommonResp{
		Status: Fail,
		Code:   RedisCommonError,
		Info:   RedisCommonErrorInfo,
	}
)

func NewPBResp(status string, code int, info string) *sdkws.CommonResp {
	resp := &sdkws.CommonResp{
		Status: status,
		Code:   int32(code),
		Info:   info,
	}
	return resp
}

func NewRespNoData(status string, code int, info string) *base_info.CommonResp {
	resp := &base_info.CommonResp{
		Status: status,
		Code:   code,
		Info:   info,
	}
	return resp
}

func NewRespWithData(status string, code int, info string, data interface{}) *base_info.CommonResp {
	resp := &base_info.CommonResp{
		Status: status,
		Code:   code,
		Info:   info,
		Data:   data,
	}
	return resp
}

func NewSuccessRespWithData(data interface{}) *base_info.CommonResp {
	resp := &base_info.CommonResp{
		Status: Success,
		Code:   NoError,
		Info:   SuccessInfo,
		Data:   data,
	}
	return resp
}

func NewBindJSONErrorRespWithInfo(info string) *base_info.CommonResp {
	resp := &base_info.CommonResp{
		Status: Fail,
		Code:   RequestBindJSONError,
		Info:   info,
	}
	return resp
}
