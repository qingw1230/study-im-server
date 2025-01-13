package base_info

type RegisterReq struct {
	Email       string `json:"email" binding:"required,email"`
	NickName    string `json:"nickName" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CheckCode   string `json:"checkCode" binding:"required"`
	CheckCodeID string `json:"checkCodeId" binding:"required"`
}
