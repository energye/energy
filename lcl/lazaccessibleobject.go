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

// ILazAccessibleObject Parent: IObject
type ILazAccessibleObject interface {
	IObject
	AccessibleName() string                                                          // property
	SetAccessibleName(AValue string)                                                 // property
	AccessibleDescription() string                                                   // property
	SetAccessibleDescription(AValue string)                                          // property
	AccessibleValue() string                                                         // property
	SetAccessibleValue(AValue string)                                                // property
	AccessibleRole() TLazAccessibilityRole                                           // property
	SetAccessibleRole(AValue TLazAccessibilityRole)                                  // property
	Position() (resultPoint TPoint)                                                  // property
	SetPosition(AValue *TPoint)                                                      // property
	Size() (resultSize TSize)                                                        // property
	SetSize(AValue *TSize)                                                           // property
	Handle() uint32                                                                  // property
	SetHandle(AValue uint32)                                                         // property
	HandleAllocated() bool                                                           // function
	FindOwnerWinControl() IWinControl                                                // function
	AddChildAccessibleObject(ADataObject IObject) ILazAccessibleObject               // function
	GetChildAccessibleObjectWithDataObject(ADataObject IObject) ILazAccessibleObject // function
	GetChildAccessibleObjectsCount() int32                                           // function
	GetChildAccessibleObject(AIndex int32) ILazAccessibleObject                      // function
	GetFirstChildAccessibleObject() ILazAccessibleObject                             // function
	GetNextChildAccessibleObject() ILazAccessibleObject                              // function
	GetSelectedChildAccessibleObject() ILazAccessibleObject                          // function
	GetChildAccessibleObjectAtPos(APos *TPoint) ILazAccessibleObject                 // function
	GetEnumerator() ILazAccessibleObjectEnumerator                                   // function
	InitializeHandle()                                                               // procedure
	SetAccessibleName1(AName string)                                                 // procedure
	SetAccessibleDescription1(ADescription string)                                   // procedure
	SetAccessibleValue1(AValue string)                                               // procedure
	SetAccessibleRole1(ARole TLazAccessibilityRole)                                  // procedure
	InsertChildAccessibleObject(AObject ILazAccessibleObject)                        // procedure
	ClearChildAccessibleObjects()                                                    // procedure
	RemoveChildAccessibleObject(AObject ILazAccessibleObject, AFreeObject bool)      // procedure
}

// TLazAccessibleObject Parent: TObject
type TLazAccessibleObject struct {
	TObject
}

