// Copyright © yanghy. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and limitations under the License.

package wg

import (
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/types"
	"github.com/energye/lcl/types/colors"
	"math"
)

// 按钮方向
type TButtonBorderDirection = int32

const (
	BbdNone TButtonBorderDirection = iota
	BbdLeft
	BbdTop
	BbdRight
	BbdBottom
)

// 按钮方向集合
type TButtonBorderDirections = types.TSet

// TButtonColor 按钮颜色
type TButtonColor struct {
	start    colors.TColor     // 按钮起始渐变颜色
	end      colors.TColor     // 按钮结束渐变颜色
	Border   TButtonBorder     // 按钮边框
	img      lcl.ILazIntfImage // 缓存
	bitMap   lcl.IBitmap       // 缓存
	type_    int32             // 按钮类型, 自定义, 区分类型
	canPaint bool              // 是否绘制
}

// 按钮边框
type TButtonBorder struct {
	color       colors.TColor           // 按钮边框颜色, 启用边框方向才有作用
	colorLeft   colors.TColor           // 按钮边框颜色, Left
	colorTop    colors.TColor           // 按钮边框颜色, Top
	colorRight  colors.TColor           // 按钮边框颜色, Right
	colorBottom colors.TColor           // 按钮边框颜色, Bottom
	Direction   TButtonBorderDirections // 按钮显示边框方向
	width       int32                   // 边框宽度
	leftWidth   int32                   // 左宽度
	topWidth    int32                   // 上宽度
	rightWidth  int32                   // 右宽度
	bottomWidth int32                   // 下宽度
}

func NewButtonColor() *TButtonColor {
	m := &TButtonColor{
		img:    lcl.NewLazIntfImageWithIntX2RIQFlags(0, 0, types.NewSet(types.RiqfRGB, types.RiqfAlpha)),
		bitMap: lcl.NewBitmap(),
	}
	m.bitMap.SetPixelFormat(types.Pf32bit)
	return m
}

func (m *TButtonColor) Free() {
	if m.img != nil && m.img.IsValid() {
		m.img.Free()
	}
	if m.bitMap != nil && m.bitMap.IsValid() {
		m.bitMap.Free()
	}
}

// SetBorderWidth 设置按钮指定方向的边框宽度
// direction: 边框方向，指定要设置哪一边的边框宽度
// width: 边框宽度值，单位为像素
func (m *TButtonColor) SetBorderWidth(direction TButtonBorderDirection, width int32) {
	switch direction {
	case BbdLeft:
		m.Border.leftWidth = width
	case BbdTop:
		m.Border.topWidth = width
	case BbdRight:
		m.Border.rightWidth = width
	case BbdBottom:
		m.Border.bottomWidth = width
	default:
		m.Border.width = width
	}
	m.canPaint = true
}

// BorderWidth 根据指定的边框方向返回对应的边框宽度
//
//	direction: 边框方向枚举值，指定要获取哪个方向的边框宽度
//	int32: 指定方向的边框宽度值，如果方向无效则返回默认宽度
func (m *TButtonColor) BorderWidth(direction TButtonBorderDirection) (width int32) {
	switch direction {
	case BbdLeft:
		width = m.Border.leftWidth
	case BbdTop:
		width = m.Border.topWidth
	case BbdRight:
		width = m.Border.rightWidth
	case BbdBottom:
		width = m.Border.bottomWidth
	}
	if width == 0 {
		width = m.Border.width
	}
	return
}

// SetBorderColor 设置按钮边框颜色
// borderColor: 指定要设置的边框方向
// color: 要设置的边框颜色值
func (m *TButtonColor) SetBorderColor(direction TButtonBorderDirection, color colors.TColor) {
	switch direction {
	case BbdLeft:
		m.Border.colorLeft = color
	case BbdTop:
		m.Border.colorTop = color
	case BbdRight:
		m.Border.colorRight = color
	case BbdBottom:
		m.Border.colorBottom = color
	default:
		m.Border.color = color
	}
	m.canPaint = true
}

