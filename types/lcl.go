//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

/*
  注意：Free Pascal中所有集合这里全部使用TSet(uint32)表示，也就是说最多32个元素
*/

// TAlign ENUM
type TAlign = int32

const (
	AlNone TAlign = iota
	AlTop
	AlBottom
	AlLeft
	AlRight
	AlClient
	AlCustom
)

// TAlignSet SET
type TAlignSet = TSet

// TBorderStyle ENUM
type TBorderStyle = int32

const (
	BsNone TBorderStyle = iota
	BsSingle
	BsSizeable
	BsDialog
	BsToolWindow
	BsSizeToolWin
)

// TFormBorderStyle ENUM
type TFormBorderStyle = TBorderStyle

// TFormStyle ENUM
type TFormStyle = int32

const (
	FsNormal TFormStyle = iota
	FsMDIChild
	FsMDIForm
	FsStayOnTop
	FsSplash
	FsSystemStayOnTop
)

// TPosition ENUM
type TPosition = int32

const (
	PoDesigned        TPosition = iota // use bounds from the designer (read from stream)
	PoDefault                          // LCL decision (normally window manager decides)
	PoDefaultPosOnly                   // designed size and LCL position
	PoDefaultSizeOnly                  // designed position and LCL size
	PoScreenCenter                     // center form on screen (depends on DefaultMonitor)
	PoDesktopCenter                    // center form on desktop (total of all screens)
	PoMainFormCenter                   // center form on main form (depends on DefaultMonitor)
	PoOwnerFormCenter                  // center form on owner form (depends on DefaultMonitor)
	PoWorkAreaCenter                   // center form on working area (depends on DefaultMonitor)
)

// TCursor = -32768..32767;
// 相关常量见 lcl/types/cussors.go
type TCursor int16

// TSeekOrigin ENUM
type TSeekOrigin = int32

const (
	SoBeginning TSeekOrigin = iota
	SoCurrent
	SoEnd
)

// TPixelFormat ENUM
type TPixelFormat = int32

const (
	PfDevice TPixelFormat = iota
	Pf1bit
	Pf4bit
	Pf8bit
	Pf15bit
	Pf16bit
	Pf24bit
	Pf32bit
	PfCustom
)

// TAlignment ENUM
type TAlignment = int32

const (
	TaLeftJustify TAlignment = iota
	TaRightJustify
	TaCenter
)

// TLinkAlignment ENUM
type TLinkAlignment = TAlignment

// TLeftRight = TAlignment.taLeftJustify..TAlignment.taRightJustify;
type TLeftRight = int32

// TBiDiMode ENUM
type TBiDiMode = int32

const (
	BdLeftToRight TBiDiMode = iota
	BdRightToLeft
	BdRightToLeftNoAlign
	BdRightToLeftReadingOnly
)

// TVerticalAlignment ENUM
type TVerticalAlignment = int32

const (
	TaAlignTop TVerticalAlignment = iota
	TaAlignBottom
	TaVerticalCenter
)

// TComboBoxStyle ENUM
type TComboBoxStyle = int32

const (
	CsDropDown                  TComboBoxStyle = iota // like an TEdit plus a button to drop down the list, default
	CsSimple                                          // like an TEdit plus a TListBox
	CsDropDownList                                    // like TLabel plus a button to drop down the list
	CsOwnerDrawFixed                                  // like csDropDownList, but custom drawn
	CsOwnerDrawVariable                               // like csDropDownList, but custom drawn and with each item can have another height
	CsOwnerDrawEditableFixed                          // like csOwnerDrawFixed, but with TEdit
	CsOwnerDrawEditableVariable                       // like csOwnerDrawVariable, but with TEdit
)

type TColorBoxStyle = TComboBoxStyle

// TWindowState ENUM
type TWindowState = int32

const (
	WsNormal TWindowState = iota
	WsMinimized
	WsMaximized
	WsFullScreen
)

// TTextLayout ENUM
type TTextLayout = int32

const (
	TlTop TTextLayout = iota
	TlCenter
	TlBottom
)

// TEllipsisPosition ENUM
type TEllipsisPosition = int32

const (
	EpNone TEllipsisPosition = iota
	EpPathEllipsis
	EpEndEllipsis
	EpWordEllipsis
)

// TListBoxStyle ENUM
type TListBoxStyle = int32

const (
	LbStandard TListBoxStyle = iota
	LbOwnerDrawFixed
	LbOwnerDrawVariable
	LbVirtual
	//LbVirtualOwnerDraw
)

// TPopupAlignment ENUM
type TPopupAlignment = int32

const (
	PaLeft TPopupAlignment = iota
	PaRight
	PaCenter
)

// TTrackButton ENUM
type TTrackButton = int32

const (
	TbRightButton TTrackButton = iota
	TbLeftButton
)

// TProgressBarOrientation ENUM
type TProgressBarOrientation = int32

const (
	PbHorizontal TProgressBarOrientation = iota
	PbVertical
	PbRightToLeft
	PbTopDown
)

// TProgressBarStyle ENUM
type TProgressBarStyle = int32

const (
	PbstNormal TProgressBarStyle = iota
	PbstMarquee
)

// TProgressBarState ENUM
type TProgressBarState = int32

const (
	PbsNormal TProgressBarState = iota
	PbsError
	PbsPaused
)

// TButtonLayout ENUM
type TButtonLayout = int32

const (
	BlGlyphLeft TButtonLayout = iota
	BlGlyphRight
	BlGlyphTop
	BlGlyphBottom
)

// TButtonState ENUM
type TButtonState = int32

const (
	BsUp        TButtonState = iota // button is up
	BsDisabled                      // button disabled (grayed)
	BsDown                          // button is down
	BsExclusive                     // button is the only down in his group
	BsHot                           // button is under mouse
)

// TButtonStyle ENUM
type TButtonStyle = int32

const (
	BsAutoDetect TButtonStyle = iota
	BsWin31
	BsNew
)

// TNumGlyphs = 1..4;
type TNumGlyphs = int32

// TStaticBorderStyle ENUM
type TStaticBorderStyle = int32

const (
	SbsNone TStaticBorderStyle = iota
	SbsSingle
	SbsSunken
)

// TFontStyle ENUM
type TFontStyle = int32

const (
	FsBold TFontStyle = iota
	FsItalic
	FsUnderline
	FsStrikeOut
)

// TFontStyles SET TFontStyle
type TFontStyles = TSet

// TFontStylesBase SET TFontStyle
type TFontStylesBase = TSet

// TScrollStyle ENUM
type TScrollStyle = int32

const (
	SsNone TScrollStyle = iota
	SsHorizontal
	SsVertical
	SsBoth
	SsAutoHorizontal
	SsAutoVertical
	SsAutoBoth
)

// TSortType ENUM
type TSortType = int32

const (
	StNone TSortType = iota
	StData
	StText
	StBoth
)

// TListItemsSortType ENUM
type TListItemsSortType = TSortType

// TListArrangement ENUM
type TListArrangement = int32

const (
	ArAlignBottom TListArrangement = iota
	ArAlignLeft
	ArAlignRight
	ArAlignTop
	ArDefault
	ArSnapToGrid
)

// TViewStyle ENUM
type TViewStyle = int32

const (
	VsIcon TViewStyle = iota
	VsSmallIcon
	VsList
	VsReport
)

// TItemState ENUM
type TItemState = int32

const (
	IsNone TItemState = iota
	IsCut
	IsDropHilited
	IsFocused
	IsSelected
	IsActivating
)

// TItemStates SET
type TItemStates = TSet

// TItemChange ENUM
type TItemChange = int32

const (
	CtText TItemChange = iota
	CtImage
	CtState
)

// TItemFind ENUM
type TItemFind = int32

const (
	IfData TItemFind = iota
	IfPartialString
	IfExactString
	IfNearest
)

// TSearchDirection ENUM
type TSearchDirection = int32

const (
	SdLeft TSearchDirection = iota
	SdRight
	SdAbove
	SdBelow
	SdAll
)

// TListHotTrackStyle ENUM
type TListHotTrackStyle = int32

const (
	HtHandPoint TListHotTrackStyle = iota
	HtUnderlineCold
	HtUnderlineHot
)

// TListHotTrackStyles SET TListHotTrackStyle
type TListHotTrackStyles = TSet

// TItemRequests ENUM
type TItemRequests = int32

const (
	IrText TItemRequests = iota
	IrImage
	IrParam
	IrState
	IrIndent
)

// TBrushStyle ENUM
type TBrushStyle = int32

const (
	BsSolid TBrushStyle = iota
	BsClear
	BsHorizontal
	BsVertical
	BsFDiagonal
	BsBDiagonal
	BsCross
	BsDiagCross
	BsImage
	BsPattern
)

// TFPBrushStyle ENUM
type TFPBrushStyle = TBrushStyle

// TPenStyle ENUM
type TPenStyle = int32

const (
	PsSolid TPenStyle = iota
	PsDash
	PsDot
	PsDashDot
	PsDashDotDot
	PsinsideFrame
	PsPattern
	PsClear
)

// TFPPenStyle ENUM
type TFPPenStyle TPenStyle

// TFPPenStyleSet SET TFPPenStyle
type TFPPenStyleSet = TSet

// TUDBtnType ENUM
type TUDBtnType = int32

const (
	BtNext TUDBtnType = iota
	BtPrev
)

// TTabPosition ENUM
type TTabPosition = int32

const (
	TpTop TTabPosition = iota
	TpBottom
	TpLeft
	TpRight
)

// TTabStyle ENUM
type TTabStyle = int32

const (
	TsTabs TTabStyle = iota
	TsButtons
	TsFlatButtons
)

// TFontPitch ENUM
type TFontPitch = int32

const (
	FpDefault TFontPitch = iota
	FpVariable
	FpFixed
)

// TPenMode ENUM
type TPenMode = int32

const (
	PmBlack TPenMode = iota
	PmWhite
	PmNop
	PmNot
	PmCopy
	PmNotCopy
	PmMergePenNot
	PmMaskPenNot
	PmMergeNotPen
	PmMaskNotPen
	PmMerge
	PmNotMerge
	PmMask
	PmNotMask
	PmXor
	PmNotXor
)

// TFPPenMode ENUM
type TFPPenMode TPenMode

// TTrackBarOrientation ENUM
type TTrackBarOrientation = int32

const (
	TrHorizontal TTrackBarOrientation = iota
	TrVertical
)

// TUDOrientation ENUM
type TUDOrientation = int32

const (
	UdHorizontal TUDOrientation = iota
	UdVertical
)

// TFontQuality ENUM
type TFontQuality = int32

const (
	FqDefault TFontQuality = iota
	FqDraft
	FqProof
	FqNonAntialiased
	FqAntialiased
	FqClearType
	FqClearTypeNatural
)

// TCloseAction ENUM
type TCloseAction = int32

const (
	CaNone TCloseAction = iota
	CaHide
	CaFree
	CaMinimize
)

// TBalloonFlags ENUM
type TBalloonFlags = int32

const (
	BfNone TBalloonFlags = iota
	BfInfo
	BfWarning
	BfError
)

// TMsgDlgType ENUM
type TMsgDlgType = int32

const (
	MtWarning TMsgDlgType = iota
	MtError
	MtInformation
	MtConfirmation
	MtCustom
)

// TMsgDlgBtn ENUM
type TMsgDlgBtn = int32

const (
	MbYes TMsgDlgBtn = iota
	MbNo
	MbOK
	MbCancel
	MbAbort
	MbRetry
	MbIgnore
	MbAll
	MbNoToAll
	MbYesToAll
	MbHelp
	MbClose
)

// TMsgDlgButtons TMsgDlgBtn SET
type TMsgDlgButtons = TSet

// TSysLinkType ENUM
type TSysLinkType = int32

const (
	SltURL TSysLinkType = iota
	SltID
)

// TStatusPanelStyle ENUM
type TStatusPanelStyle = int32

const (
	PsText TStatusPanelStyle = iota
	PsOwnerDraw
)

// TJPEGPerformance ENUM
type TJPEGPerformance = int32

const (
	JpBestQuality TJPEGPerformance = iota
	JpBestSpeed
)

type TJPEGPixelFormat = TPixelFormat

//const (
//	Jf24Bit = iota
//	Jf8Bit
//)

type TShortCut uint16

// TNodeState ENUM
type TNodeState = int32

const (
	NsCut              TNodeState = iota // = Node.Cut
	NsDropHilite                         // = Node.DropTarget
	NsFocused                            // = Node.Focused
	NsSelected                           // = Node.Selected
	NsMultiSelected                      // = Node.MultiSelected
	NsExpanded                           // = Node.Expanded
	NsHasChildren                        // = Node.HasChildren
	NsDeleting                           // = Node.Deleting, set on Destroy
	NsVisible                            // = Node.Visible
	NsBound                              // bound to a tree, e.g. has Parent or is top lvl node
	NsValidHasChildren                   // Node.HasChildren has been assigned
)

// TNodeStates SET TNodeState
type TNodeStates = TSet

// TNodeAttachMode ENUM
type TNodeAttachMode = int32

const (
	NaAdd           TNodeAttachMode = iota // add as last sibling of Destination
	NaAddFirst                             // add as first sibling of Destination
	NaAddChild                             // add as last child of Destination
	NaAddChildFirst                        // add as first child of Destination
	NaInsert                               // insert in front of Destination
	NaInsertBehind                         // insert behind Destination
)

// TAddMode ENUM
type TAddMode = int32

const (
	TaAddFirst TAddMode = iota
	TaAdd
	TaInsert
)

// TMultiSelectStyle ENUM
type TMultiSelectStyle = int32

const (
	MsControlSelect TMultiSelectStyle = iota
	MsShiftSelect
	MsVisibleOnly
	MsSiblingOnly
)

// TMultiSelectStyles SET
type TMultiSelectStyles = TSet

// TActionListState ENUM
type TActionListState = int32

const (
	AsNormal TActionListState = iota
	AsSuspended
	AsSuspendedEnabled
)

// TGradientDirection ENUM
type TGradientDirection = int32

const (
	GdHorizontal TGradientDirection = iota
	GdVertical
)

// TDrawingStyle ENUM
type TDrawingStyle = int32

const (
	DSFocus TDrawingStyle = iota
	DSSelected
	DSNormal
	DSTransparent
)

// TImageType ENUM
type TImageType = int32

const (
	ItImage TImageType = iota
	ItMask
)

// TResType ENUM
type TResType = int32

const (
	RtBitmap TResType = iota
	RtCursor
	RtIcon
)

// TLoadResource ENUM
type TLoadResource = int32

const (
	LrDefaultColor TLoadResource = iota
	LrDefaultSize
	LrFromFile
	LrMap3DColors
	LrTransparent
	LrMonoChrome
)

// TLoadResources SET
type TLoadResources = TSet

// TColorDepth ENUM
type TColorDepth = int32

const (
	CdDefault TColorDepth = iota
	CdDeviceDependent
	Cd4Bit
	Cd8Bit
	Cd16Bit
	Cd24Bit
	Cd32Bit
)

// TCheckBoxState ENUM
type TCheckBoxState = int32

const (
	CbUnchecked TCheckBoxState = iota
	CbChecked
	CbGrayed
)

// TToolButtonStyle ENUM
type TToolButtonStyle = int32

const (
	TbsButton     TToolButtonStyle = iota // button (can be clicked)
	TbsCheck                              // check item (click to toggle state, can be grouped)
	TbsDropDown                           // button with dropdown button to show a popup menu
	TbsSeparator                          // space holder
	TbsDivider                            // space holder with line
	TbsButtonDrop                         // button with arrow (not separated from each other)
)

// TTBGradientDrawingOption ENUM
type TTBGradientDrawingOption = int32

const (
	GdoHotTrack TTBGradientDrawingOption = iota
	GdoGradient
)

// TTBGradientDrawingOptions SET
type TTBGradientDrawingOptions = TSet

// ENUM
type TColorDialogOption = int32

const (
	CdFullOpen TColorDialogOption = iota
	CdPreventFullOpen
	CdShowHelp
	CdSolidColor
	CdAnyColor
)

// TColorDialogOptions SET
type TColorDialogOptions = TSet

// TBorderIcon ENUM
type TBorderIcon = int32

const (
	BiSystemMenu TBorderIcon = iota
	BiMinimize
	BiMaximize
	BiHelp
)

// TBorderIcons SET
type TBorderIcons = TSet

// TFontDialogOption ENUM
type TFontDialogOption = int32

const (
	FdAnsiOnly TFontDialogOption = iota
	FdTrueTypeOnly
	FdEffects
	FdFixedPitchOnly
	FdForceFontExist
	FdNoFaceSel
	FdNoOEMFonts
	FdNoSimulations
	FdNoSizeSel
	FdNoStyleSel
	FdNoVectorFonts
	FdShowHelp
	FdWysiwyg
	FdLimitSize
	FdScalableOnly
	FdApplyButton
)

// TFontDialogOptions SET
type TFontDialogOptions = TSet

// TOpenOption ENUM
type TOpenOption = int32

