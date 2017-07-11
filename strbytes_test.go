package alphago

import (
	"testing"
)

var su = "hello,world!"
var bu = []byte(su)

func Test_Str2bytes(t *testing.T) {
	_ = Str2Bytes(su)
}

func Test_Bytes2str(t *testing.T) {
	_ = Bytes2Str(bu)
}
