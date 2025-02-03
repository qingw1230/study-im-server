package constant

import (
	"github.com/qingw1230/study-im-server/pkg/base_info"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
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
	PBCommonSuccessResp = pbPublic.CommonResp{
		Status: Success,
		Code:   NoError,
		Info:   SuccessInfo,
	}

	PBCommonFailResp = pbPublic.CommonResp{
		Status: Fail,
		Code:   CommonError,
		Info:   FailInfo,
	}

	PBTokenAccessErrorResp = pbPublic.CommonResp{
		Status: Fail,
		Code:   RequestTokenAccessError,
		Info:   RequestTokenAccessErrorInfo,
	}

	PBMySQLCommonFailResp = pbPublic.CommonResp{
		Status: Fail,
		Code:   MySQLCommonError,
		Info:   MySQLCommonErrorInfo,
	}

	PBMongoCommonFailResp = pbPublic.CommonResp{
		Status: Fail,
		Code:   MongoCommonError,
		Info:   MongoCommonErrorInfo,
	}

	PBRedisCommonFailResp = pbPublic.CommonResp{
		Status: Fail,
		Code:   RedisCommonError,
		Info:   RedisCommonErrorInfo,
	}
)

func NewPBResp(status string, code int, info string) *pbPublic.CommonResp {
	resp := &pbPublic.CommonResp{
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
