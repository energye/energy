//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

func (m *TPicture) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}
