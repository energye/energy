//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/energy/v2/api"
	. "github.com/energye/energy/v2/types"
)

// ILazVirtualDrawTree Parent: ICustomVirtualDrawTree
type ILazVirtualDrawTree interface {
	ICustomVirtualDrawTree
	LastDragEffect() uint32                                                // property
	Alignment() TAlignment                                                 // property
	SetAlignment(AValue TAlignment)                                        // property
	AnimationDuration() uint32                                             // property
	SetAnimationDuration(AValue uint32)                                    // property
	AutoExpandDelay() uint32                                               // property
	SetAutoExpandDelay(AValue uint32)                                      // property
	AutoScrollDelay() uint32                                               // property
	SetAutoScrollDelay(AValue uint32)                                      // property
	AutoScrollInterval() TAutoScrollInterval                               // property
	SetAutoScrollInterval(AValue TAutoScrollInterval)                      // property
	Background() IPicture                                                  // property
	SetBackground(AValue IPicture)                                         // property
	BackgroundOffsetX() int32                                              // property
	SetBackgroundOffsetX(AValue int32)                                     // property
	BackgroundOffsetY() int32                                              // property
	SetBackgroundOffsetY(AValue int32)                                     // property
	BottomSpace() uint32                                                   // property
	SetBottomSpace(AValue uint32)                                          // property
	ButtonFillMode() TVTButtonFillMode                                     // property
	SetButtonFillMode(AValue TVTButtonFillMode)                            // property
	ButtonStyle() TVTButtonStyle                                           // property
	SetButtonStyle(AValue TVTButtonStyle)                                  // property
	ChangeDelay() uint32                                                   // property
	SetChangeDelay(AValue uint32)                                          // property
	CheckImageKind() TCheckImageKind                                       // property
	SetCheckImageKind(AValue TCheckImageKind)                              // property
	ClipboardFormats() IClipboardFormats                                   // property
	SetClipboardFormats(AValue IClipboardFormats)                          // property
	Colors() IVTColors                                                     // property
	SetColors(AValue IVTColors)                                            // property
	CustomCheckImages() ICustomImageList                                   // property
	SetCustomCheckImages(AValue ICustomImageList)                          // property
	DefaultNodeHeight() uint32                                             // property
	SetDefaultNodeHeight(AValue uint32)                                    // property
	DefaultPasteMode() TVTNodeAttachMode                                   // property
	SetDefaultPasteMode(AValue TVTNodeAttachMode)                          // property
	DragCursor() TCursor                                                   // property
	SetDragCursor(AValue TCursor)                                          // property
	DragHeight() int32                                                     // property
	SetDragHeight(AValue int32)                                            // property
	DragKind() TDragKind                                                   // property
	SetDragKind(AValue TDragKind)                                          // property
	DragImageKind() TVTDragImageKind                                       // property
	SetDragImageKind(AValue TVTDragImageKind)                              // property
	DragMode() TDragMode                                                   // property
	SetDragMode(AValue TDragMode)                                          // property
	DragOperations() TDragOperations                                       // property
	SetDragOperations(AValue TDragOperations)                              // property
	DragType() TVTDragType                                                 // property
	SetDragType(AValue TVTDragType)                                        // property
	DragWidth() int32                                                      // property
	SetDragWidth(AValue int32)                                             // property
	DrawSelectionMode() TVTDrawSelectionMode                               // property
	SetDrawSelectionMode(AValue TVTDrawSelectionMode)                      // property
	EditDelay() uint32                                                     // property
	SetEditDelay(AValue uint32)                                            // property
	Header() IVTHeader                                                     // property
	SetHeader(AValue IVTHeader)                                            // property
	HintMode() TVTHintMode                                                 // property
	SetHintMode(AValue TVTHintMode)                                        // property
	HotCursor() TCursor                                                    // property
	SetHotCursor(AValue TCursor)                                           // property
	Images() ICustomImageList                                              // property
	SetImages(AValue ICustomImageList)                                     // property
	IncrementalSearch() TVTIncrementalSearch                               // property
	SetIncrementalSearch(AValue TVTIncrementalSearch)                      // property
	IncrementalSearchDirection() TVTSearchDirection                        // property
	SetIncrementalSearchDirection(AValue TVTSearchDirection)               // property
	IncrementalSearchStart() TVTSearchStart                                // property
	SetIncrementalSearchStart(AValue TVTSearchStart)                       // property
	IncrementalSearchTimeout() uint32                                      // property
	SetIncrementalSearchTimeout(AValue uint32)                             // property
	Indent() uint32                                                        // property
	SetIndent(AValue uint32)                                               // property
	LineMode() TVTLineMode                                                 // property
	SetLineMode(AValue TVTLineMode)                                        // property
	LineStyle() TVTLineStyle                                               // property
	SetLineStyle(AValue TVTLineStyle)                                      // property
	Margin() int32                                                         // property
	SetMargin(AValue int32)                                                // property
	NodeAlignment() TVTNodeAlignment                                       // property
	SetNodeAlignment(AValue TVTNodeAlignment)                              // property
	NodeDataSize() int32                                                   // property
	SetNodeDataSize(AValue int32)                                          // property
	OperationCanceled() bool                                               // property
	ParentColor() bool                                                     // property
	SetParentColor(AValue bool)                                            // property
	ParentFont() bool                                                      // property
	SetParentFont(AValue bool)                                             // property
	ParentShowHint() bool                                                  // property
	SetParentShowHint(AValue bool)                                         // property
	RootNodeCount() uint32                                                 // property
	SetRootNodeCount(AValue uint32)                                        // property
	ScrollBarOptions() IScrollBarOptions                                   // property
	SetScrollBarOptions(AValue IScrollBarOptions)                          // property
	SelectionBlendFactor() Byte                                            // property
	SetSelectionBlendFactor(AValue Byte)                                   // property
	SelectionCurveRadius() uint32                                          // property
	SetSelectionCurveRadius(AValue uint32)                                 // property
	StateImages() ICustomImageList                                         // property
	SetStateImages(AValue ICustomImageList)                                // property
	TextMargin() int32                                                     // property
	SetTextMargin(AValue int32)                                            // property
	TreeOptions() IVirtualTreeOptions                                      // property
	SetTreeOptions(AValue IVirtualTreeOptions)                             // property
	WantTabs() bool                                                        // property
	SetWantTabs(AValue bool)                                               // property
	SetOnAddToSelection(fn TVTAddToSelectionEvent)                         // property event
	SetOnAdvancedHeaderDraw(fn TVTAdvancedHeaderPaintEvent)                // property event
	SetOnAfterAutoFitColumn(fn TVTAfterAutoFitColumnEvent)                 // property event
	SetOnAfterAutoFitColumns(fn TVTAfterAutoFitColumnsEvent)               // property event
	SetOnAfterCellPaint(fn TVTAfterCellPaintEvent)                         // property event
	SetOnAfterColumnExport(fn TVTColumnExportEvent)                        // property event
	SetOnAfterColumnWidthTracking(fn TVTAfterColumnWidthTrackingEvent)     // property event
	SetOnAfterGetMaxColumnWidth(fn TVTAfterGetMaxColumnWidthEvent)         // property event
	SetOnAfterHeaderExport(fn TVTTreeExportEvent)                          // property event
	SetOnAfterHeaderHeightTracking(fn TVTAfterHeaderHeightTrackingEvent)   // property event
	SetOnAfterItemErase(fn TVTAfterItemEraseEvent)                         // property event
	SetOnAfterItemPaint(fn TVTAfterItemPaintEvent)                         // property event
	SetOnAfterNodeExport(fn TVTNodeExportEvent)                            // property event
	SetOnAfterPaint(fn TVTPaintEvent)                                      // property event
	SetOnAfterTreeExport(fn TVTTreeExportEvent)                            // property event
	SetOnBeforeAutoFitColumn(fn TVTBeforeAutoFitColumnEvent)               // property event
	SetOnBeforeAutoFitColumns(fn TVTBeforeAutoFitColumnsEvent)             // property event
	SetOnBeforeCellPaint(fn TVTBeforeCellPaintEvent)                       // property event
	SetOnBeforeColumnExport(fn TVTColumnExportEvent)                       // property event
	SetOnBeforeColumnWidthTracking(fn TVTBeforeColumnWidthTrackingEvent)   // property event
	SetOnBeforeDrawTreeLine(fn TVTBeforeDrawLineImageEvent)                // property event
	SetOnBeforeGetMaxColumnWidth(fn TVTBeforeGetMaxColumnWidthEvent)       // property event
	SetOnBeforeHeaderExport(fn TVTTreeExportEvent)                         // property event
	SetOnBeforeHeaderHeightTracking(fn TVTBeforeHeaderHeightTrackingEvent) // property event
	SetOnBeforeItemErase(fn TVTBeforeItemEraseEvent)                       // property event
	SetOnBeforeItemPaint(fn TVTBeforeItemPaintEvent)                       // property event
	SetOnBeforeNodeExport(fn TVTNodeExportEvent)                           // property event
	SetOnBeforePaint(fn TVTPaintEvent)                                     // property event
	SetOnBeforeTreeExport(fn TVTTreeExportEvent)                           // property event
	SetOnCanSplitterResizeColumn(fn TVTCanSplitterResizeColumnEvent)       // property event
	SetOnCanSplitterResizeHeader(fn TVTCanSplitterResizeHeaderEvent)       // property event
	SetOnCanSplitterResizeNode(fn TVTCanSplitterResizeNodeEvent)           // property event
	SetOnChange(fn TVTChangeEvent)                                         // property event
	SetOnChecked(fn TVTChangeEvent)                                        // property event
	SetOnChecking(fn TVTCheckChangingEvent)                                // property event
	SetOnCollapsed(fn TVTChangeEvent)                                      // property event
	SetOnCollapsing(fn TVTChangingEvent)                                   // property event
	SetOnColumnClick(fn TVTColumnClickEvent)                               // property event
	SetOnColumnDblClick(fn TVTColumnDblClickEvent)                         // property event
	SetOnColumnExport(fn TVTColumnExportEvent)                             // property event
	SetOnColumnResize(fn TVTHeaderNotifyEvent)                             // property event
	SetOnColumnWidthDblClickResize(fn TVTColumnWidthDblClickResizeEvent)   // property event
	SetOnColumnWidthTracking(fn TVTColumnWidthTrackingEvent)               // property event
	SetOnCompareNodes(fn TVTCompareEvent)                                  // property event
	SetOnContextPopup(fn TContextPopupEvent)                               // property event
	SetOnCreateDataObject(fn TVTCreateDataObjectEvent)                     // property event
	SetOnCreateDragManager(fn TVTCreateDragManagerEvent)                   // property event
	SetOnCreateEditor(fn TVTCreateEditorEvent)                             // property event
	SetOnDblClick(fn TNotifyEvent)                                         // property event
	SetOnDragAllowed(fn TVTDragAllowedEvent)                               // property event
	SetOnDragOver(fn TVTDragOverEvent)                                     // property event
	SetOnDragDrop(fn TVTDragDropEvent)                                     // property event
	SetOnDrawHint(fn TVTDrawHintEvent)                                     // property event
	SetOnDrawNode(fn TVTDrawNodeEvent)                                     // property event
	SetOnEdited(fn TVTEditChangeEvent)                                     // property event
	SetOnEditing(fn TVTEditChangingEvent)                                  // property event
	SetOnEndDock(fn TEndDragEvent)                                         // property event
	SetOnEndDrag(fn TEndDragEvent)                                         // property event
	SetOnEndOperation(fn TVTOperationEvent)                                // property event
	SetOnExpanded(fn TVTChangeEvent)                                       // property event
	SetOnExpanding(fn TVTChangingEvent)                                    // property event
	SetOnFocusChanged(fn TVTFocusChangeEvent)                              // property event
	SetOnFocusChanging(fn TVTFocusChangingEvent)                           // property event
	SetOnFreeNode(fn TVTFreeNodeEvent)                                     // property event
	SetOnGetCellIsEmpty(fn TVTGetCellIsEmptyEvent)                         // property event
	SetOnGetCursor(fn TVTGetCursorEvent)                                   // property event
	SetOnGetHeaderCursor(fn TVTGetHeaderCursorEvent)                       // property event
	SetOnGetHelpContext(fn TVTHelpContextEvent)                            // property event
	SetOnGetHintKind(fn TVTHintKindEvent)                                  // property event
	SetOnGetHintSize(fn TVTGetHintSizeEvent)                               // property event
	SetOnGetImageIndex(fn TVTGetImageEvent)                                // property event
	SetOnGetImageIndexEx(fn TVTGetImageExEvent)                            // property event
	SetOnGetLineStyle(fn TVTGetLineStyleEvent)                             // property event
	SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent)                       // property event
	SetOnGetNodeWidth(fn TVTGetNodeWidthEvent)                             // property event
	SetOnGetPopupMenu(fn TVTPopupEvent)                                    // property event
	SetOnHeaderClick(fn TVTHeaderClickEvent)                               // property event
	SetOnHeaderDblClick(fn TVTHeaderClickEvent)                            // property event
	SetOnHeaderDragged(fn TVTHeaderDraggedEvent)                           // property event
	SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent)                     // property event
	SetOnHeaderDragging(fn TVTHeaderDraggingEvent)                         // property event
	SetOnHeaderDraw(fn TVTHeaderPaintEvent)                                // property event
	SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent)      // property event
	SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent)             // property event
	SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) // property event
	SetOnHeaderMouseDown(fn TVTHeaderMouseEvent)                           // property event
	SetOnHeaderMouseMove(fn TVTHeaderMouseMoveEvent)                       // property event
	SetOnHeaderMouseUp(fn TVTHeaderMouseEvent)                             // property event
	SetOnHotChange(fn TVTHotNodeChangeEvent)                               // property event
	SetOnIncrementalSearch(fn TVTIncrementalSearchEvent)                   // property event
	SetOnInitChildren(fn TVTInitChildrenEvent)                             // property event
	SetOnInitNode(fn TVTInitNodeEvent)                                     // property event
	SetOnKeyAction(fn TVTKeyActionEvent)                                   // property event
	SetOnLoadNode(fn TVTSaveNodeEvent)                                     // property event
	SetOnLoadTree(fn TVTSaveTreeEvent)                                     // property event
	SetOnMeasureItem(fn TVTMeasureItemEvent)                               // property event
	SetOnMouseDown(fn TMouseEvent)                                         // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                     // property event
	SetOnMouseUp(fn TMouseEvent)                                           // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                   // property event
	SetOnNodeClick(fn TVTNodeClickEvent)                                   // property event
	SetOnNodeCopied(fn TVTNodeCopiedEvent)                                 // property event
	SetOnNodeCopying(fn TVTNodeCopyingEvent)                               // property event
	SetOnNodeDblClick(fn TVTNodeClickEvent)                                // property event
	SetOnNodeExport(fn TVTNodeExportEvent)                                 // property event
	SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent)                 // property event
	SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent)     // property event
	SetOnNodeMoved(fn TVTNodeMovedEvent)                                   // property event
	SetOnNodeMoving(fn TVTNodeMovingEvent)                                 // property event
	SetOnPaintBackground(fn TVTBackgroundPaintEvent)                       // property event
	SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent)               // property event
	SetOnResetNode(fn TVTChangeEvent)                                      // property event
	SetOnSaveNode(fn TVTSaveNodeEvent)                                     // property event
	SetOnSaveTree(fn TVTSaveTreeEvent)                                     // property event
	SetOnScroll(fn TVTScrollEvent)                                         // property event
	SetOnShowScrollBar(fn TVTScrollBarShowEvent)                           // property event
	SetOnStartDock(fn TStartDockEvent)                                     // property event
	SetOnStartDrag(fn TStartDragEvent)                                     // property event
	SetOnStartOperation(fn TVTOperationEvent)                              // property event
	SetOnStateChange(fn TVTStateChangeEvent)                               // property event
	SetOnStructureChange(fn TVTStructureChangeEvent)                       // property event
	SetOnUpdating(fn TVTUpdatingEvent)                                     // property event
}

