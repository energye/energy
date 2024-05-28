package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/api"
	"github.com/energye/energy/v2/api/libname"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/pkgs/assetserve"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/colors"
	"github.com/energye/energy/v2/types/messages"
	"github.com/energye/energy/v2/wv"
	"path/filepath"
	"unsafe"
)

var mainForm TMainForm
var load wv.IWVLoader

//go:embed resources
var resources embed.FS

func main() {
	fmt.Println("Go ENERGY Run Main")
	wv.Init(nil, nil)
	// GlobalWebView2Loader
	load = wv.GlobalWebView2Loader(nil)
	liblcl := libname.LibName
	webView2Loader, _ := filepath.Split(liblcl)
	webView2Loader = filepath.Join(webView2Loader, "WebView2Loader.dll")
	fmt.Println("当前目录:", types.CurrentExecuteDir)
	fmt.Println("liblcl.dll目录:", liblcl)
	fmt.Println("WebView2Loader.dll目录:", webView2Loader)
	fmt.Println("用户缓存目录:", filepath.Join(types.CurrentExecuteDir, "EnergyCache"))
	load.SetUserDataFolder(filepath.Join(types.CurrentExecuteDir, "EnergyCache"))
	load.SetLoaderDllPath(webView2Loader)
	r := load.StartWebView2()
	fmt.Println("StartWebView2", r)
	// 启动 内置http server
	startHttpServer()
	// 底层库全局异常
	lcl.Application.SetOnException(func(sender lcl.IObject, e lcl.IException) {
		fmt.Println("底层库异常:", e.ToString())
	})
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&mainForm)
	lcl.Application.Run()
}

type TMainForm struct {
	lcl.TForm
	windowParent         wv.IWVWindowParent
	browser              wv.IWVBrowser
	messageEdit          lcl.IEdit
	messageBtn           lcl.IButton
	postMessageMem       lcl.IMemo
	postSharedBufferEdit lcl.IEdit
	postSharedBufferBtn  lcl.IButton
	sharedBuffer         wv.ICoreWebView2SharedBuffer
	menuExitMemory       lcl.IMemoryStream
	enableContextMenuChk bool
}

var CUSTOM_SHARED_BUFFER_SIZE int64 = 1024

