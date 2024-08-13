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
	Caption                    string             `json:"-"`
	DefaultURL                 string             `json:"-"`
	ICON                       []byte             `json:"-"`
	LocalLoad                  *LocalLoad         `json:"-"`
	X                          int32              `json:"x"`
	Y                          int32              `json:"y"`
	Width                      int32              `json:"width"`
	Height                     int32              `json:"height"`
	DefaultWindowStatus        types.TWindowState `json:"-"`
	DisableDevTools            bool               `json:"-"`
	DisableContextMenu         bool               `json:"-"`
	DisableWebkitAppRegionDClk bool               `json:"disableWebkitAppRegionDClk"`
	DisableResize              bool               `json:"disableResize"`
	DisableMinimize            bool               `json:"-"`
	DisableMaximize            bool               `json:"-"`
	DisableSystemMenu          bool               `json:"-"`
	Frameless                  bool               `json:"frameless"`
}
