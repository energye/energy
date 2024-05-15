//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package cef

// TCefRect struct type: base
//
//	Structure representing a rectangle.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_geometry.h">CEF source file: /include/internal/cef_types_geometry.h (cef_rect_t))</a>
type TCefRect struct {
	X      int32 //
	Y      int32 //
	Width  int32 //
	Height int32 //
}

// TCefPoint struct type: base
//
//	Structure representing a point.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_geometry.h">CEF source file: /include/internal/cef_types_geometry.h (cef_point_t))</a>
type TCefPoint struct {
	X int32 //
	Y int32 //
}

// TCefSize struct type: base
//
//	Structure representing a size.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_geometry.h">CEF source file: /include/internal/cef_types_geometry.h (cef_size_t))</a>
type TCefSize struct {
	Width  int32 //
	Height int32 //
}

// TCefRange struct type: base
//
//	Structure representing a range.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_range_t))</a>
type TCefRange struct {
	From uint32 //
	To   uint32 //
}

// TCefCursorInfo struct type: unrecognized
//
//	Structure representing cursor information. |buffer| will be |size.width|*|size.height|*4 bytes in size and represents a BGRA image with an upper-left origin.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cursor_info_t))</a>
type TCefCursorInfo struct {
	Hotspot          TCefPoint //
	ImageScaleFactor float32   //
	Buffer           Pointer   //
	Size             TCefSize  //
}

// TCefUrlParts struct type: complex
//
//	URL component parts.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_urlparts_t))</a>
type TCefUrlParts struct {
	instance *tCefUrlParts
	Spec     string //  The complete URL specification.
	Scheme   string //  Scheme component not including the colon (e.g., "http").
	Username string //  User name component.
	Password string //  Password component.
	Host     string //  Host component. This may be a hostname, an IPv4 address or an IPv6 literal surrounded by square brackets (e.g., "[2001:db8::1]").
	Port     string //  Port number component.
	Origin   string //  Origin contains just the scheme, host, and port from a URL. Equivalent to clearing any username and password, replacing the path with a slash, and clearing everything after that. This value will be empty for non-standard URLs.
	Path     string //  Path component including the first slash following the host.
	Query    string //  Query string component (i.e., everything following the '?').
	Fragment string //  Fragment (hash) identifier component (i.e., the string following the '#').
}

// TUrlParts struct type: complex
//
//	String version of TCefUrlParts
type TUrlParts struct {
	instance *tUrlParts
	Spec     string //
	Scheme   string //
	Username string //
	Password string //
	Host     string //
	Port     string //
	Origin   string //
	Path     string //
	Query    string //
	Fragment string //
}

// TCefInsets struct type: base
//
//	Structure representing insets.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types_geometry.h">CEF source file: /include/internal/cef_types_geometry.h (cef_insets_t))</a>
type TCefInsets struct {
	Top    int32 //
	Left   int32 //
	Bottom int32 //
	Right  int32 //
}

// TCefTouchHandleState struct type: base
//
//	Touch handle state.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_touch_handle_state_t))</a>
type TCefTouchHandleState struct {
	TouchHandleId    int32                   //  Touch handle id. Increments for each new touch handle.
	Flags            uint32                  //  Combination of TCefTouchHandleStateFlags values indicating what state is set.
	Enabled          int32                   //  Enabled state. Only set if |flags| contains CEF_THS_FLAG_ENABLED.
	Orientation      TCefHorizontalAlignment //  Orientation state. Only set if |flags| contains CEF_THS_FLAG_ORIENTATION.
	MirrorVertical   int32                   //
	MirrorHorizontal int32                   //
	Origin           TCefPoint               //  Origin state. Only set if |flags| contains CEF_THS_FLAG_ORIGIN.
	Alpha            float32                 //  Alpha state. Only set if |flags| contains CEF_THS_FLAG_ALPHA.
}

// TCefCompositionUnderline struct type: base
//
//	Structure representing IME composition underline information. This is a thin wrapper around Blink's WebCompositionUnderline class and should be kept in sync with that.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_composition_underline_t))</a>
type TCefCompositionUnderline struct {
	Range           TCefRange                     //  Underline character range.
	Color           TCefColor                     //  Text color.
	BackgroundColor TCefColor                     //  Background color.
	Thick           int32                         //  Set to true (1) for thick underline.
	Style           TCefCompositionUnderlineStyle //  Style.
}

