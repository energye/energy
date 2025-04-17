//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// CEF TBufferPanel

package cef

import (
	"github.com/cyber-xxm/energy/v2/cef/internal/def"
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type TBufferPanel struct {
	lcl.IWinControl
	instance unsafe.Pointer
}

// NewBufferPanel
//
//	创建一个新的对象。
func NewBufferPanel(owner lcl.IComponent) *TBufferPanel {
	m := new(TBufferPanel)
	var result uintptr
	imports.SysCallN(def.BufferPanel_Create, owner.Instance(), uintptr(unsafe.Pointer(&result)))
	m.instance = unsafe.Pointer(result)
	return m
}

// AsBufferPanel
//
// 动态转换一个已存在的对象实例。
//
// Dynamically convert an existing object instance.
func AsBufferPanel(obj interface{}) *TBufferPanel {
	instance := getInstance(obj)
	if instance == nil {
		return nil
	}
	return &TBufferPanel{instance: instance}
}

// Free
//
// 释放对象。
//
// Free object.
func (m *TBufferPanel) Free() {
	if m.instance != nil {
		imports.SysCallN(def.BufferPanel_Free, m.Instance())
		m.instance = nil
	}
}

// Instance
//
// 返回对象实例指针。
//
// Return object instance pointer.
func (m *TBufferPanel) Instance() uintptr {
	return uintptr(m.instance)
}

// IsValid
//
// 检测地址是否为空。
//
// Check if the address is empty.
func (m *TBufferPanel) IsValid() bool {
	return m.instance != nil
}

// TBufferPanelClass
//
// 获取类信息指针。
//
// Get class information pointer.
func TBufferPanelClass() types.TClass {
	return types.TClass(imports.SysCallN(def.BufferPanel_StaticClassType))
}

// CanFocus
//
// 是否可以获得焦点。
func (m *TBufferPanel) CanFocus() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_CanFocus, m.Instance()))
}

// ContainsControl
//
// 返回是否包含指定控件。
//
// it's contain a specified control.
func (m *TBufferPanel) ContainsControl(Control lcl.IControl) bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_ContainsControl, m.Instance(), lcl.CheckPtr(Control)))
}

// ControlAtPos
//
// 返回指定坐标及相关属性位置控件。
//
// Returns the specified coordinate and the relevant attribute position control..
func (m *TBufferPanel) ControlAtPos(Pos types.TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *lcl.TControl {
	return lcl.AsControl(imports.SysCallN(def.BufferPanel_ControlAtPos, m.Instance(), uintptr(unsafe.Pointer(&Pos)), api.PascalBool(AllowDisabled), api.PascalBool(AllowWinControls), api.PascalBool(AllLevels)))
}

// DisableAlign
//
// 禁用控件的对齐。
//
// Disable control alignment.
func (m *TBufferPanel) DisableAlign() {
	imports.SysCallN(def.BufferPanel_DisableAlign, m.Instance())
}

// EnableAlign
//
// 启用控件对齐。
//
// Enabled control alignment.
func (m *TBufferPanel) EnableAlign() {
	imports.SysCallN(def.BufferPanel_EnableAlign, m.Instance())
}

// FindChildControl
//
// 查找子控件。
//
// Find sub controls.
func (m *TBufferPanel) FindChildControl(ControlName string) *lcl.TControl {
	return lcl.AsControl(imports.SysCallN(def.BufferPanel_FindChildControl, m.Instance(), api.PascalStr(ControlName)))
}

func (m *TBufferPanel) FlipChildren(AllLevels bool) {
	imports.SysCallN(def.BufferPanel_FlipChildren, m.Instance(), api.PascalBool(AllLevels))
}

// Focused
//
// 返回是否获取焦点。
//
// Return to get focus.
func (m *TBufferPanel) Focused() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_Focused, m.Instance()))
}

// HandleAllocated
//
// 句柄是否已经分配。
//
// Is the handle already allocated.
func (m *TBufferPanel) HandleAllocated() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_HandleAllocated, m.Instance()))
}

// InsertControl
//
// 插入一个控件。
//
// Insert a control.
func (m *TBufferPanel) InsertControl(AControl lcl.IControl) {
	imports.SysCallN(def.BufferPanel_InsertControl, m.Instance(), lcl.CheckPtr(AControl))
}

// Invalidate
//
// 要求重绘。
//
// Redraw.
func (m *TBufferPanel) Invalidate() {
	imports.SysCallN(def.BufferPanel_Invalidate, m.Instance())
}

// PaintTo
//
// 绘画至指定DC。
//
// Painting to the specified DC.
func (m *TBufferPanel) PaintTo(DC types.HDC, X int32, Y int32) {
	imports.SysCallN(def.BufferPanel_PaintTo, m.Instance(), DC, uintptr(X), uintptr(Y))
}

// RemoveControl
//
// 移除一个控件。
//
// Remove a control.
func (m *TBufferPanel) RemoveControl(AControl lcl.IControl) {
	imports.SysCallN(def.BufferPanel_RemoveControl, m.Instance(), lcl.CheckPtr(AControl))
}

// Realign
//
// 重新对齐。
//
// Realign.
func (m *TBufferPanel) Realign() {
	imports.SysCallN(def.BufferPanel_Realign, m.Instance())
}

// Repaint
//
// 重绘。
//
// Repaint.
func (m *TBufferPanel) Repaint() {
	imports.SysCallN(def.BufferPanel_Repaint, m.Instance())
}

// ScaleBy
//
// 按比例缩放。
//
// Scale by.
func (m *TBufferPanel) ScaleBy(M int32, D int32) {
	imports.SysCallN(def.BufferPanel_ScaleBy, m.Instance(), uintptr(M), uintptr(D))
}

// ScrollBy
//
// 滚动至指定位置。
//
// Scroll by.
func (m *TBufferPanel) ScrollBy(DeltaX int32, DeltaY int32) {
	imports.SysCallN(def.BufferPanel_ScrollBy, m.Instance(), uintptr(DeltaX), uintptr(DeltaY))
}

// SetBounds
//
// 设置组件边界。
//
// Set component boundaries.
func (m *TBufferPanel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	imports.SysCallN(def.BufferPanel_SetBounds, m.Instance(), uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

// SetFocus
//
// 设置控件焦点。
//
// Set control focus.
func (m *TBufferPanel) SetFocus() {
	imports.SysCallN(def.BufferPanel_SetFocus, m.Instance())
}

// Update
//
// 控件更新。
//
// Update.
func (m *TBufferPanel) Update() {
	imports.SysCallN(def.BufferPanel_Update, m.Instance())
}

// BringToFront
//
// 将控件置于最前。
//
// Bring the control to the front.
func (m *TBufferPanel) BringToFront() {
	imports.SysCallN(def.BufferPanel_BringToFront, m.Instance())
}

// ClientToScreen
//
// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (m *TBufferPanel) ClientToScreen(Point types.TPoint) (result types.TPoint) {
	imports.SysCallN(def.BufferPanel_ClientToScreen, m.Instance(), uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&result)))
	return
}

