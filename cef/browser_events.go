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
	"bytes"
	"encoding/json"
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/core"
	"github.com/energye/energy/v3/ipc"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/tool/exec"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/keys"
	"net/url"
	"path/filepath"
	"runtime"
	"unsafe"
)

func (m *TBrowser) initBrowserDefaultEvent() {
	logger.Debug("Browser.initBrowserDefaultEvent")

	m.chromium.SetOnProcessMessageReceived(m.chromiumOnProcessMessageReceived)

	m.chromium.SetOnGetResourceHandler(m.chromiumOnGetResourceHandler)
	m.chromium.SetOnResourceLoadComplete(m.chromiumOnResourceLoadComplete)

	m.chromium.SetOnLoadStart(m.chromiumOnLoadStart)
	m.chromium.SetOnLoadEnd(m.chromiumOnLoadEnd)
	m.chromium.SetOnTitleChange(m.chromiumOnTitleChange)

	m.chromium.SetOnBeforeContextMenu(m.chromiumOnBeforeContextMenu)   // ContextMenu
	m.chromium.SetOnContextMenuCommand(m.chromiumOnContextMenuCommand) // ContextMenuCommand

	m.chromium.SetOnBeforeDownload(m.chromiumOnBeforeDownload)

	m.chromium.SetOnDragEnter(m.chromiumOnDragEnter)
	//m.chromium.SetOnStartDragging(m.chromiumOnStartDragging)

	m.chromium.SetOnKeyEvent(m.chromiumOnKeyEvent)

	// new tab or popup browser
	m.chromium.SetOnOpenUrlFromTab(m.chromiumOnOpenUrlFromTab)
	m.chromium.SetOnBeforePopup(m.chromiumOnBeforePopup)
}

func (m *TBrowser) chromiumOnProcessMessageReceived(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, sourceProcess cefTypes.TCefProcessId,
	message cef.ICefProcessMessage, outResult *bool) {
	name := message.GetName()
	args := message.GetArgumentList()
	logger.Debug("Chromium.OnProcessMessageReceived name:", name)
	defer func() {
		message.Release()
		args.Release()
	}()
	if ok := m.doProcessMessagePostMessage(name, args); ok {
		*outResult = true
	} else if ok := m.doProcessMessageExecuteScriptResult(name, args); ok {
		*outResult = true
	}
}

func (m *TBrowser) chromiumOnGetResourceHandler(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
	resourceHandler *cef.IEngResourceHandler) {
	logger.Debug("Chromium.OnGetResourceHandler ProcessType:", ProcessType(GApplication.ProcessType()))

	var (
		uri          = request.GetUrl()
		reqUrl, err  = url.Parse(uri)
		path, method string
		resource     string
		handle       bool
		data         []byte
	)

	if err == nil {
		if GApplication == nil || GApplication.LocalLoad == nil || reqUrl.Scheme != GApplication.LocalLoad.Scheme {
			return
		}
		if m.resourceHandlerList == nil {
			m.resourceHandlerList = make(map[string]*source)
		}
		if m.onResourceRequest != nil {
			header := make(map[string]string)
			path = reqUrl.Path
			method = request.GetMethod()
			headerMap := cef.NewStringMultimapOwn()
			intfHeaderMap := cef.AsCefStringMultimapOwn(headerMap.AsIntfStringMultimap())
			request.GetHeaderMap(intfHeaderMap)
			for i := 0; i < int(intfHeaderMap.GetSize()); i++ {
				key := intfHeaderMap.GetKey(uint32(i))
				value := intfHeaderMap.GetValue(uint32(i))
				header[key] = value
			}
			intfHeaderMap.Release()
			resource, handle = m.onResourceRequest(uri, path, method, header)
		}
		if handle && resource != "" {
			data = []byte(resource)
		}
		src, err := makeSource(browser, frame, request)
		if err != nil {
			logger.Error("Chromium.OnGetResourceHandler makeSource:", err.Error())
			return
		}
		src.data = data
		*resourceHandler = cef.AsEngResourceHandler(src.resourceHandler.AsIntfResourceHandler())
		m.resourceHandlerList[uri] = src
	}
}

