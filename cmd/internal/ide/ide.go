package ide

import (
	"embed"
	"fmt"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

func Run(resources *embed.FS) {
	inits.Init(nil, resources)
	lcl.RunApp(&Ide)
}

const (
	Title   = "Energy 自定义安装程序图形化制作 - %s"
	Version = "1.0.0"
)
const (
	HTLEFT        = 10
	HTRIGHT       = 11
	HTTOP         = 12
	HTTOPLEFT     = 13
	HTTOPRIGHT    = 14
	HTBOTTOM      = 15
	HTBOTTOMLEFT  = 16
	HTBOTTOMRIGHT = 17
)

const (
	BorderRange  int32 = 8
	BorderMargin       = 4
	Border             = BorderMargin / 2
	PointW             = 6
	PointWC            = PointW / 2
	MinW, MinH         = 24, 8
)

type ComponentType int8

const (
	CtForm ComponentType = iota
	CtImage
	CtButton
	CtLabel
	CtEdit
	CtOpenDialog
)

var (
	Ide            *IDE //= &IDE{forms: make([]*IDEForm, 0, 0)}
	Fx, Fy, Fw, Fh int32
)

type IDE struct {
	*lcl.TForm
	formCount            int
	statusBar            *lcl.TStatusBar
	topBox               *lcl.TPanel
	leftBox              *lcl.TPanel
	rightBox             *lcl.TPanel
	imageList            *lcl.TImageList
	actionList           *lcl.TActionList
	mainMenu             *lcl.TMainMenu
	topToolBar           *lcl.TToolBar
	topToolButton        *lcl.TToolButton
	forms                []*IDEForm
	active               *IDEForm
	pageControl          *lcl.TPageControl
	pageControlPopupMenu *lcl.TPopupMenu
}

func (m *IDE) OnFormCreate(sender lcl.IObject) {
	lcl.Application.Icon().LoadFromFSFile("resources/icon.ico")
	m.forms = make([]*IDEForm, 0, 0)
	m.SetCaption(fmt.Sprintf(Title, Version))
	m.SetWidth(1200)
	m.SetHeight(800)
	m.SetDoubleBuffered(true)
	m.SetColor(colors.ClWhite)
	m.SetShowHint(true)

	m.statusBar = lcl.NewStatusBar(m)
	m.statusBar.SetParent(m)
	m.statusBar.SetAlign(types.AlBottom)
	m.statusBar.SetAnchors(types.NewSet(types.AkLeft, types.AkBottom, types.AkRight))
	m.statusBar.SetAutoHint(true)
	m.statusBar.SetSimplePanel(true)

	//ide Top
	m.topBox = lcl.NewPanel(m)
	m.topBox.SetParent(m)
	m.topBox.SetWidth(m.Width())
	m.topBox.SetHeight(100)
	m.topBox.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkRight))
	m.initTopMainMenu()

	//ide Left
	m.leftBox = lcl.NewPanel(m)
	m.leftBox.SetParent(m)
	m.leftBox.SetTop(m.topBox.Height())
	m.leftBox.SetWidth(250)
	m.leftBox.SetHeight(m.Height() - m.leftBox.Top() - 20)
	m.leftBox.SetColor(colors.ClAntiquewhite)
	m.leftBox.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkBottom))

	//ide box
	m.rightBox = lcl.NewPanel(m)
	m.rightBox.SetParent(m)
	m.rightBox.SetTop(m.topBox.Height())
	m.rightBox.SetLeft(m.leftBox.Width())
	m.rightBox.SetWidth(m.Width() - m.rightBox.Left())
	m.rightBox.SetHeight(m.Height() - m.rightBox.Top() - 20)
	m.rightBox.SetBevelInner(types.BvNone)
	m.rightBox.SetBevelOuter(types.BvNone)
	m.rightBox.SetBorderStyle(types.BsNone)
	m.rightBox.SetColor(colors.ClBlanchedalmond)
	m.rightBox.SetAnchors(types.NewSet(types.AkLeft, types.AkTop, types.AkBottom, types.AkRight))

	m.pageControl = lcl.NewPageControl(m.rightBox)
	m.pageControl.SetParent(m.rightBox)
	m.pageControl.SetAlign(types.AlClient)
	//m.pageControlPopupMenu = lcl.NewPopupMenu(m.pageControl)
	//item := lcl.NewMenuItem(m.pageControl)
	//item.SetCaption("修改名称")
	//item.SetOnClick(func(lcl.IObject) {
	//	fmt.Println("修改名称")
	//})
	//m.pageControlPopupMenu.Items().Add(item)
	//item = lcl.NewMenuItem(m.pageControl)
	//item.SetCaption("关闭")
	//item.SetOnClick(func(lcl.IObject) {
	//	fmt.Println("关闭")
	//	m.removeForm(m.active.Id)
	//})
	//m.pageControlPopupMenu.Items().Add(item)
	//m.pageControl.SetPopupMenu(m.pageControlPopupMenu)
	//m.pageControl.SetOnContextPopup(func(sender lcl.IObject, mousePos types.TPoint, handled *bool) {
	//	pageIndex := m.pageControl.IndexOfPageAt(mousePos.X, mousePos.Y)
	//	//sheet := pageControl.Pages(pageIndex)
	//	fmt.Println("pageIndex", pageIndex)
	//	if pageIndex >= 2 {
	//		m.pageControl.SetActivePageIndex(pageIndex)
	//		m.active = m.forms[int(pageIndex)]
	//		m.active.Id = int(pageIndex)
	//	} else {
	//		*handled = true
	//	}
	//})

	app(m.CreateForm("应用程序信息配置"))
	install(m.CreateForm("安装程序-欢迎界面"))
	install(m.CreateForm("安装程序-目录选择界面"))
	install(m.CreateForm("安装程序-安装过程界面"))
	install(m.CreateForm("安装程序-完成界面"))
	install(m.CreateForm("安装程序-卸载界面"))
	m.ScreenCenter()
}

