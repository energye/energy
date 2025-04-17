package main

import (
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/ipc"
	"github.com/cyber-xxm/energy/v2/consts"
	demoCommon "github.com/cyber-xxm/energy/v2/examples/common"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/rtl"
	"github.com/energye/golcl/lcl/types"
	"path/filepath"
)

func main() {
	//Global initialization must be called
	cef.GlobalInit(nil, demoCommon.ResourcesFS())
	//Create an application
	app := cef.NewApplication()
	//Local load resources

	cef.BrowserWindow.Config.LocalResource(cef.LocalLoadConfig{
		ResRootDir: "resources",
		FS:         demoCommon.ResourcesFS(),
		Home:       "screenshot.html",
	}.Build())
	cef.BrowserWindow.Config.Width = 600
	cef.BrowserWindow.Config.Height = 400
	// run main process and main thread
	cef.BrowserWindow.SetBrowserInit(browserInit)
	//run app
	cef.Run(app)
}

// run main process and main thread
func browserInit(event *cef.BrowserEvent, window cef.IBrowserWindow) {
	var (
		schotForm *lcl.TForm
		image     *lcl.TImage
	)
	if window.IsLCL() {
		// 创建一个窗口显示截屏图片
		schotForm = lcl.NewForm(window.AsLCLBrowserWindow().BrowserWindow())
		// 窗口透明
		schotForm.SetAlphaBlend(true)
		// 无边框窗口
		//schotForm.SetBorderStyle(types.BsNone)
		// 窗口透明度
		//schotForm.SetAlphaBlendValue(155)
		// 窗口大小是整个显示器大小
		//schotForm.SetBoundsRect(window.AsLCLBrowserWindow().BrowserWindow().Monitor().BoundsRect())
		// 显示截屏图片
		image = lcl.NewImage(schotForm)
		image.SetParent(schotForm)
		image.SetAlign(types.AlClient)
		// 可以使用一些事件来处理截图.
		image.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
			println("MouseMove")
		})
		image.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {

		})
		image.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {

		})
	}
	// 屏幕截图
	ipc.On("screenshot", func() {
		println("screenshot")
		window.RunOnMainThread(func() {
			dc := rtl.GetDC(0)
			// 一定要释放
			defer rtl.ReleaseDC(0, dc)

			// 位图
			bmp := lcl.NewBitmap()
			defer bmp.Free()
			bmp.LoadFromDevice(dc)

			// 可以 保存到本地 bmp
			wd := consts.CurrentExecuteDir
			bmp.SaveToFile(filepath.Join(wd, "sc.bmp"))

			// 可以 显示截屏图片窗口
			if schotForm != nil {
				schotForm.Show()
				image.Picture().Assign(bmp)
			}
		})
	})
}
