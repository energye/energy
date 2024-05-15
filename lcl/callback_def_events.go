//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/types"
)

type TNotifyEvent func(sender IObject)
type TUDClickEvent func(sender IObject, button TUDBtnType)
type TCloseEvent func(sender IObject, action *TCloseAction)
type TCloseQueryEvent func(sender IObject, canClose *bool)
type TMenuChangeEvent func(sender IObject, source IMenuItem, rebuild bool)
type TSysLinkEvent func(sender IObject, link string, linkType TSysLinkType)
type TExceptionEvent func(sender IObject, e IException)
type TKeyEvent func(sender IObject, key *Char, shift TShiftState)
type TKeyPressEvent func(sender IObject, key *Char)
type TMouseEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)
type TMouseMoveEvent func(sender IObject, shift TShiftState, x, y int32)
type TMouseWheelEvent func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)
type TDrawItemEvent func(control IWinControl, index int32, rect TRect, state TOwnerDrawState)
type TLVColumnClickEvent func(sender IObject, column IListColumn)
type TLVColumnRClickEvent func(sender IObject, column IListColumn, point TPoint)
type TLVSelectItemEvent func(sender IObject, item IListItem, selected bool)
type TLVCheckedItemEvent func(sender IObject, item IListItem)
type TLVCompareEvent func(sender IObject, item1, item2 IListItem, data int32, compare *int32)
type TLVChangeEvent func(sender IObject, item IListItem, change TItemChange)
type TLVNotifyEvent func(sender IObject, item IListItem)
type TLVAdvancedCustomDrawEvent func(sender IListView, rect TRect, stage TCustomDrawStage, defaultDraw *bool)
type TLVAdvancedCustomDrawItemEvent func(sender IListView, item IListItem, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)
type TLVAdvancedCustomDrawSubItemEvent func(sender IListView, item IListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)
type TTVCompareEvent func(sender IObject, node1, node2 ITreeNode, data int32, compare *int32)
type TTVExpandedEvent func(sender IObject, node ITreeNode)
type TTVChangedEvent func(sender IObject, node ITreeNode)
type TTVAdvancedCustomDrawEvent func(sender ITreeView, rect TRect, stage TCustomDrawStage, defaultDraw *bool)
type TTVAdvancedCustomDrawItemEvent func(sender ITreeView, node ITreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)
type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)
type TTBAdvancedCustomDrawEvent func(sender IToolBar, rect TRect, stage TCustomDrawStage, defaultDraw *bool)

// 主线程运行回调函数
type TMainThreadSyncProc func()
type TMainThreadAsyncProc func(id uint32)
type TWndProcEvent func(msg *TMessage)

