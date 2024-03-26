//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

// AsException Convert a pointer object to an existing class object
func AsException(obj interface{}) IException {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	exception := new(Exception)
	SetObjectInstance(exception, instance)
	return exception
}

// AsATGauge Convert a pointer object to an existing class object
func AsATGauge(obj interface{}) IATGauge {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	aTGauge := new(TATGauge)
	SetObjectInstance(aTGauge, instance)
	return aTGauge
}

// AsAVLTree Convert a pointer object to an existing class object
func AsAVLTree(obj interface{}) IAVLTree {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	aVLTree := new(TAVLTree)
	SetObjectInstance(aVLTree, instance)
	return aVLTree
}

// AsAVLTreeNode Convert a pointer object to an existing class object
func AsAVLTreeNode(obj interface{}) IAVLTreeNode {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	aVLTreeNode := new(TAVLTreeNode)
	SetObjectInstance(aVLTreeNode, instance)
	return aVLTreeNode
}

// AsAVLTreeNodeEnumerator Convert a pointer object to an existing class object
func AsAVLTreeNodeEnumerator(obj interface{}) IAVLTreeNodeEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	aVLTreeNodeEnumerator := new(TAVLTreeNodeEnumerator)
	SetObjectInstance(aVLTreeNodeEnumerator, instance)
	return aVLTreeNodeEnumerator
}

// AsAction Convert a pointer object to an existing class object
func AsAction(obj interface{}) IAction {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	action := new(TAction)
	SetObjectInstance(action, instance)
	return action
}

// AsActionList Convert a pointer object to an existing class object
func AsActionList(obj interface{}) IActionList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	actionList := new(TActionList)
	SetObjectInstance(actionList, instance)
	return actionList
}

// AsActionListEnumerator Convert a pointer object to an existing class object
func AsActionListEnumerator(obj interface{}) IActionListEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	actionListEnumerator := new(TActionListEnumerator)
	SetObjectInstance(actionListEnumerator, instance)
	return actionListEnumerator
}

// AsAnchorSide Convert a pointer object to an existing class object
func AsAnchorSide(obj interface{}) IAnchorSide {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	anchorSide := new(TAnchorSide)
	SetObjectInstance(anchorSide, instance)
	return anchorSide
}

// AsApplication Convert a pointer object to an existing class object
func AsApplication(obj interface{}) IApplication {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	application := new(TApplication)
	SetObjectInstance(application, instance)
	return application
}

// AsBasicAction Convert a pointer object to an existing class object
func AsBasicAction(obj interface{}) IBasicAction {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	basicAction := new(TBasicAction)
	SetObjectInstance(basicAction, instance)
	return basicAction
}

// AsBasicActionLink Convert a pointer object to an existing class object
func AsBasicActionLink(obj interface{}) IBasicActionLink {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	basicActionLink := new(TBasicActionLink)
	SetObjectInstance(basicActionLink, instance)
	return basicActionLink
}

// AsBevel Convert a pointer object to an existing class object
func AsBevel(obj interface{}) IBevel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	bevel := new(TBevel)
	SetObjectInstance(bevel, instance)
	return bevel
}

// AsBitBtn Convert a pointer object to an existing class object
func AsBitBtn(obj interface{}) IBitBtn {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	bitBtn := new(TBitBtn)
	SetObjectInstance(bitBtn, instance)
	return bitBtn
}

// AsBitmap Convert a pointer object to an existing class object
func AsBitmap(obj interface{}) IBitmap {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	bitmap := new(TBitmap)
	SetObjectInstance(bitmap, instance)
	return bitmap
}

// AsBoundLabel Convert a pointer object to an existing class object
func AsBoundLabel(obj interface{}) IBoundLabel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	boundLabel := new(TBoundLabel)
	SetObjectInstance(boundLabel, instance)
	return boundLabel
}

// AsBrush Convert a pointer object to an existing class object
func AsBrush(obj interface{}) IBrush {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	brush := new(TBrush)
	SetObjectInstance(brush, instance)
	return brush
}

// AsButton Convert a pointer object to an existing class object
func AsButton(obj interface{}) IButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	button := new(TButton)
	SetObjectInstance(button, instance)
	return button
}

// AsButtonControl Convert a pointer object to an existing class object
func AsButtonControl(obj interface{}) IButtonControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	buttonControl := new(TButtonControl)
	SetObjectInstance(buttonControl, instance)
	return buttonControl
}

// AsCalendar Convert a pointer object to an existing class object
func AsCalendar(obj interface{}) ICalendar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	calendar := new(TCalendar)
	SetObjectInstance(calendar, instance)
	return calendar
}

// AsCanvas Convert a pointer object to an existing class object
func AsCanvas(obj interface{}) ICanvas {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	canvas := new(TCanvas)
	SetObjectInstance(canvas, instance)
	return canvas
}

// AsChangeLink Convert a pointer object to an existing class object
func AsChangeLink(obj interface{}) IChangeLink {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	changeLink := new(TChangeLink)
	SetObjectInstance(changeLink, instance)
	return changeLink
}

// AsCheckBox Convert a pointer object to an existing class object
func AsCheckBox(obj interface{}) ICheckBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	checkBox := new(TCheckBox)
	SetObjectInstance(checkBox, instance)
	return checkBox
}

// AsCheckComboBox Convert a pointer object to an existing class object
func AsCheckComboBox(obj interface{}) ICheckComboBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	checkComboBox := new(TCheckComboBox)
	SetObjectInstance(checkComboBox, instance)
	return checkComboBox
}

// AsCheckGroup Convert a pointer object to an existing class object
func AsCheckGroup(obj interface{}) ICheckGroup {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	checkGroup := new(TCheckGroup)
	SetObjectInstance(checkGroup, instance)
	return checkGroup
}

// AsCheckListBox Convert a pointer object to an existing class object
func AsCheckListBox(obj interface{}) ICheckListBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	checkListBox := new(TCheckListBox)
	SetObjectInstance(checkListBox, instance)
	return checkListBox
}

// AsClipboard Convert a pointer object to an existing class object
func AsClipboard(obj interface{}) IClipboard {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	clipboard := new(TClipboard)
	SetObjectInstance(clipboard, instance)
	return clipboard
}

// AsCollection Convert a pointer object to an existing class object
func AsCollection(obj interface{}) ICollection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	collection := new(TCollection)
	SetObjectInstance(collection, instance)
	return collection
}

// AsCollectionEnumerator Convert a pointer object to an existing class object
func AsCollectionEnumerator(obj interface{}) ICollectionEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	collectionEnumerator := new(TCollectionEnumerator)
	SetObjectInstance(collectionEnumerator, instance)
	return collectionEnumerator
}

// AsCollectionItem Convert a pointer object to an existing class object
func AsCollectionItem(obj interface{}) ICollectionItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	collectionItem := new(TCollectionItem)
	SetObjectInstance(collectionItem, instance)
	return collectionItem
}

// AsColorBox Convert a pointer object to an existing class object
func AsColorBox(obj interface{}) IColorBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	colorBox := new(TColorBox)
	SetObjectInstance(colorBox, instance)
	return colorBox
}

// AsColorButton Convert a pointer object to an existing class object
func AsColorButton(obj interface{}) IColorButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	colorButton := new(TColorButton)
	SetObjectInstance(colorButton, instance)
	return colorButton
}

// AsColorDialog Convert a pointer object to an existing class object
func AsColorDialog(obj interface{}) IColorDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	colorDialog := new(TColorDialog)
	SetObjectInstance(colorDialog, instance)
	return colorDialog
}

// AsColorListBox Convert a pointer object to an existing class object
func AsColorListBox(obj interface{}) IColorListBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	colorListBox := new(TColorListBox)
	SetObjectInstance(colorListBox, instance)
	return colorListBox
}

// AsComboBox Convert a pointer object to an existing class object
func AsComboBox(obj interface{}) IComboBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	comboBox := new(TComboBox)
	SetObjectInstance(comboBox, instance)
	return comboBox
}

// AsComboBoxEx Convert a pointer object to an existing class object
func AsComboBoxEx(obj interface{}) IComboBoxEx {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	comboBoxEx := new(TComboBoxEx)
	SetObjectInstance(comboBoxEx, instance)
	return comboBoxEx
}

// AsComboExItem Convert a pointer object to an existing class object
func AsComboExItem(obj interface{}) IComboExItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	comboExItem := new(TComboExItem)
	SetObjectInstance(comboExItem, instance)
	return comboExItem
}

// AsComboExItems Convert a pointer object to an existing class object
func AsComboExItems(obj interface{}) IComboExItems {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	comboExItems := new(TComboExItems)
	SetObjectInstance(comboExItems, instance)
	return comboExItems
}

// AsCommonDialog Convert a pointer object to an existing class object
func AsCommonDialog(obj interface{}) ICommonDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	commonDialog := new(TCommonDialog)
	SetObjectInstance(commonDialog, instance)
	return commonDialog
}

