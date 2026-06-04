//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------


#include <stdlib.h>
#include <string.h>
#include <webkit2/webkit2.h>

#ifdef __cplusplus
extern "C" {
#endif

void WebkitSetBackgroundColor(WebKitWebView *webview, gdouble r, gdouble g, gdouble b, gdouble a);
void WebkitOpenDevTools(WebKitWebView *webview);

void WebkitExecuteEditingCommand(WebKitWebView *webview, const char *command);
gboolean WebkitCanUndo(WebKitWebView *webview);
gboolean WebkitCanRedo(WebKitWebView *webview);
gboolean WebkitCanCut(WebKitWebView *webview);
gboolean WebkitCanCopy(WebKitWebView *webview);
gboolean WebkitCanPaste(WebKitWebView *webview);

#ifdef __cplusplus
}
#endif


// impl

void WebkitSetBackgroundColor(WebKitWebView *webview, gdouble r, gdouble g, gdouble b, gdouble a) {
	if (webview != NULL && WEBKIT_IS_WEB_VIEW(webview))
    {
		GdkRGBA colour = {r, g, b, a};
        webkit_web_view_set_background_color(WEBKIT_WEB_VIEW(webview), &colour);
    }
}

void WebkitOpenDevTools(WebKitWebView *webview) {
    if (!webview || !WEBKIT_IS_WEB_VIEW(webview)) return;

    WebKitSettings *settings = webkit_web_view_get_settings(webview);

    if (!webkit_settings_get_enable_developer_extras(settings)) {
        return;
    }

    WebKitWebInspector *inspector = webkit_web_view_get_inspector(webview);
    webkit_web_inspector_show(inspector);
}

void WebkitExecuteEditingCommand(WebKitWebView *webview, const char *command) {
    if (!webview || !WEBKIT_IS_WEB_VIEW(webview) || !command) return;
    webkit_web_view_execute_editing_command(WEBKIT_WEB_VIEW(webview), command);
}

gboolean WebkitCanUndo(WebKitWebView *webview) {
    if (!webview || !WEBKIT_IS_WEB_VIEW(webview)) return FALSE;
    WebKitEditorState *state = webkit_web_view_get_editor_state(webview);
    if (!state) return FALSE;
    return webkit_editor_state_is_undo_available(state);
}

gboolean WebkitCanRedo(WebKitWebView *webview) {
    if (!webview || !WEBKIT_IS_WEB_VIEW(webview)) return FALSE;
    WebKitEditorState *state = webkit_web_view_get_editor_state(webview);
    if (!state) return FALSE;
    return webkit_editor_state_is_redo_available(state);
}

gboolean WebkitCanCut(WebKitWebView *webview) {
    if (!webview || !WEBKIT_IS_WEB_VIEW(webview)) return FALSE;
    WebKitEditorState *state = webkit_web_view_get_editor_state(webview);
    if (!state) return FALSE;
    return webkit_editor_state_is_cut_available(state);
}

gboolean WebkitCanCopy(WebKitWebView *webview) {
    if (!webview || !WEBKIT_IS_WEB_VIEW(webview)) return FALSE;
    WebKitEditorState *state = webkit_web_view_get_editor_state(webview);
    if (!state) return FALSE;
    return webkit_editor_state_is_copy_available(state);
}

gboolean WebkitCanPaste(WebKitWebView *webview) {
    if (!webview || !WEBKIT_IS_WEB_VIEW(webview)) return FALSE;
    WebKitEditorState *state = webkit_web_view_get_editor_state(webview);
    if (!state) return FALSE;
    return webkit_editor_state_is_paste_available(state);
}