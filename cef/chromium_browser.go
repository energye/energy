//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/pkgs/assetserve"
	et "github.com/energye/energy/v2/types"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

// ICEFChromiumBrowser
//
//	CEFChromium浏览器接口
type ICEFChromiumBrowser interface {
	SetCreateBrowserExtraInfo(windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) //
	CreateBrowser()                                                                                           // 创建浏览器
	Chromium() IChromium                                                                                      // 返回 chromium
	WindowParent() ICEFWindowParent                                                                           // 返回 chromium window 组件
	IsCreated() bool                                                                                          // 创建浏览器是否成功
	SetSelfWindow(selfWindow IBrowserWindow)                                                                  // 设置当前Chromium自己所属的window对象
	// RegisterDefaultEvent 注册默认Chromium事件
	//  如果希望所有默认实现的事件都被注册成功，在使用 NewChromiumBrowser 创建时 owner 参数非 IBrowserWindow
	//  需要在创建Browser之前设置当前Chromium的所属Window, 使用 SetSelfWindow 函数
	RegisterDefaultEvent()
	// RegisterDefaultPopupEvent 注册弹出子窗口事件
	//  如果希望所有默认实现的事件都被注册成功，在使用 NewChromiumBrowser 创建时 owner 参数非 IBrowserWindow
	//  需要在创建Browser之前设置当前Chromium的所属Window, 使用 SetSelfWindow 函数
	RegisterDefaultPopupEvent()
	BroderDirectionAdjustments() et.BroderDirectionAdjustments       // 返回可以调整窗口大小的边框方向, 默认所有方向
	SetBroderDirectionAdjustments(val et.BroderDirectionAdjustments) // 设置可以调整窗口大小的边框方向, 默认所有方向
	Rgn() *et.HRGN                                                   // Rgn 返回区域
	Regions() *TCefDraggableRegions                                  // Regions 返回区域坐标
	FreeRgn()                                                        // FreeRgn 释放掉rgn
	FreeRegions()                                                    // FreeRegions 释放掉regions
}

// TCEFChromiumBrowser
//
//	CEFChromium浏览器包装结构
type TCEFChromiumBrowser struct {
	window                     IBrowserWindow                // chromium 所属窗口
	chromium                   IChromium                     // chromium
	windowParent               ICEFWindowParent              // chromium window 组件
	isCreated                  bool                          // chromium browser is created
	createTimer                *lcl.TTimer                   // loop check and create chromium browser
	windowName                 string                        //
	context                    *ICefRequestContext           //
	extraInfo                  *ICefDictionaryValue          //
	broderDirectionAdjustments et.BroderDirectionAdjustments //可以调整窗口大小的边框方向, 默认所有方向
	regions                    *TCefDraggableRegions         //窗口内html拖拽区域
	rgn                        *et.HRGN                      //
}

// NewChromiumBrowser
//
//	初始创建一个 chromium 浏览器
//	当 owner 参数是 IBrowserWindow 类型时，将此参数设置为当前chromium所属的窗口
func NewChromiumBrowser(owner lcl.IWinControl, config *TCefChromiumConfig) ICEFChromiumBrowser {
	var m = new(TCEFChromiumBrowser)
	// owner是窗口直接设置到selfWindow
	if window, ok := owner.(IBrowserWindow); ok {
		m.window = window
	}
	m.chromium = NewChromium(owner, config)
	m.windowParent = NewCEFWindowParent(owner)
	m.windowParent.SetParent(owner)
	m.windowParent.SetChromium(m.chromium, 0)
	m.windowParent.SetWidth(100)
	m.windowParent.SetHeight(100)
	m.windowParent.SetAlign(types.AlNone)
	m.createTimer = lcl.NewTimer(owner)
	m.createTimer.SetInterval(200)
	m.createTimer.SetEnabled(false)
	m.createTimer.SetOnTimer(m.checkAndCreateBrowser)
	m.broderDirectionAdjustments = et.NewSet(et.BdaTop, et.BdaTopRight, et.BdaRight, et.BdaBottomRight, et.BdaBottom, et.BdaBottomLeft, et.BdaLeft, et.BdaTopLeft)
	return m
}

