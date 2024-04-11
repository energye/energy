package rod

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/cdp"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/defaults"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/proto"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/utils"
	"github.com/energye/golcl/lcl"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Result struct {
	Msg json.RawMessage
	Err error
}

type OnBeforePopup func(chromium *Chromium)

type Chromium struct {
	rodBrowser      *Browser
	chromium        cef.IChromium
	chromiumBrowser cef.ICEFChromiumBrowser
	window          cef.IBrowserWindow
	title           string
	url             string
	targetId        proto.TargetTargetID
	created         bool

	onBeforePopup OnBeforePopup

	loadSuccess     bool
	enablePageCheck bool
	timer           *time.Timer
	pageLoadProcess float64

	page    *Page
	count   uint64
	pending *sync.Map       // pending requests
	event   chan *cdp.Event // events from browser
	logger  utils.Logger
}

// NewChromium Create a chrome and layout it in the current main window
func NewChromium(owner lcl.IWinControl, config *cef.TCefChromiumConfig) *Chromium {
	chromium := &Chromium{
		event:           make(chan *cdp.Event),
		logger:          defaults.CDP,
		pending:         new(sync.Map),
		enablePageCheck: true,
	}
	chromium.rodBrowser = New()
	chromium.rodBrowser.client = chromium
	chromium.chromiumBrowser = cef.NewChromiumBrowser(owner, config)
	chromium.chromiumBrowser.RegisterDefaultEvent()
	chromium.chromium = chromium.chromiumBrowser.Chromium()
	chromium.listen()
	return chromium
}

// NewWindow creates a window
func NewWindow(config *cef.TCefChromiumConfig, windowProperty cef.WindowProperty, owner lcl.IComponent) *Chromium {
	chromium := &Chromium{
		event:           make(chan *cdp.Event),
		logger:          defaults.CDP,
		pending:         new(sync.Map),
		enablePageCheck: true,
	}
	chromium.rodBrowser = New()
	chromium.rodBrowser.client = chromium
	chromium.window = cef.NewBrowserWindow(config, windowProperty, owner)
	chromium.window.EnableAllDefaultEvent()
	chromium.chromium = chromium.window.Chromium()
	chromium.listen()
	return chromium
}

// ReadData Read pointer data to [] byte
func ReadData(data uintptr, count uint32) []byte {
	result := make([]byte, count, count)
	var n uint32 = 0
	for n < count {
		result[n] = *(*byte)(unsafe.Pointer(data + uintptr(n)))
		n = n + 1
	}
	return result
}

func (m *Chromium) SetOnBeforePopup(fn OnBeforePopup) {
	m.onBeforePopup = fn
}

// TargetInfo Return current target info
func (m *Chromium) TargetInfo() *proto.TargetTargetInfo {
	result, err := proto.TargetGetTargetInfo{TargetID: m.targetId}.Call(m)
	m.rodBrowser.e(err)
	return result.TargetInfo
}

// Targets Return All Targets Info
func (m *Chromium) Targets() []*proto.TargetTargetInfo {
	result, err := proto.TargetGetTargets{}.Call(m)
	m.rodBrowser.e(err)
	return result.TargetInfos
}

// RodBrowser return RodBrowser
//
// Note that the devtools and rod for operating CEF in energy are different, and some functions cannot be directly used through rod
// For example, window state management or chrome closure requires obtaining window objects and chrome objects directly for use
func (m *Chromium) RodBrowser() *Browser {
	return m.rodBrowser
}

// ChromiumBrowser return chromium
func (m *Chromium) ChromiumBrowser() cef.ICEFChromiumBrowser {
	return m.chromiumBrowser
}

// BrowserWindow return Window
func (m *Chromium) BrowserWindow() cef.IBrowserWindow {
	return m.window
}

// Page Return the current Chromium Page
func (m *Chromium) Page() *Page {
	if m.page == nil {
		if m.targetId == "" {
			m.targetId = m.TargetInfo().TargetID
		}
		p, err := m.rodBrowser.PageFromTarget(m.targetId)
		if err != nil {
			return nil
		}
		m.page = p
	}
	return m.page
}

// CreateBrowser Call this function to create a browser after creating chrome or window
func (m *Chromium) CreateBrowser() {
	if !m.created {
		m.created = true
		if m.chromiumBrowser != nil {
			m.chromiumBrowser.CreateBrowser()
		} else if m.window != nil {
			m.window.Show()
		}
		m.rodBrowser.initEvents()
	}
}

