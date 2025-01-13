package utils

import (
	"math/rand"
	"time"
)

var (
	r *rand.Rand
)

func init() {
	r.Seed(time.Now().Unix())
}

func GenerateRandomID(length int) string {
	const digits = "0123456789"
	id := make([]byte, length)
	for i := range id {
		id[i] = digits[r.Intn(len(digits))]
	}
	return "U" + string(id)
}