// TDropFilesEvent
// 注意，当在Windows上使用时如果使用了UAC，则无法收到消息
// 需要使用未公开的winapi   ChangeWindowMessageFilter 或 ChangeWindowMessageFilterEx 根据系统版本不同使用其中的，然后添加
// ChangeWindowMessageFilterEx(pnl_Drag.Handle, WM_DROPFILES, MSGFLT_ALLOW, LChangeFilterStruct),消息
type TDropFilesEvent func(sender IObject, fileNames []string)
type TConstrainedResizeEvent func(sender IObject, minWidth, minHeight, maxWidth, maxHeight *int32)
type THelpEvent func(command uint16, data THelpEventData, callHelp, result *bool)
type TShortCutEvent func(msg *TWMKey, handled *bool)
type TContextPopupEvent func(sender IObject, mousePos TPoint, handled *bool)
type TDragOverEvent func(sender, source IObject, x, y int32, state TDragState, accept *bool)
type TDragDropEvent func(sender, source IObject, x, y int32)
type TEndDragEvent func(sender, target IObject, x, y int32)
type TDockDropEvent func(sender IObject, source IDragDockObject, x, y int32)
type TDockOverEvent func(sender IObject, source IDragDockObject, x, y int32, state TDragState, accept *bool)
type TUnDockEvent func(sender IObject, client IControl, newTarget IControl, allow *bool)
type TStartDockEvent func(sender IObject, dragObject *IDragDockObject)
type TGetSiteInfoEvent func(sender IObject, dockClient IControl, resultInfluenceRect *TRect, mousePos TPoint, canDock *bool)
type TMouseWheelUpDownEvent func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)
type TGridOperationEvent func(sender IObject, isColumn bool, sIndex, tIndex int32)
type TDrawCellEvent func(sender IObject, col, row int32, rect TRect, state TGridDrawState)
type TFixedCellClickEvent func(sender IObject, col, row int32)
type TGetEditEvent func(sender IObject, col, row int32, value *string)
type TSelectCellEvent func(sender IObject, col, row int32, canSelect *bool)
type TSetEditEvent func(sender IObject, col, row int32, value string)
type TDrawSectionEvent func(headerControl IHeaderControl, section IHeaderSection, rect TRect, pressed bool)
type TSectionNotifyEvent func(headerControl IHeaderControl, section IHeaderSection)
type TSectionTrackEvent func(headerControl IHeaderControl, section IHeaderSection, width int32, state TSectionTrackState)
type TSectionDragEvent func(sender IObject, fromSection, toSection IHeaderSection, allowDrag *bool)
type TCustomSectionNotifyEvent func(headerControl IHeaderControl, section IHeaderSection)

// TGestureEvent = procedure(sender IObject, EventInfo: TGestureEventInfo, Handled: bool)
//type TGestureEvent func(sender IObject, eventInfo TGestureEventInfo, handled *bool)

type TStartDragEvent func(sender IObject, dragObject *IDragObject)
type TMouseActivateEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32, hitTest int32, mouseActivate *TMouseActivate)
type TLBGetDataEvent func(control IWinControl, index int32, data *string)
type TLBGetDataObjectEvent func(control IWinControl, index int32, dataObject IObject)
type TLBFindDataEvent func(control IWinControl, findString string) int32
type TMeasureItemEvent func(control IWinControl, index int32, height *int32)
type TLVChangingEvent func(sender IObject, item IListItem, change TItemChange, allowChange *bool)
type TLVDataFindEvent func(sender IObject, find TItemFind, findString string, findPosition TPoint, findData TCustomData, startIndex int32, direction TSearchDirection, warp bool, index *int32)
type TLVEditingEvent func(sender IObject, item IListItem, allowEdit *bool)
type TLVEditedEvent func(sender IObject, item IListItem, s *string)
type TMenuMeasureItemEvent func(sender IObject, canvas ICanvas, width, height *int32)
type TTabChangingEvent func(sender IObject, allowChange *bool)
type TTVChangingEvent func(sender IObject, node ITreeNode, allowChange *bool)
type TTVCollapsingEvent func(sender IObject, node ITreeNode, allowCollapse *bool)
type TTVEditedEvent func(sender IObject, node ITreeNode, s *string)
type TTVEditingEvent func(sender IObject, node ITreeNode, allowEdit *bool)
type TTVExpandingEvent func(sender IObject, node ITreeNode, allowExpansion *bool)
type TTVHintEvent func(sender IObject, node ITreeNode, hint *string)
type TUDChangingEvent func(sender IObject, allowChange *bool)
type TCreatingListErrorEvent func(sender IObject, winErrorCode uint32, errorDescription string, handled *bool)
type TLVCustomDrawEvent func(sender IListView, rect TRect, defaultDraw *bool)
type TLVCustomDrawItemEvent func(sender IListView, item IListItem, state TCustomDrawState, defaultDraw *bool)
type TLVCustomDrawSubItemEvent func(sender IListView, item IListItem, subItem int32, state TCustomDrawState, defaultDraw *bool)
type TLVDrawItemEvent func(sender IListView, item IListItem, rect TRect, state TOwnerDrawState)
type TLVDataHintEvent func(sender IObject, startIndex, endIndex int32)
type TTVCustomDrawEvent func(sender ITreeView, rect TRect, defaultDraw *bool)
type TTVCustomDrawItemEvent func(sender ITreeView, node ITreeNode, state TCustomDrawState, defaultDraw *bool)
type TWebTitleChangeEvent func(sender IObject, text string)
type TWebJSExternalEvent func(sender IObject, funcName, args string, retVal *string)
type TTaskDlgClickEvent func(sender IObject, modalResult TModalResult, canClose *bool)
type TTaskDlgTimerEvent func(sender IObject, tickCount uint32, reset *bool)
type TAlignPositionEvent func(sender IWinControl, control IControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *TRect, alignInfo TAlignInfo)
type TCheckGroupClicked func(sender IObject, index int32)
type TOnSelectEvent func(sender IObject, col, row int32)
type TToggledCheckboxEvent func(sender IObject, col, row int32, state TCheckBoxState)
type TOnCompareCells func(sender IObject, col, row, col1, row1 int32, result *int32)
type TGetCellHintEvent func(sender IObject, col, row int32, hintText *string)
type TGetCheckboxStateEvent func(sender IObject, col, row int32, value *TCheckBoxState)
type TSetCheckboxStateEvent func(sender IObject, col, row int32, Value TCheckBoxState)
type THdrEvent func(sender IObject, isColumn bool, index int32)
type THeaderSizingEvent func(sender IObject, isColumn bool, index, size int32)
type TSelectEditorEvent func(sender IObject, col, row int32, editor *IWinControl)
type TUserCheckBoxBitmapEvent func(sender IObject, col, row int32, CheckedState TCheckBoxState, bitmap *IBitmap)
type TValidateEntryEvent func(sender IObject, col, row int32, oldValue string, newValue *string)
type TOnPrepareCanvasEvent func(sender IObject, col, row int32, state TGridDrawState)
type TAcceptFileNameEvent func(sender IObject, value *string)
type TCheckItemChange func(sender IObject, index int32)
type TUTF8KeyPressEvent func(sender IObject, utf8key *TUTF8Char)
type TMenuDrawItemEvent func(sender IObject, canvas ICanvas, rect TRect, state TOwnerDrawState)
type TImagePaintBackgroundEvent func(sender IObject, canvas ICanvas, rect TRect)