func (m *TBrowser) chromiumOnResourceLoadComplete(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, request cef.ICefRequest,
	response cef.ICefResponse, status cefTypes.TCefUrlRequestStatus, receivedContentLength int64) {
	uri := request.GetUrl()
	logger.Debug("Chromium.OnResourceLoadComplete uri:", uri, "status:", status, "receivedContentLength:", receivedContentLength)
	if src, ok := m.resourceHandlerList[uri]; ok {
		src.resourceHandler.Free()
		delete(m.resourceHandlerList, uri)
	}
}

func (m *TBrowser) chromiumOnBeforeContextMenu(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams, model cef.ICefMenuModel) {
	logger.Debug("Chromium.OnBeforeContextMenu")
	menuItemClear := func(menuItems cef.ICefMenuModel) {
		menuItems.Clear()
	}
	if m.window != nil && m.window.Options().DisableContextMenu {
		menuItemClear(model)
		return
	}
	if m.onContextMenu != nil {
		nextMenuId := cefTypes.MENU_ID_USER_FIRST
		var nextContextMenuCommandId = func() int32 {
			nextMenuId++
			return int32(nextMenuId)
		}
		var add func(text string, kind core.TContextMenuKind, menuItems cef.ICefMenuModel) (*core.TContextMenuItem, int32)
		add = func(text string, kind core.TContextMenuKind, menuItems cef.ICefMenuModel) (*core.TContextMenuItem, int32) {
			var subContextMenu cef.ICefMenuModel
			commandId := nextContextMenuCommandId()
			switch kind {
			case core.CmkCommand:
				menuItems.AddItem(commandId, text)
			case core.CmkSub:
				subContextMenu = menuItems.AddSubMenu(commandId, text)
			case core.CmkSeparator:
				menuItems.AddSeparator()
			default:
				return nil, 0
			}
			childContextMenu := &core.TContextMenuItem{
				Clear: func() {
					menuItemClear(subContextMenu)
				},
				Add: func(text string, kind core.TContextMenuKind) (*core.TContextMenuItem, int32) {
					newMenuItem, newCommandId := add(text, kind, subContextMenu)
					return newMenuItem, newCommandId
				}}
			return childContextMenu, commandId
		}
		contextMenu := &core.TContextMenuItem{
			Clear: func() {
				menuItemClear(model)
			},
			Add: func(text string, kind core.TContextMenuKind) (*core.TContextMenuItem, int32) {
				return add(text, kind, model)
			},
		}
		m.onContextMenu(contextMenu)
	}
}

func (m *TBrowser) chromiumOnContextMenuCommand(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, params cef.ICefContextMenuParams,
	commandId int32, eventFlags cefTypes.TCefEventFlags, outResult *bool) {
	logger.Debug("Chromium.OnContextMenuCommand", "commandId:", commandId, *outResult)
	if m.onContextMenuCommand != nil {
		m.onContextMenuCommand(commandId, outResult)
	}
}

func (m *TBrowser) chromiumOnKeyEvent(sender lcl.IObject, browser cef.ICefBrowser, event cef.TCefKeyEvent, osEvent cefTypes.TCefEventHandle, outResult *bool) {
	if event.WindowsKeyCode == keys.VkF12 {
		m.chromium.ShowDevToolsWithPointWinControl(types.Point(0, 0), nil)
		*outResult = true
	}
}

func (m *TBrowser) chromiumOnOpenUrlFromTab(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, targetUrl string,
	targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, outResult *bool) {
	logger.Debug("Chromium.OnOpenUrlFromTab", "targetUrl:", targetUrl)
	*outResult = true
}

func (m *TBrowser) chromiumOnAdapterBeforePopup(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, popupId int32, targetUrl string,
	targetFrameName string, targetDisposition cefTypes.TCefWindowOpenDisposition, userGesture bool, popupFeatures cef.TCefPopupFeatures,
	windowInfo *cef.TCefWindowInfo, client *cef.IEngClient, settings *cef.TCefBrowserSettings, extraInfo *cef.ICefDictionaryValue,
	noJavascriptAccess *bool, result *bool) {
	logger.Debug("Chromium.OnAdapterBeforePopup", "popupId:", popupId, "targetUrl:", targetUrl)
	*result = true
	var handle bool
	if m.onPopupWindow != nil {
		handle = m.onPopupWindow(targetUrl)
	}
	if m.kind == bkEmbedded && !handle && m.window != nil {
		options := m.window.Options()
		if options.AutoPopupWindow && gPrePopupWindow != nil {
			lcl.RunOnMainThreadAsync(func(id uint32) {
				gPrePopupWindow.Browser().Chromium().SetDefaultUrl(targetUrl)
				gPrePopupWindow.Show()
				gPrePopupWindow = nil
			})
		}
	} else if GApplication.Options.AutoPopupWindow && m.kind == bkViews && !handle {
		nViewsBrowser := NewViewsBrowser(nil)
		nViewsBrowser.SetDefaultURL(targetUrl)
		nViewsBrowser.CreateBrowser()
		nViewsBrowser.Show()
	}
}

