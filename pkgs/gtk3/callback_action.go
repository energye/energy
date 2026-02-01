package gtk3

/*
#cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
#include <stdlib.h>
#include <string.h>
#include <gtk/gtk.h>
#include <gdk/gdk.h>
#include <cairo/cairo.h>

extern void go_on_event_handler(GtkWidget* widget, gpointer user_data);
extern gboolean go_on_key_press_handler(GtkWidget* entry, GdkEventKey* event, gpointer user_data);
extern void go_on_button_press(GtkWidget* widget, GdkEventButton* event, gpointer user_data);
extern void go_on_leave_enter_notify(GtkWidget* widget, GdkEventCrossing* event, gpointer user_data);
extern gboolean go_on_window_configure(GtkWidget *window, GdkEventConfigure *event, gpointer user_data);
extern gboolean go_on_window_draw(GtkWidget *window, cairo_t *cr, gpointer user_data);
extern void go_on_drag_data_received(GtkWidget *widget, GdkDragContext *context, gint x, gint y, GtkSelectionData *data, guint info, guint32 time, gpointer user_data);
extern gboolean go_on_drag_drop (GtkWidget* self, GdkDragContext* context, gint x, gint y, guint time, gpointer user_data);

static void remove_signal_handler(GtkWidget* widget, gulong handler_id) {
  	g_print("尝试移除信号处理器: handler_id=%lu, widget=%p\n", handler_id, widget);
    if (handler_id > 0 && widget != NULL) {
        g_signal_handler_disconnect(widget, handler_id);
        g_print("移除信号处理器成功: handler_id=%lu\n", handler_id);
    }
}
*/
import "C"
import (
	"sync"
)

//export go_on_event_handler
func go_on_event_handler(widget *C.GtkWidget, user_data C.gpointer) {
	doOnEventHandler(unsafePointer(widget), unsafePointer(user_data))
}

//export go_on_key_press_handler
func go_on_key_press_handler(widget *C.GtkWidget, event *C.GdkEventKey, user_data C.gpointer) C.gboolean {
	result := doOnKeyPressEventHandler(unsafePointer(widget), unsafePointer(event), unsafePointer(user_data))
	return CBool(result)
}

//export go_on_button_press
func go_on_button_press(widget *C.GtkWidget, event *C.GdkEventButton, user_data C.gpointer) {
	doButtonPress(unsafePointer(widget), unsafePointer(event), unsafePointer(user_data))
}

//export go_on_leave_enter_notify
func go_on_leave_enter_notify(widget *C.GtkWidget, event *C.GdkEventCrossing, user_data C.gpointer) {
	doLeaveEnter(unsafePointer(widget), unsafePointer(event), unsafePointer(user_data))
}

//export go_on_window_configure
func go_on_window_configure(widget *C.GtkWidget, event *C.GdkEventConfigure, user_data C.gpointer) C.gboolean {
	result := doOnWindowConfigure(unsafePointer(widget), unsafePointer(event), unsafePointer(user_data))
	return CBool(result)
}

//export go_on_window_draw
func go_on_window_draw(widget *C.GtkWidget, cr *C.cairo_t, user_data C.gpointer) C.gboolean {
	result := doOnWindowDraw(unsafePointer(widget), unsafePointer(cr), unsafePointer(user_data))
	return CBool(result)
}

//export go_on_drag_data_received
func go_on_drag_data_received(widget *C.GtkWidget, context *C.GdkDragContext, x C.gint, y C.gint,
	data *C.GtkSelectionData, info C.guint, time C.guint, user_data C.gpointer) {
	doOnDragDataReceived(unsafePointer(widget), unsafePointer(context), int(x), int(y),
		unsafePointer(data), uint(info), uint(time), unsafePointer(user_data))
}

//export go_on_drag_drop
func go_on_drag_drop(widget *C.GtkWidget, context *C.GdkDragContext, x C.gint, y C.gint, time C.guint, user_data C.gpointer) C.gboolean {
	result := doOnDragDrop(unsafePointer(widget), unsafePointer(context), int(x), int(y), uint(time), unsafePointer(user_data))
	return CBool(result)
}

func RegisterAction(widget IWidget, signal EventSignalName, cb *Callback) *SignalHandler {
	return registerAction(widget, signal, cb)
}
func registerAction(widget IWidget, signal EventSignalName, cb *Callback) *SignalHandler {
	var cCb C.GCallback
	switch signal {
	case EsnKeyPressEvent, EsnKeyReleaseEvent:
		cCb = C.GCallback(C.go_on_key_press_handler)
	case EsnButtonPressEvent:
		cCb = C.GCallback(C.go_on_button_press)
	case EsnEnterNotifyEvent, EsnLeaveNotifyEvent:
		cCb = C.GCallback(C.go_on_leave_enter_notify)
	case EsnConfigureEvent:
		cCb = C.GCallback(C.go_on_window_configure)
	case EsnDrawEvent:
		cCb = C.GCallback(C.go_on_window_draw)
	case EsnDragDataReceivedEvent:
		cCb = C.GCallback(C.go_on_drag_data_received)
	case EsnDragDropEvent:
		cCb = C.GCallback(C.go_on_drag_drop)
	default:
		cCb = C.GCallback(C.go_on_event_handler)
	}
	cWidget := GtkWidget(widget)
	sh := registerSignal(cWidget, cCb, signal)
	RegisterEvent(sh.id, cb)
	return sh
}

// 事件列表
var (
	eventList = make(map[uintptr]*Callback)          // event list
	eventLock sync.Mutex                             // register event lock
	gEventId  uint                          = 100000 // event id
)

// RegisterEvent 事件注册，使用控件唯一标识 + 事件类型做为事件唯一id
func RegisterEvent(id uintptr, fn *Callback) {
	eventLock.Lock()
	defer eventLock.Unlock()
	eventList[id] = fn
}

type SignalHandler struct {
	widget    *C.GtkWidget
	handlerID C.gulong
	id        uintptr
}

func (m *SignalHandler) Disconnect() {
	if m != nil && m.handlerID > 0 {
		C.remove_signal_handler(m.widget, m.handlerID)
		m.handlerID = 0
		delete(eventList, m.id)
	}
}

func (m *SignalHandler) HandlerID() uint64 {
	return uint64(m.handlerID)
}

func (m *SignalHandler) ID() int {
	return int(m.id)
}

func registerSignal(widget *C.GtkWidget, cb C.GCallback, signal EventSignalName) *SignalHandler {
	eventLock.Lock()
	defer eventLock.Unlock()
	gEventId++
	nextEventId := uintptr(gEventId)
	name := C.CString(signal)
	defer C.free(unsafePointer(name))
	pointer := C.gpointer(widget)
	handlerId := C.g_signal_connect_data(pointer, name, cb, C.gpointer(nextEventId), nil, 0)
	if handlerId == 0 {
		println("[ERROR] 连接信号失败 signal:", signal)
		return nil
	}
	return &SignalHandler{
		widget:    widget,
		handlerID: handlerId,
		id:        nextEventId,
	}
}