func (m *TMainForm) createRightBoxLayout() {
	messagePanel := lcl.NewPanel(m)
	messagePanel.SetParent(m)
	messagePanel.SetParentDoubleBuffered(true)
	messagePanel.SetLeft(m.Width() / 2)
	messagePanel.SetHeight(m.Height())
	messagePanel.SetWidth(m.Width() / 2)
	messagePanel.SetBorderStyle(types.BsNone)
	messagePanel.SetBorderWidth(1)
	messagePanel.SetColor(colors.ClWhite)
	messagePanel.SetAnchors(types.NewSet(types.AkTop, types.AkRight, types.AkBottom))

	uiLbl := lcl.NewLabel(messagePanel)
	uiLbl.SetParent(messagePanel)
	uiLbl.SetLeft(5)
	uiLbl.SetTop(5)
	uiLbl.SetCaption("Webview2 此部分客户端UI组件")

	// send message
	m.messageEdit = lcl.NewEdit(messagePanel)
	m.messageEdit.SetParent(messagePanel)
	m.messageEdit.SetWidth(messagePanel.Width() - 10)
	m.messageEdit.SetTop(35)
	m.messageEdit.SetLeft(5)
	m.messageEdit.SetText("你好，来自主机应用程序！")

	m.messageBtn = lcl.NewButton(messagePanel)
	m.messageBtn.SetParent(messagePanel)
	m.messageBtn.SetWidth(messagePanel.Width() - 10)
	m.messageBtn.SetTop(m.messageEdit.Top() + 25)
	m.messageBtn.SetLeft(5)
	m.messageBtn.SetCaption("<<<向web浏览器发送消息")
	m.messageBtn.SetOnClick(func(sender lcl.IObject) {
		m.browser.PostWebMessageAsString(m.messageEdit.Text())
	})

	msgLal := lcl.NewLabel(messagePanel)
	msgLal.SetParent(messagePanel)
	msgLal.SetCaption("来自web浏览器的消息：")
	msgLal.SetTop(m.messageBtn.Top() + 35)
	msgLal.SetLeft(5)

	m.postMessageMem = lcl.NewMemo(messagePanel)
	m.postMessageMem.SetParent(messagePanel)
	m.postMessageMem.SetTop(msgLal.Top() + 35)
	m.postMessageMem.SetLeft(5)
	m.postMessageMem.SetHeight(200)
	m.postMessageMem.SetWidth(messagePanel.Width() - 10)
	m.postMessageMem.SetReadOnly(true)

	m.postSharedBufferEdit = lcl.NewEdit(messagePanel)
	m.postSharedBufferEdit.SetParent(messagePanel)
	m.postSharedBufferEdit.SetWidth(messagePanel.Width() - 10)
	m.postSharedBufferEdit.SetTop(m.postMessageMem.Top() + m.postMessageMem.Height() + 10)
	m.postSharedBufferEdit.SetLeft(5)
	m.postSharedBufferEdit.SetText("共享缓冲区自定义内容")

	m.postSharedBufferBtn = lcl.NewButton(messagePanel)
	m.postSharedBufferBtn.SetParent(messagePanel)
	m.postSharedBufferBtn.SetTop(m.postSharedBufferEdit.Top() + m.postSharedBufferEdit.Height() + 10)
	m.postSharedBufferBtn.SetLeft(5)
	m.postSharedBufferBtn.SetWidth(messagePanel.Width() - 10)
	m.postSharedBufferBtn.SetCaption("<<<将共享缓冲区发布到web浏览器")
	// 仅存放指针, 未实现
	var tempSharedBufferItf wv.ICoreWebView2SharedBuffer
	m.postSharedBufferBtn.SetOnClick(func(sender lcl.IObject) {
		if m.sharedBuffer == nil {
			if m.browser.CreateSharedBuffer(CUSTOM_SHARED_BUFFER_SIZE, &tempSharedBufferItf) {
				// 创建buffer实现对象
				m.sharedBuffer = wv.NewCoreWebView2SharedBuffer(tempSharedBufferItf)
			}
		}
		text := m.postSharedBufferEdit.Text()
		// 清空缓冲区, 填充0
		m.fullChar(m.sharedBuffer.Buffer(), uintptr(CUSTOM_SHARED_BUFFER_SIZE), 0)
		// 将新数据放入缓冲区
		api.DMove(api.PascalStr(text), m.sharedBuffer.Buffer(), len(text))
		// 发送缓冲数据到页面监听事件 "sharedbufferreceived"
		m.browser.PostSharedBufferToScript(tempSharedBufferItf, wv.COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_WRITE, "")
	})

	enableContextMenuChk := lcl.NewCheckBox(messagePanel)
	enableContextMenuChk.SetParent(messagePanel)
	enableContextMenuChk.SetLeft(5)
	enableContextMenuChk.SetTop(m.postSharedBufferBtn.Top() + m.postSharedBufferBtn.Height() + 10)
	enableContextMenuChk.SetCaption("启用菜单")
	enableContextMenuChk.SetChecked(true)
	enableContextMenuChk.SetOnChange(func(sender lcl.IObject) {
		m.enableContextMenuChk = enableContextMenuChk.Checked()
		if m.enableContextMenuChk {
			enableContextMenuChk.SetCaption("启用菜单")
		} else {
			enableContextMenuChk.SetCaption("禁用菜单")
		}
	})
	m.enableContextMenuChk = enableContextMenuChk.Checked()
}

// 填充
func (m *TMainForm) fullChar(buf types.PByte, size, value uintptr) {
	for i := 0; i < int(size); i++ {
		*(*uintptr)(unsafe.Pointer(buf + uintptr(i))) = value
	}
}

