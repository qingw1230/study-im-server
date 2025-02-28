package msg_gateway

import (
	"encoding/json"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbConversation "github.com/qingw1230/study-im-server/pkg/proto/conversation"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
)

type Req struct {
	ReqIdentifier int32  `json:"reqIdentifier" validate:"required"`
	Token         string `json:"token"`
	SendId        string `json:"sendId" validate:"required"`
	Data          []byte `json:"data"`
}

type Resp struct {
	ReqIdentifier int32  `json:"reqIdentifier"`
	Code          int32  `json:"code"`
	Info          string `json:"info"`
	Data          []byte `json:"data"`
}

// argsValidate 解析 Req 的 Data 部分，解析成相应 rpc 请求结构
func (ws *WsServer) argsValidate(m *Req, r int32) (isPass bool, code int32, info string, returnData interface{}) {
	switch r {
	case constant.WSPullMsgBySeqList:
		data := pbMsg.PullMessageBySeqListReq{}
		if err := json.Unmarshal(m.Data, &data); err != nil {
			log.Error("Unmarshal failed", err.Error(), "reqIdentifier", r)
			return false, constant.WSUnmarshalError, constant.WSUnmarshalErrorInfo, nil
		}
		if err := validate.Struct(&data); err != nil {
			log.Error("validate failed", err.Error(), "reqIdentifier", r)
			return false, constant.WSValidateError, constant.WSValidateErrorInfo, nil
		}
		return true, constant.NoError, constant.SuccessInfo, &data
	case constant.WSSendMsg:
		data := sdkws.MsgData{}
		if err := json.Unmarshal(m.Data, &data); err != nil {
			log.Error("Unmarshal failed", err.Error(), "reqIdentifier", r)
			return false, constant.WSUnmarshalError, constant.WSUnmarshalErrorInfo, nil
		}
		if err := validate.Struct(&data); err != nil {
			log.Error("validate failed", err.Error(), "reqIdentifier", r)
			return false, constant.WSValidateError, constant.WSValidateErrorInfo, nil
		}
		return true, constant.NoError, constant.SuccessInfo, &data
	case constant.WSPullConversationList:
		data := pbConversation.GetConversationListReq{}
		if err := json.Unmarshal(m.Data, &data); err != nil {
			log.Error("Unmarshal failed", err.Error(), "reqIdentifier", r)
			return false, constant.WSUnmarshalError, constant.WSUnmarshalErrorInfo, nil
		}
		if err := validate.Struct(&data); err != nil {
			log.Error("validate failed", err.Error(), "reqIdentifier", r)
			return false, constant.WSValidateError, constant.WSValidateErrorInfo, nil
		}
		return true, constant.NoError, constant.SuccessInfo, &data
	default:
	}

	return false, constant.WSValidateError, constant.WSValidateErrorInfo, nil
}
