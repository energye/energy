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

// ICefDragData Parent: ICefBaseRefCounted
//
//	Interface used to represent drag data. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_drag_data_capi.h">CEF source file: /include/capi/cef_drag_data_capi.h (cef_drag_data_t))</a>
type ICefDragData interface {
	ICefBaseRefCounted
	// Clone
	//  Returns a copy of the current object.
	Clone() ICefDragData // function
	// IsReadOnly
	//  Returns true (1) if this object is read-only.
	IsReadOnly() bool // function
	// IsLink
	//  Returns true (1) if the drag data is a link.
	IsLink() bool // function
	// IsFragment
	//  Returns true (1) if the drag data is a text or html fragment.
	IsFragment() bool // function
	// IsFile
	//  Returns true (1) if the drag data is a file.
	IsFile() bool // function
	// GetLinkUrl
	//  Return the link URL that is being dragged.
	GetLinkUrl() string // function
	// GetLinkTitle
	//  Return the title associated with the link being dragged.
	GetLinkTitle() string // function
	// GetLinkMetadata
	//  Return the metadata, if any, associated with the link being dragged.
	GetLinkMetadata() string // function
	// GetFragmentText
	//  Return the plain text fragment that is being dragged.
	GetFragmentText() string // function
	// GetFragmentHtml
	//  Return the text/html fragment that is being dragged.
	GetFragmentHtml() string // function
	// GetFragmentBaseUrl
	//  Return the base URL that the fragment came from. This value is used for resolving relative URLs and may be NULL.
	GetFragmentBaseUrl() string // function
	// GetFileName
	//  Return the name of the file being dragged out of the browser window.
	GetFileName() string // function
	// GetFileContents
	//  Write the contents of the file being dragged out of the web view into |writer|. Returns the number of bytes sent to |writer|. If |writer| is NULL this function will return the size of the file contents in bytes. Call get_file_name() to get a suggested name for the file.
	GetFileContents(writer ICefStreamWriter) NativeUInt // function
	// GetFileNames
	//  Retrieve the list of file names that are being dragged into the browser window.
	GetFileNames(names *IStrings) int32 // function
	// GetFilePaths
	//  Retrieve the list of file paths that are being dragged into the browser window.
	GetFilePaths(paths *IStrings) int32 // function
	// GetImage
	//  Get the image representation of drag data. May return NULL if no image representation is available.
	GetImage() ICefImage // function
	// GetImageHotspot
	//  Get the image hotspot (drag start location relative to image dimensions).
	GetImageHotspot() (resultCefPoint TCefPoint) // function
	// HasImage
	//  Returns true (1) if an image representation of drag data is available.
	HasImage() bool // function
	// SetLinkUrl
	//  Set the link URL that is being dragged.
	SetLinkUrl(url string) // procedure
	// SetLinkTitle
	//  Set the title associated with the link being dragged.
	SetLinkTitle(title string) // procedure
	// SetLinkMetadata
	//  Set the metadata associated with the link being dragged.
	SetLinkMetadata(data string) // procedure
	// SetFragmentText
	//  Set the plain text fragment that is being dragged.
	SetFragmentText(text string) // procedure
	// SetFragmentHtml
	//  Set the text/html fragment that is being dragged.
	SetFragmentHtml(html string) // procedure
	// SetFragmentBaseUrl
	//  Set the base URL that the fragment came from.
	SetFragmentBaseUrl(baseUrl string) // procedure
	// ResetFileContents
	//  Reset the file contents. You should do this before calling ICefBrowserHost.DragTargetDragEnter as the web view does not allow us to drag in this kind of data.
	ResetFileContents() // procedure
	// AddFile
	//  Add a file that is being dragged into the webview.
	AddFile(path, displayName string) // procedure
	// ClearFilenames
	//  Clear list of filenames.
	ClearFilenames() // procedure
}

// TCefDragData Parent: TCefBaseRefCounted
//
//	Interface used to represent drag data. The functions of this interface may be called on any thread.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_drag_data_capi.h">CEF source file: /include/capi/cef_drag_data_capi.h (cef_drag_data_t))</a>
type TCefDragData struct {
	TCefBaseRefCounted
}

// DragDataRef -> ICefDragData
var DragDataRef dragData