const (
	OfReadOnly        TOpenOption = iota
	OfOverwritePrompt             // if selected file exists shows a message, that file
	// will be overwritten
	OfHideReadOnly // hide read only file
	OfNoChangeDir  // do not change current directory
	OfShowHelp     // show a help button
	OfNoValidate
	OfAllowMultiSelect // allow multiselection
	OfExtensionDifferent
	OfPathMustExist // shows an error message if selected path does not exist
	OfFileMustExist // shows an error message if selected file does not exist
	OfCreatePrompt
	OfShareAware
	OfNoReadOnlyReturn // do not return filenames that are readonly
	OfNoTestFileCreate
	OfNoNetworkButton
	OfNoLongNames
	OfOldStyleDialog
	OfNoDereferenceLinks // do not resolve links while dialog is shown (only on Windows, see OFN_NODEREFERENCELINKS)
	OfNoResolveLinks     // do not resolve links after Execute
	OfEnableIncludeNotify
	OfEnableSizing    // dialog can be resized, e.g. via the mouse
	OfDontAddToRecent // do not add the path to the history list
	OfForceShowHidden // show hidden files
	OfViewDetail      // details are OS and interface dependent
	OfAutoPreview     // details are OS and interface dependent
)

// TOpenOptions SET:TOpenOption
type TOpenOptions = TSet

// TOpenOptionEx ENUM
type TOpenOptionEx = int32

const (
	OfExNoPlacesBar TOpenOptionEx = iota
)

// TOpenOptionsEx SET:TOpenOptionEx
type TOpenOptionsEx = TSet

// TPrintRange ENUM
type TPrintRange = int32

const (
	PrAllPages TPrintRange = iota
	PrSelection
	PrPageNums
	PrCurrentPage
)

// TPrintDialogOption ENUM
type TPrintDialogOption = int32

const (
	PoPrintToFile TPrintDialogOption = iota
	PoPageNums
	PoSelection
	PoWarning
	PoHelp
	PoDisablePrintToFile
	PoBeforeBeginDoc
)

// TPrintDialogOptions SET:TPrintDialogOption
type TPrintDialogOptions = TSet

// TPageSetupDialogOption ENUM
type TPageSetupDialogOption = int32

const (
	PsoDefaultMinMargins TPageSetupDialogOption = iota
	PsoDisableMargins
	PsoDisableOrientation
	PsoDisablePagePainting
	PsoDisablePaper
	PsoDisablePrinter
	PsoMargins
	PsoMinMargins
	PsoShowHelp
	PsoWarning
	PsoNoNetworkButton
)

// TPageSetupDialogOptions SET:TPageSetupDialogOption
type TPageSetupDialogOptions = TSet

// TPrinterKind ENUM
type TPrinterKind = int32

const (
	PkDotMatrix TPrinterKind = iota
	PkHPPCL
)

// TPageType ENUM
type TPageType = int32

const (
	PtEnvelope TPageType = iota
	PtPaper
)

// TPageMeasureUnits ENUM
type TPageMeasureUnits = int32

const (
	PmDefault TPageMeasureUnits = iota
	PmMillimeters
	PmInches
)

// TStringsOption ENUM
type TStringsOption = int32

const (
	SoStrictDelimiter TStringsOption = iota
	SoWriteBOM
	SoTrailingLineBreak
	SoUseLocale
)

// TStringsOptions SET:TStringsOption
type TStringsOptions = TSet

// TShiftStateEnum ENUM
type TShiftStateEnum = int32

const (
	SsShift TShiftStateEnum = iota
	SsAlt
	SsCtrl
	SsLeft
	SsRight
	SsMiddle
	SsDouble
	// Extra additions
	SsMeta
	SsSuper
	SsHyper
	SsAltGr
	SsCaps
	SsNum
	SsScroll
	SsTriple
	SsQuad
	SsExtra1
	SsExtra2
)

// TShiftState SET:TShiftStateEnum
type TShiftState = TSet

// TMouseButton ENUM
type TMouseButton = int32

const (
	MbLeft TMouseButton = iota
	MbRight
	MbMiddle
	MbExtra1
	MbExtra2
)

// TCaptureMouseButtons SET TMouseButton
type TCaptureMouseButtons = TSet

// TFillStyle ENUM
type TFillStyle = int32

const (
	FsSurface TFillStyle = iota
	FsBorder
)

// TFillMode ENUM
type TFillMode = int32

const (
	FmAlternate TFillMode = iota
	FmWinding
)

// TCanvasStates ENUM
type TCanvasStates = int32

const (
	CsHandleValid TCanvasStates = iota
	CsFontValid
	CsPenValid
	CsBrushValid
	CsRegionValid
)

// TCanvasState SET:TCanvasStates
type TCanvasState = TSet

// TCanvasOrientation ENUM
type TCanvasOrientation = int32

const (
	CoLeftToRight TCanvasOrientation = iota
	CoRightToLeft
)

// TTextFormats ENUM
type TTextFormats = int32

const (
	TfBottom TTextFormats = iota
	TfCalcRect
	TfCenter
	TfEditControl
	TfEndEllipsis
	TfPathEllipsis
	TfExpandTabs
	TfExternalLeading
	TfLeft
	TfModifyString
	TfNoClip
	TfNoPrefix
	TfRight
	TfRtlReading
	TfSingleLine
	TfTop
	TfVerticalCenter
	TfWordBreak
	TfHidePrefix
	TfNoFullWidthCharBreak
	TfPrefixOnly
	TfTabStop
	TfWordEllipsis
	TfComposited
)

// TTextFormat SET:TTextFormats
type TTextFormat = TSet

// TBevelCut ENUM
type TBevelCut = int32

const (
	BvNone TBevelCut = iota
	BvLowered
	BvRaised
	BvSpace
)

// TGraphicsBevelCut ENUM
type TGraphicsBevelCut = TBevelCut

// TPanelBevel ENUM
type TPanelBevel = TBevelCut

// TBevelWidth ENUM
type TBevelWidth = int32

// ENUM
type TBevelEdge = int32

const (
	BeLeft TBevelEdge = iota
	BeTop
	BeRight
	BeBottom
)

// TBevelEdges SET:TBevelEdge
type TBevelEdges = TSet

// TBevelKind ENUM
type TBevelKind = int32

const (
	BkNone TBevelKind = iota
	BkTile
	BkSoft
	BkFlat
)

// TTickMark ENUM
type TTickMark = int32

const (
	TmBottomRight TTickMark = iota
	TmTopLeft
	TmBoth
)

// TTickStyle ENUM
type TTickStyle = int32

const (
	TsNone TTickStyle = iota
	TsAuto
	TsManual
)

// TPositionToolTip ENUM
type TPositionToolTip = int32

const (
	PtNone TPositionToolTip = iota
	PtTop
	PtLeft
	PtBottom
	PtRight
)

// TDateTimeKind ENUM
type TDateTimeKind = int32

const (
	DtkDate TDateTimeKind = iota
	DtkTime
	DtkDateTime
)

// TDTDateMode ENUM
type TDTDateMode = int32

const (
	DmComboBox TDTDateMode = iota
	DmUpDown
	DmNone
)

// TDTDateFormat ENUM
type TDTDateFormat = int32

const (
	DfShort TDTDateFormat = iota
	DfLong
)

// TDTCalAlignment ENUM
type TDTCalAlignment = int32

const (
	DtaLeft TDTCalAlignment = iota
	DtaRight
	DtaDefault
)

// TCalDayOfWeek ENUM
type TCalDayOfWeek = int32

const (
	DowMonday TCalDayOfWeek = iota
	DowTuesday
	DowWednesday
	DowThursday
	DowFriday
	DowSaturday
	DowSunday
	DowLocaleDefault
)

// TSearchType ENUM
type TSearchType = int32

const (
	StWholeWord TSearchType = iota
	StMatchCase
)

// TSearchTypes SET: TSearchType
type TSearchTypes = TSet

// TNumberingStyle ENUM
type TNumberingStyle = int32

const (
	NsNone TNumberingStyle = iota
	NsBullte
)

// TAttributeType ENUM
type TAttributeType = int32

const (
	AtSelected TAttributeType = iota
	AtDefaultText
)

// TConsistentAttribute ENUM
type TConsistentAttribute = int32

const (
	CaBold TConsistentAttribute = iota
	CaColor
	CaFace
	CaItalic
	CaSize
	CaStrikeOut
	CaUnderline
	CaProtected
)

// TConsistentAttributes SET: TConsistentAttribute
type TConsistentAttributes = TSet

// TIconArrangement ENUM
type TIconArrangement = int32

const (
	IaTop TIconArrangement = iota
	IaLeft
)

// THeaderStyle ENUM
type THeaderStyle = int32

const (
	HsGradient THeaderStyle = iota
	HsImage
	HsThemed
)

// TImageAlignment ENUM
type TImageAlignment = int32

// IaTop有冲突，所以增加一个i
const (
	IiaLeft TImageAlignment = iota
	IiaRight
	IiaTop
	IiaBottom
	IiaCenter
)

// TAnchorKind ENUM
type TAnchorKind = int32

const (
	AkTop TAnchorKind = iota
	AkLeft
	AkRight
	AkBottom
)

// TAnchors SET: TAnchorKind
type TAnchors = TSet

// TOwnerDrawStateType ENUM
type TOwnerDrawStateType = int32

const (
	OdSelected TOwnerDrawStateType = iota
	OdGrayed
	OdDisabled
	OdChecked
	OdFocused
	OdDefault
	OdHotLight
	OdInactive
	OdNoAccel
	OdNoFocusRect
	OdReserved1
	OdReserved2
	OdComboBoxEdit
	OdBackgroundPainted // item background already painted
)

// TOwnerDrawState SET: TOwnerDrawStateType
type TOwnerDrawState = TSet

// TBitBtnKind ENUM
type TBitBtnKind = int32

const (
	BkCustom TBitBtnKind = iota
	BkOK
	BkCancel
	BkHelp
	BkYes
	BkNo
	BkClose
	BkAbort
	BkRetry
	BkIgnore
	BkAll
	BkNoToAll
	BkYesToAll
)

// TScrollBarKind ENUM
type TScrollBarKind = int32

const (
	SbHorizontal TScrollBarKind = iota
	SbVertical
)

// TScrollBarInc = 1..32767;
type TScrollBarInc int32

// TScrollBarStyle ENUM
type TScrollBarStyle = int32

const (
	SsRegular TScrollBarStyle = iota
	SsFlat
	SsHotTrack
)

// TShapeType ENUM
type TShapeType = int32

const (
	StRectangle TShapeType = iota
	StSquare
	StRoundRect
	StRoundSquare
	StEllipse
	StCircle
	StSquaredDiamond
	StDiamond
	StTriangle
	StTriangleLeft
	StTriangleRight
	StTriangleDown
	StStar
	StStarDown
)

// TBevelStyle = (bsLowered, bsRaised);
type TBevelStyle = int32

const (
	BsLowered TBevelStyle = iota
	BsRaised
)

// TBevelShape ENUM
type TBevelShape = int32

const (
	BsBox TBevelShape = iota
	BsFrame
	BsTopLine
	BsBottomLine
	BsLeftLine
	BsRightLine
	BsSpacer
)

// TGaugeKind ENUM
type TGaugeKind = int32

const (
	GkText TGaugeKind = iota
	GkHorizontalBar
	GkVerticalBar
	GkPie
	GkNeedle
	GkHalfPie
)

// TATGaugeKind ENUM
type TATGaugeKind = TGaugeKind

// TCustomDrawTarget ENUM
type TCustomDrawTarget = int32

const (
	DtControl TCustomDrawTarget = iota
	DtItem
	DtSubItem
)

// TCustomDrawStage ENUM
type TCustomDrawStage = int32

const (
	CdPrePaint TCustomDrawStage = iota
	CdPostPaint
	CdPreErase
	CdPostErase
)

// TCustomDrawStateFlag ENUM
type TCustomDrawStateFlag = int32

const (
	CdsSelected TCustomDrawStateFlag = iota
	CdsGrayed
	CdsDisabled
	CdsChecked
	CdsFocused
	CdsDefault
	CdsHot
	CdsMarked
	CdsIndeterminate
)

// TCustomDrawState SET: TCustomDrawStateFlag
type TCustomDrawState = TSet

// TDisplayCode ENUM
type TDisplayCode = int32

const (
	DrBounds TDisplayCode = iota
	DrIcon
	DrLabel
	DrSelectBounds
)

// TSelectDirOpt ENUM
type TSelectDirOpt = int32

const (
	SdAllowCreate TSelectDirOpt = iota
	SdPerformCreate
	SdPrompt
)

// TSelectDirOpts SET: TSelectDirOpt
type TSelectDirOpts = TSet

// TFindOption ENUM
type TFindOption = int32

const (
	FrDown TFindOption = iota
	FrFindNext
	FrHideMatchCase
	FrHideWholeWord
	FrHideUpDown
	FrMatchCase
	FrDisableMatchCase
	FrDisableUpDown
	FrDisableWholeWord
	FrReplace
	FrReplaceAll
	FrWholeWord
	FrShowHelp
	FrEntireScope
	FrHideEntireScope
	FrPromptOnReplace
	FrHidePromptOnReplace
	FrButtonsAtBottom
)

// TFindOptions SET: TFindOption
type TFindOptions = TSet

// TDragMode ENUM
type TDragMode = int32

const (
	DmManual TDragMode = iota
	DmAutomatic
)

// TDragState ENUM
type TDragState = int32

const (
	DsDragEnter TDragState = iota
	DsDragLeave
	DsDragMove
)

// TDragKind ENUM
type TDragKind = int32

const (
	DkDrag TDragKind = iota
	DkDock
)

// TEditCharCase ENUM
type TEditCharCase = int32

const (
	EcNormal TEditCharCase = iota
	EcUpperCase
	EcLowerCase
)

// TEdgeBorder ENUM
type TEdgeBorder = int32

const (
	EbLeft TEdgeBorder = iota
	EbTop
	EbRight
	EbBottom
)

// TEdgeBorders SET: TEdgeBorder
type TEdgeBorders = TSet

// TEdgeStyle ENUM
type TEdgeStyle = int32

const (
	EsNone TEdgeStyle = iota
	EsRaised
	EsLowered
)

// TGridDrawingStyle ENUM
type TGridDrawingStyle = int32

const (
	GdsClassic TGridDrawingStyle = iota
	GdsThemed
	GdsGradient
)

// TGridOption ENUM
type TGridOption = int32

const (
	GoFixedVertLine TGridOption = iota
	GoFixedHorzLine
	GoVertLine
	GoHorzLine
	GoRangeSelect
	GoDrawFocusSelected
	GoRowSizing
	GoColSizing
	GoRowMoving
	GoColMoving
	GoEditing
	GoAutoAddRows
	GoTabs
	GoRowSelect
	GoAlwaysShowEditor
	GoThumbTracking
	// Additional Options
	GoColSpanning                 // Enable cellextent calcs
	GoRelaxedRowSelect            // User can see focused cell on goRowSelect
	GoDblClickAutoSize            // dblclicking columns borders (on hdrs) resize col.
	GoSmoothScroll                // Switch scrolling mode (pixel scroll is by default)
	GoFixedRowNumbering           // Ya
	GoScrollKeepVisible           // keeps focused cell visible while scrolling
	GoHeaderHotTracking           // Header cells change look when mouse is over them
	GoHeaderPushedLook            // Header cells looks pushed when clicked
	GoSelectionActive             // Setting grid.Selection moves also cell cursor
	GoFixedColSizing              // Allow to resize fixed columns
	GoDontScrollPartCell          // clicking partially visible cells will not scroll
	GoCellHints                   // show individual cell hints
	GoTruncCellHints              // show cell hints if cell text is too long
	GoCellEllipsis                // show "..." if cell text is too long
	GoAutoAddRowsSkipContentCheck //BB Also add a row (if AutoAddRows in Options) if last row is empty
	GoRowHighlight                // Highlight the current Row
)

// TGridOptions SET: TGridOption
type TGridOptions = TSet

// TGridDrawStates ENUM
type TGridDrawStates = int32

const (
	GdSelected TGridDrawStates = iota
	GdFocused
	GdFixed
	GdHot
	GdPushed
	GdRowHighlight
)

// TGridDrawState SET: TGridDrawStates
type TGridDrawState = TSet

// THeaderSectionStyle ENUM
type THeaderSectionStyle = int32

const (
	HsText THeaderSectionStyle = iota
	HsOwnerDraw
)

// TLabelPosition ENUM
type TLabelPosition = int32

const (
	LpAbove TLabelPosition = iota
	LpBelow
	LpLeft
	LpRight
)

// TFlowStyle ENUM
type TFlowStyle = int32

const (
	FsLeftRightTopBottom TFlowStyle = iota
	FsRightLeftTopBottom
	FsLeftRightBottomTop
	FsRightLeftBottomTop
	FsTopBottomLeftRight
	FsBottomTopLeftRight
	FsTopBottomRightLeft
	FsBottomTopRightLeft
)

