//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package gifplay

import (
	"bytes"
	"github.com/energye/energy/v2/pkgs/ext"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"unsafe"
)

type TGIFLoader struct {
	gifHeader          *TGIFHeader
	gifDescriptor      *TGIFImageDescriptor
	gifGraphicsCtrlExt *TGIFGraphicsControlExtension
	gifUseGraphCtrlExt bool
	gifBackgroundColor byte
	interlaced         bool
	scanLine           []byte
	lineSize           int32
	disposalMethod     byte
	empty              bool
	filePath           string
	gifStream          *lcl.TMemoryStream
	height             int32
	isTransparent      bool
	width              int32
	palette            *TPalette
	localHeight        int32
	localWidth         int32
}

func (m *TGIFLoader) LoadAllBitmap(gifBitmaps *TGIFList) bool {
	if m.filePath == "" && m.gifStream == nil {
		return false
	}
	var gifStream *lcl.TMemoryStream
	if m.gifStream != nil {
		gifStream = m.gifStream
	} else {
		gifStream = lcl.NewMemoryStream()
		gifStream.LoadFromFile(m.filePath)
		gifStream.SetPosition(0)
	}
	m.ReadHeader(gifStream)
	if m.gifHeader.Version[0] != '8' && m.gifHeader.Version[1] != '9' && m.gifHeader.Version[2] != 'a' {
		return false
	}
	var introducer byte
	// 跳过第一个块扩展（如果存在）
	for {
		introducer = m.SkipBlock(gifStream)
		if introducer == ID_IMAGE_DESCRIPTOR || introducer == ID_TRAILER {
			break
		}
	}
	for {
		m.ReadGifBitmap(gifStream)
		// 解码扫描线缓冲区中的Gif位图
		m.ReadScanLine(gifStream)
		// 为放置扫描线像素创建临时Fp图像
		intfImage := ext.NewLazIntfImage(m.localWidth, m.localHeight)
		intfImage.DataDescription(ext.Init_BPP32_B8G8R8A8_M1_BIO_TTB, m.localWidth, m.localHeight)
		m.WriteScanLine(intfImage)

		gifBitmap := &TGIFImage{
			Bitmap: ext.NewBitmap(),
			PosX:   m.gifDescriptor.Left,
			PosY:   m.gifDescriptor.Top,
			Method: m.disposalMethod,
			Delay:  m.gifGraphicsCtrlExt.DelayTime,
		}
		gifBitmap.Bitmap.LoadFromIntfImage(intfImage)
		gifBitmaps.Add(gifBitmap)
		intfImage.Free()
		m.scanLine = nil
		m.gifUseGraphCtrlExt = false
		for {
			introducer = m.SkipBlock(gifStream)
			if introducer == ID_IMAGE_DESCRIPTOR || introducer == ID_TRAILER {
				break
			}
		}
		if introducer == ID_TRAILER {
			break
		}
	}
	return true
}

func (m *TGIFLoader) ReadPalette(stream *lcl.TMemoryStream, size int) {
	m.palette.Clear()
	m.palette.count = 0
	var entry TRGB
	sizeOf := int32(unsafe.Sizeof(entry)) //3
	for i := 0; i < size; i++ {
		_, d := stream.Read(sizeOf)
		entry.BytesToConvert(d)
		r := uint16(entry.Red)
		g := uint16(entry.Green)
		b := uint16(entry.Blue)
		color := &ext.TColor{
			Red:   r | (r << 8),
			Green: g | (g << 8),
			Blue:  b | (b << 8),
			Alpha: AlphaOpaque,
		}
		m.palette.Add(color)
	}
}

