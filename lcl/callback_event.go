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
	"unsafe"
)

// TExtEventCallback
//
// 外部回调事件
// 参数一：函数地址
// 参数二：获取参数值的函数
// 返回值：如果为true则终止事件传递，如果为false则继续向后转发事件。
type TExtEventCallback func(fn interface{}, getVal func(idx int) uintptr) bool

// 外部扩展的事件回调，先不管重复注册的问题
var extEventCallback []TExtEventCallback

// RegisterExtEventCallback
//
// 注册外部扩展回调事件
//
// Registering external extension callback events.
func RegisterExtEventCallback(callback TExtEventCallback) {
	extEventCallback = append(extEventCallback, callback)
}

// getParam 从指定索引和地址获取事件中的参数
// 不再使用FreePascal导出的了，直接在这处理
func getParamOf(index int, ptr uintptr) uintptr {
	return *(*uintptr)(getPointer(ptr + uintptr(index)*unsafe.Sizeof(ptr)))
}

// 移除事件，释放相关的引用
func removeEventCallbackProc(f uintptr) uintptr {
	//RemoveEventElement(PtrToElementPtr(f))
	RemoveEventElement(f)
	return 0
}

func getPointer(ptr uintptr) unsafePointer {
	return unsafePointer(ptr)
}

// 回调过程
func eventCallbackProc(f uintptr, args uintptr, _ int) uintptr {
	fn := PtrToElementValue(f)
	if fn != nil {
		// 获取值
		getVal := func(i int) uintptr {
			return getParamOf(i, args)
		}
		// 指针
		getPtr := func(i int) unsafePointer {
			return getPointer(getVal(i))
		}
		// 指针的值
		getPtrVal := func(i int) uintptr {
			return *(*uintptr)(getPtr(i))
		}
		setPtrVal := func(i int, val uintptr) {
			*(*uintptr)(getPtr(i)) = val
		}
		getBoolPtr := func(i int) *bool {
			return (*bool)(getPtr(i))
		}
		getI32Ptr := func(i int) *int32 {
			return (*int32)(getPtr(i))
		}
		getRectPtr := func(i int) *TRect {
			return (*TRect)(getPtr(i))
		}
		getPointPtr := func(i int) *TPoint {
			return (*TPoint)(getPtr(i))
		}

		// 调用外部注册的事件回调过程
		for n := 0; n < len(extEventCallback); n++ {
			// 外部返回True则不继续下去了
			if extEventCallback[n](fn, getVal) {
				return 0
			}
		}

		switch fn.(type) {
		// func(sender IObject)
		case TNotifyEvent:
			fn.(TNotifyEvent)(AsObject(getVal(0)))

		// func(sender IObject, button TUDBtnType)
		case TUDClickEvent:
			fn.(TUDClickEvent)(AsObject(getVal(0)), TUDBtnType(getVal(1)))

		// func(sender IObject, item *TListItem, change int32)
		case TLVChangeEvent:
			fn.(TLVChangeEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), TItemChange(getVal(2)))

		// func(sender IObject, action *TCloseAction) // Action *uintptr
		case TCloseEvent:
			fn.(TCloseEvent)(AsObject(getVal(0)), (*TCloseAction)(getPtr(1)))

		// func(sender IObject, canClose *bool) //CanClose *uintptr
		case TCloseQueryEvent:
			fn.(TCloseQueryEvent)(AsObject(getVal(0)), getBoolPtr(1))

		// func(sender IObject, source *TMenuItem, rebuild bool)
		case TMenuChangeEvent:
			fn.(TMenuChangeEvent)(AsObject(getVal(0)), AsMenuItem(getVal(1)), GoBool(getVal(2)))

		// func(sender IObject, node *TreeNode)
		case TTVChangedEvent:
			fn.(TTVChangedEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)))

		// func(sender IObject, link string, linkType TSysLinkType)
		case TSysLinkEvent:
			fn.(TSysLinkEvent)(AsObject(getVal(0)), GoStr(getVal(1)), TSysLinkType(getVal(2)))

		// func(sender, e IObject)
		case TExceptionEvent:
			fn.(TExceptionEvent)(AsObject(getVal(0)), AsException(getVal(1)))

		// func(sender IObject, key *Char, shift TShiftState)
		case TKeyEvent:
			fn.(TKeyEvent)(AsObject(getVal(0)), (*Char)(getPtr(1)), TShiftState(getVal(2)))

		// func(sender IObject, key *Char)
		case TKeyPressEvent:
			fn.(TKeyPressEvent)(AsObject(getVal(0)), (*Char)(getPtr(1)))

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)
		case TMouseEvent:
			fn.(TMouseEvent)(AsObject(getVal(0)), TMouseButton(getVal(1)), TShiftState(getVal(2)), int32(getVal(3)), int32(getVal(4)))

		// func(sender IObject, shift TShiftState, x, y int32)
		case TMouseMoveEvent:
			fn.(TMouseMoveEvent)(AsObject(getVal(0)), TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)
		case TMouseWheelEvent:
			fn.(TMouseWheelEvent)(AsObject(getVal(0)), TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)), int32(getVal(4)), getBoolPtr(5))

		// func(control IWinControl, index int32, aRect TRect, state TOwnerDrawState)
		case TDrawItemEvent:
			fn.(TDrawItemEvent)(AsWinControl(getVal(0)), int32(getVal(1)), *getRectPtr(2), TOwnerDrawState(getVal(3)))

		// func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)
		case TMenuDrawItemEvent:
			fn.(TMenuDrawItemEvent)(AsObject(getVal(0)), AsCanvas(getVal(1)), *getRectPtr(2), TOwnerDrawState(getVal(3)))

		// func(sender IObject, item *TListItem)
		case TLVNotifyEvent:
			fn.(TLVNotifyEvent)(AsObject(getVal(0)), AsListItem(getVal(1)))

		// func(sender IObject, column *TListColumn)
		case TLVColumnClickEvent:
			fn.(TLVColumnClickEvent)(AsObject(getVal(0)), AsListColumn(getVal(1)))

		// func(sender IObject, column *TListColumn, point TPoint)
		case TLVColumnRClickEvent:
			fn.(TLVColumnRClickEvent)(AsObject(getVal(0)), AsListColumn(getVal(1)), TPoint{X: int32(getVal(2)), Y: int32(getVal(3))})

		// func(sender IObject, item *TListItem, selected bool)
		case TLVSelectItemEvent:
			fn.(TLVSelectItemEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), GoBool(getVal(2)))

		//  func(sender IObject, item *TListItem)
		case TLVCheckedItemEvent:
			fn.(TLVCheckedItemEvent)(AsObject(getVal(0)), AsListItem(getVal(1)))

		// func(sender IObject, tabIndex int32, imageIndex *int32)
		case TTabGetImageEvent:
			fn.(TTabGetImageEvent)(AsObject(getVal(0)), int32(getVal(1)), getI32Ptr(2))

		// func(sender IObject, node *TTreeNode)
		case TTVExpandedEvent:
			fn.(TTVExpandedEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)))

		// func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)
		case TLVCompareEvent:
			fn.(TLVCompareEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), AsListItem(getVal(2)), int32(getVal(3)), getI32Ptr(4))

		// func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)
		case TTVCompareEvent:
			fn.(TTVCompareEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), AsTreeNode(getVal(2)), int32(getVal(3)), getI32Ptr(4))

		// func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTVAdvancedCustomDrawEvent:
			fn.(TTVAdvancedCustomDrawEvent)(AsTreeView(getVal(0)), *getRectPtr(1), TCustomDrawStage(getVal(2)), getBoolPtr(3))

		// func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)
		case TTVAdvancedCustomDrawItemEvent:
			fn.(TTVAdvancedCustomDrawItemEvent)(AsTreeView(getVal(0)), AsTreeNode(getVal(1)), TCustomDrawState(getVal(2)), TCustomDrawStage(getVal(3)),
				getBoolPtr(4), getBoolPtr(5))

		// func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawEvent:
			fn.(TLVAdvancedCustomDrawEvent)(AsListView(getVal(0)), *getRectPtr(1), TCustomDrawStage(getVal(2)), getBoolPtr(3))

		// func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawItemEvent:
			fn.(TLVAdvancedCustomDrawItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), TCustomDrawState(getVal(2)), TCustomDrawStage(getVal(3)),
				getBoolPtr(4))

		// func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawSubItemEvent:
			fn.(TLVAdvancedCustomDrawSubItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), int32(getVal(2)), TCustomDrawState(getVal(3)),
				TCustomDrawStage(getVal(4)), getBoolPtr(5))

		// func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTBAdvancedCustomDrawEvent:
			fn.(TTBAdvancedCustomDrawEvent)(AsToolBar(getVal(0)), *getRectPtr(1), TCustomDrawStage(getVal(2)), getBoolPtr(3))

		// func(sender IObject, aFileNames []string)
		case TDropFilesEvent:
			nLen := int(getVal(2))
			tempArr := make([]string, nLen)
			p := getVal(1)
			for i := 0; i < nLen; i++ {
				tempArr[i] = DGetStringArrOf(p, i)
			}
			fn.(TDropFilesEvent)(AsObject(getVal(0)), tempArr)

		// func(sender IObject, minWidth, minHeight, maxWidth, maxHeight *int32)
		case TConstrainedResizeEvent:
			fn.(TConstrainedResizeEvent)(AsObject(getVal(0)), getI32Ptr(1), getI32Ptr(2), getI32Ptr(3), getI32Ptr(4))

		// func(command uint16, data THelpEventData, callHelp *bool) bool
		case THelpEvent:
			fn.(THelpEvent)(uint16(getVal(0)), THelpEventData(getVal(1)), getBoolPtr(2), getBoolPtr(3))

		// func(msg *TWMKey, handled *bool)
		case TShortCutEvent:
			fn.(TShortCutEvent)((*TWMKey)(getPtr(0)), getBoolPtr(1))

		// func(sender IObject, mousePos TPoint, handled *bool)
		case TContextPopupEvent:
			fn.(TContextPopupEvent)(AsObject(getVal(0)), *getPointPtr(1), getBoolPtr(2))

		// func(sender, source IObject, x, y int32, state TDragState, accept *bool)
		case TDragOverEvent:
			fn.(TDragOverEvent)(AsObject(getVal(0)), AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)), TDragState(getVal(4)), getBoolPtr(5))

		//func(sender, source IObject, x, y int32)
		case TDragDropEvent:
			fn.(TDragDropEvent)(AsObject(getVal(0)), AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		//func(sender IObject, dragObject *TDragObject)
		case TStartDragEvent:
			obj := AsDragObject(getVal(1))
			fn.(TStartDragEvent)(AsObject(getVal(0)), &obj)
			if obj != nil {
				*(*uintptr)(unsafePointer(getVal(1))) = obj.Instance()
			}

		//func(sender, target IObject, x, y int32)
		case TEndDragEvent:
			fn.(TEndDragEvent)(AsObject(getVal(0)), AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, source *TDragDockObject, x, y int32)
		case TDockDropEvent:
			fn.(TDockDropEvent)(AsObject(getVal(0)), AsDragDockObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		//func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)
		case TDockOverEvent:
			fn.(TDockOverEvent)(AsObject(getVal(0)), AsDragDockObject(getVal(1)), int32(getVal(2)), int32(getVal(3)),
				TDragState(getVal(4)), getBoolPtr(5))

		//func(sender IObject, client *TControl, newTarget *TControl, allow *bool)
		case TUnDockEvent:
			fn.(TUnDockEvent)(AsObject(getVal(0)), AsControl(getVal(1)), AsControl(getVal(2)), getBoolPtr(3))

		//func(sender IObject, dragObject **TDragDockObject)
		case TStartDockEvent:
			obj := AsDragDockObject(getPtrVal(1))
			fn.(TStartDockEvent)(AsObject(getVal(0)), &obj)
			if obj != nil {
				setPtrVal(1, obj.Instance())
			}

		//func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)
		case TGetSiteInfoEvent:
			fn.(TGetSiteInfoEvent)(AsObject(getVal(0)), AsControl(getVal(1)), getRectPtr(2), *getPointPtr(3), getBoolPtr(4))

		//func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)
		case TMouseWheelUpDownEvent:
			fn.(TMouseWheelUpDownEvent)(AsObject(getVal(0)), TShiftState(getVal(1)), *getPointPtr(2), getBoolPtr(3))

		// func(sender IObject, isColumn bool, sIndex, tIndex int32)
		case TGridOperationEvent:
			fn.(TGridOperationEvent)(AsObject(getVal(0)), GoBool(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, aCol, aRow int32, aRect TRect, state TGridDrawState)
		case TDrawCellEvent:
			fn.(TDrawCellEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), *getRectPtr(3), TGridDrawState(getVal(4)))

		// func(sender IObject, aCol, aRow int32)
		case TFixedCellClickEvent:
			fn.(TFixedCellClickEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// func(sender IObject, aCol, aRow int32, value *string)
		case TGetEditEvent:
			str := GoStr(getPtrVal(3))
			fn.(TGetEditEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), &str)
			setPtrVal(3, PascalStr(str))

		// func(sender IObject, aCol, aRow int32, canSelect *bool)
		case TSelectCellEvent:
			fn.(TSelectCellEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), getBoolPtr(3))

		// func(sender IObject, aCol, aRow int32, value string)
		case TSetEditEvent:
			fn.(TSetEditEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), GoStr(getVal(3)))

		// func(headerControl *THeaderControl, section *THeaderSection, aRect TRect, pressed bool)
		case TDrawSectionEvent:
			fn.(TDrawSectionEvent)(AsHeaderControl(getVal(0)), AsHeaderSection(getVal(1)), *getRectPtr(2), getVal(3) != 0)

		// func(headerControl *THeaderControl, section *THeaderSection)
		case TSectionNotifyEvent:
			fn.(TSectionNotifyEvent)(AsHeaderControl(getVal(0)), AsHeaderSection(getVal(1)))

		// func(headerControl *THeaderControl, section *THeaderSection, width int32, state TSectionTrackState)
		case TSectionTrackEvent:
			fn.(TSectionTrackEvent)(AsHeaderControl(getVal(0)), AsHeaderSection(getVal(1)), int32(getVal(2)), TSectionTrackState(getVal(3)))

		// func(sender IObject, fromSection, toSection *THeaderSection, allowDrag *bool)
		case TSectionDragEvent:
			fn.(TSectionDragEvent)(AsObject(getVal(0)), AsHeaderSection(getVal(1)), AsHeaderSection(getVal(2)), getBoolPtr(3))

		// func(headerControl *THeaderControl, section *THeaderSection)
		case TCustomSectionNotifyEvent:
			fn.(TCustomSectionNotifyEvent)(AsHeaderControl(getVal(0)), AsHeaderSection(getVal(1)))

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32, hitTest int32, mouseActivate *TMouseActivate)
		case TMouseActivateEvent:
			fn.(TMouseActivateEvent)(AsObject(getVal(0)), TMouseButton(getVal(1)), TShiftState(getVal(2)), int32(getVal(3)), int32(getVal(4)),
				int32(getVal(5)), (*TMouseActivate)(getPtr(6)))

		// func(control *TWinControl, index int32, data *string)
		case TLBGetDataEvent:
			str := GoStr(getPtrVal(2))
			fn.(TLBGetDataEvent)(AsWinControl(getVal(0)), int32(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(control *TWinControl, index int32, dataObject IObject)
		case TLBGetDataObjectEvent:
			fn.(TLBGetDataObjectEvent)(AsWinControl(getVal(0)), int32(getVal(1)), AsObject(getVal(2))) // 这个参数要改，先这样

		// func(control *TWinControl, findString string) int32
		case TLBFindDataEvent:
			result := fn.(TLBFindDataEvent)(AsWinControl(getVal(0)), GoStr(getVal(1)))
			*getI32Ptr(2) = result

		// func(control *TWinControl, index int32, height *int32)
		case TMeasureItemEvent:
			fn.(TMeasureItemEvent)(AsWinControl(getVal(0)), int32(getVal(1)), getI32Ptr(2))

		// func(sender IObject, item *TListItem, change TItemChange, allowChange *bool)
		case TLVChangingEvent:
			fn.(TLVChangingEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), TItemChange(getVal(2)), getBoolPtr(3))

		// func(sender IObject, item *TListItem)
		case TLVDataEvent:
			fn.(TLVDataEvent)(AsObject(getVal(0)), AsListItem(getVal(1)))

		// func(sender IObject, item *TListItem)
		//case TLVDeletedEvent:
		//	fn.(TLVDeletedEvent)(AsObject(getVal(0)), AsListItem(getVal(1)))

		// func(sender IObject, find TItemFind, findString string, findPosition TPoint, findData TCustomData, startIndex int32,
		//	direction TSearchDirection, warp bool, index *int32)
		case TLVDataFindEvent:
			fn.(TLVDataFindEvent)(AsObject(getVal(0)), TItemFind(getVal(1)), GoStr(getVal(2)), *getPointPtr(3), TCustomData(getVal(4)),
				int32(getVal(5)), TSearchDirection(getVal(6)), GoBool(getVal(7)), getI32Ptr(8))

		// func(sender IObject, item *TListItem, allowEdit *bool)
		case TLVEditingEvent:
			fn.(TLVEditingEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), getBoolPtr(2))

		// func(sender IObject, item *TListItem, s *string)
		case TLVEditedEvent:
			str := GoStr(getPtrVal(2))
			fn.(TLVEditedEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(sender IObject, aCanvas *TCanvas, width, height *int32)
		case TMenuMeasureItemEvent:
			fn.(TMenuMeasureItemEvent)(AsObject(getVal(0)), AsCanvas(getVal(1)), getI32Ptr(2), getI32Ptr(3))

		//type func(sender IObject, allowChange *bool)
		case TTabChangingEvent:
			fn.(TTabChangingEvent)(AsObject(getVal(0)), getBoolPtr(1))

		// func(sender IObject, node *TTreeNode, allowChange *bool)
		case TTVChangingEvent:
			fn.(TTVChangingEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), getBoolPtr(2))

		// func(sender IObject, node *TTreeNode, allowCollapse *bool)
		case TTVCollapsingEvent:
			fn.(TTVCollapsingEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), getBoolPtr(2))

		// func(sender IObject, node *TTreeNode, s *string)
		case TTVEditedEvent:
			str := GoStr(getPtrVal(2))
			fn.(TTVEditedEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(sender IObject, node *TTreeNode, allowEdit *bool)
		case TTVEditingEvent:
			fn.(TTVEditingEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), getBoolPtr(2))

		// func(sender IObject, node *TTreeNode, allowExpansion *bool)
		case TTVExpandingEvent:
			fn.(TTVExpandingEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), getBoolPtr(2))

		// func(sender IObject, node *TTreeNode, hint *string)
		case TTVHintEvent:
			str := GoStr(getPtrVal(2))
			fn.(TTVHintEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(sender IObject, allowChange *bool)
		case TUDChangingEvent:
			fn.(TUDChangingEvent)(AsObject(getVal(0)), getBoolPtr(1))

		// func(sender IObject, winErrorCode uint32, errorDescription string, handled *bool)
		case TCreatingListErrorEvent:
			fn.(TCreatingListErrorEvent)(AsObject(getVal(0)), uint32(getVal(1)), GoStr(getVal(2)), getBoolPtr(3))

		// func(sender *TListView, aRect TRect, defaultDraw *bool)
		case TLVCustomDrawEvent:
			fn.(TLVCustomDrawEvent)(AsListView(getVal(0)), *getRectPtr(1), getBoolPtr(2))

		// func(sender *TListView, item *TListItem, state TCustomDrawStage, defaultDraw *bool)
		case TLVCustomDrawItemEvent:
			fn.(TLVCustomDrawItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), TCustomDrawState(getVal(2)), getBoolPtr(3))

		// func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawStage, defaultDraw *bool)
		case TLVCustomDrawSubItemEvent:
			fn.(TLVCustomDrawSubItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), int32(getVal(2)), TCustomDrawState(getVal(3)), getBoolPtr(4))

		// func(sender *TListView, item *TListItem, rect TRect, state TOwnerDrawState)
		case TLVDrawItemEvent:
			fn.(TLVDrawItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), *getRectPtr(2), TOwnerDrawState(getVal(3)))

		// func(sender IObject, startIndex, endIndex int32)
		case TLVDataHintEvent:
			fn.(TLVDataHintEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// func(sender *TTreeView, aRect TRect, defaultDraw *bool)
		case TTVCustomDrawEvent:
			fn.(TTVCustomDrawEvent)(AsTreeView(getVal(0)), *getRectPtr(1), getBoolPtr(2))

		// func(sender *TTreeView, node *TTreeNode, state TCustomDrawStage, defaultDraw *bool)
		case TTVCustomDrawItemEvent:
			fn.(TTVCustomDrawItemEvent)(AsTreeView(getVal(0)), AsTreeNode(getVal(1)), TCustomDrawState(getVal(2)), getBoolPtr(3))

		// func(sender IObject, text string)
		case TWebTitleChangeEvent:
			fn.(TWebTitleChangeEvent)(AsObject(getVal(0)), GoStr(getVal(1)))

		// func(sender IObject, funcName, args string, retVal *string)
		case TWebJSExternalEvent:
			str := GoStr(getPtrVal(3))
			fn.(TWebJSExternalEvent)(AsObject(getVal(0)), GoStr(getVal(1)), GoStr(getVal(2)), &str)
			setPtrVal(3, PascalStr(str))

		// func(sender IObject, modalResult TModalResult, canClose *bool)
		case TTaskDlgClickEvent:
			fn.(TTaskDlgClickEvent)(AsObject(getVal(0)), TModalResult(getVal(1)), getBoolPtr(2))

		// func(sender IObject, tickCount uint32, reset *bool)
		case TTaskDlgTimerEvent:
			fn.(TTaskDlgTimerEvent)(AsObject(getVal(0)), uint32(getVal(1)), getBoolPtr(2))

		// func(sender *TWinControl, control *TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *TRect, alignInfo TAlignInfo)
		case TAlignPositionEvent:
			fn.(TAlignPositionEvent)(AsWinControl(getVal(0)), AsControl(getVal(1)), getI32Ptr(2), getI32Ptr(3), getI32Ptr(4),
				getI32Ptr(5), getRectPtr(6), *(*TAlignInfo)(getPtr(7)))

		// func(sender IObject, index int32)
		case TCheckGroupClicked:
			fn.(TCheckGroupClicked)(AsObject(getVal(0)), int32(getVal(1)))

		// func(sender IObject, aCol, aRow int32)
		case TOnSelectEvent:
			fn.(TOnSelectEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// func(sender IObject, aCol, aRow int32, aState TCheckBoxState)
		case TToggledCheckboxEvent:
			fn.(TToggledCheckboxEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCheckBoxState(getVal(3)))

		// func(sender IObject, ACol, ARow, BCol, BRow int32, result *int32)
		case TOnCompareCells:
			fn.(TOnCompareCells)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), int32(getVal(3)), int32(getVal(4)), getI32Ptr(5))

		// func(sender IObject, ACol, ARow int32, hintText *string)
		case TGetCellHintEvent:
			str := GoStr(getPtrVal(3))
			fn.(TGetCellHintEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), &str)
			setPtrVal(3, PascalStr(str))

		// func(sender IObject, ACol, ARow int32, value *TCheckBoxState)
		case TGetCheckboxStateEvent:
			fn.(TGetCheckboxStateEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), (*TCheckBoxState)(getPtr(3)))

		// func(sender IObject, ACol, ARow int32, Value TCheckBoxState)
		case TSetCheckboxStateEvent:
			fn.(TSetCheckboxStateEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCheckBoxState(getVal(3)))

		// func(sender IObject, isColumn bool, index int32)
		case THdrEvent:
			fn.(THdrEvent)(AsObject(getVal(0)), GoBool(getVal(1)), int32(getVal(2)))

		// func(sender IObject, isColumn bool, aIndex, aSize int32)
		case THeaderSizingEvent:
			fn.(THeaderSizingEvent)(AsObject(getVal(0)), GoBool(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, aCol, aRow int32, editor **TWinControl)
		case TSelectEditorEvent:
			obj := AsWinControl(getPtrVal(3))
			fn.(TSelectEditorEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), &obj)
			if obj != nil {
				setPtrVal(3, obj.Instance())
			}

		// func(sender IObject, aCol, aRow int32, CheckedState TCheckBoxState, aBitmap **TBitmap)
		case TUserCheckBoxBitmapEvent:
			obj := AsBitmap(getPtrVal(4))
			fn.(TUserCheckBoxBitmapEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCheckBoxState(getVal(3)), &obj)
			if obj != nil {
				setPtrVal(4, obj.Instance())
			}

		// func(sender IObject, aCol, aRow int32, oldValue string, newValue *string)
		case TValidateEntryEvent:
			str := GoStr(getPtrVal(4))
			fn.(TValidateEntryEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), GoStr(getVal(3)), &str)
			setPtrVal(4, PascalStr(str))

		// func(sender IObject, aCol, aRow int32, aState TGridDrawState)
		case TOnPrepareCanvasEvent:
			fn.(TOnPrepareCanvasEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TGridDrawState(getVal(3)))

		// func(sender IObject, value *string)
		case TAcceptFileNameEvent:
			str := GoStr(getPtrVal(1))
			fn.(TAcceptFileNameEvent)(AsObject(getVal(0)), &str)
			setPtrVal(1, PascalStr(str))

		// func(sender IObject, index int32)
		case TCheckItemChange:
			fn.(TCheckItemChange)(AsObject(getVal(0)), int32(getVal(1)))

		// func(sender IObject, utf8key *TUTF8Char)
		case TUTF8KeyPressEvent:
			fn.(TUTF8KeyPressEvent)(AsObject(getVal(0)), (*TUTF8Char)(getPtr(1)))

		// type  func(sender IObject, aCanvas *TCanvas, aRect TRect)
		case TImagePaintBackgroundEvent:
			fn.(TImagePaintBackgroundEvent)(AsObject(getVal(0)), AsCanvas(getVal(1)), *getRectPtr(2))

		// func(sender IObject, Action IBasicAction, handled *bool)
		case TActionEvent:
			fn.(TActionEvent)(AsObject(getVal(0)), AsBasicAction(getVal(1)), getBoolPtr(2))

		// func(handle *HWND)
		case TGetHandleEvent:
			fn.(TGetHandleEvent)((*HWND)(getPtr(0)))

		// func(sender IObject, result int32)
		case TModalDialogFinished:
			fn.(TModalDialogFinished)(AsObject(getVal(0)), int32(getVal(1)))

		// func(sender IObject, control IControl, caption *string)
		case TGetDockCaptionEvent:
			str := GoStr(getPtrVal(2))
			fn.(TGetDockCaptionEvent)(AsObject(getVal(0)), AsControl(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(sender IObject, aCol, aRow int32, aRect TRect, aState TGridDrawState)
		case TOnDrawCell:
			fn.(TOnDrawCell)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), *getRectPtr(3), TGridDrawState(getVal(4)))

		// func(sender IObject, aCol, aRow int32, checkedState TCheckBoxState, imageList ICustomImageList, imageIndex TImageIndex)
		case TUserCheckBoxImageEvent:
			fn.(TUserCheckBoxImageEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCheckBoxState(getVal(3)), AsCustomImageList(getVal(4)),
				TImageIndex(getVal(5)))

		// func(sender IWinControl, control1, control2 IControl) bool
		case TAlignInsertBeforeEvent:
			fn.(TAlignInsertBeforeEvent)(AsWinControl(getVal(0)), AsControl(getVal(1)), AsControl(getVal(2)))

		// func(sender IObject, AllowChange *bool, newValue int16, direction TUpDownDirection)
		case TUDChangingEventEx:
			fn.(TUDChangingEventEx)(AsObject(getVal(0)), getBoolPtr(1), int16(getVal(2)), TUpDownDirection(getVal(3)))

		// func(sender IStatusBar, panelClass *TStatusPanelClass)
		case TSBCreatePanelClassEvent:
			fn.(TSBCreatePanelClassEvent)(AsStatusBar(getVal(0)), (*TStatusPanelClass)(getPtr(1)))

		// func(statusBar IStatusBar, panel IStatusPanel, rect TRect)
		case TDrawPanelEvent:
			fn.(TDrawPanelEvent)(AsStatusBar(getVal(0)), AsStatusPanel(getVal(1)), *getRectPtr(2))

		// func(sender IObject, success bool)
		case TDialogResultEvent:
			fn.(TDialogResultEvent)(AsObject(getVal(0)), *getBoolPtr(1))

		// func(sender ICustomColorBox, items IStrings)
		case TGetColorsEvent:
			fn.(TGetColorsEvent)(AsCustomColorBox(getVal(0)), AsStrings(getVal(1)))

		// func(hintStr *string, canShow bool)
		case THintEvent:
			str := GoStr(getPtrVal(0))
			fn.(THintEvent)(&str, *getBoolPtr(1))
			setPtrVal(0, PascalStr(str))

		// func(sender IObject, aCol, aRow int32, canSelect *bool)
		case TOnSelectCellEvent:
			fn.(TOnSelectCellEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), getBoolPtr(3))

		// func(sender IObject, user bool)
		case TSelectionChangeEvent:
			fn.(TSelectionChangeEvent)(AsObject(getVal(0)), *getBoolPtr(1))

		// func(color1, color2 TFPColor) TFPColor
		case TFPCanvasCombineColors:
			fn.(TFPCanvasCombineColors)(*(*TFPColor)(getPtr(0)), *(*TFPColor)(getPtr(1)), (*TFPColor)(getPtr(2)))

		// func(sender IObject, stage TFPImgProgressStage, percentDone byte, redrawNow bool, rect TRect, msg string, continue_ *bool)
		case TFPImgProgressEvent:
			fn.(TFPImgProgressEvent)(AsObject(getVal(0)), TFPImgProgressStage(getVal(1)), byte(getVal(2)), *getBoolPtr(3), *getRectPtr(4),
				GoStr(getVal(5)), getBoolPtr(6))

		// func(sender ICustomColorListBox, items IStrings)
		case TLBGetColorsEvent:
			fn.(TLBGetColorsEvent)(AsCustomColorListBox(getVal(0)), AsStrings(getVal(1)))

		// func(sender ICustomHeaderControl, sectionClass *THeaderSectionClass)
		case TCustomHCCreateSectionClassEvent:
			fn.(TCustomHCCreateSectionClassEvent)(AsCustomHeaderControl(getVal(0)), (*THeaderSectionClass)(getPtr(1)))

		// func(headerControl ICustomHeaderControl, section IHeaderSection, width int32, state TSectionTrackState)
		case TCustomSectionTrackEvent:
			fn.(TCustomSectionTrackEvent)(AsCustomHeaderControl(getVal(0)), AsHeaderSection(getVal(1)), int32(getVal(2)), TSectionTrackState(getVal(3)))

		// func(sender IObject, newOffset *int32, accept *bool)
		case TCanOffsetEvent:
			fn.(TCanOffsetEvent)(AsObject(getVal(0)), getI32Ptr(1), getBoolPtr(2))

		// func(sender IObject, newSize *int32, accept *bool)
		case TCanResizeEvent:
			fn.(TCanResizeEvent)(AsObject(getVal(0)), getI32Ptr(1), getBoolPtr(2))

		// func(sender IObject, keyName string, values IStrings)
		case TGetPickListEvent:
			fn.(TGetPickListEvent)(AsObject(getVal(0)), GoStr(getVal(1)), AsStrings(getVal(2)))

		// func(sender IObject, aCol, aRow int32, keyName, keyValue string)
		case TOnValidateEvent:
			fn.(TOnValidateEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), GoStr(getVal(3)), GoStr(getVal(4)))

		// func(requestedFormatID TClipboardFormat, data IStream)
		case TClipboardRequestEvent:
			fn.(TClipboardRequestEvent)(getVal(0), AsStream(getVal(1)))

		// func(sender IToolButton, state int32)
		case TToolBarOnPaintButton:
			fn.(TToolBarOnPaintButton)(AsToolButton(getVal(0)), int32(getVal(1)))

		// func(sender IObject, aCol, aRow int32, processType TCellProcessType, aValue *string)
		case TCellProcessEvent:
			str := GoStr(getPtrVal(4))
			fn.(TCellProcessEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCellProcessType(getVal(3)), &str)
			setPtrVal(4, PascalStr(str))

		// func(sender IObject, scrollCode TScrollCode, scrollPos *int32)
		case TScrollEvent:
			fn.(TScrollEvent)(AsObject(getVal(0)), TScrollCode(getVal(1)), getI32Ptr(2))

		// func(aList IListControlItems, aItem1, aItem2 IListControlItem) int32
		case TListCompareEvent:
			fn.(TListCompareEvent)(AsListControlItems(getVal(0)), AsListControlItem(getVal(1)), AsListControlItem(getVal(2)), getI32Ptr(3))

		// func(sender ICustomImageList, aImageWidth, aPPI int32, aResultWidth *int32)
		case TCustomImageListGetWidthForPPI:
			fn.(TCustomImageListGetWidthForPPI)(AsCustomImageList(getVal(0)), int32(getVal(1)), int32(getVal(2)), getI32Ptr(3))

		// func(sender ICustomListView, itemClass *TListItemClass)
		case TLVCreateItemClassEvent:
			fn.(TLVCreateItemClassEvent)(AsCustomListView(getVal(0)), (*TListItemClass)(getPtr(1)))

		// func(sender IObject, startIndex, endIndex int32, oldState, newState TListItemStates)
		case TLVDataStateChangeEvent:
			fn.(TLVDataStateChangeEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TListItemStates(getVal(3)), TListItemStates(getVal(4)))

		// func(sender ICustomTreeView, nodeClass *TTreeNodeClass)
		case TTVCreateNodeClassEvent:
			fn.(TTVCreateNodeClassEvent)(AsCustomTreeView(getVal(0)), (*TTreeNodeClass)(getPtr(1)))

		// func(sender ICustomTreeView, aTreeNode ITreeNode)
		case TTVCustomCreateNodeEvent:
			fn.(TTVCustomCreateNodeEvent)(AsCustomTreeView(getVal(0)), AsTreeNode(getVal(1)))

		// func(sender ICustomTreeView, ARect TRect, aCollapsed bool)
		case TTVCustomDrawArrowEvent:
			fn.(TTVCustomDrawArrowEvent)(AsCustomTreeView(getVal(0)), *getRectPtr(1), *getBoolPtr(2))

		// func(sender IObject, node ITreeNode, cancel bool)
		case TTVEditingEndEvent:
			fn.(TTVEditingEndEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), *getBoolPtr(2))

		// func(sender ICustomTreeView, aNode ITreeNode) bool
		case TTVHasChildrenEvent:
			result := getBoolPtr(2)
			*result = fn.(TTVHasChildrenEvent)(AsCustomTreeView(getVal(0)), AsTreeNode(getVal(1)))

		// func(sender IObject, node ITreeNode, changeReason TTreeNodeChangeReason)
		case TTVNodeChangedEvent:
			fn.(TTVNodeChangedEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), TTreeNodeChangeReason(getVal(2)))

		// func(hintStr *string, canShow *bool, hintInfo *THintInfo) // var hintInfo
		case TShowHintEvent:
			str := GoStr(getPtrVal(0))
			hitInfoPtr := (*tHintInfo)(getPtr(2))
			hitInfo := hitInfoPtr.Convert()
			fn.(TShowHintEvent)(&str, getBoolPtr(1), hitInfo)
			setPtrVal(0, PascalStr(str))
			hitInfo.SetInstanceValue()

		// func(sender IObject, hintInfo *HintInfo)
		case TControlShowHintEvent:
			hitInfoPtr := (*tHintInfo)(getPtr(1))
			hitInfo := hitInfoPtr.Convert()
			fn.(TControlShowHintEvent)(AsObject(getVal(0)), *hitInfo)

		// func(sender ICustomImageList, aWidth int32, aReferenceHandle TLCLHandle)
		case TDestroyResolutionHandleEvent:
			fn.(TDestroyResolutionHandleEvent)(AsCustomImageList(getVal(0)), int32(getVal(1)), TLCLHandle(getVal(2)))

		// type TTreeNodeCompare func(node1, node2 ITreeNode) int32
		case TTreeNodeCompare:
			var result = getI32Ptr(2)
			*result = fn.(TTreeNodeCompare)(AsTreeNode(getVal(0)), AsTreeNode(getVal(1)))

		// type TStringListSortCompare func(list IStringList, index1, index2 int32) int32
		case TStringListSortCompare:
			var result = getI32Ptr(3)
			*result = fn.(TStringListSortCompare)(AsStringList(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// type TListItemsCompare func(list IListControlItems, item1, item2 int32) int32
		case TListItemsCompare:
			var result = getI32Ptr(3)
			*result = fn.(TListItemsCompare)(AsListControlItems(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// type TLVCompare func(item1, item2 IListItem, optionalParam int64) int32 //stdcall
		case TLVCompare:
			var result = getI32Ptr(3)
			*result = fn.(TLVCompare)(AsListItem(getVal(0)), AsListItem(getVal(1)), uint32(getVal(2)))

		// type TListSortCompare func(item1, item2 uintptr) int32
		case TListSortCompare:
			var result = getI32Ptr(2)
			*result = fn.(TListSortCompare)(getVal(0), getVal(1))

		// type TCollectionSortCompare func(item1, item2 ICollectionItem) int32
		case TCollectionSortCompare:
			var result = getI32Ptr(2)
			*result = fn.(TCollectionSortCompare)(AsCollectionItem(getVal(0)), AsCollectionItem(getVal(1)))

		case TVTChangingEvent:
			fn.(TVTChangingEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), getBoolPtr(2))

		case TVTCheckChangingEvent:
			fn.(TVTCheckChangingEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), (*TCheckState)(getPtr(2)), getBoolPtr(3))

		case TVTChangeEvent:
			fn.(TVTChangeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)))

		case TVTStructureChangeEvent:
			fn.(TVTStructureChangeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TChangeReason(getVal(2)))

		case TVTEditCancelEvent:
			fn.(TVTEditCancelEvent)(AsBaseVirtualTree(getVal(0)), TColumnIndex(getVal(1)))

		case TVTEditChangingEvent:
			fn.(TVTEditChangingEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), getBoolPtr(3))

		case TVTEditChangeEvent:
			fn.(TVTEditChangeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)))

		case TVTFreeNodeEvent:
			fn.(TVTFreeNodeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)))

		case TVTFocusChangingEvent:
			fn.(TVTFocusChangingEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), AsVirtualNode(getVal(2)),
				TColumnIndex(getVal(3)), TColumnIndex(getVal(4)), getBoolPtr(5))

		case TVTFocusChangeEvent:
			fn.(TVTFocusChangeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)))

		case TVTAddToSelectionEvent:
			fn.(TVTAddToSelectionEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)))

		case TVTRemoveFromSelectionEvent:
			fn.(TVTRemoveFromSelectionEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)))

		case TVTGetImageEvent:
			fn.(TVTGetImageEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TVTImageKind(getVal(2)), TColumnIndex(getVal(3)),
				getBoolPtr(4), getI32Ptr(5))

		case TVTGetImageExEvent:
			imageList := AsCustomImageList(getPtrVal(6))
			fn.(TVTGetImageExEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TVTImageKind(getVal(2)), TColumnIndex(getVal(3)),
				getBoolPtr(4), getI32Ptr(5), &imageList)
			if imageList != nil {
				setPtrVal(6, imageList.Instance())
			}

		case TVTGetImageTextEvent:
			imageText := GoStr(getPtrVal(4))
			fn.(TVTGetImageTextEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TVTImageKind(getVal(2)), TColumnIndex(getVal(3)),
				&imageText)
			setPtrVal(4, PascalStr(imageText))

		case TVTHotNodeChangeEvent:
			fn.(TVTHotNodeChangeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), AsVirtualNode(getVal(2)))

		case TVTInitChildrenEvent:
			fn.(TVTInitChildrenEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), (*Cardinal)(getPtr(2)))

		case TVTInitNodeEvent:
			fn.(TVTInitNodeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), AsVirtualNode(getVal(2)), (*TVirtualNodeInitStates)(getPtr(3)))

		case TVTPopupEvent:
			popupMenu := AsPopupMenu(getVal(5))
			fn.(TVTPopupEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), *getPointPtr(3), getBoolPtr(4),
				&popupMenu)
			if popupMenu != nil {
				setPtrVal(5, popupMenu.Instance())
			}

		case TVTHelpContextEvent:
			fn.(TVTHelpContextEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), getI32Ptr(3))

		case TVTCreateEditorEvent:
			outEditLink := &IVTEditLink{instance: getPtr(3)}
			fn.(TVTCreateEditorEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), outEditLink)

		case TVTSaveTreeEvent:
			fn.(TVTSaveTreeEvent)(AsBaseVirtualTree(getVal(0)), AsStream(getVal(1)))

		case TVTSaveNodeEvent:
			fn.(TVTSaveNodeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), AsStream(getVal(2)))

		case TVTNodeExportEvent:
			fn.(TVTNodeExportEvent)(AsBaseVirtualTree(getVal(0)), TVTExportType(getVal(1)), AsVirtualNode(getVal(2)))

		case TVTColumnExportEvent:
			fn.(TVTColumnExportEvent)(AsBaseVirtualTree(getVal(0)), TVTExportType(getVal(1)), AsVirtualTreeColumn(getVal(2)))

		case TVTTreeExportEvent:
			fn.(TVTTreeExportEvent)(AsBaseVirtualTree(getVal(0)), TVTExportType(getVal(1)))

		case TVTDrawNodeEvent:
			paintInfoPtr := (*tVTPaintInfo)(getPtr(1))
			paintInfo := paintInfoPtr.Convert()
			fn.(TVTDrawNodeEvent)(AsBaseVirtualTree(getVal(0)), *paintInfo)

		case TVTGetCellContentMarginEvent:
			fn.(TVTGetCellContentMarginEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)),
				TVTCellContentMarginType(getVal(4)), getPointPtr(5))

		case TVTGetNodeWidthEvent:
			fn.(TVTGetNodeWidthEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)),
				getI32Ptr(4))

		case TVTHeaderClickEvent:
			fn.(TVTHeaderClickEvent)(AsVTHeader(getVal(0)), *(*TVTHeaderHitInfo)(getPtr(1)))

		case TVTHeaderMouseEvent:
			fn.(TVTHeaderMouseEvent)(AsVTHeader(getVal(0)), TMouseButton(getVal(1)), TShiftState(getVal(2)), int32(getVal(3)), int32(getVal(4)))

		case TVTHeaderMouseMoveEvent:
			fn.(TVTHeaderMouseMoveEvent)(AsVTHeader(getVal(0)), TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		case TVTBeforeHeaderHeightTrackingEvent:
			fn.(TVTBeforeHeaderHeightTrackingEvent)(AsVTHeader(getVal(0)), TShiftState(getVal(1)))

		case TVTAfterHeaderHeightTrackingEvent:
			fn.(TVTAfterHeaderHeightTrackingEvent)(AsVTHeader(getVal(0)))

		case TVTHeaderHeightTrackingEvent:
			fn.(TVTHeaderHeightTrackingEvent)(AsVTHeader(getVal(0)), getPointPtr(1), TShiftState(getVal(2)), getBoolPtr(3))

		case TVTHeaderHeightDblClickResizeEvent:
			fn.(TVTHeaderHeightDblClickResizeEvent)(AsVTHeader(getVal(0)), getPointPtr(1), TShiftState(getVal(2)), getBoolPtr(3))

		case TVTHeaderNotifyEvent:
			fn.(TVTHeaderNotifyEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)))

		case TVTHeaderDraggingEvent:
			fn.(TVTHeaderDraggingEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), getBoolPtr(2))

		case TVTHeaderDraggedEvent:
			fn.(TVTHeaderDraggedEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), int32(getVal(2)))

		case TVTHeaderDraggedOutEvent:
			fn.(TVTHeaderDraggedOutEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), *getPointPtr(2))

		case TVTHeaderPaintEvent:
			fn.(TVTHeaderPaintEvent)(AsVTHeader(getVal(0)), AsCanvas(getVal(1)), AsVirtualTreeColumn(getVal(2)), *getRectPtr(3), *getBoolPtr(4),
				*getBoolPtr(5), TVTDropMarkMode(getVal(6)))

		case TVTHeaderPaintQueryElementsEvent:
			paintInfoPtr := (*tHeaderPaintInfo)(getPtr(1))
			paintInfo := paintInfoPtr.Convert()
			fn.(TVTHeaderPaintQueryElementsEvent)(AsVTHeader(getVal(0)), paintInfo, (*THeaderPaintElements)(getPtr(2)))
			paintInfo.SetInstanceValue()

		case TVTAdvancedHeaderPaintEvent:
			paintInfoPtr := (*tHeaderPaintInfo)(getPtr(1))
			paintInfo := paintInfoPtr.Convert()
			fn.(TVTAdvancedHeaderPaintEvent)(AsVTHeader(getVal(0)), paintInfo, *(*THeaderPaintElements)(getPtr(2)))
			paintInfo.SetInstanceValue()

		case TVTBeforeAutoFitColumnsEvent:
			fn.(TVTBeforeAutoFitColumnsEvent)(AsVTHeader(getVal(0)), (*TSmartAutoFitType)(getPtr(1)))

		case TVTBeforeAutoFitColumnEvent:
			fn.(TVTBeforeAutoFitColumnEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), (*TSmartAutoFitType)(getPtr(2)), getBoolPtr(3))

		case TVTAfterAutoFitColumnEvent:
			fn.(TVTAfterAutoFitColumnEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)))

		case TVTAfterAutoFitColumnsEvent:
			fn.(TVTAfterAutoFitColumnsEvent)(AsVTHeader(getVal(0)))

		case TVTColumnClickEvent:
			fn.(TVTColumnClickEvent)(AsBaseVirtualTree(getVal(0)), TColumnIndex(getVal(1)), TShiftState(getVal(2)))

		case TVTColumnDblClickEvent:
			fn.(TVTColumnDblClickEvent)(AsBaseVirtualTree(getVal(0)), TColumnIndex(getVal(1)), TShiftState(getVal(2)))

		case TVTColumnWidthDblClickResizeEvent:
			fn.(TVTColumnWidthDblClickResizeEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), TShiftState(getVal(2)), *getPointPtr(3),
				getBoolPtr(4))

		case TVTBeforeColumnWidthTrackingEvent:
			fn.(TVTBeforeColumnWidthTrackingEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), TShiftState(getVal(2)))

		case TVTAfterColumnWidthTrackingEvent:
			fn.(TVTAfterColumnWidthTrackingEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)))

		case TVTColumnWidthTrackingEvent:
			fn.(TVTColumnWidthTrackingEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), TShiftState(getVal(2)), getPointPtr(3),
				*getPointPtr(4), getBoolPtr(5))

		case TVTGetHeaderCursorEvent:
			fn.(TVTGetHeaderCursorEvent)(AsVTHeader(getVal(0)), (*HCURSOR)(getPtr(1)))

		case TVTBeforeGetMaxColumnWidthEvent:
			fn.(TVTBeforeGetMaxColumnWidthEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), getBoolPtr(2))

		case TVTAfterGetMaxColumnWidthEvent:
			fn.(TVTAfterGetMaxColumnWidthEvent)(AsVTHeader(getVal(0)), TColumnIndex(getVal(1)), getI32Ptr(2))

		case TVTCanSplitterResizeColumnEvent:
			fn.(TVTCanSplitterResizeColumnEvent)(AsVTHeader(getVal(0)), *getPointPtr(1), TColumnIndex(getVal(2)), getBoolPtr(3))

		case TVTCanSplitterResizeHeaderEvent:
			fn.(TVTCanSplitterResizeHeaderEvent)(AsVTHeader(getVal(0)), *getPointPtr(1), getBoolPtr(2))

		case TVTNodeMovedEvent:
			fn.(TVTNodeMovedEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)))

		case TVTNodeMovingEvent:
			fn.(TVTNodeMovingEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), AsVirtualNode(getVal(2)), getBoolPtr(3))

		case TVTNodeCopiedEvent:
			fn.(TVTNodeCopiedEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)))

		case TVTNodeCopyingEvent:
			fn.(TVTNodeCopyingEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), AsVirtualNode(getVal(2)), getBoolPtr(3))

		case TVTNodeClickEvent:
			hitInfoPtr := (*tHitInfo)(getPtr(1))
			hitInfo := hitInfoPtr.Convert()
			fn.(TVTNodeClickEvent)(AsBaseVirtualTree(getVal(0)), *hitInfo)

		case TVTNodeHeightTrackingEvent:
			fn.(TVTNodeHeightTrackingEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), TShiftState(getVal(3)),
				getPointPtr(4), *getPointPtr(5), getBoolPtr(6))

		case TVTNodeHeightDblClickResizeEvent:
			fn.(TVTNodeHeightDblClickResizeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), TShiftState(getVal(3)),
				*getPointPtr(4), getBoolPtr(5))

		case TVTCanSplitterResizeNodeEvent:
			fn.(TVTCanSplitterResizeNodeEvent)(AsBaseVirtualTree(getVal(0)), *getPointPtr(1), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)), getBoolPtr(4))

		case TVTCreateDragManagerEvent:
			outDragManager := &IVTDragManager{instance: getPtr(1)}
			fn.(TVTCreateDragManagerEvent)(AsBaseVirtualTree(getVal(0)), outDragManager)

		case TVTCreateDataObjectEvent:
			outIDataObject := AsDataObject(getVal(1))
			fn.(TVTCreateDataObjectEvent)(AsBaseVirtualTree(getVal(0)), &outIDataObject)

		case TVTDragAllowedEvent:
			fn.(TVTDragAllowedEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), getBoolPtr(3))

		case TVTDragOverEvent:
			fn.(TVTDragOverEvent)(AsBaseVirtualTree(getVal(0)), AsObject(getVal(1)), TShiftState(getVal(2)), TDragState(getVal(3)), *getPointPtr(4),
				TDropMode(getVal(5)), (*LongWord)(getPtr(6)), getBoolPtr(7))

		case TVTDragDropEvent:
			formats := &TFormatArray{instance: getPtr(3), count: int(getVal(4))} // @formats
			fn.(TVTDragDropEvent)(AsBaseVirtualTree(getVal(0)), AsObject(getVal(1)), AsDataObject(getVal(2)), formats, TShiftState(getVal(5)),
				*getPointPtr(6), (*LongWord)(getPtr(7)), TDropMode(getVal(8)))

		case TVTBeforeItemEraseEvent:
			fn.(TVTBeforeItemEraseEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), *getRectPtr(3), (*TColor)(getPtr(4)),
				(*TItemEraseAction)(getPtr(5)))

		case TVTAfterItemEraseEvent:
			fn.(TVTAfterItemEraseEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), *getRectPtr(3))

		case TVTBeforeItemPaintEvent:
			fn.(TVTBeforeItemPaintEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), *getRectPtr(3), getBoolPtr(4))

		case TVTAfterItemPaintEvent:
			fn.(TVTAfterItemPaintEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), *getRectPtr(3))

		case TVTBeforeCellPaintEvent:
			fn.(TVTBeforeCellPaintEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)),
				TVTCellPaintMode(getVal(4)), *getRectPtr(5), getRectPtr(6))

		case TVTAfterCellPaintEvent:
			fn.(TVTAfterCellPaintEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)),
				*getRectPtr(4))

		case TVTPaintEvent:
			fn.(TVTPaintEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)))

		case TVTBackgroundPaintEvent:
			fn.(TVTBackgroundPaintEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), *getRectPtr(2), getBoolPtr(3))

		case TVTGetLineStyleEvent:
			fn.(TVTGetLineStyleEvent)(AsBaseVirtualTree(getVal(0)), (*Pointer)(getPtr(1)))

		case TVTMeasureItemEvent:
			fn.(TVTMeasureItemEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), getI32Ptr(3))

		case TVTCompareEvent:
			fn.(TVTCompareEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)), getI32Ptr(4))

		case TVTIncrementalSearchEvent:
			fn.(TVTIncrementalSearchEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), GoStr(getVal(2)), getI32Ptr(3))

		case TVTOperationEvent:
			fn.(TVTOperationEvent)(AsBaseVirtualTree(getVal(0)), TVTOperationKind(getVal(1)))

		case TVTHintKindEvent:
			fn.(TVTHintKindEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), (*TVTHintKind)(getPtr(3)))

		case TVTDrawHintEvent:
			fn.(TVTDrawHintEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), *getRectPtr(3), TColumnIndex(getVal(4)))

		case TVTGetHintSizeEvent:
			fn.(TVTGetHintSizeEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), getRectPtr(3))

		case TVTBeforeDrawLineImageEvent:
			fn.(TVTBeforeDrawLineImageEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), int32(getVal(2)), getI32Ptr(3))

		case TVTGetNodeDataSizeEvent:
			fn.(TVTGetNodeDataSizeEvent)(AsBaseVirtualTree(getVal(0)), getI32Ptr(1))

		case TVTKeyActionEvent:
			fn.(TVTKeyActionEvent)(AsBaseVirtualTree(getVal(0)), (*Word)(getPtr(1)), (*TShiftState)(getPtr(2)), getBoolPtr(3))

		case TVTScrollEvent:
			fn.(TVTScrollEvent)(AsBaseVirtualTree(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		case TVTUpdatingEvent:
			fn.(TVTUpdatingEvent)(AsBaseVirtualTree(getVal(0)), TVTUpdateState(getVal(1)))

		case TVTGetCursorEvent:
			fn.(TVTGetCursorEvent)(AsBaseVirtualTree(getVal(0)), (*TCursor)(getPtr(1)))

		case TVTStateChangeEvent:
			fn.(TVTStateChangeEvent)(AsBaseVirtualTree(getVal(0)), TVirtualTreeStates(getVal(1)), TVirtualTreeStates(getVal(2)))

		case TVTGetCellIsEmptyEvent:
			fn.(TVTGetCellIsEmptyEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), getBoolPtr(3))

		case TVTScrollBarShowEvent:
			fn.(TVTScrollBarShowEvent)(AsBaseVirtualTree(getVal(0)), int32(getVal(1)), *getBoolPtr(2))

		case TVTDrawTextEvent:
			fn.(TVTDrawTextEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)), GoStr(getVal(4)),
				*getRectPtr(5), *getBoolPtr(6))

		case TVSTGetTextEvent:
			cellText := GoStr(getPtrVal(4))
			fn.(TVSTGetTextEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), TVSTTextType(getVal(3)), &cellText)
			setPtrVal(4, PascalStr(cellText))

		case TVTPaintText:
			fn.(TVTPaintText)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)), TVSTTextType(getVal(3)))

		case TVSTGetHintEvent:
			hintText := GoStr(getPtrVal(4))
			fn.(TVSTGetHintEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), (*TVTTooltipLineBreakStyle)(getPtr(3)),
				&hintText)
			setPtrVal(4, PascalStr(hintText))

		case TVTMeasureTextEvent:
			fn.(TVTMeasureTextEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)),
				GoStr(getVal(4)), getI32Ptr(5))

		case TVSTNewTextEvent:
			fn.(TVSTNewTextEvent)(AsBaseVirtualTree(getVal(0)), AsVirtualNode(getVal(1)), TColumnIndex(getVal(2)), GoStr(getVal(3)))

		case TVSTShortenstringEvent:
			result := GoStr(getPtrVal(6))
			fn.(TVSTShortenstringEvent)(AsBaseVirtualTree(getVal(0)), AsCanvas(getVal(1)), AsVirtualNode(getVal(2)), TColumnIndex(getVal(3)),
				GoStr(getVal(4)), *getI32Ptr(5), &result, getBoolPtr(6))
			setPtrVal(6, PascalStr(result))

		default:
		}
	}
	return 0
}