// checkAndCreateBrowser
//
//	创建浏览器
//	创建时如果未创建成功, 使用定时器创建直到成功
func (m *TCEFChromiumBrowser) checkAndCreateBrowser(sender lcl.IObject) {
	if m.isCreated || m.chromium == nil || m.createTimer == nil {
		return
	}
	m.createTimer.SetEnabled(false)
	if m.isCreated { // 成功创建 释放定时器
		m.createTimer.Free()
		m.createTimer = nil
		return
	}
	m.chromium.Initialized()
	m.isCreated = m.chromium.CreateBrowser(m.windowParent, m.windowName, m.context, m.extraInfo)
	if !m.isCreated {
		m.createTimer.SetEnabled(true)
	} else {
		m.windowParent.UpdateSize()
	}
}

func (m *TCEFChromiumBrowser) SetCreateBrowserExtraInfo(windowName string, context *ICefRequestContext, extraInfo *ICefDictionaryValue) {
	m.windowName = windowName
	m.context = context
	m.extraInfo = extraInfo
}

// BroderDirectionAdjustments 可以调整窗口大小的边框方向, 默认所有方向
func (m *TCEFChromiumBrowser) BroderDirectionAdjustments() et.BroderDirectionAdjustments {
	return m.broderDirectionAdjustments
}

// SetBroderDirectionAdjustments 设置可以调整窗口大小的边框方向, 默认所有方向
func (m *TCEFChromiumBrowser) SetBroderDirectionAdjustments(val et.BroderDirectionAdjustments) {
	m.broderDirectionAdjustments = val
}

// CreateBrowser
//
//	创建浏览器
//	创建时如果未创建成功, 使用定时任务创建直到成功
func (m *TCEFChromiumBrowser) CreateBrowser() {
	m.checkAndCreateBrowser(nil)
}

// Chromium
//
//	返回 chromium
func (m *TCEFChromiumBrowser) Chromium() IChromium {
	return m.chromium
}

// WindowParent
//
//	返回 chromium window 组件
func (m *TCEFChromiumBrowser) WindowParent() ICEFWindowParent {
	return m.windowParent
}

// IsCreated
//
//	创建浏览器是否成功
func (m *TCEFChromiumBrowser) IsCreated() bool {
	return m.isCreated
}

// SetSelfWindow
//
// 设置当前Chromium自己所属的window对象
func (m *TCEFChromiumBrowser) SetSelfWindow(selfWindow IBrowserWindow) {
	m.window = selfWindow
}

