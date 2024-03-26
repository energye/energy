//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ITrayIcon Parent: ICustomTrayIcon
type ITrayIcon interface {
	ICustomTrayIcon
}

// TTrayIcon Parent: TCustomTrayIcon
type TTrayIcon struct {
	TCustomTrayIcon
}

func NewTrayIcon(TheOwner IComponent) ITrayIcon {
	r1 := LCL().SysCallN(4913, GetObjectUintptr(TheOwner))
	return AsTrayIcon(r1)
}

func TrayIconClass() TClass {
	ret := LCL().SysCallN(4912)
	return TClass(ret)
}