// TCoolBandMaximize ENUM
type TCoolBandMaximize = int32

const (
	BmNone TCoolBandMaximize = iota
	BmClick
	BmDblClick
)

// TMenuBreak ENUM
type TMenuBreak = int32

const (
	MbNone TMenuBreak = iota
	MbBreak
	MbBarBreak
)

// TSectionTrackState ENUM
type TSectionTrackState = int32

const (
	TsTrackBegin TSectionTrackState = iota
	TsTrackMove
	TsTrackEnd
)

// TControlStateType ENUM
type TControlStateType = int32

const (
	CsLButtonDown TControlStateType = iota
	CsClicked
	CsPalette
	CsReadingState
	CsFocusing
	CsCreating // not used, exists for Delphi compatibility
	CsPaintCopy
	CsCustomPaint
	CsDestroyingHandle
	CsDocking
	CsVisibleSetInLoading
)

// TControlState SET: TControlStateType
type TControlState = TSet

// TControlStyleType ENUM
type TControlStyleType = int32

const (
	CsAcceptsControls            TControlStyleType = iota // can have children in the designer
	CsCaptureMouse                                        // auto capture mouse when clicked
	CsDesignInteractive                                   // wants mouse events in design mode
	CsClickEvents                                         // handles mouse events
	CsFramed                                              // not implemented, has 3d frame
	CsSetCaption                                          // if Name=Caption, changing the Name changes the Caption
	CsOpaque                                              // the control paints its area completely
	CsDoubleClicks                                        // understands mouse double clicks
	CsTripleClicks                                        // understands mouse triple clicks
	CsQuadClicks                                          // understands mouse quad clicks
	CsFixedWidth                                          // cannot change its width
	CsFixedHeight                                         // cannot change its height (for example combobox)
	CsNoDesignVisible                                     // is invisible in the designer
	CsReplicatable                                        // PaintTo works
	CsNoStdEvents                                         // standard events such as mouse, key, and click events are ignored.
	CsDisplayDragImage                                    // display images from dragimagelist during drag operation over control
	CsReflector                                           // not implemented, the controls respond to size, focus and dlg messages - it can be used as ActiveX control under Windows
	CsActionClient                                        // Action is set
	CsMenuEvents                                          // not implemented
	CsNoFocus                                             // control will not take focus when clicked with mouse.
	CsNeedsBorderPaint                                    // not implemented
	CsParentBackground                                    // tells WinXP to paint the theme background of parent on controls background
	CsDesignNoSmoothResize                                // when resizing control in the designer do not SetBounds while dragging
	CsDesignFixedBounds                                   // can not be moved nor resized in designer
	CsHasDefaultAction                                    // implements useful ExecuteDefaultAction
	CsHasCancelAction                                     // implements useful ExecuteCancelAction
	CsNoDesignSelectable                                  // can not be selected at design time
	CsOwnedChildrenNotSelectable                          // child controls owned by this control are NOT selectable in the designer
	CsAutoSize0x0                                         // if the preferred size is 0x0 then control is shrinked ot 0x0
	CsAutoSizeKeepChildLeft                               // when AutoSize=true do not move children horizontally
	CsAutoSizeKeepChildTop                                // when AutoSize=true do not move children vertically
	CsRequiresKeyboardInput                               // If the device has no physical keyboard then show the virtual keyboard when this control gets focus (therefore available only to TWinControl descendents)
)

// TControlStyle SET: TControlStyleType
type TControlStyle = TSet

// TMouseActivate ENUM
type TMouseActivate = int32

const (
	MaDefault TMouseActivate = iota
	MaActivate
	MaActivateAndEat
	MaNoActivate
	MaNoActivateAndEat
)

// TTaskBarProgressState ENUM
type TTaskBarProgressState = int32

const (
	None TTaskBarProgressState = iota
	Indeterminate
	Normal
	Error
	Paused
)

// TBitmapHandleType ENUM
type TBitmapHandleType = int32

const (
	BmDIB TBitmapHandleType = iota
	BmDDB
)

// TPrinterState ENUM
type TPrinterState = int32

const (
	PsNoDefine TPrinterState = iota
	PsReady
	PsPrinting
	PsStopped
)

// TPrinterOrientation ENUM
type TPrinterOrientation = int32

const (
	PoPortrait TPrinterOrientation = iota
	PoLandscape
	PoReverseLandscape
	PoReversePortrait
)

// TPrinterCapability ENUM
type TPrinterCapability = int32

const (
	PcCopies TPrinterCapability = iota
	PcOrientation
	PcCollation
)

// TPrinterCapabilities SET: TPrinterCapability
type TPrinterCapabilities = TSet

// TPrinterType ENUM
type TPrinterType = int32

const (
	PtLocal TPrinterType = iota
	PtNetWork
)

// TReadyState ENUM
type TReadyState = int32

const (
	RsUninitialized TReadyState = iota
	RsLoading
	RsLoaded
	RsInterActive
	RsComplete
)

// TStringEncoding ENUM
type TStringEncoding = int32

const (
	SeUnknown TStringEncoding = iota
	SeANSI
	SeUnicode
	SeUTF8
)

// TShowInTaskBar ENUM
type TShowInTaskBar = int32

const (
	StDefault TShowInTaskBar = iota // use default rules for showing taskbar item
	StAlways                        // always show taskbar item for the form
	StNever                         // never show taskbar item for the form
)

// TTaskDialogCommonButton ENUM
type TTaskDialogCommonButton = int32

const (
	TcbOk TTaskDialogCommonButton = iota
	TcbYes
	TcbNo
	TcbCancel
	TcbRetry
	TcbClose
)

// TTaskDialogCommonButtons SET: TTaskDialogCommonButton
type TTaskDialogCommonButtons = TSet

// TTaskDialogFlag ENUM
type TTaskDialogFlag = int32

const (
	TfEnableHyperlinks TTaskDialogFlag = iota
	TfUseHiconMain
	TfUseHiconFooter
	TfAllowDialogCancellation
	TfUseCommandLinks
	TfUseCommandLinksNoIcon
	TfExpandFooterArea
	TfExpandedByDefault
	TfVerificationFlagChecked
	TfShowProgressBar
	TfShowMarqueeProgressBar
	TfCallbackTimer
	TfPositionRelativeToWindow
	TfRtlLayout
	TfNoDefaultRadioButton
	TfCanBeMinimized
)

// TTaskDialogFlags SET: TTaskDialogFlag
type TTaskDialogFlags = TSet

// TTaskDialogIcon ENUM
type TTaskDialogIcon = int32

const (
	TdiNone TTaskDialogIcon = iota
	TdiWarning
	TdiError
	TdiInformation
	TdiShield
	TdiQuestion
)

// TComboBoxExStyle ENUM
type TComboBoxExStyle = int32

const (
	CsExDropDown TComboBoxExStyle = iota
	CsExSimple
	CsExDropDownList
)

// TComboBoxExStyleEx ENUM
type TComboBoxExStyleEx = int32

const (
	CsExCaseSensitive TComboBoxExStyleEx = iota
	CsExNoEditImage
	CsExNoEditImageIndent
	CsExNoSizeLimit
	CsExPathWordBreak
)

// TComboBoxExStyles SET: TComboBoxExStyleEx
type TComboBoxExStyles = TSet

// TAutoCompleteOption ENUM
type TAutoCompleteOption = int32

const (
	AcoAutoSuggest TAutoCompleteOption = iota
	AcoAutoAppend
	AcoSearch
	AcoFilterPrefixes
	AcoUseTab
	AcoUpDownKeyDropsList
	AcoRtlReading
)

// TAutoCompleteOptions SET: TAutoCompleteOption
type TAutoCompleteOptions = TSet

// TDefaultMonitor ENUM
type TDefaultMonitor = int32

const (
	DmDesktop TDefaultMonitor = iota
	DmPrimary
	DmMainForm
	DmActiveForm
)

// TTransparentMode ENUM
type TTransparentMode = int32

const (
	TmAuto TTransparentMode = iota
	TmFixed
)

// TDrawImageMode ENUM
type TDrawImageMode = int32

const (
	DimNormal TDrawImageMode = iota
	DimCenter
	DimStretch
)

// TListBoxOption ENUM
type TListBoxOption = int32

const (
	LboDrawFocusRect TListBoxOption = iota // draw focus rect in case of owner drawing
)

// TListBoxOptions SET: TListBoxOption
type TListBoxOptions = TSet

// TAntialiasingMode ENUM
type TAntialiasingMode = int32

const (
	AmDontCare TAntialiasingMode = iota // default antialiasing
	AmOn                                // enabled
	AmOff                               // disabled
)

// TSortDirection ENUM
type TSortDirection = int32

const (
	SdAscending TSortDirection = iota
	SdDescending
)

// TTreeViewExpandSignType ENUM
type TTreeViewExpandSignType = int32

const (
	TvestTheme     TTreeViewExpandSignType = iota // use themed sign
	TvestPlusMinus                                // use +/- sign
	TvestArrow                                    // use blank arrow
	TvestArrowFill                                // use filled arrow
)

// TTreeViewOption ENUM
type TTreeViewOption = int32

const (
	TvoAllowMultiselect TTreeViewOption = iota
	TvoAutoExpand
	TvoAutoInsertMark
	TvoAutoItemHeight
	TvoHideSelection
	TvoHotTrack
	TvoKeepCollapsedNodes
	TvoReadOnly
	TvoRightClickSelect
	TvoRowSelect
	TvoShowButtons
	TvoShowLines
	TvoShowRoot
	TvoShowSeparators
	TvoToolTips
	TvoNoDoubleClickExpand
	TvoThemedDraw
)

// TTreeViewOptions SET: TTreeViewOption
type TTreeViewOptions = TSet

// TGlyphShowMode ENUM
type TGlyphShowMode = int32

const (
	GsmAlways      TGlyphShowMode = iota // always show
	GsmNever                             // never show
	GsmApplication                       // depends on application settings
	GsmSystem                            // depends on system settings
)

// TCTabControlOption ENUM
type TCTabControlOption = int32

const (
	NboShowCloseButtons TCTabControlOption = iota
	NboMultiLine
	NboHidePageListPopup
	NboKeyboardTabSwitch
	NboShowAddTabButton
	NboDoChangeOnSetIndex
)

// TCTabControlOptions SET: TCTabControlOption
type TCTabControlOptions = TSet

// TAnchorSideReference ENUM
type TAnchorSideReference = int32

const (
	AsrTop TAnchorSideReference = iota
	AsrBottom
	AsrCenter
)

// TControlCellAlign ENUM
type TControlCellAlign = int32

const (
	CcaFill TControlCellAlign = iota
	CcaLeftTop
	CcaRightBottom
	CcaCenter
)

// TControlCellAligns SET: TControlCellAlign
type TControlCellAligns = TSet

// TChildControlResizeStyle ENUM
type TChildControlResizeStyle = int32

const (
	CrsAnchorAligning        TChildControlResizeStyle = iota // (like Delphi)
	CrsScaleChilds                                           // scale children equally, keep space between children fixed
	CrsHomogenousChildResize                                 // enlarge children equally (i.e. by the same amount of pixel)
	CrsHomogenousSpaceResize                                 // enlarge space between children equally
	//{$IFDEF EnablecrsSameSize}
	//,CrsSameSize  // each child gets the same size (maybe one pixel difference)
	//{$ENDIF}
)

// TControlChildrenLayout ENUM
type TControlChildrenLayout = int32

const (
	CclNone                       TControlChildrenLayout = iota
	CclLeftToRightThenTopToBottom                        // if BiDiMode <> bdLeftToRight then it becomes RightToLeft
	CclTopToBottomThenLeftToRight
)

// TColumnLayout ENUM
type TColumnLayout = int32

const (
	ClHorizontalThenVertical TColumnLayout = iota
	ClVerticalThenHorizontal
)

// TSortIndicator ENUM
type TSortIndicator = int32

const (
	SiNone TSortIndicator = iota
	SiAscending
	SiDescending
)

// TLibType VCL或者LCL，只是用于引入的
type TLibType = int32

const (
	LtVCL TLibType = iota
	LtLCL
)

// TColumnButtonStyle ENUM
type TColumnButtonStyle = int32

const (
	CbsAuto TColumnButtonStyle = iota
	CbsEllipsis
	CbsNone
	CbsPickList
	CbsCheckboxColumn
	CbsButton
	CbsButtonColumn
)

// TGridZone ENUM
type TGridZone = int32

const (
	GzNormal TGridZone = iota
	GzFixedCols
	GzFixedRows
	GzFixedCells
	GzInvalid
)

// TGridZoneSet SET: TGridZone
type TGridZoneSet = TSet

// TSortOrder ENUM
type TSortOrder = int32

const (
	SoAscending TSortOrder = iota
	SoDescending
)

// TAutoAdvance ENUM
type TAutoAdvance = int32

const (
	AaNone TAutoAdvance = iota
	AaDown
	AaRight
	AaLeft
	AaRightDown
	AaLeftDown
	AaRightUp
	AaLeftUp
)

// TCellHintPriority ENUM
type TCellHintPriority = int32

const (
	ChpAll TCellHintPriority = iota
	ChpAllNoDefault
	ChpTruncOnly
)

// TMouseWheelOption ENUM
type TMouseWheelOption = int32

const (
	MwCursor TMouseWheelOption = iota
	MwGrid
)

// TGridOption2 ENUM
type TGridOption2 = int32

const (
	GoScrollToLastCol TGridOption2 = iota // allow scrolling to last column (so that last column can be leftcol)
	GoScrollToLastRow                     // allow scrolling to last row (so that last row can be toprow)
)

// TGridOptions2 SET: TGridOption2
type TGridOptions2 = TSet

// TRangeSelectMode ENUM
type TRangeSelectMode = int32

const (
	RsmSingle TRangeSelectMode = iota
	RsmMulti
)

// TTitleStyle ENUM
type TTitleStyle = int32

const (
	TsLazarus TTitleStyle = iota
	TsStandard
	TsNative
)

// TPrefixOption ENUM
type TPrefixOption = int32

const (
	PoNone TPrefixOption = iota
	PoHeaderClick
)

// TDisplaySetting ENUM
type TDisplaySetting = int32

const (
	DsShowHeadings TDisplaySetting = iota
	DsShowDayNames
	DsNoMonthChange
	DsShowWeekNumbers
	DsStartMonday
)

// TDisplaySettings SET TDisplaySetting
type TDisplaySettings = TSet

// TTimeFormat ENUM
type TTimeFormat = int32

const (
	Tf12 TTimeFormat = iota // 12 hours format, with am/pm string
	Tf24                    // 24 hours format
)

// TTimeDisplay ENUM
type TTimeDisplay = int32

const (
	TdHM    TTimeDisplay = iota // hour and minute
	TdHMS                       // hour Minute and second
	TdHMSMs                     // hour Minute Second and milisecond
)

// TArrowShape ENUM
type TArrowShape = int32

const (
	AsClassicSmaller TArrowShape = iota
	AsClassicLarger
	AsModernSmaller
	AsModernLarger
	AsYetAnotherShape
	AsTheme
)

// TDateDisplayOrder ENUM
type TDateDisplayOrder = int32

const (
	DdoDMY TDateDisplayOrder = iota
	DdoMDY
	DdoYMD
	DdoTryDefault
)

// TDateTimePart ENUM
type TDateTimePart = int32

const (
	DtpDay TDateTimePart = iota
	DtpMonth
	DtpYear
	DtpHour
	DtpMinute
	DtpSecond
	DtpMiliSec
	DtpAMPM
)

// TDateTimeParts SET: TDateTimePart
type TDateTimeParts = TSet

// TDateTimePickerOption ENUM
type TDateTimePickerOption = int32

const (
	DtpoDoChangeOnSetDateTime TDateTimePickerOption = iota
	DtpoEnabledIfUnchecked
	DtpoAutoCheck
	DtpoFlatButton
)

// TDateTimePickerOptions SET: TDateTimePickerOption
type TDateTimePickerOptions = TSet

// TImageOrientation ENUM
type TImageOrientation = int32

const (
	IoHorizontal TImageOrientation = iota
	IoVertical
)

// TLayoutAdjustmentPolicy ENUM
type TLayoutAdjustmentPolicy = int32

const (
	LapDefault                              TLayoutAdjustmentPolicy = iota // widgetset dependent
	LapFixedLayout                                                         // A fixed absolute layout in all platforms
	LapAutoAdjustWithoutHorizontalScrolling                                // Smartphone platforms use this one,
	// the x axis is stretched to fill the screen and
	// the y is scaled to fit the DPI
	LapAutoAdjustForDPI // For desktops using High DPI, scale x and y to fit the DPI
)

// THitTest ENUM
type THitTest = int32

const (
	HtAbove THitTest = iota
	HtBelow
	HtNowhere
	HtOnItem
	HtOnButton
	HtOnIcon
	HtOnIndent
	HtOnLabel
	HtOnRight
	HtOnStateIcon
	HtToLeft
	HtToRight
)

