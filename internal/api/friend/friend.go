package friend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qingw1230/study-im-server/pkg/base_info"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
)

func AddFriend(c *gin.Context) {
	params := base_info.AddFriendReq{}
	if err := c.BindJSON(&params); err != nil {
		log.Error("BindJSON failed ", err.Error())
		c.JSON(http.StatusBadRequest, base_info.CommonResp{
			Status: constant.Fail,
			Code:   constant.RequestBindJSONError,
			Info:   err.Error(),
		})
		return
	}
	log.Info("AddFriend BindJSON success")
}
