//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"unsafe"
)

func (m *TLazIntfImage) DataDescription() *TRawImageDescription {
	return nil
}

func (m *TLazIntfImage) SetDataDescription(AValue *TRawImageDescription) {
}

func (m *TLazIntfImage) CheckDescription(ADescription *TRawImageDescription, ExceptionOnError bool) bool {
	return false
}

func (m *TLazIntfImage) SetDataDescriptionKeepData(ADescription *TRawImageDescription) {

}

// TRawImageDescription TODO no impl
//
//	Descriptor object for the image format of devices and raw (uncompressed) image data.
//	This effectively is a record with some attached methods. More related procedures exist outside the object.
//	The object describes the presence and exact storage of the RGBA image and mask pixels, of a device or image. The color information is stored in aligned scanlines.
//	Note: palettes, BitOrder and ByteOrder seem not to be implemented yet. The meaning of the PaletteXXX values and of BitOrder is undefined, so far.
//	xxxBitsPerPixel and xxxPrecisionMask applies to color data. For masked images, the pixels and the mask are two different arrays, containing different elements.
type TRawImageDescription struct {
	instance unsafe.Pointer
}

func NewRawImageDescription() *TRawImageDescription {
	return &TRawImageDescription{}
}

func (m *TRawImageDescription) Width() int32 {
	return 0
}

func (m *TRawImageDescription) Height() int32 {
	return 0
}

// 1-bit mono format

func (m *TRawImageDescription) Init_BPP1(width, height int32) {
}

// 16-bits formats

func (m *TRawImageDescription) Init_BPP16_R5G6B5(width, height int32) {
}

// Formats in RGB order

func (m *TRawImageDescription) Init_BPP24_R8G8B8_BIO_TTB(width, height int32) {
}

func (m *TRawImageDescription) Init_BPP24_R8G8B8_BIO_TTB_UpsideDown(width, height int32) {
}

func (m *TRawImageDescription) Init_BPP32_A8R8G8B8_BIO_TTB(width, height int32) {
}

func (m *TRawImageDescription) Init_BPP32_R8G8B8A8_BIO_TTB(width, height int32) {
}

// Formats in Windows pixels order: BGR

func (m *TRawImageDescription) Init_BPP24_B8G8R8_BIO_TTB(width, height int32) {
}

func (m *TRawImageDescription) Init_BPP24_B8G8R8_M1_BIO_TTB(width, height int32) {
}

func (m *TRawImageDescription) Init_BPP32_B8G8R8_BIO_TTB(width, height int32) {
}

func (m *TRawImageDescription) Init_BPP32_B8G8R8_M1_BIO_TTB(width, height int32) {
}

func (m *TRawImageDescription) Init_BPP32_B8G8R8A8_BIO_TTB(width, height int32) {
}

func (m *TRawImageDescription) Init_BPP32_B8G8R8A8_M1_BIO_TTB(width, height int32) {
}