// AsComponent Convert a pointer object to an existing class object
func AsComponent(obj interface{}) IComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	component := new(TComponent)
	SetObjectInstance(component, instance)
	return component
}

// AsComponentEnumerator Convert a pointer object to an existing class object
func AsComponentEnumerator(obj interface{}) IComponentEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	componentEnumerator := new(TComponentEnumerator)
	SetObjectInstance(componentEnumerator, instance)
	return componentEnumerator
}

// AsContainedAction Convert a pointer object to an existing class object
func AsContainedAction(obj interface{}) IContainedAction {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	containedAction := new(TContainedAction)
	SetObjectInstance(containedAction, instance)
	return containedAction
}

// AsControl Convert a pointer object to an existing class object
func AsControl(obj interface{}) IControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	control := new(TControl)
	SetObjectInstance(control, instance)
	return control
}

// AsControlBorderSpacing Convert a pointer object to an existing class object
func AsControlBorderSpacing(obj interface{}) IControlBorderSpacing {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	controlBorderSpacing := new(TControlBorderSpacing)
	SetObjectInstance(controlBorderSpacing, instance)
	return controlBorderSpacing
}

// AsControlChildSizing Convert a pointer object to an existing class object
func AsControlChildSizing(obj interface{}) IControlChildSizing {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	controlChildSizing := new(TControlChildSizing)
	SetObjectInstance(controlChildSizing, instance)
	return controlChildSizing
}

// AsControlScrollBar Convert a pointer object to an existing class object
func AsControlScrollBar(obj interface{}) IControlScrollBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	controlScrollBar := new(TControlScrollBar)
	SetObjectInstance(controlScrollBar, instance)
	return controlScrollBar
}

// AsCoolBand Convert a pointer object to an existing class object
func AsCoolBand(obj interface{}) ICoolBand {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coolBand := new(TCoolBand)
	SetObjectInstance(coolBand, instance)
	return coolBand
}

// AsCoolBands Convert a pointer object to an existing class object
func AsCoolBands(obj interface{}) ICoolBands {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coolBands := new(TCoolBands)
	SetObjectInstance(coolBands, instance)
	return coolBands
}

// AsCoolBar Convert a pointer object to an existing class object
func AsCoolBar(obj interface{}) ICoolBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	coolBar := new(TCoolBar)
	SetObjectInstance(coolBar, instance)
	return coolBar
}

// AsCustomAbstractGroupedEdit Convert a pointer object to an existing class object
func AsCustomAbstractGroupedEdit(obj interface{}) ICustomAbstractGroupedEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customAbstractGroupedEdit := new(TCustomAbstractGroupedEdit)
	SetObjectInstance(customAbstractGroupedEdit, instance)
	return customAbstractGroupedEdit
}

// AsCustomAction Convert a pointer object to an existing class object
func AsCustomAction(obj interface{}) ICustomAction {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customAction := new(TCustomAction)
	SetObjectInstance(customAction, instance)
	return customAction
}

// AsCustomActionList Convert a pointer object to an existing class object
func AsCustomActionList(obj interface{}) ICustomActionList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customActionList := new(TCustomActionList)
	SetObjectInstance(customActionList, instance)
	return customActionList
}

// AsCustomApplication Convert a pointer object to an existing class object
func AsCustomApplication(obj interface{}) ICustomApplication {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customApplication := new(TCustomApplication)
	SetObjectInstance(customApplication, instance)
	return customApplication
}

// AsCustomBitBtn Convert a pointer object to an existing class object
func AsCustomBitBtn(obj interface{}) ICustomBitBtn {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customBitBtn := new(TCustomBitBtn)
	SetObjectInstance(customBitBtn, instance)
	return customBitBtn
}

// AsCustomBitmap Convert a pointer object to an existing class object
func AsCustomBitmap(obj interface{}) ICustomBitmap {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customBitmap := new(TCustomBitmap)
	SetObjectInstance(customBitmap, instance)
	return customBitmap
}

// AsCustomButton Convert a pointer object to an existing class object
func AsCustomButton(obj interface{}) ICustomButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customButton := new(TCustomButton)
	SetObjectInstance(customButton, instance)
	return customButton
}

// AsCustomCalendar Convert a pointer object to an existing class object
func AsCustomCalendar(obj interface{}) ICustomCalendar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCalendar := new(TCustomCalendar)
	SetObjectInstance(customCalendar, instance)
	return customCalendar
}

// AsCustomCheckBox Convert a pointer object to an existing class object
func AsCustomCheckBox(obj interface{}) ICustomCheckBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCheckBox := new(TCustomCheckBox)
	SetObjectInstance(customCheckBox, instance)
	return customCheckBox
}

// AsCustomCheckCombo Convert a pointer object to an existing class object
func AsCustomCheckCombo(obj interface{}) ICustomCheckCombo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCheckCombo := new(TCustomCheckCombo)
	SetObjectInstance(customCheckCombo, instance)
	return customCheckCombo
}

// AsCustomCheckGroup Convert a pointer object to an existing class object
func AsCustomCheckGroup(obj interface{}) ICustomCheckGroup {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCheckGroup := new(TCustomCheckGroup)
	SetObjectInstance(customCheckGroup, instance)
	return customCheckGroup
}

// AsCustomCheckListBox Convert a pointer object to an existing class object
func AsCustomCheckListBox(obj interface{}) ICustomCheckListBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCheckListBox := new(TCustomCheckListBox)
	SetObjectInstance(customCheckListBox, instance)
	return customCheckListBox
}

// AsCustomColorBox Convert a pointer object to an existing class object
func AsCustomColorBox(obj interface{}) ICustomColorBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customColorBox := new(TCustomColorBox)
	SetObjectInstance(customColorBox, instance)
	return customColorBox
}

// AsCustomColorListBox Convert a pointer object to an existing class object
func AsCustomColorListBox(obj interface{}) ICustomColorListBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customColorListBox := new(TCustomColorListBox)
	SetObjectInstance(customColorListBox, instance)
	return customColorListBox
}

// AsCustomComboBox Convert a pointer object to an existing class object
func AsCustomComboBox(obj interface{}) ICustomComboBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customComboBox := new(TCustomComboBox)
	SetObjectInstance(customComboBox, instance)
	return customComboBox
}

// AsCustomComboBoxEx Convert a pointer object to an existing class object
func AsCustomComboBoxEx(obj interface{}) ICustomComboBoxEx {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customComboBoxEx := new(TCustomComboBoxEx)
	SetObjectInstance(customComboBoxEx, instance)
	return customComboBoxEx
}

// AsCustomControl Convert a pointer object to an existing class object
func AsCustomControl(obj interface{}) ICustomControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customControl := new(TCustomControl)
	SetObjectInstance(customControl, instance)
	return customControl
}

// AsCustomCoolBar Convert a pointer object to an existing class object
func AsCustomCoolBar(obj interface{}) ICustomCoolBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customCoolBar := new(TCustomCoolBar)
	SetObjectInstance(customCoolBar, instance)
	return customCoolBar
}

// AsCustomDateTimePicker Convert a pointer object to an existing class object
func AsCustomDateTimePicker(obj interface{}) ICustomDateTimePicker {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customDateTimePicker := new(TCustomDateTimePicker)
	SetObjectInstance(customDateTimePicker, instance)
	return customDateTimePicker
}

// AsCustomDesignControl Convert a pointer object to an existing class object
func AsCustomDesignControl(obj interface{}) ICustomDesignControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customDesignControl := new(TCustomDesignControl)
	SetObjectInstance(customDesignControl, instance)
	return customDesignControl
}

// AsCustomDrawGrid Convert a pointer object to an existing class object
func AsCustomDrawGrid(obj interface{}) ICustomDrawGrid {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customDrawGrid := new(TCustomDrawGrid)
	SetObjectInstance(customDrawGrid, instance)
	return customDrawGrid
}

// AsCustomEdit Convert a pointer object to an existing class object
func AsCustomEdit(obj interface{}) ICustomEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customEdit := new(TCustomEdit)
	SetObjectInstance(customEdit, instance)
	return customEdit
}

// AsCustomEditButton Convert a pointer object to an existing class object
func AsCustomEditButton(obj interface{}) ICustomEditButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customEditButton := new(TCustomEditButton)
	SetObjectInstance(customEditButton, instance)
	return customEditButton
}

// AsCustomFloatSpinEdit Convert a pointer object to an existing class object
func AsCustomFloatSpinEdit(obj interface{}) ICustomFloatSpinEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customFloatSpinEdit := new(TCustomFloatSpinEdit)
	SetObjectInstance(customFloatSpinEdit, instance)
	return customFloatSpinEdit
}

// AsCustomFlowPanel Convert a pointer object to an existing class object
func AsCustomFlowPanel(obj interface{}) ICustomFlowPanel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customFlowPanel := new(TCustomFlowPanel)
	SetObjectInstance(customFlowPanel, instance)
	return customFlowPanel
}