// TLazVirtualDrawTree Parent: TCustomVirtualDrawTree
type TLazVirtualDrawTree struct {
	TCustomVirtualDrawTree
	addToSelectionPtr             uintptr
	advancedHeaderDrawPtr         uintptr
	afterAutoFitColumnPtr         uintptr
	afterAutoFitColumnsPtr        uintptr
	afterCellPaintPtr             uintptr
	afterColumnExportPtr          uintptr
	afterColumnWidthTrackingPtr   uintptr
	afterGetMaxColumnWidthPtr     uintptr
	afterHeaderExportPtr          uintptr
	afterHeaderHeightTrackingPtr  uintptr
	afterItemErasePtr             uintptr
	afterItemPaintPtr             uintptr
	afterNodeExportPtr            uintptr
	afterPaintPtr                 uintptr
	afterTreeExportPtr            uintptr
	beforeAutoFitColumnPtr        uintptr
	beforeAutoFitColumnsPtr       uintptr
	beforeCellPaintPtr            uintptr
	beforeColumnExportPtr         uintptr
	beforeColumnWidthTrackingPtr  uintptr
	beforeDrawTreeLinePtr         uintptr
	beforeGetMaxColumnWidthPtr    uintptr
	beforeHeaderExportPtr         uintptr
	beforeHeaderHeightTrackingPtr uintptr
	beforeItemErasePtr            uintptr
	beforeItemPaintPtr            uintptr
	beforeNodeExportPtr           uintptr
	beforePaintPtr                uintptr
	beforeTreeExportPtr           uintptr
	canSplitterResizeColumnPtr    uintptr
	canSplitterResizeHeaderPtr    uintptr
	canSplitterResizeNodePtr      uintptr
	changePtr                     uintptr
	checkedPtr                    uintptr
	checkingPtr                   uintptr
	collapsedPtr                  uintptr
	collapsingPtr                 uintptr
	columnClickPtr                uintptr
	columnDblClickPtr             uintptr
	columnExportPtr               uintptr
	columnResizePtr               uintptr
	columnWidthDblClickResizePtr  uintptr
	columnWidthTrackingPtr        uintptr
	compareNodesPtr               uintptr
	contextPopupPtr               uintptr
	createDataObjectPtr           uintptr
	createDragManagerPtr          uintptr
	createEditorPtr               uintptr
	dblClickPtr                   uintptr
	dragAllowedPtr                uintptr
	dragOverPtr                   uintptr
	dragDropPtr                   uintptr
	drawHintPtr                   uintptr
	drawNodePtr                   uintptr
	editedPtr                     uintptr
	editingPtr                    uintptr
	endDockPtr                    uintptr
	endDragPtr                    uintptr
	endOperationPtr               uintptr
	expandedPtr                   uintptr
	expandingPtr                  uintptr
	focusChangedPtr               uintptr
	focusChangingPtr              uintptr
	freeNodePtr                   uintptr
	getCellIsEmptyPtr             uintptr
	getCursorPtr                  uintptr
	getHeaderCursorPtr            uintptr
	getHelpContextPtr             uintptr
	getHintKindPtr                uintptr
	getHintSizePtr                uintptr
	getImageIndexPtr              uintptr
	getImageIndexExPtr            uintptr
	getLineStylePtr               uintptr
	getNodeDataSizePtr            uintptr
	getNodeWidthPtr               uintptr
	getPopupMenuPtr               uintptr
	headerClickPtr                uintptr
	headerDblClickPtr             uintptr
	headerDraggedPtr              uintptr
	headerDraggedOutPtr           uintptr
	headerDraggingPtr             uintptr
	headerDrawPtr                 uintptr
	headerDrawQueryElementsPtr    uintptr
	headerHeightTrackingPtr       uintptr
	headerHeightDblClickResizePtr uintptr
	headerMouseDownPtr            uintptr
	headerMouseMovePtr            uintptr
	headerMouseUpPtr              uintptr
	hotChangePtr                  uintptr
	incrementalSearchPtr          uintptr
	initChildrenPtr               uintptr
	initNodePtr                   uintptr
	keyActionPtr                  uintptr
	loadNodePtr                   uintptr
	loadTreePtr                   uintptr
	measureItemPtr                uintptr
	mouseDownPtr                  uintptr
	mouseMovePtr                  uintptr
	mouseUpPtr                    uintptr
	mouseWheelPtr                 uintptr
	nodeClickPtr                  uintptr
	nodeCopiedPtr                 uintptr
	nodeCopyingPtr                uintptr
	nodeDblClickPtr               uintptr
	nodeExportPtr                 uintptr
	nodeHeightTrackingPtr         uintptr
	nodeHeightDblClickResizePtr   uintptr
	nodeMovedPtr                  uintptr
	nodeMovingPtr                 uintptr
	paintBackgroundPtr            uintptr
	removeFromSelectionPtr        uintptr
	resetNodePtr                  uintptr
	saveNodePtr                   uintptr
	saveTreePtr                   uintptr
	scrollPtr                     uintptr
	showScrollBarPtr              uintptr
	startDockPtr                  uintptr
	startDragPtr                  uintptr
	startOperationPtr             uintptr
	stateChangePtr                uintptr
	structureChangePtr            uintptr
	updatingPtr                   uintptr
}

