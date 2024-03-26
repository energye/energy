//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ICoreWebView2ClientCertificateCollection Parent: IObject
//
//	A collection of client certificate object.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificatecollection">See the ICoreWebView2ClientCertificateCollection article.</a>
type ICoreWebView2ClientCertificateCollection interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2ClientCertificateCollection // property
	// Count
	//  The number of client certificates contained in the
	//  ICoreWebView2ClientCertificateCollection.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificatecollection#get_count">See the ICoreWebView2ClientCertificateCollection article.</a>
	Count() uint32 // property
	// Items
	//  Gets the certificate object at the given index.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificatecollection#getvalueatindex">See the ICoreWebView2ClientCertificateCollection article.</a>
	Items(idx uint32) ICoreWebView2ClientCertificate // property
}

// TCoreWebView2ClientCertificateCollection Parent: TObject
//
//	A collection of client certificate object.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2clientcertificatecollection">See the ICoreWebView2ClientCertificateCollection article.</a>
type TCoreWebView2ClientCertificateCollection struct {
	TObject
}

func NewCoreWebView2ClientCertificateCollection(aBaseIntf ICoreWebView2ClientCertificateCollection) ICoreWebView2ClientCertificateCollection {
	r1 := WV().SysCallN(72, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2ClientCertificateCollection(r1)
}

func (m *TCoreWebView2ClientCertificateCollection) Initialized() bool {
	r1 := WV().SysCallN(73, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2ClientCertificateCollection) BaseIntf() ICoreWebView2ClientCertificateCollection {
	var resultCoreWebView2ClientCertificateCollection uintptr
	WV().SysCallN(69, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2ClientCertificateCollection)))
	return AsCoreWebView2ClientCertificateCollection(resultCoreWebView2ClientCertificateCollection)
}

func (m *TCoreWebView2ClientCertificateCollection) Count() uint32 {
	r1 := WV().SysCallN(71, m.Instance())
	return uint32(r1)
}

func (m *TCoreWebView2ClientCertificateCollection) Items(idx uint32) ICoreWebView2ClientCertificate {
	var resultCoreWebView2ClientCertificate uintptr
	WV().SysCallN(74, m.Instance(), uintptr(idx), uintptr(unsafePointer(&resultCoreWebView2ClientCertificate)))
	return AsCoreWebView2ClientCertificate(resultCoreWebView2ClientCertificate)
}

func CoreWebView2ClientCertificateCollectionClass() TClass {
	ret := WV().SysCallN(70)
	return TClass(ret)
}
