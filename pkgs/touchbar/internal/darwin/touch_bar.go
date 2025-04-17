//go:build darwin

package darwin

import (
	"sync"
	"unsafe"

	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
)

type touchBar struct {
	options  barbuilder.Options
	handlers *handlers
	context  unsafe.Pointer
	lock     sync.Mutex
}

func NewTouchBar(options barbuilder.Options) barbuilder.TouchBar {
	return &touchBar{
		options: options,
	}
}
