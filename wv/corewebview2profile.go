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
	. "github.com/energye/energy/v2/types"
)

// ICoreWebView2Profile Parent: IObject
//
//	Provides a set of properties to configure a Profile object.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile">See the ICoreWebView2Profile article.</a>
type ICoreWebView2Profile interface {
	IObject
	// Initialized
	//  Returns true when the interface implemented by this class is fully initialized.
	Initialized() bool // property
	// BaseIntf
	//  Returns the interface implemented by this class.
	BaseIntf() ICoreWebView2Profile // property
	// SetBaseIntf Set BaseIntf
	SetBaseIntf(AValue ICoreWebView2Profile) // property
	// ProfileName
	//  Name of the profile.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_profilename">See the ICoreWebView2Profile article.</a>
	ProfileName() string // property
	// IsInPrivateModeEnabled
	//  InPrivate mode is enabled or not.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_isinprivatemodeenabled">See the ICoreWebView2Profile article.</a>
	IsInPrivateModeEnabled() bool // property
	// ProfilePath
	//  Full path of the profile directory.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_profilepath">See the ICoreWebView2Profile article.</a>
	ProfilePath() string // property
	// DefaultDownloadFolderPath
	//  Gets the `DefaultDownloadFolderPath` property. The default value is the
	//  system default download folder path for the user.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_defaultdownloadfolderpath">See the ICoreWebView2Profile article.</a>
	DefaultDownloadFolderPath() string // property
	// SetDefaultDownloadFolderPath Set DefaultDownloadFolderPath
	SetDefaultDownloadFolderPath(AValue string) // property
	// PreferredColorScheme
	//  The PreferredColorScheme property sets the overall color scheme of the
	//  WebView2s associated with this profile. This sets the color scheme for
	//  WebView2 UI like dialogs, prompts, and context menus by setting the
	//  media feature `prefers-color-scheme` for websites to respond to.
	//  The default value for this is COREWEBVIEW2_PREFERRED_COLOR_AUTO,
	//  which will follow whatever theme the OS is currently set to.
	//  Returns the value of the `PreferredColorScheme` property.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile#get_preferredcolorscheme">See the ICoreWebView2Profile article.</a>
	PreferredColorScheme() TWVPreferredColorScheme // property
	// SetPreferredColorScheme Set PreferredColorScheme
	SetPreferredColorScheme(AValue TWVPreferredColorScheme) // property
	// PreferredTrackingPreventionLevel
	//  The `PreferredTrackingPreventionLevel` property allows you to control levels of tracking prevention for WebView2
	//  which are associated with a profile. This level would apply to the context of the profile. That is, all WebView2s
	//  sharing the same profile will be affected and also the value is persisted in the user data folder.
	//  See `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL` for descriptions of levels.
	//  If tracking prevention feature is enabled when creating the WebView2 environment, you can also disable tracking
	//  prevention later using this property and `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_NONE` value but that doesn't
	//  improves runtime performance.
	//  There is `ICoreWebView2EnvironmentOptions5::EnableTrackingPrevention` property to enable/disable tracking prevention feature
	//  for all the WebView2's created in the same environment. If enabled, `PreferredTrackingPreventionLevel` is set to
	//  `COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED` by default for all the WebView2's and profiles created in the same
	//  environment or is set to the level whatever value was last changed/persisted to the profile. If disabled
	//  `PreferredTrackingPreventionLevel` is not respected by WebView2. If `PreferredTrackingPreventionLevel` is set when the
	//  feature is disabled, the property value get changed and persisted but it will takes effect only if
	//  `ICoreWebView2EnvironmentOptions5::EnableTrackingPrevention` is true.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile3#get_preferredtrackingpreventionlevel">See the ICoreWebView2Profile3 article.</a>
	PreferredTrackingPreventionLevel() TWVTrackingPreventionLevel // property
	// SetPreferredTrackingPreventionLevel Set PreferredTrackingPreventionLevel
	SetPreferredTrackingPreventionLevel(AValue TWVTrackingPreventionLevel) // property
	// CookieManager
	//  Get the cookie manager for the profile. All CoreWebView2s associated with this
	//  profile share the same cookie values. Changes to cookies in this cookie manager apply to all
	//  CoreWebView2s associated with this profile. See ICoreWebView2CookieManager.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile5#get_cookiemanager">See the ICoreWebView2Profile5 article.</a>
	CookieManager() ICoreWebView2CookieManager // property
	// IsPasswordAutosaveEnabled
	//  IsPasswordAutosaveEnabled controls whether autosave for password
	//  information is enabled. The IsPasswordAutosaveEnabled property behaves
	//  independently of the IsGeneralAutofillEnabled property. When IsPasswordAutosaveEnabled is
	//  false, no new password data is saved and no Save/Update Password prompts are displayed.
	//  However, if there was password data already saved before disabling this setting,
	//  then that password information is auto-populated, suggestions are shown and clicking on
	//  one will populate the fields.
	//  When IsPasswordAutosaveEnabled is true, password information is auto-populated,
	//  suggestions are shown and clicking on one will populate the fields, new data
	//  is saved, and a Save/Update Password prompt is displayed.
	//  It will take effect immediately after setting. The default value is `FALSE`.
	//  This property has the same value as
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsPasswordAutosaveEnabled` and
	//  `CoreWebView2Profile.IsPasswordAutosaveEnabled` will always have the same
	//  value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile6#get_ispasswordautosaveenabled">See the ICoreWebView2Profile6 article.</a>
	IsPasswordAutosaveEnabled() bool // property
	// SetIsPasswordAutosaveEnabled Set IsPasswordAutosaveEnabled
	SetIsPasswordAutosaveEnabled(AValue bool) // property
	// IsGeneralAutofillEnabled
	//  IsGeneralAutofillEnabled controls whether autofill for information
	//  like names, street and email addresses, phone numbers, and arbitrary input
	//  is enabled. This excludes password and credit card information. When
	//  IsGeneralAutofillEnabled is false, no suggestions appear, and no new information
	//  is saved. When IsGeneralAutofillEnabled is true, information is saved, suggestions
	//  appear and clicking on one will populate the form fields.
	//  It will take effect immediately after setting. The default value is `TRUE`.
	//  This property has the same value as
	//  `CoreWebView2Settings.IsGeneralAutofillEnabled`, and changing one will
	//  change the other. All `CoreWebView2`s with the same `CoreWebView2Profile`
	//  will share the same value for this property, so for the `CoreWebView2`s
	//  with the same profile, their
	//  `CoreWebView2Settings.IsGeneralAutofillEnabled` and
	//  `CoreWebView2Profile.IsGeneralAutofillEnabled` will always have the same
	//  value.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile6#get_isgeneralautofillenabled">See the ICoreWebView2Profile6 article.</a>
	IsGeneralAutofillEnabled() bool // property
	// SetIsGeneralAutofillEnabled Set IsGeneralAutofillEnabled
	SetIsGeneralAutofillEnabled(AValue bool) // property
	// AddAllBrowserEvents
	//  Adds all the events of this class to an existing TWVBrowserBase instance.
	//  <param name="aBrowserComponent">The TWVBrowserBase instance.</param>
	AddAllBrowserEvents(aBrowserComponent IComponent) bool // function
	// ClearBrowsingData
	//  Clear browsing data based on a data type. This method takes two parameters,
	//  the first being a mask of one or more `COREWEBVIEW2_BROWSING_DATA_KINDS`. OR
	//  operation(s) can be applied to multiple `COREWEBVIEW2_BROWSING_DATA_KINDS` to
	//  create a mask representing those data types. The browsing data kinds that are
	//  supported are listed below. These data kinds follow a hierarchical structure in
	//  which nested bullet points are included in their parent bullet point's data kind.
	//  Ex: All DOM storage is encompassed in all site data which is encompassed in
	//  all profile data.<code>
	//  * All Profile
	//  * All Site Data
	//  * All DOM Storage: File Systems, Indexed DB, Local Storage, Web SQL, Cache
	//  Storage
	//  * Cookies
	//  * Disk Cache
	//  * Download History
	//  * General Autofill
	//  * Password Autosave
	//  * Browsing History
	//  * Settings</code>
	//  The completed handler will be invoked when the browsing data has been cleared and
	//  will indicate if the specified data was properly cleared. In the case in which
	//  the operation is interrupted and the corresponding data is not fully cleared
	//  the handler will return `E_ABORT` and otherwise will return `S_OK`.
	//  Because this is an asynchronous operation, code that is dependent on the cleared
	//  data must be placed in the callback of this operation.
	//  If the WebView object is closed before the clear browsing data operation
	//  has completed, the handler will be released, but not invoked. In this case
	//  the clear browsing data operation may or may not be completed.
	//  ClearBrowsingData clears the `dataKinds` regardless of timestamp.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdata">See the ICoreWebView2Profile2 article.</a>
	ClearBrowsingData(dataKinds TWVBrowsingDataKinds, handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool // function
	// ClearBrowsingDataInTimeRange
	//  ClearBrowsingDataInTimeRange behaves like ClearBrowsingData except that it
	//  takes in two additional parameters for the start and end time for which it
	//  should clear the data between. The `startTime` and `endTime`
	//  parameters correspond to the number of seconds since the UNIX epoch.
	//  `startTime` is inclusive while `endTime` is exclusive, therefore the data will
	//  be cleared between [startTime, endTime).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdataintimerange">See the ICoreWebView2Profile2 article.</a>
	ClearBrowsingDataInTimeRange(dataKinds TWVBrowsingDataKinds, startTime, endTime TDateTime, handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool // function
	// ClearBrowsingDataAll
	//  ClearBrowsingDataAll behaves like ClearBrowsingData except that it
	//  clears the entirety of the data associated with the profile it is called on.
	//  It clears the data regardless of timestamp.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile2#clearbrowsingdataall">See the ICoreWebView2Profile2 article.</a>
	ClearBrowsingDataAll(handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool // function
	// SetPermissionState
	//  Sets permission state for the given permission kind and origin
	//  asynchronously. The change persists across sessions until it is changed by
	//  another call to `SetPermissionState`, or by setting the `State` property
	//  in `PermissionRequestedEventArgs`. Setting the state to
	//  `COREWEBVIEW2_PERMISSION_STATE_DEFAULT` will erase any state saved in the
	//  profile and restore the default behavior.
	//  The origin should have a valid scheme and host(e.g. "https://www.example.com"),
	//  otherwise the method fails with `E_INVALIDARG`. Additional URI parts like
	//  path and fragment are ignored. For example, "https://wwww.example.com/app1/index.html/"
	//  is treated the same as "https://wwww.example.com". See the
	//  [MDN origin definition](https://developer.mozilla.org/en-US/docs/Glossary/Origin)
	//  for more details.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#setpermissionstate">See the ICoreWebView2Profile4 article.</a>
	SetPermissionState(PermissionKind TWVPermissionKind, origin string, State TWVPermissionState, completedHandler ICoreWebView2SetPermissionStateCompletedHandler) bool // function
	// GetNonDefaultPermissionSettings
	//  Invokes the handler with a collection of all nondefault permission settings.
	//  Use this method to get the permission state set in the current and previous
	//  sessions.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile4#getnondefaultpermissionsettings">See the ICoreWebView2Profile4 article.</a>
	GetNonDefaultPermissionSettings(completedHandler ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) bool // function
	// AddBrowserExtension
	//  Adds the [browser extension](https://developer.mozilla.org/en-US/docs/Mozilla/Add-ons/WebExtensions)
	//  using the extension path for unpacked extensions from the local device. Extension is
	//  running right after installation.
	//  The extension folder path is the topmost folder of an unpacked browser extension and
	//  contains the browser extension manifest file.
	//  If the `extensionFolderPath` is an invalid path or doesn't contain the extension manifest.json
	//  file, this function will return `ERROR_FILE_NOT_FOUND` to callers.
	//  Installed extension will default `IsEnabled` to true.
	//  When `AreBrowserExtensionsEnabled` is `FALSE`, `AddBrowserExtension` will fail and return
	//  HRESULT `ERROR_NOT_SUPPORTED`.
	//  During installation, the content of the extension is not copied to the user data folder.
	//  Once the extension is installed, changing the content of the extension will cause the
	//  extension to be removed from the installed profile.
	//  When an extension is added the extension is persisted in the corresponding profile. The
	//  extension will still be installed the next time you use this profile.
	//  When an extension is installed from a folder path, adding the same extension from the same
	//  folder path means reinstalling this extension. When two extensions with the same Id are
	//  installed, only the later installed extension will be kept.
	//  Extensions that are designed to include any UI interactions(e.g. icon, badge, pop up, etc.)
	//  can be loaded and used but will have missing UI entry points due to the lack of browser
	//  UI elements to host these entry points in WebView2.
	//  The following summarizes the possible error values that can be returned from
	//  `AddBrowserExtension` and a description of why these errors occur.
	//  <code>
	//  Error value | Description
	//  ----------------------------------------------- | --------------------------
	//  `HRESULT_FROM_WIN32(ERROR_NOT_SUPPORTED)` | Extensions are disabled.
	//  `HRESULT_FROM_WIN32(ERROR_FILE_NOT_FOUND)` | Cannot find `manifest.json` file or it is not a valid extension manifest.
	//  `E_ACCESSDENIED` | Cannot load extension with file or directory name starting with \"_\", reserved for use by the system.
	//  `E_FAIL` | Extension failed to install with other unknown reasons.
	//  </code>
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#addbrowserextension">See the ICoreWebView2Profile7 article.</a>
	AddBrowserExtension(extensionFolderPath string, completedHandler ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) bool // function
	// GetBrowserExtensions
	//  Gets a snapshot of the set of extensions installed at the time `GetBrowserExtensions` is
	//  called. If an extension is installed or uninstalled after `GetBrowserExtensions` completes,
	//  the list returned by `GetBrowserExtensions` remains the same.
	//  When `AreBrowserExtensionsEnabled` is `FALSE`, `GetBrowserExtensions` won't return any
	//  extensions on current user profile.
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile7#getbrowserextensions">See the ICoreWebView2Profile7 article.</a>
	GetBrowserExtensions(completedHandler ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) bool // function
	// Delete
	//  After the API is called, the profile will be marked for deletion. The
	//  local profile's directory will be deleted at browser process exit. If it
	//  fails to delete, because something else is holding the files open,
	//  WebView2 will try to delete the profile at all future browser process
	//  starts until successful.
	//  The corresponding CoreWebView2s will be closed and the
	//  ICoreWebView2Profile.Deleted event will be raised. See
	//  `ICoreWebView2Profile.Deleted` for more information.
	//  If you try to create a new profile with the same name as an existing
	//  profile that has been marked as deleted but hasn't yet been deleted,
	//  profile creation will fail with HRESULT_FROM_WIN32(ERROR_DELETE_PENDING).
	//  <a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile8#delete">See the ICoreWebView2Profile8 article.</a>
	Delete() bool // function
}

// TCoreWebView2Profile Parent: TObject
//
//	Provides a set of properties to configure a Profile object.
//	<a href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2profile">See the ICoreWebView2Profile article.</a>
type TCoreWebView2Profile struct {
	TObject
}

func NewCoreWebView2Profile(aBaseIntf ICoreWebView2Profile) ICoreWebView2Profile {
	r1 := WV().SysCallN(568, GetObjectUintptr(aBaseIntf))
	return AsCoreWebView2Profile(r1)
}

func (m *TCoreWebView2Profile) Initialized() bool {
	r1 := WV().SysCallN(573, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) BaseIntf() ICoreWebView2Profile {
	var resultCoreWebView2Profile uintptr
	WV().SysCallN(562, 0, m.Instance(), 0, uintptr(unsafePointer(&resultCoreWebView2Profile)))
	return AsCoreWebView2Profile(resultCoreWebView2Profile)
}

func (m *TCoreWebView2Profile) SetBaseIntf(AValue ICoreWebView2Profile) {
	WV().SysCallN(562, 1, m.Instance(), GetObjectUintptr(AValue), GetObjectUintptr(AValue))
}

func (m *TCoreWebView2Profile) ProfileName() string {
	r1 := WV().SysCallN(579, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Profile) IsInPrivateModeEnabled() bool {
	r1 := WV().SysCallN(575, m.Instance())
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) ProfilePath() string {
	r1 := WV().SysCallN(580, m.Instance())
	return GoStr(r1)
}

func (m *TCoreWebView2Profile) DefaultDownloadFolderPath() string {
	r1 := WV().SysCallN(569, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TCoreWebView2Profile) SetDefaultDownloadFolderPath(AValue string) {
	WV().SysCallN(569, 1, m.Instance(), PascalStr(AValue))
}

func (m *TCoreWebView2Profile) PreferredColorScheme() TWVPreferredColorScheme {
	r1 := WV().SysCallN(577, 0, m.Instance(), 0)
	return TWVPreferredColorScheme(r1)
}

func (m *TCoreWebView2Profile) SetPreferredColorScheme(AValue TWVPreferredColorScheme) {
	WV().SysCallN(577, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Profile) PreferredTrackingPreventionLevel() TWVTrackingPreventionLevel {
	r1 := WV().SysCallN(578, 0, m.Instance(), 0)
	return TWVTrackingPreventionLevel(r1)
}

func (m *TCoreWebView2Profile) SetPreferredTrackingPreventionLevel(AValue TWVTrackingPreventionLevel) {
	WV().SysCallN(578, 1, m.Instance(), uintptr(AValue))
}

func (m *TCoreWebView2Profile) CookieManager() ICoreWebView2CookieManager {
	var resultCoreWebView2CookieManager uintptr
	WV().SysCallN(567, m.Instance(), uintptr(unsafePointer(&resultCoreWebView2CookieManager)))
	return AsCoreWebView2CookieManager(resultCoreWebView2CookieManager)
}

func (m *TCoreWebView2Profile) IsPasswordAutosaveEnabled() bool {
	r1 := WV().SysCallN(576, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) SetIsPasswordAutosaveEnabled(AValue bool) {
	WV().SysCallN(576, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Profile) IsGeneralAutofillEnabled() bool {
	r1 := WV().SysCallN(574, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) SetIsGeneralAutofillEnabled(AValue bool) {
	WV().SysCallN(574, 1, m.Instance(), PascalBool(AValue))
}

func (m *TCoreWebView2Profile) AddAllBrowserEvents(aBrowserComponent IComponent) bool {
	r1 := WV().SysCallN(560, m.Instance(), GetObjectUintptr(aBrowserComponent))
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) ClearBrowsingData(dataKinds TWVBrowsingDataKinds, handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool {
	r1 := WV().SysCallN(564, m.Instance(), uintptr(dataKinds), GetObjectUintptr(handler))
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) ClearBrowsingDataInTimeRange(dataKinds TWVBrowsingDataKinds, startTime, endTime TDateTime, handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool {
	r1 := WV().SysCallN(566, m.Instance(), uintptr(dataKinds), uintptr(startTime), uintptr(endTime), GetObjectUintptr(handler))
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) ClearBrowsingDataAll(handler ICoreWebView2ClearBrowsingDataCompletedHandler) bool {
	r1 := WV().SysCallN(565, m.Instance(), GetObjectUintptr(handler))
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) SetPermissionState(PermissionKind TWVPermissionKind, origin string, State TWVPermissionState, completedHandler ICoreWebView2SetPermissionStateCompletedHandler) bool {
	r1 := WV().SysCallN(581, m.Instance(), uintptr(PermissionKind), PascalStr(origin), uintptr(State), GetObjectUintptr(completedHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) GetNonDefaultPermissionSettings(completedHandler ICoreWebView2GetNonDefaultPermissionSettingsCompletedHandler) bool {
	r1 := WV().SysCallN(572, m.Instance(), GetObjectUintptr(completedHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) AddBrowserExtension(extensionFolderPath string, completedHandler ICoreWebView2ProfileAddBrowserExtensionCompletedHandler) bool {
	r1 := WV().SysCallN(561, m.Instance(), PascalStr(extensionFolderPath), GetObjectUintptr(completedHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) GetBrowserExtensions(completedHandler ICoreWebView2ProfileGetBrowserExtensionsCompletedHandler) bool {
	r1 := WV().SysCallN(571, m.Instance(), GetObjectUintptr(completedHandler))
	return GoBool(r1)
}

func (m *TCoreWebView2Profile) Delete() bool {
	r1 := WV().SysCallN(570, m.Instance())
	return GoBool(r1)
}

func CoreWebView2ProfileClass() TClass {
	ret := WV().SysCallN(563)
	return TClass(ret)
}
