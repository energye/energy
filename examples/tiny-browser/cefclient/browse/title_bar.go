package browse

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl"
)

const titleHeight = 40

const (
	ID_WINDOW_MIN_BTN   = 101
	ID_WINDOW_MAX_BTN   = 102
	ID_WINDOW_CLOSE_BTN = 103
)

type TitleBar struct {
	chromium           cef.IChromium
	browserView        *cef.TCEFBrowserViewComponent
	titlePanel         *cef.ICefPanel
	titlePanelLayout   *cef.ICefBoxLayout
	titlePanelDelegate *cef.ICefPanelDelegate
	menuButtonDelegate *cef.ICefMenuButtonDelegate
	buttonDelegate     *cef.ICefButtonDelegate
	window             *cef.TCEFWindowComponent
	minBtn             *ImageButton
	maxBtn             *ImageButton
	closeBtn           *ImageButton
}

func NewTitleBar(window *cef.TCEFWindowComponent) *TitleBar {
	return &TitleBar{window: window}
}

func (m *TitleBar) EnsureTitlePanel() *cef.ICefPanel {
	if m.titlePanel == nil {
		updateWindowButtonState := func() {
			if m.maxBtn != nil {
				if m.window.IsMaximized() {
					m.maxBtn.btn.SetImage(consts.CEF_BUTTON_STATE_NORMAL, m.maxBtn.disable)
				} else {
					m.maxBtn.btn.SetImage(consts.CEF_BUTTON_STATE_NORMAL, m.maxBtn.enable)
				}
			}
		}
		m.titlePanelDelegate = cef.PanelDelegateRef.New()
		m.titlePanelDelegate.SetOnGetPreferredSize(func(view *cef.ICefView, result *cef.TCefSize) {
			fmt.Println("titlePanelDelegate.SetOnGetPreferredSize")
			m.titlePanel.SetBackgroundColor(cef.CefColorSetARGB(255, 237, 237, 237))
			regions := make([]cef.TCefDraggableRegion, 1)
			regions[0] = cef.TCefDraggableRegion{Bounds: cef.TCefRect{X: m.GetWindowButtonAllWidth(), Y: 0, Width: m.window.GetBounds().Width, Height: titleHeight}, Draggable: true}
			m.window.SetDraggableRegions(regions)
			updateWindowButtonState()
		})

		m.menuButtonDelegate = cef.MenuButtonDelegateRef.New()
		m.menuButtonDelegate.SetOnGetPreferredSize(func(view *cef.ICefView, result *cef.TCefSize) {
			*result = cef.TCefSize{Height: 40, Width: 35}
		})
		m.menuButtonDelegate.SetOnMenuButtonPressed(func(menuButton *cef.ICefMenuButton, screenPoint cef.TCefPoint, buttonPressedLock *cef.ICefMenuButtonPressedLock) {
			switch menuButton.GetID() {
			case ID_WINDOW_MIN_BTN:
				m.window.Minimize()
			case ID_WINDOW_MAX_BTN:
				if m.window.IsMaximized() {
					m.window.Restore()
				} else {
					m.window.Maximize()
				}
				updateWindowButtonState()
			case ID_WINDOW_CLOSE_BTN:
				window.chromium.CloseBrowser(true)
			}
		})
		m.buttonDelegate = cef.ButtonDelegateRef.New()
		m.titlePanelDelegate.SetOnWindowChanged(func(view *cef.ICefView, added bool) {
			fmt.Println("titlePanelDelegate.SetOnWindowChanged added:", added)
			if added {
				m.titlePanel.SetBackgroundColor(cef.CefColorSetARGB(255, 77, 177, 177))
			}
		})
		m.titlePanelDelegate.SetOnThemeChanged(func(view *cef.ICefView) {
			fmt.Println("titlePanelDelegate.SetOnThemeChanged")
		})
		m.titlePanel = cef.PanelRef.New(m.titlePanelDelegate)
		//m.titlePanel.SetSize(cef.TCefSize{Height: titleHeight})
		m.titlePanel.SetBackgroundColor(cef.CefColorSetARGB(255, 77, 177, 177))
		//m.titlePanel.SetToFillLayout()
		m.titlePanelLayout = m.titlePanel.SetToBoxLayout(cef.TCefBoxLayoutSettings{
			Horizontal:         1,
			DefaultFlex:        1,
			CrossAxisAlignment: consts.CEF_AXIS_ALIGNMENT_CENTER,
		})
	}
	return m.titlePanel
}
func (m *TitleBar) CreateButton(tooltip, iconEnable, iconDisable string, id int32) *ImageButton {
	return CreateImageButton("", tooltip, iconEnable, iconDisable, id, m.menuButtonDelegate)
}