func (m *TGIFLoader) ReadScanLine(stream *lcl.TMemoryStream) {
	var (
		unpackedSize, packedSize           int64
		data, bits, code                   uint32
		inCode                             uint32
		codeSize                           uint32
		codeMask                           uint32
		freeCode                           uint32
		oldCode                            uint32
		prefix                             [CODE_TABLE_SIZE]uint32
		suffix, stack                      = make([]byte, CODE_TABLE_SIZE), make([]byte, CODE_TABLE_SIZE)
		b, initialCodeSize, firstChar      byte
		clearCode, eoiCode                 uint16
		targetIdx, stackIdx, sourceDataIdx int
	)
	// 解压缩字典的初始化
	_, d := stream.Read(1)
	initialCodeSize = d[0]
	// 压缩尾部
	oldPos := stream.Position()
	for {
		_, d = stream.Read(1)
		b = d[0]
		if b > 0 {
			packedSize += int64(b)
			stream.Seek(int64(b), types.SoCurrent)
		}
		if b == 0 {
			break
		}
	}
	stream.SetPosition(oldPos)
	var sourceBuf bytes.Buffer
	for {
		_, d = stream.Read(1)
		b = d[0]
		if b > 0 {
			_, d = stream.Read(int32(b))
			sourceBuf.Write(d)
		}
		if b == 0 {
			break
		}
	}
	sourceData := sourceBuf.Bytes()
	sourceBuf.Reset()
	codeSize = uint32(initialCodeSize + 1)
	clearCode = 1 << initialCodeSize
	eoiCode = clearCode + 1
	freeCode = uint32(clearCode + 2)
	oldCode = CODE_TABLE_SIZE
	codeMask = (1 << codeSize) - 1
	unpackedSize = int64(m.localWidth * m.localHeight)
	for I := 0; I < int(clearCode); I++ {
		prefix[I] = CODE_TABLE_SIZE
		suffix[I] = byte(I)
	}
	firstChar = 0
	data = 0
	bits = 0
	//解压缩LZW gif
	for unpackedSize > 0 && packedSize > 0 {
		source := uint32(sourceData[sourceDataIdx])
		data += source << bits
		bits += 8
		for bits >= codeSize {
			code = data & codeMask
			data >>= codeSize
			bits -= codeSize
			if code == uint32(eoiCode) {
				break
			}
			if code == uint32(clearCode) {
				codeSize = uint32(initialCodeSize + 1)
				codeMask = 1<<codeSize - 1
				freeCode = uint32(clearCode + 2)
				oldCode = CODE_TABLE_SIZE
				continue
			}
			if code > freeCode {
				break
			}
			if oldCode == CODE_TABLE_SIZE {
				firstChar = suffix[code]
				m.scanLine[targetIdx] = firstChar
				targetIdx++
				unpackedSize--
				oldCode = code
				continue
			}
			inCode = code
			if code == freeCode {
				stack[stackIdx] = firstChar
				stackIdx++
				code = oldCode
			}
			for code > uint32(clearCode) {
				stack[stackIdx] = suffix[code]
				stackIdx++
				code = prefix[code]
			}
			firstChar = suffix[code]
			stack[stackIdx] = firstChar
			stackIdx++
			prefix[freeCode] = oldCode
			suffix[freeCode] = firstChar
			if freeCode == codeMask && codeSize < 12 {
				codeSize++
				codeMask = (1 << codeSize) - 1
			}
			if freeCode < CODE_TABLE_SIZE-1 {
				freeCode++
			}
			oldCode = inCode
			for {
				stackIdx--
				m.scanLine[targetIdx] = stack[stackIdx]
				targetIdx++
				unpackedSize--
				if stackIdx == 0 {
					break
				}
			}
		}
		sourceDataIdx++
		packedSize--
	}
}
func (m *TGIFLoader) ReadHeader(stream *lcl.TMemoryStream) {
	var gifHeader TGIFHeader
	sizeOf := int32(unsafe.Sizeof(gifHeader)) - 1 //13
	_, d := stream.Read(sizeOf)
	gifHeader.BytesToConvert(d)
	m.gifHeader = &gifHeader
	m.gifBackgroundColor = gifHeader.BackgroundColor
	m.width = int32(gifHeader.ScreenWidth)
	m.height = int32(gifHeader.ScreenHeight)
	m.localWidth = int32(gifHeader.ScreenWidth)
	m.localHeight = int32(gifHeader.ScreenHeight)
	m.isTransparent = false
	m.ReadGlobalPalette(stream)
}

func (m *TGIFLoader) ReadGlobalPalette(stream *lcl.TMemoryStream) {
	if m.gifHeader.Packedbit&ID_COLOR_TABLE != 0 {
		colorTableSize := m.gifHeader.Packedbit&ID_COLOR_TABLE_SIZE + 1
		m.ReadPalette(stream, 1<<colorTableSize)
	}
}

