//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/energye/energy/v2/api/internal/lcl"
	. "github.com/energye/energy/v2/types"
	"unsafe"
)

// -------------------- TApplication ---------------------------

func Application_Instance() uintptr {
	return defSyscallN(lcl.Application_Instance)
}

func Application_CreateForm(app uintptr) uintptr {
	return defSyscallN(lcl.Application_CreateForm, app)
}

func Application_Run(app uintptr) {
	defer func() {
		LibRelease()
	}()
	defSyscallN(lcl.Application_Run, app)
}

func Application_Initialize(obj uintptr) {
	defSyscallN(lcl.Application_Initialize, obj)
}

func Application_SetRunLoopReceived(obj uintptr, proc uintptr) {
	defSyscallN(lcl.Application_SetRunLoopReceived, obj, proc)
}

// -------------------- TForm ---------------------------

func Form_InheritedWndProc(obj uintptr, message *TMessage) uintptr {
	return defSyscallN(lcl.Form_InheritedWndProc, obj, uintptr(unsafe.Pointer(message)))
}

func Form_Create2(owner uintptr) uintptr {
	return defSyscallN(lcl.Form_Create2, owner)
}

func Form_SetOnWndProc(obj uintptr, fn interface{}) {
	defSyscallN(lcl.Form_SetOnWndProc, obj, MakeEventDataPtr(fn))
}

func Form_SetGoPtr(obj uintptr, ptr uintptr) {
	defSyscallN(lcl.Form_SetGoPtr, obj, ptr)
}

// -------------------- Form Resource ---------------------------
//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 需要配合窗口设计器使用
//----------------------------------------

// ResFormLoadFromStream
//
// 从流中加载窗口资源
func ResFormLoadFromStream(obj, root uintptr) {
	defSyscallN(lcl.ResFormLoadFromStream, obj, root)
}

// ResFormLoadFromFile
//
// 从文件中加载窗口资源
func ResFormLoadFromFile(filename string, root uintptr) {
	defSyscallN(lcl.ResFormLoadFromFile, PascalStr(filename), root)
}

// ResFormLoadFromResourceName
//
// 从指定资源中加载窗口资源
func ResFormLoadFromResourceName(instance uintptr, resName string, root uintptr) {
	defSyscallN(lcl.ResFormLoadFromResourceName, instance, PascalStr(resName), root)
}

// -------------------- TClipboard ---------------------------

func Clipboard_Instance() uintptr {
	return defSyscallN(lcl.CLIPBOARD_INSTANCE)
}

func DSetClipboard(obj uintptr) uintptr {
	return defSyscallN(lcl.DSETCLIPBOARD, obj)
}

func DRegisterClipboardFormat(aFormat string) TClipboardFormat {
	return TClipboardFormat(defSyscallN(lcl.DREGISTERCLIPBOARDFORMAT, PascalStr(aFormat)))
}

func DPredefinedClipboardFormat(aFormat TPredefinedClipboardFormat) TClipboardFormat {
	return TClipboardFormat(defSyscallN(lcl.DPREDEFINEDCLIPBOARDFORMAT, uintptr(aFormat)))
}

// -------------------- TMouse ---------------------------

func Mouse_Instance() uintptr {
	return defSyscallN(lcl.MOUSE_INSTANCE)
}

// -------------------- TPrinter ---------------------------

func Printer_Instance() uintptr {
	return defSyscallN(lcl.PRINTER_INSTANCE)
}

func Printer_SetPrinter(obj uintptr, aName string) {
	defSyscallN(lcl.PRINTER_SETPRINTER, obj, PascalStr(aName))
}

// -------------------- TScreen ---------------------------

func Screen_Instance() uintptr {
	return defSyscallN(lcl.SCREEN_INSTANCE)
}

// -------------------- Procedures/Functions ---------------------------

func SetEventCallback(ptr uintptr) {
	defSyscallN(lcl.SetEventCallback, ptr)
}

func SetMessageCallback(ptr uintptr) {
	defSyscallN(lcl.SetMessageCallback, ptr)
}

func SetRequestCallCreateParamsCallback(ptr uintptr) {
	defSyscallN(lcl.SetRequestCallCreateParamsCallback, ptr)
}

func SetRemoveEventCallback(ptr uintptr) {
	defSyscallN(lcl.SetRemoveEventCallback, ptr)
}

func SetIPCEventCallback(ptr uintptr) {
	defSyscallN(lcl.SetIPCEventCallback, ptr)
}

func SetExceptionHandlerCallback(ptr uintptr) {
	defSyscallN(lcl.SetExceptionHandlerCallback, ptr)
}

