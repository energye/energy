// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0

//go:build windows

package notification

import (
	. "github.com/energye/energy/v3/platform/notification/types"
	win32notification "github.com/energye/energy/v3/platform/win32/notification"
)

func New() INotification {
	return win32notification.New()
}