// new

type TActionEvent func(sender IObject, action IBasicAction, handled *bool)
type TGetHandleEvent func(handle *HWND)
type TModalDialogFinished func(sender IObject, result int32)
type TGetDockCaptionEvent func(sender IObject, control IControl, caption *string)
type TOnDrawCell func(sender IObject, col, row int32, rect TRect, state TGridDrawState)
type TUserCheckBoxImageEvent func(sender IObject, col, row int32, checkedState TCheckBoxState, imageList ICustomImageList, imageIndex TImageIndex)

//type TSaveColumnEvent func(sender, aColumn IObject, aColIndex int32, aCfg IXMLConfig, aVersion int32, aPath string)

type TAlignInsertBeforeEvent func(sender IWinControl, control1, control2 IControl) bool
type TUDChangingEventEx func(sender IObject, allowChange *bool, newValue int16, direction TUpDownDirection)
type TSBCreatePanelClassEvent func(sender IStatusBar, panelClass *TStatusPanelClass)
type TDrawPanelEvent func(statusBar IStatusBar, panel IStatusPanel, rect TRect)
type TDialogResultEvent func(sender IObject, success bool)
type TGetColorsEvent func(sender ICustomColorBox, items IStrings)
type THintEvent func(hintStr *string, canShow bool)
type TOnSelectCellEvent func(sender IObject, col, row int32, canSelect *bool)
type TSelectionChangeEvent func(sender IObject, user bool)
type TFPCanvasCombineColors func(color1, color2 TFPColor, result *TFPColor)
type TFPImgProgressEvent func(sender IObject, stage TFPImgProgressStage, percentDone byte, redrawNow bool, rect TRect, msg string, continue_ *bool)
type TProgressEvent = TFPImgProgressEvent
type TLBGetColorsEvent func(sender ICustomColorListBox, items IStrings)
type TCustomHCCreateSectionClassEvent func(sender ICustomHeaderControl, sectionClass *THeaderSectionClass)
type TCustomSectionTrackEvent func(headerControl ICustomHeaderControl, section IHeaderSection, width int32, state TSectionTrackState)
type TCanOffsetEvent func(sender IObject, newOffset *int32, accept *bool)
type TCanResizeEvent func(sender IObject, newSize *int32, accept *bool)
type TGetPickListEvent func(sender IObject, keyName string, values IStrings)
type TOnValidateEvent func(sender IObject, col, row int32, keyName, keyValue string)
type TClipboardRequestEvent func(requestedFormatID TClipboardFormat, data IStream)
type TToolBarOnPaintButton func(sender IToolButton, state int32)
type TCellProcessEvent func(sender IObject, col, row int32, processType TCellProcessType, value *string)
type TScrollEvent func(sender IObject, scrollCode TScrollCode, scrollPos *int32)
type TListCompareEvent func(aList IListControlItems, item1, item2 IListControlItem, result *int32)
type TCustomImageListGetWidthForPPI func(sender ICustomImageList, imageWidth, PPI int32, resultWidth *int32)
type TLVCreateItemClassEvent func(sender ICustomListView, itemClass *TListItemClass)
type TLVDataStateChangeEvent func(sender IObject, startIndex, endIndex int32, oldState, newState TListItemStates)
type TLVDeletedEvent func(sender IObject, item IListItem)
type TLVInsertEvent = TLVDeletedEvent
type TLVDataEvent = TLVDeletedEvent
type TTVCreateNodeClassEvent func(sender ICustomTreeView, nodeClass *TTreeNodeClass)
type TTVCustomCreateNodeEvent func(sender ICustomTreeView, treeNode ITreeNode)
type TTVCustomDrawArrowEvent func(sender ICustomTreeView, rect TRect, collapsed bool)
type TTVEditingEndEvent func(sender IObject, node ITreeNode, cancel bool)
type TTVHasChildrenEvent func(sender ICustomTreeView, node ITreeNode) bool
type TTVNodeChangedEvent func(sender IObject, node ITreeNode, changeReason TTreeNodeChangeReason)
type TShowHintEvent func(hintStr *string, canShow *bool, hintInfo *THintInfo)
type TControlShowHintEvent func(sender IObject, hintInfo THintInfo)
type TDestroyResolutionHandleEvent func(sender ICustomImageList, width int32, referenceHandle TLCLHandle)

