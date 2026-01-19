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

package widget

/*
#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "go_arguments.h"

*/
import "C"
import (
	"github.com/energye/lcl/tool"
	"unsafe"
)

type OCGoArgumentsType = C.GoArgumentsType

const (
	GoArgsType_None    = OCGoArgumentsType(C.ArgsType_None)    // 未使用类型
	GoArgsType_Int     = OCGoArgumentsType(C.ArgsType_Int)     // 基础类型 int
	GoArgsType_Float   = OCGoArgumentsType(C.ArgsType_Float)   // 基础类型 float64
	GoArgsType_Bool    = OCGoArgumentsType(C.ArgsType_Bool)    // 基础类型 bool
	GoArgsType_String  = OCGoArgumentsType(C.ArgsType_String)  // 基础类型 string
	GoArgsType_Struct  = OCGoArgumentsType(C.ArgsType_Struct)  // C 结构体类型
	GoArgsType_Object  = OCGoArgumentsType(C.ArgsType_Object)  // 对象类型 NS 里创建的对象指针 (void*)[obj retain]
	GoArgsType_Pointer = OCGoArgumentsType(C.ArgsType_Pointer) // 指针类型 NS 里创建的 [NSValue valueWithPointer:customData]
)

type CStructPointer Pointer

// OCGoArgsItem 参数项 对应到Go需要手动正确处理
type OCGoArgsItem struct {
	item Pointer
}

// OCGoArguments 在OC设置的常用数据类型
// 对应到 Go 需要手动对应数组顺序正确处理
type OCGoArguments struct {
	arguments Pointer
	count     int
}

// GoArguments 在Go设置的常用数据类型
// 对应到 OC 需要手动对应数组顺序正确处理
// 基础类型 直接设置
// 结构 需要转为对应OC的结构，并使用 CStructPointer(xx) 做为参数
// OC对象 需要转为对应到 C 的 Pointer(xx)
type GoArguments struct {
	Items []any
}

func (m *GoArguments) Add(v any) {
	m.Items = append(m.Items, v)
}

func (m *GoArguments) AddCStruct(v CStructPointer) {
	m.Items = append(m.Items, v)
}

func (m *GoArguments) ToOC() *C.GoArguments {
	if len(m.Items) == 0 {
		return nil
	}
	goArgs := (*C.GoArguments)(C.malloc(C.sizeof_GoArguments))
	goArgs.Count = C.int(len(m.Items))
	itemSize := C.size_t(unsafe.Sizeof(C.GoArgsItem{}))
	goArgs.Items = (*C.GoArgsItem)(C.malloc(C.size_t(goArgs.Count) * itemSize))
	for i := 0; i < int(goArgs.Count); i++ {
		itemPtr := (*C.GoArgsItem)(Pointer(
			uintptr(Pointer(goArgs.Items)) + uintptr(i)*uintptr(itemSize),
		))
		arg := m.Items[i]
		switch v := arg.(type) {
		case CStructPointer:
			itemPtr.Type = GoArgsType_Struct
			itemPtr.Value = Pointer(v) // 对应 C 结构体
		case int, int8, uint8, int16, uint16, int32, uint32, int64, uint64:
			itemPtr.Type = GoArgsType_Int
			val := (*C.int)(C.malloc(C.sizeof_int))
			*val = C.int(tool.ToInt(v)) // 假设tool.ToInt能统一转换为int
			itemPtr.Value = Pointer(val)
		case float32, float64:
			itemPtr.Type = GoArgsType_Float
			val := (*C.double)(C.malloc(C.sizeof_double))
			*val = C.double(tool.ToDouble(v)) // 统一转换为double
			itemPtr.Value = Pointer(val)
		case bool:
			itemPtr.Type = GoArgsType_Bool
			val := (*C.bool)(C.malloc(C.sizeof_bool))
			*val = C.bool(v)
			itemPtr.Value = Pointer(val)
		case string:
			itemPtr.Type = GoArgsType_String
			cVal := C.CString(v)
			itemPtr.Value = Pointer(cVal) // C.CString内部使用malloc，与OC的strdup行为一致（需外部释放）
		case uintptr:
			itemPtr.Type = GoArgsType_Pointer
			itemPtr.Value = Pointer(v) // 对应OC中NSValue包装的指针或其他对象（需确保生命周期正确）
		case Pointer:
			itemPtr.Type = GoArgsType_Object
			itemPtr.Value = v // 对应 类对象
		default:
			println("[警告] 不支持的类型")
		}
	}
	return goArgs
}

