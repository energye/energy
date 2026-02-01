//go:build linux

package systray

import (
	"github.com/godbus/dbus/v5"
)

func (m *Tray) applyItemToLayout(item *MenuItem) {
	item.layout.V1["enabled"] = dbus.MakeVariant(!item.disabled)
	item.layout.V1["label"] = dbus.MakeVariant(item.title)
	if item.isCheckable {
		item.layout.V1["toggle-type"] = dbus.MakeVariant("checkmark")
		if item.checked {
			item.layout.V1["toggle-state"] = dbus.MakeVariant(1)
		} else {
			item.layout.V1["toggle-state"] = dbus.MakeVariant(0)
		}
	} else if item.isRadio {
		item.layout.V1["toggle-type"] = dbus.MakeVariant("radio")
		if item.checked {
			item.layout.V1["toggle-state"] = dbus.MakeVariant(1)
		} else {
			item.layout.V1["toggle-state"] = dbus.MakeVariant(0)
		}
	} else {
		item.layout.V1["toggle-type"] = dbus.MakeVariant("")
		item.layout.V1["toggle-state"] = dbus.MakeVariant(0)
	}
	if item.parent != nil {
		partLayout := item.parent.layout
		partLayout.V1["children-display"] = dbus.MakeVariant("submenu")
	}
	m.menuItems[item.id] = item
}
