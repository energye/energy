//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

// IDataObject ActiveX > IDataObject
//
//	ole: github.com/go-ole/go-ole
//	UUID: '{0000010e-0000-0000-C000-000000000046}'
//	COM https://learn.microsoft.com/zh-cn/windows/win32/api/objidl/
type IDataObject interface {
	IUnknown
}

// DataObject ActiveX > IDataObject
type DataObject struct {
	Unknown
}
