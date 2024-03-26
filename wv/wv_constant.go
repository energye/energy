//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package wv

type TOleEnum = uint32

// TWVKeyEventKind
//
//	Specifies the key event type that triggered an AcceleratorKeyPressed event.
//	Renamed COREWEBVIEW2_KEY_EVENT_KIND type.
type TWVKeyEventKind = COREWEBVIEW2_KEY_EVENT_KIND

// TWVMoveFocusReason
//
//	Specifies the reason for moving focus.
//	Renamed COREWEBVIEW2_MOVE_FOCUS_REASON type.
type TWVMoveFocusReason = COREWEBVIEW2_MOVE_FOCUS_REASON

// TWVWebErrorStatus
//
//	Indicates the error status values for web navigations.
//	Renamed COREWEBVIEW2_WEB_ERROR_STATUS type.
type TWVWebErrorStatus = COREWEBVIEW2_WEB_ERROR_STATUS

// TWVScriptDialogKind
//
//	Specifies the JavaScript dialog type used in the ICoreWebView2ScriptDialogOpeningEventHandler interface.
//	Renamed COREWEBVIEW2_SCRIPT_DIALOG_KIND type.
type TWVScriptDialogKind = COREWEBVIEW2_SCRIPT_DIALOG_KIND

// TWVPermissionState
//
//	Specifies the response to a permission request.
//	Renamed COREWEBVIEW2_PERMISSION_STATE type.
type TWVPermissionState = COREWEBVIEW2_PERMISSION_STATE

// TWVPermissionKind
//
//	Indicates the type of a permission request.
//	Renamed COREWEBVIEW2_PERMISSION_KIND type.
type TWVPermissionKind = COREWEBVIEW2_PERMISSION_KIND

// TWVProcessFailedKind
//
//	Specifies the process failure type used in the `ICoreWebView2ProcessFailedEventArgs` interface. The values in this enum make reference to the process kinds in the Chromium architecture. For more information about what these processes are and what they do, see [Browser Architecture - Inside look at modern web browser](https://developers.google.com/web/updates/2018/09/inside-browser-part1).
//	Renamed COREWEBVIEW2_PROCESS_FAILED_KIND type.
type TWVProcessFailedKind = COREWEBVIEW2_PROCESS_FAILED_KIND

// TWVCapturePreviewImageFormat
//
//	Specifies the image format for the ICoreWebView2.CapturePreview method.
//	Renamed COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT type.
type TWVCapturePreviewImageFormat = COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT

// TWVWebResourceContext
//
//	Specifies the web resource request contexts.
//	Renamed COREWEBVIEW2_WEB_RESOURCE_CONTEXT type.
type TWVWebResourceContext = COREWEBVIEW2_WEB_RESOURCE_CONTEXT

// TWVCookieSameSiteKind
//
//	Kind of cookie SameSite status used in the ICoreWebView2Cookie interface. These fields match those as specified in https://developer.mozilla.org/docs/Web/HTTP/Cookies#. Learn more about SameSite cookies here: https://tools.ietf.org/html/draft-west-first-party-cookies-07
//	Renamed COREWEBVIEW2_COOKIE_SAME_SITE_KIND type.
type TWVCookieSameSiteKind = COREWEBVIEW2_COOKIE_SAME_SITE_KIND

// TWVHostResourceAcccessKind
//
//	Kind of cross origin resource access allowed for host resources during download. Note that other normal access checks like same origin DOM access check and [Content Security Policy](https://developer.mozilla.org/docs/Web/HTTP/CSP) still apply. The following table illustrates the host resource cross origin access according to access context and COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND. Cross Origin Access Context | DENY | ALLOW | DENY_CORS --- | --- | --- | --- From DOM like src of img, script or iframe element| Deny | Allow | Allow From Script like Fetch or XMLHttpRequest| Deny | Allow | Deny
//	Renamed COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND type.
type TWVHostResourceAcccessKind = COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND

// TWVDownloadState
//
//	State of the download operation.
//	Renamed COREWEBVIEW2_DOWNLOAD_STATE type.
type TWVDownloadState = COREWEBVIEW2_DOWNLOAD_STATE

// TWVDownloadInterruptReason
//
//	Reason why a download was interrupted.
//	Renamed COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON type.
type TWVDownloadInterruptReason = COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON

// TWVClientCertificateKind
//
//	Specifies the client certificate kind.
//	Renamed COREWEBVIEW2_CLIENT_CERTIFICATE_KIND type.
type TWVClientCertificateKind = COREWEBVIEW2_CLIENT_CERTIFICATE_KIND

// TWVBrowserProcessExitKind
//
//	Specifies the browser process exit type used in the ICoreWebView2BrowserProcessExitedEventArgs interface.
//	Renamed COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND type.
type TWVBrowserProcessExitKind = COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND

// TWVMouseEventKind
//
//	Mouse event type used by SendMouseInput to convey the type of mouse event being sent to WebView. The values of this enum align with the matching WM_* window messages.
//	Renamed COREWEBVIEW2_MOUSE_EVENT_KIND type.
type TWVMouseEventKind = COREWEBVIEW2_MOUSE_EVENT_KIND

// TWVMouseEventVirtualKeys
//
//	Mouse event virtual keys associated with a COREWEBVIEW2_MOUSE_EVENT_KIND for SendMouseInput. These values can be combined into a bit flag if more than one virtual key is pressed for the event. The values of this enum align with the matching MK_* mouse keys.
//	Renamed COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS type.
type TWVMouseEventVirtualKeys = COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS

// TWVPointerEventKind
//
//	Pointer event type used by SendPointerInput to convey the type of pointer event being sent to WebView. The values of this enum align with the matching WM_POINTER* window messages.
//	Renamed COREWEBVIEW2_POINTER_EVENT_KIND type.
type TWVPointerEventKind = COREWEBVIEW2_POINTER_EVENT_KIND

// TWVBoundsMode
//
//	Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
//	Renamed COREWEBVIEW2_BOUNDS_MODE type.
type TWVBoundsMode = COREWEBVIEW2_BOUNDS_MODE

// TWVProcessFailedReason
//
//	Specifies the process failure reason used in the ICoreWebView2ProcessFailedEventHandler interface.
//	Renamed COREWEBVIEW2_PROCESS_FAILED_REASON type.
type TWVProcessFailedReason = COREWEBVIEW2_PROCESS_FAILED_REASON

// TWVPrintOrientation
//
//	The orientation for printing, used by the Orientation property on ICoreWebView2PrintSettings.
//	Renamed COREWEBVIEW2_PRINT_ORIENTATION type.
type TWVPrintOrientation = COREWEBVIEW2_PRINT_ORIENTATION

// TWVColor
//
//	A value representing RGBA color (Red, Green, Blue, Alpha) for WebView2. Each component takes a value from 0 to 255, with 0 being no intensity and 255 being the highest intensity.
//	Renamed COREWEBVIEW2_COLOR type.
type TWVColor = COREWEBVIEW2_COLOR

// TWVDefaultDownloadDialogCornerAlignment
//
//	The default download dialog can be aligned to any of the WebView corners by setting the DefaultDownloadDialogCornerAlignment property. The default position is top-right corner.
//	Renamed COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT type.
type TWVDefaultDownloadDialogCornerAlignment = COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT

// TWVProcessKind
//
//	Indicates the process type used in the ICoreWebView2ProcessInfo interface.
//	Renamed COREWEBVIEW2_PROCESS_KIND type.
type TWVProcessKind = COREWEBVIEW2_PROCESS_KIND

// TWVMenuItemKind
//
//	Specifies the menu item kind for the ICoreWebView2ContextMenuItem.get_Kind method
//	Renamed COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND type.
type TWVMenuItemKind = COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND

// TWVMenuTargetKind
//
//	Indicates the kind of context for which the context menu was created for the `ICoreWebView2ContextMenuTarget::get_Kind` method. This enum will always represent the active element that caused the context menu request. If there is a selection with multiple images, audio and text, for example, the element that the end user right clicks on within this selection will be the option represented by this enum.
//	Renamed COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND type.
type TWVMenuTargetKind = COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND

// TWVPDFToolbarItems
//
//	PDF toolbar item. This enum must be in sync with ToolBarItem in pdf-store-data-types.ts Specifies the PDF toolbar item types used for the ICoreWebView2Settings.put_HiddenPdfToolbarItems method.
//	Renamed COREWEBVIEW2_PDF_TOOLBAR_ITEMS type.
type TWVPDFToolbarItems = COREWEBVIEW2_PDF_TOOLBAR_ITEMS

// TWVPreferredColorScheme
//
//	An enum to represent the options for WebView2 color scheme: auto, light, or dark.
//	Renamed COREWEBVIEW2_PREFERRED_COLOR_SCHEME type.
type TWVPreferredColorScheme = COREWEBVIEW2_PREFERRED_COLOR_SCHEME

// TWVBrowsingDataKinds
//
//	Specifies the datatype for the ICoreWebView2Profile2.ClearBrowsingData method.
//	Renamed COREWEBVIEW2_BROWSING_DATA_KINDS type.
type TWVBrowsingDataKinds = COREWEBVIEW2_BROWSING_DATA_KINDS

// TWVServerCertificateErrorAction
//
//	Specifies the action type when server certificate error is detected to be used in the ICoreWebView2ServerCertificateErrorDetectedEventArgs interface.
//	Renamed COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION type.
type TWVServerCertificateErrorAction = COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION

// TWVFaviconImageFormat
//
//	Specifies the image format to use for favicon.
//	Renamed COREWEBVIEW2_FAVICON_IMAGE_FORMAT type.
type TWVFaviconImageFormat = COREWEBVIEW2_FAVICON_IMAGE_FORMAT

// TWVPrintCollation
//
//	Specifies the collation for a print.
//	Renamed COREWEBVIEW2_PRINT_COLLATION type.
type TWVPrintCollation = COREWEBVIEW2_PRINT_COLLATION

// TWVPrintColorMode
//
//	Specifies the color mode for a print.
//	Renamed COREWEBVIEW2_PRINT_COLOR_MODE type.
type TWVPrintColorMode = COREWEBVIEW2_PRINT_COLOR_MODE

// TWVPrintDuplex
//
//	Specifies the duplex option for a print.
//	Renamed COREWEBVIEW2_PRINT_DUPLEX type.
type TWVPrintDuplex = COREWEBVIEW2_PRINT_DUPLEX

// TWVPrintMediaSize
//
//	Specifies the media size for a print.
//	Renamed COREWEBVIEW2_PRINT_MEDIA_SIZE type.
type TWVPrintMediaSize = COREWEBVIEW2_PRINT_MEDIA_SIZE

// TWVPrintStatus
//
//	Indicates the status for printing.
//	Renamed COREWEBVIEW2_PRINT_STATUS type.
type TWVPrintStatus = COREWEBVIEW2_PRINT_STATUS

// TWVPrintDialogKind
//
//	Specifies the print dialog kind.
//	Renamed COREWEBVIEW2_PRINT_DIALOG_KIND type.
type TWVPrintDialogKind = COREWEBVIEW2_PRINT_DIALOG_KIND

// TWVSharedBufferAccess
//
//	Specifies the desired access from script to CoreWebView2SharedBuffer.
//	Renamed COREWEBVIEW2_SHARED_BUFFER_ACCESS type.
type TWVSharedBufferAccess = COREWEBVIEW2_SHARED_BUFFER_ACCESS

// TWVTrackingPreventionLevel
//
//	Tracking prevention levels.
//	Renamed COREWEBVIEW2_TRACKING_PREVENTION_LEVEL type.
type TWVTrackingPreventionLevel = COREWEBVIEW2_TRACKING_PREVENTION_LEVEL

// TWVMemoryUsageTargetLevel
//
//	Specifies memory usage target level of WebView.
//	Renamed COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL type.
type TWVMemoryUsageTargetLevel = COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL

// TWVNavigationKind
//
//	Specifies the navigation kind of each navigation.
//	Renamed COREWEBVIEW2_NAVIGATION_KIND type.
type TWVNavigationKind = COREWEBVIEW2_NAVIGATION_KIND

// TWVFrameKind
//
//	Indicates the frame type used in the `ICoreWebView2FrameInfo` interface.
//	Renamed COREWEBVIEW2_FRAME_KIND type.
type TWVFrameKind = COREWEBVIEW2_FRAME_KIND

// TWV2LoaderStatus
//
//	TWVLoader status values
type TWV2LoaderStatus = int32

const (
	WVLSCREATEDvlsCreated TWV2LoaderStatus = iota
	WVLSLOADINGvlsLoading
	WVLSLOADEDvlsLoaded
	WVLSIMPORTEDvlsImported
	WVLSINITIALIZEDvlsInitialized
	WVLSERRORvlsError
	WVLSUNLOADEDvlsUnloaded
)

// TWV2KeyEventType
//
//	Event type used by TWVBrowserBase.SimulateKeyEvent
type TWV2KeyEventType = int32

const (
	KETKEYDOWNetKeyDown TWV2KeyEventType = iota
	KETKEYUPetKeyUp
	KETRAWKEYDOWNetRawKeyDown
	KETCHARetChar
)

// TWV2DebugLog
//
//	Debug log values used by TWVLoader.DebugLog
type TWV2DebugLog = int32

