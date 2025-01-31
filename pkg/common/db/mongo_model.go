package db

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"gopkg.in/mgo.v2/bson"
)

const (
	singleChat      = "msg"
	groupChat       = "group"
	singleGocMsgNum = 5000
)

type MsgInfo struct {
	SendTime int64
	Msg      []byte
}

type UserChat struct {
	UserId string
	Msg    []MsgInfo
}

func (d *DataBases) SaveUserChat(userId string, sendTime int64, m *pbMsg.MsgDataToDb) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(singleChat)
	seqUid := getSeqUid(userId, m.MsgData.Seq)
	filter := bson.M{"uid": seqUid}
	singleMsg := MsgInfo{}
	singleMsg.SendTime = sendTime
	if singleMsg.Msg, err = json.Marshal(m.MsgData); err != nil {
		return err
	}
	err = c.FindOneAndUpdate(ctx, filter, bson.M{"$push": bson.M{"msg": singleMsg}}).Err()
	if err != nil {
		singleChat := UserChat{}
		singleChat.UserId = seqUid
		singleChat.Msg = append(singleChat.Msg, singleMsg)
		if _, err = c.InsertOne(ctx, &singleChat); err != nil {
			return err
		}
		return err
	}
	return nil
}

func getSeqUid(userId string, seq uint32) string {
	seqSuffix := seq / singleGocMsgNum
	return indexGen(userId, seqSuffix)
}

func indexGen(userId string, seqSuffix uint32) string {
	return userId + ":" + strconv.FormatInt(int64(seqSuffix), 10)
}
