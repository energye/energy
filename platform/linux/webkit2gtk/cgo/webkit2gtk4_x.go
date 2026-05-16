package cgo

import (
	"github.com/energye/energy/v3/platform/linux/callback"
	"github.com/energye/energy/v3/platform/linux/gtk3/cgo"
	. "github.com/energye/energy/v3/platform/linux/types"
	"github.com/energye/lcl/types/colors"
	"unsafe"
)

type Webkit2 struct {
	cgo.Widget
}

func AsWebkit2(ptr unsafe.Pointer) IWebkit2 {
	if ptr == nil {
		return nil
	}
	m := new(Webkit2)
	m.Object = cgo.ToGoObject(ptr)
	return m
}

func (m *Webkit2) OpenDevTools() {
	WebkitOpenDevTools(m.Instance())
}

func (m *Webkit2) SetBackgroundColor(color *colors.TARGB) {
	if color == nil {
		return
	}
	WebkitSetBackgroundColor(m.Instance(), color)
}

func (m *Webkit2) SetOnDragDataReceived(fn TDragDataReceivedEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDataReceivedEvent,
		callback.C_trampoline_8_void_drag_data_received, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragDrop(fn TDragDropEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDropEvent,
		callback.C_trampoline_6_gboolean_drag_drop_motion, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragMotion(fn TDragMotionEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragMotionEvent,
		callback.C_trampoline_6_gboolean_drag_drop_motion, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragLeave(fn TDragLeaveEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragLeaveEvent,
		callback.C_trampoline_4_void, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragDataDelete(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragDataDeleteEvent,
		callback.C_trampoline_3_void, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragBegin(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragBeginEvent,
		callback.C_trampoline_3_void, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnDragEnd(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnDragEndEvent,
		callback.C_trampoline_3_void, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnFocusIn(fn TFocusInEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnFocusInEvent, callback.C_trampoline_3_gboolean, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) SetOnFocusOut(fn TFocusOutEvent) ISignalHandlerID {
	signalHandlerID := callback.Connect(m.Instance(), EsnFocusOutEvent, callback.C_trampoline_3_gboolean, fn, 0)
	return signalHandlerID
}

func (m *Webkit2) Undo() {
	WebkitUndo(m.Instance())
}

func (m *Webkit2) CanUndo() bool {
	return WebkitCanUndo(m.Instance())
}

func (m *Webkit2) Redo() {
	WebkitRedo(m.Instance())
}

func (m *Webkit2) CanRedo() bool {
	return WebkitCanRedo(m.Instance())
}

func (m *Webkit2) Cut() {
	WebkitCut(m.Instance())
}

func (m *Webkit2) CanCut() bool {
	return WebkitCanCut(m.Instance())
}

func (m *Webkit2) Copy() {
	WebkitCopy(m.Instance())
}

func (m *Webkit2) CanCopy() bool {
	return WebkitCanCopy(m.Instance())
}

func (m *Webkit2) Paste() {
	WebkitPaste(m.Instance())
}

func (m *Webkit2) CanPaste() bool {
	return WebkitCanPaste(m.Instance())
}

func (m *Webkit2) Delete() {
	WebkitDelete(m.Instance())
}

func (m *Webkit2) CanDelete() bool {
	return WebkitCanDelete(m.Instance())
}

func (m *Webkit2) SelectAll() {
	WebkitSelectAll(m.Instance())
}