// AsCustomForm Convert a pointer object to an existing class object
func AsCustomForm(obj interface{}) ICustomForm {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customForm := new(TCustomForm)
	SetObjectInstance(customForm, instance)
	return customForm
}

// AsCustomFrame Convert a pointer object to an existing class object
func AsCustomFrame(obj interface{}) ICustomFrame {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customFrame := new(TCustomFrame)
	SetObjectInstance(customFrame, instance)
	return customFrame
}

// AsCustomGrid Convert a pointer object to an existing class object
func AsCustomGrid(obj interface{}) ICustomGrid {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customGrid := new(TCustomGrid)
	SetObjectInstance(customGrid, instance)
	return customGrid
}

// AsCustomGroupBox Convert a pointer object to an existing class object
func AsCustomGroupBox(obj interface{}) ICustomGroupBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customGroupBox := new(TCustomGroupBox)
	SetObjectInstance(customGroupBox, instance)
	return customGroupBox
}

// AsCustomHeaderControl Convert a pointer object to an existing class object
func AsCustomHeaderControl(obj interface{}) ICustomHeaderControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customHeaderControl := new(TCustomHeaderControl)
	SetObjectInstance(customHeaderControl, instance)
	return customHeaderControl
}

// AsCustomIcon Convert a pointer object to an existing class object
func AsCustomIcon(obj interface{}) ICustomIcon {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customIcon := new(TCustomIcon)
	SetObjectInstance(customIcon, instance)
	return customIcon
}

// AsCustomImage Convert a pointer object to an existing class object
func AsCustomImage(obj interface{}) ICustomImage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customImage := new(TCustomImage)
	SetObjectInstance(customImage, instance)
	return customImage
}

// AsCustomImageList Convert a pointer object to an existing class object
func AsCustomImageList(obj interface{}) ICustomImageList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customImageList := new(TCustomImageList)
	SetObjectInstance(customImageList, instance)
	return customImageList
}

// AsCustomImageListResolution Convert a pointer object to an existing class object
func AsCustomImageListResolution(obj interface{}) ICustomImageListResolution {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customImageListResolution := new(TCustomImageListResolution)
	SetObjectInstance(customImageListResolution, instance)
	return customImageListResolution
}

// AsCustomImageListResolutionEnumerator Convert a pointer object to an existing class object
func AsCustomImageListResolutionEnumerator(obj interface{}) ICustomImageListResolutionEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customImageListResolutionEnumerator := new(TCustomImageListResolutionEnumerator)
	SetObjectInstance(customImageListResolutionEnumerator, instance)
	return customImageListResolutionEnumerator
}

// AsCustomIniFile Convert a pointer object to an existing class object
func AsCustomIniFile(obj interface{}) ICustomIniFile {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customIniFile := new(TCustomIniFile)
	SetObjectInstance(customIniFile, instance)
	return customIniFile
}

// AsCustomLabel Convert a pointer object to an existing class object
func AsCustomLabel(obj interface{}) ICustomLabel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customLabel := new(TCustomLabel)
	SetObjectInstance(customLabel, instance)
	return customLabel
}

// AsCustomLabeledEdit Convert a pointer object to an existing class object
func AsCustomLabeledEdit(obj interface{}) ICustomLabeledEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customLabeledEdit := new(TCustomLabeledEdit)
	SetObjectInstance(customLabeledEdit, instance)
	return customLabeledEdit
}

// AsCustomListBox Convert a pointer object to an existing class object
func AsCustomListBox(obj interface{}) ICustomListBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customListBox := new(TCustomListBox)
	SetObjectInstance(customListBox, instance)
	return customListBox
}

// AsCustomListView Convert a pointer object to an existing class object
func AsCustomListView(obj interface{}) ICustomListView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customListView := new(TCustomListView)
	SetObjectInstance(customListView, instance)
	return customListView
}

// AsCustomMaskEdit Convert a pointer object to an existing class object
func AsCustomMaskEdit(obj interface{}) ICustomMaskEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customMaskEdit := new(TCustomMaskEdit)
	SetObjectInstance(customMaskEdit, instance)
	return customMaskEdit
}

// AsCustomMemo Convert a pointer object to an existing class object
func AsCustomMemo(obj interface{}) ICustomMemo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customMemo := new(TCustomMemo)
	SetObjectInstance(customMemo, instance)
	return customMemo
}

// AsCustomMemoryStream Convert a pointer object to an existing class object
func AsCustomMemoryStream(obj interface{}) ICustomMemoryStream {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customMemoryStream := new(TCustomMemoryStream)
	SetObjectInstance(customMemoryStream, instance)
	return customMemoryStream
}

// AsCustomPage Convert a pointer object to an existing class object
func AsCustomPage(obj interface{}) ICustomPage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customPage := new(TCustomPage)
	SetObjectInstance(customPage, instance)
	return customPage
}

// AsCustomPanel Convert a pointer object to an existing class object
func AsCustomPanel(obj interface{}) ICustomPanel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customPanel := new(TCustomPanel)
	SetObjectInstance(customPanel, instance)
	return customPanel
}

// AsCustomPrintDialog Convert a pointer object to an existing class object
func AsCustomPrintDialog(obj interface{}) ICustomPrintDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customPrintDialog := new(TCustomPrintDialog)
	SetObjectInstance(customPrintDialog, instance)
	return customPrintDialog
}

// AsCustomPrinterSetupDialog Convert a pointer object to an existing class object
func AsCustomPrinterSetupDialog(obj interface{}) ICustomPrinterSetupDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customPrinterSetupDialog := new(TCustomPrinterSetupDialog)
	SetObjectInstance(customPrinterSetupDialog, instance)
	return customPrinterSetupDialog
}

// AsCustomProgressBar Convert a pointer object to an existing class object
func AsCustomProgressBar(obj interface{}) ICustomProgressBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customProgressBar := new(TCustomProgressBar)
	SetObjectInstance(customProgressBar, instance)
	return customProgressBar
}

// AsCustomRadioGroup Convert a pointer object to an existing class object
func AsCustomRadioGroup(obj interface{}) ICustomRadioGroup {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customRadioGroup := new(TCustomRadioGroup)
	SetObjectInstance(customRadioGroup, instance)
	return customRadioGroup
}

// AsCustomRichMemo Convert a pointer object to an existing class object
func AsCustomRichMemo(obj interface{}) ICustomRichMemo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customRichMemo := new(TCustomRichMemo)
	SetObjectInstance(customRichMemo, instance)
	return customRichMemo
}

// AsCustomScrollBar Convert a pointer object to an existing class object
func AsCustomScrollBar(obj interface{}) ICustomScrollBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customScrollBar := new(TCustomScrollBar)
	SetObjectInstance(customScrollBar, instance)
	return customScrollBar
}

// AsCustomSpeedButton Convert a pointer object to an existing class object
func AsCustomSpeedButton(obj interface{}) ICustomSpeedButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customSpeedButton := new(TCustomSpeedButton)
	SetObjectInstance(customSpeedButton, instance)
	return customSpeedButton
}

// AsCustomSpinEdit Convert a pointer object to an existing class object
func AsCustomSpinEdit(obj interface{}) ICustomSpinEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customSpinEdit := new(TCustomSpinEdit)
	SetObjectInstance(customSpinEdit, instance)
	return customSpinEdit
}

// AsCustomSplitter Convert a pointer object to an existing class object
func AsCustomSplitter(obj interface{}) ICustomSplitter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customSplitter := new(TCustomSplitter)
	SetObjectInstance(customSplitter, instance)
	return customSplitter
}

// AsCustomStaticText Convert a pointer object to an existing class object
func AsCustomStaticText(obj interface{}) ICustomStaticText {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customStaticText := new(TCustomStaticText)
	SetObjectInstance(customStaticText, instance)
	return customStaticText
}

// AsCustomStringGrid Convert a pointer object to an existing class object
func AsCustomStringGrid(obj interface{}) ICustomStringGrid {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customStringGrid := new(TCustomStringGrid)
	SetObjectInstance(customStringGrid, instance)
	return customStringGrid
}

// AsCustomTabControl Convert a pointer object to an existing class object
func AsCustomTabControl(obj interface{}) ICustomTabControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customTabControl := new(TCustomTabControl)
	SetObjectInstance(customTabControl, instance)
	return customTabControl
}

// AsCustomTaskDialog Convert a pointer object to an existing class object
func AsCustomTaskDialog(obj interface{}) ICustomTaskDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customTaskDialog := new(TCustomTaskDialog)
	SetObjectInstance(customTaskDialog, instance)
	return customTaskDialog
}

// AsCustomTimer Convert a pointer object to an existing class object
func AsCustomTimer(obj interface{}) ICustomTimer {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customTimer := new(TCustomTimer)
	SetObjectInstance(customTimer, instance)
	return customTimer
}

