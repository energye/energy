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

import "github.com/energye/lcl/types"

// Options App config option
type Options struct {
	Caption                    string                `json:"-"`                          // window title
	DefaultURL                 string                `json:"-"`                          // load url in window
	X                          int32                 `json:"x"`                          // initial X position of the window.
	Y                          int32                 `json:"y"`                          // initial Y position of the window.
	Width                      int32                 `json:"width"`                      // initial width of the window
	Height                     int32                 `json:"height"`                     // initial height of the window
	MinWidth                   types.TConstraintSize `json:"-"`                          // min width of the window
	MinHeight                  types.TConstraintSize `json:"-"`                          // min height of the window
	MaxWidth                   types.TConstraintSize `json:"-"`                          // max width of the window
	MaxHeight                  types.TConstraintSize `json:"-"`                          // max height of the window
	DefaultWindowStatus        types.TWindowState    `json:"-"`                          // initial window state
	DisableDevTools            bool                  `json:"-"`                          //
	DisableContextMenu         bool                  `json:"-"`                          //
	DisableWebkitAppRegionDClk bool                  `json:"disableWebkitAppRegionDClk"` //
	DisableResize              bool                  `json:"disableResize"`              //
	DisableMinimize            bool                  `json:"-"`                          //
	DisableMaximize            bool                  `json:"-"`                          //
	DisableSystemMenu          bool                  `json:"-"`                          //
	Frameless                  bool                  `json:"frameless"`                  //
	Windows                    Windows               `json:"-"`                          //
	MacOS                      MacOS                 `json:"-"`                          //
	Linux                      Linux                 `json:"-"`                          //
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