// ClientToParent
//
// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (m *TBufferPanel) ClientToParent(Point types.TPoint, AParent lcl.IWinControl) (result types.TPoint) {
	imports.SysCallN(def.BufferPanel_ClientToParent, m.Instance(), uintptr(unsafe.Pointer(&Point)), lcl.CheckPtr(AParent), uintptr(unsafe.Pointer(&result)))
	return
}

// Dragging
//
// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (m *TBufferPanel) Dragging() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_Dragging, m.Instance()))
}

// HasParent
//
// 是否有父容器。
//
// Is there a parent container.
func (m *TBufferPanel) HasParent() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_HasParent, m.Instance()))
}

// Hide
//
// 隐藏控件。
//
// Hidden control.
func (m *TBufferPanel) Hide() {
	imports.SysCallN(def.BufferPanel_Hide, m.Instance())
}

// Perform
//
// 发送一个消息。
//
// Send a message.
func (m *TBufferPanel) Perform(Msg uint32, WParam uintptr, LParam int) int {
	return int(imports.SysCallN(def.BufferPanel_Perform, m.Instance(), uintptr(Msg), WParam, uintptr(LParam)))
}

// Refresh
//
// 刷新控件。
//
// Refresh control.
func (m *TBufferPanel) Refresh() {
	imports.SysCallN(def.BufferPanel_Refresh, m.Instance())
}

// ScreenToClient
//
// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (m *TBufferPanel) ScreenToClient(Point types.TPoint) (result types.TPoint) {
	imports.SysCallN(def.BufferPanel_ScreenToClient, m.Instance(), uintptr(unsafe.Pointer(&Point)))
	return
}

// ParentToClient
//
// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (m *TBufferPanel) ParentToClient(Point types.TPoint, AParent lcl.IWinControl) (result types.TPoint) {
	imports.SysCallN(def.BufferPanel_ParentToClient, m.Instance(), uintptr(unsafe.Pointer(&Point)), lcl.CheckPtr(AParent), uintptr(unsafe.Pointer(&result)))
	return
}

// SendToBack
//
// 控件至于最后面。
//
// The control is placed at the end.
func (m *TBufferPanel) SendToBack() {
	imports.SysCallN(def.BufferPanel_SendToBack, m.Instance())
}

// Show
//
// 显示控件。
//
// Show control.
func (m *TBufferPanel) Show() {
	imports.SysCallN(def.BufferPanel_Show, m.Instance())
}

// GetTextBuf
//
// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (m *TBufferPanel) GetTextBuf(Buffer *string, BufSize int32) (sLen int32) {
	if Buffer == nil || BufSize == 0 {
		return
	}
	strPtr := make([]uint8, BufSize+1)
	sLen = int32(imports.SysCallN(def.BufferPanel_GetTextBuf, m.Instance(), uintptr(unsafe.Pointer(&strPtr[0])), uintptr(BufSize)))
	if sLen > 0 {
		*Buffer = string(strPtr[:sLen])
	}
	return
}

// GetTextLen
//
// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (m *TBufferPanel) GetTextLen() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetTextLen, m.Instance()))
}

// SetTextBuf
//
// 设置控件字符，如果有。
//
// Set control characters, if any.
func (m *TBufferPanel) SetTextBuf(Buffer string) {
	imports.SysCallN(def.BufferPanel_SetTextBuf, m.Instance(), api.PascalStr(Buffer))
}

// FindComponent
//
// 查找指定名称的组件。
//
// Find the component with the specified name.
func (m *TBufferPanel) FindComponent(AName string) *lcl.TComponent {
	return lcl.AsComponent(imports.SysCallN(def.BufferPanel_FindComponent, m.Instance(), api.PascalStr(AName)))
}

// GetNamePath
//
// 获取类名路径。
//
// Get the class name path.
func (m *TBufferPanel) GetNamePath() string {
	return api.GoStr(imports.SysCallN(def.BufferPanel_GetNamePath, m.Instance()))
}

// Assign
//
// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (m *TBufferPanel) Assign(Source lcl.IObject) {
	imports.SysCallN(def.BufferPanel_Assign, m.Instance(), lcl.CheckPtr(Source))
}

// ClassType
//
// 获取类的类型信息。
//
// Get class type information.
func (m *TBufferPanel) ClassType() types.TClass {
	return types.TClass(imports.SysCallN(def.BufferPanel_ClassType, m.Instance()))
}

// ClassName
//
// 获取当前对象类名称。
//
// Get the current object class name.
func (m *TBufferPanel) ClassName() string {
	return api.GoStr(imports.SysCallN(def.BufferPanel_ClassName, m.Instance()))
}

// InstanceSize
//
// 获取当前对象实例大小。
//
// Get the current object instance size.
func (m *TBufferPanel) InstanceSize() int32 {
	return int32(imports.SysCallN(def.BufferPanel_InstanceSize, m.Instance()))
}

// InheritsFrom
//
// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (m *TBufferPanel) InheritsFrom(AClass types.TClass) bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_InheritsFrom, m.Instance(), uintptr(AClass)))
}

// Equals
//
// 与一个对象进行比较。
//
// Compare with an object.
func (m *TBufferPanel) Equals(Obj lcl.IObject) bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_Equals, m.Instance(), lcl.CheckPtr(Obj)))
}

// GetHashCode
//
// 获取类的哈希值。
//
// Get the hash value of the class.
func (m *TBufferPanel) GetHashCode() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetHashCode, m.Instance()))
}

// ToString
//
// 文本类信息。
//
// Text information.
func (m *TBufferPanel) ToString() string {
	return api.GoStr(imports.SysCallN(def.BufferPanel_ToString, m.Instance()))
}

func (m *TBufferPanel) AnchorToNeighbour(ASide types.TAnchorKind, ASpace int32, ASibling lcl.IControl) {
	imports.SysCallN(def.BufferPanel_AnchorToNeighbour, m.Instance(), uintptr(ASide), uintptr(ASpace), lcl.CheckPtr(ASibling))
}

func (m *TBufferPanel) AnchorParallel(ASide types.TAnchorKind, ASpace int32, ASibling lcl.IControl) {
	imports.SysCallN(def.BufferPanel_AnchorParallel, m.Instance(), uintptr(ASide), uintptr(ASpace), lcl.CheckPtr(ASibling))
}

// AnchorHorizontalCenterTo
//
// 置于指定控件的横向中心。
func (m *TBufferPanel) AnchorHorizontalCenterTo(ASibling lcl.IControl) {
	imports.SysCallN(def.BufferPanel_AnchorHorizontalCenterTo, m.Instance(), lcl.CheckPtr(ASibling))
}

// AnchorVerticalCenterTo
//
// 置于指定控件的纵向中心。
func (m *TBufferPanel) AnchorVerticalCenterTo(ASibling lcl.IControl) {
	imports.SysCallN(def.BufferPanel_AnchorVerticalCenterTo, m.Instance(), lcl.CheckPtr(ASibling))
}

func (m *TBufferPanel) AnchorSame(ASide types.TAnchorKind, ASibling lcl.IControl) {
	imports.SysCallN(def.BufferPanel_AnchorSame, m.Instance(), uintptr(ASide), lcl.CheckPtr(ASibling))
}

