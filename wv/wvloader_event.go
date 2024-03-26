//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	"github.com/energye/energy/v2/api"
)

func (m *TWVLoader) SetOnEnvironmentCreated(fn TOnLoaderNotifyEvent) {
	api.WVPreDef().SysCallN(0, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnInitializationError(fn TOnLoaderNotifyEvent) {
	api.WVPreDef().SysCallN(1, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnGetCustomSchemes(fn TOnLoaderGetCustomSchemesEvent) {
	api.WVPreDef().SysCallN(2, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnNewBrowserVersionAvailable(fn TOnLoaderNewBrowserVersionAvailableEvent) {
	api.WVPreDef().SysCallN(3, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnBrowserProcessExited(fn TOnLoaderBrowserProcessExitedEvent) {
	api.WVPreDef().SysCallN(4, api.MakeEventDataPtr(fn))
}

func (m *TWVLoader) SetOnProcessInfosChanged(fn TOnLoaderProcessInfosChangedEvent) {
	api.WVPreDef().SysCallN(5, api.MakeEventDataPtr(fn))
}
