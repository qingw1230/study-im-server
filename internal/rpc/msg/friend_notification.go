package msg

import (
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbFriend "github.com/qingw1230/study-im-server/pkg/proto/friend"
	"google.golang.org/protobuf/proto"
)

func friendNotification(commonId *pbFriend.CommonId, contentType int32, m proto.Message) {
	log.Info("call friendNotification args:", commonId, contentType)

	var content []byte
	switch contentType {
	case constant.FriendRequestNotification:
		content = []byte(m.(*pbFriend.FriendRequestTips).ReqMsg)
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
func FriendRequestNotification(req *pbFriend.AddFriendReq, reqMsg string) {
	friendRequestTips := pbFriend.FriendRequestTips{
		ReqMsg: reqMsg,
	}
	friendNotification(req.CommonId, constant.FriendRequestNotification, &friendRequestTips)
}
