package alphago

import (
	"testing"
)

//  +-------------------------------------------------------------------------+
//  |  Operation  |    Cycles    | Time(ns)/op | Memory(B)/op |  Allocate/op  |
//  +=========================================================================+
//  |    Goid     |    300000    |     5374    |       192    |       3       |
//  +-------------------------------------------------------------------------+
func Benchmark_Goid_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Goid()
	}
}
