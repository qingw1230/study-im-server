package utils

import (
	"fmt"
	"testing"
)

func Test_GenerateUserId(t *testing.T) {
	fmt.Println(GenerateUserId())
}

func Test_GenerateGroupId(t *testing.T) {
	fmt.Println(GenerateGroupId())
}