// 方法参数回调

type TTreeNodeCompare func(node1, node2 ITreeNode) int32
type TStringListSortCompare func(list IStringList, index1, index2 int32) int32
type TListItemsCompare func(list IListControlItems, item1, item2 int32) int32
type TLVCompare func(item1, item2 IListItem, optionalParam uint32) int32 //stdcall
type TListSortCompare func(item1, item2 uintptr) int32
type TCollectionSortCompare func(item1, item2 ICollectionItem) int32

// node events

type TVTChangingEvent func(sender IBaseVirtualTree, node IVirtualNode, allowed *bool)
type TVTCheckChangingEvent func(sender IBaseVirtualTree, node IVirtualNode, newState *TCheckState, allowed *bool)
type TVTChangeEvent func(sender IBaseVirtualTree, node IVirtualNode)
type TVTStructureChangeEvent func(sender IBaseVirtualTree, node IVirtualNode, reason TChangeReason)
type TVTEditCancelEvent func(sender IBaseVirtualTree, column TColumnIndex)
type TVTEditChangingEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, allowed *bool)
type TVTEditChangeEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex)
type TVTFreeNodeEvent func(sender IBaseVirtualTree, node IVirtualNode)
type TVTFocusChangingEvent func(sender IBaseVirtualTree, oldNode, newNode IVirtualNode, oldColumn, newColumn TColumnIndex, allowed *bool)
type TVTFocusChangeEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex)
type TVTAddToSelectionEvent func(sender IBaseVirtualTree, node IVirtualNode)
type TVTRemoveFromSelectionEvent func(sender IBaseVirtualTree, node IVirtualNode)
type TVTGetImageEvent func(sender IBaseVirtualTree, node IVirtualNode, kind TVTImageKind, column TColumnIndex, ghosted *bool, imageIndex *int32)
type TVTGetImageExEvent func(sender IBaseVirtualTree, node IVirtualNode, kind TVTImageKind, column TColumnIndex, ghosted *bool, imageIndex *int32, imageList *ICustomImageList)
type TVTGetImageTextEvent func(sender IBaseVirtualTree, node IVirtualNode, kind TVTImageKind, column TColumnIndex, imageText *string)
type TVTHotNodeChangeEvent func(sender IBaseVirtualTree, oldNode, newNode IVirtualNode)
type TVTInitChildrenEvent func(sender IBaseVirtualTree, node IVirtualNode, childCount *Cardinal)
type TVTInitNodeEvent func(sender IBaseVirtualTree, parentNode, node IVirtualNode, initialStates *TVirtualNodeInitStates)
type TVTPopupEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, point TPoint, askParent *bool, popupMenu *IPopupMenu)
type TVTHelpContextEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, helpContext *int32)
type TVTCreateEditorEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, outEditLink *IVTEditLink)
type TVTSaveTreeEvent func(sender IBaseVirtualTree, stream IStream)
type TVTSaveNodeEvent func(sender IBaseVirtualTree, node IVirtualNode, stream IStream)

