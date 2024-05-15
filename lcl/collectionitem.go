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

// ICollectionItem Parent: IPersistent
type ICollectionItem interface {
	IPersistent
	Collection() ICollection          // property
	SetCollection(AValue ICollection) // property
	ID() int32                        // property
	Index() int32                     // property
	SetIndex(AValue int32)            // property
	DisplayName() string              // property
	SetDisplayName(AValue string)     // property
}

// TCollectionItem Parent: TPersistent
type TCollectionItem struct {
	TPersistent
}

func NewCollectionItem(ACollection ICollection) ICollectionItem {
	r1 := LCL().SysCallN(694, GetObjectUintptr(ACollection))
	return AsCollectionItem(r1)
}

func (m *TCollectionItem) Collection() ICollection {
	r1 := LCL().SysCallN(693, 0, m.Instance(), 0)
	return AsCollection(r1)
}

func (m *TCollectionItem) SetCollection(AValue ICollection) {
	LCL().SysCallN(693, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCollectionItem) ID() int32 {
	r1 := LCL().SysCallN(696, m.Instance())
	return int32(r1)
}

func (m *TCollectionItem) Index() int32 {
	r1 := LCL().SysCallN(697, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCollectionItem) SetIndex(AValue int32) {
	LCL().SysCallN(697, 1, m.Instance(), uintptr(AValue))
}

func (m *TCollectionItem) DisplayName() string {
	r1 := LCL().SysCallN(695, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCollectionItem) SetDisplayName(AValue string) {
	LCL().SysCallN(695, 1, m.Instance(), PascalStr(AValue))
}

func CollectionItemClass() TClass {
	ret := LCL().SysCallN(692)
	return TClass(ret)
}