func (m *TBufferPanel) AnchorAsAlign(ATheAlign types.TAlign, ASpace int32) {
	imports.SysCallN(def.BufferPanel_AnchorAsAlign, m.Instance(), uintptr(ATheAlign), uintptr(ASpace))
}

func (m *TBufferPanel) AnchorClient(ASpace int32) {
	imports.SysCallN(def.BufferPanel_AnchorClient, m.Instance(), uintptr(ASpace))
}

func (m *TBufferPanel) ScaleDesignToForm(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_ScaleDesignToForm, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) ScaleFormToDesign(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_ScaleFormToDesign, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) Scale96ToForm(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_Scale96ToForm, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) ScaleFormTo96(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_ScaleFormTo96, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) Scale96ToFont(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_Scale96ToFont, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) ScaleFontTo96(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_ScaleFontTo96, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) ScaleScreenToFont(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_ScaleScreenToFont, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) ScaleFontToScreen(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_ScaleFontToScreen, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) Scale96ToScreen(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_Scale96ToScreen, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) ScaleScreenTo96(ASize int32) int32 {
	return int32(imports.SysCallN(def.BufferPanel_ScaleScreenTo96, m.Instance(), uintptr(ASize)))
}

func (m *TBufferPanel) AutoAdjustLayout(AMode types.TLayoutAdjustmentPolicy, AFromPPI int32, AToPPI int32, AOldFormWidth int32, ANewFormWidth int32) {
	imports.SysCallN(def.BufferPanel_AutoAdjustLayout, m.Instance(), uintptr(AMode), uintptr(AFromPPI), uintptr(AToPPI), uintptr(AOldFormWidth), uintptr(ANewFormWidth))
}

func (m *TBufferPanel) FixDesignFontsPPI(ADesignTimePPI int32) {
	imports.SysCallN(def.BufferPanel_FixDesignFontsPPI, m.Instance(), uintptr(ADesignTimePPI))
}

func (m *TBufferPanel) ScaleFontsPPI(AToPPI int32, AProportion float64) {
	imports.SysCallN(def.BufferPanel_ScaleFontsPPI, m.Instance(), uintptr(AToPPI), uintptr(unsafe.Pointer(&AProportion)))
}

// Canvas
//
// 获取画布。
func (m *TBufferPanel) Canvas() *lcl.TCanvas {
	return lcl.AsCanvas(imports.SysCallN(def.BufferPanel_GetCanvas, m.Instance()))
}

// SetCanvas
//
// 设置画布。
func (m *TBufferPanel) SetCanvas(value *lcl.TCanvas) {
	imports.SysCallN(def.BufferPanel_SetCanvas, m.Instance(), lcl.CheckPtr(value))
}

// SetOnPaint
//
// 设置绘画事件。
func (m *TBufferPanel) SetOnPaint(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnPaint, m.Instance(), api.MakeEventDataPtr(fn))
}

// Align
//
// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (m *TBufferPanel) Align() types.TAlign {
	return types.TAlign(imports.SysCallN(def.BufferPanel_GetAlign, m.Instance()))
}

// SetAlign
//
// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (m *TBufferPanel) SetAlign(value types.TAlign) {
	imports.SysCallN(def.BufferPanel_SetAlign, m.Instance(), uintptr(value))
}

// Alignment
//
// 获取文字对齐。
//
// Get Text alignment.
func (m *TBufferPanel) Alignment() types.TAlignment {
	return types.TAlignment(imports.SysCallN(def.BufferPanel_GetAlignment, m.Instance()))
}

// SetAlignment
//
// 设置文字对齐。
//
// Set Text alignment.
func (m *TBufferPanel) SetAlignment(value types.TAlignment) {
	imports.SysCallN(def.BufferPanel_SetAlignment, m.Instance(), uintptr(value))
}

// Anchors
//
// 获取四个角位置的锚点。
func (m *TBufferPanel) Anchors() types.TAnchors {
	return types.TAnchors(imports.SysCallN(def.BufferPanel_GetAnchors, m.Instance()))
}

// SetAnchors
//
// 设置四个角位置的锚点。
func (m *TBufferPanel) SetAnchors(value types.TAnchors) {
	imports.SysCallN(def.BufferPanel_SetAnchors, m.Instance(), uintptr(value))
}

// AutoSize
//
// 获取自动调整大小。
func (m *TBufferPanel) AutoSize() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetAutoSize, m.Instance()))
}

// SetAutoSize
//
// 设置自动调整大小。
func (m *TBufferPanel) SetAutoSize(value bool) {
	imports.SysCallN(def.BufferPanel_SetAutoSize, m.Instance(), api.PascalBool(value))
}

func (m *TBufferPanel) BevelInner() types.TBevelCut {
	return types.TBevelCut(imports.SysCallN(def.BufferPanel_GetBevelInner, m.Instance()))
}

func (m *TBufferPanel) SetBevelInner(value types.TBevelCut) {
	imports.SysCallN(def.BufferPanel_SetBevelInner, m.Instance(), uintptr(value))
}

func (m *TBufferPanel) BevelOuter() types.TBevelCut {
	return types.TBevelCut(imports.SysCallN(def.BufferPanel_GetBevelOuter, m.Instance()))
}

func (m *TBufferPanel) SetBevelOuter(value types.TBevelCut) {
	imports.SysCallN(def.BufferPanel_SetBevelOuter, m.Instance(), uintptr(value))
}

func (m *TBufferPanel) BiDiMode() types.TBiDiMode {
	return types.TBiDiMode(imports.SysCallN(def.BufferPanel_GetBiDiMode, m.Instance()))
}

func (m *TBufferPanel) SetBiDiMode(value types.TBiDiMode) {
	imports.SysCallN(def.BufferPanel_SetBiDiMode, m.Instance(), uintptr(value))
}

// BorderWidth
//
// 获取边框的宽度。
func (m *TBufferPanel) BorderWidth() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetBorderWidth, m.Instance()))
}

// SetBorderWidth
//
// 设置边框的宽度。
func (m *TBufferPanel) SetBorderWidth(value int32) {
	imports.SysCallN(def.BufferPanel_SetBorderWidth, m.Instance(), uintptr(value))
}

// BorderStyle
//
// 获取窗口边框样式。比如：无边框，单一边框等。
func (m *TBufferPanel) BorderStyle() types.TBorderStyle {
	return types.TBorderStyle(imports.SysCallN(def.BufferPanel_GetBorderStyle, m.Instance()))
}

// SetBorderStyle
//
// 设置窗口边框样式。比如：无边框，单一边框等。
func (m *TBufferPanel) SetBorderStyle(value types.TBorderStyle) {
	imports.SysCallN(def.BufferPanel_SetBorderStyle, m.Instance(), uintptr(value))
}

// Caption
//
// 获取控件标题。
//
// Get the control title.
func (m *TBufferPanel) Caption() string {
	return api.GoStr(imports.SysCallN(def.BufferPanel_GetCaption, m.Instance()))
}

