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

var application ICEFApplication

// ICEFApplication
//	CEF 应用内部保留接口
type ICEFApplication interface {
	SingleProcess() bool
}

func SetApplication(app ICEFApplication) {
	if application == nil {
		application = app
	}
}

// Application
//	返回 IApplication
func Application() ICEFApplication {
	return application
}