// TCefTime struct type: base
//
//	Time information. Values should always be in UTC.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_time.h">CEF source file: /include/internal/cef_time.h (cef_time_t))</a>
type TCefTime struct {
	Year        int32 //  Four or five digit year "2007" (1601 to 30827 on Windows, 1970 to 2038 on 32-bit POSIX)
	Month       int32 //  1-based month (values 1 = January, etc.)
	DayOfWeek   int32 //  0-based day of week (0 = Sunday, etc.)
	DayOfMonth  int32 //  1-based day of month (1-31)
	Hour        int32 //  Hour within the current day (0-23)
	Minute      int32 //  Minute within the current hour (0-59)
	Second      int32 //  Second within the current minute (0-59 plus leap seconds which may take it up to 60).
	Millisecond int32 //  Milliseconds within the current second (0-999)
}

// TCefBoxLayoutSettings struct type: base
//
//	Initialization settings. Specify NULL or 0 to get the recommended default values. Many of these and other settings can also configured using command- line switches.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_box_layout_settings_t))</a>
type TCefBoxLayoutSettings struct {
	Horizontal                    int32                  //  If true (1) the layout will be horizontal, otherwise the layout will be vertical.
	InsideBorderHorizontalSpacing int32                  //  Adds additional horizontal space between the child view area and the host view border.
	InsideBorderVerticalSpacing   int32                  //  Adds additional vertical space between the child view area and the host view border.
	InsideBorderInsets            TCefInsets             //  Adds additional space around the child view area.
	BetweenChildSpacing           int32                  //  Adds additional space between child views.
	MainAxisAlignment             TCefMainAxisAlignment  //  Specifies where along the main axis the child views should be laid out.
	CrossAxisAlignment            TCefCrossAxisAlignment //  Specifies where along the cross axis the child views should be laid out.
	MinimumCrossAxisSize          int32                  //  Minimum cross axis size.
	DefaultFlex                   int32                  //  Default flex for views when none is specified via CefBoxLayout methods. Using the preferred size as the basis, free space along the main axis is distributed to views in the ratio of their flex weights. Similarly, if the views will overflow the parent, space is subtracted in these ratios. A flex of 0 means this view is not resized. Flex values must not be negative.
}