func (m *TBrowser) chromiumOnAdapterBeforeDownload(sender lcl.IObject, browser cef.ICefBrowser, downloadItem cef.ICefDownloadItem, suggestedName string,
	callback cef.ICefBeforeDownloadCallback, result *bool) {
	callback.Cont(filepath.Join(exec.AppDir(), suggestedName), true)
	*result = true
}

func (m *TBrowser) chromiumOnDragEnter(sender lcl.IObject, browser cef.ICefBrowser, dragData cef.ICefDragData, mask cefTypes.TCefDragOperations,
	outResult *bool) {
	logger.Debug("Chromium.OnDragEnter", browser.GetIdentifier())
	if dragData.IsFile() {
		m.dragFilePathCache = make(map[string]string) // clear and new
		names := lcl.AsStrings(lcl.NewStringList())
		dragData.GetFileNames(&names)
		for i := int32(0); i < names.Count(); i++ {
			dragFilePath := names.Strings(i)
			_, fileName := filepath.Split(dragFilePath)
			m.dragFilePathCache[fileName] = dragFilePath
		}
		names.Free()
	}
}

func (m *TBrowser) chromiumOnLoadStart(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, transitionType cefTypes.TCefTransitionType) {
	m.loadingURL = frame.GetUrl()
	m.loadingState = core.LcStart
	logger.Debug("Chromium.OnLoadStart", m.loadingURL)
	if m.onLoadChange != nil {
		m.onLoadChange(m.loadingURL, m.loadingTitle, m.loadingState)
	}
}

func (m *TBrowser) chromiumOnLoadEnd(sender lcl.IObject, browser cef.ICefBrowser, frame cef.ICefFrame, httpStatusCode int32) {
	logger.Debug("Chromium.OnLoadEnd")
	m.loadingState = core.LcFinish
	m.createEnergyJavasScript()
	if m.onLoadChange != nil {
		m.onLoadChange(m.loadingURL, m.loadingTitle, m.loadingState)
	}
}

func (m *TBrowser) chromiumOnTitleChange(sender lcl.IObject, browser cef.ICefBrowser, title string) {
	logger.Debug("Chromium.OnTitleChange title:", title)
	m.loadingTitle = title
	if m.kind == bkEmbedded {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			if m.window.Caption() == "" {
				m.window.SetCaption(title)
			}
		})
	}
}

func (m *TBrowser) doDragDrop(message ipc.ProcessMessage, args cef.ICefListValue) {
	data, ok := message.Data.(map[string]any)
	if !ok {
		return
	}
	dataType := core.TDragType(tool.ToInt(data["type"]))
	x := int32(tool.ToInt(data["x"]))
	y := int32(tool.ToInt(data["y"]))
	switch message.Type {
	case ipc.MT_DRAG_DROP_ENTER:
		if m.onDragEnter != nil {
			m.onDragEnter(dataType, x, y)
		}
	case ipc.MT_DRAG_DROP_LEAVE:
		if m.onDragLeave != nil {
			m.onDragLeave()
		}
	case ipc.MT_DRAG_DROP_OVER:
		if m.onDragOver != nil {
			dragData := &core.TDragData{Type: dataType}
			if dataType == core.DragTypeData {
				dragData.Data = []byte(data["text"].(string))
			} else if dataType == core.DragTypeFile {
				objectsDataBin := args.GetBinary(1)
				defer func() {
					objectsDataBin.Release()
				}()
				objectsDataBytes := make([]byte, int(objectsDataBin.GetSize()))
				objectsDataBin.GetData(uintptr(unsafe.Pointer(&objectsDataBytes[0])), objectsDataBin.GetSize(), 0)

				var (
					objectFiles []tObjectFile
					files       []string
				)
				err := json.Unmarshal(objectsDataBytes, &objectFiles)
				if err == nil && m.dragFilePathCache != nil {
					for _, file := range objectFiles {
						if filePath, ok := m.dragFilePathCache[file.Name]; ok {
							files = append(files, filePath)
						}
					}
				}

				dragData.Filenames = files
				m.dragFilePathCache = nil // clear
			}
			m.onDragOver(dragData, x, y)
		}
	}
}