// dragData TCefDragData Ref
type dragData uintptr

func (m *dragData) UnWrap(data uintptr) ICefDragData {
	var resultCefDragData uintptr
	CEF().SysCallN(931, uintptr(data), uintptr(unsafePointer(&resultCefDragData)))
	return AsCefDragData(resultCefDragData)
}

func (m *dragData) New() ICefDragData {
	var resultCefDragData uintptr
	CEF().SysCallN(923, uintptr(unsafePointer(&resultCefDragData)))
	return AsCefDragData(resultCefDragData)
}

func (m *TCefDragData) Clone() ICefDragData {
	var resultCefDragData uintptr
	CEF().SysCallN(905, m.Instance(), uintptr(unsafePointer(&resultCefDragData)))
	return AsCefDragData(resultCefDragData)
}

func (m *TCefDragData) IsReadOnly() bool {
	r1 := CEF().SysCallN(922, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) IsLink() bool {
	r1 := CEF().SysCallN(921, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) IsFragment() bool {
	r1 := CEF().SysCallN(920, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) IsFile() bool {
	r1 := CEF().SysCallN(919, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) GetLinkUrl() string {
	r1 := CEF().SysCallN(917, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetLinkTitle() string {
	r1 := CEF().SysCallN(916, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetLinkMetadata() string {
	r1 := CEF().SysCallN(915, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFragmentText() string {
	r1 := CEF().SysCallN(912, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFragmentHtml() string {
	r1 := CEF().SysCallN(911, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFragmentBaseUrl() string {
	r1 := CEF().SysCallN(910, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFileName() string {
	r1 := CEF().SysCallN(907, m.Instance())
	return GoStr(r1)
}

func (m *TCefDragData) GetFileContents(writer ICefStreamWriter) NativeUInt {
	r1 := CEF().SysCallN(906, m.Instance(), GetObjectUintptr(writer))
	return NativeUInt(r1)
}

func (m *TCefDragData) GetFileNames(names *IStrings) int32 {
	var result0 uintptr
	r1 := CEF().SysCallN(908, m.Instance(), uintptr(unsafePointer(&result0)))
	*names = AsStrings(result0)
	return int32(r1)
}

func (m *TCefDragData) GetFilePaths(paths *IStrings) int32 {
	var result0 uintptr
	r1 := CEF().SysCallN(909, m.Instance(), uintptr(unsafePointer(&result0)))
	*paths = AsStrings(result0)
	return int32(r1)
}

func (m *TCefDragData) GetImage() ICefImage {
	var resultCefImage uintptr
	CEF().SysCallN(913, m.Instance(), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefDragData) GetImageHotspot() (resultCefPoint TCefPoint) {
	CEF().SysCallN(914, m.Instance(), uintptr(unsafePointer(&resultCefPoint)))
	return
}

func (m *TCefDragData) HasImage() bool {
	r1 := CEF().SysCallN(918, m.Instance())
	return GoBool(r1)
}

func (m *TCefDragData) SetLinkUrl(url string) {
	CEF().SysCallN(930, m.Instance(), PascalStr(url))
}

func (m *TCefDragData) SetLinkTitle(title string) {
	CEF().SysCallN(929, m.Instance(), PascalStr(title))
}

func (m *TCefDragData) SetLinkMetadata(data string) {
	CEF().SysCallN(928, m.Instance(), PascalStr(data))
}

func (m *TCefDragData) SetFragmentText(text string) {
	CEF().SysCallN(927, m.Instance(), PascalStr(text))
}

func (m *TCefDragData) SetFragmentHtml(html string) {
	CEF().SysCallN(926, m.Instance(), PascalStr(html))
}

func (m *TCefDragData) SetFragmentBaseUrl(baseUrl string) {
	CEF().SysCallN(925, m.Instance(), PascalStr(baseUrl))
}

func (m *TCefDragData) ResetFileContents() {
	CEF().SysCallN(924, m.Instance())
}

func (m *TCefDragData) AddFile(path, displayName string) {
	CEF().SysCallN(903, m.Instance(), PascalStr(path), PascalStr(displayName))
}

func (m *TCefDragData) ClearFilenames() {
	CEF().SysCallN(904, m.Instance())
}