// LoadSuccess Returns whether the current page was successfully loaded
func (m *Chromium) LoadSuccess() bool {
	return m.loadSuccess
}

// EnableCheckPageLoad  Enable page loading detection
func (m *Chromium) EnableCheckPageLoad(v bool) {
	m.enablePageCheck = v
}

// PageLoadProcess Return to page loading progress
func (m *Chromium) PageLoadProcess() float64 {
	return m.pageLoadProcess
}

// URL return current url
func (m *Chromium) URL() string {
	return m.url
}

// CheckWaitPageLoad Detect and wait for the current page to load until it is successfully loaded
func (m *Chromium) CheckWaitPageLoad() {
	if !m.loadSuccess && m.enablePageCheck {
		if m.timer == nil {
			m.timer = time.NewTimer(time.Second / 100)
		} else {
			m.timer.Reset(time.Second / 100)
		}
		defer m.timer.Stop()
		select {
		case <-m.timer.C:
			m.CheckWaitPageLoad()
		}
	}
}

// Call a method and wait for its response.
func (m *Chromium) Call(ctx context.Context, sessionID, method string, params interface{}) ([]byte, error) {
	m.CheckWaitPageLoad()
	req := &cdp.Request{
		ID:        int(atomic.AddUint64(&m.count, 1)),
		SessionID: sessionID,
		Method:    method,
		Params:    params,
	}

	m.logger.Println(req)

	data, err := json.Marshal(req)
	utils.E(err)

	done := make(chan Result)
	m.pending.Store(req.ID, func(res Result) {
		done <- res
	})
	defer m.pending.Delete(req.ID)
	//m.logger.Println("send-data:", string(data))
	m.chromium.SendDevToolsMessage(string(data))
	res := <-done
	return res.Msg, res.Err
	//return nil, nil
}

// Event returns a channel that will emit browser devtools protocol events. Must be consumed or will block producer.
func (m *Chromium) Event() <-chan *cdp.Event {
	return m.event
}

// Pending Each message event result map
func (m *Chromium) Pending() *sync.Map {
	return m.pending
}

func (m *Chromium) listen() {
	// 消息接收，energy中使用 CEF 回调函数接收消息
	m.chromium.SetOnDevToolsRawMessage(func(sender lcl.IObject, browser *cef.ICefBrowser, message uintptr, messageSize uint32) (handled bool) {
		data := ReadData(message, messageSize)
		var id struct {
			ID int `json:"id"`
		}
		err := json.Unmarshal(data, &id)
		utils.E(err)
		var res cdp.Response
		err = json.Unmarshal(data, &res)
		utils.E(err)
		m.logger.Println(&res)
		val, ok := m.pending.Load(id.ID)
		if !ok {
			return false
		}
		if res.Error == nil {
			val.(func(Result))(Result{res.Result, nil})
		} else {
			val.(func(Result))(Result{nil, res.Error})
		}
		return true
	})
	//TODO 还未实现补全
	m.chromium.SetOnDevToolsEvent(func(sender lcl.IObject, browser *cef.ICefBrowser, method string, params *cef.ICefValue) {
		fmt.Println("OnDevToolsEvent", method, "params:", params.GetType())
	})
	//TODO 还未实现补全
	m.chromium.SetOnDevToolsMethodRawResult(func(sender lcl.IObject, browser *cef.ICefBrowser, messageId int32, success bool, result uintptr, resultSize uint32) {
		fmt.Println("OnDevToolsMethodRawResult messageId:", messageId, "success:", success, "result:", result, "resultSize:", resultSize)
	})
	m.chromium.SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
		m.title = title
		if m.window != nil {
			m.window.SetTitle(title)
		}
	})
	m.chromium.SetOnLoadStart(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, transitionType consts.TCefTransitionType) {
		m.url = frame.Url()
	})
	m.chromium.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
		m.pageLoadProcess = progress
		m.loadSuccess = int(progress*100) == 100
	})

	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
		if m.onBeforePopup != nil {
			wp := cef.NewWindowProperty()
			wp.Url = beforePopupInfo.TargetUrl
			rodWindow := NewWindow(m.chromium.Config(), wp, nil)
			cef.RunOnMainThread(func() {
				rodWindow.CreateBrowser()
				go m.onBeforePopup(rodWindow)
			})
		}
		return true
	})
}
