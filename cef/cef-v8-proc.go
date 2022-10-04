//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/golcl/lcl/api"
)

func _CEFV8ValueRef_SetCommonRootName(commonRootName string) {
	Proc("CEFV8ValueRef_SetCommonRootName").Call(api.GoStrToDStr(commonRootName))
}

func _CEFV8ValueRef_SetObjectRootName(objectRootName string) {
	Proc("CEFV8ValueRef_SetObjectRootName").Call(api.GoStrToDStr(objectRootName))
}

func _CEFV8ValueRef_CommonValueBindInfo(binds uintptr) {
	Proc("CEFV8ValueRef_CommonValueBindInfo").Call(binds)
}

func _CEFV8ValueRef_ObjectValueBindInfo(CefObject uintptr) {
	Proc("CEFV8ValueRef_ObjectValueBindInfo").Call(CefObject)
}