// TCefSettings struct type: complex
//
//	Initialization settings. Specify NULL or 0 to get the recommended default values. Many of these and other settings can also configured using command- line switches.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_settings_t))</a>
type TCefSettings struct {
	instance                         *tCefSettings
	NoSandbox                        int32           //  Set to true (1) to disable the sandbox for sub-processes. See cef_sandbox_win.h for requirements to enable the sandbox on Windows. Also configurable using the "no-sandbox" command-line switch.
	BrowserSubprocessPath            string          //  The path to a separate executable that will be launched for sub-processes. If this value is empty on Windows or Linux then the main process executable will be used. If this value is empty on macOS then a helper executable must exist at "Contents/Frameworks/<app> Helper.app/Contents/MacOS/<app> Helper" in the top-level app bundle. See the comments on CefExecuteProcess() for details. If this value is non-empty then it must be an absolute path. Also configurable using the "browser-subprocess-path" command-line switch.
	FrameworkDirPath                 string          //  The path to the CEF framework directory on macOS. If this value is empty then the framework must exist at "Contents/Frameworks/Chromium Embedded Framework.framework" in the top-level app bundle. If this value is non-empty then it must be an absolute path. Also configurable using the "framework-dir-path" command-line switch.
	MainBundlePath                   string          //  The path to the main bundle on macOS. If this value is empty then it defaults to the top-level app bundle. If this value is non-empty then it must be an absolute path. Also configurable using the "main-bundle-path" command-line switch.
	ChromeRuntime                    int32           //  Set to true (1) to enable use of the Chrome runtime in CEF. This feature is considered experimental and is not recommended for most users at this time. See issue #2969 for details.
	MultiThreadedMessageLoop         int32           //  Set to true (1) to have the browser process message loop run in a separate thread. If false (0) then the CefDoMessageLoopWork() function must be called from your application message loop. This option is only supported on Windows and Linux.
	ExternalMessagePump              int32           //  Set to true (1) to control browser process main (UI) thread message pump scheduling via the ICefBrowserProcessHandler.OnScheduleMessagePumpWork() callback. This option is recommended for use in combination with the CefDoMessageLoopWork() function in cases where the CEF message loop must be integrated into an existing application message loop (see additional comments and warnings on CefDoMessageLoopWork). Enabling this option is not recommended for most users; leave this option disabled and use either the CefRunMessageLoop() function or multi_threaded_message_loop if possible.
	WindowlessRenderingEnabled       int32           //  Set to true (1) to enable windowless (off-screen) rendering support. Do not enable this value if the application does not use windowless rendering as it may reduce rendering performance on some systems.
	CommandLineArgsDisabled          int32           //  Set to true (1) to disable configuration of browser process features using standard CEF and Chromium command-line arguments. Configuration can still be specified using CEF data structures or via the ICefApp.OnBeforeCommandLineProcessing() method.
	CachePath                        string          //  The location where data for the global browser cache will be stored on disk. If this value is non-empty then it must be an absolute path that is either equal to or a child directory of TCefSettings.root_cache_path. If this value is empty then browsers will be created in "incognito mode" where in-memory caches are used for storage and no data is persisted to disk. HTML5 databases such as localStorage will only persist across sessions if a cache path is specified. Can be overridden for individual CefRequestContext instances via the TCefRequestContextSettings.cache_path value. When using the Chrome runtime the "default" profile will be used if |cache_path| and |root_cache_path| have the same value.
	RootCachePath                    string          //  The root directory that all TCefSettings.cache_path and TCefRequestContextSettings.cache_path values must have in common. If this value is empty and TCefSettings.cache_path is non-empty then it will default to the TCefSettings.cache_path value. If both values are empty then the default platform-specific directory will be used ("~/.config/cef_user_data" directory on Linux, "~/Library/Application Support/CEF/User Data" directory on MacOS, "AppData\Local\CEF\User Data" directory under the user profile directory on Windows). If this value is non-empty then it must be an absolute path. Failure to set this value correctly may result in the sandbox blocking read/write access to certain files.
	PersistSessionCookies            int32           //  To persist session cookies (cookies without an expiry date or validity interval) by default when using the global cookie manager set this value to true (1). Session cookies are generally intended to be transient and most Web browsers do not persist them. A |cache_path| value must also be specified to enable this feature. Also configurable using the "persist-session-cookies" command-line switch. Can be overridden for individual CefRequestContext instances via the TCefRequestContextSettings.persist_session_cookies value.
	PersistUserPreferences           int32           //  To persist user preferences as a JSON file in the cache path directory set this value to true (1). A |cache_path| value must also be specified to enable this feature. Also configurable using the "persist-user-preferences" command-line switch. Can be overridden for individual CefRequestContext instances via the TCefRequestContextSettings.persist_user_preferences value.
	UserAgent                        string          //  Value that will be returned as the User-Agent HTTP header. If empty the default User-Agent string will be used. Also configurable using the "user-agent" command-line switch.
	UserAgentProduct                 string          //  Value that will be inserted as the product portion of the default User-Agent string. If empty the Chromium product version will be used. If |userAgent| is specified this value will be ignored. Also configurable using the "user-agent-product" command-line switch.
	Locale                           string          //  The locale string that will be passed to WebKit. If empty the default locale of "en-US" will be used. This value is ignored on Linux where locale is determined using environment variable parsing with the precedence order: LANGUAGE, LC_ALL, LC_MESSAGES and LANG. Also configurable using the "lang" command-line switch.
	LogFile                          string          //  The directory and file name to use for the debug log. If empty a default log file name and location will be used. On Windows and Linux a "debug.log" file will be written in the main executable directory. On MacOS a "~/Library/Logs/[app name]_debug.log" file will be written where [app name] is the name of the main app executable. Also configurable using the "log-file" command-line switch.
	LogSeverity                      TCefLogSeverity //  The log severity. Only messages of this severity level or higher will be logged. When set to DISABLE no messages will be written to the log file, but FATAL messages will still be output to stderr. Also configurable using the "log-severity" command-line switch with a value of "verbose", "info", "warning", "error", "fatal" or "disable".
	LogItems                         TCefLogItems    //  The log items prepended to each log line. If not set the default log items will be used. Also configurable using the "log-items" command-line switch with a value of "none" for no log items, or a comma-delimited list of values "pid", "tid", "timestamp" or "tickcount" for custom log items.
	JavascriptFlags                  string          //  Custom flags that will be used when initializing the V8 JavaScript engine. The consequences of using custom flags may not be well tested. Also configurable using the "js-flags" command-line switch.
	ResourcesDirPath                 string          //  The fully qualified path for the resources directory. If this value is empty the *.pak files must be located in the module directory on Windows/Linux or the app bundle Resources directory on MacOS. If this value is non-empty then it must be an absolute path. Also configurable using the "resources-dir-path" command-line switch.
	LocalesDirPath                   string          //  The fully qualified path for the locales directory. If this value is empty the locales directory must be located in the module directory. If this value is non-empty then it must be an absolute path. This value is ignored on MacOS where pack files are always loaded from the app bundle Resources directory. Also configurable using the "locales-dir-path" command-line switch.
	PackLoadingDisabled              int32           //  Set to true (1) to disable loading of pack files for resources and locales. A resource bundle handler must be provided for the browser and render processes via ICefApp.GetResourceBundleHandler() if loading of pack files is disabled. Also configurable using the "disable-pack-loading" command- line switch.
	RemoteDebuggingPort              int32           //  Set to a value between 1024 and 65535 to enable remote debugging on the specified port. Also configurable using the "remote-debugging-port" command-line switch. Remote debugging can be accessed by loading the chrome://inspect page in Google Chrome. Port numbers 9222 and 9229 are discoverable by default. Other port numbers may need to be configured via "Discover network targets" on the Devices tab.
	UncaughtExceptionStackSize       int32           //  The number of stack trace frames to capture for uncaught exceptions. Specify a positive value to enable the ICefRenderProcessHandler.OnUncaughtException() callback. Specify 0 (default value) and OnUncaughtException() will not be called. Also configurable using the "uncaught-exception-stack-size" command-line switch.
	BackgroundColor                  TCefColor       //  Background color used for the browser before a document is loaded and when no document color is specified. The alpha component must be either fully opaque (0xFF) or fully transparent (0x00). If the alpha component is fully opaque then the RGB components will be used as the background color. If the alpha component is fully transparent for a windowed browser then the default value of opaque white be used. If the alpha component is fully transparent for a windowless (off-screen) browser then transparent painting will be enabled.
	AcceptLanguageList               string          //  Comma delimited ordered list of language codes without any whitespace that will be used in the "Accept-Language" HTTP request header and "navigator.language" JS attribute. Can be overridden for individual ICefRequestContext instances via the TCefRequestContextSettingsCefRequestContextSettings.accept_language_list value.
	CookieableSchemesList            string          //  Comma delimited list of schemes supported by the associated ICefCookieManager. If |cookieable_schemes_exclude_defaults| is false (0) the default schemes ("http", "https", "ws" and "wss") will also be supported. Not specifying a |cookieable_schemes_list| value and setting |cookieable_schemes_exclude_defaults| to true (1) will disable all loading and saving of cookies. These settings will only impact the global ICefRequestContext. Individual ICefRequestContext instances can be configured via the TCefRequestContextSettings.cookieable_schemes_list and TCefRequestContextSettings.cookieable_schemes_exclude_defaults values.
	CookieableSchemesExcludeDefaults int32           //
	ChromePolicyId                   string          //  Specify an ID to enable Chrome policy management via Platform and OS-user policies. On Windows, this is a registry key like "SOFTWARE\\Policies\\Google\\Chrome". On MacOS, this is a bundle ID like "com.google.Chrome". On Linux, this is an absolute directory path like "/etc/opt/chrome/policies". Only supported with the Chrome runtime. See https://support.google.com/chrome/a/answer/9037717 for details. Chrome Browser Cloud Management integration, when enabled via the "enable-chrome-browser-cloud-management" command-line flag, will also use the specified ID. See https://support.google.com/chrome/a/answer/9116814 for details.
}

