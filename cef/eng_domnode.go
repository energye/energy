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

// ICefDomNode Parent: ICefBaseRefCounted
//
//	Interface used to represent a DOM node. The functions of this interface should only be called on the render process main thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domnode_t))
type ICefDomNode interface {
	ICefBaseRefCounted
	// GetType
	//  Returns the type for this node.
	GetType() TCefDomNodeType // function
	// IsText
	//  Returns true (1) if this is a text node.
	IsText() bool // function
	// IsElement
	//  Returns true (1) if this is an element node.
	IsElement() bool // function
	// IsEditable
	//  Returns true (1) if this is an editable node.
	IsEditable() bool // function
	// IsFormControlElement
	//  Returns true (1) if this is a form control element node.
	IsFormControlElement() bool // function
	// GetFormControlElementType
	//  Returns the type of this form control element node.
	GetFormControlElementType() string // function
	// IsSame
	//  Returns true (1) if this object is pointing to the same handle as |that| object.
	IsSame(that ICefDomNode) bool // function
	// GetName
	//  Returns the name of this node.
	GetName() string // function
	// GetValue
	//  Returns the value of this node.
	GetValue() string // function
	// SetValue
	//  Set the value of this node. Returns true (1) on success.
	SetValue(value string) bool // function
	// GetAsMarkup
	//  Returns the contents of this node as markup.
	GetAsMarkup() string // function
	// GetDocument
	//  Returns the document associated with this node.
	GetDocument() ICefDomDocument // function
	// GetParent
	//  Returns the parent node.
	GetParent() ICefDomNode // function
	// GetPreviousSibling
	//  Returns the previous sibling node.
	GetPreviousSibling() ICefDomNode // function
	// GetNextSibling
	//  Returns the next sibling node.
	GetNextSibling() ICefDomNode // function
	// HasChildren
	//  Returns true (1) if this node has child nodes.
	HasChildren() bool // function
	// GetFirstChild
	//  Return the first child node.
	GetFirstChild() ICefDomNode // function
	// GetLastChild
	//  Returns the last child node.
	GetLastChild() ICefDomNode // function
	// GetElementTagName
	//  Returns the tag name of this element.
	GetElementTagName() string // function
	// HasElementAttributes
	//  Returns true (1) if this element has attributes.
	HasElementAttributes() bool // function
	// HasElementAttribute
	//  Returns true (1) if this element has an attribute named |attrName|.
	HasElementAttribute(attrName string) bool // function
	// GetElementAttribute
	//  Returns the element attribute named |attrName|.
	GetElementAttribute(attrName string) string // function
	// SetElementAttribute
	//  Set the value for the element attribute named |attrName|. Returns true (1) on success.
	SetElementAttribute(attrName, value string) bool // function
	// GetElementInnerText
	//  Returns the inner text of the element.
	GetElementInnerText() string // function
	// GetElementBounds
	//  Returns the bounds of the element in device pixels. Use "window.devicePixelRatio" to convert to/from CSS pixels.
	GetElementBounds() (resultCefRect TCefRect) // function
	// GetElementAttributes
	//  Returns a ICefStringMap of all element attributes.
	GetElementAttributes(attrMap ICefStringMap) // procedure
	// GetElementAttributes1
	//  Returns a ICefStringMap of all element attributes.
	GetElementAttributes1(attrList *IStrings) // procedure
}

// TCefDomNode Parent: TCefBaseRefCounted
//
//	Interface used to represent a DOM node. The functions of this interface should only be called on the render process main thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_dom_capi.h">CEF source file: /include/capi/cef_dom_capi.h (cef_domnode_t))
type TCefDomNode struct {
	TCefBaseRefCounted
}

// DomNodeRef -> ICefDomNode
var DomNodeRef domNode

// domNode TCefDomNode Ref
type domNode uintptr

func (m *domNode) UnWrap(data uintptr) ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(878, uintptr(data), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetType() TCefDomNodeType {
	r1 := CEF().SysCallN(866, m.Instance())
	return TCefDomNodeType(r1)
}

func (m *TCefDomNode) IsText() bool {
	r1 := CEF().SysCallN(875, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) IsElement() bool {
	r1 := CEF().SysCallN(872, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) IsEditable() bool {
	r1 := CEF().SysCallN(871, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) IsFormControlElement() bool {
	r1 := CEF().SysCallN(873, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) GetFormControlElementType() string {
	r1 := CEF().SysCallN(860, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) IsSame(that ICefDomNode) bool {
	r1 := CEF().SysCallN(874, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefDomNode) GetName() string {
	r1 := CEF().SysCallN(862, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) GetValue() string {
	r1 := CEF().SysCallN(867, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) SetValue(value string) bool {
	r1 := CEF().SysCallN(877, m.Instance(), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefDomNode) GetAsMarkup() string {
	r1 := CEF().SysCallN(851, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) GetDocument() ICefDomDocument {
	var resultCefDomDocument uintptr
	CEF().SysCallN(852, m.Instance(), uintptr(unsafePointer(&resultCefDomDocument)))
	return AsCefDomDocument(resultCefDomDocument)
}

func (m *TCefDomNode) GetParent() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(864, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetPreviousSibling() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(865, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetNextSibling() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(863, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) HasChildren() bool {
	r1 := CEF().SysCallN(868, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) GetFirstChild() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(859, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetLastChild() ICefDomNode {
	var resultCefDomNode uintptr
	CEF().SysCallN(861, m.Instance(), uintptr(unsafePointer(&resultCefDomNode)))
	return AsCefDomNode(resultCefDomNode)
}

func (m *TCefDomNode) GetElementTagName() string {
	r1 := CEF().SysCallN(858, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) HasElementAttributes() bool {
	r1 := CEF().SysCallN(870, m.Instance())
	return GoBool(r1)
}

func (m *TCefDomNode) HasElementAttribute(attrName string) bool {
	r1 := CEF().SysCallN(869, m.Instance(), PascalStr(attrName))
	return GoBool(r1)
}

func (m *TCefDomNode) GetElementAttribute(attrName string) string {
	r1 := CEF().SysCallN(853, m.Instance(), PascalStr(attrName))
	return GoStr(r1)
}

func (m *TCefDomNode) SetElementAttribute(attrName, value string) bool {
	r1 := CEF().SysCallN(876, m.Instance(), PascalStr(attrName), PascalStr(value))
	return GoBool(r1)
}

func (m *TCefDomNode) GetElementInnerText() string {
	r1 := CEF().SysCallN(857, m.Instance())
	return GoStr(r1)
}

func (m *TCefDomNode) GetElementBounds() (resultCefRect TCefRect) {
	CEF().SysCallN(856, m.Instance(), uintptr(unsafePointer(&resultCefRect)))
	return
}

func (m *TCefDomNode) GetElementAttributes(attrMap ICefStringMap) {
	CEF().SysCallN(854, m.Instance(), GetObjectUintptr(attrMap))
}

func (m *TCefDomNode) GetElementAttributes1(attrList *IStrings) {
	var result0 uintptr
	CEF().SysCallN(855, m.Instance(), uintptr(unsafePointer(&result0)))
	*attrList = AsStrings(result0)
}
