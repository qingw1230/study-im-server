package logic

import (
	"context"

	"github.com/qingw1230/study-im-server/pkg/common/log"
	pbPush "github.com/qingw1230/study-im-server/pkg/proto/push"
	pbReply "github.com/qingw1230/study-im-server/pkg/proto/reply"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func MsgToUser(req *pbPush.PushMsgReq) {
	conn, err := grpc.NewClient("127.0.0.1:10400", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("NewClient failed", err.Error())
		return
	}
	client := pbReply.NewOnlineMessageRelayServiceClient(conn)
	_, err = client.OnlinePushMsg(context.Background(), &pbReply.OnlinePushMsgReq{MsgData: req.MsgData, PushToUserId: req.PushToUserId})
	if err != nil {
		log.Error("push data to client rpc failed", err.Error())
		return
	}
}