func (m *TBrowser) doProcessMessagePostMessage(name string, arguments cef.ICefListValue) bool {
	if name == internalPostMessageName {
		var handle bool
		messageData := ""
		dataBin := arguments.GetBinary(0)
		defer func() {
			dataBin.Release()
		}()
		messageDataBytes := make([]byte, int(dataBin.GetSize()))
		dataBin.GetData(uintptr(unsafe.Pointer(&messageDataBytes[0])), dataBin.GetSize(), 0)
		messageData = string(messageDataBytes)
		if m.messageReceivedDelegate != nil {
			// ipc message
			var pMessage ipc.ProcessMessage
			err := json.Unmarshal(messageDataBytes, &pMessage)
			if err == nil {
				switch pMessage.Type {
				case ipc.MT_READY:
					// ipc ready
					handle = true
				case ipc.MT_EVENT_GO_EMIT, ipc.MT_EVENT_JS_EMIT, ipc.MT_EVENT_GO_EMIT_CALLBACK, ipc.MT_EVENT_JS_EMIT_CALLBACK:
					// ipc on, emit event
					handle = m.messageReceivedDelegate.Received(m.BrowserId(), &pMessage)
				case ipc.MT_DRAG_MOVE, ipc.MT_DRAG_DOWN, ipc.MT_DRAG_UP, ipc.MT_DRAG_DBLCLICK:
					// ipc drag window
					handle = m.drag(pMessage)
				case ipc.MT_DRAG_RESIZE:
					// border drag resize
					ht := pMessage.Data.(string)
					handle = m.resize(ht)
				case ipc.MT_DRAG_BORDER_WMSZ:
				case ipc.MT_DRAG_DROP_ENTER, ipc.MT_DRAG_DROP_LEAVE, ipc.MT_DRAG_DROP_OVER:
					m.doDragDrop(pMessage, arguments)
				}
			} else {
				println("MessageReceived-ERROR：", err.Error())
			}
		}
		logger.Debug("Chromium.OnProcessMessageReceived messageData:", messageData)
		if !handle && m.onProcessMessage != nil {
			m.onProcessMessage(messageData)
		}
		return true
	}
	return false
}

func (m *TBrowser) doProcessMessageExecuteScriptResult(name string, arguments cef.ICefListValue) bool {
	if name == internalExecuteScriptResultName {
		dataBin := arguments.GetBinary(0)
		defer func() {
			dataBin.Release()
		}()
		messageDataBytes := make([]byte, int(dataBin.GetSize()))
		dataBin.GetData(uintptr(unsafe.Pointer(&messageDataBytes[0])), dataBin.GetSize(), 0)
		executeScriptResult := tExecuteScriptResultMessage{}
		_ = json.Unmarshal(messageDataBytes, &executeScriptResult)
		executionID := executeScriptResult.Id
		if callback, ok := m.executeScriptCallback.Load(executionID); ok {
			m.executeScriptCallback.Delete(executionID)
			callback.(core.TOnEvaluateScriptCallbackEvent)(executeScriptResult.Data, executeScriptResult.Error)
		}
		return true
	}
	return false
}

func (m *TBrowser) createEnergyJavasScript() {
	jsCode := &bytes.Buffer{}
	var envJS = func(json string) {
		jsCode.WriteString(`window.energy.setOptionsEnv(`)
		jsCode.WriteString(json)
		jsCode.WriteString(`);`)
	}
	optionsJSON, err := json.Marshal(GApplication.Options)
	if err == nil {
		envJS(string(optionsJSON))
	}
	browser := make(map[string]any)
	browser["id"] = m.BrowserId()
	env := make(map[string]any)
	env["frameWidth"] = frameWidth
	env["frameHeight"] = frameHeight
	env["frameCorner"] = frameCorner
	env["os"] = runtime.GOOS
	env["browser"] = browser
	envJSON, err := json.Marshal(env)
	if err == nil {
		envJS(string(envJSON))
	}
	m.ExecuteScript(jsCode.String())
	m.ExecuteScript(`window.energy.drag.setup();`)
}
