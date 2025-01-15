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

const (
	Success     = "success"
	SuccessInfo = "请求成功"
	Fail        = "fail"
	FailInfo    = "请求失败"

	CheckCodeInfo            = "验证码错误"
	EmailAlreadyRegisterInfo = "邮箱已注册"
)

// 数据库表名
const (
	DBTableUserInfo       = "user_info"
	DBTableUserInfoBeauty = "user_info_beauty"
)

const (
	LENGTH_10 = 10
	LENGTH_11 = 11
	LENGTH_20 = 20
)
