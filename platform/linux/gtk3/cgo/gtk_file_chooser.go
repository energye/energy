package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// static GtkWidget* _gtk_file_chooser_dialog_new(const char *title, GtkWindow *parent, GtkFileChooserAction action) {
//     return gtk_file_chooser_dialog_new(title, parent, action, NULL, NULL);
// }
import "C"
import (
	. "github.com/energye/energy/v3/platform/linux/types"
	"runtime"
	"unsafe"
)

// FileChooserDialog is a representation of GTK's GtkFileChooserDialog.
type FileChooserDialog struct {
	Dialog
}

func (v *FileChooserDialog) native() *C.GtkFileChooserDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkFileChooserDialog(unsafe.Pointer(v.GObject))
}

func wrapFileChooserDialog(obj *Object) *FileChooserDialog {
	return &FileChooserDialog{Dialog{Window{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}}
}

// NewFileChooserDialog is a wrapper around gtk_file_chooser_dialog_new().
func NewFileChooserDialog(title string, parent IWindow, action FileChooserAction) *FileChooserDialog {
	cstr := C.CString(title)
	defer C.free(unsafe.Pointer(cstr))
	var w *C.GtkWindow
	if parent != nil {
		w = (*C.GtkWindow)(unsafe.Pointer(parent.Instance()))
	}
	c := C._gtk_file_chooser_dialog_new((*C.gchar)(cstr), w, C.GtkFileChooserAction(action))
	if c == nil {
		return nil
	}
	dlg := wrapFileChooserDialog(ToGoObject(unsafe.Pointer(c)))
	dlg.AddButton("取消", int(RESPONSE_CANCEL))
	if action == FILE_CHOOSER_ACTION_SAVE {
		dlg.AddButton("保存", int(RESPONSE_ACCEPT))
	} else {
		dlg.AddButton("打开", int(RESPONSE_ACCEPT))
	}
	return dlg
}

// GetFilename is a wrapper around gtk_file_chooser_get_filename().
func (v *FileChooserDialog) GetFilename() string {
	c := C.gtk_file_chooser_get_filename(C.toGtkFileChooser(unsafe.Pointer(v.GObject)))
	if c == nil {
		return ""
	}
	s := C.GoString((*C.char)(c))
	C.g_free(C.gpointer(c))
	return s
}

// SetFilename is a wrapper around gtk_file_chooser_set_filename().
func (v *FileChooserDialog) SetFilename(filename string) bool {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	return GoBool(C.gtk_file_chooser_set_filename(C.toGtkFileChooser(unsafe.Pointer(v.GObject)), cstr))
}

// SetCurrentFolder is a wrapper around gtk_file_chooser_set_current_folder().
func (v *FileChooserDialog) SetCurrentFolder(folder string) bool {
	cstr := C.CString(folder)
	defer C.free(unsafe.Pointer(cstr))
	return GoBool(C.gtk_file_chooser_set_current_folder(C.toGtkFileChooser(unsafe.Pointer(v.GObject)), cstr))
}

// SetAction is a wrapper around gtk_file_chooser_set_action().
func (v *FileChooserDialog) SetAction(action FileChooserAction) {
	C.gtk_file_chooser_set_action(C.toGtkFileChooser(unsafe.Pointer(v.GObject)), C.GtkFileChooserAction(action))
}

// GetAction is a wrapper around gtk_file_chooser_get_action().
func (v *FileChooserDialog) GetAction() FileChooserAction {
	return FileChooserAction(C.gtk_file_chooser_get_action(C.toGtkFileChooser(unsafe.Pointer(v.GObject))))
}

// SetSelectMultiple is a wrapper around gtk_file_chooser_set_select_multiple().
func (v *FileChooserDialog) SetSelectMultiple(selectMultiple bool) {
	C.gtk_file_chooser_set_select_multiple(C.toGtkFileChooser(unsafe.Pointer(v.GObject)), CBool(selectMultiple))
}

// GetSelectMultiple is a wrapper around gtk_file_chooser_get_select_multiple().
func (v *FileChooserDialog) GetSelectMultiple() bool {
	return GoBool(C.gtk_file_chooser_get_select_multiple(C.toGtkFileChooser(unsafe.Pointer(v.GObject))))
}

// SetCurrentName is a wrapper around gtk_file_chooser_set_current_name().
func (v *FileChooserDialog) SetCurrentName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_file_chooser_set_current_name(C.toGtkFileChooser(unsafe.Pointer(v.GObject)), (*C.gchar)(cstr))
}

// GetCurrentFolder is a wrapper around gtk_file_chooser_get_current_folder().
func (v *FileChooserDialog) GetCurrentFolder() string {
	c := C.gtk_file_chooser_get_current_folder(C.toGtkFileChooser(unsafe.Pointer(v.GObject)))
	if c == nil {
		return ""
	}
	s := C.GoString((*C.char)(c))
	C.g_free(C.gpointer(c))
	return s
}

// SetFilter is a wrapper around gtk_file_chooser_set_filter().
func (v *FileChooserDialog) SetFilter(filter IFileFilter) {
	f := filter.(*FileFilter)
	if f == nil || f.Object == nil || f.GObject == nil {
		return
	}
	C.gtk_file_chooser_set_filter(C.toGtkFileChooser(unsafe.Pointer(v.GObject)),
		C.toGtkFileFilter(unsafe.Pointer(f.GObject)))
	runtime.KeepAlive(f)
}

// AddFilter is a wrapper around gtk_file_chooser_add_filter().
func (v *FileChooserDialog) AddFilter(filter IFileFilter) {
	f := filter.(*FileFilter)
	if f == nil || f.Object == nil || f.GObject == nil {
		return
	}
	C.gtk_file_chooser_add_filter(C.toGtkFileChooser(unsafe.Pointer(v.GObject)),
		C.toGtkFileFilter(unsafe.Pointer(f.GObject)))
	runtime.KeepAlive(f)
}

// GetFilter is a wrapper around gtk_file_chooser_get_filter().
func (v *FileChooserDialog) GetFilter() IFileFilter {
	c := C.gtk_file_chooser_get_filter(C.toGtkFileChooser(unsafe.Pointer(v.GObject)))
	if c == nil {
		return nil
	}
	return wrapFileFilter(ToGoObject(unsafe.Pointer(c)))
}
