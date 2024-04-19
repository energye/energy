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
)

type TGIFHeader struct {
	Signature       [3]byte // 页眉签名（始终为“GIF”）
	Version         [3]byte // GIF格式版本（“87a”或“89a”）
	ScreenWidth     uint16  // 以像素为单位的显示屏宽度
	ScreenHeight    uint16  // 以像素为单位的显示屏高度
	Packedbit       byte    // 屏幕和彩色地图信息
	BackgroundColor byte    // 背景色索引
	AspectRatio     byte    // 像素纵横比
}

func (m *TGIFHeader) BytesToConvert(data []byte) {
	var buf bytes.Buffer
	buf.Write(data)
	binary.Read(&buf, binary.LittleEndian, m)
}
