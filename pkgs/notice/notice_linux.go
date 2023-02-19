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
// +build linux

package notice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/godbus/dbus/v5"
)

var once sync.Once
var scriptNum int

func SendNotification(n *Notification) {
	conn, err := dbus.SessionBus() // shared connection, don't close
	if err != nil {
		println("Failed to send message to bus", err)
		return
	}
	var appIcon string
	var appName = UniqueID()
	if n.Icon != nil && len(n.Icon) > 0 {
		bh := md5.Sum(n.Icon)
		dataHash := hex.EncodeToString(bh[:])
		fileName := fmt.Sprintf("notice-%s-%s", dataHash, n.iconExt)
		appIcon = filepath.Join(os.TempDir(), fileName)
		if _, err := os.Stat(appIcon); os.IsNotExist(err) {
			err = ioutil.WriteFile(appIcon, n.Icon, 0600)
			if err != nil {
				return
			}
		}
	}

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, appName, uint32(0), appIcon, n.Title, n.Content, []string{}, map[string]dbus.Variant{}, n.Timeout)
	if call.Err != nil {
		println("Failed to send message to bus", call.Err)
	}
}