// TCefDraggableRegion struct type: base
//
//	Structure representing a draggable region.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_draggable_region_t))</a>
type TCefDraggableRegion struct {
	Bounds    TCefRect //  Bounds of the region.
	Draggable int32    //  True (1) this this region is draggable and false (0) otherwise.
}

// TCefKeyEvent struct type: base
//
//	Structure representing keyboard event information.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_key_event_t))</a>
type TCefKeyEvent struct {
	Kind                 TCefKeyEventType //  The type of keyboard event. It's called 'type' in the original CEF source code.
	Modifiers            TCefEventFlags   //  Bit flags describing any pressed modifier keys. See TCefEventFlags for values.
	WindowsKeyCode       int32            //  The Windows key code for the key event. This value is used by the DOM specification. Sometimes it comes directly from the event (i.e. on Windows) and sometimes it's determined using a mapping function. See WebCore/platform/chromium/KeyboardCodes.h for the list of values.
	NativeKeyCode        int32            //  The actual key code genenerated by the platform.
	IsSystemKey          int32            //  Indicates whether the event is considered a "system key" event (see http://msdn.microsoft.com/en-us/library/ms646286(VS.85).aspx for details). This value will always be false on non-Windows platforms.
	Character            uint16           //  The character generated by the keystroke.
	UnmodifiedCharacter  uint16           //  Same as |character| but unmodified by any concurrently-held modifiers (except shift). This is useful for working out shortcut keys.
	FocusOnEditableField int32            //  True if the focus is currently on an editable field on the page. This is useful for determining if standard key events should be intercepted.
}

