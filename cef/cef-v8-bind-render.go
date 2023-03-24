//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// 基于IPC的字段数据绑定 - 渲染(子)进程
package cef

import (
	"fmt"
	"github.com/energye/energy/pkgs/json"
)

var bindRender *bindRenderProcess

type bindRenderProcess struct {
	handler *ICefV8Handler
}

func (m *bindRenderProcess) initBindIPC() {
	renderIPC.addCallback(func(channelId int64, data json.JSON) bool {
		if data != nil {
			//messageJSON := data.JSONObject()
			//messageId := messageJSON.GetIntByKey(ipc_id)// messageId: 同步永远是1
			//name := messageJSON.GetStringByKey(ipc_name)
			//argumentList := messageJSON.GetArrayByKey(ipc_argumentList)
			//if name == internalIPCJSExecuteGoSyncEventReplay {
			//	return true
			//}
		}
		return false
	})
}

func (m *bindRenderProcess) makeBind() {
	m.handler = V8HandlerRef.New()
	m.handler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		fmt.Println("v8Handler.Execute", name, renderIPC)
		if renderIPC != nil {

		}
		return false
	})
	//注册js
	var jsCode = `
            let bind;
            if (!bind) {
                bind = {};
            }
            (function () {
				Object.defineProperty(bind, 'myparam', {
					get(){
						native function GetMyParam();
						return GetMyParam();
					},
					set(v){
                    	native function SetMyParam();
						SetMyParam(v);
					}
				});
            })();
`
	registerExtension(internalBind, jsCode, m.handler)
}
