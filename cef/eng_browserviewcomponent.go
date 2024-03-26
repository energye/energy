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

// ICEFBrowserViewComponent Parent: ICEFViewComponent
//
//	Component hosting a ICefBrowserView instance. Used in Chrome runtime mode only.
type ICEFBrowserViewComponent interface {
	ICEFViewComponent
	// Browser
	//  Returns the ICefBrowser hosted by this BrowserView. Will return NULL if
	//  the browser has not yet been created or has already been destroyed.
	Browser() ICefBrowser // property
	// BrowserView
	//  ICefBrowserView assiciated to this component.
	BrowserView() ICefBrowserView // property
	// CreateBrowserView
	//  Create a new ICefBrowserView. The underlying ICefBrowser will not be created
	//  until this view is added to the views hierarchy. The optional |extra_info|
	//  parameter provides an opportunity to specify extra information specific to
	//  the created browser that will be passed to
	//  ICefRenderProcessHandler.OnBrowserCreated in the render process.
	CreateBrowserView(client ICefClient, url string, settings *TCefBrowserSettings, extrainfo ICefDictionaryValue, requestcontext ICefRequestContext) bool // function
	// GetForBrowser
	//  Updates the internal ICefBrowserView with the ICefBrowserView associated with |browser|.
	GetForBrowser(browser ICefBrowser) bool // function
	// SetPreferAccelerators
	//  Sets whether accelerators registered with ICefWindow.SetAccelerator are
	//  triggered before or after the event is sent to the ICefBrowser. If
	//  |prefer_accelerators| is true(1) then the matching accelerator will be
	//  triggered immediately and the event will not be sent to the ICefBrowser.
	//  If |prefer_accelerators| is false(0) then the matching accelerator will
	//  only be triggered if the event is not handled by web content or by
	//  ICefKeyboardHandler. The default value is false(0).
	SetPreferAccelerators(preferaccelerators bool) // procedure
	// SetOnBrowserCreated
	//  Called when |browser| associated with |browser_view| is created. This
	//  function will be called after ICefLifeSpanHandler.OnAfterCreated()
	//  is called for |browser| and before OnPopupBrowserViewCreated() is
	//  called for |browser|'s parent delegate if |browser| is a popup.
	SetOnBrowserCreated(fn TOnBrowserCreatedBvc) // property event
	// SetOnBrowserDestroyed
	//  Called when |browser| associated with |browser_view| is destroyed. Release
	//  all references to |browser| and do not attempt to execute any functions on
	//  |browser| after this callback returns. This function will be called before
	//  ICefLifeSpanHandler.OnBeforeClose() is called for |browser|.
	SetOnBrowserDestroyed(fn TOnBrowserDestroyedBvc) // property event
	// SetOnGetDelegateForPopupBrowserView
	//  Called before a new popup BrowserView is created. The popup originated
	//  from |browser_view|. |settings| and |client| are the values returned from
	//  ICefLifeSpanHandler.OnBeforePopup(). |is_devtools| will be true(1)
	//  if the popup will be a DevTools browser. Return the delegate that will be
	//  used for the new popup BrowserView.
	SetOnGetDelegateForPopupBrowserView(fn TOnGetDelegateForPopupBrowserViewBvc) // property event
	// SetOnPopupBrowserViewCreated
	//  Called after |popup_browser_view| is created. This function will be called
	//  after ICefLifeSpanHandler.OnAfterCreated() and OnBrowserCreated()
	//  are called for the new popup browser. The popup originated from
	//  |browser_view|. |is_devtools| will be true(1) if the popup is a DevTools
	//  browser. Optionally add |popup_browser_view| to the views hierarchy
	//  yourself and return true(1). Otherwise return false(0) and a default
	//  ICefWindow will be created for the popup.
	SetOnPopupBrowserViewCreated(fn TOnPopupBrowserViewCreatedBvc) // property event
	// SetOnGetChromeToolbarType
	//  Returns the Chrome toolbar type that will be available via
	//  ICefBrowserView.GetChromeToolbar(). See that function for related
	//  documentation.
	SetOnGetChromeToolbarType(fn TOnGetChromeToolbarTypeBvc) // property event
	// SetOnUseFramelessWindowForPictureInPicture
	//  Return true(1) to create frameless windows for Document picture-in-
	//  picture popups. Content in frameless windows should specify draggable
	//  regions using "-webkit-app-region: drag" CSS.
	SetOnUseFramelessWindowForPictureInPicture(fn TOnUseFramelessWindowForPictureInPictureBvc) // property event
	// SetOnGestureCommand
	//  Called when |browser_view| receives a gesture command. Return true(1) to
	//  handle(or disable) a |gesture_command| or false(0) to propagate the
	//  gesture to the browser for default handling. With the Chrome runtime these
	//  commands can also be handled via cef_command_handler_t::OnChromeCommand.
	SetOnGestureCommand(fn TOnGestureCommandBvc) // property event
}

