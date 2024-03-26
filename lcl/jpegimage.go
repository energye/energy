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
)

// IJPEGImage Parent: IFPImageBitmap
type IJPEGImage interface {
	IFPImageBitmap
	LoadFromBytes(data []byte)
	LoadFromFSFile(Filename string) error
	CompressionQuality() TJPEGQualityRange          // property
	SetCompressionQuality(AValue TJPEGQualityRange) // property
	GrayScale() bool                                // property
	MinHeight() int32                               // property
	SetMinHeight(AValue int32)                      // property
	MinWidth() int32                                // property
	SetMinWidth(AValue int32)                       // property
	ProgressiveEncoding() bool                      // property
	SetProgressiveEncoding(AValue bool)             // property
	Performance() TJPEGPerformance                  // property
	SetPerformance(AValue TJPEGPerformance)         // property
	Scale() TJPEGScale                              // property
	SetScale(AValue TJPEGScale)                     // property
	Smoothing() bool                                // property
	SetSmoothing(AValue bool)                       // property
	Compress()                                      // procedure
}

// TJPEGImage Parent: TFPImageBitmap
type TJPEGImage struct {
	TFPImageBitmap
}

func NewJPEGImage() IJPEGImage {
	r1 := LCL().SysCallN(3170)
	return AsJPEGImage(r1)
}

func (m *TJPEGImage) CompressionQuality() TJPEGQualityRange {
	r1 := LCL().SysCallN(3169, 0, m.Instance(), 0)
	return TJPEGQualityRange(r1)
}

func (m *TJPEGImage) SetCompressionQuality(AValue TJPEGQualityRange) {
	LCL().SysCallN(3169, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) GrayScale() bool {
	r1 := LCL().SysCallN(3171, m.Instance())
	return GoBool(r1)
}

func (m *TJPEGImage) MinHeight() int32 {
	r1 := LCL().SysCallN(3172, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TJPEGImage) SetMinHeight(AValue int32) {
	LCL().SysCallN(3172, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) MinWidth() int32 {
	r1 := LCL().SysCallN(3173, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TJPEGImage) SetMinWidth(AValue int32) {
	LCL().SysCallN(3173, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) ProgressiveEncoding() bool {
	r1 := LCL().SysCallN(3175, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TJPEGImage) SetProgressiveEncoding(AValue bool) {
	LCL().SysCallN(3175, 1, m.Instance(), PascalBool(AValue))
}

func (m *TJPEGImage) Performance() TJPEGPerformance {
	r1 := LCL().SysCallN(3174, 0, m.Instance(), 0)
	return TJPEGPerformance(r1)
}

func (m *TJPEGImage) SetPerformance(AValue TJPEGPerformance) {
	LCL().SysCallN(3174, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) Scale() TJPEGScale {
	r1 := LCL().SysCallN(3176, 0, m.Instance(), 0)
	return TJPEGScale(r1)
}

func (m *TJPEGImage) SetScale(AValue TJPEGScale) {
	LCL().SysCallN(3176, 1, m.Instance(), uintptr(AValue))
}

func (m *TJPEGImage) Smoothing() bool {
	r1 := LCL().SysCallN(3177, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TJPEGImage) SetSmoothing(AValue bool) {
	LCL().SysCallN(3177, 1, m.Instance(), PascalBool(AValue))
}

func JPEGImageClass() TClass {
	ret := LCL().SysCallN(3167)
	return TClass(ret)
}

func (m *TJPEGImage) Compress() {
	LCL().SysCallN(3168, m.Instance())
}
