// Package systray is a cross-platform Go library to place an icon and menu in the notification area.
package systray

import (
	"fmt"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
	"github.com/godbus/dbus/v5"
	"log"
	"sync"
	"sync/atomic"
)

var (
	menuItems             = make(map[uint32]*MenuItem)
	menuItemsLock         sync.RWMutex
	currentID                   = uint32(0)
	dClickTimeMinInterval int64 = 500
)

type IMenu interface {
	ShowMenu() error
}

// MenuItem is used to keep track each menu item of systray.
// Don't create it directly, use the one systray.AddMenuItem() returned
type MenuItem struct {
	// ClickedCh is the channel which will be notified when the menu item is clicked
	click func()

	// id uniquely identify a menu item, not supposed to be modified
	id uint32
	// title is the text shown on menu item
	title string
	// tooltip is the text shown when pointing to menu item
	tooltip string
	// shortcutKey Menu shortcut key
	shortcutKey string
	// disabled menu item is grayed out and has no effect when clicked
	disabled bool
	// checked menu item has a tick before the title
	checked bool
	// has the menu item a checkbox (Linux)
	isCheckable bool
	isRadio     bool
	// parent item, for sub menus
	parent *MenuItem
}

func (item *MenuItem) Click(fn func()) {
	item.click = fn
}

func (item *MenuItem) String() string {
	if item.parent == nil {
		return fmt.Sprintf("MenuItem[%d, %q]", item.id, item.title)
	}
	return fmt.Sprintf("MenuItem[%d, parent %d, %q]", item.id, item.parent.id, item.title)
}

// newMenuItem returns a populated MenuItem object
func newMenuItem(title string, tooltip string, parent *MenuItem) *MenuItem {
	return &MenuItem{
		id:          atomic.AddUint32(&currentID, 1),
		title:       title,
		tooltip:     tooltip,
		shortcutKey: "",
		disabled:    false,
		checked:     false,
		isCheckable: false,
		parent:      parent,
	}
}

// 设置鼠标左键双击事件的时间间隔 默认500毫秒
func SetDClickTimeMinInterval(value int64) {
	dClickTimeMinInterval = value
}

// 设置托盘鼠标左键点击事件
func SetOnClick(fn func()) {
	setOnClick(fn)
}

// 设置托盘鼠标左键双击事件
func SetOnDClick(fn func()) {
	setOnDClick(fn)
}

// ResetMenu will remove all menu items
func ResetMenu() {
	resetMenu()
}

// AddMenuItem adds a menu item with the designated title and tooltip.
// It can be safely invoked from different goroutines.
// Created menu items are checkable on Windows and OSX by default. For Linux you have to use AddMenuItemCheckbox
func AddMenuItem(title string, tooltip string) *MenuItem {
	item := newMenuItem(title, tooltip, nil)
	item.update()
	return item
}

// AddMenuItemCheckbox adds a menu item with the designated title and tooltip and a checkbox for Linux.
// It can be safely invoked from different goroutines.
// On Windows and OSX this is the same as calling AddMenuItem
func AddMenuItemCheckbox(title string, tooltip string, checked bool) *MenuItem {
	item := newMenuItem(title, tooltip, nil)
	item.isCheckable = true
	item.checked = checked
	item.update()
	return item
}

// AddSeparator adds a separator bar to the menu
func AddSeparator() {
	addSeparator(atomic.AddUint32(&currentID, 1))
}

func (item *MenuItem) AddSeparator() {
	instance.menuLock.Lock()
	defer instance.menuLock.Unlock()
	m, exists := findLayout(int32(item.id))
	if exists {
		id := atomic.AddUint32(&currentID, 1)
		layout := &menuLayout{
			V0: int32(id),
			V1: map[string]dbus.Variant{
				"type": dbus.MakeVariant("separator"),
			},
			V2: []dbus.Variant{},
		}
		m.V2 = append(m.V2, dbus.MakeVariant(layout))
		refresh()
	}
}

