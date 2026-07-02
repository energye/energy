//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Image is a representation of GTK's GtkImage.
type Image struct {
	Widget
}

func AsImage(ptr unsafe.Pointer) IImage {
	if ptr == nil {
		return nil
	}
	m := new(Image)
	m.instance = ptr
	return m
}

// NewImage is a wrapper around gtk_image_new().
func NewImage() IImage {
	r := gtk3.SysCall("gtk_image_new")
	if r == 0 {
		return nil
	}
	return AsImage(unsafe.Pointer(r))
}

// NewImageFromFile is a wrapper around gtk_image_new_from_file().
func NewImageFromFile(filename string) IImage {
	r := gtk3.SysCall("gtk_image_new_from_file", CStr(filename))
	if r == 0 {
		return nil
	}
	return AsImage(unsafe.Pointer(r))
}

// NewImageFromResource is a wrapper around gtk_image_new_from_resource().
func NewImageFromResource(resourcePath string) IImage {
	r := gtk3.SysCall("gtk_image_new_from_resource", CStr(resourcePath))
	if r == 0 {
		return nil
	}
	return AsImage(unsafe.Pointer(r))
}

// NewImageFromIconName is a wrapper around gtk_image_new_from_icon_name().
func NewImageFromIconName(iconName string, size IconSize) IImage {
	r := gtk3.SysCall("gtk_image_new_from_icon_name", CStr(iconName), uintptr(size))
	if r == 0 {
		return nil
	}
	return AsImage(unsafe.Pointer(r))
}

// Clear is a wrapper around gtk_image_clear().
func (m *Image) Clear() {
	gtk3.SysCall("gtk_image_clear", m.Instance())
}

// SetFromFile is a wrapper around gtk_image_set_from_file().
func (m *Image) SetFromFile(filename string) {
	gtk3.SysCall("gtk_image_set_from_file", m.Instance(), CStr(filename))
}

// SetFromResource is a wrapper around gtk_image_set_from_resource().
func (m *Image) SetFromResource(resourcePath string) {
	gtk3.SysCall("gtk_image_set_from_resource", m.Instance(), CStr(resourcePath))
}

// SetFromIconName is a wrapper around gtk_image_set_from_icon_name().
func (m *Image) SetFromIconName(iconName string, size IconSize) {
	gtk3.SysCall("gtk_image_set_from_icon_name", m.Instance(), CStr(iconName), uintptr(size))
}

// SetPixelSize is a wrapper around gtk_image_set_pixel_size().
func (m *Image) SetPixelSize(pixelSize int) {
	gtk3.SysCall("gtk_image_set_pixel_size", m.Instance(), uintptr(pixelSize))
}

// GetStorageType is a wrapper around gtk_image_get_storage_type().
func (m *Image) GetStorageType() ImageType {
	r := gtk3.SysCall("gtk_image_get_storage_type", m.Instance())
	return ImageType(r)
}

// GetIconName is a wrapper around gtk_image_get_icon_name().
func (m *Image) GetIconName() (string, IconSize) {
	var iconName uintptr
	var size int32
	gtk3.SysCall("gtk_image_get_icon_name", m.Instance(), uintptr(unsafe.Pointer(&iconName)), uintptr(unsafe.Pointer(&size)))
	return GoStr(iconName), IconSize(size)
}

// GetPixelSize is a wrapper around gtk_image_get_pixel_size().
func (m *Image) GetPixelSize() int {
	r := gtk3.SysCall("gtk_image_get_pixel_size", m.Instance())
	return int(r)
}

// SetFromGIcon is a wrapper around gtk_image_set_from_gicon().
func (m *Image) SetFromGIcon(icon *Icon, size IconSize) {
	gtk3.SysCall("gtk_image_set_from_gicon", m.Instance(), icon.NativePrivate(), uintptr(size))
}

// GetGIcon is a wrapper around gtk_image_get_gicon().
func (m *Image) GetGIcon() (*Icon, IconSize, error) {
	var gicon uintptr
	var size int32
	gtk3.SysCall("gtk_image_get_gicon", m.Instance(),
		uintptr(unsafe.Pointer(&gicon)), uintptr(unsafe.Pointer(&size)))
	if gicon == 0 {
		return nil, ICON_SIZE_INVALID, errNilPtr
	}
	return AsIcon(unsafe.Pointer(gicon)), IconSize(size), nil
}
