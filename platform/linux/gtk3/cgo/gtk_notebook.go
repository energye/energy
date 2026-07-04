package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

// Notebook is a representation of GTK's GtkNotebook.
type Notebook struct {
	Container
}

func (v *Notebook) native() *C.GtkNotebook {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkNotebook(p)
}

func wrapNotebook(obj *Object) *Notebook {
	return &Notebook{Container{Widget{InitiallyUnowned{obj}}}}
}

// NewNotebook is a wrapper around gtk_notebook_new().
func NewNotebook() *Notebook {
	c := C.gtk_notebook_new()
	if c == nil {
		return nil
	}
	return wrapNotebook(ToGoObject(unsafe.Pointer(c)))
}

// AppendPage is a wrapper around gtk_notebook_append_page().
func (v *Notebook) AppendPage(child IWidget, tabLabel IWidget) int {
	var tab *C.GtkWidget
	if tabLabel != nil {
		tab = GtkWidget(tabLabel)
	}
	return int(C.gtk_notebook_append_page(v.native(), GtkWidget(child), tab))
}

// RemovePage is a wrapper around gtk_notebook_remove_page().
func (v *Notebook) RemovePage(pageNum int) {
	C.gtk_notebook_remove_page(v.native(), C.gint(pageNum))
}

// GetCurrentPage is a wrapper around gtk_notebook_get_current_page().
func (v *Notebook) GetCurrentPage() int {
	return int(C.gtk_notebook_get_current_page(v.native()))
}

// SetCurrentPage is a wrapper around gtk_notebook_set_current_page().
func (v *Notebook) SetCurrentPage(pageNum int) {
	C.gtk_notebook_set_current_page(v.native(), C.gint(pageNum))
}

// GetNPages is a wrapper around gtk_notebook_get_n_pages().
func (v *Notebook) GetNPages() int {
	return int(C.gtk_notebook_get_n_pages(v.native()))
}

// SetShowTabs is a wrapper around gtk_notebook_set_show_tabs().
func (v *Notebook) SetShowTabs(showTabs bool) {
	C.gtk_notebook_set_show_tabs(v.native(), CBool(showTabs))
}

// SetShowBorder is a wrapper around gtk_notebook_set_show_border().
func (v *Notebook) SetShowBorder(showBorder bool) {
	C.gtk_notebook_set_show_border(v.native(), CBool(showBorder))
}

// SetTabPos is a wrapper around gtk_notebook_set_tab_pos().
func (v *Notebook) SetTabPos(pos PositionType) {
	C.gtk_notebook_set_tab_pos(v.native(), C.GtkPositionType(pos))
}

func (v *Notebook) SetScrollable(scrollable bool) {
	C.gtk_notebook_set_scrollable(v.native(), CBool(scrollable))
}

func (v *Notebook) GetScrollable() bool {
	return GoBool(C.gtk_notebook_get_scrollable(v.native()))
}

func (m *Notebook) SetOnSwitchPage(fn TSwitchPageEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnSwitchPage, callback.C_trampoline_4_void, fn, 0)
}

func (m *Notebook) SetOnPageAdded(fn TPageEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnPageAdded, callback.C_trampoline_3_void, fn, 0)
}

func (m *Notebook) SetOnPageRemoved(fn TPageEvent) ISignalHandlerID {
	return callback.Connect(m.Instance(), EsnPageRemoved, callback.C_trampoline_3_void, fn, 0)
}
