package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomId(length int) string {
	return "U" + generateRandomNum(length)
}

func generateRandomNum(length int) string {
	const digits = "0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := make([]byte, length)
	for i := range id {
		id[i] = digits[r.Intn(len(digits))]
	}
	return string(id)
}

func GenerateRandomStr(length int) string {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	str := make([]byte, length)
	for i := range str {
		str[i] = letters[r.Intn(len(letters))]
	}
	return string(str)
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
