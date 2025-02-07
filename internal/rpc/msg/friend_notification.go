package msg

import (
	"encoding/json"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/db/controller"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbFriend "github.com/qingw1230/study-im-server/pkg/proto/friend"
	"google.golang.org/protobuf/proto"
)

func getFromToUserNickNameAndFaceUrl(fromUserID, toUserID string) (string, string, string, string, error) {
	from, err := controller.GetUserById(fromUserID)
	if err != nil {
		return "", "", "", "", err
	}
	to, err := controller.GetUserById(toUserID)
	if err != nil {
		return "", "", "", "", err
	}
	return from.NickName, from.FaceUrl, to.NickName, to.FaceUrl, nil
}

func friendNotification(commonId *pbFriend.CommonId, contentType int32, m proto.Message) {
	log.Info("call friendNotification args:", commonId, contentType)

	var content []byte
	var err error
	switch contentType {
	case constant.FriendRequestNotification:
		content, err = json.Marshal(m)
		if err != nil {
			log.Error("json.Marshal failed")
			return
		}
	default:
		log.Error("contentType failed", contentType)
		return
	}

	n := NotificationMsg{
		SendId:      commonId.FromUserId,
		RecvId:      commonId.ToUserId,
		MsgFrom:     constant.SysMsgType,
		SessionType: constant.SingleChatType,
		ContentType: contentType,
		Content:     content,
	}
	Notification(&n)
	log.Info("friendNotification return")
}

// FriendRequestNotification 添加好友的通知
func FriendRequestNotification(req *pbFriend.AddFriendReq, fr *db.FriendRequest) {
	friendRequestTips := pbFriend.FriendRequestTips{
		FromUserId: fr.FromUserId,
		ToUserId:   fr.ToUserId,
		ReqMsg:     fr.ReqMsg,
	}
	fromUser, err := controller.GetUserById(fr.FromUserId)
	if err != nil {
		log.Error("GetUserById failed", err.Error(), fr.FromUserId)
		return
	}
	friendRequestTips.FromNickName = fromUser.NickName
	friendRequestTips.FromFaceURL = fromUser.FaceUrl
	friendNotification(req.CommonId, constant.FriendRequestNotification, &friendRequestTips)
}
