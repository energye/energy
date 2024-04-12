package rod

import (
	"context"
	"encoding/json"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/cdp"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/defaults"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/proto"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/utils"
	engJSON "github.com/energye/energy/v2/pkgs/json"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Result struct {
	Msg json.RawMessage
	Err error
}

type OnBeforePopup func(energy *Energy)
type OnClose func(energy *Energy)
type OnLoadingProgressChange func(energy *Energy, progress float64)
type OnDevToolsRawMessage func(data []byte)

// Energy Devtools message processing structure for rod extension encapsulation
type Energy struct {
	rodBrowser      *Browser
	chromium        cef.IChromium
	chromiumBrowser cef.ICEFChromiumBrowser
	window          cef.IBrowserWindow
	targetId        proto.TargetTargetID
	created         bool

	onBeforePopup           OnBeforePopup
	onLoadingProgressChange OnLoadingProgressChange
	onClose                 OnClose
	onDevToolsRawMessage    OnDevToolsRawMessage

	loadSuccess     bool
	timer           *time.Timer
	pageLoadProcess float64

	page    *Page
	count   uint64
	pending *sync.Map       // pending requests
	event   chan *cdp.Event // events from browser
	logger  utils.Logger
}

// NewEnergyChromium Create a chrome and layout it in the current main window
func NewEnergyChromium(owner lcl.IWinControl, config *cef.TCefChromiumConfig) *Energy {
	m := &Energy{
		event:   make(chan *cdp.Event),
		logger:  defaults.CDP,
		pending: new(sync.Map),
	}
	m.rodBrowser = New()
	m.rodBrowser.client = m
	m.chromiumBrowser = cef.NewChromiumBrowser(owner, config)
	m.chromiumBrowser.RegisterDefaultEvent()
	m.chromium = m.chromiumBrowser.Chromium()
	m.listen()
	return m
}

// NewEnergyWindow creates a window
func NewEnergyWindow(config *cef.TCefChromiumConfig, windowProperty cef.WindowProperty, owner lcl.IComponent) *Energy {
	m := &Energy{
		event:   make(chan *cdp.Event),
		logger:  defaults.CDP,
		pending: new(sync.Map),
	}
	m.rodBrowser = New()
	m.rodBrowser.client = m
	m.window = cef.NewBrowserWindow(config, windowProperty, owner)
	m.window.EnableAllDefaultEvent()
	m.chromium = m.window.Chromium()
	m.listen()
	return m
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

// Chromium return current chromium instance
func (m *Energy) Chromium() cef.IChromium {
	return m.chromium
}

// SetOnBeforePopup energy rod popup callback
func (m *Energy) SetOnBeforePopup(fn OnBeforePopup) {
	m.onBeforePopup = fn
}

// SetOnLoadingProgressChange page load process
func (m *Energy) SetOnLoadingProgressChange(fn OnLoadingProgressChange) {
	m.onLoadingProgressChange = fn
}

// SetOnClose window close callback
func (m *Energy) SetOnClose(fn OnClose) {
	m.onClose = fn
}

// SetOnDevToolsRawMessage Call SendDevToolsMessage or ExecuteDevToolsMethod. If successfully validated, the callback function will be executed and it will return the execution result
func (m *Energy) SetOnDevToolsRawMessage(fn OnDevToolsRawMessage) {
	m.onDevToolsRawMessage = fn
}

// TargetInfo Return current target info
func (m *Energy) TargetInfo() *proto.TargetTargetInfo {
	result, err := proto.TargetGetTargetInfo{TargetID: m.targetId}.Call(m)
	m.rodBrowser.e(err)
	return result.TargetInfo
}

// Targets Return All Targets Info
func (m *Energy) Targets() []*proto.TargetTargetInfo {
	result, err := proto.TargetGetTargets{}.Call(m)
	m.rodBrowser.e(err)
	return result.TargetInfos
}

// RodBrowser return RodBrowser
//
// Note that the devtools and rod for operating CEF in energy are different, and some functions cannot be directly used through rod
// For example, window state management or chrome closure requires obtaining window objects and chrome objects directly for use
func (m *Energy) RodBrowser() *Browser {
	return m.rodBrowser
}

// ChromiumBrowser return chromium
func (m *Energy) ChromiumBrowser() cef.ICEFChromiumBrowser {
	return m.chromiumBrowser
}

// BrowserWindow return Window
func (m *Energy) BrowserWindow() cef.IBrowserWindow {
	return m.window
}

// Page Return the current Chromium Page
func (m *Energy) Page() *Page {
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
func (m *Energy) CreateBrowser() {
	if !m.created {
		m.created = true
		// chromium
		if m.chromiumBrowser != nil {
			m.chromiumBrowser.CreateBrowser()
		} else if m.window != nil {
			// window
			if m.window.IsLCL() {
				cef.RunOnMainThread(func() {
					m.window.Show()
				})
			} else {
				m.window.Show()
			}
		}
		m.rodBrowser.initEvents()
	}
}

// LoadSuccess Returns whether the current page was successfully loaded
func (m *Energy) LoadSuccess() bool {
	return m.loadSuccess
}

// PageLoadProcess Return to page loading progress
func (m *Energy) PageLoadProcess() float64 {
	return m.pageLoadProcess
}

// Call a method and wait for its response.
func (m *Energy) Call(ctx context.Context, sessionID, method string, params interface{}) ([]byte, error) {
	//m.CheckWaitPageLoad()
	req := &cdp.Request{
		ID:        int(atomic.AddUint64(&m.count, 1)),
		SessionID: sessionID,
		Method:    method,
		Params:    params,
	}

	m.logger.Println(req)
	data, err := json.Marshal(params)
	utils.E(err)

	done := make(chan Result)
	m.pending.Store(req.ID, func(res Result) {
		done <- res
	})
	defer m.pending.Delete(req.ID)
	//m.logger.Println("send-data:", string(data))
	//fmt.Println("send-data:", string(data))
	//m.chromium.SendDevToolsMessage(string(data))// Linux cannot be used
	dict := JSONParse(data)
	m.chromium.ExecuteDevToolsMethod(int32(req.ID), req.Method, dict)
	res := <-done
	dict.Free()
	return res.Msg, res.Err
}

// Event returns a channel that will emit browser devtools protocol events. Must be consumed or will block producer.
func (m *Energy) Event() <-chan *cdp.Event {
	return m.event
}

// Pending Each message event result map
func (m *Energy) Pending() *sync.Map {
	return m.pending
}

func (m *Energy) listen() {
	if m.window.IsLCL() {
		bw := m.window.AsLCLBrowserWindow().BrowserWindow()
		bw.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) bool {
			if m.onClose != nil {
				m.onClose(m)
			}
			return false
		})
	} else {
		m.chromium.SetOnClose(func(sender lcl.IObject, browser *cef.ICefBrowser, aAction *consts.TCefCloseBrowserAction) {
			if m.onClose != nil {
				m.onClose(m)
			}
		})
	}
	// 消息接收，energy中使用 CEF 回调函数接收消息
	m.chromium.SetOnDevToolsRawMessage(func(sender lcl.IObject, browser *cef.ICefBrowser, message uintptr, messageSize uint32) (handled bool) {
		data := ReadData(message, messageSize)
		if m.onDevToolsRawMessage != nil {
			m.onDevToolsRawMessage(data)
		}
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
	//m.chromium.SetOnDevToolsEvent(func(sender lcl.IObject, browser *cef.ICefBrowser, method string, params *cef.ICefValue) {
	//	//fmt.Println("OnDevToolsEvent", method, "params:", params.GetType())
	//})
	//TODO 还未实现补全
	//m.chromium.SetOnDevToolsMethodRawResult(func(sender lcl.IObject, browser *cef.ICefBrowser, messageId int32, success bool, result uintptr, resultSize uint32) {
	//	//fmt.Println("OnDevToolsMethodRawResult messageId:", messageId, "success:", success, "result:", result, "resultSize:", resultSize)
	//})
	m.chromium.SetOnTitleChange(func(sender lcl.IObject, browser *cef.ICefBrowser, title string) {
		if m.window != nil {
			m.window.SetTitle(title)
		}
	})
	m.chromium.SetOnLoadingProgressChange(func(sender lcl.IObject, browser *cef.ICefBrowser, progress float64) {
		m.pageLoadProcess = progress
		m.loadSuccess = int(progress*100) == 100
		if m.onLoadingProgressChange != nil {
			m.onLoadingProgressChange(m, progress)
		}
	})

	m.chromium.SetOnBeforePopup(func(sender lcl.IObject, browser *cef.ICefBrowser, frame *cef.ICefFrame, beforePopupInfo *cef.BeforePopupInfo, popupFeatures *cef.TCefPopupFeatures, windowInfo *cef.TCefWindowInfo, resultClient *cef.ICefClient, settings *cef.TCefBrowserSettings, resultExtraInfo *cef.ICefDictionaryValue, noJavascriptAccess *bool) bool {
		if m.onBeforePopup != nil {
			wp := cef.NewWindowProperty()
			wp.Url = beforePopupInfo.TargetUrl
			window := NewEnergyWindow(m.chromium.Config(), wp, nil)
			cef.RunOnMainThread(func() {
				window.CreateBrowser()
				go m.onBeforePopup(window)
			})
		}
		return true
	})
}

func parseObject(object engJSON.JSONObject) *cef.ICefDictionaryValue {
	obj := cef.DictionaryValueRef.New()
	keys := object.Keys()
	for _, key := range keys {
		if key == "id" || key == "method" {
			continue
		}
		val := object.GetByKey(key)
		if val.IsInt() {
			obj.SetInt(key, int32(val.Int()))
		} else if val.IsBool() {
			obj.SetBool(key, val.Bool())
		} else if val.IsString() {
			obj.SetString(key, val.String())
		} else if val.IsFloat() {
			obj.SetDouble(key, val.Float())
		} else if val.IsObject() {
			obj.SetDictionary(key, parseObject(val.JSONObject()))
		} else if val.IsArray() {
			obj.SetList(key, parseArray(val.JSONArray()))
		}
	}
	return obj
}

func parseArray(array engJSON.JSONArray) *cef.ICefListValue {
	arr := cef.ListValueRef.New()
	for i := 0; i < array.Size(); i++ {
		val := array.GetByIndex(i)
		if val.IsInt() {
			arr.SetInt(uint32(i), int32(val.Int()))
		} else if val.IsBool() {
			arr.SetBool(uint32(i), val.Bool())
		} else if val.IsString() {
			arr.SetString(uint32(i), val.String())
		} else if val.IsFloat() {
			arr.SetDouble(uint32(i), val.Float())
		} else if val.IsObject() {
			arr.SetDictionary(uint32(i), parseObject(val.JSONObject()))
		} else if val.IsArray() {
			arr.SetList(uint32(i), parseArray(val.JSONArray()))
		}
	}
	return arr
}

func JSONParse(jsonByte []byte) (result *cef.ICefDictionaryValue) {
	obj := engJSON.NewJSONObject(jsonByte)
	result = parseObject(obj)
	return
}
