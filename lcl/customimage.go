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

// ICustomImage Parent: IGraphicControl
type ICustomImage interface {
	IGraphicControl
	HasGraphic() bool                                   // property
	AntialiasingMode() TAntialiasingMode                // property
	SetAntialiasingMode(AValue TAntialiasingMode)       // property
	Center() bool                                       // property
	SetCenter(AValue bool)                              // property
	KeepOriginXWhenClipped() bool                       // property
	SetKeepOriginXWhenClipped(AValue bool)              // property
	KeepOriginYWhenClipped() bool                       // property
	SetKeepOriginYWhenClipped(AValue bool)              // property
	ImageIndex() int32                                  // property
	SetImageIndex(AValue int32)                         // property
	ImageWidth() int32                                  // property
	SetImageWidth(AValue int32)                         // property
	Images() ICustomImageList                           // property
	SetImages(AValue ICustomImageList)                  // property
	Picture() IPicture                                  // property
	SetPicture(AValue IPicture)                         // property
	Stretch() bool                                      // property
	SetStretch(AValue bool)                             // property
	StretchOutEnabled() bool                            // property
	SetStretchOutEnabled(AValue bool)                   // property
	StretchInEnabled() bool                             // property
	SetStretchInEnabled(AValue bool)                    // property
	Transparent() bool                                  // property
	SetTransparent(AValue bool)                         // property
	Proportional() bool                                 // property
	SetProportional(AValue bool)                        // property
	DestRect() (resultRect TRect)                       // function
	SetOnMouseDown(fn TMouseEvent)                      // property event
	SetOnMouseEnter(fn TNotifyEvent)                    // property event
	SetOnMouseLeave(fn TNotifyEvent)                    // property event
	SetOnMouseMove(fn TMouseMoveEvent)                  // property event
	SetOnMouseUp(fn TMouseEvent)                        // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                // property event
	SetOnMouseWheelDown(fn TMouseWheelUpDownEvent)      // property event
	SetOnMouseWheelUp(fn TMouseWheelUpDownEvent)        // property event
	SetOnPictureChanged(fn TNotifyEvent)                // property event
	SetOnPaintBackground(fn TImagePaintBackgroundEvent) // property event
}

// TCustomImage Parent: TGraphicControl
type TCustomImage struct {
	TGraphicControl
	mouseDownPtr       uintptr
	mouseEnterPtr      uintptr
	mouseLeavePtr      uintptr
	mouseMovePtr       uintptr
	mouseUpPtr         uintptr
	mouseWheelPtr      uintptr
	mouseWheelDownPtr  uintptr
	mouseWheelUpPtr    uintptr
	pictureChangedPtr  uintptr
	paintBackgroundPtr uintptr
}

func NewCustomImage(AOwner IComponent) ICustomImage {
	r1 := LCL().SysCallN(1732, GetObjectUintptr(AOwner))
	return AsCustomImage(r1)
}

func (m *TCustomImage) HasGraphic() bool {
	r1 := LCL().SysCallN(1734, m.Instance())
	return GoBool(r1)
}

func (m *TCustomImage) AntialiasingMode() TAntialiasingMode {
	r1 := LCL().SysCallN(1729, 0, m.Instance(), 0)
	return TAntialiasingMode(r1)
}

func (m *TCustomImage) SetAntialiasingMode(AValue TAntialiasingMode) {
	LCL().SysCallN(1729, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImage) Center() bool {
	r1 := LCL().SysCallN(1730, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetCenter(AValue bool) {
	LCL().SysCallN(1730, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) KeepOriginXWhenClipped() bool {
	r1 := LCL().SysCallN(1738, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetKeepOriginXWhenClipped(AValue bool) {
	LCL().SysCallN(1738, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) KeepOriginYWhenClipped() bool {
	r1 := LCL().SysCallN(1739, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetKeepOriginYWhenClipped(AValue bool) {
	LCL().SysCallN(1739, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) ImageIndex() int32 {
	r1 := LCL().SysCallN(1735, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImage) SetImageIndex(AValue int32) {
	LCL().SysCallN(1735, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImage) ImageWidth() int32 {
	r1 := LCL().SysCallN(1736, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCustomImage) SetImageWidth(AValue int32) {
	LCL().SysCallN(1736, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomImage) Images() ICustomImageList {
	r1 := LCL().SysCallN(1737, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TCustomImage) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(1737, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomImage) Picture() IPicture {
	r1 := LCL().SysCallN(1740, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TCustomImage) SetPicture(AValue IPicture) {
	LCL().SysCallN(1740, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TCustomImage) Stretch() bool {
	r1 := LCL().SysCallN(1752, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetStretch(AValue bool) {
	LCL().SysCallN(1752, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) StretchOutEnabled() bool {
	r1 := LCL().SysCallN(1754, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetStretchOutEnabled(AValue bool) {
	LCL().SysCallN(1754, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) StretchInEnabled() bool {
	r1 := LCL().SysCallN(1753, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetStretchInEnabled(AValue bool) {
	LCL().SysCallN(1753, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) Transparent() bool {
	r1 := LCL().SysCallN(1755, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetTransparent(AValue bool) {
	LCL().SysCallN(1755, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) Proportional() bool {
	r1 := LCL().SysCallN(1741, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomImage) SetProportional(AValue bool) {
	LCL().SysCallN(1741, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomImage) DestRect() (resultRect TRect) {
	LCL().SysCallN(1733, m.Instance(), uintptr(unsafe.Pointer(&resultRect)))
	return
}

func CustomImageClass() TClass {
	ret := LCL().SysCallN(1731)
	return TClass(ret)
}

func (m *TCustomImage) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1742, m.Instance(), m.mouseDownPtr)
}

func (m *TCustomImage) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1743, m.Instance(), m.mouseEnterPtr)
}

func (m *TCustomImage) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1744, m.Instance(), m.mouseLeavePtr)
}

func (m *TCustomImage) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1745, m.Instance(), m.mouseMovePtr)
}

func (m *TCustomImage) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1746, m.Instance(), m.mouseUpPtr)
}

func (m *TCustomImage) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1747, m.Instance(), m.mouseWheelPtr)
}

func (m *TCustomImage) SetOnMouseWheelDown(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelDownPtr != 0 {
		RemoveEventElement(m.mouseWheelDownPtr)
	}
	m.mouseWheelDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1748, m.Instance(), m.mouseWheelDownPtr)
}

func (m *TCustomImage) SetOnMouseWheelUp(fn TMouseWheelUpDownEvent) {
	if m.mouseWheelUpPtr != 0 {
		RemoveEventElement(m.mouseWheelUpPtr)
	}
	m.mouseWheelUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1749, m.Instance(), m.mouseWheelUpPtr)
}

func (m *TCustomImage) SetOnPictureChanged(fn TNotifyEvent) {
	if m.pictureChangedPtr != 0 {
		RemoveEventElement(m.pictureChangedPtr)
	}
	m.pictureChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1751, m.Instance(), m.pictureChangedPtr)
}

func (m *TCustomImage) SetOnPaintBackground(fn TImagePaintBackgroundEvent) {
	if m.paintBackgroundPtr != 0 {
		RemoveEventElement(m.paintBackgroundPtr)
	}
	m.paintBackgroundPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(1750, m.Instance(), m.paintBackgroundPtr)
}
