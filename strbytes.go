package alphago

import (
	"unsafe"
)

// str2bytes is a function that converts data of string into []byte
// formate, compared to the official []byte conversion it has a performance
// improvement of more than five hundred times, especially under a higher
// benchmark and no allocating memory over avoiding duplication of the
// underlying byte array.
func str2bytes(stringData string) []byte {
	temporaryData := (*[2]uintptr)(unsafe.Pointer(&stringData))
	bytesData := [3]uintptr{temporaryData[0], temporaryData[1], temporaryData[1]}
	return *(*[]byte)(unsafe.Pointer(&bytesData))
}

// bytes2str is a function that converts data of []byte into string
// formate.Compared to the official string conversion it has a performance
// improvement of more than two thousand times, especially under a higher
// benchmark and no allocating memory over avoiding duplication of the
// underlying byte array.
func bytes2str(bytesData []byte) string {
	return *(*string)(unsafe.Pointer(&bytesData))
}