func (m *IDE) addForm(form *IDEForm) int {
	form.Id = int(m.pageControl.ControlCount()) - 1
	m.forms = append(m.forms, form)
	m.active = form
	m.pageControl.SetActivePageIndex(int32(form.Id))
	m.formCount++
	return form.Id
}

func (m *IDE) removeForm(index int) {
	form := m.forms[index]
	form.TabSheet.Free()
	m.forms = append(m.forms[:index], m.forms[index+1:]...)
}

func (m *IDE) formsSyncSize(id int) {
	for _, form := range m.forms {
		if form.Id != id {
			form.BorderPanel.SetBounds(Fx-Border, Fy-Border, Fw+Border, Fh+Border)
			form.ParentToPanel().SetBounds(Fx-Border/2, Fy-Border/2, Fw, Fh)
		}
	}
}

func (m *IDE) initPopupMenu() {
	pm := lcl.NewPopupMenu(m)
	item := lcl.NewMenuItem(m)
	item.SetCaption("退出(&E)")
	item.SetOnClick(func(lcl.IObject) {
		m.Close()
	})
	pm.Items().Add(item)
	m.SetPopupMenu(pm)
}

func (m *IDE) initTopMainMenu() {
	//创建Image Icon集合
	m.imageList = lcl.NewImageList(m)
	//加载图标
	icon0 := lcl.NewIcon()
	icon0.LoadFromFSFile("resources/icon.ico")
	m.imageList.AddIcon(icon0)
	m.actionList = lcl.NewActionList(m)
	m.actionList.SetImages(m.imageList)

	m.mainMenu = lcl.NewMainMenu(m)
	m.mainMenu.SetImages(m.imageList)
	//m.mainMenu.SetOnMeasureItem(func(sender lcl.IObject, aCanvas *lcl.TCanvas, width, height *int32) {
	//	//*height = 44
	//})

	m.topToolBar = lcl.NewToolBar(m)
	m.topToolBar.SetParent(m)
	m.topToolBar.SetImages(m.imageList)

	action := lcl.NewAction(m)
	action.SetCaption("新建(&F)")
	action.SetImageIndex(0)
	action.SetHint("新建Form窗口|新建一个Form窗口")
	action.SetOnExecute(func(sender lcl.IObject) {
		var form = m.CreateForm(fmt.Sprintf("Form%d", m.formCount))
		form.CreateDialogOpen()
		form.CreateEdit()
		form.CreateImage()
		form.CreateLabel()
		form.CreateButton()
		println("CreateForm:", form)
	})

	m.topToolButton = lcl.NewToolButton(m)
	m.topToolButton.SetParent(m.topToolBar)
	m.topToolButton.SetAction(action)

	fileMenuItem := lcl.NewMenuItem(m)
	fileMenuItem.SetCaption("文件(&F)")

	fileMenuItemSubNew := lcl.NewMenuItem(m)
	fileMenuItemSubNew.SetCaption("新建(&N)")
	fileMenuItemSubNew.SetAction(action)
	fileMenuItemSubNew.SetShortCutFromString("Ctrl+N")
	//fileMenuItemSubNew.SetOnClick(func(lcl.IObject) {
	//	fmt.Println("新建")
	//})
	fileMenuItem.Add(fileMenuItemSubNew)
	fileMenuItemSubSave := lcl.NewMenuItem(m)
	fileMenuItemSubSave.SetCaption("保存(&S)")
	fileMenuItemSubSave.SetShortCutFromString("Ctrl+S")
	fileMenuItemSubSave.SetOnClick(func(lcl.IObject) {
		fmt.Println("保存")
	})
	fileMenuItem.Add(fileMenuItemSubSave)
	separate := lcl.NewMenuItem(m)
	separate.SetCaption("-")
	fileMenuItem.Add(separate)
	fileMenuItemSubExit := lcl.NewMenuItem(m)
	fileMenuItemSubExit.SetCaption("退出(&E)")
	fileMenuItemSubExit.SetOnClick(func(lcl.IObject) {
		fmt.Println("退出")
	})
	fileMenuItem.Add(fileMenuItemSubExit)
	//add file
	m.mainMenu.Items().Add(fileMenuItem)

	editMenuItem := lcl.NewMenuItem(m)
	editMenuItem.SetCaption("编辑(&E)")
	//add edit
	m.mainMenu.Items().Add(editMenuItem)

	configMenuItem := lcl.NewMenuItem(m)
	configMenuItem.SetCaption("设置(&C)")
	configMenuSubEnv := lcl.NewMenuItem(m)
	configMenuSubEnv.SetCaption("环境配置")
	configMenuSubEnv.SetOnClick(func(lcl.IObject) {
		fmt.Println("环境配置")
	})
	configMenuItem.Add(configMenuSubEnv)
	//add config
	m.mainMenu.Items().Add(configMenuItem)
}

