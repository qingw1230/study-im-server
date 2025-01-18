package utils

import (
	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/proto/public"
)

func FriendDBCopyIM(dst *public.FriendInfo, src *db.Friend) error {
	copier.Copy(dst, src)
	user, err := controller.FindUserByID(src.FriendUserID)
	if err != nil {
		return err
	}
	copier.Copy(dst.FriendInfo, user)
	dst.CreateTime = uint32(src.CreateTime.Unix())
	return nil
}
