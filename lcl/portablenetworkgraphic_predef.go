//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

func (m *TPortableNetworkGraphic) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

type IPngImage interface {
	IPortableNetworkGraphic
}

type TPngImage struct {
	TPortableNetworkGraphic
}

// NewPngImage type TPortableNetworkGraphic
func NewPngImage() IPngImage {
	return NewPortableNetworkGraphic()
}
