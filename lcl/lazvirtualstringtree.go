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

// ILazVirtualStringTree Parent: ICustomVirtualStringTree
type ILazVirtualStringTree interface {
	ICustomVirtualStringTree
	RangeX() uint32                                                        // property
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
	DefaultText() string                                                   // property
	SetDefaultText(AValue string)                                          // property
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
	TreeOptions() IStringTreeOptions                                       // property
	SetTreeOptions(AValue IStringTreeOptions)                              // property
	WantTabs() bool                                                        // property
	SetWantTabs(AValue bool)                                               // property
	ImagesWidth() int32                                                    // property
	SetImagesWidth(AValue int32)                                           // property
	StateImagesWidth() int32                                               // property
	SetStateImagesWidth(AValue int32)                                      // property
	CustomCheckImagesWidth() int32                                         // property
	SetCustomCheckImagesWidth(AValue int32)                                // property
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
	SetOnDrawText(fn TVTDrawTextEvent)                                     // property event
	SetOnEditCancelled(fn TVTEditCancelEvent)                              // property event
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
	SetOnGetText(fn TVSTGetTextEvent)                                      // property event
	SetOnPaintText(fn TVTPaintText)                                        // property event
	SetOnGetHelpContext(fn TVTHelpContextEvent)                            // property event
	SetOnGetHintKind(fn TVTHintKindEvent)                                  // property event
	SetOnGetHintSize(fn TVTGetHintSizeEvent)                               // property event
	SetOnGetImageIndex(fn TVTGetImageEvent)                                // property event
	SetOnGetImageIndexEx(fn TVTGetImageExEvent)                            // property event
	SetOnGetImageText(fn TVTGetImageTextEvent)                             // property event
	SetOnGetHint(fn TVSTGetHintEvent)                                      // property event
	SetOnGetLineStyle(fn TVTGetLineStyleEvent)                             // property event
	SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent)                       // property event
	SetOnGetPopupMenu(fn TVTPopupEvent)                                    // property event
	SetOnHeaderClick(fn TVTHeaderClickEvent)                               // property event
	SetOnHeaderDblClick(fn TVTHeaderClickEvent)                            // property event
	SetOnHeaderDragged(fn TVTHeaderDraggedEvent)                           // property event
	SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent)                     // property event
	SetOnHeaderDragging(fn TVTHeaderDraggingEvent)                         // property event
	SetOnHeaderDraw(fn TVTHeaderPaintEvent)                                // property event
	SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent)      // property event
	SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) // property event
	SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent)             // property event
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
	SetOnMeasureTextWidth(fn TVTMeasureTextEvent)                          // property event
	SetOnMeasureTextHeight(fn TVTMeasureTextEvent)                         // property event
	SetOnMouseDown(fn TMouseEvent)                                         // property event
	SetOnMouseMove(fn TMouseMoveEvent)                                     // property event
	SetOnMouseUp(fn TMouseEvent)                                           // property event
	SetOnMouseWheel(fn TMouseWheelEvent)                                   // property event
	SetOnMouseEnter(fn TNotifyEvent)                                       // property event
	SetOnMouseLeave(fn TNotifyEvent)                                       // property event
	SetOnNewText(fn TVSTNewTextEvent)                                      // property event
	SetOnNodeClick(fn TVTNodeClickEvent)                                   // property event
	SetOnNodeCopied(fn TVTNodeCopiedEvent)                                 // property event
	SetOnNodeCopying(fn TVTNodeCopyingEvent)                               // property event
	SetOnNodeDblClick(fn TVTNodeClickEvent)                                // property event
	SetOnNodeExport(fn TVTNodeExportEvent)                                 // property event
	SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent)     // property event
	SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent)                 // property event
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

