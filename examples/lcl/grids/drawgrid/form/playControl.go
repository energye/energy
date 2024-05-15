package form

import (
	"fmt"

	. "github.com/energye/energy/v2/lcl"
	. "github.com/energye/energy/v2/types"
)

type TPlayListItem struct {
	Caption  string // 歌曲名
	Singer   string // 歌手
	Length   int32  // 歌曲长度
	Lyric    string // 歌词文件
	FileName string // 歌曲文件名
}

type TPlayControl struct {
	IDrawGrid
	datas          []TPlayListItem
	focusedColor   TColor
	playColor      TColor
	mouseMoveColor TColor
	mouseMoveIndex int32
	playingIndex   int32
	singerPicR     TRect
	singerPic      *TBitmap
}

func NewPlayControl(owner IComponent) *TPlayControl {
	m := new(TPlayControl)
	m.IDrawGrid = NewDrawGrid(owner)
	m.IDrawGrid.SetDefaultDrawing(false)
	m.IDrawGrid.SetDefaultRowHeight(24)
	m.IDrawGrid.SetOptions(NewSet(GoRangeSelect, GoRowSelect))
	m.IDrawGrid.SetRowCount(1)
	m.IDrawGrid.SetColCount(4)
	m.IDrawGrid.SetFixedRows(0)
	m.IDrawGrid.SetFixedCols(0)
	m.IDrawGrid.SetGridLineWidth(0)
	m.IDrawGrid.SetBorderStyle(BsNone)
	m.IDrawGrid.SetScrollBars(SsVertical)
	m.IDrawGrid.SetWidth(536)
	m.IDrawGrid.SetHeight(397)
	// 加载时取消第一行永远被选中
	m.IDrawGrid.SetSelection(&TGridRect{24, 24, 24, 24})
	m.IDrawGrid.SetColWidths(0, int32(float32(m.Width())*0.1))
	m.IDrawGrid.SetColWidths(1, int32(float32(m.Width())*0.4))
	m.IDrawGrid.SetColWidths(2, int32(float32(m.Width())*0.2))
	m.IDrawGrid.SetColWidths(3, int32(float32(m.Width())*0.2))
	m.IDrawGrid.SetColor(0x00EDEEF9)
	m.IDrawGrid.SetDoubleBuffered(true)
	m.focusedColor = 0x00C8CBEB
	m.playColor = m.focusedColor + 12
	m.mouseMoveColor = m.focusedColor - 12
	m.mouseMoveIndex = -1
	m.playingIndex = -1
	m.IDrawGrid.SetOnDblClick(m.onDblClick)
	m.IDrawGrid.SetOnMouseMove(m.onMouseMove)
	//m.IDrawGrid.SetOnMouseDown(m.onMouseDown)
	m.IDrawGrid.SetOnDrawCell(m.onDrawCell)
	//m.TDrawGrid.SetOnMouseEnter(m.onMouseEnter)
	m.IDrawGrid.SetOnMouseLeave(m.onMouseLeave)

	return m
}

func (p *TPlayControl) Add(item TPlayListItem) int32 {
	p.datas = append(p.datas, item)
	p.SetRowCount(int32(len(p.datas)))
	return int32(len(p.datas)) - 1
}

func (p *TPlayControl) onDrawCell(sender IObject, aCol, aRow int32, rect TRect, state TGridDrawState) {
	if len(p.datas) > 0 {
		canvas := p.Canvas()
		brush := AsBrush(canvas.Brush())
		font := AsFont(canvas.Font())
		if aRow < int32(len(p.datas)) {
			drawFlags := NewSet(TfVerticalCenter, TfSingleLine, TfEndEllipsis)
			item := p.datas[int(aRow)]
			if p.mouseMoveIndex == aRow && p.playingIndex != aRow && !state.In(GdFocused) && !state.In(GdSelected) {
				brush.SetColor(p.focusedColor - 12)
			} else if state.In(GdFocused) || state.In(GdSelected) {
				brush.SetColor(p.focusedColor)
			} else {
				brush.SetColor(p.Color())
			}

			if p.playingIndex == aRow {
				brush.SetColor(p.focusedColor + 12)
				p.SetRowHeights(aRow, 50)
				if p.RowHeights(aRow) < 50 {
					p.ScrollBy(0, aRow+2)
				}
			} else {
				p.SetRowHeights(aRow, 24)
			}
			canvas.FillRect(&rect)
			r := p.CellRect(aCol, aRow)
			switch aCol {
			case 0:

				if aRow == p.playingIndex {
					if !p.singerPicR.IsEmpty() {
						r.Left += 1
						r.Top += +1
						r.Bottom -= -1
						//canvas.StretchDraw(r, FSingerPic);
					}
				} else {
					r.Inflate(-10, 0)
					s := fmt.Sprintf("%d.", aRow+1)
					canvas.TextRect2(&r, s, drawFlags.Include(TfRight))
				}

			case 1:
				if aRow == p.playingIndex {
					r.Inflate(-10, 0)
					canvas.Font().SetSize(12)
					font.SetStyle(NewSet(FsBold))
					canvas.TextRect2(&r, item.Caption, drawFlags)
				} else {
					r.Inflate(-5, 0)
					canvas.TextRect2(&r, item.Caption, drawFlags)
				}
				canvas.Font().SetSize(9)
				font.SetStyle(0)
			case 2:
				r.Inflate(-5, 0)
				canvas.TextRect2(&r, item.Singer, drawFlags)
			case 3:
				r.Inflate(-5, 0)
				canvas.TextRect2(&r, p.mediaLengthToTimeStr(item.Length), drawFlags.Include(TfRight))
			}
		}

	} else {
		AsBrush(p.Canvas().Brush()).SetColor(p.Color())
		p.Canvas().FillRect(&rect)
	}
}

func (p *TPlayControl) onMouseMove(sender IObject, shift TShiftState, x, y int32) {
	if !p.Enabled() {
		return
	}
	var col, row int32
	p.MouseToCell1(x, y, &col, &row)
	p.mouseMoveIndex = row
	if p.mouseMoveIndex == -1 {
		return
	}
	p.Invalidate()
}

func (p *TPlayControl) onMouseDown(sender IObject, button TMouseButton, shift TShiftState, x, y int32) {

}

func (p *TPlayControl) onDblClick(sender IObject) {
	if !p.Enabled() {
		return
	}
	row := p.Row()
	if row == -1 {
		return
	}
	p.playingIndex = row
	p.Invalidate()
}

func (p *TPlayControl) onMouseEnter(sender IObject) {

}

func (p *TPlayControl) onMouseLeave(sender IObject) {
	if !p.Enabled() {
		return
	}
	p.mouseMoveIndex = -1
	p.Invalidate()
}

func (p *TPlayControl) mediaLengthToTimeStr(alen int32) string {
	return fmt.Sprintf("%.2d:%.2d", int(float32(alen)/1000.0)/60, int(float32(alen)/1000.0)%60)
}
