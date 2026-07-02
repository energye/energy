//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package nocgo

import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"math"
	"unsafe"
)

type Context struct {
	instance unsafe.Pointer
}

func AsContext(ptr unsafe.Pointer) IContext {
	if ptr == nil {
		return nil
	}
	m := new(Context)
	m.instance = ptr
	return m
}

func (m *Context) Instance() uintptr {
	return uintptr(m.instance)
}

// Status is a wrapper around cairo_status().
func (m *Context) Status() Status {
	r := cairo.SysCall("cairo_status", m.Instance())
	return Status(r)
}

// Close closes the context. The context must not be used afterwards.
func (m *Context) Close() {
	cairo.SysCall("cairo_destroy", m.Instance())
}

// Save is a wrapper around cairo_save().
func (m *Context) Save() {
	cairo.SysCall("cairo_save", m.Instance())
}

// Restore is a wrapper around cairo_restore().
func (m *Context) Restore() {
	cairo.SysCall("cairo_restore", m.Instance())
}

// SetSourceRGB is a wrapper around cairo_set_source_rgb().
func (m *Context) SetSourceRGB(red, green, blue float64) {
	registerCairoFloatFuncs()
	cairoSetSourceRGB(m.Instance(), red, green, blue)
}

// SetSourceRGBA is a wrapper around cairo_set_source_rgba().
func (m *Context) SetSourceRGBA(red, green, blue, alpha float64) {
	registerCairoFloatFuncs()
	cairoSetSourceRGBA(m.Instance(), red, green, blue, alpha)
}

// SetLineWidth is a wrapper around cairo_set_line_width().
func (m *Context) SetLineWidth(width float64) {
	registerCairoFloatFuncs()
	cairoSetLineWidth(m.Instance(), width)
}

// GetLineWidth is a wrapper around cairo_get_line_width().
func (m *Context) GetLineWidth() float64 {
	registerCairoFloatFuncs()
	return cairoGetLineWidth(m.Instance())
}

// Clip is a wrapper around cairo_clip().
func (m *Context) Clip() {
	cairo.SysCall("cairo_clip", m.Instance())
}

// ClipPreserve is a wrapper around cairo_clip_preserve().
func (m *Context) ClipPreserve() {
	cairo.SysCall("cairo_clip_preserve", m.Instance())
}

// ClipExtents is a wrapper around cairo_clip_extents().
func (m *Context) ClipExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 float64
	cairo.SysCall("cairo_clip_extents", m.Instance(),
		uintptr(unsafe.Pointer(&cx1)), uintptr(unsafe.Pointer(&cy1)),
		uintptr(unsafe.Pointer(&cx2)), uintptr(unsafe.Pointer(&cy2)))
	return cx1, cy1, cx2, cy2
}

// InClip is a wrapper around cairo_in_clip().
func (m *Context) InClip(x, y float64) bool {
	r := cairo.SysCall("cairo_in_clip", m.Instance(), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
	return ToGoBool(r)
}

// ResetClip is a wrapper around cairo_reset_clip().
func (m *Context) ResetClip() {
	cairo.SysCall("cairo_reset_clip", m.Instance())
}

// Rectangle is a wrapper around cairo_rectangle().
func (m *Context) Rectangle(x, y, w, h float64) {
	registerCairoFloatFuncs()
	cairoRectangle(m.Instance(), x, y, w, h)
}

// Arc is a wrapper around cairo_arc().
func (m *Context) Arc(xc, yc, radius, angle1, angle2 float64) {
	registerCairoFloatFuncs()
	cairoArc(m.Instance(), xc, yc, radius, angle1, angle2)
}

// ArcNegative is a wrapper around cairo_arc_negative().
func (m *Context) ArcNegative(xc, yc, radius, angle1, angle2 float64) {
	cairo.SysCall("cairo_arc_negative", m.Instance(),
		uintptr(math.Float64bits(xc)), uintptr(math.Float64bits(yc)), uintptr(math.Float64bits(radius)),
		uintptr(math.Float64bits(angle1)), uintptr(math.Float64bits(angle2)))
}

// LineTo is a wrapper around cairo_line_to().
func (m *Context) LineTo(x, y float64) {
	registerCairoFloatFuncs()
	cairoLineTo(m.Instance(), x, y)
}

// CurveTo is a wrapper around cairo_curve_to().
func (m *Context) CurveTo(x1, y1, x2, y2, x3, y3 float64) {
	cairo.SysCall("cairo_curve_to", m.Instance(),
		uintptr(math.Float64bits(x1)), uintptr(math.Float64bits(y1)),
		uintptr(math.Float64bits(x2)), uintptr(math.Float64bits(y2)),
		uintptr(math.Float64bits(x3)), uintptr(math.Float64bits(y3)))
}

// MoveTo is a wrapper around cairo_move_to().
func (m *Context) MoveTo(x, y float64) {
	registerCairoFloatFuncs()
	cairoMoveTo(m.Instance(), x, y)
}

// Fill is a wrapper around cairo_fill().
func (m *Context) Fill() {
	cairo.SysCall("cairo_fill", m.Instance())
}

// FillPreserve is a wrapper around cairo_fill_preserve().
func (m *Context) FillPreserve() {
	cairo.SysCall("cairo_fill_preserve", m.Instance())
}

// FillExtents is a wrapper around cairo_fill_extents().
func (m *Context) FillExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 float64
	cairo.SysCall("cairo_fill_extents", m.Instance(),
		uintptr(unsafe.Pointer(&cx1)), uintptr(unsafe.Pointer(&cy1)),
		uintptr(unsafe.Pointer(&cx2)), uintptr(unsafe.Pointer(&cy2)))
	return cx1, cy1, cx2, cy2
}

