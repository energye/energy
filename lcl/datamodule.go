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

// IDataModule Parent: IComponent
type IDataModule interface {
	IComponent
	DesignOffset() (resultPoint TPoint) // property
	SetDesignOffset(AValue *TPoint)     // property
	DesignSize() (resultPoint TPoint)   // property
	SetDesignSize(AValue *TPoint)       // property
	DesignPPI() int32                   // property
	SetDesignPPI(AValue int32)          // property
	OldCreateOrder() bool               // property
	SetOldCreateOrder(AValue bool)      // property
	SetOnCreate(fn TNotifyEvent)        // property event
	SetOnDestroy(fn TNotifyEvent)       // property event
}

// TDataModule Parent: TComponent
type TDataModule struct {
	TComponent
	createPtr  uintptr
	destroyPtr uintptr
}

func NewDataModule(AOwner IComponent) IDataModule {
	r1 := LCL().SysCallN(2504, GetObjectUintptr(AOwner))
	return AsDataModule(r1)
}

func NewDataModuleNew(AOwner IComponent) IDataModule {
	r1 := LCL().SysCallN(2505, GetObjectUintptr(AOwner))
	return AsDataModule(r1)
}

func NewDataModuleNew1(AOwner IComponent, CreateMode int32) IDataModule {
	r1 := LCL().SysCallN(2506, GetObjectUintptr(AOwner), uintptr(CreateMode))
	return AsDataModule(r1)
}

func (m *TDataModule) DesignOffset() (resultPoint TPoint) {
	LCL().SysCallN(2507, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDataModule) SetDesignOffset(AValue *TPoint) {
	LCL().SysCallN(2507, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDataModule) DesignSize() (resultPoint TPoint) {
	LCL().SysCallN(2509, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TDataModule) SetDesignSize(AValue *TPoint) {
	LCL().SysCallN(2509, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TDataModule) DesignPPI() int32 {
	r1 := LCL().SysCallN(2508, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDataModule) SetDesignPPI(AValue int32) {
	LCL().SysCallN(2508, 1, m.Instance(), uintptr(AValue))
}

func (m *TDataModule) OldCreateOrder() bool {
	r1 := LCL().SysCallN(2510, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDataModule) SetOldCreateOrder(AValue bool) {
	LCL().SysCallN(2510, 1, m.Instance(), PascalBool(AValue))
}

func DataModuleClass() TClass {
	ret := LCL().SysCallN(2503)
	return TClass(ret)
}

func (m *TDataModule) SetOnCreate(fn TNotifyEvent) {
	if m.createPtr != 0 {
		RemoveEventElement(m.createPtr)
	}
	m.createPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2511, m.Instance(), m.createPtr)
}

func (m *TDataModule) SetOnDestroy(fn TNotifyEvent) {
	if m.destroyPtr != 0 {
		RemoveEventElement(m.destroyPtr)
	}
	m.destroyPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2512, m.Instance(), m.destroyPtr)
}