// THitTests SET
type THitTests = TSet

// TListItemState ENUM
type TListItemState = int32

const (
	LisCut TListItemState = iota
	LisDropTarget
	LisFocused
	LisSelected
)

// TListItemStates SET TListItemState
type TListItemStates = TSet

// TPredefinedClipboardFormat ENUM
type TPredefinedClipboardFormat = int32

const (
	PcfText TPredefinedClipboardFormat = iota
	PcfBitmap
	PcfPixmap
	PcfIcon
	PcfPicture
	PcfMetaFilePict
	PcfObject
	PcfComponent
	PcfCustomData
)

// TWrapAfter ENUM
type TWrapAfter = int32

const (
	WaAuto   TWrapAfter = iota // auto
	WaForce                    // always wrap after this control
	WaAvoid                    // try not to wrap after this control, if the control is already at the beginning of the row, wrap though
	WaForbid                   // never wrap after this control
)

// TGraphicsDrawEffect ENUM
type TGraphicsDrawEffect = int32

const (
	GdeNormal      TGraphicsDrawEffect = iota // no effect
	GdeDisabled                               // grayed image
	GdeHighlighted                            // a bit highlighted image
	GdeShadowed                               // a bit shadowed image
	Gde1Bit                                   // 1 Bit image (for non-XP windows buttons)
)

// TLazAccessibilityRole ENUM
type TLazAccessibilityRole = int32

const (
	LarIgnore               TLazAccessibilityRole = iota // Default value. Something to be ignored. For example a blank space between other objects.
	LarAnimation                                         // An object that displays an animation.
	LarButton                                            // A button.
	LarCell                                              // A cell in a table.
	LarChart                                             // An object that displays a graphical representation of data.
	LarCheckBox                                          // An object that can be checked or unchecked, or sometimes in an intermediary state
	LarClock                                             // A clock displaying time.
	LarColorPicker                                       // A control which allows selecting a color.
	LarColumn                                            // A generic column that goes in a table.
	LarComboBox                                          // A list of choices that the user can select from.
	LarDateField                                         // A controls which displays and possibly allows one to choose a date.
	LarGrid                                              // A grid control which displays cells
	LarGroup                                             // A control which groups others, such as a TGroupBox.
	LarImage                                             // A graphic or picture or an icon.
	LarLabel                                             // A text label as usually placed near other widgets.
	LarListBox                                           // A list of items, from which the user can select one or more items.
	LarListItem                                          // An item in a list of items.
	LarMenuBar                                           // A main menu bar.
	LarMenuItem                                          // A item in a menu.
	LarProgressIndicator                                 // A control which shows a progress indication.
	LarRadioButton                                       // A radio button, see for example TRadioButton.
	LarResizeGrip                                        // A grip that the user can drag to change the size of widgets.
	LarRow                                               // A generic row that goes in a table.
	LarScrollBar                                         // A control to scroll another one
	LarSpinner                                           // A control which allows one to increment / decrement a value.
	LarTabControl                                        // A control with tabs, like TPageControl.
	LarText                                              // Text inside of a control, like text in a row cell
	LarTextEditorMultiline                               // A multi-line text editor (for example: TMemo, SynEdit)
	LarTextEditorSingleline                              // A single-line text editor (for example: TEdit)
	LarToolBar                                           // A control that holds ToolButtons
	LarToolBarButton                                     // A button on a ToolBar
	LarTrackBar                                          // A control which allows one to drag a slider.
	LarTreeView                                          // A list of items in a tree structure.
	LarTreeItem                                          // An item in a tree structure.
	LarUnknown                                           // An item that doesn't fit any of the other categories.
	LarWindow                                            // A top level window.
)

// THelpContext ENUM
type THelpContext = int32

// TApplicationType ENUM
type TApplicationType = int32

const (
	AtDefault        TApplicationType = iota // The widgetset will attempt to auto-detect the device type
	AtDesktop                                // For common desktops and notebooks
	AtPDA                                    // For smartphones and other devices with touch screen and a small screen
	AtKeyPadDevice                           // Devices without any pointing device, such as keypad feature phones or kiosk machines
	AtTablet                                 // Similar to a PDA/Smartphone, but with a large screen
	AtTV                                     // The device is a television
	AtMobileEmulator                         // For desktop platforms. It will create a main windows of 240x320 and place all forms there to immitate a mobile platform
)

// TApplicationDoubleBuffered ENUM
// what Forms.DoubleBuffered with ParentDoubleBuffered=True will gain when created
type TApplicationDoubleBuffered = int32

const (
	AdbDefault TApplicationDoubleBuffered = iota // widgetset dependent (LCLWin32: True unless in remote desktop connection; other WSs: False)
	AdbFalse                                     // False
	AdbTrue                                      // True
)

// TApplicationExceptionDlg ENUM
type TApplicationExceptionDlg = int32

const (
	AedOkCancelDialog TApplicationExceptionDlg = iota // Exception handler window will be a dialog with Ok/Cancel buttons
	AedOkMessageBox                                   // Exception handler window will be a simple message box
)

// TApplicationFlag ENUM
type TApplicationFlag = int32

const (
	AppWaiting TApplicationFlag = iota
	AppIdleEndSent
	AppNoExceptionMessages
	AppActive // application has focus
	AppDestroying
	AppDoNotCallAsyncQueue
	AppInitialized // initialization of application was done
)

// TApplicationFlags SET TApplicationFlag
type TApplicationFlags = TSet

// TApplicationNavigationOption ENUM
type TApplicationNavigationOption = int32

const (
	AnoTabToSelectNext TApplicationNavigationOption = iota
	AnoReturnForDefaultControl
	AnoEscapeForCancelControl
	AnoF1ForHelp
	AnoArrowToSelectNextInParent
)

// TApplicationNavigationOptions SET TApplicationNavigationOption
type TApplicationNavigationOptions = TSet

// TGraphicsFillStyle ENUM
type TGraphicsFillStyle = int32

const (
	GfsSurface TGraphicsFillStyle = iota // fill till the color (it fills all except this color)
	GfsBorder                            // fill this color (it fills only connected pixels of this color)
)

// TTaskBarBehavior ENUM
type TTaskBarBehavior = int32

const (
	TbDefault      TTaskBarBehavior = iota // widgetset dependent
	TbMultiButton                          // show buttons for Forms with ShowTaskBar = stDefault
	TbSingleButton                         // hide buttons for Forms with ShowTaskBar = stDefault.
	// Some Linux window managers do not support it. For example Cinnamon.
)

// TApplicationShowGlyphs ENUM
type TApplicationShowGlyphs = int32

const (
	SbgAlways TApplicationShowGlyphs = iota // show them always (default)
	SbgNever                                // show them never
	SbgSystem                               // show them depending on OS
)

// TATButtonOverlayPosition ENUM
type TATButtonOverlayPosition = int32

const (
	BopLeftTop TATButtonOverlayPosition = iota
	BopRightTop
	BopLeftBottom
	BopRightBottom
)

// TFPDrawingMode ENUM
type TFPDrawingMode = int32

const (
	DmOpaque TFPDrawingMode = iota
	DmAlphaBlend
	DmCustom
)

// TFPPenEndCap ENUM
type TFPPenEndCap = int32

const (
	PecRound TFPPenEndCap = iota
	PecSquare
	PecFlat
)

// TPenEndCap ENUM
type TPenEndCap = TFPPenEndCap

// TFPPenJoinStyle ENUM
type TFPPenJoinStyle = int32

const (
	PjsRound TFPPenJoinStyle = iota
	PjsBevel
	PjsMiter
)

// TPenJoinStyle ENUM
type TPenJoinStyle = TFPPenJoinStyle

// THeaderSectionState ENUM
type THeaderSectionState = int32

const (
	HsNormal THeaderSectionState = iota
	HsHot
	HsPressed
)

// TJPEGScale ENUM
type TJPEGScale = int32

const (
	JsFullSize TJPEGScale = iota
	JsHalf
	JsQuarter
	JsEighth
)

// TFormStateType ENUM
type TFormStateType = int32

const (
	FsCreating           TFormStateType = iota // initializing (form streaming)
	FsVisible                                  // form should be shown
	FsShowing                                  // form handling WM_SHOWWINDOW message
	FsModal                                    // form is modal
	FsCreatedMDIChild                          // todo: not mplemented
	FsBorderStyleChanged                       // border style is changed before window handle creation
	FsFormStyleChanged                         // form style is changed before window handle creation
	FsFirstShow                                // form is shown for the first time
	FsDisableAutoSize                          // disable autosize
)

// TFormState SET TFormStateType
type TFormState = TSet

// TPopupMode ENUM
type TPopupMode = int32

const (
	PmNone     TPopupMode = iota // modal: popup to active form or if not available, to main form; non-modal: no window parent
	PmAuto                       // modal & non-modal: popup to active form or if not available, to main form
	PmExplicit                   // modal & non-modal: popup to PopupParent or if not available, to main form
)

// TGridCursorState ENUM
type TGridCursorState = int32

const (
	GcsDefault TGridCursorState = iota
	GcsColWidthChanging
	GcsRowHeightChanging
	GcsDragging
)

// TIniFileOption ENUM
type TIniFileOption = int32

const (
	IfoStripComments        TIniFileOption = iota // Strip comments when reading file
	IfoStripInvalid                               // Strip invalid lines when reading file.
	IfoEscapeLineFeeds                            // Escape linefeeds when reading file.
	IfoCaseSensitive                              // Use Case sensitive section/key names
	IfoStripQuotes                                // Strip quotes when reading string values.
	IfoFormatSettingsActive                       // Use format settings when writing date/float etc.
	IfoWriteStringBoolean                         // Write booleans as string
)

// TIniFileOptions SET TIniFileOption
type TIniFileOptions = TSet

// TResizeStyle ENUM
type TResizeStyle = int32

const (
	RsLine    TResizeStyle = iota // draw a line, don't update splitter position during moving
	RsNone                        // draw nothing and don't update splitter position during moving
	RsPattern                     // draw a dot pattern, don't update splitter position during moving
	RsUpdate                      // draw nothing, update splitter position during moving
)

// TTrackBarScalePos ENUM
type TTrackBarScalePos = int32

const (
	TrLeft TTrackBarScalePos = iota
	TrRight
	TrTop
	TrBottom
)

// TTreeViewInsertMarkType ENUM
type TTreeViewInsertMarkType = int32

const (
	TvimNone         TTreeViewInsertMarkType = iota
	TvimAsFirstChild                         // or as root
	TvimAsNextSibling
	TvimAsPrevSibling
)

// TMonthDisplay ENUM
type TMonthDisplay = int32

const (
	MdShort TMonthDisplay = iota
	MdLong
	MdCustom
)

// TDockOrientation ENUM
type TDockOrientation = int32

const (
	DoNoOrient   TDockOrientation = iota // zone contains a TControl and no child zones.
	DoHorizontal                         // zone's children are stacked top-to-bottom.
	DoVertical                           // zone's children are arranged left-to-right.
	DoPages                              // zone's children are pages arranged left-to-right.
)

// TListItemsFlag ENUM
type TListItemsFlag = int32

const (
	LisfWSItemsCreated TListItemsFlag = iota
)

// TListItemsFlags SET TListItemsFlag
type TListItemsFlags = TSet

// TMaskEditValidationErrorMode ENUM
type TMaskEditValidationErrorMode = int32

const (
	MvemException TMaskEditValidationErrorMode = iota
	MvemEvent
)

// TSizeConstraintsOption ENUM
type TSizeConstraintsOption = int32

const (
	ScoAdviceWidthAsMin TSizeConstraintsOption = iota
	ScoAdviceWidthAsMax
	ScoAdviceHeightAsMin
	ScoAdviceHeightAsMax
)

// TSizeConstraintsOptions SET  TSizeConstraintsOption
type TSizeConstraintsOptions = TSet

// TStatusPanelBevel ENUM
type TStatusPanelBevel = int32

const (
	PbNone TStatusPanelBevel = iota
	PbLowered
	PbRaised
)

// TDuplicates ENUM
type TDuplicates = int32

const (
	DupIgnore TDuplicates = iota
	DupAccept
	DupError
)

// TStringsSortStyle ENUM
type TStringsSortStyle = int32

const (
	SslNone TStringsSortStyle = iota
	SslUser
	SslAuto
)

// TStringsSortStyles SET TStringsSortStyle
type TStringsSortStyles = TSet

// TMissingNameValueSeparatorAction ENUM
type TMissingNameValueSeparatorAction = int32

const (
	MnvaValue TMissingNameValueSeparatorAction = iota
	MnvaName
	MnvaEmpty
	MnvaError
)

// TMissingNameValueSeparatorActions SET TMissingNameValueSeparatorAction
type TMissingNameValueSeparatorActions = TSet

// TTextLineBreakStyle ENUM
type TTextLineBreakStyle = int32

const (
	TlbsLF TTextLineBreakStyle = iota
	TlbsCRLF
	TlbsCR
)

// TUDAlignButton ENUM
type TUDAlignButton = int32

const (
	udLeft TUDAlignButton = iota
	udRight
	udTop
	udBottom
)

// TDisplayOption ENUM
type TDisplayOption = int32

const (
	DoColumnTitles TDisplayOption = iota
	DoAutoColResize
	DoKeyColFixed
)

// TDisplayOptions SET TDisplayOption
type TDisplayOptions = TSet

// TKeyOption ENUM
type TKeyOption = int32

const (
	KeyEdit TKeyOption = iota
	KeyAdd
	KeyDelete
	KeyUnique
)

// TKeyOptions SET TKeyOption
type TKeyOptions = TSet

// TClipboardType ENUM
type TClipboardType = int32

const (
	CtPrimarySelection TClipboardType = iota
	CtSecondarySelection
	CtClipboard
)

// TComponentState ENUM
type TComponentState = int32

const (
	CsLoading TComponentState = iota
	CsReading
	CsWriting
	CsDestroying
	CsDesigning
	CsAncestor
	CsUpdating
	CsFixups
	CsFreeNotification
	CsInline
	CsDesignInstance
)

// TComponentStates SET TComponentState
type TComponentStates = TSet

// TComponentStyle ENUM
type TComponentStyle = int32

const (
	CsInheritable TComponentStyle = iota
	CsCheckPropAvail
	CsSubComponent
	CsTransient
)

// TComponentStyles SET TComponentStyle
type TComponentStyles = TSet

// THelpType ENUM
type THelpType = int32

const (
	htKeyword THelpType = iota
	htContext
)

// TEventType ENUM
type TEventType = int32

const (
	EtCustom TEventType = iota
	EtInfo
	EtWarning
	EtError
	EtDebug
)

// TEventTypes SET TEventType
type TEventTypes = TSet

// TEventLogTypes SET TEventType
type TEventLogTypes = TSet

// TComboBoxAutoCompleteTextOption ENUM
type TComboBoxAutoCompleteTextOption = int32

const (
	CbactEnabled             TComboBoxAutoCompleteTextOption = iota //Enable Auto-Completion Feature
	CbactEndOfLineComplete                                          //Perform Auto-Complete only when cursor is at end of line
	CbactRetainPrefixCase                                           //Retains the case of characters user has typed if is cbactEndOfLineComplete
	CbactSearchCaseSensitive                                        //Search Text with CaseSensitivity
	CbactSearchAscending                                            //Search Text from top of the list
)

// TComboBoxAutoCompleteText SET TComboBoxAutoCompleteTextOption
type TComboBoxAutoCompleteText = TSet

// TEmulatedTextHintStatus ENUM
type TEmulatedTextHintStatus = int32

const (
	ThsHidden TEmulatedTextHintStatus = iota
	ThsShowing
	ThsChanging
)

// TEchoMode ENUM
type TEchoMode = int32

const (
	EmNormal TEchoMode = iota
	EmNone
	EmPassword
)

// TGrabStyle ENUM
type TGrabStyle = int32

const (
	GsSimple TGrabStyle = iota
	GsDouble
	GsHorLines
	GsVerLines
	GsGripper
	GsButton
)

// TGridSaveOptions ENUM
type TGridSaveOptions = int32

const (
	SoDesign     TGridSaveOptions = iota // Save grid structure (col/row count and Options)
	SoAttributes                         // Save grid attributes (Font,Brush,TextStyle)
	SoContent                            // Save Grid Content (Text in StringGrid)
	SoPosition                           // Save Grid cursor and selection position
)

// TSaveOptions SET TGridSaveOptions
type TSaveOptions = TSet

// TEditStyle ENUM
type TEditStyle = int32

const (
	EsSimple TEditStyle = iota
	EsEllipsis
	EsPickList
)

// TFPObservedOperation ENUM
type TFPObservedOperation = int32

const (
	OoChange TFPObservedOperation = iota
	OoFree
	OoAddItem
	OoDeleteItem
	OoCustom
)

// TVleSortCol ENUM
type TVleSortCol = int32

const (
	ColKey TVleSortCol = iota
	ColValue
)

// TControlAtPosFlag ENUM
type TControlAtPosFlag = int32

