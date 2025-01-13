package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/captcha"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
)

func Register(c *gin.Context) {
	params := base_info.RegisterReq{}
	if err := c.BindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestParamsError,
			Info:   constant.FailInfo,
		})
		return
	}

	if !captcha.Captcha.Verify(params.CheckCodeID, params.CheckCode) {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestCheckCodeError,
			Info:   constant.FailInfo,
		})
		return
	}

}

type Data struct {
	ID   string `json:"id"`
	B64S string `json:"b64s"`
}

func GetCheckCode(c *gin.Context) {
	id, b64s := captcha.Captcha.Generate()

	resp := base_info.CommonResp{
		Status: constant.Success,
		Code:   constant.NoError,
		Info:   constant.SuccessInfo,
		Data: Data{
			ID:   id,
			B64S: b64s,
		},
	}
	c.JSON(http.StatusOK, resp)
}
