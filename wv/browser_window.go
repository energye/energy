package wv

import (
	"github.com/energye/lcl/lcl"
	"github.com/energye/wv/wv"
)

type OnWindowCreate func(window IBrowserWindow)

type IBrowserWindow interface {
	lcl.IForm
	WindowParent() wv.IWVWindowParent
	Browser() wv.IWVBrowser
	SetOnAfterCreated(fn wv.TNotifyEvent)
	SetOnWebMessageReceived(fn wv.TOnWebMessageReceivedEvent)
}

type BrowserWindow struct {
	lcl.TForm
	windowParent         wv.IWVWindowParent
	browser              wv.IWVBrowser
	options              Options
	onWindowCreate       OnWindowCreate
	onAfterCreated       wv.TNotifyEvent
	onWebMessageReceived wv.TOnWebMessageReceivedEvent
}

func (m *BrowserWindow) FormCreate(sender lcl.IObject) {
	m.SetCaption(m.options.Name)
	m.ScreenCenter()
	m.SetWidth(m.options.Width)
	m.SetHeight(m.options.Height)
	m.SetDoubleBuffered(true)
	m.windowParent = wv.NewWVWindowParent(m)
	m.windowParent.SetParent(m)
	m.browser = wv.NewWVBrowser(m)
	m.windowParent.SetBrowser(m.browser)
	if m.options.DefaultURL != "" {
		m.browser.SetDefaultURL(m.options.DefaultURL)
	}
	// call window main form create callback
	if m.onWindowCreate != nil {
		m.onWindowCreate(m)
	}
}

func (m *BrowserWindow) WindowParent() wv.IWVWindowParent {
	return m.windowParent
}

func (m *BrowserWindow) Browser() wv.IWVBrowser {
	return m.browser
}

func (m *BrowserWindow) SetOnAfterCreated(fn wv.TNotifyEvent) {
	m.onAfterCreated = fn
}

func (m *BrowserWindow) SetOnWebMessageReceived(fn wv.TOnWebMessageReceivedEvent) {
	m.onWebMessageReceived = fn
}
