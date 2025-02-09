package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qingw1230/study-im-server/internal/api/account"
	"github.com/qingw1230/study-im-server/internal/api/conversation"
	"github.com/qingw1230/study-im-server/internal/api/friend"
	"github.com/qingw1230/study-im-server/internal/api/group"
	"github.com/qingw1230/study-im-server/internal/api/msg"
)

func Router() *gin.Engine {
	r := gin.Default()

	accountRouterGroup := r.Group("/api/account")
	{
		accountRouterGroup.POST("/register", account.Register)
		accountRouterGroup.POST("/login", account.Login)
		accountRouterGroup.POST("/get_check_code", account.GetCheckCode)
		accountRouterGroup.POST("/update_user_info", account.UpdateUserInfo)
		accountRouterGroup.POST("/get_user_info", account.GetUserInfo)
		accountRouterGroup.POST("/get_self_user_info", account.GetSelfUserInfo)
	}

	friendRouterGroup := r.Group("/api/friend")
	{
		friendRouterGroup.POST("/get_friend_list", friend.GetFriendList)
		friendRouterGroup.POST("/get_friend_apply_list", friend.GetFriendApplyList)
		friendRouterGroup.POST("/add_friend", friend.AddFriend)
		friendRouterGroup.POST("/add_friend_response", friend.AddFriendResponse)
		friendRouterGroup.POST("/delete_friend", friend.DeleteFriend)

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

	msgGroup := r.Group("/api/msg")
	{
		msgGroup.POST("/get_newest_seq", msg.GetNewestSeq)
		msgGroup.POST("/send_msg", msg.SendMsg)
		msgGroup.POST("/pull_msg_by_seq_list", msg.PullMsgBySeqList)
	}

	conversationGroup := r.Group("/api/conversation")
	{
		conversationGroup.POST("/get_conversation_list", conversation.GetConversationList)
	}

	return r
}
