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

package toolbar

/*

#cgo CFLAGS: -mmacosx-version-min=10.15 -x objective-c
#cgo LDFLAGS: -mmacosx-version-min=10.15 -framework Cocoa
#include "config.h"

extern GoArguments* onDelegateEvent(ToolbarCallbackContext *cContext);

*/
import "C"
import (
	"github.com/energye/lcl/lcl"
	"strconv"
	"sync"
)

//export onDelegateEvent
func onDelegateEvent(cContext *C.ToolbarCallbackContext) *C.GoArguments {
	ctx := ToolbarCallbackContext{
		Type:       TccType(cContext.type_),
		Identifier: C.GoString(cContext.identifier),
		Value:      C.GoString(cContext.value),
		Index:      int(cContext.index),
		Owner:      cContext.owner,
		Sender:     cContext.sender,
	}
	cArguments := cContext.arguments
	if cArguments != nil {
		ctx.Arguments = &OCGoArguments{arguments: Pointer(cArguments), count: int(cArguments.Count)}
	}
	//fmt.Printf("onDelegateEvent event: %+v\n", ctx)
	eventId := ctx.Identifier
	eventId = eventId + strconv.Itoa(int(ctx.Type))
	cb := eventList[eventId]
	if cb == nil {
		return nil
	}
	if result := cb.cb(&ctx); result != nil {
		return result.ToOC()
	} else {
		return nil
	}
}

func cControlEventCallback() C.ControlEventCallback {
	return (C.ControlEventCallback)(C.onDelegateEvent)
}

// 事件列表
var (
	eventList = make(map[string]*Callback)
	eventLock sync.Mutex
)

// RegisterEvent 事件注册，使用控件唯一标识 + 事件类型做为事件唯一id
func RegisterEvent(identifier string, fn *Callback) {
	eventLock.Lock()
	defer eventLock.Unlock()
	identifier = identifier + strconv.Itoa(int(fn.type_))
	eventList[identifier] = fn
}

func IsARCMode() bool {
	return bool(C.isARCMode())
}

// SetWindowBackgroundColor 公开方法 设置窗口背景色
func SetWindowBackgroundColor(owner lcl.IForm, color Color) {
	nsWindow := uintptr(lcl.PlatformWindow(owner.Instance()))
	if nsWindow == 0 {
		return
	}
	cColor := color.ToOC()
	C.SetWindowBackgroundColor(C.ulong(nsWindow), cColor)
}

// GetStringConstValue 获取 C 字符串常昊值
func GetStringConstValue(constName *C.NSString) string {
	cValue := C.GetStringConstValue(Pointer(constName))
	return C.GoString(cValue)
}

//func AddToolbarButton(nsWindowHandle uintptr, identifier, title, tooltip string, property ControlProperty) {
//	cIdentifier := C.CString(identifier)
//	defer C.free(Pointer(cIdentifier))
//	cTitle := C.CString(title)
//	defer C.free(Pointer(cTitle))
//	var cTooltip *C.char
//	if tooltip != "" {
//		cTooltip = C.CString(tooltip)
//		defer C.free(Pointer(cTooltip))
//	}
//	cProperty := property.ToOC()
//	C.AddToolbarButton(C.ulong(nsWindowHandle), cIdentifier, cTitle, cTooltip, cProperty)
//}

//func AddToolbarImageButton(nsWindowHandle uintptr, identifier, imageName, tooltip string, property ControlProperty) {
//	cIdentifier := C.CString(identifier)
//	defer C.free(Pointer(cIdentifier))
//
//	cImageName := C.CString(imageName)
//	defer C.free(Pointer(cImageName))
//
//	var cTooltip *C.char
//	if tooltip != "" {
//		cTooltip = C.CString(tooltip)
//		defer C.free(Pointer(cTooltip))
//	}
//
//	cProperty := property.ToOC()
//
//	C.AddToolbarImageButton(C.ulong(nsWindowHandle), cIdentifier, cImageName, cTooltip, cProperty)
//}

//func AddToolbarTextField(nsWindowHandle uintptr, identifier, placeholder string, property ControlProperty) {
//	cIdentifier := C.CString(identifier)
//	defer C.free(Pointer(cIdentifier))
//
//	var cPlaceholder *C.char
//	if placeholder != "" {
//		cPlaceholder = C.CString(placeholder)
//		defer C.free(Pointer(cPlaceholder))
//	}
//
//	cProperty := property.ToOC()
//
//	C.AddToolbarTextField(C.ulong(nsWindowHandle), cIdentifier, cPlaceholder, cProperty)
//}