func DGetStringArrOf(p uintptr, index int) string {
	return GoStr(defSyscallN(lcl.DGETSTRINGARROF, p, uintptr(index)))
}

func DStrLen(p uintptr) int {
	return int(defSyscallN(lcl.DSTRLEN, p))
}

func DMove(src, dest uintptr, nLen int) {
	defSyscallN(lcl.DMOVE, src, dest, uintptr(nLen))
}

func DShowMessage(s string) {
	defSyscallN(lcl.DSHOWMESSAGE, PascalStr(s))
}

func DMessageDlg(Msg string, DlgType TMsgDlgType, Buttons TMsgDlgButtons, HelpCtx int32) int32 {
	return int32(defSyscallN(lcl.DMESSAGEDLG, PascalStr(Msg), uintptr(DlgType), uintptr(Buttons), uintptr(HelpCtx)))
}

func DTextToShortCut(val string) TShortCut {
	return TShortCut(defSyscallN(lcl.DTEXTTOSHORTCUT, PascalStr(val)))
}

func DShortCutToText(val TShortCut) string {
	return GoStr(defSyscallN(lcl.DSHORTCUTTOTEXT, uintptr(val)))
}

func DSysOpen(filename string) {
	defSyscallN(lcl.DSYSOPEN, PascalStr(filename))
}

func DExtractFilePath(filename string) string {
	return GoStr(defSyscallN(lcl.DEXTRACTFILEPATH, PascalStr(filename)))
}

func DFileExists(filename string) bool {
	return GoBool(defSyscallN(lcl.DFILEEXISTS, PascalStr(filename)))
}

func DSelectDirectory1(options TSelectDirOpts) (bool, string) {
	var ptr uintptr
	v := GoBool(defSyscallN(lcl.DSELECTDIRECTORY1, uintptr(unsafe.Pointer(&ptr)), uintptr(options), 0))
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func DSelectDirectory2(caption, root string, showHidden bool) (bool, string) {
	var ptr uintptr
	v := GoBool(defSyscallN(lcl.DSELECTDIRECTORY2, PascalStr(caption), PascalStr(root), PascalBool(showHidden), uintptr(unsafe.Pointer(&ptr))))
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func SetThreadAsyncCallback(ptr uintptr) {
	defSyscallN(lcl.SetThreadAsyncCallback, ptr)
}

func SetThreadSyncCallback(ptr uintptr) {
	defSyscallN(lcl.SetThreadSyncCallback, ptr)
}

func DRunMainAsyncCall(id uintptr) {
	defSyscallN(lcl.DRunMainAsyncCall, id)
}

func DRunMainSyncCall(fn func()) {
	threadSyncRef.Lock()
	defer threadSyncRef.Unlock()
	threadSyncProc = fn
	defSyscallN(lcl.DRunMainSyncCall)
	threadSyncProc = nil
}

func DInputBox(aCaption, aPrompt, aDefault string) string {
	return GoStr(defSyscallN(lcl.DINPUTBOX, PascalStr(aCaption), PascalStr(aPrompt), PascalStr(aDefault)))
}

func DInputQuery(aCaption, aPrompt string, value *string) bool {
	if value == nil {
		return false
	}
	var strPtr uintptr
	r := defSyscallN(lcl.DINPUTQUERY, PascalStr(aCaption), PascalStr(aPrompt), PascalStr(*value), uintptr(unsafe.Pointer(&strPtr)))
	if strPtr != 0 {
		*value = GoStr(strPtr)
	}
	return GoBool(r)
}

func DPasswordBox(aCaption, aPrompt string) string {
	return GoStr(defSyscallN(lcl.DPASSWORDBOX, PascalStr(aCaption), PascalStr(aPrompt)))
}

func DInputCombo(aCaption, aPrompt string, aList uintptr) int32 {
	return int32(defSyscallN(lcl.DINPUTCOMBO, PascalStr(aCaption), PascalStr(aPrompt), aList))
}

func DInputComboEx(aCaption, aPrompt string, aList uintptr, allowCustomText bool) string {
	return GoStr(defSyscallN(lcl.DINPUTCOMBOEX, PascalStr(aCaption), PascalStr(aPrompt), aList, PascalBool(allowCustomText)))
}

func DSysLocale(aInfo *TSysLocale) {
	defSyscallN(lcl.DSYSLOCALE, uintptr(unsafe.Pointer(aInfo)))
}

func DCreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	defSyscallN(lcl.DCREATEURLSHORTCUT, PascalStr(aDestPath), PascalStr(aShortCutName), PascalStr(aURL))
}

func DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	return GoBool(defSyscallN(lcl.DCREATESHORTCUT, PascalStr(aDestPath), PascalStr(aShortCutName), PascalStr(aSrcFileName),
		PascalStr(aIconFileName), PascalStr(aDescription), PascalStr(aCmdArgs)))
}

