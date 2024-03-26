package main

import (
	"fmt"
	"github.com/energye/energy/v2/examples/lcl/trayicon/data"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/types"
	"runtime"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)

	mMenu := lcl.NewMainMenu(mainForm)
	mmItem := lcl.NewMenuItem(mainForm)
	mmItem.SetCaption("File")
	mMenu.Items().Add(mmItem)

	trayicon := lcl.NewTrayIcon(mainForm)

	btn := lcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetCaption("ShowBalloon")
	btn.SetLeft(20)
	btn.SetTop(30)
	btn.SetWidth(100)
	btn.SetOnClick(func(lcl.IObject) {
		trayicon.SetBalloonTitle("test")
		trayicon.SetBalloonTimeout(1000)
		trayicon.SetBalloonHint("我是提示正文啦")
		trayicon.ShowBalloonHint()
	})

	pm := lcl.NewPopupMenu(mainForm)
	item := lcl.NewMenuItem(mainForm)
	item.SetCaption("显示(&S)")
	item.SetOnClick(func(lcl.IObject) {

		mainForm.Show()
		// Windows上为了最前面显示，有时候要调用SetForegroundWindow
		//rtl.SetForegroundWindow(mainForm.Handle())
		//lcl.Application.Restore()
		//lcl.Application.BringToFront()
	})
	pm.Items().Add(item)

	item = lcl.NewMenuItem(mainForm)
	item.SetCaption("退出(&E)")
	item.SetOnClick(func(lcl.IObject) {
		// 主窗口关闭
		mainForm.Close()
		// 或者使用
		//		lcl.Application.Terminate()
	})
	pm.Items().Add(item)
	trayicon.SetPopUpMenu(pm)
	// lcl库得手指定，在windows下，如果实例资源中存在一个名为 MAINICON 的图标资源，则会自动加载，下面只是应对于linux与macOS下

	if runtime.GOOS != "windows" {
		// 这是使用文件加载方法，不考虑外部文件的话，可以用新的内存方法加载
		//icon := lcl.NewIcon()
		//icon.LoadFromFile(rtl.ExtractFilePath(lcl.Application.ExeName()) + "/2.ico")
		//trayicon.SetIcon(icon)
		//icon.Free()
		// 将图标应用到Application中的Icon中，到时候随时可以使用
		// 但也可不使用
		//loadMainIconFromStream(lcl.Application.Icon())
		loadMainIconFromStream(trayicon.Icon())

	} else {
		//trayicon.SetIcon(lcl.Application.Icon())
	}

	trayicon.SetHint(mainForm.Caption())
	trayicon.SetVisible(true)

	// 捕捉最小化
	lcl.Application.SetOnMinimize(func(sender lcl.IObject) {
		mainForm.Hide() // 主窗口最隐藏掉
	})

	// 这里写啥好呢，macOS下似乎这些事件跟PopupMenu有冲突
	if runtime.GOOS != "darwin" {
		trayicon.SetOnDblClick(func(lcl.IObject) {
			// macOS似乎不支持双击
			trayicon.SetBalloonTitle("test")
			trayicon.SetBalloonTimeout(10000)
			trayicon.SetBalloonHint("我是提示正文啦")
			trayicon.ShowBalloonHint()
			fmt.Println("TrayIcon DClick.")
		})
	}

	// 托盘图片可闪烁 1 秒闪一次
	tmr1 := lcl.NewTimer(mainForm)
	tmr1.SetOnTimer(func(sender lcl.IObject) {
		if trayicon.Icon().Empty() {
			trayicon.SetIcon(lcl.Application.Icon())
		} else {
			trayicon.SetIcon(nil)
		}
	})
	tmr1.SetEnabled(true)

	//加载其他格式的ico方式一
	lcl.NewPortableNetworkGraphic()
	//png := lcl.NewPortableNetworkGraphic()
	//png.LoadFromFile("bow.png")
	//trayicon.Icon().Assign(png)
	//png.Free()
	//方式二可以通过imagelist操作
	//imglist := lcl.NewImageList(mainForm)
	//png := lcl.NewPortableNetworkGraphic()
	//png.LoadFromFile("bow.png")
	//imglist.Add(png, nil)
	//png.Free()
	//imglist.GetIcon(0, trayicon.Icon(), types.GdeNormal)

	lcl.Application.Run()
}

// 主要是用于linux跟macOS下，因为不能像Windows一样直接内置到资源中
func loadMainIconFromStream(outIcon lcl.IIcon) {
	if outIcon.IsValid() {
		//mem := lcl.NewMemoryStreamFromBytes(mainIconBytes)
		//defer mem.Free() // 不要在阻塞的时候使用defer不然会一直到阻塞结束才释放，这里使用是因为这个函数结束了就释放了
		//mem.SetPosition(0)
		//outIcon.LoadFromStream(mem)
		outIcon.LoadFromBytes(data.MainIconBytes)
	}
}
