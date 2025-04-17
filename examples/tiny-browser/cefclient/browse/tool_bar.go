package browse

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/winapi"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/energye/golcl/lcl/api"
)

type ToolBar struct {
	toolPanel           *cef.ICefPanel
	toolPanelDelegate   *cef.ICefPanelDelegate
	toolPanelLayout     *cef.ICefBoxLayout
	buttons             []*cef.ICefLabelButton
	buttonDelegate      *cef.ICefButtonDelegate
	menuButtonDelegate  *cef.ICefMenuButtonDelegate
	locationBarDelegate *cef.ICefTextFieldDelegate
	locationBar         *cef.ICefTextfield
	rightMenuButton     *cef.ICefMenuButton
	window              *cef.TCEFWindowComponent
	btnBack             *ImageButton
	btnForward          *cef.ICefMenuButton
	btnReload           *cef.ICefMenuButton
	btnStop             *cef.ICefMenuButton
}

func NewToolBar(window *cef.TCEFWindowComponent) *ToolBar {
	return &ToolBar{window: window, buttons: make([]*cef.ICefLabelButton, 0)}
}

func (m *ToolBar) EnsureToolPanel() *cef.ICefPanel {
	if m.toolPanel == nil {
		m.toolPanelDelegate = cef.PanelDelegateRef.New()
		m.toolPanel = cef.PanelRef.New(m.toolPanelDelegate)
		m.toolPanelDelegate.SetOnGetPreferredSize(func(view *cef.ICefView, result *cef.TCefSize) {
			//m.toolPanel.SetBackgroundColor(cef.CefColorSetARGB(255, 237, 237, 237))
		})
		m.locationBarDelegate = cef.TextFieldDelegateRef.New()
		m.locationBarDelegate.SetOnKeyEvent(func(textField *cef.ICefTextfield, event *cef.TCefKeyEvent) bool {
			if event.Kind == consts.KEYEVENT_RAW_KEYDOWN && event.WindowsKeyCode == winapi.VK_RETURN {
				fmt.Println("OnKeyEvent", textField.GetID(), event.KeyDown(), textField.GetText())
				window.chromium.LoadUrl(textField.GetText())
				return true
			}
			return false
		})
		m.buttonDelegate = cef.ButtonDelegateRef.New()
		m.buttonDelegate.SetOnButtonPressed(func(button *cef.ICefButton) {
			fmt.Println("OnButtonPressed", button.GetID())
		})
		m.menuButtonDelegate = cef.MenuButtonDelegateRef.New()
		m.menuButtonDelegate.SetOnGetPreferredSize(func(view *cef.ICefView, result *cef.TCefSize) {
			*result = cef.TCefSize{Height: 40, Width: 40}
		})
		m.menuButtonDelegate.SetOnMenuButtonPressed(func(menuButton *cef.ICefMenuButton, screenPoint cef.TCefPoint, buttonPressedLock *cef.ICefMenuButtonPressedLock) {
			fmt.Println("OnMenuButtonPressed", menuButton.GetID(), "IsMainThread:", api.DCurrentThreadId() == api.DMainThreadId())
			browser := window.browserView.GetBrowser()
			switch menuButton.GetID() {
			case ID_BACK_BUTTON:
				browser.GoBack()
			case ID_FORWARD_BUTTON:
				browser.GoForward()
			case ID_STOP_BUTTON:
				browser.StopLoad()
			case ID_RELOAD_BUTTON:
				browser.Reload()
			case ID_MENU_BUTTON:
				buttonBounds := menuButton.GetBoundsInScreen()
				point := screenPoint
				if cef.CefIsRTL() {
					point.X += buttonBounds.Width - 4
				} else {
					point.X -= buttonBounds.Width - 4
				}
				testMenu := window.menuBar.menuModels[menuButton.GetID()]
				if testMenu != nil {
					displayBounds := menuButton.GetWindow().Display().WorkArea()
					availableHeight := displayBounds.Y + displayBounds.Height - buttonBounds.Y - buttonBounds.Height
					menuHeight := int32(testMenu.GetCount()) * buttonBounds.Height
					if menuHeight > availableHeight {
						point.Y -= buttonBounds.Height - 8
					}
					menuButton.ShowMenu(testMenu, point, consts.CEF_MENU_ANCHOR_TOPLEFT)
				}
			}
		})
		m.toolPanelLayout = m.toolPanel.SetToBoxLayout(cef.TCefBoxLayoutSettings{
			Horizontal:         1,
			DefaultFlex:        1,
			CrossAxisAlignment: consts.CEF_AXIS_ALIGNMENT_CENTER,
		})
	}
	return m.toolPanel

}

func (m *ToolBar) CreateBrowseButton(label string, id int32) *cef.ICefLabelButton {
	button := cef.LabelButtonRef.New(m.buttonDelegate, label)
	button.SetID(id)
	button.SetInkDropEnabled(true)
	button.SetEnabled(true)    // 默认为关闭
	button.SetFocusable(false) // 不要把焦点放在按钮上
	button.SetMinimumSize(cef.TCefSize{})
	m.buttons = append(m.buttons, button)
	return button
}

func (m *ToolBar) CreateButton(label, icon string, id int32) *cef.ICefMenuButton {
	button := cef.MenuButtonRef.New(m.menuButtonDelegate, "")
	button.SetID(id)
	button.SetInkDropEnabled(true)
	button.SetEnabled(true)    // 默认为关闭
	button.SetFocusable(false) // 不要把焦点放在按钮上
	button.SetMinimumSize(cef.TCefSize{})
	button.SetImage(consts.CEF_BUTTON_STATE_NORMAL, LoadImage(icon))
	button.SetTooltipText(label)
	return button
}

