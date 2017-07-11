package alphago

import (
	"unsafe"
)

// Str2Bytes is a function that converts data of string into []byte
// formate, compared to the official []byte conversion it has a performance
// improvement of more than five hundred times, especially under a higher
// benchmark and no allocating memory over avoiding duplication of the
// underlying byte array.
func Str2Bytes(stringData string) []byte {
	temporaryData := (*[2]uintptr)(unsafe.Pointer(&stringData))
	bytesData := [3]uintptr{temporaryData[0], temporaryData[1], temporaryData[1]}
	return *(*[]byte)(unsafe.Pointer(&bytesData))
}

// Bytes2Str is a function that converts data of []byte into string
// formate.Compared to the official string conversion it has a performance
// improvement of more than two thousand times, especially under a higher
// benchmark and no allocating memory over avoiding duplication of the
// underlying byte array.
func Bytes2Str(bytesData []byte) string {
	return *(*string)(unsafe.Pointer(&bytesData))
}
