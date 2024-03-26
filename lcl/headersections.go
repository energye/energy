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

// IHeaderSections Parent: ICollection
type IHeaderSections interface {
	ICollection
	ItemsForHeaderSection(Index int32) IHeaderSection            // property
	SetItemsForHeaderSection(Index int32, AValue IHeaderSection) // property
	AddForHeaderSection() IHeaderSection                         // function
	AddItem(Item IHeaderSection, Index int32) IHeaderSection     // function
	InsertForHeaderSection(Index int32) IHeaderSection           // function
}

// THeaderSections Parent: TCollection
type THeaderSections struct {
	TCollection
}

func NewHeaderSections(HeaderControl ICustomHeaderControl) IHeaderSections {
	r1 := LCL().SysCallN(3084, GetObjectUintptr(HeaderControl))
	return AsHeaderSections(r1)
}

func (m *THeaderSections) ItemsForHeaderSection(Index int32) IHeaderSection {
	r1 := LCL().SysCallN(3086, 0, m.Instance(), uintptr(Index))
	return AsHeaderSection(r1)
}

func (m *THeaderSections) SetItemsForHeaderSection(Index int32, AValue IHeaderSection) {
	LCL().SysCallN(3086, 1, m.Instance(), uintptr(Index), GetObjectUintptr(AValue))
}

func (m *THeaderSections) AddForHeaderSection() IHeaderSection {
	r1 := LCL().SysCallN(3081, m.Instance())
	return AsHeaderSection(r1)
}

func (m *THeaderSections) AddItem(Item IHeaderSection, Index int32) IHeaderSection {
	r1 := LCL().SysCallN(3082, m.Instance(), GetObjectUintptr(Item), uintptr(Index))
	return AsHeaderSection(r1)
}

func (m *THeaderSections) InsertForHeaderSection(Index int32) IHeaderSection {
	r1 := LCL().SysCallN(3085, m.Instance(), uintptr(Index))
	return AsHeaderSection(r1)
}

func HeaderSectionsClass() TClass {
	ret := LCL().SysCallN(3083)
	return TClass(ret)
}
