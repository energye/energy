package cgo

/*
#cgo pkg-config: gio-2.0 glib-2.0 gobject-2.0
#include <gio/gio.h>
#include <stdlib.h>
#include <glib.h>
#include <glib-object.h>

extern void go_remove_source_func(gpointer data);
extern gboolean go_source_func(gpointer data);
*/
import "C"
import (
	"sync"
	"unsafe"
)

type SourceHandle uint

// Priority is the enumerated type for GLib priority event sources.
type Priority int

const (
	PRIORITY_HIGH         Priority = C.G_PRIORITY_HIGH
	PRIORITY_DEFAULT      Priority = C.G_PRIORITY_DEFAULT // TimeoutAdd
	PRIORITY_HIGH_IDLE    Priority = C.G_PRIORITY_HIGH_IDLE
	PRIORITY_DEFAULT_IDLE Priority = C.G_PRIORITY_DEFAULT_IDLE // IdleAdd
	PRIORITY_LOW          Priority = C.G_PRIORITY_LOW
)

type idleFunc func() bool

var (
	registry sync.Map
)

func idlePut(callback idleFunc) uintptr {
	id := uintptr(unsafe.Pointer(&callback))
	registry.Store(id, callback)
	return id
}

func idleGet(id uintptr) idleFunc {
	if fn, ok := registry.Load(id); ok {
		return fn.(idleFunc)
	}
	return nil
}

func idleDelete(id uintptr) {
	registry.Delete(id)
}

// sourceFunc is the callback for g_idle_add_full and g_timeout_add_full that
// replaces the GClosure API.
//
//export go_source_func
func go_source_func(data C.gpointer) C.gboolean {
	fn := idleGet(uintptr(data))
	if fn != nil {
		return CBool(fn())
	}
	return C.FALSE
}

//export go_remove_source_func
func go_remove_source_func(data C.gpointer) {
	idleDelete(uintptr(data))
}

// IdleAdd adds an idle source to the default main event loop context with the
// DefaultIdle priority. If f is not a function with no parameter, then IdleAdd
// will panic.
//
// After running once, the source func will be removed from the main event loop,
// unless f returns a single bool true.
func IdleAdd(fn idleFunc) SourceHandle {
	return idleAdd(PRIORITY_DEFAULT_IDLE, fn)
}

// IdleAddPriority adds an idle source to the default main event loop context
// with the given priority. Its behavior is the same as IdleAdd.
func IdleAddPriority(priority Priority, fn idleFunc) SourceHandle {
	return idleAdd(priority, fn)
}

func idleAdd(priority Priority, fn idleFunc) SourceHandle {
	id := idlePut(fn)
	h := C.g_idle_add_full(C.gint(priority), C.GSourceFunc(C.go_source_func), C.gpointer(id), C.GDestroyNotify(C.go_remove_source_func))
	return SourceHandle(h)
}
