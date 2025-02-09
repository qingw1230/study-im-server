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

// IsContainString 检查 list 中是否包含 target
func IsContainString(target string, list []string) bool {
	for _, l := range list {
		if target == l {
			return true
		}
	}
	return false
}

func IsCotainUInt32(target uint32, list []uint32) bool {
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

func Intersect(slice1, slice2 []uint64) []uint64 {
	m := make(map[uint64]bool)
	n := make([]uint64, 0)
	for _, v := range slice1 {
		m[v] = true
	}
	for _, v := range slice2 {
		flag, _ := m[v]
		if flag {
			n = append(n, v)
		}
	}
	return n
}

func Difference(slice1, slice2 []uint64) []uint64 {
	m := make(map[uint64]bool)
	n := make([]uint64, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v] = true
	}
	for _, v := range slice1 {
		if !m[v] {
			n = append(n, v)
		}
	}

	for _, v := range slice2 {
		if !m[v] {
			n = append(n, v)
		}
	}
	return n
}
