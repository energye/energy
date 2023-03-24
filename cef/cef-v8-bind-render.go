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
)

var bindRender *bindRenderProcess

type bindRenderProcess struct {
	handler *ICefV8Handler
}

func (m *bindRenderProcess) makeBind() {
	m.handler = V8HandlerRef.New()
	m.handler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		fmt.Println("v8Handler.Execute", name, renderIPC)
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
						//return ipc.emitSync("testEmitSync", ["同步参数", 1, 2, 3, ["aaaa", "bbb", 6666]]);
						return GetMyParam();
					},
					set(v){
                    	native function SetMyParam();
						SetMyParam(v);
					}
				});
            })();
`
	RegisterExtension("v8/bind", jsCode, m.handler)
}
