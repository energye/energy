package main

import (
	"fmt"
	_ "github.com/energye/energy/v2/examples/syso"
	"github.com/energye/energy/v2/inits"
	"github.com/energye/energy/v2/lcl"
	"github.com/energye/energy/v2/lcl/rtl"
	"github.com/energye/energy/v2/types"
	"github.com/energye/energy/v2/types/colors"
)

type TPoint struct {
	X, Y int32
	Down bool
}

var (
	isMouseDown bool
	points      = make([]TPoint, 0)
)

func main() {
	inits.Init(nil, nil)
	lcl.Application.Initialize()
	lcl.Application.SetMainFormOnTaskBar(true)

	jpgFileName := "1.jpg"
	canLoad := rtl.FileExists(jpgFileName)
	var jpgimg lcl.IJPEGImage
	if canLoad {
		jpgimg = lcl.NewJPEGImage()
		jpgimg.LoadFromFile(jpgFileName)
	}

	mainForm := lcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(600)
	mainForm.SetHeight(600)
	mainForm.SetDoubleBuffered(true)
	mainForm.SetOnDestroy(func(sender lcl.IObject) {
		if jpgimg != nil {
			jpgimg.Free()
		}
	})

	mainForm.SetOnPaint(func(lcl.IObject) {
		canvas := mainForm.Canvas()
		canvas.MoveTo(10, 10)
		canvas.LineTo(50, 10)
		s := "这是一段文字"
		font := lcl.AsFont(canvas.Font())
		brush := lcl.AsBrush(canvas.Brush())
		pen := lcl.AsPen(canvas.Pen())
		font.SetColor(colors.ClRed) // red
		font.SetSize(20)
		//style := font.Style()
		canvas.Brush().SetStyle(types.BsClear)
		canvas.TextOut(100, 30, s)
		fmt.Println("canvas.Font()", font.Height(), canvas.Font().Size())

		r := &types.TRect{0, 0, 80, 80}

		// 计算文字
		//fmt.Println("TfSingleLine: ", types.TfSingleLine)
		s = "由于现有第三方的Go UI库不是太庞大就是用的不习惯，或者组件太少。"
		canvas.TextRect2(r, s, types.NewSet(types.TfCenter, types.TfVerticalCenter, types.TfSingleLine))
		//fmt.Println("r: ", r, ", s: ", s)

		s = "测试输出"
		r = &types.TRect{0, 0, 80, 80}
		// brush
		brush.SetColor(colors.ClGreen)
		canvas.FillRect(r)

		// font
		font.SetStyle(0)
		font.SetSize(9)
		font.SetColor(colors.ClBlue)

		// pen
		pen.SetColor(colors.ClFuchsia)
		canvas.Rectangle1(r.Left, r.Top, r.Right, r.Bottom)

		textFmt := types.NewSet(types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
		canvas.TextRect2(r, s, textFmt)

		if jpgimg != nil {
			canvas.DrawForGraphic(0, 80, jpgimg)
		}

		// 画多边形
		brush.SetColor(colors.ClYellow)
		//canvas.Polygon([]types.TPoint{{15, 40}, {43, 123}, {81, 42}, {45, 11}}, false)
		//canvas.Polyline([]types.TPoint{{15 + 100, 40}, {43 + 100, 123}, {81 + 100, 42}, {45 + 100, 11}})
	})

	paintbox := lcl.NewPaintBox(mainForm)
	paintbox.SetParent(mainForm)
	paintbox.SetAlign(types.AlBottom)
	paintbox.SetHeight(mainForm.Height() - 280)
	//paintbox.SetColor(colors.ClRed)
	paintbox.SetOnPaint(func(lcl.IObject) {
		canvas := paintbox.Canvas()
		pen := lcl.AsPen(canvas.Pen())
		pen.SetColor(colors.ClGreen)
		font := lcl.AsFont(canvas.Font())
		r := paintbox.ClientRect()
		canvas.Rectangle(&r)

		font.SetColor(colors.ClSkyblue)
		rect := paintbox.ClientRect()
		s := "在这可以用鼠标绘制"
		textFmt := types.NewSet(types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
		canvas.TextRect2(&rect, s, textFmt)

		for _, p := range points {
			if p.Down {
				canvas.MoveTo(p.X, p.Y)
			} else {
				canvas.LineTo(p.X, p.Y)
			}
		}

	})

	paintbox.SetOnMouseDown(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("mouse down")
		if button == types.MbLeft {
			points = append(points, TPoint{X: x, Y: y, Down: true})
			isMouseDown = true
		}
	})

	paintbox.SetOnMouseMove(func(sender lcl.IObject, shift types.TShiftState, x, y int32) {
		if isMouseDown {
			points = append(points, TPoint{X: x, Y: y, Down: false})
			paintbox.Repaint()
		}
	})

	paintbox.SetOnMouseUp(func(sender lcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("mouse SetOnMouseUp")
		if button == types.MbLeft {
			isMouseDown = false
		}
	})

	btnClear := lcl.NewButton(mainForm)
	btnClear.SetParent(mainForm)
	btnClear.SetCaption("清除绘制")
	btnClear.SetLeft(mainForm.Width() - btnClear.Width() - 20)
	btnClear.SetTop(10)
	btnClear.SetOnClick(func(lcl.IObject) {
		points = make([]TPoint, 0)
		paintbox.Repaint()
	})

	lcl.Application.Run()
}
