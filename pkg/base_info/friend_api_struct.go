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
	HandleResult int32 `json:"handleResult" binding:"required,oneof=-1 0 1"`
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
	UserId     string `json:"userId" binding:"required"`
	Pagination `json:"pagination"`
}

type GetFriendApplyListResp struct {
	CommonResp        `json:"commonResp"`
	FriendRequestList []*friend.FriendRequest `json:"-"`
	Total             int                     `json:"total"`
}

type AddBlacklistReq struct {
	CommonId
}

type AddBlacklistResp struct {
	CommonResp
}
