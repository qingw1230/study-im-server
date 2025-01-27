package utils

import (
	"strconv"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
)

func IntToString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func StringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

// IsContain 检查 list 中是否包含 target
func IsContain(target string, list []string) bool {
	for _, l := range list {
		if target == l {
			return true
		}
	}
	return false
}

// GetConversationIdBySessionType 类型前缀 + sourceId
func GetConversationIdBySessionType(sourceId string, sessionType int) string {
	switch sessionType {
	case constant.SingleChatType:
		return constant.STR_SINGLE_ + sourceId
	case constant.GroupChatType:
		return constant.STR_GROUP_ + sourceId
	}
	return ""
}
