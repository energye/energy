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
	"unsafe"
)

// Editable is a representation of GTK's GtkEditable.
type Editable struct {
	Widget
}

func AsEditable(ptr unsafe.Pointer) IEditable {
	if ptr == nil {
		return nil
	}
	m := new(Editable)
	m.instance = ptr
	return m
}

// SelectRegion is a wrapper around gtk_editable_select_region().
func (m *Editable) SelectRegion(startPos, endPos int) {
	gtk3.SysCall("gtk_editable_select_region", m.Instance(), uintptr(startPos), uintptr(endPos))
}

// GetSelectionBounds is a wrapper around gtk_editable_get_selection_bounds().
func (m *Editable) GetSelectionBounds() (start, end int, ok bool) {
	var cstart, cend int32
	r := gtk3.SysCall("gtk_editable_get_selection_bounds", m.Instance(), uintptr(unsafe.Pointer(&cstart)), uintptr(unsafe.Pointer(&cend)))
	return int(cstart), int(cend), ToGoBool(r)
}

// InsertText is a wrapper around gtk_editable_insert_text().
func (m *Editable) InsertText(newText string, position int) int {
	pos := int32(position)
	gtk3.SysCall("gtk_editable_insert_text", m.Instance(), CStr(newText), uintptr(len(newText)), uintptr(unsafe.Pointer(&pos)))
	return int(pos)
}

// DeleteText is a wrapper around gtk_editable_delete_text().
func (m *Editable) DeleteText(startPos, endPos int) {
	gtk3.SysCall("gtk_editable_delete_text", m.Instance(), uintptr(startPos), uintptr(endPos))
}

// GetChars is a wrapper around gtk_editable_get_chars().
func (m *Editable) GetChars(startPos, endPos int) string {
	r := gtk3.SysCall("gtk_editable_get_chars", m.Instance(), uintptr(startPos), uintptr(endPos))
	if r == 0 {
		return ""
	}
	defer GFree(r)
	return GoStr(r)
}

// CutClipboard is a wrapper around gtk_editable_cut_clipboard().
func (m *Editable) CutClipboard() {
	gtk3.SysCall("gtk_editable_cut_clipboard", m.Instance())
}

// CopyClipboard is a wrapper around gtk_editable_copy_clipboard().
func (m *Editable) CopyClipboard() {
	gtk3.SysCall("gtk_editable_copy_clipboard", m.Instance())
}

// PasteClipboard is a wrapper around gtk_editable_paste_clipboard().
func (m *Editable) PasteClipboard() {
	gtk3.SysCall("gtk_editable_paste_clipboard", m.Instance())
}

// DeleteSelection is a wrapper around gtk_editable_delete_selection().
func (m *Editable) DeleteSelection() {
	gtk3.SysCall("gtk_editable_delete_selection", m.Instance())
}

// SetPosition is a wrapper around gtk_editable_set_position().
func (m *Editable) SetPosition(position int) {
	gtk3.SysCall("gtk_editable_set_position", m.Instance(), uintptr(position))
}

// GetPosition is a wrapper around gtk_editable_get_position().
func (m *Editable) GetPosition() int {
	r := gtk3.SysCall("gtk_editable_get_position", m.Instance())
	return int(r)
}

// SetEditable is a wrapper around gtk_editable_set_editable().
func (m *Editable) SetEditable(isEditable bool) {
	gtk3.SysCall("gtk_editable_set_editable", m.Instance(), ToCBool(isEditable))
}

// GetEditable is a wrapper around gtk_editable_get_editable().
func (m *Editable) GetEditable() bool {
	r := gtk3.SysCall("gtk_editable_get_editable", m.Instance())
	return ToGoBool(r)
}

// CellEditable is a representation of GTK's GtkCellEditable.
type CellEditable struct {
	Widget
}

func AsCellEditable(ptr unsafe.Pointer) *CellEditable {
	if ptr == nil {
		return nil
	}
	m := new(CellEditable)
	m.instance = ptr
	return m
}

// ToEntry returns the CellEditable as an Entry.
func (m *CellEditable) ToEntry() *Entry {
	entry := new(Entry)
	entry.instance = m.instance
	return entry
}

// StartEditing is a wrapper around gtk_cell_editable_start_editing().
func (m *CellEditable) StartEditing(event *Event) {
	gtk3.SysCall("gtk_cell_editable_start_editing", m.Instance(), event.Instance())
}

// EditingDone is a wrapper around gtk_cell_editable_editing_done().
func (m *CellEditable) EditingDone() {
	gtk3.SysCall("gtk_cell_editable_editing_done", m.Instance())
}

// RemoveWidget is a wrapper around gtk_cell_editable_remove_widget().
func (m *CellEditable) RemoveWidget() {
	gtk3.SysCall("gtk_cell_editable_remove_widget", m.Instance())
}
