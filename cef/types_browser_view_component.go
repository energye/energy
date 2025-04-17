//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// VF 窗口组件 BrowserView

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// TCEFBrowserViewComponent
type TCEFBrowserViewComponent struct {
	*TCEFViewComponent
}

// BrowserViewComponentRef -> TCEFBrowserViewComponent
var BrowserViewComponentRef browserViewComponent

type browserViewComponent uintptr

// New 创建BrowserViewComponent组件
func (*browserViewComponent) New(AOwner lcl.IComponent) *TCEFBrowserViewComponent {
	var result uintptr
	imports.Proc(def.CEFBrowserViewComponent_Create).Call(lcl.CheckPtr(AOwner), uintptr(unsafe.Pointer(&result)))
	return &TCEFBrowserViewComponent{&TCEFViewComponent{instance: getInstance(result)}}
}

// Create a new ICefBrowserView. The underlying ICefBrowser will not be created
// until this view is added to the views hierarchy. The optional |extra_info|
// parameter provides an opportunity to specify extra information specific to
// the created browser that will be passed to
// ICefRenderProcessHandler.OnBrowserCreated in the render process.
func (m *TCEFBrowserViewComponent) CreateBrowserView(client *ICefClient, url string, browserSettings TCefBrowserSettings,
	extraInfo *ICefDictionaryValue, requestContext *ICefRequestContext) {
	browserSettingsPtr := browserSettings.ToPtr()
	imports.Proc(def.CEFBrowserViewComponent_CreateBrowserView).Call(m.Instance(), uintptr(client.instance), api.PascalStr(url),
		uintptr(unsafe.Pointer(browserSettingsPtr)), extraInfo.Instance(), requestContext.Instance())
}

// Instance
func (m *TCEFBrowserViewComponent) Instance() uintptr {
	return uintptr(m.instance)
}

func (m *TCEFBrowserViewComponent) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *TCEFBrowserViewComponent) Free() {
	if m.instance != nil {
		m.instance = nil
	}
}

// Updates the internal ICefBrowserView with the ICefBrowserView associated with |browser|.
func (m *TCEFBrowserViewComponent) GetForBrowser(browser *ICefBrowser) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_GetForBrowser).Call(m.Instance(), browser.Instance())
}

// Sets whether normal priority accelerators are first forwarded to the web
// content (`keydown` event handler) or ICefKeyboardHandler. Normal priority
// accelerators can be registered via ICefWindow.SetAccelerator (with
// |high_priority|=false) or internally for standard accelerators supported
// by Chrome style. If |prefer_accelerators| is true then the matching
// accelerator will be triggered immediately (calling
// ICefWindowDelegate.OnAccelerator or ICefCommandHandler.OnChromeCommand
// respectively) and the event will not be forwarded to the web content or
// ICefKeyboardHandler first. If |prefer_accelerators| is false then the
// matching accelerator will only be triggered if the event is not handled by
// web content (`keydown` event handler that calls `event.preventDefault()`)
// or by ICefKeyboardHandler. The default value is false.
func (m *TCEFBrowserViewComponent) SetPreferAccelerators(preferAccelerators bool) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_SetPreferAccelerators).Call(m.Instance(), api.PascalBool(preferAccelerators))
}

// Request keyboard focus. If this View is focusable it will become the focused View.
func (m *TCEFBrowserViewComponent) RequestFocus() {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_RequestFocus).Call(m.Instance())
}

// Returns the ICefBrowser hosted by this BrowserView. Will return NULL if
// the browser has not yet been created or has already been destroyed.
func (m *TCEFBrowserViewComponent) GetBrowser() *ICefBrowser {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFBrowserViewComponent_Browser).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBrowser{instance: unsafe.Pointer(result)}
}

func (m *TCEFBrowserViewComponent) BrowserView() *ICefBrowserView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFBrowserViewComponent_BrowserView).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &ICefBrowserView{&ICefView{instance: unsafe.Pointer(result)}}
}

// Returns the Chrome toolbar associated with this BrowserView. Only
// supported when using Chrome style. The ICefBrowserViewDelegate.GetChromeToolbarType
// function must return a value other than
// CEF_CTT_NONE and the toolbar will not be available until after this
// BrowserView is added to a ICefWindow and
// ICefViewDelegate.OnWindowChanged() has been called.
func (m *TCEFBrowserViewComponent) GetChromeToolbar() *ICefView {
	if !m.IsValid() {
		return nil
	}
	var result uintptr
	imports.Proc(def.CEFBrowserViewComponent_ChromeToolbar).Call(m.Instance(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefView{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Returns the runtime style for this BrowserView (ALLOY or CHROME). See
// TCefRuntimeStyle documentation for details.
func (m *TCEFBrowserViewComponent) GetRuntimeStyle() consts.TCefRuntimeStyle {
	if !m.IsValid() {
		return 0
	}
	r1, _, _ := imports.Proc(def.CEFBrowserViewComponent_RuntimeStyle).Call(m.Instance())
	return consts.TCefRuntimeStyle(r1)
}

// SetOnBrowserCreated
func (m *TCEFBrowserViewComponent) SetOnBrowserCreated(fn browserViewOnBrowserCreated) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_SetOnBrowserCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnBrowserDestroyed
func (m *TCEFBrowserViewComponent) SetOnBrowserDestroyed(fn browserViewOnBrowserDestroyed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_SetOnBrowserDestroyed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGetDelegateForPopupBrowserView
func (m *TCEFBrowserViewComponent) SetOnGetDelegateForPopupBrowserView(fn browserViewOnGetDelegateForPopupBrowserView) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_SetOnGetDelegateForPopupBrowserView).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnPopupBrowserViewCreated
func (m *TCEFBrowserViewComponent) SetOnPopupBrowserViewCreated(fn browserViewOnPopupBrowserViewCreated) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_SetOnPopupBrowserViewCreated).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGetChromeToolbarType
func (m *TCEFBrowserViewComponent) SetOnGetChromeToolbarType(fn browserViewOnGetChromeToolbarType) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_SetOnGetChromeToolbarType).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnUseFramelessWindowForPictureInPicture
func (m *TCEFBrowserViewComponent) SetOnUseFramelessWindowForPictureInPicture(fn browserViewOnUseFramelessWindowForPictureInPicture) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_SetOnUseFramelessWindowForPictureInPicture).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnGestureCommand
func (m *TCEFBrowserViewComponent) SetOnGestureCommand(fn browserViewOnGestureCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.CEFBrowserViewComponent_SetOnGestureCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TCEFBrowserViewComponent) SetOnGetBrowserRuntimeStyle(fn browserViewOnGetBrowserRuntimeStyle) {
	imports.Proc(def.CEFBrowserViewComponent_SetOnGetBrowserRuntimeStyle).Call(m.Instance(), api.MakeEventDataPtr(fn))
}
