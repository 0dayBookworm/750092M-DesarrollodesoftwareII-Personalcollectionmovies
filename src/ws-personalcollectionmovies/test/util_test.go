package test

import (
    "fmt"
    "testing"
	"ws-personalcollectionmovies/util"
    )

func TestEncryptMD5(t *testing.T) {
    fmt.Println("EncryptMD5 [TestSuit]")
    expected := "dd4b21e9ef71e1291183a46b913ae6f2"
    result := util.EncryptMD5("00000000")
    if result != expected {
        t.Error("Expected "+expected+", got ", result)
    }
}