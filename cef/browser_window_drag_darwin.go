//go:build darwin
// +build darwin

package cef

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa

*/
import "C"
import "unsafe"

//export GoLog
func GoLog(message *C.char) {
	msg := C.GoString(message)
	println(msg)
}

//export ShouldDrag
func ShouldDrag(window unsafe.Pointer) bool {
	println("ShouldDrag", uintptr(window))
	return false
}

//export SetShouldDrag
func SetShouldDrag(window unsafe.Pointer, value bool) {
	println("SetShouldDrag", uintptr(window), "value", value)
}
