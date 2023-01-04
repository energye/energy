//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/consts"
	"github.com/energye/energy/ipc"
	t "github.com/energye/energy/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// 通用Result bool
type ChromiumEventOnResult func(sender lcl.IObject, aResultOK bool)

// 通用Result float
type ChromiumEventOnResultFloat func(sender lcl.IObject, result float64)

// chromiumEvent Beforebrowser - 主进程执行每创建一个新的浏览器进程都会调用
type ChromiumEventOnBeforeBrowser func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame) bool

// chromiumEvent 地址改变事件
type ChromiumEventOnAddressChange func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, url string)

// chromiumEvent browser TitleChange
type ChromiumEventOnTitleChange func(sender lcl.IObject, browser *ICefBrowser, title string)

// render
type ChromiumEventOnRenderProcessTerminated func(sender lcl.IObject, browser *ICefBrowser, status consts.TCefTerminationStatus)
type ChromiumEventOnRenderCompMsg func(sender lcl.IObject, message types.TMessage, aHandled bool)

// Event CefBrowse
type ChromiumEventOnCefBrowser func(sender lcl.IObject, browser *ICefBrowser)

// mainChromium Create close
type ChromiumEventOnAfterCreated ChromiumEventOnCefBrowser
type ChromiumEventOnBeforeClose ChromiumEventOnCefBrowser
type ChromiumEventOnClose func(sender lcl.IObject, browser *ICefBrowser, aAction *TCefCloseBrowsesAction)

type ChromiumEventOnScrollOffsetChanged func(sender lcl.IObject, browser *ICefBrowser, x, y float64)

// chromiumEvent 加载
type ChromiumEventOnLoadStart func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)
type ChromiumEventOnLoadingStateChange func(sender lcl.IObject, browser *ICefBrowser, isLoading, canGoBack, canGoForward bool)
type ChromiumEventOnLoadingProgressChange func(sender lcl.IObject, browser *ICefBrowser, progress float64)
type ChromiumEventOnLoadError func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, errorCode consts.CEF_NET_ERROR, errorText, failedUrl string)
type ChromiumEventOnLoadEnd func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32)

// chromiumEvent cookie
type ChromiumEventOnCookieSet func(sender lcl.IObject, success bool, ID int32)
type ChromiumEventOnCookiesDeleted func(sender lcl.IObject, numDeleted int32)
type ChromiumEventOnCookiesFlushed func(sender lcl.IObject)
type ChromiumEventOnCookiesVisited func(sender lcl.IObject, cookie *ICefCookie)
type ChromiumEventOnCookieVisitorDestroyed func(sender lcl.IObject, ID int32)

// chromiumEvent context menu
type ChromiumEventOnBeforeContextMenu func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel)
type ChromiumEventOnContextMenuCommand func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32, result *bool)
type ChromiumEventOnContextMenuDismissed func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)

// chromiumEvent 全屏模式
type ChromiumEventOnFullScreenModeChange func(sender lcl.IObject, browser *ICefBrowser, fullscreen bool)

// 下载之前
type ChromiumEventOnBeforeDownload func(sender lcl.IObject, browser *ICefBrowser, beforeDownloadItem *DownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback)

// 下载中
type ChromiumEventOnDownloadUpdated func(sender lcl.IObject, browser *ICefBrowser, downloadItem *DownloadItem, callback *ICefDownloadItemCallback)

// 键盘
type ChromiumEventOnKeyEvent func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, result *bool)

// 资源加载
type ChromiumEventOnBeforeResourceLoad func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue)
type ChromiumEventOnResourceResponse func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse, result *bool)
type ChromiumEventOnResourceRedirect func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse, newUrl *t.TString)
type ChromiumEventOnResourceLoadComplete func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, response *ICefResponse, status consts.TCefUrlRequestStatus, receivedContentLength int64)

// 渲染进程执行每个浏览器上下文时都会调用
type GlobalCEFAppEventOnContextCreated func(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) bool
type GlobalCEFAppEventOnWebKitInitialized func()
type GlobalCEFAppEventOnContextInitialized func()
type GlobalCEFAppEventOnBeforeChildProcessLaunch func(commandLine *TCefCommandLine)
type GlobalCEFAppEventOnBrowserDestroyed func(browser *ICefBrowser)
type GlobalCEFAppEventOnRenderLoadStart func(browser *ICefBrowser, frame *ICefFrame, transitionType consts.TCefTransitionType)
type GlobalCEFAppEventOnRenderLoadEnd func(browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32)
type GlobalCEFAppEventOnRenderLoadError func(browser *ICefBrowser, frame *ICefFrame, errorCode consts.TCefErrorCode, errorText, failedUrl string)
type GlobalCEFAppEventOnRenderLoadingStateChange func(browser *ICefBrowser, frame *ICefFrame, isLoading, canGoBack, canGoForward bool)
type GlobalCEFAppEventOnGetDefaultClient func(client *ICefClient)