type TVTNodeExportEvent func(sender IBaseVirtualTree, exportType TVTExportType, Node IVirtualNode) bool
type TVTColumnExportEvent func(sender IBaseVirtualTree, exportType TVTExportType, column IVirtualTreeColumn)
type TVTTreeExportEvent func(sender IBaseVirtualTree, exportType TVTExportType)
type TVTDrawNodeEvent func(sender IBaseVirtualTree, paintInfo TVTPaintInfo)
type TVTGetCellContentMarginEvent func(sender IBaseVirtualTree, hintCanvas ICanvas, node IVirtualNode, column TColumnIndex, cellContentMarginType TVTCellContentMarginType, cellContentMargin *TPoint)
type TVTGetNodeWidthEvent func(sender IBaseVirtualTree, hintCanvas ICanvas, node IVirtualNode, column TColumnIndex, nodeWidth *int32)

// header/column events

type TVTHeaderClickEvent func(sender IVTHeader, hitInfo TVTHeaderHitInfo)
type TVTHeaderMouseEvent func(sender IVTHeader, button TMouseButton, shift TShiftState, x, y int32)
type TVTHeaderMouseMoveEvent func(sender IVTHeader, shift TShiftState, x, y int32)
type TVTBeforeHeaderHeightTrackingEvent func(sender IVTHeader, shift TShiftState)
type TVTAfterHeaderHeightTrackingEvent func(sender IVTHeader)
type TVTHeaderHeightTrackingEvent func(sender IVTHeader, point *TPoint, shift TShiftState, allowed *bool)
type TVTHeaderHeightDblClickResizeEvent func(sender IVTHeader, point *TPoint, shift TShiftState, allowed *bool)
type TVTHeaderNotifyEvent func(sender IVTHeader, column TColumnIndex)
type TVTHeaderDraggingEvent func(sender IVTHeader, column TColumnIndex, allowed *bool)
type TVTHeaderDraggedEvent func(sender IVTHeader, column TColumnIndex, oldPosition int32)
type TVTHeaderDraggedOutEvent func(sender IVTHeader, column TColumnIndex, dropPosition TPoint)
type TVTHeaderPaintEvent func(sender IVTHeader, headerCanvas ICanvas, column IVirtualTreeColumn, rect TRect, hHover, pressed bool, dropMark TVTDropMarkMode)
type TVTHeaderPaintQueryElementsEvent func(sender IVTHeader, paintInfo *THeaderPaintInfo, elements *THeaderPaintElements)
type TVTAdvancedHeaderPaintEvent func(sender IVTHeader, paintInfo *THeaderPaintInfo, elements THeaderPaintElements)
type TVTBeforeAutoFitColumnsEvent func(sender IVTHeader, smartAutoFitType *TSmartAutoFitType)
type TVTBeforeAutoFitColumnEvent func(sender IVTHeader, column TColumnIndex, smartAutoFitType *TSmartAutoFitType, allowed *bool)
type TVTAfterAutoFitColumnEvent func(sender IVTHeader, column TColumnIndex)
type TVTAfterAutoFitColumnsEvent func(sender IVTHeader)
type TVTColumnClickEvent func(sender IBaseVirtualTree, column TColumnIndex, shift TShiftState)
type TVTColumnDblClickEvent func(sender IBaseVirtualTree, column TColumnIndex, shift TShiftState)
type TVTColumnWidthDblClickResizeEvent func(sender IVTHeader, column TColumnIndex, shift TShiftState, point TPoint, allowed *bool)
type TVTBeforeColumnWidthTrackingEvent func(sender IVTHeader, column TColumnIndex, shift TShiftState)
type TVTAfterColumnWidthTrackingEvent func(sender IVTHeader, column TColumnIndex)
type TVTColumnWidthTrackingEvent func(sender IVTHeader, column TColumnIndex, shift TShiftState, trackPoint *TPoint, point TPoint, allowed *bool)
type TVTGetHeaderCursorEvent func(sender IVTHeader, cursor *HCURSOR)
type TVTBeforeGetMaxColumnWidthEvent func(sender IVTHeader, column TColumnIndex, useSmartColumnWidth *bool)
type TVTAfterGetMaxColumnWidthEvent func(sender IVTHeader, column TColumnIndex, maxWidth *int32)
type TVTCanSplitterResizeColumnEvent func(sender IVTHeader, point TPoint, column TColumnIndex, allowed *bool)
type TVTCanSplitterResizeHeaderEvent func(sender IVTHeader, point TPoint, allowed *bool)

