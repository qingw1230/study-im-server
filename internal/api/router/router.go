package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qingw1230/study-im-server/internal/api/account"
	"github.com/qingw1230/study-im-server/internal/api/friend"
	"github.com/qingw1230/study-im-server/internal/api/group"
)

func Router() *gin.Engine {
	r := gin.Default()

	accountRouterGroup := r.Group("/api/account")
	{
		accountRouterGroup.POST("/register", account.Register)
		accountRouterGroup.POST("/login", account.Login)
		accountRouterGroup.POST("/get_check_code", account.GetCheckCode)
		accountRouterGroup.POST("/get_user_info", account.GetUserInfo)
	}

	friendRouterGroup := r.Group("/api/friend")
	{
		friendRouterGroup.POST("/add_friend", friend.AddFriend)
		friendRouterGroup.POST("/add_friend_response", friend.AddFriendResponse)
		friendRouterGroup.POST("/delete_friend", friend.DeleteFriend)
		friendRouterGroup.POST("/get_friend_list", friend.GetFriendList)
		friendRouterGroup.POST("/get_friend_apply_list", friend.GetFriendApplyList)

		friendRouterGroup.POST("/add_blacklist", friend.AddBlacklist)
	}

	groupRouterGroup := r.Group("/api/group")
	{
		groupRouterGroup.POST("/create_group", group.CreateGroup)
		groupRouterGroup.POST("/delete_group", group.DeleteGroup)
		groupRouterGroup.POST("/quit_group", group.QuitGroup)
		groupRouterGroup.POST("/get_joined_group_list", group.GetJoinedGroupList)
		groupRouterGroup.POST("/get_group_info", group.GetGroupInfo)
		groupRouterGroup.POST("/set_group_info", group.SetGroupInfo)
	}

	return r
}
