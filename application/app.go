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
//	本地加载资源，在浏览器发起xhr请求时的代理协议
//	http, https
type LocalProxyScheme int

const (
	LpsHttp  LocalProxyScheme = iota // http
	LpsHttps                         // https
	//LpsTcp                           // tcp
)

// SetOptions 设置应用程序的选项配置
func (m *Application) SetOptions(options Options) {
	m.Options = options
}

// SetLocalLoad 设置本地负载配置
func (m *Application) SetLocalLoad(localLoad LocalLoad) {
	m.LocalLoad = NewLocalLoadResource(&localLoad)
	m.LocalLoad.LocalLoad = &localLoad
}
