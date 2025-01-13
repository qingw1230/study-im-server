package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/captcha"
	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/utils"
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

	_, err := controller.GetUserByEmail(params.Email)
	if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RecordAlreadyExists,
			Info:   constant.FailInfo,
		})
		return
	}

	controller.CreateUser(params.Email, params.NickName, params.Password)
}

func Login(c *gin.Context) {
	params := base_info.LoginReq{}
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

	user, err := controller.GetUserByEmail(params.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RecordNotExists,
			Info:   constant.FailInfo,
		})
		return
	}

	if !utils.ValidPassword(params.Password, user.Password) {
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RecordAccountORPwdError,
			Info:   constant.FailInfo,
		})
		return
	}

	flag := false
	for _, str := range config.Config.Admin.Emails {
		if str == params.Email {
			flag = true
			break
		}
	}

	token := utils.Md5Encode(user.UserID + utils.GenerateRandomStr(constant.LENGTH_20))

	// TODO(qingw1230): 多设备登录检测
	tokenStruct := base_info.TokenToRedis{
		Token:    token,
		UserID:   user.UserID,
		NickName: user.NickName,
		Admin:    flag,
	}
	db.DB.SetUserToken(tokenStruct)

	resp := base_info.LoginResp{
		Token:             token,
		Admin:             flag,
		UserID:            user.UserID,
		NickName:          user.NickName,
		PersonalSignature: user.PersonalSignature,
		JoinType:          user.JoinType,
		Sex:               user.Sex,
		AreaName:          user.AreaName,
		AreaCode:          user.AreaCode,
	}
	c.JSON(http.StatusOK, base_info.CommonResp{
		Status: constant.Success,
		Code:   constant.NoError,
		Info:   constant.SuccessInfo,
		Data:   resp,
	})
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
