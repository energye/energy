//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import "github.com/energye/energy/v2/lcl"

// IDataObject ActiveX > IDataObject
//
//	ole: github.com/go-ole/go-ole
//	UUID: '{0000010e-0000-0000-C000-000000000046}'
//	COM https://learn.microsoft.com/zh-cn/windows/win32/api/objidl/
type IDataObject = lcl.IDataObject

// DataObject ActiveX > IDataObject
type DataObject = lcl.DataObject
