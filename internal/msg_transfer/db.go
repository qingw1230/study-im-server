package msg_transfer

import (
	"github.com/qingw1230/study-im-server/pkg/common/db"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
)

func saveUserChat(userId string, msg *pbMsg.MsgDataToMq) error {
	pbSaveData := pbMsg.MsgDataToDb{
		MsgData: msg.MsgData,
	}
	return db.DB.SaveUserChat(userId, pbSaveData.MsgData.SendTime, &pbSaveData)
}
