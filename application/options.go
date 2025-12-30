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

// Options App config option
type Options struct {
	Caption                    string  `json:"-"`                          // window title
	DefaultURL                 string  `json:"-"`                          // load url in window
	DisableDevTools            bool    `json:"-"`                          //
	DisableContextMenu         bool    `json:"-"`                          //
	DisableWebkitAppRegionDClk bool    `json:"disableWebkitAppRegionDClk"` //
	DisableResize              bool    `json:"disableResize"`              //
	DisableMinimize            bool    `json:"-"`                          //
	DisableMaximize            bool    `json:"-"`                          //
	DisableSystemMenu          bool    `json:"-"`                          //
	Frameless                  bool    `json:"frameless"`                  //
	Windows                    Windows `json:"-"`                          //
	MacOS                      MacOS   `json:"-"`                          //
	Linux                      Linux   `json:"-"`                          //
}

type Theme uintptr

const (
	SystemDefault Theme = iota // SystemDefault will use whatever the system theme is. The application will follow system theme changes.
	Dark                       // Dark Mode
	Light                      // Light Mode
)

type Windows struct {
	ICON []byte `json:"-"` // window icon
	// Theme (Dark / Light / SystemDefault)
	// Default: SystemDefault - The application will follow system theme changes.
	Theme Theme
}

type MacOS struct {
}

type Linux struct {
	ICON []byte `json:"-"` // window icon
}
