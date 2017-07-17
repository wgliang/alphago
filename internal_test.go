package alphago

import "testing"

func Test_Goid(t *testing.T) {
	t.Logf("This is goroutine %v", Goid())
}
