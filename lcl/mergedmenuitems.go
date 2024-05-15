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

// IMergedMenuItems Parent: IObject
type IMergedMenuItems interface {
	IObject
	VisibleCount() int32                  // property
	VisibleItems(Index int32) IMenuItem   // property
	InvisibleCount() int32                // property
	InvisibleItems(Index int32) IMenuItem // property
}

// TMergedMenuItems Parent: TObject
type TMergedMenuItems struct {
	TObject
}

func NewMergedMenuItems(aParent IMenuItem) IMergedMenuItems {
	r1 := LCL().SysCallN(4333, GetObjectUintptr(aParent))
	return AsMergedMenuItems(r1)
}

func (m *TMergedMenuItems) VisibleCount() int32 {
	r1 := LCL().SysCallN(4336, m.Instance())
	return int32(r1)
}

func (m *TMergedMenuItems) VisibleItems(Index int32) IMenuItem {
	r1 := LCL().SysCallN(4337, m.Instance(), uintptr(Index))
	return AsMenuItem(r1)
}

func (m *TMergedMenuItems) InvisibleCount() int32 {
	r1 := LCL().SysCallN(4334, m.Instance())
	return int32(r1)
}

func (m *TMergedMenuItems) InvisibleItems(Index int32) IMenuItem {
	r1 := LCL().SysCallN(4335, m.Instance(), uintptr(Index))
	return AsMenuItem(r1)
}

func MergedMenuItemsClass() TClass {
	ret := LCL().SysCallN(4332)
	return TClass(ret)
}