const (
	DLDISABLEDlDisabled TWV2DebugLog = iota
	DLENABLEDlEnabled
	DLENABLEDSTDOUTlEnabledStdOut
	DLENABLEDSTDERRlEnabledStdErr
)

// TWV2DebugLogLevel
//
//	Debug log level used when the logging is enabled
type TWV2DebugLogLevel = int32

const (
	DLLDEFAULTllDefault TWV2DebugLogLevel = iota
	DLLINFOllInfo
	DLLWARNINGllWarning
	DLLERRORllError
	DLLFATALllFatal
)

// TWV2EditingCommand
//
//	Blink editing commands used by the "Input.dispatchKeyEvent" DevTools method.
//	<see href="https://chromedevtools.github.io/devtools-protocol/1-3/Input/#method-dispatchKeyEvent">See the "Input.dispatchKeyEvent" DevTools method.) <see href="https://source.chromium.org/chromium/chromium/src/+/master:third_party/blink/renderer/core/editing/commands/editor_command_names.h">See the Chromium sources.)
type TWV2EditingCommand = int32

const (
	ECALIGNCENTERcAlignCenter TWV2EditingCommand = iota
	ECALIGNJUSTIFIEDcAlignJustified
	ECALIGNLEFTcAlignLeft
	ECALIGNRIGHTcAlignRight
	ECBACKCOLORcBackColor
	ECBACKWARDDELETEcBackwardDelete
	ECBOLDcBold
	ECCOPYcCopy
	ECCREATELINKcCreateLink
	ECCUTcCut
	ECDEFAULTPARAGRAPHSEPARATORcDefaultParagraphSeparator
	ECDELETEcDelete
	ECDELETEBACKWARDcDeleteBackward
	ECDELETEBACKWARDBYDECOMPOSINGPREVIOUSCHARACTERcDeleteBackwardByDecomposingPreviousCharacter
	ECDELETEFORWARDcDeleteForward
	ECDELETETOBEGINNINGOFLINEcDeleteToBeginningOfLine
	ECDELETETOBEGINNINGOFPARAGRAPHcDeleteToBeginningOfParagraph
	ECDELETETOENDOFLINEcDeleteToEndOfLine
	ECDELETETOENDOFPARAGRAPHcDeleteToEndOfParagraph
	ECDELETETOMARKcDeleteToMark
	ECDELETEWORDBACKWARDcDeleteWordBackward
	ECDELETEWORDFORWARDcDeleteWordForward
	ECFINDSTRINGcFindString
	ECFONTNAMEcFontName
	ECFONTSIZEcFontSize
	ECFONTSIZEDELTAcFontSizeDelta
	ECFORECOLORcForeColor
	ECFORMATBLOCKcFormatBlock
	ECFORWARDDELETEcForwardDelete
	ECHILITECOLORcHiliteColor
	ECIGNORESPELLINGcIgnoreSpelling
	ECINDENTcIndent
	ECINSERTBACKTABcInsertBacktab
	ECINSERTHORIZONTALRULEcInsertHorizontalRule
	ECINSERTHTMLcInsertHTML
	ECINSERTIMAGEcInsertImage
	ECINSERTLINEBREAKcInsertLineBreak
	ECINSERTNEWLINEcInsertNewline
	ECINSERTNEWLINEINQUOTEDCONTENTcInsertNewlineInQuotedContent
	ECINSERTORDEREDLISTcInsertOrderedList
	ECINSERTPARAGRAPHcInsertParagraph
	ECINSERTTABcInsertTab
	ECINSERTTEXTcInsertText
	ECINSERTUNORDEREDLISTcInsertUnorderedList
	ECITALICcItalic
	ECJUSTIFYCENTERcJustifyCenter
	ECJUSTIFYFULLcJustifyFull
	ECJUSTIFYLEFTcJustifyLeft
	ECJUSTIFYNONEcJustifyNone
	ECJUSTIFYRIGHTcJustifyRight
	ECMAKETEXTWRITINGDIRECTIONLEFTTORIGHTcMakeTextWritingDirectionLeftToRight
	ECMAKETEXTWRITINGDIRECTIONNATURALcMakeTextWritingDirectionNatural
	ECMAKETEXTWRITINGDIRECTIONRIGHTTOLEFTcMakeTextWritingDirectionRightToLeft
	ECMOVEBACKWARDcMoveBackward
	ECMOVEBACKWARDANDMODIFYSELECTIONcMoveBackwardAndModifySelection
	ECMOVEDOWNcMoveDown
	ECMOVEDOWNANDMODIFYSELECTIONcMoveDownAndModifySelection
	ECMOVEFORWARDcMoveForward
	ECMOVEFORWARDANDMODIFYSELECTIONcMoveForwardAndModifySelection
	ECMOVELEFTcMoveLeft
	ECMOVELEFTANDMODIFYSELECTIONcMoveLeftAndModifySelection
	ECMOVEPAGEDOWNcMovePageDown
	ECMOVEPAGEDOWNANDMODIFYSELECTIONcMovePageDownAndModifySelection
	ECMOVEPAGEUPcMovePageUp
	ECMOVEPAGEUPANDMODIFYSELECTIONcMovePageUpAndModifySelection
	ECMOVEPARAGRAPHBACKWARDcMoveParagraphBackward
	ECMOVEPARAGRAPHBACKWARDANDMODIFYSELECTIONcMoveParagraphBackwardAndModifySelection
	ECMOVEPARAGRAPHFORWARDcMoveParagraphForward
	ECMOVEPARAGRAPHFORWARDANDMODIFYSELECTIONcMoveParagraphForwardAndModifySelection
	ECMOVERIGHTcMoveRight
	ECMOVERIGHTANDMODIFYSELECTIONcMoveRightAndModifySelection
	ECMOVETOBEGINNINGOFDOCUMENTcMoveToBeginningOfDocument
	ECMOVETOBEGINNINGOFDOCUMENTANDMODIFYSELECTIONcMoveToBeginningOfDocumentAndModifySelection
	ECMOVETOBEGINNINGOFLINEcMoveToBeginningOfLine
	ECMOVETOBEGINNINGOFLINEANDMODIFYSELECTIONcMoveToBeginningOfLineAndModifySelection
	ECMOVETOBEGINNINGOFPARAGRAPHcMoveToBeginningOfParagraph
	ECMOVETOBEGINNINGOFPARAGRAPHANDMODIFYSELECTIONcMoveToBeginningOfParagraphAndModifySelection
	ECMOVETOBEGINNINGOFSENTENCEcMoveToBeginningOfSentence
	ECMOVETOBEGINNINGOFSENTENCEANDMODIFYSELECTIONcMoveToBeginningOfSentenceAndModifySelection
	ECMOVETOENDOFDOCUMENTcMoveToEndOfDocument
	ECMOVETOENDOFDOCUMENTANDMODIFYSELECTIONcMoveToEndOfDocumentAndModifySelection
	ECMOVETOENDOFLINEcMoveToEndOfLine
	ECMOVETOENDOFLINEANDMODIFYSELECTIONcMoveToEndOfLineAndModifySelection
	ECMOVETOENDOFPARAGRAPHcMoveToEndOfParagraph
	ECMOVETOENDOFPARAGRAPHANDMODIFYSELECTIONcMoveToEndOfParagraphAndModifySelection
	ECMOVETOENDOFSENTENCEcMoveToEndOfSentence
	ECMOVETOENDOFSENTENCEANDMODIFYSELECTIONcMoveToEndOfSentenceAndModifySelection
	ECMOVETOLEFTENDOFLINEcMoveToLeftEndOfLine
	ECMOVETOLEFTENDOFLINEANDMODIFYSELECTIONcMoveToLeftEndOfLineAndModifySelection
	ECMOVETORIGHTENDOFLINEcMoveToRightEndOfLine
	ECMOVETORIGHTENDOFLINEANDMODIFYSELECTIONcMoveToRightEndOfLineAndModifySelection
	ECMOVEUPcMoveUp
	ECMOVEUPANDMODIFYSELECTIONcMoveUpAndModifySelection
	ECMOVEWORDBACKWARDcMoveWordBackward
	ECMOVEWORDBACKWARDANDMODIFYSELECTIONcMoveWordBackwardAndModifySelection
	ECMOVEWORDFORWARDcMoveWordForward
	ECMOVEWORDFORWARDANDMODIFYSELECTIONcMoveWordForwardAndModifySelection
	ECMOVEWORDLEFTcMoveWordLeft
	ECMOVEWORDLEFTANDMODIFYSELECTIONcMoveWordLeftAndModifySelection
	ECMOVEWORDRIGHTcMoveWordRight
	ECMOVEWORDRIGHTANDMODIFYSELECTIONcMoveWordRightAndModifySelection
	ECOUTDENTcOutdent
	ECOVERWRITEcOverWrite
	ECPASTEcPaste
	ECPASTEANDMATCHSTYLEcPasteAndMatchStyle
	ECPASTEGLOBALSELECTIONcPasteGlobalSelection
	ECPRINTcPrint
	ECREDOcRedo
	ECREMOVEFORMATcRemoveFormat
	ECSCROLLLINEDOWNcScrollLineDown
	ECSCROLLLINEUPcScrollLineUp
	ECSCROLLPAGEBACKWARDcScrollPageBackward
	ECSCROLLPAGEFORWARDcScrollPageForward
	ECSCROLLTOBEGINNINGOFDOCUMENTcScrollToBeginningOfDocument
	ECSCROLLTOENDOFDOCUMENTcScrollToEndOfDocument
	ECSELECTALLcSelectAll
	ECSELECTLINEcSelectLine
	ECSELECTPARAGRAPHcSelectParagraph
	ECSELECTSENTENCEcSelectSentence
	ECSELECTTOMARKcSelectToMark
	ECSELECTWORDcSelectWord
	ECSETMARKcSetMark
	ECSTRIKETHROUGHcStrikethrough
	ECSTYLEWITHCSScStyleWithCSS
	ECSUBSCRIPTcSubscript
	ECSUPERSCRIPTcSuperscript
	ECSWAPWITHMARKcSwapWithMark
	ECTOGGLEBOLDcToggleBold
	ECTOGGLEITALICcToggleItalic
	ECTOGGLEUNDERLINEcToggleUnderline
	ECTRANSPOSEcTranspose
	ECUNDERLINEcUnderline
	ECUNDOcUndo
	ECUNLINKcUnlink
	ECUNSCRIPTcUnscript
	ECUNSELECTcUnselect
	ECUSECSScUseCSS
	ECYANKcYank
	ECYANKANDSELECTcYankAndSelect
)

// TWVClearDataStorageTypes
//
//	Used by TWVBrowserBase.ClearDataForOrigin to clear the storage
type TWVClearDataStorageTypes = int32

const (
	CDSTAPPCACHEdstAppCache TWVClearDataStorageTypes = iota
	CDSTCOOKIESdstCookies
	CDSTFILESYSTEMSdstFileSystems
	CDSTINDEXEDDBdstIndexeddb
	CDSTLOCALSTORAGEdstLocalStorage
	CDSTSHADERCACHEdstShaderCache
	CDSTWEBSQLdstWebsql
	CDSTSERVICEWORKERSdstServiceWorkers
	CDSTCACHESTORAGEdstCacheStorage
	CDSTALLdstAll
)

// TWVState
//
//	Represents the state of a setting.
type TWVState = int32

const (
	STATE_DEFAULTTATE_DEFAULT TWVState = iota
	STATE_ENABLEDTATE_ENABLED
	STATE_DISABLEDTATE_DISABLED
)

// TWVAutoplayPolicy
//
//	Autoplay policy types used by TWVLoader.AutoplayPolicy. See the --autoplay-policy switch.
type TWVAutoplayPolicy = int32

const (
	APPDEFAULTppDefault TWVAutoplayPolicy = iota
	APPDOCUMENTUSERACTIVATIONREQUIREDppDocumentUserActivationRequired
	APPNOUSERGESTUREREQUIREDppNoUserGestureRequired
	APPUSERGESTUREREQUIREDppUserGestureRequired
)

// TWVCustomSchemeInfoArray
type TWVCustomSchemeInfoArray = []*TWVCustomSchemeInfo

// COREWEBVIEW2_KEY_EVENT_KIND
//
//	Specifies the key event type that triggered an AcceleratorKeyPressed event.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_key_event_kind">See the Globals article.)
type COREWEBVIEW2_KEY_EVENT_KIND = TOleEnum

const (
	// COREWEBVIEW2_KEY_EVENT_KIND_KEY_DOWN
	//
	//	Specifies that the key event type corresponds to window message WM_KEYDOWN.
	//	This is one of the COREWEBVIEW2_KEY_EVENT_KIND values.
	COREWEBVIEW2_KEY_EVENT_KIND_KEY_DOWN COREWEBVIEW2_KEY_EVENT_KIND = 0x0000000
	// COREWEBVIEW2_KEY_EVENT_KIND_KEY_UP
	//
	//	Specifies that the key event type corresponds to window message WM_KEYUP.
	//	This is one of the COREWEBVIEW2_KEY_EVENT_KIND values.
	COREWEBVIEW2_KEY_EVENT_KIND_KEY_UP COREWEBVIEW2_KEY_EVENT_KIND = 0x0000001
	// COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_DOWN
	//
	//	Specifies that the key event type corresponds to window message WM_SYSKEYDOWN.
	//	This is one of the COREWEBVIEW2_KEY_EVENT_KIND values.
	COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_DOWN COREWEBVIEW2_KEY_EVENT_KIND = 0x0000002
	// COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_UP
	//
	//	Specifies that the key event type corresponds to window message WM_SYSKEYUP.
	//	This is one of the COREWEBVIEW2_KEY_EVENT_KIND values.
	COREWEBVIEW2_KEY_EVENT_KIND_SYSTEM_KEY_UP COREWEBVIEW2_KEY_EVENT_KIND = 0x0000003
)

