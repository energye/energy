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

// IPersistent Parent: IObject
type IPersistent interface {
	IObject
	GetNamePath() string                                                               // function
	Assign(Source IPersistent)                                                         // procedure
	FPOAttachObserver(AObserver IObject)                                               // procedure
	FPODetachObserver(AObserver IObject)                                               // procedure
	FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) // procedure
}

// TPersistent Parent: TObject
type TPersistent struct {
	TObject
}

func NewPersistent() IPersistent {
	r1 := LCL().SysCallN(3880)
	return AsPersistent(r1)
}

func (m *TPersistent) GetNamePath() string {
	r1 := LCL().SysCallN(3884, m.Instance())
	return GoStr(r1)
}

func PersistentClass() TClass {
	ret := LCL().SysCallN(3879)
	return TClass(ret)
}

func (m *TPersistent) Assign(Source IPersistent) {
	LCL().SysCallN(3878, m.Instance(), GetObjectUintptr(Source))
}

func (m *TPersistent) FPOAttachObserver(AObserver IObject) {
	LCL().SysCallN(3881, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TPersistent) FPODetachObserver(AObserver IObject) {
	LCL().SysCallN(3882, m.Instance(), GetObjectUintptr(AObserver))
}

func (m *TPersistent) FPONotifyObservers(ASender IObject, AOperation TFPObservedOperation, Data uintptr) {
	LCL().SysCallN(3883, m.Instance(), GetObjectUintptr(ASender), uintptr(AOperation), uintptr(Data))
}
