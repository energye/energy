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

// ICustomActionList Parent: ILCLComponent
type ICustomActionList interface {
	ILCLComponent
	Actions(Index int32) IContainedAction                        // property
	SetActions(Index int32, AValue IContainedAction)             // property
	ActionCount() int32                                          // property
	Images() ICustomImageList                                    // property
	SetImages(AValue ICustomImageList)                           // property
	State() TActionListState                                     // property
	SetState(AValue TActionListState)                            // property
	ActionByName(ActionName string) IContainedAction             // function
	GetEnumeratorForActionListEnumerator() IActionListEnumerator // function
	IndexOfName(ActionName string) int32                         // function
	IsShortCut(Message *TLMKey) bool                             // function
}

// TCustomActionList Parent: TLCLComponent
type TCustomActionList struct {
	TLCLComponent
}

func NewCustomActionList(AOwner IComponent) ICustomActionList {
	r1 := LCL().SysCallN(1057, GetObjectUintptr(AOwner))
	return AsCustomActionList(r1)
}

func (m *TCustomActionList) Actions(Index int32) IContainedAction {
	r1 := LCL().SysCallN(1055, 0, m.Instance(), uintptr(Index))
	return AsContainedAction(r1)
}

func (m *TCustomActionList) SetActions(Index int32, AValue IContainedAction) {
	LCL().SysCallN(1055, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *TCustomActionList) ActionCount() int32 {
	r1 := LCL().SysCallN(1054, m.Instance())
	return int32(r1)
}

func (m *TCustomActionList) Images() ICustomImageList {
	r1 := LCL().SysCallN(1059, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomActionList) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(1059, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomActionList) State() TActionListState {
	r1 := LCL().SysCallN(1062, 0, m.Instance(), 0)
	return TActionListState(r1)
}

func (m *TCustomActionList) SetState(AValue TActionListState) {
	LCL().SysCallN(1062, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomActionList) ActionByName(ActionName string) IContainedAction {
	r1 := LCL().SysCallN(1053, m.Instance(), PascalStr(ActionName))
	return AsContainedAction(r1)
}

func (m *TCustomActionList) GetEnumeratorForActionListEnumerator() IActionListEnumerator {
	r1 := LCL().SysCallN(1058, m.Instance())
	return AsActionListEnumerator(r1)
}

func (m *TCustomActionList) IndexOfName(ActionName string) int32 {
	r1 := LCL().SysCallN(1060, m.Instance(), PascalStr(ActionName))
	return int32(r1)
}

func (m *TCustomActionList) IsShortCut(Message *TLMKey) bool {
	var result0 uintptr
	r1 := LCL().SysCallN(1061, m.Instance(), uintptr(unsafe.Pointer(&result0)))
	*Message = *(*TLMKey)(getPointer(result0))
	return GoBool(r1)
}

func CustomActionListClass() TClass {
	ret := LCL().SysCallN(1056)
	return TClass(ret)
}
