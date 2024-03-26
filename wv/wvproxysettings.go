//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IWVProxySettings Parent: IObject
//
//	Class used by the TWVLoader.ProxySettigns property to configure
//	a custom proxy server using the following command line switches:
//	--no-proxy-server, --proxy-auto-detect, --proxy-bypass-list,
//	--proxy-pac-url and --proxy-server.
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-proxy-server</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-auto-detect</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-bypass-list</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-pac-url</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-server</a>
type IWVProxySettings interface {
	IObject
	NoProxyServer() bool          // property
	SetNoProxyServer(AValue bool) // property
	AutoDetect() bool             // property
	SetAutoDetect(AValue bool)    // property
	ByPassList() string           // property
	SetByPassList(AValue string)  // property
	PacUrl() string               // property
	SetPacUrl(AValue string)      // property
	Server() string               // property
	SetServer(AValue string)      // property
}

// TWVProxySettings Parent: TObject
//
//	Class used by the TWVLoader.ProxySettigns property to configure
//	a custom proxy server using the following command line switches:
//	--no-proxy-server, --proxy-auto-detect, --proxy-bypass-list,
//	--proxy-pac-url and --proxy-server.
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --no-proxy-server</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-auto-detect</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-bypass-list</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-pac-url</a>
//	<a href="https://peter.sh/experiments/chromium-command-line-switches/">Uses the following command line switch: --proxy-server</a>
type TWVProxySettings struct {
	TObject
}

func NewWVProxySettings() IWVProxySettings {
	r1 := WV().SysCallN(1121)
	return AsWVProxySettings(r1)
}

func (m *TWVProxySettings) NoProxyServer() bool {
	r1 := WV().SysCallN(1122, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVProxySettings) SetNoProxyServer(AValue bool) {
	WV().SysCallN(1122, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVProxySettings) AutoDetect() bool {
	r1 := WV().SysCallN(1118, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TWVProxySettings) SetAutoDetect(AValue bool) {
	WV().SysCallN(1118, 1, m.Instance(), PascalBool(AValue))
}

func (m *TWVProxySettings) ByPassList() string {
	r1 := WV().SysCallN(1119, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVProxySettings) SetByPassList(AValue string) {
	WV().SysCallN(1119, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVProxySettings) PacUrl() string {
	r1 := WV().SysCallN(1123, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVProxySettings) SetPacUrl(AValue string) {
	WV().SysCallN(1123, 1, m.Instance(), PascalStr(AValue))
}

func (m *TWVProxySettings) Server() string {
	r1 := WV().SysCallN(1124, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TWVProxySettings) SetServer(AValue string) {
	WV().SysCallN(1124, 1, m.Instance(), PascalStr(AValue))
}

func WVProxySettingsClass() TClass {
	ret := WV().SysCallN(1120)
	return TClass(ret)
}
