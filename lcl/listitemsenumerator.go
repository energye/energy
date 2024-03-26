//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// IListItemsEnumerator Parent: IObject
type IListItemsEnumerator interface {
	IObject
	Current() IListItem // property
	MoveNext() bool     // function
}

// TListItemsEnumerator Parent: TObject
type TListItemsEnumerator struct {
	TObject
}

func NewListItemsEnumerator(AItems IListItems) IListItemsEnumerator {
	r1 := LCL().SysCallN(3407, GetObjectUintptr(AItems))
	return AsListItemsEnumerator(r1)
}

func (m *TListItemsEnumerator) Current() IListItem {
	r1 := LCL().SysCallN(3408, m.Instance())
	return AsListItem(r1)
}

func (m *TListItemsEnumerator) MoveNext() bool {
	r1 := LCL().SysCallN(3409, m.Instance())
	return GoBool(r1)
}

func ListItemsEnumeratorClass() TClass {
	ret := LCL().SysCallN(3406)
	return TClass(ret)
}