// TCEFBrowserViewComponent Parent: TCEFViewComponent
//
//	Component hosting a ICefBrowserView instance. Used in Chrome runtime mode only.
type TCEFBrowserViewComponent struct {
	TCEFViewComponent
	browserCreatedPtr                        uintptr
	browserDestroyedPtr                      uintptr
	getDelegateForPopupBrowserViewPtr        uintptr
	popupBrowserViewCreatedPtr               uintptr
	getChromeToolbarTypePtr                  uintptr
	useFramelessWindowForPictureInPicturePtr uintptr
	gestureCommandPtr                        uintptr
}

func NewCEFBrowserViewComponent(aOwner IComponent) ICEFBrowserViewComponent {
	r1 := CEF().SysCallN(89, GetObjectUintptr(aOwner))
	return AsCEFBrowserViewComponent(r1)
}

func (m *TCEFBrowserViewComponent) Browser() ICefBrowser {
	var resultCefBrowser uintptr
	CEF().SysCallN(86, m.Instance(), uintptr(unsafePointer(&resultCefBrowser)))
	return AsCefBrowser(resultCefBrowser)
}

func (m *TCEFBrowserViewComponent) BrowserView() ICefBrowserView {
	var resultCefBrowserView uintptr
	CEF().SysCallN(87, m.Instance(), uintptr(unsafePointer(&resultCefBrowserView)))
	return AsCefBrowserView(resultCefBrowserView)
}

func (m *TCEFBrowserViewComponent) CreateBrowserView(client ICefClient, url string, settings *TCefBrowserSettings, extrainfo ICefDictionaryValue, requestcontext ICefRequestContext) bool {
	inArgs2 := settings.Pointer()
	r1 := CEF().SysCallN(90, m.Instance(), GetObjectUintptr(client), PascalStr(url), uintptr(unsafePointer(inArgs2)), GetObjectUintptr(extrainfo), GetObjectUintptr(requestcontext))
	return GoBool(r1)
}

func (m *TCEFBrowserViewComponent) GetForBrowser(browser ICefBrowser) bool {
	r1 := CEF().SysCallN(91, m.Instance(), GetObjectUintptr(browser))
	return GoBool(r1)
}

func CEFBrowserViewComponentClass() TClass {
	ret := CEF().SysCallN(88)
	return TClass(ret)
}

func (m *TCEFBrowserViewComponent) SetPreferAccelerators(preferaccelerators bool) {
	CEF().SysCallN(99, m.Instance(), PascalBool(preferaccelerators))
}

func (m *TCEFBrowserViewComponent) SetOnBrowserCreated(fn TOnBrowserCreatedBvc) {
	if m.browserCreatedPtr != 0 {
		RemoveEventElement(m.browserCreatedPtr)
	}
	m.browserCreatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(92, m.Instance(), m.browserCreatedPtr)
}

func (m *TCEFBrowserViewComponent) SetOnBrowserDestroyed(fn TOnBrowserDestroyedBvc) {
	if m.browserDestroyedPtr != 0 {
		RemoveEventElement(m.browserDestroyedPtr)
	}
	m.browserDestroyedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(93, m.Instance(), m.browserDestroyedPtr)
}

func (m *TCEFBrowserViewComponent) SetOnGetDelegateForPopupBrowserView(fn TOnGetDelegateForPopupBrowserViewBvc) {
	if m.getDelegateForPopupBrowserViewPtr != 0 {
		RemoveEventElement(m.getDelegateForPopupBrowserViewPtr)
	}
	m.getDelegateForPopupBrowserViewPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(96, m.Instance(), m.getDelegateForPopupBrowserViewPtr)
}

func (m *TCEFBrowserViewComponent) SetOnPopupBrowserViewCreated(fn TOnPopupBrowserViewCreatedBvc) {
	if m.popupBrowserViewCreatedPtr != 0 {
		RemoveEventElement(m.popupBrowserViewCreatedPtr)
	}
	m.popupBrowserViewCreatedPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(97, m.Instance(), m.popupBrowserViewCreatedPtr)
}

func (m *TCEFBrowserViewComponent) SetOnGetChromeToolbarType(fn TOnGetChromeToolbarTypeBvc) {
	if m.getChromeToolbarTypePtr != 0 {
		RemoveEventElement(m.getChromeToolbarTypePtr)
	}
	m.getChromeToolbarTypePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(95, m.Instance(), m.getChromeToolbarTypePtr)
}

func (m *TCEFBrowserViewComponent) SetOnUseFramelessWindowForPictureInPicture(fn TOnUseFramelessWindowForPictureInPictureBvc) {
	if m.useFramelessWindowForPictureInPicturePtr != 0 {
		RemoveEventElement(m.useFramelessWindowForPictureInPicturePtr)
	}
	m.useFramelessWindowForPictureInPicturePtr = MakeEventDataPtr(fn)
	CEF().SysCallN(98, m.Instance(), m.useFramelessWindowForPictureInPicturePtr)
}

func (m *TCEFBrowserViewComponent) SetOnGestureCommand(fn TOnGestureCommandBvc) {
	if m.gestureCommandPtr != 0 {
		RemoveEventElement(m.gestureCommandPtr)
	}
	m.gestureCommandPtr = MakeEventDataPtr(fn)
	CEF().SysCallN(94, m.Instance(), m.gestureCommandPtr)
}
