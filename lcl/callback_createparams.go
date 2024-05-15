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

// TOnCreateParams TForm CreateParams 函数
type TOnCreateParams func(params *types.TCreateParams)

var requestCreateParamsMap sync.Map

func addToRequestCreateParamsMap(ptr uintptr, proc TOnCreateParams) {
	requestCreateParamsMap.Store(ptr, proc)
}

func requestCallCreateParamsCallbackProc(ptr uintptr, params uintptr) uintptr {
	if val, ok := requestCreateParamsMap.Load(ptr); ok {
		//val.(reflect.Value).Call([]reflect.Value{reflect.ValueOf((*types.TCreateParams)(unsafe.Pointer(params)))})
		if fn, ok := val.(TOnCreateParams); ok {
			fn((*types.TCreateParams)(unsafePointer(params)))
		}
	}
	return 0
}
