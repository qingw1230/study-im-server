package msg

import (
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
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
	case constant.FriendRequestApprovedNotification:
		contentType = constant.Text
		content = []byte(m.(*pbFriend.FriendRequestApprovedTips).Msg)
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
	friendRequestTips := &pbFriend.FriendRequestTips{
		ReqMsg: reqMsg,
	}
	friendNotification(req.CommonId, constant.FriendRequestNotification, friendRequestTips)
}

func FriendRequestApproveNotification(req *pbFriend.AddFriendResponseReq, fr *db.FriendRequest) {
	log.Info("call FriendRequestApproveNotification")
	friendRequestApproved := &pbFriend.FriendRequestApprovedTips{
		Msg: fr.ReqMsg,
	}
	commonId := &pbFriend.CommonId{
		OpUserId:   fr.ToUserId,
		FromUserId: fr.FromUserId,
		ToUserId:   fr.ToUserId,
	}
	friendNotification(commonId, constant.FriendRequestApprovedNotification, friendRequestApproved)
	friendRequestApproved.Msg = constant.ADD_FRIEND_DEFAULT_MSG
	commonId.FromUserId, commonId.ToUserId = commonId.ToUserId, commonId.FromUserId
	friendNotification(commonId, constant.FriendRequestApprovedNotification, friendRequestApproved)
}
