package cgo

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"github.com/energye/energy/v3/platform/linux/callback"
	. "github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

type InfoBar struct {
	Box
}

func (v *InfoBar) native() *C.GtkInfoBar {
	if v == nil || v.GObject == nil {
		return nil
	}
	return C.toGtkInfoBar(unsafe.Pointer(v.GObject))
}

func wrapInfoBar(obj *Object) *InfoBar {
	return &InfoBar{Box{Container{Widget{InitiallyUnowned{obj}}}}}
}

func AsInfoBar(ptr unsafe.Pointer) *InfoBar {
	return wrapInfoBar(ToGoObject(ptr))
}

func NewInfoBar() *InfoBar {
	c := C.gtk_info_bar_new()
	if c == nil {
		return nil
	}
	return wrapInfoBar(ToGoObject(unsafe.Pointer(c)))
}

func (v *InfoBar) AddActionWidget(child IWidget, responseId int) {
	C.gtk_info_bar_add_action_widget(v.native(), GtkWidget(child), C.gint(responseId))
}

func (v *InfoBar) AddButton(buttonText string, responseId int) {
	cstr := C.CString(buttonText)
	defer C.free(unsafe.Pointer(cstr))
	C.gtk_info_bar_add_button(v.native(), (*C.gchar)(cstr), C.gint(responseId))
}

func (v *InfoBar) SetResponseSensitive(responseId int, setting bool) {
	C.gtk_info_bar_set_response_sensitive(v.native(), C.gint(responseId), CBool(setting))
}

func (v *InfoBar) SetDefaultResponse(responseId int) {
	C.gtk_info_bar_set_default_response(v.native(), C.gint(responseId))
}

func (v *InfoBar) SetMessageType(messageType MessageType) {
	C.gtk_info_bar_set_message_type(v.native(), C.GtkMessageType(messageType))
}

func (v *InfoBar) GetMessageType() MessageType {
	return MessageType(C.gtk_info_bar_get_message_type(v.native()))
}

func (v *InfoBar) GetActionArea() IWidget {
	c := C.gtk_info_bar_get_action_area(v.native())
	if c == nil {
		return nil
	}
	return wrapWidget(ToGoObject(unsafe.Pointer(c)))
}

func (v *InfoBar) GetContentArea() IBox {
	c := C.gtk_info_bar_get_content_area(v.native())
	if c == nil {
		return nil
	}
	return wrapBox(ToGoObject(unsafe.Pointer(c)))
}

func (v *InfoBar) SetShowCloseButton(setting bool) {
	C.gtk_info_bar_set_show_close_button(v.native(), CBool(setting))
}

func (v *InfoBar) GetShowCloseButton() bool {
	return GoBool(C.gtk_info_bar_get_show_close_button(v.native()))
}

func (v *InfoBar) SetOnResponse(fn TResponseEvent) ISignalHandlerID {
	return callback.Connect(v.Instance(), EsnResponse, callback.C_trampoline_3_void, fn, 0)
}
