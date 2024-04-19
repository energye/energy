//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// GIF Play component

package gifplay

import (
	"bytes"
	"encoding/binary"
	"github.com/energye/energy/v2/pkgs/ext"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
)

const (
	EXT_INTRODUCER       byte = 0x21
	EXT_GRAPHICS_CONTROL byte = 0xF9
	EXT_PLAIN_TEXT       byte = 0x01
	EXT_APPLICATION      byte = 0xFF
	EXT_COMMENT          byte = 0xFE
	DSC_LOCAL_IMAGE      byte = 0x2C
	ID_TRANSPARENT       byte = 0x01
	ID_COLOR_TABLE_SIZE  byte = 0x07
	ID_SORT              byte = 0x20
	ID_INTERLACED        byte = 0x40
	ID_COLOR_TABLE       byte = 0x80
	ID_IMAGE_DESCRIPTOR  byte = 0x2C
	ID_TRAILER           byte = 0x3B
	CODE_TABLE_SIZE           = 4096
)

const (
	AlphaTransparent uint16 = 0x0000
	AlphaOpaque      uint16 = 0xFFFF
)

type TRGB struct {
	Red, Green, Blue byte
}

//type buffer struct {
//	buf []byte
//}
//
//func (b *buffer) Read(p []byte) (n int, err error) {
//	if b.buf == nil {
//		if len(p) == 0 {
//			return 0, nil
//		}
//		return 0, io.EOF
//	}
//	n = copy(p, b.buf)
//	return n, nil
//}

func (m *TRGB) BytesToConvert(data []byte) {
	var buf bytes.Buffer
	buf.Write(data)
	binary.Read(&buf, binary.LittleEndian, m)
}

type TGIFPlay struct {
	*lcl.TPanel
	animate        bool
	empty          bool
	filePath       string
	gifBitmaps     *TGIFList
	onFrameChanged lcl.TNotifyEvent
	onStart        lcl.TNotifyEvent
	onStop         lcl.TNotifyEvent
	wait           *lcl.TTimer
	currentImage   int32
	gifHeightidth  int32
	gifWidth       int32
	bufferImg      *ext.TBitmap
	currentView    *ext.TBitmap
}

func NewGIFPlay(owner lcl.IComponent) (gifPlay *TGIFPlay) {
	control := lcl.NewPanel(owner)
	control.SetAutoSize(true)
	control.SetBounds(0, 0, 90, 90)
	gifPlay = &TGIFPlay{
		TPanel: control,
	}
	gifPlay.empty = true
	gifPlay.currentImage = 0
	gifPlay.currentView = ext.NewBitmap()
	gifPlay.bufferImg = ext.NewBitmap()
	gifPlay.wait = lcl.NewTimer(owner)
	gifPlay.wait.SetInterval(100)
	gifPlay.wait.SetEnabled(false)
	gifPlay.wait.SetOnTimer(gifPlay.OnTime)
	gifPlay.SetAnimate(true)
	gifPlay.SetOnPaint(gifPlay.OnPaint)
	return
}

func (m *TGIFPlay) OnPaint(sender lcl.IObject) {
	if !m.empty && m.IsValid() {
		if m.currentImage < m.gifBitmaps.count {
			m.bufferImg.Canvas().Brush().SetColor(m.Color())
			currentImage := m.gifBitmaps.GetItem(m.currentImage)
			if m.currentImage == 0 {
				m.bufferImg.Canvas().FillRect(types.Rect(0, 0, m.Width(), m.Height()))
			}
			if currentImage.Delay != 0 {
				m.wait.SetInterval(uint32(currentImage.Delay * 10))
			}
			m.bufferImg.Canvas().Draw(int32(currentImage.PosX), int32(currentImage.PosY), currentImage.Bitmap)
			m.currentView.Assign(m.bufferImg)
			switch currentImage.Method {
			//case 0: 未指定...
			//case 1: 无更改背景
			case 2:
				m.bufferImg.Canvas().FillRect(types.Rect(int32(currentImage.PosX), int32(currentImage.PosY),
					currentImage.Bitmap.Width()+int32(currentImage.PosX), currentImage.Bitmap.Height()+int32(currentImage.PosY)))
			case 3:
				m.bufferImg.Canvas().FillRect(types.Rect(0, 0, m.Width(), m.Height()))
			}
		}
		m.Canvas().Draw(0, 0, m.currentView)
	}
}

func (m *TGIFPlay) OnTime(sender lcl.IObject) {
	if !m.empty && m.IsValid() {
		if m.currentImage >= m.gifBitmaps.count {
			m.currentImage = 0
		} else {
			m.currentImage++
		}
		m.Repaint()
	}
}

func (m *TGIFPlay) GifChanged() {
	canvas := m.bufferImg.Canvas()
	canvas.Brush().SetColor(m.Color())
	canvas.FillRect(types.Rect(0, 0, m.Width(), m.Height()))
	currentImage := m.gifBitmaps.GetItem(m.currentImage)
	canvas.Draw(int32(currentImage.PosX), int32(currentImage.PosY), currentImage.Bitmap)
	m.currentView.Assign(m.bufferImg)
	m.Invalidate()
}

func (m *TGIFPlay) ResetImage() {
	if m.gifBitmaps != nil {
		m.gifBitmaps.items = make([]*TGIFImage, 0)
		m.gifBitmaps.count = 0
	}
	m.currentImage = 0
	canvas := m.currentView.Canvas()
	canvas.Brush().SetColor(m.Color())
	canvas.FillRect(types.Rect(0, 0, m.Width(), m.Height()))
}

func (m *TGIFPlay) LoadFromFile(filePath string) {
	if m.filePath == filePath {
		return
	}
	m.filePath = filePath
	m.ResetImage()
	if m.filePath == "" {
		return
	}
	if !m.empty {
		m.GifChanged()
	}
	m.empty = true
	gifLoader := &TGIFLoader{
		filePath:           filePath,
		gifStream:          nil,
		gifUseGraphCtrlExt: false,
		width:              20,
		height:             20,
		palette:            &TPalette{},
	}
	if m.currentView == nil {
		m.currentView = ext.NewBitmap()
	}
	if m.gifBitmaps == nil {
		m.gifBitmaps = &TGIFList{}
	}
	m.empty = !gifLoader.LoadAllBitmap(m.gifBitmaps)
	m.DefineSize(gifLoader.width, gifLoader.height)
	gifLoader.Free()
}

func (m *TGIFPlay) DefineSize(width, height int32) {
	if width == m.gifWidth && height == m.gifHeightidth {
		return
	}
	m.gifWidth = width
	m.gifHeightidth = height
	m.SetWidth(m.gifWidth)
	m.SetHeight(m.gifHeightidth)
	m.bufferImg.SetWidth(m.gifWidth)
	m.bufferImg.SetHeight(m.gifHeightidth)
}

func (m *TGIFPlay) SetAnimate(v bool) {
	if m.animate == v {
		return
	}
	m.animate = v
	m.wait.SetEnabled(v)
	if m.animate {
		m.doStart()
	} else {
		m.doStop()
	}
}

func (m *TGIFPlay) doStart() {

}

func (m *TGIFPlay) doStop() {

}
