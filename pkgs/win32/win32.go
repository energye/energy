//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build windows

package win32

// --------------------------  磨砂效果实现核心函数  --------------------------
// EnableWindowBlur：为已有 HWND 窗口启用磨砂玻璃效果
// hwnd：Lazarus 窗口的原生句柄（需转换为 syscall.Handle）
// alpha：窗口透明度（0-255，240 为半透明，效果更优）
//func EnableWindowBlur(hwnd types.HWND, alpha byte) error {
//	// 步骤 1：获取窗口当前扩展样式，添加 WS_EX_LAYERED 分层样式
//	currentExStyle := win.GetWindowLongPtr(hwnd, win.GWL_EXSTYLE)
//	// 追加 WS_EX_LAYERED 样式（按位或）
//	newExStyle := currentExStyle | uintptr(win.WS_EX_LAYERED)
//	win.SetWindowLongPtr(hwnd, win.GWL_EXSTYLE, newExStyle)
//
//	// 步骤 2：设置窗口半透明（可选，让磨砂效果更明显）
//	win.SetLayeredWindowAttributes(
//		uintptr(hwnd),
//		0,             // 透明色（此处不使用，传 0）
//		alpha,         // 透明度（0 全透，255 不透明，推荐 240）
//		win.LWA_ALPHA, // 启用透明度配置
//	)
//
//	// 步骤 3：创建全屏模糊区域（CreateRectRgn(0,0,-1,-1) 表示整个窗口）
//	rgn, _, err := createRectRgn.Call(0, 0, ^uintptr(0)>>1, ^uintptr(0)>>1) // 对应 0,0,-1,-1
//	if err != syscall.Errno(0) {
//		return err
//	}
//	defer deleteObject.Call(rgn) // 延迟释放区域句柄，避免内存泄漏
//
//	// 步骤 4：配置 DWM 模糊背景结构体
//	blurBehind := DWMBlurBehind{
//		dwFlags:                DWM_BB_ENABLE | DWM_BB_BLURREGION,
//		fEnable:                true,
//		hRgnBlur:               syscall.Handle(rgn),
//		fTransitionOnMaximized: false,
//	}
//
//	// 步骤 5：调用 DwmEnableBlurBehindWindow 启用磨砂效果
//	_, _, err = dwmEnableBlurBehindWindow.Call(
//		uintptr(hwnd),
//		uintptr(unsafe.Pointer(&blurBehind)),
//	)
//	if err != syscall.Errno(0) {
//		return err
//	}
//
//	return nil
//}