// BorderColor 根据指定的边框方向返回对应的边框颜色
// direction: 边框方向枚举值，指定要获取哪个方向的边框颜色
// colors.TColor: 指定方向的边框颜色值，如果方向无效则返回0
func (m *TButtonColor) BorderColor(direction TButtonBorderDirection) (color colors.TColor) {
	switch direction {
	case BbdLeft:
		color = m.Border.colorLeft
	case BbdTop:
		color = m.Border.colorTop
	case BbdRight:
		color = m.Border.colorRight
	case BbdBottom:
		color = m.Border.colorBottom
	}
	if color == 0 {
		color = m.Border.color
	}
	return
}

// CanPaint 检查按钮颜色是否可以绘制
// rect: 绘制区域的矩形尺寸
// 返回值: 当尺寸发生变化或允许绘制时返回true，否则返回false
func (m *TButtonColor) CanPaint(rect types.TRect) bool {
	isPaint := false
	if m.img.Width() != rect.Width() || m.img.Height() != rect.Height() {
		m.img.SetSize(rect.Width(), rect.Height())
		isPaint = true
	}
	if m.bitMap.Width() != rect.Width() || m.bitMap.Height() != rect.Height() {
		m.bitMap.SetSize(rect.Width(), rect.Height())
		isPaint = true
	}
	return isPaint || m.canPaint
}

// paint 绘制按钮颜色
// roundedCorners: 圆角设置
// rect: 绘制区域矩形
// alpha: 透明度值
// radius: 圆角半径
func (m *TButtonColor) tryPaint(roundedCorners TRoundedCorners, rect types.TRect, alpha byte, radius int32) {
	if !m.CanPaint(rect) {
		return
	}
	m.canPaint = false
	m.doPaint(roundedCorners, rect, alpha, radius)
}

// doPaint 绘制带有圆角和透明度的垂直渐变按钮图像。
// 参数:
//
//	roundedCorners: 指定哪些角落需要绘制为圆角
//	alpha: 图像的整体透明度，取值范围 0-255
//	radius: 圆角的半径大小
func (m *TButtonColor) doPaint(roundedCorners TRoundedCorners, rect types.TRect, alpha byte, radius int32) {
	w, h := rect.Width(), rect.Height()
	// 提取起始颜色和结束颜色的 RGB 分量，用于计算渐变过程中的颜色插值
	startR := colors.Red(m.start)
	startG := colors.Green(m.start)
	startB := colors.Blue(m.start)
	// 获取结束颜色分量
	endR := colors.Red(m.end)
	endG := colors.Green(m.end)
	endB := colors.Blue(m.end)
	// 处理垂直渐变（带抗锯齿圆角）
	// 遍历图像每一行，根据当前行位置计算颜色渐变比例，并逐像素设置颜色与透明度
	imgHeight := m.img.Height()
	imgWidth := m.img.Width()
	for y := int32(0); y < imgHeight; y++ {
		// 计算颜色渐变
		ratio := float64(y) / float64(imgHeight-1)
		r := round(float64(startR)*(1-ratio) + float64(endR)*ratio)
		g := round(float64(startG)*(1-ratio) + float64(endG)*ratio)
		b := round(float64(startB)*(1-ratio) + float64(endB)*ratio)
		color := lcl.TFPColor{Red: uint16(r) << 8, Green: uint16(g) << 8, Blue: uint16(b) << 8}
		borderFPColor := color
		// 注意：Alpha会在内循环中为每个像素单独设置
		for x := int32(0); x < imgWidth; x++ {
			alphaFactor, corners := m.calculateRoundedAlpha(roundedCorners, x, y, imgWidth, imgHeight, radius)
			_ = corners
			isBorder := false
			borderColor := colors.TColor(0)
			if m.Border.Direction != 0 { // 启用了边框
				if alphaFactor < 1.0 { // 圆角范围内
					// TODO 有点小问题. 1: 圆角的边框绘制有瑕疵, 2: 圆角的边框宽度判断未增加
					if m.Border.Direction.In(BbdLeft) && (corners == RcLeftTop || corners == RcLeftBottom) {
						// 左边框
						borderColor = m.BorderColor(BbdLeft)
						isBorder = true
					} else if m.Border.Direction.In(BbdTop) && (corners == RcRightTop || corners == RcLeftTop) {
						// 上边框
						borderColor = m.BorderColor(BbdTop)
						isBorder = true
					} else if m.Border.Direction.In(BbdRight) && (corners == RcRightTop || corners == RcRightBottom) {
						// 右边框
						borderColor = m.BorderColor(BbdRight)
						isBorder = true
					} else if m.Border.Direction.In(BbdBottom) && (corners == RcLeftBottom || corners == RcRightBottom) {
						// 下边框
						borderColor = m.BorderColor(BbdBottom)
						isBorder = true
					}
				} else { // 圆角外
					if m.Border.Direction.In(BbdLeft) && x < m.BorderWidth(BbdLeft) && y < h {
						// 左边框
						isBorder = true
						borderColor = m.BorderColor(BbdLeft)
					} else if m.Border.Direction.In(BbdTop) && y < m.BorderWidth(BbdTop) && x <= w {
						// 上边框
						isBorder = true
						borderColor = m.BorderColor(BbdTop)
					} else if m.Border.Direction.In(BbdRight) && x >= w-m.BorderWidth(BbdRight) && y < h {
						// 右边框
						isBorder = true
						borderColor = m.BorderColor(BbdRight)
					} else if m.Border.Direction.In(BbdBottom) && x < w && y >= h-m.BorderWidth(BbdBottom) {
						// 下边框
						isBorder = true
						borderColor = m.BorderColor(BbdBottom)
					}
				}
				if borderColor != 0 {
					borderFPColor = ColorToFPColor(borderColor, ratio)
				}
			}
			actualAlpha := round(float64(alpha) * float64(alphaFactor))
			if isBorder {
				borderFPColor.Alpha = uint16(actualAlpha) << 8
				m.img.SetColors(x, y, borderFPColor)
			} else {
				color.Alpha = uint16(actualAlpha) << 8
				m.img.SetColors(x, y, color)
			}
		}
	}
	// 将处理好的图像数据加载到位图对象中，供后续使用
	m.bitMap.LoadFromIntfImage(m.img)
}