// SetCaption
//
// 设置控件标题。
//
// Set the control title.
func (m *TBufferPanel) SetCaption(value string) {
	imports.SysCallN(def.BufferPanel_SetCaption, m.Instance(), api.PascalStr(value))
}

// Color
//
// 获取颜色。
//
// Get color.
func (m *TBufferPanel) Color() types.TColor {
	return types.TColor(imports.SysCallN(def.BufferPanel_GetColor, m.Instance()))
}

// SetColor
//
// 设置颜色。
//
// Set color.
func (m *TBufferPanel) SetColor(value types.TColor) {
	imports.SysCallN(def.BufferPanel_SetColor, m.Instance(), uintptr(value))
}

// Constraints
//
// 获取约束控件大小。
func (m *TBufferPanel) Constraints() *lcl.TSizeConstraints {
	return lcl.AsSizeConstraints(imports.SysCallN(def.BufferPanel_GetConstraints, m.Instance()))
}

// SetConstraints
//
// 设置约束控件大小。
func (m *TBufferPanel) SetConstraints(value *lcl.TSizeConstraints) {
	imports.SysCallN(def.BufferPanel_SetConstraints, m.Instance(), lcl.CheckPtr(value))
}

// UseDockManager
//
// 获取使用停靠管理。
func (m *TBufferPanel) UseDockManager() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetUseDockManager, m.Instance()))
}

// SetUseDockManager
//
// 设置使用停靠管理。
func (m *TBufferPanel) SetUseDockManager(value bool) {
	imports.SysCallN(def.BufferPanel_SetUseDockManager, m.Instance(), api.PascalBool(value))
}

// DockSite
//
// 获取停靠站点。
//
// Get Docking site.
func (m *TBufferPanel) DockSite() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetDockSite, m.Instance()))
}

// SetDockSite
//
// 设置停靠站点。
//
// Set Docking site.
func (m *TBufferPanel) SetDockSite(value bool) {
	imports.SysCallN(def.BufferPanel_SetDockSite, m.Instance(), api.PascalBool(value))
}

// DoubleBuffered
//
// 获取设置控件双缓冲。
//
// Get Set control double buffering.
func (m *TBufferPanel) DoubleBuffered() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetDoubleBuffered, m.Instance()))
}

// SetDoubleBuffered
//
// 设置设置控件双缓冲。
//
// Set Set control double buffering.
func (m *TBufferPanel) SetDoubleBuffered(value bool) {
	imports.SysCallN(def.BufferPanel_SetDoubleBuffered, m.Instance(), api.PascalBool(value))
}

// DragCursor
//
// 获取设置控件拖拽时的光标。
//
// Get Set the cursor when the control is dragged.
func (m *TBufferPanel) DragCursor() types.TCursor {
	return types.TCursor(imports.SysCallN(def.BufferPanel_GetDragCursor, m.Instance()))
}

// SetDragCursor
//
// 设置设置控件拖拽时的光标。
//
// Set Set the cursor when the control is dragged.
func (m *TBufferPanel) SetDragCursor(value types.TCursor) {
	imports.SysCallN(def.BufferPanel_SetDragCursor, m.Instance(), uintptr(value))
}

// DragKind
//
// 获取拖拽方式。
//
// Get Drag and drom.
func (m *TBufferPanel) DragKind() types.TDragKind {
	return types.TDragKind(imports.SysCallN(def.BufferPanel_GetDragKind, m.Instance()))
}

// SetDragKind
//
// 设置拖拽方式。
//
// Set Drag and drom.
func (m *TBufferPanel) SetDragKind(value types.TDragKind) {
	imports.SysCallN(def.BufferPanel_SetDragKind, m.Instance(), uintptr(value))
}

// DragMode
//
// 获取拖拽模式。
//
// Get Drag mode.
func (m *TBufferPanel) DragMode() types.TDragMode {
	return types.TDragMode(imports.SysCallN(def.BufferPanel_GetDragMode, m.Instance()))
}

// SetDragMode
//
// 设置拖拽模式。
//
// Set Drag mode.
func (m *TBufferPanel) SetDragMode(value types.TDragMode) {
	imports.SysCallN(def.BufferPanel_SetDragMode, m.Instance(), uintptr(value))
}

// Enabled
//
// 获取控件启用。
//
// Get the control enabled.
func (m *TBufferPanel) Enabled() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetEnabled, m.Instance()))
}

// SetEnabled
//
// 设置控件启用。
//
// Set the control enabled.
func (m *TBufferPanel) SetEnabled(value bool) {
	imports.SysCallN(def.BufferPanel_SetEnabled, m.Instance(), api.PascalBool(value))
}

func (m *TBufferPanel) FullRepaint() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetFullRepaint, m.Instance()))
}

func (m *TBufferPanel) SetFullRepaint(value bool) {
	imports.SysCallN(def.BufferPanel_SetFullRepaint, m.Instance(), api.PascalBool(value))
}

// Font
//
// 获取字体。
//
// Get Font.
func (m *TBufferPanel) Font() *lcl.TFont {
	return lcl.AsFont(imports.SysCallN(def.BufferPanel_GetFont, m.Instance()))
}

// SetFont
//
// 设置字体。
//
// Set Font.
func (m *TBufferPanel) SetFont(value *lcl.TFont) {
	imports.SysCallN(def.BufferPanel_SetFont, m.Instance(), lcl.CheckPtr(value))
}

func (m *TBufferPanel) ParentBackground() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetParentBackground, m.Instance()))
}

func (m *TBufferPanel) SetParentBackground(value bool) {
	imports.SysCallN(def.BufferPanel_SetParentBackground, m.Instance(), api.PascalBool(value))
}

// ParentColor
//
// 获取使用父容器颜色。
//
// Get parent color.
func (m *TBufferPanel) ParentColor() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetParentColor, m.Instance()))
}

// SetParentColor
//
// 设置使用父容器颜色。
//
// Set parent color.
func (m *TBufferPanel) SetParentColor(value bool) {
	imports.SysCallN(def.BufferPanel_SetParentColor, m.Instance(), api.PascalBool(value))
}

// ParentDoubleBuffered
//
// 获取使用父容器双缓冲。
//
// Get Parent container double buffering.
func (m *TBufferPanel) ParentDoubleBuffered() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetParentDoubleBuffered, m.Instance()))
}

// SetParentDoubleBuffered
//
// 设置使用父容器双缓冲。
//
// Set Parent container double buffering.
func (m *TBufferPanel) SetParentDoubleBuffered(value bool) {
	imports.SysCallN(def.BufferPanel_SetParentDoubleBuffered, m.Instance(), api.PascalBool(value))
}

// ParentFont
//
// 获取使用父容器字体。
//
// Get Parent container font.
func (m *TBufferPanel) ParentFont() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetParentFont, m.Instance()))
}

// SetParentFont
//
// 设置使用父容器字体。
//
// Set Parent container font.
func (m *TBufferPanel) SetParentFont(value bool) {
	imports.SysCallN(def.BufferPanel_SetParentFont, m.Instance(), api.PascalBool(value))
}

