package main

import (
	"fmt"
	"github.com/energye/energy/v2/api"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/pkgs/win"
	"github.com/energye/energy/v2/pkgs/winapi"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/colors"
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)
	lcl.Application.CreateForm(&MainForm, true)
	lcl.Application.Run()
}

type TMainForm struct {
	*lcl.TForm
}

var MainForm *TMainForm

func (m *TMainForm) OnFormCreate(sender lcl.IObject) {
	//m.SetBorderStyleForFormBorderStyle(types.BsNone) // 边框窗口无边框
	m.SetColor(colors.ClWhite) // 边框窗口的填充颜色, 做为透明色
	//m.SetFormStyle(types.FsStayOnTop)
	m.ScreenCenter()
	//m.Paint()
	//winapi.SetWindowLongPtr(et.HWND(m.Handle()), win.GWL_EXSTYLE, winapi.GetWindowLongPtr(et.HWND(m.Handle()), win.GWL_EXSTYLE)^win.WS_EX_CLIENTEDGE) // 移除默认的边框样式
	//winapi.SetWindowPos(et.HWND(m.Handle()), 0, 0, 0, 0, 0, win.SWP_FRAMECHANGED)                                                                     // 重新绘制窗口边框
	//hwnd := m.Handle()
	//gwlStyle := win.GetWindowLong(hwnd, win.GWL_STYLE)
	//win.SetWindowLong(hwnd, win.GWL_STYLE, uintptr(gwlStyle&^win.WS_CAPTION&^win.WS_THICKFRAME|win.WS_BORDER))
	//win.SetWindowPos(hwnd, 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))

	//// 设置窗口扩展样式以实现透明效果和边框阴影颜色设置
	//win.SetWindowLong(hwnd, win.GWL_EXSTYLE, uintptr(win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED|win.WS_EX_TRANSPARENT))
	//// 设置窗口的透明色为红色（RGB值为255,0,0）
	//win.SetLayeredWindowAttributes(hwnd, 0x00FF0000, 255, win.LWA_COLORKEY)

	gwlStyle := win.GetWindowLong(m.Handle(), win.GWL_STYLE)
	win.SetWindowLong(m.Handle(), win.GWL_STYLE, uintptr(gwlStyle^win.WS_CAPTION^win.WS_BORDER^win.WS_MAXIMIZEBOX^win.WS_MINIMIZEBOX^win.WS_SIZEBOX|win.WS_THICKFRAME))
	win.SetWindowPos(m.Handle(), 0, 0, 0, 0, 0, uint32(win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED))

	m.SetOnPaint(func(sender lcl.IObject) {
		fmt.Println("paint")
		//m.SetClientHeight(m.ClientHeight() + 8)
	})
}

// CreateParams
func (m *TMainForm) CreateParams(params *types.TCreateParams) {
	fmt.Println(params.Style, api.GoStr(params.Caption))
	fmt.Println(params.ExStyle)
	//params.Style = params.Style & win.WS_MAXIMIZEBOX
	//params.Style = params.Style&^win.WS_CAPTION&^win.WS_BORDER | win.WS_THICKFRAME
}

// Paint 根据设置的border颜色和四边圆角画边框, 在Show之前调用该函数
func (m *TMainForm) Paint() {
	//var borderColors = []types.TColor{0xFEFEFE, 0xFDFDFD, 0xFCFCFC, 0xFBFBFB, 0xFAFAFA, 0xF7F7F7, 0xEDEDED, 0xE3E3E3, 0xD9D9D9}
	var bitmap = lcl.NewBitmap()
	bitmap.SetTransparent(false)
	var png = lcl.NewPngImage()
	png.LoadFromFile("D:\\Energy-Doc\\Go-Energy.png")
	bitmap.Assign(png)
	m.SetOnPaint(func(sender lcl.IObject) {
		//r := m.ClientRect()
		canvas := m.Canvas()
		lcl.AsBrush(canvas.Brush()).SetBitmap(bitmap) //.Bitmap().SetTransparent(false)
		//canvas.Brush().SetStyle(types.BsClear)
		//pen := canvas.Pen()
		//pen.SetWidth(1)
		////pen.SetMode(types.PmNotMask)
		////pen.SetStyle(types.PsClear)
		//for i := 1; i <= len(borderColors); i++ {
		//	pen.SetColor(borderColors[i-1])
		//	canvas.Rectangle(r.Left+int32(i), r.Top+int32(i), r.Right-int32(i), r.Bottom-int32(i))
		//}
	})
	//m.framelessWindowBroderTransparent(et.HWND(m.Handle()))
}

// framelessWindowBroderTransparent 无边框窗口自定义窗口透明
func (m *TMainForm) framelessWindowBroderTransparent(hWnd types.HWND) {
	exStyle := winapi.GetWindowLong(hWnd, win.GWL_EXSTYLE)
	exStyle = exStyle | win.WS_EX_LAYERED
	winapi.SetWindowLong(hWnd, win.GWL_EXSTYLE, exStyle)
	win.SetLayeredWindowAttributes(hWnd, colors.ClWhite, 100, win.LWA_ALPHA|win.LWA_COLORKEY)
}
