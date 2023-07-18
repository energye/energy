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

package cef

// 窗口拖拽JS扩展
// 在这里执行并启用JS拖拽
func dragExtensionJS(frame *ICefFrame) {

}

// 窗口拖拽JS扩展处理器
//  1. 注册JS扩展到CEF, 注册鼠标事件，通过本地函数在Go里处理鼠标事件
//  2. 通过IPC将鼠标消息发送到主进程，主进程监听到消息处理鼠标事件
//  3. linux 默认使用VF窗口组件, 所以不需要使用该实现
func dragExtensionHandler() {

}

func (m *drag) drag() {

}
