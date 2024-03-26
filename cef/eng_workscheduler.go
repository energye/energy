//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

import (
	. "github.com/energye/energy/v2/api"
)

// ICEFWorkScheduler Parent: IComponent
//
//	Implementation of an external message pump for VCL and LCL.
//	Read the GlobalCEFApp.OnScheduleMessagePumpWork documentation for all the details.
type ICEFWorkScheduler interface {
	IComponent
	// Priority
	//  Priority of TCEFWorkSchedulerThread in Windows.
	Priority() TThreadPriority // property
	// SetPriority Set Priority
	SetPriority(AValue TThreadPriority) // property
	// DefaultInterval
	//  Default interval in milliseconds to do the next GlobalCEFApp.DoMessageLoopWork call.
	DefaultInterval() int32 // property
	// SetDefaultInterval Set DefaultInterval
	SetDefaultInterval(AValue int32) // property
	// DepleteWorkCycles
	//  Number of cycles used to deplete the remaining messages in the work loop.
	DepleteWorkCycles() uint32 // property
	// SetDepleteWorkCycles Set DepleteWorkCycles
	SetDepleteWorkCycles(AValue uint32) // property
	// DepleteWorkDelay
	//  Delay in milliseconds between the cycles used to deplete the remaining messages in the work loop.
	DepleteWorkDelay() uint32 // property
	// SetDepleteWorkDelay Set DepleteWorkDelay
	SetDepleteWorkDelay(AValue uint32) // property
	// UseQueueThread
	//  Use a custom queue thread instead of Windows messages or any other way to schedule the next pump work.
	UseQueueThread() bool // property
	// SetUseQueueThread Set UseQueueThread
	SetUseQueueThread(AValue bool) // property
	// ScheduleMessagePumpWork
	//  TCEFWorkScheduler destructor.
	//  Called from GlobalCEFApp.OnScheduleMessagePumpWork to schedule
	//  a GlobalCEFApp.DoMessageLoopWork call asynchronously to perform a single
	//  iteration of CEF message loop processing.
	//  <param name="delay_ms">Requested delay in milliseconds.</param>
	ScheduleMessagePumpWork(delayms int64) // procedure
	// StopScheduler
	//  Stop the scheduler. This function must be called after the destruction of all the forms in the application.
	StopScheduler() // procedure
	// CreateThread
	//  Creates all the internal threads used by TCEFWorkScheduler.
	CreateThread() // procedure
}

// TCEFWorkScheduler Parent: TComponent
//
//	Implementation of an external message pump for VCL and LCL.
//	Read the GlobalCEFApp.OnScheduleMessagePumpWork documentation for all the details.
type TCEFWorkScheduler struct {
	TComponent
}

func NewCEFWorkScheduler(aOwner IComponent) ICEFWorkScheduler {
	r1 := CEF().SysCallN(406, GetObjectUintptr(aOwner))
	return AsCEFWorkScheduler(r1)
}

func NewCEFWorkSchedulerDelayed() ICEFWorkScheduler {
	r1 := CEF().SysCallN(407)
	return AsCEFWorkScheduler(r1)
}

func (m *TCEFWorkScheduler) Priority() TThreadPriority {
	r1 := CEF().SysCallN(412, 0, m.Instance(), 0)
	return TThreadPriority(r1)
}

func (m *TCEFWorkScheduler) SetPriority(AValue TThreadPriority) {
	CEF().SysCallN(412, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWorkScheduler) DefaultInterval() int32 {
	r1 := CEF().SysCallN(409, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TCEFWorkScheduler) SetDefaultInterval(AValue int32) {
	CEF().SysCallN(409, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWorkScheduler) DepleteWorkCycles() uint32 {
	r1 := CEF().SysCallN(410, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFWorkScheduler) SetDepleteWorkCycles(AValue uint32) {
	CEF().SysCallN(410, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWorkScheduler) DepleteWorkDelay() uint32 {
	r1 := CEF().SysCallN(411, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TCEFWorkScheduler) SetDepleteWorkDelay(AValue uint32) {
	CEF().SysCallN(411, 1, m.Instance(), uintptr(AValue))
}

func (m *TCEFWorkScheduler) UseQueueThread() bool {
	r1 := CEF().SysCallN(415, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCEFWorkScheduler) SetUseQueueThread(AValue bool) {
	CEF().SysCallN(415, 1, m.Instance(), PascalBool(AValue))
}

func CEFWorkSchedulerClass() TClass {
	ret := CEF().SysCallN(405)
	return TClass(ret)
}

func (m *TCEFWorkScheduler) ScheduleMessagePumpWork(delayms int64) {
	CEF().SysCallN(413, m.Instance(), uintptr(unsafePointer(&delayms)))
}

func (m *TCEFWorkScheduler) StopScheduler() {
	CEF().SysCallN(414, m.Instance())
}

func (m *TCEFWorkScheduler) CreateThread() {
	CEF().SysCallN(408, m.Instance())
}
