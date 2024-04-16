package layout

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"os"
	"path/filepath"
	"strconv"
)

type MainWindowLayout struct {
	extensionChromium cef.ICEFChromiumBrowser
	fExtension        *cef.ICefExtension
	addressPnl        *lcl.TPanel
	addressEdit       *lcl.TEdit
	addressBtn        *lcl.TButton
	mainPnl           *lcl.TPanel
	extensionsPnl     *lcl.TPanel
	loadExtensBtn     *lcl.TButton
	loadPopupBtn      *lcl.TButton
	unloadExtensBtn   *lcl.TButton
	extensionMem      *lcl.TMemo
}

func (m *MainWindowLayout) updateExtension() {
	if m.extensionChromium.Chromium().Initialized() {
		m.loadExtensBtn.SetEnabled(false)
		m.loadPopupBtn.SetEnabled(false)
		m.unloadExtensBtn.SetEnabled(m.fExtension.IsValid())
	} else {
		m.loadExtensBtn.SetEnabled(!m.fExtension.IsValid())
		m.loadPopupBtn.SetEnabled(m.fExtension.IsValid())
		m.unloadExtensBtn.SetEnabled(m.fExtension.IsValid())
	}
}

func WindowLayout(window cef.IBrowserWindow) {
	mwl := &MainWindowLayout{}
	// 主窗口
	bw := window.AsLCLBrowserWindow().BrowserWindow()
	bw.SetWidth(1550)

	// 区域一、地址栏panel
	mwl.createAddress(bw)
	// 区域二、主页面panel
	mwl.createMainBrowser(bw)

	// 区域三、右侧扩展功能panel
	mwl.createExtension(bw)

	// 最后在初始化更新一次按钮状态
	mwl.updateExtension()

	bw.SetOnCloseQuery(func(sender lcl.IObject, canClose *bool) bool {
		if bw.IsClosing() {
			return false
		}
		if mwl.extensionChromium != nil {
			mwl.extensionChromium.Chromium().CloseBrowser(true)
		}
		return false
	})
}

func (mwl *MainWindowLayout) createAddress(bw *cef.LCLBrowserWindow) {
	mwl.addressPnl = lcl.NewPanel(bw)
	mwl.addressPnl.SetParent(bw)
	mwl.addressPnl.SetWidth(bw.Width())
	mwl.addressPnl.SetHeight(35)
	mwl.addressPnl.SetBevelOuter(types.BvNone)
	mwl.addressPnl.SetBevelInner(types.BvNone)
	mwl.addressPnl.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
	mwl.addressPnl.SetBorderStyle(types.BsNone)
	mwl.addressPnl.SetDoubleBuffered(true)
	// 地址栏功能组件
	// edit
	mwl.addressEdit = lcl.NewEdit(bw)
	mwl.addressEdit.SetParent(mwl.addressPnl)
	mwl.addressEdit.SetWidth(mwl.addressPnl.Width() - 80)
	mwl.addressEdit.SetText("https://gitee.com/energye/energy")
	mwl.addressEdit.SetAnchors(types.NewSet(types.AkLeft, types.AkRight))
	mwl.addressEdit.SetTop(5)
	mwl.addressEdit.SetDoubleBuffered(true)
	// button
	mwl.addressBtn = lcl.NewButton(bw)
	mwl.addressBtn.SetParent(mwl.addressPnl)
	mwl.addressBtn.SetLeft(mwl.addressEdit.Width() + 10)
	mwl.addressBtn.SetWidth(70)
	mwl.addressBtn.SetTop(5)
	mwl.addressBtn.SetCaption("GO")
	mwl.addressBtn.SetAnchors(types.NewSet(types.AkRight))
	mwl.addressBtn.SetDoubleBuffered(true)
	mwl.addressBtn.SetOnClick(func(sender lcl.IObject) {
		url := mwl.addressEdit.Text()
		if url != "" {
			bw.Chromium().LoadUrl(url)
			mwl.extensionChromium.Chromium().LoadUrl(url)
		}
	})
}

func (mwl *MainWindowLayout) createMainBrowser(bw *cef.LCLBrowserWindow) {
	mwl.mainPnl = lcl.NewPanel(bw)
	mwl.mainPnl.SetParent(bw)
	mwl.mainPnl.SetTop(mwl.addressPnl.Height())
	mwl.mainPnl.SetWidth(bw.Width() - 260)
	mwl.mainPnl.SetHeight(bw.Height() - mwl.addressPnl.Height())
	mwl.mainPnl.SetBevelOuter(types.BvNone)
	mwl.mainPnl.SetBevelInner(types.BvNone)
	mwl.mainPnl.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))
	mwl.mainPnl.SetBorderStyle(types.BsNone)
	mwl.mainPnl.SetDoubleBuffered(true)
	// 修改默认的主浏览器布局属性
	mainParent := bw.WindowParent()
	mainParent.RevertCustomAnchors()
	mainParent.SetParent(mwl.mainPnl)
	mainParent.SetAlign(types.AlClient)
}