// AsCustomTrackBar Convert a pointer object to an existing class object
func AsCustomTrackBar(obj interface{}) ICustomTrackBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customTrackBar := new(TCustomTrackBar)
	SetObjectInstance(customTrackBar, instance)
	return customTrackBar
}

// AsCustomTrayIcon Convert a pointer object to an existing class object
func AsCustomTrayIcon(obj interface{}) ICustomTrayIcon {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customTrayIcon := new(TCustomTrayIcon)
	SetObjectInstance(customTrayIcon, instance)
	return customTrayIcon
}

// AsCustomTreeView Convert a pointer object to an existing class object
func AsCustomTreeView(obj interface{}) ICustomTreeView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customTreeView := new(TCustomTreeView)
	SetObjectInstance(customTreeView, instance)
	return customTreeView
}

// AsCustomUpDown Convert a pointer object to an existing class object
func AsCustomUpDown(obj interface{}) ICustomUpDown {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	customUpDown := new(TCustomUpDown)
	SetObjectInstance(customUpDown, instance)
	return customUpDown
}

// AsDataModule Convert a pointer object to an existing class object
func AsDataModule(obj interface{}) IDataModule {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dataModule := new(TDataModule)
	SetObjectInstance(dataModule, instance)
	return dataModule
}

// AsDateTimePicker Convert a pointer object to an existing class object
func AsDateTimePicker(obj interface{}) IDateTimePicker {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dateTimePicker := new(TDateTimePicker)
	SetObjectInstance(dateTimePicker, instance)
	return dateTimePicker
}

// AsDirectoryEdit Convert a pointer object to an existing class object
func AsDirectoryEdit(obj interface{}) IDirectoryEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	directoryEdit := new(TDirectoryEdit)
	SetObjectInstance(directoryEdit, instance)
	return directoryEdit
}

// AsDockManager Convert a pointer object to an existing class object
func AsDockManager(obj interface{}) IDockManager {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dockManager := new(TDockManager)
	SetObjectInstance(dockManager, instance)
	return dockManager
}

// AsDockTree Convert a pointer object to an existing class object
func AsDockTree(obj interface{}) IDockTree {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dockTree := new(TDockTree)
	SetObjectInstance(dockTree, instance)
	return dockTree
}

// AsDockZone Convert a pointer object to an existing class object
func AsDockZone(obj interface{}) IDockZone {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dockZone := new(TDockZone)
	SetObjectInstance(dockZone, instance)
	return dockZone
}

// AsDragDockObject Convert a pointer object to an existing class object
func AsDragDockObject(obj interface{}) IDragDockObject {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dragDockObject := new(TDragDockObject)
	SetObjectInstance(dragDockObject, instance)
	return dragDockObject
}

// AsDragImageList Convert a pointer object to an existing class object
func AsDragImageList(obj interface{}) IDragImageList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dragImageList := new(TDragImageList)
	SetObjectInstance(dragImageList, instance)
	return dragImageList
}

// AsDragImageListResolution Convert a pointer object to an existing class object
func AsDragImageListResolution(obj interface{}) IDragImageListResolution {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dragImageListResolution := new(TDragImageListResolution)
	SetObjectInstance(dragImageListResolution, instance)
	return dragImageListResolution
}

// AsDragObject Convert a pointer object to an existing class object
func AsDragObject(obj interface{}) IDragObject {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	dragObject := new(TDragObject)
	SetObjectInstance(dragObject, instance)
	return dragObject
}

// AsDrawGrid Convert a pointer object to an existing class object
func AsDrawGrid(obj interface{}) IDrawGrid {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	drawGrid := new(TDrawGrid)
	SetObjectInstance(drawGrid, instance)
	return drawGrid
}

// AsEbEdit Convert a pointer object to an existing class object
func AsEbEdit(obj interface{}) IEbEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	ebEdit := new(TEbEdit)
	SetObjectInstance(ebEdit, instance)
	return ebEdit
}

// AsEdit Convert a pointer object to an existing class object
func AsEdit(obj interface{}) IEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	edit := new(TEdit)
	SetObjectInstance(edit, instance)
	return edit
}

// AsEditButton Convert a pointer object to an existing class object
func AsEditButton(obj interface{}) IEditButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	editButton := new(TEditButton)
	SetObjectInstance(editButton, instance)
	return editButton
}

// AsFPCanvasHelper Convert a pointer object to an existing class object
func AsFPCanvasHelper(obj interface{}) IFPCanvasHelper {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCanvasHelper := new(TFPCanvasHelper)
	SetObjectInstance(fPCanvasHelper, instance)
	return fPCanvasHelper
}

// AsFPCustomBrush Convert a pointer object to an existing class object
func AsFPCustomBrush(obj interface{}) IFPCustomBrush {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomBrush := new(TFPCustomBrush)
	SetObjectInstance(fPCustomBrush, instance)
	return fPCustomBrush
}

// AsFPCustomCanvas Convert a pointer object to an existing class object
func AsFPCustomCanvas(obj interface{}) IFPCustomCanvas {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomCanvas := new(TFPCustomCanvas)
	SetObjectInstance(fPCustomCanvas, instance)
	return fPCustomCanvas
}

// AsFPCustomFont Convert a pointer object to an existing class object
func AsFPCustomFont(obj interface{}) IFPCustomFont {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomFont := new(TFPCustomFont)
	SetObjectInstance(fPCustomFont, instance)
	return fPCustomFont
}

// AsFPCustomImage Convert a pointer object to an existing class object
func AsFPCustomImage(obj interface{}) IFPCustomImage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomImage := new(TFPCustomImage)
	SetObjectInstance(fPCustomImage, instance)
	return fPCustomImage
}

// AsFPCustomImageHandler Convert a pointer object to an existing class object
func AsFPCustomImageHandler(obj interface{}) IFPCustomImageHandler {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomImageHandler := new(TFPCustomImageHandler)
	SetObjectInstance(fPCustomImageHandler, instance)
	return fPCustomImageHandler
}

// AsFPCustomImageReader Convert a pointer object to an existing class object
func AsFPCustomImageReader(obj interface{}) IFPCustomImageReader {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomImageReader := new(TFPCustomImageReader)
	SetObjectInstance(fPCustomImageReader, instance)
	return fPCustomImageReader
}

// AsFPCustomImageWriter Convert a pointer object to an existing class object
func AsFPCustomImageWriter(obj interface{}) IFPCustomImageWriter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomImageWriter := new(TFPCustomImageWriter)
	SetObjectInstance(fPCustomImageWriter, instance)
	return fPCustomImageWriter
}

// AsFPCustomPen Convert a pointer object to an existing class object
func AsFPCustomPen(obj interface{}) IFPCustomPen {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomPen := new(TFPCustomPen)
	SetObjectInstance(fPCustomPen, instance)
	return fPCustomPen
}

// AsFPCustomRegion Convert a pointer object to an existing class object
func AsFPCustomRegion(obj interface{}) IFPCustomRegion {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPCustomRegion := new(TFPCustomRegion)
	SetObjectInstance(fPCustomRegion, instance)
	return fPCustomRegion
}

// AsFPImageBitmap Convert a pointer object to an existing class object
func AsFPImageBitmap(obj interface{}) IFPImageBitmap {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPImageBitmap := new(TFPImageBitmap)
	SetObjectInstance(fPImageBitmap, instance)
	return fPImageBitmap
}

// AsFPList Convert a pointer object to an existing class object
func AsFPList(obj interface{}) IFPList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPList := new(TFPList)
	SetObjectInstance(fPList, instance)
	return fPList
}

// AsFPListEnumerator Convert a pointer object to an existing class object
func AsFPListEnumerator(obj interface{}) IFPListEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPListEnumerator := new(TFPListEnumerator)
	SetObjectInstance(fPListEnumerator, instance)
	return fPListEnumerator
}

// AsFPMemoryImage Convert a pointer object to an existing class object
func AsFPMemoryImage(obj interface{}) IFPMemoryImage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPMemoryImage := new(TFPMemoryImage)
	SetObjectInstance(fPMemoryImage, instance)
	return fPMemoryImage
}

// AsFPPalette Convert a pointer object to an existing class object
func AsFPPalette(obj interface{}) IFPPalette {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPPalette := new(TFPPalette)
	SetObjectInstance(fPPalette, instance)
	return fPPalette
}

// AsFPRectRegion Convert a pointer object to an existing class object
func AsFPRectRegion(obj interface{}) IFPRectRegion {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fPRectRegion := new(TFPRectRegion)
	SetObjectInstance(fPRectRegion, instance)
	return fPRectRegion
}

// AsFileDialog Convert a pointer object to an existing class object
func AsFileDialog(obj interface{}) IFileDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fileDialog := new(TFileDialog)
	SetObjectInstance(fileDialog, instance)
	return fileDialog
}

