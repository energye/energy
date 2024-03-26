//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICustomResponseFilter Parent: ICefResponseFilter
type ICustomResponseFilter interface {
	ICefResponseFilter
	SetOnFilter(fn TOnFilter)         // property event
	SetOnInitFilter(fn TOnInitFilter) // property event
}

// TCustomResponseFilter Parent: TCefResponseFilter
type TCustomResponseFilter struct {
	TCefResponseFilter
	filterPtr     uintptr
	initFilterPtr uintptr
}

func NewCustomResponseFilter() ICustomResponseFilter {
	r1 := CEF().SysCallN(2163)
	return AsCustomResponseFilter(r1)
}

func (m *TCustomResponseFilter) SetOnFilter(fn TOnFilter) {
	if m.filterPtr != 0 {
		RemoveEventElement(m.filterPtr)
	}
	m.filterPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2164, m.Instance(), m.filterPtr)
}

func (m *TCustomResponseFilter) SetOnInitFilter(fn TOnInitFilter) {
	if m.initFilterPtr != 0 {
		RemoveEventElement(m.initFilterPtr)
	}
	m.initFilterPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2165, m.Instance(), m.initFilterPtr)
}