//func AddToolbarSearchField(nsWindowHandle uintptr, identifier, placeholder string, property ControlProperty) *NSSearchField {
//	cIdentifier := C.CString(identifier)
//	defer C.free(Pointer(cIdentifier))
//	var cPlaceholder *C.char
//	if placeholder != "" {
//		cPlaceholder = C.CString(placeholder)
//		defer C.free(Pointer(cPlaceholder))
//	}
//	cProperty := property.ToOC()
//	cSF := C.AddToolbarSearchField(C.ulong(nsWindowHandle), cIdentifier, cPlaceholder, cProperty)
//	return &NSSearchField{instance: Pointer(cSF)}
//}

func AddToolbarCombobox(nsWindowHandle uintptr, identifier string, items []string, property ControlProperty) {
	cIdentifier := C.CString(identifier)
	defer C.free(Pointer(cIdentifier))

	// 转换Go字符串切片为C字符串数组
	cItems := make([]*C.char, len(items))
	for i, item := range items {
		cItems[i] = C.CString(item)
	}
	cProperty := property.ToOC()
	C.AddToolbarCombobox(C.ulong(nsWindowHandle), cIdentifier, (**C.char)(Pointer(&cItems[0])), C.int(len(items)), cProperty)
	for i, _ := range items {
		C.free(Pointer(cItems[i]))
	}
}

func AddToolbarCustomView(nsWindowHandle uintptr, identifier string, property ControlProperty) {
	cIdentifier := C.CString(identifier)
	defer C.free(Pointer(cIdentifier))
	cProperty := property.ToOC()
	C.AddToolbarCustomView(C.ulong(nsWindowHandle), cIdentifier, cProperty)
}

func RemoveToolbarItem(nsWindowHandle uintptr, identifier string) {
	cIdentifier := C.CString(identifier)
	defer C.free(Pointer(cIdentifier))

	C.RemoveToolbarItem(C.ulong(nsWindowHandle), cIdentifier)
}

func GetToolbarControlValue(nsWindowHandle uintptr, identifier string) string {
	cIdentifier := C.CString(identifier)
	defer C.free(Pointer(cIdentifier))

	cValue := C.GetToolbarControlValue(C.ulong(nsWindowHandle), cIdentifier)
	if cValue == nil {
		return ""
	}
	return C.GoString(cValue)
}

func SetToolbarControlValue(nsWindowHandle uintptr, identifier, value string) {
	cIdentifier := C.CString(identifier)
	defer C.free(Pointer(cIdentifier))

	cValue := C.CString(value)
	defer C.free(Pointer(cValue))

	C.SetToolbarControlValue(C.ulong(nsWindowHandle), cIdentifier, cValue)
}

func SetToolbarControlEnabled(nsWindowHandle uintptr, identifier string, enabled bool) {
	cIdentifier := C.CString(identifier)
	defer C.free(Pointer(cIdentifier))
	C.SetToolbarControlEnabled(C.ulong(nsWindowHandle), cIdentifier, C.bool(enabled))
}

func SetToolbarControlHidden(nsWindowHandle uintptr, identifier string, hidden bool) {
	cIdentifier := C.CString(identifier)
	defer C.free(Pointer(cIdentifier))
	C.SetToolbarControlHidden(C.ulong(nsWindowHandle), cIdentifier, C.bool(hidden))
}

func CreateDefaultControlProperty() ControlProperty {
	cProperty := C.CreateDefaultControlProperty()
	return ControlProperty{
		Width:              float64(cProperty.width),
		Height:             float64(cProperty.height),
		MinWidth:           float64(cProperty.minWidth),
		MaxWidth:           float64(cProperty.maxWidth),
		BezelStyle:         NSBezelStyle(cProperty.bezelStyle),
		ControlSize:        NSControlSize(cProperty.controlSize),
		Font:               Pointer(cProperty.font),
		VisibilityPriority: int(cProperty.VisibilityPriority),
	}
}

//现代 macOS 工具栏开发最佳实践总结
//
//理解“统一工具栏”：从 macOS 11 (Big Sur) 开始，工具栏和标题栏在视觉上融合。使用 isNavigational 和 allowedAligned 属性来正确放置你的项。
//明确项的角色：
//导航类 (isNavigational = true)：如前进、后退、侧边栏切换。靠左放置。
//主要操作/搜索 (principalItem)：如搜索栏。居中放置。
//内容相关操作 (allowedAligned = .trailing)：如分享、排序、查看选项。靠右放置。
//灵活空间 (.flexibleSpace, .space)：用于布局和对齐。
//优先使用 SF Symbols：确保图标在不同主题和状态下的一致性。
//善用分组：对于相关的操作（如视图切换：列表、图标、分栏），使用 NSToolbarItemGroup 并以 collapsed 模式显示，以节省空间。
//响应式显示：正确设置 visibilityPriority，确保在窗口变窄时，最重要的项仍然可见，不重要的项会被自动隐藏到溢出菜单中。
