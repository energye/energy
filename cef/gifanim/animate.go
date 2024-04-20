//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// GIF play animate

package gifanim

import (
	"bytes"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"image/gif"
	"os"
)

type TGIFAnimate struct {
	*lcl.TImage
	task           *lcl.TTimer
	filePath       string
	cache          bool
	gif            *gif.GIF
	w, h           int32
	frames         []*Frame
	index          int
	count          int
	onStart        func()
	onStop         func()
	onFrameChanged func(frame *Frame)
}

func NewGIFAnimate(owner lcl.IComponent) *TGIFAnimate {
	m := new(TGIFAnimate)
	m.TImage = lcl.NewImage(owner)
	m.task = lcl.NewTimer(owner)
	m.task.SetEnabled(false)
	m.task.SetInterval(100)
	m.task.SetOnTimer(m.onTimer)
	return m
}

func (m *TGIFAnimate) onTimer(sender lcl.IObject) {
	if m.count == 0 {
		return
	}
	frame := m.frames[m.index]
	m.doFrameChanged(frame)
	if m.index >= m.count-1 {
		m.index = 0
	} else {
		m.index++
	}
	nextDelay := m.nextDelay()
	if nextDelay != frame.delay {
		m.task.SetInterval(nextDelay)
	}
	m.scan(frame)
}

func (m *TGIFAnimate) nextDelay() uint32 {
	return m.frames[m.index].delay
}

func (m *TGIFAnimate) first() {
	m.w = int32(m.gif.Config.Width)
	m.h = int32(m.gif.Config.Height)
	m.SetWidth(m.w)
	m.SetHeight(m.h)
	m.Canvas().FillRect(types.Rect(0, 0, m.w, m.h))
	m.scan(m.frames[0])
}

func (m *TGIFAnimate) scan(frame *Frame) {
	frame.scan()
	m.Canvas().Draw(frame.x, frame.y, frame.image)
	if !m.cache {
		frame.reset()
	}
}

func (m *TGIFAnimate) reset() {
	for _, frame := range m.frames {
		frame.free()
	}
	m.w, m.h = 0, 0
	m.frames = nil
	m.index = 0
	m.count = 0
}

func (m *TGIFAnimate) load() {
	m.reset()
	m.count = len(m.gif.Image)
	m.frames = make([]*Frame, m.count)
	for i, img := range m.gif.Image {
		bounds := img.Bounds()
		var buf = new(bytes.Buffer)
		err := gif.Encode(buf, img, nil)
		if err != nil {
			panic(err)
		}
		m.frames[i] = &Frame{
			index: i,
			x:     int32(bounds.Min.X),
			y:     int32(bounds.Min.Y),
			w:     int32(bounds.Dx()),
			h:     int32(bounds.Dy()),
			delay: uint32(m.gif.Delay[i] * 10),
			data:  buf.Bytes(),
		}
		buf.Reset()
	}
	m.first()
	m.gif = nil
}

func (m *TGIFAnimate) SetAnimate(v bool) {
	m.task.SetEnabled(v)
	if v {
		m.doStart()
	} else {
		m.doStop()
	}
}

func (m *TGIFAnimate) CurrentFrameIndex() int {
	return m.index
}

func (m *TGIFAnimate) FrameCount() int {
	return m.count
}

func (m *TGIFAnimate) Frame(index int) *Frame {
	if index < m.count {
		return m.frames[index]
	}
	return nil
}

func (m *TGIFAnimate) Stop() {
	m.SetAnimate(false)
}

func (m *TGIFAnimate) Start() {
	m.SetAnimate(true)
}

func (m *TGIFAnimate) doStop() {
	if m.onStop != nil {
		m.onStop()
	}
}

func (m *TGIFAnimate) doStart() {
	if m.onStop != nil {
		m.onStop()
	}
}

func (m *TGIFAnimate) doFrameChanged(frame *Frame) {
	if m.onFrameChanged != nil {
		m.onFrameChanged(frame)
	}
}

func (m *TGIFAnimate) SetOnStop(fn func()) {
	m.onStop = fn
}

func (m *TGIFAnimate) SetOnStart(fn func()) {
	m.onStart = fn
}

func (m *TGIFAnimate) SetOnFrameChanged(fn func(frame *Frame)) {
	m.onFrameChanged = fn
}

func (m *TGIFAnimate) EnableCache(v bool) {
	m.cache = v
}

func (m *TGIFAnimate) LoadFromFile(filePath string) {
	if m.filePath == filePath {
		return
	}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	g, err := gif.DecodeAll(file)
	if err != nil {
		panic(err)
	}
	m.gif = g
	m.load()
}

func (m *TGIFAnimate) LoadFromBytes(data []byte) {
	g, err := gif.DecodeAll(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	m.gif = g
	m.load()
}
