package msg_gateway

type Req struct {
	ReqIdentifier int32  `json:"reqIdentifier" validate:"required"`
	Token         string `json:"token"`
	SendId        string `json:"sendId" validate:"required"`
	MsgIncr       string `json:"msgIncr" validate:"required"`
	Data          []byte `json:"data"`
}

type Resp struct {
	ReqIdentifier int32  `json:"reqIdentifier"`
	MsgIncr       string `json:"msgIncr"`
	Code          int32  `json:"code"`
	Info          string `json:"info"`
	Data          []byte `json:"data"`
}

// func (ws *WsServer) argsValidate(m *Req, r int32) (isPass bool, code int32, info string, returnData interface{}) {
// 	switch r {
// 	case constant.WSSendMsg:
// 		data := pbPublic.MsgData{}
// 		if err := proto.Unmarshal(m.Data, &data); err != nil {
// 			log.Error("Unmarshal failed", err.Error(), "reqIdentifier", r)
// 			return false, constant.WSUnmarshalError, constant.WSUnmarshalErrorInfo, nil
// 		}
// 		if err := validate.Struct(&data); err != nil {
// 			log.Error("validate failed", err.Error(), "reqIdentifier", r)
// 			return false, constant.WSValidateError, constant.WSValidateErrorInfo, nil
// 		}
// 		return true, constant.NoError, constant.SuccessInfo, &data
// 	default:
// 	}

// 	return false, constant.WSValidateError, constant.WSValidateErrorInfo, nil
// }
