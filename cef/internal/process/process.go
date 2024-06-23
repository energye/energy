//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package process

var Current = &current{}

type current struct {
	browserId int32
	frameId   int64
}

func (m *current) BrowserId() int32 {
	return m.browserId
}

func (m *current) FrameId() int64 {
	return m.frameId
}

func (m *current) SetBrowserId(v int32) {
	m.browserId = v
}

func (m *current) SetFrameId(v int64) {
	m.frameId = v
}
