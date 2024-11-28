package browse

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef"
	"github.com/energye/energy/v2/cef/winapi"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/examples/tiny-browser/cefclient/views_style"
)

var Assets embed.FS

// Control IDs for Views in the top-level Window.
const (
	ID_WINDOW int32 = iota
	ID_BROWSER_VIEW
	ID_BACK_BUTTON
	ID_FORWARD_BUTTON
	ID_STOP_BUTTON
	ID_RELOAD_BUTTON
	ID_URL_TEXTFIELD
	ID_MENU_BUTTON
	// Reserved range of top menu button IDs.
	ID_TOP_MENU_FIRST
	ID_TOP_MENU_LAST = ID_TOP_MENU_FIRST + 10
)
const kMenuBarGroupId int32 = 100

type MenuBar struct {
	menuPanel          *cef.ICefPanel
	menuPanelLayout    *cef.ICefBoxLayout
	menuPanelDelegate  *cef.ICefPanelDelegate
	menuModeDelegate   *cef.ICefMenuModelDelegate
	menuButtonDelegate *cef.ICefMenuButtonDelegate
	idNext             int32
	menuModels         map[int32]*cef.ICefMenuModel
	window             *cef.TCEFWindowComponent
}

func NewMenuBar(window *cef.TCEFWindowComponent) *MenuBar {
	return &MenuBar{menuModels: make(map[int32]*cef.ICefMenuModel), window: window}
}

func (m *MenuBar) EnsureMenuPanel() *cef.ICefPanel {
	if m.menuPanel == nil {
		m.menuPanelDelegate = cef.PanelDelegateRef.New()
		m.menuPanelDelegate.SetOnGetPreferredSize(func(view *cef.ICefView, result *cef.TCefSize) {
			//m.menuPanel.SetBackgroundColor(cef.CefColorSetARGB(255, 237, 237, 237))
		})
		m.menuModeDelegate = cef.MenuModelDelegateRef.New()
		m.menuModeDelegate.SetOnExecuteCommand(func(menuModel *cef.ICefMenuModel, commandId int32, eventFlags consts.TCefEventFlags) {
			fmt.Println("OnExecuteCommand commandId:", commandId, eventFlags)
		})
		m.menuButtonDelegate = cef.MenuButtonDelegateRef.New()
		m.menuButtonDelegate.SetOnMenuButtonPressed(func(menuButton *cef.ICefMenuButton, screenPoint cef.TCefPoint, buttonPressedLock *cef.ICefMenuButtonPressedLock) {
			buttonBounds := menuButton.GetBoundsInScreen()
			fmt.Printf("OnMenuButtonPressed screenPoint: %+v, ID: %+v, Text: %+v, bounds: %+v, buttonPressedLock: %+v \n", screenPoint, menuButton.GetID(), menuButton.GetText(), buttonBounds, buttonPressedLock)
			point := screenPoint
			if cef.CefIsRTL() {
				point.X += buttonBounds.Width - 4
			} else {
				point.X -= buttonBounds.Width - 4
			}
			displayBounds := menuButton.GetWindow().Display().WorkArea()
			availableHeight := displayBounds.Y + displayBounds.Height - buttonBounds.Y - buttonBounds.Height
			menuHeight := int32(m.menuModels[menuButton.GetID()].GetCount()) * buttonBounds.Height
			if menuHeight > availableHeight {
				point.Y -= buttonBounds.Height - 8
			}
			menuButton.ShowMenu(m.menuModels[menuButton.GetID()], point, consts.CEF_MENU_ANCHOR_TOPLEFT)
		})
		//m.menuButtonDelegate.SetOnGetPreferredSize(func(view *cef.ICefView, result *cef.TCefSize) {
		//	fmt.Println("OnGetPreferredSize", result)
		//})
		m.menuPanel = cef.PanelRef.New(m.menuPanelDelegate)
		m.menuPanelLayout = m.menuPanel.SetToBoxLayout(cef.TCefBoxLayoutSettings{
			BetweenChildSpacing: 2,
			Horizontal:          1,
			CrossAxisAlignment:  consts.CEF_AXIS_ALIGNMENT_CENTER,
		})
		m.idNext = ID_TOP_MENU_FIRST
	}
	return m.menuPanel
}

func (m *MenuBar) CreateMenuModel(label, tooltipText string, menuId *int32) *cef.ICefMenuModel {
	m.EnsureMenuPanel()
	m.idNext++
	if menuId != nil {
		*menuId = m.idNext
	}
	// 创建新的MenuModel。
	model := cef.MenuModelRef.New(m.menuModeDelegate)
	views_style.ApplyTo(model)

	// 创建新的MenuButton.
	button := cef.MenuButtonRef.New(m.menuButtonDelegate, label)
	button.SetID(m.idNext)
	button.SetInkDropEnabled(true)
	// 指定一个组ID，以便在不显示菜单时使用箭头键在MenuButtons之间进行焦点遍历。
	button.SetGroupID(kMenuBarGroupId)
	button.SetTooltipText(tooltipText)
	button.SetBackgroundColor(cef.CefColorSetARGB(255, 237, 237, 237))

	// 添加新的菜单按钮到平面上。
	m.menuPanel.AddChildView(button.AsView())
	return model
}

