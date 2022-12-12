package cef

import (
	. "github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
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
	Proc(internale_CEFWindowComponent_ShowMenu).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) CancelMenu() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SetDraggableRegions() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SendKeyPress() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SendMouseMove() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SendMouseEvents() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) SetAccelerator() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) RemoveAccelerator() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) RemoveAllAccelerators() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Title() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) WindowIcon() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) WindowAppIcon() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) Display() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}

func (m *TCEFWindowComponent) ClientAreaBoundsInScreen() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) WindowHandle() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) IsClosed() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) IsActive() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) IsAlwaysOnTop() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) IsFullscreen() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) IsMaximized() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) IsMinimized() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnWindowCreated() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnWindowDestroyed() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnWindowActivationChanged() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnGetParentWindow() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnGetInitialBounds() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnGetInitialShowState() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnIsFrameless() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnCanResize() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnCanMaximize() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnCanMinimize() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnCanClose() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnAccelerator() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
func (m *TCEFWindowComponent) SetOnKeyEvent() {
	Proc(internale_CEFWindowComponent_BringToTop).Call(uintptr(m.instance))
}
