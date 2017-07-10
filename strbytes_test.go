package alphago

import (
	"testing"
)

var su = "hello,world!"
var bu = []byte(su)

func Test_str2bytes_1(t *testing.T) {
	_ = str2bytes(su)
}

func Test_bytes2str_1(t *testing.T) {
	_ = bytes2str(bu)
}