func (m *OCGoArguments) GetItem(index int) *OCGoArgsItem {
	item := C.GetItemFromGoArguments((*C.GoArguments)(m.arguments), C.int(index))
	return &OCGoArgsItem{item: Pointer(item)}
}

func (m *OCGoArgsItem) Type() OCGoArgumentsType {
	item := (*C.GoArgsItem)(m.item)
	return OCGoArgumentsType(item.Type)
}

func (m *OCGoArgsItem) Value() Pointer {
	item := (*C.GoArgsItem)(m.item)
	return Pointer(item.Value)
}

func (m *OCGoArgsItem) IntValue() int {
	if m.Type() == GoArgsType_Int {
		item := (*C.GoArgsItem)(m.item)
		return int(*(*C.int)(item.Value))
	}
	return 0
}

func (m *OCGoArguments) GetInt(index int) int {
	item := C.GetItemFromGoArguments((*C.GoArguments)(m.arguments), C.int(index))
	if item == nil || item.Type != GoArgsType_Int {
		return 0
	}
	return int(*(*C.int)(item.Value))
}

func (m *OCGoArguments) GetFloat(index int) float64 {
	item := C.GetItemFromGoArguments((*C.GoArguments)(m.arguments), C.int(index))
	if item == nil || item.Type != GoArgsType_Float {
		return 0
	}
	return float64(*(*C.double)(item.Value))
}

func (m *OCGoArguments) GetBool(index int) bool {
	item := C.GetItemFromGoArguments((*C.GoArguments)(m.arguments), C.int(index))
	if item == nil || item.Type != GoArgsType_Bool {
		return false
	}
	return bool(*(*C.bool)(item.Value))
}

func (m *OCGoArguments) GetString(index int) string {
	item := C.GetItemFromGoArguments((*C.GoArguments)(m.arguments), C.int(index))
	if item == nil || item.Type != GoArgsType_String {
		return ""
	}
	return C.GoString((*C.char)(item.Value))
}

func (m *OCGoArguments) GetPointer(index int) Pointer {
	item := C.GetItemFromGoArguments((*C.GoArguments)(m.arguments), C.int(index))
	if item == nil || item.Type != GoArgsType_Pointer {
		return nil
	}
	return Pointer(item.Value)
}

func (m *OCGoArguments) GetObject(index int) Pointer {
	item := C.GetItemFromGoArguments((*C.GoArguments)(m.arguments), C.int(index))
	if item == nil || item.Type != GoArgsType_Object {
		return nil
	}
	return Pointer(item.Value)
}

//export GoFreeGoArguments
func GoFreeGoArguments(data *C.GoArguments) {
	if data == nil {
		return
	}
	C.FreeGoArguments(data)

	//count := int(data.Count)
	//for i := 0; i < count; i++ {
	//	cItem := C.GetItemFromGoArguments(data, C.int(i))
	//	cType := OCGoArgumentsType(cItem.Type)
	//	println("FreeGoArguments: 释放第", i, "个参数, 类型:", cType)
	//	switch cType {
	//	case GoArgsType_Int, GoArgsType_Float, GoArgsType_Bool, GoArgsType_String:
	//		if cItem.Value != nil {
	//			C.free(Pointer(cItem.Value))
	//			cItem.Value = nil
	//		}
	//	case GoArgsType_Object, GoArgsType_Pointer: // 只在OC创建
	//	}
	//}
	//println("FreeGoArguments: 释放参数数组 Items")
	//C.free(Pointer(data.Items))
	//println("FreeGoArguments: 释放参数 data")
	//C.free(Pointer(data))
}
