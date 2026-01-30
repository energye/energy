package gtk3

/*
#cgo pkg-config: gdk-3.0 gio-2.0 glib-2.0 gobject-2.0 gtk+-3.0
#include <stdlib.h>
#include <string.h>
#include <gtk/gtk.h>
#include <cairo/cairo.h>

extern void go_on_event_handler(GtkWidget* widget, gpointer user_data);
extern gboolean go_on_key_press_handler(GtkWidget* entry, GdkEventKey* event, gpointer user_data);
extern void go_on_button_press(GtkWidget* widget, GdkEventButton* event, gpointer user_data);
extern void go_on_leave_enter_notify(GtkWidget* widget, GdkEventCrossing* event, gpointer user_data);
extern gboolean go_on_window_configure(GtkWidget *window, GdkEventConfigure *event, gpointer user_data);
extern gboolean go_on_window_draw(GtkWidget *window, cairo_t *cr, gpointer user_data);

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
	"fmt"
	"sync"
	"unsafe"
)

func doOnEventHandler(widget, userData unsafe.Pointer) {
	fmt.Println("doOnEventHandler userData:", uintptr(userData))
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget}
		cb.cb(context)
	}
}

func doOnKeyPressEventHandler(widget, event, userData unsafe.Pointer) bool {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: event}
		cb.cb(context)
		result, ok := context.result.(bool)
		if ok {
			return result
		}
	}
	return false
}

func doButtonPress(widget, event, userData unsafe.Pointer) {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: event}
		cb.cb(context)
	}
}

func doLeaveEnter(widget, event, userData unsafe.Pointer) {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: event}
		cb.cb(context)
	}
}

func doOnWindowConfigure(widget, event, userData unsafe.Pointer) bool {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: event}
		cb.cb(context)
		result, ok := context.result.(bool)
		if ok {
			return result
		}
	}
	return false
}

func doOnWindowDraw(widget, cr, userData unsafe.Pointer) bool {
	id := uintptr(userData)
	if cb, ok := eventList[id]; ok {
		context := &CallbackContext{widget: widget, input: cr}
		cb.cb(context)
		result, ok := context.result.(bool)
		if ok {
			return result
		}
	}
	return false
}

//export go_on_event_handler
func go_on_event_handler(widget *C.GtkWidget, user_data C.gpointer) {
	doOnEventHandler(unsafe.Pointer(widget), unsafe.Pointer(user_data))
}

//export go_on_key_press_handler
func go_on_key_press_handler(widget *C.GtkWidget, event *C.GdkEventKey, user_data C.gpointer) C.gboolean {
	result := doOnKeyPressEventHandler(unsafe.Pointer(widget), unsafe.Pointer(event), unsafe.Pointer(user_data))
	return CBool(result)
}

//export go_on_button_press
func go_on_button_press(widget *C.GtkWidget, event *C.GdkEventButton, user_data C.gpointer) {
	doButtonPress(unsafe.Pointer(widget), unsafe.Pointer(event), unsafe.Pointer(user_data))
}

//export go_on_leave_enter_notify
func go_on_leave_enter_notify(widget *C.GtkWidget, event *C.GdkEventCrossing, user_data C.gpointer) {
	doLeaveEnter(unsafe.Pointer(widget), unsafe.Pointer(event), unsafe.Pointer(user_data))
}

//export go_on_window_configure
func go_on_window_configure(widget *C.GtkWidget, event *C.GdkEventConfigure, user_data C.gpointer) C.gboolean {
	result := doOnWindowConfigure(unsafe.Pointer(widget), unsafe.Pointer(event), unsafe.Pointer(user_data))
	return CBool(result)
}

//export go_on_window_draw
func go_on_window_draw(widget *C.GtkWidget, cr *C.cairo_t, user_data C.gpointer) C.gboolean {
	result := doOnWindowDraw(unsafe.Pointer(widget), unsafe.Pointer(cr), unsafe.Pointer(user_data))
	return CBool(result)
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
	defer C.free(unsafe.Pointer(name))
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
	default:
		cCb = C.GCallback(C.go_on_event_handler)
	}
	cWidget := GtkWidget(widget)
	sh := registerSignal(cWidget, cCb, signal)
	RegisterEvent(sh.id, cb)
	return sh
}
