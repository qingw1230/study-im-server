package base_info

import "github.com/qingw1230/study-im-server/pkg/proto/public"

type CreateGroupReq struct {
	OwnerUserId  string `json:"ownerUserId" binding:"required"`
	GroupName    string `json:"groupName" binding:"required"`
	GroupType    int32  `json:"groupType"`
	FaceUrl      string `json:"faceUrl"`
	Introduction string `json:"introduction"`
	Notification string `json:"notification"`
}

type CreateGroupResp struct {
	CommonResp
}

type GetJoinedGroupListReq struct {
	FromUserId string `json:"fromUserId" binding:"required"`
	RoleLevel  int32  `json:"roleLevel" binding:"required"`
}

type GetJoinedGroupListResp struct {
	CommonResp
	GroupInfoList []*public.GroupInfo `json:"-"`
}

type GetGroupInfoReq struct {
	GroupId string `json:"groupId" binding:"required"`
}

type GetGroupInfoResp struct {
	CommonResp
	GroupInfo *public.GroupInfo `json:"-"`
}
