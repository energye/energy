package cgo

// #cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"errors"
	. "github.com/energye/energy/v3/pkgs/linux/types"
	"unsafe"
)

// IWidget is an interface type implemented by all structs
// embedding a Widget.  It is meant to be used as an argument type
// for wrapper functions that wrap around a C GTK function taking a
// GtkWidget.
type _IWidget interface {
	toWidget() *C.GtkWidget
}

// Widget is a representation of GTK's GtkWidget.
type Widget struct {
	InitiallyUnowned
}

func GtkWidget(widget IWidget) *C.GtkWidget {
	if widget == nil {
		return (*C.GtkWidget)(nil)
	}
	return widget.(_IWidget).toWidget()
}

func wrapWidget(obj *Object) *Widget {
	if obj == nil {
		return nil
	}

	return &Widget{InitiallyUnowned{obj}}
}

func AsWidget(p unsafe.Pointer) IWidget {
	return &Widget{InitiallyUnowned{ToGoObject(p)}}
}

// native returns a pointer to the underlying GtkWidget.
func (v *Widget) native() *C.GtkWidget {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkWidget(p)
}

func (v *Widget) toWidget() *C.GtkWidget {
	if v == nil {
		return nil
	}
	return v.native()
}

// ToWidget is a helper getter, e.g.: it returns *gtk.Label as a *gtk.Widget.
// In other cases, where you have a gtk.IWidget, use the type assertion.
func (v *Widget) ToWidget() *Widget {
	return v
}

// GetHAlign is a wrapper around gtk_widget_get_halign().
func (v *Widget) GetHAlign() Align {
	c := C.gtk_widget_get_halign(v.native())
	return Align(c)
}

// SetHAlign is a wrapper around gtk_widget_set_halign().
func (v *Widget) SetHAlign(align Align) {
	C.gtk_widget_set_halign(v.native(), C.GtkAlign(align))
}

// GetVAlign is a wrapper around gtk_widget_get_valign().
func (v *Widget) GetVAlign() Align {
	c := C.gtk_widget_get_valign(v.native())
	return Align(c)
}

// SetVAlign is a wrapper around gtk_widget_set_valign().
func (v *Widget) SetVAlign(align Align) {
	C.gtk_widget_set_valign(v.native(), C.GtkAlign(align))
}

// GetMarginTop is a wrapper around gtk_widget_get_margin_top().
func (v *Widget) GetMarginTop() int {
	c := C.gtk_widget_get_margin_top(v.native())
	return int(c)
}

// SetMarginTop is a wrapper around gtk_widget_set_margin_top().
func (v *Widget) SetMarginTop(margin int) {
	C.gtk_widget_set_margin_top(v.native(), C.gint(margin))
}

// GetMarginBottom is a wrapper around gtk_widget_get_margin_bottom().
func (v *Widget) GetMarginBottom() int {
	c := C.gtk_widget_get_margin_bottom(v.native())
	return int(c)
}

// SetMarginBottom is a wrapper around gtk_widget_set_margin_bottom().
func (v *Widget) SetMarginBottom(margin int) {
	C.gtk_widget_set_margin_bottom(v.native(), C.gint(margin))
}

// GetHExpand is a wrapper around gtk_widget_get_hexpand().
func (v *Widget) GetHExpand() bool {
	c := C.gtk_widget_get_hexpand(v.native())
	return GoBool(c)
}

// SetHExpand is a wrapper around gtk_widget_set_hexpand().
func (v *Widget) SetHExpand(expand bool) {
	C.gtk_widget_set_hexpand(v.native(), CBool(expand))
}

// GetVExpand is a wrapper around gtk_widget_get_vexpand().
func (v *Widget) GetVExpand() bool {
	c := C.gtk_widget_get_vexpand(v.native())
	return GoBool(c)
}

// SetVExpand is a wrapper around gtk_widget_set_vexpand().
func (v *Widget) SetVExpand(expand bool) {
	C.gtk_widget_set_vexpand(v.native(), CBool(expand))
}

// GetRealized is a wrapper around gtk_widget_get_realized().
func (v *Widget) GetRealized() bool {
	c := C.gtk_widget_get_realized(v.native())
	return GoBool(c)
}

