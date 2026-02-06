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
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type TTrayImageList struct {
	imageList  [][]byte
	imageIndex map[string]int32
}

func (m *TTrayImageList) ImageIndex(imageName string) int32 {
	index, ok := m.imageIndex[strings.ToLower(imageName)]
	if ok {
		return index
	}
	return -1
}

func (m *TTrayImageList) setImageListData(data []byte, name string, index int32) {
	pic := lcl.NewPicture()
	defer pic.Free()
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, data)
	mem.SetPosition(0)
	pic.LoadFromStream(mem)
	m.imageList = append(m.imageList, data)
	if name != "" && index != -1 {
		m.imageIndex[name] = index
	}
}

type TTrayIcon struct {
	tray     *systray.Tray
	trayMenu *TTrayMenu
	visible  bool
}

type TTrayMenu struct {
	imageList *TTrayImageList
	trayIcon  *TTrayIcon
}

type TTrayMenuItem struct {
	item *systray.MenuItem
	menu *TTrayMenu
}

// NewTrayIcon 创建并初始化一个新的系统托盘图标实例
func NewTrayIcon() *TTrayIcon {
	tray := systray.NativeStart()
	m := &TTrayIcon{tray: tray}
	api.SetOnReleaseCallback(func() {
		m.Close()
	})
	return m
}

func (m *TTrayIcon) Close() {
	if m.tray != nil {
		m.tray.NativeEnd()
		m.tray = nil
	}
}

// tray

func (m *TTrayIcon) SetOnClick(fn func()) {
	m.tray.SetOnClick(fn)
}

func (m *TTrayIcon) SetOnDblClick(fn func()) {
	m.tray.SetOnDClick(fn)
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
	m.tray.SetVisible(m.visible)
}

func (m *TTrayIcon) Hide() {
	m.visible = false
	m.tray.SetVisible(m.visible)
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
	m.tray.SetIcon(data)
}

func (m *TTrayIcon) SetHint(hint string) {
	m.tray.SetTooltip(hint)
}

// tray menu

func (m *TTrayIcon) Menu() *TTrayMenu {
	if m.trayMenu == nil {
		m.trayMenu = &TTrayMenu{trayIcon: m}
	}
	return m.trayMenu
}

func (m *TTrayMenu) mustImageList() {
	if m.imageList == nil {
		m.imageList = &TTrayImageList{imageList: make([][]byte, 0), imageIndex: make(map[string]int32)}
	}
}

// SetImageList 设置托盘菜单的图像列表
//   - pngImagePathList: PNG图像文件路径列表
//   - size: 图片尺寸
func (m *TTrayMenu) SetImageList(pngImagePathList []string) *TTrayImageList {
	m.mustImageList()
	imageListAddPng := func(filePath string, name string, index int32) {
		data, err := os.ReadFile(filePath)
		if data != nil && err == nil {
			m.imageList.setImageListData(data, name, index)
		}
	}
	for index, image := range pngImagePathList {
		_, name := filepath.Split(image)
		name = strings.ToLower(name)
		imageListAddPng(image, name, int32(index))
	}
	return m.imageList
}

// SetImageListEmbed 设置嵌入式图片列表到托盘菜单中
//   - embed: 嵌入文件系统接口，用于读取嵌入的图片资源
//   - pngImageEmbedPathList: PNG图片的嵌入路径列表
//   - size: 图片尺寸
func (m *TTrayMenu) SetImageListEmbed(embed emfs.IEmbedFS, pngImageEmbedPathList []string) *TTrayImageList {
	m.mustImageList()
	imageListAddPng := func(imagePath string, name string, index int32) {
		data, err := embed.ReadFile(imagePath)
		if data != nil && err == nil {
			m.imageList.setImageListData(data, name, index)
		}
	}
	for index, image := range pngImageEmbedPathList {
		_, name := filepath.Split(image)
		name = strings.ToLower(name)
		imageListAddPng(image, name, int32(index))
	}
	return m.imageList
}

// SetImageListDataBytes 设置图像列表的数据字节
//   - pngImageDataList: PNG图像数据字节数组的切片
//   - size: 图像尺寸
func (m *TTrayMenu) SetImageListDataBytes(pngImageDataList [][]byte) {
	m.mustImageList()
	for _, data := range pngImageDataList {
		if data != nil && len(data) > 0 {
			m.imageList.setImageListData(data, "", -1)
		}
	}
}

// tray menu item

func (m *TTrayMenu) tray() *systray.Tray {
	return m.trayIcon.tray
}

