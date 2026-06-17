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
	"github.com/energye/energy/v3/application"
	"github.com/energye/energy/v3/logger"
)

type tSchemeHandlerFactory struct {
	factory cef.IEngSchemeHandlerFactory
}

func createSchemeHandlerFactory(browser cef.ICefBrowser) *tSchemeHandlerFactory {
	if application.GApplication == nil || application.GApplication.LocalLoad == nil {
		return nil
	}
	localLoad := application.GApplication.LocalLoad
	m := &tSchemeHandlerFactory{}

	m.factory = cef.NewEngSchemeHandlerFactory(0)
	m.factory.SetOnSchemeFactoryNew(m.onSchemeFactoryNew)
	logger.Debug("Chromium.OnAfterCreated > createSchemeHandlerFactory Scheme:", localLoad.Scheme, "Domain:", localLoad.Domain, "factory-IsValid:", m.factory.IsValid())
	intf := cef.AsEngSchemeHandlerFactory(m.factory.AsIntfSchemeHandlerFactory())
	ok := browser.GetHost().GetRequestContext().RegisterSchemeHandlerFactory(localLoad.Scheme, localLoad.Domain, intf)
	logger.Debug("Chromium.OnAfterCreated > createSchemeHandlerFactory RegisterSchemeHandlerFactory:", ok)
	return m
}

func (m *tSchemeHandlerFactory) onSchemeFactoryNew(browser cef.ICefBrowser, frame cef.ICefFrame, schemeName string, request cef.ICefRequest) cef.IEngResourceHandler {
	logger.Debug("SchemeHandlerFactory.OnNew")

	return nil
}
