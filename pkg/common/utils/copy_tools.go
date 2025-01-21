package utils

import (
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	pbPublic "github.com/qingw1230/study-im-server/pkg/proto/public"
)

func FriendDBCopyIM(dst *pbPublic.FriendInfo, src *db.Friend) error {
	copier.Copy(dst, src)
	user, err := controller.FindUserById(src.FriendUserId)
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