// ParentShowHint
//
// 获取以父容器的ShowHint属性为准。
func (m *TBufferPanel) ParentShowHint() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetParentShowHint, m.Instance()))
}

// SetParentShowHint
//
// 设置以父容器的ShowHint属性为准。
func (m *TBufferPanel) SetParentShowHint(value bool) {
	imports.SysCallN(def.BufferPanel_SetParentShowHint, m.Instance(), api.PascalBool(value))
}

// PopupMenu
//
// 获取右键菜单。
//
// Get Right click menu.
func (m *TBufferPanel) PopupMenu() *lcl.TPopupMenu {
	return lcl.AsPopupMenu(imports.SysCallN(def.BufferPanel_GetPopupMenu, m.Instance()))
}

// SetPopupMenu
//
// 设置右键菜单。
//
// Set Right click menu.
func (m *TBufferPanel) SetPopupMenu(value lcl.IComponent) {
	imports.SysCallN(def.BufferPanel_SetPopupMenu, m.Instance(), lcl.CheckPtr(value))
}

// ShowHint
//
// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (m *TBufferPanel) ShowHint() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetShowHint, m.Instance()))
}

// SetShowHint
//
// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (m *TBufferPanel) SetShowHint(value bool) {
	imports.SysCallN(def.BufferPanel_SetShowHint, m.Instance(), api.PascalBool(value))
}

// TabOrder
//
// 获取Tab切换顺序序号。
//
// Get Tab switching sequence number.
func (m *TBufferPanel) TabOrder() types.TTabOrder {
	return types.TTabOrder(imports.SysCallN(def.BufferPanel_GetTabOrder, m.Instance()))
}

// SetTabOrder
//
// 设置Tab切换顺序序号。
//
// Set Tab switching sequence number.
func (m *TBufferPanel) SetTabOrder(value types.TTabOrder) {
	imports.SysCallN(def.BufferPanel_SetTabOrder, m.Instance(), uintptr(value))
}

// TabStop
//
// 获取Tab可停留。
//
// Get Tab can stay.
func (m *TBufferPanel) TabStop() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetTabStop, m.Instance()))
}

// SetTabStop
//
// 设置Tab可停留。
//
// Set Tab can stay.
func (m *TBufferPanel) SetTabStop(value bool) {
	imports.SysCallN(def.BufferPanel_SetTabStop, m.Instance(), api.PascalBool(value))
}

// Visible
//
// 获取控件可视。
//
// Get the control visible.
func (m *TBufferPanel) Visible() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetVisible, m.Instance()))
}

// SetVisible
//
// 设置控件可视。
//
// Set the control visible.
func (m *TBufferPanel) SetVisible(value bool) {
	imports.SysCallN(def.BufferPanel_SetVisible, m.Instance(), api.PascalBool(value))
}

