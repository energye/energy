//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build darwin

package wv

import (
	"github.com/energye/energy/v3/application"
	"github.com/energye/lcl/lcl"
	wv "github.com/energye/wv/darwin"
)

var (
	gApplication *Application
)

// Init WkWebkit Global initialization, invoked at application startup in main
func Init() *Application {
	lcl.Init()
	wv.Init()
	return NewApplication()
}

type Application struct {
	application.Application
	onCustomSchemes TApplicationOnCustomSchemesEvent
}

// NewApplication creates and returns an Application instance
// If the global Application instance is not initialized yet, it will perform initialization setup
func NewApplication() *Application {
	if gApplication == nil {
		gApplication = &Application{}
		application.GApplication = &gApplication.Application
	}
	return gApplication
}

func DestroyGlobalLoader() {
}

// Start starts the application
func (m *Application) Start() bool {
	return true
}