// AddMenuItem 向托盘菜单中添加一个新的菜单项
//   - label - 菜单项显示的文本标签
//   - fn - 菜单项被点击时执行的回调函数，可以为nil表示无点击事件
func (m *TTrayMenu) AddMenuItem(label string) *TTrayMenuItem {
	newMenuItem := m.tray().Menu().AddMenuItem(m.tray(), label, "")
	menuItem := &TTrayMenuItem{menu: m, item: newMenuItem}
	return menuItem
}

// AddSeparator 向系统托盘菜单中添加一个分隔符
func (m *TTrayMenu) AddSeparator() {
	m.trayIcon.tray.Menu().AddSeparator(m.trayIcon.tray)
}

func (m *TTrayMenuItem) tray() *systray.Tray {
	return m.menu.trayIcon.tray
}

// AddSubMenuItem 添加子菜单项到当前菜单项
//   - label - 菜单项显示的标签文本
//   - fn - 点击菜单项时执行的回调函数，可以为nil表示无点击事件
func (m *TTrayMenuItem) AddSubMenuItem(label string) *TTrayMenuItem {
	newMenuItem := m.item.AddMenuItem(m.tray(), label, "")
	menuItem := &TTrayMenuItem{menu: m.menu, item: newMenuItem}
	return menuItem
}

// AddSeparator 向系统托盘菜单中添加一个分隔符
func (m *TTrayMenuItem) AddSeparator() {
	m.item.AddSeparator(m.tray())
}

// SetImage 设置菜单项的图标
//   - imageName: 图标名称，用于在图像列表中查找对应的图标索引
//
// 说明:
//   - 该方法会检查菜单及其图像列表是否存在，如果存在则根据图标名称获取索引并设置到菜单项上
func (m *TTrayMenuItem) SetImage(imageName string) *TTrayMenuItem {
	if m.menu != nil && m.menu.imageList != nil {
		if imageIndex := m.menu.imageList.ImageIndex(imageName); imageIndex != -1 && int(imageIndex) < len(m.menu.imageList.imageList) {
			data := m.menu.imageList.imageList[imageIndex]
			m.item.SetIcon(m.tray(), data)
		}
	}
	return m
}

func (m *TTrayMenuItem) SetBitmap(image []byte) *TTrayMenuItem {
	if image == nil {
		return
	}
	m.item.SetIcon(m.tray(), image)
	return m
}

// SetImageIndex 设置菜单项的图像索引
//
// - index - 要设置的图像索引，必须为非负整数
//
// 说明:
//   - 该方法会检查菜单及其图像列表是否存在，如果存在则根据图标名称获取索引并设置到菜单项上
func (m *TTrayMenuItem) SetImageIndex(index int32) *TTrayMenuItem {
	if m.menu != nil && m.menu.imageList != nil && index >= 0 && int(index) < len(m.menu.imageList.imageList) {
		data := m.menu.imageList.imageList[index]
		m.item.SetIcon(m.tray(), data)
	}
	return m
}

func (m *TTrayMenuItem) SetChecked(checked bool) *TTrayMenuItem {
	m.item.SetChecked(m.tray(), checked)
	return m
}

func (m *TTrayMenuItem) SetRadio(radio bool) *TTrayMenuItem {
	m.item.SetRadio(m.tray(), radio)
	return m
}

func (m *TTrayMenuItem) SetEnabled(enabled bool) *TTrayMenuItem {
	m.item.SetEnabled(m.tray(), enabled)
	return m
}

func (m *TTrayMenuItem) Enabled() bool {
	return m.item.Enabled(m.tray())
}

func (m *TTrayMenuItem) Checked() bool {
	return m.item.Checked()
}

func (m *TTrayMenuItem) Clear() {
	m.item.Clear(m.tray())
}

func (m *TTrayMenuItem) SetOnMeasureItem(fn lcl.TMenuMeasureItemEvent) *TTrayMenuItem {
	log.Println("SetOnMeasureItem No Implementation")
	return m
}

func (m *TTrayMenuItem) SetOnDrawItem(fn lcl.TMenuDrawItemEvent) *TTrayMenuItem {
	log.Println("SetOnDrawItem No Implementation")
	return m
}

func (m *TTrayMenuItem) SetOnClick(fn func()) *TTrayMenuItem {
	if fn != nil {
		m.item.Click(func() {
			lcl.RunOnMainThreadAsync(func(id uint32) {
				fn()
			})
		})
	}
	return m
}
