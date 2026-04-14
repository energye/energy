//go:build !cgo

package webkit2

import (
	. "github.com/energye/energy/v3/platform/darwin/types"
	"unsafe"
)

func AsWkWebView(ptr unsafe.Pointer) IWkWebView {
	return nil
}
