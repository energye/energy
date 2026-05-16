//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/tool"
)

type IMenuEditing interface {
	Undo()
	CanUndo() bool
	Redo()
	CanRedo() bool
	Cut()
	CanCut() bool
	Copy()
	CanCopy() bool
	Paste()
	CanPaste() bool
	Delete()
	CanDelete() bool
	SelectAll()
}

type TMenuEditing struct {
	ActionList      lcl.IActionList
	CutAction       lcl.IEditCut
	CopyAction      lcl.IEditCopy
	PasteAction     lcl.IEditPaste
	SelectAllAction lcl.IEditSelectAll
	UndoAction      lcl.IEditUndo
	RedoAction      lcl.IAction
	DeleteAction    lcl.IEditDelete
}

func NewMenuEditing(owner lcl.IComponent) *TMenuEditing {
	m := &TMenuEditing{ActionList: lcl.NewActionList(owner)}

	m.CutAction = lcl.NewEditCut(m.ActionList)
	m.CutAction.SetShortCut(api.TextToShortCut(PlatformControl() + "+X"))
	m.CutAction.SetCaption("剪切")

	m.CopyAction = lcl.NewEditCopy(m.ActionList)
	m.CopyAction.SetShortCut(api.TextToShortCut(PlatformControl() + "+C"))
	m.CopyAction.SetCaption("复制")

	m.PasteAction = lcl.NewEditPaste(m.ActionList)
	m.PasteAction.SetShortCut(api.TextToShortCut(PlatformControl() + "+V"))
	m.PasteAction.SetCaption("粘贴")

	m.SelectAllAction = lcl.NewEditSelectAll(m.ActionList)
	m.SelectAllAction.SetShortCut(api.TextToShortCut(PlatformControl() + "+A"))
	m.SelectAllAction.SetCaption("全选")

	m.UndoAction = lcl.NewEditUndo(m.ActionList)
	m.UndoAction.SetShortCut(api.TextToShortCut(PlatformControl() + "+Z"))
	m.UndoAction.SetCaption("撤销")

	m.RedoAction = lcl.NewAction(m.ActionList)
	m.RedoAction.SetShortCut(api.TextToShortCut(PlatformControl() + "+Shift+Z"))
	m.RedoAction.SetCaption("恢复")

	m.DeleteAction = lcl.NewEditDelete(m.ActionList)
	m.DeleteAction.SetShortCut(api.TextToShortCut(PlatformControl() + "+Del"))
	m.DeleteAction.SetCaption("删除")

	m.init()
	return m
}

func PlatformControl() string {
	if tool.IsDarwin() {
		return "Meta"
	}
	return "Ctrl"
}

func (m *TMenuEditing) init() {
	m.CutAction.SetOnExecute(m.CutActionOnExecute)
	m.CutAction.SetOnUpdate(m.CutActionOnUpdate)

	m.CopyAction.SetOnExecute(m.CopyActionOnExecute)
	m.CopyAction.SetOnUpdate(m.CopyActionOnUpdate)

	m.PasteAction.SetOnExecute(m.PasteActionOnExecute)
	m.PasteAction.SetOnUpdate(m.PasteActionOnUpdate)

	m.SelectAllAction.SetOnExecute(m.SelectAllActionOnExecute)
	m.SelectAllAction.SetOnUpdate(m.SelectAllActionOnUpdate)

	m.UndoAction.SetOnExecute(m.UndoActionOnExecute)
	m.UndoAction.SetOnUpdate(m.UndoActionOnUpdate)

	m.RedoAction.SetOnExecute(m.RedoActionOnExecute)
	m.RedoAction.SetOnUpdate(m.RedoActionOnUpdate)
}

