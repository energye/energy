package cef

import (
	. "github.com/energye/energy/common"
	"github.com/energye/energy/types"
	"unsafe"
)

func (m *ICefDisplay) ID() (result int64) {
	Proc(internale_CEFDisplay_ID).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) DeviceScaleFactor() (result types.Single) {
	Proc(internale_CEFDisplay_DeviceScaleFactor).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) Rotation() int32 {
	r1, _, _ := Proc(internale_CEFDisplay_Rotation).Call(uintptr(m.instance))
	return int32(r1)
}

func (m *ICefDisplay) Bounds() (result TCefRect) {
	Proc(internale_CEFDisplay_Bounds).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefDisplay) WorkArea() (result TCefRect) {
	Proc(internale_CEFDisplay_WorkArea).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}
