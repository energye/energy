//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
	"unsafe"
)

// IMenuItem Parent: ILCLComponent
type IMenuItem interface {
	ILCLComponent
	Merged() IMenuItem                                                // property
	MergedWith() IMenuItem                                            // property
	Count() int32                                                     // property
	Handle() HMENU                                                    // property
	SetHandle(AValue HMENU)                                           // property
	Items(Index int32) IMenuItem                                      // property
	MergedItems() IMergedMenuItems                                    // property
	MenuIndex() int32                                                 // property
	SetMenuIndex(AValue int32)                                        // property
	Menu() IMenu                                                      // property
	Parent() IMenuItem                                                // property
	MergedParent() IMenuItem                                          // property
	Command() Word                                                    // property
	Action() IBasicAction                                             // property
	SetAction(AValue IBasicAction)                                    // property
	AutoCheck() bool                                                  // property
	SetAutoCheck(AValue bool)                                         // property
	Caption() string                                                  // property
	SetCaption(AValue string)                                         // property
	Checked() bool                                                    // property
	SetChecked(AValue bool)                                           // property
	Default() bool                                                    // property
	SetDefault(AValue bool)                                           // property
	Enabled() bool                                                    // property
	SetEnabled(AValue bool)                                           // property
	Bitmap() IBitmap                                                  // property
	SetBitmap(AValue IBitmap)                                         // property
	GroupIndex() Byte                                                 // property
	SetGroupIndex(AValue Byte)                                        // property
	GlyphShowMode() TGlyphShowMode                                    // property
	SetGlyphShowMode(AValue TGlyphShowMode)                           // property
	HelpContext() THelpContext                                        // property
	SetHelpContext(AValue THelpContext)                               // property
	Hint() string                                                     // property
	SetHint(AValue string)                                            // property
	ImageIndex() TImageIndex                                          // property
	SetImageIndex(AValue TImageIndex)                                 // property
	RadioItem() bool                                                  // property
	SetRadioItem(AValue bool)                                         // property
	RightJustify() bool                                               // property
	SetRightJustify(AValue bool)                                      // property
	ShortCut() TShortCut                                              // property
	SetShortCut(AValue TShortCut)                                     // property
	ShortCutKey2() TShortCut                                          // property
	SetShortCutKey2(AValue TShortCut)                                 // property
	ShowAlwaysCheckable() bool                                        // property
	SetShowAlwaysCheckable(AValue bool)                               // property
	SubMenuImages() ICustomImageList                                  // property
	SetSubMenuImages(AValue ICustomImageList)                         // property
	SubMenuImagesWidth() int32                                        // property
	SetSubMenuImagesWidth(AValue int32)                               // property
	Visible() bool                                                    // property
	SetVisible(AValue bool)                                           // property
	Find(ACaption string) IMenuItem                                   // function
	GetEnumeratorForMenuItemEnumerator() IMenuItemEnumerator          // function
	GetImageList() ICustomImageList                                   // function
	GetParentMenu() IMenu                                             // function
	GetMergedParentMenu() IMenu                                       // function
	GetIsRightToLeft() bool                                           // function
	HandleAllocated() bool                                            // function
	HasIcon() bool                                                    // function
	IndexOf(Item IMenuItem) int32                                     // function
	IndexOfCaption(ACaption string) int32                             // function
	VisibleIndexOf(Item IMenuItem) int32                              // function
	IsCheckItem() bool                                                // function
	IsLine() bool                                                     // function
	IsInMenuBar() bool                                                // function
	HasBitmap() bool                                                  // function
	GetIconSize(ADC HDC, DPI int32) (resultPoint TPoint)              // function
	MenuVisibleIndex() int32                                          // function
	GetImageList1(OutImages *ICustomImageList, OutImagesWidth *int32) // procedure
	InitiateAction()                                                  // procedure
	IntfDoSelect()                                                    // procedure
	InvalidateMergedItems()                                           // procedure
	Add(Item IMenuItem)                                               // procedure
	AddSeparator()                                                    // procedure
	Click()                                                           // procedure
	Delete(Index int32)                                               // procedure
	HandleNeeded()                                                    // procedure
	Insert(Index int32, Item IMenuItem)                               // procedure
	RecreateHandle()                                                  // procedure
	Remove(Item IMenuItem)                                            // procedure
	UpdateImage(forced bool)                                          // procedure
	UpdateImages(forced bool)                                         // procedure
	Clear()                                                           // procedure
	WriteDebugReport(Prefix string)                                   // procedure
	SetOnClick(fn TNotifyEvent)                                       // property event
	SetOnDrawItem(fn TMenuDrawItemEvent)                              // property event
	SetOnMeasureItem(fn TMenuMeasureItemEvent)                        // property event
}

