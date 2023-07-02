//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF 右键菜单

package cef

import (
	"github.com/energye/energy/v2/cef/internal/def"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// MenuModelDelegateRef -> ICefMenuModelDelegate
var MenuModelDelegateRef menuModelDelegate

type menuModelDelegate uintptr

func (*menuModelDelegate) New() *ICefMenuModelDelegate {
	var result uintptr
	imports.Proc(def.MenuModelDelegateRef_Create).Call(uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		return &ICefMenuModelDelegate{instance: unsafe.Pointer(result)}
	}
	return nil
}

// Instance 实例
func (m *ICefMenuModelDelegate) Instance() uintptr {
	if m == nil {
		return 0
	}
	return uintptr(m.instance)
}

func (m *ICefMenuModelDelegate) Free() {
	if m.instance != nil {
		m.base.Free(m.Instance())
		m.instance = nil
	}
}

func (m *ICefMenuModelDelegate) IsValid() bool {
	if m == nil || m.instance == nil {
		return false
	}
	return m.instance != nil
}

func (m *ICefRunFileDialogCallback) ExecuteCommand(fn executeCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_ExecuteCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRunFileDialogCallback) MouseOutsideMenu(fn mouseOutsideMenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_MouseOutsideMenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRunFileDialogCallback) UnhandledOpenSubmenu(fn unhandledOpenSubmenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_UnhandledOpenSubmenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRunFileDialogCallback) UnhandledCloseSubmenu(fn unhandledCloseSubmenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_UnhandledCloseSubmenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRunFileDialogCallback) MenuWillShow(fn menuWillShow) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_MenuWillShow).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRunFileDialogCallback) MenuClosed(fn menuClosed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_MenuClosed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *ICefRunFileDialogCallback) FormatLabel(fn formatLabel) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_FormatLabel).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type executeCommand func(menuModel *ICefMenuModel, commandId int32, eventFlags consts.TCefEventFlags)
type mouseOutsideMenu func(menuModel *ICefMenuModel, screenPoint *TCefPoint)
type unhandledOpenSubmenu func(menuModel *ICefMenuModel, isRTL bool)
type unhandledCloseSubmenu func(menuModel *ICefMenuModel, isRTL bool)
type menuWillShow func(menuModel *ICefMenuModel)
type menuClosed func(menuModel *ICefMenuModel)
type formatLabel func(menuModel *ICefMenuModel) (label string, result bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case executeCommand:
			fn.(executeCommand)(&ICefMenuModel{instance: getPtr(0)}, int32(getVal(1)), consts.TCefEventFlags(getVal(2)))
		case mouseOutsideMenu:
			fn.(mouseOutsideMenu)(&ICefMenuModel{instance: getPtr(0)}, (*TCefPoint)(getPtr(1)))
		case unhandledOpenSubmenu:
			fn.(unhandledOpenSubmenu)(&ICefMenuModel{instance: getPtr(0)}, api.GoBool(getVal(1)))
		case unhandledCloseSubmenu:
			fn.(unhandledCloseSubmenu)(&ICefMenuModel{instance: getPtr(0)}, api.GoBool(getVal(1)))
		case menuWillShow:
			fn.(menuWillShow)(&ICefMenuModel{instance: getPtr(0)})
		case menuClosed:
			fn.(menuClosed)(&ICefMenuModel{instance: getPtr(0)})
		case formatLabel:
			labelPtr := (*uintptr)(getPtr(1))
			resultPtr := (*bool)(getPtr(2))
			label, result := fn.(formatLabel)(&ICefMenuModel{instance: getPtr(0)})
			*labelPtr = api.PascalStr(label)
			*resultPtr = result
		default:
			return false
		}
		return true
	})
}
