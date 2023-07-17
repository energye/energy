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
	var executeJS = `
energyExtension.drag.setGOOS('darwin');
energyExtension.drag.setWindowType('LCL');
energyExtension.drag.setEnableDrag(true);
energyExtension.drag.setup();`
	frame.ExecuteJavaScript(executeJS, "", 0)
}

// appMainRunCallback 应用运行 - 默认实现
func appMainRunCallback() {
	ipcBrowser.registerEvent() // browser ipc
}

// appWebKitInitialized - webkit - 默认实现
func appWebKitInitialized() {
	//return
	var (
		dx, dy int32
		mx, my int32
	)
	energyExtensionHandler := V8HandlerRef.New()
	energyExtensionHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		//fmt.Println("Execute", name, application.IsMessageLoop(), application.SingleProcess())
		if name == "mouseUp" {
			dx = 0
			dy = 0
			mx = 0
			my = 0
		} else if name == "mouseDown" {
			if arguments.Size() > 0 {
				point := arguments.Get(0)
				v8ValX := point.getValueByKey("x")
				v8ValY := point.getValueByKey("y")
				dx = v8ValX.GetIntValue()
				dy = v8ValY.GetIntValue()
				v8ValX.Free()
				v8ValY.Free()
				point.Free()
			}
			fmt.Println("down xy:", dx, dy)
		} else if name == "mouseMove" {
			if arguments.Size() > 0 {
				point := arguments.Get(0)
				v8ValX := point.getValueByKey("x")
				v8ValY := point.getValueByKey("y")
				mx = v8ValX.GetIntValue()
				my = v8ValY.GetIntValue()
				v8ValX.Free()
				v8ValY.Free()
				point.Free()
			}
			fmt.Println("move xy:", mx, my)

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
                    defaultCursor: null,
					goos: "windows", // windows, linux, darwin
					windowType: "LCL", // LCL, VF
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
				//console.log('mouseMove:', e);
                if (!energyExtension.drag.enableDrag || !energyExtension.drag.shouldDrag) {
                    return
                }
				if (energyExtension.drag.goos === "windows") {
					energyExtension.drag.shouldDrag = false;
				}
				native function mouseMove();
				mouseMove({x: e.clientX, y: e.clientY});
            }
            energyExtension.drag.mouseUp = function (e) {
                if (!energyExtension.drag.enableDrag || (energyExtension.drag.goos === "darwin" && !energyExtension.drag.shouldDrag)) {
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
                    e.preventDefault();
                    energyExtension.drag.shouldDrag = true;
                    native function mouseDown();
					console.log(e);
                    mouseDown({x:0, y:0});
                } else {
                    energyExtension.drag.shouldDrag = false;
                }
            }
            energyExtension.drag.setEnableDrag = function (v) {
				// macos enable
				if (energyExtension.drag.goos === "darwin"){
                	energyExtension.drag.enableDrag = v;
				}
            }
            energyExtension.drag.setGOOS = function (v) {
                energyExtension.drag.goos = v;
            }
            energyExtension.drag.setWindowType = function (v) {
                energyExtension.drag.windowType = v;
            }
            energyExtension.drag.setup = function () {
				if (!energyExtension.drag.enableDrag) {
					return;
				}
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