// SetRealized is a wrapper around gtk_widget_set_realized().
func (v *Widget) SetRealized(realized bool) {
	C.gtk_widget_set_realized(v.native(), CBool(realized))
}

// GetHasWindow is a wrapper around gtk_widget_get_has_window().
func (v *Widget) GetHasWindow() bool {
	c := C.gtk_widget_get_has_window(v.native())
	return GoBool(c)
}

// SetHasWindow is a wrapper around gtk_widget_set_has_window().
func (v *Widget) SetHasWindow(hasWindow bool) {
	C.gtk_widget_set_has_window(v.native(), CBool(hasWindow))
}

// ShowNow is a wrapper around gtk_widget_show_now().
func (v *Widget) ShowNow() {
	C.gtk_widget_show_now(v.native())
}

// ShowAll is a wrapper around gtk_widget_show_all().
func (v *Widget) ShowAll() {
	C.gtk_widget_show_all(v.native())
}

// SetNoShowAll is a wrapper around gtk_widget_set_no_show_all().
func (v *Widget) SetNoShowAll(noShowAll bool) {
	C.gtk_widget_set_no_show_all(v.native(), CBool(noShowAll))
}

// GetNoShowAll is a wrapper around gtk_widget_get_no_show_all().
func (v *Widget) GetNoShowAll() bool {
	c := C.gtk_widget_get_no_show_all(v.native())
	return GoBool(c)
}

// Map is a wrapper around gtk_widget_map().
func (v *Widget) Map() {
	C.gtk_widget_map(v.native())
}

// Unmap is a wrapper around gtk_widget_unmap().
func (v *Widget) Unmap() {
	C.gtk_widget_unmap(v.native())
}

// Realize is a wrapper around gtk_widget_realize().
func (v *Widget) Realize() {
	C.gtk_widget_realize(v.native())
}

// Unrealize is a wrapper around gtk_widget_unrealize().
func (v *Widget) Unrealize() {
	C.gtk_widget_unrealize(v.native())
}

// Event is a wrapper around gtk_widget_event().
func (v *Widget) Event(event *Event) bool {
	c := C.gtk_widget_event(v.native(),
		(*C.GdkEvent)(unsafe.Pointer(event.Native())))
	return GoBool(c)
}

// Activate is a wrapper around gtk_widget_activate().
func (v *Widget) Activate() bool {
	return GoBool(C.gtk_widget_activate(v.native()))
}

// IsFocus is a wrapper around gtk_widget_is_focus().
func (v *Widget) IsFocus() bool {
	return GoBool(C.gtk_widget_is_focus(v.native()))
}

// GrabFocus is a wrapper around gtk_widget_grab_focus().
func (v *Widget) GrabFocus() {
	C.gtk_widget_grab_focus(v.native())
}

// GrabDefault is a wrapper around gtk_widget_grab_default().
func (v *Widget) GrabDefault() {
	C.gtk_widget_grab_default(v.native())
}

// SetName is a wrapper around gtk_widget_set_name().
func (v *Widget) SetName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_name(v.native(), (*C.gchar)(cstr))
}

