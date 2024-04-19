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

import "github.com/energye/energy/v2/pkgs/ext"

type TPalette struct {
	colors []*ext.TColor
	count  int
}

func (m *TPalette) Clear() {
	m.colors = make([]*ext.TColor, 0)
}

func (m *TPalette) Free() {
	m.colors = nil
	m.count = 0
}

func (m *TPalette) Add(color *ext.TColor) {
	m.colors = append(m.colors, color)
	m.count++
}

func (m *TPalette) Get(index int) *ext.TColor {
	return m.colors[index]
}

func (m *TPalette) Set(index int, color *ext.TColor) {
	m.colors[index] = color
}

func (m *TPalette) Count() int {
	return m.count
}