// TCefPopupFeatures struct type: base
//
//	Popup window features.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_popup_features_t))</a>
type TCefPopupFeatures struct {
	X         int32 //
	XSet      int32 //
	Y         int32 //
	YSet      int32 //
	Width     int32 //
	WidthSet  int32 //
	Height    int32 //
	HeightSet int32 //
	IsPopup   int32 //  True (1) if browser interface elements should be hidden.
}

// TCefBrowserSettings struct type: complex
//
//	Browser initialization settings. Specify NULL or 0 to get the recommended default values. The consequences of using custom values may not be well tested. Many of these and other settings can also configured using command- line switches.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_browser_settings_t))</a>
type TCefBrowserSettings struct {
	instance                   *tCefBrowserSettings
	WindowlessFrameRate        int32     //  The maximum rate in frames per second (fps) that ICefRenderHandler.OnPaint will be called for a windowless browser. The actual fps may be lower if the browser cannot generate frames at the requested rate. The minimum value is 1 and the maximum value is 60 (default 30). This value can also be changed dynamically via ICefBrowserHost.SetWindowlessFrameRate.
	StandardFontFamily         string    //  Font settings.
	FixedFontFamily            string    //
	SerifFontFamily            string    //
	SansSerifFontFamily        string    //
	CursiveFontFamily          string    //
	FantasyFontFamily          string    //
	DefaultFontSize            int32     //
	DefaultFixedFontSize       int32     //
	MinimumFontSize            int32     //
	MinimumLogicalFontSize     int32     //
	DefaultEncoding            string    //  Default encoding for Web content. If empty "ISO-8859-1" will be used. Also configurable using the "default-encoding" command-line switch.
	RemoteFonts                TCefState //  Controls the loading of fonts from remote sources. Also configurable using the "disable-remote-fonts" command-line switch.
	Javascript                 TCefState //  Controls whether JavaScript can be executed. Also configurable using the "disable-javascript" command-line switch.
	JavascriptCloseWindows     TCefState //  Controls whether JavaScript can be used to close windows that were not opened via JavaScript. JavaScript can still be used to close windows that were opened via JavaScript or that have no back/forward history. Also configurable using the "disable-javascript-close-windows" command-line switch.
	JavascriptAccessClipboard  TCefState //  Controls whether JavaScript can access the clipboard. Also configurable using the "disable-javascript-access-clipboard" command-line switch.
	JavascriptDomPaste         TCefState //  Controls whether DOM pasting is supported in the editor via execCommand("paste"). The |javascript_access_clipboard| setting must also be enabled. Also configurable using the "disable-javascript-dom-paste" command-line switch.
	ImageLoading               TCefState //  Controls whether image URLs will be loaded from the network. A cached image will still be rendered if requested. Also configurable using the "disable-image-loading" command-line switch.
	ImageShrinkStandaloneToFit TCefState //  Controls whether standalone images will be shrunk to fit the page. Also configurable using the "image-shrink-standalone-to-fit" command-line switch.
	TextAreaResize             TCefState //  Controls whether text areas can be resized. Also configurable using the "disable-text-area-resize" command-line switch.
	TabToLinks                 TCefState //  Controls whether the tab key can advance focus to links. Also configurable using the "disable-tab-to-links" command-line switch.
	LocalStorage               TCefState //  Controls whether local storage can be used. Also configurable using the "disable-local-storage" command-line switch.
	Databases                  TCefState //  Controls whether databases can be used. Also configurable using the "disable-databases" command-line switch.
	Webgl                      TCefState //  Controls whether WebGL can be used. Note that WebGL requires hardware support and may not work on all systems even when enabled. Also configurable using the "disable-webgl" command-line switch.
	BackgroundColor            TCefColor //  Background color used for the browser before a document is loaded and when no document color is specified. The alpha component must be either fully opaque (0xFF) or fully transparent (0x00). If the alpha component is fully opaque then the RGB components will be used as the background color. If the alpha component is fully transparent for a windowed browser then the TCefSettings.background_color value will be used. If the alpha component is fully transparent for a windowless (off-screen) browser then transparent painting will be enabled.
	ChromeStatusBubble         TCefState //  Controls whether the Chrome status bubble will be used. Only supported with the Chrome runtime. For details about the status bubble see https://www.chromium.org/user-experience/status-bubble/
	ChromeZoomBubble           TCefState //  Controls whether the Chrome zoom bubble will be shown when zooming. Only supported with the Chrome runtime.
}

