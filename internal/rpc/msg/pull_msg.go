package msg

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
)

func (s *msgServer) PullMessageBySeqList(_ context.Context, req *pbMsg.PullMessageBySeqListReq) (*pbMsg.PullMessageBySeqListResp, error) {
	log.Info("call rpc PullMessageBySeqList args:", req.String())
	if !token_verify.CheckAccess(req.OpUserId, req.UserId) {
		log.Error("CheckAccess failed", req.OpUserId, req.UserId)
		return &pbMsg.PullMessageBySeqListResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	resp := &pbMsg.PullMessageBySeqListResp{CommonResp: &constant.PBCommonSuccessResp}
	msgList, err := db.DB.GetMsgBySeqList(req.UserId, req.SeqList)
	if err != nil {
		log.Error("mongo GetMsgBySeqList failed", err.Error(), req.UserId, req.SeqList)
		copier.Copy(&resp.CommonResp, &constant.MongoCommonFailResp)
		return resp, nil
	}

	resp.List = msgList
	return resp, nil
}