// 进程消息接收
type RenderProcessMessageReceived func(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ipc.ICefProcessMessage) bool
type BrowseProcessMessageReceived func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ipc.ICefProcessMessage) bool
type ChromiumEventOnFindResult func(sender lcl.IObject, browser *ICefBrowser, identifier, count int32, selectionRect *TCefRect, activeMatchOrdinal int32, finalUpdate bool)

// frame
type ChromiumEventOnFrameAttached func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, reattached bool)
type ChromiumEventOnFrameCreated func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)
type ChromiumEventOnFrameDetached func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)
type ChromiumEventOnMainFrameChanged func(sender lcl.IObject, browser *ICefBrowser, oldFrame *ICefFrame, newFrame *ICefFrame)

// windowParent popup
type ChromiumEventOnBeforePopup func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, client *ICefClient, noJavascriptAccess *bool) bool

// windowParent popup
type ChromiumEventOnBeforePopupForWindowInfo func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, popupWindow *LCLBrowserWindow, noJavascriptAccess *bool) bool

// windowParent open url from tab
type ChromiumEventOnOpenUrlFromTab func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame)

// window event
type TCloseEvent func(sender lcl.IObject, action *types.TCloseAction) bool
type TNotifyEvent func(sender lcl.IObject) bool
type TCloseQueryEvent func(sender lcl.IObject, canClose *bool) bool

//TCEFWindowComponent
type WindowComponentOnWindowCreated func(sender lcl.IObject, window *ICefWindow)
type WindowComponentOnWindowDestroyed func(sender lcl.IObject, window *ICefWindow)
type WindowComponentOnWindowActivationChanged func(sender lcl.IObject, window *ICefWindow, active bool)
type WindowComponentOnGetParentWindow func(sender lcl.IObject, window *ICefWindow, isMenu, canActivateMenu *bool, aResult *ICefWindow)
type WindowComponentOnGetInitialBounds func(sender lcl.IObject, window *ICefWindow, aResult *TCefRect)
type WindowComponentOnGetInitialShowState func(sender lcl.IObject, window *ICefWindow, aResult *consts.TCefShowState)
type WindowComponentOnIsFrameless func(sender lcl.IObject, window *ICefWindow, aResult *bool)
type WindowComponentOnCanResize func(sender lcl.IObject, window *ICefWindow, aResult *bool)
type WindowComponentOnCanMaximize func(sender lcl.IObject, window *ICefWindow, aResult *bool)
type WindowComponentOnCanMinimize func(sender lcl.IObject, window *ICefWindow, aResult *bool)
type WindowComponentOnCanClose func(sender lcl.IObject, window *ICefWindow, aResult *bool)
type WindowComponentOnAccelerator func(sender lcl.IObject, window *ICefWindow, commandId int32, aResult *bool)
type WindowComponentOnKeyEvent func(sender lcl.IObject, window *ICefWindow, event *TCefKeyEvent, aResult *bool)

//TCEFBrowserViewComponent
type BrowserViewComponentOnBrowserCreated func(sender lcl.IObject, browserView *ICefBrowserView, browser *ICefBrowser)
type BrowserViewComponentOnBrowserDestroyed func(sender lcl.IObject, browserView *ICefBrowserView, browser *ICefBrowser)
type BrowserViewComponentOnGetDelegateForPopupBrowserView func(sender lcl.IObject, browserView *ICefBrowserView, browserSettings *TCefBrowserSettings, client *ICefClient, isDevtools bool, aResult *ICefBrowserViewDelegate)
type BrowserViewComponentOnPopupBrowserViewCreated func(sender lcl.IObject, browserView, popupBrowserView *ICefBrowserView, isDevtools bool, aResult *bool)
type BrowserViewComponentOnGetChromeToolbarType func(sender lcl.IObject, aResult *consts.TCefChromeToolbarType)
