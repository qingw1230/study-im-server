package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qingw1230/study-im-server/internal/api/account"
)

func Router() *gin.Engine {
	r := gin.Default()

	accountRouterGroup := r.Group("/account")
	{
		accountRouterGroup.POST("/register", account.Register)
		accountRouterGroup.POST("/login", account.Login)
		accountRouterGroup.POST("/get_check_code", account.GetCheckCode)
	}

	return r
}
