//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows
// +build windows

package cef

import (
	"github.com/energye/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/energye/energy/v2/cef/ipc/argument"
	"strconv"
)

// 窗口拖拽JS扩展
// 在这里执行并启用JS拖拽
func dragExtensionJS(frame *ICefFrame, drag bool) {
	// Windows2种方式,自定义+webkit 只在LCL窗口中使用自定义窗口拖拽, VF窗口默认已实现
	var executeJS = `
energyExtension.drag.setEnableDrag(` + strconv.FormatBool(drag) + `);
energyExtension.drag.setup();`
	frame.ExecuteJavaScript(executeJS, "", 0)
}

// 窗口拖拽JS扩展处理器
//  1. 注册JS扩展到CEF, 注册鼠标事件，通过本地函数在Go里处理鼠标事件
//  2. 通过IPC将鼠标消息发送到主进程，主进程监听到消息处理鼠标事件
//  3. windows 使用win.api实现窗口拖拽
func dragExtensionHandler() {
	energyExtensionHandler := V8HandlerRef.New()
	energyExtensionHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		if name == mouseDown {
			return true
		} else if name == mouseUp {
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragUp},
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
			return true
		} else if name == mouseMove {
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
                if (!energyExtension.drag.enableDrag || !energyExtension.drag.shouldDrag) {
                    return
                }
				energyExtension.drag.shouldDrag = false;
				native function mouseMove();
				mouseMove();
            }
            energyExtension.drag.mouseUp = function (e) {
                if (!energyExtension.drag.enableDrag) {
                    return
                }
                energyExtension.drag.shouldDrag = false;
				if (energyExtension.drag.war(e)) {
                    e.preventDefault();
					native function mouseUp();
					mouseUp();
				}
            }
            energyExtension.drag.mouseDown = function (e) {
                if (!energyExtension.drag.enableDrag || ((e.offsetX > e.target.clientWidth || e.offsetY > e.target.clientHeight))) {
                    return
                }
                if (energyExtension.drag.war(e)) {
                    e.preventDefault();
                    energyExtension.drag.shouldDrag = true;
                    native function mouseDown();
                    mouseDown();
                } else {
                    energyExtension.drag.shouldDrag = false;
                }
            }
            energyExtension.drag.setEnableDrag = function (v) {
				energyExtension.drag.enableDrag = v;
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
	RegisterExtension("energyExtension", code, energyExtensionHandler)
}

func (m *drag) drag() {
	if m.T == dragUp {
		if m.window.IsLCL() {
			m.window.AsLCLBrowserWindow().BrowserWindow().cwcap.canCaption = false
		}
	}
}
