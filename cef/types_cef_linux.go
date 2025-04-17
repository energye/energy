//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build linux
// +build linux

// cef -> energy 所有结构类型定义 linux

package cef

import (
	"github.com/cyber-xxm/energy/v2/consts"
	. "github.com/cyber-xxm/energy/v2/types"
)

// TCefWindowInfo /include/internal/cef_types_win.h (cef_window_info_t)
type TCefWindowInfo struct {
	instance                   *tCefWindowInfoPtr
	WindowName                 TCefString
	X                          Integer
	Y                          Integer
	Width                      Integer
	Height                     Integer
	Hidden                     Integer
	ParentWindow               TCefWindowHandle
	WindowlessRenderingEnabled Integer
	SharedTextureEnabled       Integer
	ExternalBeginFrameEnabled  Integer
	Window                     TCefWindowHandle
}

// SetInstanceValue 实例指针设置值
func (m *TCefWindowInfo) setInstanceValue() {
	if m.instance == nil {
		return
	}
	// 字段指针引用赋值, 如果是字符串类型需直接赋值
	m.instance.WindowName = UIntptr(m.WindowName.ToPtr())                               // TCefString
	m.instance.X.SetValue(int32(m.X))                                                   // Integer
	m.instance.Y.SetValue(int32(m.Y))                                                   // Integer
	m.instance.Width.SetValue(int32(m.Width))                                           // Integer
	m.instance.Height.SetValue(int32(m.Height))                                         // Integer
	m.instance.Hidden.SetValue(int32(m.Hidden))                                         // Integer
	m.instance.ParentWindow.SetValue(uintptr(m.ParentWindow))                           // TCefWindowHandle
	m.instance.WindowlessRenderingEnabled.SetValue(int32(m.WindowlessRenderingEnabled)) // Integer
	m.instance.SharedTextureEnabled.SetValue(int32(m.SharedTextureEnabled))             // Integer
	m.instance.ExternalBeginFrameEnabled.SetValue(int32(m.ExternalBeginFrameEnabled))   // Integer
	m.instance.Window.SetValue(uintptr(m.Window))                                       // TCefWindowHandle
}

// / Structure containing the plane information of the shared texture.
// / Sync with native_pixmap_handle.h
// / <para><see href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_linux.h">CEF source file: /include/internal/cef_types_linux.h (cef_accelerated_paint_native_pixmap_plane_t)</see></para>
type TCefAcceleratedPaintNativePixmapPlaneInfo struct {
	/// The strides in bytes to be used when accessing the buffers via
	/// a memory mapping. One per plane per entry.
	stride Cardinal
	/// The offsets in bytes to be used when accessing the buffers via
	/// a memory mapping. One per plane per entry.
	offset uint64
	/// Size in bytes of the plane is necessary to map the buffers.
	size uint64
	/// File descriptor for the underlying memory object (usually dmabuf).
	fd int32
}

const CEF_KACCELERATEDPAINTMAXPLANES = 4

type TCefAcceleratedPaintInfo struct {
	/// Planes of the shared texture, usually file descriptors of dmabufs.
	planes [CEF_KACCELERATEDPAINTMAXPLANES]TCefAcceleratedPaintNativePixmapPlaneInfo
	/// Plane count.
	plane_count int32
	/// Modifier could be used with EGL driver.
	modifier uint64
	/// The pixel format of the texture.
	format consts.TCefColorType
}
