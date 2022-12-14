package cef

import (
	. "github.com/energye/energy/common"
	"github.com/energye/energy/logger"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api"
	"unsafe"
)

type TCEFBrowserViewComponent struct {
	instance unsafe.Pointer
}

func NewBrowserViewComponent(AOwner lcl.TComponent) *TCEFBrowserViewComponent {
	r1, _, _ := Proc(internale_CEFBrowserViewComponent_Create).Call(lcl.CheckPtr(AOwner))
	return &TCEFBrowserViewComponent{
		instance: unsafe.Pointer(r1),
	}
}

func (m *TCEFBrowserViewComponent) CreateBrowserView(client *ICefClient, url string, requestContextSettings *TCefRequestContextSettings, browserSettings *TCefBrowserSettings, extraInfo *ICefDictionaryValue) {
	contextSettingsPtr := requestContextSettings.ToPtr()
	browserSettingsPtr := browserSettings.ToPtr()
	var dataBytes = []byte{}
	var dataBytesPtr unsafe.Pointer
	var dataBytesLen int = 0
	var argsLen int = 0
	if extraInfo != nil && extraInfo.dataLen > 0 {
		defer func() {
			extraInfo.Clear()
			extraInfo = nil
			dataBytes = nil
			dataBytesPtr = nil
		}()
		dataBytes = extraInfo.Package()
		argsLen = extraInfo.dataLen
		dataBytesPtr = unsafe.Pointer(&dataBytes[0])
		dataBytesLen = len(dataBytes) - 1
	} else {
		dataBytesPtr = unsafe.Pointer(&dataBytes)
	}
	Proc(internale_CEFBrowserViewComponent_CreateBrowserView).Call(uintptr(m.instance), uintptr(client.instance), api.PascalStr(url), uintptr(unsafe.Pointer(&contextSettingsPtr)), uintptr(unsafe.Pointer(&browserSettingsPtr)), uintptr(argsLen), uintptr(dataBytesPtr), uintptr(dataBytesLen))
}

func (m *TCEFBrowserViewComponent) GetForBrowser(browser *ICefBrowser) {
	Proc(internale_CEFBrowserViewComponent_CreateBrowserView).Call(uintptr(m.instance), browser.Instance())
}

func (m *TCEFBrowserViewComponent) SetPreferAccelerators() {

}

func (m *TCEFBrowserViewComponent) RequestFocus() {

}

func (m *TCEFBrowserViewComponent) Browser() {

}

func (m *TCEFBrowserViewComponent) BrowserView() {

}

func (m *TCEFBrowserViewComponent) SetOnBrowserCreated() {

}

func (m *TCEFBrowserViewComponent) SetOnBrowserDestroyed() {

}

func (m *TCEFBrowserViewComponent) SetOnGetDelegateForPopupBrowserView() {

}

func (m *TCEFBrowserViewComponent) SetOnPopupBrowserViewCreated() {

}

func (m *TCEFBrowserViewComponent) SetOnGetChromeToolbarType() {

}

func init() {
	lcl.RegisterExtEventCallback(func(fn interface{}, getVal func(idx int) uintptr) bool {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("v8event Error:", err)
			}
		}()
		//getPtr := func(i int) unsafe.Pointer {
		//	return unsafe.Pointer(getVal(i))
		//}
		switch fn.(type) {

		default:
			return false
		}
		return true
	})
}