// GetName is a wrapper around gtk_widget_get_name().  A non-nil
// error is returned in the case that gtk_widget_get_name returns NULL to
// differentiate between NULL and an empty string.
func (v *Widget) GetName() string {
	c := C.gtk_widget_get_name(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// GetSensitive is a wrapper around gtk_widget_get_sensitive().
func (v *Widget) GetSensitive() bool {
	c := C.gtk_widget_get_sensitive(v.native())
	return GoBool(c)
}

// IsSensitive is a wrapper around gtk_widget_is_sensitive().
func (v *Widget) IsSensitive() bool {
	c := C.gtk_widget_is_sensitive(v.native())
	return GoBool(c)
}

// SetSensitive is a wrapper around gtk_widget_set_sensitive().
func (v *Widget) SetSensitive(sensitive bool) {
	C.gtk_widget_set_sensitive(v.native(), CBool(sensitive))
}

// GetVisible is a wrapper around gtk_widget_get_visible().
func (v *Widget) GetVisible() bool {
	c := C.gtk_widget_get_visible(v.native())
	return GoBool(c)
}

// SetVisible is a wrapper around gtk_widget_set_visible().
func (v *Widget) SetVisible(visible bool) {
	C.gtk_widget_set_visible(v.native(), CBool(visible))
}

// SetParent is a wrapper around gtk_widget_set_parent().
func (v *Widget) SetParent(parent IWidget) {
	C.gtk_widget_set_parent(v.native(), GtkWidget(parent))
}

// GetParent is a wrapper around gtk_widget_get_parent().
func (v *Widget) GetParent() *Widget {
	c := C.gtk_widget_get_parent(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(ToGoObject(unsafe.Pointer(c)))
}

// SetSizeRequest is a wrapper around gtk_widget_set_size_request().
func (v *Widget) SetSizeRequest(width, height int) {
	C.gtk_widget_set_size_request(v.native(), C.gint(width), C.gint(height))
}

// GetSizeRequest is a wrapper around gtk_widget_get_size_request().
func (v *Widget) GetSizeRequest() (width, height int) {
	var w, h C.gint
	C.gtk_widget_get_size_request(v.native(), &w, &h)
	return int(w), int(h)
}

// GetAllocatedWidth is a wrapper around gtk_widget_get_allocated_width().
func (v *Widget) GetAllocatedWidth() int {
	return int(C.gtk_widget_get_allocated_width(v.native()))
}

// GetAllocatedHeight is a wrapper around gtk_widget_get_allocated_height().
func (v *Widget) GetAllocatedHeight() int {
	return int(C.gtk_widget_get_allocated_height(v.native()))
}

// SetEvents is a wrapper around gtk_widget_set_events().
func (v *Widget) SetEvents(events int) {
	C.gtk_widget_set_events(v.native(), C.gint(events))
}

// GetEvents is a wrapper around gtk_widget_get_events().
func (v *Widget) GetEvents() int {
	return int(C.gtk_widget_get_events(v.native()))
}

// AddEvents is a wrapper around gtk_widget_add_events().
func (v *Widget) AddEvents(events EventMask) {
	C.gtk_widget_add_events(v.native(), C.gint(events))
}

// FreezeChildNotify is a wrapper around gtk_widget_freeze_child_notify().
func (v *Widget) FreezeChildNotify() {
	C.gtk_widget_freeze_child_notify(v.native())
}

// ThawChildNotify is a wrapper around gtk_widget_thaw_child_notify().
func (v *Widget) ThawChildNotify() {
	C.gtk_widget_thaw_child_notify(v.native())
}

// HasDefault is a wrapper around gtk_widget_has_default().
func (v *Widget) HasDefault() bool {
	c := C.gtk_widget_has_default(v.native())
	return GoBool(c)
}

// HasFocus is a wrapper around gtk_widget_has_focus().
func (v *Widget) HasFocus() bool {
	c := C.gtk_widget_has_focus(v.native())
	return GoBool(c)
}

// HasVisibleFocus is a wrapper around gtk_widget_has_visible_focus().
func (v *Widget) HasVisibleFocus() bool {
	c := C.gtk_widget_has_visible_focus(v.native())
	return GoBool(c)
}

// SetOpacity is a wrapper around gtk_widget_set_opacity()
func (v *Widget) SetOpacity(opacity float64) {
	C.gtk_widget_set_opacity(v.native(), C.double(opacity))
}

// GetOpacity is a wrapper around gtk_widget_get_opacity()
func (v *Widget) GetOpacity() float64 {
	c := C.gtk_widget_get_opacity(v.native())
	return float64(c)
}

// HasGrab is a wrapper around gtk_widget_has_grab().
func (v *Widget) HasGrab() bool {
	c := C.gtk_widget_has_grab(v.native())
	return GoBool(c)
}

// IsDrawable is a wrapper around gtk_widget_is_drawable().
func (v *Widget) IsDrawable() bool {
	c := C.gtk_widget_is_drawable(v.native())
	return GoBool(c)
}

// IsToplevel is a wrapper around gtk_widget_is_toplevel().
func (v *Widget) IsToplevel() bool {
	c := C.gtk_widget_is_toplevel(v.native())
	return GoBool(c)
}

// GetToplevel is a wrapper around gtk_widget_get_toplevel().
func (v *Widget) GetToplevel() *Widget {
	c := C.gtk_widget_get_toplevel(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(ToGoObject(unsafe.Pointer(c)))
}

// GetTooltipMarkup is a wrapper around gtk_widget_get_tooltip_markup().
// A non-nil error is returned in the case that gtk_widget_get_tooltip_markup
// returns NULL to differentiate between NULL and an empty string.
func (v *Widget) GetTooltipMarkup() string {
	c := C.gtk_widget_get_tooltip_markup(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// SetTooltipMarkup is a wrapper around gtk_widget_set_tooltip_markup().
func (v *Widget) SetTooltipMarkup(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_tooltip_markup(v.native(), (*C.gchar)(cstr))
}

// GetTooltipText is a wrapper around gtk_widget_get_tooltip_text().
// A non-nil error is returned in the case that
// gtk_widget_get_tooltip_text returns NULL to differentiate between NULL
// and an empty string.
func (v *Widget) GetTooltipText() string {
	c := C.gtk_widget_get_tooltip_text(v.native())
	if c == nil {
		return ""
	}
	return C.GoString((*C.char)(c))
}

// SetTooltipText is a wrapper around gtk_widget_set_tooltip_text().
func (v *Widget) SetTooltipText(text string) {
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_widget_set_tooltip_text(v.native(), (*C.gchar)(cstr))
}

// TranslateCoordinates is a wrapper around gtk_widget_translate_coordinates().
func (v *Widget) TranslateCoordinates(dest IWidget, srcX, srcY int) (destX, destY int, e error) {
	var cdest *C.GtkWidget = nil
	if dest != nil {
		cdest = GtkWidget(dest)
	}
	var cdestX, cdestY C.gint
	c := C.gtk_widget_translate_coordinates(v.native(), cdest, C.gint(srcX), C.gint(srcY), &cdestX, &cdestY)
	if !GoBool(c) {
		return 0, 0, errors.New("translate coordinates failed")
	}
	return int(cdestX), int(cdestY), nil
}

// SetVisual is a wrapper around gtk_widget_set_visual().
func (v *Widget) SetVisual(visual IVisual) {
	C.gtk_widget_set_visual(v.native(), (*C.GdkVisual)(unsafe.Pointer(visual.(*Visual).Native())))
}

// SetAppPaintable is a wrapper around gtk_widget_set_app_paintable().
func (v *Widget) SetAppPaintable(paintable bool) {
	C.gtk_widget_set_app_paintable(v.native(), CBool(paintable))
}

// GetAppPaintable is a wrapper around gtk_widget_get_app_paintable().
func (v *Widget) GetAppPaintable() bool {
	c := C.gtk_widget_get_app_paintable(v.native())
	return GoBool(c)
}

// QueueDraw is a wrapper around gtk_widget_queue_draw().
func (v *Widget) QueueDraw() {
	C.gtk_widget_queue_draw(v.native())
}

// GetAllocation is a wrapper around gtk_widget_get_allocation().
func (v *Widget) GetAllocation() IRectangle {
	var a Allocation
	C.gtk_widget_get_allocation(v.native(), a.native())
	return &a
}

// SetAllocation is a wrapper around gtk_widget_set_allocation().
func (v *Widget) SetAllocation(allocation *Allocation) {
	C.gtk_widget_set_allocation(v.native(), allocation.native())
}

// SizeAllocate is a wrapper around gtk_widget_size_allocate().
func (v *Widget) SizeAllocate(allocation *Allocation) {
	C.gtk_widget_size_allocate(v.native(), allocation.native())
}

// SetStateFlags is a wrapper around gtk_widget_set_state_flags().
func (v *Widget) SetStateFlags(stateFlags StateFlags, clear bool) {
	C.gtk_widget_set_state_flags(v.native(), C.GtkStateFlags(stateFlags), CBool(clear))
}

func (v *Widget) UnsetStateFlags(stateFlags StateFlags) {
	C.gtk_widget_unset_state_flags(v.native(), C.GtkStateFlags(stateFlags))
}

func (v *Widget) GetStateFlags() StateFlags {
	return StateFlags(C.gtk_widget_get_state_flags(v.native()))
}

// GetDisplay is a wrapper around gtk_widget_get_display().
func (v *Widget) GetDisplay() *Display {
	c := C.gtk_widget_get_display(v.native())
	if c == nil {
		return nil
	}
	s := &Display{ToGoObject(unsafe.Pointer(c))}
	return s
}

// GetScreen is a wrapper around gtk_widget_get_screen().
func (v *Widget) GetScreen() IScreen {
	c := C.gtk_widget_get_screen(v.native())
	if c == nil {
		return nil
	}
	s := &Screen{ToGoObject(unsafe.Pointer(c))}
	return s
}

// SetMarginStart is a wrapper around gtk_widget_set_margin_start().
func (v *Widget) SetMarginStart(margin int) {
	C.gtk_widget_set_margin_start(v.native(), C.gint(margin))
}

// GetMarginStart is a wrapper around gtk_widget_get_margin_start().
func (v *Widget) GetMarginStart() int {
	c := C.gtk_widget_get_margin_start(v.native())
	return int(c)
}

// SetMarginEnd is a wrapper around gtk_widget_set_margin_end().
func (v *Widget) SetMarginEnd(margin int) {
	C.gtk_widget_set_margin_end(v.native(), C.gint(margin))
}

// GetMarginEnd is a wrapper around gtk_widget_get_margin_end().
func (v *Widget) GetMarginEnd() int {
	c := C.gtk_widget_get_margin_end(v.native())
	return int(c)
}

// GetFocusOnClick is a wrapper around gtk_widget_get_focus_on_click().
func (v *Widget) GetFocusOnClick() bool {
	c := C.gtk_widget_get_focus_on_click(v.native())
	return GoBool(c)
}

// SetFocusOnClick is a wrapper around gtk_widget_set_focus_on_click().
func (v *Widget) SetFocusOnClick(focusOnClick bool) {
	C.gtk_widget_set_focus_on_click(v.native(), CBool(focusOnClick))
}

// ResetStyle is a wrapper around gtk_widget_reset_style().
func (v *Widget) ResetStyle() {
	C.gtk_widget_reset_style(v.native())
}

// InDestruction is a wrapper around gtk_widget_in_destruction().
func (v *Widget) InDestruction() bool {
	return GoBool(C.gtk_widget_in_destruction(v.native()))
}

// Unparent is a wrapper around gtk_widget_unparent().
func (v *Widget) Unparent() {
	C.gtk_widget_unparent(v.native())
}

// Show is a wrapper around gtk_widget_show().
func (v *Widget) Show() {
	C.gtk_widget_show(v.native())
}

// Hide is a wrapper around gtk_widget_hide().
func (v *Widget) Hide() {
	C.gtk_widget_hide(v.native())
}

// GetCanFocus is a wrapper around gtk_widget_get_can_focus().
func (v *Widget) GetCanFocus() bool {
	c := C.gtk_widget_get_can_focus(v.native())
	return GoBool(c)
}

// SetCanFocus is a wrapper around gtk_widget_set_can_focus().
func (v *Widget) SetCanFocus(canFocus bool) {
	C.gtk_widget_set_can_focus(v.native(), CBool(canFocus))
}

// GetCanDefault is a wrapper around gtk_widget_get_can_default().
func (v *Widget) GetCanDefault() bool {
	c := C.gtk_widget_get_can_default(v.native())
	return GoBool(c)
}

// SetCanDefault is a wrapper around gtk_widget_set_can_default().
func (v *Widget) SetCanDefault(canDefault bool) {
	C.gtk_widget_set_can_default(v.native(), CBool(canDefault))
}

// SetMapped is a wrapper around gtk_widget_set_mapped().
func (v *Widget) SetMapped(mapped bool) {
	C.gtk_widget_set_mapped(v.native(), CBool(mapped))
}

// GetMapped is a wrapper around gtk_widget_get_mapped().
func (v *Widget) GetMapped() bool {
	c := C.gtk_widget_get_mapped(v.native())
	return GoBool(c)
}

// GetPreferredHeight is a wrapper around gtk_widget_get_preferred_height().
func (v *Widget) GetPreferredHeight() (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_height(v.native(), &minimum, &natural)
	return int(minimum), int(natural)
}

// GetPreferredWidth is a wrapper around gtk_widget_get_preferred_width().
func (v *Widget) GetPreferredWidth() (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_width(v.native(), &minimum, &natural)
	return int(minimum), int(natural)
}

// GetPreferredHeightForWidth is a wrapper around gtk_widget_get_preferred_height_for_width().
func (v *Widget) GetPreferredHeightForWidth(width int) (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_height_for_width(
		v.native(),
		C.gint(width),
		&minimum,
		&natural)
	return int(minimum), int(natural)
}

// GetPreferredWidthForHeight is a wrapper around gtk_widget_get_preferred_width_for_height().
func (v *Widget) GetPreferredWidthForHeight(height int) (int, int) {
	var minimum, natural C.gint
	C.gtk_widget_get_preferred_width_for_height(
		v.native(),
		C.gint(height),
		&minimum,
		&natural)
	return int(minimum), int(natural)
}

// GetRequestMode is a wrapper around gtk_widget_get_request_mode().
func (v *Widget) GetRequestMode() SizeRequestMode {
	return SizeRequestMode(C.gtk_widget_get_request_mode(v.native()))
}

// GetPreferredSize is a wrapper around gtk_widget_get_preferred_size().
func (v *Widget) GetPreferredSize() (*Requisition, *Requisition) {
	minimum_size := new(C.GtkRequisition)
	natural_size := new(C.GtkRequisition)
	C.gtk_widget_get_preferred_size(v.native(), minimum_size, natural_size)
	minR := requisitionFromNative(minimum_size)
	natR := requisitionFromNative(natural_size)
	return minR, natR
}

// DragDestSet is a wrapper around gtk_drag_dest_set().
func (v *Widget) DragDestSet(flags DestDefaults, targets []TargetEntry, actions DragAction) {
	C.gtk_drag_dest_set(v.native(), C.GtkDestDefaults(flags), (*C.GtkTargetEntry)(&targets[0]), C.gint(len(targets)), C.GdkDragAction(actions))
}

// DragSourceSet is a wrapper around gtk_drag_source_set().
func (v *Widget) DragSourceSet(startButtonMask ModifierType, targets []TargetEntry, actions DragAction) {
	C.gtk_drag_source_set(v.native(), C.GdkModifierType(startButtonMask), (*C.GtkTargetEntry)(&targets[0]), C.gint(len(targets)), C.GdkDragAction(actions))
}

func (v *Widget) DragGetData(context IDragContext, target IAtom, time uint) {
	C.gtk_drag_get_data(v.native(), context.(*DragContext).native(), Atom(target.Atom()).native(), C.uint(time))
}

func (m *Widget) IsContainer() bool {
	containerGType := TypeFromName("GtkContainer")
	return m.TypeFromInstance().IsA(containerGType)
}

// Allocation is a representation of GTK's GtkAllocation type.
type Allocation struct {
	Rectangle
}

// Native returns a pointer to the underlying GtkAllocation.
func (v *Allocation) native() *C.GtkAllocation {
	return (*C.GtkAllocation)(unsafe.Pointer(&v.GdkRectangle))
}

// TargetEntry is a representation of GTK's GtkTargetEntry
type TargetEntry C.GtkTargetEntry

func (v *TargetEntry) native() *C.GtkTargetEntry {
	return (*C.GtkTargetEntry)(unsafe.Pointer(v))
}

// NewTargetEntry is a wrapper around gtk_target_entry_new().
func NewTargetEntry(target string, flags TargetFlags, info uint) *TargetEntry {
	cstr := C.CString(target)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_target_entry_new((*C.gchar)(cstr), C.guint(flags), C.guint(info))
	if c == nil {
		return nil
	}
	t := (*TargetEntry)(unsafe.Pointer(c))
	return t
}

func (v *TargetEntry) Free() {
	C.gtk_target_entry_free(v.native())
}
