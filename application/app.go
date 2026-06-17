//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package application

var GApplication *Application

type Application struct {
	Options   Options
	LocalLoad *LocalLoadResource
}

// LocalProxyScheme
//
// SetupXHRProxy configures proxy protocol(http, https) for XHR requests
// Supports loading local resources when webview sends XMLHttpRequest
type LocalProxyScheme int

const (
	LpsHttp  LocalProxyScheme = iota // http
	LpsHttps                         // https
	//LpsTcp                           // tcp
)

// SetOptions sets configuration options for the application
func (m *Application) SetOptions(options Options) {
	m.Options = options
}

// SetLocalLoad sets configuration for local resource loading
func (m *Application) SetLocalLoad(localLoad LocalLoad) {
	m.LocalLoad = NewLocalLoadResource(&localLoad)
	m.LocalLoad.LocalLoad = &localLoad
}
