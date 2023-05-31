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
	"bytes"
	"fmt"
	"github.com/energye/energy/v2/pkgs/json"
	"text/template"
)

// registerBindExtension
//  注册绑定JS扩展
func (m *bindRenderProcess) registerBindExtension(bindObject json.JSONObject) {
	temp, err := template.New("v8bind").Parse(bindTemplate)
	if err != nil {
		panic(err)
	}
	var data = map[string]any{}
	data["bind"] = internalBind
	data["JSONObject"] = bindObject.ToJSONString()
	var tempResult = &bytes.Buffer{}
	err = temp.Execute(tempResult, data)
	if err != nil {
		panic(err)
	}
	//fmt.Println("tempResult:", tempResult.String())
	registerExtension(fmt.Sprintf("%s/%s", internalV8Bind, internalBind), tempResult.String(), m.handler)
	tempResult.Reset()
}

var bindTemplate = `let {{.bind}};
(function () {
	const __energyBind = {
		util: {
			isObject(value) {
				if (value && value.t) {
					return false
				}
				return Object.prototype.toString.call(value) === '[object Object]'
			},
			isArray(value) {
				if (value && value.t) {
					return false
				}
				return Object.prototype.toString.call(value) === '[object Array]';
			},
			isProxy(value) {
				if (value && value.t) {
					return false
				}
				return Object.prototype.toString.call(value) === '[object Proxy]';
			},
			isFunction(value) {
				if (!value) {
					return false
				}
				return value.t === 19
			}
		},
		createProxy(object) {
			return new Proxy(object, __energyBind.handler);
		},
		handler: {
			get(target, key, receiver) {
				//console.log("GET target:", target, "key:", key, "value:", target[key], "receiver:", receiver);
				let value = target[key];
				if (!value) {
					if (key === 'toJSON') {
						//native function toJSON();
						//let result = toJSON(target);
						//return JSON.stringify(target);
						return;
					} else if (key === 'length') {
						return target.length;
					}
					throw new Error("error " + key);
				}
				if (__energyBind.util.isObject(value) || __energyBind.util.isArray(value)) {
					return __energyBind.createProxy(value);
				} else if (__energyBind.util.isFunction(value)) {
					return function (...arguments) {
						// call
						native function call();
						let result = call(value.p.toString(), arguments);
						return result;
					}
				}
				// get
				native function get();
				let result = get(value.toString());
				// console.log('native get result:', result);
				return result;
				// return Reflect.get(target, key, receiver);
			},
			set(target, key, newValue, receiver) {
				let value = target[key];
				if (!value) {
					throw new Error("binding object does not allow extension of new fields. " + target + " - " + key);
				} else if (__energyBind.util.isObject(value)) {
					throw new Error("binding object cannot be assigned a value. " + target + " - " + key);
				} else if (__energyBind.util.isFunction(value)) {
					throw new Error("binding object functions cannot be assigned. " + target + " - " + key);
				} else if (__energyBind.util.isArray(value)) {
					throw new Error("binding object arrays cannot be assigned. " + target + " - " + key);
				} else {
					native function set();
					set(value.toString(), newValue);
				}
				console.log("SET target:", target, "key:", key, "value:", value, "newValue:", newValue, "receiver:", receiver);
				return true;
			},
			deleteProperty(target, key) {
				throw new Error("binding object not allowed to be deleted. " + target + " - " + key);
			},
			isExtensible(target) {
				throw new Error("binding object does not allow extension of new fields. " + target);
			},
			// defineProperty() {
			//     throw new Error("error defineProperty");
			// },
			preventExtensions() {
				throw new Error("error preventExtensions");
			},
			setPrototypeOf(){
				throw new Error("error setPrototypeOf");
			}
		}
	}
	{{.bind}} = __energyBind.createProxy({{.JSONObject}});
})();
`
