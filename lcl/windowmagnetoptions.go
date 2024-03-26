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

// IWindowMagnetOptions Parent: IPersistent
type IWindowMagnetOptions interface {
	IPersistent
	SnapToMonitor() bool           // property
	SetSnapToMonitor(AValue bool)  // property
	SnapToForms() bool             // property
	SetSnapToForms(AValue bool)    // property
	SnapFormTarget() bool          // property
	SetSnapFormTarget(AValue bool) // property
	Distance() int32               // property
	SetDistance(AValue int32)      // property
	AssignTo(Dest IPersistent)     // procedure
}

// TWindowMagnetOptions Parent: TPersistent
type TWindowMagnetOptions struct {
	TPersistent
}

func NewWindowMagnetOptions() IWindowMagnetOptions {
	r1 := LCL().SysCallN(5290)
	return AsWindowMagnetOptions(r1)
}

func (m *TWindowMagnetOptions) SnapToMonitor() bool {
	r1 := LCL().SysCallN(5294, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWindowMagnetOptions) SetSnapToMonitor(AValue bool) {
	LCL().SysCallN(5294, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWindowMagnetOptions) SnapToForms() bool {
	r1 := LCL().SysCallN(5293, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWindowMagnetOptions) SetSnapToForms(AValue bool) {
	LCL().SysCallN(5293, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWindowMagnetOptions) SnapFormTarget() bool {
	r1 := LCL().SysCallN(5292, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWindowMagnetOptions) SetSnapFormTarget(AValue bool) {
	LCL().SysCallN(5292, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWindowMagnetOptions) Distance() int32 {
	r1 := LCL().SysCallN(5291, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TWindowMagnetOptions) SetDistance(AValue int32) {
	LCL().SysCallN(5291, 1, m.Instance(), uintptr(AValue))
}

func WindowMagnetOptionsClass() TClass {
	ret := LCL().SysCallN(5289)
	return TClass(ret)
}

func (m *TWindowMagnetOptions) AssignTo(Dest IPersistent) {
	LCL().SysCallN(5288, m.Instance(), GetObjectUintptr(Dest))
}
