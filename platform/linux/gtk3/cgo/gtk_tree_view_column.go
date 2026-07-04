package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// TreeViewColumn is a representation of GTK's GtkTreeViewColumn.
type TreeViewColumn struct {
	*Object
}

func (v *TreeViewColumn) native() *C.GtkTreeViewColumn {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkTreeViewColumn(unsafe.Pointer(v.GObject))
}

func wrapTreeViewColumn(obj *Object) *TreeViewColumn {
	return &TreeViewColumn{obj}
}

// NewTreeViewColumn is a wrapper around gtk_tree_view_column_new().
func NewTreeViewColumn() *TreeViewColumn {
	c := C.gtk_tree_view_column_new()
	if c == nil {
		return nil
	}
	return wrapTreeViewColumn(ToGoObject(unsafe.Pointer(c)))
}

// SetTitle is a wrapper around gtk_tree_view_column_set_title().
func (v *TreeViewColumn) SetTitle(title string) {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_tree_view_column_set_title(v.native(), (*C.gchar)(cstr))
}

// GetTitle is a wrapper around gtk_tree_view_column_get_title().
func (v *TreeViewColumn) GetTitle() string {
	c := C.gtk_tree_view_column_get_title(v.native())
	return C.GoString((*C.char)(c))
}

// PackStart is a wrapper around gtk_tree_view_column_pack_start().
func (v *TreeViewColumn) PackStart(cell ICellRenderer, expand bool) {
	cr := cell.(cellRenderer)
	C.gtk_tree_view_column_pack_start(v.native(), cr.toCellRenderer(), CBool(expand))
}

// AddAttribute is a wrapper around gtk_tree_view_column_add_attribute().
func (v *TreeViewColumn) AddAttribute(renderer ICellRenderer, attribute string, column int) {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	cr := renderer.(cellRenderer)
	C.gtk_tree_view_column_add_attribute(v.native(), cr.toCellRenderer(), (*C.gchar)(cstr), C.gint(column))
}

func (m *TreeViewColumn) SetOnClicked(fn TNotifyEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnClicked, callback.C_trampoline_2_void, fn, 0)
}

// PackEnd is a wrapper around gtk_tree_view_column_pack_end().
func (v *TreeViewColumn) PackEnd(cell ICellRenderer, expand bool) {
	cr := cell.(cellRenderer)
	C.gtk_tree_view_column_pack_end(v.native(), cr.toCellRenderer(), CBool(expand))
}

// SetResizable is a wrapper around gtk_tree_view_column_set_resizable().
func (v *TreeViewColumn) SetResizable(resizable bool) {
	C.gtk_tree_view_column_set_resizable(v.native(), CBool(resizable))
}

// GetResizable is a wrapper around gtk_tree_view_column_get_resizable().
func (v *TreeViewColumn) GetResizable() bool {
	return GoBool(C.gtk_tree_view_column_get_resizable(v.native()))
}

// SetSizing is a wrapper around gtk_tree_view_column_set_sizing().
func (v *TreeViewColumn) SetSizing(sizing TreeViewColumnSizing) {
	C.gtk_tree_view_column_set_sizing(v.native(), C.GtkTreeViewColumnSizing(sizing))
}

// GetSizing is a wrapper around gtk_tree_view_column_get_sizing().
func (v *TreeViewColumn) GetSizing() TreeViewColumnSizing {
	return TreeViewColumnSizing(C.gtk_tree_view_column_get_sizing(v.native()))
}

// SetFixedWidth is a wrapper around gtk_tree_view_column_set_fixed_width().
func (v *TreeViewColumn) SetFixedWidth(fixedWidth int) {
	C.gtk_tree_view_column_set_fixed_width(v.native(), C.gint(fixedWidth))
}

// GetFixedWidth is a wrapper around gtk_tree_view_column_get_fixed_width().
func (v *TreeViewColumn) GetFixedWidth() int {
	return int(C.gtk_tree_view_column_get_fixed_width(v.native()))
}

