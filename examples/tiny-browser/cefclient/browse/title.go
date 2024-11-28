package browse

import (
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/golcl/lcl"
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
}

func NewTitleBar(window *cef.TCEFWindowComponent) *TitleBar {
	return &TitleBar{window: window}
}

func (m *TitleBar) EnsureTitlePanel() *cef.ICefPanel {
	if m.titlePanel == nil {
		m.titlePanelDelegate = cef.PanelDelegateRef.New()
		m.titlePanelDelegate.SetOnGetPreferredSize(func(view *cef.ICefView, result *cef.TCefSize) {
			fmt.Println("titlePanelDelegate.SetOnGetPreferredSize")
			m.titlePanel.SetBackgroundColor(cef.CefColorSetARGB(255, 237, 237, 237))
			regions := make([]cef.TCefDraggableRegion, 1)
			regions[0] = cef.TCefDraggableRegion{Bounds: cef.TCefRect{X: 0, Y: 0, Width: m.window.GetBounds().Width, Height: 30}, Draggable: true}
			fmt.Println("OnGetPreferredSize RegionsCount:", regions, m.window.GetBounds())
			m.window.SetDraggableRegions(regions)
		})

		m.menuButtonDelegate = cef.MenuButtonDelegateRef.New()
		m.menuButtonDelegate.SetOnWindowChanged(func(view *cef.ICefView, added bool) {
		})
		m.buttonDelegate = cef.ButtonDelegateRef.New()
		m.titlePanelDelegate.SetOnWindowChanged(func(view *cef.ICefView, added bool) {
			fmt.Println("titlePanelDelegate.SetOnWindowChanged added:", added)
			if added {
				m.titlePanel.SetBackgroundColor(cef.CefColorSetARGB(255, 77, 177, 177))
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
		})
		m.titlePanelDelegate.SetOnThemeChanged(func(view *cef.ICefView) {
			fmt.Println("titlePanelDelegate.SetOnThemeChanged")
		})
		m.titlePanel = cef.PanelRef.New(m.titlePanelDelegate)
		m.titlePanel.SetSize(cef.TCefSize{Height: 40})
		m.titlePanel.SetBackgroundColor(cef.CefColorSetARGB(255, 77, 177, 177))
		m.titlePanel.SetToFillLayout()
		m.titlePanelLayout = m.titlePanel.SetToBoxLayout(cef.TCefBoxLayoutSettings{
			BetweenChildSpacing: 2,
			Horizontal:          1,
			CrossAxisAlignment:  consts.CEF_AXIS_ALIGNMENT_CENTER,
		})
	}
	return m.titlePanel
}

func (m *TitleBar) CreateTitleBrowser(owner lcl.IComponent) {
	m.EnsureTitlePanel()
	m.chromium = cef.NewChromium(owner, nil)
	m.browserView = cef.BrowserViewComponentRef.New(owner)
	fmt.Println("CreateTitleBrowser ")
	if m.chromium.CreateBrowserByBrowserViewComponent("http://localhost:22022", m.browserView, nil, nil) {
		fmt.Println("CreateTitleBrowser 1")
		m.titlePanel.AddChildView(m.browserView.AsView())
	}
}
