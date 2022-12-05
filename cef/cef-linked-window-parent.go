package cef

import (
	. "github.com/energye/energy/common"
	"github.com/energye/energy/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type TCEFLinkedWindowParent struct {
	BaseWinControl
}

func NewCEFLinkedWindowParent(owner lcl.IComponent) *TCEFLinkedWindowParent {
	m := new(TCEFLinkedWindowParent)
	r1, _, _ := Proc(internale_CEFLinkedWindow_Create).Call(lcl.CheckPtr(owner))
	m.instance = unsafe.Pointer(r1)
	return m
}

func (m *TCEFLinkedWindowParent) Handle() types.HWND {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetHandle).Call(uintptr(m.Instance()))
	return types.HWND(ret)
}

func (m *TCEFLinkedWindowParent) UpdateSize() {
	Proc(internale_CEFLinkedWindow_UpdateSize).Call(uintptr(m.Instance()))
}

func (m *TCEFLinkedWindowParent) Type() consts.TCefWindowHandleType {
	return consts.Wht_LinkedWindowParent
}

func (m *TCEFLinkedWindowParent) SetChromium(chromium IChromium, tag int32) {
	Proc(internale_CEFLinkedWindow_SetChromium).Call(uintptr(m.instance), chromium.Instance(), uintptr(tag))
}

func (m *TCEFLinkedWindowParent) HandleAllocated() bool {
	ret, _, _ := Proc(internale_CEFLinkedWindow_HandleAllocated).Call(uintptr(m.Instance()))
	return api.GoBool(ret)
}

func (m *TCEFLinkedWindowParent) CreateHandle() {
	Proc(internale_CEFLinkedWindow_CreateHandle).Call(uintptr(m.Instance()))
}

func (m *TCEFLinkedWindowParent) DestroyChildWindow() bool {
	ret, _, _ := Proc(internale_CEFLinkedWindow_DestroyChildWindow).Call(uintptr(m.Instance()))
	return api.GoBool(ret)
}

func (m *TCEFLinkedWindowParent) SetOnEnter(fn lcl.TNotifyEvent) {
	Proc(internale_CEFLinkedWindow_OnEnter).Call(uintptr(m.Instance()), api.MakeEventDataPtr(fn))
}

func (m *TCEFLinkedWindowParent) SetOnExit(fn lcl.TNotifyEvent) {
	Proc(internale_CEFLinkedWindow_OnExit).Call(uintptr(m.Instance()), api.MakeEventDataPtr(fn))
}

func (m *TCEFLinkedWindowParent) Free() {
	if m.IsValid() {
		Proc(internale_CEFLinkedWindow_Free).Call(uintptr(m.Instance()))
		m.instance = nullptr
	}
}

// 获取组件名称。
func (m *TCEFLinkedWindowParent) Name() string {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetName).Call(uintptr(m.Instance()))
	return api.GoStr(ret)
}

// 设置组件名称。
func (m *TCEFLinkedWindowParent) SetName(value string) {
	Proc(internale_CEFLinkedWindow_SetName).Call(uintptr(m.Instance()), api.PascalStr(value))
}

// 设置控件父容器。
func (m *TCEFLinkedWindowParent) SetParent(value lcl.IWinControl) {
	Proc(internale_CEFLinkedWindow_SetParent).Call(uintptr(m.Instance()), lcl.CheckPtr(value))
}

//Align 获取控件自动调整。
func (m *TCEFLinkedWindowParent) Align() types.TAlign {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetAlign).Call(uintptr(m.Instance()))
	return types.TAlign(ret)
}

//SetAlign 设置控件自动调整。
func (m *TCEFLinkedWindowParent) SetAlign(value types.TAlign) {
	Proc(internale_CEFLinkedWindow_SetAlign).Call(uintptr(m.Instance()), uintptr(value))
}

//Anchors 获取四个角位置的锚点。
func (m *TCEFLinkedWindowParent) Anchors() types.TAnchors {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetAnchors).Call(uintptr(m.Instance()))
	return types.TAnchors(ret)
}

//SetAnchors 设置四个角位置的锚点。
func (m *TCEFLinkedWindowParent) SetAnchors(value types.TAnchors) {
	Proc(internale_CEFLinkedWindow_SetAnchors).Call(uintptr(m.Instance()), uintptr(value))
}

//Visible 获取控件可视。
func (m *TCEFLinkedWindowParent) Visible() bool {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetVisible).Call(uintptr(m.Instance()))
	return api.GoBool(ret)
}

//SetVisible 设置控件可视。
func (m *TCEFLinkedWindowParent) SetVisible(value bool) {
	Proc(internale_CEFLinkedWindow_SetVisible).Call(uintptr(m.Instance()), api.PascalBool(value))
}

//Enabled 获取是否启用
func (m *TCEFLinkedWindowParent) Enabled() bool {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetEnabled).Call(uintptr(m.Instance()))
	return api.GoBool(ret)
}

//SetEnabled 设置是否启用
func (m *TCEFLinkedWindowParent) SetEnabled(value bool) {
	Proc(internale_CEFLinkedWindow_SetEnabled).Call(uintptr(m.Instance()), api.PascalBool(value))
}

//Left 获取左边距
func (m *TCEFLinkedWindowParent) Left() int32 {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetLeft).Call(uintptr(m.Instance()))
	return int32(ret)
}

//SetLeft 设置左边距
func (m *TCEFLinkedWindowParent) SetLeft(value int32) {
	Proc(internale_CEFLinkedWindow_SetLeft).Call(uintptr(m.Instance()), uintptr(value))
}

//Top 获取上边距
func (m *TCEFLinkedWindowParent) Top() int32 {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetTop).Call(uintptr(m.Instance()))
	return int32(ret)
}

//SetTop 设置上边距
func (m *TCEFLinkedWindowParent) SetTop(value int32) {
	Proc(internale_CEFLinkedWindow_SetTop).Call(uintptr(m.Instance()), uintptr(value))
}

//Width 获取宽度
func (m *TCEFLinkedWindowParent) Width() int32 {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetWidth).Call(uintptr(m.Instance()))
	return int32(ret)
}

//SetWidth 设置宽度
func (m *TCEFLinkedWindowParent) SetWidth(value int32) {
	Proc(internale_CEFLinkedWindow_SetWidth).Call(uintptr(m.Instance()), uintptr(value))
}

//Height 获取高度
func (m *TCEFLinkedWindowParent) Height() int32 {
	ret, _, _ := Proc(internale_CEFLinkedWindow_GetHeight).Call(uintptr(m.Instance()))
	return int32(ret)
}

//SetHeight 设置高度
func (m *TCEFLinkedWindowParent) SetHeight(value int32) {
	Proc(internale_CEFLinkedWindow_SetHeight).Call(uintptr(m.Instance()), uintptr(value))
}

func (m *TCEFLinkedWindowParent) BoundsRect() (result types.TRect) {
	Proc(internale_CEFLinkedWindow_GetBoundsRect).Call(uintptr(m.Instance()), uintptr(unsafe.Pointer(&result)))
	return result
}

func (m *TCEFLinkedWindowParent) SetBoundsRect(value types.TRect) {
	Proc(internale_CEFLinkedWindow_SetBoundsRect).Call(uintptr(m.Instance()), uintptr(unsafe.Pointer(&value)))
}