func (m *TButtonColor) SetColor(start, end colors.TColor) {
	m.start = start
	m.end = end
	m.canPaint = true
}

// calculateRoundedAlpha 根据给定的圆角信息和像素位置，计算该点在按钮背景中的 alpha 值以及所属的圆角类型。
// 用于实现抗锯齿效果的圆角绘制逻辑。
//
// 参数说明：
//
//	roundedCorners: 指示哪些角落需要绘制为圆角（TRoundedCorners 类型）
//	x, y: 当前像素点相对于控件左上角的坐标
//	width, height: 控件的宽高尺寸
//	radius: 预期设置的圆角半径
//
// 返回值说明：
//
//	alphaValue: 当前点的透明度值，范围 [0.0, 1.0]
//	corners: 所属的圆角类型（如 RcLeftTop、RcRightBottom 等），若不在圆角区域则为默认值
func (m *TButtonColor) calculateRoundedAlpha(roundedCorners TRoundedCorners, x, y, width, height, radius int32) (alphaValue float32, corners RoundedCorner) {
	// 计算实际可用最大半径（不超过尺寸限制）
	maxRadius := min(width/2, height/2)
	if radius > maxRadius {
		radius = maxRadius
	}
	// 如果半径被限制为0，直接返回不透明
	if radius <= 0 {
		alphaValue = 1.0
		return
	}
	var (
		cornerX, cornerY int32   // 圆角顶点坐标
		d                float32 // 距离圆心的距离
		inCorner         bool    // 是否在圆角内
	)
	// 判断当前点是否位于某个圆角区域内，并计算其到对应圆心的距离
	// 左上角区域
	if roundedCorners.In(RcLeftTop) && x < radius && y < radius {
		cornerX = radius
		cornerY = radius
		inCorner = true
		corners = RcLeftTop
	} else if roundedCorners.In(RcRightTop) && x >= width-radius && y < radius {
		// 右上角区域
		cornerX = width - radius - 1
		cornerY = radius
		inCorner = true
		corners = RcRightTop
	} else if roundedCorners.In(RcLeftBottom) && x < radius && y >= height-radius {
		// 左下角区域
		cornerX = radius
		cornerY = height - radius - 1
		inCorner = true
		corners = RcLeftBottom
	} else if roundedCorners.In(RcRightBottom) && x >= width-radius && y >= height-radius {
		// 右下角区域
		cornerX = width - radius - 1
		cornerY = height - radius - 1
		inCorner = true
		corners = RcRightBottom
	}
	if !inCorner {
		alphaValue = 1.0
		return
	}
	d = sqrt(float64(sqr(x-cornerX) + sqr(y-cornerY)))
	// 抗锯齿过渡处理：根据距离决定 alpha 渐变值
	const transition = 1.0
	innerRadius := float32(radius) - transition
	// 完全在圆角内
	if d <= innerRadius {
		alphaValue = 1.0
		return
	}
	// 完全在圆角外
	if d >= float32(radius)+transition {
		alphaValue = 0.0
		return
	}
	// 在过渡区域内（平滑渐变）
	alphaValue = 1.0 - (d-innerRadius)/(2*transition)
	return
}

