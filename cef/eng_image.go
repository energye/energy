//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICefImage Parent: ICefBaseRefCounted
//
//	Container for a single image represented at different scale factors. All image representations should be the same size in density independent pixel (DIP) units. For example, if the image at scale factor 1.0 is 100x100 pixels then the image at scale factor 2.0 should be 200x200 pixels -- both images will display with a DIP size of 100x100 units. The functions of this interface can be called on any browser process thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_image_capi.h">CEF source file: /include/capi/cef_image_capi.h (cef_image_t))
type ICefImage interface {
	ICefBaseRefCounted
	// IsEmpty
	//  Returns true (1) if this Image is NULL.
	IsEmpty() bool // function
	// IsSame
	//  Returns true (1) if this Image and |that| Image share the same underlying storage. Will also return true (1) if both images are NULL.
	IsSame(that ICefImage) bool // function
	// AddBitmap
	//  Add a bitmap image representation for |scale_factor|. Only 32-bit RGBA/BGRA formats are supported. |pixel_width| and |pixel_height| are the bitmap representation size in pixel coordinates. |pixel_data| is the array of pixel data and should be |pixel_width| x |pixel_height| x 4 bytes in size. |color_type| and |alpha_type| values specify the pixel format.
	AddBitmap(scaleFactor float32, pixelWidth, pixelHeight int32, colorType TCefColorType, alphaType TCefAlphaType, pixelData uintptr, pixelDataSize NativeUInt) bool // function
	// AddPng
	//  Add a PNG image representation for |scale_factor|. |png_data| is the image data of size |png_data_size|. Any alpha transparency in the PNG data will be maintained.
	AddPng(scaleFactor float32, pngData uintptr, pngDataSize NativeUInt) bool // function
	// AddJpeg
	//  Create a JPEG image representation for |scale_factor|. |jpeg_data| is the image data of size |jpeg_data_size|. The JPEG format does not support transparency so the alpha byte will be set to 0xFF for all pixels.
	AddJpeg(scaleFactor float32, jpegData uintptr, jpegDataSize NativeUInt) bool // function
	// GetWidth
	//  Returns the image width in density independent pixel (DIP) units.
	GetWidth() NativeUInt // function
	// GetHeight
	//  Returns the image height in density independent pixel (DIP) units.
	GetHeight() NativeUInt // function
	// HasRepresentation
	//  Returns true (1) if this image contains a representation for |scale_factor|.
	HasRepresentation(scaleFactor float32) bool // function
	// RemoveRepresentation
	//  Removes the representation for |scale_factor|. Returns true (1) on success.
	RemoveRepresentation(scaleFactor float32) bool // function
	// GetRepresentationInfo
	//  Returns information for the representation that most closely matches |scale_factor|. |actual_scale_factor| is the actual scale factor for the representation. |pixel_width| and |pixel_height| are the representation size in pixel coordinates. Returns true (1) on success.
	GetRepresentationInfo(scaleFactor float32, actualScaleFactor *float32, pixelWidth, pixelHeight *int32) bool // function
	// GetAsBitmap
	//  Returns the bitmap representation that most closely matches |scale_factor|. Only 32-bit RGBA/BGRA formats are supported. |color_type| and |alpha_type| values specify the desired output pixel format. |pixel_width| and |pixel_height| are the output representation size in pixel coordinates. Returns a ICefBinaryValue containing the pixel data on success or NULL on failure.
	GetAsBitmap(scaleFactor float32, colorType TCefColorType, alphaType TCefAlphaType, pixelWidth, pixelHeight *int32) ICefBinaryValue // function
	// GetAsPng
	//  Returns the PNG representation that most closely matches |scale_factor|. If |with_transparency| is true (1) any alpha transparency in the image will be represented in the resulting PNG data. |pixel_width| and |pixel_height| are the output representation size in pixel coordinates. Returns a ICefBinaryValue containing the PNG image data on success or NULL on failure.
	GetAsPng(scaleFactor float32, withTransparency bool, pixelWidth, pixelHeight *int32) ICefBinaryValue // function
	// GetAsJpeg
	//  Returns the JPEG representation that most closely matches |scale_factor|. |quality| determines the compression level with 0 == lowest and 100 == highest. The JPEG format does not support alpha transparency and the alpha channel, if any, will be discarded. |pixel_width| and |pixel_height| are the output representation size in pixel coordinates. Returns a ICefBinaryValue containing the JPEG image data on success or NULL on failure.
	GetAsJpeg(scaleFactor float32, quality int32, pixelWidth, pixelHeight *int32) ICefBinaryValue // function
}

// TCefImage Parent: TCefBaseRefCounted
//
//	Container for a single image represented at different scale factors. All image representations should be the same size in density independent pixel (DIP) units. For example, if the image at scale factor 1.0 is 100x100 pixels then the image at scale factor 2.0 should be 200x200 pixels -- both images will display with a DIP size of 100x100 units. The functions of this interface can be called on any browser process thread.
//	 <a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/capi/cef_image_capi.h">CEF source file: /include/capi/cef_image_capi.h (cef_image_t))
type TCefImage struct {
	TCefBaseRefCounted
}

