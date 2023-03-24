//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package bind

import (
	"bytes"
	"github.com/energye/golcl/lcl/api"
	"strconv"
)

// ICefV8Context bindGoToJS
//
// 主进程创建完之后和渲染进程每次创建之后调用
//
// 潜在问题，如果函数名包含数字可能会引起函数冲突，入参或出参类型不正确，导致调用失败
func bindGoToJS() {
	//通过回调函数绑定到CEF
	//VariableBind.callVariableBind(browser, frame)
	//通过直接绑定到CEF
	var valueBindInfos []*valueBindInfo
	for _, value := range VariableBind.binds() {
		if !value.isCommon() {
			continue
		}
		jsValue := value.(JSValue)
		var vBind = &valueBindInfo{
			BindType: uintptr(int32(jsValue.ValueType().Jsv)),
		}
		vBind.Name = api.PascalStr(jsValue.Name())
		vBind.EventId = jsValue.getEventId()
		valueBindInfos = append(valueBindInfos, vBind)
		if jsValue.IsFunction() {
			var inParamBuf bytes.Buffer
			var outParamBuf bytes.Buffer
			fnInfo := jsValue.getFuncInfo()
			fnInNum := len(fnInfo.InParam)
			fnOutNum := len(fnInfo.OutParam)
			vBind.FnInNum = uintptr(fnInNum)
			vBind.FnOutNum = uintptr(fnOutNum)
			for i, inParamType := range fnInfo.InParam {
				if i > 0 {
					inParamBuf.WriteString(",")
				}
				inParamBuf.WriteString(strconv.Itoa(int(inParamType.Jsv)))
			}
			vBind.FnInParamType = api.PascalStr(inParamBuf.String())
			for i, outParamType := range fnInfo.OutParam {
				if i > 0 {
					outParamBuf.WriteString(",")
				}
				outParamBuf.WriteString(strconv.Itoa(int(outParamType.Jsv)))
			}
			vBind.FnOutParamType = api.PascalStr(outParamBuf.String())
		}
	}
	if len(valueBindInfos) > 0 {
		for i := 0; i < len(valueBindInfos); i++ {
			//imports.Proc(internale_CEFV8ValueRef_CommonValueBindInfo).Call(uintptr(unsafe.Pointer(valueBindInfos[i])))
		}
	}
}
