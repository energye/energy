//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package progressbar

type BarColor struct {
	Graph   int
	Back    int
	Ratio   int
	Percent int
	Notice  int
}

type Color struct {
	Black  int
	Blue   int
	Green  int
	Aqua   int
	Red    int
	Purple int
	Yellow int
	White  int
	Gray   int
}

var (
	FontColor Color
	BackColor Color
)

func init() {
	initColor()
}

func initColor() {
	FontColor = Color{
		Black:  30,
		Blue:   34,
		Green:  32,
		Aqua:   36,
		Red:    31,
		Purple: 35,
		Yellow: 33,
		White:  37,
		Gray:   37,
	}
	BackColor = Color{
		Black:  40,
		Blue:   44,
		Green:  42,
		Aqua:   46,
		Red:    41,
		Purple: 45,
		Yellow: 43,
		White:  47,
		Gray:   47,
	}
}

func (m *Bar) SetProgressGraphColor(color int) {
	m.color.Graph = color
}

func (m *Bar) SetColor(color BarColor) {
	m.color = color
}

func (m *Bar) SetGraphColor(color int) {
	m.color.Graph = color
}

func (m *Bar) SetBackColor(color int) {
	m.color.Back = color
}

func (m *Bar) SetRatioColor(color int) {
	m.color.Ratio = color
}

func (m *Bar) SetPercentColor(color int) {
	m.color.Percent = color
}

func (m *Bar) SetNoticeColor(color int) {
	m.color.Notice = color
}
