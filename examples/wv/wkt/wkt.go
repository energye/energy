package main

import (
	"fmt"
	"github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/api/imports"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"unsafe"
)

type TMainForm struct {
	lcl.TForm
	webview *TWebview
}

var MainForm TMainForm

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetScaled(true)
	lcl.Application.CreateForm(&MainForm)
	lcl.Application.Run()
	// sudo apt-get update

	// gtk2 webkit
	// sudo apt-get install libwebkitgtk-1.0-0

	// gtk3 webkit
	// sudo apt-get install libwebkitgtk-3.0-0

	// gtk3 webkit2
	// sudo apt-get install libwebkit2gtk-4.0-37
	// sudo apt-get install libwebkit2gtk-4.0-0 ??

	// sudo apt-get install libwebkitgtk-3.0-dev

	// ldconfig -p | grep libwebkitgtk-1.0.so

	// 解决 WebVTT 编码器问题
	// sudo apt install gstreamer1.0-plugins-bad

	// sudo find / -name libwebkit2gtk
	// apt-cache search libwebkit2gtk-4.0
}

func (m *TMainForm) FormCreate(sender lcl.IObject) {
	fmt.Println("main create")
	m.SetCaption("Main")
	m.SetWidth(100)
	m.SetHeight(100)
	btn := lcl.NewButton(m)
	btn.SetParent(m)
	btn.SetCaption("test")
	btn.SetOnClick(func(sender lcl.IObject) {
		fmt.Println("SetOnClick")
	})
	m.SetOnShow(func(sender lcl.IObject) {
		m.SetWidth(1024)
		m.SetHeight(600)
		m.SetPosition(types.PoScreenCenter)
	})

	mainMenu := lcl.NewMainMenu(m)
	item := lcl.NewMenuItem(m)
	item.SetCaption("文件(&F)")
	mainMenu.Items().Add(item)

	//m.webview = NewWebview(m)
	//m.webview.SetParent(m)
	//m.webview.SetAlign(types.AlClient)
	//url := "https://www.baidu.com"
	//url = "https://gitee.com/energye/energy"
	//url = "https://energy.yanghy.cn/"
	//url = "https://vfx.mtime.cn/Video/2019/03/21/mp4/190321153853126488.mp4"
	//m.SetOnShow(func(sender lcl.IObject) {
	//	m.webview.Navigate(url)
	//})
}

func (m *TMainForm) CreateParams(params *types.TCreateParams) {
	fmt.Println("调用此过程  TMainForm.CreateParams:", *params)
	//mainMenu := lcl.NewMainMenu(m)
	//item := lcl.NewMenuItem(m)
	//item.SetCaption("文件(&F)")
	//mainMenu.Items().Add(item)
}

var webkitImport = new(imports.Imports) // 自定义组件初始化导入

func init() {
	webkitImports := []*imports.Table{
		imports.NewTable("Webview_Create", 0),
		imports.NewTable("Webview_Free", 0),
		imports.NewTable("Webview_Navigate", 0),
		imports.NewTable("Webview_GoBack", 0),
		imports.NewTable("Webview_GoForward", 0),
		imports.NewTable("Webview_GoHome", 0),
		imports.NewTable("Webview_GoSearch", 0),
		imports.NewTable("Webview_Refresh", 0),
		imports.NewTable("Webview_Stop", 0),
		imports.NewTable("Webview_SetBounds", 0),
		imports.NewTable("Webview_ExecuteScript", 0),
		imports.NewTable("Webview_ExecuteJS", 0),
		imports.NewTable("Webview_LoadHTML", 0),
		imports.NewTable("Webview_CanFocus", 0),
		imports.NewTable("Webview_SetParent", 0),
		imports.NewTable("Webview_SetAlign", 0),
	}
	webkitImport.SetImportTable(webkitImports)
	webkitImport.SetOk(true)
}

type TWebview struct {
	lcl.ICustomControl
	instance unsafe.Pointer
}

func NewWebview(owner lcl.IComponent) *TWebview {
	m := new(TWebview)
	m.instance = unsafe.Pointer(Webview_Create(owner.Instance()))
	return m
}

func (m *TWebview) Free() {
	if m.instance != nil {
		Webview_Free(m._instance())
		m.instance = nil
	}
}

func (m *TWebview) _instance() uintptr {
	return uintptr(m.instance)
}

func (m *TWebview) Instance() uintptr {
	return m._instance()
}

func (m *TWebview) UnsafeAddr() unsafe.Pointer {
	return m.instance
}

func (m *TWebview) IsValid() bool {
	return m.instance != nil
}

func (m *TWebview) Navigate(AURL string) {
	Webview_Navigate(m._instance(), AURL)
}

func (m *TWebview) GoBack() {
	Webview_GoBack(m._instance())
}

func (m *TWebview) GoForward() {
	Webview_GoForward(m._instance())
}

func (m *TWebview) GoHome() {
	Webview_GoHome(m._instance())
}

func (m *TWebview) GoSearch() {
	Webview_GoSearch(m._instance())
}

func (m *TWebview) Refresh() {
	Webview_Refresh(m._instance())
}

func (m *TWebview) Stop() {
	Webview_Stop(m._instance())
}

func (m *TWebview) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	Webview_SetBounds(m._instance(), ALeft, ATop, AWidth, AHeight)
}

func (m *TWebview) ExecuteScript(AScriptText string, AScriptType string) string {
	return Webview_ExecuteScript(m._instance(), AScriptText, AScriptType)
}

func (m *TWebview) ExecuteJS(AScriptText string) string {
	return Webview_ExecuteJS(m._instance(), AScriptText)
}

func (m *TWebview) LoadHTML(AStr string) {
	Webview_LoadHTML(m._instance(), AStr)
}

func syscallN(trap int, args ...uintptr) uintptr {
	return webkitImport.SysCallN(trap, args...)
}

func Webview_Create(obj uintptr) uintptr {
	return syscallN(0, obj)
}

func Webview_Free(obj uintptr) {
	syscallN(1, obj)
}

func Webview_Navigate(obj uintptr, AURL string) {
	syscallN(2, obj, api.PascalStr(AURL))
}

func Webview_GoBack(obj uintptr) {
	syscallN(3, obj)
}

func Webview_GoForward(obj uintptr) {
	syscallN(4, obj)
}

func Webview_GoHome(obj uintptr) {
	syscallN(5, obj)
}

func Webview_GoSearch(obj uintptr) {
	syscallN(6, obj)
}

func Webview_Refresh(obj uintptr) {
	syscallN(7, obj)
}

func Webview_Stop(obj uintptr) {
	syscallN(8, obj)
}

func Webview_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	syscallN(9, obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func Webview_ExecuteScript(obj uintptr, AScriptText string, AScriptType string) string {
	return api.GoStr(syscallN(10, obj, api.PascalStr(AScriptText), api.PascalStr(AScriptType)))
}

func Webview_ExecuteJS(obj uintptr, AScriptText string) string {
	return api.GoStr(syscallN(11, obj, api.PascalStr(AScriptText)))
}

func Webview_LoadHTML(obj uintptr, AStr string) {
	syscallN(12, obj, api.PascalStr(AStr))
}

func Webview_CanFocus(obj uintptr) bool {
	return api.GoBool(syscallN(13, obj))
}

func (m *TWebview) SetParent(value lcl.IWinControl) {
	syscallN(14, m.Instance(), value.Instance())
}

func (m *TWebview) SetAlign(value types.TAlign) {
	syscallN(15, m.Instance(), uintptr(value))
}
