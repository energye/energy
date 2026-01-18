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

type Theme uintptr

const (
	SystemDefault Theme = iota // SystemDefault will use whatever the system theme is. The application will follow system theme changes.
	Dark                       // Dark Mode
	Light                      // Light Mode
)

type TBackdropType int32

const (
	BtAuto TBackdropType = iota
	BtNone
	BtMica
	BtAcrylic
	BtTabbed
)

type ThemeSetting struct {
	DarkTitleBar           int32
	DarkTitleBarInactive   int32
	DarkTitleText          int32
	DarkTitleTextInactive  int32
	DarkBorder             int32
	DarkBorderInactive     int32
	LightTitleBar          int32
	LightTitleBarInactive  int32
	LightTitleText         int32
	LightTitleTextInactive int32
	LightBorder            int32
	LightBorderInactive    int32
}