func DSetPropertyValue(instance uintptr, propName, value string) {
	defSyscallN(lcl.DSETPROPERTYVALUE, instance, PascalStr(propName), PascalStr(value))
}

func DSetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	defSyscallN(lcl.DSETPROPERTYSECVALUE, instance, PascalStr(propName), PascalStr(secPropName), PascalStr(value))
}

func DGUIDToString(guid TGUID) string {
	return GoStr(defSyscallN(lcl.DGUIDTOSTRING, uintptr(unsafe.Pointer(&guid))))
}

func DStringToGUID(str string) (guid TGUID) {
	defSyscallN(lcl.DSTRINGTOGUID, PascalStr(str), uintptr(unsafe.Pointer(&guid)))
	return
}

func DCreateGUID() (guid TGUID) {
	defSyscallN(lcl.DCREATEGUID, uintptr(unsafe.Pointer(&guid)))
	return
}

func DGetLibResourceCount() int32 {
	return int32(defSyscallN(lcl.DGETLIBRESOURCECOUNT))
}

func DGetLibResourceItem(aIndex int32) (ret TLibResource) {
	item := struct {
		Name     uintptr
		ValuePtr uintptr
	}{}
	defSyscallN(lcl.DGETLIBRESOURCEITEM, uintptr(aIndex), uintptr(unsafe.Pointer(&item)))
	ret.Name = GoStr(item.Name)
	ret.Ptr = item.ValuePtr
	return
}

func DModifyLibResource(aPtr uintptr, aValue string) {
	defSyscallN(lcl.DMODIFYLIBRESOURCE, aPtr, PascalStr(aValue))
}

func DLibStringEncoding() TStringEncoding {
	return TStringEncoding(defSyscallN(lcl.DLIBSTRINGENCODING))
}

func DLibVersion() uint32 {
	return uint32(defSyscallN(lcl.DLIBVERSION))
}

func DLibAbout() string {
	return GoStr(defSyscallN(lcl.DLIBABOUT))
}

func DMainThreadId() uintptr {
	return defSyscallN(lcl.DMAINTHREADID)
}

func DCurrentThreadId() uintptr {
	return defSyscallN(lcl.DCURRENTTHREADID)
}

func DInitGoDll(aMainThreadId uintptr) {
	defSyscallN(lcl.DINITGODLL, aMainThreadId)
}

func DFindControl(handle HWND) uintptr {
	return defSyscallN(lcl.DFINDCONTROL, handle)
}

func DFindLCLControl(screenPos TPoint) uintptr {
	return defSyscallN(lcl.DFINDLCLCONTROL, uintptr(unsafe.Pointer(&screenPos)))
}

func DFindOwnerControl(handle HWND) uintptr {
	return defSyscallN(lcl.DFINDOWNERCONTROL, handle)
}

func DFindControlAtPosition(position TPoint, allowDisabled bool) uintptr {
	return defSyscallN(lcl.DFINDCONTROLATPOSITION, uintptr(unsafe.Pointer(&position)), PascalBool(allowDisabled))
}

func DFindLCLWindow(screenPos TPoint, allowDisabled bool) uintptr {
	return defSyscallN(lcl.DFINDLCLWINDOW, uintptr(unsafe.Pointer(&screenPos)), PascalBool(allowDisabled))
}

func DFindDragTarget(position TPoint, allowDisabled bool) uintptr {
	return defSyscallN(lcl.DFINDDRAGTARGET, uintptr(unsafe.Pointer(&position)), PascalBool(allowDisabled))
}

func DFreeAndNil(obj uintptr) {
	defSyscallN(lcl.DFreeAndNil, obj)
}

func DDateTimeToUnix(dateTime float64) (result int64) {
	defSyscallN(lcl.DToUnixTime, uintptr(unsafe.Pointer(&dateTime)), uintptr(unsafe.Pointer(&result)))
	return
}

func DUnixToDateTime(dateTime int64) (result float64) {
	defSyscallN(lcl.DUnixToTime, uintptr(unsafe.Pointer(&dateTime)), uintptr(unsafe.Pointer(&result)))
	return
}

func SetCEFEventCallback(ptr uintptr) {
	defSyscallN(lcl.SetCEFEventCallback, ptr)
}

