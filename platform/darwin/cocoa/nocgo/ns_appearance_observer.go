//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	"github.com/ebitengine/purego/objc"
	"reflect"
	"runtime"
	"unsafe"
)

var (
	appearanceObserverClass objc.Class
	appearanceObserverKey   uintptr
	sel_themeChanged        objc.SEL
)

func init() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	initAppearanceObserverSelectors()
	registerAppearanceObserverClass()
}

func initAppearanceObserverSelectors() {
	sel_themeChanged = objc.RegisterName("themeChanged:")
}

func registerAppearanceObserverClass() {
	var err error
	appearanceObserverClass, err = objc.RegisterClass(
		"TAppearanceObserver",
		objc.GetClass("NSObject"),
		nil, // no protocols
		[]objc.FieldDef{
			{
				Name:      "window",
				Type:      reflect.TypeOf(uintptr(0)),
				Attribute: objc.ReadWrite,
			},
		},
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("startObserving"),
				Fn:  appearanceObserverStartObserving,
			},
			{
				Cmd: objc.RegisterName("stopObserving"),
				Fn:  appearanceObserverStopObserving,
			},
			{
				Cmd: objc.RegisterName("themeChanged:"),
				Fn:  appearanceObserverThemeChanged,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	// 创建关联对象 key (使用一个全局变量的地址作为唯一 key)
	appearanceObserverKey = uintptr(unsafe.Pointer(&appearanceObserverClass))
}

// appearanceObserverStartObserving 开始监听系统主题变更通知
func appearanceObserverStartObserving(self objc.ID, _cmd objc.SEL) {
	// 获取 NSDistributedNotificationCenter
	notifCenterClass := objc.GetClass("NSDistributedNotificationCenter")
	notifCenter := objc.ID(notifCenterClass).Send(objc.RegisterName("defaultCenter"))

	// addObserver:selector:name:object:
	notifCenter.Send(objc.RegisterName("addObserver:selector:name:object:"),
		self,
		sel_themeChanged,
		// AppleInterfaceThemeChangedNotification
		objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"),
			"AppleInterfaceThemeChangedNotification"),
		0) // nil object
}

// appearanceObserverStopObserving 停止监听
func appearanceObserverStopObserving(self objc.ID, _cmd objc.SEL) {
	notifCenterClass := objc.GetClass("NSDistributedNotificationCenter")
	notifCenter := objc.ID(notifCenterClass).Send(objc.RegisterName("defaultCenter"))
	notifCenter.Send(objc.RegisterName("removeObserver:"), self)
}

// appearanceObserverThemeChanged 主题变更回调
func appearanceObserverThemeChanged(self objc.ID, _cmd objc.SEL, notification objc.ID) {
	windowID := self.Send(objc.RegisterName("window"))
	if windowID == 0 {
		return
	}
	nsWindow := (*NSWindow)(unsafe.Pointer(windowID))
	isDark := isCurrentlyDarkMode()
	nsWindow.doAppearanceChanged(isDark)
}

// isCurrentlyDarkMode 检测当前系统是否为暗色模式
func isCurrentlyDarkMode() bool {
	nsApp := objc.ID(objc.GetClass("NSApplication")).Send(objc.RegisterName("sharedApplication"))
	effectiveAppearance := nsApp.Send(objc.RegisterName("effectiveAppearance"))
	if effectiveAppearance == 0 {
		return false
	}

	// 创建外观名称数组
	nsArray := objc.ID(objc.GetClass("NSMutableArray")).Send(objc.RegisterName("array"))
	nsArray.Send(objc.RegisterName("addObject:"),
		objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"),
			"NSAppearanceNameAqua"))
	nsArray.Send(objc.RegisterName("addObject:"),
		objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"),
			"NSAppearanceNameDarkAqua"))

	// bestMatchFromAppearancesWithNames:
	bestMatch := effectiveAppearance.Send(objc.RegisterName("bestMatchFromAppearancesWithNames:"), nsArray)
	if bestMatch == 0 {
		return false
	}

	// 检查是否为 DarkAqua
	darkAqua := objc.ID(objc.GetClass("NSString")).Send(objc.RegisterName("stringWithUTF8String:"),
		"NSAppearanceNameDarkAqua")
	return bestMatch.Send(objc.RegisterName("isEqualToString:"), darkAqua) != 0
}

// IsDarkMode 导出函数：检测当前系统是否为暗色模式
func IsDarkMode() bool {
	return isCurrentlyDarkMode()
}
