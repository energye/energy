//go:build darwin
// +build darwin

package cef

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa

*/
import "C"
import (
	"github.com/energye/golcl/lcl"
	"unsafe"
)

//export GoLog
func GoLog(message *C.char) {
	msg := C.GoString(message)
	println(msg)
}

//export CanDrag
func CanDrag(nsWindow unsafe.Pointer) bool {
	var window = currentWindow(lcl.NSWindow(nsWindow))
	if window != nil {
		return window.drag.canDrag
	}
	return false
}

//export SetCanDrag
func SetCanDrag(nsWindow unsafe.Pointer, value bool) {
	var window = currentWindow(lcl.NSWindow(nsWindow))
	if window != nil {
		window.drag.canDrag = value
	}
}

//export GetTitlebarHeight
func GetTitlebarHeight(nsWindow unsafe.Pointer) int32 {
	var window = currentWindow(lcl.NSWindow(nsWindow))
	if window != nil {
		// TODO no impl
	}
	return 0
}

//export CheckDraggableRegions
func CheckDraggableRegions(nsWindow unsafe.Pointer, mouseX, mouseY int32) {
	var window = currentWindow(lcl.NSWindow(nsWindow))
	if window != nil {
		window.drag.canDrag = false
		regions := window.ChromiumBrowser().Regions()
		if regions != nil {
			var isDrag bool
			for i := 0; i < regions.RegionsCount(); i++ {
				region := regions.Region(i)
				if region.Draggable {
					isDrag = PtInRegion(mouseX, mouseY, region.Bounds.X, region.Bounds.Y, region.Bounds.Width, region.Bounds.Height)
					if isDrag {
						break
					}
				}
			}
			if isDrag {
				for i := 0; i < regions.RegionsCount(); i++ {
					region := regions.Region(i)
					if !region.Draggable {
						tempIsDrag := PtInRegion(mouseX, mouseY, region.Bounds.X, region.Bounds.Y, region.Bounds.Width, region.Bounds.Height)
						if tempIsDrag {
							isDrag = false
							break
						}
					}
				}
			}
			window.drag.canDrag = isDrag
		}
	}
}

func currentWindow(windowHandle lcl.NSWindow) *LCLBrowserWindow {
	var window *LCLBrowserWindow
	for _, win := range BrowserWindow.GetWindowInfos() {
		if win.IsClosing() {
			break
		}
		windowPtr := win.AsLCLBrowserWindow().BrowserWindow().TForm.PlatformWindow()
		if windowPtr == windowHandle {
			window = win.AsLCLBrowserWindow().BrowserWindow()
			break
		}
	}
	return window
}