// COREWEBVIEW2_MOVE_FOCUS_REASON
//
//	Specifies the reason for moving focus.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_move_focus_reason">See the Globals article.)
type COREWEBVIEW2_MOVE_FOCUS_REASON = TOleEnum

const (
	// COREWEBVIEW2_MOVE_FOCUS_REASON_PROGRAMMATIC
	//
	//	This is one of the COREWEBVIEW2_MOVE_FOCUS_REASON values.
	//	Specifies that the code is setting focus into WebView.
	COREWEBVIEW2_MOVE_FOCUS_REASON_PROGRAMMATIC COREWEBVIEW2_MOVE_FOCUS_REASON = 0x0000000
	// COREWEBVIEW2_MOVE_FOCUS_REASON_NEXT
	//
	//	Specifies that the focus is moving due to Tab traversal forward.
	//	This is one of the COREWEBVIEW2_MOVE_FOCUS_REASON values.
	COREWEBVIEW2_MOVE_FOCUS_REASON_NEXT COREWEBVIEW2_MOVE_FOCUS_REASON = 0x0000001
	// COREWEBVIEW2_MOVE_FOCUS_REASON_PREVIOUS
	//
	//	Specifies that the focus is moving due to Tab traversal backward.
	//	This is one of the COREWEBVIEW2_MOVE_FOCUS_REASON values.
	COREWEBVIEW2_MOVE_FOCUS_REASON_PREVIOUS COREWEBVIEW2_MOVE_FOCUS_REASON = 0x0000002
)

// COREWEBVIEW2_WEB_ERROR_STATUS
//
//	Indicates the error status values for web navigations.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_web_error_status">See the Globals article.)
type COREWEBVIEW2_WEB_ERROR_STATUS = TOleEnum

const (
	// COREWEBVIEW2_WEB_ERROR_STATUS_UNKNOWN
	//
	//	Indicates that an unknown error occurred.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_UNKNOWN COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000000
	// COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_COMMON_NAME_IS_INCORRECT
	//
	//	Indicates that the SSL certificate common name does not match the web address.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_COMMON_NAME_IS_INCORRECT COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000001
	// COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_EXPIRED
	//
	//	Indicates that the SSL certificate has expired.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_EXPIRED COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000002
	// COREWEBVIEW2_WEB_ERROR_STATUS_CLIENT_CERTIFICATE_CONTAINS_ERRORS
	//
	//	Indicates that the SSL client certificate contains errors.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_CLIENT_CERTIFICATE_CONTAINS_ERRORS COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000003
	// COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_REVOKED
	//
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	//	Indicates that the SSL certificate has been revoked.
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_REVOKED COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000004
	// COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_IS_INVALID
	//
	//	Indicates that the SSL certificate is not valid. The certificate may not match the public key pins for the host name, the certificate is signed by an untrusted authority or using a weak sign algorithm, the certificate claimed DNS names violate name constraints, the certificate contains a weak key, the validity period of the certificate is too long, lack of revocation information or revocation mechanism, non-unique host name, lack of certificate transparency information, or the certificate is chained to a [legacy Symantec root](https://security.googleblog.com/2018/03/distrust-of-symantec-pki-immediate.html).
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_CERTIFICATE_IS_INVALID COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000005
	// COREWEBVIEW2_WEB_ERROR_STATUS_SERVER_UNREACHABLE
	//
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	//	Indicates that the host is unreachable.
	COREWEBVIEW2_WEB_ERROR_STATUS_SERVER_UNREACHABLE COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000006
	// COREWEBVIEW2_WEB_ERROR_STATUS_TIMEOUT
	//
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	//	Indicates that the connection has timed out.
	COREWEBVIEW2_WEB_ERROR_STATUS_TIMEOUT COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000007
	// COREWEBVIEW2_WEB_ERROR_STATUS_ERROR_HTTP_INVALID_SERVER_RESPONSE
	//
	//	Indicates that the server returned an invalid or unrecognized response.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_ERROR_HTTP_INVALID_SERVER_RESPONSE COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000008
	// COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_ABORTED
	//
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	//	Indicates that the connection was stopped.
	COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_ABORTED COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000009
	// COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_RESET
	//
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	//	Indicates that the connection was reset.
	COREWEBVIEW2_WEB_ERROR_STATUS_CONNECTION_RESET COREWEBVIEW2_WEB_ERROR_STATUS = 0x000000A
	// COREWEBVIEW2_WEB_ERROR_STATUS_DISCONNECTED
	//
	//	Indicates that the Internet connection has been lost.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_DISCONNECTED COREWEBVIEW2_WEB_ERROR_STATUS = 0x000000B
	// COREWEBVIEW2_WEB_ERROR_STATUS_CANNOT_CONNECT
	//
	//	Indicates that a connection to the destination was not established.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_CANNOT_CONNECT COREWEBVIEW2_WEB_ERROR_STATUS = 0x000000C
	// COREWEBVIEW2_WEB_ERROR_STATUS_HOST_NAME_NOT_RESOLVED
	//
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	//	Indicates that the provided host name was not able to be resolved.
	COREWEBVIEW2_WEB_ERROR_STATUS_HOST_NAME_NOT_RESOLVED COREWEBVIEW2_WEB_ERROR_STATUS = 0x000000D
	// COREWEBVIEW2_WEB_ERROR_STATUS_OPERATION_CANCELED
	//
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	//	Indicates that the operation was canceled. This status code is also used in the following cases: - When the app cancels a navigation via NavigationStarting event. - For original navigation if the app navigates the WebView2 in a rapid succession away after the load for original navigation commenced, but before it completed.
	COREWEBVIEW2_WEB_ERROR_STATUS_OPERATION_CANCELED COREWEBVIEW2_WEB_ERROR_STATUS = 0x000000E
	// COREWEBVIEW2_WEB_ERROR_STATUS_REDIRECT_FAILED
	//
	//	Indicates that the request redirect failed.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_REDIRECT_FAILED COREWEBVIEW2_WEB_ERROR_STATUS = 0x000000F
	// COREWEBVIEW2_WEB_ERROR_STATUS_UNEXPECTED_ERROR
	//
	//	Indicates that an unexpected error occurred.
	COREWEBVIEW2_WEB_ERROR_STATUS_UNEXPECTED_ERROR COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000010
	// COREWEBVIEW2_WEB_ERROR_STATUS_VALID_AUTHENTICATION_CREDENTIALS_REQUIRED
	//
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	//	Indicates that user is prompted with a login, waiting on user action. Initial navigation to a login site will always return this even if app provides credential using BasicAuthenticationRequested. HTTP response status code in this case is 401. See status code reference here: https://developer.mozilla.org/docs/Web/HTTP/Status.
	COREWEBVIEW2_WEB_ERROR_STATUS_VALID_AUTHENTICATION_CREDENTIALS_REQUIRED COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000011
	// COREWEBVIEW2_WEB_ERROR_STATUS_VALID_PROXY_AUTHENTICATION_REQUIRED
	//
	//	Indicates that user lacks proper authentication credentials for a proxy server. HTTP response status code in this case is 407. See status code reference here: https://developer.mozilla.org/docs/Web/HTTP/Status.
	//	This is one of the COREWEBVIEW2_WEB_ERROR_STATUS values.
	COREWEBVIEW2_WEB_ERROR_STATUS_VALID_PROXY_AUTHENTICATION_REQUIRED COREWEBVIEW2_WEB_ERROR_STATUS = 0x0000012
)

// COREWEBVIEW2_SCRIPT_DIALOG_KIND
//
//	Specifies the JavaScript dialog type used in the ICoreWebView2ScriptDialogOpeningEventHandler interface.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_script_dialog_kind">See the Globals article.)
type COREWEBVIEW2_SCRIPT_DIALOG_KIND = TOleEnum

const (
	// COREWEBVIEW2_SCRIPT_DIALOG_KIND_ALERT
	//
	//	Indicates that the dialog uses the window.alert JavaScript function.
	//	This is one of the COREWEBVIEW2_SCRIPT_DIALOG_KIND values.
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_ALERT COREWEBVIEW2_SCRIPT_DIALOG_KIND = 0x0000000
	// COREWEBVIEW2_SCRIPT_DIALOG_KIND_CONFIRM
	//
	//	Indicates that the dialog uses the window.confirm JavaScript function.
	//	This is one of the COREWEBVIEW2_SCRIPT_DIALOG_KIND values.
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_CONFIRM COREWEBVIEW2_SCRIPT_DIALOG_KIND = 0x0000001
	// COREWEBVIEW2_SCRIPT_DIALOG_KIND_PROMPT
	//
	//	Indicates that the dialog uses the window.prompt JavaScript function.
	//	This is one of the COREWEBVIEW2_SCRIPT_DIALOG_KIND values.
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_PROMPT COREWEBVIEW2_SCRIPT_DIALOG_KIND = 0x0000002
	// COREWEBVIEW2_SCRIPT_DIALOG_KIND_BEFOREUNLOAD
	//
	//	Indicates that the dialog uses the beforeunload JavaScript event.
	//	This is one of the COREWEBVIEW2_SCRIPT_DIALOG_KIND values.
	COREWEBVIEW2_SCRIPT_DIALOG_KIND_BEFOREUNLOAD COREWEBVIEW2_SCRIPT_DIALOG_KIND = 0x0000003
)

// COREWEBVIEW2_PERMISSION_KIND
//
//	Indicates the type of a permission request.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_permission_kind">See the Globals article.)
type COREWEBVIEW2_PERMISSION_KIND = TOleEnum

const (
	// COREWEBVIEW2_PERMISSION_KIND_UNKNOWN_PERMISSION
	//
	//	Indicates an unknown permission.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_UNKNOWN_PERMISSION COREWEBVIEW2_PERMISSION_KIND = 0x0000000
	// COREWEBVIEW2_PERMISSION_KIND_MICROPHONE
	//
	//	Indicates permission to capture audio.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_MICROPHONE COREWEBVIEW2_PERMISSION_KIND = 0x0000001
	// COREWEBVIEW2_PERMISSION_KIND_CAMERA
	//
	//	Indicates permission to capture video.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_CAMERA COREWEBVIEW2_PERMISSION_KIND = 0x0000002
	// COREWEBVIEW2_PERMISSION_KIND_GEOLOCATION
	//
	//	Indicates permission to access geolocation.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_GEOLOCATION COREWEBVIEW2_PERMISSION_KIND = 0x0000003
	// COREWEBVIEW2_PERMISSION_KIND_NOTIFICATIONS
	//
	//	Indicates permission to send web notifications. Apps that would like to show notifications should handle PermissionRequested events and no browser permission prompt will be shown for notification requests. Note that push notifications are currently unavailable in WebView2.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_NOTIFICATIONS COREWEBVIEW2_PERMISSION_KIND = 0x0000004
	// COREWEBVIEW2_PERMISSION_KIND_OTHER_SENSORS
	//
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	//	Indicates permission to access generic sensor. Generic Sensor covering ambient-light-sensor, accelerometer, gyroscope, and magnetometer.
	COREWEBVIEW2_PERMISSION_KIND_OTHER_SENSORS COREWEBVIEW2_PERMISSION_KIND = 0x0000005
	// COREWEBVIEW2_PERMISSION_KIND_CLIPBOARD_READ
	//
	//	Indicates permission to read the system clipboard without a user gesture.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_CLIPBOARD_READ COREWEBVIEW2_PERMISSION_KIND = 0x0000006
	// COREWEBVIEW2_PERMISSION_KIND_MULTIPLE_AUTOMATIC_DOWNLOADS
	//
	//	Indicates permission to automatically download multiple files. Permission is requested when multiple downloads are triggered in quick succession.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_MULTIPLE_AUTOMATIC_DOWNLOADS COREWEBVIEW2_PERMISSION_KIND = 0x0000007
	// COREWEBVIEW2_PERMISSION_KIND_FILE_READ_WRITE
	//
	//	Indicates permission to read and write to files or folders on the device. Permission is requested when developers use the [File System Access API](https://developer.mozilla.org/en-US/docs/Web/API/File_System_Access_API) to show the file or folder picker to the end user, and then request "readwrite" permission for the user's selection.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_FILE_READ_WRITE COREWEBVIEW2_PERMISSION_KIND = 0x0000008
	// COREWEBVIEW2_PERMISSION_KIND_AUTOPLAY
	//
	//	Indicates permission to play audio and video automatically on sites. This permission affects the autoplay attribute and play method of the audio and video HTML elements, and the start method of the Web Audio API. See the [Autoplay guide for media and Web Audio APIs](https://developer.mozilla.org/en-US/docs/Web/Media/Autoplay_guide) for details.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_AUTOPLAY COREWEBVIEW2_PERMISSION_KIND = 0x0000009
	// COREWEBVIEW2_PERMISSION_KIND_LOCAL_FONTS
	//
	//	Indicates permission to use fonts on the device. Permission is requested when developers use the [Local Font Access API](https://wicg.github.io/local-font-access/) to query the system fonts available for styling web content.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_LOCAL_FONTS COREWEBVIEW2_PERMISSION_KIND = 0x000000A
	// COREWEBVIEW2_PERMISSION_KIND_MIDI_SYSTEM_EXCLUSIVE_MESSAGES
	//
	//	Indicates permission to send and receive system exclusive messages to/from MIDI (Musical Instrument Digital Interface) devices. Permission is requested when developers use the [Web MIDI API](https://developer.mozilla.org/en-US/docs/Web/API/Web_MIDI_API) to request access to system exclusive MIDI messages.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_MIDI_SYSTEM_EXCLUSIVE_MESSAGES COREWEBVIEW2_PERMISSION_KIND = 0x000000B
	// COREWEBVIEW2_PERMISSION_KIND_WINDOW_MANAGEMENT
	//
	//	Indicates permission to open and place windows on the screen. Permission is requested when developers use the [Multi-Screen Window Placement API](https://www.w3.org/TR/window-placement/) to get screen details.
	//	This is one of the COREWEBVIEW2_PERMISSION_KIND values.
	COREWEBVIEW2_PERMISSION_KIND_WINDOW_MANAGEMENT COREWEBVIEW2_PERMISSION_KIND = 0x000000C
)

