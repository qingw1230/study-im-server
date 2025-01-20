package utils

import (
	"strconv"
	"time"

	"github.com/qingw1230/study-im-server/pkg/common/constant"
)

// GenerateRandomStrWithLength 使用 timestamp 和 md5 随机生成，最大长度为 32
func GenerateRandomStrWithLength(length int) string {
	if length > 32 {
		return ""
	}
	return Md5Encode(strconv.FormatInt(time.Now().UnixNano(), 10))[:length]
}

func GenerateUserId() string {
	return "U" + Md5Encode(strconv.FormatInt(time.Now().UnixNano(), 10))[:constant.LENGTH_11]
}

func GenerateGroupId() string {
	return "G" + Md5Encode(strconv.FormatInt(time.Now().UnixNano(), 10))[:constant.LENGTH_11]
}
