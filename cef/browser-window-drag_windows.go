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

package cef

func dragExtensionJS(browser *ICefBrowser, frame *ICefFrame) {
	// Windows2种方式,自定义+webkit 只在LCL窗口中使用自定义窗口拖拽, VF窗口默认已实现
	var executeJS = `
energyExtension.drag.setEnableDrag(true);
energyExtension.drag.setup();`
	frame.ExecuteJavaScript(executeJS, "", 0)
}

func dragExtensionHandler() {
	energyExtensionHandler := V8HandlerRef.New()
	energyExtensionHandler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		//fmt.Println("Execute", name, application.IsMessageLoop(), application.SingleProcess())
		if name == mouseUp {
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
				Data: &drag{T: dragUp},
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
		} else if name == mouseDown {
			message := &ipcArgument.List{
				Id:   -1,
				BId:  ipc.RenderChan().BrowserId(),
				Name: internalIPCDRAG,
			}
			ipc.RenderChan().IPC().Send(message.Bytes())
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
	// 注册 EnergyExtension JS
	RegisterExtension("energyExtension", code, energyExtensionHandler)
}

func (m *drag) drag() {

}
