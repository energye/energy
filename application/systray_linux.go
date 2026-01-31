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

import (
	"github.com/energye/energy/v3/application/internal/systray"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"log"
	"os"
)

type TTrayIcon struct {
	trayMenu *TTrayMenu
	visible  bool
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
	systray.NativeStart()
	return m
}

func (m *TTrayIcon) Close() {
	systray.NativeEnd()
}

// tray

func (m *TTrayIcon) SetOnClick(fn func()) {
	systray.SetOnClick(fn)
}

func (m *TTrayIcon) SetOnDblClick(fn func()) {
	systray.SetOnDClick(fn)
}

func (m *TTrayIcon) SetOnMouseUp(fn func(button types.TMouseButton, shift types.TShiftState, x, y int32)) {
	log.Println("SetOnMouseUp No Implementation")
}

func (m *TTrayIcon) SetOnMouseDown(fn func(button types.TMouseButton, shift types.TShiftState, x, y int32)) {
	log.Println("SetOnMouseDown No Implementation")
}

func (m *TTrayIcon) SetOnMouseMove(fn func(shift types.TShiftState, x, y int32)) {
	log.Println("SetOnMouseMove No Implementation")
}

func (m *TTrayIcon) Show() {
	m.visible = true
	systray.SetVisible(m.visible)
}

func (m *TTrayIcon) Hide() {
	m.visible = false
	systray.SetVisible(m.visible)
}

func (m *TTrayIcon) Visible() bool {
	return m.visible
}

func (m *TTrayIcon) SetIcon(png string) {
	if data, err := os.ReadFile(png); err == nil {
		m.SetIconBytes(data)
	}
}

func (m *TTrayIcon) SetIconBytes(data []byte) {
	if data == nil || len(data) == 0 {
		return
	}
	pic := lcl.NewPicture()
	defer pic.Free()
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, data)
	mem.SetPosition(0)
	pic.LoadFromStream(mem)
	//m.trayIcon.Icon().Assign(pic.Bitmap())
}

func (m *TTrayIcon) SetHint(hint string) {
	systray.SetTooltip(hint)
}
