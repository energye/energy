//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package progressbar

import (
	"fmt"
	"runtime"
	"strings"
)

const defaultCount = 50
const (
	s_bit = 1024
	s_kb  = s_bit * 1024
	s_mb  = s_kb * 1024
	s_gb  = s_mb * 1024
)

type Bar struct {
	totalValue         int
	currentValue       int
	currentRate        int
	progressGraph      string
	progressGraphTotal int
	backGraph          string
	isShowPercent      bool
	isShowRatio        bool
	isShowBar          bool
	notice             string
	progressEnds       Ends
	color              BarColor
}

type Ends struct {
	Start string
	End   string
}

func NewBar(totalVale int) *Bar {
	res := &Bar{
		totalValue:         totalVale,
		currentValue:       0,
		progressGraphTotal: defaultCount,
		progressGraph:      "█",
		backGraph:          " ",
		progressEnds:       Ends{Start: "[", End: "]"},
		isShowPercent:      true,
		isShowRatio:        true,
		isShowBar:          true,
	}
	res.CountCurrentRate()
	return res
}

func (m *Bar) SetEnds(start, end string) {
	ends := Ends{Start: start, End: end}
	m.progressEnds = ends
}

func (m *Bar) CountCurrentRate() {
	if m.currentValue == 0 {
		m.currentRate = 0
	} else {
		m.currentRate = m.currentValue * 100 / m.totalValue
	}
}

func (m *Bar) SetGraph(graph string) {
	m.progressGraph = graph
}

func (m *Bar) SetCurrentValue(value int) {
	m.currentValue = value
	m.CountCurrentRate()
}

func (m *Bar) CurrentPrintGraphNumber() int {
	if m.currentRate == 100 {
		return m.progressGraphTotal
	} else {
		return int(float64(m.currentRate) * (float64(m.progressGraphTotal) / float64(100)))
	}
}

func (m *Bar) HideProgressBar() {
	m.isShowBar = false
}

func (m *Bar) HidePercent() {
	m.isShowPercent = false
}

func (m *Bar) HideRatio() {
	m.isShowRatio = false
}

func (m *Bar) SetNotice(notice string) {
	m.notice = notice
}

func (m *Bar) SetProgressGraphTotal(totalGraph int) {
	m.progressGraphTotal = totalGraph
}

func (m *Bar) SetBackGraph(graph string) {
	m.backGraph = graph
}

func (m *Bar) PrintBar(currValue int) {
	m.SetCurrentValue(currValue)
	printStr := "\r" + m.NoticePrintString()
	if m.isShowBar {
		printStr += m.ProgressPrintString()
	}
	if m.isShowPercent {
		printStr += m.PercentPrintString()
	}
	if m.isShowRatio {
		printStr += m.RatioPrintString()
	}
	fmt.Print(printStr)
}

func (m *Bar) PrintSizeBar(size int64) {
	printStr := "\r" + m.NoticePrintString()
	var sizeStr = fmt.Sprintf("%f MB", float64(size)/s_kb)
	if runtime.GOOS != "windows" {
		printStr += fmt.Sprintf(" %c[%vm%v%c[0m", 0x1B, m.color.Notice, sizeStr, 0x1B)
	} else {
		printStr += sizeStr
	}
	fmt.Print(printStr)
}

func (m *Bar) NoticePrintString() string {
	if runtime.GOOS == "windows" {
		return m.notice
	} else {
		return fmt.Sprintf(" %c[%vm%v%c[0m", 0x1B, m.color.Notice, m.notice, 0x1B)
	}
}

func (m *Bar) ProgressPrintString() string {
	back := m.GetBackString()
	printStr := m.progressEnds.Start
	printStr += strings.Replace(back, m.backGraph, m.progressGraph, m.CurrentPrintGraphNumber())
	printStr += m.progressEnds.End
	if runtime.GOOS == "windows" {
		return printStr
	} else {
		return fmt.Sprintf("%c[%v;%vm%s%c[0m", 0x1B, m.color.Graph, m.color.Back, printStr, 0x1B)
	}
}

func (m *Bar) PercentPrintString() string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf(" %v%%", m.currentRate)
	}
	return fmt.Sprintf(" %c[%vm%v%%%c[0m", 0x1B, m.color.Percent, m.currentRate, 0x1B)
}

func (m *Bar) RatioPrintString() string {
	if runtime.GOOS == "windows" {
		return fmt.Sprintf(" %v/%v", m.currentValue, m.totalValue)
	}
	return fmt.Sprintf(" %c[%vm\t%v/%v%c[0m", 0x1B, m.color.Ratio, m.currentValue, m.totalValue, 0x1B)
}

func (m *Bar) PrintEnd(tip ...string) {
	fmt.Printf("\n")
	if len(tip) > 0 {
		fmt.Println(tip[0])
	}
}

func (m *Bar) GetBackString() string {
	res := ""
	for i := 0; i < m.progressGraphTotal; i++ {
		res += m.backGraph
	}
	return res
}
