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
	. "github.com/energye/energy/v3/platform/darwin/types"
	"reflect"
	"unsafe"
)

var (
	appDelegateClass                    objc.Class
	sel_applicationOpenURLs             objc.SEL
	sel_applicationContinueUserActivity objc.SEL
	sel_originalDelegate                objc.SEL
	sel_superclass                      objc.SEL
	sel_class                           objc.SEL
	sel_respondsToSelector              objc.SEL
)

type NSAppDelegate struct {
	NSObject
	originalDelegate objc.ID
	callback         func(eventName string, data string)
}

func init() {
	initAppDelegateSelectors()
	registerAppDelegateClass()
}

func initAppDelegateSelectors() {
	sel_applicationOpenURLs = objc.RegisterName("application:openURLs:")
	sel_applicationContinueUserActivity = objc.RegisterName("application:continueUserActivity:restorationHandler:")
	sel_originalDelegate = objc.RegisterName("originalDelegate")
	sel_superclass = objc.RegisterName("superclass")
	sel_class = objc.RegisterName("class")
	sel_respondsToSelector = objc.RegisterName("respondsToSelector:")
}

func registerAppDelegateClass() {
	var err error
	protoApplicationDelegate := objc.GetProtocol("NSApplicationDelegate")
	appDelegateClass, err = objc.RegisterClass(
		"TAppDelegateWrapper",
		objc.GetClass("NSObject"),
		[]*objc.Protocol{protoApplicationDelegate},
		[]objc.FieldDef{
			{
				Name:      "originalDelegate",
				Type:      reflect.TypeOf(objc.ID(0)),
				Attribute: objc.ReadWrite,
			},
			{
				Name:      "nsApp",
				Type:      reflect.TypeOf(uintptr(0)),
				Attribute: objc.ReadWrite,
			},
		},
		[]objc.MethodDef{
			{
				Cmd: objc.RegisterName("application:openURLs:"),
				Fn:  applicationOpenURLs,
			},
			{
				Cmd: objc.RegisterName("application:continueUserActivity:restorationHandler:"),
				Fn:  applicationContinueUserActivity,
			},
			{
				Cmd: objc.RegisterName("forwardingTargetForSelector:"),
				Fn:  forwardingTargetForSelector,
			},
			{
				Cmd: sel_respondsToSelector,
				Fn:  respondsToSelector,
			},
		},
	)
	if err != nil {
		panic(err)
	}
}

func applicationOpenURLs(self objc.ID, _cmd objc.SEL, application objc.ID, urls objc.ID) {
	originalDelegate := objc.ID(self.Send(sel_originalDelegate))
	appID := self.Send(objc.RegisterName("nsApp"))
	if appID != 0 {
		app := (*NSApp)(unsafe.Pointer(appID))
		urlArray := objc.ID(urls)
		count := urlArray.Send(objc.RegisterName("count"))
		items := make([]string, 0, count)
		for i := uintptr(0); i < uintptr(count); i++ {
			url := urlArray.Send(objc.RegisterName("objectAtIndex:"), i)
			absoluteString := objc.ID(url).Send(objc.RegisterName("absoluteString"))
			goString := NSStringToGoString(absoluteString)
			items = append(items, goString)
		}
		app.doOpenURLs(items)
	}

	if originalDelegate != 0 {
		if originalDelegate.Send(sel_respondsToSelector, sel_applicationOpenURLs) != 0 {
			originalDelegate.Send(sel_applicationOpenURLs, application, urls)
		}
	}
}

func applicationContinueUserActivity(self objc.ID, _cmd objc.SEL,
	application objc.ID, userActivity objc.ID, restorationHandler objc.ID) bool {
	originalDelegate := self.Send(sel_originalDelegate)
	appID := self.Send(objc.RegisterName("nsApp"))
	if appID != 0 {
		app := (*NSApp)(unsafe.Pointer(appID))
		activityType := userActivity.Send(objc.RegisterName("activityType"))
		activityTypeStr := NSStringToGoString(activityType)

		if activityTypeStr == "NSUserActivityTypeBrowsingWeb" {
			webpageURL := userActivity.Send(objc.RegisterName("webpageURL"))
			if webpageURL != 0 {
				absoluteString := webpageURL.Send(objc.RegisterName("absoluteString"))
				data := NSStringToGoString(absoluteString)
				app.doUniversalLink(data)
			}
		}
	}

	if originalDelegate != 0 {
		if originalDelegate.Send(sel_respondsToSelector, sel_applicationContinueUserActivity) != 0 {
			result := originalDelegate.Send(sel_applicationContinueUserActivity, application, userActivity, restorationHandler)
			return result != 0
		}
	}

	return true
}

func forwardingTargetForSelector(self objc.ID, _cmd objc.SEL, aSelector objc.SEL) objc.ID {
	originalDelegate := objc.ID(self.Send(sel_originalDelegate))
	if originalDelegate != 0 {
		if originalDelegate.Send(sel_respondsToSelector, aSelector) != 0 {
			return originalDelegate
		}
	}
	return objc.ID(0)
}

func respondsToSelector(self objc.ID, _cmd objc.SEL, aSelector objc.SEL) bool {
	// 1. 检查父类是否实现（模拟 super 的行为）
	class := self.Send(sel_class)
	if class != 0 {
		superClass := class.Send(sel_superclass)
		if superClass != 0 {
			method := superClass.Send(sel_respondsToSelector, aSelector)
			if method != 0 {
				return true
			}
		}
	}
	// 2. 检查原代理是否实现
	originalDelegate := objc.ID(self.Send(sel_originalDelegate))
	if originalDelegate != 0 {
		if originalDelegate.Send(sel_respondsToSelector, aSelector) != 0 {
			return true
		}
	}
	return false
}

// InitAppDelegate 初始化 macOS 应用程序代理
func (m *NSApp) InitAppDelegate() {
	if m.initializationAppDelegate {
		return
	}
	m.initializationAppDelegate = true

	nsAppClass := objc.GetClass("NSApplication")
	nsApp := objc.ID(nsAppClass).Send(objc.RegisterName("sharedApplication"))

	currentDelegate := nsApp.Send(objc.RegisterName("delegate"))
	if currentDelegate == 0 {
		return
	}

	delegate := objc.ID(appDelegateClass).Send(objc.RegisterName("new"))
	delegate.Send(objc.RegisterName("setOriginalDelegate:"), currentDelegate)
	println("[DEBUG] New delegate class:", getClassName(delegate))

	app := m
	delegate.Send(objc.RegisterName("setNsApp:"), uintptr(unsafe.Pointer(app)))

	println("[DEBUG] Old delegate class:", getClassName(currentDelegate))

	nsApp.Send(objc.RegisterName("setDelegate:"), delegate)

	println("[DEBUG] New delegate class:", getClassName(delegate))
}

func getClassName(obj objc.ID) string {
	class := obj.Send(sel_class)
	className := class.Send(objc.RegisterName("className"))
	return NSStringToGoString(className)
}

func (m *NSApp) SetOnOpenURLs(fn TOpenURLsEvent) {
	m.onOpenURLs = fn
}

func (m *NSApp) SetOnUniversalLink(fn TUniversalLinkEvent) {
	m.onUniversalLink = fn
}
