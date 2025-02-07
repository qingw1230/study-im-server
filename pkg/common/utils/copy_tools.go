package utils

import (
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	pbFriend "github.com/qingw1230/study-im-server/pkg/proto/friend"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
)

func FriendDBCopyIM(dst *pbPublic.FriendInfo, src *db.Friend) error {
	copier.Copy(dst, src)
	user, err := controller.GetUserById(src.FriendUserId)
	if err != nil {
		return err
	}
	copier.Copy(dst.FriendInfo, user)
	dst.CreateTime = uint32(src.CreateTime.Unix())
	return nil
}

func GroupDBCopyIM(dst *pbPublic.GroupInfo, src *db.Group) error {
	copier.Copy(dst, src)
	user, err := controller.GetGroupOwnerInfoByGroupId(src.GroupId)
	if err != nil {
		return err
	}
	dst.OwnerUserId = user.UserId
	dst.MemberCount, err = controller.GetGroupMemberNumByGroupId(src.GroupId)
	if err != nil {
		return err
	}
	dst.CreateTime = uint32(src.CreateTime.Unix())
	return nil
}

// FriendRequestDBCopyIM 将 DB 的 FriendRequest 的数据拷贝到 pb 的 FriendRequest 中
func FriendRequestDBCopyIM(dst *pbFriend.FriendRequest, src *db.FriendRequest) error {
	copier.Copy(dst, src)
	user, err := controller.GetUserById(src.FromUserId)
	if err != nil {
		return err
	}
	dst.FromNickName = user.NickName
	dst.FromFaceURL = user.FaceUrl
	user, err = controller.GetUserById(src.ToUserId)
	dst.ToNickName = user.NickName
	dst.ToFaceURL = user.FaceUrl
	dst.CreateTime = src.CreateTime
	dst.HandleTime = src.HandleTime
	return nil
}
