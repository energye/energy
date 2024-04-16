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

import (
	"bytes"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl"
	"image/gif"
	"os"
)

// TGIFPlay GIF 图片播放组件
type TGIFPlay struct {
	*lcl.TImage              // 展示每帧图像
	cache        [][]byte    // gif缓存
	cacheCount   int         //
	autoPlay     bool        // 自动默认，默认 true
	currentFrame int         // 当前帧
	playTask     *lcl.TTimer // 播放任务
}

// NewGIFPlay 创建一个GIFPlay
func NewGIFPlay(owner lcl.IComponent) *TGIFPlay {
	m := &TGIFPlay{
		playTask: lcl.NewTimer(owner),
	}
	m.TImage = lcl.NewImage(owner)
	m.playTask.SetInterval(66)
	m.playTask.SetEnabled(false)
	m.play()
	return m
}

// IsValid 返回GIFPlay是否有效
func (m *TGIFPlay) IsValid() bool {
	return m.playTask != nil && m.TImage.IsValid()
}

func (m *TGIFPlay) play() {
	m.playTask.SetOnTimer(func(sender lcl.IObject) {
		if m.cacheCount == 0 || m.cache == nil {
			return
		}
		if m.currentFrame >= m.cacheCount {
			m.currentFrame = 0
		}
		m.Picture().LoadFromBytes(m.cache[m.currentFrame])
		m.currentFrame++
	})
}

// Free 停止播放并释放掉这个GIFPlay, 释放后将不可用
func (m *TGIFPlay) Free() {
	m.Stop()
	m.playTask = nil
	m.TImage.Free()
	m.cache = nil
}

// Stop 在当前帧停止播放
func (m *TGIFPlay) Stop() {
	if m.playTask != nil {
		m.playTask.SetEnabled(false)
	}
}

// Start 在当前帧开始播放
func (m *TGIFPlay) Start() {
	if m.playTask != nil {
		m.playTask.SetEnabled(true)
	}
}

// PlaybackSpeed 设置播放速度，毫秒 默认 66
func (m *TGIFPlay) PlaybackSpeed(playbackSpeed uint32) {
	if m.playTask != nil {
		m.playTask.SetInterval(playbackSpeed)
	}
}

// LoadFile 在本地加载GIF
func (m *TGIFPlay) LoadFile(filePath string) {
	if m.playTask == nil {
		return
	}
	file, err := os.Open(filePath)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	frames, err := gif.DecodeAll(file)
	if err != nil {
		panic(err)
	}
	m.cacheGIFFrames(frames)
}

// LoadFSFile 在内置FS中加载GIF
func (m *TGIFPlay) LoadFSFile(filePath string) {
	if m.playTask == nil {
		return
	}
	data, err := emfs.GetResources(filePath)
	if err != nil {
		panic(err.Error())
	}
	m.LoadBytes(data)
}

// LoadBytes 在图片字节里加载GIF
func (m *TGIFPlay) LoadBytes(data []byte) {
	if m.playTask == nil {
		return
	}
	var buf bytes.Buffer
	buf.Write(data)
	frames, err := gif.DecodeAll(&buf)
	if err != nil {
		panic(err)
	}
	defer buf.Reset()
	m.cacheGIFFrames(frames)
}

// CacheFrames 返回所有帧
func (m *TGIFPlay) CacheFrames() [][]byte {
	return m.cache
}

func (m *TGIFPlay) AppendFrame(data []byte) {

}

func (m *TGIFPlay) RemoveFrame(index int) {

}

func (m *TGIFPlay) InsertFrame(index int, data []byte) {

}

func (m *TGIFPlay) cacheGIFFrames(frames *gif.GIF) {
	if m.playTask == nil {
		return
	}
	m.cacheCount = len(frames.Image)
	m.cache = make([][]byte, m.cacheCount)
	for i, img := range frames.Image {
		var buf bytes.Buffer
		err := gif.Encode(&buf, img, nil)
		if err != nil {
			panic(err.Error())
		}
		m.cache[i] = buf.Bytes()
	}
}
