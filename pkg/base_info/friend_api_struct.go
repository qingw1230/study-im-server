package base_info

import (
	"github.com/qingw1230/study-im-server/pkg/proto/friend"
	"github.com/qingw1230/study-im-server/pkg/proto/public"
)

type CommonId struct {
	FromUserId string `json:"fromUserId" binding:"required"`
	ToUserId   string `json:"toUserId" binding:"required"`
}

type AddFriendReq struct {
	CommonId
	ReqMsg string `json:"reqMsg" binding:"required"`
}

type AddFriendResp struct {
	CommonResp
}

type AddFriendResponseReq struct {
	CommonId
	HandleResult int32  `json:"handleResult" binding:"required,oneof=-1 0 1"`
	HandleMsg    string `json:"handleMsg"`
}

type AddFriendResponseResp struct {
	CommonResp
}

type DeleteFriendReq struct {
	CommonId
}

type DeleteFriendResp struct {
	CommonResp
}

type GetFriendListReq struct {
	FromUserId string `json:"fromUserId" binding:"required"`
}

type GetFriendListResp struct {
	CommonResp
	FriendInfoList []*public.FriendInfo `json:"-"`
}

type GetFriendApplyListReq struct {
	FromUserID string `json:"fromUserId" binding:"required"`
}

type GetFriendApplyListResp struct {
	CommonResp
	FriendRequestList []*friend.FriendRequest `json:"-"`
}

type AddBlacklistReq struct {
	CommonId
}

type AddBlacklistResp struct {
	CommonResp
}
