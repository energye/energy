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
	"encoding/binary"
	"github.com/energye/energy/v2/pkgs/ext"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"image"
	"image/gif"
	"io/ioutil"
)

const (
	DisposalNone = iota + 1
	DisposalBackground
	DisposalPrevious
)

type TGIFAnimate struct {
	*ext.TImage
	buffImg        *lcl.TBitmap
	task           *lcl.TTimer
	delay          uint32
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
	m.delay = 100
	m.buffImg = lcl.NewBitmap()
	m.TImage = ext.NewImage(owner)
	//m.SetOnPaint(m.onPaint)
	m.task = lcl.NewTimer(owner)
	m.task.SetEnabled(false)
	m.task.SetInterval(m.delay)
	m.task.SetOnTimer(m.onTimer)
	return m
}

func (m *TGIFAnimate) Free() {
	m.Stop()
	m.reset()
	m.TImage.Free()
	m.buffImg.Free()
	m.task.Free()
}

func (m *TGIFAnimate) onTimer(sender lcl.IObject) {
	if m.count == 0 {
		return
	}
	m.scan()
	if m.index >= m.count-1 {
		m.index = 0
	} else {
		m.index++
	}
	nextDelay := m.frames[m.index].delay
	if nextDelay != m.delay {
		m.delay = nextDelay
		m.task.SetEnabled(false)
		m.task.SetInterval(m.delay)
		m.task.SetEnabled(true)
	}
	//m.Repaint() 使用OnPaint有些问题, 这里直接绘制
}

func (m *TGIFAnimate) onPaint(sender lcl.IObject) {
	m.scan()
}

func (m *TGIFAnimate) initialed() {
	m.w = int32(m.gif.Config.Width)
	m.h = int32(m.gif.Config.Height)
	m.SetWidth(m.w)
	m.SetHeight(m.h)
	m.Picture().LoadFromBytes(m.frames[0].data)
	m.buffImg.SetSize(m.w, m.h)
	m.buffImg.SetPixelFormat(types.Pf32bit)
	m.buffImg.SetHandleType(types.BmDIB)
	// 填充整个画布
	//m.Canvas().FillRect(types.Rect(0, 0, m.w, m.h))
	//m.Repaint()
}

func (m *TGIFAnimate) scan() {
	frame := m.currentFrame()
	m.doFrameChanged(frame)
	frame.scan()
	canvas := m.Canvas()
	m.buffImg.Canvas().Draw(frame.x, frame.y, frame.image)
	//canvas.Draw(frame.x, frame.y, frame.image)
	switch frame.method {
	case DisposalNone: // 不处理，图形留在原处
	case DisposalBackground: // 删除当前帧，恢复为上一个帧的内容. 显示图形的区域必须要恢复成背景颜色
		// TOdO 先不处理，当前只处理简单的GIF
		//pFrame := m.priorFrame(frame.index)
		//pFrame.scan()
		////canvas.Draw(pFrame.x, pFrame.y, pFrame.image)
		//m.buffImg.Canvas().Draw(pFrame.x, pFrame.y, pFrame.image)
		//m.task.SetEnabled(false) // debug
	case DisposalPrevious: // 删除当前帧，恢复为 GIF 开始时的状态
		canvas.FillRect(Rect(0, 0, m.w, m.h))
	}
	canvas.Draw(0, 0, m.buffImg)
	//canvas.Draw(frame.x, frame.y, frame.image)
	if !m.cache {
		frame.reset()
	}
}

func (m *TGIFAnimate) load() {
	m.reset()
	m.count = len(m.gif.Image)
	m.frames = make([]*Frame, m.count)
	var frameToBytes = func(frame *image.Paletted) (result []byte) {
		var buf bytes.Buffer
		//err := png.Encode(&buf, frame)
		err := gif.Encode(&buf, frame, nil)
		if err != nil {
			panic(err)
		}
		result = buf.Bytes()
		buf.Reset()
		return
	}
	for i, frame := range m.gif.Image {
		bounds := frame.Bounds()
		m.frames[i] = &Frame{
			index:  i,
			method: m.gif.Disposal[i],
			x:      int32(bounds.Min.X),
			y:      int32(bounds.Min.Y),
			w:      int32(bounds.Dx()),
			h:      int32(bounds.Dy()),
			delay:  uint32(m.gif.Delay[i] * 10),
			data:   frameToBytes(frame),
		}
	}
	m.initialed()
	m.gif = nil
}

//
//func background(width, height int, data []byte) lcl.IBitmap {
//	bmp := lcl.NewBitmap()
//	bmp.SetSize(int32(width), int32(height))
//	bmp.SetPixelFormat(types.Pf32bit)
//	bmp.SetHandleType(types.BmDIB)
//	//bmp.SetTransparent(true)
//
//	img := lcl.NewPngImage()
//	//img.SetTransparent(true)
//	img.LoadFromBytes(data)
//	bmp.Assign(img)
//
//	return bmp
//}

func (m *TGIFAnimate) reset() {
	for _, frame := range m.frames {
		frame.free()
	}
	m.w, m.h = 0, 0
	m.frames = nil
	m.index = 0
	m.count = 0
}

func (m *TGIFAnimate) currentFrame() *Frame {
	return m.frames[m.index]
}

func (m *TGIFAnimate) PrevFrame() {
	m.index--
	if m.index < 0 {
		m.index = m.count - 1
	}
	m.scan()
}

func (m *TGIFAnimate) NextFrame() {
	if m.index >= m.count-1 {
		m.index = 0
	} else {
		m.index++
	}
	m.scan()
}

func (m *TGIFAnimate) priorFrame(index int) *Frame {
	index--
	if index < 0 {
		index = m.count - 1
	}
	return m.frames[index]
}

func (m *TGIFAnimate) SetAnimate(v bool) {
	m.task.SetEnabled(v)
	if v {
		m.doStart()
	} else {
		m.doStop()
	}
}

func (m *TGIFAnimate) Animate() bool {
	return m.task.Enabled()
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
	if m.onStart != nil {
		m.onStart()
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
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	m.LoadFromBytes(data)
}

func (m *TGIFAnimate) LoadFromBytes(data []byte) {
	header := m.ReadHeader(data)
	if header == nil || !header.IsGIF() || !header.Is89a() {
		return
	}
	g, err := gif.DecodeAll(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	m.gif = g
	m.load()
}

func (m *TGIFAnimate) ReadHeader(data []byte) *TGIFHeader {
	if data != nil && len(data) > 13 {
		header := new(TGIFHeader)
		var buf bytes.Buffer
		buf.Write(data[:13])
		binary.Read(&buf, binary.LittleEndian, header)
		return header
	}
	return nil
}

func Rect(left, top, right, bottom int32) types.TRect {
	return types.TRect{Left: left, Top: top, Right: right, Bottom: bottom}
}