func (mwl *MainWindowLayout) createExtension(bw *cef.LCLBrowserWindow) {
	wd, _ := os.Getwd()
	// 扩展功能组件
	var (
		extensionsPath = filepath.Join(wd, "examples", "extensions")
		hello          = filepath.Join(extensionsPath, "resources", "hello")
	)
	fmt.Println("extensionsPath:", extensionsPath)

	mwl.extensionsPnl = lcl.NewPanel(bw)
	mwl.extensionsPnl.SetParent(bw)
	mwl.extensionsPnl.SetTop(mwl.addressPnl.Height())
	mwl.extensionsPnl.SetLeft(mwl.mainPnl.Width())
	mwl.extensionsPnl.SetHeight(bw.Height() - mwl.addressPnl.Height())
	mwl.extensionsPnl.SetWidth(bw.Width() - mwl.mainPnl.Width())
	mwl.extensionsPnl.SetBevelOuter(types.BvNone)
	mwl.extensionsPnl.SetBevelInner(types.BvNone)
	mwl.extensionsPnl.SetAnchors(types.NewSet(types.AkTop, types.AkRight, types.AkBottom))
	mwl.extensionsPnl.SetBorderStyle(types.BsNone)
	mwl.extensionsPnl.SetDoubleBuffered(true)
	// 加载按钮
	mwl.loadExtensBtn = lcl.NewButton(bw)
	mwl.loadExtensBtn.SetParent(mwl.extensionsPnl)
	mwl.loadExtensBtn.SetWidth(mwl.extensionsPnl.Width())
	mwl.loadExtensBtn.SetCaption("1. 加载插件")
	mwl.loadExtensBtn.SetDoubleBuffered(true)
	mwl.loadExtensBtn.SetOnClick(func(sender lcl.IObject) {
		if mwl.fExtension == nil {
			fmt.Println("load:", hello)
			mwl.extensionChromium.Chromium().LoadExtension(hello, nil, nil, nil)
		}
	})
	// 加载弹出页面按钮
	mwl.loadPopupBtn = lcl.NewButton(bw)
	mwl.loadPopupBtn.SetParent(mwl.extensionsPnl)
	mwl.loadPopupBtn.SetWidth(mwl.extensionsPnl.Width())
	mwl.loadPopupBtn.SetTop(mwl.loadExtensBtn.Top() + mwl.loadExtensBtn.Height() + 10)
	mwl.loadPopupBtn.SetCaption("2. 加载弹出页面")
	mwl.loadPopupBtn.SetDoubleBuffered(true)
	mwl.loadPopupBtn.SetOnClick(func(sender lcl.IObject) {
		if !mwl.fExtension.IsValid() || mwl.extensionChromium.Chromium().Initialized() {
			fmt.Println("load popup exit")
			return
		}
		url := mwl.fExtension.GetURL() + mwl.fExtension.GetBrowserActionPopup()
		fmt.Println("loadURL:", url)
		mwl.extensionChromium.Chromium().SetDefaultURL(url)
		mwl.extensionChromium.CreateBrowser()
	})
	// 卸载扩展按钮
	mwl.unloadExtensBtn = lcl.NewButton(bw)
	mwl.unloadExtensBtn.SetParent(mwl.extensionsPnl)
	mwl.unloadExtensBtn.SetWidth(mwl.extensionsPnl.Width())
	mwl.unloadExtensBtn.SetTop(mwl.loadPopupBtn.Top() + mwl.loadPopupBtn.Height() + 10)
	mwl.unloadExtensBtn.SetCaption("3. 卸载扩展")
	mwl.unloadExtensBtn.SetDoubleBuffered(true)
	mwl.unloadExtensBtn.SetOnClick(func(sender lcl.IObject) {
		if mwl.fExtension.IsValid() {
			mwl.fExtension.Unload()
		}
	})
	// memo
	mwl.extensionMem = lcl.NewMemo(bw)
	mwl.extensionMem.SetParent(mwl.extensionsPnl)
	mwl.extensionMem.SetWidth(mwl.extensionsPnl.Width())
	mwl.extensionMem.SetTop(mwl.unloadExtensBtn.Top() + mwl.unloadExtensBtn.Height() + 10)
	mwl.extensionMem.SetHeight(230)
	mwl.extensionMem.SetDoubleBuffered(true)
	// 扩展 chromium
	mwl.extensionChromium = cef.NewChromiumBrowser(mwl.extensionsPnl, nil)
	//extensionChromium.RegisterDefaultEvent()
	extensParent := mwl.extensionChromium.WindowParent()
	extensParent.SetTop(mwl.extensionMem.Top() + mwl.extensionMem.Height() + 10)
	extensParent.SetWidth(mwl.extensionsPnl.Width())
	extensParent.SetHeight(mwl.extensionsPnl.Height() - (mwl.extensionMem.Top() + mwl.extensionMem.Height() + 10))
	extensParent.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight, types.AkBottom))

	// chromium 事件
	mwl.extensionEvent(bw)
}
func (mwl *MainWindowLayout) extensionEvent(bw *cef.LCLBrowserWindow) {
	mwl.extensionChromium.Chromium().SetOnBeforeBrowser(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, request *cef.ICefRequest, userGesture, isRedirect bool) bool {
		mwl.extensionChromium.WindowParent().UpdateSize()
		return false
	})
	mwl.extensionChromium.Chromium().SetOnClose(func(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
		// 当前只对windows的做了处理
		*aAction = consts.CbaDelay
		cef.RunOnMainThread(func() {
			mwl.extensionChromium.WindowParent().Free()
		})
	})
	// 扩展事件
	mwl.extensionChromium.Chromium().SetOnExtensionGetActiveBrowser(func(sender lcl.IObject, extension *cef.ICefExtension, browser *cef.ICefBrowser,
		includeIncognito bool, resultBrowser *cef.ICefBrowser) {
		fmt.Println("OnExtensionGetActiveBrowser:", includeIncognito, "id:", extension.GetIdentifier(), "Initialized:", mwl.extensionChromium.Chromium().Initialized())
		if mwl.extensionChromium.Chromium().Initialized() {
			//*resultBrowser = *mwl.extensionChromium.Chromium().Browser()
			*resultBrowser = *bw.Chromium().Browser()
		}
	})
	mwl.extensionChromium.Chromium().SetOnExtensionLoaded(func(sender lcl.IObject, extension *cef.ICefExtension) {
		fmt.Println("OnExtensionLoaded:", mwl.fExtension.GetIdentifier())
		// 保留引用
		mwl.fExtension = cef.ExtensionRef.UnWrap(extension)
		cef.RunOnMainThread(func() {
			// CEF_EXT_LOADED
			mwl.extensionMem.Lines().Add("--------------------------------")
			mwl.extensionMem.Lines().Add("Extension loaded successfully!")
			mwl.extensionMem.Lines().Add("Identifier: " + mwl.fExtension.GetIdentifier())
			mwl.extensionMem.Lines().Add("Path: " + mwl.fExtension.GetPath())
			mwl.extensionMem.Lines().Add("IsLoaded: " + strconv.FormatBool(mwl.fExtension.IsLoaded()))
			mwl.extensionMem.Lines().Add("Popup: " + mwl.fExtension.GetBrowserActionPopup())
			mwl.extensionMem.Lines().Add("Icon: " + mwl.fExtension.GetBrowserActionIcon())
			mwl.extensionMem.Lines().Add("URL: " + mwl.fExtension.GetURL())
			manifest := mwl.fExtension.GetManifest()
			keys := manifest.GetKeys()
			for i := 0; i < keys.Count(); i++ {
				typ := manifest.GetType(keys.Get(i))
				fmt.Println("manifset typ:", typ)
			}
			manifest.Free()
			mwl.updateExtension()
		})
	})

	mwl.extensionChromium.Chromium().SetOnExtensionLoadFailed(func(sender lcl.IObject, result consts.TCefErrorCode) {
		fmt.Println("OnExtensionLoadFailed:", result)
		cef.RunOnMainThread(func() {
			// CEF_EXT_ERROR
			mwl.extensionMem.Lines().Add("--------------------------------")
			mwl.extensionMem.Lines().Add("Extension load failed. Result : " + strconv.Itoa(int(result)))
			mwl.updateExtension()
		})
	})
	mwl.extensionChromium.Chromium().SetOnExtensionGetExtensionResource(func(sender lcl.IObject, extension *cef.ICefExtension, browser *cef.ICefBrowser, file string, callback *cef.ICefGetExtensionResourceCallback) bool {
		fmt.Println("OnExtensionGetExtensionResource:", file)
		return false
	})
	mwl.extensionChromium.Chromium().SetOnExtensionUnloaded(func(sender lcl.IObject, extension *cef.ICefExtension) {
		fmt.Println("OnExtensionUnloaded:", extension.GetIdentifier(), "IsSame:", extension.IsSame(mwl.fExtension))
		if extension.IsValid() && mwl.fExtension.IsValid() && extension.IsSame(mwl.fExtension) {
			cef.RunOnMainThread(func() {
				// CEF_EXT_UNLOADED
				mwl.fExtension.Free()
				mwl.fExtension = nil
				mwl.extensionMem.Lines().Add("--------------------------------")
				mwl.extensionMem.Lines().Add("Extension unloaded successfully!")
				mwl.updateExtension()
			})
		}
	})
	mwl.extensionChromium.Chromium().SetOnLoadEnd(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, httpStatusCode int32) {
		fmt.Println("OnLoadEnd:", httpStatusCode)
		cef.RunOnMainThread(func() {
			//CEF_EXT_POPUP_LOADED
		})
	})
	mwl.extensionChromium.Chromium().SetOnLoadError(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame,
		errorCode consts.CEF_NET_ERROR, errorText, failedUrl string) {
		fmt.Println("OnLoadError:", errorCode, errorText, failedUrl)
		//cef.RunOnMainThread(func() {
		// CEF_EXT_POPUP_ERROR
		//})
	})
}
