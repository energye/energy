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

type TGIFList struct {
	items []*TGIFImage
	count int32
}

func (m *TGIFList) Add(AGifImage *TGIFImage) int32 {
	m.items = append(m.items, AGifImage)
	m.count++
	return m.count
}

//func (m *TGIFList) Extract(Item *TGIFImage) *TGIFImage {
//
//}
//
//func (m *TGIFList) Remove(AGifImage *TGIFImage) int32 {
//
//}
//
//func (m *TGIFList) IndexOf(AGifImage *TGIFImage) int32 {
//
//}

func (m *TGIFList) First() *TGIFImage {
	return m.items[0]
}

func (m *TGIFList) Last() *TGIFImage {
	return m.items[len(m.items)-1]
}

func (m *TGIFList) Count() int32 {
	return m.count
}

func (m *TGIFList) Insert(index int32, AGifImage *TGIFImage) {
	m.items[index] = AGifImage
}

func (m *TGIFList) GetItem(index int32) *TGIFImage {
	if index < m.count {
		return m.items[index]
	}
	return nil
}

func (m *TGIFList) SetItem(index int32, AGifImage *TGIFImage) {
	m.items[index] = AGifImage
}