// SetMinWidth is a wrapper around gtk_tree_view_column_set_min_width().
func (v *TreeViewColumn) SetMinWidth(minWidth int) {
	C.gtk_tree_view_column_set_min_width(v.native(), C.gint(minWidth))
}

// GetMinWidth is a wrapper around gtk_tree_view_column_get_min_width().
func (v *TreeViewColumn) GetMinWidth() int {
	return int(C.gtk_tree_view_column_get_min_width(v.native()))
}

// SetMaxWidth is a wrapper around gtk_tree_view_column_set_max_width().
func (v *TreeViewColumn) SetMaxWidth(maxWidth int) {
	C.gtk_tree_view_column_set_max_width(v.native(), C.gint(maxWidth))
}

// GetMaxWidth is a wrapper around gtk_tree_view_column_get_max_width().
func (v *TreeViewColumn) GetMaxWidth() int {
	return int(C.gtk_tree_view_column_get_max_width(v.native()))
}

// SetExpand is a wrapper around gtk_tree_view_column_set_expand().
func (v *TreeViewColumn) SetExpand(expand bool) {
	C.gtk_tree_view_column_set_expand(v.native(), CBool(expand))
}

// GetExpand is a wrapper around gtk_tree_view_column_get_expand().
func (v *TreeViewColumn) GetExpand() bool {
	return GoBool(C.gtk_tree_view_column_get_expand(v.native()))
}

// SetSortColumnId is a wrapper around gtk_tree_view_column_set_sort_column_id().
func (v *TreeViewColumn) SetSortColumnId(sortColumnId int) {
	C.gtk_tree_view_column_set_sort_column_id(v.native(), C.gint(sortColumnId))
}

// GetSortColumnId is a wrapper around gtk_tree_view_column_get_sort_column_id().
func (v *TreeViewColumn) GetSortColumnId() int {
	return int(C.gtk_tree_view_column_get_sort_column_id(v.native()))
}

// SetSortIndicator is a wrapper around gtk_tree_view_column_set_sort_indicator().
func (v *TreeViewColumn) SetSortIndicator(setting bool) {
	C.gtk_tree_view_column_set_sort_indicator(v.native(), CBool(setting))
}

// GetSortIndicator is a wrapper around gtk_tree_view_column_get_sort_indicator().
func (v *TreeViewColumn) GetSortIndicator() bool {
	return GoBool(C.gtk_tree_view_column_get_sort_indicator(v.native()))
}

// SetReorderable is a wrapper around gtk_tree_view_column_set_reorderable().
func (v *TreeViewColumn) SetReorderable(reorderable bool) {
	C.gtk_tree_view_column_set_reorderable(v.native(), CBool(reorderable))
}

// GetReorderable is a wrapper around gtk_tree_view_column_get_reorderable().
func (v *TreeViewColumn) GetReorderable() bool {
	return GoBool(C.gtk_tree_view_column_get_reorderable(v.native()))
}

// SetAlignment is a wrapper around gtk_tree_view_column_set_alignment().
func (v *TreeViewColumn) SetAlignment(xalign float32) {
	C.gtk_tree_view_column_set_alignment(v.native(), C.gfloat(xalign))
}

// GetAlignment is a wrapper around gtk_tree_view_column_get_alignment().
func (v *TreeViewColumn) GetAlignment() float32 {
	return float32(C.gtk_tree_view_column_get_alignment(v.native()))
}

// GetWidth is a wrapper around gtk_tree_view_column_get_width().
func (v *TreeViewColumn) GetWidth() int {
	return int(C.gtk_tree_view_column_get_width(v.native()))
}

// SetSpacing is a wrapper around gtk_tree_view_column_set_spacing().
func (v *TreeViewColumn) SetSpacing(spacing int) {
	C.gtk_tree_view_column_set_spacing(v.native(), C.gint(spacing))
}

// GetSpacing is a wrapper around gtk_tree_view_column_get_spacing().
func (v *TreeViewColumn) GetSpacing() int {
	return int(C.gtk_tree_view_column_get_spacing(v.native()))
}
