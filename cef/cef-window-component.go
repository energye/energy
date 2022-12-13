package cef

import (
	"fmt"
	. "github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type TCEFWindowComponent struct {
	instance unsafe.Pointer
}

func NewWindowComponent(AOwner lcl.IComponent) *TCEFWindowComponent {
	r1, _, _ := Proc(internale_CEFWindowComponent_Create).Call(lcl.CheckPtr(AOwner))
	return &TCEFWindowComponent{
		instance: unsafe.Pointer(r1),
	}
}

func (m *TCEFWindowComponent) CreateTopLevelWindow() {
	Proc(internale_CEFWindowComponent_CreateTopLevelWindow).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Show() {
	Proc(internale_CEFWindowComponent_Show).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Hide() {
	Proc(internale_CEFWindowComponent_Hide).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) CenterWindow(size TCefSize) {
	Proc(internale_CEFWindowComponent_CenterWindow).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&size)))
}

func (m *TCEFWindowComponent) Close() {
	Proc(internale_CEFWindowComponent_Close).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Activate() {
	Proc(internale_CEFWindowComponent_Activate).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Deactivate() {
	Proc(internale_CEFWindowComponent_Deactivate).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) BringToTop() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Maximize() {
	Proc(internale_CEFWindowComponent_Maximize).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Minimize() {
	Proc(internale_CEFWindowComponent_Minimize).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Restore() {
	Proc(internale_CEFWindowComponent_Restore).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) AddOverlayView() {
	//Proc(internale_CEFWindowComponent_AddOverlayView).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) ShowMenu(menuModel *ICefMenuModel, point TCefPoint, anchorPosition consts.TCefMenuAnchorPosition) {
	Proc(internale_CEFWindowComponent_ShowMenu).Call(uintptr(m.instance), uintptr(menuModel.instance), uintptr(unsafe.Pointer(&point)), uintptr(anchorPosition))
}

func (m *TCEFWindowComponent) CancelMenu() {
	Proc(internale_CEFWindowComponent_CancelMenu).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SetDraggableRegions(regionsCount int32, regions []TCefDraggableRegion) {
	Proc(internale_CEFWindowComponent_SetDraggableRegions).Call(uintptr(m.instance), uintptr(regionsCount), uintptr(unsafe.Pointer(&regions[0])), uintptr(int32(len(regions))))
}

func (m *TCEFWindowComponent) SendKeyPress(keyCode int32, eventFlags uint32) {
	Proc(internale_CEFWindowComponent_SendKeyPress).Call(uintptr(m.instance), uintptr(keyCode), uintptr(eventFlags))
}

func (m *TCEFWindowComponent) SendMouseMove(screenX, screenY int32) {
	Proc(internale_CEFWindowComponent_SendMouseMove).Call(uintptr(m.instance), uintptr(screenX), uintptr(screenY))
}

func (m *TCEFWindowComponent) SendMouseEvents(button consts.TCefMouseButtonType, mouseDown, mouseUp bool) {
	Proc(internale_CEFWindowComponent_SendMouseEvents).Call(uintptr(m.instance), uintptr(button), api.PascalBool(mouseDown), api.PascalBool(mouseUp))
}

func (m *TCEFWindowComponent) SetAccelerator(commandId, keyCode int32, shiftPressed, ctrlPressed, altPressed bool) {
	Proc(internale_CEFWindowComponent_SetAccelerator).Call(uintptr(m.instance), uintptr(commandId), uintptr(keyCode), api.PascalBool(shiftPressed), api.PascalBool(ctrlPressed), api.PascalBool(altPressed))
}

func (m *TCEFWindowComponent) RemoveAccelerator(commandId int32) {
	Proc(internale_CEFWindowComponent_RemoveAccelerator).Call(uintptr(m.instance), uintptr(commandId))
}

func (m *TCEFWindowComponent) RemoveAllAccelerators() {
	Proc(internale_CEFWindowComponent_RemoveAllAccelerators).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SetTitle(title string) {
	Proc(internale_CEFWindowComponent_SetTitle).Call(uintptr(m.instance), api.PascalStr(title))
}

func (m *TCEFWindowComponent) Title() string {
	r1, _, _ := Proc(internale_CEFWindowComponent_Title).Call(uintptr(m.instance))
	return api.GoStr(r1)
}

func (m *TCEFWindowComponent) WindowIcon() *ICefImage {
	var ret uintptr
	Proc(internale_CEFWindowComponent_WindowIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefImage{
		instance: unsafe.Pointer(ret),
	}
}

func (m *TCEFWindowComponent) WindowAppIcon() *ICefImage {
	var ret uintptr
	Proc(internale_CEFWindowComponent_WindowAppIcon).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefImage{
		instance: unsafe.Pointer(ret),
	}
}

func (m *TCEFWindowComponent) Display() *ICefDisplay {
	var ret uintptr
	Proc(internale_CEFWindowComponent_Display).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&ret)))
	return &ICefDisplay{
		instance: unsafe.Pointer(ret),
	}
}

func (m *TCEFWindowComponent) ClientAreaBoundsInScreen() (result TCefRect) {
	Proc(internale_CEFWindowComponent_ClientAreaBoundsInScreen).Call(uintptr(m.instance), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TCEFWindowComponent) WindowHandle() consts.TCefWindowHandle {
	r1, _, _ := Proc(internale_CEFWindowComponent_WindowHandle).Call(uintptr(m.instance))
	return consts.TCefWindowHandle(r1)
}

func (m *TCEFWindowComponent) IsClosed() bool {
	r1, _, _ := Proc(internale_CEFWindowComponent_IsClosed).Call(uintptr(m.instance))
	fmt.Println(r1)
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsActive() bool {
	r1, _, _ := Proc(internale_CEFWindowComponent_IsActive).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsAlwaysOnTop() bool {
	r1, _, _ := Proc(internale_CEFWindowComponent_IsAlwaysOnTop).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsFullscreen() bool {
	r1, _, _ := Proc(internale_CEFWindowComponent_IsFullscreen).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsMaximized() bool {
	r1, _, _ := Proc(internale_CEFWindowComponent_IsMaximized).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) IsMinimized() bool {
	r1, _, _ := Proc(internale_CEFWindowComponent_IsMinimized).Call(uintptr(m.instance))
	return api.GoBool(r1)
}

func (m *TCEFWindowComponent) SetOnWindowCreated() {
	Proc(internale_CEFWindowComponent_SetOnWindowCreated).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnWindowDestroyed() {
	Proc(internale_CEFWindowComponent_SetOnWindowDestroyed).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnWindowActivationChanged() {
	Proc(internale_CEFWindowComponent_SetOnWindowActivationChanged).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnGetParentWindow() {
	Proc(internale_CEFWindowComponent_SetOnGetParentWindow).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnGetInitialBounds() {
	Proc(internale_CEFWindowComponent_SetOnGetInitialBounds).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnGetInitialShowState() {
	Proc(internale_CEFWindowComponent_SetOnGetInitialShowState).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnIsFrameless() {
	Proc(internale_CEFWindowComponent_SetOnIsFrameless).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnCanResize() {
	Proc(internale_CEFWindowComponent_SetOnCanResize).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnCanMaximize() {
	Proc(internale_CEFWindowComponent_SetOnCanMaximize).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnCanMinimize() {
	Proc(internale_CEFWindowComponent_SetOnCanMinimize).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnCanClose() {
	Proc(internale_CEFWindowComponent_SetOnCanClose).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnAccelerator() {
	Proc(internale_CEFWindowComponent_SetOnAccelerator).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnKeyEvent() {
	Proc(internale_CEFWindowComponent_SetOnKeyEvent).Call(uintptr(m.instance))
}