// TLazVirtualStringTree Parent: TCustomVirtualStringTree
type TLazVirtualStringTree struct {
	TCustomVirtualStringTree
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
	drawTextPtr                   uintptr
	editCancelledPtr              uintptr
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
	getTextPtr                    uintptr
	paintTextPtr                  uintptr
	getHelpContextPtr             uintptr
	getHintKindPtr                uintptr
	getHintSizePtr                uintptr
	getImageIndexPtr              uintptr
	getImageIndexExPtr            uintptr
	getImageTextPtr               uintptr
	getHintPtr                    uintptr
	getLineStylePtr               uintptr
	getNodeDataSizePtr            uintptr
	getPopupMenuPtr               uintptr
	headerClickPtr                uintptr
	headerDblClickPtr             uintptr
	headerDraggedPtr              uintptr
	headerDraggedOutPtr           uintptr
	headerDraggingPtr             uintptr
	headerDrawPtr                 uintptr
	headerDrawQueryElementsPtr    uintptr
	headerHeightDblClickResizePtr uintptr
	headerHeightTrackingPtr       uintptr
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
	measureTextWidthPtr           uintptr
	measureTextHeightPtr          uintptr
	mouseDownPtr                  uintptr
	mouseMovePtr                  uintptr
	mouseUpPtr                    uintptr
	mouseWheelPtr                 uintptr
	mouseEnterPtr                 uintptr
	mouseLeavePtr                 uintptr
	newTextPtr                    uintptr
	nodeClickPtr                  uintptr
	nodeCopiedPtr                 uintptr
	nodeCopyingPtr                uintptr
	nodeDblClickPtr               uintptr
	nodeExportPtr                 uintptr
	nodeHeightDblClickResizePtr   uintptr
	nodeHeightTrackingPtr         uintptr
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

func NewLazVirtualStringTree(AOwner IComponent) ILazVirtualStringTree {
	r1 := LCL().SysCallN(3772, GetObjectUintptr(AOwner))
	return AsLazVirtualStringTree(r1)
}

func (m *TLazVirtualStringTree) RangeX() uint32 {
	r1 := LCL().SysCallN(3808, m.Instance())
	return uint32(r1)
}

func (m *TLazVirtualStringTree) LastDragEffect() uint32 {
	r1 := LCL().SysCallN(3798, m.Instance())
	return uint32(r1)
}

func (m *TLazVirtualStringTree) Alignment() TAlignment {
	r1 := LCL().SysCallN(3756, 0, m.Instance(), 0)
	return TAlignment(r1)
}

func (m *TLazVirtualStringTree) SetAlignment(AValue TAlignment) {
	LCL().SysCallN(3756, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) AnimationDuration() uint32 {
	r1 := LCL().SysCallN(3757, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetAnimationDuration(AValue uint32) {
	LCL().SysCallN(3757, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) AutoExpandDelay() uint32 {
	r1 := LCL().SysCallN(3758, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetAutoExpandDelay(AValue uint32) {
	LCL().SysCallN(3758, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) AutoScrollDelay() uint32 {
	r1 := LCL().SysCallN(3759, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetAutoScrollDelay(AValue uint32) {
	LCL().SysCallN(3759, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) AutoScrollInterval() TAutoScrollInterval {
	r1 := LCL().SysCallN(3760, 0, m.Instance(), 0)
	return TAutoScrollInterval(r1)
}

func (m *TLazVirtualStringTree) SetAutoScrollInterval(AValue TAutoScrollInterval) {
	LCL().SysCallN(3760, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Background() IPicture {
	r1 := LCL().SysCallN(3761, 0, m.Instance(), 0)
	return AsPicture(r1)
}

func (m *TLazVirtualStringTree) SetBackground(AValue IPicture) {
	LCL().SysCallN(3761, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) BackgroundOffsetX() int32 {
	r1 := LCL().SysCallN(3762, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetBackgroundOffsetX(AValue int32) {
	LCL().SysCallN(3762, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) BackgroundOffsetY() int32 {
	r1 := LCL().SysCallN(3763, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetBackgroundOffsetY(AValue int32) {
	LCL().SysCallN(3763, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) BottomSpace() uint32 {
	r1 := LCL().SysCallN(3764, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetBottomSpace(AValue uint32) {
	LCL().SysCallN(3764, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ButtonFillMode() TVTButtonFillMode {
	r1 := LCL().SysCallN(3765, 0, m.Instance(), 0)
	return TVTButtonFillMode(r1)
}

func (m *TLazVirtualStringTree) SetButtonFillMode(AValue TVTButtonFillMode) {
	LCL().SysCallN(3765, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ButtonStyle() TVTButtonStyle {
	r1 := LCL().SysCallN(3766, 0, m.Instance(), 0)
	return TVTButtonStyle(r1)
}

func (m *TLazVirtualStringTree) SetButtonStyle(AValue TVTButtonStyle) {
	LCL().SysCallN(3766, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ChangeDelay() uint32 {
	r1 := LCL().SysCallN(3767, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetChangeDelay(AValue uint32) {
	LCL().SysCallN(3767, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) CheckImageKind() TCheckImageKind {
	r1 := LCL().SysCallN(3768, 0, m.Instance(), 0)
	return TCheckImageKind(r1)
}

func (m *TLazVirtualStringTree) SetCheckImageKind(AValue TCheckImageKind) {
	LCL().SysCallN(3768, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ClipboardFormats() IClipboardFormats {
	r1 := LCL().SysCallN(3770, 0, m.Instance(), 0)
	return AsClipboardFormats(r1)
}

func (m *TLazVirtualStringTree) SetClipboardFormats(AValue IClipboardFormats) {
	LCL().SysCallN(3770, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) Colors() IVTColors {
	r1 := LCL().SysCallN(3771, 0, m.Instance(), 0)
	return AsVTColors(r1)
}

func (m *TLazVirtualStringTree) SetColors(AValue IVTColors) {
	LCL().SysCallN(3771, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) CustomCheckImages() ICustomImageList {
	r1 := LCL().SysCallN(3773, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualStringTree) SetCustomCheckImages(AValue ICustomImageList) {
	LCL().SysCallN(3773, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) DefaultNodeHeight() uint32 {
	r1 := LCL().SysCallN(3775, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetDefaultNodeHeight(AValue uint32) {
	LCL().SysCallN(3775, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DefaultPasteMode() TVTNodeAttachMode {
	r1 := LCL().SysCallN(3776, 0, m.Instance(), 0)
	return TVTNodeAttachMode(r1)
}

func (m *TLazVirtualStringTree) SetDefaultPasteMode(AValue TVTNodeAttachMode) {
	LCL().SysCallN(3776, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DefaultText() string {
	r1 := LCL().SysCallN(3777, 0, m.Instance(), 0)
	return GoStr(r1)
}

func (m *TLazVirtualStringTree) SetDefaultText(AValue string) {
	LCL().SysCallN(3777, 1, m.Instance(), PascalStr(AValue))
}

func (m *TLazVirtualStringTree) DragCursor() TCursor {
	r1 := LCL().SysCallN(3778, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLazVirtualStringTree) SetDragCursor(AValue TCursor) {
	LCL().SysCallN(3778, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragHeight() int32 {
	r1 := LCL().SysCallN(3779, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetDragHeight(AValue int32) {
	LCL().SysCallN(3779, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragKind() TDragKind {
	r1 := LCL().SysCallN(3781, 0, m.Instance(), 0)
	return TDragKind(r1)
}

func (m *TLazVirtualStringTree) SetDragKind(AValue TDragKind) {
	LCL().SysCallN(3781, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragImageKind() TVTDragImageKind {
	r1 := LCL().SysCallN(3780, 0, m.Instance(), 0)
	return TVTDragImageKind(r1)
}

func (m *TLazVirtualStringTree) SetDragImageKind(AValue TVTDragImageKind) {
	LCL().SysCallN(3780, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragMode() TDragMode {
	r1 := LCL().SysCallN(3782, 0, m.Instance(), 0)
	return TDragMode(r1)
}

func (m *TLazVirtualStringTree) SetDragMode(AValue TDragMode) {
	LCL().SysCallN(3782, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragOperations() TDragOperations {
	r1 := LCL().SysCallN(3783, 0, m.Instance(), 0)
	return TDragOperations(r1)
}

func (m *TLazVirtualStringTree) SetDragOperations(AValue TDragOperations) {
	LCL().SysCallN(3783, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragType() TVTDragType {
	r1 := LCL().SysCallN(3784, 0, m.Instance(), 0)
	return TVTDragType(r1)
}

func (m *TLazVirtualStringTree) SetDragType(AValue TVTDragType) {
	LCL().SysCallN(3784, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DragWidth() int32 {
	r1 := LCL().SysCallN(3785, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetDragWidth(AValue int32) {
	LCL().SysCallN(3785, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) DrawSelectionMode() TVTDrawSelectionMode {
	r1 := LCL().SysCallN(3786, 0, m.Instance(), 0)
	return TVTDrawSelectionMode(r1)
}

func (m *TLazVirtualStringTree) SetDrawSelectionMode(AValue TVTDrawSelectionMode) {
	LCL().SysCallN(3786, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) EditDelay() uint32 {
	r1 := LCL().SysCallN(3787, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetEditDelay(AValue uint32) {
	LCL().SysCallN(3787, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Header() IVTHeader {
	r1 := LCL().SysCallN(3788, 0, m.Instance(), 0)
	return AsVTHeader(r1)
}

func (m *TLazVirtualStringTree) SetHeader(AValue IVTHeader) {
	LCL().SysCallN(3788, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) HintMode() TVTHintMode {
	r1 := LCL().SysCallN(3789, 0, m.Instance(), 0)
	return TVTHintMode(r1)
}

func (m *TLazVirtualStringTree) SetHintMode(AValue TVTHintMode) {
	LCL().SysCallN(3789, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) HotCursor() TCursor {
	r1 := LCL().SysCallN(3790, 0, m.Instance(), 0)
	return TCursor(r1)
}

func (m *TLazVirtualStringTree) SetHotCursor(AValue TCursor) {
	LCL().SysCallN(3790, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Images() ICustomImageList {
	r1 := LCL().SysCallN(3791, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualStringTree) SetImages(AValue ICustomImageList) {
	LCL().SysCallN(3791, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) IncrementalSearch() TVTIncrementalSearch {
	r1 := LCL().SysCallN(3793, 0, m.Instance(), 0)
	return TVTIncrementalSearch(r1)
}

func (m *TLazVirtualStringTree) SetIncrementalSearch(AValue TVTIncrementalSearch) {
	LCL().SysCallN(3793, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) IncrementalSearchDirection() TVTSearchDirection {
	r1 := LCL().SysCallN(3794, 0, m.Instance(), 0)
	return TVTSearchDirection(r1)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchDirection(AValue TVTSearchDirection) {
	LCL().SysCallN(3794, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) IncrementalSearchStart() TVTSearchStart {
	r1 := LCL().SysCallN(3795, 0, m.Instance(), 0)
	return TVTSearchStart(r1)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchStart(AValue TVTSearchStart) {
	LCL().SysCallN(3795, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) IncrementalSearchTimeout() uint32 {
	r1 := LCL().SysCallN(3796, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetIncrementalSearchTimeout(AValue uint32) {
	LCL().SysCallN(3796, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Indent() uint32 {
	r1 := LCL().SysCallN(3797, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetIndent(AValue uint32) {
	LCL().SysCallN(3797, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) LineMode() TVTLineMode {
	r1 := LCL().SysCallN(3799, 0, m.Instance(), 0)
	return TVTLineMode(r1)
}

func (m *TLazVirtualStringTree) SetLineMode(AValue TVTLineMode) {
	LCL().SysCallN(3799, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) LineStyle() TVTLineStyle {
	r1 := LCL().SysCallN(3800, 0, m.Instance(), 0)
	return TVTLineStyle(r1)
}

func (m *TLazVirtualStringTree) SetLineStyle(AValue TVTLineStyle) {
	LCL().SysCallN(3800, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) Margin() int32 {
	r1 := LCL().SysCallN(3801, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetMargin(AValue int32) {
	LCL().SysCallN(3801, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) NodeAlignment() TVTNodeAlignment {
	r1 := LCL().SysCallN(3802, 0, m.Instance(), 0)
	return TVTNodeAlignment(r1)
}

func (m *TLazVirtualStringTree) SetNodeAlignment(AValue TVTNodeAlignment) {
	LCL().SysCallN(3802, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) NodeDataSize() int32 {
	r1 := LCL().SysCallN(3803, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetNodeDataSize(AValue int32) {
	LCL().SysCallN(3803, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) OperationCanceled() bool {
	r1 := LCL().SysCallN(3804, m.Instance())
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) ParentColor() bool {
	r1 := LCL().SysCallN(3805, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) SetParentColor(AValue bool) {
	LCL().SysCallN(3805, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualStringTree) ParentFont() bool {
	r1 := LCL().SysCallN(3806, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) SetParentFont(AValue bool) {
	LCL().SysCallN(3806, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualStringTree) ParentShowHint() bool {
	r1 := LCL().SysCallN(3807, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) SetParentShowHint(AValue bool) {
	LCL().SysCallN(3807, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualStringTree) RootNodeCount() uint32 {
	r1 := LCL().SysCallN(3809, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetRootNodeCount(AValue uint32) {
	LCL().SysCallN(3809, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) ScrollBarOptions() IScrollBarOptions {
	r1 := LCL().SysCallN(3810, 0, m.Instance(), 0)
	return AsScrollBarOptions(r1)
}

func (m *TLazVirtualStringTree) SetScrollBarOptions(AValue IScrollBarOptions) {
	LCL().SysCallN(3810, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) SelectionBlendFactor() Byte {
	r1 := LCL().SysCallN(3811, 0, m.Instance(), 0)
	return Byte(r1)
}

func (m *TLazVirtualStringTree) SetSelectionBlendFactor(AValue Byte) {
	LCL().SysCallN(3811, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) SelectionCurveRadius() uint32 {
	r1 := LCL().SysCallN(3812, 0, m.Instance(), 0)
	return uint32(r1)
}

func (m *TLazVirtualStringTree) SetSelectionCurveRadius(AValue uint32) {
	LCL().SysCallN(3812, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) StateImages() ICustomImageList {
	r1 := LCL().SysCallN(3944, 0, m.Instance(), 0)
	return AsCustomImageList(r1)
}

func (m *TLazVirtualStringTree) SetStateImages(AValue ICustomImageList) {
	LCL().SysCallN(3944, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) TextMargin() int32 {
	r1 := LCL().SysCallN(3946, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetTextMargin(AValue int32) {
	LCL().SysCallN(3946, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) TreeOptions() IStringTreeOptions {
	r1 := LCL().SysCallN(3947, 0, m.Instance(), 0)
	return AsStringTreeOptions(r1)
}

func (m *TLazVirtualStringTree) SetTreeOptions(AValue IStringTreeOptions) {
	LCL().SysCallN(3947, 1, m.Instance(), GetObjectUintptr(AValue))
}

func (m *TLazVirtualStringTree) WantTabs() bool {
	r1 := LCL().SysCallN(3948, 0, m.Instance(), 0)
	return GoBool(r1)
}

func (m *TLazVirtualStringTree) SetWantTabs(AValue bool) {
	LCL().SysCallN(3948, 1, m.Instance(), PascalBool(AValue))
}

func (m *TLazVirtualStringTree) ImagesWidth() int32 {
	r1 := LCL().SysCallN(3792, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetImagesWidth(AValue int32) {
	LCL().SysCallN(3792, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) StateImagesWidth() int32 {
	r1 := LCL().SysCallN(3945, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetStateImagesWidth(AValue int32) {
	LCL().SysCallN(3945, 1, m.Instance(), uintptr(AValue))
}

func (m *TLazVirtualStringTree) CustomCheckImagesWidth() int32 {
	r1 := LCL().SysCallN(3774, 0, m.Instance(), 0)
	return int32(r1)
}

func (m *TLazVirtualStringTree) SetCustomCheckImagesWidth(AValue int32) {
	LCL().SysCallN(3774, 1, m.Instance(), uintptr(AValue))
}

func LazVirtualStringTreeClass() TClass {
	ret := LCL().SysCallN(3769)
	return TClass(ret)
}

func (m *TLazVirtualStringTree) SetOnAddToSelection(fn TVTAddToSelectionEvent) {
	if m.addToSelectionPtr != 0 {
		RemoveEventElement(m.addToSelectionPtr)
	}
	m.addToSelectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3813, m.Instance(), m.addToSelectionPtr)
}

func (m *TLazVirtualStringTree) SetOnAdvancedHeaderDraw(fn TVTAdvancedHeaderPaintEvent) {
	if m.advancedHeaderDrawPtr != 0 {
		RemoveEventElement(m.advancedHeaderDrawPtr)
	}
	m.advancedHeaderDrawPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3814, m.Instance(), m.advancedHeaderDrawPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterAutoFitColumn(fn TVTAfterAutoFitColumnEvent) {
	if m.afterAutoFitColumnPtr != 0 {
		RemoveEventElement(m.afterAutoFitColumnPtr)
	}
	m.afterAutoFitColumnPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3815, m.Instance(), m.afterAutoFitColumnPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterAutoFitColumns(fn TVTAfterAutoFitColumnsEvent) {
	if m.afterAutoFitColumnsPtr != 0 {
		RemoveEventElement(m.afterAutoFitColumnsPtr)
	}
	m.afterAutoFitColumnsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3816, m.Instance(), m.afterAutoFitColumnsPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterCellPaint(fn TVTAfterCellPaintEvent) {
	if m.afterCellPaintPtr != 0 {
		RemoveEventElement(m.afterCellPaintPtr)
	}
	m.afterCellPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3817, m.Instance(), m.afterCellPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterColumnExport(fn TVTColumnExportEvent) {
	if m.afterColumnExportPtr != 0 {
		RemoveEventElement(m.afterColumnExportPtr)
	}
	m.afterColumnExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3818, m.Instance(), m.afterColumnExportPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterColumnWidthTracking(fn TVTAfterColumnWidthTrackingEvent) {
	if m.afterColumnWidthTrackingPtr != 0 {
		RemoveEventElement(m.afterColumnWidthTrackingPtr)
	}
	m.afterColumnWidthTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3819, m.Instance(), m.afterColumnWidthTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterGetMaxColumnWidth(fn TVTAfterGetMaxColumnWidthEvent) {
	if m.afterGetMaxColumnWidthPtr != 0 {
		RemoveEventElement(m.afterGetMaxColumnWidthPtr)
	}
	m.afterGetMaxColumnWidthPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3820, m.Instance(), m.afterGetMaxColumnWidthPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterHeaderExport(fn TVTTreeExportEvent) {
	if m.afterHeaderExportPtr != 0 {
		RemoveEventElement(m.afterHeaderExportPtr)
	}
	m.afterHeaderExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3821, m.Instance(), m.afterHeaderExportPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterHeaderHeightTracking(fn TVTAfterHeaderHeightTrackingEvent) {
	if m.afterHeaderHeightTrackingPtr != 0 {
		RemoveEventElement(m.afterHeaderHeightTrackingPtr)
	}
	m.afterHeaderHeightTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3822, m.Instance(), m.afterHeaderHeightTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterItemErase(fn TVTAfterItemEraseEvent) {
	if m.afterItemErasePtr != 0 {
		RemoveEventElement(m.afterItemErasePtr)
	}
	m.afterItemErasePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3823, m.Instance(), m.afterItemErasePtr)
}

func (m *TLazVirtualStringTree) SetOnAfterItemPaint(fn TVTAfterItemPaintEvent) {
	if m.afterItemPaintPtr != 0 {
		RemoveEventElement(m.afterItemPaintPtr)
	}
	m.afterItemPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3824, m.Instance(), m.afterItemPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterNodeExport(fn TVTNodeExportEvent) {
	if m.afterNodeExportPtr != 0 {
		RemoveEventElement(m.afterNodeExportPtr)
	}
	m.afterNodeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3825, m.Instance(), m.afterNodeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterPaint(fn TVTPaintEvent) {
	if m.afterPaintPtr != 0 {
		RemoveEventElement(m.afterPaintPtr)
	}
	m.afterPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3826, m.Instance(), m.afterPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnAfterTreeExport(fn TVTTreeExportEvent) {
	if m.afterTreeExportPtr != 0 {
		RemoveEventElement(m.afterTreeExportPtr)
	}
	m.afterTreeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3827, m.Instance(), m.afterTreeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeAutoFitColumn(fn TVTBeforeAutoFitColumnEvent) {
	if m.beforeAutoFitColumnPtr != 0 {
		RemoveEventElement(m.beforeAutoFitColumnPtr)
	}
	m.beforeAutoFitColumnPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3828, m.Instance(), m.beforeAutoFitColumnPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeAutoFitColumns(fn TVTBeforeAutoFitColumnsEvent) {
	if m.beforeAutoFitColumnsPtr != 0 {
		RemoveEventElement(m.beforeAutoFitColumnsPtr)
	}
	m.beforeAutoFitColumnsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3829, m.Instance(), m.beforeAutoFitColumnsPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeCellPaint(fn TVTBeforeCellPaintEvent) {
	if m.beforeCellPaintPtr != 0 {
		RemoveEventElement(m.beforeCellPaintPtr)
	}
	m.beforeCellPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3830, m.Instance(), m.beforeCellPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeColumnExport(fn TVTColumnExportEvent) {
	if m.beforeColumnExportPtr != 0 {
		RemoveEventElement(m.beforeColumnExportPtr)
	}
	m.beforeColumnExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3831, m.Instance(), m.beforeColumnExportPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeColumnWidthTracking(fn TVTBeforeColumnWidthTrackingEvent) {
	if m.beforeColumnWidthTrackingPtr != 0 {
		RemoveEventElement(m.beforeColumnWidthTrackingPtr)
	}
	m.beforeColumnWidthTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3832, m.Instance(), m.beforeColumnWidthTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeDrawTreeLine(fn TVTBeforeDrawLineImageEvent) {
	if m.beforeDrawTreeLinePtr != 0 {
		RemoveEventElement(m.beforeDrawTreeLinePtr)
	}
	m.beforeDrawTreeLinePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3833, m.Instance(), m.beforeDrawTreeLinePtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeGetMaxColumnWidth(fn TVTBeforeGetMaxColumnWidthEvent) {
	if m.beforeGetMaxColumnWidthPtr != 0 {
		RemoveEventElement(m.beforeGetMaxColumnWidthPtr)
	}
	m.beforeGetMaxColumnWidthPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3834, m.Instance(), m.beforeGetMaxColumnWidthPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeHeaderExport(fn TVTTreeExportEvent) {
	if m.beforeHeaderExportPtr != 0 {
		RemoveEventElement(m.beforeHeaderExportPtr)
	}
	m.beforeHeaderExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3835, m.Instance(), m.beforeHeaderExportPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeHeaderHeightTracking(fn TVTBeforeHeaderHeightTrackingEvent) {
	if m.beforeHeaderHeightTrackingPtr != 0 {
		RemoveEventElement(m.beforeHeaderHeightTrackingPtr)
	}
	m.beforeHeaderHeightTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3836, m.Instance(), m.beforeHeaderHeightTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeItemErase(fn TVTBeforeItemEraseEvent) {
	if m.beforeItemErasePtr != 0 {
		RemoveEventElement(m.beforeItemErasePtr)
	}
	m.beforeItemErasePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3837, m.Instance(), m.beforeItemErasePtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeItemPaint(fn TVTBeforeItemPaintEvent) {
	if m.beforeItemPaintPtr != 0 {
		RemoveEventElement(m.beforeItemPaintPtr)
	}
	m.beforeItemPaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3838, m.Instance(), m.beforeItemPaintPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeNodeExport(fn TVTNodeExportEvent) {
	if m.beforeNodeExportPtr != 0 {
		RemoveEventElement(m.beforeNodeExportPtr)
	}
	m.beforeNodeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3839, m.Instance(), m.beforeNodeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforePaint(fn TVTPaintEvent) {
	if m.beforePaintPtr != 0 {
		RemoveEventElement(m.beforePaintPtr)
	}
	m.beforePaintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3840, m.Instance(), m.beforePaintPtr)
}

func (m *TLazVirtualStringTree) SetOnBeforeTreeExport(fn TVTTreeExportEvent) {
	if m.beforeTreeExportPtr != 0 {
		RemoveEventElement(m.beforeTreeExportPtr)
	}
	m.beforeTreeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3841, m.Instance(), m.beforeTreeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeColumn(fn TVTCanSplitterResizeColumnEvent) {
	if m.canSplitterResizeColumnPtr != 0 {
		RemoveEventElement(m.canSplitterResizeColumnPtr)
	}
	m.canSplitterResizeColumnPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3842, m.Instance(), m.canSplitterResizeColumnPtr)
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeHeader(fn TVTCanSplitterResizeHeaderEvent) {
	if m.canSplitterResizeHeaderPtr != 0 {
		RemoveEventElement(m.canSplitterResizeHeaderPtr)
	}
	m.canSplitterResizeHeaderPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3843, m.Instance(), m.canSplitterResizeHeaderPtr)
}

func (m *TLazVirtualStringTree) SetOnCanSplitterResizeNode(fn TVTCanSplitterResizeNodeEvent) {
	if m.canSplitterResizeNodePtr != 0 {
		RemoveEventElement(m.canSplitterResizeNodePtr)
	}
	m.canSplitterResizeNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3844, m.Instance(), m.canSplitterResizeNodePtr)
}

func (m *TLazVirtualStringTree) SetOnChange(fn TVTChangeEvent) {
	if m.changePtr != 0 {
		RemoveEventElement(m.changePtr)
	}
	m.changePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3845, m.Instance(), m.changePtr)
}

func (m *TLazVirtualStringTree) SetOnChecked(fn TVTChangeEvent) {
	if m.checkedPtr != 0 {
		RemoveEventElement(m.checkedPtr)
	}
	m.checkedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3846, m.Instance(), m.checkedPtr)
}

func (m *TLazVirtualStringTree) SetOnChecking(fn TVTCheckChangingEvent) {
	if m.checkingPtr != 0 {
		RemoveEventElement(m.checkingPtr)
	}
	m.checkingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3847, m.Instance(), m.checkingPtr)
}

func (m *TLazVirtualStringTree) SetOnCollapsed(fn TVTChangeEvent) {
	if m.collapsedPtr != 0 {
		RemoveEventElement(m.collapsedPtr)
	}
	m.collapsedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3848, m.Instance(), m.collapsedPtr)
}

func (m *TLazVirtualStringTree) SetOnCollapsing(fn TVTChangingEvent) {
	if m.collapsingPtr != 0 {
		RemoveEventElement(m.collapsingPtr)
	}
	m.collapsingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3849, m.Instance(), m.collapsingPtr)
}

func (m *TLazVirtualStringTree) SetOnColumnClick(fn TVTColumnClickEvent) {
	if m.columnClickPtr != 0 {
		RemoveEventElement(m.columnClickPtr)
	}
	m.columnClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3850, m.Instance(), m.columnClickPtr)
}

func (m *TLazVirtualStringTree) SetOnColumnDblClick(fn TVTColumnDblClickEvent) {
	if m.columnDblClickPtr != 0 {
		RemoveEventElement(m.columnDblClickPtr)
	}
	m.columnDblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3851, m.Instance(), m.columnDblClickPtr)
}

func (m *TLazVirtualStringTree) SetOnColumnExport(fn TVTColumnExportEvent) {
	if m.columnExportPtr != 0 {
		RemoveEventElement(m.columnExportPtr)
	}
	m.columnExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3852, m.Instance(), m.columnExportPtr)
}

func (m *TLazVirtualStringTree) SetOnColumnResize(fn TVTHeaderNotifyEvent) {
	if m.columnResizePtr != 0 {
		RemoveEventElement(m.columnResizePtr)
	}
	m.columnResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3853, m.Instance(), m.columnResizePtr)
}

func (m *TLazVirtualStringTree) SetOnColumnWidthDblClickResize(fn TVTColumnWidthDblClickResizeEvent) {
	if m.columnWidthDblClickResizePtr != 0 {
		RemoveEventElement(m.columnWidthDblClickResizePtr)
	}
	m.columnWidthDblClickResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3854, m.Instance(), m.columnWidthDblClickResizePtr)
}

func (m *TLazVirtualStringTree) SetOnColumnWidthTracking(fn TVTColumnWidthTrackingEvent) {
	if m.columnWidthTrackingPtr != 0 {
		RemoveEventElement(m.columnWidthTrackingPtr)
	}
	m.columnWidthTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3855, m.Instance(), m.columnWidthTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnCompareNodes(fn TVTCompareEvent) {
	if m.compareNodesPtr != 0 {
		RemoveEventElement(m.compareNodesPtr)
	}
	m.compareNodesPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3856, m.Instance(), m.compareNodesPtr)
}

func (m *TLazVirtualStringTree) SetOnContextPopup(fn TContextPopupEvent) {
	if m.contextPopupPtr != 0 {
		RemoveEventElement(m.contextPopupPtr)
	}
	m.contextPopupPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3857, m.Instance(), m.contextPopupPtr)
}

func (m *TLazVirtualStringTree) SetOnCreateDataObject(fn TVTCreateDataObjectEvent) {
	if m.createDataObjectPtr != 0 {
		RemoveEventElement(m.createDataObjectPtr)
	}
	m.createDataObjectPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3858, m.Instance(), m.createDataObjectPtr)
}

func (m *TLazVirtualStringTree) SetOnCreateDragManager(fn TVTCreateDragManagerEvent) {
	if m.createDragManagerPtr != 0 {
		RemoveEventElement(m.createDragManagerPtr)
	}
	m.createDragManagerPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3859, m.Instance(), m.createDragManagerPtr)
}

func (m *TLazVirtualStringTree) SetOnCreateEditor(fn TVTCreateEditorEvent) {
	if m.createEditorPtr != 0 {
		RemoveEventElement(m.createEditorPtr)
	}
	m.createEditorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3860, m.Instance(), m.createEditorPtr)
}

func (m *TLazVirtualStringTree) SetOnDblClick(fn TNotifyEvent) {
	if m.dblClickPtr != 0 {
		RemoveEventElement(m.dblClickPtr)
	}
	m.dblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3861, m.Instance(), m.dblClickPtr)
}

func (m *TLazVirtualStringTree) SetOnDragAllowed(fn TVTDragAllowedEvent) {
	if m.dragAllowedPtr != 0 {
		RemoveEventElement(m.dragAllowedPtr)
	}
	m.dragAllowedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3862, m.Instance(), m.dragAllowedPtr)
}

func (m *TLazVirtualStringTree) SetOnDragOver(fn TVTDragOverEvent) {
	if m.dragOverPtr != 0 {
		RemoveEventElement(m.dragOverPtr)
	}
	m.dragOverPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3864, m.Instance(), m.dragOverPtr)
}

func (m *TLazVirtualStringTree) SetOnDragDrop(fn TVTDragDropEvent) {
	if m.dragDropPtr != 0 {
		RemoveEventElement(m.dragDropPtr)
	}
	m.dragDropPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3863, m.Instance(), m.dragDropPtr)
}

func (m *TLazVirtualStringTree) SetOnDrawHint(fn TVTDrawHintEvent) {
	if m.drawHintPtr != 0 {
		RemoveEventElement(m.drawHintPtr)
	}
	m.drawHintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3865, m.Instance(), m.drawHintPtr)
}

func (m *TLazVirtualStringTree) SetOnDrawText(fn TVTDrawTextEvent) {
	if m.drawTextPtr != 0 {
		RemoveEventElement(m.drawTextPtr)
	}
	m.drawTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3866, m.Instance(), m.drawTextPtr)
}

func (m *TLazVirtualStringTree) SetOnEditCancelled(fn TVTEditCancelEvent) {
	if m.editCancelledPtr != 0 {
		RemoveEventElement(m.editCancelledPtr)
	}
	m.editCancelledPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3867, m.Instance(), m.editCancelledPtr)
}

func (m *TLazVirtualStringTree) SetOnEdited(fn TVTEditChangeEvent) {
	if m.editedPtr != 0 {
		RemoveEventElement(m.editedPtr)
	}
	m.editedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3868, m.Instance(), m.editedPtr)
}

func (m *TLazVirtualStringTree) SetOnEditing(fn TVTEditChangingEvent) {
	if m.editingPtr != 0 {
		RemoveEventElement(m.editingPtr)
	}
	m.editingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3869, m.Instance(), m.editingPtr)
}

func (m *TLazVirtualStringTree) SetOnEndDock(fn TEndDragEvent) {
	if m.endDockPtr != 0 {
		RemoveEventElement(m.endDockPtr)
	}
	m.endDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3870, m.Instance(), m.endDockPtr)
}

func (m *TLazVirtualStringTree) SetOnEndDrag(fn TEndDragEvent) {
	if m.endDragPtr != 0 {
		RemoveEventElement(m.endDragPtr)
	}
	m.endDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3871, m.Instance(), m.endDragPtr)
}

func (m *TLazVirtualStringTree) SetOnEndOperation(fn TVTOperationEvent) {
	if m.endOperationPtr != 0 {
		RemoveEventElement(m.endOperationPtr)
	}
	m.endOperationPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3872, m.Instance(), m.endOperationPtr)
}

func (m *TLazVirtualStringTree) SetOnExpanded(fn TVTChangeEvent) {
	if m.expandedPtr != 0 {
		RemoveEventElement(m.expandedPtr)
	}
	m.expandedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3873, m.Instance(), m.expandedPtr)
}

func (m *TLazVirtualStringTree) SetOnExpanding(fn TVTChangingEvent) {
	if m.expandingPtr != 0 {
		RemoveEventElement(m.expandingPtr)
	}
	m.expandingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3874, m.Instance(), m.expandingPtr)
}

func (m *TLazVirtualStringTree) SetOnFocusChanged(fn TVTFocusChangeEvent) {
	if m.focusChangedPtr != 0 {
		RemoveEventElement(m.focusChangedPtr)
	}
	m.focusChangedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3875, m.Instance(), m.focusChangedPtr)
}

func (m *TLazVirtualStringTree) SetOnFocusChanging(fn TVTFocusChangingEvent) {
	if m.focusChangingPtr != 0 {
		RemoveEventElement(m.focusChangingPtr)
	}
	m.focusChangingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3876, m.Instance(), m.focusChangingPtr)
}

func (m *TLazVirtualStringTree) SetOnFreeNode(fn TVTFreeNodeEvent) {
	if m.freeNodePtr != 0 {
		RemoveEventElement(m.freeNodePtr)
	}
	m.freeNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3877, m.Instance(), m.freeNodePtr)
}

func (m *TLazVirtualStringTree) SetOnGetCellIsEmpty(fn TVTGetCellIsEmptyEvent) {
	if m.getCellIsEmptyPtr != 0 {
		RemoveEventElement(m.getCellIsEmptyPtr)
	}
	m.getCellIsEmptyPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3878, m.Instance(), m.getCellIsEmptyPtr)
}

func (m *TLazVirtualStringTree) SetOnGetCursor(fn TVTGetCursorEvent) {
	if m.getCursorPtr != 0 {
		RemoveEventElement(m.getCursorPtr)
	}
	m.getCursorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3879, m.Instance(), m.getCursorPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHeaderCursor(fn TVTGetHeaderCursorEvent) {
	if m.getHeaderCursorPtr != 0 {
		RemoveEventElement(m.getHeaderCursorPtr)
	}
	m.getHeaderCursorPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3880, m.Instance(), m.getHeaderCursorPtr)
}

func (m *TLazVirtualStringTree) SetOnGetText(fn TVSTGetTextEvent) {
	if m.getTextPtr != 0 {
		RemoveEventElement(m.getTextPtr)
	}
	m.getTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3891, m.Instance(), m.getTextPtr)
}

func (m *TLazVirtualStringTree) SetOnPaintText(fn TVTPaintText) {
	if m.paintTextPtr != 0 {
		RemoveEventElement(m.paintTextPtr)
	}
	m.paintTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3931, m.Instance(), m.paintTextPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHelpContext(fn TVTHelpContextEvent) {
	if m.getHelpContextPtr != 0 {
		RemoveEventElement(m.getHelpContextPtr)
	}
	m.getHelpContextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3881, m.Instance(), m.getHelpContextPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHintKind(fn TVTHintKindEvent) {
	if m.getHintKindPtr != 0 {
		RemoveEventElement(m.getHintKindPtr)
	}
	m.getHintKindPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3883, m.Instance(), m.getHintKindPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHintSize(fn TVTGetHintSizeEvent) {
	if m.getHintSizePtr != 0 {
		RemoveEventElement(m.getHintSizePtr)
	}
	m.getHintSizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3884, m.Instance(), m.getHintSizePtr)
}

func (m *TLazVirtualStringTree) SetOnGetImageIndex(fn TVTGetImageEvent) {
	if m.getImageIndexPtr != 0 {
		RemoveEventElement(m.getImageIndexPtr)
	}
	m.getImageIndexPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3885, m.Instance(), m.getImageIndexPtr)
}

func (m *TLazVirtualStringTree) SetOnGetImageIndexEx(fn TVTGetImageExEvent) {
	if m.getImageIndexExPtr != 0 {
		RemoveEventElement(m.getImageIndexExPtr)
	}
	m.getImageIndexExPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3886, m.Instance(), m.getImageIndexExPtr)
}

func (m *TLazVirtualStringTree) SetOnGetImageText(fn TVTGetImageTextEvent) {
	if m.getImageTextPtr != 0 {
		RemoveEventElement(m.getImageTextPtr)
	}
	m.getImageTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3887, m.Instance(), m.getImageTextPtr)
}

func (m *TLazVirtualStringTree) SetOnGetHint(fn TVSTGetHintEvent) {
	if m.getHintPtr != 0 {
		RemoveEventElement(m.getHintPtr)
	}
	m.getHintPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3882, m.Instance(), m.getHintPtr)
}

func (m *TLazVirtualStringTree) SetOnGetLineStyle(fn TVTGetLineStyleEvent) {
	if m.getLineStylePtr != 0 {
		RemoveEventElement(m.getLineStylePtr)
	}
	m.getLineStylePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3888, m.Instance(), m.getLineStylePtr)
}

func (m *TLazVirtualStringTree) SetOnGetNodeDataSize(fn TVTGetNodeDataSizeEvent) {
	if m.getNodeDataSizePtr != 0 {
		RemoveEventElement(m.getNodeDataSizePtr)
	}
	m.getNodeDataSizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3889, m.Instance(), m.getNodeDataSizePtr)
}

func (m *TLazVirtualStringTree) SetOnGetPopupMenu(fn TVTPopupEvent) {
	if m.getPopupMenuPtr != 0 {
		RemoveEventElement(m.getPopupMenuPtr)
	}
	m.getPopupMenuPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3890, m.Instance(), m.getPopupMenuPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderClick(fn TVTHeaderClickEvent) {
	if m.headerClickPtr != 0 {
		RemoveEventElement(m.headerClickPtr)
	}
	m.headerClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3892, m.Instance(), m.headerClickPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDblClick(fn TVTHeaderClickEvent) {
	if m.headerDblClickPtr != 0 {
		RemoveEventElement(m.headerDblClickPtr)
	}
	m.headerDblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3893, m.Instance(), m.headerDblClickPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDragged(fn TVTHeaderDraggedEvent) {
	if m.headerDraggedPtr != 0 {
		RemoveEventElement(m.headerDraggedPtr)
	}
	m.headerDraggedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3894, m.Instance(), m.headerDraggedPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDraggedOut(fn TVTHeaderDraggedOutEvent) {
	if m.headerDraggedOutPtr != 0 {
		RemoveEventElement(m.headerDraggedOutPtr)
	}
	m.headerDraggedOutPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3895, m.Instance(), m.headerDraggedOutPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDragging(fn TVTHeaderDraggingEvent) {
	if m.headerDraggingPtr != 0 {
		RemoveEventElement(m.headerDraggingPtr)
	}
	m.headerDraggingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3896, m.Instance(), m.headerDraggingPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDraw(fn TVTHeaderPaintEvent) {
	if m.headerDrawPtr != 0 {
		RemoveEventElement(m.headerDrawPtr)
	}
	m.headerDrawPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3897, m.Instance(), m.headerDrawPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderDrawQueryElements(fn TVTHeaderPaintQueryElementsEvent) {
	if m.headerDrawQueryElementsPtr != 0 {
		RemoveEventElement(m.headerDrawQueryElementsPtr)
	}
	m.headerDrawQueryElementsPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3898, m.Instance(), m.headerDrawQueryElementsPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderHeightDblClickResize(fn TVTHeaderHeightDblClickResizeEvent) {
	if m.headerHeightDblClickResizePtr != 0 {
		RemoveEventElement(m.headerHeightDblClickResizePtr)
	}
	m.headerHeightDblClickResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3899, m.Instance(), m.headerHeightDblClickResizePtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderHeightTracking(fn TVTHeaderHeightTrackingEvent) {
	if m.headerHeightTrackingPtr != 0 {
		RemoveEventElement(m.headerHeightTrackingPtr)
	}
	m.headerHeightTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3900, m.Instance(), m.headerHeightTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseDown(fn TVTHeaderMouseEvent) {
	if m.headerMouseDownPtr != 0 {
		RemoveEventElement(m.headerMouseDownPtr)
	}
	m.headerMouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3901, m.Instance(), m.headerMouseDownPtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseMove(fn TVTHeaderMouseMoveEvent) {
	if m.headerMouseMovePtr != 0 {
		RemoveEventElement(m.headerMouseMovePtr)
	}
	m.headerMouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3902, m.Instance(), m.headerMouseMovePtr)
}

func (m *TLazVirtualStringTree) SetOnHeaderMouseUp(fn TVTHeaderMouseEvent) {
	if m.headerMouseUpPtr != 0 {
		RemoveEventElement(m.headerMouseUpPtr)
	}
	m.headerMouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3903, m.Instance(), m.headerMouseUpPtr)
}

func (m *TLazVirtualStringTree) SetOnHotChange(fn TVTHotNodeChangeEvent) {
	if m.hotChangePtr != 0 {
		RemoveEventElement(m.hotChangePtr)
	}
	m.hotChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3904, m.Instance(), m.hotChangePtr)
}

func (m *TLazVirtualStringTree) SetOnIncrementalSearch(fn TVTIncrementalSearchEvent) {
	if m.incrementalSearchPtr != 0 {
		RemoveEventElement(m.incrementalSearchPtr)
	}
	m.incrementalSearchPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3905, m.Instance(), m.incrementalSearchPtr)
}

func (m *TLazVirtualStringTree) SetOnInitChildren(fn TVTInitChildrenEvent) {
	if m.initChildrenPtr != 0 {
		RemoveEventElement(m.initChildrenPtr)
	}
	m.initChildrenPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3906, m.Instance(), m.initChildrenPtr)
}

func (m *TLazVirtualStringTree) SetOnInitNode(fn TVTInitNodeEvent) {
	if m.initNodePtr != 0 {
		RemoveEventElement(m.initNodePtr)
	}
	m.initNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3907, m.Instance(), m.initNodePtr)
}

func (m *TLazVirtualStringTree) SetOnKeyAction(fn TVTKeyActionEvent) {
	if m.keyActionPtr != 0 {
		RemoveEventElement(m.keyActionPtr)
	}
	m.keyActionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3908, m.Instance(), m.keyActionPtr)
}

func (m *TLazVirtualStringTree) SetOnLoadNode(fn TVTSaveNodeEvent) {
	if m.loadNodePtr != 0 {
		RemoveEventElement(m.loadNodePtr)
	}
	m.loadNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3909, m.Instance(), m.loadNodePtr)
}

func (m *TLazVirtualStringTree) SetOnLoadTree(fn TVTSaveTreeEvent) {
	if m.loadTreePtr != 0 {
		RemoveEventElement(m.loadTreePtr)
	}
	m.loadTreePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3910, m.Instance(), m.loadTreePtr)
}

func (m *TLazVirtualStringTree) SetOnMeasureItem(fn TVTMeasureItemEvent) {
	if m.measureItemPtr != 0 {
		RemoveEventElement(m.measureItemPtr)
	}
	m.measureItemPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3911, m.Instance(), m.measureItemPtr)
}

func (m *TLazVirtualStringTree) SetOnMeasureTextWidth(fn TVTMeasureTextEvent) {
	if m.measureTextWidthPtr != 0 {
		RemoveEventElement(m.measureTextWidthPtr)
	}
	m.measureTextWidthPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3913, m.Instance(), m.measureTextWidthPtr)
}

func (m *TLazVirtualStringTree) SetOnMeasureTextHeight(fn TVTMeasureTextEvent) {
	if m.measureTextHeightPtr != 0 {
		RemoveEventElement(m.measureTextHeightPtr)
	}
	m.measureTextHeightPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3912, m.Instance(), m.measureTextHeightPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseDown(fn TMouseEvent) {
	if m.mouseDownPtr != 0 {
		RemoveEventElement(m.mouseDownPtr)
	}
	m.mouseDownPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3914, m.Instance(), m.mouseDownPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseMove(fn TMouseMoveEvent) {
	if m.mouseMovePtr != 0 {
		RemoveEventElement(m.mouseMovePtr)
	}
	m.mouseMovePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3917, m.Instance(), m.mouseMovePtr)
}

func (m *TLazVirtualStringTree) SetOnMouseUp(fn TMouseEvent) {
	if m.mouseUpPtr != 0 {
		RemoveEventElement(m.mouseUpPtr)
	}
	m.mouseUpPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3918, m.Instance(), m.mouseUpPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseWheel(fn TMouseWheelEvent) {
	if m.mouseWheelPtr != 0 {
		RemoveEventElement(m.mouseWheelPtr)
	}
	m.mouseWheelPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3919, m.Instance(), m.mouseWheelPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseEnter(fn TNotifyEvent) {
	if m.mouseEnterPtr != 0 {
		RemoveEventElement(m.mouseEnterPtr)
	}
	m.mouseEnterPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3915, m.Instance(), m.mouseEnterPtr)
}

func (m *TLazVirtualStringTree) SetOnMouseLeave(fn TNotifyEvent) {
	if m.mouseLeavePtr != 0 {
		RemoveEventElement(m.mouseLeavePtr)
	}
	m.mouseLeavePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3916, m.Instance(), m.mouseLeavePtr)
}

func (m *TLazVirtualStringTree) SetOnNewText(fn TVSTNewTextEvent) {
	if m.newTextPtr != 0 {
		RemoveEventElement(m.newTextPtr)
	}
	m.newTextPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3920, m.Instance(), m.newTextPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeClick(fn TVTNodeClickEvent) {
	if m.nodeClickPtr != 0 {
		RemoveEventElement(m.nodeClickPtr)
	}
	m.nodeClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3921, m.Instance(), m.nodeClickPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeCopied(fn TVTNodeCopiedEvent) {
	if m.nodeCopiedPtr != 0 {
		RemoveEventElement(m.nodeCopiedPtr)
	}
	m.nodeCopiedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3922, m.Instance(), m.nodeCopiedPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeCopying(fn TVTNodeCopyingEvent) {
	if m.nodeCopyingPtr != 0 {
		RemoveEventElement(m.nodeCopyingPtr)
	}
	m.nodeCopyingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3923, m.Instance(), m.nodeCopyingPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeDblClick(fn TVTNodeClickEvent) {
	if m.nodeDblClickPtr != 0 {
		RemoveEventElement(m.nodeDblClickPtr)
	}
	m.nodeDblClickPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3924, m.Instance(), m.nodeDblClickPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeExport(fn TVTNodeExportEvent) {
	if m.nodeExportPtr != 0 {
		RemoveEventElement(m.nodeExportPtr)
	}
	m.nodeExportPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3925, m.Instance(), m.nodeExportPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeHeightDblClickResize(fn TVTNodeHeightDblClickResizeEvent) {
	if m.nodeHeightDblClickResizePtr != 0 {
		RemoveEventElement(m.nodeHeightDblClickResizePtr)
	}
	m.nodeHeightDblClickResizePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3926, m.Instance(), m.nodeHeightDblClickResizePtr)
}

func (m *TLazVirtualStringTree) SetOnNodeHeightTracking(fn TVTNodeHeightTrackingEvent) {
	if m.nodeHeightTrackingPtr != 0 {
		RemoveEventElement(m.nodeHeightTrackingPtr)
	}
	m.nodeHeightTrackingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3927, m.Instance(), m.nodeHeightTrackingPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeMoved(fn TVTNodeMovedEvent) {
	if m.nodeMovedPtr != 0 {
		RemoveEventElement(m.nodeMovedPtr)
	}
	m.nodeMovedPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3928, m.Instance(), m.nodeMovedPtr)
}

func (m *TLazVirtualStringTree) SetOnNodeMoving(fn TVTNodeMovingEvent) {
	if m.nodeMovingPtr != 0 {
		RemoveEventElement(m.nodeMovingPtr)
	}
	m.nodeMovingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3929, m.Instance(), m.nodeMovingPtr)
}

func (m *TLazVirtualStringTree) SetOnPaintBackground(fn TVTBackgroundPaintEvent) {
	if m.paintBackgroundPtr != 0 {
		RemoveEventElement(m.paintBackgroundPtr)
	}
	m.paintBackgroundPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3930, m.Instance(), m.paintBackgroundPtr)
}

func (m *TLazVirtualStringTree) SetOnRemoveFromSelection(fn TVTRemoveFromSelectionEvent) {
	if m.removeFromSelectionPtr != 0 {
		RemoveEventElement(m.removeFromSelectionPtr)
	}
	m.removeFromSelectionPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3932, m.Instance(), m.removeFromSelectionPtr)
}

func (m *TLazVirtualStringTree) SetOnResetNode(fn TVTChangeEvent) {
	if m.resetNodePtr != 0 {
		RemoveEventElement(m.resetNodePtr)
	}
	m.resetNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3933, m.Instance(), m.resetNodePtr)
}

func (m *TLazVirtualStringTree) SetOnSaveNode(fn TVTSaveNodeEvent) {
	if m.saveNodePtr != 0 {
		RemoveEventElement(m.saveNodePtr)
	}
	m.saveNodePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3934, m.Instance(), m.saveNodePtr)
}

func (m *TLazVirtualStringTree) SetOnSaveTree(fn TVTSaveTreeEvent) {
	if m.saveTreePtr != 0 {
		RemoveEventElement(m.saveTreePtr)
	}
	m.saveTreePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3935, m.Instance(), m.saveTreePtr)
}

func (m *TLazVirtualStringTree) SetOnScroll(fn TVTScrollEvent) {
	if m.scrollPtr != 0 {
		RemoveEventElement(m.scrollPtr)
	}
	m.scrollPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3936, m.Instance(), m.scrollPtr)
}

func (m *TLazVirtualStringTree) SetOnShowScrollBar(fn TVTScrollBarShowEvent) {
	if m.showScrollBarPtr != 0 {
		RemoveEventElement(m.showScrollBarPtr)
	}
	m.showScrollBarPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3937, m.Instance(), m.showScrollBarPtr)
}

func (m *TLazVirtualStringTree) SetOnStartDock(fn TStartDockEvent) {
	if m.startDockPtr != 0 {
		RemoveEventElement(m.startDockPtr)
	}
	m.startDockPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3938, m.Instance(), m.startDockPtr)
}

func (m *TLazVirtualStringTree) SetOnStartDrag(fn TStartDragEvent) {
	if m.startDragPtr != 0 {
		RemoveEventElement(m.startDragPtr)
	}
	m.startDragPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3939, m.Instance(), m.startDragPtr)
}

func (m *TLazVirtualStringTree) SetOnStartOperation(fn TVTOperationEvent) {
	if m.startOperationPtr != 0 {
		RemoveEventElement(m.startOperationPtr)
	}
	m.startOperationPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3940, m.Instance(), m.startOperationPtr)
}

func (m *TLazVirtualStringTree) SetOnStateChange(fn TVTStateChangeEvent) {
	if m.stateChangePtr != 0 {
		RemoveEventElement(m.stateChangePtr)
	}
	m.stateChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3941, m.Instance(), m.stateChangePtr)
}

func (m *TLazVirtualStringTree) SetOnStructureChange(fn TVTStructureChangeEvent) {
	if m.structureChangePtr != 0 {
		RemoveEventElement(m.structureChangePtr)
	}
	m.structureChangePtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3942, m.Instance(), m.structureChangePtr)
}

func (m *TLazVirtualStringTree) SetOnUpdating(fn TVTUpdatingEvent) {
	if m.updatingPtr != 0 {
		RemoveEventElement(m.updatingPtr)
	}
	m.updatingPtr = MakeEventDataPtr(fn)
	LCL().SysCallN(3943, m.Instance(), m.updatingPtr)
}
