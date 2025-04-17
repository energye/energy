package browse

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/consts"
	"github.com/cyber-xxm/energy/v2/examples/tiny-browser/cefclient/views_style"
)

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
			if commandId == ID_QUIT {
				window.chromium.CloseBrowser(true)
			}
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
	if menuId != nil && *menuId == 0 {
		*menuId = m.idNext
	} else {
		m.idNext = *menuId
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
	var id = ID_MENU_BUTTON
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
	menuModel.AddItem(ID_QUIT, "&Exit")
	// 在菜单中显示快捷键文本。
	menuModel.SetAcceleratorAt(menuModel.GetCount()-1, 'X', true, true, false)
	// 还要在窗口添加快捷键，否则不触发 SetOnAccelerator
	m.window.SetAccelerator(ID_QUIT, 'X', true, true, false, false)
	m.menuModels[id] = menuModel
	fmt.Println("file menu Id:", id)
	subMenuModel := menuModel.AddSubMenu(0, "子菜单")
	subMenuModel.AddItem(ID_TESTS_GETSOURCE, "这是一个子菜单项")
	views_style.ApplyTo(menuModel)
}
