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

// ListStore is a representation of GTK's GtkListStore.
type ListStore struct {
	Object
}

func AsListStore(ptr unsafe.Pointer) *ListStore {
	if ptr == nil {
		return nil
	}
	m := new(ListStore)
	m.instance = ptr
	return m
}

// NewListStore is a wrapper around gtk_list_store_new().
// Creates a new list store with the given column types.
func NewListStore(columnTypes ...Type) *ListStore {
	args := make([]uintptr, len(columnTypes)+1)
	args[0] = uintptr(len(columnTypes))
	for i, t := range columnTypes {
		args[i+1] = uintptr(t)
	}
	r := gtk3.SysCall("gtk_list_store_new", args...)
	if r == 0 {
		return nil
	}
	return AsListStore(unsafe.Pointer(r))
}

// Append is a wrapper around gtk_list_store_append().
func (m *ListStore) Append() ITreeIter {
	var iter TreeIter
	gtk3.SysCall("gtk_list_store_append", m.Instance(), uintptr(unsafe.Pointer(&iter)))
	return &iter
}

// SetValue is a wrapper around gtk_list_store_set_value() for string values.
func (m *ListStore) SetValue(iter ITreeIter, column int, value string) {
	// gtk_list_store_set(list_store, iter, column, value, -1);
	gValue := NewStringGValue(value)
	if gValue == nil {
		return
	}
	defer gValue.Free()
	gtk3.SysCall("gtk_list_store_set_value", m.Instance(), iter.Instance(), uintptr(column), uintptr(unsafe.Pointer(gValue)))
}

// Remove is a wrapper around gtk_list_store_remove().
func (m *ListStore) Remove(iter ITreeIter) bool {
	r := gtk3.SysCall("gtk_list_store_remove", m.Instance(), iter.Instance())
	return ToGoBool(r)
}

// Clear is a wrapper around gtk_list_store_clear().
func (m *ListStore) Clear() {
	gtk3.SysCall("gtk_list_store_clear", m.Instance())
}

// Insert is a wrapper around gtk_list_store_insert().
func (m *ListStore) Insert(position int) ITreeIter {
	var iter TreeIter
	gtk3.SysCall("gtk_list_store_insert", m.Instance(), uintptr(unsafe.Pointer(&iter)), uintptr(position))
	return &iter
}

// InsertBefore is a wrapper around gtk_list_store_insert_before().
func (m *ListStore) InsertBefore(sibling ITreeIter) ITreeIter {
	var iter TreeIter
	sibPtr := uintptr(0)
	if sibling != nil {
		sibPtr = sibling.Instance()
	}
	gtk3.SysCall("gtk_list_store_insert_before", m.Instance(), uintptr(unsafe.Pointer(&iter)), sibPtr)
	return &iter
}

// InsertAfter is a wrapper around gtk_list_store_insert_after().
func (m *ListStore) InsertAfter(sibling ITreeIter) ITreeIter {
	var iter TreeIter
	sibPtr := uintptr(0)
	if sibling != nil {
		sibPtr = sibling.Instance()
	}
	gtk3.SysCall("gtk_list_store_insert_after", m.Instance(), uintptr(unsafe.Pointer(&iter)), sibPtr)
	return &iter
}

// MoveBefore is a wrapper around gtk_list_store_move_before().
func (m *ListStore) MoveBefore(iter, position ITreeIter) {
	gtk3.SysCall("gtk_list_store_move_before", m.Instance(), iter.Instance(), position.Instance())
}

// MoveAfter is a wrapper around gtk_list_store_move_after().
func (m *ListStore) MoveAfter(iter, position ITreeIter) {
	gtk3.SysCall("gtk_list_store_move_after", m.Instance(), iter.Instance(), position.Instance())
}

// Swap is a wrapper around gtk_list_store_swap().
func (m *ListStore) Swap(a, b ITreeIter) {
	gtk3.SysCall("gtk_list_store_swap", m.Instance(), a.Instance(), b.Instance())
}

// IterIsValid is a wrapper around gtk_list_store_iter_is_valid().
func (m *ListStore) IterIsValid(iter ITreeIter) bool {
	r := gtk3.SysCall("gtk_list_store_iter_is_valid", m.Instance(), iter.Instance())
	return ToGoBool(r)
}
