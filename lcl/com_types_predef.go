//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/types"
	"unsafe"
)

// ISequentialStream Types > ISequentialStream
type ISequentialStream interface {
	IUnknown
	Read(pv uintptr, cb types.DWORD, pcbRead uintptr) types.HRESULT
	Write(pv uintptr, cb types.DWORD, pcbWritten uintptr) types.HRESULT
}

// SequentialStream Types > ISequentialStream
type SequentialStream struct {
	Unknown
}

// ICOMStream Types > IStream
type ICOMStream interface {
	ISequentialStream
	Seek(dlibMove int64, dwOrigin types.DWORD, outLibNewPosition *int64) types.HRESULT
	SetSize(libNewSize int64) types.HRESULT
	CopyTo(stm IStream, cb int64, outCbRead *int64, outCbWritten *int64) types.HRESULT
	Commit(grfCommitFlags types.DWORD) types.HRESULT
	Revert() types.HRESULT
	LockRegion(libOffset int64, cb int64, dwLockType types.DWORD) types.HRESULT
	UnlockRegion(libOffset int64, cb int64, dwLockType types.DWORD) types.HRESULT
	Stat(outStatStg *TStatStg, grfStatFlag types.DWORD) types.HRESULT
	Clone(outStm *IStream) types.HRESULT
}

// COMStream Types > IStream
type COMStream struct {
	SequentialStream
}

func (m *SequentialStream) Read(pv uintptr, cb types.DWORD, pcbRead uintptr) types.HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.COMSequentialStreamRead(), pv, uintptr(cb), pcbRead)
	return types.HRESULT(r1)
}

func (m *SequentialStream) Write(pv uintptr, cb types.DWORD, pcbWritten uintptr) types.HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.COMSequentialStreamWrite(), pv, uintptr(cb), pcbWritten)
	return types.HRESULT(r1)
}

func (m *COMStream) Seek(dlibMove int64, dwOrigin types.DWORD, outLibNewPosition *int64) types.HRESULT {
	if outLibNewPosition == nil {
		return 0
	}
	r1 := api.LCLPreDef().SysCallN(api.COMStreamSeek(), uintptr(unsafe.Pointer(&dlibMove)), uintptr(dwOrigin), uintptr(unsafe.Pointer(outLibNewPosition)))
	return types.HRESULT(r1)
}

func (m *COMStream) SetSize(libNewSize int64) types.HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.COMStreamSetSize(), uintptr(unsafe.Pointer(&libNewSize)))
	return types.HRESULT(r1)
}

func (m *COMStream) CopyTo(stm IStream, cb int64, outCbRead *int64, outCbWritten *int64) types.HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.COMStreamCopyTo(), stm.Instance(), uintptr(unsafe.Pointer(&cb)), uintptr(unsafe.Pointer(outCbRead)), uintptr(unsafe.Pointer(outCbWritten)))
	return types.HRESULT(r1)
}

func (m *COMStream) Commit(grfCommitFlags types.DWORD) types.HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.COMStreamCommit(), uintptr(grfCommitFlags))
	return types.HRESULT(r1)
}

func (m *COMStream) Revert() types.HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.COMStreamRevert())
	return types.HRESULT(r1)
}

func (m *COMStream) LockRegion(libOffset int64, cb int64, dwLockType types.DWORD) types.HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.COMStreamLockRegion(), uintptr(unsafe.Pointer(&libOffset)), uintptr(unsafe.Pointer(&cb)), uintptr(dwLockType))
	return types.HRESULT(r1)
}

func (m *COMStream) UnlockRegion(libOffset int64, cb int64, dwLockType types.DWORD) types.HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.COMStreamUnlockRegion(), uintptr(unsafe.Pointer(&libOffset)), uintptr(unsafe.Pointer(&cb)), uintptr(dwLockType))
	return types.HRESULT(r1)
}

func (m *COMStream) Stat(outStatStg *TStatStg, grfStatFlag types.DWORD) types.HRESULT {
	var resultStatStg uintptr
	r1 := api.LCLPreDef().SysCallN(api.COMStreamStat(), uintptr(unsafe.Pointer(&resultStatStg)), uintptr(grfStatFlag))
	if resultStatStg != 0 {
		statStg := (*tStatStgPtr)(getPointer(resultStatStg))
		if outStatStg != nil {
			outStatStg.PwcsName = api.GoStr(statStg.PwcsName)
			outStatStg.DwType = *(*types.DWORD)(getPointer(statStg.DwType))
			outStatStg.CbSize = *(*int64)(getPointer(statStg.CbSize))
			outStatStg.Mtime = api.GoStr(statStg.Mtime)
			outStatStg.Ctime = api.GoStr(statStg.Ctime)
			outStatStg.Atime = api.GoStr(statStg.Atime)
			outStatStg.GrfMode = *(*types.DWORD)(getPointer(statStg.GrfMode))
			outStatStg.GrfLocksSupported = *(*types.DWORD)(getPointer(statStg.GrfLocksSupported))
			outStatStg.Clsid = api.GoStr(statStg.Clsid)
			outStatStg.GrfStateBits = *(*types.DWORD)(getPointer(statStg.GrfStateBits))
			outStatStg.Reserved = *(*types.DWORD)(getPointer(statStg.Reserved))
		}
	}
	return types.HRESULT(r1)
}

func (m *COMStream) Clone(outStm *IStream) types.HRESULT {
	var result uintptr
	r1 := api.LCLPreDef().SysCallN(api.COMStreamClone(), uintptr(unsafe.Pointer(&result)))
	if result != 0 {
		*outStm = AsStream(result)
	}
	return types.HRESULT(r1)
}
