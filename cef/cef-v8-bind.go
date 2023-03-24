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

import "fmt"

const (
	internalBind = "bind"
)

var v8bind *v8BindHandler

type v8BindHandler struct {
	handler *ICefV8Handler
}

// bindInit 初始化
func bindInit() {
	isSingleProcess := application.SingleProcess()
	if isSingleProcess {

	} else {

	}
	v8bind = &v8BindHandler{}
}

func (m *v8BindHandler) makeBind() {
	fmt.Println("SetOnWebKitInitialized")
	v8Handler := V8HandlerRef.New()
	v8Handler.Execute(func(name string, object *ICefV8Value, arguments *TCefV8ValueArray, retVal *ResultV8Value, exception *ResultString) bool {
		fmt.Println("v8Handler.Execute", name)
		return true
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
	RegisterExtension("v8/bind", jsCode, v8Handler)
}
