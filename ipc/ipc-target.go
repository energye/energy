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
	GetBrowserId() int32
	GetFrameId() int64
}

// Target Go IPC 事件的接收目标
type Target struct {
	BrowseId int32
	FrameId  int64
}

// NewTarget 创建一个新的Emit目标
func NewTarget(browserId int32, frameId int64) *Target {
	return &Target{
		BrowseId: browserId,
		FrameId:  frameId,
	}
}

// GetBrowserId 返回BrowserId
func (m *Target) GetBrowserId() int32 {
	return m.BrowseId
}

// GetFrameId 返回FrameId
func (m *Target) GetFrameId() int64 {
	return m.FrameId
}
