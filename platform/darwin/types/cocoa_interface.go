//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

import (
	"github.com/energye/lcl/types"
	"unsafe"
)

type INSObject interface {
	Instance() uintptr
}

type INSResponder interface {
	INSObject
}

type INSToolBar interface {
	INSObject
}

type INSView interface {
	INSResponder
}

type INSVisualEffectView interface {
	INSView
	SetAutoresizingMask(mask NSAutoresizingMaskOptions)
	SetBlendingMode(mode NSVisualEffectBlendingMode)
	SetState(state NSVisualEffectState)
}

type INSWindow interface {
	INSResponder
	Restore()
	Maximize()
	Minimized()
	ExitMinimized()
	EnterFullScreen()
	ExitFullScreen()
	Drag()
	SetBackgroundColor(r, g, b, alpha uint8)
	SetRadius(radius float32)
	SetTransparent() INSVisualEffectView
	SwitchFrostedMaterial(appearanceName string)
	AddSubview(view INSView, x, y, width, height float32)
	ContentViewFrame() (rect types.TRect)
}

type INSWindowDelegate interface {
	INSObject
}

type INSApp interface {
	DockHide()
	DockShow()
	SetPresentationOptions(options NSApplicationPresentationOptions)
	SetMainMenu(nsMenu unsafe.Pointer)
	GetActivationPolicy() int
	GetPresentationOptions() NSApplicationPresentationOptions
	Activate()
	Deactivate()
	Hide()
	UnHide()
	Terminate()
	InitAppDelegate()
	SetOnOpenURLs(fn TOpenURLsEvent)
	SetOnUniversalLink(fn TUniversalLinkEvent)
}

type IWkWebView interface {
	INSView
	SetWebviewTransparent(isTransparent bool)
	UpdateBounds(window INSWindow, x, y, width, height float32)
	BecomeFirstResponder()
	Undo()
	Redo()
	Cut()
	Copy()
	Paste()
	SelectAll()
	RegisterPerformKeyEquivalentMethod()
	ConvertPoint(inPoint types.TPoint) (point types.TPoint)
	ExecuteScriptCallback(script string, callback TEvaluateScriptCallbackEvent)
}

type INSPasteboard interface {
	Types() NSTypes
	PasteboardData() *PasteboardData
}

type INSDraggingInfo interface {
	DraggingPasteboard() INSPasteboard
	DraggingLocation() (point types.TPoint)
}
