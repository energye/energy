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

// getUintptr 获取指针地址值
func getUintptr(ptr uintptr) uintptr {
	return *(*uintptr)(getPointer(ptr))
}

func getPointer(ptr uintptr) unsafe.Pointer {
	return unsafe.Pointer(ptr)
}

func getPointerPtr(ptr *uintptr) unsafe.Pointer {
	return unsafe.Pointer(ptr)
}

// 回调过程
func eventCallbackProc(f uintptr, args uintptr, _ int) uintptr {
	v := PtrToElementValue(f)
	if v != nil {
		// 获取值
		getVal := func(i int) uintptr {
			return getParamOf(i, args)
		}

		// 指针
		getPtr := func(i int) unsafe.Pointer {
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
			if extEventCallback[n](v, getVal) {
				return 0
			}
		}

		switch v.(type) {
		// func(sender IObject)
		case TNotifyEvent:
			v.(TNotifyEvent)(AsObject(getVal(0)))

		// func(sender IObject, button TUDBtnType)
		case TUDClickEvent:
			v.(TUDClickEvent)(AsObject(getVal(0)), TUDBtnType(getVal(1)))

		// func(sender IObject, item *TListItem, change int32)
		case TLVChangeEvent:
			v.(TLVChangeEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), TItemChange(getVal(2)))

		// func(sender IObject, action *TCloseAction) // Action *uintptr
		case TCloseEvent:
			v.(TCloseEvent)(AsObject(getVal(0)), (*TCloseAction)(getPtr(1)))

		// func(sender IObject, canClose *bool) //CanClose *uintptr
		case TCloseQueryEvent:
			v.(TCloseQueryEvent)(AsObject(getVal(0)), getBoolPtr(1))

		// func(sender IObject, source *TMenuItem, rebuild bool)
		case TMenuChangeEvent:
			v.(TMenuChangeEvent)(AsObject(getVal(0)), AsMenuItem(getVal(1)), GoBool(getVal(2)))

		// func(sender IObject, node *TreeNode)
		case TTVChangedEvent:
			v.(TTVChangedEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)))

		// func(sender IObject, link string, linkType TSysLinkType)
		case TSysLinkEvent:
			v.(TSysLinkEvent)(AsObject(getVal(0)), GoStr(getVal(1)), TSysLinkType(getVal(2)))

		// func(sender, e IObject)
		case TExceptionEvent:
			v.(TExceptionEvent)(AsObject(getVal(0)), AsException(getVal(1)))

		// func(sender IObject, key *Char, shift TShiftState)
		case TKeyEvent:
			v.(TKeyEvent)(AsObject(getVal(0)), (*Char)(getPtr(1)), TShiftState(getVal(2)))

		// func(sender IObject, key *Char)
		case TKeyPressEvent:
			v.(TKeyPressEvent)(AsObject(getVal(0)), (*Char)(getPtr(1)))

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)
		case TMouseEvent:
			v.(TMouseEvent)(AsObject(getVal(0)), TMouseButton(getVal(1)), TShiftState(getVal(2)), int32(getVal(3)), int32(getVal(4)))

		// func(sender IObject, shift TShiftState, x, y int32)
		case TMouseMoveEvent:
			v.(TMouseMoveEvent)(AsObject(getVal(0)), TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)
		case TMouseWheelEvent:
			v.(TMouseWheelEvent)(AsObject(getVal(0)), TShiftState(getVal(1)), int32(getVal(2)), int32(getVal(3)), int32(getVal(4)), getBoolPtr(5))

		// func(control IWinControl, index int32, aRect TRect, state TOwnerDrawState)
		case TDrawItemEvent:
			v.(TDrawItemEvent)(AsWinControl(getVal(0)), int32(getVal(1)), *getRectPtr(2), TOwnerDrawState(getVal(3)))

		// func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)
		case TMenuDrawItemEvent:
			v.(TMenuDrawItemEvent)(AsObject(getVal(0)), AsCanvas(getVal(1)), *getRectPtr(2), TOwnerDrawState(getVal(3)))

		// func(sender IObject, item *TListItem)
		case TLVNotifyEvent:
			v.(TLVNotifyEvent)(AsObject(getVal(0)), AsListItem(getVal(1)))

		// func(sender IObject, column *TListColumn)
		case TLVColumnClickEvent:
			v.(TLVColumnClickEvent)(AsObject(getVal(0)), AsListColumn(getVal(1)))

		// func(sender IObject, column *TListColumn, point TPoint)
		case TLVColumnRClickEvent:
			v.(TLVColumnRClickEvent)(AsObject(getVal(0)), AsListColumn(getVal(1)), TPoint{X: int32(getVal(2)), Y: int32(getVal(3))})

		// func(sender IObject, item *TListItem, selected bool)
		case TLVSelectItemEvent:
			v.(TLVSelectItemEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), GoBool(getVal(2)))

		//  func(sender IObject, item *TListItem)
		case TLVCheckedItemEvent:
			v.(TLVCheckedItemEvent)(AsObject(getVal(0)), AsListItem(getVal(1)))

		// func(sender IObject, tabIndex int32, imageIndex *int32)
		case TTabGetImageEvent:
			v.(TTabGetImageEvent)(AsObject(getVal(0)), int32(getVal(1)), getI32Ptr(2))

		// func(sender IObject, node *TTreeNode)
		case TTVExpandedEvent:
			v.(TTVExpandedEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)))

		// func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)
		case TLVCompareEvent:
			v.(TLVCompareEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), AsListItem(getVal(2)), int32(getVal(3)), getI32Ptr(4))

		// func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)
		case TTVCompareEvent:
			v.(TTVCompareEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), AsTreeNode(getVal(2)), int32(getVal(3)), getI32Ptr(4))

		// func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTVAdvancedCustomDrawEvent:
			v.(TTVAdvancedCustomDrawEvent)(AsTreeView(getVal(0)), *getRectPtr(1), TCustomDrawStage(getVal(2)), getBoolPtr(3))

		// func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)
		case TTVAdvancedCustomDrawItemEvent:
			v.(TTVAdvancedCustomDrawItemEvent)(AsTreeView(getVal(0)), AsTreeNode(getVal(1)), TCustomDrawState(getVal(2)), TCustomDrawStage(getVal(3)),
				getBoolPtr(4), getBoolPtr(5))

		// func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawEvent:
			v.(TLVAdvancedCustomDrawEvent)(AsListView(getVal(0)), *getRectPtr(1), TCustomDrawStage(getVal(2)), getBoolPtr(3))

		// func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawItemEvent:
			v.(TLVAdvancedCustomDrawItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), TCustomDrawState(getVal(2)), TCustomDrawStage(getVal(3)),
				getBoolPtr(4))

		// func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawSubItemEvent:
			v.(TLVAdvancedCustomDrawSubItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), int32(getVal(2)), TCustomDrawState(getVal(3)),
				TCustomDrawStage(getVal(4)), getBoolPtr(5))

		// func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTBAdvancedCustomDrawEvent:
			v.(TTBAdvancedCustomDrawEvent)(AsToolBar(getVal(0)), *getRectPtr(1), TCustomDrawStage(getVal(2)), getBoolPtr(3))

		// func(sender IObject, aFileNames []string)
		case TDropFilesEvent:
			nLen := int(getVal(2))
			tempArr := make([]string, nLen)
			p := getVal(1)
			for i := 0; i < nLen; i++ {
				tempArr[i] = DGetStringArrOf(p, i)
			}
			v.(TDropFilesEvent)(AsObject(getVal(0)), tempArr)

		// func(sender IObject, minWidth, minHeight, maxWidth, maxHeight *int32)
		case TConstrainedResizeEvent:
			v.(TConstrainedResizeEvent)(AsObject(getVal(0)), getI32Ptr(1), getI32Ptr(2), getI32Ptr(3), getI32Ptr(4))

		// func(command uint16, data THelpEventData, callHelp *bool) bool
		case THelpEvent:
			v.(THelpEvent)(uint16(getVal(0)), THelpEventData(getVal(1)), getBoolPtr(2), getBoolPtr(3))

		// func(msg *TWMKey, handled *bool)
		case TShortCutEvent:
			v.(TShortCutEvent)((*TWMKey)(getPtr(0)), getBoolPtr(1))

		// func(sender IObject, mousePos TPoint, handled *bool)
		case TContextPopupEvent:
			v.(TContextPopupEvent)(AsObject(getVal(0)), *getPointPtr(1), getBoolPtr(2))

		// func(sender, source IObject, x, y int32, state TDragState, accept *bool)
		case TDragOverEvent:
			v.(TDragOverEvent)(AsObject(getVal(0)), AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)), TDragState(getVal(4)), getBoolPtr(5))

		//func(sender, source IObject, x, y int32)
		case TDragDropEvent:
			v.(TDragDropEvent)(AsObject(getVal(0)), AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		//func(sender IObject, dragObject *TDragObject)
		case TStartDragEvent:
			obj := AsDragObject(getVal(1))
			v.(TStartDragEvent)(AsObject(getVal(0)), &obj)
			if obj != nil {
				*(*uintptr)(unsafe.Pointer(getVal(1))) = obj.Instance()
			}

		//func(sender, target IObject, x, y int32)
		case TEndDragEvent:
			v.(TEndDragEvent)(AsObject(getVal(0)), AsObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, source *TDragDockObject, x, y int32)
		case TDockDropEvent:
			v.(TDockDropEvent)(AsObject(getVal(0)), AsDragDockObject(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		//func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)
		case TDockOverEvent:
			v.(TDockOverEvent)(AsObject(getVal(0)), AsDragDockObject(getVal(1)), int32(getVal(2)), int32(getVal(3)),
				TDragState(getVal(4)), getBoolPtr(5))

		//func(sender IObject, client *TControl, newTarget *TControl, allow *bool)
		case TUnDockEvent:
			v.(TUnDockEvent)(AsObject(getVal(0)), AsControl(getVal(1)), AsControl(getVal(2)), getBoolPtr(3))

		//func(sender IObject, dragObject **TDragDockObject)
		case TStartDockEvent:
			obj := AsDragDockObject(getPtrVal(1))
			v.(TStartDockEvent)(AsObject(getVal(0)), &obj)
			if obj != nil {
				setPtrVal(1, obj.Instance())
			}

		//func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)
		case TGetSiteInfoEvent:
			v.(TGetSiteInfoEvent)(AsObject(getVal(0)), AsControl(getVal(1)), getRectPtr(2), *getPointPtr(3), getBoolPtr(4))

		//func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)
		case TMouseWheelUpDownEvent:
			v.(TMouseWheelUpDownEvent)(AsObject(getVal(0)), TShiftState(getVal(1)), *getPointPtr(2), getBoolPtr(3))

		// func(sender IObject, isColumn bool, sIndex, tIndex int32)
		case TGridOperationEvent:
			v.(TGridOperationEvent)(AsObject(getVal(0)), GoBool(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, aCol, aRow int32, aRect TRect, state TGridDrawState)
		case TDrawCellEvent:
			v.(TDrawCellEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), *getRectPtr(3), TGridDrawState(getVal(4)))

		// func(sender IObject, aCol, aRow int32)
		case TFixedCellClickEvent:
			v.(TFixedCellClickEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// func(sender IObject, aCol, aRow int32, value *string)
		case TGetEditEvent:
			str := GoStr(getPtrVal(3))
			v.(TGetEditEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), &str)
			setPtrVal(3, PascalStr(str))

		// func(sender IObject, aCol, aRow int32, canSelect *bool)
		case TSelectCellEvent:
			v.(TSelectCellEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), getBoolPtr(3))

		// func(sender IObject, aCol, aRow int32, value string)
		case TSetEditEvent:
			v.(TSetEditEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), GoStr(getVal(3)))

		// func(headerControl *THeaderControl, section *THeaderSection, aRect TRect, pressed bool)
		case TDrawSectionEvent:
			v.(TDrawSectionEvent)(AsHeaderControl(getVal(0)), AsHeaderSection(getVal(1)), *getRectPtr(2), getVal(3) != 0)

		// func(headerControl *THeaderControl, section *THeaderSection)
		case TSectionNotifyEvent:
			v.(TSectionNotifyEvent)(AsHeaderControl(getVal(0)), AsHeaderSection(getVal(1)))

		// func(headerControl *THeaderControl, section *THeaderSection, width int32, state TSectionTrackState)
		case TSectionTrackEvent:
			v.(TSectionTrackEvent)(AsHeaderControl(getVal(0)), AsHeaderSection(getVal(1)), int32(getVal(2)), TSectionTrackState(getVal(3)))

		// func(sender IObject, fromSection, toSection *THeaderSection, allowDrag *bool)
		case TSectionDragEvent:
			v.(TSectionDragEvent)(AsObject(getVal(0)), AsHeaderSection(getVal(1)), AsHeaderSection(getVal(2)), getBoolPtr(3))

		// func(headerControl *THeaderControl, section *THeaderSection)
		case TCustomSectionNotifyEvent:
			v.(TCustomSectionNotifyEvent)(AsHeaderControl(getVal(0)), AsHeaderSection(getVal(1)))

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32, hitTest int32, mouseActivate *TMouseActivate)
		case TMouseActivateEvent:
			v.(TMouseActivateEvent)(AsObject(getVal(0)), TMouseButton(getVal(1)), TShiftState(getVal(2)), int32(getVal(3)), int32(getVal(4)),
				int32(getVal(5)), (*TMouseActivate)(getPtr(6)))

		// func(control *TWinControl, index int32, data *string)
		case TLBGetDataEvent:
			str := GoStr(getPtrVal(2))
			v.(TLBGetDataEvent)(AsWinControl(getVal(0)), int32(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(control *TWinControl, index int32, dataObject IObject)
		case TLBGetDataObjectEvent:
			v.(TLBGetDataObjectEvent)(AsWinControl(getVal(0)), int32(getVal(1)), AsObject(getVal(2))) // 这个参数要改，先这样

		// func(control *TWinControl, findString string) int32
		case TLBFindDataEvent:
			result := v.(TLBFindDataEvent)(AsWinControl(getVal(0)), GoStr(getVal(1)))
			*getI32Ptr(2) = result

		// func(control *TWinControl, index int32, height *int32)
		case TMeasureItemEvent:
			v.(TMeasureItemEvent)(AsWinControl(getVal(0)), int32(getVal(1)), getI32Ptr(2))

		// func(sender IObject, item *TListItem, change TItemChange, allowChange *bool)
		case TLVChangingEvent:
			v.(TLVChangingEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), TItemChange(getVal(2)), getBoolPtr(3))

		// func(sender IObject, item *TListItem)
		case TLVDataEvent:
			v.(TLVDataEvent)(AsObject(getVal(0)), AsListItem(getVal(1)))

		// func(sender IObject, item *TListItem)
		//case TLVDeletedEvent:
		//	v.(TLVDeletedEvent)(AsObject(getVal(0)), AsListItem(getVal(1)))

		// func(sender IObject, find TItemFind, findString string, findPosition TPoint, findData TCustomData, startIndex int32,
		//	direction TSearchDirection, warp bool, index *int32)
		case TLVDataFindEvent:
			v.(TLVDataFindEvent)(AsObject(getVal(0)), TItemFind(getVal(1)), GoStr(getVal(2)), *getPointPtr(3), TCustomData(getVal(4)),
				int32(getVal(5)), TSearchDirection(getVal(6)), GoBool(getVal(7)), getI32Ptr(8))

		// func(sender IObject, item *TListItem, allowEdit *bool)
		case TLVEditingEvent:
			v.(TLVEditingEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), getBoolPtr(2))

		// func(sender IObject, item *TListItem, s *string)
		case TLVEditedEvent:
			str := GoStr(getPtrVal(2))
			v.(TLVEditedEvent)(AsObject(getVal(0)), AsListItem(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(sender IObject, aCanvas *TCanvas, width, height *int32)
		case TMenuMeasureItemEvent:
			v.(TMenuMeasureItemEvent)(AsObject(getVal(0)), AsCanvas(getVal(1)), getI32Ptr(2), getI32Ptr(3))

		//type func(sender IObject, allowChange *bool)
		case TTabChangingEvent:
			v.(TTabChangingEvent)(AsObject(getVal(0)), getBoolPtr(1))

		// func(sender IObject, node *TTreeNode, allowChange *bool)
		case TTVChangingEvent:
			v.(TTVChangingEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), getBoolPtr(2))

		// func(sender IObject, node *TTreeNode, allowCollapse *bool)
		case TTVCollapsingEvent:
			v.(TTVCollapsingEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), getBoolPtr(2))

		// func(sender IObject, node *TTreeNode, s *string)
		case TTVEditedEvent:
			str := GoStr(getPtrVal(2))
			v.(TTVEditedEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(sender IObject, node *TTreeNode, allowEdit *bool)
		case TTVEditingEvent:
			v.(TTVEditingEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), getBoolPtr(2))

		// func(sender IObject, node *TTreeNode, allowExpansion *bool)
		case TTVExpandingEvent:
			v.(TTVExpandingEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), getBoolPtr(2))

		// func(sender IObject, node *TTreeNode, hint *string)
		case TTVHintEvent:
			str := GoStr(getPtrVal(2))
			v.(TTVHintEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(sender IObject, allowChange *bool)
		case TUDChangingEvent:
			v.(TUDChangingEvent)(AsObject(getVal(0)), getBoolPtr(1))

		// func(sender IObject, winErrorCode uint32, errorDescription string, handled *bool)
		case TCreatingListErrorEvent:
			v.(TCreatingListErrorEvent)(AsObject(getVal(0)), uint32(getVal(1)), GoStr(getVal(2)), getBoolPtr(3))

		// func(sender *TListView, aRect TRect, defaultDraw *bool)
		case TLVCustomDrawEvent:
			v.(TLVCustomDrawEvent)(AsListView(getVal(0)), *getRectPtr(1), getBoolPtr(2))

		// func(sender *TListView, item *TListItem, state TCustomDrawStage, defaultDraw *bool)
		case TLVCustomDrawItemEvent:
			v.(TLVCustomDrawItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), TCustomDrawState(getVal(2)), getBoolPtr(3))

		// func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawStage, defaultDraw *bool)
		case TLVCustomDrawSubItemEvent:
			v.(TLVCustomDrawSubItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), int32(getVal(2)), TCustomDrawState(getVal(3)), getBoolPtr(4))

		// func(sender *TListView, item *TListItem, rect TRect, state TOwnerDrawState)
		case TLVDrawItemEvent:
			v.(TLVDrawItemEvent)(AsListView(getVal(0)), AsListItem(getVal(1)), *getRectPtr(2), TOwnerDrawState(getVal(3)))

		// func(sender IObject, startIndex, endIndex int32)
		case TLVDataHintEvent:
			v.(TLVDataHintEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// func(sender *TTreeView, aRect TRect, defaultDraw *bool)
		case TTVCustomDrawEvent:
			v.(TTVCustomDrawEvent)(AsTreeView(getVal(0)), *getRectPtr(1), getBoolPtr(2))

		// func(sender *TTreeView, node *TTreeNode, state TCustomDrawStage, defaultDraw *bool)
		case TTVCustomDrawItemEvent:
			v.(TTVCustomDrawItemEvent)(AsTreeView(getVal(0)), AsTreeNode(getVal(1)), TCustomDrawState(getVal(2)), getBoolPtr(3))

		// func(sender IObject, text string)
		case TWebTitleChangeEvent:
			v.(TWebTitleChangeEvent)(AsObject(getVal(0)), GoStr(getVal(1)))

		// func(sender IObject, funcName, args string, retVal *string)
		case TWebJSExternalEvent:
			str := GoStr(getPtrVal(3))
			v.(TWebJSExternalEvent)(AsObject(getVal(0)), GoStr(getVal(1)), GoStr(getVal(2)), &str)
			setPtrVal(3, PascalStr(str))

		// func(sender IObject, modalResult TModalResult, canClose *bool)
		case TTaskDlgClickEvent:
			v.(TTaskDlgClickEvent)(AsObject(getVal(0)), TModalResult(getVal(1)), getBoolPtr(2))

		// func(sender IObject, tickCount uint32, reset *bool)
		case TTaskDlgTimerEvent:
			v.(TTaskDlgTimerEvent)(AsObject(getVal(0)), uint32(getVal(1)), getBoolPtr(2))

		// func(sender *TWinControl, control *TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *TRect, alignInfo TAlignInfo)
		case TAlignPositionEvent:
			v.(TAlignPositionEvent)(AsWinControl(getVal(0)), AsControl(getVal(1)), getI32Ptr(2), getI32Ptr(3), getI32Ptr(4),
				getI32Ptr(5), getRectPtr(6), *(*TAlignInfo)(getPtr(7)))

		// func(sender IObject, index int32)
		case TCheckGroupClicked:
			v.(TCheckGroupClicked)(AsObject(getVal(0)), int32(getVal(1)))

		// func(sender IObject, aCol, aRow int32)
		case TOnSelectEvent:
			v.(TOnSelectEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// func(sender IObject, aCol, aRow int32, aState TCheckBoxState)
		case TToggledCheckboxEvent:
			v.(TToggledCheckboxEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCheckBoxState(getVal(3)))

		// func(sender IObject, ACol, ARow, BCol, BRow int32, result *int32)
		case TOnCompareCells:
			v.(TOnCompareCells)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), int32(getVal(3)), int32(getVal(4)), getI32Ptr(5))

		// func(sender IObject, ACol, ARow int32, hintText *string)
		case TGetCellHintEvent:
			str := GoStr(getPtrVal(3))
			v.(TGetCellHintEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), &str)
			setPtrVal(3, PascalStr(str))

		// func(sender IObject, ACol, ARow int32, value *TCheckBoxState)
		case TGetCheckboxStateEvent:
			v.(TGetCheckboxStateEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), (*TCheckBoxState)(getPtr(3)))

		// func(sender IObject, ACol, ARow int32, Value TCheckBoxState)
		case TSetCheckboxStateEvent:
			v.(TSetCheckboxStateEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCheckBoxState(getVal(3)))

		// func(sender IObject, isColumn bool, index int32)
		case THdrEvent:
			v.(THdrEvent)(AsObject(getVal(0)), GoBool(getVal(1)), int32(getVal(2)))

		// func(sender IObject, isColumn bool, aIndex, aSize int32)
		case THeaderSizingEvent:
			v.(THeaderSizingEvent)(AsObject(getVal(0)), GoBool(getVal(1)), int32(getVal(2)), int32(getVal(3)))

		// func(sender IObject, aCol, aRow int32, editor **TWinControl)
		case TSelectEditorEvent:
			obj := AsWinControl(getPtrVal(3))
			v.(TSelectEditorEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), &obj)
			if obj != nil {
				setPtrVal(3, obj.Instance())
			}

		// func(sender IObject, aCol, aRow int32, CheckedState TCheckBoxState, aBitmap **TBitmap)
		case TUserCheckBoxBitmapEvent:
			obj := AsBitmap(getPtrVal(4))
			v.(TUserCheckBoxBitmapEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCheckBoxState(getVal(1)), &obj)
			if obj != nil {
				setPtrVal(4, obj.Instance())
			}

		// func(sender IObject, aCol, aRow int32, oldValue string, newValue *string)
		case TValidateEntryEvent:
			str := GoStr(getPtrVal(4))
			v.(TValidateEntryEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), GoStr(getVal(3)), &str)
			setPtrVal(4, PascalStr(str))

		// func(sender IObject, aCol, aRow int32, aState TGridDrawState)
		case TOnPrepareCanvasEvent:
			v.(TOnPrepareCanvasEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TGridDrawState(getVal(3)))

		// func(sender IObject, value *string)
		case TAcceptFileNameEvent:
			str := GoStr(getPtrVal(1))
			v.(TAcceptFileNameEvent)(AsObject(getVal(0)), &str)
			setPtrVal(1, PascalStr(str))

		// func(sender IObject, index int32)
		case TCheckItemChange:
			v.(TCheckItemChange)(AsObject(getVal(0)), int32(getVal(1)))

		// func(sender IObject, utf8key *TUTF8Char)
		case TUTF8KeyPressEvent:
			v.(TUTF8KeyPressEvent)(AsObject(getVal(0)), (*TUTF8Char)(getPtr(1)))

		// type  func(sender IObject, aCanvas *TCanvas, aRect TRect)
		case TImagePaintBackgroundEvent:
			v.(TImagePaintBackgroundEvent)(AsObject(getVal(0)), AsCanvas(getVal(1)), *getRectPtr(2))

		// func(sender IObject, Action IBasicAction, handled *bool)
		case TActionEvent:
			v.(TActionEvent)(AsObject(getVal(0)), AsBasicAction(getVal(1)), getBoolPtr(2))

		// func(handle *HWND)
		case TGetHandleEvent:
			v.(TGetHandleEvent)((*HWND)(getPtr(1)))

		// func(sender IObject, result int32)
		case TModalDialogFinished:
			v.(TModalDialogFinished)(AsObject(getVal(0)), int32(getVal(1)))

		// func(sender IObject, control IControl, caption *string)
		case TGetDockCaptionEvent:
			str := GoStr(getPtrVal(2))
			v.(TGetDockCaptionEvent)(AsObject(getVal(0)), AsControl(getVal(1)), &str)
			setPtrVal(2, PascalStr(str))

		// func(sender IObject, aCol, aRow int32, aRect TRect, aState TGridDrawState)
		case TOnDrawCell:
			v.(TOnDrawCell)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), *getRectPtr(3), TGridDrawState(getVal(4)))

		// func(sender IObject, aCol, aRow int32, checkedState TCheckBoxState, imageList ICustomImageList, imageIndex TImageIndex)
		case TUserCheckBoxImageEvent:
			v.(TUserCheckBoxImageEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCheckBoxState(getVal(3)), AsCustomImageList(getVal(4)),
				TImageIndex(getVal(5)))

		// func(sender IWinControl, control1, control2 IControl) bool
		case TAlignInsertBeforeEvent:
			v.(TAlignInsertBeforeEvent)(AsWinControl(getVal(0)), AsControl(getVal(1)), AsControl(getVal(2)))

		// func(sender IObject, AllowChange *bool, newValue int16, direction TUpDownDirection)
		case TUDChangingEventEx:
			v.(TUDChangingEventEx)(AsObject(getVal(0)), getBoolPtr(1), int16(getVal(2)), TUpDownDirection(getVal(3)))

		// func(sender IStatusBar, panelClass *TStatusPanelClass)
		case TSBCreatePanelClassEvent:
			v.(TSBCreatePanelClassEvent)(AsStatusBar(getVal(0)), (*TStatusPanelClass)(getPtr(1)))

		// func(statusBar IStatusBar, panel IStatusPanel, rect TRect)
		case TDrawPanelEvent:
			v.(TDrawPanelEvent)(AsStatusBar(getVal(0)), AsStatusPanel(getVal(1)), *getRectPtr(2))

		// func(sender IObject, success bool)
		case TDialogResultEvent:
			v.(TDialogResultEvent)(AsObject(getVal(0)), *getBoolPtr(1))

		// func(sender ICustomColorBox, items IStrings)
		case TGetColorsEvent:
			v.(TGetColorsEvent)(AsCustomColorBox(getVal(0)), AsStrings(getVal(1)))

		// func(hintStr *string, canShow bool)
		case THintEvent:
			str := GoStr(getPtrVal(0))
			v.(THintEvent)(&str, *getBoolPtr(1))
			setPtrVal(0, PascalStr(str))

		// func(sender IObject, aCol, aRow int32, canSelect *bool)
		case TOnSelectCellEvent:
			v.(TOnSelectCellEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), getBoolPtr(3))

		// func(sender IObject, user bool)
		case TSelectionChangeEvent:
			v.(TSelectionChangeEvent)(AsObject(getVal(0)), *getBoolPtr(1))

		// func(color1, color2 TFPColor) TFPColor
		case TFPCanvasCombineColors:
			v.(TFPCanvasCombineColors)(*(*TFPColor)(getPtr(0)), *(*TFPColor)(getPtr(1)), (*TFPColor)(getPtr(2)))

		// func(sender IObject, stage TFPImgProgressStage, percentDone byte, redrawNow bool, rect TRect, msg string, continue_ *bool)
		case TFPImgProgressEvent:
			v.(TFPImgProgressEvent)(AsObject(getVal(0)), TFPImgProgressStage(getVal(1)), byte(getVal(2)), *getBoolPtr(3), *getRectPtr(4),
				GoStr(getVal(5)), getBoolPtr(6))

		// func(sender ICustomColorListBox, items IStrings)
		case TLBGetColorsEvent:
			v.(TLBGetColorsEvent)(AsCustomColorListBox(getVal(0)), AsStrings(getVal(1)))

		// func(sender ICustomHeaderControl, sectionClass *THeaderSectionClass)
		case TCustomHCCreateSectionClassEvent:
			v.(TCustomHCCreateSectionClassEvent)(AsCustomHeaderControl(getVal(0)), (*THeaderSectionClass)(getPtr(1)))

		// func(headerControl ICustomHeaderControl, section IHeaderSection, width int32, state TSectionTrackState)
		case TCustomSectionTrackEvent:
			v.(TCustomSectionTrackEvent)(AsCustomHeaderControl(getVal(0)), AsHeaderSection(getVal(1)), int32(getVal(2)), TSectionTrackState(getVal(3)))

		// func(sender IObject, newOffset *int32, accept *bool)
		case TCanOffsetEvent:
			v.(TCanOffsetEvent)(AsObject(getVal(0)), getI32Ptr(1), getBoolPtr(2))

		// func(sender IObject, newSize *int32, accept *bool)
		case TCanResizeEvent:
			v.(TCanResizeEvent)(AsObject(getVal(0)), getI32Ptr(1), getBoolPtr(2))

		// func(sender IObject, keyName string, values IStrings)
		case TGetPickListEvent:
			v.(TGetPickListEvent)(AsObject(getVal(0)), GoStr(getVal(1)), AsStrings(getVal(2)))

		// func(sender IObject, aCol, aRow int32, keyName, keyValue string)
		case TOnValidateEvent:
			v.(TOnValidateEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), GoStr(getVal(3)), GoStr(getVal(4)))

		// func(requestedFormatID TClipboardFormat, data IStream)
		case TClipboardRequestEvent:
			v.(TClipboardRequestEvent)(getVal(0), AsStream(getVal(1)))

		// func(sender IToolButton, state int32)
		case TToolBarOnPaintButton:
			v.(TToolBarOnPaintButton)(AsToolButton(getVal(0)), int32(getVal(1)))

		// func(sender IObject, aCol, aRow int32, processType TCellProcessType, aValue *string)
		case TCellProcessEvent:
			str := GoStr(getPtrVal(4))
			v.(TCellProcessEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TCellProcessType(getVal(3)), &str)
			setPtrVal(4, PascalStr(str))

		// func(sender IObject, scrollCode TScrollCode, scrollPos *int32)
		case TScrollEvent:
			v.(TScrollEvent)(AsObject(getVal(0)), TScrollCode(getVal(1)), getI32Ptr(2))

		// func(aList IListControlItems, aItem1, aItem2 IListControlItem) int32
		case TListCompareEvent:
			v.(TListCompareEvent)(AsListControlItems(getVal(0)), AsListControlItem(getVal(1)), AsListControlItem(getVal(2)), getI32Ptr(3))

		// func(sender ICustomImageList, aImageWidth, aPPI int32, aResultWidth *int32)
		case TCustomImageListGetWidthForPPI:
			v.(TCustomImageListGetWidthForPPI)(AsCustomImageList(getVal(0)), int32(getVal(1)), int32(getVal(2)), getI32Ptr(3))

		// func(sender ICustomListView, itemClass *TListItemClass)
		case TLVCreateItemClassEvent:
			v.(TLVCreateItemClassEvent)(AsCustomListView(getVal(0)), (*TListItemClass)(getPtr(1)))

		// func(sender IObject, startIndex, endIndex int32, oldState, newState TListItemStates)
		case TLVDataStateChangeEvent:
			v.(TLVDataStateChangeEvent)(AsObject(getVal(0)), int32(getVal(1)), int32(getVal(2)), TListItemStates(getVal(3)), TListItemStates(getVal(4)))

		// func(sender ICustomTreeView, nodeClass *TTreeNodeClass)
		case TTVCreateNodeClassEvent:
			v.(TTVCreateNodeClassEvent)(AsCustomTreeView(getVal(0)), (*TTreeNodeClass)(getPtr(1)))

		// func(sender ICustomTreeView, aTreeNode ITreeNode)
		case TTVCustomCreateNodeEvent:
			v.(TTVCustomCreateNodeEvent)(AsCustomTreeView(getVal(0)), AsTreeNode(getVal(1)))

		// func(sender ICustomTreeView, ARect TRect, aCollapsed bool)
		case TTVCustomDrawArrowEvent:
			v.(TTVCustomDrawArrowEvent)(AsCustomTreeView(getVal(0)), *getRectPtr(1), *getBoolPtr(2))

		// func(sender IObject, node ITreeNode, cancel bool)
		case TTVEditingEndEvent:
			v.(TTVEditingEndEvent)(AsObject(getVal(0)), AsTreeNode(getVal(1)), *getBoolPtr(2))

		// func(sender ICustomTreeView, aNode ITreeNode) bool
		case TTVHasChildrenEvent:
			v.(TTVHasChildrenEvent)(AsCustomTreeView(getVal(0)), AsTreeNode(getVal(0)))

		// func(sender IObject, node ITreeNode, changeReason TTreeNodeChangeReason)
		case TTVNodeChangedEvent:
			v.(TTVNodeChangedEvent)(AsObject(getVal(0)), AsTreeNode(getVal(0)), TTreeNodeChangeReason(getVal(0)))

		// func(hintStr *string, canShow *bool, hintInfo *THintInfo) // var hintInfo
		case TShowHintEvent:
			str := GoStr(getPtrVal(0))
			hitInfoPtr := (*PHintInfo)(getPtr(2))
			hitInfo := hitInfoPtr.THintInfo()
			v.(TShowHintEvent)(&str, getBoolPtr(1), hitInfo)
			setPtrVal(0, PascalStr(str))
			setPtrVal(2, uintptr(unsafe.Pointer(hitInfo.PHintInfo())))

		// func(sender IObject, hintInfo *HintInfo)
		case TControlShowHintEvent:
			hitInfo := (*PHintInfo)(getPtr(1))
			v.(TControlShowHintEvent)(AsObject(getVal(0)), hitInfo.THintInfo())

		// func(sender ICustomImageList, aWidth int32, aReferenceHandle TLCLHandle)
		case TDestroyResolutionHandleEvent:
			v.(TDestroyResolutionHandleEvent)(AsCustomImageList(getVal(0)), int32(getVal(1)), TLCLHandle(getVal(2)))

		// type TTreeNodeCompare func(node1, node2 ITreeNode) int32
		case TTreeNodeCompare:
			var result = getI32Ptr(2)
			*result = v.(TTreeNodeCompare)(AsTreeNode(getVal(0)), AsTreeNode(getVal(1)))

		// type TStringListSortCompare func(list IStringList, index1, index2 int32) int32
		case TStringListSortCompare:
			var result = getI32Ptr(3)
			*result = v.(TStringListSortCompare)(AsStringList(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// type TListItemsCompare func(list IListControlItems, item1, item2 int32) int32
		case TListItemsCompare:
			var result = getI32Ptr(3)
			*result = v.(TListItemsCompare)(AsListControlItems(getVal(0)), int32(getVal(1)), int32(getVal(2)))

		// type TLVCompare func(item1, item2 IListItem, optionalParam int64) int32 //stdcall
		case TLVCompare:
			var result = getI32Ptr(3)
			*result = v.(TLVCompare)(AsListItem(getVal(0)), AsListItem(getVal(1)), uint32(getVal(2)))

		// type TListSortCompare func(item1, item2 uintptr) int32
		case TListSortCompare:
			var result = getI32Ptr(2)
			*result = v.(TListSortCompare)(getVal(0), getVal(1))

		// type TCollectionSortCompare func(item1, item2 ICollectionItem) int32
		case TCollectionSortCompare:
			var result = getI32Ptr(2)
			*result = v.(TCollectionSortCompare)(AsCollectionItem(getVal(0)), AsCollectionItem(getVal(1)))
		default:
		}
	}
	return 0
}
