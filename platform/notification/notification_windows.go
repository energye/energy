//go:build windows

package notification

import (
	. "github.com/energye/energy/v3/platform/notification/types"
	win32notification "github.com/energye/energy/v3/platform/win32/notification"
)

func New() INotification {
	return win32notification.New()
}