// AsFindDialog Convert a pointer object to an existing class object
func AsFindDialog(obj interface{}) IFindDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	findDialog := new(TFindDialog)
	SetObjectInstance(findDialog, instance)
	return findDialog
}

// AsFloatSpinEdit Convert a pointer object to an existing class object
func AsFloatSpinEdit(obj interface{}) IFloatSpinEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	floatSpinEdit := new(TFloatSpinEdit)
	SetObjectInstance(floatSpinEdit, instance)
	return floatSpinEdit
}

// AsFlowPanel Convert a pointer object to an existing class object
func AsFlowPanel(obj interface{}) IFlowPanel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	flowPanel := new(TFlowPanel)
	SetObjectInstance(flowPanel, instance)
	return flowPanel
}

// AsFlowPanelControl Convert a pointer object to an existing class object
func AsFlowPanelControl(obj interface{}) IFlowPanelControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	flowPanelControl := new(TFlowPanelControl)
	SetObjectInstance(flowPanelControl, instance)
	return flowPanelControl
}

// AsFlowPanelControlList Convert a pointer object to an existing class object
func AsFlowPanelControlList(obj interface{}) IFlowPanelControlList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	flowPanelControlList := new(TFlowPanelControlList)
	SetObjectInstance(flowPanelControlList, instance)
	return flowPanelControlList
}

// AsFont Convert a pointer object to an existing class object
func AsFont(obj interface{}) IFont {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	font := new(TFont)
	SetObjectInstance(font, instance)
	return font
}

// AsFontDialog Convert a pointer object to an existing class object
func AsFontDialog(obj interface{}) IFontDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	fontDialog := new(TFontDialog)
	SetObjectInstance(fontDialog, instance)
	return fontDialog
}

// AsForm Convert a pointer object to an existing class object
func AsForm(obj interface{}) IForm {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	form := new(TForm)
	SetObjectInstance(form, instance)
	return form
}

// AsFrame Convert a pointer object to an existing class object
func AsFrame(obj interface{}) IFrame {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	frame := new(TFrame)
	SetObjectInstance(frame, instance)
	return frame
}

// AsGEEdit Convert a pointer object to an existing class object
func AsGEEdit(obj interface{}) IGEEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	gEEdit := new(TGEEdit)
	SetObjectInstance(gEEdit, instance)
	return gEEdit
}

// AsGIFImage Convert a pointer object to an existing class object
func AsGIFImage(obj interface{}) IGIFImage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	gIFImage := new(TGIFImage)
	SetObjectInstance(gIFImage, instance)
	return gIFImage
}

// AsGauge Convert a pointer object to an existing class object
func AsGauge(obj interface{}) IGauge {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	gauge := new(TGauge)
	SetObjectInstance(gauge, instance)
	return gauge
}

// AsGraphic Convert a pointer object to an existing class object
func AsGraphic(obj interface{}) IGraphic {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	graphic := new(TGraphic)
	SetObjectInstance(graphic, instance)
	return graphic
}

// AsGraphicControl Convert a pointer object to an existing class object
func AsGraphicControl(obj interface{}) IGraphicControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	graphicControl := new(TGraphicControl)
	SetObjectInstance(graphicControl, instance)
	return graphicControl
}

// AsGraphicsObject Convert a pointer object to an existing class object
func AsGraphicsObject(obj interface{}) IGraphicsObject {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	graphicsObject := new(TGraphicsObject)
	SetObjectInstance(graphicsObject, instance)
	return graphicsObject
}

// AsGridColumn Convert a pointer object to an existing class object
func AsGridColumn(obj interface{}) IGridColumn {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	gridColumn := new(TGridColumn)
	SetObjectInstance(gridColumn, instance)
	return gridColumn
}

// AsGridColumnTitle Convert a pointer object to an existing class object
func AsGridColumnTitle(obj interface{}) IGridColumnTitle {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	gridColumnTitle := new(TGridColumnTitle)
	SetObjectInstance(gridColumnTitle, instance)
	return gridColumnTitle
}

// AsGridColumns Convert a pointer object to an existing class object
func AsGridColumns(obj interface{}) IGridColumns {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	gridColumns := new(TGridColumns)
	SetObjectInstance(gridColumns, instance)
	return gridColumns
}

// AsGroupBox Convert a pointer object to an existing class object
func AsGroupBox(obj interface{}) IGroupBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	groupBox := new(TGroupBox)
	SetObjectInstance(groupBox, instance)
	return groupBox
}

// AsHeaderControl Convert a pointer object to an existing class object
func AsHeaderControl(obj interface{}) IHeaderControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	headerControl := new(THeaderControl)
	SetObjectInstance(headerControl, instance)
	return headerControl
}

// AsHeaderSection Convert a pointer object to an existing class object
func AsHeaderSection(obj interface{}) IHeaderSection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	headerSection := new(THeaderSection)
	SetObjectInstance(headerSection, instance)
	return headerSection
}

// AsHeaderSections Convert a pointer object to an existing class object
func AsHeaderSections(obj interface{}) IHeaderSections {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	headerSections := new(THeaderSections)
	SetObjectInstance(headerSections, instance)
	return headerSections
}

// AsIDesigner Convert a pointer object to an existing class object
func AsIDesigner(obj interface{}) IIDesigner {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	iDesigner := new(TIDesigner)
	SetObjectInstance(iDesigner, instance)
	return iDesigner
}

// AsIcon Convert a pointer object to an existing class object
func AsIcon(obj interface{}) IIcon {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	icon := new(TIcon)
	SetObjectInstance(icon, instance)
	return icon
}

// AsIconOptions Convert a pointer object to an existing class object
func AsIconOptions(obj interface{}) IIconOptions {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	iconOptions := new(TIconOptions)
	SetObjectInstance(iconOptions, instance)
	return iconOptions
}

// AsImage Convert a pointer object to an existing class object
func AsImage(obj interface{}) IImage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	image := new(TImage)
	SetObjectInstance(image, instance)
	return image
}

// AsImageButton Convert a pointer object to an existing class object
func AsImageButton(obj interface{}) IImageButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	imageButton := new(TImageButton)
	SetObjectInstance(imageButton, instance)
	return imageButton
}

// AsImageList Convert a pointer object to an existing class object
func AsImageList(obj interface{}) IImageList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	imageList := new(TImageList)
	SetObjectInstance(imageList, instance)
	return imageList
}

// AsIniFile Convert a pointer object to an existing class object
func AsIniFile(obj interface{}) IIniFile {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	iniFile := new(TIniFile)
	SetObjectInstance(iniFile, instance)
	return iniFile
}

// AsItemProp Convert a pointer object to an existing class object
func AsItemProp(obj interface{}) IItemProp {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	itemProp := new(TItemProp)
	SetObjectInstance(itemProp, instance)
	return itemProp
}

// AsJPEGImage Convert a pointer object to an existing class object
func AsJPEGImage(obj interface{}) IJPEGImage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	jPEGImage := new(TJPEGImage)
	SetObjectInstance(jPEGImage, instance)
	return jPEGImage
}

// AsLCLComponent Convert a pointer object to an existing class object
func AsLCLComponent(obj interface{}) ILCLComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lCLComponent := new(TLCLComponent)
	SetObjectInstance(lCLComponent, instance)
	return lCLComponent
}

// AsLCLReferenceComponent Convert a pointer object to an existing class object
func AsLCLReferenceComponent(obj interface{}) ILCLReferenceComponent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lCLReferenceComponent := new(TLCLReferenceComponent)
	SetObjectInstance(lCLReferenceComponent, instance)
	return lCLReferenceComponent
}

// AsLabel Convert a pointer object to an existing class object
func AsLabel(obj interface{}) ILabel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	label := new(TLabel)
	SetObjectInstance(label, instance)
	return label
}

// AsLabeledEdit Convert a pointer object to an existing class object
func AsLabeledEdit(obj interface{}) ILabeledEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	labeledEdit := new(TLabeledEdit)
	SetObjectInstance(labeledEdit, instance)
	return labeledEdit
}

// AsLazAccessibleObject Convert a pointer object to an existing class object
func AsLazAccessibleObject(obj interface{}) ILazAccessibleObject {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lazAccessibleObject := new(TLazAccessibleObject)
	SetObjectInstance(lazAccessibleObject, instance)
	return lazAccessibleObject
}

// AsLazAccessibleObjectEnumerator Convert a pointer object to an existing class object
func AsLazAccessibleObjectEnumerator(obj interface{}) ILazAccessibleObjectEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lazAccessibleObjectEnumerator := new(TLazAccessibleObjectEnumerator)
	SetObjectInstance(lazAccessibleObjectEnumerator, instance)
	return lazAccessibleObjectEnumerator
}

// AsLazDockForm Convert a pointer object to an existing class object
func AsLazDockForm(obj interface{}) ILazDockForm {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lazDockForm := new(TLazDockForm)
	SetObjectInstance(lazDockForm, instance)
	return lazDockForm
}

