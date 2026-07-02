//----------------------------------------
// Copyright © yanghy. All Rights Reserved.
// Licensed under Apache License Version 2.0, January 2004
// https://www.apache.org/licenses/LICENSE-2.0
//----------------------------------------

package nocgo

import "unsafe"

// AboutDialog is a representation of GTK's GtkAboutDialog.
type AboutDialog struct {
	Dialog
}

func AsAboutDialog(ptr unsafe.Pointer) *AboutDialog {
	if ptr == nil {
		return nil
	}
	m := new(AboutDialog)
	m.instance = ptr
	return m
}

// NewAboutDialog is a wrapper around gtk_about_dialog_new().
func NewAboutDialog() *AboutDialog {
	r := gtk3.SysCall("gtk_about_dialog_new")
	if r == 0 {
		return nil
	}
	return AsAboutDialog(unsafe.Pointer(r))
}

// SetProgramName is a wrapper around gtk_about_dialog_set_program_name().
func (m *AboutDialog) SetProgramName(name string) {
	gtk3.SysCall("gtk_about_dialog_set_program_name", m.Instance(), CStr(name))
}

// GetProgramName is a wrapper around gtk_about_dialog_get_program_name().
func (m *AboutDialog) GetProgramName() string {
	r := gtk3.SysCall("gtk_about_dialog_get_program_name", m.Instance())
	return GoStr(r)
}

// SetVersion is a wrapper around gtk_about_dialog_set_version().
func (m *AboutDialog) SetVersion(version string) {
	gtk3.SysCall("gtk_about_dialog_set_version", m.Instance(), CStr(version))
}

// GetVersion is a wrapper around gtk_about_dialog_get_version().
func (m *AboutDialog) GetVersion() string {
	r := gtk3.SysCall("gtk_about_dialog_get_version", m.Instance())
	return GoStr(r)
}

// SetComments is a wrapper around gtk_about_dialog_set_comments().
func (m *AboutDialog) SetComments(comments string) {
	gtk3.SysCall("gtk_about_dialog_set_comments", m.Instance(), CStr(comments))
}

// GetComments is a wrapper around gtk_about_dialog_get_comments().
func (m *AboutDialog) GetComments() string {
	r := gtk3.SysCall("gtk_about_dialog_get_comments", m.Instance())
	return GoStr(r)
}

// SetWebsite is a wrapper around gtk_about_dialog_set_website().
func (m *AboutDialog) SetWebsite(website string) {
	gtk3.SysCall("gtk_about_dialog_set_website", m.Instance(), CStr(website))
}

// GetWebsite is a wrapper around gtk_about_dialog_get_website().
func (m *AboutDialog) GetWebsite() string {
	r := gtk3.SysCall("gtk_about_dialog_get_website", m.Instance())
	return GoStr(r)
}

// SetWebsiteLabel is a wrapper around gtk_about_dialog_set_website_label().
func (m *AboutDialog) SetWebsiteLabel(label string) {
	gtk3.SysCall("gtk_about_dialog_set_website_label", m.Instance(), CStr(label))
}

// GetWebsiteLabel is a wrapper around gtk_about_dialog_get_website_label().
func (m *AboutDialog) GetWebsiteLabel() string {
	r := gtk3.SysCall("gtk_about_dialog_get_website_label", m.Instance())
	return GoStr(r)
}

// SetLicense is a wrapper around gtk_about_dialog_set_license().
func (m *AboutDialog) SetLicense(license string) {
	gtk3.SysCall("gtk_about_dialog_set_license", m.Instance(), CStr(license))
}

// GetLicense is a wrapper around gtk_about_dialog_get_license().
func (m *AboutDialog) GetLicense() string {
	r := gtk3.SysCall("gtk_about_dialog_get_license", m.Instance())
	return GoStr(r)
}

// SetAuthors is a wrapper around gtk_about_dialog_set_authors().
func (m *AboutDialog) SetAuthors(authors []string) {
	// Build a NULL-terminated char** array
	count := len(authors)
	arraySize := uintptr(count+1) * ptrSize
	array := glib2_0.SysCall("g_malloc", arraySize)
	if array == 0 {
		return
	}
	for i, a := range authors {
		cstr := CStr(a)
		ptrAddr := array + uintptr(i)*ptrSize
		*(*uintptr)(unsafe.Pointer(ptrAddr)) = cstr
	}
	// NULL terminator
	nullAddr := array + uintptr(count)*ptrSize
	*(*uintptr)(unsafe.Pointer(nullAddr)) = 0
	gtk3.SysCall("gtk_about_dialog_set_authors", m.Instance(), array)
	// Free the array (not the strings, GTK takes ownership)
	glib2_0.SysCall("g_free", array)
}

// SetTranslatorCredits is a wrapper around gtk_about_dialog_set_translator_credits().
func (m *AboutDialog) SetTranslatorCredits(credits string) {
	gtk3.SysCall("gtk_about_dialog_set_translator_credits", m.Instance(), CStr(credits))
}
