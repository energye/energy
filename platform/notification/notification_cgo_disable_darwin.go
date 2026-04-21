//go:build darwin && !cgo

package notification

import (
	"github.com/energye/energy/v3/platform/darwin/cocoa/nocgo/notification"
	. "github.com/energye/energy/v3/platform/notification/types"
)

func New() INotification {
	return notification.New()
}