// COREWEBVIEW2_PERMISSION_STATE
//
//	Specifies the response to a permission request.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_permission_state">See the Globals article.)
type COREWEBVIEW2_PERMISSION_STATE = TOleEnum

const (
	// COREWEBVIEW2_PERMISSION_STATE_DEFAULT
	//
	//	Specifies that the default browser behavior is used, which normally prompt users for decision.
	//	This is one of the COREWEBVIEW2_PERMISSION_STATE values.
	COREWEBVIEW2_PERMISSION_STATE_DEFAULT COREWEBVIEW2_PERMISSION_STATE = 0x0000000
	// COREWEBVIEW2_PERMISSION_STATE_ALLOW
	//
	//	Specifies that the permission request is granted.
	//	This is one of the COREWEBVIEW2_PERMISSION_STATE values.
	COREWEBVIEW2_PERMISSION_STATE_ALLOW COREWEBVIEW2_PERMISSION_STATE = 0x0000001
	// COREWEBVIEW2_PERMISSION_STATE_DENY
	//
	//	Specifies that the permission request is denied.
	//	This is one of the COREWEBVIEW2_PERMISSION_STATE values.
	COREWEBVIEW2_PERMISSION_STATE_DENY COREWEBVIEW2_PERMISSION_STATE = 0x0000002
)

// COREWEBVIEW2_PROCESS_FAILED_KIND
//
//	Specifies the process failure type used in the `ICoreWebView2ProcessFailedEventArgs` interface. The values in this enum make reference to the process kinds in the Chromium architecture. For more information about what these processes are and what they do, see [Browser Architecture - Inside look at modern web browser](https://developers.google.com/web/updates/2018/09/inside-browser-part1).
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_process_failed_kind">See the Globals article.)
type COREWEBVIEW2_PROCESS_FAILED_KIND = TOleEnum

const (
	// COREWEBVIEW2_PROCESS_FAILED_KIND_BROWSER_PROCESS_EXITED
	//
	//	Indicates that the browser process ended unexpectedly. The WebView automatically moves to the Closed state. The app has to recreate a new WebView to recover from this failure.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_BROWSER_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000000
	// COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_EXITED
	//
	//	Indicates that the main frame's render process ended unexpectedly. Any subframes in the WebView will be gone too. A new render process is created automatically and navigated to an error page. You can use the `Reload` method to try to recover from this failure. Alternatively, you can `Close` and recreate the WebView.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000001
	// COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE
	//
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	//	Indicates that the main frame's render process is unresponsive. Renderer process unresponsiveness can happen for the following reasons: * There is a **long-running script** being executed. For example, the web content in your WebView might be performing a synchronous XHR, or have entered an infinite loop. * The **system is busy**. The `ProcessFailed` event will continue to be raised every few seconds until the renderer process has become responsive again. The application can consider taking action if the event keeps being raised. For example, the application might show UI for the user to decide to keep waiting or reload the page, or navigate away.
	COREWEBVIEW2_PROCESS_FAILED_KIND_RENDER_PROCESS_UNRESPONSIVE COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000002
	// COREWEBVIEW2_PROCESS_FAILED_KIND_FRAME_RENDER_PROCESS_EXITED
	//
	//	Indicates that a frame-only render process ended unexpectedly. The process exit does not affect the top-level document, only a subset of the subframes within it. The content in these frames is replaced with an error page in the frame. Your application can communicate with the main frame to recover content in the impacted frames, using `ICoreWebView2ProcessFailedEventArgs2.FrameInfosForFailedProcess` to get information about the impacted frames.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_FRAME_RENDER_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000003
	// COREWEBVIEW2_PROCESS_FAILED_KIND_UTILITY_PROCESS_EXITED
	//
	//	Indicates that a utility process ended unexpectedly. The failed process is recreated automatically. Your application does **not** need to handle recovery for this event, but can use `ICoreWebView2ProcessFailedEventArgs` and `ICoreWebView2ProcessFailedEventArgs2` to collect information about the failure, including `ProcessDescription`.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_UTILITY_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000004
	// COREWEBVIEW2_PROCESS_FAILED_KIND_SANDBOX_HELPER_PROCESS_EXITED
	//
	//	Indicates that a sandbox helper process ended unexpectedly. This failure is not fatal. Your application does **not** need to handle recovery for this event, but can use `ICoreWebView2ProcessFailedEventArgs` and `ICoreWebView2ProcessFailedEventArgs2` to collect information about the failure.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_SANDBOX_HELPER_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000005
	// COREWEBVIEW2_PROCESS_FAILED_KIND_GPU_PROCESS_EXITED
	//
	//	Indicates that the GPU process ended unexpectedly. The failed process is recreated automatically. Your application does **not** need to handle recovery for this event, but can use `ICoreWebView2ProcessFailedEventArgs` and `ICoreWebView2ProcessFailedEventArgs2` to collect information about the failure.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_GPU_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000006
	// COREWEBVIEW2_PROCESS_FAILED_KIND_PPAPI_PLUGIN_PROCESS_EXITED
	//
	//	Indicates that a PPAPI plugin process ended unexpectedly. This failure is not fatal. Your application does **not** need to handle recovery for this event, but can use `ICoreWebView2ProcessFailedEventArgs` and `ICoreWebView2ProcessFailedEventArgs2` to collect information about the failure, including `ProcessDescription`.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_PPAPI_PLUGIN_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000007
	// COREWEBVIEW2_PROCESS_FAILED_KIND_PPAPI_BROKER_PROCESS_EXITED
	//
	//	Indicates that a PPAPI plugin broker process ended unexpectedly. This failure is not fatal. Your application does **not** need to handle recovery for this event, but can use `ICoreWebView2ProcessFailedEventArgs` and `ICoreWebView2ProcessFailedEventArgs2` to collect information about the failure.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_PPAPI_BROKER_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000008
	// COREWEBVIEW2_PROCESS_FAILED_KIND_UNKNOWN_PROCESS_EXITED
	//
	//	Indicates that a process of unspecified kind ended unexpectedly. Your application can use `ICoreWebView2ProcessFailedEventArgs` and `ICoreWebView2ProcessFailedEventArgs2` to collect information about the failure.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_KIND values.
	COREWEBVIEW2_PROCESS_FAILED_KIND_UNKNOWN_PROCESS_EXITED COREWEBVIEW2_PROCESS_FAILED_KIND = 0x0000009
)

// COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT
//
//	Specifies the image format for the ICoreWebView2.CapturePreview method.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_capture_preview_image_format">See the Globals article.)
type COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT = TOleEnum

const (
	// COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG
	//
	//	Indicates that the PNG image format is used.
	//	This is one of the COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT values.
	COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_PNG COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT = 0x0000000
	// COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_JPEG
	//
	//	Indicates the JPEG image format is used.
	//	This is one of the COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT values.
	COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT_JPEG COREWEBVIEW2_CAPTURE_PREVIEW_IMAGE_FORMAT = 0x0000001
)

// COREWEBVIEW2_WEB_RESOURCE_CONTEXT
//
//	Specifies the web resource request contexts.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_web_resource_context">See the Globals article.)
type COREWEBVIEW2_WEB_RESOURCE_CONTEXT = TOleEnum

const (
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL
	//
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	//	Specifies all resources.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_ALL COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000000
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_DOCUMENT
	//
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	//	Specifies a document resource.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_DOCUMENT COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000001
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_STYLESHEET
	//
	//	Specifies a CSS resource.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_STYLESHEET COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000002
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_IMAGE
	//
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	//	Specifies an image resource.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_IMAGE COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000003
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_MEDIA
	//
	//	Specifies another media resource such as a video.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_MEDIA COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000004
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_FONT
	//
	//	Specifies a font resource.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_FONT COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000005
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SCRIPT
	//
	//	Specifies a script resource.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SCRIPT COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000006
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_XML_HTTP_REQUEST
	//
	//	Specifies an XML HTTP request, Fetch and EventSource API communication.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_XML_HTTP_REQUEST COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000007
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_FETCH
	//
	//	Specifies a Fetch API communication.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values. Note that this isn't working. Fetch API requests are fired as a part of COREWEBVIEW2_WEB_RESOURCE_CONTEXT_XML_HTTP_REQUEST.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_FETCH COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000008
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_TEXT_TRACK
	//
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	//	Specifies a TextTrack resource.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_TEXT_TRACK COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000009
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_EVENT_SOURCE
	//
	//	Specifies an EventSource API communication.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_EVENT_SOURCE COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x000000A
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_WEBSOCKET
	//
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values. Note that this isn't working. EventSource API requests are fired as a part of COREWEBVIEW2_WEB_RESOURCE_CONTEXT_XML_HTTP_REQUEST.
	//	Specifies a WebSocket API communication.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_WEBSOCKET COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x000000B
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_MANIFEST
	//
	//	Specifies a Web App Manifest.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_MANIFEST COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x000000C
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SIGNED_EXCHANGE
	//
	//	Specifies a Signed HTTP Exchange.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_SIGNED_EXCHANGE COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x000000D
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_PING
	//
	//	Specifies a Ping request.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_PING COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x000000E
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_CSP_VIOLATION_REPORT
	//
	//	Specifies a CSP Violation Report.
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_CSP_VIOLATION_REPORT COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x000000F
	// COREWEBVIEW2_WEB_RESOURCE_CONTEXT_OTHER
	//
	//	This is one of the COREWEBVIEW2_WEB_RESOURCE_CONTEXT values.
	//	Specifies an other resource.
	COREWEBVIEW2_WEB_RESOURCE_CONTEXT_OTHER COREWEBVIEW2_WEB_RESOURCE_CONTEXT = 0x0000010
)

// COREWEBVIEW2_COOKIE_SAME_SITE_KIND
//
//	Kind of cookie SameSite status used in the ICoreWebView2Cookie interface. These fields match those as specified in https://developer.mozilla.org/docs/Web/HTTP/Cookies#. Learn more about SameSite cookies here: https://tools.ietf.org/html/draft-west-first-party-cookies-07
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_cookie_same_site_kind">See the Globals article.)
type COREWEBVIEW2_COOKIE_SAME_SITE_KIND = TOleEnum

const (
	// COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE
	//
	//	This is one of the COREWEBVIEW2_COOKIE_SAME_SITE_KIND values.
	//	None SameSite type. No restrictions on cross-site requests.
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_NONE COREWEBVIEW2_COOKIE_SAME_SITE_KIND = 0x0000000
	// COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX
	//
	//	Lax SameSite type. The cookie will be sent with "same-site" requests, and with "cross-site" top level navigation.
	//	This is one of the COREWEBVIEW2_COOKIE_SAME_SITE_KIND values.
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_LAX COREWEBVIEW2_COOKIE_SAME_SITE_KIND = 0x0000001
	// COREWEBVIEW2_COOKIE_SAME_SITE_KIND_STRICT
	//
	//	This is one of the COREWEBVIEW2_COOKIE_SAME_SITE_KIND values.
	//	Strict SameSite type. The cookie will only be sent along with "same-site" requests.
	COREWEBVIEW2_COOKIE_SAME_SITE_KIND_STRICT COREWEBVIEW2_COOKIE_SAME_SITE_KIND = 0x0000002
)

// COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND
//
//	Kind of cross origin resource access allowed for host resources during download. Note that other normal access checks like same origin DOM access check and [Content Security Policy](https://developer.mozilla.org/docs/Web/HTTP/CSP) still apply. The following table illustrates the host resource cross origin access according to access context and COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND. Cross Origin Access Context | DENY | ALLOW | DENY_CORS --- | --- | --- | --- From DOM like src of img, script or iframe element| Deny | Allow | Allow From Script like Fetch or XMLHttpRequest| Deny | Allow | Deny
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_host_resource_access_kind">See the Globals article.)
type COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND = TOleEnum