func (m *TGIFLoader) ReadGraphCtrlExt() {
	m.isTransparent = (m.gifGraphicsCtrlExt.PackedBit & ID_TRANSPARENT) != 0
	m.disposalMethod = (m.gifGraphicsCtrlExt.PackedBit & 0x1C) >> 2
	if m.isTransparent {
		// 如果透明位图更改alpha通道
		m.gifBackgroundColor = m.gifGraphicsCtrlExt.ColorIndex
		color := m.palette.Get(int(m.gifBackgroundColor))
		color.Alpha = AlphaTransparent
		m.palette.Set(int(m.gifBackgroundColor), color)
	}
}

func (m *TGIFLoader) SkipBlock(stream *lcl.TMemoryStream) byte {
	var introducer, labels, skipByte byte
	_, d := stream.Read(1)
	introducer = d[0]
	if introducer == EXT_INTRODUCER {
		_, d = stream.Read(1)
		labels = d[0]
		switch labels {
		case EXT_COMMENT, EXT_APPLICATION:
			for {
				_, d = stream.Read(1)
				skipByte = d[0]
				if skipByte == 0 {
					break
				}
				stream.Seek(int64(skipByte), types.SoCurrent)
			}
		case EXT_GRAPHICS_CONTROL:
			var gifGraphicsCtrlExt TGIFGraphicsControlExtension
			_, d = stream.Read(int32(unsafe.Sizeof(gifGraphicsCtrlExt))) //6
			gifGraphicsCtrlExt.BytesToConvert(d)
			m.gifGraphicsCtrlExt = &gifGraphicsCtrlExt
			m.gifUseGraphCtrlExt = true
		case EXT_PLAIN_TEXT:
			_, d = stream.Read(1)
			skipByte = d[0]
			stream.Seek(int64(skipByte), types.SoCurrent)
			for {
				_, d = stream.Read(1)
				skipByte = d[0]
				if skipByte == 0 {
					break
				}
				stream.Seek(int64(skipByte), types.SoCurrent)
			}
		}
	}
	return introducer
}

func (m *TGIFLoader) WriteScanLine(intfImage *ext.TLazIntfImage) {
	var (
		row, col    int
		pass, every byte
		scanLineIdx int
	)
	if m.interlaced {
		for pass = 1; pass <= 4; pass++ {
			switch pass {
			case 1:
				row = 0
				every = 8
			case 2:
				row = 4
				every = 8
			case 3:
				row = 2
				every = 4
			case 4:
				row = 1
				every = 2
			}
			for {
				for col = 0; col < int(m.localWidth); col++ {
					color := *m.palette.Get(int(m.scanLine[scanLineIdx]))
					intfImage.Colors(int32(col), int32(row), color)
					scanLineIdx++
				}
				row += int(every)
				if row >= int(m.localHeight) {
					break
				}
			}
		}
	} else {
		for row = 0; row < int(m.localHeight); row++ {
			for col = 0; col < int(m.localWidth); col++ {
				color := *m.palette.Get(int(m.scanLine[scanLineIdx]))
				intfImage.Colors(int32(col), int32(row), color)
				scanLineIdx++
			}
		}
	}
}

func (m *TGIFLoader) ReadGifBitmap(stream *lcl.TMemoryStream) {
	var gifDescriptor TGIFImageDescriptor
	sizeOf := int32(unsafe.Sizeof(gifDescriptor)) - 1 //9
	_, d := stream.Read(sizeOf)
	gifDescriptor.BytesToConvert(d)
	m.gifDescriptor = &gifDescriptor
	m.localWidth = int32(m.gifDescriptor.Width)
	m.localHeight = int32(m.gifDescriptor.Height)
	m.interlaced = m.gifDescriptor.Packedbit&ID_INTERLACED == ID_INTERLACED

	m.lineSize = m.localWidth * (m.localHeight + 1)
	m.scanLine = make([]byte, m.lineSize)

	if m.gifDescriptor.Packedbit&ID_COLOR_TABLE != 0 {
		colorTableSize := m.gifDescriptor.Packedbit&ID_COLOR_TABLE_SIZE + 1
		m.ReadPalette(stream, 1<<colorTableSize)
	}

	if m.gifUseGraphCtrlExt {
		m.ReadGraphCtrlExt()
	}
}

func (m *TGIFLoader) Free() {
	m.palette.Free()
}
