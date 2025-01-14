package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qingw1230/study-im-server/internal/api/account"
	"github.com/qingw1230/study-im-server/internal/api/group"
)

func Router() *gin.Engine {
	r := gin.Default()

	accountRouterGroup := r.Group("/api/account")
	{
		accountRouterGroup.POST("/register", account.Register)
		accountRouterGroup.POST("/login", account.Login)
		accountRouterGroup.POST("/get_check_code", account.GetCheckCode)
	}

	groupRouterGroup := r.Group("/api/group")
	{
		groupRouterGroup.POST("/create_group", group.CreateGroup)
	}

	return r
}
