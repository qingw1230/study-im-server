package base_info

type CommonResp struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}
