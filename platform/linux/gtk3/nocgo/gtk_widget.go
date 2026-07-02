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
	"errors"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type Widget struct {
	Object
}

func AsWidget(ptr unsafe.Pointer) IWidget {
	if ptr == nil {
		return nil
	}
	m := new(Widget)
	m.instance = ptr
	return m
}

// GetScreen is a wrapper around gtk_widget_get_screen().
func (m *Widget) GetScreen() IScreen {
	r := gtk3.SysCall("gtk_widget_get_screen", m.Instance())
	if r == 0 {
		return nil
	}
	s := &Screen{}
	s.instance = unsafe.Pointer(r)
	return s
}

// SetVisual is a wrapper around gtk_widget_set_visual().
func (m *Widget) SetVisual(visual IVisual) {
	gtk3.SysCall("gtk_widget_set_visual", m.Instance(), visual.Instance())
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (m *Widget) SetAppPaintable(paintable bool) {
	gtk3.SysCall("gtk_widget_set_app_paintable", m.Instance(), ToCBool(paintable))
}

// GetName is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (m *Widget) GetName() string {
	r := gtk3.SysCall("gtk_widget_get_name", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// GetAllocation is a wrapper around gtk_widget_get_allocation().
func (m *Widget) GetAllocation() IRectangle {
	var rect Rectangle
	gtk3.SysCall("gtk_widget_get_allocation", m.Instance(), uintptr(unsafe.Pointer(&rect)))
	return &rect
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (m *Widget) SetSizeRequest(width, height int) {
	gtk3.SysCall("gtk_widget_set_size_request", m.Instance(), uintptr(width), uintptr(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (m *Widget) GetSizeRequest() (width, height int) {
	gtk3.SysCall("gtk_widget_get_size_request", m.Instance(), uintptr(unsafe.Pointer(&width)), uintptr(unsafe.Pointer(&height)))
	return
}

// GetStyleContext is a wrapper around gtk_widget_get_style_context().
func (m *Widget) GetStyleContext() IStyleContext {
	r := gtk3.SysCall("gtk_widget_get_style_context", m.Instance())
	return AsStyleContext(unsafe.Pointer(r))
}

// GrabFocus is a wrapper around gtk_widget_grab_focus().
func (m *Widget) GrabFocus() {
	gtk3.SysCall("gtk_widget_grab_focus", m.Instance())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (m *Widget) ShowAll() {
	gtk3.SysCall("gtk_widget_show_all", m.Instance())
}

// Realize is a wrapper around gtk_widget_realize().
func (m *Widget) Realize() {
	gtk3.SysCall("gtk_widget_realize", m.Instance())
}

func (m *Widget) DragGetData(context IDragContext, target IAtom, time uint) {
	gtk3.SysCall("gtk_drag_get_data", m.Instance(), context.Instance(), uintptr(target.Atom()), uintptr(time))
}

func (m *Widget) IsContainer() bool {
	containerGType := GTypeFromName(CStr("GtkContainer"))
	widgetGType := GTypeFormInstance(m.Instance())
	return GTypeIsA(widgetGType, containerGType)
}

// ToWidget is a helper getter, e.g.: it returns *gtk.Label as a *gtk.Widget.
func (m *Widget) ToWidget() *Widget {
	return m
}

// GetHAlign is a wrapper around gtk_widget_get_halign().
func (m *Widget) GetHAlign() Align {
	r := gtk3.SysCall("gtk_widget_get_halign", m.Instance())
	return Align(r)
}

// SetHAlign is a wrapper around gtk_widget_set_halign().
func (m *Widget) SetHAlign(align Align) {
	gtk3.SysCall("gtk_widget_set_halign", m.Instance(), uintptr(align))
}

// GetVAlign is a wrapper around gtk_widget_get_valign().
func (m *Widget) GetVAlign() Align {
	r := gtk3.SysCall("gtk_widget_get_valign", m.Instance())
	return Align(r)
}

// SetVAlign is a wrapper around gtk_widget_set_valign().
func (m *Widget) SetVAlign(align Align) {
	gtk3.SysCall("gtk_widget_set_valign", m.Instance(), uintptr(align))
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (m *Widget) GetMarginTop() int {
	return int(gtk3.SysCall("gtk_widget_get_margin_top", m.Instance()))
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (m *Widget) SetMarginTop(margin int) {
	gtk3.SysCall("gtk_widget_set_margin_top", m.Instance(), uintptr(margin))
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (m *Widget) GetMarginBottom() int {
	return int(gtk3.SysCall("gtk_widget_get_margin_bottom", m.Instance()))
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (m *Widget) SetMarginBottom(margin int) {
	gtk3.SysCall("gtk_widget_set_margin_bottom", m.Instance(), uintptr(margin))
}

// GetHExpand is a wrapper around gtk_widget_get_hexpand().
func (m *Widget) GetHExpand() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_hexpand", m.Instance()))
}

// SetHExpand is a wrapper around gtk_widget_set_hexpand().
func (m *Widget) SetHExpand(expand bool) {
	gtk3.SysCall("gtk_widget_set_hexpand", m.Instance(), ToCBool(expand))
}

// GetVExpand is a wrapper around gtk_widget_get_vexpand().
func (m *Widget) GetVExpand() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_vexpand", m.Instance()))
}

// SetVExpand is a wrapper around gtk_widget_set_vexpand().
func (m *Widget) SetVExpand(expand bool) {
	gtk3.SysCall("gtk_widget_set_vexpand", m.Instance(), ToCBool(expand))
}

// GetRealized is a wrapper around gtk_widget_get_realized().
func (m *Widget) GetRealized() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_realized", m.Instance()))
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (m *Widget) SetRealized(realized bool) {
	gtk3.SysCall("gtk_widget_set_realized", m.Instance(), ToCBool(realized))
}

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (m *Widget) GetHasWindow() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_has_window", m.Instance()))
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (m *Widget) SetHasWindow(hasWindow bool) {
	gtk3.SysCall("gtk_widget_set_has_window", m.Instance(), ToCBool(hasWindow))
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (m *Widget) ShowNow() {
	gtk3.SysCall("gtk_widget_show_now", m.Instance())
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (m *Widget) SetNoShowAll(noShowAll bool) {
	gtk3.SysCall("gtk_widget_set_no_show_all", m.Instance(), ToCBool(noShowAll))
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (m *Widget) GetNoShowAll() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_no_show_all", m.Instance()))
}

// Map is a wrapper around gtk_widget_map().
func (m *Widget) Map() {
	gtk3.SysCall("gtk_widget_map", m.Instance())
}

// Unmap is a wrapper around gtk_widget_unmap().
func (m *Widget) Unmap() {
	gtk3.SysCall("gtk_widget_unmap", m.Instance())
}

// Unrealize is a wrapper around gtk_widget_unrealize().
func (m *Widget) Unrealize() {
	gtk3.SysCall("gtk_widget_unrealize", m.Instance())
}

// Event is a wrapper around gtk_widget_event().
func (m *Widget) Event(event *Event) bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_event", m.Instance(), event.Instance()))
}

// Activate is a wrapper around gtk_widget_activate().
func (m *Widget) Activate() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_activate", m.Instance()))
}

// IsFocus is a wrapper around gtk_widget_is_focus().
func (m *Widget) IsFocus() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_is_focus", m.Instance()))
}

// GrabDefault is a wrapper around gtk_widget_grab_default().
func (m *Widget) GrabDefault() {
	gtk3.SysCall("gtk_widget_grab_default", m.Instance())
}

// SetName is a wrapper around gtk_widget_set_name().
func (m *Widget) SetName(name string) {
	gtk3.SysCall("gtk_widget_set_name", m.Instance(), CStr(name))
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (m *Widget) GetSensitive() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_sensitive", m.Instance()))
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (m *Widget) IsSensitive() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_is_sensitive", m.Instance()))
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (m *Widget) SetSensitive(sensitive bool) {
	gtk3.SysCall("gtk_widget_set_sensitive", m.Instance(), ToCBool(sensitive))
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (m *Widget) GetVisible() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_visible", m.Instance()))
}

// SetVisible is a wrapper around gtk_widget_set_visible().
func (m *Widget) SetVisible(visible bool) {
	gtk3.SysCall("gtk_widget_set_visible", m.Instance(), ToCBool(visible))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (m *Widget) SetParent(parent IWidget) {
	gtk3.SysCall("gtk_widget_set_parent", m.Instance(), parent.Instance())
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (m *Widget) GetParent() *Widget {
	r := gtk3.SysCall("gtk_widget_get_parent", m.Instance())
	if r == 0 {
		return nil
	}
	widget := new(Widget)
	widget.instance = unsafe.Pointer(r)
	return widget
}

// GetAllocatedWidth is a wrapper around gtk_widget_get_allocated_width().
func (m *Widget) GetAllocatedWidth() int {
	return int(gtk3.SysCall("gtk_widget_get_allocated_width", m.Instance()))
}

// GetAllocatedHeight is a wrapper around gtk_widget_get_allocated_height().
func (m *Widget) GetAllocatedHeight() int {
	return int(gtk3.SysCall("gtk_widget_get_allocated_height", m.Instance()))
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (m *Widget) SetEvents(events EventMask) {
	gtk3.SysCall("gtk_widget_set_events", m.Instance(), uintptr(events))
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (m *Widget) GetEvents() EventMask {
	return EventMask(gtk3.SysCall("gtk_widget_get_events", m.Instance()))
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (m *Widget) AddEvents(events EventMask) {
	gtk3.SysCall("gtk_widget_add_events", m.Instance(), uintptr(events))
}

// FreezeChildNotify is a wrapper around gtk_widget_freeze_child_notify().
func (m *Widget) FreezeChildNotify() {
	gtk3.SysCall("gtk_widget_freeze_child_notify", m.Instance())
}

// ThawChildNotify is a wrapper around gtk_widget_thaw_child_notify().
func (m *Widget) ThawChildNotify() {
	gtk3.SysCall("gtk_widget_thaw_child_notify", m.Instance())
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (m *Widget) HasDefault() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_has_default", m.Instance()))
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (m *Widget) HasFocus() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_has_focus", m.Instance()))
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (m *Widget) HasVisibleFocus() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_has_visible_focus", m.Instance()))
}

// SetOpacity is a wrapper around gtk_widget_set_opacity().
func (m *Widget) SetOpacity(opacity float64) {
	registerGtkFloatFuncs()
	gtkWidgetSetOpacity(m.Instance(), opacity)
}

// GetOpacity is a wrapper around gtk_widget_get_opacity().
func (m *Widget) GetOpacity() float64 {
	registerGtkFloatFuncs()
	return gtkWidgetGetOpacity(m.Instance())
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (m *Widget) HasGrab() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_has_grab", m.Instance()))
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (m *Widget) IsDrawable() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_is_drawable", m.Instance()))
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (m *Widget) IsToplevel() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_is_toplevel", m.Instance()))
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (m *Widget) GetToplevel() *Widget {
	r := gtk3.SysCall("gtk_widget_get_toplevel", m.Instance())
	if r == 0 {
		return nil
	}
	widget := new(Widget)
	widget.instance = unsafe.Pointer(r)
	return widget
}

// GetTooltipMarkup is a wrapper around gtk_widget_get_tooltip_markup().
func (m *Widget) GetTooltipMarkup() string {
	r := gtk3.SysCall("gtk_widget_get_tooltip_markup", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// SetTooltipMarkup is a wrapper around gtk_widget_set_tooltip_markup().
func (m *Widget) SetTooltipMarkup(text string) {
	gtk3.SysCall("gtk_widget_set_tooltip_markup", m.Instance(), CStr(text))
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
func (m *Widget) GetTooltipText() string {
	r := gtk3.SysCall("gtk_widget_get_tooltip_text", m.Instance())
	if r == 0 {
		return ""
	}
	return GoStr(r)
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (m *Widget) SetTooltipText(text string) {
	gtk3.SysCall("gtk_widget_set_tooltip_text", m.Instance(), CStr(text))
}

// TranslateCoordinates is a wrapper around gtk_widget_translate_coordinates().
func (m *Widget) TranslateCoordinates(dest IWidget, srcX, srcY int) (destX, destY int, e error) {
	var destPtr uintptr
	if dest != nil {
		destPtr = dest.Instance()
	}
	var cdestX, cdestY int
	r := gtk3.SysCall("gtk_widget_translate_coordinates", m.Instance(), destPtr, uintptr(srcX), uintptr(srcY), uintptr(unsafe.Pointer(&cdestX)), uintptr(unsafe.Pointer(&cdestY)))
	if !ToGoBool(r) {
		return 0, 0, errors.New("translate coordinates failed")
	}
	return cdestX, cdestY, nil
}

// GetAppPaintable is a wrapper around gtk_widget_get_app_paintable().
func (m *Widget) GetAppPaintable() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_app_paintable", m.Instance()))
}

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (m *Widget) QueueDraw() {
	gtk3.SysCall("gtk_widget_queue_draw", m.Instance())
}

// SetAllocation is a wrapper around gtk_widget_set_allocation().
func (m *Widget) SetAllocation(allocation *Rectangle) {
	gtk3.SysCall("gtk_widget_set_allocation", m.Instance(), uintptr(unsafe.Pointer(allocation)))
}

// SizeAllocate is a wrapper around gtk_widget_size_allocate().
func (m *Widget) SizeAllocate(allocation *Rectangle) {
	gtk3.SysCall("gtk_widget_size_allocate", m.Instance(), uintptr(unsafe.Pointer(allocation)))
}

// SetStateFlags is a wrapper around gtk_widget_set_state_flags().
func (m *Widget) SetStateFlags(stateFlags StateFlags, clear bool) {
	gtk3.SysCall("gtk_widget_set_state_flags", m.Instance(), uintptr(stateFlags), ToCBool(clear))
}

// UnsetStateFlags is a wrapper around gtk_widget_unset_state_flags().
func (m *Widget) UnsetStateFlags(stateFlags StateFlags) {
	gtk3.SysCall("gtk_widget_unset_state_flags", m.Instance(), uintptr(stateFlags))
}

// GetStateFlags is a wrapper around gtk_widget_get_state_flags().
func (m *Widget) GetStateFlags() StateFlags {
	return StateFlags(gtk3.SysCall("gtk_widget_get_state_flags", m.Instance()))
}

// GetDisplay is a wrapper around gtk_widget_get_display().
func (m *Widget) GetDisplay() *Display {
	r := gtk3.SysCall("gtk_widget_get_display", m.Instance())
	if r == 0 {
		return nil
	}
	d := new(Display)
	d.instance = unsafe.Pointer(r)
	return d
}

// SetMarginStart is a wrapper around gtk_widget_set_margin_start().
func (m *Widget) SetMarginStart(margin int) {
	gtk3.SysCall("gtk_widget_set_margin_start", m.Instance(), uintptr(margin))
}

// GetMarginStart is a wrapper around gtk_widget_get_margin_start().
func (m *Widget) GetMarginStart() int {
	return int(gtk3.SysCall("gtk_widget_get_margin_start", m.Instance()))
}

// SetMarginEnd is a wrapper around gtk_widget_set_margin_end().
func (m *Widget) SetMarginEnd(margin int) {
	gtk3.SysCall("gtk_widget_set_margin_end", m.Instance(), uintptr(margin))
}

// GetMarginEnd is a wrapper around gtk_widget_get_margin_end().
func (m *Widget) GetMarginEnd() int {
	return int(gtk3.SysCall("gtk_widget_get_margin_end", m.Instance()))
}

// GetFocusOnClick is a wrapper around gtk_widget_get_focus_on_click().
func (m *Widget) GetFocusOnClick() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_focus_on_click", m.Instance()))
}

// SetFocusOnClick is a wrapper around gtk_widget_set_focus_on_click().
func (m *Widget) SetFocusOnClick(focusOnClick bool) {
	gtk3.SysCall("gtk_widget_set_focus_on_click", m.Instance(), ToCBool(focusOnClick))
}

// ResetStyle is a wrapper around gtk_widget_reset_style().
func (m *Widget) ResetStyle() {
	gtk3.SysCall("gtk_widget_reset_style", m.Instance())
}

// InDestruction is a wrapper around gtk_widget_in_destruction().
func (m *Widget) InDestruction() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_in_destruction", m.Instance()))
}

// Unparent is a wrapper around gtk_widget_unparent().
func (m *Widget) Unparent() {
	gtk3.SysCall("gtk_widget_unparent", m.Instance())
}

// Show is a wrapper around gtk_widget_show().
func (m *Widget) Show() {
	gtk3.SysCall("gtk_widget_show", m.Instance())
}

// Hide is a wrapper around gtk_widget_hide().
func (m *Widget) Hide() {
	gtk3.SysCall("gtk_widget_hide", m.Instance())
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (m *Widget) GetCanFocus() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_can_focus", m.Instance()))
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (m *Widget) SetCanFocus(canFocus bool) {
	gtk3.SysCall("gtk_widget_set_can_focus", m.Instance(), ToCBool(canFocus))
}

// GetCanDefault is a wrapper around gtk_widget_get_can_default().
func (m *Widget) GetCanDefault() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_can_default", m.Instance()))
}

// SetCanDefault is a wrapper around gtk_widget_set_can_default().
func (m *Widget) SetCanDefault(canDefault bool) {
	gtk3.SysCall("gtk_widget_set_can_default", m.Instance(), ToCBool(canDefault))
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (m *Widget) SetMapped(mapped bool) {
	gtk3.SysCall("gtk_widget_set_mapped", m.Instance(), ToCBool(mapped))
}

// GetMapped is a wrapper around gtk_widget_get_mapped().
func (m *Widget) GetMapped() bool {
	return ToGoBool(gtk3.SysCall("gtk_widget_get_mapped", m.Instance()))
}

// GetPreferredHeight is a wrapper around gtk_widget_get_preferred_height().
func (m *Widget) GetPreferredHeight() (int, int) {
	var minimum, natural int
	gtk3.SysCall("gtk_widget_get_preferred_height", m.Instance(), uintptr(unsafe.Pointer(&minimum)), uintptr(unsafe.Pointer(&natural)))
	return minimum, natural
}

// GetPreferredWidth is a wrapper around gtk_widget_get_preferred_width().
func (m *Widget) GetPreferredWidth() (int, int) {
	var minimum, natural int
	gtk3.SysCall("gtk_widget_get_preferred_width", m.Instance(), uintptr(unsafe.Pointer(&minimum)), uintptr(unsafe.Pointer(&natural)))
	return minimum, natural
}

// GetPreferredHeightForWidth is a wrapper around gtk_widget_get_preferred_height_for_width().
func (m *Widget) GetPreferredHeightForWidth(width int) (int, int) {
	var minimum, natural int
	gtk3.SysCall("gtk_widget_get_preferred_height_for_width", m.Instance(), uintptr(width), uintptr(unsafe.Pointer(&minimum)), uintptr(unsafe.Pointer(&natural)))
	return minimum, natural
}

// GetPreferredWidthForHeight is a wrapper around gtk_widget_get_preferred_width_for_height().
func (m *Widget) GetPreferredWidthForHeight(height int) (int, int) {
	var minimum, natural int
	gtk3.SysCall("gtk_widget_get_preferred_width_for_height", m.Instance(), uintptr(height), uintptr(unsafe.Pointer(&minimum)), uintptr(unsafe.Pointer(&natural)))
	return minimum, natural
}

// GetRequestMode is a wrapper around gtk_widget_get_request_mode().
func (m *Widget) GetRequestMode() SizeRequestMode {
	return SizeRequestMode(gtk3.SysCall("gtk_widget_get_request_mode", m.Instance()))
}

// GetPreferredSize is a wrapper around gtk_widget_get_preferred_size().
func (m *Widget) GetPreferredSize() (*Requisition, *Requisition) {
	minimumSize := new(Requisition)
	naturalSize := new(Requisition)
	gtk3.SysCall("gtk_widget_get_preferred_size", m.Instance(), uintptr(unsafe.Pointer(minimumSize)), uintptr(unsafe.Pointer(naturalSize)))
	return minimumSize, naturalSize
}

// Destroy is a wrapper around gtk_widget_destroy().
func (m *Widget) Destroy() {
	gtk3.SysCall("gtk_widget_destroy", m.Instance())
}

// QueueDrawArea is a wrapper around gtk_widget_queue_draw_area().
func (m *Widget) QueueDrawArea(x, y, width, height int) {
	gtk3.SysCall("gtk_widget_queue_draw_area", m.Instance(),
		uintptr(x), uintptr(y), uintptr(width), uintptr(height))
}

// TargetEntry is a representation of GTK's GtkTargetEntry.
// The struct layout matches the C GtkTargetEntry: {gchar *target, guint flags, guint info}.
type TargetEntry struct {
	target uintptr
	flags  uint32
	info   uint32
}

// NewTargetEntry creates a new TargetEntry.
func NewTargetEntry(target string, flags TargetFlags, info uint) *TargetEntry {
	return &TargetEntry{
		target: CStr(target),
		flags:  uint32(flags),
		info:   uint32(info),
	}
}

// DragDestSet is a wrapper around gtk_drag_dest_set().
func (m *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions DragAction) {
	if len(targets) == 0 {
		gtk3.SysCall("gtk_drag_dest_set", m.Instance(), uintptr(flags), 0, 0, uintptr(actions))
	} else {
		gtk3.SysCall("gtk_drag_dest_set", m.Instance(), uintptr(flags),
			uintptr(unsafe.Pointer(&targets[0])), uintptr(len(targets)), uintptr(actions))
	}
}

// DragSourceSet is a wrapper around gtk_drag_source_set().
func (m *Widget) DragSourceSet(startButtonMask uint, targets []TargetEntry, actions DragAction) {
	if len(targets) == 0 {
		gtk3.SysCall("gtk_drag_source_set", m.Instance(), uintptr(startButtonMask), 0, 0, uintptr(actions))
	} else {
		gtk3.SysCall("gtk_drag_source_set", m.Instance(), uintptr(startButtonMask),
			uintptr(unsafe.Pointer(&targets[0])), uintptr(len(targets)), uintptr(actions))
	}
}