// AsLazDockPage Convert a pointer object to an existing class object
func AsLazDockPage(obj interface{}) ILazDockPage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lazDockPage := new(TLazDockPage)
	SetObjectInstance(lazDockPage, instance)
	return lazDockPage
}

// AsLazDockPages Convert a pointer object to an existing class object
func AsLazDockPages(obj interface{}) ILazDockPages {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lazDockPages := new(TLazDockPages)
	SetObjectInstance(lazDockPages, instance)
	return lazDockPages
}

// AsLazDockSplitter Convert a pointer object to an existing class object
func AsLazDockSplitter(obj interface{}) ILazDockSplitter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lazDockSplitter := new(TLazDockSplitter)
	SetObjectInstance(lazDockSplitter, instance)
	return lazDockSplitter
}

// AsLazDockTree Convert a pointer object to an existing class object
func AsLazDockTree(obj interface{}) ILazDockTree {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lazDockTree := new(TLazDockTree)
	SetObjectInstance(lazDockTree, instance)
	return lazDockTree
}

// AsLazDockZone Convert a pointer object to an existing class object
func AsLazDockZone(obj interface{}) ILazDockZone {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	lazDockZone := new(TLazDockZone)
	SetObjectInstance(lazDockZone, instance)
	return lazDockZone
}

// AsLinkLabel Convert a pointer object to an existing class object
func AsLinkLabel(obj interface{}) ILinkLabel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	linkLabel := new(TLinkLabel)
	SetObjectInstance(linkLabel, instance)
	return linkLabel
}

// AsList Convert a pointer object to an existing class object
func AsList(obj interface{}) IList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	list := new(TList)
	SetObjectInstance(list, instance)
	return list
}

// AsListBox Convert a pointer object to an existing class object
func AsListBox(obj interface{}) IListBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listBox := new(TListBox)
	SetObjectInstance(listBox, instance)
	return listBox
}

// AsListColumn Convert a pointer object to an existing class object
func AsListColumn(obj interface{}) IListColumn {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listColumn := new(TListColumn)
	SetObjectInstance(listColumn, instance)
	return listColumn
}

// AsListColumns Convert a pointer object to an existing class object
func AsListColumns(obj interface{}) IListColumns {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listColumns := new(TListColumns)
	SetObjectInstance(listColumns, instance)
	return listColumns
}

// AsListControlItem Convert a pointer object to an existing class object
func AsListControlItem(obj interface{}) IListControlItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listControlItem := new(TListControlItem)
	SetObjectInstance(listControlItem, instance)
	return listControlItem
}

// AsListControlItems Convert a pointer object to an existing class object
func AsListControlItems(obj interface{}) IListControlItems {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listControlItems := new(TListControlItems)
	SetObjectInstance(listControlItems, instance)
	return listControlItems
}

// AsListEnumerator Convert a pointer object to an existing class object
func AsListEnumerator(obj interface{}) IListEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listEnumerator := new(TListEnumerator)
	SetObjectInstance(listEnumerator, instance)
	return listEnumerator
}

// AsListItem Convert a pointer object to an existing class object
func AsListItem(obj interface{}) IListItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listItem := new(TListItem)
	SetObjectInstance(listItem, instance)
	return listItem
}

// AsListItems Convert a pointer object to an existing class object
func AsListItems(obj interface{}) IListItems {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listItems := new(TListItems)
	SetObjectInstance(listItems, instance)
	return listItems
}

// AsListItemsEnumerator Convert a pointer object to an existing class object
func AsListItemsEnumerator(obj interface{}) IListItemsEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listItemsEnumerator := new(TListItemsEnumerator)
	SetObjectInstance(listItemsEnumerator, instance)
	return listItemsEnumerator
}

// AsListView Convert a pointer object to an existing class object
func AsListView(obj interface{}) IListView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	listView := new(TListView)
	SetObjectInstance(listView, instance)
	return listView
}

// AsMainMenu Convert a pointer object to an existing class object
func AsMainMenu(obj interface{}) IMainMenu {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	mainMenu := new(TMainMenu)
	SetObjectInstance(mainMenu, instance)
	return mainMenu
}

// AsMaskEdit Convert a pointer object to an existing class object
func AsMaskEdit(obj interface{}) IMaskEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	maskEdit := new(TMaskEdit)
	SetObjectInstance(maskEdit, instance)
	return maskEdit
}

// AsMemo Convert a pointer object to an existing class object
func AsMemo(obj interface{}) IMemo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	memo := new(TMemo)
	SetObjectInstance(memo, instance)
	return memo
}

// AsMemoScrollBar Convert a pointer object to an existing class object
func AsMemoScrollBar(obj interface{}) IMemoScrollBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	memoScrollBar := new(TMemoScrollBar)
	SetObjectInstance(memoScrollBar, instance)
	return memoScrollBar
}

// AsMemoryStream Convert a pointer object to an existing class object
func AsMemoryStream(obj interface{}) IMemoryStream {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	memoryStream := new(TMemoryStream)
	SetObjectInstance(memoryStream, instance)
	return memoryStream
}

// AsMenu Convert a pointer object to an existing class object
func AsMenu(obj interface{}) IMenu {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	menu := new(TMenu)
	SetObjectInstance(menu, instance)
	return menu
}

// AsMenuItem Convert a pointer object to an existing class object
func AsMenuItem(obj interface{}) IMenuItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	menuItem := new(TMenuItem)
	SetObjectInstance(menuItem, instance)
	return menuItem
}

// AsMenuItemEnumerator Convert a pointer object to an existing class object
func AsMenuItemEnumerator(obj interface{}) IMenuItemEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	menuItemEnumerator := new(TMenuItemEnumerator)
	SetObjectInstance(menuItemEnumerator, instance)
	return menuItemEnumerator
}

// AsMergedMenuItems Convert a pointer object to an existing class object
func AsMergedMenuItems(obj interface{}) IMergedMenuItems {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	mergedMenuItems := new(TMergedMenuItems)
	SetObjectInstance(mergedMenuItems, instance)
	return mergedMenuItems
}

// AsMonitor Convert a pointer object to an existing class object
func AsMonitor(obj interface{}) IMonitor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	monitor := new(TMonitor)
	SetObjectInstance(monitor, instance)
	return monitor
}

// AsMouse Convert a pointer object to an existing class object
func AsMouse(obj interface{}) IMouse {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	mouse := new(TMouse)
	SetObjectInstance(mouse, instance)
	return mouse
}

// AsObject Convert a pointer object to an existing class object
func AsObject(obj interface{}) IObject {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	object := new(TObject)
	SetObjectInstance(object, instance)
	return object
}

// AsOpenDialog Convert a pointer object to an existing class object
func AsOpenDialog(obj interface{}) IOpenDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	openDialog := new(TOpenDialog)
	SetObjectInstance(openDialog, instance)
	return openDialog
}

// AsOpenPictureDialog Convert a pointer object to an existing class object
func AsOpenPictureDialog(obj interface{}) IOpenPictureDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	openPictureDialog := new(TOpenPictureDialog)
	SetObjectInstance(openPictureDialog, instance)
	return openPictureDialog
}

// AsOwnedCollection Convert a pointer object to an existing class object
func AsOwnedCollection(obj interface{}) IOwnedCollection {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	ownedCollection := new(TOwnedCollection)
	SetObjectInstance(ownedCollection, instance)
	return ownedCollection
}

// AsPageControl Convert a pointer object to an existing class object
func AsPageControl(obj interface{}) IPageControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	pageControl := new(TPageControl)
	SetObjectInstance(pageControl, instance)
	return pageControl
}

// AsPageSetupDialog Convert a pointer object to an existing class object
func AsPageSetupDialog(obj interface{}) IPageSetupDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	pageSetupDialog := new(TPageSetupDialog)
	SetObjectInstance(pageSetupDialog, instance)
	return pageSetupDialog
}

// AsPaintBox Convert a pointer object to an existing class object
func AsPaintBox(obj interface{}) IPaintBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	paintBox := new(TPaintBox)
	SetObjectInstance(paintBox, instance)
	return paintBox
}

// AsPanel Convert a pointer object to an existing class object
func AsPanel(obj interface{}) IPanel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	panel := new(TPanel)
	SetObjectInstance(panel, instance)
	return panel
}

// AsPaperSize Convert a pointer object to an existing class object
func AsPaperSize(obj interface{}) IPaperSize {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	paperSize := new(TPaperSize)
	SetObjectInstance(paperSize, instance)
	return paperSize
}

// AsParaAttributes Convert a pointer object to an existing class object
func AsParaAttributes(obj interface{}) IParaAttributes {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	paraAttributes := new(TParaAttributes)
	SetObjectInstance(paraAttributes, instance)
	return paraAttributes
}

// AsPen Convert a pointer object to an existing class object
func AsPen(obj interface{}) IPen {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	pen := new(TPen)
	SetObjectInstance(pen, instance)
	return pen
}

