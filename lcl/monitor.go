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
	r1 := LCL().SysCallN(4340)
	return AsMonitor(r1)
}

func (m *TMonitor) Handle() HMONITOR {
	r1 := LCL().SysCallN(4341, m.Instance())
	return HMONITOR(r1)
}

func (m *TMonitor) MonitorNum() int32 {
	r1 := LCL().SysCallN(4344, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Left() int32 {
	r1 := LCL().SysCallN(4343, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Height() int32 {
	r1 := LCL().SysCallN(4342, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Top() int32 {
	r1 := LCL().SysCallN(4347, m.Instance())
	return int32(r1)
}

func (m *TMonitor) Width() int32 {
	r1 := LCL().SysCallN(4348, m.Instance())
	return int32(r1)
}

func (m *TMonitor) BoundsRect() (resultRect TRect) {
	LCL().SysCallN(4338, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TMonitor) WorkareaRect() (resultRect TRect) {
	LCL().SysCallN(4349, m.Instance(), uintptr(unsafePointer(&resultRect)))
	return
}

func (m *TMonitor) Primary() bool {
	r1 := LCL().SysCallN(4346, m.Instance())
	return GoBool(r1)
}

func (m *TMonitor) PixelsPerInch() int32 {
	r1 := LCL().SysCallN(4345, m.Instance())
	return int32(r1)
}

func MonitorClass() TClass {
	ret := LCL().SysCallN(4339)
	return TClass(ret)
}
