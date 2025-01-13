package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/captcha"
)

type Data struct {
	ID   string `json:"id"`
	B64S string `json:"b64s"`
}

func GetCheckCode(c *gin.Context) {
	cap := captcha.NewCaptcha()
	id, b64s := cap.Generate()

	resp := base_info.CommonResp{
		Status: "success",
		Code:   200,
		Info:   "请求成功",
		Data: Data{
			ID:   id,
			B64S: b64s,
		},
	}
	c.JSON(http.StatusOK, resp)
}
