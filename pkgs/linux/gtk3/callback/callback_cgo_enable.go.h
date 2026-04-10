//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------


typedef enum {
    RET_VOID,
    RET_GBOOLEAN,
    RET_GINT,
    RET_POINTER
} ReturnType;

typedef struct {
    void* args[10];   // 支持最多10个参数
    int nargs;
    void* user_data;
    void* ret;
    ReturnType retType;
} SignalCall;

// Go 导出的通用处理函数
extern void go_signal_handler_generic(SignalCall* call);

// ----------------- trampoline -----------------

// 2 参数 + void
static void c_trampoline_2_void(gpointer a1, gpointer user_data) {
    SignalCall call;
    call.nargs = 2;
    call.args[0] = a1;
    call.args[1] = user_data;
    call.user_data = user_data;
    call.retType = RET_VOID;
    call.ret = NULL;
    go_signal_handler_generic(&call);
}

// 3 参数 + void
static void c_trampoline_3_void(gpointer a1, gpointer a2, gpointer user_data) {
    SignalCall call;
    call.nargs = 3;
    call.args[0] = a1;
    call.args[1] = a2;
    call.args[2] = user_data;
    call.user_data = user_data;
    call.retType = RET_VOID;
    call.ret = NULL;
    go_signal_handler_generic(&call);
}

// 4 参数 + void
static void c_trampoline_4_void(gpointer a1, gpointer a2, gpointer a3, gpointer user_data) {
    SignalCall call;
    call.nargs = 4;
    call.args[0] = a1;
    call.args[1] = a2;
    call.args[2] = a3;
    call.args[3] = user_data;
    call.user_data = user_data;
    call.retType = RET_VOID;
    call.ret = NULL;
    go_signal_handler_generic(&call);
}

// 3 参数 + gboolean
static gboolean c_trampoline_3_gboolean(gpointer a1, gpointer a2, gpointer user_data) {
    SignalCall call;
    call.nargs = 3;
    call.args[0] = a1;
    call.args[1] = a2;
    call.args[2] = user_data;
    call.user_data = user_data;
    call.retType = RET_GBOOLEAN;
    call.ret = NULL;
    go_signal_handler_generic(&call);
    return (gboolean)(uintptr_t)(call.ret);
}

// 4 参数 + gboolean
static gboolean c_trampoline_4_gboolean(gpointer a1, gpointer a2, gpointer a3, gpointer user_data) {
    SignalCall call;
    call.nargs = 4;
    call.args[0] = a1;
    call.args[1] = a2;
    call.args[2] = a3;
    call.args[3] = user_data;
    call.user_data = user_data;
    call.retType = RET_GBOOLEAN;
    call.ret = NULL;
    go_signal_handler_generic(&call);
    return (gboolean)(uintptr_t)(call.ret);
}

// 8 参数 + void
static void c_trampoline_8_void_drag_data_received(GtkWidget *widget, GdkDragContext *context, gint x, gint y,
    GtkSelectionData *data, guint info, guint32 time, gpointer user_data) {
    SignalCall call;
    call.nargs = 8;
    call.args[0] = widget;
    call.args[1] = context;
    call.args[2] = (void*)(uintptr_t)x;
    call.args[3] = (void*)(uintptr_t)y;
    call.args[4] = data;
    call.args[5] = (void*)(uintptr_t)info;
    call.args[6] = (void*)(uintptr_t)time;
    call.args[7] = user_data;
    call.user_data = user_data;
    call.retType = RET_VOID;
    call.ret = NULL;
    go_signal_handler_generic(&call);
}

// 6 参数 + gboolean
static gboolean c_trampoline_6_void_drag_drop_motion(GtkWidget* widget, GdkDragContext* context, gint x, gint y,
    guint time, gpointer user_data) {
    SignalCall call;
    call.nargs = 6;
    call.args[0] = widget;
    call.args[1] = context;
    call.args[2] = (void*)(uintptr_t)x;
    call.args[3] = (void*)(uintptr_t)y;
    call.args[4] = (void*)(uintptr_t)time;
    call.args[5] = user_data;
    call.user_data = user_data;
    call.retType = RET_GBOOLEAN;
    call.ret = NULL;
    go_signal_handler_generic(&call);
    return (gboolean)(uintptr_t)(call.ret);
}

// 4 参数 + void
static void c_trampoline_4_void_drag_leave(GtkWidget *widget, GdkDragContext* context, guint time, gpointer user_data) {
    SignalCall call;
    call.nargs = 4;
    call.args[0] = widget;
    call.args[1] = context;
    call.args[2] = (void*)(uintptr_t)time;
    call.args[3] = user_data;
    call.user_data = user_data;
    call.retType = RET_VOID;
    call.ret = NULL;
    go_signal_handler_generic(&call);
}


// 可以继续扩展 static void  c_trampoline ,... 参数


// ----------------- 辅助函数 -----------------

typedef struct {
    const char* signal_name;
    GCallback cb;
} TrampolineMap;

// 全局表：通用函数名或指定函数名 → 对应C函数
static const TrampolineMap trampoline_table[] = {
    {"c_trampoline_2_void",                                 (GCallback)c_trampoline_2_void},
    {"c_trampoline_3_void",                                 (GCallback)c_trampoline_3_void},
    {"c_trampoline_4_void",                                 (GCallback)c_trampoline_4_void},
    {"c_trampoline_3_gboolean",                             (GCallback)c_trampoline_3_gboolean},
    {"c_trampoline_4_gboolean",                             (GCallback)c_trampoline_4_gboolean},
    {"c_trampoline_8_void_drag_data_received",              (GCallback)c_trampoline_8_void_drag_data_received},
    {"c_trampoline_6_void_drag_drop_motion",                (GCallback)c_trampoline_6_void_drag_drop_motion},
    {"c_trampoline_4_void_drag_leave",                      (GCallback)c_trampoline_4_void_drag_leave},
};

static inline GCallback get_trampoline(const char* signal_name) {
       int count = sizeof(trampoline_table) / sizeof(TrampolineMap);
       for (int i = 0; i < count; i++) {
           const TrampolineMap *map = &trampoline_table[i];
           if (strcmp(map->signal_name, signal_name) == 0) {
               return map->cb;
           }
       }
   return NULL;
}

static inline gulong go_g_signal_connect(gpointer instance,
                                         const gchar *signal,
                                         GCallback handler,
                                         gpointer data) {
    return g_signal_connect_data(instance, signal, handler, data, NULL, 0);
}



