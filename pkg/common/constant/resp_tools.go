package constant

import "github.com/qingw1230/study-im-server/pkg/base_info"

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
		Code:   RecordMySQLCommonError,
		Info:   RecordMySQLCommonErrorInfo,
	}
)

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