func (m *TMainForm) FormCreate(sender lcl.IObject) {
	m.SetCaption("Webview2 浏览器主机应用程序通信")
	m.SetPosition(types.PoScreenCenter)
	m.SetWidth(1024)
	m.SetHeight(768)
	m.SetDoubleBuffered(true)
	m.createRightBoxLayout()

	m.windowParent = wv.NewWVWindowParent(m)
	m.windowParent.SetParent(m)
	//重新调整browser窗口的Parent属性
	//重新设置了上边距，宽，高
	m.windowParent.SetAlign(types.AlCustom) //重置对齐,默认是整个客户端
	m.windowParent.SetHeight(m.Height())
	m.windowParent.SetWidth(m.Width() / 2)
	m.windowParent.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))

	m.browser = wv.NewWVBrowser(m)
	m.browser.SetDefaultURL("http://localhost:22022/index.html")
	m.browser.SetTargetCompatibleBrowserVersion("95.0.1020.44") // 设置
	fmt.Println("TargetCompatibleBrowserVersion:", m.browser.TargetCompatibleBrowserVersion())
	m.browser.SetOnAfterCreated(func(sender lcl.IObject) {
		fmt.Println("回调函数 WVBrowser => SetOnAfterCreated")
		m.windowParent.UpdateSize()
	})
	m.browser.SetOnDocumentTitleChanged(func(sender lcl.IObject) {
		fmt.Println("回调函数 WVBrowser => SetOnDocumentTitleChanged:", m.browser.DocumentTitle())
	})
	m.browser.SetOnWebMessageReceived(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2WebMessageReceivedEventArgs) {
		fmt.Println("回调函数 WVBrowser => SetOnWebMessageReceived")
		args = wv.NewCoreWebView2WebMessageReceivedEventArgs(args)
		messageData := args.WebMessageAsString()
		m.postMessageMem.Lines().Add(messageData)
		if m.sharedBuffer != nil && messageData == "SharedBufferDataUpdated" {
			buf := m.sharedBuffer.Buffer()
			bufData := api.GoStr(buf)
			fmt.Println("bufData:", bufData, "SIZE:", len(bufData))
			m.postMessageMem.Lines().Add("新缓冲区内容: " + bufData)
		}
		// free
		args.Free()
	})
	// 右键菜单图标
	menuExit, _ := resources.ReadFile("resources/menu_exit.png")
	m.menuExitMemory = lcl.NewMemoryStream()
	m.menuExitMemory.LoadFromBytes(menuExit)
	menuExitStreamAdapter := lcl.NewStreamAdapter(m.menuExitMemory, types.SoOwned)
	// 右键菜单退出项ID
	var exitItemId int32
	m.browser.SetOnContextMenuRequested(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2ContextMenuRequestedEventArgs) {
		var tmpMenuItemPtr wv.ICoreWebView2ContextMenuItem
		tmpArgs := wv.NewCoreWebView2ContextMenuRequestedEventArgs(args)
		menuItems := tmpArgs.MenuItems()
		tmpCollection := wv.NewCoreWebView2ContextMenuItemCollection(menuItems)
		if !m.enableContextMenuChk {
			tmpCollection.RemoveAllMenuItems()
			return
		}
		if m.browser.CoreWebView2Environment().CreateContextMenuItem("Exit", menuExitStreamAdapter, wv.COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_COMMAND, &tmpMenuItemPtr) {
			tmpMenuItem := wv.NewCoreWebView2ContextMenuItem(tmpMenuItemPtr)
			exitItemId = tmpMenuItem.CommandId()
			// 设置菜单事件触发对象为delegateEvents, 触发 SetOnCustomItemSelected 事件
			tmpMenuItem.AddAllBrowserEvents(m.browser)
			tmpCollection.InsertValueAtIndex(tmpCollection.Count(), tmpMenuItemPtr)
		}
		// free
		tmpCollection.Free()
		tmpArgs.Free()
	})
	// 自定义菜单项选择事件回调
	m.browser.SetOnCustomItemSelected(func(sender wv.IObject, menuItem wv.ICoreWebView2ContextMenuItem) {
		menuItem = wv.NewCoreWebView2ContextMenuItem(menuItem)
		fmt.Println("SetOnCustomItemSelected", menuItem.CommandId())
		if exitItemId == menuItem.CommandId() {
			m.Close()
		}
		// free
		menuItem.Free()
	})
	m.browser.SetOnNewWindowRequested(func(sender wv.IObject, webView wv.ICoreWebView2, args wv.ICoreWebView2NewWindowRequestedEventArgs) {
		args = wv.NewCoreWebView2NewWindowRequestedEventArgs(args)
		// 阻止新窗口
		args.SetHandled(true)
		// 可以自己创建窗口

		// 当前页面打开新链接
		//m.browser.Navigate(args.URI())
		//free
		args.Free()
	})
	// 设置browser到window parent
	m.windowParent.SetBrowser(m.browser)

	// 窗口显示时创建browser
	m.SetOnShow(func(sender lcl.IObject) {
		if load.InitializationError() {
			fmt.Println("回调函数 => SetOnShow 初始化失败")
		} else {
			if load.Initialized() {
				fmt.Println("回调函数 => SetOnShow 初始化成功")
				m.browser.CreateBrowser(m.windowParent.Handle(), true)
			}
		}
	})
	m.SetOnWndProc(func(msg *types.TMessage) {
		m.InheritedWndProc(msg)
		switch msg.Msg {
		case messages.WM_SIZE, messages.WM_MOVE:
			m.browser.NotifyParentWindowPositionChanged()
		}
	})
	m.SetOnDestroy(m.OnFormDestroy)
}

func (m *TMainForm) OnFormDestroy(sender lcl.IObject) {
	fmt.Println("OnFormDestroy")
}

func startHttpServer() {
	//通过内置http服务加载资源
	server := assetserve.NewAssetsHttpServer()
	server.PORT = 22022               //服务端口号
	server.AssetsFSName = "resources" //必须设置目录名和资源文件夹同名
	server.Assets = resources
	go server.StartHttpServer()
}
