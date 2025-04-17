//go:build darwin

package darwin

import (
	"encoding/json"
	"fmt"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/mattn/go-pointer"
	"unsafe"
)

//#cgo CFLAGS: -x objective-c -std=c2x
//#cgo LDFLAGS: -framework Foundation -framework Cocoa
//#include "entry.h"
import "C"

func serializeConfig(config *barbuilder.Configuration) (*C.char, *handlers, error) {
	data, handlers, err := processConfig(config)
	if err != nil {
		return nil, nil, err
	}
	buffer, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}
	return C.CString(string(buffer)), handlers, nil
}

func transformError(err *C.char) error {
	return fmt.Errorf("darwin glue: %v", C.GoString(err))
}

func handleError(result C.ErrorResult) error {
	if result.err != nil {
		return transformError(result.err)
	}
	return nil
}

//export handleEvent
func handleEvent(raw unsafe.Pointer, event *C.char) {
	me := (pointer.Restore(raw)).(*touchBar)
	me.handleEvent(C.GoString(event))
}

func (me *touchBar) install(debug bool, configuration barbuilder.Configuration) error {
	if me.context != nil {
		return fmt.Errorf("touch bar already initialized")
	}
	mode := C.kMainWindow
	if debug {
		mode = C.kDebug
	}
	data, handlers, err := serializeConfig(&configuration)
	if err != nil {
		return err
	}
	defer C.free(unsafe.Pointer(data))
	me.handlers = handlers
	result := C.initTouchBar(C.AttachMode(mode), data, pointer.Save(me))
	if result.err != nil {
		return transformError(result.err)
	}
	me.context = result.result
	return nil
}

func (me *touchBar) Install(configuration barbuilder.Configuration) error {
	me.lock.Lock()
	defer me.lock.Unlock()
	return me.install(false, configuration)
}

func (me *touchBar) Debug(configuration barbuilder.Configuration) error {
	me.lock.Lock()
	err := me.install(true, configuration)
	me.lock.Unlock()
	if err != nil {
		return err
	}
	return handleError(C.runDebug(me.context))
}

func (me *touchBar) Update(configuration barbuilder.Configuration) error {
	me.lock.Lock()
	defer me.lock.Unlock()
	if me.context == nil {
		return fmt.Errorf("touch bar has not been initialized")
	}
	data, handlers, err := serializeConfig(&configuration)
	if err != nil {
		return err
	}
	defer C.free(unsafe.Pointer(data))
	me.handlers = handlers
	return handleError(C.updateTouchBar(me.context, data))
}

func (me *touchBar) Uninstall() error {
	me.lock.Lock()
	defer me.lock.Unlock()
	if me.context == nil {
		return fmt.Errorf("touch bar has not been initialized")
	}
	return handleError(C.destroyTouchBar(me.context))
}
