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

// linux => dbus tray

package application

type TTrayIcon struct {
	trayMenu *TTrayMenu
}

type TTrayMenu struct {
	imageList *TTrayImageList
}

type TTrayMenuItem struct {
	menu *TTrayMenu
}

// NewTrayIcon 创建并初始化一个新的系统托盘图标实例
func NewTrayIcon() *TTrayIcon {
	m := &TTrayIcon{}

	return m
}
