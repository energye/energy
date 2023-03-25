//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// IPC-事件 到 指定浏览器
package ipc

// ITarget 指定目标
//
// ipc.NewTarget() *Target
type ITarget interface {
	BrowserId() int32
	FrameId() int64
}

// Target Go IPC 事件的接收目标
type Target struct {
	browseId int32
	frameId  int64
}

// NewTarget 创建一个新的Emit目标
func NewTarget(browserId int32, frameId int64) ITarget {
	return &Target{
		browseId: browserId,
		frameId:  frameId,
	}
}

// BrowserId 返回BrowserId
func (m *Target) BrowserId() int32 {
	return m.browseId
}

// FrameId 返回FrameId
func (m *Target) FrameId() int64 {
	return m.frameId
}