// InFill is a wrapper around cairo_in_fill().
func (m *Context) InFill(x, y float64) bool {
	r := cairo.SysCall("cairo_in_fill", m.Instance(), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
	return ToGoBool(r)
}

// ClosePath is a wrapper around cairo_close_path().
func (m *Context) ClosePath() {
	cairo.SysCall("cairo_close_path", m.Instance())
}

// NewPath is a wrapper around cairo_new_path().
func (m *Context) NewPath() {
	cairo.SysCall("cairo_new_path", m.Instance())
}

// GetCurrentPoint is a wrapper around cairo_get_current_point().
func (m *Context) GetCurrentPoint() (x, y float64) {
	cairo.SysCall("cairo_get_current_point", m.Instance(),
		uintptr(unsafe.Pointer(&x)), uintptr(unsafe.Pointer(&y)))
	return
}

// Paint is a wrapper around cairo_paint().
func (m *Context) Paint() {
	cairo.SysCall("cairo_paint", m.Instance())
}

// PaintWithAlpha is a wrapper around cairo_paint_with_alpha().
func (m *Context) PaintWithAlpha(alpha float64) {
	registerCairoFloatFuncs()
	cairoPaintWithAlpha(m.Instance(), alpha)
}

// Stroke is a wrapper around cairo_stroke().
func (m *Context) Stroke() {
	cairo.SysCall("cairo_stroke", m.Instance())
}

// StrokePreserve is a wrapper around cairo_stroke_preserve().
func (m *Context) StrokePreserve() {
	cairo.SysCall("cairo_stroke_preserve", m.Instance())
}

// StrokeExtents is a wrapper around cairo_stroke_extents().
func (m *Context) StrokeExtents() (x1, y1, x2, y2 float64) {
	var cx1, cy1, cx2, cy2 float64
	cairo.SysCall("cairo_stroke_extents", m.Instance(),
		uintptr(unsafe.Pointer(&cx1)), uintptr(unsafe.Pointer(&cy1)),
		uintptr(unsafe.Pointer(&cx2)), uintptr(unsafe.Pointer(&cy2)))
	return cx1, cy1, cx2, cy2
}

// InStroke is a wrapper around cairo_in_stroke().
func (m *Context) InStroke(x, y float64) bool {
	r := cairo.SysCall("cairo_in_stroke", m.Instance(), uintptr(math.Float64bits(x)), uintptr(math.Float64bits(y)))
	return ToGoBool(r)
}

// CopyPage is a wrapper around cairo_copy_page().
func (m *Context) CopyPage() {
	cairo.SysCall("cairo_copy_page", m.Instance())
}

// ShowPage is a wrapper around cairo_show_page().
func (m *Context) ShowPage() {
	cairo.SysCall("cairo_show_page", m.Instance())
}

// SetFillRule is a wrapper around cairo_set_fill_rule().
func (m *Context) SetFillRule(fillRule FillRule) {
	cairo.SysCall("cairo_set_fill_rule", m.Instance(), uintptr(fillRule))
}

// GetFillRule is a wrapper around cairo_get_fill_rule().
func (m *Context) GetFillRule() FillRule {
	r := cairo.SysCall("cairo_get_fill_rule", m.Instance())
	return FillRule(r)
}

// SetLineCap is a wrapper around cairo_set_line_cap().
func (m *Context) SetLineCap(lineCap LineCap) {
	cairo.SysCall("cairo_set_line_cap", m.Instance(), uintptr(lineCap))
}

// GetLineCap is a wrapper around cairo_get_line_cap().
func (m *Context) GetLineCap() LineCap {
	r := cairo.SysCall("cairo_get_line_cap", m.Instance())
	return LineCap(r)
}

// SetLineJoin is a wrapper around cairo_set_line_join().
func (m *Context) SetLineJoin(lineJoin LineJoin) {
	cairo.SysCall("cairo_set_line_join", m.Instance(), uintptr(lineJoin))
}

// GetLineJoin is a wrapper around cairo_get_line_join().
func (m *Context) GetLineJoin() LineJoin {
	r := cairo.SysCall("cairo_get_line_join", m.Instance())
	return LineJoin(r)
}

// SetOperator is a wrapper around cairo_set_operator().
func (m *Context) SetOperator(op Operator) {
	cairo.SysCall("cairo_set_operator", m.Instance(), uintptr(op))
}

// GetOperator is a wrapper around cairo_get_operator().
func (m *Context) GetOperator() Operator {
	r := cairo.SysCall("cairo_get_operator", m.Instance())
	return Operator(r)
}

// SetDash is a wrapper around cairo_set_dash().
func (m *Context) SetDash(dashes []float64, offset float64) {
	count := len(dashes)
	if count == 0 {
		cairo.SysCall("cairo_set_dash", m.Instance(), 0, 0, uintptr(math.Float64bits(offset)))
		return
	}
	cairo.SysCall("cairo_set_dash", m.Instance(),
		uintptr(unsafe.Pointer(&dashes[0])), uintptr(count), uintptr(math.Float64bits(offset)))
}

// GetDashCount is a wrapper around cairo_get_dash_count().
func (m *Context) GetDashCount() int {
	r := cairo.SysCall("cairo_get_dash_count", m.Instance())
	return int(r)
}

// GetDash is a wrapper around cairo_get_dash().
func (m *Context) GetDash() (dashes []float64, offset float64) {
	dashCount := m.GetDashCount()
	if dashCount == 0 {
		return nil, 0
	}
	cdashes := make([]float64, dashCount)
	var coffset float64
	cairo.SysCall("cairo_get_dash", m.Instance(),
		uintptr(unsafe.Pointer(&cdashes[0])),
		uintptr(unsafe.Pointer(&coffset)))
	return cdashes, coffset
}

// SetTolerance is a wrapper around cairo_set_tolerance().
func (m *Context) SetTolerance(tolerance float64) {
	cairo.SysCall("cairo_set_tolerance", m.Instance(), uintptr(math.Float64bits(tolerance)))
}

// GetTolerance is a wrapper around cairo_get_tolerance().
func (m *Context) GetTolerance() float64 {
	r := cairo.SysCall("cairo_get_tolerance", m.Instance())
	return math.Float64frombits(uint64(r))
}

// SetMiterLimit is a wrapper around cairo_set_miter_limit().
func (m *Context) SetMiterLimit(limit float64) {
	cairo.SysCall("cairo_set_miter_limit", m.Instance(), uintptr(math.Float64bits(limit)))
}

// GetMiterLimit is a wrapper around cairo_get_miter_limit().
func (m *Context) GetMiterLimit() float64 {
	r := cairo.SysCall("cairo_get_miter_limit", m.Instance())
	return math.Float64frombits(uint64(r))
}

// PushGroup is a wrapper around cairo_push_group().
func (m *Context) PushGroup() {
	cairo.SysCall("cairo_push_group", m.Instance())
}

// PopGroupToSource is a wrapper around cairo_pop_group_to_source().
func (m *Context) PopGroupToSource() {
	cairo.SysCall("cairo_pop_group_to_source", m.Instance())
}
