// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0

//go:build linux

package notification

import (
	"github.com/energye/energy/v3/platform/linux/notification"
	. "github.com/energye/energy/v3/platform/notification/types"
)

func New() INotification {
	return notification.New()
}
