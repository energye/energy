//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package systray

/*
#cgo darwin CFLAGS: -DDARWIN -x objective-c -fobjc-arc
#cgo darwin LDFLAGS: -framework Cocoa

#include <stdbool.h>
#include "systray.h"

void setInternalLoop(bool);
*/
import "C"

import (
	"time"
	"unsafe"
)

var st = &systray{}

type systray struct {
}

func (m *systray) ShowMenu() error {
	C.show_menu()
	return nil
}

// SetTemplateIcon sets the systray icon as a template icon (on Mac), falling back
// to a regular icon on other platforms.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
	cstr := (*C.char)(unsafe.Pointer(&templateIconBytes[0]))
	C.setIcon(cstr, (C.int)(len(templateIconBytes)), true)
}

// SetIcon sets the icon of a menu item. Only works on macOS and Windows.
// iconBytes should be the content of .ico/.jpg/.png
func (item *MenuItem) SetIcon(iconBytes []byte) {
	cstr := (*C.char)(unsafe.Pointer(&iconBytes[0]))
	C.setMenuItemIcon(cstr, (C.int)(len(iconBytes)), C.int(item.id), false)
}

// SetTemplateIcon sets the icon of a menu item as a template icon (on macOS). On Windows, it
// falls back to the regular icon bytes and on Linux it does nothing.
// templateIconBytes and regularIconBytes should be the content of .ico for windows and
// .ico/.jpg/.png for other platforms.
func (item *MenuItem) SetTemplateIcon(templateIconBytes []byte, regularIconBytes []byte) {
	cstr := (*C.char)(unsafe.Pointer(&templateIconBytes[0]))
	C.setMenuItemIcon(cstr, (C.int)(len(templateIconBytes)), C.int(item.id), true)
}

func registerSystray() {
	C.registerSystray()
}

func nativeLoop() {
	C.nativeLoop()
}

func nativeEnd() {
	C.nativeEnd()
}

func nativeStart() {
	C.nativeStart()
}

func quit() {
	C.quit()
}

func setInternalLoop(internal bool) {
	C.setInternalLoop(C.bool(internal))
}

var (
	onClick         func()
	onDClick        func()
	onRClick        func(menu IMenu)
	dClickTime      int64
	isEnableOnClick = false
)

func setOnClick(fn func()) {
	enableOnClick()
	onClick = fn
}

func setOnDClick(fn func()) {
	enableOnClick()
	onDClick = fn
}

func setOnRClick(fn func(menu IMenu)) {
	enableOnClick()
	onRClick = fn
}

// CreateMenu 创建托盘菜单, 如果托盘菜单是空, 把菜单项添加到托盘
// 该方法主动调用后 如果托盘菜单已创建则添加进去, 之后鼠标事件失效
//
// 仅MacOSX平台
func CreateMenu() {
	createMenu()
}

// SetMenuNil 托盘菜单设置为nil, 如果托盘菜单不是空, 把菜单项设置为nil
// 该方法主动调用后 将移除托盘菜单, 之后鼠标事件生效
//
// 仅MacOSX平台
func SetMenuNil() {
	setMenuNil()
}

// SetIcon sets the systray icon.
// iconBytes should be the content of .ico for windows and .ico/.jpg/.png
// for other platforms.
func SetIcon(iconBytes []byte) {
	cstr := (*C.char)(unsafe.Pointer(&iconBytes[0]))
	C.setIcon(cstr, (C.int)(len(iconBytes)), false)
}

// SetTitle sets the systray title, only available on Mac and Linux.
func SetTitle(title string) {
	C.setTitle(C.CString(title))
}

// SetTooltip sets the systray tooltip to display on mouse hover of the tray icon,
// only available on Mac and Windows.
func SetTooltip(tooltip string) {
	C.setTooltip(C.CString(tooltip))
}

func addOrUpdateMenuItem(item *MenuItem) {
	var disabled C.short
	if item.disabled {
		disabled = 1
	}
	var checked C.short
	if item.checked {
		checked = 1
	}
	var isCheckable C.short
	if item.isCheckable {
		isCheckable = 1
	}
	var parentID uint32 = 0
	if item.parent != nil {
		parentID = item.parent.id
	}
	C.add_or_update_menu_item(
		C.int(item.id),
		C.int(parentID),
		C.CString(item.title),
		C.CString(item.tooltip),
		C.CString(item.shortcutKey),
		disabled,
		checked,
		isCheckable,
	)
}

func addSeparator(id uint32) {
	C.add_separator(C.int(id))
}

func hideMenuItem(item *MenuItem) {
	C.hide_menu_item(
		C.int(item.id),
	)
}

func showMenuItem(item *MenuItem) {
	C.show_menu_item(
		C.int(item.id),
	)
}

func removeMenuItem(item *MenuItem) {
	C.remove_menu_item(
		C.int(item.id),
	)
}

func resetMenu() {
	C.reset_menu()
}

func createMenu() {
	C.create_menu()
}

func setMenuNil() {
	C.set_menu_nil()
}
func enableOnClick() {
	if !isEnableOnClick {
		isEnableOnClick = true
		C.enable_on_click()
	}
}

//export systray_ready
func systray_ready() {
	if systrayReady != nil {
		systrayReady()
	}
}

//export systray_on_exit
func systray_on_exit() {
	systrayExit()
}

//export systray_menu_item_selected
func systray_menu_item_selected(cID C.int) {
	systrayMenuItemSelected(uint32(cID))
}

//export systray_on_click
func systray_on_click() {
	if dClickTime == 0 {
		dClickTime = time.Now().UnixMilli()
	} else {
		nowMilli := time.Now().UnixMilli()
		if nowMilli-dClickTime < dClickTimeMinInterval {
			dClickTime = dClickTimeMinInterval
			if onDClick != nil {
				onDClick()
				return
			}
		} else {
			dClickTime = nowMilli
		}
	}
	if onClick != nil {
		onClick()
	}
}

//export systray_on_rclick
func systray_on_rclick() {
	if onRClick != nil {
		onRClick(st)
	} else {
		C.show_menu()
	}
}