func (m *TitleBar) GetWindowButtonAllWidth() int32 {
	var getWidth = func(btn *ImageButton) int32 {
		if btn == nil {
			return 0
		}
		return btn.btn.GetBounds().Width
	}
	return getWidth(m.closeBtn) + getWidth(m.maxBtn) + getWidth(m.minBtn)
}
func (m *TitleBar) CreateWindowButton() {
	m.closeBtn = m.CreateButton("关闭", "btn-close.png", "", ID_WINDOW_CLOSE_BTN)
	m.titlePanel.AddChildView(m.closeBtn.btn.AsView())
	m.titlePanelLayout.SetFlexForView(m.closeBtn.btn.AsView(), 0)
	m.maxBtn = m.CreateButton("最大化", "btn-max.png", "btn-max-re.png", ID_WINDOW_MAX_BTN)
	m.titlePanel.AddChildView(m.maxBtn.btn.AsView())
	m.titlePanelLayout.SetFlexForView(m.maxBtn.btn.AsView(), 0)
	m.minBtn = m.CreateButton("最小化", "btn-min.png", "", ID_WINDOW_MIN_BTN)
	m.titlePanel.AddChildView(m.minBtn.btn.AsView())
	m.titlePanelLayout.SetFlexForView(m.minBtn.btn.AsView(), 0)
}

func (m *TitleBar) CreateICON() {
	button := cef.MenuButtonRef.New(m.menuButtonDelegate, "")
	button.SetID(188)
	button.SetInkDropEnabled(true)
	button.SetEnabled(false)   // 默认为关闭
	button.SetFocusable(false) // 不要把焦点放在按钮上
	button.SetMinimumSize(cef.TCefSize{Height: 35})
	button.SetImage(consts.CEF_BUTTON_STATE_NORMAL, LoadImage("app-icon-min.png"))
	button.SetTooltipText("Go ENERGY")
	button.SetBackgroundColor(cef.CefColorSetARGB(255, 77, 177, 177))
	m.titlePanel.AddChildView(button.AsView())
}

func (m *TitleBar) CreateTitleBrowser(owner lcl.IComponent) {
	owner = lcl.NewComponent(nil)
	m.EnsureTitlePanel()
	//m.CreateICON()
	m.CreateWindowButton()
	m.chromium = cef.NewChromium(owner, nil)
	m.browserView = cef.BrowserViewComponentRef.New(owner)
	fmt.Println("CreateTitleBrowser ")
}

func (m *TitleBar) CreateChromiumBrowser(chromium cef.IChromium) {
	m.chromium.SetOnAfterCreated(func(sender lcl.IObject, browser *cef.ICefBrowser) {
		fmt.Println("TitleBar OnAfterCreated")
	})
	if m.chromium.CreateBrowserByBrowserViewComponent("http://localhost:22022", m.browserView, nil, nil) {
		fmt.Println("TitleBar CreateTitleBrowser 1")
		m.titlePanel.AddChildView(m.browserView.AsView())
		m.titlePanelLayout.SetFlexForView(m.browserView.AsView(), 1)
		m.browserView.RequestFocus()
	}
}