func (m *TCEFChromiumBrowser) RegisterDefaultEvent() {
	var bwEvent = BrowserWindow.browserEvent
	m.Chromium().SetOnProcessMessageReceived(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) bool {
		if bwEvent.onProcessMessageReceived != nil {
			return bwEvent.onProcessMessageReceived(sender, browser, frame, sourceProcess, message, m.window)
		}
		return false
	})
	m.Chromium().SetOnBeforeResourceLoad(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, callback *ICefCallback, result *consts.TCefReturnValue) {
		if assetserve.AssetsServerHeaderKeyValue != "" {
			if application.IsSpecVer49() {
				headerMap := request.GetHeaderMap()
				headerMap.Append(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue)
				request.SetHeaderMap(headerMap)
				headerMap.Free()
			} else {
				request.SetHeaderByName(assetserve.AssetsServerHeaderKeyName, assetserve.AssetsServerHeaderKeyValue, true)
			}
		}
		if bwEvent.onBeforeResourceLoad != nil {
			bwEvent.onBeforeResourceLoad(sender, browser, frame, request, callback, result, m.window)
		}
	})
	//事件可以被覆盖
	m.Chromium().SetOnBeforeDownload(func(sender lcl.IObject, browser *ICefBrowser, beforeDownloadItem *ICefDownloadItem, suggestedName string, callback *ICefBeforeDownloadCallback) {
		if bwEvent.onBeforeDownload != nil {
			bwEvent.onBeforeDownload(sender, browser, beforeDownloadItem, suggestedName, callback, m.window)
		} else {
			// 默认保存到当前执行文件所在目录
			callback.Cont(consts.ExeDir+consts.Separator+suggestedName, true)
		}
	})
	m.Chromium().SetOnLoadEnd(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, httpStatusCode int32) {
		if bwEvent.onLoadEnd != nil {
			bwEvent.onLoadEnd(sender, browser, frame, httpStatusCode, m.window)
		}
	})
	if localLoadRes.enable() {
		m.Chromium().SetOnGetResourceHandler(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest) (resourceHandler *ICefResourceHandler) {
			//var flag bool
			if bwEvent.onGetResourceHandler != nil {
				resourceHandler, _ = bwEvent.onGetResourceHandler(sender, browser, frame, request, m.window)
			}
			//if !flag {
			//	resourceHandler = localLoadRes.getResourceHandler(browser, frame, request)
			//}
			return
		})
	}
	if m.window != nil {
		m.Chromium().SetOnBeforeContextMenu(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, model *ICefMenuModel) {
			var flag bool
			if bwEvent.onBeforeContextMenu != nil {
				flag = bwEvent.onBeforeContextMenu(sender, browser, frame, params, model, m.window)
			}
			if !flag {
				chromiumOnBeforeContextMenu(m.window, browser, frame, params, model)
			}
		})
		m.Chromium().SetOnContextMenuCommand(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, params *ICefContextMenuParams, commandId consts.MenuId, eventFlags uint32) bool {
			var result bool
			if bwEvent.onContextMenuCommand != nil {
				result = bwEvent.onContextMenuCommand(sender, browser, frame, params, commandId, eventFlags, m.window)
			}
			if !result {
				result = chromiumOnContextMenuCommand(m.window, m, browser, frame, params, commandId, eventFlags)
			}
			return result
		})
		m.Chromium().SetOnAfterCreated(func(sender lcl.IObject, browser *ICefBrowser) {
			var flag bool
			if bwEvent.onAfterCreated != nil {
				flag = bwEvent.onAfterCreated(sender, browser, m.window)
			}
			if !flag {
				chromiumOnAfterCreate(m.window, browser)
			}
		})
		//事件可以被覆盖
		m.Chromium().SetOnKeyEvent(func(sender lcl.IObject, browser *ICefBrowser, event *TCefKeyEvent, osEvent *consts.TCefEventHandle, result *bool) {
			if bwEvent.onKeyEvent != nil {
				bwEvent.onKeyEvent(sender, browser, event, osEvent, m.window, result)
			}
			if !*result {
				if m.window == nil || m.window.WindowType() == consts.WT_DEV_TOOLS || m.window.WindowType() == consts.WT_VIEW_SOURCE {
					return
				}
				if m.Chromium().Config().EnableDevTools() {
					if event.WindowsKeyCode == consts.VkF12 && event.Kind == consts.KEYEVENT_RAW_KEYDOWN {
						browser.ShowDevTools(m.window, m)
						*result = true
					} else if event.WindowsKeyCode == consts.VkF12 && event.Kind == consts.KEYEVENT_KEYUP {
						*result = true
					}
				}
				if KeyAccelerator.accelerator(browser, event, result) {
					return
				}
			}
		})
		m.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, request *ICefRequest, userGesture, isRedirect bool) bool {
			m.WindowParent().UpdateSize()
			chromiumOnBeforeBrowser(m.window, browser, frame, request) // default impl
			if bwEvent.onBeforeBrowser != nil {
				return bwEvent.onBeforeBrowser(sender, browser, frame, request, userGesture, isRedirect, m.window)
			}
			return false
		})
		m.Chromium().SetOnTitleChange(func(sender lcl.IObject, browser *ICefBrowser, title string) {
			updateBrowserDevTools(m.window, browser, title)
			if bwEvent.onTitleChange != nil {
				bwEvent.onTitleChange(sender, browser, title, m.window)
			}
		})
		m.Chromium().SetOnDragEnter(func(sender lcl.IObject, browser *ICefBrowser, dragData *ICefDragData, mask consts.TCefDragOperations, result *bool) {
			*result = !m.window.WindowProperty().EnableDragFile
			if bwEvent.onDragEnter != nil {
				bwEvent.onDragEnter(sender, browser, dragData, mask, m.window, result)
			}
		})
		if m.window.WindowProperty().EnableWebkitAppRegion {
			m.Chromium().SetOnDraggableRegionsChanged(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, regions *TCefDraggableRegions) {
				if bwEvent.onDraggableRegionsChanged != nil {
					bwEvent.onDraggableRegionsChanged(sender, browser, frame, regions, m.window)
				}
				if m.window.IsLCL() {
					m.regions = regions
					m.setDraggableRegions()
				} else {
					m.window.AsViewsFrameworkBrowserWindow().BrowserWindow().regions = regions
					m.window.AsViewsFrameworkBrowserWindow().WindowComponent().SetDraggableRegions(regions.Regions())
				}
			})
		}
	}
	// 注册windows下CompMsg事件, CEF 组件消息事件
	m.registerWindowsCompMsgEvent()
}