// SetOnAlignPosition
//
// 设置对齐位置事件，当Align为alCustom时Parent会收到这个消息。
func (m *TBufferPanel) SetOnAlignPosition(fn lcl.TAlignPositionEvent) {
	imports.SysCallN(def.BufferPanel_SetOnAlignPosition, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnClick
//
// 设置控件单击事件。
//
// Set control click event.
func (m *TBufferPanel) SetOnClick(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnClick, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnContextPopup
//
// 设置上下文弹出事件，一般是右键时弹出。
//
// Set Context popup event, usually pop up when right click.
func (m *TBufferPanel) SetOnContextPopup(fn lcl.TContextPopupEvent) {
	imports.SysCallN(def.BufferPanel_SetOnContextPopup, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnDockDrop(fn lcl.TDockDropEvent) {
	imports.SysCallN(def.BufferPanel_SetOnDockDrop, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnDblClick
//
// 设置双击事件。
func (m *TBufferPanel) SetOnDblClick(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnDblClick, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnDragDrop
//
// 设置拖拽下落事件。
//
// Set Drag and drop event.
func (m *TBufferPanel) SetOnDragDrop(fn lcl.TDragDropEvent) {
	imports.SysCallN(def.BufferPanel_SetOnDragDrop, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnDragOver
//
// 设置拖拽完成事件。
//
// Set Drag and drop completion event.
func (m *TBufferPanel) SetOnDragOver(fn lcl.TDragOverEvent) {
	imports.SysCallN(def.BufferPanel_SetOnDragOver, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnEndDock
//
// 设置停靠结束事件。
//
// Set Dock end event.
func (m *TBufferPanel) SetOnEndDock(fn lcl.TEndDragEvent) {
	imports.SysCallN(def.BufferPanel_SetOnEndDock, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnEndDrag
//
// 设置拖拽结束。
//
// Set End of drag.
func (m *TBufferPanel) SetOnEndDrag(fn lcl.TEndDragEvent) {
	imports.SysCallN(def.BufferPanel_SetOnEndDrag, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnEnter
//
// 设置焦点进入。
//
// Set Focus entry.
func (m *TBufferPanel) SetOnEnter(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnEnter, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnExit
//
// 设置焦点退出。
//
// Set Focus exit.
func (m *TBufferPanel) SetOnExit(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnExit, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnGetSiteInfo(fn lcl.TGetSiteInfoEvent) {
	imports.SysCallN(def.BufferPanel_SetOnGetSiteInfo, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnMouseDown
//
// 设置鼠标按下事件。
//
// Set Mouse down event.
func (m *TBufferPanel) SetOnMouseDown(fn lcl.TMouseEvent) {
	imports.SysCallN(def.BufferPanel_SetOnMouseDown, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnMouseEnter
//
// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (m *TBufferPanel) SetOnMouseEnter(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnMouseEnter, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnMouseLeave
//
// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (m *TBufferPanel) SetOnMouseLeave(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnMouseLeave, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnMouseMove
//
// 设置鼠标移动事件。
func (m *TBufferPanel) SetOnMouseMove(fn lcl.TMouseMoveEvent) {
	imports.SysCallN(def.BufferPanel_SetOnMouseMove, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnMouseUp
//
// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (m *TBufferPanel) SetOnMouseUp(fn lcl.TMouseEvent) {
	imports.SysCallN(def.BufferPanel_SetOnMouseUp, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnResize
//
// 设置大小被改变事件。
func (m *TBufferPanel) SetOnResize(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnResize, m.Instance(), api.MakeEventDataPtr(fn))
}

// SetOnStartDock
//
// 设置启动停靠。
func (m *TBufferPanel) SetOnStartDock(fn lcl.TStartDockEvent) {
	imports.SysCallN(def.BufferPanel_SetOnStartDock, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnUnDock(fn lcl.TUnDockEvent) {
	imports.SysCallN(def.BufferPanel_SetOnUnDock, m.Instance(), api.MakeEventDataPtr(fn))
}

// DockClientCount
//
// 获取依靠客户端总数。
func (m *TBufferPanel) DockClientCount() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetDockClientCount, m.Instance()))
}

// MouseInClient
//
// 获取鼠标是否在客户端，仅VCL有效。
//
// Get Whether the mouse is on the client, only VCL is valid.
func (m *TBufferPanel) MouseInClient() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetMouseInClient, m.Instance()))
}

// VisibleDockClientCount
//
// 获取当前停靠的可视总数。
//
// Get The total number of visible calls currently docked.
func (m *TBufferPanel) VisibleDockClientCount() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetVisibleDockClientCount, m.Instance()))
}

// Brush
//
// 获取画刷对象。
//
// Get Brush.
func (m *TBufferPanel) Brush() *lcl.TBrush {
	return lcl.AsBrush(imports.SysCallN(def.BufferPanel_GetBrush, m.Instance()))
}

// ControlCount
//
// 获取子控件数。
//
// Get Number of child controls.
func (m *TBufferPanel) ControlCount() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetControlCount, m.Instance()))
}

// Handle
//
// 获取控件句柄。
//
// Get Control handle.
func (m *TBufferPanel) Handle() types.HWND {
	return types.HWND(imports.SysCallN(def.BufferPanel_GetHandle, m.Instance()))
}

// ParentWindow
//
// 获取父容器句柄。
//
// Get Parent container handle.
func (m *TBufferPanel) ParentWindow() types.HWND {
	return imports.SysCallN(def.BufferPanel_GetParentWindow, m.Instance())
}

// SetParentWindow
//
// 设置父容器句柄。
//
// Set Parent container handle.
func (m *TBufferPanel) SetParentWindow(value types.HWND) {
	imports.SysCallN(def.BufferPanel_SetParentWindow, m.Instance(), value)
}

func (m *TBufferPanel) Showing() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetShowing, m.Instance()))
}

func (m *TBufferPanel) Action() *lcl.TAction {
	return lcl.AsAction(imports.SysCallN(def.BufferPanel_GetAction, m.Instance()))
}

func (m *TBufferPanel) SetAction(value lcl.IComponent) {
	imports.SysCallN(def.BufferPanel_SetAction, m.Instance(), lcl.CheckPtr(value))
}

func (m *TBufferPanel) BoundsRect() (result types.TRect) {
	imports.SysCallN(def.BufferPanel_GetBoundsRect, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TBufferPanel) SetBoundsRect(value types.TRect) {
	imports.SysCallN(def.BufferPanel_SetBoundsRect, m.Instance(), uintptr(unsafe.Pointer(&value)))
}

// ClientHeight
//
// 获取客户区高度。
//
// Get client height.
func (m *TBufferPanel) ClientHeight() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetClientHeight, m.Instance()))
}

// SetClientHeight
//
// 设置客户区高度。
//
// Set client height.
func (m *TBufferPanel) SetClientHeight(value int32) {
	imports.SysCallN(def.BufferPanel_SetClientHeight, m.Instance(), uintptr(value))
}

func (m *TBufferPanel) ClientOrigin() (result types.TPoint) {
	imports.SysCallN(def.BufferPanel_GetClientOrigin, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// ClientRect
//
// 获取客户区矩形。
//
// Get client rectangle.
func (m *TBufferPanel) ClientRect() (result types.TRect) {
	imports.SysCallN(def.BufferPanel_GetClientRect, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

// ClientWidth
//
// 获取客户区宽度。
//
// Get client width.
func (m *TBufferPanel) ClientWidth() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetClientWidth, m.Instance()))
}

// SetClientWidth
//
// 设置客户区宽度。
//
// Set client width.
func (m *TBufferPanel) SetClientWidth(value int32) {
	imports.SysCallN(def.BufferPanel_SetClientWidth, m.Instance(), uintptr(value))
}

// ControlState
//
// 获取控件状态。
//
// Get control state.
func (m *TBufferPanel) ControlState() types.TControlState {
	return types.TControlState(imports.SysCallN(def.BufferPanel_GetControlState, m.Instance()))
}

// SetControlState
//
// 设置控件状态。
//
// Set control state.
func (m *TBufferPanel) SetControlState(value types.TControlState) {
	imports.SysCallN(def.BufferPanel_SetControlState, m.Instance(), uintptr(value))
}

// ControlStyle
//
// 获取控件样式。
//
// Get control style.
func (m *TBufferPanel) ControlStyle() types.TControlStyle {
	return types.TControlStyle(imports.SysCallN(def.BufferPanel_GetControlStyle, m.Instance()))
}

// SetControlStyle
//
// 设置控件样式。
//
// Set control style.
func (m *TBufferPanel) SetControlStyle(value types.TControlStyle) {
	imports.SysCallN(def.BufferPanel_SetControlStyle, m.Instance(), uintptr(value))
}

func (m *TBufferPanel) Floating() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetFloating, m.Instance()))
}

// Parent
//
// 获取控件父容器。
//
// Get control parent container.
func (m *TBufferPanel) Parent() *lcl.TWinControl {
	return lcl.AsWinControl(imports.SysCallN(def.BufferPanel_GetParent, m.Instance()))
}

// SetParent
//
// 设置控件父容器。
//
// Set control parent container.
func (m *TBufferPanel) SetParent(value lcl.IWinControl) {
	imports.SysCallN(def.BufferPanel_SetParent, m.Instance(), lcl.CheckPtr(value))
}

// Left
//
// 获取左边位置。
//
// Get Left position.
func (m *TBufferPanel) Left() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetLeft, m.Instance()))
}

// SetLeft
//
// 设置左边位置。
//
// Set Left position.
func (m *TBufferPanel) SetLeft(value int32) {
	imports.SysCallN(def.BufferPanel_SetLeft, m.Instance(), uintptr(value))
}

// Top
//
// 获取顶边位置。
//
// Get Top position.
func (m *TBufferPanel) Top() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetTop, m.Instance()))
}

// SetTop
//
// 设置顶边位置。
//
// Set Top position.
func (m *TBufferPanel) SetTop(value int32) {
	imports.SysCallN(def.BufferPanel_SetTop, m.Instance(), uintptr(value))
}

// Width
//
// 获取宽度。
//
// Get width.
func (m *TBufferPanel) Width() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetWidth, m.Instance()))
}

// SetWidth
//
// 设置宽度。
//
// Set width.
func (m *TBufferPanel) SetWidth(value int32) {
	imports.SysCallN(def.BufferPanel_SetWidth, m.Instance(), uintptr(value))
}

// Height
//
// 获取高度。
//
// Get height.
func (m *TBufferPanel) Height() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetHeight, m.Instance()))
}

// SetHeight
//
// 设置高度。
//
// Set height.
func (m *TBufferPanel) SetHeight(value int32) {
	imports.SysCallN(def.BufferPanel_SetHeight, m.Instance(), uintptr(value))
}

// Cursor
//
// 获取控件光标。
//
// Get control cursor.
func (m *TBufferPanel) Cursor() types.TCursor {
	return types.TCursor(imports.SysCallN(def.BufferPanel_GetCursor, m.Instance()))
}

// SetCursor
//
// 设置控件光标。
//
// Set control cursor.
func (m *TBufferPanel) SetCursor(value types.TCursor) {
	imports.SysCallN(def.BufferPanel_SetCursor, m.Instance(), uintptr(value))
}

// Hint
//
// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (m *TBufferPanel) Hint() string {
	return api.GoStr(imports.SysCallN(def.BufferPanel_GetHint, m.Instance()))
}

