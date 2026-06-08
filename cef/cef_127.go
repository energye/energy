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
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
)

type ICEFApplication interface {
	cef.ICefApplication
}

type ICEFWorkScheduler interface {
	cef.ICEFWorkScheduler
}

type ICEFWindowParent interface {
	cef.ICEFWinControl
}

type ICEFChromium interface {
	cef.IChromium
}

func (m *Application) IsMainProcess() bool {
	return m.ProcessType() == types.PtBrowser
}

func NewCEFApplication() ICEFApplication {
	return cef.NewApplication()
}

func NewCEFWorkScheduler(owner lcl.IComponent) ICEFWorkScheduler {
	return cef.NewWorkScheduler(owner)
}

func NewCEFChromium(owner lcl.IComponent) ICEFChromium {
	return cef.NewChromium(owner)
}

func NewCEFWindowParent(chromium ICEFChromium, value lcl.IWinControl) ICEFWindowParent {
	if tool.IsWindows() {
		return cef.NewWindowParent(value)
	} else {
		windowParent := cef.NewLinkedWindowParent(value)
		windowParent.SetChromium(chromium)
		return windowParent
	}
}