const (
	// COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY
	//
	//	All cross origin resource access is denied, including normal sub resource access as src of a script or image element.
	//	This is one of the COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND values.
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND = 0x0000000
	// COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW
	//
	//	All cross origin resource access is allowed, including accesses that are subject to Cross-Origin Resource Sharing(CORS) check. The behavior is similar to a web site sends back http header Access-Control-Allow-Origin: *.
	//	This is one of the COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND values.
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_ALLOW COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND = 0x0000001
	// COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY_CORS
	//
	//	Cross origin resource access is allowed for normal sub resource access like as src of a script or image element, while any access that subjects to CORS check will be denied. See [Cross-Origin Resource Sharing](https://developer.mozilla.org/docs/Web/HTTP/CORS) for more information.
	//	This is one of the COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND values.
	COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND_DENY_CORS COREWEBVIEW2_HOST_RESOURCE_ACCESS_KIND = 0x0000002
)

// COREWEBVIEW2_DOWNLOAD_STATE
//
//	State of the download operation.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_download_state">See the Globals article.)
type COREWEBVIEW2_DOWNLOAD_STATE = TOleEnum

const (
	// COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS
	//
	//	This is one of the COREWEBVIEW2_DOWNLOAD_STATE values.
	//	The download is in progress.
	COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS COREWEBVIEW2_DOWNLOAD_STATE = 0x0000000
	// COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED
	//
	//	The connection with the file host was broken. The InterruptReason property can be accessed from ICoreWebView2DownloadOperation. See COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON for descriptions of kinds of interrupt reasons. Host can check whether an interrupted download can be resumed with the CanResume property on the ICoreWebView2DownloadOperation. Once resumed, a download is in the COREWEBVIEW2_DOWNLOAD_STATE_IN_PROGRESS state.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_STATE values.
	COREWEBVIEW2_DOWNLOAD_STATE_INTERRUPTED COREWEBVIEW2_DOWNLOAD_STATE = 0x0000001
	// COREWEBVIEW2_DOWNLOAD_STATE_COMPLETED
	//
	//	The download completed successfully.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_STATE values.
	COREWEBVIEW2_DOWNLOAD_STATE_COMPLETED COREWEBVIEW2_DOWNLOAD_STATE = 0x0000002
)

// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON
//
//	Reason why a download was interrupted.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_download_interrupt_reason">See the Globals article.)
type COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = TOleEnum

const (
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NONE
	//
	//	No reason.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NONE COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000000
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_FAILED
	//
	//	Generic file error.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_FAILED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000001
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_ACCESS_DENIED
	//
	//	Access denied due to security restrictions.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_ACCESS_DENIED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000002
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_NO_SPACE
	//
	//	Disk full. User should free some space or choose a different location to store the file.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_NO_SPACE COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000003
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_NAME_TOO_LONG
	//
	//	Result file path with file name is too long.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_NAME_TOO_LONG COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000004
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_LARGE
	//
	//	File is too large for file system.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_LARGE COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000005
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_MALICIOUS
	//
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	//	Microsoft Defender Smartscreen detected a virus in the file.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_MALICIOUS COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000006
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TRANSIENT_ERROR
	//
	//	File was in use, too many files opened, or out of memory.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TRANSIENT_ERROR COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000007
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_BLOCKED_BY_POLICY
	//
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	//	File blocked by local policy.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_BLOCKED_BY_POLICY COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000008
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_SECURITY_CHECK_FAILED
	//
	//	Security check failed unexpectedly. Microsoft Defender SmartScreen could not scan this file.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_SECURITY_CHECK_FAILED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000009
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT
	//
	//	Seeking past the end of a file in opening a file, as part of resuming an interrupted download. The file did not exist or was not as large as expected. Partially downloaded file was truncated or deleted, and download will be restarted automatically.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_TOO_SHORT COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000000A
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH
	//
	//	Partial file did not match the expected hash and was deleted. Download will be restarted automatically.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_FILE_HASH_MISMATCH COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000000B
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_FAILED
	//
	//	Generic network error. User can retry the download manually.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_FAILED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000000C
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_TIMEOUT
	//
	//	Network operation timed out.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_TIMEOUT COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000000D
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_DISCONNECTED
	//
	//	Network connection lost. User can retry the download manually.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_DISCONNECTED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000000E
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_SERVER_DOWN
	//
	//	Server has gone down. User can retry the download manually.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_SERVER_DOWN COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000000F
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_INVALID_REQUEST
	//
	//	Network request invalid because original or redirected URI is invalid, has an unsupported scheme, or is disallowed by network policy.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_NETWORK_INVALID_REQUEST COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000010
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_FAILED
	//
	//	Generic server error. User can retry the download manually.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_FAILED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000011
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE
	//
	//	Server does not support range requests.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_NO_RANGE COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000012
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_BAD_CONTENT
	//
	//	Server does not have the requested data.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_BAD_CONTENT COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000013
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_UNAUTHORIZED
	//
	//	Server did not authorize access to resource.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_UNAUTHORIZED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000014
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CERTIFICATE_PROBLEM
	//
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	//	Server certificate problem.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CERTIFICATE_PROBLEM COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000015
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_FORBIDDEN
	//
	//	Server access forbidden.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_FORBIDDEN COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000016
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_UNEXPECTED_RESPONSE
	//
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	//	Unexpected server response. Responding server may not be intended server. User can retry the download manually.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_UNEXPECTED_RESPONSE COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000017
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CONTENT_LENGTH_MISMATCH
	//
	//	Server sent fewer bytes than the Content-Length header. Content-length header may be invalid or connection may have closed. Download is treated as complete unless there are [strong validators](https://tools.ietf.org/html/rfc7232#section-2) present to interrupt the download.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CONTENT_LENGTH_MISMATCH COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000018
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CROSS_ORIGIN_REDIRECT
	//
	//	Unexpected cross-origin redirect.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_SERVER_CROSS_ORIGIN_REDIRECT COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x0000019
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED
	//
	//	User canceled the download.
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_CANCELED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000001A
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_SHUTDOWN
	//
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	//	User shut down the WebView. Resuming downloads that were interrupted during shutdown is not yet supported.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_SHUTDOWN COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000001B
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_PAUSED
	//
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	//	User paused the download.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_USER_PAUSED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000001C
	// COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_DOWNLOAD_PROCESS_CRASHED
	//
	//	This is one of the COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON values.
	//	WebView crashed.
	COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON_DOWNLOAD_PROCESS_CRASHED COREWEBVIEW2_DOWNLOAD_INTERRUPT_REASON = 0x000001D
)

// COREWEBVIEW2_CLIENT_CERTIFICATE_KIND
//
//	Specifies the client certificate kind.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_client_certificate_kind">See the Globals article.)
type COREWEBVIEW2_CLIENT_CERTIFICATE_KIND = TOleEnum

const (
	// COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_SMART_CARD
	//
	//	Specifies smart card certificate.
	//	This is one of the COREWEBVIEW2_CLIENT_CERTIFICATE_KIND values.
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_SMART_CARD COREWEBVIEW2_CLIENT_CERTIFICATE_KIND = 0x0000000
	// COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_PIN
	//
	//	Specifies PIN certificate.
	//	This is one of the COREWEBVIEW2_CLIENT_CERTIFICATE_KIND values.
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_PIN COREWEBVIEW2_CLIENT_CERTIFICATE_KIND = 0x0000001
	// COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_OTHER
	//
	//	This is one of the COREWEBVIEW2_CLIENT_CERTIFICATE_KIND values.
	//	Specifies other certificate.
	COREWEBVIEW2_CLIENT_CERTIFICATE_KIND_OTHER COREWEBVIEW2_CLIENT_CERTIFICATE_KIND = 0x0000002
)

// COREWEBVIEW2_PRINT_ORIENTATION
//
//	The orientation for printing, used by the Orientation property on ICoreWebView2PrintSettings.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_print_orientation">See the Globals article.)
type COREWEBVIEW2_PRINT_ORIENTATION = TOleEnum

const (
	// COREWEBVIEW2_PRINT_ORIENTATION_PORTRAIT
	//
	//	This is one of the COREWEBVIEW2_PRINT_ORIENTATION values.
	//	Print the page(s) in portrait orientation.
	COREWEBVIEW2_PRINT_ORIENTATION_PORTRAIT COREWEBVIEW2_PRINT_ORIENTATION = 0x0000000
	// COREWEBVIEW2_PRINT_ORIENTATION_LANDSCAPE
	//
	//	This is one of the COREWEBVIEW2_PRINT_ORIENTATION values.
	//	Print the page(s) in landscape orientation.
	COREWEBVIEW2_PRINT_ORIENTATION_LANDSCAPE COREWEBVIEW2_PRINT_ORIENTATION = 0x0000001
)

// COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT
//
//	The default download dialog can be aligned to any of the WebView corners by setting the DefaultDownloadDialogCornerAlignment property. The default position is top-right corner.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_default_download_dialog_corner_alignment">See the Globals article.)
type COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT = TOleEnum

const (
	// COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_TOP_LEFT
	//
	//	Top-left corner of the WebView.
	//	This is one of the COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT values.
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_TOP_LEFT COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT = 0x0000000
	// COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_TOP_RIGHT
	//
	//	Top-right corner of the WebView.
	//	This is one of the COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT values.
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_TOP_RIGHT COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT = 0x0000001
	// COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_BOTTOM_LEFT
	//
	//	Bottom-left corner of the WebView.
	//	This is one of the COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT values.
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_BOTTOM_LEFT COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT = 0x0000002
	// COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_BOTTOM_RIGHT
	//
	//	Bottom-right corner of the WebView.
	//	This is one of the COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT values.
	COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT_BOTTOM_RIGHT COREWEBVIEW2_DEFAULT_DOWNLOAD_DIALOG_CORNER_ALIGNMENT = 0x0000003
)

// COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND
//
//	Specifies the menu item kind for the ICoreWebView2ContextMenuItem.get_Kind method
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_context_menu_item_kind">See the Globals article.)
type COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND = TOleEnum

const (
	// COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_COMMAND
	//
	//	Specifies a command menu item kind.
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND values.
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_COMMAND COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND = 0x0000000
	// COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_CHECK_BOX
	//
	//	Specifies a check box menu item kind. ContextMenuItem objects of this kind will need the IsChecked property to determine current state of the check box.
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND values.
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_CHECK_BOX COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND = 0x0000001
	// COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_RADIO
	//
	//	Specifies a radio button menu item kind. ContextMenuItem objects of this kind will need the IsChecked property to determine current state of the radio button.
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND values.
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_RADIO COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND = 0x0000002
	// COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SEPARATOR
	//
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND values.
	//	Specifies a separator menu item kind. ContextMenuItem objects of this kind are used to signal a visual separator with no functionality.
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SEPARATOR COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND = 0x0000003
	// COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SUBMENU
	//
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND values.
	//	Specifies a submenu menu item kind. ContextMenuItem objects of this kind will contain a ContextMenuItemCollection of its children ContextMenuItem objects.
	COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND_SUBMENU COREWEBVIEW2_CONTEXT_MENU_ITEM_KIND = 0x0000004
)

// COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND
//
//	Indicates the kind of context for which the context menu was created for the ICoreWebView2ContextMenuTarget::get_Kind method. This enum will always represent the active element that caused the context menu request. If there is a selection with multiple images, audio and text, for example, the element that the end user right clicks on within this selection will be the option represented by this enum.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_context_menu_target_kind">See the Globals article.)
type COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND = TOleEnum

const (
	// COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_PAGE
	//
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND values.
	//	Indicates that the context menu was created for the page without any additional content.
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_PAGE COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND = 0x0000000
	// COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_IMAGE
	//
	//	Indicates that the context menu was created for an image element.
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND values.
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_IMAGE COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND = 0x0000001
	// COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_SELECTED_TEXT
	//
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND values.
	//	Indicates that the context menu was created for selected text.
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_SELECTED_TEXT COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND = 0x0000002
	// COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_AUDIO
	//
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND values.
	//	Indicates that the context menu was created for an audio element.
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_AUDIO COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND = 0x0000003
	// COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_VIDEO
	//
	//	This is one of the COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND values.
	//	Indicates that the context menu was created for a video element.
	COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND_VIDEO COREWEBVIEW2_CONTEXT_MENU_TARGET_KIND = 0x0000004
)

// COREWEBVIEW2_PREFERRED_COLOR_SCHEME
//
//	An enum to represent the options for WebView2 color scheme: auto, light, or dark.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_preferred_color_scheme">See the Globals article.)
type COREWEBVIEW2_PREFERRED_COLOR_SCHEME = TOleEnum

const (
	// COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO
	//
	//	Auto color scheme.
	//	This is one of the COREWEBVIEW2_PREFERRED_COLOR_SCHEME values.
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_AUTO COREWEBVIEW2_PREFERRED_COLOR_SCHEME = 0x0000000
	// COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT
	//
	//	Light color scheme.
	//	This is one of the COREWEBVIEW2_PREFERRED_COLOR_SCHEME values.
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_LIGHT COREWEBVIEW2_PREFERRED_COLOR_SCHEME = 0x0000001
	// COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK
	//
	//	Dark color scheme.
	//	This is one of the COREWEBVIEW2_PREFERRED_COLOR_SCHEME values.
	COREWEBVIEW2_PREFERRED_COLOR_SCHEME_DARK COREWEBVIEW2_PREFERRED_COLOR_SCHEME = 0x0000002
)

// COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION
//
//	Specifies the action type when server certificate error is detected to be used in the ICoreWebView2ServerCertificateErrorDetectedEventArgs interface.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_server_certificate_error_action">See the Globals article.)
type COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION = TOleEnum

const (
	// COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_ALWAYS_ALLOW
	//
	//	Indicates to ignore the warning and continue the request with the TLS certificate. This decision is cached for the RequestUri's host and the server certificate in the session.
	//	This is one of the COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION values.
	COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_ALWAYS_ALLOW COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION = 0x0000000
	// COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_CANCEL
	//
	//	Indicates to reject the certificate and cancel the request.
	//	This is one of the COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION values.
	COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_CANCEL COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION = 0x0000001
	// COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_DEFAULT
	//
	//	Indicates to display the default TLS interstitial error page to user for page navigations. For others TLS certificate is rejected and the request is cancelled.
	//	This is one of the COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION values.
	COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION_DEFAULT COREWEBVIEW2_SERVER_CERTIFICATE_ERROR_ACTION = 0x0000002
)

// COREWEBVIEW2_FAVICON_IMAGE_FORMAT
//
//	Specifies the image format to use for favicon.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_favicon_image_format">See the Globals article.)
type COREWEBVIEW2_FAVICON_IMAGE_FORMAT = TOleEnum

const (
	// COREWEBVIEW2_FAVICON_IMAGE_FORMAT_PNG
	//
	//	Indicates that the PNG image format is used.
	//	This is one of the COREWEBVIEW2_FAVICON_IMAGE_FORMAT values.
	COREWEBVIEW2_FAVICON_IMAGE_FORMAT_PNG COREWEBVIEW2_FAVICON_IMAGE_FORMAT = 0x0000000
	// COREWEBVIEW2_FAVICON_IMAGE_FORMAT_JPEG
	//
	//	Indicates the JPEG image format is used.
	//	This is one of the COREWEBVIEW2_FAVICON_IMAGE_FORMAT values.
	COREWEBVIEW2_FAVICON_IMAGE_FORMAT_JPEG COREWEBVIEW2_FAVICON_IMAGE_FORMAT = 0x0000001
)

// COREWEBVIEW2_PRINT_STATUS
//
//	Indicates the status for printing.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_print_status">See the Globals article.)
type COREWEBVIEW2_PRINT_STATUS = TOleEnum

const (
	// COREWEBVIEW2_PRINT_STATUS_SUCCEEDED
	//
	//	Indicates that the print operation is succeeded.
	//	This is one of the COREWEBVIEW2_PRINT_STATUS values.
	COREWEBVIEW2_PRINT_STATUS_SUCCEEDED COREWEBVIEW2_PRINT_STATUS = 0x0000000
	// COREWEBVIEW2_PRINT_STATUS_PRINTER_UNAVAILABLE
	//
	//	Indicates that the printer is not available.
	//	This is one of the COREWEBVIEW2_PRINT_STATUS values.
	COREWEBVIEW2_PRINT_STATUS_PRINTER_UNAVAILABLE COREWEBVIEW2_PRINT_STATUS = 0x0000001
	// COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR
	//
	//	Indicates that the print operation is failed.
	//	This is one of the COREWEBVIEW2_PRINT_STATUS values.
	COREWEBVIEW2_PRINT_STATUS_OTHER_ERROR COREWEBVIEW2_PRINT_STATUS = 0x0000002
)

// COREWEBVIEW2_PRINT_DIALOG_KIND
//
//	Specifies the print dialog kind.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_print_dialog_kind">See the Globals article.)
type COREWEBVIEW2_PRINT_DIALOG_KIND = TOleEnum

const (
	// COREWEBVIEW2_PRINT_DIALOG_KIND_BROWSER
	//
	//	Opens the browser print preview dialog.
	//	This is one of the COREWEBVIEW2_PRINT_DIALOG_KIND values.
	COREWEBVIEW2_PRINT_DIALOG_KIND_BROWSER COREWEBVIEW2_PRINT_DIALOG_KIND = 0x0000000
	// COREWEBVIEW2_PRINT_DIALOG_KIND_SYSTEM
	//
	//	Opens the system print dialog.
	//	This is one of the COREWEBVIEW2_PRINT_DIALOG_KIND values.
	COREWEBVIEW2_PRINT_DIALOG_KIND_SYSTEM COREWEBVIEW2_PRINT_DIALOG_KIND = 0x0000001
)

// COREWEBVIEW2_SHARED_BUFFER_ACCESS
//
//	Specifies the desired access from script to CoreWebView2SharedBuffer.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_shared_buffer_access">See the Globals article.)
type COREWEBVIEW2_SHARED_BUFFER_ACCESS = TOleEnum

const (
	// COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY
	//
	//	This is one of the COREWEBVIEW2_SHARED_BUFFER_ACCESS values.
	//	Script from web page only has read access to the shared buffer.
	COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_ONLY COREWEBVIEW2_SHARED_BUFFER_ACCESS = 0x0000000
	// COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_WRITE
	//
	//	Script from web page has read and write access to the shared buffer.
	//	This is one of the COREWEBVIEW2_SHARED_BUFFER_ACCESS values.
	COREWEBVIEW2_SHARED_BUFFER_ACCESS_READ_WRITE COREWEBVIEW2_SHARED_BUFFER_ACCESS = 0x0000001
)

// COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND
//
//	Specifies the browser process exit type used in the ICoreWebView2BrowserProcessExitedEventArgs interface.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_browser_process_exit_kind">See the Globals article.)
type COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND = TOleEnum

const (
	// COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_NORMAL
	//
	//	Indicates that the browser process ended normally.
	//	This is one of the COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND values.
	COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_NORMAL COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND = 0x0000000
	// COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_FAILED
	//
	//	Indicates that the browser process ended unexpectedly. A ProcessFailed event will also be sent to listening WebViews from the ICoreWebView2Environment associated to the failed process.
	//	This is one of the COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND values.
	COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND_FAILED COREWEBVIEW2_BROWSER_PROCESS_EXIT_KIND = 0x0000001
)

// COREWEBVIEW2_MOUSE_EVENT_KIND
//
//	Mouse event type used by SendMouseInput to convey the type of mouse event being sent to WebView. The values of this enum align with the matching WM_* window messages.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_mouse_event_kind">See the Globals article.)
type COREWEBVIEW2_MOUSE_EVENT_KIND = TOleEnum

const (
	// COREWEBVIEW2_MOUSE_EVENT_KIND_HORIZONTAL_WHEEL
	//
	//	Mouse horizontal wheel scroll event, WM_MOUSEHWHEEL.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_HORIZONTAL_WHEEL COREWEBVIEW2_MOUSE_EVENT_KIND = 0x000020E
	// COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_DOUBLE_CLICK
	//
	//	Left button double click mouse event, WM_LBUTTONDBLCLK.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_DOUBLE_CLICK COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000203
	// COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_DOWN
	//
	//	Left button down mouse event, WM_LBUTTONDOWN.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_DOWN COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000201
	// COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_UP
	//
	//	Left button up mouse event, WM_LBUTTONUP.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEFT_BUTTON_UP COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000202
	// COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE
	//
	//	Mouse leave event, WM_MOUSELEAVE.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_LEAVE COREWEBVIEW2_MOUSE_EVENT_KIND = 0x00002A3
	// COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_DOUBLE_CLICK
	//
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	//	Middle button double click mouse event, WM_MBUTTONDBLCLK.
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_DOUBLE_CLICK COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000209
	// COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_DOWN
	//
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	//	Middle button down mouse event, WM_MBUTTONDOWN.
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_DOWN COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000207
	// COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_UP
	//
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	//	Middle button up mouse event, WM_MBUTTONUP.
	COREWEBVIEW2_MOUSE_EVENT_KIND_MIDDLE_BUTTON_UP COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000208
	// COREWEBVIEW2_MOUSE_EVENT_KIND_MOVE
	//
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	//	Mouse move event, WM_MOUSEMOVE.
	COREWEBVIEW2_MOUSE_EVENT_KIND_MOVE COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000200
	// COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_DOUBLE_CLICK
	//
	//	Right button double click mouse event, WM_RBUTTONDBLCLK.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_DOUBLE_CLICK COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000206
	// COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_DOWN
	//
	//	Right button down mouse event, WM_RBUTTONDOWN.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_DOWN COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000204
	// COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_UP
	//
	//	Right button up mouse event, WM_RBUTTONUP.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_RIGHT_BUTTON_UP COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0000205
	// COREWEBVIEW2_MOUSE_EVENT_KIND_WHEEL
	//
	//	Mouse wheel scroll event, WM_MOUSEWHEEL.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_WHEEL COREWEBVIEW2_MOUSE_EVENT_KIND = 0x000020A
	// COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOUBLE_CLICK
	//
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	//	First or second X button double click mouse event, WM_XBUTTONDBLCLK.
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOUBLE_CLICK COREWEBVIEW2_MOUSE_EVENT_KIND = 0x000020D
	// COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOWN
	//
	//	First or second X button down mouse event, WM_XBUTTONDOWN.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_DOWN COREWEBVIEW2_MOUSE_EVENT_KIND = 0x000020B
	// COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_UP
	//
	//	First or second X button up mouse event, WM_XBUTTONUP.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_X_BUTTON_UP COREWEBVIEW2_MOUSE_EVENT_KIND = 0x000020C
	// COREWEBVIEW2_MOUSE_EVENT_KIND_NON_CLIENT_RIGHT_BUTTON_DOWN
	//
	//	Mouse Right Button Down event over a nonclient area, WM_NCRBUTTONDOWN.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_NON_CLIENT_RIGHT_BUTTON_DOWN COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0A4
	// COREWEBVIEW2_MOUSE_EVENT_KIND_NON_CLIENT_RIGHT_BUTTON_UP
	//
	//	Mouse Right Button up event over a nonclient area, WM_NCRBUTTONUP.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_KIND values.
	COREWEBVIEW2_MOUSE_EVENT_KIND_NON_CLIENT_RIGHT_BUTTON_UP COREWEBVIEW2_MOUSE_EVENT_KIND = 0x0A5
)

// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS
//
//	Mouse event virtual keys associated with a COREWEBVIEW2_MOUSE_EVENT_KIND for SendMouseInput. These values can be combined into a bit flag if more than one virtual key is pressed for the event. The values of this enum align with the matching MK_* mouse keys.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_mouse_event_virtual_keys">See the Globals article.)
type COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = TOleEnum

const (
	// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_NONE
	//
	//	No additional keys pressed.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS values.
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_NONE COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = 0x0000000
	// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_LEFT_BUTTON
	//
	//	Left mouse button is down, MK_LBUTTON.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS values.
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_LEFT_BUTTON COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = 0x0000001
	// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_RIGHT_BUTTON
	//
	//	Right mouse button is down, MK_RBUTTON.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS values.
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_RIGHT_BUTTON COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = 0x0000002
	// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_SHIFT
	//
	//	SHIFT key is down, MK_SHIFT.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS values.
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_SHIFT COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = 0x0000004
	// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_CONTROL
	//
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS values.
	//	CTRL key is down, MK_CONTROL.
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_CONTROL COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = 0x0000008
	// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_MIDDLE_BUTTON
	//
	//	Middle mouse button is down, MK_MBUTTON.
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS values.
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_MIDDLE_BUTTON COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = 0x0000010
	// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_X_BUTTON1
	//
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS values.
	//	First X button is down, MK_XBUTTON1
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_X_BUTTON1 COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = 0x0000020
	// COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_X_BUTTON2
	//
	//	Second X button is down, MK_XBUTTON2
	//	This is one of the COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS values.
	COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS_X_BUTTON2 COREWEBVIEW2_MOUSE_EVENT_VIRTUAL_KEYS = 0x0000040
)

// COREWEBVIEW2_POINTER_EVENT_KIND
//
//	Pointer event type used by SendPointerInput to convey the type of pointer event being sent to WebView. The values of this enum align with the matching WM_POINTER* window messages.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_pointer_event_kind">See the Globals article.)
type COREWEBVIEW2_POINTER_EVENT_KIND = TOleEnum