// SetHint
//
// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (m *TBufferPanel) SetHint(value string) {
	imports.SysCallN(def.BufferPanel_SetHint, m.Instance(), api.PascalStr(value))
}

// ComponentCount
//
// 获取组件总数。
//
// Get the total number of components.
func (m *TBufferPanel) ComponentCount() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetComponentCount, m.Instance()))
}

// ComponentIndex
//
// 获取组件索引。
//
// Get component index.
func (m *TBufferPanel) ComponentIndex() int32 {
	return int32(imports.SysCallN(def.BufferPanel_GetComponentIndex, m.Instance()))
}

// SetComponentIndex
//
// 设置组件索引。
//
// Set component index.
func (m *TBufferPanel) SetComponentIndex(value int32) {
	imports.SysCallN(def.BufferPanel_SetComponentIndex, m.Instance(), uintptr(value))
}

// Owner
//
// 获取组件所有者。
//
// Get component owner.
func (m *TBufferPanel) Owner() *lcl.TComponent {
	return lcl.AsComponent(imports.SysCallN(def.BufferPanel_GetOwner, m.Instance()))
}

// Name
//
// 获取组件名称。
//
// Get the component name.
func (m *TBufferPanel) Name() string {
	return api.GoStr(imports.SysCallN(def.BufferPanel_GetName, m.Instance()))
}

// SetName
//
// 设置组件名称。
//
// Set the component name.
func (m *TBufferPanel) SetName(value string) {
	imports.SysCallN(def.BufferPanel_SetName, m.Instance(), api.PascalStr(value))
}

// Tag
//
// 获取对象标记。
//
// Get the control tag.
func (m *TBufferPanel) Tag() int {
	return int(imports.SysCallN(def.BufferPanel_GetTag, m.Instance()))
}

// SetTag
//
// 设置对象标记。
//
// Set the control tag.
func (m *TBufferPanel) SetTag(value int) {
	imports.SysCallN(def.BufferPanel_SetTag, m.Instance(), uintptr(value))
}

// AnchorSideLeft
//
// 获取左边锚点。
func (m *TBufferPanel) AnchorSideLeft() *lcl.TAnchorSide {
	return lcl.AsAnchorSide(imports.SysCallN(def.BufferPanel_GetAnchorSideLeft, m.Instance()))
}

// SetAnchorSideLeft
//
// 设置左边锚点。
func (m *TBufferPanel) SetAnchorSideLeft(value *lcl.TAnchorSide) {
	imports.SysCallN(def.BufferPanel_SetAnchorSideLeft, m.Instance(), lcl.CheckPtr(value))
}

// AnchorSideTop
//
// 获取顶边锚点。
func (m *TBufferPanel) AnchorSideTop() *lcl.TAnchorSide {
	return lcl.AsAnchorSide(imports.SysCallN(def.BufferPanel_GetAnchorSideTop, m.Instance()))
}

// SetAnchorSideTop
//
// 设置顶边锚点。
func (m *TBufferPanel) SetAnchorSideTop(value *lcl.TAnchorSide) {
	imports.SysCallN(def.BufferPanel_SetAnchorSideTop, m.Instance(), lcl.CheckPtr(value))
}

// AnchorSideRight
//
// 获取右边锚点。
func (m *TBufferPanel) AnchorSideRight() *lcl.TAnchorSide {
	return lcl.AsAnchorSide(imports.SysCallN(def.BufferPanel_GetAnchorSideRight, m.Instance()))
}

// SetAnchorSideRight
//
// 设置右边锚点。
func (m *TBufferPanel) SetAnchorSideRight(value *lcl.TAnchorSide) {
	imports.SysCallN(def.BufferPanel_SetAnchorSideRight, m.Instance(), lcl.CheckPtr(value))
}

// AnchorSideBottom
//
// 获取底边锚点。
func (m *TBufferPanel) AnchorSideBottom() *lcl.TAnchorSide {
	return lcl.AsAnchorSide(imports.SysCallN(def.BufferPanel_GetAnchorSideBottom, m.Instance()))
}

// SetAnchorSideBottom
//
// 设置底边锚点。
func (m *TBufferPanel) SetAnchorSideBottom(value *lcl.TAnchorSide) {
	imports.SysCallN(def.BufferPanel_SetAnchorSideBottom, m.Instance(), lcl.CheckPtr(value))
}

func (m *TBufferPanel) ChildSizing() *lcl.TControlChildSizing {
	return lcl.AsControlChildSizing(imports.SysCallN(def.BufferPanel_GetChildSizing, m.Instance()))
}

func (m *TBufferPanel) SetChildSizing(value *lcl.TControlChildSizing) {
	imports.SysCallN(def.BufferPanel_SetChildSizing, m.Instance(), lcl.CheckPtr(value))
}

// BorderSpacing
//
// 获取边框间距。
func (m *TBufferPanel) BorderSpacing() *lcl.TControlBorderSpacing {
	return lcl.AsControlBorderSpacing(imports.SysCallN(def.BufferPanel_GetBorderSpacing, m.Instance()))
}

// SetBorderSpacing
//
// 设置边框间距。
func (m *TBufferPanel) SetBorderSpacing(value *lcl.TControlBorderSpacing) {
	imports.SysCallN(def.BufferPanel_SetBorderSpacing, m.Instance(), lcl.CheckPtr(value))
}

// DockClients
//
// 获取指定索引停靠客户端。
func (m *TBufferPanel) DockClients(Index int32) *lcl.TControl {
	return lcl.AsControl(imports.SysCallN(def.BufferPanel_GetDockClients, m.Instance(), uintptr(Index)))
}

// Controls
//
// 获取指定索引子控件。
func (m *TBufferPanel) Controls(Index int32) *lcl.TControl {
	return lcl.AsControl(imports.SysCallN(def.BufferPanel_GetControls, m.Instance(), uintptr(Index)))
}

// Components
//
// 获取指定索引组件。
//
// Get the specified index component.
func (m *TBufferPanel) Components(Index int32) *lcl.TComponent {
	return lcl.AsComponent(imports.SysCallN(def.BufferPanel_GetComponents, m.Instance(), uintptr(Index)))
}

// AnchorSide
//
// 获取锚侧面。
func (m *TBufferPanel) AnchorSide(AKind types.TAnchorKind) *lcl.TAnchorSide {
	return lcl.AsAnchorSide(imports.SysCallN(def.BufferPanel_GetAnchorSide, m.Instance(), uintptr(AKind)))
}

func (m *TBufferPanel) GetTransparent() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetTransparent, m.Instance()))
}

func (m *TBufferPanel) SetTransparent(value bool) {
	imports.SysCallN(def.BufferPanel_SetTransparent, m.Instance(), api.PascalBool(value))
}

func (m *TBufferPanel) SaveToFile(fileName string) bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_SaveToFile, m.Instance(), api.PascalStr(fileName)))
}

func (m *TBufferPanel) InvalidatePanel() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_InvalidatePanel, m.Instance()))
}