// TCefScreenInfo struct type: base
//
//	Screen information used when window rendering is disabled. This structure is passed as a parameter to ICefRenderHandler.GetScreenInfo and should be filled in by the client.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_screen_info_t))</a>
type TCefScreenInfo struct {
	DeviceScaleFactor float32  //  Device scale factor. Specifies the ratio between physical and logical pixels.
	Depth             int32    //  The screen depth in bits per pixel.
	DepthPerComponent int32    //  The bits per color component. This assumes that the colors are balanced equally.
	IsMonochrome      int32    //  This can be true for black and white printers.
	Rect              TCefRect //  This is set from the rcMonitor member of MONITORINFOEX, to whit: "A RECT structure that specifies the display monitor rectangle, expressed in virtual-screen coordinates. Note that if the monitor is not the primary display monitor, some of the rectangle's coordinates may be negative values." The |rect| and |available_rect| properties are used to determine the available surface for rendering popup views.
	AvailableRect     TCefRect //  This is set from the rcWork member of MONITORINFOEX, to whit: "A RECT structure that specifies the work area rectangle of the display monitor that can be used by applications, expressed in virtual-screen coordinates. Windows uses this rectangle to maximize an application on the monitor. The rest of the area in rcMonitor contains system windows such as the task bar and side bars. Note that if the monitor is not the primary display monitor, some of the rectangle's coordinates may be negative values". The |rect| and |available_rect| properties are used to determine the available surface for rendering popup views.
}