const (
	// COREWEBVIEW2_POINTER_EVENT_KIND_ACTIVATE
	//
	//	Corresponds to WM_POINTERACTIVATE.
	//	This is one of the COREWEBVIEW2_POINTER_EVENT_KIND values.
	COREWEBVIEW2_POINTER_EVENT_KIND_ACTIVATE COREWEBVIEW2_POINTER_EVENT_KIND = 0x000024B
	// COREWEBVIEW2_POINTER_EVENT_KIND_DOWN
	//
	//	Corresponds to WM_POINTERDOWN.
	//	This is one of the COREWEBVIEW2_POINTER_EVENT_KIND values.
	COREWEBVIEW2_POINTER_EVENT_KIND_DOWN COREWEBVIEW2_POINTER_EVENT_KIND = 0x0000246
	// COREWEBVIEW2_POINTER_EVENT_KIND_ENTER
	//
	//	This is one of the COREWEBVIEW2_POINTER_EVENT_KIND values.
	//	Corresponds to WM_POINTERENTER.
	COREWEBVIEW2_POINTER_EVENT_KIND_ENTER COREWEBVIEW2_POINTER_EVENT_KIND = 0x0000249
	// COREWEBVIEW2_POINTER_EVENT_KIND_LEAVE
	//
	//	This is one of the COREWEBVIEW2_POINTER_EVENT_KIND values.
	//	Corresponds to WM_POINTERLEAVE.
	COREWEBVIEW2_POINTER_EVENT_KIND_LEAVE COREWEBVIEW2_POINTER_EVENT_KIND = 0x000024A
	// COREWEBVIEW2_POINTER_EVENT_KIND_UP
	//
	//	This is one of the COREWEBVIEW2_POINTER_EVENT_KIND values.
	//	Corresponds to WM_POINTERUP.
	COREWEBVIEW2_POINTER_EVENT_KIND_UP COREWEBVIEW2_POINTER_EVENT_KIND = 0x0000247
	// COREWEBVIEW2_POINTER_EVENT_KIND_UPDATE
	//
	//	Corresponds to WM_POINTERUPDATE.
	//	This is one of the COREWEBVIEW2_POINTER_EVENT_KIND values.
	COREWEBVIEW2_POINTER_EVENT_KIND_UPDATE COREWEBVIEW2_POINTER_EVENT_KIND = 0x0000245
)

// COREWEBVIEW2_BOUNDS_MODE
//
//	Mode for how the Bounds property is interpreted in relation to the RasterizationScale property.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_bounds_mode">See the Globals article.)
type COREWEBVIEW2_BOUNDS_MODE = TOleEnum

const (
	// COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS
	//
	//	Bounds property represents raw pixels. Physical size of Webview is not impacted by RasterizationScale.
	//	This is one of the COREWEBVIEW2_BOUNDS_MODE values.
	COREWEBVIEW2_BOUNDS_MODE_USE_RAW_PIXELS COREWEBVIEW2_BOUNDS_MODE = 0x0000000
	// COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE
	//
	//	Bounds property represents logical pixels and the RasterizationScale property is used to get the physical size of the WebView.
	//	This is one of the COREWEBVIEW2_BOUNDS_MODE values.
	COREWEBVIEW2_BOUNDS_MODE_USE_RASTERIZATION_SCALE COREWEBVIEW2_BOUNDS_MODE = 0x0000001
)

// COREWEBVIEW2_PROCESS_KIND
//
//	Indicates the process type used in the ICoreWebView2ProcessInfo interface.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_process_kind">See the Globals article.)
type COREWEBVIEW2_PROCESS_KIND = TOleEnum

const (
	// COREWEBVIEW2_PROCESS_KIND_BROWSER
	//
	//	Indicates the browser process kind.
	//	This is one of the COREWEBVIEW2_PROCESS_KIND values.
	COREWEBVIEW2_PROCESS_KIND_BROWSER COREWEBVIEW2_PROCESS_KIND = 0x0000000
	// COREWEBVIEW2_PROCESS_KIND_RENDERER
	//
	//	This is one of the COREWEBVIEW2_PROCESS_KIND values.
	//	Indicates the render process kind.
	COREWEBVIEW2_PROCESS_KIND_RENDERER COREWEBVIEW2_PROCESS_KIND = 0x0000001
	// COREWEBVIEW2_PROCESS_KIND_UTILITY
	//
	//	Indicates the utility process kind.
	//	This is one of the COREWEBVIEW2_PROCESS_KIND values.
	COREWEBVIEW2_PROCESS_KIND_UTILITY COREWEBVIEW2_PROCESS_KIND = 0x0000002
	// COREWEBVIEW2_PROCESS_KIND_SANDBOX_HELPER
	//
	//	Indicates the sandbox helper process kind.
	//	This is one of the COREWEBVIEW2_PROCESS_KIND values.
	COREWEBVIEW2_PROCESS_KIND_SANDBOX_HELPER COREWEBVIEW2_PROCESS_KIND = 0x0000003
	// COREWEBVIEW2_PROCESS_KIND_GPU
	//
	//	Indicates the GPU process kind.
	//	This is one of the COREWEBVIEW2_PROCESS_KIND values.
	COREWEBVIEW2_PROCESS_KIND_GPU COREWEBVIEW2_PROCESS_KIND = 0x0000004
	// COREWEBVIEW2_PROCESS_KIND_PPAPI_PLUGIN
	//
	//	Indicates the PPAPI plugin process kind.
	//	This is one of the COREWEBVIEW2_PROCESS_KIND values.
	COREWEBVIEW2_PROCESS_KIND_PPAPI_PLUGIN COREWEBVIEW2_PROCESS_KIND = 0x0000005
	// COREWEBVIEW2_PROCESS_KIND_PPAPI_BROKER
	//
	//	Indicates the PPAPI plugin broker process kind.
	//	This is one of the COREWEBVIEW2_PROCESS_KIND values.
	COREWEBVIEW2_PROCESS_KIND_PPAPI_BROKER COREWEBVIEW2_PROCESS_KIND = 0x0000006
)

// COREWEBVIEW2_FRAME_KIND
//
//	Indicates the frame type used in the `ICoreWebView2FrameInfo` interface.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_frame_kind">See the Globals article.)
type COREWEBVIEW2_FRAME_KIND = TOleEnum

const (
	// COREWEBVIEW2_FRAME_KIND_UNKNOWN
	//
	//	Indicates that the frame is an unknown type frame. We may extend this enum type to identify more frame kinds in the future.
	//	This is one of the COREWEBVIEW2_FRAME_KIND values.
	COREWEBVIEW2_FRAME_KIND_UNKNOWN COREWEBVIEW2_FRAME_KIND = 0x0000000
	// COREWEBVIEW2_FRAME_KIND_MAIN_FRAME
	//
	//	Indicates that the frame is a primary main frame(webview).
	//	This is one of the COREWEBVIEW2_FRAME_KIND values.
	COREWEBVIEW2_FRAME_KIND_MAIN_FRAME COREWEBVIEW2_FRAME_KIND = 0x0000001
	// COREWEBVIEW2_FRAME_KIND_IFRAME
	//
	//	Indicates that the frame is an iframe.
	//	This is one of the COREWEBVIEW2_FRAME_KIND values.
	COREWEBVIEW2_FRAME_KIND_IFRAME COREWEBVIEW2_FRAME_KIND = 0x0000002
	// COREWEBVIEW2_FRAME_KIND_EMBED
	//
	//	Indicates that the frame is an embed element.
	//	This is one of the COREWEBVIEW2_FRAME_KIND values.
	COREWEBVIEW2_FRAME_KIND_EMBED COREWEBVIEW2_FRAME_KIND = 0x0000003
	// COREWEBVIEW2_FRAME_KIND_OBJECT
	//
	//	Indicates that the frame is an object element.
	//	This is one of the COREWEBVIEW2_FRAME_KIND values.
	COREWEBVIEW2_FRAME_KIND_OBJECT COREWEBVIEW2_FRAME_KIND = 0x0000004
)

// COREWEBVIEW2_NAVIGATION_KIND
//
//	Specifies the navigation kind of each navigation.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_navigation_kind">See the Globals article.)
type COREWEBVIEW2_NAVIGATION_KIND = TOleEnum

const (
	// COREWEBVIEW2_NAVIGATION_KIND_RELOAD
	//
	//	A navigation caused by CoreWebView2.Reload(), location.reload(), the end user using F5 or other UX, or other reload mechanisms to reload the current document without modifying the navigation history.
	//	This is one of the COREWEBVIEW2_NAVIGATION_KIND values.
	COREWEBVIEW2_NAVIGATION_KIND_RELOAD COREWEBVIEW2_NAVIGATION_KIND = 0x0000000
	// COREWEBVIEW2_NAVIGATION_KIND_BACK_OR_FORWARD
	//
	//	A navigation back or forward to a different entry in the session navigation history, like via CoreWebView2.Back(), location.back(), the end user pressing Alt+Left or other UX, or other mechanisms to navigate back or forward in the current session navigation history.
	//	This is one of the COREWEBVIEW2_NAVIGATION_KIND values. This kind doesn't distinguish back or forward, because we can't distinguish it from origin source blink.mojom.NavigationType.
	COREWEBVIEW2_NAVIGATION_KIND_BACK_OR_FORWARD COREWEBVIEW2_NAVIGATION_KIND = 0x0000001
	// COREWEBVIEW2_NAVIGATION_KIND_NEW_DOCUMENT
	//
	//	A navigation to another document, which can be caused by CoreWebView2.Navigate(), window.location.href = ..., or other WebView2 or DOM APIs that navigate to a new URI.
	//	This is one of the COREWEBVIEW2_NAVIGATION_KIND values.
	COREWEBVIEW2_NAVIGATION_KIND_NEW_DOCUMENT COREWEBVIEW2_NAVIGATION_KIND = 0x0000002
)

// COREWEBVIEW2_PRINT_COLLATION
//
//	Specifies the collation for a print.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_print_collation">See the Globals article.)
type COREWEBVIEW2_PRINT_COLLATION = TOleEnum

const (
	// COREWEBVIEW2_PRINT_COLLATION_DEFAULT
	//
	//	The default collation for a printer.
	//	This is one of the COREWEBVIEW2_PRINT_COLLATION values.
	COREWEBVIEW2_PRINT_COLLATION_DEFAULT COREWEBVIEW2_PRINT_COLLATION = 0x0000000
	// COREWEBVIEW2_PRINT_COLLATION_COLLATED
	//
	//	Indicate that the collation has been selected for the printed output.
	//	This is one of the COREWEBVIEW2_PRINT_COLLATION values.
	COREWEBVIEW2_PRINT_COLLATION_COLLATED COREWEBVIEW2_PRINT_COLLATION = 0x0000001
	// COREWEBVIEW2_PRINT_COLLATION_UNCOLLATED
	//
	//	Indicate that the collation has not been selected for the printed output.
	//	This is one of the COREWEBVIEW2_PRINT_COLLATION values.
	COREWEBVIEW2_PRINT_COLLATION_UNCOLLATED COREWEBVIEW2_PRINT_COLLATION = 0x0000002
)

// COREWEBVIEW2_PRINT_COLOR_MODE
//
//	Specifies the color mode for a print.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_print_color_mode">See the Globals article.)
type COREWEBVIEW2_PRINT_COLOR_MODE = TOleEnum

const (
	// COREWEBVIEW2_PRINT_COLOR_MODE_DEFAULT
	//
	//	This is one of the COREWEBVIEW2_PRINT_COLOR_MODE values.
	//	The default color mode for a printer.
	COREWEBVIEW2_PRINT_COLOR_MODE_DEFAULT COREWEBVIEW2_PRINT_COLOR_MODE = 0x0000000
	// COREWEBVIEW2_PRINT_COLOR_MODE_COLOR
	//
	//	Indicate that the printed output will be in color.
	//	This is one of the COREWEBVIEW2_PRINT_COLOR_MODE values.
	COREWEBVIEW2_PRINT_COLOR_MODE_COLOR COREWEBVIEW2_PRINT_COLOR_MODE = 0x0000001
	// COREWEBVIEW2_PRINT_COLOR_MODE_GRAYSCALE
	//
	//	Indicate that the printed output will be in shades of gray.
	//	This is one of the COREWEBVIEW2_PRINT_COLOR_MODE values.
	COREWEBVIEW2_PRINT_COLOR_MODE_GRAYSCALE COREWEBVIEW2_PRINT_COLOR_MODE = 0x0000002
)

// COREWEBVIEW2_PRINT_DUPLEX
//
//	Specifies the duplex option for a print.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_print_duplex">See the Globals article.)
type COREWEBVIEW2_PRINT_DUPLEX = TOleEnum

const (
	// COREWEBVIEW2_PRINT_DUPLEX_DEFAULT
	//
	//	The default duplex for a printer.
	//	This is one of the COREWEBVIEW2_PRINT_DUPLEX values.
	COREWEBVIEW2_PRINT_DUPLEX_DEFAULT COREWEBVIEW2_PRINT_DUPLEX = 0x0000000
	// COREWEBVIEW2_PRINT_DUPLEX_ONE_SIDED
	//
	//	Print on only one side of the sheet.
	//	This is one of the COREWEBVIEW2_PRINT_DUPLEX values.
	COREWEBVIEW2_PRINT_DUPLEX_ONE_SIDED COREWEBVIEW2_PRINT_DUPLEX = 0x0000001
	// COREWEBVIEW2_PRINT_DUPLEX_TWO_SIDED_LONG_EDGE
	//
	//	Print on both sides of the sheet, flipped along the long edge.
	//	This is one of the COREWEBVIEW2_PRINT_DUPLEX values.
	COREWEBVIEW2_PRINT_DUPLEX_TWO_SIDED_LONG_EDGE COREWEBVIEW2_PRINT_DUPLEX = 0x0000002
	// COREWEBVIEW2_PRINT_DUPLEX_TWO_SIDED_SHORT_EDGE
	//
	//	Print on both sides of the sheet, flipped along the short edge.
	//	This is one of the COREWEBVIEW2_PRINT_DUPLEX values.
	COREWEBVIEW2_PRINT_DUPLEX_TWO_SIDED_SHORT_EDGE COREWEBVIEW2_PRINT_DUPLEX = 0x0000003
)