const (
	CapfAllowDisabled    TControlAtPosFlag = iota // include controls with Enabled=false
	CapfAllowWinControls                          // include TWinControls
	CapfOnlyClientAreas                           // use the client areas, not the whole child area
	CapfRecursive                                 // search recursively in grand childrens
	CapfHasScrollOffset                           // do not add the scroll offset to Pos (already included)
	CapfOnlyWinControls                           // include only TWinControls (ignore TControls)
)

// TControlAtPosFlags SET TControlAtPosFlag
type TControlAtPosFlags = TSet

// TDefaultColorType ENUM
type TDefaultColorType = int32

const (
	DctBrush TDefaultColorType = iota
	DctFont
)

// TControlAutoSizePhase ENUM
type TControlAutoSizePhase = int32

const (
	CaspNone TControlAutoSizePhase = iota
	CaspChangingProperties
	CaspCreatingHandles // create/destroy handles
	CaspComputingBounds
	CaspRealizingBounds
	CaspShowing // make handles visible
)

// TControlAutoSizePhases SET TControlAutoSizePhase
type TControlAutoSizePhases = TSet

// TFindItemKind ENUM
type TFindItemKind = int32

const (
	FkCommand TFindItemKind = iota
	FkHandle
	FkShortCut
)

// TControlRoleForForm ENUM
type TControlRoleForForm = int32

const (
	CrffDefault TControlRoleForForm = iota // this control is notified when user presses Return
	CrffCancel                             // this control is notified when user presses Escape
)

// TControlRolesForForm SET TControlRoleForForm
type TControlRolesForForm = TSet

// TOperation ENUM
type TOperation = int32

const (
	OpInsert TOperation = iota
	OpRemove
)

// TThemeOption ENUM
type TThemeOption = int32

const (
	ToShowButtonImages TThemeOption = iota // show images on buttons
	ToShowMenuImages                       // show images on menus
	ToUseGlyphEffects                      // use hot/down effects on (button) glyphs
)

// TThemedElement ENUM
// These are all elements which can be themed.
type TThemedElement = int32

const (
	TeButton TThemedElement = iota
	TeClock
	TeComboBox
	TeEdit
	TeExplorerBar
	TeHeader
	TeListView
	TeMenu
	TePage
	TeProgress
	TeRebar
	TeScrollBar
	TeSpin
	TeStartPanel
	TeStatus
	TeTab
	TeTaskBand
	TeTaskBar
	TeToolBar
	TeToolTip
	TeTrackBar
	TeTrayNotify
	TeTreeview
	TeWindow
)

// TThemedButton ENUM
// 'Button' theme data
type TThemedButton = int32

const (
	TbButtonDontCare TThemedButton = iota
	TbButtonRoot                   // The root part of each element is sometimes used for special painting and does not belong to a certain state.
	TbPushButtonNormal
	TbPushButtonHot
	TbPushButtonPressed
	TbPushButtonDisabled
	TbPushButtonDefaulted
	TbRadioButtonUncheckedNormal
	TbRadioButtonUncheckedHot
	TbRadioButtonUncheckedPressed
	TbRadioButtonUncheckedDisabled
	TbRadioButtonCheckedNormal
	TbRadioButtonCheckedHot
	TbRadioButtonCheckedPressed
	TbRadioButtonCheckedDisabled
	TbCheckBoxUncheckedNormal
	TbCheckBoxUncheckedHot
	TbCheckBoxUncheckedPressed
	TbCheckBoxUncheckedDisabled
	TbCheckBoxCheckedNormal
	TbCheckBoxCheckedHot
	TbCheckBoxCheckedPressed
	TbCheckBoxCheckedDisabled
	TbCheckBoxMixedNormal
	TbCheckBoxMixedHot
	TbCheckBoxMixedPressed
	TbCheckBoxMixedDisabled
	TbGroupBoxNormal
	TbGroupBoxDisabled
	TbUserButton
)

// TThemedClock ENUM
// 'Clock' theme data
type TThemedClock = int32

const (
	TcClockDontCare TThemedClock = iota
	TcClockRoot
	TcTimeNormal
)

// TThemedComboBox ENUM
// 'ComboBox' theme data
type TThemedComboBox = int32

const (
	TcComboBoxDontCare TThemedComboBox = iota
	TcComboBoxRoot
	TcDropDownButtonNormal
	TcDropDownButtonHot
	TcDropDownButtonPressed
	TcDropDownButtonDisabled
)

// TThemedEdit ENUM
// 'Edit' theme data
type TThemedEdit = int32

const (
	TeEditDontCare TThemedEdit = iota
	TeEditRoot
	TeEditTextNormal
	TeEditTextHot
	TeEditTextSelected
	TeEditTextDisabled
	TeEditTextFocused
	TeEditTextReadOnly
	TeEditTextAssist
	TeEditCaret
)

// TThemedExplorerBar ENUM
// 'ExplorerBar' theme data
type TThemedExplorerBar = int32

const (
	TebExplorerBarDontCare TThemedExplorerBar = iota
	TebExplorerBarRoot
	TebHeaderBackgroundNormal
	TebHeaderBackgroundHot
	TebHeaderBackgroundPressed
	TebHeaderCloseNormal
	TebHeaderCloseHot
	TebHeaderClosePressed
	TebHeaderPinNormal
	TebHeaderPinHot
	TebHeaderPinPressed
	TebHeaderPinSelectedNormal
	TebHeaderPinSelectedHot
	TebHeaderPinSelectedPressed
	TebIEBarMenuNormal
	TebIEBarMenuHot
	TebIEBarMenuPressed
	TebNormalGroupBackground
	TebNormalGroupCollapseNormal
	TebNormalGroupCollapseHot
	TebNormalGroupCollapsePressed
	TebNormalGroupExpandNormal
	TebNormalGroupExpandHot
	TebNormalGroupExpandPressed
	TebNormalGroupHead
	TebSpecialGroupBackground
	TebSpecialGroupCollapseSpecial
	TebSpecialGroupCollapseHot
	TebSpecialGroupCollapsePressed
	TebSpecialGroupExpandSpecial
	TebSpecialGroupExpandHot
	TebSpecialGroupExpandPressed
	TebSpecialGroupHead
)

// TThemedHeader ENUM
// 'Header' theme data
type TThemedHeader = int32

const (
	ThHeaderDontCare TThemedHeader = iota
	ThHeaderRoot
	ThHeaderItemNormal
	ThHeaderItemHot
	ThHeaderItemPressed
	ThHeaderItemLeftNormal
	ThHeaderItemLeftHot
	ThHeaderItemLeftPressed
	ThHeaderItemRightNormal
	ThHeaderItemRightHot
	ThHeaderItemRightPressed
	ThHeaderSortArrowSortedUp
	ThHeaderSortArrowSortedDown
)

// TThemedListView ENUM
// 'ListView' theme data
type TThemedListView = int32

const (
	TlListviewDontCare TThemedListView = iota
	TlListviewRoot
	TlListItemNormal
	TlListItemHot
	TlListItemSelected
	TlListItemDisabled
	TlListItemSelectedNotFocus
	TlListGroup
	TlListDetail
	TlListSortDetail
	TlEmptyText
)

// TThemedMenu ENUM
// 'Menu' theme data
type TThemedMenu = int32

const (
	TmMenuDontCare TThemedMenu = iota
	TmMenuRoot
	TmMenuItemNormal
	TmMenuItemSelected
	TmMenuItemDemoted
	TmMenuDropDown
	TmMenuBarItem
	TmMenuBarDropDown
	TmChevron
	TmSeparator
	TmBarBackgroundActive
	TmBarBackgroundInactive
	TmBarItemNormal
	TmBarItemHot
	TmBarItemPushed
	TmBarItemDisabled
	TmBarItemDisabledHot
	TmBarItemDisabledPushed
	TmPopupBackground
	TmPopupBorders
	TmPopupCheckMarkNormal
	TmPopupCheckMarkDisabled
	TmPopupBulletNormal
	TmPopupBulletDisabled
	TmPopupCheckBackgroundDisabled
	TmPopupCheckBackgroundNormal
	TmPopupCheckBackgroundBitmap
	TmPopupGutter
	TmPopupItemNormal
	TmPopupItemHot
	TmPopupItemDisabled
	TmPopupItemDisabledHot
	TmPopupSeparator
	TmPopupSubmenuNormal
	TmPopupSubmenuDisabled
	TmSystemCloseNormal
	TmSystemCloseDisabled
	TmSystemMaximizeNormal
	TmSystemMaximizeDisabled
	TmSystemMinimizeNormal
	TmSystemMinimizeDisabled
	TmSystemRestoreNormal
	TmSystemRestoreDisabled
)

// TThemedPage ENUM
// 'Page' theme data
type TThemedPage = int32

const (
	TpPageDontCare TThemedPage = iota
	TpPageRoot
	TpUpNormal
	TpUpHot
	TpUpPressed
	TpUpDisabled
	TpDownNormal
	TpDownHot
	TpDownPressed
	TpDownDisabled
	TpUpHorzNormal
	TpUpHorzHot
	TpUpHorzPressed
	TpUpHorzDisabled
	TpDownHorzNormal
	TpDownHorzHot
	TpDownHorzPressed
	TpDownHorzDisabled
)

// TThemedProgress ENUM
// 'Progress' theme data
type TThemedProgress = int32

const (
	TpProgressDontCare TThemedProgress = iota
	TpProgressRoot
	TpBar
	TpBarVert
	TpChunk
	TpChunkVert
)

// TThemedRebar ENUM
// 'Rebar' theme data
type TThemedRebar = int32

const (
	TrRebarDontCare TThemedRebar = iota
	TrRebarRoot
	TrGripper
	TrGripperVert
	TrBandNormal
	TrBandHot
	TrBandPressed
	TrBandDisabled
	TrBandChecked
	TrBandHotChecked
	TrChevronNormal
	TrChevronHot
	TrChevronPressed
	TrChevronDisabled
	TrChevronVertNormal
	TrChevronVertHot
	TrChevronVertPressed
	TrChevronVertDisabled
)

// TThemedScrollBar ENUM
// 'ScrollBar' theme data
type TThemedScrollBar = int32

const (
	TsScrollBarDontCare TThemedScrollBar = iota
	TsScrollBarRoot
	TsArrowBtnUpNormal
	TsArrowBtnUpHot
	TsArrowBtnUpPressed
	TsArrowBtnUpDisabled
	TsArrowBtnDownNormal
	TsArrowBtnDownHot
	TsArrowBtnDownPressed
	TsArrowBtnDownDisabled
	TsArrowBtnLeftNormal
	TsArrowBtnLeftHot
	TsArrowBtnLeftPressed
	TsArrowBtnLeftDisabled
	TsArrowBtnRightNormal
	TsArrowBtnRightHot
	TsArrowBtnRightPressed
	TsArrowBtnRightDisabled
	TsThumbBtnHorzNormal
	TsThumbBtnHorzHot
	TsThumbBtnHorzPressed
	TsThumbBtnHorzDisabled
	TsThumbBtnVertNormal
	TsThumbBtnVertHot
	TsThumbBtnVertPressed
	TsThumbBtnVertDisabled
	TsLowerTrackHorzNormal
	TsLowerTrackHorzHot
	TsLowerTrackHorzPressed
	TsLowerTrackHorzDisabled
	TsUpperTrackHorzNormal
	TsUpperTrackHorzHot
	TsUpperTrackHorzPressed
	TsUpperTrackHorzDisabled
	TsLowerTrackVertNormal
	TsLowerTrackVertHot
	TsLowerTrackVertPressed
	TsLowerTrackVertDisabled
	TsUpperTrackVertNormal
	TsUpperTrackVertHot
	TsUpperTrackVertPressed
	TsUpperTrackVertDisabled
	TsGripperHorzNormal
	TsGripperHorzHot
	TsGripperHorzPressed
	TsGripperHorzDisabled
	TsGripperVertNormal
	TsGripperVertHot
	TsGripperVertPressed
	TsGripperVertDisabled
	TsSizeBoxRightAlign
	TsSizeBoxLeftAlign
)

// TThemedSpin ENUM
// 'Spin' theme data
type TThemedSpin = int32

const (
	TsSpinDontCare TThemedSpin = iota
	TsSpinRoot
	TsUpNormal
	TsUpHot
	TsUpPressed
	TsUpDisabled
	TsDownNormal
	TsDownHot
	TsDownPressed
	TsDownDisabled
	TsUpHorzNormal
	TsUpHorzHot
	TsUpHorzPressed
	TsUpHorzDisabled
	TsDownHorzNormal
	TsDownHorzHot
	TsDownHorzPressed
	TsDownHorzDisabled
)

// TThemedStartPanel ENUM
// 'StartPanel' theme data
type TThemedStartPanel = int32

const (
	TspStartPanelDontCare TThemedStartPanel = iota
	TspStartPanelRoot
	TspUserPane
	TspMorePrograms
	TspMoreProgramsArrowNormal
	TspMoreProgramsArrowHot
	TspMoreProgramsArrowPressed
	TspProgList
	TspProgListSeparator
	TspPlacesList
	TspPlacesListSeparator
	TspLogOff
	TspLogOffButtonsNormal
	TspLogOffButtonsHot
	TspLogOffButtonsPressed
	TspUserPicture
	TspPreview
)

// TThemedStatus ENUM
// 'Status' theme data
type TThemedStatus = int32

const (
	TsStatusDontCare TThemedStatus = iota
	TsStatusRoot
	TsPane
	TsGripperPane
	TsGripper
)

// TThemedTab ENUM
// 'Tab' theme data
type TThemedTab = int32

const (
	TtTabDontCare TThemedTab = iota
	TtTabRoot
	TtTabItemNormal
	TtTabItemHot
	TtTabItemSelected
	TtTabItemDisabled
	TtTabItemFocused
	TtTabItemLeftEdgeNormal
	TtTabItemLeftEdgeHot
	TtTabItemLeftEdgeSelected
	TtTabItemLeftEdgeDisabled
	TtTabItemLeftEdgeFocused
	TtTabItemRightEdgeNormal
	TtTabItemRightEdgeHot
	TtTabItemRightEdgeSelected
	TtTabItemRightEdgeDisabled
	TtTabItemRightEdgeFocused
	TtTabItemBothEdgeNormal
	TtTabItemBothEdgeHot
	TtTabItemBothEdgeSelected
	TtTabItemBothEdgeDisabled
	TtTabItemBothEdgeFocused
	TtTopTabItemNormal
	TtTopTabItemHot
	TtTopTabItemSelected
	TtTopTabItemDisabled
	TtTopTabItemFocused
	TtTopTabItemLeftEdgeNormal
	TtTopTabItemLeftEdgeHot
	TtTopTabItemLeftEdgeSelected
	TtTopTabItemLeftEdgeDisabled
	TtTopTabItemLeftEdgeFocused
	TtTopTabItemRightEdgeNormal
	TtTopTabItemRightEdgeHot
	TtTopTabItemRightEdgeSelected
	TtTopTabItemRightEdgeDisabled
	TtTopTabItemRightEdgeFocused
	TtTopTabItemBothEdgeNormal
	TtTopTabItemBothEdgeHot
	TtTopTabItemBothEdgeSelected
	TtTopTabItemBothEdgeDisabled
	TtTopTabItemBothEdgeFocused
	TtPane
	TtBody
)

// TThemedTaskBand ENUM
// 'TaskBand' theme data
type TThemedTaskBand = int32

const (
	TtbTaskBandDontCare TThemedTaskBand = iota
	TtbTaskBandRoot
	TtbGroupCount
	TtbFlashButton
	TtpFlashButtonGroupMenu
)

// TThemedTaskBar ENUM
// 'TaskBar' theme data
type TThemedTaskBar = int32

const (
	TtTaskBarDontCare TThemedTaskBar = iota
	TtTaskBarRoot
	TtbTimeNormal
)

// TThemedToolBar ENUM
// 'ToolBar' theme data
type TThemedToolBar = int32

const (
	TtbToolBarDontCare TThemedToolBar = iota
	TtbToolBarRoot
	TtbButtonNormal
	TtbButtonHot
	TtbButtonPressed
	TtbButtonDisabled
	TtbButtonChecked
	TtbButtonCheckedHot
	TtbDropDownButtonNormal
	TtbDropDownButtonHot
	TtbDropDownButtonPressed
	TtbDropDownButtonDisabled
	TtbDropDownButtonChecked
	TtbDropDownButtonCheckedHot
	TtbSplitButtonNormal
	TtbSplitButtonHot
	TtbSplitButtonPressed
	TtbSplitButtonDisabled
	TtbSplitButtonChecked
	TtbSplitButtonCheckedHot
	TtbSplitButtonDropDownNormal
	TtbSplitButtonDropDownHot
	TtbSplitButtonDropDownPressed
	TtbSplitButtonDropDownDisabled
	TtbSplitButtonDropDownChecked
	TtbSplitButtonDropDownCheckedHot
	TtbSeparatorNormal
	TtbSeparatorHot
	TtbSeparatorPressed
	TtbSeparatorDisabled
	TtbSeparatorChecked
	TtbSeparatorCheckedHot
	TtbSeparatorVertNormal
	TtbSeparatorVertHot
	TtbSeparatorVertPressed
	TtbSeparatorVertDisabled
	TtbSeparatorVertChecked
	TtbSeparatorVertCheckedHot
)

