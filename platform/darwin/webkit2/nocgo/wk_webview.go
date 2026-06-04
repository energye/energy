// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

package nocgo

import (
	"github.com/ebitengine/purego/objc"
	"github.com/energye/energy/v3/platform/darwin/cocoa/nocgo"
	. "github.com/energye/energy/v3/platform/darwin/types"
	"github.com/energye/lcl/types"
	"unsafe"
)

type WkWebView struct {
	nocgo.NSView
}

func AsWkWebView(ptr unsafe.Pointer) IWkWebView {
	if ptr == nil {
		return nil
	}
	m := &WkWebView{}
	m.SetInstance(ptr)
	return m
}

func (m *WkWebView) SetWebviewTransparent(isTransparent bool) {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	value := !isTransparent
	webview.Send(objc.RegisterName("setValue:forKey:"),
		objc.ID(objc.GetClass("NSNumber")).Send(objc.RegisterName("numberWithBool:"), value),
		objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"), "drawsBackground"),
	)
}

func (m *WkWebView) UpdateBounds(window INSWindow, x, y, width, height float32) {
	if m.Instance() == 0 || window == nil {
		return
	}
	nsWindow := objc.ID(window.Instance())
	contentView := nsWindow.Send(objc.RegisterName("contentView"))
	if contentView == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	frame := CGRect{
		X:      float64(x),
		Y:      float64(y),
		Width:  float64(width),
		Height: float64(height),
	}
	webview.Send(objc.RegisterName("setFrame:"), frame)
}

func (m *WkWebView) BecomeFirstResponder() {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	webview.Send(objc.RegisterName("becomeFirstResponder"))
}

func (m *WkWebView) Undo() {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	undoManager := webview.Send(objc.RegisterName("undoManager"))
	if undoManager != 0 {
		canUndo := undoManager.Send(objc.RegisterName("canUndo"))
		if canUndo != 0 {
			undoManager.Send(objc.RegisterName("undo"))
		}
	}
}

func (m *WkWebView) Redo() {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	undoManager := webview.Send(objc.RegisterName("undoManager"))
	if undoManager != 0 {
		canRedo := undoManager.Send(objc.RegisterName("canRedo"))
		if canRedo != 0 {
			undoManager.Send(objc.RegisterName("redo"))
		}
	}
}

func (m *WkWebView) Cut() {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	webview.Send(objc.RegisterName("cut:"), webview)
}

func (m *WkWebView) Copy() {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	webview.Send(objc.RegisterName("copy:"), webview)
}

func (m *WkWebView) Paste() {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	webview.Send(objc.RegisterName("paste:"), webview)
}

func (m *WkWebView) SelectAll() {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	webview.Send(objc.RegisterName("selectAll:"), webview)
}

func (m *WkWebView) RegisterPerformKeyEquivalentMethod() {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	webviewClass := webview.Send(objc.RegisterName("class"))

	selector := objc.RegisterName("performKeyEquivalent:")
	methodType := "B@:@"

	class := objc.Class(webviewClass)
	exists := class.AddMethod(selector, objc.NewIMP(performKeyEquivalent), methodType)
	_ = exists
}

func performKeyEquivalent(self objc.ID, _cmd objc.SEL, event objc.ID) bool {
	modifierFlags := event.Send(objc.RegisterName("modifierFlags"))
	keyCode := event.Send(objc.RegisterName("keyCode"))

	const (
		NSEventModifierFlagCommand = 1 << 20
		NSEventModifierFlagShift   = 1 << 17
		NSEventModifierFlagOption  = 1 << 19
	)

	cmd := modifierFlags & NSEventModifierFlagCommand
	shift := modifierFlags & NSEventModifierFlagShift
	option := modifierFlags & NSEventModifierFlagOption

	var selector objc.SEL

	if cmd != 0 && shift == 0 && option == 0 {
		switch keyCode {
		case 0:
			selector = objc.RegisterName("selectAll:")
		case 8:
			selector = objc.RegisterName("copy:")
		case 9:
			selector = objc.RegisterName("paste:")
		case 7:
			selector = objc.RegisterName("cut:")
		case 6:
			self.Send(objc.RegisterName("undo"))
			return true
		case 15:
			selector = objc.RegisterName("reload:")
		}
	} else if cmd != 0 && shift != 0 && option == 0 {
		switch keyCode {
		case 6:
			self.Send(objc.RegisterName("redo"))
			return true
		}
	}

	if selector != 0 && self.Send(objc.RegisterName("respondsToSelector:"), selector) != 0 {
		self.Send(selector, self)
		return true
	}

	return false
}

func (m *WkWebView) ConvertPoint(inPoint types.TPoint) (point types.TPoint) {
	if m.Instance() == 0 {
		return
	}
	webview := objc.ID(m.Instance())
	inNsPoint := CGPoint{X: float64(inPoint.X), Y: float64(inPoint.Y)}
	outNsPoint := objc.Send[CGPoint](webview, objc.RegisterName("convertPoint:fromView:"), inNsPoint, objc.ID(0))
	point.X = int32(outNsPoint.X)
	point.Y = int32(outNsPoint.Y)
	return
}
