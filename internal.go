package alphago

import (
	"runtime"
	"strconv"
	"strings"
)

// Goid is a function that gives an access to the goroutine id, will
// return the id of the current goroutine.If you call this function
// you will go straight to hell.
func Goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, _ := strconv.Atoi(idField)

	return id
}