// COREWEBVIEW2_PRINT_MEDIA_SIZE
//
//	Specifies the media size for a print.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_print_media_size">See the Globals article.)
type COREWEBVIEW2_PRINT_MEDIA_SIZE = TOleEnum

const (
	// COREWEBVIEW2_PRINT_MEDIA_SIZE_DEFAULT
	//
	//	The default media size for a printer.
	//	This is one of the COREWEBVIEW2_PRINT_MEDIA_SIZE values.
	COREWEBVIEW2_PRINT_MEDIA_SIZE_DEFAULT COREWEBVIEW2_PRINT_MEDIA_SIZE = 0x0000000
	// COREWEBVIEW2_PRINT_MEDIA_SIZE_CUSTOM
	//
	//	Indicate custom media size that is specific to the printer.
	//	This is one of the COREWEBVIEW2_PRINT_MEDIA_SIZE values.
	COREWEBVIEW2_PRINT_MEDIA_SIZE_CUSTOM COREWEBVIEW2_PRINT_MEDIA_SIZE = 0x0000001
)

// COREWEBVIEW2_PROCESS_FAILED_REASON
//
//	Specifies the process failure reason used in the ICoreWebView2ProcessFailedEventHandler interface.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_process_failed_reason">See the Globals article.)
type COREWEBVIEW2_PROCESS_FAILED_REASON = TOleEnum

const (
	// COREWEBVIEW2_PROCESS_FAILED_REASON_UNEXPECTED
	//
	//	An unexpected process failure occurred.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_REASON values.
	COREWEBVIEW2_PROCESS_FAILED_REASON_UNEXPECTED COREWEBVIEW2_PROCESS_FAILED_REASON = 0x0000000
	// COREWEBVIEW2_PROCESS_FAILED_REASON_UNRESPONSIVE
	//
	//	The process became unresponsive. This only applies to the main frame's render process.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_REASON values.
	COREWEBVIEW2_PROCESS_FAILED_REASON_UNRESPONSIVE COREWEBVIEW2_PROCESS_FAILED_REASON = 0x0000001
	// COREWEBVIEW2_PROCESS_FAILED_REASON_TERMINATED
	//
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_REASON values.
	//	The process was terminated. For example, from Task Manager.
	COREWEBVIEW2_PROCESS_FAILED_REASON_TERMINATED COREWEBVIEW2_PROCESS_FAILED_REASON = 0x0000002
	// COREWEBVIEW2_PROCESS_FAILED_REASON_CRASHED
	//
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_REASON values.
	//	The process crashed. Most crashes will generate dumps in the location indicated by `ICoreWebView2Environment11.get_FailureReportFolderPath`.
	COREWEBVIEW2_PROCESS_FAILED_REASON_CRASHED COREWEBVIEW2_PROCESS_FAILED_REASON = 0x0000003
	// COREWEBVIEW2_PROCESS_FAILED_REASON_LAUNCH_FAILED
	//
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_REASON values.
	//	The process failed to launch.
	COREWEBVIEW2_PROCESS_FAILED_REASON_LAUNCH_FAILED COREWEBVIEW2_PROCESS_FAILED_REASON = 0x0000004
	// COREWEBVIEW2_PROCESS_FAILED_REASON_OUT_OF_MEMORY
	//
	//	The process terminated due to running out of memory.
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_REASON values.
	COREWEBVIEW2_PROCESS_FAILED_REASON_OUT_OF_MEMORY COREWEBVIEW2_PROCESS_FAILED_REASON = 0x0000005
	// COREWEBVIEW2_PROCESS_FAILED_REASON_PROFILE_DELETED
	//
	//	This is one of the COREWEBVIEW2_PROCESS_FAILED_REASON values.
	//	The process exited because its corresponding profile was deleted.
	COREWEBVIEW2_PROCESS_FAILED_REASON_PROFILE_DELETED COREWEBVIEW2_PROCESS_FAILED_REASON = 0x0000006
)

// COREWEBVIEW2_BROWSING_DATA_KINDS
//
//	Specifies the datatype for the ICoreWebView2Profile2.ClearBrowsingData method.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_browsing_data_kinds">See the Globals article.)
type COREWEBVIEW2_BROWSING_DATA_KINDS = TOleEnum

const (
	// COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS
	//
	//	Specifies file systems data.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 0
	// COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB
	//
	//	Specifies data stored by the IndexedDB DOM feature.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 1
	// COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE
	//
	//	Specifies data stored by the localStorage DOM API.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 2
	// COREWEBVIEW2_BROWSING_DATA_KINDS_WEB_SQL
	//
	//	Specifies data stored by the Web SQL database DOM API.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_WEB_SQL COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 3
	// COREWEBVIEW2_BROWSING_DATA_KINDS_CACHE_STORAGE
	//
	//	Specifies data stored by the CacheStorage DOM API.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_CACHE_STORAGE COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 4
	// COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_DOM_STORAGE
	//
	//	Specifies DOM storage data, now and future. This browsing data kind is inclusive of COREWEBVIEW2_BROWSING_DATA_KINDS_FILE_SYSTEMS, COREWEBVIEW2_BROWSING_DATA_KINDS_INDEXED_DB, COREWEBVIEW2_BROWSING_DATA_KINDS_LOCAL_STORAGE, COREWEBVIEW2_BROWSING_DATA_KINDS_WEB_SQL, COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS, COREWEBVIEW2_BROWSING_DATA_KINDS_CACHE_STORAGE, and some other data kinds not listed yet to keep consistent with [DOM-accessible storage](https://www.w3.org/TR/clear-site-data/#storage).
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_DOM_STORAGE COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 5
	// COREWEBVIEW2_BROWSING_DATA_KINDS_COOKIES
	//
	//	Specifies HTTP cookies data.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_COOKIES COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 6
	// COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_SITE
	//
	//	Specifies all site data, now and future. This browsing data kind is inclusive of COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_DOM_STORAGE and COREWEBVIEW2_BROWSING_DATA_KINDS_COOKIES. New site data types may be added to this data kind in the future.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_SITE COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 7
	// COREWEBVIEW2_BROWSING_DATA_KINDS_DISK_CACHE
	//
	//	Specifies disk cache.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_DISK_CACHE COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 8
	// COREWEBVIEW2_BROWSING_DATA_KINDS_DOWNLOAD_HISTORY
	//
	//	Specifies download history data.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_DOWNLOAD_HISTORY COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 9
	// COREWEBVIEW2_BROWSING_DATA_KINDS_GENERAL_AUTOFILL
	//
	//	Specifies general autofill form data. This excludes password information and includes information like: names, street and email addresses, phone numbers, and arbitrary input. This also includes payment data.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_GENERAL_AUTOFILL COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 10
	// COREWEBVIEW2_BROWSING_DATA_KINDS_PASSWORD_AUTOSAVE
	//
	//	Specifies password autosave data.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_PASSWORD_AUTOSAVE COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 11
	// COREWEBVIEW2_BROWSING_DATA_KINDS_BROWSING_HISTORY
	//
	//	Specifies browsing history data.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_BROWSING_HISTORY COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 12
	// COREWEBVIEW2_BROWSING_DATA_KINDS_SETTINGS
	//
	//	Specifies settings data.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_SETTINGS COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 13
	// COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_PROFILE
	//
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	//	Specifies profile data that should be wiped to make it look like a new profile. This does not delete account-scoped data like passwords but will remove access to account-scoped data by signing the user out. Specifies all profile data, now and future. New profile data types may be added to this data kind in the future. This browsing data kind is inclusive of COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_SITE, COREWEBVIEW2_BROWSING_DATA_KINDS_DISK_CACHE, COREWEBVIEW2_BROWSING_DATA_KINDS_DOWNLOAD_HISTORY, COREWEBVIEW2_BROWSING_DATA_KINDS_GENERAL_AUTOFILL, COREWEBVIEW2_BROWSING_DATA_KINDS_PASSWORD_AUTOSAVE, COREWEBVIEW2_BROWSING_DATA_KINDS_BROWSING_HISTORY, and COREWEBVIEW2_BROWSING_DATA_KINDS_SETTINGS.
	COREWEBVIEW2_BROWSING_DATA_KINDS_ALL_PROFILE COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 14
	// COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS
	//
	//	Specifies service workers registered for an origin, and clear will result in termination and deregistration of them.
	//	This is one of the COREWEBVIEW2_BROWSING_DATA_KINDS values.
	COREWEBVIEW2_BROWSING_DATA_KINDS_SERVICE_WORKERS COREWEBVIEW2_BROWSING_DATA_KINDS = 1 << 15
)

// COREWEBVIEW2_TRACKING_PREVENTION_LEVEL
//
//	Tracking prevention levels.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_tracking_prevention_level">See the Globals article.)
type COREWEBVIEW2_TRACKING_PREVENTION_LEVEL = TOleEnum

const (
	// COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_NONE
	//
	//	Tracking prevention is turned off.
	//	This is one of the COREWEBVIEW2_TRACKING_PREVENTION_LEVEL values.
	COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_NONE COREWEBVIEW2_TRACKING_PREVENTION_LEVEL = 0x0000000
	// COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BASIC
	//
	//	This is one of the COREWEBVIEW2_TRACKING_PREVENTION_LEVEL values.
	//	The least restrictive level of tracking prevention. Set to this level to protect against malicious trackers but allows most other trackers and personalize content and ads. See [Current tracking prevention behavior](/microsoft-edge/web-platform/tracking-prevention#current-tracking-prevention-behavior) for fine-grained information on what is being blocked with this level and can change with different Edge versions.
	COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BASIC COREWEBVIEW2_TRACKING_PREVENTION_LEVEL = 0x0000001
	// COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED
	//
	//	The default level of tracking prevention. Set to this level to protect against social media tracking on top of malicious trackers. Content and ads will likely be less personalized. See [Current tracking prevention behavior](/microsoft-edge/web-platform/tracking-prevention#current-tracking-prevention-behavior) for fine-grained information on what is being blocked with this level and can change with different Edge versions.
	//	This is one of the COREWEBVIEW2_TRACKING_PREVENTION_LEVEL values.
	COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_BALANCED COREWEBVIEW2_TRACKING_PREVENTION_LEVEL = 0x0000002
	// COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_STRICT
	//
	//	The most restrictive level of tracking prevention. Set to this level to protect against malicious trackers and most trackers across sites. Content and ads will likely have minimal personalization. This level blocks the most trackers but could cause some websites to not behave as expected. See [Current tracking prevention behavior](/microsoft-edge/web-platform/tracking-prevention#current-tracking-prevention-behavior) for fine-grained information on what is being blocked with this level and can change with different Edge versions.
	//	This is one of the COREWEBVIEW2_TRACKING_PREVENTION_LEVEL values.
	COREWEBVIEW2_TRACKING_PREVENTION_LEVEL_STRICT COREWEBVIEW2_TRACKING_PREVENTION_LEVEL = 0x0000003
)

// COREWEBVIEW2_PDF_TOOLBAR_ITEMS
//
//	PDF toolbar item. This enum must be in sync with ToolBarItem in pdf-store-data-types.ts Specifies the PDF toolbar item types used for the ICoreWebView2Settings.put_HiddenPdfToolbarItems method.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_pdf_toolbar_items">See the Globals article.)
type COREWEBVIEW2_PDF_TOOLBAR_ITEMS = TOleEnum

const (
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_NONE
	//
	//	No item
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_NONE COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000000
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE
	//
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	//	The save button
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000001
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PRINT
	//
	//	The print button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PRINT COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000002
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE_AS
	//
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	//	The save as button
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SAVE_AS COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000004
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_IN
	//
	//	The zoom in button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_IN COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000008
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_OUT
	//
	//	The zoom out button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ZOOM_OUT COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000010
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ROTATE
	//
	//	The rotate button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_ROTATE COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000020
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FIT_PAGE
	//
	//	The fit page button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FIT_PAGE COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000040
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_LAYOUT
	//
	//	The page layout button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_LAYOUT COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000080
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_BOOKMARKS
	//
	//	The bookmarks button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_BOOKMARKS COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000100
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_SELECTOR
	//
	//	The page select button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_PAGE_SELECTOR COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000200
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SEARCH
	//
	//	The search button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_SEARCH COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000400
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FULL_SCREEN
	//
	//	The full screen button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_FULL_SCREEN COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0000800
	// COREWEBVIEW2_PDF_TOOLBAR_ITEMS_MORE_SETTINGS
	//
	//	The more settings button
	//	This is one of the COREWEBVIEW2_PDF_TOOLBAR_ITEMS values.
	COREWEBVIEW2_PDF_TOOLBAR_ITEMS_MORE_SETTINGS COREWEBVIEW2_PDF_TOOLBAR_ITEMS = 0x0001000
)

// COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL
//
//	Specifies memory usage target level of WebView.
//	<see href="https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/webview2-idl#corewebview2_memory_usage_target_level">See the Globals article.)
type COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL = TOleEnum

const (
	// COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_NORMAL
	//
	//	Specifies normal memory usage target level.
	//	This is one of the COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL values.
	COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_NORMAL COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL = 0x0000000
	// COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_LOW
	//
	//	Specifies low memory usage target level. Used for inactivate WebView for reduced memory consumption.
	//	This is one of the COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL values.
	COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL_LOW COREWEBVIEW2_MEMORY_USAGE_TARGET_LEVEL = 0x0000001
)
