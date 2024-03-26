//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import "github.com/energye/energy/v2/emfs"

// LoadFromFSFile 从FS文件加载。
func (m *TIcon) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

func (m *TIcon) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

// LoadFromFSFile 从FS文件加载。
func (m *TTreeView) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

// LoadFromFSFile 从FS文件加载。
func (m *TBitmap) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

func (m *TBitmap) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

// LoadFromFSFile 从FS文件加载。
func (m *TGIFImage) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

func (m *TGIFImage) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

// LoadFromBytes 从字节加载
func (m *TGraphic) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

// LoadFromFSFile 从FS文件加载。
func (m *TJPEGImage) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

func (m *TJPEGImage) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

// LoadFromFSFile 从FS文件加载。
func (m *TMemoryStream) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

// LoadFromFSFile 从FS文件加载。
func (m *TPicture) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

// LoadFromFSFile 从FS文件加载。
func (m *TPortableNetworkGraphic) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

// LoadFromFSFile 从FS文件加载。
func (m *TStringList) LoadFromFSFile(Filename string) error {
	bytes, err := emfs.GetResources(Filename)
	if err != nil {
		return err
	}
	m.LoadFromBytes(bytes)
	return nil
}

// LoadFromBytes 文件流加载。
func (m *TTreeView) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}

// LoadFromBytes 文件流加载。
func (m *TMemoryStream) LoadFromBytes(data []byte) {
	if len(data) == 0 {
		return
	}
	mem := NewMemoryStream()
	defer mem.Free()
	mem.Write(data)
	mem.SetPosition(0)
	m.LoadFromStream(mem)
}
