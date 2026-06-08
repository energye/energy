//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build CEF127

package cef

import (
	"github.com/energye/cef/127/cef"
	"github.com/energye/cef/127/types"
)

type ICefApplication interface {
	cef.ICefApplication
}

type ICEFWorkScheduler interface {
	cef.ICEFWorkScheduler
}

func NewCEFApplication() ICefApplication {
	return cef.NewApplication()
}

func (m *Application) IsMainProcess() bool {
	return m.ProcessType() == types.PtBrowser
}

func NewWorkScheduler(owner lcl.IComponent) ICEFWorkScheduler {
	return cef.NewWorkScheduler(owner)
}
