//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/energy/v2/pkgs/win"
)

// NewRegistryAllAccess 所有访问权限
func NewRegistryAllAccess() IRegistry {
	return NewRegistry1(win.KEY_ALL_ACCESS)
}
