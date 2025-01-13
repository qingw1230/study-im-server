package base_info

type RegisterReq struct {
	Email       string `json:"email" binding:"required,email"`
	NickName    string `json:"nickName" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CheckCode   string `json:"checkCode" binding:"required"`
	CheckCodeID string `json:"checkCodeId" binding:"required"`
}

type LoginReq struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	CheckCode   string `json:"checkCode" binding:"required"`
	CheckCodeID string `json:"checkCodeId" binding:"required"`
}

type LoginResp struct {
	Token             string `json:"token"`
	Admin             bool   `json:"admin"`
	UserID            string `json:"userId"`
	NickName          string `json:"nickName"`
	PersonalSignature string `json:"personalSignature"`
	JoinType          int    `json:"joinType"`
	Sex               int    `json:"sex"`
	AreaName          string `json:"areaName"`
	AreaCode          string `json:"areaCode"`
}

type TokenToRedis struct {
	Token    string `json:"token"`
	UserID   string `json:"userId"`
	NickName string `json:"nickName"`
	Admin    bool   `json:"admin"`
}
