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

// ICefDomDocument Parent: ICefBaseRefCounted
//
//	Interface used to represent a DOM document. The functions of this interface should only be called on the render process main thread thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domdocument_t))
type ICefDomDocument interface {
	ICefBaseRefCounted
	// GetType
	//  Returns the document type.
	GetType() TCefDomDocumentType // function
	// GetDocument
	//  Returns the root document node.
	GetDocument() ICefDomNode // function
	// GetBody
	//  Returns the BODY node of an HTML document.
	GetBody() ICefDomNode // function
	// GetHead
	//  Returns the HEAD node of an HTML document.
	GetHead() ICefDomNode // function
	// GetTitle
	//  Returns the title of an HTML document.
	GetTitle() string // function
	// GetElementById
	//  Returns the document element with the specified ID value.
	GetElementById(id string) ICefDomNode // function
	// GetFocusedNode
	//  Returns the node that currently has keyboard focus.
	GetFocusedNode() ICefDomNode // function
	// HasSelection
	//  Returns true (1) if a portion of the document is selected.
	HasSelection() bool // function
	// GetSelectionStartOffset
	//  Returns the selection offset within the start node.
	GetSelectionStartOffset() int32 // function
	// GetSelectionEndOffset
	//  Returns the selection offset within the end node.
	GetSelectionEndOffset() int32 // function
	// GetSelectionAsMarkup
	//  Returns the contents of this selection as markup.
	GetSelectionAsMarkup() string // function
	// GetSelectionAsText
	//  Returns the contents of this selection as text.
	GetSelectionAsText() string // function
	// GetBaseUrl
	//  Returns the base URL for the document.
	GetBaseUrl() string // function
	// GetCompleteUrl
	//  Returns a complete URL based on the document base URL and the specified partial URL.
	GetCompleteUrl(partialURL string) string // function
}

// TCefDomDocument Parent: TCefBaseRefCounted
//
//	Interface used to represent a DOM document. The functions of this interface should only be called on the render process main thread thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domdocument_t))
type TCefDomDocument struct {
	TCefBaseRefCounted
}

// DomDocumentRef -> ICefDomDocument
var DomDocumentRef domDocument

// domDocument TCefDomDocument Ref
type domDocument uintptr

func (m *domDocument) UnWrap(data uintptr) ICefDomDocument {
	var resultCefDomDocument uintptr
	CEF().SysCallN(850, uintptr(data), uintptr(unsafePointer(&resultCefDomDocument)))
	return AsCefDomDocument(resultCefDomDocument)
}

func (m *TCefDomDocument) GetType() TCefDomDocumentType {
	r1 := CEF().SysCallN(848, m.Instance())
	return TCefDomDocumentType(r1)
}

func (m *TCefDomDocument) GetDocument() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(839, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) GetBody() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(837, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) GetHead() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(842, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) GetTitle() string {
	r1 := CEF().SysCallN(847, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomDocument) GetElementById(id string) ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(840, m.Instance(), PascalStr(id), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) GetFocusedNode() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(841, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomDocument) HasSelection() bool {
	r1 := CEF().SysCallN(849, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomDocument) GetSelectionStartOffset() int32 {
	r1 := CEF().SysCallN(846, m.Instance())
	return int32(r1)
}

func (m *TCefDomDocument) GetSelectionEndOffset() int32 {
	r1 := CEF().SysCallN(845, m.Instance())
	return int32(r1)
}

func (m *TCefDomDocument) GetSelectionAsMarkup() string {
	r1 := CEF().SysCallN(843, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomDocument) GetSelectionAsText() string {
	r1 := CEF().SysCallN(844, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomDocument) GetBaseUrl() string {
	r1 := CEF().SysCallN(836, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomDocument) GetCompleteUrl(partialURL string) string {
	r1 := CEF().SysCallN(838, m.Instance(), PascalStr(partialURL))
	return GoStr(r1)
}
