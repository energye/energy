//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

// ICefMenuModelDelegate
//
//	Implement this interface to handle menu model events. The functions of this
//	interface will be called on the browser process UI thread unless otherwise indicated.
//	<para><see cref="uCEFTypes|TCefMenuModelDelegate">Implements TCefMenuModelDelegate</see></para>
//	<para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_menu_model_delegate_capi.h">CEF source file: /include/capi/cef_menu_model_delegate_capi.h (cef_menu_model_delegate_t)</see></para>
type ICefMenuModelDelegate struct {
	base     TCefBaseRefCounted
	instance unsafe.Pointer
}

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

// Perform the action associated with the specified |command_id| and optional |event_flags|.
func (m *ICefMenuModelDelegate) SetOnExecuteCommand(fn menuModelDelegateOnExecuteCommand) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_ExecuteCommand).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called when the user moves the mouse outside the menu and over the owning window.
func (m *ICefMenuModelDelegate) SetOnMouseOutsideMenu(fn menuModelDelegateOnMouseOutsideMenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_MouseOutsideMenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called on unhandled open submenu keyboard commands. |is_rtl| will be true
// (1) if the menu is displaying a right-to-left language.
func (m *ICefMenuModelDelegate) SetOnUnhandledOpenSubmenu(fn menuModelDelegateOnUnhandledOpenSubmenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_UnhandledOpenSubmenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Called on unhandled close submenu keyboard commands. |is_rtl| will be true
// (1) if the menu is displaying a right-to-left language.
func (m *ICefMenuModelDelegate) SetOnUnhandledCloseSubmenu(fn menuModelDelegateOnUnhandledCloseSubmenu) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_UnhandledCloseSubmenu).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// The menu is about to show.
func (m *ICefMenuModelDelegate) SetOnMenuWillShow(fn menuModelDelegateOnMenuWillShow) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_MenuWillShow).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// The menu has closed.
func (m *ICefMenuModelDelegate) SetOnMenuClosed(fn menuModelDelegateOnMenuClosed) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_MenuClosed).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

// Optionally modify a menu item label. Return true (1) if |label| was modified.
func (m *ICefMenuModelDelegate) SetOnFormatLabel(fn menuModelDelegateOnFormatLabel) {
	if !m.IsValid() {
		return
	}
	imports.Proc(def.MenuModelDelegate_FormatLabel).Call(m.Instance(), api.MakeEventDataPtr(fn))
}

type menuModelDelegateOnExecuteCommand func(menuModel *ICefMenuModel, commandId int32, eventFlags consts.TCefEventFlags)
type menuModelDelegateOnMouseOutsideMenu func(menuModel *ICefMenuModel, screenPoint *TCefPoint)
type menuModelDelegateOnUnhandledOpenSubmenu func(menuModel *ICefMenuModel, isRTL bool)
type menuModelDelegateOnUnhandledCloseSubmenu func(menuModel *ICefMenuModel, isRTL bool)
type menuModelDelegateOnMenuWillShow func(menuModel *ICefMenuModel)
type menuModelDelegateOnMenuClosed func(menuModel *ICefMenuModel)
type menuModelDelegateOnFormatLabel func(menuModel *ICefMenuModel) (label string, result bool)

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case menuModelDelegateOnExecuteCommand:
			fn.(menuModelDelegateOnExecuteCommand)(&ICefMenuModel{instance: getPtr(0)}, int32(getVal(1)), consts.TCefEventFlags(getVal(2)))
		case menuModelDelegateOnMouseOutsideMenu:
			fn.(menuModelDelegateOnMouseOutsideMenu)(&ICefMenuModel{instance: getPtr(0)}, (*TCefPoint)(getPtr(1)))
		case menuModelDelegateOnUnhandledOpenSubmenu:
			fn.(menuModelDelegateOnUnhandledOpenSubmenu)(&ICefMenuModel{instance: getPtr(0)}, api.GoBool(getVal(1)))
		case menuModelDelegateOnUnhandledCloseSubmenu:
			fn.(menuModelDelegateOnUnhandledCloseSubmenu)(&ICefMenuModel{instance: getPtr(0)}, api.GoBool(getVal(1)))
		case menuModelDelegateOnMenuWillShow:
			fn.(menuModelDelegateOnMenuWillShow)(&ICefMenuModel{instance: getPtr(0)})
		case menuModelDelegateOnMenuClosed:
			fn.(menuModelDelegateOnMenuClosed)(&ICefMenuModel{instance: getPtr(0)})
		case menuModelDelegateOnFormatLabel:
			labelPtr := (*uintptr)(getPtr(1))
			resultPtr := (*bool)(getPtr(2))
			label, result := fn.(menuModelDelegateOnFormatLabel)(&ICefMenuModel{instance: getPtr(0)})
			*labelPtr = api.PascalStr(label)
			*resultPtr = result
		default:
			return false
		}
		return true
	})
}
