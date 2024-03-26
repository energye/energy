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
//
//	NOTE: While simple types, `IDispatch` and array are supported, and
//	`IUnknown` objects that also implement `IDispatch` are treated as `IDispatch`,
//	generic `IUnknown`, `VT_DECIMAL`, or `VT_RECORD` variant is not supported.
//	Remote JavaScript objects like callback functions are represented as an
//	`VT_DISPATCH` `VARIANT` with the object implementing `IDispatch`.  The
//	JavaScript callback method may be invoked using `DISPID_VALUE` for the
//	`DISPID`.  Such callback method invocations will return immediately and will
//	not wait for the JavaScript function to run and so will not provide the return
//	value of the JavaScript function.
//	Nested arrays are supported up to a depth of 3.  Arrays of by
//	reference types are not supported. `VT_EMPTY` and `VT_NULL` are mapped
//	into JavaScript as `null`.  In JavaScript, `null` and undefined are
//	mapped to `VT_EMPTY`.
//
//	Additionally, all host objects are exposed as
//	`window.chrome.webview.hostObjects.sync.{name}`.  Here the host objects
//	are exposed as synchronous host object proxies. These are not promises
//	and function runtimes or property access synchronously block running
//	script waiting to communicate cross process for the host code to run.
//	Accordingly the result may have reliability issues and it is recommended
//	that you use the promise-based asynchronous
//	`window.chrome.webview.hostObjects.{name}` API.
//
//	Synchronous host object proxies and asynchronous host object proxies may
//	both use a proxy to the same host object.  Remote changes made by one
//	proxy propagates to any other proxy of that same host object whether
//	the other proxies and synchronous or asynchronous.
//
//	While JavaScript is blocked on a synchronous run to native code, that
//	native code is unable to run back to JavaScript.  Attempts to do so fail
//	 with `HRESULT_FROM_WIN32(ERROR_POSSIBLE_DEADLOCK)`.
//
//	Host object proxies are JavaScript Proxy objects that intercept all
//	property get, property set, and method invocations. Properties or methods
//	 that are a part of the Function or Object prototype are run locally.
//	Additionally any property or method in the
//	`chrome.webview.hostObjects.options.forceLocalProperties`
//	array are also run locally.  This defaults to including optional methods
//	that have meaning in JavaScript like `toJSON` and `Symbol.toPrimitive`.
//	Add more to the array as required.
//
//	The `chrome.webview.hostObjects.cleanupSome` method performs a best
//	effort garbage collection on host object proxies.
//
//	The `chrome.webview.hostObjects.options` object provides the ability to
//	change some functionality of host objects.
//	<code>
//	Options property | Details
//	---|---
//	`forceLocalProperties` | This is an array of host object property names that will be run locally, instead of being called on the native host object. This defaults to `then`, `toJSON`, `Symbol.toString`, and `Symbol.toPrimitive`. You can add other properties to specify that they should be run locally on the javascript host object proxy.
//	`log` | This is a callback that will be called with debug information. For example, you can set this to `console.log.bind(console)` to have it print debug information to the console to help when troubleshooting host object usage. By default this is null.
//	`shouldSerializeDates` | By default this is false, and javascript Date objects will be sent to host objects as a string using `JSON.stringify`. You can set this property to true to have Date objects properly serialize as a `VT_DATE` when sending to the native host object, and have `VT_DATE` properties and return values create a javascript Date object.
//	`defaultSyncProxy` | When calling a method on a synchronous proxy, the result should also be a synchronous proxy. But in some cases, the sync/async context is lost (for example, when providing to native code a reference to a function, and then calling that function in native code). In these cases, the proxy will be asynchronous, unless this property is set.
//	`forceAsyncMethodMatches ` | This is an array of regular expressions. When calling a method on a synchronous proxy, the method call will be performed asynchronously if the method name matches a string or regular expression in this array. Setting this value to `Async` will make any method that ends with Async be an asynchronous method call. If an async method doesn't match here and isn't forced to be asynchronous, the method will be invoked synchronously, blocking execution of the calling JavaScript and then returning the resolution of the promise, rather than returning a promise.
//	`ignoreMemberNotFoundError` | By default, an exception is thrown when attempting to get the value of a proxy property that doesn't exist on the corresponding native class. Setting this property to `true` switches the behavior to match Chakra WinRT projection (and general JavaScript) behavior of returning `undefined` with no error.
//	</code>
//	Host object proxies additionally have the following methods which run
//	locally.
//	<code>
//	Method name | Details
//	---|---
//
// `applyHostFunction`, `getHostProperty`, `setHostProperty` | Perform a method invocation, property get, or property set on the host object. Use the methods to explicitly force a method or property to run remotely if a conflicting local method or property exists.  For instance, `proxy.toString()` runs the local `toString` method on the proxy object. But proxy.applyHostFunction('toString') runs `toString` on the host proxied object instead.
// `getLocalProperty`, `setLocalProperty` | Perform property get, or property set locally.  Use the methods to force getting or setting a property on the host object proxy rather than on the host object it represents. For instance, `proxy.unknownProperty` gets the property named `unknownProperty` from the host proxied object.  But proxy.getLocalProperty('unknownProperty') gets the value of the property `unknownProperty` on the proxy object.
// `sync` | Asynchronous host object proxies expose a sync method which returns a promise for a synchronous host object proxy for the same host object.  For example, `chrome.webview.hostObjects.sample.methodCall()` returns an asynchronous host object proxy.  Use the `sync` method to obtain a synchronous host object proxy instead: `const syncProxy = await chrome.webview.hostObjects.sample.methodCall().sync()`.
// `async` | Synchronous host object proxies expose an async method which blocks and returns an asynchronous host object proxy for the same host object.  For example, `chrome.webview.hostObjects.sync.sample.methodCall()` returns a synchronous host object proxy.  Running the `async` method on this blocks and then returns an asynchronous host object proxy for the same host object: `const asyncProxy = chrome.webview.hostObjects.sync.sample.methodCall().async()`.
// `then` | Asynchronous host object proxies have a `then` method.  Allows proxies to be awaitable.  `then` returns a promise that resolves with a representation of the host object.  If the proxy represents a JavaScript literal, a copy of that is returned locally.  If the proxy represents a function, a non-awaitable proxy is returned.  If the proxy represents a JavaScript object with a mix of literal properties and function properties, the a copy of the object is returned with some properties as host object proxies.
//
//	</code>
//	All other property and method invocations (other than the above Remote
//	object proxy methods, `forceLocalProperties` list, and properties on
//	Function and Object prototypes) are run remotely.  Asynchronous host
//	object proxies return a promise representing asynchronous completion of
//	remotely invoking the method, or getting the property.  The promise
//	resolves after the remote operations complete and the promises resolve to
//	 the resulting value of the operation.  Synchronous host object proxies
//	work similarly, but block running JavaScript and wait for the remote
//	operation to complete.
//
//	Setting a property on an asynchronous host object proxy works slightly
//	differently.  The set returns immediately and the return value is the
//	value that is set.  This is a requirement of the JavaScript Proxy object.
//	If you need to asynchronously wait for the property set to complete, use
//	the `setHostProperty` method which returns a promise as described above.
//	Synchronous object property set property synchronously blocks until the
//	property is set.
func (m *TCoreWebView2) AddHostObjectToScript(name string, object OleVariant) bool {
	r1 := WVPreDef().SysCallN(-1, m.Instance())
	return GoBool(r1)
}