// TCefRequestContextSettings struct type: complex
//
//	Request context initialization settings. Specify NULL or 0 to get the recommended default values.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_request_context_settings_t))</a>
type TCefRequestContextSettings struct {
	instance                         *tCefRequestContextSettings
	CachePath                        string //  The location where cache data for this request context will be stored on disk. If this value is non-empty then it must be an absolute path that is either equal to or a child directory of TCefSettings.root_cache_path. If this value is empty then browsers will be created in "incognito mode" where in-memory caches are used for storage and no data is persisted to disk. HTML5 databases such as localStorage will only persist across sessions if a cache path is specified. To share the global browser cache and related configuration set this value to match the TCefSettings.cache_path value.
	PersistSessionCookies            int32  //  To persist session cookies (cookies without an expiry date or validity interval) by default when using the global cookie manager set this value to true (1). Session cookies are generally intended to be transient and most Web browsers do not persist them. Can be set globally using the TCefSettings.persist_session_cookies value. This value will be ignored if |cache_path| is empty or if it matches the TCefSettings.cache_path value.
	PersistUserPreferences           int32  //  To persist user preferences as a JSON file in the cache path directory set this value to true (1). Can be set globally using the TCefSettings.persist_user_preferences value. This value will be ignored if |cache_path| is empty or if it matches the TCefSettings.cache_path value.
	AcceptLanguageList               string //  Comma delimited ordered list of language codes without any whitespace that will be used in the "Accept-Language" HTTP header. Can be set globally using the TCefSettings.accept_language_list value or overridden on a per- browser basis using the TCefBrowserSettings.accept_language_list value. If all values are empty then "en-US,en" will be used. This value will be ignored if |cache_path| matches the TCefSettings.cache_path value.
	CookieableSchemesList            string //  Comma delimited list of schemes supported by the associated CefCookieManager. If |cookieable_schemes_exclude_defaults| is false (0) the default schemes ("http", "https", "ws" and "wss") will also be supported. Not specifying a |cookieable_schemes_list| value and setting |cookieable_schemes_exclude_defaults| to true (1) will disable all loading and saving of cookies. These values will be ignored if |cache_path| matches the TCefSettings.cache_path value.
	CookieableSchemesExcludeDefaults int32  //
}

// TCefCookie struct type: complex
//
//	Cookie information.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_cookie_t))</a>
type TCefCookie struct {
	instance   *tCefCookie
	Name       string             //  The cookie name.
	Value      string             //  The cookie value.
	Domain     string             //  If |domain| is empty a host cookie will be created instead of a domain cookie. Domain cookies are stored with a leading "." and are visible to sub-domains whereas host cookies are not.
	Path       string             //  If |path| is non-empty only URLs at or below the path will get the cookie value.
	Secure     int32              //  If |secure| is true the cookie will only be sent for HTTPS requests.
	Httponly   int32              //  If |httponly| is true the cookie will only be sent for HTTP requests.
	Creation   TCefBaseTime       //  The cookie creation date. This is automatically populated by the system on cookie creation.
	LastAccess TCefBaseTime       //  The cookie last access date. This is automatically populated by the system on access.
	HasExpires int32              //  The cookie expiration date is only valid if |has_expires| is true.
	Expires    TCefBaseTime       //
	SameSite   TCefCookieSameSite //  Same site.
	Priority   TCefCookiePriority //  Priority.
}

// TCookie struct type: complex
//
//	Cookie information.
type TCookie struct {
	instance   *tCookie
	Name       string             //
	Value      string             //
	Domain     string             //
	Path       string             //
	Creation   TDateTime          //
	LastAccess TDateTime          //
	Expires    TDateTime          //
	Secure     bool               //
	Httponly   bool               //
	HasExpires bool               //
	SameSite   TCefCookieSameSite //
	Priority   TCefCookiePriority //
}

