package db

import (
	"context"
	"strconv"
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
	"github.com/qingw1230/study-im-server/pkg/proto/sdkws"
	"github.com/qingw1230/study-im-server/pkg/utils"
	"google.golang.org/protobuf/proto"
	"gopkg.in/mgo.v2/bson"
)

const (
	singleChat      = "msg"
	groupChat       = "group"
	singleGocMsgNum = 5000
)

type UserChat struct {
	Uid      string
	Seq      uint64
	SendTime int64
	Msg      []byte // 使用 proto 序列化后的 *pbPublic.MsgData
}

func (d *DataBases) GetMsgBySeqList(userId string, seqList []uint64) (seqMsg []*sdkws.MsgData, err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(singleChat)

	table := func(userId string, seqList []uint64) map[string][]uint64 {
		m := make(map[string][]uint64)
		for _, seq := range seqList {
			seqUid := getSeqUid(userId, seq)
			m[seqUid] = append(m[seqUid], seq)
		}
		return m
	}(userId, seqList)

	var hasSeqList []uint64
	userChat := UserChat{}
	for seqUid, groupSeqList := range table {
		for _, seq := range groupSeqList {
			if err = c.FindOne(ctx, bson.M{"uid": seqUid, "seq": seq}).Decode(&userChat); err != nil {
				log.Error("not find seqUid", seqUid, err.Error())
				continue
			}

			msg := &sdkws.MsgData{}
			if err = proto.Unmarshal(userChat.Msg, msg); err != nil {
				log.Error("mongo get Unmarshal failed", err.Error(), seqUid, seq)
				return
			}
			seqMsg = append(seqMsg, msg)
			hasSeqList = append(hasSeqList, seq)
		}

	}

	if len(hasSeqList) != len(seqList) {
		diff := utils.Difference(hasSeqList, seqList)
		exceptionMSg := genExceptionMessageBySeqList(diff)
		seqMsg = append(seqMsg, exceptionMSg...)
	}
	return seqMsg, nil
}

func genExceptionMessageBySeqList(seqList []uint64) (exceptionMsg []*sdkws.MsgData) {
	for _, seq := range seqList {
		msg := &sdkws.MsgData{}
		msg.Seq = seq
		exceptionMsg = append(exceptionMsg, msg)
	}
	return exceptionMsg
}

func (d *DataBases) SaveUserChat(userId string, sendTime int64, m *pbMsg.MsgDataToDb) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Config.Mongo.DBTimeout)*time.Second)
	c := d.mongoClient.Database(config.Config.Mongo.DBDatabase).Collection(singleChat)

	seqUid := getSeqUid(userId, m.MsgData.Seq)
	singleChatLog := UserChat{
		Uid:      seqUid,
		Seq:      m.MsgData.Seq,
		SendTime: sendTime,
	}
	if singleChatLog.Msg, err = proto.Marshal(m.MsgData); err != nil {
		return err
	}
	if _, err = c.InsertOne(ctx, &singleChatLog); err != nil {
		return err
	}
	return nil
}

func getSeqUid(userId string, seq uint64) string {
	seqSuffix := seq / singleGocMsgNum
	return indexGen(userId, seqSuffix)
}

func indexGen(userId string, seqSuffix uint64) string {
	return userId + ":" + strconv.FormatInt(int64(seqSuffix), 10)
}
