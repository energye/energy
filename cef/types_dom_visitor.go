//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefDomVisitor
type ICefDomVisitor struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

// DomVisitorRef -> ICefDomVisitor
var DomVisitorRef domVisitor

type domVisitor uintptr

func (*domVisitor) New() *ICefDomVisitor {
	var result uintptr
	imports.Proc(def.CefDomVisitorRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefDomVisitor{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefDomVisitor) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefDomVisitor) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefDomVisitor) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefDomVisitor) SetOnVisit(fn onVisit) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CefDomVisitor_Visit).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// ************************** events ************************** //

type onVisit func(document *ICefDomDocument)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case onVisit:
			document := &ICefDomDocument{instance: getPtr(0)}
			fn.(onVisit)(document)
		default:
			return false
		}
		return true
	})
}
