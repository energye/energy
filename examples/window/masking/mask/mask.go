package mask

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef"
	"github.com/cyber-xxm/energy/v2/cef/gifanim"
	"github.com/energye/golcl/energy/emfs"
	"github.com/energye/golcl/lcl"
	"github.com/energye/golcl/lcl/types"
	"github.com/energye/golcl/lcl/types/colors"
)

type Mask struct {
	maskForm      *lcl.TForm
	gifPlay       *gifanim.TGIFAnimate
	progressLabel *lcl.TLabel
}

func Create(window *cef.LCLBrowserWindow) *Mask {
	var mask = new(Mask)
	// 创建一个 form 或 panel、或其它任意 IWinControl 组件模拟遮罩
	// form 可以设置透明度
	mask.maskForm = lcl.NewForm(window)        // form有窗口句柄
	mask.maskForm.SetParent(window)            // 显示在浏览器窗口组件里的一个子窗口
	mask.maskForm.SetBorderStyle(types.BsNone) // 因为是窗口所以要去掉标签栏，效果和panel差不多了
	mask.maskForm.SetAlign(types.AlClient)     //铺满整个主窗口
	//mask.maskForm.SetColor(colors.ClSkyblue)
	// 这是透明设置，只能form, 但是它的子组件也会跟着半透明
	mask.maskForm.SetAlphaBlend(true)     //透明 MacOS好像没效果啊
	mask.maskForm.SetAlphaBlendValue(200) //透明度
	mask.maskForm.BringToFront()
	// 创建一个gif播放组件
	mask.gifPlay = gifanim.NewGIFAnimate(mask.maskForm)
	mask.gifPlay.SetParent(mask.maskForm)
	mask.gifPlay.SetAnimate(true)
	//在内置FS中读取gif资源
	data, _ := emfs.GetResources("resources/loading.gif")
	//play.LoadFromFile("本地加载") //或本地加载
	mask.gifPlay.LoadFromBytes(data)
	mask.gifPlay.SetOnFrameChanged(func(frame *gifanim.Frame) {
		fmt.Println("OnFrameChanged CurrentImageIndex:", frame.Index(), mask.gifPlay.CurrentFrameIndex())
	})
	//关闭时释放掉 play 占内存啊
	mask.maskForm.SetOnClose(func(sender lcl.IObject, action *types.TCloseAction) {
		mask.gifPlay.Free()
	})
	// 进度显示
	mask.progressLabel = lcl.NewLabel(mask.maskForm)
	mask.progressLabel.SetParent(mask.maskForm)
	mask.progressLabel.Font().SetColor(colors.ClRed)
	mask.progressLabel.Font().SetSize(18)
	// 居中显示
	var center = func(sender lcl.IObject) {
		mask.gifPlay.SetLeft(window.Width()/2 - mask.gifPlay.Width()/2)
		mask.gifPlay.SetTop(window.Height()/2 - mask.gifPlay.Height()/2)
		mask.progressLabel.SetLeft(window.Width()/2 + 80)
		mask.progressLabel.SetTop(window.Height()/2 - 20)
	}
	mask.maskForm.SetOnShow(center)
	mask.maskForm.SetOnResize(center)
	return mask
}

func (m *Mask) Progress(v int) {
	m.progressLabel.SetCaption(fmt.Sprintf("加载进度: %v", v))
}

func (m *Mask) Show() {
	m.maskForm.Show()
}

func (m *Mask) Hide() {
	m.maskForm.Hide()
}

func (m *Mask) IsValid() bool {
	return m.maskForm.IsValid()
}

func (m *Mask) Stop() {
	m.gifPlay.Stop()
}

func (m *Mask) Start() {
	m.gifPlay.Start()
}
