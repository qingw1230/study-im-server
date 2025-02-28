package base_info

import "github.com/qingw1230/study-im-server/pkg/proto/sdkws"

type RegisterReq struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	NickName    string `json:"nickName" binding:"required"`
	CheckCode   string `json:"checkCode" binding:"required"`
	CheckCodeId string `json:"checkCodeId" binding:"required"`
}

type LoginReq struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	CheckCode   string `json:"checkCode" binding:"required"`
	CheckCodeId string `json:"checkCodeId" binding:"required"`
}

type LoginResp struct {
	Token             string `json:"token"`
	Admin             bool   `json:"admin"`
	UserId            string `json:"userId"`
	NickName          string `json:"nickName"`
	PersonalSignature string `json:"personalSignature"`
	JoinType          int    `json:"joinType"`
	Sex               int    `json:"sex"`
	AreaName          string `json:"areaName"`
	AreaCode          string `json:"areaCode"`
}

type UpdateUserInfoReq struct {
	UserId            string `json:"userId" binding:"required,min=1,max=12"`
	NickName          string `json:"nickName" binding:"omitempty,min=1,max=20"`
	FaceUrl           string `json:"faceUrl" binding:"omitempty,max=255"`
	PersonalSignature string `json:"personalSignature" binding:"omitempty,max=50"`
	Sex               int32  `json:"sex" binding:"omitempty,oneof=0 1 2"`
	AreaName          string `json:"areaName" binding:"omitempty,max=50"`
	AreaCode          string `json:"areaCode" binding:"omitempty,max=50"`
	Email             string `json:"email" binding:"omitempty,max=50"`
	Password          string `json:"password" binding:"omitempty,min=8,max=32"`
}

type UpdateUserInfoResp struct {
	CommonResp
}

type GetUserInfoReq struct {
	UserId string `json:"userId" binding:"required"`
}

type GetUserInfoResp struct {
	PublicUserInfo *sdkws.PublicUserInfo
}

type GetSelfUserInfoReq struct {
	UserId string `json:"userId" binding:"required"`
}

type GetSelfUserInfoResp struct {
	CommonResp
	UserInfo *sdkws.UserInfo `json:"-"`
}
