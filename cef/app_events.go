//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package cef

import (
	"github.com/energye/cef/cef"
	cefTypes "github.com/energye/cef/cef/types"
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/logger"
	"strings"
)

func (m *Application) initDefaultEvent() {
	logger.Debug("Application.initDefaultEvent")
	m.SetOnContextCreated(m.applicationOnContextCreated)
	m.SetOnProcessMessageReceived(m.applicationOnProcessMessageReceived)
	m.SetOnWebKitInitialized(m.applicationOnWebKitInitialized)
	m.SetOnRegCustomSchemes(m.applicationOnRegCustomSchemes)
}

func (m *Application) applicationOnContextCreated(browser cef.ICefBrowser, frame cef.ICefFrame, context cef.ICefv8Context) {
}

func (m *Application) applicationOnProcessMessageReceived(browser cef.ICefBrowser, frame cef.ICefFrame, sourceProcess cefTypes.TCefProcessId,
	message cef.ICefProcessMessage, handled *bool) {

}

func (m *Application) applicationOnWebKitInitialized() {

}

func (m *Application) applicationOnRegCustomSchemes(registrar cef.ICefSchemeRegistrarRef) {
	logger.Debug("Application.OnRegCustomSchemes")
	gApp := application.GApplication
	if gApp == nil || gApp.LocalLoad == nil {
		return
	}
	switch strings.ToUpper(gApp.LocalLoad.Scheme) {
	case "HTTP", "HTTPS", "FILE", "FTP", "ABOUT", "DATA":
		return
	}
	registrar.AddCustomScheme(gApp.LocalLoad.Scheme,
		cefTypes.CEF_SCHEME_OPTION_STANDARD|cefTypes.CEF_SCHEME_OPTION_CORS_ENABLED|cefTypes.CEF_SCHEME_OPTION_SECURE|cefTypes.CEF_SCHEME_OPTION_FETCH_ENABLED)
}
