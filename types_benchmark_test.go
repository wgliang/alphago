package alphago

import (
	"strings"
	"testing"
)

var stringBench = strings.Repeat("hello", 1000)
var bytesBench = []byte(stringBench)

//  +-------------------------------------------------------------------------+
//  |  Operation  |    Cycles    | Time(ns)/op | Memory(B)/op |  Allocate/op  |
//  +=========================================================================+
//  |   []byte    |  2000000     |     885     |     5376     |      1        |
//  +-------------------------------------------------------------------------+
//  |  Str2Bytes  | 1000000000   |     2.61    |      0       |      0        |
//  +-------------------------------------------------------------------------+
func Benchmark_official_bytes_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(stringBench)
	}
}

func Benchmark_alphago_Str2Bytes_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Str2Bytes(stringBench)
	}
}

//  +-------------------------------------------------------------------------+
//  |  Operation  |    Cycles    | Time(ns)/op | Memory(B)/op |  Allocate/op  |
//  +=========================================================================+
//  |   []byte    |  1000000     |     1041    |     5376     |      1        |
//  +-------------------------------------------------------------------------+
//  |  Bytes2Str  | 2000000000   |     1.02    |      0       |      0        |
//  +-------------------------------------------------------------------------+
func Benchmark_official_string_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(bytesBench)
	}
}

func Benchmark_alphago_Bytes2Str_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Bytes2Str(bytesBench)
	}
}
