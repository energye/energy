//go:build linux
// +build linux

package app

import (
	"sync"

	"github.com/godbus/dbus/v5"
)

var once sync.Once

func SendNotification(n *Notification) {
	conn, err := dbus.SessionBus() // shared connection, don't close
	if err != nil {
		println("Failed to send message to bus", err)
		return
	}

	appName := UniqueID()
	appIcon := a.cachedIconPath()
	timeout := int32(0)

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, appName, uint32(0), appIcon, n.Title, n.Content, []string{}, map[string]dbus.Variant{}, timeout)
	if call.Err != nil {
		println("Failed to send message to bus", call.Err)
	}
}
