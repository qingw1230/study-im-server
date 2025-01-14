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
)

func NewSuccessRespWithData(data interface{}) *base_info.CommonResp {
	resp := &base_info.CommonResp{
		Status: Success,
		Code:   NoError,
		Info:   SuccessInfo,
	}
	resp.Data = data
	return resp
}

const (
	Success     = "success"
	SuccessInfo = "请求成功"
	Fail        = "fail"
	FailInfo    = "请求失败"
)

const (
	RequestCheckCodeErrorInfo = "验证码错误"

	RecordEmailAlreadyRegisterErrorInfo = "邮箱已注册"
	RecordAccountORPwdErrorInfo         = "账号或密码错误"
	RecordAccountNotExistsInfo          = "账号不存在"
)
