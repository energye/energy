//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/energy/v2/types"
	"sync"
)

var (
	createParamsMap sync.Map
	formCreate      sync.Map
)

func addRequestCreateParamsMap(ptr uintptr, proc IForm) {
	createParamsMap.Store(ptr, proc)
}

func addRequestFormCreateMap(ptr uintptr, proc IForm) {
	formCreate.Store(ptr, proc)
}

func requestCallCreateParamsCallbackProc(ptr uintptr, sender, params uintptr) uintptr {
	if val, ok := createParamsMap.Load(ptr); ok {
		//val.(reflect.Value).Call([]reflect.Value{reflect.ValueOf((*types.TCreateParams)(unsafe.Pointer(params)))})
		if form, ok := val.(IForm); ok {
			form.SetInstance(unsafePointer(sender))
			switch form.(type) {
			case IOnCreateParams:
				form.(IOnCreateParams).CreateParams((*types.TCreateParams)(unsafePointer(params)))
			}
		}
		createParamsMap.Delete(ptr)
	}
	return 0
}

func requestCallFormCreateCallbackProc(ptr uintptr, sender uintptr) uintptr {
	if val, ok := formCreate.Load(ptr); ok {
		if form, ok := val.(IForm); ok {
			form.SetInstance(unsafePointer(sender))
			switch form.(type) {
			case IOnCreate:
				form.(IOnCreate).FormCreate(form)
			}
		}
		formCreate.Delete(ptr)
	}
	return 0
}
