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

// 本地加载资源
var localLoadRes *LocalLoadResource

// LocalLoadResource
type LocalLoadResource struct {
	*LocalLoad
}

// 初始化本地加载配置对象
func localLoadResourceInit(ll *LocalLoad) {
	if ll != nil {
		localLoadRes = &LocalLoadResource{
			LocalLoad: ll,
		}
		localLoadRes.LocalLoad.defaultInit()
	}
}