func SetCEFRemoveEventCallback(ptr uintptr) {
	defSyscallN(lcl.SetCEFRemoveEventCallback, ptr)
}

func SetWVEventCallback(ptr uintptr) {
	defSyscallN(lcl.SetWVEventCallback, ptr)
}

func SetWVRemoveEventCallback(ptr uintptr) {
	defSyscallN(lcl.SetWVRemoveEventCallback, ptr)
}

// 以下是直接返回 api 索引

func COMSequentialStreamRead() int {
	return lcl.COMSequentialStream_Read
}
func COMSequentialStreamWrite() int {
	return lcl.COMSequentialStream_Write
}
func COMStreamSeek() int {
	return lcl.COMStream_Seek
}
func COMStreamSetSize() int {
	return lcl.COMStream_SetSize
}
func COMStreamCopyTo() int {
	return lcl.COMStream_CopyTo
}
func COMStreamCommit() int {
	return lcl.COMStream_Commit
}
func COMStreamRevert() int {
	return lcl.COMStream_Revert
}
func COMStreamLockRegion() int {
	return lcl.COMStream_LockRegion
}
func COMStreamUnlockRegion() int {
	return lcl.COMStream_UnlockRegion
}
func COMStreamStat() int {
	return lcl.COMStream_Stat
}
func COMStreamClone() int {
	return lcl.COMStream_Clone
}
func StreamAdapterClone() int {
	return lcl.StreamAdapter_Clone
}
func StreamAdapterCommit() int {
	return lcl.StreamAdapter_Commit
}
func StreamAdapterCopyTo() int {
	return lcl.StreamAdapter_CopyTo
}
func StreamAdapterCreate() int {
	return lcl.StreamAdapter_Create
}
func StreamAdapterLockRegion() int {
	return lcl.StreamAdapter_LockRegion
}
func StreamAdapterRead() int {
	return lcl.StreamAdapter_Read
}
func StreamAdapterRevert() int {
	return lcl.StreamAdapter_Revert
}
func StreamAdapterSeek() int {
	return lcl.StreamAdapter_Seek
}
func StreamAdapterSetSize() int {
	return lcl.StreamAdapter_SetSize
}
func StreamAdapterStat() int {
	return lcl.StreamAdapter_Stat
}
func StreamAdapterStream() int {
	return lcl.StreamAdapter_Stream
}
func StreamAdapterStreamOwnership() int {
	return lcl.StreamAdapter_StreamOwnership
}
func StreamAdapterUnlockRegion() int {
	return lcl.StreamAdapter_UnlockRegion
}
func StreamAdapterWrite() int {
	return lcl.StreamAdapter_Write
}
func InterfacedObjectCreate() int {
	return lcl.InterfacedObject_Create
}
func InterfacedObjectRefCount() int {
	return lcl.InterfacedObject_RefCount
}

func VirtualTreeColumnsGetVisibleColumns() int {
	return lcl.VirtualTreeColumns_GetVisibleColumns
}

func VTImageInfoGet() int {
	return lcl.VTImageInfo_Get
}

func VirtualNodeCreate() int {
	return lcl.VirtualNode_Create
}

func VirtualNodeIndex() int {
	return lcl.VirtualNode_Index
}
func VirtualNodeChildCount() int {
	return lcl.VirtualNode_ChildCount
}
func VirtualNodeNodeHeight() int {
	return lcl.VirtualNode_NodeHeight
}
func VirtualNodeStates() int {
	return lcl.VirtualNode_States
}
func VirtualNodeAlign() int {
	return lcl.VirtualNode_Align
}
func VirtualNodeCheckState() int {
	return lcl.VirtualNode_CheckState
}
func VirtualNodeCheckType() int {
	return lcl.VirtualNode_CheckType
}
func VirtualNodeDummy() int {
	return lcl.VirtualNode_Dummy
}
func VirtualNodeTotalCount() int {
	return lcl.VirtualNode_TotalCount
}
func VirtualNodeTotalHeight() int {
	return lcl.VirtualNode_TotalHeight
}
func VirtualNodeParent() int {
	return lcl.VirtualNode_Parent
}
func VirtualNodePrevSibling() int {
	return lcl.VirtualNode_PrevSibling
}
func VirtualNodeNextSibling() int {
	return lcl.VirtualNode_NextSibling
}
func VirtualNodeFirstChild() int {
	return lcl.VirtualNode_FirstChild
}
func VirtualNodeLastChild() int {
	return lcl.VirtualNode_LastChild
}
func VirtualNodeData() int {
	return lcl.VirtualNode_Data
}
