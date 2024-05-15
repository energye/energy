//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

// IVTEditLink TODO no impl
//
//	Communication interface between a tree editor and the tree itself (declared as using stdcall in case it
//	is implemented in a (C/C++) DLL). The GUID is not nessecary in Delphi but important for BCB users
//	to allow QueryInterface and _uuidof calls.
//	{2BE3EAFA-5ACB-45B4-9D9A-B58BCC496E17}
type IVTEditLink struct {
	instance unsafePointer
}