// AddSubMenuItem adds a nested sub-menu item with the designated title and tooltip.
// It can be safely invoked from different goroutines.
// Created menu items are checkable on Windows and OSX by default. For Linux you have to use AddSubMenuItemCheckbox
func (item *MenuItem) AddSubMenuItem(title string, tooltip string) *MenuItem {
	child := newMenuItem(title, tooltip, item)
	child.update()
	return child
}

// AddSubMenuItemCheckbox adds a nested sub-menu item with the designated title and tooltip and a checkbox for Linux.
// It can be safely invoked from different goroutines.
// On Windows and OSX this is the same as calling AddSubMenuItem
func (item *MenuItem) AddSubMenuItemCheckbox(title string, tooltip string, checked bool) *MenuItem {
	child := newMenuItem(title, tooltip, item)
	child.isCheckable = true
	child.checked = checked
	child.update()
	return child
}

// SetTitle set the text to display on a menu item
func (item *MenuItem) SetTitle(title string) {
	item.title = title
	item.update()
}

// SetTooltip set the tooltip to show when mouse hover
func (item *MenuItem) SetTooltip(tooltip string) {
	item.tooltip = tooltip
	item.update()
}

// Disabled checks if the menu item is disabled
func (item *MenuItem) Disabled() bool {
	return item.disabled
}

// Enable a menu item regardless if it's previously enabled or not
func (item *MenuItem) Enable() {
	item.disabled = false
	item.update()
}

// Disable a menu item regardless if it's previously disabled or not
func (item *MenuItem) Disable() {
	item.disabled = true
	item.update()
}

// Clear menu item
func (item *MenuItem) Clear() {
	instance.menuLock.Lock()
	defer instance.menuLock.Unlock()
	m, exists := findLayout(int32(item.id))
	if exists {
		for _, sub := range m.V2 {
			menu := sub.Value().(*menuLayout)
			delete(menuItems, uint32(menu.V0))
		}
		m.V2 = []dbus.Variant{}
		m.V1["children-display"] = dbus.MakeVariant("menu")
		refresh()
	}
}

// Hide hides a menu item
func (item *MenuItem) Hide() {
	instance.menuLock.Lock()
	defer instance.menuLock.Unlock()
	m, exists := findLayout(int32(item.id))
	if exists {
		m.V1["visible"] = dbus.MakeVariant(false)
		refresh()
	}
}

// Show shows a previously hidden menu item
func (item *MenuItem) Show() {
	instance.menuLock.Lock()
	defer instance.menuLock.Unlock()
	m, exists := findLayout(int32(item.id))
	if exists {
		m.V1["visible"] = dbus.MakeVariant(true)
		refresh()
	}
}

// Checked returns if the menu item has a check mark
func (item *MenuItem) Checked() bool {
	return item.checked
}

// SetChecked a menu item regardless if it's previously checked or not
func (item *MenuItem) SetChecked(value bool) {
	item.checked = value
	item.isCheckable = true
	item.isRadio = false
	item.update()
}

// SetRadio a menu item regardless if it's previously radio or not
func (item *MenuItem) SetRadio(value bool) {
	item.checked = value
	item.isCheckable = false
	item.isRadio = true
	item.update()
}

// update propagates changes on a menu item to systray
func (item *MenuItem) update() {
	menuItemsLock.Lock()
	menuItems[item.id] = item
	menuItemsLock.Unlock()
	addOrUpdateMenuItem(item)
}

func systrayMenuItemSelected(id uint32) {
	menuItemsLock.RLock()
	item, ok := menuItems[id]
	menuItemsLock.RUnlock()
	if !ok {
		log.Printf("systray error: no menu item with ID %d\n", id)
		return
	}
	if item.click != nil {
		lcl.RunOnMainThreadAsync(func(id uint32) {
			item.click()
		})
	}
}