func (m *TCEFChromiumBrowser) RegisterDefaultPopupEvent() {
	var bwEvent = BrowserWindow.browserEvent
	if m.window != nil {
		if m.window.IsViewsFramework() {
			isMain := m.window.WindowType() == consts.WT_MAIN_BROWSER
			if !isMain {
				// 子窗口关闭流程
				m.Chromium().SetOnBeforeClose(func(sender lcl.IObject, browser *ICefBrowser) {
					var flag bool
					if bwEvent.onBeforeClose != nil {
						flag = bwEvent.onBeforeClose(sender, browser, m.window)
					}
					if !flag {
						chromiumOnBeforeClose(m.window, browser)
						m.window.AsViewsFrameworkBrowserWindow().BrowserWindow().TryCloseWindowAndTerminate()
					}
				})
				m.Chromium().SetOnClose(func(sender lcl.IObject, browser *ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
					if bwEvent.onClose != nil {
						bwEvent.onClose(sender, browser, aAction, m.window)
					}
				})
			}
		}
		m.Chromium().SetOnOpenUrlFromTab(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, targetUrl string, targetDisposition consts.TCefWindowOpenDisposition, userGesture bool) bool {
			if !m.Chromium().Config().EnableOpenUrlTab() {
				return true
			}
			return false
		})
		m.Chromium().SetOnBeforePopup(func(sender lcl.IObject, browser *ICefBrowser, frame *ICefFrame, beforePopupInfo *BeforePopupInfo, popupFeatures *TCefPopupFeatures, windowInfo *TCefWindowInfo, client *ICefClient, settings *TCefBrowserSettings, resultExtraInfo *ICefDictionaryValue, noJavascriptAccess *bool) bool {
			if !m.Chromium().Config().EnableWindowPopup() {
				return true
			}
			return m.window.doBeforePopup(sender, browser, frame, beforePopupInfo, popupFeatures, windowInfo, client, settings, resultExtraInfo, noJavascriptAccess)
		})
	}
}

// setDraggableRegions
// 每一次拖拽区域改变都需要重新设置
func (m *TCEFChromiumBrowser) setDraggableRegions() {
	var scp float32
	// Windows 10 版本 1607 [仅限桌面应用]
	// Windows Server 2016 [仅限桌面应用]
	// 可动态调整
	dpi, err := winapi.GetDpiForWindow(et.HWND(m.window.Handle()))
	if err == nil {
		scp = float32(dpi) / 96.0
	} else {
		// 使用默认的，但不能动态调整
		scp = winapi.ScalePercent()
	}
	//在主线程中运行
	RunOnMainThread(func() {
		if m.rgn == nil {
			//第一次时创建RGN
			m.rgn = winapi.CreateRectRgn(0, 0, 0, 0)
		} else {
			//每次重置RGN
			winapi.SetRectRgn(m.rgn, 0, 0, 0, 0)
		}
		// 重新根据缩放比计算新的区域位置
		for i := 0; i < m.regions.RegionsCount(); i++ {
			region := m.regions.Region(i)
			x := int32(float32(region.Bounds.X) * scp)
			y := int32(float32(region.Bounds.Y) * scp)
			w := int32(float32(region.Bounds.Width) * scp)
			h := int32(float32(region.Bounds.Height) * scp)
			creRGN := winapi.CreateRectRgn(x, y, x+w, y+h)
			if region.Draggable {
				winapi.CombineRgn(m.rgn, m.rgn, creRGN, consts.RGN_OR)
			} else {
				winapi.CombineRgn(m.rgn, m.rgn, creRGN, consts.RGN_DIFF)
			}
			winapi.DeleteObject(creRGN)
		}
	})
}

// Rgn 返回区域
func (m *TCEFChromiumBrowser) Rgn() *et.HRGN {
	return m.rgn
}

// Regions 返回区域坐标
func (m *TCEFChromiumBrowser) Regions() *TCefDraggableRegions {
	return m.regions
}

// FreeRgn 释放掉rgn
func (m *TCEFChromiumBrowser) FreeRgn() {
	if m.rgn != nil {
		winapi.SetRectRgn(m.rgn, 0, 0, 0, 0)
		winapi.DeleteObject(m.rgn)
		m.rgn.Free()
	}
}

// FreeRegions 释放掉regions
func (m *TCEFChromiumBrowser) FreeRegions() {
	if m.regions != nil {
		m.regions.regions = nil
		m.regions = nil
	}
}
