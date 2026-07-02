package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

// AboutDialog is a representation of GTK's GtkAboutDialog.
type AboutDialog struct {
	Dialog
}

func (v *AboutDialog) native() *C.GtkAboutDialog {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkAboutDialog(unsafe.Pointer(v.GObject))
}

func wrapAboutDialog(obj *Object) *AboutDialog {
	return &AboutDialog{Dialog{Window{Bin{Container{Widget{InitiallyUnowned{obj}}}}}}}
}

// NewAboutDialog is a wrapper around gtk_about_dialog_new().
func NewAboutDialog() *AboutDialog {
	c := C.gtk_about_dialog_new()
	if c == nil {
		return nil
	}
	return wrapAboutDialog(ToGoObject(unsafe.Pointer(c)))
}

// SetProgramName is a wrapper around gtk_about_dialog_set_program_name().
func (v *AboutDialog) SetProgramName(name string) {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_program_name(v.native(), (*C.gchar)(cstr))
}

// GetProgramName is a wrapper around gtk_about_dialog_get_program_name().
func (v *AboutDialog) GetProgramName() string {
	return C.GoString((*C.char)(C.gtk_about_dialog_get_program_name(v.native())))
}

// SetVersion is a wrapper around gtk_about_dialog_set_version().
func (v *AboutDialog) SetVersion(version string) {
	cstr := C.CString(version)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_version(v.native(), (*C.gchar)(cstr))
}

// GetVersion is a wrapper around gtk_about_dialog_get_version().
func (v *AboutDialog) GetVersion() string {
	return C.GoString((*C.char)(C.gtk_about_dialog_get_version(v.native())))
}

// SetComments is a wrapper around gtk_about_dialog_set_comments().
func (v *AboutDialog) SetComments(comments string) {
	cstr := C.CString(comments)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_comments(v.native(), (*C.gchar)(cstr))
}

// GetComments is a wrapper around gtk_about_dialog_get_comments().
func (v *AboutDialog) GetComments() string {
	return C.GoString((*C.char)(C.gtk_about_dialog_get_comments(v.native())))
}

// SetWebsite is a wrapper around gtk_about_dialog_set_website().
func (v *AboutDialog) SetWebsite(website string) {
	cstr := C.CString(website)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_website(v.native(), (*C.gchar)(cstr))
}

// GetWebsite is a wrapper around gtk_about_dialog_get_website().
func (v *AboutDialog) GetWebsite() string {
	return C.GoString((*C.char)(C.gtk_about_dialog_get_website(v.native())))
}

// SetWebsiteLabel is a wrapper around gtk_about_dialog_set_website_label().
func (v *AboutDialog) SetWebsiteLabel(label string) {
	cstr := C.CString(label)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_website_label(v.native(), (*C.gchar)(cstr))
}

// GetWebsiteLabel is a wrapper around gtk_about_dialog_get_website_label().
func (v *AboutDialog) GetWebsiteLabel() string {
	return C.GoString((*C.char)(C.gtk_about_dialog_get_website_label(v.native())))
}

// SetLicense is a wrapper around gtk_about_dialog_set_license().
func (v *AboutDialog) SetLicense(license string) {
	cstr := C.CString(license)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_license(v.native(), (*C.gchar)(cstr))
}

// GetLicense is a wrapper around gtk_about_dialog_get_license().
func (v *AboutDialog) GetLicense() string {
	return C.GoString((*C.char)(C.gtk_about_dialog_get_license(v.native())))
}

// SetAuthors is a wrapper around gtk_about_dialog_set_authors().
func (v *AboutDialog) SetAuthors(authors []string) {
	cAuthors := make([]*C.gchar, len(authors)+1)
	for i, a := range authors {
		cAuthors[i] = (*C.gchar)(C.CString(a))
		defer C.free(unsafe.Pointer(cAuthors[i]))
	}
	cAuthors[len(authors)] = nil
	C.gtk_about_dialog_set_authors(v.native(), &cAuthors[0])
}

// SetTranslatorCredits is a wrapper around gtk_about_dialog_set_translator_credits().
func (v *AboutDialog) SetTranslatorCredits(credits string) {
	cstr := C.CString(credits)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_about_dialog_set_translator_credits(v.native(), (*C.gchar)(cstr))
}