// TCefPdfPrintSettings struct type: complex
//
//	Structure representing PDF print settings. These values match the parameters supported by the DevTools Page.printToPDF function. See https://chromedevtools.github.io/devtools-protocol/tot/Page/#method-printToPDF
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_pdf_print_settings_t))</a>
type TCefPdfPrintSettings struct {
	instance            *tCefPdfPrintSettings
	Landscape           int32                  //  Set to true (1) for landscape mode or false (0) for portrait mode.
	PrintBackground     int32                  //  Set to true (1) to print background graphics.
	Scale               float64                //  The percentage to scale the PDF by before printing (e.g. .5 is 50%). If this value is less than or equal to zero the default value of 1.0 will be used.
	PaperWidth          float64                //  Output paper size in inches. If either of these values is less than or equal to zero then the default paper size (letter, 8.5 x 11 inches) will be used.
	PaperHeight         float64                //
	PreferCssPageSize   int32                  //  Set to true (1) to prefer page size as defined by css. Defaults to false (0), in which case the content will be scaled to fit the paper size.
	MarginType          TCefPdfPrintMarginType //  Margin type.
	MarginTop           float64                //  Margins in inches. Only used if |margin_type| is set to PDF_PRINT_MARGIN_CUSTOM.
	MarginRight         float64                //
	MarginBottom        float64                //
	MarginLeft          float64                //
	PageRanges          string                 //  Paper ranges to print, one based, e.g., '1-5, 8, 11-13'. Pages are printed in the document order, not in the order specified, and no more than once. Defaults to empty string, which implies the entire document is printed. The page numbers are quietly capped to actual page count of the document, and ranges beyond the end of the document are ignored. If this results in no pages to print, an error is reported. It is an error to specify a range with start greater than end.
	DisplayHeaderFooter int32                  //  Set to true (1) to display the header and/or footer. Modify |header_template| and/or |footer_template| to customize the display.
	HeaderTemplate      string                 //  HTML template for the print header. Only displayed if |display_header_footer| is true (1). Should be valid HTML markup with the following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document For example, "<span class=title></span>" would generate a span containing the title.
	FooterTemplate      string                 //  HTML template for the print footer. Only displayed if |display_header_footer| is true (1). Uses the same format as |header_template|.
	GenerateTaggedPdf   int32                  //  Set to true (1) to generate tagged (accessible) PDF.
}

// TCefMouseEvent struct type: base
//
//	Structure representing mouse event information.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_mouse_event_t))</a>
type TCefMouseEvent struct {
	X         int32          //  X coordinate relative to the left side of the view.
	Y         int32          //  Y coordinate relative to the top side of the view.
	Modifiers TCefEventFlags //  Bit flags describing any pressed modifier keys. See TCefEventFlags for values.
}

// TCefTouchEvent struct type: base
//
//	Structure representing touch event information.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_touch_event_t))</a>
type TCefTouchEvent struct {
	Id            int32               //  Id of a touch point. Must be unique per touch, can be any number except -1. Note that a maximum of 16 concurrent touches will be tracked; touches beyond that will be ignored.
	X             float32             //  X coordinate relative to the left side of the view.
	Y             float32             //  Y coordinate relative to the top side of the view.
	RadiusX       float32             //  X radius in pixels. Set to 0 if not applicable.
	RadiusY       float32             //  Y radius in pixels. Set to 0 if not applicable.
	RotationAngle float32             //  Rotation angle in radians. Set to 0 if not applicable.
	Pressure      float32             //  The normalized pressure of the pointer input in the range of [0,1]. Set to 0 if not applicable.
	Type          TCefTouchEeventType //  The state of the touch point. Touches begin with one CEF_TET_PRESSED event followed by zero or more CEF_TET_MOVED events and finally one CEF_TET_RELEASED or CEF_TET_CANCELLED event. Events not respecting this order will be ignored.
	Modifiers     TCefEventFlags      //  Bit flags describing any pressed modifier keys. See TCefEventFlags for values.
	PointerType   TCefPointerType     //  The device type that caused the event.
}

// TCefAudioParameters struct type: base
//
//	Structure representing the audio parameters for setting up the audio handler.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_audio_parameters_t))</a>
type TCefAudioParameters struct {
	ChannelLayout   TCefChannelLayout //  Layout of the audio channels
	SampleRate      int32             //  Sample rate
	FramesPerBuffer int32             //  Number of frames per buffer
}

// TCefMediaSinkDeviceInfo struct type: complex
//
//	Device information for a MediaSink object. handler.
//	<a href="https://bitbucket.org/chromiumembedded/cef/src/master/include/internal/cef_types.h">CEF source file: /include/internal/cef_types.h (cef_media_sink_device_info_t))</a>
type TCefMediaSinkDeviceInfo struct {
	instance  *tCefMediaSinkDeviceInfo
	IpAddress string //
	Port      int32  //
	ModelName string //
}
