//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package gifanim

type TGIFHeader struct {
	Signature       [3]byte // 页眉签名（始终为“GIF”）
	Version         [3]byte // GIF格式版本（“89a”）
	ScreenWidth     uint16  // 以像素为单位的显示屏宽度
	ScreenHeight    uint16  // 以像素为单位的显示屏高度
	Packedbit       byte    // 屏幕和彩色地图信息
	BackgroundColor byte    // 背景色索引
	AspectRatio     byte    // 像素纵横比
}

func (m *TGIFHeader) Is89a() bool {
	if m == nil {
		return false
	}
	return m.Version[0] == '8' && m.Version[1] == '9' && m.Version[2] == 'a'
}

func (m *TGIFHeader) IsGIF() bool {
	if m == nil {
		return false
	}
	return m.Signature[0] == 'G' && m.Signature[1] == 'I' && m.Signature[2] == 'F'
}
