//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

func (m *TStrings) AddStrings2(sArr []string) {
	m.BeginUpdate()
	defer m.EndUpdate()
	for _, v := range sArr {
		m.Add(v)
	}
}

func (m *TStrings) AddStrings3(list IStrings, clearFirst bool) {
	if list == nil {
		return
	}
	m.BeginUpdate()
	defer m.EndUpdate()
	if clearFirst {
		m.Clear()
	}
	if m.Count()+list.Count() > m.Capacity() {
		m.SetCapacity(m.Count() + list.Count())
		for i := int32(0); i < list.Count(); i++ {
			m.AddObject(list.Strings(i), list.Objects(i))
		}
	}
}

func (m *TStrings) AddPair2(name, value string, object IObject) IStrings {
	m.AddObject(name+string(m.NameValueSeparator())+value, object)
	return m
}

func (m *TStrings) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}
