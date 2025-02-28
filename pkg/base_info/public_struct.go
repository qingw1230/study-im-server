package base_info

type CommonResp struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

type Pagination struct {
	PageNumber int32 `json:"pageNumber" binding:"required"`
	ShowNumber int32 `json:"showNumber" binding:"required"`
}

type GroupAddMemberInfo struct {
	UserId    string `json:"userId" binding:"required"`
	RoleLevel int32  `json:"roleLevel" binding:"required"`
}
