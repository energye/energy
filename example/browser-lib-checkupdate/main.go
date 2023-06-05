// auto update application

package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef/autoupdate"
	"github.com/energye/energy/v2/cef/i18n"
	"github.com/energye/energy/v2/common"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/consts"
	"github.com/energye/energy/v2/example/browser-lib-checkupdate/form"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"runtime"
	"strconv"
	"strings"
)

/*
  检测liblcl库的更新
  下面的代码演示是未通过cef库导入的liblcl,而是单独导入,这样可以减少执行文件的大小
    []*dllimports.ImportTable
  通过imports.SetEnergyImportDefs(version)将proc导入到执行文件中, 为保持和cef的proc下标一致,第0个下标为空导入
*/

var (
	//go:embed resources
	resources embed.FS
	// form
	updateForm *form.UpdateForm
	version    = []*dllimports.ImportTable{
		dllimports.NewEnergyImport("", 0),                //空导入
		dllimports.NewEnergyImport("LibVersion", 0),      //获取lib库的版本号
		dllimports.NewEnergyImport("LibBuildVersion", 0), //获取lib库的构建工具版本
	}
)

func main() {
	// 注入到 imports
	imports.SetEnergyImportDefs(version)
	// 初始化golcl
	inits.Init(nil, &resources)
	i18n.SetLocalFS(&resources, "resources")
	//i18n.Switch(consts.LANGUAGE_en_US)
	i18n.Switch(consts.LANGUAGE_zh_CN)
	// 开启自动更新检查
	autoupdate.IsCheckUpdate(true)
	// 如果 liblcl 有更新该函数将被回调
	autoupdate.CanUpdateLiblcl = func(model *autoupdate.Model, level int) {
		fmt.Println(*model)
		updateVersion := model.Versions[model.Latest]
		fmt.Println(model.Versions[model.Latest])
		fmt.Println(level)

		var energyLiblcl = func() (string, bool) {
			if common.IsWindows() {
				return fmt.Sprintf("Windows %d bits", strconv.IntSize), true
			} else if common.IsLinux() {
				return "Linux x86 64 bits", true
			} else if common.IsDarwin() {
				return "MacOSX x86 64 bits", true
			}
			//not support download
			return fmt.Sprintf("%v %v", runtime.GOOS, runtime.GOARCH), false
		}

		// 这里使用 窗口形式展示更新
		// 运行应用后窗口创建时回调
		form.OnCreate = func(m *form.UpdateForm) {
			// 应用图标
			if runtime.GOOS == "windows" {
				lcl.Application.Icon().LoadFromFSFile("resources/icon.ico")
			} else {
				lcl.Application.Icon().LoadFromFSFile("resources/icon.png")
			}
			// 窗口一些属性配置
			m.SetDoubleBuffered(true)
			m.EnabledMinimize(false)
			m.EnabledMaximize(false)
			m.SetShowHint(true)
			//m.SetFormStyle(types.FsSystemStayOnTop)
			m.SetPosition(types.PoDesktopCenter)
			//m.SetBorderStyle(types.BsSingle)
			m.SetBorderStyle(types.BsNone)
			//m.SetShowInTaskBar(types.StNever)
			m.SetColor(colors.ClWhite)
			m.SetWidth(590)
			m.SetHeight(390)
			m.SetCaption(i18n.Resource("title")) // 自定义窗口标题

			// 自定义窗口标题栏
			m.TitlePanel = m.NewPanel()
			m.TitlePanel.SetColor(colors.ClTeal)
			m.TitlePanel.SetHeight(32) // icon 的高
			m.TitlePanel.SetWidth(m.Width())
			// 模拟标题栏移动窗口
			var (
				isDown bool
				dx, dy int32
			)
			m.TitlePanel.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
				isDown = true
				dx, dy = x, y
			})
			m.TitlePanel.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
				isDown = false
			})
			m.TitlePanel.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
				if isDown { //鼠标按下时计算移动坐标
					m.SetLeft(m.Left() - (dx - x))
					m.SetTop(m.Top() - (dy - y))
				}
			})

			// title -> icon
			titleIcon := lcl.NewImage(m.TitlePanel)
			titleIcon.SetParent(m.TitlePanel)
			titleIcon.SetWidth(32)  // icon 的宽
			titleIcon.SetHeight(32) // icon 的高
			titleIcon.Picture().LoadFromFSFile("resources/icon.png")

			// title -> text
			titleText := lcl.NewLabel(m.TitlePanel)
			titleText.SetParent(m.TitlePanel)
			titleText.SetTop(3)
			titleText.SetLeft(40)
			titleText.Font().SetSize(12)
			titleText.Font().SetColor(colors.ClWhite)
			titleText.SetCaption(m.Caption())

			// title -> close button
			titleClose := lcl.NewImageButton(m.TitlePanel)
			titleClose.SetParent(m.TitlePanel)
			titleClose.SetImageCount(4)
			titleClose.SetAutoSize(true)
			titleClose.SetCursor(types.CrHandPoint)
			titleClose.Picture().LoadFromFSFile("resources/btn_close.png")
			titleClose.SetLeft(m.Width() - 40) //关闭按钮位置 left = 窗口宽 - 按钮图片宽
			titleClose.SetHint(i18n.Resource("close"))
			titleClose.SetOnClick(func(lcl.IObject) {
				m.Close() // 关闭窗口
			})

			// background
			bgImage := lcl.NewImage(m)
			bgImage.SetParent(m)
			bgImage.SetTop(m.TitlePanel.Height())
			bgImage.SetWidth(271)                              // 图片宽
			bgImage.SetHeight(60)                              // 图片高
			bgImage.SetLeft((m.Width() - bgImage.Width()) / 2) // 设置以窗口居中
			bgImage.Picture().LoadFromFSFile("resources/bg.png")

			// 更新提醒 panel
			m.UpdatePromptPanel = m.NewPanel()
			m.UpdatePromptPanel.SetTop(bgImage.Top() + bgImage.Height())
			m.UpdatePromptPanel.SetWidth(m.Width())
			m.UpdatePromptPanel.SetHeight(m.Height() - m.UpdatePromptPanel.Top())

			// 更新进度 panel
			m.UpdateProgressPanel = m.NewPanel()
			m.UpdateProgressPanel.SetTop(bgImage.Top() + bgImage.Height())
			m.UpdateProgressPanel.SetWidth(m.Width())
			m.UpdateProgressPanel.SetHeight(m.Height() - m.UpdateProgressPanel.Top())
			m.UpdateProgressPanel.SetVisible(false)

			// 更新内容
			updateContent := lcl.NewMemo(m.UpdatePromptPanel)
			updateContent.SetParent(m.UpdatePromptPanel)
			ucw := m.Width() / 4
			updateContent.SetWidth(ucw * 3)
			updateContent.SetLeft((m.Width() - updateContent.Width()) / 2)
			updateContent.SetHeight(180)
			updateContent.SetReadOnly(true)
			updateContent.SetColor(colors.ClWhite)
			updateContent.SetScrollBars(types.SsAutoBoth)
			updateContent.Lines().Add(i18n.Resource("updateContent") + " " + model.Latest)
			for i, content := range updateVersion.Content {
				updateContent.Lines().Add("  " + strconv.Itoa(i+1) + ". " + content)
			}
			liblclZipName, _ := energyLiblcl()
			downUrl := strings.Replace(model.Download.Url, "{url}", model.Download.Source[model.Download.SourceSelect], -1)
			downUrl = strings.Replace(downUrl, "{version}", updateVersion.EnergyVersion, -1)
			downUrl = strings.Replace(downUrl, "{OSARCH}", liblclZipName, -1)
			fmt.Println("downUrl", downUrl)
			updateContent.Lines().Add("")
			updateContent.Lines().Add(i18n.Resource("downloadURL"))
			updateContent.Lines().Add(downUrl)

			//
			ok := lcl.NewButton(m.UpdatePromptPanel)
			ok.SetParent(m.UpdatePromptPanel)
			ok.SetCaption("提示")
			ok.SetOnClick(func(sender lcl.IObject) {
				fmt.Println("提示")
				m.UpdatePromptPanel.SetVisible(false)
				m.UpdateProgressPanel.SetVisible(true)
			})

			process := lcl.NewButton(m.UpdateProgressPanel)
			process.SetParent(m.UpdateProgressPanel)
			process.SetCaption("进度")
			process.SetOnClick(func(sender lcl.IObject) {
				fmt.Println("进度")
				m.UpdatePromptPanel.SetVisible(true)
				m.UpdateProgressPanel.SetVisible(false)
			})
		}
		// run and create update form
		lcl.RunApp(&updateForm)
	}
	// 检查 liblcl 库是否有更新
	autoupdate.CheckUpdate()
}
