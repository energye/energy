//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package opengl

import (
	"github.com/cyber-xxm/energy/v2/common/imports"
	"github.com/energye/golcl/lcl"
	. "github.com/energye/golcl/lcl/api"
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

// TOpenGLControlOption ENUM
type TOpenGLControlOption = int32

const (
	OcoMacRetinaMode TOpenGLControlOption = iota
	OcoRenderAtDesignTime
)

// TOpenGLControlOptions SET: TOpenGLControlOption
type TOpenGLControlOptions = types.TSet

type TOpenGlCtrlMakeCurrentEvent func(sender lcl.IObject, allow *bool)

// TCustomOpenGLControl Parent: TWinControl
type TCustomOpenGLControl struct {
	*lcl.TWinControl
	makeCurrentPtr uintptr
	paintPtr       uintptr
}

func (m *TCustomOpenGLControl) SharingControls(Index int32) *TCustomOpenGLControl {
	r1 := customOpenGLDllTableAPI().SysCallN(0, m.Instance(), uintptr(Index))
	r := &TCustomOpenGLControl{TWinControl: lcl.AsWinControl(r1)}
	return r
}

func (m *TCustomOpenGLControl) FrameDiffTimeInMSecs() int32 {
	r1 := customOpenGLDllTableAPI().SysCallN(1, m.Instance())
	return int32(r1)
}

func (m *TCustomOpenGLControl) SharedControl() *TCustomOpenGLControl {
	r1 := customOpenGLDllTableAPI().SysCallN(2, 0, m.Instance(), 0)
	r := &TCustomOpenGLControl{TWinControl: lcl.AsWinControl(r1)}
	return r
}

func (m *TCustomOpenGLControl) SetSharedControl(AValue *TCustomOpenGLControl) {
	customOpenGLDllTableAPI().SysCallN(2, 1, m.Instance(), AValue.Instance())
}

func (m *TCustomOpenGLControl) AutoResizeViewport() bool {
	r1 := customOpenGLDllTableAPI().SysCallN(3, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetAutoResizeViewport(AValue bool) {
	customOpenGLDllTableAPI().SysCallN(3, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) DebugContext() bool {
	r1 := customOpenGLDllTableAPI().SysCallN(4, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetDebugContext(AValue bool) {
	customOpenGLDllTableAPI().SysCallN(4, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) RGBA() bool {
	r1 := customOpenGLDllTableAPI().SysCallN(5, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SetRGBA(AValue bool) {
	customOpenGLDllTableAPI().SysCallN(5, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCustomOpenGLControl) RedBits() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(6, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetRedBits(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(6, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) GreenBits() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(7, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetGreenBits(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(7, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) BlueBits() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(8, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetBlueBits(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(8, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) OpenGLMajorVersion() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(9, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetOpenGLMajorVersion(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(9, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) OpenGLMinorVersion() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(10, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetOpenGLMinorVersion(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(10, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) MultiSampling() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(11, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetMultiSampling(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(11, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) AlphaBits() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(12, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetAlphaBits(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(12, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) DepthBits() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(13, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetDepthBits(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(13, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) StencilBits() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(14, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetStencilBits(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(14, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) AUXBuffers() uint32 {
	r1 := customOpenGLDllTableAPI().SysCallN(15, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCustomOpenGLControl) SetAUXBuffers(AValue uint32) {
	customOpenGLDllTableAPI().SysCallN(15, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) Options() TOpenGLControlOptions {
	r1 := customOpenGLDllTableAPI().SysCallN(16, 0, m.Instance(), 0)
	return TOpenGLControlOptions(r1)
}

func (m *TCustomOpenGLControl) SetOptions(AValue TOpenGLControlOptions) {
	customOpenGLDllTableAPI().SysCallN(16, 1, m.Instance(), uintptr(AValue))
}

func (m *TCustomOpenGLControl) MakeCurrent(SaveOldToStack bool) bool {
	r1 := customOpenGLDllTableAPI().SysCallN(17, m.Instance(), PascalBool(SaveOldToStack))
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) ReleaseContext() bool {
	r1 := customOpenGLDllTableAPI().SysCallN(18, m.Instance())
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) RestoreOldOpenGLControl() bool {
	r1 := customOpenGLDllTableAPI().SysCallN(19, m.Instance())
	return GoBool(r1)
}

func (m *TCustomOpenGLControl) SharingControlCount() int32 {
	r1 := customOpenGLDllTableAPI().SysCallN(20, m.Instance())
	return int32(r1)
}

func (m *TCustomOpenGLControl) Paint() {
	customOpenGLDllTableAPI().SysCallN(21, m.Instance())
}

func (m *TCustomOpenGLControl) RealizeBounds() {
	customOpenGLDllTableAPI().SysCallN(22, m.Instance())
}

func (m *TCustomOpenGLControl) DoOnPaint() {
	customOpenGLDllTableAPI().SysCallN(23, m.Instance())
}

func (m *TCustomOpenGLControl) SwapBuffers() {
	customOpenGLDllTableAPI().SysCallN(24, m.Instance())
}

func (m *TCustomOpenGLControl) SetOnMakeCurrent(fn TOpenGlCtrlMakeCurrentEvent) {
	if m.makeCurrentPtr != 0 {
		RemoveEventElement(m.makeCurrentPtr)
	}
	m.makeCurrentPtr = MakeEventDataPtr(fn)
	customOpenGLDllTableAPI().SysCallN(25, m.Instance(), m.makeCurrentPtr)
}

func (m *TCustomOpenGLControl) SetOnPaint(fn lcl.TNotifyEvent) {
	if m.paintPtr != 0 {
		RemoveEventElement(m.paintPtr)
	}
	m.paintPtr = MakeEventDataPtr(fn)
	customOpenGLDllTableAPI().SysCallN(26, m.Instance(), m.paintPtr)
}

var (
	customOpenGLDllTable *imports.DllTable
)

func customOpenGLDllTableAPI() *imports.DllTable {
	if customOpenGLDllTable == nil {
		var importDefs = []*dllimports.ImportTable{
			dllimports.NewEnergyImport("CustomOpenGLControl_SharingControls", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_FrameDiffTimeInMSecs", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_SharedControl", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_AutoResizeViewport", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_DebugContext", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_RGBA", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_RedBits", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_GreenBits", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_BlueBits", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_OpenGLMajorVersion", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_OpenGLMinorVersion", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_MultiSampling", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_AlphaBits", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_DepthBits", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_StencilBits", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_AUXBuffers", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_Options", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_MakeCurrent", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_ReleaseContext", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_RestoreOldOpenGLControl", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_SharingControlCount", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_Paint", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_RealizeBounds", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_DoOnPaint", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_SwapBuffers", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_SetOnMakeCurrent", 0),
			dllimports.NewEnergyImport("CustomOpenGLControl_SetOnPaint", 0),
		}
		customOpenGLDllTable = new(imports.DllTable)
		customOpenGLDllTable.SetOk(true)
		customOpenGLDllTable.SetDll(imports.LibLCLExt().Dll())
		customOpenGLDllTable.SetImportTable(importDefs)
	}
	return customOpenGLDllTable
}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}
		switch fn.(type) {
		case TOpenGlCtrlMakeCurrentEvent:
			fn.(TOpenGlCtrlMakeCurrentEvent)(lcl.AsObject(getVal(0)), (*bool)(getPtr(1)))
		default:
			return false
		}
		return true
	})
}
