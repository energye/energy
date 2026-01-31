//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// macOS, windows => LCL tray

package application

import (
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"os"
	"path/filepath"
	"strings"
)

type TTrayImageList struct {
	imageList  lcl.IImageList
	imageIndex map[string]int32
}

func (m *TTrayImageList) ImageIndex(imageName string) int32 {
	index, ok := m.imageIndex[strings.ToLower(imageName)]
	if ok {
		return index
	}
	return -1
}

type TTrayIcon struct {
	owner    lcl.IComponent
	trayIcon lcl.ITrayIcon
	trayMenu *TTrayMenu
}

type TTrayMenu struct {
	trayMenu  lcl.IPopupMenu
	imageList *TTrayImageList
}

type TTrayMenuItem struct {
	menu *TTrayMenu
	item lcl.IMenuItem
}

func NewTrayIcon() *TTrayIcon {
	m := &TTrayIcon{owner: lcl.Application}
	m.trayIcon = lcl.NewTrayIcon(m.owner)
	return m
}

// tray

func (m *TTrayIcon) Show() {
	m.trayIcon.SetVisible(true)
}

func (m *TTrayIcon) Hide() {
	m.trayIcon.SetVisible(false)
}

func (m *TTrayIcon) SetIcon(png string) {
	if data, err := os.ReadFile(png); err == nil {
		m.SetIconBytes(data)
	}
}

func (m *TTrayIcon) SetIconBytes(data []byte) {
	pic := lcl.NewPicture()
	defer pic.Free()
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, data)
	mem.SetPosition(0)
	pic.LoadFromStream(mem)
	m.trayIcon.Icon().Assign(pic)
}

func (m *TTrayIcon) SetHint(hint string) {
	m.trayIcon.SetHint(hint)
}

// tray menu

func (m *TTrayIcon) Menu() *TTrayMenu {
	if m.trayMenu == nil {
		popupMenu := lcl.NewPopupMenu(m.owner)
		m.trayMenu = &TTrayMenu{trayMenu: popupMenu}
		m.trayIcon.SetPopUpMenu(popupMenu)
	}
	return m.trayMenu
}

func (m *TTrayMenu) mustImageList(size types.TSize) {
	if m.imageList == nil {
		m.imageList = &TTrayImageList{imageList: lcl.NewImageList(m.trayMenu), imageIndex: make(map[string]int32)}
		if size.Cx > 0 {
			m.imageList.imageList.SetWidth(size.Cx)
		}
		if size.Cy > 0 {
			m.imageList.imageList.SetHeight(size.Cy)
		}
		m.trayMenu.SetImages(m.imageList.imageList)
	}
}

func (m *TTrayMenu) setImageListData(data []byte) {
	pic := lcl.NewPicture()
	defer pic.Free()
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, data)
	mem.SetPosition(0)
	pic.LoadFromStream(mem)
	m.imageList.imageList.Add(pic.Bitmap(), nil)
}

func (m *TTrayMenu) SetImageList(pngImagePathList []string, size types.TSize) *TTrayImageList {
	m.mustImageList(size)
	imageListAddPng := func(filePath string) {
		data, err := os.ReadFile(filePath)
		if data != nil && err == nil {
			m.setImageListData(data)
		}
	}
	for index, image := range pngImagePathList {
		imageListAddPng(image)
		_, name := filepath.Split(image)
		name = strings.ToLower(name)
		m.imageList.imageIndex[name] = int32(index)
	}
	return m.imageList
}

func (m *TTrayMenu) SetImageListEmbed(embed emfs.IEmbedFS, pngImageEmbedPathList []string, size types.TSize) *TTrayImageList {
	m.mustImageList(size)
	imageListAddPng := func(imagePath string) {
		data, err := embed.ReadFile(imagePath)
		if data != nil && err == nil {
			m.setImageListData(data)
		}
	}
	for index, image := range pngImageEmbedPathList {
		imageListAddPng(image)
		_, name := filepath.Split(image)
		name = strings.ToLower(name)
		m.imageList.imageIndex[name] = int32(index)
	}
	return m.imageList
}

func (m *TTrayMenu) SetImageListDataBytes(pngImageDataList [][]byte, size types.TSize) {
	m.mustImageList(size)
	for _, data := range pngImageDataList {
		if data != nil && len(data) > 0 {
			m.setImageListData(data)
		}
	}
}

// tray menu item

func (m *TTrayMenu) AddMenuItem(label string, fn func()) *TTrayMenuItem {
	newMenuItem := lcl.NewMenuItem(m.trayMenu)
	menuItem := &TTrayMenuItem{menu: m, item: newMenuItem}
	newMenuItem.SetCaption(label)
	m.trayMenu.Items().Add(newMenuItem)
	if fn != nil {
		newMenuItem.SetOnClick(func(sender lcl.IObject) {
			fn()
		})
	}
	return menuItem
}

func (m *TTrayMenuItem) AddSubMenuItem(label string, fn func()) *TTrayMenuItem {
	newMenuItem := lcl.NewMenuItem(m.item)
	menuItem := &TTrayMenuItem{menu: m.menu, item: newMenuItem}
	newMenuItem.SetCaption(label)
	m.item.Add(newMenuItem)
	if fn != nil {
		newMenuItem.SetOnClick(func(sender lcl.IObject) {
			fn()
		})
	}
	return menuItem
}

func (m *TTrayMenuItem) SetImage(imageName string) {
	if m.menu != nil && m.menu.imageList != nil {
		if imageIndex := m.menu.imageList.ImageIndex(imageName); imageIndex != -1 {
			m.item.SetImageIndex(imageIndex)
		}
	}
}

func (m *TTrayMenuItem) SetImageIndex(index int32) {
	if m.menu != nil && m.menu.imageList != nil && index >= 0 {
		m.item.SetImageIndex(index)
	}
}

func (m *TTrayMenuItem) SetChecked(checked bool) {
	m.item.SetChecked(checked)
}

func (m *TTrayMenuItem) Checked() bool {
	return m.item.Checked()
}

func (m *TTrayMenuItem) Clear() {
	m.item.Clear()
}