func (m *TBufferPanel) BeginBufferDraw() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_BeginBufferDraw, m.Instance()))
}

func (m *TBufferPanel) EndBufferDraw() {
	imports.SysCallN(def.BufferPanel_EndBufferDraw, m.Instance())
}

func (m *TBufferPanel) BufferDrawPoint(x, y int32, bitmap lcl.TBitmap) {
	imports.SysCallN(def.BufferPanel_BufferDrawPoint, m.Instance(), uintptr(x), uintptr(y), bitmap.Instance())
}

func (m *TBufferPanel) BufferDrawRect(bitmap lcl.TBitmap, srcRect, dstRect types.TRect) {
	imports.SysCallN(def.BufferPanel_BufferDrawRect, m.Instance(), bitmap.Instance(), uintptr(unsafe.Pointer(&srcRect)), uintptr(unsafe.Pointer(&dstRect)))
}

func (m *TBufferPanel) UpdateBufferDimensions(width, height int32) bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_UpdateBufferDimensions, m.Instance(), uintptr(width), uintptr(height)))
}

func (m *TBufferPanel) UpdateOrigBufferDimensions(width, height int32) bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_UpdateOrigBufferDimensions, m.Instance(), uintptr(width), uintptr(height)))
}

func (m *TBufferPanel) UpdateOrigPopupBufferDimensions(width, height int32) bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_UpdateOrigPopupBufferDimensions, m.Instance(), uintptr(width), uintptr(height)))
}

func (m *TBufferPanel) UpdateDeviceScaleFactor() {
	imports.SysCallN(def.BufferPanel_UpdateDeviceScaleFactor, m.Instance())
}

func (m *TBufferPanel) BufferIsResized(useMutex bool) bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_BufferIsResized, m.Instance(), api.PascalBool(useMutex)))
}

func (m *TBufferPanel) CreateIMEHandler() {
	imports.SysCallN(def.BufferPanel_CreateIMEHandler, m.Instance())
}

func (m *TBufferPanel) ChangeCompositionRange(selectionRange TCefRange, characterBounds []TCefRect) {
	imports.SysCallN(def.BufferPanel_ChangeCompositionRange, m.Instance(), uintptr(unsafe.Pointer(&selectionRange)), uintptr(unsafe.Pointer(&characterBounds[0])), uintptr(int32(len(characterBounds)-1)))
}

func (m *TBufferPanel) DrawOrigPopupBuffer(srcRect, dstRect types.TRect) {
	imports.SysCallN(def.BufferPanel_DrawOrigPopupBuffer, m.Instance(), uintptr(unsafe.Pointer(&srcRect)), uintptr(unsafe.Pointer(&dstRect)))
}

func (m *TBufferPanel) ScanlineSize() int32 {
	return int32(imports.SysCallN(def.BufferPanel_ScanlineSize, m.Instance()))
}

func (m *TBufferPanel) BufferWidth() int32 {
	return int32(imports.SysCallN(def.BufferPanel_BufferWidth, m.Instance()))
}

func (m *TBufferPanel) BufferHeight() int32 {
	return int32(imports.SysCallN(def.BufferPanel_BufferHeight, m.Instance()))
}

func (m *TBufferPanel) BufferBits() unsafe.Pointer {
	return unsafe.Pointer(imports.SysCallN(def.BufferPanel_BufferBits, m.Instance()))
}

func (m *TBufferPanel) ScreenScale() (result float32) {
	imports.SysCallN(def.BufferPanel_ScreenScale, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TBufferPanel) GetForcedDeviceScaleFactor() (result float32) {
	imports.SysCallN(def.BufferPanel_GetForcedDeviceScaleFactor, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return
}

func (m *TBufferPanel) SetForcedDeviceScaleFactor(value float32) {
	imports.SysCallN(def.BufferPanel_SetForcedDeviceScaleFactor, m.Instance(), uintptr(unsafe.Pointer(&value)))
}

func (m *TBufferPanel) GetMustInitBuffer() bool {
	return api.GoBool(imports.SysCallN(def.BufferPanel_GetMustInitBuffer, m.Instance()))
}

func (m *TBufferPanel) SetMustInitBuffer(value bool) {
	imports.SysCallN(def.BufferPanel_SetMustInitBuffer, m.Instance(), api.PascalBool(value))
}

func (m *TBufferPanel) Buffer() *lcl.TBitmap {
	var result uintptr
	imports.SysCallN(def.BufferPanel_Buffer, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return lcl.AsBitmap(result)
}

func (m *TBufferPanel) OrigBuffer() *TCEFBitmapBitBuffer {
	var result uintptr
	imports.SysCallN(def.BufferPanel_OrigBuffer, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &TCEFBitmapBitBuffer{instance: unsafe.Pointer(result)}
}

func (m *TBufferPanel) OrigBufferWidth() int32 {
	return int32(imports.SysCallN(def.BufferPanel_OrigBufferWidth, m.Instance()))
}

func (m *TBufferPanel) OrigBufferHeight() int32 {
	return int32(imports.SysCallN(def.BufferPanel_OrigBufferHeight, m.Instance()))
}

func (m *TBufferPanel) OrigPopupBuffer() *TCEFBitmapBitBuffer {
	var result uintptr
	imports.SysCallN(def.BufferPanel_OrigPopupBuffer, m.Instance(), uintptr(unsafe.Pointer(&result)))
	return &TCEFBitmapBitBuffer{instance: unsafe.Pointer(result)}
}

func (m *TBufferPanel) OrigPopupBufferWidth() int32 {
	return int32(imports.SysCallN(def.BufferPanel_OrigPopupBufferWidth, m.Instance()))
}

func (m *TBufferPanel) OrigPopupBufferHeight() int32 {
	return int32(imports.SysCallN(def.BufferPanel_OrigPopupBufferHeight, m.Instance()))
}

func (m *TBufferPanel) OrigPopupBufferBits() unsafe.Pointer {
	return unsafe.Pointer(imports.SysCallN(def.BufferPanel_OrigPopupBufferBits, m.Instance()))
}

func (m *TBufferPanel) OrigPopupScanlineSize() int32 {
	return int32(imports.SysCallN(def.BufferPanel_OrigPopupScanlineSize, m.Instance()))
}

func (m *TBufferPanel) SetOnPaintParentBkg(fn lcl.TNotifyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnPaintParentBkg, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnMouseWheel(fn lcl.TMouseWheelEvent) {
	imports.SysCallN(def.BufferPanel_SetOnMouseWheel, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnKeyDown(fn lcl.TKeyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnKeyDown, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnKeyUp(fn lcl.TKeyEvent) {
	imports.SysCallN(def.BufferPanel_SetOnKeyUp, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnKeyPress(fn lcl.TKeyPressEvent) {
	imports.SysCallN(def.BufferPanel_SetOnKeyPress, m.Instance(), api.MakeEventDataPtr(fn))
}

func (m *TBufferPanel) SetOnUTF8KeyPress(fn lcl.TUTF8KeyPressEvent) {
	imports.SysCallN(def.BufferPanel_SetOnUTF8KeyPress, m.Instance(), api.MakeEventDataPtr(fn))
}