func (m *ToolBar) CreateImageButton(tooltip, iconEnable, iconDisable string, id int32) *ImageButton {
	return CreateImageButton("", tooltip, iconEnable, iconDisable, id, m.menuButtonDelegate)
}

func (m *ToolBar) CreateLocationBar() *cef.ICefTextfield {
	m.locationBar = cef.TextFieldRef.New(m.locationBarDelegate)
	m.locationBar.SetID(ID_URL_TEXTFIELD)
	m.locationBar.SetTextColor(cef.CefColorSetARGB(255, 150, 99, 55))
	m.locationBar.SetFontList("Microsoft YaHei, 微软雅黑, Bold 16px")
	m.locationBar.SetBackgroundColor(cef.CefColorSetARGB(50, 55, 99, 150))
	return m.locationBar
}

func (m *ToolBar) CreateRightMenuButton() *cef.ICefMenuButton {
	m.rightMenuButton = cef.MenuButtonRef.New(m.menuButtonDelegate, "")
	m.rightMenuButton.SetID(ID_MENU_BUTTON)
	m.rightMenuButton.SetInkDropEnabled(true)
	m.rightMenuButton.SetMinimumSize(cef.TCefSize{})
	m.rightMenuButton.SetMaximumSize(cef.TCefSize{Height: 40, Width: 40})
	m.rightMenuButton.SetImage(consts.CEF_BUTTON_STATE_NORMAL, LoadImage("right-menu.png"))
	m.rightMenuButton.SetTooltipText("单击打开功能菜单")
	return m.rightMenuButton
}

// 使所有的|按钮|相同的大小。
func (m *ToolBar) MakeButtonsSameSize() {
	size := cef.TCefSize{}
	// 确定按钮的最大尺寸。
	for _, button := range m.buttons {
		buttonSize := button.GetPreferredSize()
		if size.Width < buttonSize.Width {
			size.Width = buttonSize.Width
		}
		if size.Height < buttonSize.Height {
			size.Height = buttonSize.Height
		}
	}
	for _, button := range m.buttons {
		//设置按钮的最小尺寸。
		button.SetMinimumSize(size)
		//重新布局按钮和所有父视图。
		button.InvalidateLayout()
	}
}

func (m *ToolBar) AllButtonWidth() int32 {
	return m.buttons[0].GetBounds().Width*int32(len(m.buttons)) + m.rightMenuButton.GetBounds().Width + 100
}

func (m *ToolBar) CreateToolComponent() {
	m.EnsureToolPanel()

	//m.toolPanel.AddChildView(m.CreateBrowseButton("Back", ID_BACK_BUTTON).AsView())
	//m.toolPanel.AddChildView(m.CreateBrowseButton("Forward", ID_FORWARD_BUTTON).AsView())
	//m.toolPanel.AddChildView(m.CreateBrowseButton("Reload", ID_RELOAD_BUTTON).AsView())
	//m.toolPanel.AddChildView(m.CreateBrowseButton("Stop", ID_STOP_BUTTON).AsView())

	m.btnBack = m.CreateImageButton("单击返回", "back.png", "back-disable.png", ID_BACK_BUTTON)
	m.btnForward = m.CreateButton("单击继续", "forward.png", ID_FORWARD_BUTTON)
	m.btnReload = m.CreateButton("单击刷新", "refresh.png", ID_RELOAD_BUTTON)
	m.btnStop = m.CreateButton("单击停止加载", "stop.png", ID_STOP_BUTTON)

	m.toolPanel.AddChildView(m.btnBack.btn.AsView())
	m.toolPanelLayout.SetFlexForView(m.btnBack.btn.AsView(), 0)
	m.toolPanel.AddChildView(m.btnForward.AsView())
	m.toolPanelLayout.SetFlexForView(m.btnForward.AsView(), 0)
	m.toolPanel.AddChildView(m.btnReload.AsView())
	m.toolPanelLayout.SetFlexForView(m.btnReload.AsView(), 0)
	m.toolPanel.AddChildView(m.btnStop.AsView())
	m.toolPanelLayout.SetFlexForView(m.btnStop.AsView(), 0)
	m.btnStop.SetVisible(false)

	m.MakeButtonsSameSize()

	m.toolPanel.AddChildView(m.CreateLocationBar().AsView())
	m.toolPanelLayout.SetFlexForView(m.locationBar.AsView(), 1)

	rightMenuButton := m.CreateRightMenuButton()
	m.toolPanel.AddChildView(rightMenuButton.AsView())
	m.toolPanelLayout.SetFlexForView(rightMenuButton.AsView(), 0)
}

func (m *ToolBar) UpdateBrowserStatus(isLoading, canGoBack, canGoForward bool) {
	btnForward := m.toolPanel.GetViewForID(ID_FORWARD_BUTTON)
	btnReload := m.toolPanel.GetViewForID(ID_RELOAD_BUTTON)
	btnStop := m.toolPanel.GetViewForID(ID_STOP_BUTTON)
	if canGoBack {
		m.btnBack.btn.SetImage(consts.CEF_BUTTON_STATE_NORMAL, m.btnBack.enable)
	} else {
		m.btnBack.btn.SetImage(consts.CEF_BUTTON_STATE_NORMAL, m.btnBack.disable)
	}
	m.btnBack.btn.SetEnabled(canGoBack)
	btnForward.SetVisible(canGoForward)
	btnReload.SetVisible(!isLoading)
	btnStop.SetVisible(isLoading)
}
