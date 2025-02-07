package logic

import (
	"context"

	"github.com/qingw1230/study-im-server/pkg/common/config"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/etcdv3"
	pbPush "github.com/qingw1230/study-im-server/pkg/proto/push"
	pbReply "github.com/qingw1230/study-im-server/pkg/proto/reply"
)

func MsgToUser(req *pbPush.PushMsgReq) {
	conn := etcdv3.GetConn(config.Config.Etcd.EtcdSchema, config.Config.Etcd.EtcdAddr, config.Config.RpcRegisterName.OnlineMessageRelayName)
	client := pbReply.NewOnlineMessageRelayServiceClient(conn)
	_, err := client.OnlinePushMsg(context.Background(), &pbReply.OnlinePushMsgReq{MsgData: req.MsgData, PushToUserId: req.PushToUserId})
	if err != nil {
		log.Error("push data to client rpc failed", err.Error())
		return
	}
}