func NewLazAccessibleObject(AOwner IControl) ILazAccessibleObject {
	r1 := LCL().SysCallN(3494, GetObjectUintptr(AOwner))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) AccessibleName() string {
	r1 := LCL().SysCallN(3488, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TLazAccessibleObject) SetAccessibleName(AValue string) {
	LCL().SysCallN(3488, 1, m.Instance(), PascalStr(AValue))
}

func (m *TLazAccessibleObject) AccessibleDescription() string {
	r1 := LCL().SysCallN(3487, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TLazAccessibleObject) SetAccessibleDescription(AValue string) {
	LCL().SysCallN(3487, 1, m.Instance(), PascalStr(AValue))
}

func (m *TLazAccessibleObject) AccessibleValue() string {
	r1 := LCL().SysCallN(3490, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TLazAccessibleObject) SetAccessibleValue(AValue string) {
	LCL().SysCallN(3490, 1, m.Instance(), PascalStr(AValue))
}

func (m *TLazAccessibleObject) AccessibleRole() TLazAccessibilityRole {
	r1 := LCL().SysCallN(3489, 0, m.Instance(), 0)
	return TLazAccessibilityRole(r1)
}

func (m *TLazAccessibleObject) SetAccessibleRole(AValue TLazAccessibilityRole) {
	LCL().SysCallN(3489, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazAccessibleObject) Position() (resultPoint TPoint) {
	LCL().SysCallN(3508, 0, m.Instance(), uintptr(unsafePointer(&resultPoint)), uintptr(unsafePointer(&resultPoint)))
	return
}

func (m *TLazAccessibleObject) SetPosition(AValue *TPoint) {
	LCL().SysCallN(3508, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TLazAccessibleObject) Size() (resultSize TSize) {
	LCL().SysCallN(3514, 0, m.Instance(), uintptr(unsafePointer(&resultSize)), uintptr(unsafePointer(&resultSize)))
	return
}

func (m *TLazAccessibleObject) SetSize(AValue *TSize) {
	LCL().SysCallN(3514, 1, m.Instance(), uintptr(unsafePointer(AValue)), uintptr(unsafePointer(AValue)))
}

func (m *TLazAccessibleObject) Handle() uint32 {
	r1 := LCL().SysCallN(3504, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazAccessibleObject) SetHandle(AValue uint32) {
	LCL().SysCallN(3504, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazAccessibleObject) HandleAllocated() bool {
	r1 := LCL().SysCallN(3505, m.Instance())
	return GoBool(r1)
}

func (m *TLazAccessibleObject) FindOwnerWinControl() IWinControl {
	r1 := LCL().SysCallN(3495, m.Instance())
	return AsWinControl(r1)
}

func (m *TLazAccessibleObject) AddChildAccessibleObject(ADataObject IObject) ILazAccessibleObject {
	r1 := LCL().SysCallN(3491, m.Instance(), GetObjectUintptr(ADataObject))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectWithDataObject(ADataObject IObject) ILazAccessibleObject {
	r1 := LCL().SysCallN(3498, m.Instance(), GetObjectUintptr(ADataObject))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectsCount() int32 {
	r1 := LCL().SysCallN(3499, m.Instance())
	return int32(r1)
}

func (m *TLazAccessibleObject) GetChildAccessibleObject(AIndex int32) ILazAccessibleObject {
	r1 := LCL().SysCallN(3496, m.Instance(), uintptr(AIndex))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetFirstChildAccessibleObject() ILazAccessibleObject {
	r1 := LCL().SysCallN(3501, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetNextChildAccessibleObject() ILazAccessibleObject {
	r1 := LCL().SysCallN(3502, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetSelectedChildAccessibleObject() ILazAccessibleObject {
	r1 := LCL().SysCallN(3503, m.Instance())
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetChildAccessibleObjectAtPos(APos *TPoint) ILazAccessibleObject {
	r1 := LCL().SysCallN(3497, m.Instance(), uintptr(unsafePointer(APos)))
	return AsLazAccessibleObject(r1)
}

func (m *TLazAccessibleObject) GetEnumerator() ILazAccessibleObjectEnumerator {
	r1 := LCL().SysCallN(3500, m.Instance())
	return AsLazAccessibleObjectEnumerator(r1)
}

func LazAccessibleObjectClass() TClass {
	ret := LCL().SysCallN(3492)
	return TClass(ret)
}

func (m *TLazAccessibleObject) InitializeHandle() {
	LCL().SysCallN(3506, m.Instance())
}

func (m *TLazAccessibleObject) SetAccessibleName1(AName string) {
	LCL().SysCallN(3511, m.Instance(), PascalStr(AName))
}

func (m *TLazAccessibleObject) SetAccessibleDescription1(ADescription string) {
	LCL().SysCallN(3510, m.Instance(), PascalStr(ADescription))
}

func (m *TLazAccessibleObject) SetAccessibleValue1(AValue string) {
	LCL().SysCallN(3513, m.Instance(), PascalStr(AValue))
}

func (m *TLazAccessibleObject) SetAccessibleRole1(ARole TLazAccessibilityRole) {
	LCL().SysCallN(3512, m.Instance(), uintptr(ARole))
}

func (m *TLazAccessibleObject) InsertChildAccessibleObject(AObject ILazAccessibleObject) {
	LCL().SysCallN(3507, m.Instance(), GetObjectUintptr(AObject))
}

func (m *TLazAccessibleObject) ClearChildAccessibleObjects() {
	LCL().SysCallN(3493, m.Instance())
}

func (m *TLazAccessibleObject) RemoveChildAccessibleObject(AObject ILazAccessibleObject, AFreeObject bool) {
	LCL().SysCallN(3509, m.Instance(), GetObjectUintptr(AObject), PascalBool(AFreeObject))
}
