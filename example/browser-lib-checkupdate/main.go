// auto update application

package main

import (
	"embed"
	"fmt"
	"github.com/energye/energy/v2/cef/autoupdate"
	"github.com/energye/energy/v2/common/imports"
	"github.com/energye/energy/v2/example/browser-lib-checkupdate/form"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/api/dllimports"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
	"runtime"
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
	// 开启自动更新检查
	autoupdate.IsCheckUpdate(true)
	// 如果 liblcl 有更新该函数将被回调
	autoupdate.CanUpdateLiblcl = func(model *autoupdate.Model, level int) {
		fmt.Println(*model)
		fmt.Println(model.Versions[model.Latest])
		fmt.Println(level)
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
			//m.SetFormStyle(types.FsSystemStayOnTop)
			m.SetPosition(types.PoDesktopCenter)
			//m.SetBorderStyle(types.BsSingle)
			m.SetBorderStyle(types.BsNone)
			//m.SetShowInTaskBar(types.StNever)
			m.SetWidth(290)
			m.SetHeight(390)
			m.SetCaption("lib-lcl 更新") // 自定义窗口标题

			// 自定义窗口标题栏
			m.TitlePanel = m.NewPanel()
			m.TitlePanel.SetColor(colors.ClTeal)
			m.TitlePanel.SetHeight(32)
			m.TitlePanel.SetWidth(m.Width())

			// title -> icon
			titleIcon := lcl.NewImage(m.TitlePanel)
			titleIcon.SetParent(m.TitlePanel)
			titleIcon.SetWidth(32)
			titleIcon.SetHeight(32)
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
			titleClose.SetLeft(250)
			titleClose.SetHint("关闭")
			titleClose.SetOnClick(func(lcl.IObject) {
				m.Close()
			})

			m.UpdatePromptPanel = m.NewPanel()
			m.UpdatePromptPanel.SetTop(m.TitlePanel.Height())
			m.UpdatePromptPanel.SetWidth(m.Width())
			m.UpdatePromptPanel.SetHeight(m.Height() - m.TitlePanel.Height())

			//m.UpdateProgressPanel = m.NewPanel()
			//m.UpdateProgressPanel.SetVisible(false)

			// background
			bgImage := lcl.NewImage(m.UpdatePromptPanel)
			bgImage.SetParent(m.UpdatePromptPanel)
			bgImage.SetWidth(277)
			bgImage.SetHeight(156)
			bgImage.Picture().LoadFromFSFile("resources/icon.png")

			// title
			//lcl.NewImage(m.UpdateProgressPanel)
			//

			//ok := lcl.NewButton(m.UpdatePromptPanel)
			//ok.SetParent(m.UpdatePromptPanel)
			//ok.SetCaption("提示")
			//ok.SetOnClick(func(sender lcl.IObject) {
			//	fmt.Println("提示")
			//	m.UpdatePromptPanel.SetVisible(false)
			//	m.UpdateProgressPanel.SetVisible(true)
			//})
			//
			//process := lcl.NewButton(m.UpdateProgressPanel)
			//process.SetParent(m.UpdateProgressPanel)
			//process.SetCaption("进度")
			//process.SetOnClick(func(sender lcl.IObject) {
			//	fmt.Println("进度")
			//	m.UpdatePromptPanel.SetVisible(true)
			//	m.UpdateProgressPanel.SetVisible(false)
			//})
		}
		// run and create update form
		lcl.RunApp(&updateForm)
	}
	// 检查 liblcl 库是否有更新
	autoupdate.CheckUpdate()
}