func (m *TMenuEditing) CutActionOnExecute(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] cut-execute class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			lcl.AsSynEdit(activeControl).CutToClipboard()
		} else if webview := isWebview(activeControl); webview != nil {
			webview.Cut()
		} else {
			m.CutAction.ExecuteTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) CutActionOnUpdate(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] cut-update class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TCustomEditClass()) {
			edit := lcl.AsCustomEdit(activeControl)
			enable := edit.Enabled() && !edit.ReadOnly()
			m.CutAction.SetEnabled(enable)
		} else if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			synEdit := lcl.AsSynEdit(activeControl)
			enable := synEdit.Enabled() && !synEdit.ReadOnly()
			m.CutAction.SetEnabled(enable)
		} else if webview := isWebview(activeControl); webview != nil {
			m.CutAction.SetEnabled(webview.CanCut())
		} else {
			//m.CutAction.SetEnabled(false)
			m.CutAction.UpdateTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) CopyActionOnExecute(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] copy-execute class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			lcl.AsSynEdit(activeControl).CopyToClipboard()
		} else if webview := isWebview(activeControl); webview != nil {
			webview.Copy()
		} else {
			m.CopyAction.ExecuteTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) CopyActionOnUpdate(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] copy-update class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TCustomEditClass()) {
			edit := lcl.AsCustomEdit(activeControl)
			enable := edit.Enabled() && !edit.ReadOnly()
			m.CopyAction.SetEnabled(enable)
		} else if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			edit := lcl.AsSynEdit(activeControl)
			enable := edit.Enabled() && !edit.ReadOnly()
			m.CopyAction.SetEnabled(enable)
		} else if webview := isWebview(activeControl); webview != nil {
			m.CopyAction.SetEnabled(webview.CanCopy())
		} else {
			//m.CopyAction.SetEnabled(false)
			m.CopyAction.UpdateTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) PasteActionOnExecute(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] paste-execute class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			lcl.AsSynEdit(activeControl).PasteFromClipboard(true)
		} else if webview := isWebview(activeControl); webview != nil {
			webview.Paste()
		} else {
			m.PasteAction.ExecuteTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) PasteActionOnUpdate(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] paste-update class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TCustomEditClass()) {
			edit := lcl.AsCustomEdit(activeControl)
			enable := edit.Enabled() && !edit.ReadOnly()
			m.PasteAction.SetEnabled(enable)
		} else if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			edit := lcl.AsSynEdit(activeControl)
			enable := edit.Enabled() && !edit.ReadOnly()
			m.PasteAction.SetEnabled(enable)
		} else if webview := isWebview(activeControl); webview != nil {
			m.PasteAction.SetEnabled(webview.CanPaste())
		} else {
			//m.PasteAction.SetEnabled(false)
			m.PasteAction.UpdateTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) SelectAllActionOnExecute(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] selectAll-execute class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			lcl.AsSynEdit(activeControl).SelectAll()
		} else if webview := isWebview(activeControl); webview != nil {
			webview.SelectAll()
			m.SelectAllAction.ExecuteTarget(activeControl)
		} else {
			m.SelectAllAction.ExecuteTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) SelectAllActionOnUpdate(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] selectAll-update class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TCustomEditClass()) {
			edit := lcl.AsCustomEdit(activeControl)
			enable := edit.Enabled()
			m.SelectAllAction.SetEnabled(enable)
		} else if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			edit := lcl.AsSynEdit(activeControl)
			enable := edit.Enabled()
			m.SelectAllAction.SetEnabled(enable)
		} else if webview := isWebview(activeControl); webview != nil {
			//windowParent := wvLinux.AsWkWebviewParent(activeControl)
			//webview := webkit2gtk.AsWebkit2(unsafe.Pointer(windowParent.Webview().WebView()))
			m.SelectAllAction.SetEnabled(true)
		} else {
			//m.SelectAllAction.SetEnabled(false)
			m.SelectAllAction.UpdateTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) UndoActionOnExecute(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] undo-execute class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			lcl.AsSynEdit(activeControl).Undo()
		} else if webview := isWebview(activeControl); webview != nil {
			webview.Undo()
		} else {
			m.UndoAction.ExecuteTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) UndoActionOnUpdate(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] undo-update class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TCustomEditClass()) {
			m.UndoAction.SetEnabled(lcl.AsCustomEdit(activeControl).CanUndo())
		} else if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			m.UndoAction.SetEnabled(lcl.AsSynEdit(activeControl).CanUndo())
		} else if webview := isWebview(activeControl); webview != nil {
			m.UndoAction.SetEnabled(webview.CanUndo())
		} else {
			//m.UndoAction.SetEnabled(false)
			m.UndoAction.UpdateTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) RedoActionOnExecute(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] redo-execute class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			lcl.AsSynEdit(activeControl).Redo()
		} else if webview := isWebview(activeControl); webview != nil {
			webview.Redo()
		} else {
			m.RedoAction.ExecuteTarget(activeControl)
		}
	}
}

func (m *TMenuEditing) RedoActionOnUpdate(sender lcl.IObject) {
	activeControl := lcl.Screen.ActiveControl()
	if activeControl != nil {
		println("[debug] redo-update class-name:", activeControl.ClassName(), activeControl.Caption())
		if activeControl.IsObjectInstanceOf(lcl.TSynEditClass()) {
			m.RedoAction.SetEnabled(lcl.AsSynEdit(activeControl).CanRedo())
		} else if webview := isWebview(activeControl); webview != nil {
			m.RedoAction.SetEnabled(webview.CanRedo())
		} else {
			//m.RedoAction.SetEnabled(m.UndoAction.Enabled())
			m.RedoAction.UpdateTarget(activeControl)
		}
	}
}
