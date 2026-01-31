//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package application

import (
	"github.com/energye/lcl/lcl"
	"strings"
)

type TTrayImageList struct {
	imageList  lcl.IImageList
	imageIndex map[string]int32
}

func (m *TTrayImageList) ImageIndex(imageName string) int32 {
	index, ok := m.imageIndex[strings.ToLower(imageName)]
	if ok {
		return index
	}
	return -1
}

func (m *TTrayImageList) setImageListData(data []byte, name string, index int32) {
	pic := lcl.NewPicture()
	defer pic.Free()
	mem := lcl.NewMemoryStream()
	defer mem.Free()
	lcl.StreamHelper.WriteBuffer(mem, data)
	mem.SetPosition(0)
	pic.LoadFromStream(mem)
	m.imageList.Add(pic.Bitmap(), nil)
	if name != "" && index != -1 {
		m.imageIndex[name] = index
	}
}
