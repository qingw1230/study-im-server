package msg

import (
	"context"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
	"github.com/qingw1230/study-im-server/pkg/common/db"
	"github.com/qingw1230/study-im-server/pkg/common/log"
	"github.com/qingw1230/study-im-server/pkg/common/token_verify"
	pbMsg "github.com/qingw1230/study-im-server/pkg/proto/msg"
)

func (s *msgServer) GetNewestSeq(_ context.Context, req *pbMsg.GetNewestSeqReq) (*pbMsg.GetNewestSeqResp, error) {
	log.Info("call rpc GetNewestSeq args:", req.String())
	if !token_verify.CheckAccess(req.OpUserId, req.UserId) {
		log.Error("CheckAccess failed", req.OpUserId, req.UserId)
		return &pbMsg.GetNewestSeqResp{CommonResp: &constant.PBTokenAccessErrorResp}, nil
	}

	newestSeq, err := db.DB.GetMaxSeq(req.UserId)
	if err != nil {
		return &pbMsg.GetNewestSeqResp{CommonResp: &constant.PBRedisCommonFailResp}, nil
	}
	resp := &pbMsg.GetNewestSeqResp{CommonResp: &constant.PBCommonSuccessResp}
	resp.NewestSeq = newestSeq
	return resp, nil
}