// AsPersistent Convert a pointer object to an existing class object
func AsPersistent(obj interface{}) IPersistent {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	persistent := new(TPersistent)
	SetObjectInstance(persistent, instance)
	return persistent
}

// AsPicture Convert a pointer object to an existing class object
func AsPicture(obj interface{}) IPicture {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	picture := new(TPicture)
	SetObjectInstance(picture, instance)
	return picture
}

// AsPixmap Convert a pointer object to an existing class object
func AsPixmap(obj interface{}) IPixmap {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	pixmap := new(TPixmap)
	SetObjectInstance(pixmap, instance)
	return pixmap
}

// AsPopupMenu Convert a pointer object to an existing class object
func AsPopupMenu(obj interface{}) IPopupMenu {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	popupMenu := new(TPopupMenu)
	SetObjectInstance(popupMenu, instance)
	return popupMenu
}

// AsPortableAnyMapGraphic Convert a pointer object to an existing class object
func AsPortableAnyMapGraphic(obj interface{}) IPortableAnyMapGraphic {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	portableAnyMapGraphic := new(TPortableAnyMapGraphic)
	SetObjectInstance(portableAnyMapGraphic, instance)
	return portableAnyMapGraphic
}

// AsPortableNetworkGraphic Convert a pointer object to an existing class object
func AsPortableNetworkGraphic(obj interface{}) IPortableNetworkGraphic {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	portableNetworkGraphic := new(TPortableNetworkGraphic)
	SetObjectInstance(portableNetworkGraphic, instance)
	return portableNetworkGraphic
}

// AsPreviewFileControl Convert a pointer object to an existing class object
func AsPreviewFileControl(obj interface{}) IPreviewFileControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	previewFileControl := new(TPreviewFileControl)
	SetObjectInstance(previewFileControl, instance)
	return previewFileControl
}

// AsPreviewFileDialog Convert a pointer object to an existing class object
func AsPreviewFileDialog(obj interface{}) IPreviewFileDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	previewFileDialog := new(TPreviewFileDialog)
	SetObjectInstance(previewFileDialog, instance)
	return previewFileDialog
}

// AsPrintDialog Convert a pointer object to an existing class object
func AsPrintDialog(obj interface{}) IPrintDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	printDialog := new(TPrintDialog)
	SetObjectInstance(printDialog, instance)
	return printDialog
}

// AsPrinter Convert a pointer object to an existing class object
func AsPrinter(obj interface{}) IPrinter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	printer := new(TPrinter)
	SetObjectInstance(printer, instance)
	return printer
}

// AsPrinterCanvas Convert a pointer object to an existing class object
func AsPrinterCanvas(obj interface{}) IPrinterCanvas {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	printerCanvas := new(TPrinterCanvas)
	SetObjectInstance(printerCanvas, instance)
	return printerCanvas
}

// AsPrinterSetupDialog Convert a pointer object to an existing class object
func AsPrinterSetupDialog(obj interface{}) IPrinterSetupDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	printerSetupDialog := new(TPrinterSetupDialog)
	SetObjectInstance(printerSetupDialog, instance)
	return printerSetupDialog
}

// AsProgressBar Convert a pointer object to an existing class object
func AsProgressBar(obj interface{}) IProgressBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	progressBar := new(TProgressBar)
	SetObjectInstance(progressBar, instance)
	return progressBar
}

// AsRadioButton Convert a pointer object to an existing class object
func AsRadioButton(obj interface{}) IRadioButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	radioButton := new(TRadioButton)
	SetObjectInstance(radioButton, instance)
	return radioButton
}

// AsRadioGroup Convert a pointer object to an existing class object
func AsRadioGroup(obj interface{}) IRadioGroup {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	radioGroup := new(TRadioGroup)
	SetObjectInstance(radioGroup, instance)
	return radioGroup
}

// AsRasterImage Convert a pointer object to an existing class object
func AsRasterImage(obj interface{}) IRasterImage {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	rasterImage := new(TRasterImage)
	SetObjectInstance(rasterImage, instance)
	return rasterImage
}

// AsRegion Convert a pointer object to an existing class object
func AsRegion(obj interface{}) IRegion {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	region := new(TRegion)
	SetObjectInstance(region, instance)
	return region
}

// AsRegistry Convert a pointer object to an existing class object
func AsRegistry(obj interface{}) IRegistry {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	registry := new(TRegistry)
	SetObjectInstance(registry, instance)
	return registry
}

// AsReplaceDialog Convert a pointer object to an existing class object
func AsReplaceDialog(obj interface{}) IReplaceDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	replaceDialog := new(TReplaceDialog)
	SetObjectInstance(replaceDialog, instance)
	return replaceDialog
}

// AsRichEdit Convert a pointer object to an existing class object
func AsRichEdit(obj interface{}) IRichEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	richEdit := new(TRichEdit)
	SetObjectInstance(richEdit, instance)
	return richEdit
}

// AsRichMemo Convert a pointer object to an existing class object
func AsRichMemo(obj interface{}) IRichMemo {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	richMemo := new(TRichMemo)
	SetObjectInstance(richMemo, instance)
	return richMemo
}

// AsRichMemoInline Convert a pointer object to an existing class object
func AsRichMemoInline(obj interface{}) IRichMemoInline {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	richMemoInline := new(TRichMemoInline)
	SetObjectInstance(richMemoInline, instance)
	return richMemoInline
}

// AsSaveDialog Convert a pointer object to an existing class object
func AsSaveDialog(obj interface{}) ISaveDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	saveDialog := new(TSaveDialog)
	SetObjectInstance(saveDialog, instance)
	return saveDialog
}

// AsSavePictureDialog Convert a pointer object to an existing class object
func AsSavePictureDialog(obj interface{}) ISavePictureDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	savePictureDialog := new(TSavePictureDialog)
	SetObjectInstance(savePictureDialog, instance)
	return savePictureDialog
}

// AsScreen Convert a pointer object to an existing class object
func AsScreen(obj interface{}) IScreen {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	screen := new(TScreen)
	SetObjectInstance(screen, instance)
	return screen
}

// AsScrollBar Convert a pointer object to an existing class object
func AsScrollBar(obj interface{}) IScrollBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	scrollBar := new(TScrollBar)
	SetObjectInstance(scrollBar, instance)
	return scrollBar
}

// AsScrollBox Convert a pointer object to an existing class object
func AsScrollBox(obj interface{}) IScrollBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	scrollBox := new(TScrollBox)
	SetObjectInstance(scrollBox, instance)
	return scrollBox
}

// AsScrollingWinControl Convert a pointer object to an existing class object
func AsScrollingWinControl(obj interface{}) IScrollingWinControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	scrollingWinControl := new(TScrollingWinControl)
	SetObjectInstance(scrollingWinControl, instance)
	return scrollingWinControl
}

// AsSelectDirectoryDialog Convert a pointer object to an existing class object
func AsSelectDirectoryDialog(obj interface{}) ISelectDirectoryDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	selectDirectoryDialog := new(TSelectDirectoryDialog)
	SetObjectInstance(selectDirectoryDialog, instance)
	return selectDirectoryDialog
}

// AsShape Convert a pointer object to an existing class object
func AsShape(obj interface{}) IShape {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	shape := new(TShape)
	SetObjectInstance(shape, instance)
	return shape
}

// AsShortCutList Convert a pointer object to an existing class object
func AsShortCutList(obj interface{}) IShortCutList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	shortCutList := new(TShortCutList)
	SetObjectInstance(shortCutList, instance)
	return shortCutList
}

// AsSizeConstraints Convert a pointer object to an existing class object
func AsSizeConstraints(obj interface{}) ISizeConstraints {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	sizeConstraints := new(TSizeConstraints)
	SetObjectInstance(sizeConstraints, instance)
	return sizeConstraints
}

// AsSpeedButton Convert a pointer object to an existing class object
func AsSpeedButton(obj interface{}) ISpeedButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	speedButton := new(TSpeedButton)
	SetObjectInstance(speedButton, instance)
	return speedButton
}

// AsSpinEdit Convert a pointer object to an existing class object
func AsSpinEdit(obj interface{}) ISpinEdit {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	spinEdit := new(TSpinEdit)
	SetObjectInstance(spinEdit, instance)
	return spinEdit
}

// AsSplitter Convert a pointer object to an existing class object
func AsSplitter(obj interface{}) ISplitter {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	splitter := new(TSplitter)
	SetObjectInstance(splitter, instance)
	return splitter
}

// AsStaticText Convert a pointer object to an existing class object
func AsStaticText(obj interface{}) IStaticText {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	staticText := new(TStaticText)
	SetObjectInstance(staticText, instance)
	return staticText
}

// AsStatusBar Convert a pointer object to an existing class object
func AsStatusBar(obj interface{}) IStatusBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	statusBar := new(TStatusBar)
	SetObjectInstance(statusBar, instance)
	return statusBar
}

