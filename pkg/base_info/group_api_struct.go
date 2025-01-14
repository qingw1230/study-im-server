package base_info

import "mime/multipart"

type CreateGroup struct {
	GroupName   string         `json:"groupName" binding:"required"`
	GroupNotice string         `json:"groupNotice"`
	JoinType    int            `json:"joinType"`
	AvatarFile  multipart.File `json:"avatarFile"`
	AvatarCover multipart.File `json:"avatarCover"`
}
