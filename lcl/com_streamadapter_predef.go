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
	. "github.com/energye/energy/v2/types"
	"unsafe"
)

// IStreamAdapter Parent: IInterfacedObject
type IStreamAdapter interface {
	IInterfacedObject
	ICOMStream
	Stream() IStream                                                             // property
	StreamOwnership() TStreamOwnership                                           // property
	SetStreamOwnership(AValue TStreamOwnership)                                  // property
	Read(pv uintptr, cb DWORD, pcbRead uintptr) HRESULT                          // function
	Write(pv uintptr, cb DWORD, pcbWritten uintptr) HRESULT                      // function
	Seek(dlibMove int64, dwOrigin DWORD, OutLibNewPosition *int64) HRESULT       // function
	SetSize(libNewSize int64) HRESULT                                            // function
	CopyTo(stm IStream, cb int64, OutCbRead *int64, OutCbWritten *int64) HRESULT // function
	Commit(grfCommitFlags DWORD) HRESULT                                         // function
	Revert() HRESULT                                                             // function
	LockRegion(libOffset int64, cb int64, dwLockType DWORD) HRESULT              // function
	UnlockRegion(libOffset int64, cb int64, dwLockType DWORD) HRESULT            // function
	Stat(outStatStg *TStatStg, grfStatFlag DWORD) HRESULT                        // function
	Clone(OutStm *IStream) HRESULT                                               // function
}

// TStreamAdapter Parent: TInterfacedObject
type TStreamAdapter struct {
	TInterfacedObject
	COMStream
}

func NewStreamAdapter(stream IStream, ownership TStreamOwnership) IStreamAdapter {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterCreate(), GetObjectUintptr(stream), uintptr(ownership))
	return AsStreamAdapter(r1)
}

func (m *TStreamAdapter) Stream() IStream {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterStream(), m.Instance())
	return AsStream(r1)
}

func (m *TStreamAdapter) StreamOwnership() TStreamOwnership {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterStreamOwnership(), 0, m.Instance(), 0)
	return TStreamOwnership(r1)
}

func (m *TStreamAdapter) SetStreamOwnership(AValue TStreamOwnership) {
	api.LCLPreDef().SysCallN(api.StreamAdapterStreamOwnership(), 1, m.Instance(), uintptr(AValue))
}

func (m *TStreamAdapter) Read(pv uintptr, cb DWORD, pcbRead uintptr) HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterRead(), m.Instance(), uintptr(pv), uintptr(cb), uintptr(pcbRead))
	return HRESULT(r1)
}

func (m *TStreamAdapter) Write(pv uintptr, cb DWORD, pcbWritten uintptr) HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterWrite(), m.Instance(), uintptr(pv), uintptr(cb), uintptr(pcbWritten))
	return HRESULT(r1)
}

func (m *TStreamAdapter) Seek(dlibMove int64, dwOrigin DWORD, OutLibNewPosition *int64) HRESULT {
	var result2 uintptr
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterSeek(), m.Instance(), uintptr(dlibMove), uintptr(dwOrigin), uintptr(unsafe.Pointer(&result2)))
	*OutLibNewPosition = int64(result2)
	return HRESULT(r1)
}

func (m *TStreamAdapter) SetSize(libNewSize int64) HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterSetSize(), m.Instance(), uintptr(unsafe.Pointer(&libNewSize)))
	return HRESULT(r1)
}

func (m *TStreamAdapter) CopyTo(stm IStream, cb int64, OutCbRead *int64, OutCbWritten *int64) HRESULT {
	var result2 uintptr
	var result3 uintptr
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterCopyTo(), m.Instance(), stm.Instance(), uintptr(unsafe.Pointer(&cb)), uintptr(unsafe.Pointer(&result2)), uintptr(unsafe.Pointer(&result3)))
	*OutCbRead = int64(result2)
	*OutCbWritten = int64(result3)
	return HRESULT(r1)
}

func (m *TStreamAdapter) Commit(grfCommitFlags DWORD) HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterCommit(), m.Instance(), uintptr(grfCommitFlags))
	return HRESULT(r1)
}

func (m *TStreamAdapter) Revert() HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterRevert(), m.Instance())
	return HRESULT(r1)
}

func (m *TStreamAdapter) LockRegion(libOffset int64, cb int64, dwLockType DWORD) HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterLockRegion(), m.Instance(), uintptr(unsafe.Pointer(&libOffset)), uintptr(unsafe.Pointer(&cb)), uintptr(dwLockType))
	return HRESULT(r1)
}

func (m *TStreamAdapter) UnlockRegion(libOffset int64, cb int64, dwLockType DWORD) HRESULT {
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterUnlockRegion(), m.Instance(), uintptr(unsafe.Pointer(&libOffset)), uintptr(unsafe.Pointer(&cb)), uintptr(dwLockType))
	return HRESULT(r1)
}

func (m *TStreamAdapter) Stat(outStatStg *TStatStg, grfStatFlag DWORD) HRESULT {
	var resultStatStg uintptr
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterStat(), m.Instance(), uintptr(unsafe.Pointer(&resultStatStg)), uintptr(grfStatFlag))
	if resultStatStg != 0 {
		statStg := (*tStatStgPtr)(getPointer(resultStatStg))
		if outStatStg != nil {
			outStatStg.PwcsName = api.GoStr(statStg.PwcsName)
			outStatStg.DwType = *(*DWORD)(getPointer(statStg.DwType))
			outStatStg.CbSize = *(*int64)(getPointer(statStg.CbSize))
			outStatStg.Mtime = api.GoStr(statStg.Mtime)
			outStatStg.Ctime = api.GoStr(statStg.Ctime)
			outStatStg.Atime = api.GoStr(statStg.Atime)
			outStatStg.GrfMode = *(*DWORD)(getPointer(statStg.GrfMode))
			outStatStg.GrfLocksSupported = *(*DWORD)(getPointer(statStg.GrfLocksSupported))
			outStatStg.Clsid = api.GoStr(statStg.Clsid)
			outStatStg.GrfStateBits = *(*DWORD)(getPointer(statStg.GrfStateBits))
			outStatStg.Reserved = *(*DWORD)(getPointer(statStg.Reserved))
		}
	}
	return HRESULT(r1)
}

func (m *TStreamAdapter) Clone(OutStm *IStream) HRESULT {
	var result0 uintptr
	r1 := api.LCLPreDef().SysCallN(api.StreamAdapterClone(), m.Instance(), uintptr(unsafe.Pointer(&result0)))
	*OutStm = AsStream(result0)
	return HRESULT(r1)
}
