package alphago

import (
	"fmt"
	"reflect"
	"testing"
)

var su = "hello,world!"
var bu = []byte(su)

func Test_Str2bytes(t *testing.T) {
	if !reflect.DeepEqual(Str2Bytes(su), []uint8{104, 101, 108, 108, 111, 44, 119, 111, 114, 108, 100, 33}) {
		t.Error(`Str2Bytes("hello,world!") != [104 101 108 108 111 44 119 111 114 108 100 33]`)
	}

}

func Test_Bytes2str(t *testing.T) {
	if !reflect.DeepEqual(Bytes2Str(bu), "hello,world!") {
		t.Error(`Bytes2Str([]uint8{104, 101, 108, 108, 111, 44, 119, 111, 114, 108, 100, 33}) != hello,world!`)
	}
}