func (m *IDE) CreateForm(name string) *IDEForm {
	var left, top, width, height int32 = Border, Border, 600, 400

	form := &IDEForm{Components: map[int]*IDEComponent{}}
	form.TabSheet = lcl.NewTabSheet(m.pageControl)
	form.TabSheet.SetPageControl(m.pageControl)
	form.TabSheet.SetAlign(types.AlClient)

	form.BorderPanel = lcl.NewPanel(form.TabSheet)
	form.BorderPanel.SetParent(form.TabSheet)
	form.BorderPanel.SetDoubleBuffered(true)
	form.BorderPanel.SetBevelInner(types.BvNone)
	form.BorderPanel.SetBevelOuter(types.BvNone)
	form.BorderPanel.SetBorderStyle(types.BsNone)
	form.BorderPanel.SetBounds(left-Border, top-Border, width+Border, height+Border)
	form.BorderPanel.SetColor(colors.ClBlack)

	form.ComponentParentPanel = lcl.NewPanel(form.TabSheet)
	form.ParentToPanel().SetParent(form.TabSheet)
	form.ParentToPanel().SetDoubleBuffered(true)
	form.ParentToPanel().SetBevelInner(types.BvNone)
	form.ParentToPanel().SetBevelOuter(types.BvNone)
	form.ParentToPanel().SetBorderStyle(types.BsNone)
	form.ParentToPanel().SetBounds(left-Border/2, top-Border/2, width, height)
	form.ParentToPanel().SetOnMouseMove(form.MouseMove)
	form.ParentToPanel().SetOnMouseDown(form.MouseDown)
	form.ParentToPanel().SetOnMouseUp(form.MouseUp)

	form.Ox, form.Oy, form.Ow, form.Oh = form.ParentToPanel().Left(), form.ParentToPanel().Top(), form.ParentToPanel().Width(), form.ParentToPanel().Height()
	m.addForm(form)
	form.Name = name
	form.ComponentType = CtForm
	form.TabSheet.SetCaption(form.Name)
	return form
}
