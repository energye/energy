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
	r1 := LCL().SysCallN(2287, GetObjectUintptr(AOwner))
	return AsDataModule(r1)
}

func NewDataModuleNew(AOwner IComponent) IDataModule {
	r1 := LCL().SysCallN(2288, GetObjectUintptr(AOwner))
	return AsDataModule(r1)
}

func NewDataModuleNew1(AOwner IComponent, CreateMode int32) IDataModule {
	r1 := LCL().SysCallN(2289, GetObjectUintptr(AOwner), uintptr(CreateMode))
	return AsDataModule(r1)
}

func (m *TDataModule) DesignOffset() (resultPoint TPoint) {
	LCL().SysCallN(2290, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TDataModule) SetDesignOffset(AValue *TPoint) {
	LCL().SysCallN(2290, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TDataModule) DesignSize() (resultPoint TPoint) {
	LCL().SysCallN(2292, 0, m.Instance(), uintptr(unsafe.Pointer(&resultPoint)), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TDataModule) SetDesignSize(AValue *TPoint) {
	LCL().SysCallN(2292, 1, m.Instance(), uintptr(unsafe.Pointer(AValue)), uintptr(unsafe.Pointer(AValue)))
}

func (m *TDataModule) DesignPPI() int32 {
	r1 := LCL().SysCallN(2291, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TDataModule) SetDesignPPI(AValue int32) {
	LCL().SysCallN(2291, 1, m.Instance(), uintptr(AValue))
}

func (m *TDataModule) OldCreateOrder() bool {
	r1 := LCL().SysCallN(2293, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TDataModule) SetOldCreateOrder(AValue bool) {
	LCL().SysCallN(2293, 1, m.Instance(), PascalBool(AValue))
}

func DataModuleClass() TClass {
	ret := LCL().SysCallN(2286)
	return TClass(ret)
}

func (m *TDataModule) SetOnCreate(fn TNotifyEvent) {
	if m.createPtr != 0 {
		RemoveEventElement(m.createPtr)
	}
	m.createPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2294, m.Instance(), m.createPtr)
}

func (m *TDataModule) SetOnDestroy(fn TNotifyEvent) {
	if m.destroyPtr != 0 {
		RemoveEventElement(m.destroyPtr)
	}
	m.destroyPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(2295, m.Instance(), m.destroyPtr)
}