// ImageRef -> ICefImage
var ImageRef image

// image TCefImage Ref
type image uintptr

func (m *image) UnWrap(data uintptr) ICefImage {
	var resultCefImage uintptr
	CEF().SysCallN(1000, uintptr(data), uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *image) New() ICefImage {
	var resultCefImage uintptr
	CEF().SysCallN(998, uintptr(unsafePointer(&resultCefImage)))
	return AsCefImage(resultCefImage)
}

func (m *TCefImage) IsEmpty() bool {
	r1 := CEF().SysCallN(996, m.Instance())
	return GoBool(r1)
}

func (m *TCefImage) IsSame(that ICefImage) bool {
	r1 := CEF().SysCallN(997, m.Instance(), GetObjectUintptr(that))
	return GoBool(r1)
}

func (m *TCefImage) AddBitmap(scaleFactor float32, pixelWidth, pixelHeight int32, colorType TCefColorType, alphaType TCefAlphaType, pixelData uintptr, pixelDataSize NativeUInt) bool {
	r1 := CEF().SysCallN(986, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(pixelWidth), uintptr(pixelHeight), uintptr(colorType), uintptr(alphaType), uintptr(pixelData), uintptr(pixelDataSize))
	return GoBool(r1)
}

func (m *TCefImage) AddPng(scaleFactor float32, pngData uintptr, pngDataSize NativeUInt) bool {
	r1 := CEF().SysCallN(988, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(pngData), uintptr(pngDataSize))
	return GoBool(r1)
}

func (m *TCefImage) AddJpeg(scaleFactor float32, jpegData uintptr, jpegDataSize NativeUInt) bool {
	r1 := CEF().SysCallN(987, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(jpegData), uintptr(jpegDataSize))
	return GoBool(r1)
}

func (m *TCefImage) GetWidth() NativeUInt {
	r1 := CEF().SysCallN(994, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefImage) GetHeight() NativeUInt {
	r1 := CEF().SysCallN(992, m.Instance())
	return NativeUInt(r1)
}

func (m *TCefImage) HasRepresentation(scaleFactor float32) bool {
	r1 := CEF().SysCallN(995, m.Instance(), uintptr(unsafePointer(&scaleFactor)))
	return GoBool(r1)
}

func (m *TCefImage) RemoveRepresentation(scaleFactor float32) bool {
	r1 := CEF().SysCallN(999, m.Instance(), uintptr(unsafePointer(&scaleFactor)))
	return GoBool(r1)
}

func (m *TCefImage) GetRepresentationInfo(scaleFactor float32, actualScaleFactor *float32, pixelWidth, pixelHeight *int32) bool {
	var result1 uintptr
	var result2 uintptr
	var result3 uintptr
	r1 := CEF().SysCallN(993, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(unsafePointer(&result1)), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)))
	*actualScaleFactor = float32(result1)
	*pixelWidth = int32(result2)
	*pixelHeight = int32(result3)
	return GoBool(r1)
}

func (m *TCefImage) GetAsBitmap(scaleFactor float32, colorType TCefColorType, alphaType TCefAlphaType, pixelWidth, pixelHeight *int32) ICefBinaryValue {
	var result3 uintptr
	var result4 uintptr
	var resultCefBinaryValue uintptr
	CEF().SysCallN(989, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(colorType), uintptr(alphaType), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&result4)), uintptr(unsafePointer(&resultCefBinaryValue)))
	*pixelWidth = int32(result3)
	*pixelHeight = int32(result4)
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefImage) GetAsPng(scaleFactor float32, withTransparency bool, pixelWidth, pixelHeight *int32) ICefBinaryValue {
	var result2 uintptr
	var result3 uintptr
	var resultCefBinaryValue uintptr
	CEF().SysCallN(991, m.Instance(), uintptr(unsafePointer(&scaleFactor)), PascalBool(withTransparency), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&resultCefBinaryValue)))
	*pixelWidth = int32(result2)
	*pixelHeight = int32(result3)
	return AsCefBinaryValue(resultCefBinaryValue)
}

func (m *TCefImage) GetAsJpeg(scaleFactor float32, quality int32, pixelWidth, pixelHeight *int32) ICefBinaryValue {
	var result2 uintptr
	var result3 uintptr
	var resultCefBinaryValue uintptr
	CEF().SysCallN(990, m.Instance(), uintptr(unsafePointer(&scaleFactor)), uintptr(quality), uintptr(unsafePointer(&result2)), uintptr(unsafePointer(&result3)), uintptr(unsafePointer(&resultCefBinaryValue)))
	*pixelWidth = int32(result2)
	*pixelHeight = int32(result3)
	return AsCefBinaryValue(resultCefBinaryValue)
}
