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
	"encoding/binary"
	"github.com/energye/energy/v2/pkgs/ext"
)

type TGIFImage struct {
	Bitmap *ext.TBitmap
	PosX   uint16
	PosY   uint16
	Delay  uint16
	Method byte
}

type TGIFImageDescriptor struct {
	Left, // 显示器上图像的X位置
	Top, // 显示器上图像的Y位置
	Width, // 图像的宽度（像素）
	Height uint16 // 图像的高度（像素）
	Packedbit byte // 图像和颜色表数据信息
}

func (m *TGIFImageDescriptor) BytesToConvert(data []byte) {
	var buf bytes.Buffer
	buf.Write(data)
	binary.Read(&buf, binary.LittleEndian, m)
}

type TGIFGraphicsControlExtension struct {
	BlockSize  byte   // 剩余字段的大小（始终为04h）
	Packedbit  byte   // 要使用的图形处理方法
	DelayTime  uint16 // 等待数十秒
	ColorIndex byte   // 透明颜色索引
	Terminator byte   // 块终止符（始终为0）
}

func (m *TGIFGraphicsControlExtension) BytesToConvert(data []byte) {
	var buf bytes.Buffer
	buf.Write(data)
	binary.Read(&buf, binary.LittleEndian, m)
}