// 辅助函数：整数最小值
func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// DarkenColor 函数用于将给定的颜色按照指定因子进行暗化处理
// 参数:
//
//	color: 原始颜色值，类型为 types.TColor
//	factor: 暗化因子，取值范围通常为 0.0-1.0，值越大颜色越暗
//
// 返回值:
//
//	返回暗化后的颜色值，类型为 types.TColor
func DarkenColor(color types.TColor, factor float64) types.TColor {
	R := colors.Red(color)
	G := colors.Green(color)
	B := colors.Blue(color)

	R = byte(round(float64(R) * (1.0 - factor)))
	G = byte(round(float64(G) * (1.0 - factor)))
	B = byte(round(float64(B) * (1.0 - factor)))
	return colors.RGBToColor(R, G, B)
}

// LightenColor 函数用于将给定的颜色按照指定因子进行亮化处理
// 参数:
//
//	color: 原始颜色值，类型为 types.TColor
//	factor: 亮化因子，取值范围通常为 0.0-1.0，值越大颜色越亮
//
// 返回值:
//
//	返回亮化后的颜色值，类型为 types.TColor
func LightenColor(color types.TColor, factor float64) types.TColor {
	R := colors.Red(color)
	G := colors.Green(color)
	B := colors.Blue(color)
	R = byte(round(float64(R) + (255.0-float64(R))*factor))
	G = byte(round(float64(G) + (255.0-float64(G))*factor))
	B = byte(round(float64(B) + (255.0-float64(B))*factor))
	return colors.RGBToColor(R, G, B)
}

// GrayColor 函数用于将给定的颜色转换为灰度颜色
// 参数:
//
//	color: 原始颜色值，类型为 types.TColor
//
// 返回值:
//
//	返回灰度颜色值，类型为 types.TColor
func GrayColor(color types.TColor) types.TColor {
	R := colors.Red(color)
	G := colors.Green(color)
	B := colors.Blue(color)
	// 使用标准的灰度转换公式: Gray = 0.299*R + 0.587*G + 0.114*B
	gray := byte(round(0.299*float64(R) + 0.587*float64(G) + 0.114*float64(B)))
	return colors.RGBToColor(gray, gray, gray)
}

// DarkenFPColor 通过给定因子降低颜色的亮度
// color: 指向TFPColor结构体的指针，表示要处理的颜色
// factor: 浮点数，表示变暗因子，取值范围通常在0.0-1.0之间，值越大颜色越暗
func DarkenFPColor(color *lcl.TFPColor, factor float64) {
	r := uint16(float64(color.Red) * (1.0 - factor))
	g := uint16(float64(color.Green) * (1.0 - factor))
	b := uint16(float64(color.Blue) * (1.0 - factor))
	color.Red = r
	color.Green = g
	color.Blue = b
}

func ColorToFPColor(color types.TColor, ratio float64) lcl.TFPColor {
	colorR := colors.Red(color)
	colorG := colors.Green(color)
	colorB := colors.Blue(color)
	r := round(float64(colorR)*(1-ratio) + float64(colorR)*ratio)
	g := round(float64(colorG)*(1-ratio) + float64(colorG)*ratio)
	b := round(float64(colorB)*(1-ratio) + float64(colorB)*ratio)
	return lcl.TFPColor{Red: uint16(r) << 8, Green: uint16(g) << 8, Blue: uint16(b) << 8}
}

func round(v float64) float64 {
	return math.Round(v)
}

func sqr(x int32) int32 {
	return x * x
}

func sqrt(v float64) float32 {
	return float32(math.Sqrt(v))
}
