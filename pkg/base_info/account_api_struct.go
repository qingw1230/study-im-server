package base_info

import "github.com/qingw1230/study-im-server/pkg/proto/public"

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

type GetUserInfoReq struct {
	UserId string `json:"userId" binding:"required"`
}

type GetUserInfoResp struct {
	PublicUserInfo *public.PublicUserInfo
}

type GetSelfUserInfoReq struct {
	UserId string `json:"userId" binding:"required"`
}

type GetSelfUserInfoResp struct {
	CommonResp
	UserInfo *public.UserInfo `json:"-"`
}