func NewLazVirtualDrawTree(AOwner IComponent) ILazVirtualDrawTree {
	r1 := LCL().SysCallN(3593, GetObjectUintptr(AOwner))
	return AsLazVirtualDrawTree(r1)
}

func (m *TLazVirtualDrawTree) LastDragEffect() uint32 {
	r1 := LCL().SysCallN(3616, m.Instance())
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) Alignment() TAlignment {
	r1 := LCL().SysCallN(3577, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TLazVirtualDrawTree) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3577, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) AnimationDuration() uint32 {
	r1 := LCL().SysCallN(3578, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetAnimationDuration(AValue uint32) {
	LCL().SysCallN(3578, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) AutoExpandDelay() uint32 {
	r1 := LCL().SysCallN(3579, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetAutoExpandDelay(AValue uint32) {
	LCL().SysCallN(3579, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) AutoScrollDelay() uint32 {
	r1 := LCL().SysCallN(3580, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetAutoScrollDelay(AValue uint32) {
	LCL().SysCallN(3580, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) AutoScrollInterval() TAutoScrollInterval {
	r1 := LCL().SysCallN(3581, 0, m.Instance(), 0)
	return TAutoScrollInterval(r1)
}

func (m *TLazVirtualDrawTree) SetAutoScrollInterval(AValue TAutoScrollInterval) {
	LCL().SysCallN(3581, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Background() IPicture {
	r1 := LCL().SysCallN(3582, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TLazVirtualDrawTree) SetBackground(AValue IPicture) {
	LCL().SysCallN(3582, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) BackgroundOffsetX() int32 {
	r1 := LCL().SysCallN(3583, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetBackgroundOffsetX(AValue int32) {
	LCL().SysCallN(3583, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) BackgroundOffsetY() int32 {
	r1 := LCL().SysCallN(3584, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetBackgroundOffsetY(AValue int32) {
	LCL().SysCallN(3584, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) BottomSpace() uint32 {
	r1 := LCL().SysCallN(3585, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetBottomSpace(AValue uint32) {
	LCL().SysCallN(3585, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ButtonFillMode() TVTButtonFillMode {
	r1 := LCL().SysCallN(3586, 0, m.Instance(), 0)
	return TVTButtonFillMode(r1)
}

func (m *TLazVirtualDrawTree) SetButtonFillMode(AValue TVTButtonFillMode) {
	LCL().SysCallN(3586, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ButtonStyle() TVTButtonStyle {
	r1 := LCL().SysCallN(3587, 0, m.Instance(), 0)
	return TVTButtonStyle(r1)
}

func (m *TLazVirtualDrawTree) SetButtonStyle(AValue TVTButtonStyle) {
	LCL().SysCallN(3587, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ChangeDelay() uint32 {
	r1 := LCL().SysCallN(3588, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetChangeDelay(AValue uint32) {
	LCL().SysCallN(3588, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) CheckImageKind() TCheckImageKind {
	r1 := LCL().SysCallN(3589, 0, m.Instance(), 0)
	return TCheckImageKind(r1)
}

func (m *TLazVirtualDrawTree) SetCheckImageKind(AValue TCheckImageKind) {
	LCL().SysCallN(3589, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ClipboardFormats() IClipboardFormats {
	r1 := LCL().SysCallN(3591, 0, m.Instance(), 0)
	return AsClipboardFormats(r1)
}

func (m *TLazVirtualDrawTree) SetClipboardFormats(AValue IClipboardFormats) {
	LCL().SysCallN(3591, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) Colors() IVTColors {
	r1 := LCL().SysCallN(3592, 0, m.Instance(), 0)
	return AsVTColors(r1)
}

func (m *TLazVirtualDrawTree) SetColors(AValue IVTColors) {
	LCL().SysCallN(3592, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) CustomCheckImages() ICustomImageList {
	r1 := LCL().SysCallN(3594, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualDrawTree) SetCustomCheckImages(AValue ICustomImageList) {
	LCL().SysCallN(3594, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) DefaultNodeHeight() uint32 {
	r1 := LCL().SysCallN(3595, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetDefaultNodeHeight(AValue uint32) {
	LCL().SysCallN(3595, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DefaultPasteMode() TVTNodeAttachMode {
	r1 := LCL().SysCallN(3596, 0, m.Instance(), 0)
	return TVTNodeAttachMode(r1)
}

func (m *TLazVirtualDrawTree) SetDefaultPasteMode(AValue TVTNodeAttachMode) {
	LCL().SysCallN(3596, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragCursor() TCursor {
	r1 := LCL().SysCallN(3597, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLazVirtualDrawTree) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3597, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragHeight() int32 {
	r1 := LCL().SysCallN(3598, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetDragHeight(AValue int32) {
	LCL().SysCallN(3598, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragKind() TDragKind {
	r1 := LCL().SysCallN(3600, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TLazVirtualDrawTree) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3600, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragImageKind() TVTDragImageKind {
	r1 := LCL().SysCallN(3599, 0, m.Instance(), 0)
	return TVTDragImageKind(r1)
}

func (m *TLazVirtualDrawTree) SetDragImageKind(AValue TVTDragImageKind) {
	LCL().SysCallN(3599, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragMode() TDragMode {
	r1 := LCL().SysCallN(3601, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLazVirtualDrawTree) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3601, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragOperations() TDragOperations {
	r1 := LCL().SysCallN(3602, 0, m.Instance(), 0)
	return TDragOperations(r1)
}

func (m *TLazVirtualDrawTree) SetDragOperations(AValue TDragOperations) {
	LCL().SysCallN(3602, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragType() TVTDragType {
	r1 := LCL().SysCallN(3603, 0, m.Instance(), 0)
	return TVTDragType(r1)
}

func (m *TLazVirtualDrawTree) SetDragType(AValue TVTDragType) {
	LCL().SysCallN(3603, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DragWidth() int32 {
	r1 := LCL().SysCallN(3604, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetDragWidth(AValue int32) {
	LCL().SysCallN(3604, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) DrawSelectionMode() TVTDrawSelectionMode {
	r1 := LCL().SysCallN(3605, 0, m.Instance(), 0)
	return TVTDrawSelectionMode(r1)
}

func (m *TLazVirtualDrawTree) SetDrawSelectionMode(AValue TVTDrawSelectionMode) {
	LCL().SysCallN(3605, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) EditDelay() uint32 {
	r1 := LCL().SysCallN(3606, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetEditDelay(AValue uint32) {
	LCL().SysCallN(3606, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Header() IVTHeader {
	r1 := LCL().SysCallN(3607, 0, m.Instance(), 0)
	return AsVTHeader(r1)
}

func (m *TLazVirtualDrawTree) SetHeader(AValue IVTHeader) {
	LCL().SysCallN(3607, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) HintMode() TVTHintMode {
	r1 := LCL().SysCallN(3608, 0, m.Instance(), 0)
	return TVTHintMode(r1)
}

func (m *TLazVirtualDrawTree) SetHintMode(AValue TVTHintMode) {
	LCL().SysCallN(3608, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) HotCursor() TCursor {
	r1 := LCL().SysCallN(3609, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLazVirtualDrawTree) SetHotCursor(AValue TCursor) {
	LCL().SysCallN(3609, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Images() ICustomImageList {
	r1 := LCL().SysCallN(3610, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualDrawTree) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(3610, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) IncrementalSearch() TVTIncrementalSearch {
	r1 := LCL().SysCallN(3611, 0, m.Instance(), 0)
	return TVTIncrementalSearch(r1)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearch(AValue TVTIncrementalSearch) {
	LCL().SysCallN(3611, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) IncrementalSearchDirection() TVTSearchDirection {
	r1 := LCL().SysCallN(3612, 0, m.Instance(), 0)
	return TVTSearchDirection(r1)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchDirection(AValue TVTSearchDirection) {
	LCL().SysCallN(3612, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) IncrementalSearchStart() TVTSearchStart {
	r1 := LCL().SysCallN(3613, 0, m.Instance(), 0)
	return TVTSearchStart(r1)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchStart(AValue TVTSearchStart) {
	LCL().SysCallN(3613, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) IncrementalSearchTimeout() uint32 {
	r1 := LCL().SysCallN(3614, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetIncrementalSearchTimeout(AValue uint32) {
	LCL().SysCallN(3614, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Indent() uint32 {
	r1 := LCL().SysCallN(3615, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetIndent(AValue uint32) {
	LCL().SysCallN(3615, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) LineMode() TVTLineMode {
	r1 := LCL().SysCallN(3617, 0, m.Instance(), 0)
	return TVTLineMode(r1)
}

func (m *TLazVirtualDrawTree) SetLineMode(AValue TVTLineMode) {
	LCL().SysCallN(3617, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) LineStyle() TVTLineStyle {
	r1 := LCL().SysCallN(3618, 0, m.Instance(), 0)
	return TVTLineStyle(r1)
}

func (m *TLazVirtualDrawTree) SetLineStyle(AValue TVTLineStyle) {
	LCL().SysCallN(3618, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) Margin() int32 {
	r1 := LCL().SysCallN(3619, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetMargin(AValue int32) {
	LCL().SysCallN(3619, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) NodeAlignment() TVTNodeAlignment {
	r1 := LCL().SysCallN(3620, 0, m.Instance(), 0)
	return TVTNodeAlignment(r1)
}

func (m *TLazVirtualDrawTree) SetNodeAlignment(AValue TVTNodeAlignment) {
	LCL().SysCallN(3620, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) NodeDataSize() int32 {
	r1 := LCL().SysCallN(3621, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetNodeDataSize(AValue int32) {
	LCL().SysCallN(3621, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) OperationCanceled() bool {
	r1 := LCL().SysCallN(3622, m.Instance())
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) ParentColor() bool {
	r1 := LCL().SysCallN(3623, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) SetParentColor(AValue bool) {
	LCL().SysCallN(3623, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualDrawTree) ParentFont() bool {
	r1 := LCL().SysCallN(3624, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) SetParentFont(AValue bool) {
	LCL().SysCallN(3624, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualDrawTree) ParentShowHint() bool {
	r1 := LCL().SysCallN(3625, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3625, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualDrawTree) RootNodeCount() uint32 {
	r1 := LCL().SysCallN(3626, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetRootNodeCount(AValue uint32) {
	LCL().SysCallN(3626, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) ScrollBarOptions() IScrollBarOptions {
	r1 := LCL().SysCallN(3627, 0, m.Instance(), 0)
	return AsScrollBarOptions(r1)
}

func (m *TLazVirtualDrawTree) SetScrollBarOptions(AValue IScrollBarOptions) {
	LCL().SysCallN(3627, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) SelectionBlendFactor() Byte {
	r1 := LCL().SysCallN(3628, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TLazVirtualDrawTree) SetSelectionBlendFactor(AValue Byte) {
	LCL().SysCallN(3628, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) SelectionCurveRadius() uint32 {
	r1 := LCL().SysCallN(3629, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualDrawTree) SetSelectionCurveRadius(AValue uint32) {
	LCL().SysCallN(3629, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) StateImages() ICustomImageList {
	r1 := LCL().SysCallN(3752, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualDrawTree) SetStateImages(AValue ICustomImageList) {
	LCL().SysCallN(3752, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) TextMargin() int32 {
	r1 := LCL().SysCallN(3753, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualDrawTree) SetTextMargin(AValue int32) {
	LCL().SysCallN(3753, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualDrawTree) TreeOptions() IVirtualTreeOptions {
	r1 := LCL().SysCallN(3754, 0, m.Instance(), 0)
	return AsVirtualTreeOptions(r1)
}

func (m *TLazVirtualDrawTree) SetTreeOptions(AValue IVirtualTreeOptions) {
	LCL().SysCallN(3754, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualDrawTree) WantTabs() bool {
	r1 := LCL().SysCallN(3755, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualDrawTree) SetWantTabs(AValue bool) {
	LCL().SysCallN(3755, 1, m.Instance(), PascalBool(AValue))
}

func LazVirtualDrawTreeClass() TClass {
	ret := LCL().SysCallN(3590)
	return TClass(ret)
}

func (m *TLazVirtualDrawTree) SetOnAddToSelection(fn TVTAddToSelectionEvent) {
	if m.addToSelectionPtr != 0 {
		RemoveEventElement(m.addToSelectionPtr)
	}
	m.addToSelectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3630, m.Instance(), m.addToSelectionPtr)
}

func (m *TLazVirtualDrawTree) SetOnAdvancedHeaderDraw(fn TVTAdvancedHeaderPaintEvent) {
	if m.advancedHeaderDrawPtr != 0 {
		RemoveEventElement(m.advancedHeaderDrawPtr)
	}
	m.advancedHeaderDrawPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3631, m.Instance(), m.advancedHeaderDrawPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterAutoFitColumn(fn TVTAfterAutoFitColumnEvent) {
	if m.afterAutoFitColumnPtr != 0 {
		RemoveEventElement(m.afterAutoFitColumnPtr)
	}
	m.afterAutoFitColumnPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3632, m.Instance(), m.afterAutoFitColumnPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterAutoFitColumns(fn TVTAfterAutoFitColumnsEvent) {
	if m.afterAutoFitColumnsPtr != 0 {
		RemoveEventElement(m.afterAutoFitColumnsPtr)
	}
	m.afterAutoFitColumnsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3633, m.Instance(), m.afterAutoFitColumnsPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterCellPaint(fn TVTAfterCellPaintEvent) {
	if m.afterCellPaintPtr != 0 {
		RemoveEventElement(m.afterCellPaintPtr)
	}
	m.afterCellPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3634, m.Instance(), m.afterCellPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterColumnExport(fn TVTColumnExportEvent) {
	if m.afterColumnExportPtr != 0 {
		RemoveEventElement(m.afterColumnExportPtr)
	}
	m.afterColumnExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3635, m.Instance(), m.afterColumnExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterColumnWidthTracking(fn TVTAfterColumnWidthTrackingEvent) {
	if m.afterColumnWidthTrackingPtr != 0 {
		RemoveEventElement(m.afterColumnWidthTrackingPtr)
	}
	m.afterColumnWidthTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3636, m.Instance(), m.afterColumnWidthTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterGetMaxColumnWidth(fn TVTAfterGetMaxColumnWidthEvent) {
	if m.afterGetMaxColumnWidthPtr != 0 {
		RemoveEventElement(m.afterGetMaxColumnWidthPtr)
	}
	m.afterGetMaxColumnWidthPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3637, m.Instance(), m.afterGetMaxColumnWidthPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterHeaderExport(fn TVTTreeExportEvent) {
	if m.afterHeaderExportPtr != 0 {
		RemoveEventElement(m.afterHeaderExportPtr)
	}
	m.afterHeaderExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3638, m.Instance(), m.afterHeaderExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterHeaderHeightTracking(fn TVTAfterHeaderHeightTrackingEvent) {
	if m.afterHeaderHeightTrackingPtr != 0 {
		RemoveEventElement(m.afterHeaderHeightTrackingPtr)
	}
	m.afterHeaderHeightTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3639, m.Instance(), m.afterHeaderHeightTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterItemErase(fn TVTAfterItemEraseEvent) {
	if m.afterItemErasePtr != 0 {
		RemoveEventElement(m.afterItemErasePtr)
	}
	m.afterItemErasePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3640, m.Instance(), m.afterItemErasePtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterItemPaint(fn TVTAfterItemPaintEvent) {
	if m.afterItemPaintPtr != 0 {
		RemoveEventElement(m.afterItemPaintPtr)
	}
	m.afterItemPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3641, m.Instance(), m.afterItemPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterNodeExport(fn TVTNodeExportEvent) {
	if m.afterNodeExportPtr != 0 {
		RemoveEventElement(m.afterNodeExportPtr)
	}
	m.afterNodeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3642, m.Instance(), m.afterNodeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterPaint(fn TVTPaintEvent) {
	if m.afterPaintPtr != 0 {
		RemoveEventElement(m.afterPaintPtr)
	}
	m.afterPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3643, m.Instance(), m.afterPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnAfterTreeExport(fn TVTTreeExportEvent) {
	if m.afterTreeExportPtr != 0 {
		RemoveEventElement(m.afterTreeExportPtr)
	}
	m.afterTreeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3644, m.Instance(), m.afterTreeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeAutoFitColumn(fn TVTBeforeAutoFitColumnEvent) {
	if m.beforeAutoFitColumnPtr != 0 {
		RemoveEventElement(m.beforeAutoFitColumnPtr)
	}
	m.beforeAutoFitColumnPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3645, m.Instance(), m.beforeAutoFitColumnPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeAutoFitColumns(fn TVTBeforeAutoFitColumnsEvent) {
	if m.beforeAutoFitColumnsPtr != 0 {
		RemoveEventElement(m.beforeAutoFitColumnsPtr)
	}
	m.beforeAutoFitColumnsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3646, m.Instance(), m.beforeAutoFitColumnsPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeCellPaint(fn TVTBeforeCellPaintEvent) {
	if m.beforeCellPaintPtr != 0 {
		RemoveEventElement(m.beforeCellPaintPtr)
	}
	m.beforeCellPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3647, m.Instance(), m.beforeCellPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeColumnExport(fn TVTColumnExportEvent) {
	if m.beforeColumnExportPtr != 0 {
		RemoveEventElement(m.beforeColumnExportPtr)
	}
	m.beforeColumnExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3648, m.Instance(), m.beforeColumnExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeColumnWidthTracking(fn TVTBeforeColumnWidthTrackingEvent) {
	if m.beforeColumnWidthTrackingPtr != 0 {
		RemoveEventElement(m.beforeColumnWidthTrackingPtr)
	}
	m.beforeColumnWidthTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3649, m.Instance(), m.beforeColumnWidthTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeDrawTreeLine(fn TVTBeforeDrawLineImageEvent) {
	if m.beforeDrawTreeLinePtr != 0 {
		RemoveEventElement(m.beforeDrawTreeLinePtr)
	}
	m.beforeDrawTreeLinePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3650, m.Instance(), m.beforeDrawTreeLinePtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeGetMaxColumnWidth(fn TVTBeforeGetMaxColumnWidthEvent) {
	if m.beforeGetMaxColumnWidthPtr != 0 {
		RemoveEventElement(m.beforeGetMaxColumnWidthPtr)
	}
	m.beforeGetMaxColumnWidthPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3651, m.Instance(), m.beforeGetMaxColumnWidthPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeHeaderExport(fn TVTTreeExportEvent) {
	if m.beforeHeaderExportPtr != 0 {
		RemoveEventElement(m.beforeHeaderExportPtr)
	}
	m.beforeHeaderExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3652, m.Instance(), m.beforeHeaderExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeHeaderHeightTracking(fn TVTBeforeHeaderHeightTrackingEvent) {
	if m.beforeHeaderHeightTrackingPtr != 0 {
		RemoveEventElement(m.beforeHeaderHeightTrackingPtr)
	}
	m.beforeHeaderHeightTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3653, m.Instance(), m.beforeHeaderHeightTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeItemErase(fn TVTBeforeItemEraseEvent) {
	if m.beforeItemErasePtr != 0 {
		RemoveEventElement(m.beforeItemErasePtr)
	}
	m.beforeItemErasePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3654, m.Instance(), m.beforeItemErasePtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeItemPaint(fn TVTBeforeItemPaintEvent) {
	if m.beforeItemPaintPtr != 0 {
		RemoveEventElement(m.beforeItemPaintPtr)
	}
	m.beforeItemPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3655, m.Instance(), m.beforeItemPaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeNodeExport(fn TVTNodeExportEvent) {
	if m.beforeNodeExportPtr != 0 {
		RemoveEventElement(m.beforeNodeExportPtr)
	}
	m.beforeNodeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3656, m.Instance(), m.beforeNodeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforePaint(fn TVTPaintEvent) {
	if m.beforePaintPtr != 0 {
		RemoveEventElement(m.beforePaintPtr)
	}
	m.beforePaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3657, m.Instance(), m.beforePaintPtr)
}

func (m *TLazVirtualDrawTree) SetOnBeforeTreeExport(fn TVTTreeExportEvent) {
	if m.beforeTreeExportPtr != 0 {
		RemoveEventElement(m.beforeTreeExportPtr)
	}
	m.beforeTreeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3658, m.Instance(), m.beforeTreeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeColumn(fn TVTCanSplitterResizeColumnEvent) {
	if m.canSplitterResizeColumnPtr != 0 {
		RemoveEventElement(m.canSplitterResizeColumnPtr)
	}
	m.canSplitterResizeColumnPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3659, m.Instance(), m.canSplitterResizeColumnPtr)
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeHeader(fn TVTCanSplitterResizeHeaderEvent) {
	if m.canSplitterResizeHeaderPtr != 0 {
		RemoveEventElement(m.canSplitterResizeHeaderPtr)
	}
	m.canSplitterResizeHeaderPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3660, m.Instance(), m.canSplitterResizeHeaderPtr)
}

func (m *TLazVirtualDrawTree) SetOnCanSplitterResizeNode(fn TVTCanSplitterResizeNodeEvent) {
	if m.canSplitterResizeNodePtr != 0 {
		RemoveEventElement(m.canSplitterResizeNodePtr)
	}
	m.canSplitterResizeNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3661, m.Instance(), m.canSplitterResizeNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnChange(fn TVTChangeEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3662, m.Instance(), m.changePtr)
}

func (m *TLazVirtualDrawTree) SetOnChecked(fn TVTChangeEvent) {
	if m.checkedPtr != 0 {
		RemoveEventElement(m.checkedPtr)
	}
	m.checkedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3663, m.Instance(), m.checkedPtr)
}

func (m *TLazVirtualDrawTree) SetOnChecking(fn TVTCheckChangingEvent) {
	if m.checkingPtr != 0 {
		RemoveEventElement(m.checkingPtr)
	}
	m.checkingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3664, m.Instance(), m.checkingPtr)
}

func (m *TLazVirtualDrawTree) SetOnCollapsed(fn TVTChangeEvent) {
	if m.collapsedPtr != 0 {
		RemoveEventElement(m.collapsedPtr)
	}
	m.collapsedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3665, m.Instance(), m.collapsedPtr)
}

func (m *TLazVirtualDrawTree) SetOnCollapsing(fn TVTChangingEvent) {
	if m.collapsingPtr != 0 {
		RemoveEventElement(m.collapsingPtr)
	}
	m.collapsingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3666, m.Instance(), m.collapsingPtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnClick(fn TVTColumnClickEvent) {
	if m.columnClickPtr != 0 {
		RemoveEventElement(m.columnClickPtr)
	}
	m.columnClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3667, m.Instance(), m.columnClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnDblClick(fn TVTColumnDblClickEvent) {
	if m.columnDblClickPtr != 0 {
		RemoveEventElement(m.columnDblClickPtr)
	}
	m.columnDblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3668, m.Instance(), m.columnDblClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnExport(fn TVTColumnExportEvent) {
	if m.columnExportPtr != 0 {
		RemoveEventElement(m.columnExportPtr)
	}
	m.columnExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3669, m.Instance(), m.columnExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnResize(fn TVTHeaderNotifyEvent) {
	if m.columnResizePtr != 0 {
		RemoveEventElement(m.columnResizePtr)
	}
	m.columnResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3670, m.Instance(), m.columnResizePtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnWidthDblClickResize(fn TVTColumnWidthDblClickResizeEvent) {
	if m.columnWidthDblClickResizePtr != 0 {
		RemoveEventElement(m.columnWidthDblClickResizePtr)
	}
	m.columnWidthDblClickResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3671, m.Instance(), m.columnWidthDblClickResizePtr)
}

func (m *TLazVirtualDrawTree) SetOnColumnWidthTracking(fn TVTColumnWidthTrackingEvent) {
	if m.columnWidthTrackingPtr != 0 {
		RemoveEventElement(m.columnWidthTrackingPtr)
	}
	m.columnWidthTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3672, m.Instance(), m.columnWidthTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnCompareNodes(fn TVTCompareEvent) {
	if m.compareNodesPtr != 0 {
		RemoveEventElement(m.compareNodesPtr)
	}
	m.compareNodesPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3673, m.Instance(), m.compareNodesPtr)
}

func (m *TLazVirtualDrawTree) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3674, m.Instance(), m.contextPopupPtr)
}

func (m *TLazVirtualDrawTree) SetOnCreateDataObject(fn TVTCreateDataObjectEvent) {
	if m.createDataObjectPtr != 0 {
		RemoveEventElement(m.createDataObjectPtr)
	}
	m.createDataObjectPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3675, m.Instance(), m.createDataObjectPtr)
}

func (m *TLazVirtualDrawTree) SetOnCreateDragManager(fn TVTCreateDragManagerEvent) {
	if m.createDragManagerPtr != 0 {
		RemoveEventElement(m.createDragManagerPtr)
	}
	m.createDragManagerPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3676, m.Instance(), m.createDragManagerPtr)
}

func (m *TLazVirtualDrawTree) SetOnCreateEditor(fn TVTCreateEditorEvent) {
	if m.createEditorPtr != 0 {
		RemoveEventElement(m.createEditorPtr)
	}
	m.createEditorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3677, m.Instance(), m.createEditorPtr)
}

func (m *TLazVirtualDrawTree) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3678, m.Instance(), m.dblClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnDragAllowed(fn TVTDragAllowedEvent) {
	if m.dragAllowedPtr != 0 {
		RemoveEventElement(m.dragAllowedPtr)
	}
	m.dragAllowedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3679, m.Instance(), m.dragAllowedPtr)
}

func (m *TLazVirtualDrawTree) SetOnDragOver(fn TVTDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3681, m.Instance(), m.dragOverPtr)
}

func (m *TLazVirtualDrawTree) SetOnDragDrop(fn TVTDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3680, m.Instance(), m.dragDropPtr)
}

func (m *TLazVirtualDrawTree) SetOnDrawHint(fn TVTDrawHintEvent) {
	if m.drawHintPtr != 0 {
		RemoveEventElement(m.drawHintPtr)
	}
	m.drawHintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3682, m.Instance(), m.drawHintPtr)
}

func (m *TLazVirtualDrawTree) SetOnDrawNode(fn TVTDrawNodeEvent) {
	if m.drawNodePtr != 0 {
		RemoveEventElement(m.drawNodePtr)
	}
	m.drawNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3683, m.Instance(), m.drawNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnEdited(fn TVTEditChangeEvent) {
	if m.editedPtr != 0 {
		RemoveEventElement(m.editedPtr)
	}
	m.editedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3684, m.Instance(), m.editedPtr)
}

func (m *TLazVirtualDrawTree) SetOnEditing(fn TVTEditChangingEvent) {
	if m.editingPtr != 0 {
		RemoveEventElement(m.editingPtr)
	}
	m.editingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3685, m.Instance(), m.editingPtr)
}

func (m *TLazVirtualDrawTree) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3686, m.Instance(), m.endDockPtr)
}

func (m *TLazVirtualDrawTree) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3687, m.Instance(), m.endDragPtr)
}

func (m *TLazVirtualDrawTree) SetOnEndOperation(fn TVTOperationEvent) {
	if m.endOperationPtr != 0 {
		RemoveEventElement(m.endOperationPtr)
	}
	m.endOperationPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3688, m.Instance(), m.endOperationPtr)
}

func (m *TLazVirtualDrawTree) SetOnExpanded(fn TVTChangeEvent) {
	if m.expandedPtr != 0 {
		RemoveEventElement(m.expandedPtr)
	}
	m.expandedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3689, m.Instance(), m.expandedPtr)
}

func (m *TLazVirtualDrawTree) SetOnExpanding(fn TVTChangingEvent) {
	if m.expandingPtr != 0 {
		RemoveEventElement(m.expandingPtr)
	}
	m.expandingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3690, m.Instance(), m.expandingPtr)
}

func (m *TLazVirtualDrawTree) SetOnFocusChanged(fn TVTFocusChangeEvent) {
	if m.focusChangedPtr != 0 {
		RemoveEventElement(m.focusChangedPtr)
	}
	m.focusChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3691, m.Instance(), m.focusChangedPtr)
}

func (m *TLazVirtualDrawTree) SetOnFocusChanging(fn TVTFocusChangingEvent) {
	if m.focusChangingPtr != 0 {
		RemoveEventElement(m.focusChangingPtr)
	}
	m.focusChangingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3692, m.Instance(), m.focusChangingPtr)
}

func (m *TLazVirtualDrawTree) SetOnFreeNode(fn TVTFreeNodeEvent) {
	if m.freeNodePtr != 0 {
		RemoveEventElement(m.freeNodePtr)
	}
	m.freeNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3693, m.Instance(), m.freeNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnGetCellIsEmpty(fn TVTGetCellIsEmptyEvent) {
	if m.getCellIsEmptyPtr != 0 {
		RemoveEventElement(m.getCellIsEmptyPtr)
	}
	m.getCellIsEmptyPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3694, m.Instance(), m.getCellIsEmptyPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetCursor(fn TVTGetCursorEvent) {
	if m.getCursorPtr != 0 {
		RemoveEventElement(m.getCursorPtr)
	}
	m.getCursorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3695, m.Instance(), m.getCursorPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetHeaderCursor(fn TVTGetHeaderCursorEvent) {
	if m.getHeaderCursorPtr != 0 {
		RemoveEventElement(m.getHeaderCursorPtr)
	}
	m.getHeaderCursorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3696, m.Instance(), m.getHeaderCursorPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetHelpContext(fn TVTHelpContextEvent) {
	if m.getHelpContextPtr != 0 {
		RemoveEventElement(m.getHelpContextPtr)
	}
	m.getHelpContextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3697, m.Instance(), m.getHelpContextPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetHintKind(fn TVTHintKindEvent) {
	if m.getHintKindPtr != 0 {
		RemoveEventElement(m.getHintKindPtr)
	}
	m.getHintKindPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3698, m.Instance(), m.getHintKindPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetHintSize(fn TVTGetHintSizeEvent) {
	if m.getHintSizePtr != 0 {
		RemoveEventElement(m.getHintSizePtr)
	}
	m.getHintSizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3699, m.Instance(), m.getHintSizePtr)
}

func (m *TLazVirtualDrawTree) SetOnGetImageIndex(fn TVTGetImageEvent) {
	if m.getImageIndexPtr != 0 {
		RemoveEventElement(m.getImageIndexPtr)
	}
	m.getImageIndexPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3700, m.Instance(), m.getImageIndexPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetImageIndexEx(fn TVTGetImageExEvent) {
	if m.getImageIndexExPtr != 0 {
		RemoveEventElement(m.getImageIndexExPtr)
	}
	m.getImageIndexExPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3701, m.Instance(), m.getImageIndexExPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetLineStyle(fn TVTGetLineStyleEvent) {
	if m.getLineStylePtr != 0 {
		RemoveEventElement(m.getLineStylePtr)
	}
	m.getLineStylePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3702, m.Instance(), m.getLineStylePtr)
}

func (m *TLazVirtualDrawTree) SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent) {
	if m.getNodeDataSizePtr != 0 {
		RemoveEventElement(m.getNodeDataSizePtr)
	}
	m.getNodeDataSizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3703, m.Instance(), m.getNodeDataSizePtr)
}

func (m *TLazVirtualDrawTree) SetOnGetNodeWidth(fn TVTGetNodeWidthEvent) {
	if m.getNodeWidthPtr != 0 {
		RemoveEventElement(m.getNodeWidthPtr)
	}
	m.getNodeWidthPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3704, m.Instance(), m.getNodeWidthPtr)
}

func (m *TLazVirtualDrawTree) SetOnGetPopupMenu(fn TVTPopupEvent) {
	if m.getPopupMenuPtr != 0 {
		RemoveEventElement(m.getPopupMenuPtr)
	}
	m.getPopupMenuPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3705, m.Instance(), m.getPopupMenuPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderClick(fn TVTHeaderClickEvent) {
	if m.headerClickPtr != 0 {
		RemoveEventElement(m.headerClickPtr)
	}
	m.headerClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3706, m.Instance(), m.headerClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDblClick(fn TVTHeaderClickEvent) {
	if m.headerDblClickPtr != 0 {
		RemoveEventElement(m.headerDblClickPtr)
	}
	m.headerDblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3707, m.Instance(), m.headerDblClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDragged(fn TVTHeaderDraggedEvent) {
	if m.headerDraggedPtr != 0 {
		RemoveEventElement(m.headerDraggedPtr)
	}
	m.headerDraggedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3708, m.Instance(), m.headerDraggedPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent) {
	if m.headerDraggedOutPtr != 0 {
		RemoveEventElement(m.headerDraggedOutPtr)
	}
	m.headerDraggedOutPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3709, m.Instance(), m.headerDraggedOutPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDragging(fn TVTHeaderDraggingEvent) {
	if m.headerDraggingPtr != 0 {
		RemoveEventElement(m.headerDraggingPtr)
	}
	m.headerDraggingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3710, m.Instance(), m.headerDraggingPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDraw(fn TVTHeaderPaintEvent) {
	if m.headerDrawPtr != 0 {
		RemoveEventElement(m.headerDrawPtr)
	}
	m.headerDrawPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3711, m.Instance(), m.headerDrawPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent) {
	if m.headerDrawQueryElementsPtr != 0 {
		RemoveEventElement(m.headerDrawQueryElementsPtr)
	}
	m.headerDrawQueryElementsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3712, m.Instance(), m.headerDrawQueryElementsPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent) {
	if m.headerHeightTrackingPtr != 0 {
		RemoveEventElement(m.headerHeightTrackingPtr)
	}
	m.headerHeightTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3714, m.Instance(), m.headerHeightTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) {
	if m.headerHeightDblClickResizePtr != 0 {
		RemoveEventElement(m.headerHeightDblClickResizePtr)
	}
	m.headerHeightDblClickResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3713, m.Instance(), m.headerHeightDblClickResizePtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseDown(fn TVTHeaderMouseEvent) {
	if m.headerMouseDownPtr != 0 {
		RemoveEventElement(m.headerMouseDownPtr)
	}
	m.headerMouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3715, m.Instance(), m.headerMouseDownPtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseMove(fn TVTHeaderMouseMoveEvent) {
	if m.headerMouseMovePtr != 0 {
		RemoveEventElement(m.headerMouseMovePtr)
	}
	m.headerMouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3716, m.Instance(), m.headerMouseMovePtr)
}

func (m *TLazVirtualDrawTree) SetOnHeaderMouseUp(fn TVTHeaderMouseEvent) {
	if m.headerMouseUpPtr != 0 {
		RemoveEventElement(m.headerMouseUpPtr)
	}
	m.headerMouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3717, m.Instance(), m.headerMouseUpPtr)
}

func (m *TLazVirtualDrawTree) SetOnHotChange(fn TVTHotNodeChangeEvent) {
	if m.hotChangePtr != 0 {
		RemoveEventElement(m.hotChangePtr)
	}
	m.hotChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3718, m.Instance(), m.hotChangePtr)
}

func (m *TLazVirtualDrawTree) SetOnIncrementalSearch(fn TVTIncrementalSearchEvent) {
	if m.incrementalSearchPtr != 0 {
		RemoveEventElement(m.incrementalSearchPtr)
	}
	m.incrementalSearchPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3719, m.Instance(), m.incrementalSearchPtr)
}

func (m *TLazVirtualDrawTree) SetOnInitChildren(fn TVTInitChildrenEvent) {
	if m.initChildrenPtr != 0 {
		RemoveEventElement(m.initChildrenPtr)
	}
	m.initChildrenPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3720, m.Instance(), m.initChildrenPtr)
}

func (m *TLazVirtualDrawTree) SetOnInitNode(fn TVTInitNodeEvent) {
	if m.initNodePtr != 0 {
		RemoveEventElement(m.initNodePtr)
	}
	m.initNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3721, m.Instance(), m.initNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnKeyAction(fn TVTKeyActionEvent) {
	if m.keyActionPtr != 0 {
		RemoveEventElement(m.keyActionPtr)
	}
	m.keyActionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3722, m.Instance(), m.keyActionPtr)
}

func (m *TLazVirtualDrawTree) SetOnLoadNode(fn TVTSaveNodeEvent) {
	if m.loadNodePtr != 0 {
		RemoveEventElement(m.loadNodePtr)
	}
	m.loadNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3723, m.Instance(), m.loadNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnLoadTree(fn TVTSaveTreeEvent) {
	if m.loadTreePtr != 0 {
		RemoveEventElement(m.loadTreePtr)
	}
	m.loadTreePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3724, m.Instance(), m.loadTreePtr)
}

func (m *TLazVirtualDrawTree) SetOnMeasureItem(fn TVTMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3725, m.Instance(), m.measureItemPtr)
}

func (m *TLazVirtualDrawTree) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3726, m.Instance(), m.mouseDownPtr)
}

func (m *TLazVirtualDrawTree) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3727, m.Instance(), m.mouseMovePtr)
}

func (m *TLazVirtualDrawTree) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3728, m.Instance(), m.mouseUpPtr)
}

func (m *TLazVirtualDrawTree) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3729, m.Instance(), m.mouseWheelPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeClick(fn TVTNodeClickEvent) {
	if m.nodeClickPtr != 0 {
		RemoveEventElement(m.nodeClickPtr)
	}
	m.nodeClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3730, m.Instance(), m.nodeClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeCopied(fn TVTNodeCopiedEvent) {
	if m.nodeCopiedPtr != 0 {
		RemoveEventElement(m.nodeCopiedPtr)
	}
	m.nodeCopiedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3731, m.Instance(), m.nodeCopiedPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeCopying(fn TVTNodeCopyingEvent) {
	if m.nodeCopyingPtr != 0 {
		RemoveEventElement(m.nodeCopyingPtr)
	}
	m.nodeCopyingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3732, m.Instance(), m.nodeCopyingPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeDblClick(fn TVTNodeClickEvent) {
	if m.nodeDblClickPtr != 0 {
		RemoveEventElement(m.nodeDblClickPtr)
	}
	m.nodeDblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3733, m.Instance(), m.nodeDblClickPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeExport(fn TVTNodeExportEvent) {
	if m.nodeExportPtr != 0 {
		RemoveEventElement(m.nodeExportPtr)
	}
	m.nodeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3734, m.Instance(), m.nodeExportPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent) {
	if m.nodeHeightTrackingPtr != 0 {
		RemoveEventElement(m.nodeHeightTrackingPtr)
	}
	m.nodeHeightTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3736, m.Instance(), m.nodeHeightTrackingPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent) {
	if m.nodeHeightDblClickResizePtr != 0 {
		RemoveEventElement(m.nodeHeightDblClickResizePtr)
	}
	m.nodeHeightDblClickResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3735, m.Instance(), m.nodeHeightDblClickResizePtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeMoved(fn TVTNodeMovedEvent) {
	if m.nodeMovedPtr != 0 {
		RemoveEventElement(m.nodeMovedPtr)
	}
	m.nodeMovedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3737, m.Instance(), m.nodeMovedPtr)
}

func (m *TLazVirtualDrawTree) SetOnNodeMoving(fn TVTNodeMovingEvent) {
	if m.nodeMovingPtr != 0 {
		RemoveEventElement(m.nodeMovingPtr)
	}
	m.nodeMovingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3738, m.Instance(), m.nodeMovingPtr)
}

func (m *TLazVirtualDrawTree) SetOnPaintBackground(fn TVTBackgroundPaintEvent) {
	if m.paintBackgroundPtr != 0 {
		RemoveEventElement(m.paintBackgroundPtr)
	}
	m.paintBackgroundPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3739, m.Instance(), m.paintBackgroundPtr)
}

func (m *TLazVirtualDrawTree) SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent) {
	if m.removeFromSelectionPtr != 0 {
		RemoveEventElement(m.removeFromSelectionPtr)
	}
	m.removeFromSelectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3740, m.Instance(), m.removeFromSelectionPtr)
}

func (m *TLazVirtualDrawTree) SetOnResetNode(fn TVTChangeEvent) {
	if m.resetNodePtr != 0 {
		RemoveEventElement(m.resetNodePtr)
	}
	m.resetNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3741, m.Instance(), m.resetNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnSaveNode(fn TVTSaveNodeEvent) {
	if m.saveNodePtr != 0 {
		RemoveEventElement(m.saveNodePtr)
	}
	m.saveNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3742, m.Instance(), m.saveNodePtr)
}

func (m *TLazVirtualDrawTree) SetOnSaveTree(fn TVTSaveTreeEvent) {
	if m.saveTreePtr != 0 {
		RemoveEventElement(m.saveTreePtr)
	}
	m.saveTreePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3743, m.Instance(), m.saveTreePtr)
}

func (m *TLazVirtualDrawTree) SetOnScroll(fn TVTScrollEvent) {
	if m.scrollPtr != 0 {
		RemoveEventElement(m.scrollPtr)
	}
	m.scrollPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3744, m.Instance(), m.scrollPtr)
}

func (m *TLazVirtualDrawTree) SetOnShowScrollBar(fn TVTScrollBarShowEvent) {
	if m.showScrollBarPtr != 0 {
		RemoveEventElement(m.showScrollBarPtr)
	}
	m.showScrollBarPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3745, m.Instance(), m.showScrollBarPtr)
}

func (m *TLazVirtualDrawTree) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3746, m.Instance(), m.startDockPtr)
}

func (m *TLazVirtualDrawTree) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3747, m.Instance(), m.startDragPtr)
}

func (m *TLazVirtualDrawTree) SetOnStartOperation(fn TVTOperationEvent) {
	if m.startOperationPtr != 0 {
		RemoveEventElement(m.startOperationPtr)
	}
	m.startOperationPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3748, m.Instance(), m.startOperationPtr)
}

func (m *TLazVirtualDrawTree) SetOnStateChange(fn TVTStateChangeEvent) {
	if m.stateChangePtr != 0 {
		RemoveEventElement(m.stateChangePtr)
	}
	m.stateChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3749, m.Instance(), m.stateChangePtr)
}

func (m *TLazVirtualDrawTree) SetOnStructureChange(fn TVTStructureChangeEvent) {
	if m.structureChangePtr != 0 {
		RemoveEventElement(m.structureChangePtr)
	}
	m.structureChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3750, m.Instance(), m.structureChangePtr)
}

func (m *TLazVirtualDrawTree) SetOnUpdating(fn TVTUpdatingEvent) {
	if m.updatingPtr != 0 {
		RemoveEventElement(m.updatingPtr)
	}
	m.updatingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3751, m.Instance(), m.updatingPtr)
}
