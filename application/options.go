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

import (
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
)

type THardwareGPU int32

const (
	HGPUDefault THardwareGPU = iota
	HGPUEnable
	HGPUDisable
)

// Options App config option
type Options struct {
	Caption                    string             `json:"-"`                          // window title
	DefaultURL                 string             `json:"-"`                          // load url in window
	X                          int32              `json:"x"`                          // initial X position of the window.
	Y                          int32              `json:"y"`                          // initial Y position of the window.
	Width                      int32              `json:"width"`                      // initial width of the window
	Height                     int32              `json:"height"`                     // initial height of the window
	MinWidth                   int32              `json:"-"`                          // min width of the window
	MinHeight                  int32              `json:"-"`                          // min height of the window
	MaxWidth                   int32              `json:"-"`                          // max width of the window
	MaxHeight                  int32              `json:"-"`                          // max height of the window
	DefaultWindowStatus        types.TWindowState `json:"-"`                          // initial window state
	DisableDevTools            bool               `json:"disableDevTools"`            // macOS: --tags dev
	DisableContextMenu         bool               `json:"disableContextMenu"`         //
	DisableWebkitAppRegionDClk bool               `json:"disableWebkitAppRegionDClk"` //
	DisableResize              bool               `json:"disableResize"`              //
	DisableMinimize            bool               `json:"-"`                          //
	DisableMaximize            bool               `json:"-"`                          //
	DisableSystemMenu          bool               `json:"-"`                          //
	Frameless                  bool               `json:"frameless"`                  //
	WindowTransparent          bool               `json:"-"`                          //
	WebviewTransparent         bool               `json:"-"`                          //
	BackgroundColor            *colors.TARGB      `json:"-"`                          //
	Windows                    Windows            `json:"-"`                          //
	MacOS                      MacOS              `json:"-"`                          //
	Linux                      Linux              `json:"-"`                          //
}