// TThemedToolTip ENUM
// 'ToolTip' theme data
type TThemedToolTip = int32

const (
	TttToolTipDontCare TThemedToolTip = iota
	TttToolTipRoot
	TttStandardNormal
	TttStandardLink
	TttStandardTitleNormal
	TttStandardTitleLink
	TttBaloonNormal
	TttBaloonLink
	TttBaloonTitleNormal
	TttBaloonTitleLink
	TttCloseNormal
	TttCloseHot
	TttClosePressed
)

// TThemedTrackBar ENUM
// 'TrackBar' theme data
type TThemedTrackBar = int32

const (
	TtbTrackBarDontCare TThemedTrackBar = iota
	TtbTrackBarRoot
	TtbTrack
	TtbTrackVert
	TtbThumbNormal
	TtbThumbHot
	TtbThumbPressed
	TtbThumbFocused
	TtbThumbDisabled
	TtbThumbBottomNormal
	TtbThumbBottomHot
	TtbThumbBottomPressed
	TtbThumbBottomFocused
	TtbThumbBottomDisabled
	TtbThumbTopNormal
	TtbThumbTopHot
	TtbThumbTopPressed
	TtbThumbTopFocused
	TtbThumbTopDisabled
	TtbThumbVertNormal
	TtbThumbVertHot
	TtbThumbVertPressed
	TtbThumbVertFocused
	TtbThumbVertDisabled
	TtbThumbLeftNormal
	TtbThumbLeftHot
	TtbThumbLeftPressed
	TtbThumbLeftFocused
	TtbThumbLeftDisabled
	TtbThumbRightNormal
	TtbThumbRightHot
	TtbThumbRightPressed
	TtbThumbRightFocused
	TtbThumbRightDisabled
	TtbThumbTics
	TtbThumbTicsVert
)

// TThemedTrayNotify ENUM
// 'TrayNotify' theme data
type TThemedTrayNotify = int32

const (
	TtnTrayNotifyDontCare TThemedTrayNotify = iota
	TtnTrayNotifyRoot
	TtnBackground
	TtnAnimBackground
)

// TThemedTreeview ENUM
// 'Treeview' theme data
type TThemedTreeview = int32

const (
	TtTreeviewDontCare TThemedTreeview = iota
	TtTreeviewRoot
	TtItemNormal
	TtItemHot
	TtItemSelected
	TtItemDisabled
	TtItemSelectedNotFocus
	TtGlyphClosed
	TtGlyphOpened
	TtBranch
	TtHotGlyphClosed
	TtHotGlyphOpened
)

// TThemedWindow ENUM
// 'Window' theme data
type TThemedWindow = int32

const (
	TwWindowDontCare TThemedWindow = iota
	TwWindowRoot
	TwCaptionActive
	TwCaptionInactive
	TwCaptionDisabled
	TwSmallCaptionActive
	TwSmallCaptionInactive
	TwSmallCaptionDisabled
	TwMinCaptionActive
	TwMinCaptionInactive
	TwMinCaptionDisabled
	TwSmallMinCaptionActive
	TwSmallMinCaptionInactive
	TwSmallMinCaptionDisabled
	TwMaxCaptionActive
	TwMaxCaptionInactive
	TwMaxCaptionDisabled
	TwSmallMaxCaptionActive
	TwSmallMaxCaptionInactive
	TwSmallMaxCaptionDisabled
	TwFrameLeftActive
	TwFrameLeftInactive
	TwFrameRightActive
	TwFrameRightInactive
	TwFrameBottomActive
	TwFrameBottomInactive
	TwSmallFrameLeftActive
	TwSmallFrameLeftInactive
	TwSmallFrameRightActive
	TwSmallFrameRightInactive
	TwSmallFrameBottomActive
	TwSmallFrameBottomInactive
	TwSysButtonNormal
	TwSysButtonHot
	TwSysButtonPushed
	TwSysButtonDisabled
	TwSysButtonInactive
	TwMDISysButtonNormal
	TwMDISysButtonHot
	TwMDISysButtonPushed
	TwMDISysButtonDisabled
	TwMDISysButtonInactive
	TwMinButtonNormal
	TwMinButtonHot
	TwMinButtonPushed
	TwMinButtonDisabled
	TwMinButtonInactive
	TwMDIMinButtonNormal
	TwMDIMinButtonHot
	TwMDIMinButtonPushed
	TwMDIMinButtonDisabled
	TwMDIMinButtonInactive
	TwMaxButtonNormal
	TwMaxButtonHot
	TwMaxButtonPushed
	TwMaxButtonDisabled
	TwMaxButtonInactive
	TwCloseButtonNormal
	TwCloseButtonHot
	TwCloseButtonPushed
	TwCloseButtonDisabled
	TwCloseButtonInactive
	TwSmallCloseButtonNormal
	TwSmallCloseButtonHot
	TwSmallCloseButtonPushed
	TwSmallCloseButtonDisabled
	TwSmallCloseButtonInactive
	TwMDICloseButtonNormal
	TwMDICloseButtonHot
	TwMDICloseButtonPushed
	TwMDICloseButtonDisabled
	TwMDICloseButtonInactive
	TwRestoreButtonNormal
	TwRestoreButtonHot
	TwRestoreButtonPushed
	TwRestoreButtonDisabled
	TwRestoreButtonInactive
	TwMDIRestoreButtonNormal
	TwMDIRestoreButtonHot
	TwMDIRestoreButtonPushed
	TwMDIRestoreButtonDisabled
	TwMDIRestoreButtonInactive
	TwHelpButtonNormal
	TwHelpButtonHot
	TwHelpButtonPushed
	TwHelpButtonDisabled
	TwHelpButtonInactive
	TwMDIHelpButtonNormal
	TwMDIHelpButtonHot
	TwMDIHelpButtonPushed
	TwMDIHelpButtonDisabled
	TwMDIHelpButtonInactive
	TwHorzScrollNormal
	TwHorzScrollHot
	TwHorzScrollPushed
	TwHorzScrollDisabled
	TwHorzThumbNormal
	TwHorzThumbHot
	TwHorzThumbPushed
	TwHorzThumbDisabled
	TwVertScrollNormal
	TwVertScrollHot
	TwVertScrollPushed
	TwVertScrollDisabled
	TwVertThumbNormal
	TwVertThumbHot
	TwVertThumbPushed
	TwVertThumbDisabled
	TwDialog
	TwCaptionSizingTemplate
	TwSmallCaptionSizingTemplate
	TwFrameLeftSizingTemplate
	TwSmallFrameLeftSizingTemplate
	TwFrameRightSizingTemplate
	TwSmallFrameRightSizingTemplate
	TwFrameBottomSizingTemplate
	TwSmallFrameBottomSizingTemplate
)

// IdButton ENUM Stock Pixmap Types
type IdButton = int32

const (
	IdButtonBase     IdButton = 0
	IdButtonOk                = IdButtonBase + 1
	IdButtonCancel            = IdButtonBase + 2
	IdButtonHelp              = IdButtonBase + 3
	IdButtonYes               = IdButtonBase + 4
	IdButtonNo                = IdButtonBase + 5
	IdButtonClose             = IdButtonBase + 6
	IdButtonAbort             = IdButtonBase + 7
	IdButtonRetry             = IdButtonBase + 8
	IdButtonIgnore            = IdButtonBase + 9
	IdButtonAll               = IdButtonBase + 10
	IdButtonYesToAll          = IdButtonBase + 11
	IdButtonNoToAll           = IdButtonBase + 12
	IdButtonOpen              = IdButtonBase + 13
	IdButtonSave              = IdButtonBase + 14
	IdButtonShield            = IdButtonBase + 15
	IdButtonContinue          = IdButtonBase + 16
	IdButtonTryAgain          = IdButtonBase + 17
)

// TButtonImage ENUM IdButtonOk.. IdButtonNoToAll
type TButtonImage = IdButton

// TCalendarPart ENUM
type TCalendarPart = int32

const (
	CpNoWhere    TCalendarPart = iota // somewhere
	CpDate                            // date part
	CpWeekNumber                      // week number
	CpTitle                           // somewhere in the title
	CpTitleBtn                        // button in the title
	CpTitleMonth                      // month value in the title
	CpTitleYear                       // year value in the title
)

// TCalendarView ENUM
/*
In Windows since Vista native calendar control has four possible views.
In other widgetsets, as well as in older windows, calendar can only have
standard "month view" - grid with days representing a month.
*/
type TCalendarView = int32

const (
	CvMonth   TCalendarView = iota // grid with days in one month
	CvYear                         // grid with months in one year
	CvDecade                       // grid with years from one decade
	CvCentury                      // grid with decades of one century
)

// TSectionValuesOption ENUM
type TSectionValuesOption = int32

const (
	SvoIncludeComments TSectionValuesOption = iota
	SvoIncludeInvalid
	SvoIncludeQuotes
)

// TSectionValuesOptions SET TSectionValuesOption
type TSectionValuesOptions = TSet

// TVScriptPos ENUM
type TVScriptPos = int32

const (
	VpNormal TVScriptPos = iota
	VpSubScript
	VpSuperScript
)

// TParaAlignment ENUM
type TParaAlignment = int32

const (
	PraLeft TParaAlignment = iota
	PraRight
	PraCenter
	PraJustify
)

// TParaNumStyle ENUM
type TParaNumStyle = int32

const (
	PnNone TParaNumStyle = iota
	PnBullet
	PnNumber
	PnLowLetter
	PnLowRoman
	PnUpLetter
	PnUpRoman
	PnCustomChar
)

// TTabAlignment ENUM
type TTabAlignment = int32

const (
	TabLeft TTabAlignment = iota
	TabCenter
	TabRight
	TabDecimal
	TabWordBar
)

// TSearchOption ENUM
type TSearchOption = int32

const (
	SoMatchCase TSearchOption = iota
	SoWholeWord
	SoBackward
)

// TSearchOptions SET TSearchOption
type TSearchOptions = TSet

// TTextModifyMaskEnum ENUM
type TTextModifyMaskEnum = int32

const (
	TmmColor TTextModifyMaskEnum = iota
	TmmName
	TmmSize
	TmmStyles
	TmmBackColor
)

// TTextModifyMask SET TTextModifyMaskEnum
type TTextModifyMask = TSet

// TParaModifyMaskEnum ENUM
type TParaModifyMaskEnum = int32

const (
	PmmFirstLine TParaModifyMaskEnum = iota
	PmmHeadIndent
	PmmTailIndent
	PmmSpaceBefore
	PmmSpaceAfter
	PmmLineSpacing
)

// TParaModifyMask SET TParaModifyMaskEnum
type TParaModifyMask = TSet

// TCTabControlCapability ENUM
type TCTabControlCapability = int32

const (
	NbcShowCloseButtons TCTabControlCapability = iota
	NbcMultiLine
	NbcPageListPopup
	NbcShowAddTabButton
	NbcTabsSizeable
)

// TCTabControlCapabilities SET TCTabControlCapability
type TCTabControlCapabilities = TSet

type TListAssignOp = int32

const (
	LaCopy TListAssignOp = iota
	LaAnd
	LaOr
	LaXor
	LaSrcUnique
	LaDestUnique
)

// TLazDockHeaderPart ENUM
type TLazDockHeaderPart = int32

const (
	LdhpAll           TLazDockHeaderPart = iota // total header rect
	LdhpCaption                                 // header caption
	LdhpRestoreButton                           // header restore button
	LdhpCloseButton                             // header close button
)

// TDirection ENUM
type TDirection = int32

const (
	FromBeginning TDirection = iota
	FromEnd
)

// TRegDataType ENUM
type TRegDataType = int32

const (
	RdUnknown TRegDataType = iota
	RdString
	RdExpandString
	RdBinary
	RdInteger
	RdIntegerBigEndian
	RdLink
	RdMultiString
	RdResourceList
	RdFullResourceDescriptor
	RdResourceRequirementList
	RdInt64
)

// TMonitorDefaultTo ENUM
type TMonitorDefaultTo = int32

const (
	MdNearest TMonitorDefaultTo = iota
	MdNull
	MdPrimary
)

// TPanelPart ENUM
type TPanelPart = int32

const (
	PpText   TPanelPart = iota // for text and text alignment
	PpBorder                   // for bevel and style
	PpWidth                    // for width
)

// TPanelParts SET TPanelPart
type TPanelParts = TSet

// TUpDownDirection ENUM
type TUpDownDirection = int32

const (
	UpdNone TUpDownDirection = iota
	UpdUp
	UpdDown
)

// TFPImgProgressStage ENUM
type TFPImgProgressStage = int32

const (
	PsStarting TFPImgProgressStage = iota
	PsRunning
	PsEnding
)

// TCellProcessType ENUM
type TCellProcessType = int32

const (
	CpCopy TCellProcessType = iota
	CpPaste
)

// TScrollCode ENUM
type TScrollCode = int32

const (
	// !!! Beware. The position of these enums must correspond to the SB_xxx
	// values in LCLType  (Delphi compatibility, not our decision)
	// MWE: Don't know if this still is a requirement
	//      afaik have I removed all casts from the LCL
	ScLineUp    TScrollCode = iota // = SB_LINEUP
	ScLineDown                     // = SB_LINEDOWN
	ScPageUp                       // = SB_PAGEUP
	ScPageDown                     // = SB_PAGEDOWN
	ScPosition                     // = SB_THUMBPOSITION
	ScTrack                        // = SB_THUMBTRACK
	ScTop                          // = SB_TOP
	ScBottom                       // = SB_BOTTOM
	ScEndScroll                    // = SB_ENDSCROLL
)

// TTreeNodeChangeReason ENUM
type TTreeNodeChangeReason = int32

const (
	NcTextChanged   TTreeNodeChangeReason = iota //The Node's Text has changed
	NcDataChanged                                //The Node's Data has changed
	NcHeightChanged                              //The Node's Height has changed
	NcImageEffect                                //The Node's Image Effect has changed
	NcImageIndex                                 //The Node's Image Index has changed
	NcParentChanged                              //The Node's Parent has changed
	NcVisibility                                 //The Node's Visibility has changed
	NcEnablement                                 //The Node's Enabled/Disabled state has changed
	NcOverlayIndex                               //The Node's Overlay Index has Changed
	NcStateIndex                                 //The Node's State Index has Changed
	NcSelectedIndex                              //The Node's Selected Index has Changed
)

// TStreamOwnership ENUM
type TStreamOwnership = int32

const (
	SoReference TStreamOwnership = iota
	SoOwned
)

// TRawImageQueryFlag ENUM
type TRawImageQueryFlag = int32

const (
	RiqfMono    TRawImageQueryFlag = iota // Include a description for a mono image
	RiqfGrey                              // Include a description for a grey image
	RiqfRGB                               // Include a description for a RGB image
	RiqfAlpha                             // Include a description for an Alpha channel
	RiqfMask                              // Include a description for a Mask
	RiqfPalette                           // Include a description for a Palette
	RiqfUpdate                            // Update given description (instead of clearing it)
)

// TRawImageQueryFlags SET: TRawImageQueryFlag
type TRawImageQueryFlags = TSet

// TVirtualNodeState ENUM
// Be careful when adding new states as this might change the size of the type which in turn
// changes the alignment in the node record as well as the stream chunks.
// Do not reorder the states and always add new states at the end of this enumeration in order to avoid
// breaking existing code.
type TVirtualNodeState = int32

const (
	VsInitialized            TVirtualNodeState = iota // Set after the node has been initialized.
	VsChecking                                        // Node's check state is changing, avoid propagation.
	VsCutOrCopy                                       // Node is selected as cut or copy and paste source.
	VsDisabled                                        // Set if node is disabled.
	VsDeleting                                        // Set when the node is about to be freed.
	VsExpanded                                        // Set if the node is expanded.
	VsHasChildren                                     // Indicates the presence of child nodes without actually setting them.
	VsVisible                                         // Indicate whether the node is visible or not (independant of the expand states of its parents).
	VsSelected                                        // Set if the node is in the current selection.
	VsOnFreeNodeCallRequired                          // Set if user data has been set which requires OnFreeNode.
	VsAllChildrenHidden                               // Set if vsHasChildren is set and no child node has the vsVisible flag set.
	VsClearing                                        // A node's children are being deleted. Don't register structure change event.
	VsMultiline                                       // Node text is wrapped at the cell boundaries instead of being shorted.
	VsHeightMeasured                                  // Node height has been determined and does not need a recalculation.
	VsToggling                                        // Set when a node is expanded/collapsed to prevent recursive calls.
	VsFiltered                                        // Indicates that the node should not be painted (without effecting its children).
)

