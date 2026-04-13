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

package cgo

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#import "cocoa_helpers.h"

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// InspectControl 检查控件类型并打印详细信息
func InspectControl(handle uintptr) {
	if handle == 0 {
		fmt.Println("错误：无效的句柄（空指针）")
		return
	}

	cHandle := unsafe.Pointer(handle)

	// 1. 基本类型信息
	className := C.getObjectClassName(cHandle)
	fmt.Printf("控件类型: %s\n", C.GoString(className))

	// 2. 常见类型检查
	commonTypes := []string{
		"NSButton", "NSTextField", "NSComboBox",
		"NSSlider", "NSWindow", "NSView", "NSObject",
	}
	fmt.Println("\n类型匹配检查:")
	for _, typ := range commonTypes {
		cType := C.CString(typ)
		isMatch := int(C.isObjectOfClass(cHandle, cType)) > 0
		C.free(unsafe.Pointer(cType))
		fmt.Printf("  是 %-12s: %v\n", typ, isMatch)
	}

	// 3. 继承链信息
	fmt.Println("\n继承链（从自身到根类）:")
	chain := C.getObjectInheritanceChain(cHandle)
	defer C.freeInheritanceChain(&chain)

	if chain.count == 0 {
		fmt.Println("  无法获取继承链信息")
		return
	}

	for i := 0; i < int(chain.count); i++ {
		classNames := (*[1 << 20]*C.char)(unsafe.Pointer(chain.classNames))
		className := C.GoString(classNames[i])
		fmt.Printf("  第 %2d 层: %s\n", i+1, className)
	}
}

var cocoaClassNameList map[string]struct{}

// 初始化使用的 lcl 控件, 用于验证是否为 cocoa 控件
func initUseClassNameList() {
	cocoaClassNames := []string{"NSButton", "TCocoaButton"}
	cocoaClassNameList = make(map[string]struct{})
	for _, className := range cocoaClassNames {
		cocoaClassNameList[className] = struct{}{}
	}
}

func VerifyWidget(handle uintptr) bool {
	if cocoaClassNameList == nil {
		initUseClassNameList()
	}
	cHandle := unsafe.Pointer(handle)
	className := C.GoString(C.getObjectClassName(cHandle))
	_, ok := cocoaClassNameList[className]
	if !ok {
		// debug
		InspectControl(handle)
	}
	return ok
}

// 获取对象的完整继承链（包括自身）
func GetObjectInheritanceChain(handle unsafe.Pointer) (classNames []string) {
	if handle == nil {
		return nil
	}
	chain := C.getObjectInheritanceChain(handle)
	defer C.freeInheritanceChain(&chain)
	if chain.count == 0 {
		return nil
	}
	for i := 0; i < int(chain.count); i++ {
		tempClassNames := (*[1 << 20]*C.char)(unsafe.Pointer(chain.classNames))
		className := C.GoString(tempClassNames[i])
		classNames = append(classNames, className)
	}
	return
}

// 获取对象的类名
func GetObjectClassName(handle unsafe.Pointer) string {
	if handle == nil {
		return ""
	}
	className := C.getObjectClassName(handle)
	return C.GoString(className)
}
