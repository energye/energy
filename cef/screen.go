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

import (
	"github.com/cyber-xxm/energy/v2/common"
	"github.com/energye/golcl/lcl"
)

// IScreen
//
//	Two ways to obtain screen information
//	They are: 1. Using LCL, 2 Using CEF
type IScreen interface {
	Primary() (info Display)      // Returns the primary Display.
	Count() int                   // Returns display count
	Get(index int) (info Display) // Returns Obtain brief information about the specified screen
	LCLScreen() *lcl.TScreen      // Use this function when using the LCL window, If using ViewFramework window under Windows, this function can also be used
	VFScreen() *display           // Use this function when using the CEF ViewFramework
}

// Display
//
//	Screen brief information
type Display struct {
	ID                int
	WorkArea          TCefRect
	Bounds            TCefRect
	Rotation          int32   // Currently, LCL has not been obtained
	DeviceScaleFactor float32 // Currently, LCL has not been obtained
}

type Screen struct {
	window    IBrowserWindow
	lclScreen *lcl.TScreen
}

// LCLScreen
//
//	Use this function when using the LCL window
//	If using VF windows under Windows, this function can also be used
func (m *Screen) LCLScreen() *lcl.TScreen {
	if m.window == nil {
		return nil
	}
	if m.window.IsLCL() {
		if m.lclScreen == nil {
			m.lclScreen = lcl.NewScreen(m.window.AsLCLBrowserWindow().BrowserWindow())
		}
	} else if m.window.IsViewsFramework() && common.IsWindows() {
		// If using VF windows under Windows, this function can also be used
		if m.lclScreen == nil {
			m.lclScreen = lcl.NewScreen(m.window.AsViewsFrameworkBrowserWindow().WindowComponent())
		}
	}
	return m.lclScreen
}

// VFScreen
//
//	Use this function when using the CEF ViewFramework
func (m *Screen) VFScreen() *display {
	if m.window == nil {
		return nil
	}
	if m.window.IsViewsFramework() {
		return &DisplayRef
	}
	return nil
}

// Count
//
//	Returns all display count
func (m *Screen) Count() int {
	if m.window == nil {
		return 0
	}
	if m.window.IsLCL() {
		return int(m.LCLScreen().MonitorCount())
	} else if m.window.IsViewsFramework() {
		return int(m.VFScreen().GetCount())
	}
	return 0
}

// Primary
//
//	Returns the primary Display.
func (m *Screen) Primary() (info Display) {
	if m.window.IsLCL() {
		screen := m.LCLScreen()
		monitor := screen.PrimaryMonitor()
		if monitor == nil {
			return
		}
		wkr := monitor.WorkareaRect()
		bdr := monitor.BoundsRect()
		info.ID = int(monitor.MonitorNum())
		info.WorkArea = TCefRect{X: wkr.Left, Y: wkr.Top, Width: wkr.Width(), Height: wkr.Height()}
		info.Bounds = TCefRect{X: bdr.Top, Y: bdr.Left, Width: bdr.Width(), Height: bdr.Height()}
		screen.PrimaryMonitor()
	} else if m.window.IsViewsFramework() {
		screen := m.VFScreen()
		monitor := screen.Primary()
		if monitor == nil {
			return
		}
		wkr := monitor.WorkArea()
		bdr := monitor.Bounds()
		info.ID = int(monitor.ID())
		info.Rotation = monitor.Rotation()
		info.DeviceScaleFactor = monitor.DeviceScaleFactor()
		info.WorkArea = TCefRect{X: wkr.X, Y: wkr.Y, Width: wkr.Width, Height: wkr.Height}
		info.Bounds = TCefRect{X: bdr.X, Y: bdr.Y, Width: bdr.Width, Height: bdr.Height}
	}
	return
}

func (m *Screen) Get(index int) (info Display) {
	if m.window.IsLCL() {
		screen := m.LCLScreen()
		monitor := screen.Monitors(int32(index))
		if monitor == nil {
			return
		}
		wkr := monitor.WorkareaRect()
		bdr := monitor.BoundsRect()
		info.ID = int(monitor.MonitorNum())
		info.WorkArea = TCefRect{X: wkr.Left, Y: wkr.Top, Width: wkr.Width(), Height: wkr.Height()}
		info.Bounds = TCefRect{X: bdr.Top, Y: bdr.Left, Width: bdr.Width(), Height: bdr.Height()}
	} else if m.window.IsViewsFramework() {
		screen := m.VFScreen()
		alls := screen.GetAlls()
		monitor := alls.Get(uint32(index))
		if monitor == nil {
			return
		}
		wkr := monitor.WorkArea()
		bdr := monitor.Bounds()
		info.ID = int(monitor.ID())
		info.Rotation = monitor.Rotation()
		info.DeviceScaleFactor = monitor.DeviceScaleFactor()
		info.WorkArea = TCefRect{X: wkr.X, Y: wkr.Y, Width: wkr.Width, Height: wkr.Height}
		info.Bounds = TCefRect{X: bdr.X, Y: bdr.Y, Width: bdr.Width, Height: bdr.Height}
	}
	return
}