// TMenuItem Parent: TLCLComponent
type TMenuItem struct {
	TLCLComponent
	clickPtr       uintptr
	drawItemPtr    uintptr
	measureItemPtr uintptr
}

func NewMenuItem(TheOwner IComponent) IMenuItem {
	r1 := LCL().SysCallN(3610, GetObjectUintptr(TheOwner))
	return AsMenuItem(r1)
}

func (m *TMenuItem) Merged() IMenuItem {
	r1 := LCL().SysCallN(3645, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItem) MergedWith() IMenuItem {
	r1 := LCL().SysCallN(3648, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItem) Count() int32 {
	r1 := LCL().SysCallN(3609, m.Instance())
	return int32(r1)
}

func (m *TMenuItem) Handle() HMENU {
	r1 := LCL().SysCallN(3624, 0, m.Instance(), 0)
	return HMENU(r1)
}

func (m *TMenuItem) SetHandle(AValue HMENU) {
	LCL().SysCallN(3624, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) Items(Index int32) IMenuItem {
	r1 := LCL().SysCallN(3641, m.Instance(), uintptr(Index))
	return AsMenuItem(r1)
}

func (m *TMenuItem) MergedItems() IMergedMenuItems {
	r1 := LCL().SysCallN(3646, m.Instance())
	return AsMergedMenuItems(r1)
}

func (m *TMenuItem) MenuIndex() int32 {
	r1 := LCL().SysCallN(3643, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMenuItem) SetMenuIndex(AValue int32) {
	LCL().SysCallN(3643, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) Menu() IMenu {
	r1 := LCL().SysCallN(3642, m.Instance())
	return AsMenu(r1)
}

func (m *TMenuItem) Parent() IMenuItem {
	r1 := LCL().SysCallN(3649, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItem) MergedParent() IMenuItem {
	r1 := LCL().SysCallN(3647, m.Instance())
	return AsMenuItem(r1)
}

func (m *TMenuItem) Command() Word {
	r1 := LCL().SysCallN(3608, m.Instance())
	return Word(r1)
}

func (m *TMenuItem) Action() IBasicAction {
	r1 := LCL().SysCallN(3598, 0, m.Instance(), 0)
	return AsBasicAction(r1)
}

func (m *TMenuItem) SetAction(AValue IBasicAction) {
	LCL().SysCallN(3598, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenuItem) AutoCheck() bool {
	r1 := LCL().SysCallN(3601, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetAutoCheck(AValue bool) {
	LCL().SysCallN(3601, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Caption() string {
	r1 := LCL().SysCallN(3603, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TMenuItem) SetCaption(AValue string) {
	LCL().SysCallN(3603, 1, m.Instance(), PascalStr(AValue))
}

func (m *TMenuItem) Checked() bool {
	r1 := LCL().SysCallN(3604, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetChecked(AValue bool) {
	LCL().SysCallN(3604, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Default() bool {
	r1 := LCL().SysCallN(3611, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetDefault(AValue bool) {
	LCL().SysCallN(3611, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Enabled() bool {
	r1 := LCL().SysCallN(3613, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetEnabled(AValue bool) {
	LCL().SysCallN(3613, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Bitmap() IBitmap {
	r1 := LCL().SysCallN(3602, 0, m.Instance(), 0)
	return AsBitmap(r1)
}

func (m *TMenuItem) SetBitmap(AValue IBitmap) {
	LCL().SysCallN(3602, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenuItem) GroupIndex() Byte {
	r1 := LCL().SysCallN(3623, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TMenuItem) SetGroupIndex(AValue Byte) {
	LCL().SysCallN(3623, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) GlyphShowMode() TGlyphShowMode {
	r1 := LCL().SysCallN(3622, 0, m.Instance(), 0)
	return TGlyphShowMode(r1)
}

func (m *TMenuItem) SetGlyphShowMode(AValue TGlyphShowMode) {
	LCL().SysCallN(3622, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) HelpContext() THelpContext {
	r1 := LCL().SysCallN(3629, 0, m.Instance(), 0)
	return THelpContext(r1)
}

func (m *TMenuItem) SetHelpContext(AValue THelpContext) {
	LCL().SysCallN(3629, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) Hint() string {
	r1 := LCL().SysCallN(3630, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TMenuItem) SetHint(AValue string) {
	LCL().SysCallN(3630, 1, m.Instance(), PascalStr(AValue))
}

func (m *TMenuItem) ImageIndex() TImageIndex {
	r1 := LCL().SysCallN(3631, 0, m.Instance(), 0)
	return TImageIndex(r1)
}

func (m *TMenuItem) SetImageIndex(AValue TImageIndex) {
	LCL().SysCallN(3631, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) RadioItem() bool {
	r1 := LCL().SysCallN(3650, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetRadioItem(AValue bool) {
	LCL().SysCallN(3650, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) RightJustify() bool {
	r1 := LCL().SysCallN(3653, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetRightJustify(AValue bool) {
	LCL().SysCallN(3653, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) ShortCut() TShortCut {
	r1 := LCL().SysCallN(3657, 0, m.Instance(), 0)
	return TShortCut(r1)
}

func (m *TMenuItem) SetShortCut(AValue TShortCut) {
	LCL().SysCallN(3657, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) ShortCutKey2() TShortCut {
	r1 := LCL().SysCallN(3658, 0, m.Instance(), 0)
	return TShortCut(r1)
}

func (m *TMenuItem) SetShortCutKey2(AValue TShortCut) {
	LCL().SysCallN(3658, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) ShowAlwaysCheckable() bool {
	r1 := LCL().SysCallN(3659, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetShowAlwaysCheckable(AValue bool) {
	LCL().SysCallN(3659, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) SubMenuImages() ICustomImageList {
	r1 := LCL().SysCallN(3660, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TMenuItem) SetSubMenuImages(AValue ICustomImageList) {
	LCL().SysCallN(3660, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TMenuItem) SubMenuImagesWidth() int32 {
	r1 := LCL().SysCallN(3661, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TMenuItem) SetSubMenuImagesWidth(AValue int32) {
	LCL().SysCallN(3661, 1, m.Instance(), uintptr(AValue))
}

func (m *TMenuItem) Visible() bool {
	r1 := LCL().SysCallN(3664, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TMenuItem) SetVisible(AValue bool) {
	LCL().SysCallN(3664, 1, m.Instance(), PascalBool(AValue))
}

func (m *TMenuItem) Find(ACaption string) IMenuItem {
	r1 := LCL().SysCallN(3614, m.Instance(), PascalStr(ACaption))
	return AsMenuItem(r1)
}

func (m *TMenuItem) GetEnumeratorForMenuItemEnumerator() IMenuItemEnumerator {
	r1 := LCL().SysCallN(3615, m.Instance())
	return AsMenuItemEnumerator(r1)
}

func (m *TMenuItem) GetImageList() ICustomImageList {
	r1 := LCL().SysCallN(3617, m.Instance())
	return AsCustomImageList(r1)
}

func (m *TMenuItem) GetParentMenu() IMenu {
	r1 := LCL().SysCallN(3621, m.Instance())
	return AsMenu(r1)
}

func (m *TMenuItem) GetMergedParentMenu() IMenu {
	r1 := LCL().SysCallN(3620, m.Instance())
	return AsMenu(r1)
}

func (m *TMenuItem) GetIsRightToLeft() bool {
	r1 := LCL().SysCallN(3619, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) HandleAllocated() bool {
	r1 := LCL().SysCallN(3625, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) HasIcon() bool {
	r1 := LCL().SysCallN(3628, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) IndexOf(Item IMenuItem) int32 {
	r1 := LCL().SysCallN(3632, m.Instance(), GetObjectUintptr(Item))
	return int32(r1)
}

func (m *TMenuItem) IndexOfCaption(ACaption string) int32 {
	r1 := LCL().SysCallN(3633, m.Instance(), PascalStr(ACaption))
	return int32(r1)
}

func (m *TMenuItem) VisibleIndexOf(Item IMenuItem) int32 {
	r1 := LCL().SysCallN(3665, m.Instance(), GetObjectUintptr(Item))
	return int32(r1)
}

func (m *TMenuItem) IsCheckItem() bool {
	r1 := LCL().SysCallN(3638, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) IsLine() bool {
	r1 := LCL().SysCallN(3640, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) IsInMenuBar() bool {
	r1 := LCL().SysCallN(3639, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) HasBitmap() bool {
	r1 := LCL().SysCallN(3627, m.Instance())
	return GoBool(r1)
}

func (m *TMenuItem) GetIconSize(ADC HDC, DPI int32) (resultPoint TPoint) {
	LCL().SysCallN(3616, m.Instance(), uintptr(ADC), uintptr(DPI), uintptr(unsafe.Pointer(&resultPoint)))
	return
}

func (m *TMenuItem) MenuVisibleIndex() int32 {
	r1 := LCL().SysCallN(3644, m.Instance())
	return int32(r1)
}

func MenuItemClass() TClass {
	ret := LCL().SysCallN(3605)
	return TClass(ret)
}

func (m *TMenuItem) GetImageList1(OutImages *ICustomImageList, OutImagesWidth *int32) {
	var result0 uintptr
	var result1 uintptr
	LCL().SysCallN(3618, m.Instance(), uintptr(unsafe.Pointer(&result0)), uintptr(unsafe.Pointer(&result1)))
	*OutImages = AsCustomImageList(result0)
	*OutImagesWidth = int32(result1)
}

func (m *TMenuItem) InitiateAction() {
	LCL().SysCallN(3634, m.Instance())
}

func (m *TMenuItem) IntfDoSelect() {
	LCL().SysCallN(3636, m.Instance())
}

func (m *TMenuItem) InvalidateMergedItems() {
	LCL().SysCallN(3637, m.Instance())
}

func (m *TMenuItem) Add(Item IMenuItem) {
	LCL().SysCallN(3599, m.Instance(), GetObjectUintptr(Item))
}

func (m *TMenuItem) AddSeparator() {
	LCL().SysCallN(3600, m.Instance())
}

func (m *TMenuItem) Click() {
	LCL().SysCallN(3607, m.Instance())
}

func (m *TMenuItem) Delete(Index int32) {
	LCL().SysCallN(3612, m.Instance(), uintptr(Index))
}

func (m *TMenuItem) HandleNeeded() {
	LCL().SysCallN(3626, m.Instance())
}

func (m *TMenuItem) Insert(Index int32, Item IMenuItem) {
	LCL().SysCallN(3635, m.Instance(), uintptr(Index), GetObjectUintptr(Item))
}

func (m *TMenuItem) RecreateHandle() {
	LCL().SysCallN(3651, m.Instance())
}

func (m *TMenuItem) Remove(Item IMenuItem) {
	LCL().SysCallN(3652, m.Instance(), GetObjectUintptr(Item))
}

func (m *TMenuItem) UpdateImage(forced bool) {
	LCL().SysCallN(3662, m.Instance(), PascalBool(forced))
}

func (m *TMenuItem) UpdateImages(forced bool) {
	LCL().SysCallN(3663, m.Instance(), PascalBool(forced))
}

func (m *TMenuItem) Clear() {
	LCL().SysCallN(3606, m.Instance())
}

func (m *TMenuItem) WriteDebugReport(Prefix string) {
	LCL().SysCallN(3666, m.Instance(), PascalStr(Prefix))
}

func (m *TMenuItem) SetOnClick(fn TNotifyEvent) {
	if m.clickPtr != 0 {
		RemoveEventElement(m.clickPtr)
	}
	m.clickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3654, m.Instance(), m.clickPtr)
}

func (m *TMenuItem) SetOnDrawItem(fn TMenuDrawItemEvent) {
	if m.drawItemPtr != 0 {
		RemoveEventElement(m.drawItemPtr)
	}
	m.drawItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3655, m.Instance(), m.drawItemPtr)
}

func (m *TMenuItem) SetOnMeasureItem(fn TMenuMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3656, m.Instance(), m.measureItemPtr)
}
