package msg_transfer

import (
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
)

func saveUserChat(userId string, msg *pbMsg.MsgDataToMq) error {
	seq, err := db.DB.IncrUserSeq(userId)
	if err != nil {
		log.Error("redis IncrUserSeq failed", err.Error(), userId)
		return err
	}
	msg.MsgData.Seq = uint32(seq)
	pbSaveData := pbMsg.MsgDataToDb{
		MsgData: msg.MsgData,
	}
	return db.DB.SaveUserChat(userId, pbSaveData.MsgData.SendTime, &pbSaveData)
}
