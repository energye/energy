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

// TreeStore is a representation of GTK's GtkTreeStore.
type TreeStore struct {
	Object
}

func AsTreeStore(ptr unsafe.Pointer) *TreeStore {
	if ptr == nil {
		return nil
	}
	m := new(TreeStore)
	m.instance = ptr
	return m
}

// NewTreeStore is a wrapper around gtk_tree_store_new().
func NewTreeStore(columnTypes ...Type) *TreeStore {
	args := make([]uintptr, len(columnTypes)+1)
	args[0] = uintptr(len(columnTypes))
	for i, t := range columnTypes {
		args[i+1] = uintptr(t)
	}
	r := gtk3.SysCall("gtk_tree_store_new", args...)
	if r == 0 {
		return nil
	}
	return AsTreeStore(unsafe.Pointer(r))
}

// Append is a wrapper around gtk_tree_store_append().
// parent can be nil for top-level rows. Pass 0 for parent to append at top level.
func (m *TreeStore) Append(parent ITreeIter) ITreeIter {
	var iter TreeIter
	parentPtr := uintptr(0)
	if parent != nil {
		parentPtr = parent.Instance()
	}
	gtk3.SysCall("gtk_tree_store_append", m.Instance(), uintptr(unsafe.Pointer(&iter)), parentPtr)
	return &iter
}

// SetValue is a wrapper around gtk_tree_store_set_value() for string values.
func (m *TreeStore) SetValue(iter ITreeIter, column int, value string) {
	gValue := NewStringGValue(value)
	if gValue == nil {
		return
	}
	defer gValue.Free()
	gtk3.SysCall("gtk_tree_store_set_value", m.Instance(), iter.Instance(), uintptr(column), uintptr(unsafe.Pointer(gValue)))
}

// Remove is a wrapper around gtk_tree_store_remove().
func (m *TreeStore) Remove(iter ITreeIter) bool {
	r := gtk3.SysCall("gtk_tree_store_remove", m.Instance(), iter.Instance())
	return ToGoBool(r)
}

// Clear is a wrapper around gtk_tree_store_clear().
func (m *TreeStore) Clear() {
	gtk3.SysCall("gtk_tree_store_clear", m.Instance())
}

// Insert is a wrapper around gtk_tree_store_insert().
func (m *TreeStore) Insert(parent ITreeIter, position int) ITreeIter {
	var iter TreeIter
	parentPtr := uintptr(0)
	if parent != nil {
		parentPtr = parent.Instance()
	}
	gtk3.SysCall("gtk_tree_store_insert", m.Instance(), uintptr(unsafe.Pointer(&iter)), parentPtr, uintptr(position))
	return &iter
}

// InsertBefore is a wrapper around gtk_tree_store_insert_before().
func (m *TreeStore) InsertBefore(parent, sibling ITreeIter) ITreeIter {
	var iter TreeIter
	parentPtr := uintptr(0)
	sibPtr := uintptr(0)
	if parent != nil {
		parentPtr = parent.Instance()
	}
	if sibling != nil {
		sibPtr = sibling.Instance()
	}
	gtk3.SysCall("gtk_tree_store_insert_before", m.Instance(), uintptr(unsafe.Pointer(&iter)), parentPtr, sibPtr)
	return &iter
}

// InsertAfter is a wrapper around gtk_tree_store_insert_after().
func (m *TreeStore) InsertAfter(parent, sibling ITreeIter) ITreeIter {
	var iter TreeIter
	parentPtr := uintptr(0)
	sibPtr := uintptr(0)
	if parent != nil {
		parentPtr = parent.Instance()
	}
	if sibling != nil {
		sibPtr = sibling.Instance()
	}
	gtk3.SysCall("gtk_tree_store_insert_after", m.Instance(), uintptr(unsafe.Pointer(&iter)), parentPtr, sibPtr)
	return &iter
}

// MoveBefore is a wrapper around gtk_tree_store_move_before().
func (m *TreeStore) MoveBefore(iter, position ITreeIter) {
	gtk3.SysCall("gtk_tree_store_move_before", m.Instance(), iter.Instance(), position.Instance())
}

// MoveAfter is a wrapper around gtk_tree_store_move_after().
func (m *TreeStore) MoveAfter(iter, position ITreeIter) {
	gtk3.SysCall("gtk_tree_store_move_after", m.Instance(), iter.Instance(), position.Instance())
}

// Swap is a wrapper around gtk_tree_store_swap().
func (m *TreeStore) Swap(a, b ITreeIter) {
	gtk3.SysCall("gtk_tree_store_swap", m.Instance(), a.Instance(), b.Instance())
}

// IterIsValid is a wrapper around gtk_tree_store_iter_is_valid().
func (m *TreeStore) IterIsValid(iter ITreeIter) bool {
	r := gtk3.SysCall("gtk_tree_store_iter_is_valid", m.Instance(), iter.Instance())
	return ToGoBool(r)
}

// IsAncestor is a wrapper around gtk_tree_store_is_ancestor().
func (m *TreeStore) IsAncestor(iter, descendant ITreeIter) bool {
	r := gtk3.SysCall("gtk_tree_store_is_ancestor", m.Instance(), iter.Instance(), descendant.Instance())
	return ToGoBool(r)
}

// GetDepth is a wrapper around gtk_tree_store_iter_depth().
func (m *TreeStore) GetDepth(iter ITreeIter) int {
	return int(gtk3.SysCall("gtk_tree_store_iter_depth", m.Instance(), iter.Instance()))
}
