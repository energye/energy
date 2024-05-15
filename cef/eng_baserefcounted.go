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

// ICefBaseRefCounted Parent: IObject
//
//	All ref-counted framework interfaces must inherit from this interface.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_base_capi.h">CEF source file: /include/capi/cef_base_capi.h (cef_base_ref_counted_t))</a>
type ICefBaseRefCounted interface {
	IObject
	// HasOneRef
	//  Returns true (1) if the current reference count is 1.
	HasOneRef() bool // function
	// HasAtLeastOneRef
	//  Returns true (1) if the current reference count is at least 1.
	HasAtLeastOneRef() bool // function
	// SameAs
	//  Compares the aData pointer with the FData field if the current instance.
	SameAs(aData uintptr) bool // function
	// SameAs1
	//  Compares the aData pointer with the FData field if the current instance.
	SameAs1(aBaseRefCounted ICefBaseRefCounted) bool // function
	// Wrap
	//  Called to increment the reference count for the object. Should be called for every new copy of a pointer to a given object.
	Wrap() uintptr // function
	// DestroyOtherRefs
	//  Releases all other instances.
	DestroyOtherRefs() // procedure
}

// TCefBaseRefCounted Parent: TObject
//
//	All ref-counted framework interfaces must inherit from this interface.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_base_capi.h">CEF source file: /include/capi/cef_base_capi.h (cef_base_ref_counted_t))</a>
type TCefBaseRefCounted struct {
	TObject
}

func NewCefBaseRefCounted(data uintptr) ICefBaseRefCounted {
	r1 := CEF().SysCallN(587, uintptr(data))
	return AsCefBaseRefCounted(r1)
}

// BaseRefCountedRef -> ICefBaseRefCounted
var BaseRefCountedRef baseRefCounted

// baseRefCounted TCefBaseRefCounted Ref
type baseRefCounted uintptr

func (m *baseRefCounted) UnWrap(data uintptr) ICefBaseRefCounted {
	var resultCefBaseRefCounted uintptr
	CEF().SysCallN(593, uintptr(data), uintptr(unsafePointer(&resultCefBaseRefCounted)))
	return AsCefBaseRefCounted(resultCefBaseRefCounted)
}

func (m *TCefBaseRefCounted) HasOneRef() bool {
	r1 := CEF().SysCallN(590, m.Instance())
	return GoBool(r1)
}

func (m *TCefBaseRefCounted) HasAtLeastOneRef() bool {
	r1 := CEF().SysCallN(589, m.Instance())
	return GoBool(r1)
}

func (m *TCefBaseRefCounted) SameAs(aData uintptr) bool {
	r1 := CEF().SysCallN(591, m.Instance(), uintptr(aData))
	return GoBool(r1)
}

func (m *TCefBaseRefCounted) SameAs1(aBaseRefCounted ICefBaseRefCounted) bool {
	r1 := CEF().SysCallN(592, m.Instance(), GetObjectUintptr(aBaseRefCounted))
	return GoBool(r1)
}

func (m *TCefBaseRefCounted) Wrap() uintptr {
	r1 := CEF().SysCallN(594, m.Instance())
	return uintptr(r1)
}

func (m *TCefBaseRefCounted) DestroyOtherRefs() {
	CEF().SysCallN(588, m.Instance())
}
