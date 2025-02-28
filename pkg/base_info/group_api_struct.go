package base_info

import (
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
)

type CreateGroupReq struct {
	MemberList  []GroupAddMemberInfo `json:"memberList"`
	OwnerUserId string               `json:"ownerUserId" binding:"required"`
	GroupInfo   sdkws.GroupInfo      `json:"groupInfo" binding:"required"`
}

type CreateGroupResp struct {
	CommonResp
	// GroupInfo sdkws.GroupInfo `json:"-"`
}

type DeleteGroupReq struct {
	GroupId string `json:"groupId" binding:"required"`
}

type DeleteGroupResp struct {
	CommonResp
}

type QuitGroupReq struct {
	GroupId string `json:"groupId" binding:"required"`
}

type QuitGroupResp struct {
	CommonResp
}

type SetGroupInfoReq struct {
	GroupID      string `json:"groupId" binding:"required"`
	GroupName    string `json:"groupName"`
	FaceUrl      string `json:"faceUrl"`
	Introduction string `json:"introduction"`
	Notification string `json:"notification"`
}

type SetGroupInfoResp struct {
	CommonResp
}

type GetJoinedGroupListReq struct {
	FromUserId string `json:"fromUserId" binding:"required"`
	RoleLevel  int32  `json:"roleLevel" binding:"required"`
}

type GetJoinedGroupListResp struct {
	CommonResp
	GroupInfoList []*sdkws.GroupInfo `json:"-"`
}

type GetGroupInfoReq struct {
	GroupId string `json:"groupId" binding:"required"`
}

type GetGroupInfoResp struct {
	CommonResp
	GroupInfo *sdkws.GroupInfo `json:"-"`
}
