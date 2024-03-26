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

// IDomVisitor Parent: ICefDomVisitor
//
//	Interface to implement for visiting the DOM. The functions of this interface
//	will be called on the render process main thread.
//	<a cref="uCEFTypes|TCefDomVisitor">Implements TCefDomVisitor</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domvisitor_t)</a>
type IDomVisitor interface {
	ICefDomVisitor
	// SetOnDomVisitor
	//  Method executed for visiting the DOM. The document object passed to this
	//  function represents a snapshot of the DOM at the time this function is
	//  executed. DOM objects are only valid for the scope of this function. Do
	//  not keep references to or attempt to access any DOM objects outside the
	//  scope of this function.
	SetOnDomVisitor(fn TOnDomVisitor) // property event
}

// TDomVisitor Parent: TCefDomVisitor
//
//	Interface to implement for visiting the DOM. The functions of this interface
//	will be called on the render process main thread.
//	<a cref="uCEFTypes|TCefDomVisitor">Implements TCefDomVisitor</a>
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domvisitor_t)</a>
type TDomVisitor struct {
	TCefDomVisitor
	domVisitorPtr uintptr
}

func NewDomVisitor() IDomVisitor {
	r1 := CEF().SysCallN(2171)
	return AsDomVisitor(r1)
}

func DomVisitorClass() TClass {
	ret := CEF().SysCallN(2170)
	return TClass(ret)
}

func (m *TDomVisitor) SetOnDomVisitor(fn TOnDomVisitor) {
	if m.domVisitorPtr != 0 {
		RemoveEventElement(m.domVisitorPtr)
	}
	m.domVisitorPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(2172, m.Instance(), m.domVisitorPtr)
}
