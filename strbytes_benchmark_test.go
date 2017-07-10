package alphago

import (
	"strings"
	"testing"
)

var s = strings.Repeat("hello", 1000)
var b = []byte(s)

func official_bytes() {
	_ = []byte(s)
}

func official_string() {
	_ = string(b)
}

func alphago_str2bytes() {
	_ = str2bytes(s)
}

func alphago_bytes2str() {
	_ = bytes2str(b)
}

//  +-------------------------------------------------------------------------+
//  |  Operation  |    Cycles    | Time(ns)/op | Memory(B)/op |  Allocate/op  |
//  +=========================================================================+
//  |   []byte    |  2000000     |     885     |     5376     |      1        |
//  +-------------------------------------------------------------------------+
//  |  str2bytes  | 1000000000   |     2.61    |      0       |      0        |
//  +-------------------------------------------------------------------------+
func Benchmark_official_bytes_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		official_bytes()
	}
}

func Benchmark_alphago_str2bytes_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alphago_str2bytes()
	}
}

//  +-------------------------------------------------------------------------+
//  |  Operation  |    Cycles    | Time(ns)/op | Memory(B)/op |  Allocate/op  |
//  +=========================================================================+
//  |   []byte    |  1000000     |     1041    |     5376     |      1        |
//  +-------------------------------------------------------------------------+
//  |  str2bytes  | 2000000000   |     1.02    |      0       |      0        |
//  +-------------------------------------------------------------------------+
func Benchmark_official_string_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		official_string()
	}
}

func Benchmark_alphago_bytes2str_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		alphago_bytes2str()
	}
}
