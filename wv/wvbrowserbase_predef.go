//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

import (
	. "github.com/energye/energy/v2/api"
)

// AddHostObjectToScript
//
//	Add the provided host object to script running in the WebView with the
//	specified name.  Host objects are exposed as host object proxies using
//	`window.chrome.webview.hostObjects.{name}`.  Host object proxies are
//	promises and resolves to an object representing the host object.  The
//	promise is rejected if the app has not added an object with the name.
//	When JavaScript code access a property or method of the object, a promise
//	 is return, which resolves to the value returned from the host for the
//	property or method, or rejected in case of error, for example, no
//	property or method on the object or parameters are not valid.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2#addhostobjecttoscript">See the ICoreWebView2 article.</a>
func (m *TWVBrowserBase) AddHostObjectToScript(name string, object OleVariant) bool {
	r1 := WVPreDef().SysCallN(-1, m.Instance())
	return GoBool(r1)
}