// TVirtualNodeStates SET: TVirtualNodeState
type TVirtualNodeStates = TSet

// TCheckState ENUM
// The check states include both, transient and fluent (temporary) states. The only temporary state defined so
// far is the pressed state.
type TCheckState = int32

const (
	CsUncheckedNormal  TCheckState = iota // unchecked and not pressed
	CsUncheckedPressed                    // unchecked and pressed
	CsCheckedNormal                       // checked and not pressed
	CsCheckedPressed                      // checked and pressed
	CsMixedNormal                         // 3-state check box and not pressed
	CsMixedPressed                        // 3-state check box and pressed
)

// TCheckType ENUM
type TCheckType = int32

const (
	CtNone TCheckType = iota
	CtTriStateCheckBox
	CtCheckBox
	CtRadioButton
	CtButton
)

// TAutoScrollInterval ENUM
//
//	Limits the speed interval which can be used for auto scrolling (milliseconds).
//	1..1000
type TAutoScrollInterval = int32

// TVTScrollIncrement ENUM
//
//	1..10000
type TVTScrollIncrement = int32

// TVTButtonFillMode ENUM
//
//	is only used when the button style is bsRectangle and determines how to fill the interior.
type TVTButtonFillMode = int32

const (
	FmTreeColor   TVTButtonFillMode = iota // solid color, uses the tree's background color
	FmWindowColor                          // solid color, uses clWindow
	FmShaded                               // color gradient, Windows XP style (legacy code, use toThemeAware on Windows XP instead)
	FmTransparent                          // transparent color, use the item's background color
)

// TVTButtonStyle ENUM
//
//	Determines the look of a tree's buttons.
type TVTButtonStyle = int32

const (
	BsRectangle TVTButtonStyle = iota // traditional Windows look (plus/minus buttons)
	BsTriangle                        // traditional Macintosh look
)

// TCheckImageKind ENUM
type TCheckImageKind = int32

const (
	CkLightCheck    TCheckImageKind = iota // gray cross
	CkDarkCheck                            // black cross
	CkLightTick                            // gray tick mark
	CkDarkTick                             // black tick mark
	CkFlat                                 // flat images (no 3D border)
	CkXP                                   // Windows XP style
	CkCustom                               // application defined check images
	CkSystemFlat                           // Flat system defined check images.
	CkSystemDefault                        // Uses the system check images, theme aware.
)

// TVTNodeAttachMode ENUM
//
//	mode to describe a move action
type TVTNodeAttachMode = int32

const (
	AmNoWhere       TVTNodeAttachMode = iota // just for simplified tests, means to ignore the Add/Insert command
	AmInsertBefore                           // insert node just before destination (as sibling of destination)
	AmInsertAfter                            // insert node just after destionation (as sibling of destination)
	AmAddChildFirst                          // add node as first child of destination
	AmAddChildLast                           // add node as last child of destination
)

// TVTDragImageKind ENUM
//
//	determines whether and how the drag image is to show
type TVTDragImageKind = int32

const (
	DiComplete       TVTDragImageKind = iota // show a complete drag image with all columns, only visible columns are shown
	DiMainColumnOnly                         // show only the main column (the tree column)
	DiNoImage                                // don't show a drag image at all
)

// TDragOperation ENUM
//
//	operations basically allowed during drag'n drop
type TDragOperation = int32

const (
	DoCopy TDragOperation = iota
	DoMove
	DoLink
)

// TDragOperations SET: TDragOperation
type TDragOperations = TSet

// TVTDragType ENUM
//
//	Switch for OLE and VCL drag'n drop. Because it is not possible to have both simultanously.
type TVTDragType = int32

const (
	DtOLE TVTDragType = iota
	DtVCL
)

// TVTDrawSelectionMode ENUM
//
//	Determines how to draw the selection rectangle used for draw selection.
type TVTDrawSelectionMode = int32

const (
	SmDottedRectangle  TVTDrawSelectionMode = iota // same as DrawFocusRect
	SmBlendedRectangle                             // alpha blending, uses special colors (see TVTColors)
)

// TVTHintMode ENUM
type TVTHintMode = int32

const (
	HmDefault        TVTHintMode = iota // show the hint of the control
	HmHint                              // show node specific hint string returned by the application
	HmHintAndDefault                    // same as hmHint but show the control's hint if no node is concerned
	HmTooltip                           // show the text of the node if it isn't already fully shown
)

// TVTIncrementalSearch ENUM
type TVTIncrementalSearch = int32

const (
	TvtIsNone            TVTIncrementalSearch = iota // disable incremental search
	TvtIsAll                                         // search every node in tree, initialize if necessary
	TvtIsInitializedOnly                             // search only initialized nodes, skip others
	TvtIsVisibleOnly                                 // search only visible nodes, initialize if necessary
)

// TVTSearchDirection ENUM
//
//	Determines which direction to use when advancing nodes during an incremental search.
type TVTSearchDirection = int32

const (
	SdForward TVTSearchDirection = iota
	SdBackward
)

// TVTSearchStart ENUM
//
//	Determines where to start incremental searching for each key press.
type TVTSearchStart = int32

const (
	SsAlwaysStartOver TVTSearchStart = iota // always use the first/last node (depending on direction) to search from
	SsLastHit                               // use the last found node
	SsFocusedNode                           // use the currently focused node
)

// TVTLineMode ENUM
//
//	Determines how to draw tree lines.
type TVTLineMode = int32

const (
	LmNormal TVTLineMode = iota // usual tree lines (as in TTreeview)
	LmBands                     // looks similar to a Nassi-Schneidermann diagram
)

// TVTLineStyle ENUM
//
//	Determines the look of a tree's lines.
type TVTLineStyle = int32

const (
	LsCustomStyle TVTLineStyle = iota // application provides a line pattern
	LsDotted                          // usual dotted lines (default)
	LsSolid                           // simple solid lines
)

// TVTNodeAlignment ENUM
//
//	Determines how to use the align member of a node.
type TVTNodeAlignment = int32

const (
	NaFromBottom   TVTNodeAlignment = iota // the align member specifies amount of units (usually pixels) from top border of the node
	NaFromTop                              // align is to be measured from bottom
	NaProportional                         // align is to be measure in percent of the entire node height and relative to top
)

// TVirtualTreeState ENUM
//
//	Various events must be handled at different places than they were initiated or need
//	a persistent storage until they are reset.
type TVirtualTreeState = int32

const (
	TsCancelHintAnimation    TVirtualTreeState = iota // Set when a new hint is about to show but an old hint is still being animated.
	TsChangePending                                   // A selection change is pending.
	TsCheckPropagation                                // Set during automatic check state propagation.
	TsCollapsing                                      // A full collapse operation is in progress.
	TsToggleFocusedSelection                          // Node selection was modifed using Ctrl-click. Change selection state on next mouse up.
	TsClearPending                                    // Need to clear the current selection on next mouse move.
	TsClipboardFlushing                               // Set during flushing the clipboard to avoid freeing the content.
	TsCopyPending                                     // Indicates a pending copy operation which needs to be finished.
	TsCutPending                                      // Indicates a pending cut operation which needs to be finished.
	TsDrawSelPending                                  // Multiselection only. User held down the left mouse button on a free
	// area and might want to start draw selection.
	TsDrawSelecting            // Multiselection only. Draw selection has actually started.
	TsEditing                  // Indicates that an edit operation is currently in progress.
	TsEditPending              // An mouse up start edit if dragging has not started.
	TsExpanding                // A full expand operation is in progress.
	TsNodeHeightTracking       // A node height changing operation is in progress.
	TsNodeHeightTrackPending   // left button is down, user might want to start changing a node's height.
	TsHint                     // Set when our hint is visible or soon will be.
	TsInAnimation              // Set if the tree is currently in an animation loop.
	TsIncrementalSearching     // Set when the user starts incremental search.
	TsIncrementalSearchPending // Set in WM_KEYDOWN to tell to use the char in WM_CHAR for incremental search.
	TsIterating                // Set when IterateSubtree is currently in progress.
	TsKeyCheckPending          // A check operation is under way, initiated by a key press (space key). Ignore mouse.
	TsLeftButtonDown           // Set when the left mouse button is down.
	TsLeftDblClick             // Set when the left mouse button was doubly clicked.
	TsMouseCheckPending        // A check operation is under way, initiated by a mouse click. Ignore space key.
	TsMiddleButtonDown         // Set when the middle mouse button is down.
	TsMiddleDblClick           // Set when the middle mouse button was doubly clicked.
	TsNeedRootCountUpdate      // Set if while loading a root node count is set.
	TsOLEDragging              // OLE dragging in progress.
	TsOLEDragPending           // User has requested to start delayed dragging.
	TsPainting                 // The tree is currently painting itself.
	TsRightButtonDown          // Set when the right mouse button is down.
	TsRightDblClick            // Set when the right mouse button was doubly clicked.
	TsPopupMenuShown           // The user clicked the right mouse button, which might cause a popup menu to appear.
	TsScrolling                // Set when autoscrolling is active.
	TsScrollPending            // Set when waiting for the scroll delay time to elapse.
	TsSizing                   // Set when the tree window is being resized. This is used to prevent recursive calls
	// due to setting the scrollbars when sizing.
	TsStopValidation             // Cache validation can be stopped (usually because a change has occurred meanwhile).
	TsStructureChangePending     // The structure of the tree has been changed while the update was locked.
	TsSynchMode                  // Set when the tree is in synch mode, where no timer events are triggered.
	TsThumbTracking              // Stop updating the horizontal scroll bar while dragging the vertical thumb and vice versa.
	TsToggling                   // A toggle operation (for some node) is in progress.
	TsUpdateHiddenChildrenNeeded // Pending update for the hidden children flag after massive visibility changes.
	TsUpdating                   // The tree does currently not update its window because a BeginUpdate has not yet ended.
	TsUseCache                   // The tree's node caches are validated and non-empty.
	TsUserDragObject             // Signals that the application created an own drag object in OnStartDrag.
	TsUseThemes                  // The tree runs under WinXP+, is theme aware and themes are enabled.
	TsValidating                 // The tree's node caches are currently validated.
	TsPreviouslySelectedLocked   // The member FPreviouslySelected should not be changed
	TsValidationNeeded           // Something in the structure of the tree has changed. The cache needs validation.
	TsVCLDragging                // VCL drag'n drop in progress.
	TsVCLDragPending             // One-shot flag to avoid clearing the current selection on implicit mouse up for VCL drag.
	TsVCLDragFinished            // Flag to avoid triggering the OnColumnClick event twice
	TsWheelPanning               // Wheel mouse panning is active or soon will be.
	TsWheelScrolling             // Wheel mouse scrolling is active or soon will be.
	TsWindowCreating             // Set during window handle creation to avoid frequent unnecessary updates.
	TsUseExplorerTheme           // The tree runs under WinVista+ and is using the explorer theme
)

// TVirtualTreeStates SET: TVirtualTreeState
//
//	Various events must be handled at different places than they were initiated or need
//	a persistent storage until they are reset.
type TVirtualTreeStates = TSet

// TVTInternalPaintOption ENUM
//
//	options which determine what to draw in PaintTree
type TVTInternalPaintOption = int32

const (
	PoBackground    TVTInternalPaintOption = iota // draw background image if there is any and it is enabled
	PoColumnColor                                 // erase node's background with the column's color
	PoDrawFocusRect                               // draw focus rectangle around the focused node
	PoDrawSelection                               // draw selected nodes with the normal selection color
	PoDrawDropMark                                // draw drop mark if a node is currently the drop target
	PoGridLines                                   // draw grid lines if enabled
	PoMainOnly                                    // draw only the main column
	PoSelectedOnly                                // draw only selected nodes
	PoUnbuffered                                  // draw directly onto the target canvas; especially useful when printing
)

// TVTInternalPaintOptions SET: TVTInternalPaintOption
type TVTInternalPaintOptions = TSet

// TVTImageKind ENUM
type TVTImageKind = int32

const (
	IkNormal TVTImageKind = iota
	IkSelected
	IkState
	ikOverlay
)

// TVSTTextSourceType ENUM
//
//	Describes the source to use when converting a string tree into a string for clipboard etc.
type TVSTTextSourceType = int32

const (
	TstAll         TVSTTextSourceType = iota // All nodes are rendered. Initialization is done on the fly.
	TstInitialized                           // Only initialized nodes are rendered.
	TstSelected                              // Only selected nodes are rendered.
	TstCutCopySet                            // Only nodes currently marked as being in the cut/copy clipboard set are rendered.
	TstVisible                               // Only visible nodes are rendered.
	TstChecked                               // Only checked nodes are rendered
)

// TVSTTextType ENUM
//
//	Describes the type of text to return in the text and draw info retrival events.
type TVSTTextType = int32

const (
	TtNormal TVSTTextType = iota // normal label of the node, this is also the text which can be edited
	TtStatic                     // static (non-editable) text after the normal text
)

// TVTScrollBarStyle ENUM
type TVTScrollBarStyle = int32

const (
	SbmRegular TVTScrollBarStyle = iota
	Sbm3D
)

// TVTAnimationOption ENUM
//
//	Options to toggle animation support:
type TVTAnimationOption = int32

const (
	ToAnimatedToggle         TVTAnimationOption = iota // Expanding and collapsing a node is animated (quick window scroll).
	ToAdvancedAnimatedToggle                           // Do some advanced animation effects when toggling a node.
)

// TVTAnimationOptions SET: TVTAnimationOption
type TVTAnimationOptions = TSet

// TVTAutoOption ENUM
//
//	Options which toggle automatic handling of certain situations:
type TVTAutoOption = int32

const (
	ToAutoDropExpand           TVTAutoOption = iota // Expand node if it is the drop target for more than a certain time.
	ToAutoExpand                                    // Nodes are expanded (collapsed) when getting (losing) the focus.
	ToAutoScroll                                    // Scroll if mouse is near the border while dragging or selecting.
	ToAutoScrollOnExpand                            // Scroll as many child nodes in view as possible after expanding a node.
	ToAutoSort                                      // Sort tree when Header.SortColumn or Header.SortDirection change or sort node if child nodes are added.
	ToAutoSpanColumns                               // Large entries continue into next column(s) if there's no text in them (no clipping).
	ToAutoTristateTracking                          // Checkstates are automatically propagated for tri state check boxes.
	ToAutoHideButtons                               // Node buttons are hidden when there are child nodes, but all are invisible.
	ToAutoDeleteMovedNodes                          // Delete nodes which where moved in a drag operation (if not directed otherwise).
	ToDisableAutoscrollOnFocus                      // Disable scrolling a node or column into view if it gets focused.
	ToAutoChangeScale                               // Change default node height automatically if the system's font scale is set to big fonts.
	ToAutoFreeOnCollapse                            // Frees any child node after a node has been collapsed (HasChildren flag stays there).
	ToDisableAutoscrollOnEdit                       // Do not center a node horizontally when it is edited.
	ToAutoBidiColumnOrdering                        // When set then columns (if any exist) will be reordered from lowest index to highest index and vice versa when the tree's bidi mode is changed.
)

// TVTAutoOptions SET: TVTAutoOption
type TVTAutoOptions = TSet

// TVTExportMode
//
//	Options to control data export
type TVTExportMode = int32

const (
	EmAll                   TVTExportMode = iota // export all records (regardless checked state)
	EmChecked                                    // export checked records only
	EmUnchecked                                  // export unchecked records only
	EmVisibleDueToExpansion                      // Do not export nodes that are not visible because their parent is not expanded
	EmSelected                                   // export selected nodes only
)

// TVTMiscOption ENUM
// Options which do not fit into any of the other groups:
type TVTMiscOption = int32

const (
	ToAcceptOLEDrop            TVTMiscOption = iota // Register tree as OLE accepting drop target
	ToCheckSupport                                  // Show checkboxes/radio buttons.
	ToEditable                                      // Node captions can be edited.
	ToFullRepaintOnResize                           // Fully invalidate the tree when its window is resized (CS_HREDRAW/CS_VREDRAW).
	ToGridExtensions                                // Use some special enhancements to simulate and support grid behavior.
	ToInitOnSave                                    // Initialize nodes when saving a tree to a stream.
	ToReportMode                                    // Tree behaves like TListView in report mode.
	ToToggleOnDblClick                              // Toggle node expansion state when it is double clicked.
	ToWheelPanning                                  // Support for mouse panning (wheel mice only). This option and toMiddleClickSelect are mutal exclusive, where panning has precedence.
	ToReadOnly                                      // The tree does not allow to be modified in any way. No action is executed and node editing is not possible.
	ToVariableNodeHeight                            // When set then GetNodeHeight will trigger OnMeasureItem to allow variable node heights.
	ToFullRowDrag                                   // Start node dragging by clicking anywhere in it instead only on the caption or image. Must be used together with toDisableDrawSelection.
	ToNodeHeightResize                              // Allows changing a node's height via mouse.
	ToNodeHeightDblClickResize                      // Allows to reset a node's height to FDefaultNodeHeight via a double click.
	ToEditOnClick                                   // Editing mode can be entered with a single click
	ToEditOnDblClick                                // Editing mode can be entered with a double click
	ToReverseFullExpandHotKey                       // Used to define Ctrl+'+' instead of Ctrl+Shift+'+' for full expand (and similar for collapsing)
)

