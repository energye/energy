//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// IPC-事件 Emit 到 Target
package cef

// IEmitTarget 触发事件目标接口
type IEmitTarget interface {
	GetBrowserId() int32
	GetFrameId() int64
}

// EmitTarget GoEmit相关事件的接收目标
type EmitTarget struct {
	BrowseId int32
	FrameId  int64
}

// NewEmitTarget 创建一个新的Emit目标
func NewEmitTarget(browserId int32, frameId int64) *EmitTarget {
	return &EmitTarget{
		BrowseId: browserId,
		FrameId:  frameId,
	}
}

// GetBrowserId 返回BrowserId
func (m *EmitTarget) GetBrowserId() int32 {
	return m.BrowseId
}

// GetFrameId 返回FrameId
func (m *EmitTarget) GetFrameId() int64 {
	return m.FrameId
}
