//go:build linux

package systray

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"sync/atomic"
)

var (
	dClickTimeMinInterval int64 = 500
)

// 设置鼠标左键双击事件的时间间隔 默认500毫秒
func SetDClickTimeMinInterval(value int64) {
	dClickTimeMinInterval = value
}

// MenuItem is used to keep track each menu item of systray.
// Don't create it directly, use the one systray.AddMenuItem() returned
type MenuItem struct {
	// ClickedCh is the channel which will be notified when the menu item is clicked
	click func()
	// id uniquely identify a menu item, not supposed to be modified
	id int32
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
	// child
	child []*MenuItem
	// layout
	layout *menuLayout
}

func (m *MenuItem) Click(fn func()) {
	m.click = fn
}

func (m *MenuItem) String() string {
	if m.parent == nil {
		return fmt.Sprintf("MenuItem[%d, %q]", m.id, m.title)
	}
	return fmt.Sprintf("MenuItem[%d, parent %d, %q]", m.id, m.parent.id, m.title)
}

// newMenuItem returns a populated MenuItem object
func newMenuItem(title string, tooltip string, parent *MenuItem) *MenuItem {
	m := &MenuItem{
		id:          atomic.AddInt32(&currentID, 1),
		title:       title,
		tooltip:     tooltip,
		shortcutKey: "",
		disabled:    false,
		checked:     false,
		isCheckable: false,
		parent:      parent,
	}
	if parent != nil {
		parent.child = append(parent.child, m)
	}
	return m
}

func (m *MenuItem) AddSeparator(tray *Tray) {
	id := atomic.AddInt32(&currentID, 1)
	item := newMenuItem("-", "-", m)
	item.layout = &menuLayout{
		V0: id,
		V1: map[string]dbus.Variant{
			"type": dbus.MakeVariant("separator"),
		},
		V2: []dbus.Variant{},
	}
	m.layout.V2 = append(m.layout.V2, dbus.MakeVariant(item.layout))
	tray.Refresh()
}

// AddMenuItem adds a nested sub-menu item with the designated title and tooltip.
// It can be safely invoked from different goroutines.
// Created menu items are checkable on Windows and OSX by default. For Linux you have to use AddSubMenuItemCheckbox
func (m *MenuItem) AddMenuItem(tray *Tray, title string, tooltip string) *MenuItem {
	item := newMenuItem(title, tooltip, m)
	item.layout = &menuLayout{
		V0: item.id,
		V1: map[string]dbus.Variant{},
		V2: []dbus.Variant{},
	}
	tray.applyItemToLayout(item)
	m.layout.V2 = append(m.layout.V2, dbus.MakeVariant(item.layout))
	tray.Refresh()
	return item
}

// SetTitle set the text to display on a menu item
func (m *MenuItem) SetTitle(tray *Tray, title string) {
	m.title = title
	tray.applyItemToLayout(m)
	tray.Refresh()
}

// SetTooltip set the tooltip to show when mouse hover
func (m *MenuItem) SetTooltip(tray *Tray, tooltip string) {
	m.tooltip = tooltip
	tray.applyItemToLayout(m)
	tray.Refresh()
}

// Enable checks if the menu item is disabled
func (m *MenuItem) Enable() bool {
	return m.disabled
}

// SetEnable a menu item regardless if it's previously enabled or not
func (m *MenuItem) SetEnable(tray *Tray, value bool) {
	m.disabled = !value
	tray.applyItemToLayout(m)
	tray.Refresh()
}

// Checked returns if the menu item has a check mark
func (m *MenuItem) Checked() bool {
	return m.checked
}

// SetChecked a menu item regardless if it's previously checked or not
func (m *MenuItem) SetChecked(tray *Tray, value bool) {
	m.checked = value
	m.isCheckable = true
	m.isRadio = false
	tray.applyItemToLayout(m)
	tray.Refresh()
}

// SetRadio a menu item regardless if it's previously radio or not
func (m *MenuItem) SetRadio(tray *Tray, value bool) {
	m.checked = value
	m.isCheckable = false
	m.isRadio = true
	tray.applyItemToLayout(m)
	tray.Refresh()
}

// Clear menu item
func (m *MenuItem) Clear(tray *Tray) {
	var removeChild func(item *MenuItem)
	removeChild = func(item *MenuItem) {
		for _, child := range item.child {
			delete(tray.menuItems, child.id)
			removeChild(child)
		}
	}
	removeChild(m)
	m.child = []*MenuItem{}
	m.layout.V2 = []dbus.Variant{}
	m.layout.V1["children-display"] = dbus.MakeVariant("menu")
	tray.Refresh()
}

// Hide hides a menu item
func (m *MenuItem) Hide(tray *Tray) {
	m.layout.V1["visible"] = dbus.MakeVariant(false)
	tray.Refresh()
}

// Show shows a previously hidden menu item
func (m *MenuItem) Show(tray *Tray) {
	m.layout.V1["visible"] = dbus.MakeVariant(true)
	tray.Refresh()
}

// SetIcon sets the icon of a menu item.
// iconBytes should be the content of .ico/.jpg/.png
func (m *MenuItem) SetIcon(tray *Tray, iconBytes []byte) {
	m.layout.V1["icon-data"] = dbus.MakeVariant(iconBytes)
	tray.Refresh()
}