func (m *MenuBar) CreateTestMenuItems() {
	var id int32
	menuModel := m.CreateMenuModel("&Tests", "测试菜单按钮", &id)
	menuModel.AddItem(ID_TESTS_GETSOURCE, "Get Source")
	menuModel.AddItem(ID_TESTS_GETTEXT, "Get Text(获取文本)")
	menuModel.AddItem(ID_TESTS_WINDOW_NEW, "New Window")
	menuModel.AddItem(ID_TESTS_WINDOW_POPUP, "Popup Window")
	menuModel.AddItem(ID_TESTS_WINDOW_DIALOG, "Dialog Window")
	menuModel.AddItem(ID_TESTS_REQUEST, "Request")
	menuModel.AddItem(ID_TESTS_ZOOM_IN, "Zoom In")
	menuModel.AddItem(ID_TESTS_ZOOM_OUT, "Zoom Out")
	menuModel.AddItem(ID_TESTS_ZOOM_RESET, "Zoom Reset")
	menuModel.AddItem(ID_TESTS_TRACING_BEGIN, "Begin Tracing")
	menuModel.AddItem(ID_TESTS_TRACING_END, "End Tracing")
	menuModel.AddItem(ID_TESTS_PRINT, "Print")
	menuModel.AddItem(ID_TESTS_PRINT_TO_PDF, "Print to PDF")
	menuModel.AddItem(ID_TESTS_MUTE_AUDIO, "Mute Audio")
	menuModel.AddItem(ID_TESTS_UNMUTE_AUDIO, "Unmute Audio")
	menuModel.AddItem(ID_TESTS_OTHER_TESTS, "Other Tests")
	menuModel.AddItem(ID_TESTS_DUMP_WITHOUT_CRASHING, "Dump without crashing")
	views_style.ApplyTo(menuModel)
	m.menuModels[id] = menuModel
	fmt.Println("test menu Id:", id)
}

func (m *MenuBar) CreateFileMenuItems() {
	var id int32
	menuModel := m.CreateMenuModel("&File", "文件菜单按钮", &id)
	menuModel.AddItem(ID_QUIT, "E&xit")
	// 在菜单中显示快捷键文本。
	menuModel.SetAcceleratorAt(menuModel.GetCount()-1, 'X', true, true, false)
	// 还要在窗口添加快捷键，否则不触发 SetOnAccelerator
	m.window.SetAccelerator(id, 'X', true, true, false, false)
	m.menuModels[id] = menuModel
	fmt.Println("file menu Id:", id)
	subMenuModel := menuModel.AddSubMenu(0, "子菜单")
	subMenuModel.AddItem(ID_TESTS_GETSOURCE, "这是一个子菜单项")
	views_style.ApplyTo(menuModel)
}

type ToolBar struct {
	toolPanel           *cef.ICefPanel
	toolPanelDelegate   *cef.ICefPanelDelegate
	buttons             []*cef.ICefLabelButton
	buttonDelegate      *cef.ICefButtonDelegate
	menuButtonDelegate  *cef.ICefMenuButtonDelegate
	locationBarDelegate *cef.ICefTextFieldDelegate
	locationBar         *cef.ICefTextfield
	rightMenuButton     *cef.ICefMenuButton
	window              *cef.TCEFWindowComponent
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
				return true
			}
			return false
		})
		m.buttonDelegate = cef.ButtonDelegateRef.New()
		m.buttonDelegate.SetOnButtonPressed(func(button *cef.ICefButton) {
			fmt.Println("OnButtonPressed", button.GetID())
		})
		m.menuButtonDelegate = cef.MenuButtonDelegateRef.New()
		m.menuButtonDelegate.SetOnMenuButtonPressed(func(menuButton *cef.ICefMenuButton, screenPoint cef.TCefPoint, buttonPressedLock *cef.ICefMenuButtonPressedLock) {
			fmt.Println("OnMenuButtonPressed", menuButton.GetID())
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
	m.toolPanel.AddChildView(m.CreateButton("单击返回", "back.png", ID_BACK_BUTTON).AsView())
	m.toolPanel.AddChildView(m.CreateButton("单击继续", "forward.png", ID_FORWARD_BUTTON).AsView())
	m.toolPanel.AddChildView(m.CreateButton("单击刷新", "refresh.png", ID_RELOAD_BUTTON).AsView())
	m.toolPanel.AddChildView(m.CreateButton("单击停止加载", "stop.png", ID_STOP_BUTTON).AsView())
	m.MakeButtonsSameSize()

	m.toolPanel.AddChildView(m.CreateLocationBar().AsView())
	m.LayoutLocationBar()

	m.toolPanel.AddChildView(m.CreateRightMenuButton().AsView())
}

func (m *ToolBar) LayoutLocationBar() {
	// 使用水平盒布局|面板|
	layout := m.toolPanel.SetToBoxLayout(cef.TCefBoxLayoutSettings{
		BetweenChildSpacing: 2,
		Horizontal:          1,
		CrossAxisAlignment:  consts.CEF_AXIS_ALIGNMENT_CENTER,
	})
	layout.SetFlexForView(m.locationBar.AsView(), 1)
}

func LoadImage(png string) *cef.ICefImage {
	pngData, err := Assets.ReadFile("assets/icon/" + png)
	if err != nil {
		panic(err)
	}
	icon := cef.ImageRef.New()
	icon.AddPng(1, pngData)
	return icon
}
