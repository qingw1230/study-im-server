package base_info

import "github.com/qingw1230/study-im-server/pkg/proto/public"

type CommonID struct {
	FromUserID string `json:"fromUserID" binding:"required"`
	ToUserID   string `json:"toUserID" binding:"required"`
}

type AddFriendReq struct {
	CommonID
	ReqMsg string `json:"reqMsg"`
}

type AddFriendResp struct {
	CommonResp
}

type AddFriendResponseReq struct {
	CommonID
	HandleResult int32  `json:"handleResult" binding:"required,oneof=-1 0 1"`
	HandleMsg    string `json:"handleMsg"`
}

type AddFriendResponseResp struct {
	CommonResp
}

type DeleteFriendReq struct {
	CommonID
}

type DeleteFriendResp struct {
	CommonResp
}

type GetFriendListReq struct {
	FromUserID string `json:"fromUserID" binding:"required"`
}

type GetFriendListResp struct {
	CommonResp
	FriendInfoList []*public.FriendInfo `json:"-"`
}
