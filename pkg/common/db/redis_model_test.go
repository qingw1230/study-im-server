package db

import "testing"

func TestCheckCode(t *testing.T) {
	DB.SetCheckCode("test", "1002")
	ans, err := DB.GetCheckCode("test")
	if ans != "1002" {
		t.Fatal(err.Error())
	}
}