// AsStatusPanel Convert a pointer object to an existing class object
func AsStatusPanel(obj interface{}) IStatusPanel {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	statusPanel := new(TStatusPanel)
	SetObjectInstance(statusPanel, instance)
	return statusPanel
}

// AsStatusPanels Convert a pointer object to an existing class object
func AsStatusPanels(obj interface{}) IStatusPanels {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	statusPanels := new(TStatusPanels)
	SetObjectInstance(statusPanels, instance)
	return statusPanels
}

// AsStream Convert a pointer object to an existing class object
func AsStream(obj interface{}) IStream {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	stream := new(TStream)
	SetObjectInstance(stream, instance)
	return stream
}

// AsStringGrid Convert a pointer object to an existing class object
func AsStringGrid(obj interface{}) IStringGrid {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	stringGrid := new(TStringGrid)
	SetObjectInstance(stringGrid, instance)
	return stringGrid
}

// AsStringList Convert a pointer object to an existing class object
func AsStringList(obj interface{}) IStringList {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	stringList := new(TStringList)
	SetObjectInstance(stringList, instance)
	return stringList
}

// AsStrings Convert a pointer object to an existing class object
func AsStrings(obj interface{}) IStrings {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	strings := new(TStrings)
	SetObjectInstance(strings, instance)
	return strings
}

// AsStringsEnumerator Convert a pointer object to an existing class object
func AsStringsEnumerator(obj interface{}) IStringsEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	stringsEnumerator := new(TStringsEnumerator)
	SetObjectInstance(stringsEnumerator, instance)
	return stringsEnumerator
}

// AsTabSheet Convert a pointer object to an existing class object
func AsTabSheet(obj interface{}) ITabSheet {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	tabSheet := new(TTabSheet)
	SetObjectInstance(tabSheet, instance)
	return tabSheet
}

// AsTaskDialog Convert a pointer object to an existing class object
func AsTaskDialog(obj interface{}) ITaskDialog {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	taskDialog := new(TTaskDialog)
	SetObjectInstance(taskDialog, instance)
	return taskDialog
}

// AsTaskDialogBaseButtonItem Convert a pointer object to an existing class object
func AsTaskDialogBaseButtonItem(obj interface{}) ITaskDialogBaseButtonItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	taskDialogBaseButtonItem := new(TTaskDialogBaseButtonItem)
	SetObjectInstance(taskDialogBaseButtonItem, instance)
	return taskDialogBaseButtonItem
}

// AsTaskDialogButtonItem Convert a pointer object to an existing class object
func AsTaskDialogButtonItem(obj interface{}) ITaskDialogButtonItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	taskDialogButtonItem := new(TTaskDialogButtonItem)
	SetObjectInstance(taskDialogButtonItem, instance)
	return taskDialogButtonItem
}

// AsTaskDialogButtons Convert a pointer object to an existing class object
func AsTaskDialogButtons(obj interface{}) ITaskDialogButtons {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	taskDialogButtons := new(TTaskDialogButtons)
	SetObjectInstance(taskDialogButtons, instance)
	return taskDialogButtons
}

// AsTaskDialogButtonsEnumerator Convert a pointer object to an existing class object
func AsTaskDialogButtonsEnumerator(obj interface{}) ITaskDialogButtonsEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	taskDialogButtonsEnumerator := new(TTaskDialogButtonsEnumerator)
	SetObjectInstance(taskDialogButtonsEnumerator, instance)
	return taskDialogButtonsEnumerator
}

// AsTaskDialogRadioButtonItem Convert a pointer object to an existing class object
func AsTaskDialogRadioButtonItem(obj interface{}) ITaskDialogRadioButtonItem {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	taskDialogRadioButtonItem := new(TTaskDialogRadioButtonItem)
	SetObjectInstance(taskDialogRadioButtonItem, instance)
	return taskDialogRadioButtonItem
}

// AsTextAttributes Convert a pointer object to an existing class object
func AsTextAttributes(obj interface{}) ITextAttributes {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	textAttributes := new(TTextAttributes)
	SetObjectInstance(textAttributes, instance)
	return textAttributes
}

// AsThemeServices Convert a pointer object to an existing class object
func AsThemeServices(obj interface{}) IThemeServices {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	themeServices := new(TThemeServices)
	SetObjectInstance(themeServices, instance)
	return themeServices
}

// AsTimer Convert a pointer object to an existing class object
func AsTimer(obj interface{}) ITimer {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	timer := new(TTimer)
	SetObjectInstance(timer, instance)
	return timer
}

// AsToggleBox Convert a pointer object to an existing class object
func AsToggleBox(obj interface{}) IToggleBox {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	toggleBox := new(TToggleBox)
	SetObjectInstance(toggleBox, instance)
	return toggleBox
}

// AsToolBar Convert a pointer object to an existing class object
func AsToolBar(obj interface{}) IToolBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	toolBar := new(TToolBar)
	SetObjectInstance(toolBar, instance)
	return toolBar
}

// AsToolBarEnumerator Convert a pointer object to an existing class object
func AsToolBarEnumerator(obj interface{}) IToolBarEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	toolBarEnumerator := new(TToolBarEnumerator)
	SetObjectInstance(toolBarEnumerator, instance)
	return toolBarEnumerator
}

// AsToolButton Convert a pointer object to an existing class object
func AsToolButton(obj interface{}) IToolButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	toolButton := new(TToolButton)
	SetObjectInstance(toolButton, instance)
	return toolButton
}

// AsToolWindow Convert a pointer object to an existing class object
func AsToolWindow(obj interface{}) IToolWindow {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	toolWindow := new(TToolWindow)
	SetObjectInstance(toolWindow, instance)
	return toolWindow
}

// AsTrackBar Convert a pointer object to an existing class object
func AsTrackBar(obj interface{}) ITrackBar {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	trackBar := new(TTrackBar)
	SetObjectInstance(trackBar, instance)
	return trackBar
}

// AsTrayIcon Convert a pointer object to an existing class object
func AsTrayIcon(obj interface{}) ITrayIcon {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	trayIcon := new(TTrayIcon)
	SetObjectInstance(trayIcon, instance)
	return trayIcon
}

// AsTreeNode Convert a pointer object to an existing class object
func AsTreeNode(obj interface{}) ITreeNode {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	treeNode := new(TTreeNode)
	SetObjectInstance(treeNode, instance)
	return treeNode
}

// AsTreeNodes Convert a pointer object to an existing class object
func AsTreeNodes(obj interface{}) ITreeNodes {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	treeNodes := new(TTreeNodes)
	SetObjectInstance(treeNodes, instance)
	return treeNodes
}

// AsTreeNodesEnumerator Convert a pointer object to an existing class object
func AsTreeNodesEnumerator(obj interface{}) ITreeNodesEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	treeNodesEnumerator := new(TTreeNodesEnumerator)
	SetObjectInstance(treeNodesEnumerator, instance)
	return treeNodesEnumerator
}

// AsTreeView Convert a pointer object to an existing class object
func AsTreeView(obj interface{}) ITreeView {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	treeView := new(TTreeView)
	SetObjectInstance(treeView, instance)
	return treeView
}

// AsUpDown Convert a pointer object to an existing class object
func AsUpDown(obj interface{}) IUpDown {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	upDown := new(TUpDown)
	SetObjectInstance(upDown, instance)
	return upDown
}

// AsValueListEditor Convert a pointer object to an existing class object
func AsValueListEditor(obj interface{}) IValueListEditor {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	valueListEditor := new(TValueListEditor)
	SetObjectInstance(valueListEditor, instance)
	return valueListEditor
}

// AsValueListStrings Convert a pointer object to an existing class object
func AsValueListStrings(obj interface{}) IValueListStrings {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	valueListStrings := new(TValueListStrings)
	SetObjectInstance(valueListStrings, instance)
	return valueListStrings
}

// AsWinControl Convert a pointer object to an existing class object
func AsWinControl(obj interface{}) IWinControl {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	winControl := new(TWinControl)
	SetObjectInstance(winControl, instance)
	return winControl
}

// AsWinControlEnumerator Convert a pointer object to an existing class object
func AsWinControlEnumerator(obj interface{}) IWinControlEnumerator {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	winControlEnumerator := new(TWinControlEnumerator)
	SetObjectInstance(winControlEnumerator, instance)
	return winControlEnumerator
}

// AsWindowMagnetOptions Convert a pointer object to an existing class object
func AsWindowMagnetOptions(obj interface{}) IWindowMagnetOptions {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	windowMagnetOptions := new(TWindowMagnetOptions)
	SetObjectInstance(windowMagnetOptions, instance)
	return windowMagnetOptions
}

// AsXButton Convert a pointer object to an existing class object
func AsXButton(obj interface{}) IXButton {
	instance := GetInstance(obj)
	if instance == nil {
		return nil
	}
	xButton := new(TXButton)
	SetObjectInstance(xButton, instance)
	return xButton
}
