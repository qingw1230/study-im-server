package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tmpStr := h.Sum(nil)
	return hex.EncodeToString(tmpStr)
}

func MakePassword(pwd string) string {
	return Md5Encode(pwd)
}

// ValidPassword pwd 待加密的原始密码
func ValidPassword(pwd string, password string) bool {
	return Md5Encode(pwd) == password
}
