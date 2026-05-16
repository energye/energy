// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0

//go:build darwin && cgo

package notification

import (
	"github.com/energye/energy/v3/platform/darwin/cocoa/cgo/notification"
	. "github.com/energye/energy/v3/platform/notification/types"
)

func New() INotification {
	return notification.New()
}
