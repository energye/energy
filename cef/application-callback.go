//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// application event 默认事件实现

package cef

import (
	"fmt"
	"github.com/energye/energy/v2/cef/internal/ipc"
	"github.com/energye/energy/v2/cef/internal/process"
	ipcArgument "github.com/energye/energy/v2/cef/ipc/argument"
	"github.com/energye/energy/v2/consts"
)

// appOnContextCreated 创建应用上下文 - 默认实现
func appOnContextCreated(browser *ICefBrowser, frame *ICefFrame, context *ICefV8Context) {
	process.Current.SetBrowserId(browser.Identifier())                           // 当前进程 browserID
	process.Current.SetFrameId(frame.Identifier())                               // 当前进程 frameId
	ipc.RenderChan().SetRealityChannel(browser.Identifier(), frame.Identifier()) // 设置并更新真实的通道ID
	ipcRender.registerGoSyncReplayEvent()                                        // render ipc
	ipcRender.makeIPC(context)                                                   // render ipc make
	makeProcess(browser, frame, context)                                         // process make

	// 只在LCL窗口中使用自定义窗口拖拽, VF窗口默认已实现
	// 在MacOS中LCL窗口没有有效的消息事件
	//var executeJS = `energyExtension.drag.setEnableDrag(true); energyExtension.drag.setup();`
	//frame.ExecuteJavaScript(executeJS, "", 0)
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	ipcBrowser.registerEvent() // browser ipc
}

// appWebKitInitialized - webkit - 默认实现
func appWebKitInitialized() {
	return
	energyExtensionHandler := V8HandlerRef.New()
	energyExtensionHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		fmt.Println("Execute", name, consts.IsMessageLoop, application.SingleProcess())
		if name == "mouseMove" {
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
			return true
		}
		return false
	})
	var code = `
		let energyExtension;
        if (!energyExtension) {
            energyExtension = {
                drag: {
                    enableDrag: false,
                    shouldDrag: false,
                    cssDragProperty: "--webkit-app-region",
                    cssDragValue: "drag",
                    defaultCursor: null
                },
            };
        }
        (function () {
            energyExtension.drag.war = function (e) {
                let v = window.getComputedStyle(e.target).getPropertyValue(energyExtension.drag.cssDragProperty);
                if (v) {
                    v = v.trim();
                    if (v !== energyExtension.drag.cssDragValue || e.buttons !== 1) {
                        return false;
                    }
                    return e.detail === 1;
                }
                return false;
            }
            energyExtension.drag.mouseMove = function (e) {
                if (!energyExtension.drag.enableDrag && !energyExtension.drag.shouldDrag) {
                    return
                }
                if (energyExtension.drag.shouldDrag) {
                    energyExtension.drag.shouldDrag = false;
					native function mouseMove();
					mouseMove();
                }
            }
            energyExtension.drag.mouseUp = function (e) {
                if (!energyExtension.drag.enableDrag) {
                    return
                }
                energyExtension.drag.shouldDrag = false;
				//document.body.style.cursor = "default";
				native function mouseUp();
				mouseUp();
            }
            energyExtension.drag.mouseDown = function (e) {
                if (!energyExtension.drag.enableDrag || ((e.offsetX > e.target.clientWidth || e.offsetY > e.target.clientHeight))) {
                    return
                }
                if (energyExtension.drag.war(e)) {
					console.log('mouseDown');
                    e.preventDefault();
                    energyExtension.drag.shouldDrag = true;
                    native function mouseDown();
                    mouseDown();
                } else {
                    energyExtension.drag.shouldDrag = false;
                }
            }
            energyExtension.drag.setEnableDrag = function (v) {
				console.log('drag.setEnableDrag', v, energyExtension);
                energyExtension.drag.enableDrag = v;
            }
            energyExtension.drag.setup = function () {
				console.log('drag.setup', energyExtension);
                window.addEventListener("mousemove", energyExtension.drag.mouseMove);
                window.addEventListener("mousedown", energyExtension.drag.mouseDown);
                window.addEventListener("mouseup", energyExtension.drag.mouseUp);
            }
        })();
`
	// 注册 EnergyExtension JS
	RegisterExtension("energyExtension", code, energyExtensionHandler)
}

// renderProcessMessageReceived 渲染进程消息 - 默认实现
func renderProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, sourceProcess consts.CefProcessId, message *ICefProcessMessage) (result bool) {
	if message.Name() == internalIPCJSExecuteGoEventReplay {
		result = ipcRender.ipcJSExecuteGoEventMessageReply(browser, frame, sourceProcess, message)
	} else if message.Name() == internalIPCGoExecuteJSEvent {
		result = ipcRender.ipcGoExecuteJSEvent(browser, frame, sourceProcess, message)
	}
	return
}

// browserProcessMessageReceived 主进程消息 - 默认实现
func browserProcessMessageReceived(browser *ICefBrowser, frame *ICefFrame, message *ICefProcessMessage) (result bool) {
	if message.Name() == internalIPCJSExecuteGoEvent {
		result = ipcBrowser.jsExecuteGoMethodMessage(browser, frame, message)
	}
	return
}