// move, copy and node tracking events

type TVTNodeMovedEvent func(sender IBaseVirtualTree, node IVirtualNode)
type TVTNodeMovingEvent func(sender IBaseVirtualTree, node, target IVirtualNode, allowed *bool)
type TVTNodeCopiedEvent func(sender IBaseVirtualTree, node IVirtualNode)
type TVTNodeCopyingEvent func(sender IBaseVirtualTree, node, target IVirtualNode, allowed *bool)
type TVTNodeClickEvent func(sender IBaseVirtualTree, hitInfo THitInfo)
type TVTNodeHeightTrackingEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, shift TShiftState, trackPoint *TPoint, point TPoint, allowed *bool)
type TVTNodeHeightDblClickResizeEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, shift TShiftState, point TPoint, allowed *bool)
type TVTCanSplitterResizeNodeEvent func(sender IBaseVirtualTree, point TPoint, node IVirtualNode, column TColumnIndex, allowed *bool)

// drag'n drop/OLE

type TVTCreateDragManagerEvent func(sender IBaseVirtualTree, outDragManager *IVTDragManager)
type TVTCreateDataObjectEvent func(sender IBaseVirtualTree, outIDataObject *IDataObject)
type TVTDragAllowedEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, allowed *bool)
type TVTDragOverEvent func(sender IBaseVirtualTree, source IObject, shift TShiftState, state TDragState, point TPoint, mode TDropMode, effect *LongWord, accept *bool)
type TVTDragDropEvent func(sender IBaseVirtualTree, source IObject, dataObject IDataObject, formats *TFormatArray, shift TShiftState, point TPoint, effect *LongWord, mode TDropMode)