// TVTMiscOptions SET: TVTMiscOption
type TVTMiscOptions = TSet

// TVTPaintOption ENUM
//
//	There is a heap of switchable behavior in the tree. Since published properties may never exceed 4 bytes,
//	which limits sets to at most 32 members, and because for better overview tree options are splitted
//	in various sub-options and are held in a commom options class.
//	Options to customize tree appearance:
type TVTPaintOption = int32

const (
	ToHideFocusRect         TVTPaintOption = iota // Avoid drawing the dotted rectangle around the currently focused node.
	ToHideSelection                               // Selected nodes are drawn as unselected nodes if the tree is unfocused.
	ToHotTrack                                    // Track which node is under the mouse cursor.
	ToPopupMode                                   // Paint tree as would it always have the focus (useful for tree combo boxes etc.)
	ToShowBackground                              // Use the background image if there's one.
	ToShowButtons                                 // Display collapse/expand buttons left to a node.
	ToShowDropmark                                // Show the dropmark during drag'n drop operations.
	ToShowHorzGridLines                           // Display horizontal lines to simulate a grid.
	ToShowRoot                                    // Show lines also at top level (does not show the hidden/internal root node).
	ToShowTreeLines                               // Display tree lines to show hierarchy of nodes.
	ToShowVertGridLines                           // Display vertical lines (depending on columns) to simulate a grid.
	ToThemeAware                                  // Draw UI elements (header, tree buttons etc.) according to the current theme if enabled (Windows XP+ only, application must be themed).
	ToUseBlendedImages                            // Enable alpha blending for ghosted nodes or those which are being cut/copied.
	ToGhostedIfUnfocused                          // Ghosted images are still shown as ghosted if unfocused (otherwise the become non-ghosted images).
	ToFullVertGridLines                           // Display vertical lines over the full client area, not only the space occupied by nodes. This option only has an effect if toShowVertGridLines is enabled too.
	ToAlwaysHideSelection                         // Do not draw node selection, regardless of focused state.
	ToUseBlendedSelection                         // Enable alpha blending for node selections.
	ToStaticBackground                            // Show simple static background instead of a tiled one.
	ToChildrenAbove                               // Display child nodes above their parent.
	ToFixedIndent                                 // Draw the tree with a fixed indent.
	ToUseExplorerTheme                            // Use the explorer theme if run under Windows Vista (or above).
	ToHideTreeLinesIfThemed                       // Do not show tree lines if theming is used.
	ToShowFilteredNodes                           // Draw nodes even if they are filtered out.
)

// TVTPaintOptions SET: TVTPaintOption
type TVTPaintOptions = TSet

// TVTSelectionOption ENUM
//
//	Options which determine the tree's behavior when selecting nodes:
type TVTSelectionOption = int32

const (
	ToDisableDrawSelection    TVTSelectionOption = iota // Prevent user from selecting with the selection rectangle in multiselect mode.
	ToExtendedFocus                                     // Entries other than in the main column can be selected, edited etc.
	ToFullRowSelect                                     // Hit test as well as selection highlight are not constrained to the text of a node.
	ToLevelSelectConstraint                             // Constrain selection to the same level as the selection anchor.
	ToMiddleClickSelect                                 // Allow selection, dragging etc. with the middle mouse button. This and toWheelPanning are mutual exclusive.
	ToMultiSelect                                       // Allow more than one node to be selected.
	ToRightClickSelect                                  // Allow selection, dragging etc. with the right mouse button.
	ToSiblingSelectConstraint                           // Constrain selection to nodes with same parent.
	ToCenterScrollIntoView                              // Center nodes vertically in the client area when scrolling into view.
	ToSimpleDrawSelection                               // Simplifies draw selection, so a node's caption does not need to intersect with the selection rectangle.
	ToAlwaysSelectNode                                  // If this flag is set to true, the tree view tries to always have a node selected. This behavior is closer to the Windows TreeView and useful in Windows Explorer style applications.
	ToRestoreSelection                                  // Set to true if upon refill the previously selected nodes should be selected again.  The nodes will be identified by its caption only.
)

// TVTSelectionOptions SET: TVTSelectionOption
type TVTSelectionOptions = TSet

// TVTStringOption
// Options regarding strings (useful only for the string tree and descendants):
type TVTStringOption = int32

const (
	ToSaveCaptions         TVTStringOption = iota // If set then the caption is automatically saved with the tree node, regardless of what is saved in the user data.
	ToShowStaticText                              // Show static text in a caption which can be differently formatted than the caption but cannot be edited.
	ToAutoAcceptEditChange                        // Automatically accept changes during edit if the user finishes editing other then VK_RETURN or ESC. If not set then changes are cancelled.
)

// TVTStringOptions SET: TVTStringOption
type TVTStringOptions = TSet

// TVTDragMoveRestriction ENUM
//
//	Simple move limitation for the drag image.
type TVTDragMoveRestriction = int32

const (
	DmrNone TVTDragMoveRestriction = iota
	DmrHorizontalOnly
	DmrVerticalOnly
)

// TVTTransparency ENUM  0..255
// Drag image support for the tree.
type TVTTransparency = uint8

// TVTBias ENUM -128..127
type TVTBias = int8

// THeaderState ENUM
type THeaderState = int32

const (
	HsAutoSizing              THeaderState = iota // auto size chain is in progess, do not trigger again on WM_SIZE
	HsDragging                                    // header dragging is in progress (only if enabled)
	HsDragPending                                 // left button is down, user might want to start dragging a column
	HsLoading                                     // The header currently loads from stream, so updates are not necessary.
	HsColumnWidthTracking                         // column resizing is in progress
	HsColumnWidthTrackPending                     // left button is down, user might want to start resize a column
	HsHeightTracking                              // height resizing is in progress
	HsHeightTrackPending                          // left button is down, user might want to start changing height
	HsResizing                                    // multi column resizing in progress
	HsScaling                                     // the header is scaled after a change of FixedAreaConstraints or client size
	HsNeedScaling                                 // the header needs to be scaled
)

// THeaderStates SET: THeaderState
type THeaderStates = TSet

// TVTHeaderOption ENUM
type TVTHeaderOption = int32

const (
	HoAutoResize            TVTHeaderOption = iota // Adjust a column so that the header never exceeds the client width of the owner control.
	HoColumnResize                                 // Resizing columns with the mouse is allowed.
	HoDblClickResize                               // Allows a column to resize itself to its largest entry.
	HoDrag                                         // Dragging columns is allowed.
	HoHotTrack                                     // Header captions are highlighted when mouse is over a particular column.
	HoOwnerDraw                                    // Header items with the owner draw style can be drawn by the application via event.
	HoRestrictDrag                                 // Header can only be dragged horizontally.
	HoShowHint                                     // Show application defined header hint.
	HoShowImages                                   // Show header images.
	HoShowSortGlyphs                               // Allow visible sort glyphs.
	HoVisible                                      // Header is visible.
	HoAutoSpring                                   // Distribute size changes of the header to all columns, which are sizable and have the coAutoSpring option enabled.
	HoFullRepaintOnResize                          // Fully invalidate the header (instead of subsequent columns only) when a column is resized.
	HoDisableAnimatedResize                        // Disable animated resize for all columns.
	HoHeightResize                                 // Allow resizing header height via mouse.
	HoHeightDblClickResize                         // Allow the header to resize itself to its default height.
	HoHeaderClickAutoSort                          // Clicks on the header will make the clicked column the SortColumn or toggle sort direction if it already was the sort column
)

// TVTHeaderOptions SET: TVTHeaderOption
type TVTHeaderOptions = TSet

// TVTHeaderStyle ENUM
type TVTHeaderStyle = int32

const (
	HsThickButtons TVTHeaderStyle = iota // TButton look and feel
	HsFlatButtons                        // flatter look than hsThickButton, like an always raised flat TToolButton
	HsPlates                             // flat TToolButton look and feel (raise on hover etc.)
)

// TVTColumnOption
// Options per column.
type TVTColumnOption = int32

const (
	CoAllowClick            TVTColumnOption = iota // Column can be clicked (must be enabled too).
	CoDraggable                                    // Column can be dragged.
	CoEnabled                                      // Column is enabled.
	CoParentBidiMode                               // Column uses the parent's bidi mode.
	CoParentColor                                  // Column uses the parent's background color.
	CoResizable                                    // Column can be resized.
	CoShowDropMark                                 // Column shows the drop mark if it is currently the drop target.
	CoVisible                                      // Column is shown.
	CoAutoSpring                                   // Column takes part in the auto spring feature of the header (must be resizable too).
	CoFixed                                        // Column is fixed and can not be selected or scrolled etc.
	CoSmartResize                                  // Column is resized to its largest entry which is in view (instead of its largest visible entry).
	CoAllowFocus                                   // Column can be focused.
	CoDisableAnimatedResize                        // Column resizing is not animated.
	CoWrapCaption                                  // Caption could be wrapped across several header lines to fit columns width.
	CoUseCaptionAlignment                          // Column's caption has its own aligment.
	CoEditable                                     // Column can be edited
)

// TVTColumnOptions SET: TVTColumnOption
type TVTColumnOptions = TSet

// TSmartAutoFitType ENUM
type TSmartAutoFitType = int32

const (
	SmaAllColumns      TSmartAutoFitType = iota // consider nodes in view only for all columns
	SmaNoColumn                                 // consider nodes in view only for no column
	SmaUseColumnOption                          // use coSmartResize of the corresponding column
) // describes the used column resize behaviour for AutoFitColumns

// TVTConstraintPercent ENUM 0..100
type TVTConstraintPercent = uint8

// TVTHeaderColumnLayout ENUM
type TVTHeaderColumnLayout = int32

const (
	GcBlGlyphLeft TVTHeaderColumnLayout = iota
	GcBlGlyphRight
	GcBlGlyphTop
	GcBlGlyphBottom
)

// TVirtualTreeColumnStyle ENUM
type TVirtualTreeColumnStyle = int32

const (
	VsText TVirtualTreeColumnStyle = iota
	VsOwnerDraw
)

// TChangeReason ENUM
type TChangeReason = int32

const (
	CrIgnore       TChangeReason = iota // used as placeholder
	CrAccumulated                       // used for delayed changes
	CrChildAdded                        // one or more child nodes have been added
	CrChildDeleted                      // one or more child nodes have been deleted
	CrNodeAdded                         // a node has been added
	CrNodeCopied                        // a node has been duplicated
	CrNodeMoved                         // a node has been moved to a new place
) // desribes what made a structure change event happen

// TVirtualNodeInitState ENUM
//
//	States used in InitNode to indicate states a node shall initially have.
type TVirtualNodeInitState = int32

const (
	IvsDisabled TVirtualNodeInitState = iota
	IvsExpanded
	IvsHasChildren
	IvsMultiline
	IvsSelected
	IvsFiltered
	IvsReInit
)

// TVirtualNodeInitStates SET: TVirtualNodeInitState
type TVirtualNodeInitStates = TSet

// TVTExportType
// Export type
type TVTExportType = int32

const (
	TvtEtRTF    TVTExportType = iota // contentToRTF
	TvtEtHTML                        // contentToHTML
	TvtEtText                        // contentToText
	TvtEtExcel                       // supported by external tools
	TvtEtWord                        // supported by external tools
	TvtEtCustom                      // supported by external tools
)

// TVTImageInfoIndex ENUM
type TVTImageInfoIndex = int32

const (
	IiNormal TVTImageInfoIndex = iota
	IiState
	IiCheck
	IiOverlay
)

// TVTCellContentMarginType ENUM
//
//	Determines which sides of the cell content margin should be considered.
type TVTCellContentMarginType = int32

const (
	CcmtAllSides        TVTCellContentMarginType = iota // consider all sides
	CcmtTopLeftOnly                                     // consider top margin and left margin only
	CcmtBottomRightOnly                                 // consider bottom margin and right margin only
)

// TVTHeaderHitPosition ENUM
// These flags are used to indicate where a click in the header happened.
type TVTHeaderHitPosition = int32

const (
	HhiNoWhere    TVTHeaderHitPosition = iota // No column is involved (possible only if the tree is smaller than the client area).
	HhiOnColumn                               // On a column.
	HhiOnIcon                                 // On the bitmap associated with a column.
	HhiOnCheckbox                             // On the checkbox if enabled.
)

// TVTHeaderHitPositions SET: TVTHeaderHitPosition
type TVTHeaderHitPositions = TSet

// TVTDropMarkMode ENUM
//
//	Used during owner draw of the header to indicate which drop mark for the column must be drawn.
type TVTDropMarkMode = int32

const (
	DmmNone TVTDropMarkMode = iota
	DmmLeft
	DmmRight
)

// THeaderPaintElement ENUM
//
//	These elements are used both to query the application, which of them it wants to draw itself and to tell it during
//	painting, which elements must be drawn during the advanced custom draw events.
type THeaderPaintElement = int32

const (
	HpeBackground THeaderPaintElement = iota
	HpeDropMark
	HpeHeaderGlyph
	HpeSortGlyph
	HpeText
)

// THeaderPaintElements SET: THeaderPaintElement
type THeaderPaintElements = TSet

// THitPosition ENUM
// These flags are returned by the hit test method.
type THitPosition = int32

const (
	HiAbove             THitPosition = iota // above the client area (if relative) or the absolute tree area
	HiBelow                                 // below the client area (if relative) or the absolute tree area
	HiNowhere                               // no node is involved (possible only if the tree is not as tall as the client area)
	HiOnItem                                // on the bitmaps/buttons or label associated with an item
	HiOnItemButton                          // on the button associated with an item
	HiOnItemButtonExact                     // exactly on the button associated with an item
	HiOnItemCheckbox                        // on the checkbox if enabled
	HiOnItemIndent                          // in the indentation area in front of a node
	HiOnItemLabel                           // on the normal text area associated with an item
	HiOnItemLeft                            // in the area to the left of a node's text area (e.g. when right aligned or centered)
	HiOnItemRight                           // in the area to the right of a node's text area (e.g. if left aligned or centered)
	HiOnNormalIcon                          // on the "normal" image
	HiOnStateIcon                           // on the state image
	HiToLeft                                // to the left of the client area (if relative) or the absolute tree area
	HiToRight                               // to the right of the client area (if relative) or the absolute tree area
	HiUpperSplitter                         // in the upper splitter area of a node
	HiLowerSplitter                         // in the lower splitter area of a node
)

// THitPositions SET: THitPosition
type THitPositions = TSet

// TDropMode ENUM
// modes to determine drop position further
type TDropMode = int32

const (
	DmNowhere TDropMode = iota
	DmAbove
	DmOnNode
	DmBelow
)

// TItemEraseAction ENUM
// Used to describe the action to do when using the OnBeforeItemErase event.
type TItemEraseAction = int32

const (
	EaColor   TItemEraseAction = iota // Use the provided color to erase the background instead the one of the tree.
	EaDefault                         // The tree should erase the item's background (bitmap or solid).
	EaNone                            // Do nothing. Let the application paint the background.
)

// TVTCellPaintMode ENUM
// Determines for which purpose the cell paint event is called.
type TVTCellPaintMode = int32

const (
	CpmPaint            TVTCellPaintMode = iota // painting the cell
	CpmGetContentMargin                         // getting cell content margin
)

// TVTOperationKind ENUM
// Kinds of operations
type TVTOperationKind = int32

const (
	OkAutoFitColumns TVTOperationKind = iota
	OkGetMaxColumnWidth
	OkSortNode
	OkSortTree
)

// TVTOperationKinds SET: TVTOperationKind
type TVTOperationKinds = TSet

// TVTHintKind ENUM
type TVTHintKind = int32

const (
	VhkText TVTHintKind = iota
	VhkOwnerDraw
)

// TVTUpdateState ENUM
// Indicates in the OnUpdating event what state the tree is currently in.
type TVTUpdateState = int32

const (
	UsBegin      TVTUpdateState = iota // The tree just entered the update state (BeginUpdate call for the first time).
	UsBeginSynch                       // The tree just entered the synch update state (BeginSynch call for the first time).
	UsSynch                            // Begin/EndSynch has been called but the tree did not change the update state.
	UsUpdate                           // Begin/EndUpdate has been called but the tree did not change the update state.
	UsEnd                              // The tree just left the update state (EndUpdate called for the last level).
	UsEndSynch                         // The tree just left the synch update state (EndSynch called for the last level).
)

// TVTTooltipLineBreakStyle ENUM
//
//	Indicates how to format a tooltip.
type TVTTooltipLineBreakStyle = int32

const (
	HlbDefault         TVTTooltipLineBreakStyle = iota // Use multi-line style of the node.
	HlbForceSingleLine                                 // Use single line hint.
	HlbForceMultiLine                                  // Use multi line hint.
)
