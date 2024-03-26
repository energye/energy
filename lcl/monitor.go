//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
	"unsafe"
)

// IMonitor Parent: IObject
type IMonitor interface {
	IObject
	Handle() HMONITOR                 // property
	MonitorNum() int32                // property
	Left() int32                      // property
	Height() int32                    // property
	Top() int32                       // property
	Width() int32                     // property
	BoundsRect() (resultRect TRect)   // property
	WorkareaRect() (resultRect TRect) // property
	Primary() bool                    // property
	PixelsPerInch() int32             // property
}

// TMonitor Parent: TObject
type TMonitor struct {
	TObject
}

func NewMonitor() IMonitor {
	r1 := LCL().SysCallN(3698)
	return AsMonitor(r1)
}

func (m *TMonitor) Handle() HMONITOR {
	r1 := LCL().SysCallN(3699, m.Instance())
	return HMONITOR(r1)
}

func (m *TMonitor) MonitorNum() int32 {
	r1 := LCL().SysCallN(3702, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Left() int32 {
	r1 := LCL().SysCallN(3701, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Height() int32 {
	r1 := LCL().SysCallN(3700, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Top() int32 {
	r1 := LCL().SysCallN(3705, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Width() int32 {
	r1 := LCL().SysCallN(3706, m.Instance())
	return int32(r1)
}

func (m *TMonitor) BoundsRect() (resultRect TRect) {
	LCL().SysCallN(3696, m.Instance(), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TMonitor) WorkareaRect() (resultRect TRect) {
	LCL().SysCallN(3707, m.Instance(), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func (m *TMonitor) Primary() bool {
	r1 := LCL().SysCallN(3704, m.Instance())
	return GoBool(r1)
}

func (m *TMonitor) PixelsPerInch() int32 {
	r1 := LCL().SysCallN(3703, m.Instance())
	return int32(r1)
}

func MonitorClass() TClass {
	ret := LCL().SysCallN(3697)
	return TClass(ret)
}
