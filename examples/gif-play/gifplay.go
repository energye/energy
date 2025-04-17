package main

import (
	"embed"
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef/gifanim"
	_ "github.com/cyber-xxm/energy/v2/examples/syso"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/energy/inits"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/pkgs/macapp"
)

type TMainForm struct {
	*lcl.TForm
}

var (
	mainForm *TMainForm
)

//go:embed resources
var resources embed.FS

func main() {
	lcl.DEBUG = true
	macapp.MacApp.SetEnergyEnv("dev") // MacOS 开发时配置，在初始化之前
	inits.Init(nil, resources)
	lcl.Application.SetOnException(func(sender lcl.IObject, e *lcl.Exception) {
		fmt.Println("Exception:", e.Message())
	})
	lcl.RunApp(&mainForm)
}

func (m *TMainForm) OnFormCreate(sender lcl.IObject) {
	m.SetWidth(730)
	m.SetHeight(650)
	m.WorkAreaCenter()
	// 这是透明设置，只能form, 它的子组件也会跟着透明
	//m.SetAlphaBlend(true)     //透明
	//m.SetAlphaBlendValue(150) //透明度

	gifData, _ := emfs.GetResources("resources/loading.gif")
	m.NewGIFAnimate(gifData, 0, 0)
	m.NewGIFAnimate(gifData, 440, 0)
	m.NewGIFAnimate(gifData, 0, 450)
	m.NewGIFAnimate(gifData, 440, 450)
}

func (m *TMainForm) NewGIFAnimate(gifFileData []byte, left, top int32) {
	animate := gifanim.NewGIFAnimate(m)
	animate.SetParent(m)
	animate.SetAnimate(true)
	animate.SetLeft(left)
	animate.SetTop(top)
	animate.LoadFromBytes(gifFileData)
	animate.SetOnStop(func() {
		fmt.Println("stop:", animate.CurrentFrameIndex())
	})
	animate.SetOnStart(func() {
		fmt.Println("start:", animate.CurrentFrameIndex())
	})
	// 操作
	preBtn := lcl.NewButton(m)
	preBtn.SetParent(m)
	preBtn.SetCaption("上一帧")
	preBtn.SetTop(animate.Height() + top + 10)
	preBtn.SetLeft(left)
	preBtn.SetOnClick(func(sender lcl.IObject) {
		animate.PrevFrame()
	})

	nextBtn := lcl.NewButton(m)
	nextBtn.SetParent(m)
	nextBtn.SetCaption("下一帧")
	nextBtn.SetTop(animate.Height() + top + 10)
	nextBtn.SetLeft(left + 100)
	nextBtn.SetOnClick(func(sender lcl.IObject) {
		animate.NextFrame()
	})

	startBtn := lcl.NewButton(m)
	startBtn.SetParent(m)
	startBtn.SetCaption("开始 / 停止")
	startBtn.SetTop(animate.Height() + top + 10)
	startBtn.SetLeft(left + 200)
	startBtn.SetOnClick(func(sender lcl.IObject) {
		animate.SetAnimate(!animate.Animate())
	})
}
