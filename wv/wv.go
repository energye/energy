//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package wv

// LocalCustomerScheme 本地资源加载自定义固定协议
//
//	file, fs
type LocalCustomerScheme string

const (
	LcsLocal LocalCustomerScheme = "local" // 本地目录 local://energy/index.html
	LcsFS    LocalCustomerScheme = "fs"    // 内置 fs://energy/index.html
)

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
