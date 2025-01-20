package base_info

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
