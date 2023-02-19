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
	"github.com/energye/energy/common/imports"
	"github.com/energye/energy/consts"
	"github.com/energye/energy/types"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"io/ioutil"
	"unsafe"
)

type TCEFWindowComponent struct {
	lcl.IComponent
	instance unsafe.Pointer
}

func NewWindowComponent(AOwner lcl.IComponent) *TCEFWindowComponent {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_Create).Call(lcl.CheckPtr(AOwner))
	return &TCEFWindowComponent{
		instance: unsafe.Pointer(r1),
	}
}

func (m *TCEFWindowComponent) CreateTopLevelWindow() {
	imports.Proc(internale_CEFWindowComponent_CreateTopLevelWindow).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Show() {
	imports.Proc(internale_CEFWindowComponent_Show).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Hide() {
	imports.Proc(internale_CEFWindowComponent_Hide).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) CenterWindow(size *TCefSize) {
	imports.Proc(internale_CEFWindowComponent_CenterWindow).Call(uintptr(m.instance), uintptr(unsafe.Pointer(size)))
}

func (m *TCEFWindowComponent) Close() {
	imports.Proc(internale_CEFWindowComponent_Close).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Activate() {
	imports.Proc(internale_CEFWindowComponent_Activate).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Deactivate() {
	imports.Proc(internale_CEFWindowComponent_Deactivate).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) BringToTop() {
	imports.Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Maximize() {
	imports.Proc(internale_CEFWindowComponent_Maximize).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Minimize() {
	imports.Proc(internale_CEFWindowComponent_Minimize).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Restore() {
	imports.Proc(internale_CEFWindowComponent_Restore).Call(uintptr(m.instance))
}

//func (m *TCEFWindowComponent) AddOverlayView() {
//imports.Proc(internale_CEFWindowComponent_AddOverlayView).Call(uintptr(m.instance))
//}

func (m *TCEFWindowComponent) ShowMenu(menuModel *ICefMenuModel, point TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	imports.Proc(internale_CEFWindowComponent_ShowMenu).Call(uintptr(m.instance), uintptr(menuModel.instance), uintptr(unsafe.Pointer(&point)), uintptr(anchorPosition))
}

func (m *TCEFWindowComponent) CancelMenu() {
	imports.Proc(internale_CEFWindowComponent_CancelMenu).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SetDraggableRegions(regions []TCefDraggableRegion) {
	imports.Proc(internale_CEFWindowComponent_SetDraggableRegions).Call(uintptr(m.instance), uintptr(int32(len(regions))), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}

func (m *TCEFWindowComponent) SendKeyPress(keyCode int32, eventFlags uint32) {
	imports.Proc(internale_CEFWindowComponent_SendKeyPress).Call(uintptr(m.instance), uintptr(keyCode), uintptr(eventFlags))
}

func (m *TCEFWindowComponent) SendMouseMove(screenX, screenY int32) {
	imports.Proc(internale_CEFWindowComponent_SendMouseMove).Call(uintptr(m.instance), uintptr(screenX), uintptr(screenY))
}

func (m *TCEFWindowComponent) SendMouseEvents(button consts.TCefMouseButtonType, mouseDown, mouseUp bool) {
	imports.Proc(internale_CEFWindowComponent_SendMouseEvents).Call(uintptr(m.instance), uintptr(button), api.PascalBool(mouseDown), api.PascalBool(mouseUp))
}

func (m *TCEFWindowComponent) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) {
	imports.Proc(internale_CEFWindowComponent_SetAccelerator).Call(uintptr(m.instance), uintptr(commandId), uintptr(keyCode), api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed))
}

func (m *TCEFWindowComponent) RemoveAccelerator(commandId int32) {
	imports.Proc(internale_CEFWindowComponent_RemoveAccelerator).Call(uintptr(m.instance), uintptr(commandId))
}

func (m *TCEFWindowComponent) RemoveAllAccelerators() {
	imports.Proc(internale_CEFWindowComponent_RemoveAllAccelerators).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SetAlwaysOnTop(onTop bool) {
	imports.Proc(internale_CEFWindowComponent_SetAlwaysOnTop).Call(uintptr(m.instance), api.PascalBool(onTop))
}

func (m *TCEFWindowComponent) SetFullscreen(fullscreen bool) {
	imports.Proc(internale_CEFWindowComponent_SetFullscreen).Call(uintptr(m.instance), api.PascalBool(fullscreen))
}

func (m *TCEFWindowComponent) SetBackgroundColor(rect *types.TCefColor) {
	imports.Proc(internale_CEFWindowComponent_SetBackgroundColor).Call(uintptr(m.instance), rect.ToPtr())
}

func (m *TCEFWindowComponent) Bounds() (result *TCefRect) {
	imports.Proc(internale_CEFWindowComponent_Bounds).Call(uintptr(m.instance), uintptr(unsafe.Pointer(result)))
	return
}

func (m *TCEFWindowComponent) Size() (result *TCefSize) {
	imports.Proc(internale_CEFWindowComponent_Size).Call(uintptr(m.instance), uintptr(unsafe.Pointer(result)))
	return
}

func (m *TCEFWindowComponent) Position() (result *TCefPoint) {
	imports.Proc(internale_CEFWindowComponent_Position).Call(uintptr(m.instance), uintptr(unsafe.Pointer(result)))
	return
}

func (m *TCEFWindowComponent) SetBounds(rect *TCefRect) {
	imports.Proc(internale_CEFWindowComponent_SetBounds).Call(uintptr(m.instance), uintptr(unsafe.Pointer(rect)))
}

func (m *TCEFWindowComponent) SetSize(size *TCefSize) {
	imports.Proc(internale_CEFWindowComponent_SetSize).Call(uintptr(m.instance), uintptr(unsafe.Pointer(size)))
}

func (m *TCEFWindowComponent) SetPosition(point *TCefPoint) {
	imports.Proc(internale_CEFWindowComponent_SetPosition).Call(uintptr(m.instance), uintptr(unsafe.Pointer(point)))
}

func (m *TCEFWindowComponent) SetTitle(title string) {
	imports.Proc(internale_CEFWindowComponent_SetTitle).Call(uintptr(m.instance), api.PascalStr(title))
}

func (m *TCEFWindowComponent) Title() string {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_Title).Call(uintptr(m.instance))
	return api.GoStr(r1)
}

func (m *TCEFWindowComponent) WindowIcon() *ICefImage {
	var ret uintptr
	imports.Proc(internale_CEFWindowComponent_WindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefImage{
		instance: unsafe.Pointer(ret),
	}
}

func (m *TCEFWindowComponent) SetWindowIcon(scaleFactor float32, filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_CEFWindowComponent_SetWindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

func (m *TCEFWindowComponent) SetWindowIconFS(scaleFactor float32, filename string) error {
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_CEFWindowComponent_SetWindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

func (m *TCEFWindowComponent) WindowAppIcon() *ICefImage {
	var ret uintptr
	imports.Proc(internale_CEFWindowComponent_WindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefImage{
		instance: unsafe.Pointer(ret),
	}
}

func (m *TCEFWindowComponent) SetWindowAppIcon(scaleFactor float32, filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_CEFWindowComponent_SetWindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

func (m *TCEFWindowComponent) SetWindowAppIconFS(scaleFactor float32, filename string) error {
	bytes, err := emfs.GetResources(filename)
	if err != nil {
		return err
	}
	imports.Proc(internale_CEFWindowComponent_SetWindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&scaleFactor)), uintptr(unsafe.Pointer(&bytes[0])), uintptr(int32(len(bytes))))
	return nil
}

func (m *TCEFWindowComponent) Display() *ICefDisplay {
	var ret uintptr
	imports.Proc(internale_CEFWindowComponent_Display).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefDisplay{
		instance: unsafe.Pointer(ret),
	}
}

func (m *TCEFWindowComponent) ClientAreaBoundsInScreen() (result TCefRect) {
	imports.Proc(internale_CEFWindowComponent_ClientAreaBoundsInScreen).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TCEFWindowComponent) WindowHandle() consts.TCefWindowHandle {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_WindowHandle).Call(uintptr(m.instance))
	return consts.TCefWindowHandle(r1)
}

func (m *TCEFWindowComponent) IsClosed() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsClosed).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsActive() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsActive).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsAlwaysOnTop() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsAlwaysOnTop).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsFullscreen() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsFullscreen).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsMaximized() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsMaximized).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsMinimized() bool {
	r1, _, _ := imports.Proc(internale_CEFWindowComponent_IsMinimized).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) AddChildView(browserViewComponent *TCEFBrowserViewComponent) {
	imports.Proc(internale_CEFWindowComponent_AddChildView).Call(uintptr(m.instance), browserViewComponent.Instance())
}

func (m *TCEFWindowComponent) SetOnWindowCreated(fn WindowComponentOnWindowCreated) {
	imports.Proc(internale_CEFWindowComponent_SetOnWindowCreated).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnWindowDestroyed(fn WindowComponentOnWindowDestroyed) {
	imports.Proc(internale_CEFWindowComponent_SetOnWindowDestroyed).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnWindowActivationChanged(fn WindowComponentOnWindowActivationChanged) {
	imports.Proc(internale_CEFWindowComponent_SetOnWindowActivationChanged).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnGetParentWindow(fn WindowComponentOnGetParentWindow) {
	imports.Proc(internale_CEFWindowComponent_SetOnGetParentWindow).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnGetInitialBounds(fn WindowComponentOnGetInitialBounds) {
	imports.Proc(internale_CEFWindowComponent_SetOnGetInitialBounds).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnGetInitialShowState(fn WindowComponentOnGetInitialShowState) {
	imports.Proc(internale_CEFWindowComponent_SetOnGetInitialShowState).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnIsFrameless(fn WindowComponentOnIsFrameless) {
	imports.Proc(internale_CEFWindowComponent_SetOnIsFrameless).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnCanResize(fn WindowComponentOnCanResize) {
	imports.Proc(internale_CEFWindowComponent_SetOnCanResize).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnCanMaximize(fn WindowComponentOnCanMaximize) {
	imports.Proc(internale_CEFWindowComponent_SetOnCanMaximize).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnCanMinimize(fn WindowComponentOnCanMinimize) {
	imports.Proc(internale_CEFWindowComponent_SetOnCanMinimize).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnCanClose(fn WindowComponentOnCanClose) {
	imports.Proc(internale_CEFWindowComponent_SetOnCanClose).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnAccelerator(fn WindowComponentOnAccelerator) {
	imports.Proc(internale_CEFWindowComponent_SetOnAccelerator).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func (m *TCEFWindowComponent) SetOnKeyEvent(fn WindowComponentOnKeyEvent) {
	imports.Proc(internale_CEFWindowComponent_SetOnKeyEvent).Call(uintptr(m.instance), api.MakeEventDataPtr(fn))
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case WindowComponentOnWindowCreated:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnWindowCreated)(lcl.AsObject(sender), &ICefWindow{instance: window})
		case WindowComponentOnWindowDestroyed:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnWindowDestroyed)(lcl.AsObject(sender), &ICefWindow{instance: window})
		case WindowComponentOnWindowActivationChanged:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnWindowActivationChanged)(lcl.AsObject(sender), &ICefWindow{instance: window}, api.GoBool(getVal(2)))
		case WindowComponentOnGetParentWindow:
			sender := getPtr(0)
			window := getPtr(1)
			resultWindowPtr := (*uintptr)(getPtr(4))
			resultWindow := &ICefWindow{}
			fn.(WindowComponentOnGetParentWindow)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)), (*bool)(getPtr(3)), resultWindow)
			*resultWindowPtr = uintptr(resultWindow.instance)
		case WindowComponentOnGetInitialBounds:
			sender := getPtr(0)
			window := getPtr(1)
			resultRectPtr := (*tCefRectPtr)(getPtr(2))
			resultRect := new(TCefRect)
			resultRect.X = 0
			resultRect.Y = 0
			resultRect.Width = 0
			resultRect.Height = 0
			fn.(WindowComponentOnGetInitialBounds)(lcl.AsObject(sender), &ICefWindow{instance: window}, resultRect)
			resultRectPtr.X = uintptr(resultRect.X)
			resultRectPtr.Y = uintptr(resultRect.Y)
			resultRectPtr.Width = uintptr(resultRect.Width)
			resultRectPtr.Height = uintptr(resultRect.Height)
		case WindowComponentOnGetInitialShowState:
			sender := getPtr(0)
			window := getPtr(1)
			resultShowState := (*consts.TCefShowState)(getPtr(2))
			fn.(WindowComponentOnGetInitialShowState)(lcl.AsObject(sender), &ICefWindow{instance: window}, resultShowState)
		case WindowComponentOnIsFrameless:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnIsFrameless)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnCanResize:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnCanResize)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnCanMaximize:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnCanMaximize)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnCanMinimize:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnCanMinimize)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnCanClose:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnCanClose)(lcl.AsObject(sender), &ICefWindow{instance: window}, (*bool)(getPtr(2)))
		case WindowComponentOnAccelerator:
			sender := getPtr(0)
			window := getPtr(1)
			fn.(WindowComponentOnAccelerator)(lcl.AsObject(sender), &ICefWindow{instance: window}, int32(getVal(2)), (*bool)(getPtr(3)))
		case WindowComponentOnKeyEvent:
			sender := getPtr(0)
			window := getPtr(1)
			keyEvent := (*TCefKeyEvent)(getPtr(2))
			fn.(WindowComponentOnKeyEvent)(lcl.AsObject(sender), &ICefWindow{instance: window}, keyEvent, (*bool)(getPtr(3)))
		default:
			return false
		}
		return true
	})
}
