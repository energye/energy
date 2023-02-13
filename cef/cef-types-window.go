//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/types"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl/api"
	"io/ioutil"
	"unsafe"
)

func (m *ICefWindow) Show() {
	imports.Proc(internale_ICEFWindow_Show).Call(uintptr(m.instance))
}

func (m *ICefWindow) Hide() {
	imports.Proc(internale_ICEFWindow_Hide).Call(uintptr(m.instance))
}

func (m *ICefWindow) CenterWindow(size *TCefSize) {
	imports.Proc(internale_ICEFWindow_CenterWindow).Call(uintptr(m.instance), uintptr(unsafe.Pointer(size)))
}

func (m *ICefWindow) Close() {
	imports.Proc(internale_ICEFWindow_Close).Call(uintptr(m.instance))
}

func (m *ICefWindow) IsClosed() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsClosed).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *ICefWindow) Activate() {
	imports.Proc(internale_ICEFWindow_Activate).Call(uintptr(m.instance))
}

func (m *ICefWindow) Deactivate() {
	imports.Proc(internale_ICEFWindow_Deactivate).Call(uintptr(m.instance))
}

func (m *ICefWindow) IsActive() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsActive).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *ICefWindow) BringToTop() {
	imports.Proc(internale_ICEFWindow_BringToTop).Call(uintptr(m.instance))
}

func (m *ICefWindow) SetAlwaysOnTop(onTop bool) {
	imports.Proc(internale_ICEFWindow_SetAlwaysOnTop).Call(uintptr(m.instance), api.PascalBool(onTop))
}

func (m *ICefWindow) IsAlwaysOnTop() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsAlwaysOnTop).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *ICefWindow) Maximize() {
	imports.Proc(internale_ICEFWindow_Maximize).Call(uintptr(m.instance))
}

func (m *ICefWindow) Minimize() {
	imports.Proc(internale_ICEFWindow_Minimize).Call(uintptr(m.instance))
}

func (m *ICefWindow) Restore() {
	imports.Proc(internale_ICEFWindow_Restore).Call(uintptr(m.instance))
}

func (m *ICefWindow) SetFullscreen(fullscreen bool) {
	imports.Proc(internale_ICEFWindow_SetFullscreen).Call(uintptr(m.instance), api.PascalBool(fullscreen))
}

func (m *ICefWindow) SetBackgroundColor(rect *types.TCefColor) {
	imports.Proc(internale_ICEFWindow_SetBackgroundColor).Call(uintptr(m.instance), rect.ToPtr())
}

func (m *ICefWindow) SetBounds(rect *TCefRect) {
	imports.Proc(internale_ICEFWindow_SetBounds).Call(uintptr(m.instance), uintptr(unsafe.Pointer(rect)))
}

func (m *ICefWindow) SetSize(size *TCefSize) {
	imports.Proc(internale_ICEFWindow_SetSize).Call(uintptr(m.instance), uintptr(unsafe.Pointer(size)))
}

func (m *ICefWindow) SetPosition(point *TCefPoint) {
	imports.Proc(internale_ICEFWindow_SetPosition).Call(uintptr(m.instance), uintptr(unsafe.Pointer(point)))
}

func (m *ICefWindow) IsMaximized() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsMaximized).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *ICefWindow) IsMinimized() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsMinimized).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *ICefWindow) IsFullscreen() bool {
	r1, _, _ := imports.Proc(internale_ICEFWindow_IsFullscreen).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *ICefWindow) SetTitle(title string) {
	imports.Proc(internale_ICEFWindow_SetTitle).Call(uintptr(m.instance), api.PascalStr(title))
}

func (m *ICefWindow) GetTitle() string {
	r1, _, _ := imports.Proc(internale_ICEFWindow_GetTitle).Call(uintptr(m.instance))
	return api.GoStr(r1)
}

func (m *ICefWindow) SetWindowIcon(scaleFactor float32, filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_ICEFWindow_SetWindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

func (m *ICefWindow) SetWindowIconFS(scaleFactor float32, filename string) error {
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_ICEFWindow_SetWindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

func (m *ICefWindow) GetWindowIcon() *ICefImage {
	var ret uintptr
	imports.Proc(internale_ICEFWindow_GetWindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefImage{
		instance: unsafe.Pointer(ret),
	}
}

func (m *ICefWindow) SetWindowAppIcon(scaleFactor float32, filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_ICEFWindow_SetWindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

func (m *ICefWindow) SetWindowAppIconFS(scaleFactor float32, filename string) error {
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_ICEFWindow_SetWindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

func (m *ICefWindow) GetWindowAppIcon() *ICefImage {
	var ret uintptr
	imports.Proc(internale_ICEFWindow_GetWindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefImage{
		instance: unsafe.Pointer(ret),
	}
}

func (m *ICefWindow) AddOverlayView() {
	//do not implement
	//imports.Proc(internale_ICEFWindow_AddOverlayView).Call(uintptr(m.instance))
}

func (m *ICefWindow) ShowMenu(menuModel *ICefMenuModel, point TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	imports.Proc(internale_ICEFWindow_ShowMenu).Call(uintptr(m.instance), uintptr(menuModel.instance), uintptr(unsafe.Pointer(&point)), uintptr(anchorPosition))
}

func (m *ICefWindow) CancelMenu() {
	imports.Proc(internale_ICEFWindow_CancelMenu).Call(uintptr(m.instance))
}

func (m *ICefWindow) GetDisplay() *ICefDisplay {
	var ret uintptr
	imports.Proc(internale_ICEFWindow_GetDisplay).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefDisplay{
		instance: unsafe.Pointer(ret),
	}
}

func (m *ICefWindow) GetClientAreaBoundsInScreen() (result TCefRect) {
	imports.Proc(internale_ICEFWindow_GetClientAreaBoundsInScreen).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *ICefWindow) SetDraggableRegions(regions []TCefDraggableRegion) {
	imports.Proc(internale_ICEFWindow_SetDraggableRegions).Call(uintptr(m.instance), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}

func (m *ICefWindow) GetWindowHandle() consts.TCefWindowHandle {
	r1, _, _ := imports.Proc(internale_ICEFWindow_GetWindowHandle).Call(uintptr(m.instance))
	return consts.TCefWindowHandle(r1)
}

func (m *ICefWindow) SendKeyPress(keyCode int32, eventFlags uint32) {
	imports.Proc(internale_ICEFWindow_SendKeyPress).Call(uintptr(m.instance), uintptr(keyCode), uintptr(eventFlags))
}

func (m *ICefWindow) SendMouseMove(screenX, screenY int32) {
	imports.Proc(internale_ICEFWindow_SendMouseMove).Call(uintptr(m.instance), uintptr(screenX), uintptr(screenY))
}

func (m *ICefWindow) SendMouseEvents(button consts.TCefMouseButtonType, mouseDown, mouseUp bool) {
	imports.Proc(internale_ICEFWindow_SendMouseEvents).Call(uintptr(m.instance), uintptr(button), api.PascalBool(mouseDown), api.PascalBool(mouseUp))
}

func (m *ICefWindow) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) {
	imports.Proc(internale_ICEFWindow_SetAccelerator).Call(uintptr(m.instance), uintptr(commandId), uintptr(keyCode), api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed))
}

func (m *ICefWindow) RemoveAccelerator(commandId int32) {
	imports.Proc(internale_ICEFWindow_RemoveAccelerator).Call(uintptr(m.instance), uintptr(commandId))
}

func (m *ICefWindow) RemoveAllAccelerators() {
	imports.Proc(internale_ICEFWindow_RemoveAllAccelerators).Call(uintptr(m.instance))
}