// TODO 还没搞明白这种事件怎么处理 TFormatEtc TFormatEtcArray TStgMedium
//type TVTRenderOLEDataEvent func(sender IBaseVirtualTree, FormatEtcIn TFormatEtc, outMedium TStgMedium, ForClipboard bool, Result HRESULT)
//type TVTGetUserClipboardFormatsEvent func(sender IBaseVirtualTree, Formats TFormatEtcArray)

// paint events

type TVTBeforeItemEraseEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, itemRect TRect, itemColor *TColor, eraseAction *TItemEraseAction)
type TVTAfterItemEraseEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, itemRect TRect)
type TVTBeforeItemPaintEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, itemRect TRect, customDraw *bool)
type TVTAfterItemPaintEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, itemRect TRect)
type TVTBeforeCellPaintEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, column TColumnIndex, cellPaintMode TVTCellPaintMode, cellRect TRect, contentRect *TRect)
type TVTAfterCellPaintEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, column TColumnIndex, cellRect TRect)
type TVTPaintEvent func(sender IBaseVirtualTree, targetCanvas ICanvas)
type TVTBackgroundPaintEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, rect TRect, handled *bool)
type TVTGetLineStyleEvent func(sender IBaseVirtualTree, bits *Pointer)
type TVTMeasureItemEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, nodeHeight *int32)

// search, sort

type TVTCompareEvent func(sender IBaseVirtualTree, node1, node2 IVirtualNode, column TColumnIndex, result *int32)
type TVTIncrementalSearchEvent func(sender IBaseVirtualTree, node IVirtualNode, searchText string, result *int32)

// operations

type TVTOperationEvent func(sender IBaseVirtualTree, operationKind TVTOperationKind)

type TVTHintKindEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, kind *TVTHintKind)
type TVTDrawHintEvent func(sender IBaseVirtualTree, hintCanvas ICanvas, node IVirtualNode, rect TRect, column TColumnIndex)
type TVTGetHintSizeEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, rect *TRect)

// miscellaneous

type TVTBeforeDrawLineImageEvent func(sender IBaseVirtualTree, node IVirtualNode, level int32, posX *int32)
type TVTGetNodeDataSizeEvent func(sender IBaseVirtualTree, nodeDataSize *int32)
type TVTKeyActionEvent func(sender IBaseVirtualTree, charCode *Word, shift *TShiftState, doDefault *bool)
type TVTScrollEvent func(sender IBaseVirtualTree, deltaX, deltaY int32)
type TVTUpdatingEvent func(sender IBaseVirtualTree, state TVTUpdateState)
type TVTGetCursorEvent func(sender IBaseVirtualTree, cursor *TCursor)
type TVTStateChangeEvent func(sender IBaseVirtualTree, enter, leave TVirtualTreeStates)
type TVTGetCellIsEmptyEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, isEmpty *bool)
type TVTScrollBarShowEvent func(sender IBaseVirtualTree, bar int32, show bool)

type TVTDrawTextEvent func(Sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, column TColumnIndex, cellText string, cellRect TRect, defaultDraw bool)
type TVSTGetTextEvent func(sender IBaseVirtualTree, nNode IVirtualNode, column TColumnIndex, textType TVSTTextType, cellText *string)
type TVTPaintText func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, column TColumnIndex, textType TVSTTextType)
type TVSTGetHintEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, lineBreakStyle *TVTTooltipLineBreakStyle, hintText *string)
type TVTMeasureTextEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, Node IVirtualNode, column TColumnIndex, cellText string, extent *int32)
type TVSTNewTextEvent func(sender IBaseVirtualTree, node IVirtualNode, column TColumnIndex, newText string)
type TVSTShortenstringEvent func(sender IBaseVirtualTree, targetCanvas ICanvas, node IVirtualNode, column TColumnIndex, str string, textSpace int32, result *string, done *bool)
