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
	"github.com/cyber-xxm/energy/v2/pkgs/ext"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"image/png"
	"io/ioutil"
	"sync"
	"unsafe"
)

// 控制使用哪种组件绘制每帧图片
const usePNG = false

const (
	DisposalNone = iota + 1
	DisposalBackground
	DisposalPrevious
)

// TGIFAnimate
//
//	GIF 动画组件
type TGIFAnimate struct {
	*ext.TImage
	buffImg *lcl.TBitmap

	restore     bool
	restoreLock sync.WaitGroup

	task           *lcl.TTimer
	bgColor        color.Color
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

// NewGIFAnimate
//
//	创建GIF动画播放组件
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
	if m.index >= m.count-1 {
		m.index = 0
	} else {
		m.index++
	}
	m.scan()
	nextDelay := m.next(m.index).delay
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
	m.buffImg.SetSize(m.w, m.h)
	m.buffImg.SetPixelFormat(types.Pf32bit)
	m.buffImg.SetHandleType(types.BmDIB)
	//m.Repaint() // 使用OnPaint有些问题, 这里直接绘制
	m.scan()
}

func (m *TGIFAnimate) scan() {
	frame := m.currentFrame()
	m.doFrameChanged(frame)
	frame.scan()
	if m.restore {
		m.restoreLock.Wait()
		preFrame := m.prior(frame.index)
		if preFrame.background != nil {
			m.buffImg.Canvas().Draw(preFrame.x, preFrame.y, preFrame.background)
			preFrame.background.Clear()
		}
		m.restore = false
	}
	m.buffImg.Canvas().Draw(frame.x, frame.y, frame.image)
	m.Picture().Assign(m.buffImg) // Linux 怎么不能使用？ TBitmap
	switch frame.method {
	case DisposalNone:
	case DisposalBackground:
		go func() {
			m.restoreLock.Add(1)
			if frame.background == nil {
				frame.background = lcl.NewBitmap()
			}
			m.fillBackground(frame.w, frame.h, frame.background)
			m.restoreLock.Done()
		}()
		m.restore = true
	case DisposalPrevious:
		//canvas.FillRect(Rect(0, 0, m.w, m.h))
	}
	if !m.cache {
		frame.reset()
	}
}

func (m *TGIFAnimate) fillBackground(width, height int32, bmp *lcl.TBitmap) {
	bmp.SetSize(width, height)
	bmp.SetPixelFormat(types.Pf32bit)
	bmp.SetHandleType(types.BmDIB)
	//br, bg, bb, ba := m.bgColor.RGBA()
	for y := 0; y < int(height); y++ {
		ptr := bmp.ScanLine(int32(y))
		for x := 0; x < int(width); x++ {
			pixelIndex := x * 4
			r := (*byte)(unsafe.Pointer(ptr + uintptr(pixelIndex)))
			g := (*byte)(unsafe.Pointer(ptr + uintptr(pixelIndex+1)))
			b := (*byte)(unsafe.Pointer(ptr + uintptr(pixelIndex+2)))
			a := (*byte)(unsafe.Pointer(ptr + uintptr(pixelIndex+3)))
			//*r = uint8(br >> 8)
			//*g = uint8(bg >> 8)
			//*b = uint8(bb >> 8)
			//*a = uint8(ba >> 8)
			*r = 0 //uint8(br >> 8)
			*g = 0 //uint8(bg >> 8)
			*b = 0 //uint8(bb >> 8)
			*a = 0 //uint8(ba >> 8)
		}
	}
}

func (m *TGIFAnimate) load() {
	m.reset()
	m.count = len(m.gif.Image)
	m.frames = make([]*Frame, m.count)
	var frameToBytes = func(frame *image.Paletted) (result []byte) {
		var (
			buf bytes.Buffer
			err error
		)
		if usePNG {
			err = png.Encode(&buf, frame)
		} else {
			err = gif.Encode(&buf, frame, nil)
		}
		if err != nil {
			panic(err)
		}
		result = buf.Bytes()
		buf.Reset()
		return
	}
	m.bgColor = palette.Plan9[m.gif.BackgroundIndex]
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

func (m *TGIFAnimate) prior(index int) *Frame {
	index--
	if index < 0 {
		index = m.count - 1
	}
	return m.frames[index]
}

// PrevFrame 播放上一帧
func (m *TGIFAnimate) PrevFrame() {
	m.index--
	if m.index < 0 {
		m.index = m.count - 1
	}
	m.scan()
}

func (m *TGIFAnimate) next(index int) *Frame {
	index++
	if index >= m.count-1 {
		index = 0
	}
	return m.frames[index]
}

// NextFrame 播放下一帧
func (m *TGIFAnimate) NextFrame() {
	if m.index >= m.count-1 {
		m.index = 0
	} else {
		m.index++
	}
	m.scan()
}

// SetAnimate 设置GIF播放或停止
func (m *TGIFAnimate) SetAnimate(v bool) {
	m.task.SetEnabled(v)
	if v {
		m.doStart()
	} else {
		m.doStop()
	}
}

// Animate 返回当前GIF播放或停止
func (m *TGIFAnimate) Animate() bool {
	return m.task.Enabled()
}

// CurrentFrameIndex 返回当前帧索引
func (m *TGIFAnimate) CurrentFrameIndex() int {
	return m.index
}

// FrameCount 返回帧总数
func (m *TGIFAnimate) FrameCount() int {
	return m.count
}

// Frame 返回指定索引帧
func (m *TGIFAnimate) Frame(index int) *Frame {
	if index < m.count {
		return m.frames[index]
	}
	return nil
}

// Stop 停止播放
func (m *TGIFAnimate) Stop() {
	m.SetAnimate(false)
}

// Start 开始播放
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

// SetOnStop 停止播放事件
func (m *TGIFAnimate) SetOnStop(fn func()) {
	m.onStop = fn
}

// SetOnStart 开始播放事件
func (m *TGIFAnimate) SetOnStart(fn func()) {
	m.onStart = fn
}

// SetOnFrameChanged 播放每帧改变事件
func (m *TGIFAnimate) SetOnFrameChanged(fn func(frame *Frame)) {
	m.onFrameChanged = fn
}

// EnableCache 启用缓存
func (m *TGIFAnimate) EnableCache(v bool) {
	m.cache = v
}

// LoadFromFile 从本地文件加载GIF
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

// LoadFromBytes 从字节数组加载GIF
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
