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

type IEmitTarget interface {
	GetBrowserId() int32
	GetFrameId() int64
}

// GoEmit相关事件的接收目标
type EmitTarget struct {
	BrowseId int32
	FrameId  int64
}

func NewEmitTarget(browserId int32, frameId int64) *EmitTarget {
	return &EmitTarget{
		BrowseId: browserId,
		FrameId:  frameId,
	}
}

func (m *EmitTarget) GetBrowserId() int32 {
	return m.BrowseId
}

func (m *EmitTarget) GetFrameId() int64 {
	return m.FrameId
}
