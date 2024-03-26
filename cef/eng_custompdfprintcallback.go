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

// ICefCustomPDFPrintCallBack Parent: ICefPdfPrintCallback
type ICefCustomPDFPrintCallBack interface {
	ICefPdfPrintCallback
}

// TCefCustomPDFPrintCallBack Parent: TCefPdfPrintCallback
type TCefCustomPDFPrintCallBack struct {
	TCefPdfPrintCallback
}

func NewCefCustomPDFPrintCallBack(aEvents IChromiumEvents) ICefCustomPDFPrintCallBack {
	r1 := CEF().SysCallN(782, GetObjectUintptr(aEvents))
	return AsCefCustomPDFPrintCallBack(r1)
}
