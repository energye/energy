//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin
// +build darwin

package cef

import (
	"github.com/energye/energy/v2/cef/internal/ipc"
	ipcArgument "github.com/energye/energy/v2/cef/ipc/argument"
	"strconv"
)

// 窗口拖拽JS扩展
// 在这里执行并启用JS拖拽
func dragExtensionJS(frame *ICefFrame, drag bool) {
	// MacOS只在LCL窗口中使用自定义窗口拖拽, VF窗口默认已实现
	// 在MacOS中LCL窗口没有有效的消息事件
	var executeJS = `
energyExtension.drag.setEnableDrag(` + strconv.FormatBool(drag) + `);
energyExtension.drag.setup();`
	frame.ExecuteJavaScript(executeJS, "", 0)
}

// 窗口拖拽JS扩展处理器
//  1. 注册JS扩展到CEF, 注册鼠标事件，通过本地函数在Go里处理鼠标事件
//  2. 通过IPC将鼠标消息发送到主进程，主进程监听到消息处理鼠标事件
//  3. macos 使用窗口坐标实现窗口拖拽
func dragExtensionHandler() {
	energyExtensionHandler := V8HandlerRef.New()
	energyExtensionHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		if name == mouseUp {
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragUp},
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
			return true
		} else if name == mouseDown {
			var dx, dy int32
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
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragDown, X: dx, Y: dy},
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
			return true
		} else if name == mouseMove {
			var mx, my int32
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
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragMove, X: mx, Y: my},
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
				native function mouseMove();
				mouseMove({x: e.screenX, y: e.screenY});
            }
            energyExtension.drag.mouseUp = function (e) {
                if (!energyExtension.drag.enableDrag || !energyExtension.drag.shouldDrag) {
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
                    mouseDown({x: e.screenX, y: e.screenY});
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
			bw := m.window.AsLCLBrowserWindow().BrowserWindow()
			bw.drag = nil
			m.window = nil
		}
	} else if m.T == dragDown {
		m.dx = m.X
		m.dy = m.Y
		point := m.window.Point()
		m.wx = point.X
		m.wy = point.Y
	} else if m.T == dragMove {
		m.mx = m.X
		m.my = m.Y
		if m.window.IsLCL() {
			x := m.wx + (m.mx - m.dx)
			y := m.wy + (m.my - m.dy)
			m.window.SetPoint(x, y)
		}
	}
}
