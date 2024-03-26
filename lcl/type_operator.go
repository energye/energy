//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import . "github.com/energye/energy/v2/api"

// Exception Is Exception Class
func (m TIs) Exception() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ExceptionClass()))
	return GoBool(r1)
}

// ATGauge Is TATGauge Class
func (m TIs) ATGauge() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ATGaugeClass()))
	return GoBool(r1)
}

// AVLTree Is TAVLTree Class
func (m TIs) AVLTree() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(AVLTreeClass()))
	return GoBool(r1)
}

// AVLTreeNode Is TAVLTreeNode Class
func (m TIs) AVLTreeNode() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(AVLTreeNodeClass()))
	return GoBool(r1)
}

// AVLTreeNodeEnumerator Is TAVLTreeNodeEnumerator Class
func (m TIs) AVLTreeNodeEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(AVLTreeNodeEnumeratorClass()))
	return GoBool(r1)
}

// Action Is TAction Class
func (m TIs) Action() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ActionClass()))
	return GoBool(r1)
}

// ActionList Is TActionList Class
func (m TIs) ActionList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ActionListClass()))
	return GoBool(r1)
}

// ActionListEnumerator Is TActionListEnumerator Class
func (m TIs) ActionListEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ActionListEnumeratorClass()))
	return GoBool(r1)
}

// AnchorSide Is TAnchorSide Class
func (m TIs) AnchorSide() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(AnchorSideClass()))
	return GoBool(r1)
}

// Application Is TApplication Class
func (m TIs) Application() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ApplicationClass()))
	return GoBool(r1)
}

// BasicAction Is TBasicAction Class
func (m TIs) BasicAction() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(BasicActionClass()))
	return GoBool(r1)
}

// BasicActionLink Is TBasicActionLink Class
func (m TIs) BasicActionLink() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(BasicActionLinkClass()))
	return GoBool(r1)
}

// Bevel Is TBevel Class
func (m TIs) Bevel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(BevelClass()))
	return GoBool(r1)
}

// BitBtn Is TBitBtn Class
func (m TIs) BitBtn() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(BitBtnClass()))
	return GoBool(r1)
}

// Bitmap Is TBitmap Class
func (m TIs) Bitmap() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(BitmapClass()))
	return GoBool(r1)
}

// BoundLabel Is TBoundLabel Class
func (m TIs) BoundLabel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(BoundLabelClass()))
	return GoBool(r1)
}

// Brush Is TBrush Class
func (m TIs) Brush() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(BrushClass()))
	return GoBool(r1)
}

// Button Is TButton Class
func (m TIs) Button() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ButtonClass()))
	return GoBool(r1)
}

// ButtonControl Is TButtonControl Class
func (m TIs) ButtonControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ButtonControlClass()))
	return GoBool(r1)
}

// Calendar Is TCalendar Class
func (m TIs) Calendar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CalendarClass()))
	return GoBool(r1)
}

// Canvas Is TCanvas Class
func (m TIs) Canvas() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CanvasClass()))
	return GoBool(r1)
}

// ChangeLink Is TChangeLink Class
func (m TIs) ChangeLink() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ChangeLinkClass()))
	return GoBool(r1)
}

// CheckBox Is TCheckBox Class
func (m TIs) CheckBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CheckBoxClass()))
	return GoBool(r1)
}

// CheckComboBox Is TCheckComboBox Class
func (m TIs) CheckComboBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CheckComboBoxClass()))
	return GoBool(r1)
}

// CheckGroup Is TCheckGroup Class
func (m TIs) CheckGroup() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CheckGroupClass()))
	return GoBool(r1)
}

// CheckListBox Is TCheckListBox Class
func (m TIs) CheckListBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CheckListBoxClass()))
	return GoBool(r1)
}

// Clipboard Is TClipboard Class
func (m TIs) Clipboard() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ClipboardClass()))
	return GoBool(r1)
}

// Collection Is TCollection Class
func (m TIs) Collection() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CollectionClass()))
	return GoBool(r1)
}

// CollectionEnumerator Is TCollectionEnumerator Class
func (m TIs) CollectionEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CollectionEnumeratorClass()))
	return GoBool(r1)
}

// CollectionItem Is TCollectionItem Class
func (m TIs) CollectionItem() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CollectionItemClass()))
	return GoBool(r1)
}

// ColorBox Is TColorBox Class
func (m TIs) ColorBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ColorBoxClass()))
	return GoBool(r1)
}

// ColorButton Is TColorButton Class
func (m TIs) ColorButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ColorButtonClass()))
	return GoBool(r1)
}

// ColorDialog Is TColorDialog Class
func (m TIs) ColorDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ColorDialogClass()))
	return GoBool(r1)
}

// ColorListBox Is TColorListBox Class
func (m TIs) ColorListBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ColorListBoxClass()))
	return GoBool(r1)
}

// ComboBox Is TComboBox Class
func (m TIs) ComboBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ComboBoxClass()))
	return GoBool(r1)
}

// ComboBoxEx Is TComboBoxEx Class
func (m TIs) ComboBoxEx() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ComboBoxExClass()))
	return GoBool(r1)
}

// ComboExItem Is TComboExItem Class
func (m TIs) ComboExItem() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ComboExItemClass()))
	return GoBool(r1)
}

// ComboExItems Is TComboExItems Class
func (m TIs) ComboExItems() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ComboExItemsClass()))
	return GoBool(r1)
}

// CommonDialog Is TCommonDialog Class
func (m TIs) CommonDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CommonDialogClass()))
	return GoBool(r1)
}

// Component Is TComponent Class
func (m TIs) Component() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ComponentClass()))
	return GoBool(r1)
}

// ComponentEnumerator Is TComponentEnumerator Class
func (m TIs) ComponentEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ComponentEnumeratorClass()))
	return GoBool(r1)
}

// ContainedAction Is TContainedAction Class
func (m TIs) ContainedAction() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ContainedActionClass()))
	return GoBool(r1)
}

// Control Is TControl Class
func (m TIs) Control() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ControlClass()))
	return GoBool(r1)
}

// ControlBorderSpacing Is TControlBorderSpacing Class
func (m TIs) ControlBorderSpacing() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ControlBorderSpacingClass()))
	return GoBool(r1)
}

// ControlChildSizing Is TControlChildSizing Class
func (m TIs) ControlChildSizing() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ControlChildSizingClass()))
	return GoBool(r1)
}

// ControlScrollBar Is TControlScrollBar Class
func (m TIs) ControlScrollBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ControlScrollBarClass()))
	return GoBool(r1)
}

// CoolBand Is TCoolBand Class
func (m TIs) CoolBand() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CoolBandClass()))
	return GoBool(r1)
}

// CoolBands Is TCoolBands Class
func (m TIs) CoolBands() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CoolBandsClass()))
	return GoBool(r1)
}

// CoolBar Is TCoolBar Class
func (m TIs) CoolBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CoolBarClass()))
	return GoBool(r1)
}

// CustomAbstractGroupedEdit Is TCustomAbstractGroupedEdit Class
func (m TIs) CustomAbstractGroupedEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomAbstractGroupedEditClass()))
	return GoBool(r1)
}

// CustomAction Is TCustomAction Class
func (m TIs) CustomAction() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomActionClass()))
	return GoBool(r1)
}

// CustomActionList Is TCustomActionList Class
func (m TIs) CustomActionList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomActionListClass()))
	return GoBool(r1)
}

// CustomApplication Is TCustomApplication Class
func (m TIs) CustomApplication() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomApplicationClass()))
	return GoBool(r1)
}

// CustomBitBtn Is TCustomBitBtn Class
func (m TIs) CustomBitBtn() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomBitBtnClass()))
	return GoBool(r1)
}

// CustomBitmap Is TCustomBitmap Class
func (m TIs) CustomBitmap() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomBitmapClass()))
	return GoBool(r1)
}

// CustomButton Is TCustomButton Class
func (m TIs) CustomButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomButtonClass()))
	return GoBool(r1)
}

// CustomCalendar Is TCustomCalendar Class
func (m TIs) CustomCalendar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomCalendarClass()))
	return GoBool(r1)
}

// CustomCheckBox Is TCustomCheckBox Class
func (m TIs) CustomCheckBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomCheckBoxClass()))
	return GoBool(r1)
}

// CustomCheckCombo Is TCustomCheckCombo Class
func (m TIs) CustomCheckCombo() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomCheckComboClass()))
	return GoBool(r1)
}

// CustomCheckGroup Is TCustomCheckGroup Class
func (m TIs) CustomCheckGroup() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomCheckGroupClass()))
	return GoBool(r1)
}

// CustomCheckListBox Is TCustomCheckListBox Class
func (m TIs) CustomCheckListBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomCheckListBoxClass()))
	return GoBool(r1)
}

// CustomColorBox Is TCustomColorBox Class
func (m TIs) CustomColorBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomColorBoxClass()))
	return GoBool(r1)
}

// CustomColorListBox Is TCustomColorListBox Class
func (m TIs) CustomColorListBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomColorListBoxClass()))
	return GoBool(r1)
}

// CustomComboBox Is TCustomComboBox Class
func (m TIs) CustomComboBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomComboBoxClass()))
	return GoBool(r1)
}

// CustomComboBoxEx Is TCustomComboBoxEx Class
func (m TIs) CustomComboBoxEx() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomComboBoxExClass()))
	return GoBool(r1)
}

// CustomControl Is TCustomControl Class
func (m TIs) CustomControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomControlClass()))
	return GoBool(r1)
}

// CustomCoolBar Is TCustomCoolBar Class
func (m TIs) CustomCoolBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomCoolBarClass()))
	return GoBool(r1)
}

// CustomDateTimePicker Is TCustomDateTimePicker Class
func (m TIs) CustomDateTimePicker() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomDateTimePickerClass()))
	return GoBool(r1)
}

// CustomDesignControl Is TCustomDesignControl Class
func (m TIs) CustomDesignControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomDesignControlClass()))
	return GoBool(r1)
}

// CustomDrawGrid Is TCustomDrawGrid Class
func (m TIs) CustomDrawGrid() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomDrawGridClass()))
	return GoBool(r1)
}

// CustomEdit Is TCustomEdit Class
func (m TIs) CustomEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomEditClass()))
	return GoBool(r1)
}

// CustomEditButton Is TCustomEditButton Class
func (m TIs) CustomEditButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomEditButtonClass()))
	return GoBool(r1)
}

// CustomFloatSpinEdit Is TCustomFloatSpinEdit Class
func (m TIs) CustomFloatSpinEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomFloatSpinEditClass()))
	return GoBool(r1)
}

// CustomFlowPanel Is TCustomFlowPanel Class
func (m TIs) CustomFlowPanel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomFlowPanelClass()))
	return GoBool(r1)
}

// CustomForm Is TCustomForm Class
func (m TIs) CustomForm() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomFormClass()))
	return GoBool(r1)
}

// CustomFrame Is TCustomFrame Class
func (m TIs) CustomFrame() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomFrameClass()))
	return GoBool(r1)
}

// CustomGrid Is TCustomGrid Class
func (m TIs) CustomGrid() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomGridClass()))
	return GoBool(r1)
}

// CustomGroupBox Is TCustomGroupBox Class
func (m TIs) CustomGroupBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomGroupBoxClass()))
	return GoBool(r1)
}

// CustomHeaderControl Is TCustomHeaderControl Class
func (m TIs) CustomHeaderControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomHeaderControlClass()))
	return GoBool(r1)
}

// CustomIcon Is TCustomIcon Class
func (m TIs) CustomIcon() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomIconClass()))
	return GoBool(r1)
}

// CustomImage Is TCustomImage Class
func (m TIs) CustomImage() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomImageClass()))
	return GoBool(r1)
}

// CustomImageList Is TCustomImageList Class
func (m TIs) CustomImageList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomImageListClass()))
	return GoBool(r1)
}

// CustomImageListResolution Is TCustomImageListResolution Class
func (m TIs) CustomImageListResolution() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomImageListResolutionClass()))
	return GoBool(r1)
}

// CustomImageListResolutionEnumerator Is TCustomImageListResolutionEnumerator Class
func (m TIs) CustomImageListResolutionEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomImageListResolutionEnumeratorClass()))
	return GoBool(r1)
}

// CustomIniFile Is TCustomIniFile Class
func (m TIs) CustomIniFile() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomIniFileClass()))
	return GoBool(r1)
}

// CustomLabel Is TCustomLabel Class
func (m TIs) CustomLabel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomLabelClass()))
	return GoBool(r1)
}

// CustomLabeledEdit Is TCustomLabeledEdit Class
func (m TIs) CustomLabeledEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomLabeledEditClass()))
	return GoBool(r1)
}

// CustomListBox Is TCustomListBox Class
func (m TIs) CustomListBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomListBoxClass()))
	return GoBool(r1)
}

// CustomListView Is TCustomListView Class
func (m TIs) CustomListView() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomListViewClass()))
	return GoBool(r1)
}

// CustomMaskEdit Is TCustomMaskEdit Class
func (m TIs) CustomMaskEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomMaskEditClass()))
	return GoBool(r1)
}

// CustomMemo Is TCustomMemo Class
func (m TIs) CustomMemo() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomMemoClass()))
	return GoBool(r1)
}

// CustomMemoryStream Is TCustomMemoryStream Class
func (m TIs) CustomMemoryStream() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomMemoryStreamClass()))
	return GoBool(r1)
}

// CustomPage Is TCustomPage Class
func (m TIs) CustomPage() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomPageClass()))
	return GoBool(r1)
}

// CustomPanel Is TCustomPanel Class
func (m TIs) CustomPanel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomPanelClass()))
	return GoBool(r1)
}

// CustomPrintDialog Is TCustomPrintDialog Class
func (m TIs) CustomPrintDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomPrintDialogClass()))
	return GoBool(r1)
}

// CustomPrinterSetupDialog Is TCustomPrinterSetupDialog Class
func (m TIs) CustomPrinterSetupDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomPrinterSetupDialogClass()))
	return GoBool(r1)
}

// CustomProgressBar Is TCustomProgressBar Class
func (m TIs) CustomProgressBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomProgressBarClass()))
	return GoBool(r1)
}

// CustomRadioGroup Is TCustomRadioGroup Class
func (m TIs) CustomRadioGroup() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomRadioGroupClass()))
	return GoBool(r1)
}

// CustomRichMemo Is TCustomRichMemo Class
func (m TIs) CustomRichMemo() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomRichMemoClass()))
	return GoBool(r1)
}

// CustomScrollBar Is TCustomScrollBar Class
func (m TIs) CustomScrollBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomScrollBarClass()))
	return GoBool(r1)
}

// CustomSpeedButton Is TCustomSpeedButton Class
func (m TIs) CustomSpeedButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomSpeedButtonClass()))
	return GoBool(r1)
}

// CustomSpinEdit Is TCustomSpinEdit Class
func (m TIs) CustomSpinEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomSpinEditClass()))
	return GoBool(r1)
}

// CustomSplitter Is TCustomSplitter Class
func (m TIs) CustomSplitter() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomSplitterClass()))
	return GoBool(r1)
}

// CustomStaticText Is TCustomStaticText Class
func (m TIs) CustomStaticText() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomStaticTextClass()))
	return GoBool(r1)
}

// CustomStringGrid Is TCustomStringGrid Class
func (m TIs) CustomStringGrid() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomStringGridClass()))
	return GoBool(r1)
}

// CustomTabControl Is TCustomTabControl Class
func (m TIs) CustomTabControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomTabControlClass()))
	return GoBool(r1)
}

// CustomTaskDialog Is TCustomTaskDialog Class
func (m TIs) CustomTaskDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomTaskDialogClass()))
	return GoBool(r1)
}

// CustomTimer Is TCustomTimer Class
func (m TIs) CustomTimer() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomTimerClass()))
	return GoBool(r1)
}

// CustomTrackBar Is TCustomTrackBar Class
func (m TIs) CustomTrackBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomTrackBarClass()))
	return GoBool(r1)
}

// CustomTrayIcon Is TCustomTrayIcon Class
func (m TIs) CustomTrayIcon() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomTrayIconClass()))
	return GoBool(r1)
}

// CustomTreeView Is TCustomTreeView Class
func (m TIs) CustomTreeView() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomTreeViewClass()))
	return GoBool(r1)
}

// CustomUpDown Is TCustomUpDown Class
func (m TIs) CustomUpDown() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(CustomUpDownClass()))
	return GoBool(r1)
}

// DataModule Is TDataModule Class
func (m TIs) DataModule() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DataModuleClass()))
	return GoBool(r1)
}

// DateTimePicker Is TDateTimePicker Class
func (m TIs) DateTimePicker() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DateTimePickerClass()))
	return GoBool(r1)
}

// DirectoryEdit Is TDirectoryEdit Class
func (m TIs) DirectoryEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DirectoryEditClass()))
	return GoBool(r1)
}

// DockManager Is TDockManager Class
func (m TIs) DockManager() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DockManagerClass()))
	return GoBool(r1)
}

// DockTree Is TDockTree Class
func (m TIs) DockTree() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DockTreeClass()))
	return GoBool(r1)
}

// DockZone Is TDockZone Class
func (m TIs) DockZone() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DockZoneClass()))
	return GoBool(r1)
}

// DragDockObject Is TDragDockObject Class
func (m TIs) DragDockObject() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DragDockObjectClass()))
	return GoBool(r1)
}

// DragImageList Is TDragImageList Class
func (m TIs) DragImageList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DragImageListClass()))
	return GoBool(r1)
}

// DragImageListResolution Is TDragImageListResolution Class
func (m TIs) DragImageListResolution() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DragImageListResolutionClass()))
	return GoBool(r1)
}

// DragObject Is TDragObject Class
func (m TIs) DragObject() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DragObjectClass()))
	return GoBool(r1)
}

// DrawGrid Is TDrawGrid Class
func (m TIs) DrawGrid() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(DrawGridClass()))
	return GoBool(r1)
}

// EbEdit Is TEbEdit Class
func (m TIs) EbEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(EbEditClass()))
	return GoBool(r1)
}

// Edit Is TEdit Class
func (m TIs) Edit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(EditClass()))
	return GoBool(r1)
}

// EditButton Is TEditButton Class
func (m TIs) EditButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(EditButtonClass()))
	return GoBool(r1)
}

// FPCanvasHelper Is TFPCanvasHelper Class
func (m TIs) FPCanvasHelper() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCanvasHelperClass()))
	return GoBool(r1)
}

// FPCustomBrush Is TFPCustomBrush Class
func (m TIs) FPCustomBrush() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomBrushClass()))
	return GoBool(r1)
}

// FPCustomCanvas Is TFPCustomCanvas Class
func (m TIs) FPCustomCanvas() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomCanvasClass()))
	return GoBool(r1)
}

// FPCustomFont Is TFPCustomFont Class
func (m TIs) FPCustomFont() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomFontClass()))
	return GoBool(r1)
}

// FPCustomImage Is TFPCustomImage Class
func (m TIs) FPCustomImage() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomImageClass()))
	return GoBool(r1)
}

// FPCustomImageHandler Is TFPCustomImageHandler Class
func (m TIs) FPCustomImageHandler() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomImageHandlerClass()))
	return GoBool(r1)
}

// FPCustomImageReader Is TFPCustomImageReader Class
func (m TIs) FPCustomImageReader() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomImageReaderClass()))
	return GoBool(r1)
}

// FPCustomImageWriter Is TFPCustomImageWriter Class
func (m TIs) FPCustomImageWriter() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomImageWriterClass()))
	return GoBool(r1)
}

// FPCustomPen Is TFPCustomPen Class
func (m TIs) FPCustomPen() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomPenClass()))
	return GoBool(r1)
}

// FPCustomRegion Is TFPCustomRegion Class
func (m TIs) FPCustomRegion() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPCustomRegionClass()))
	return GoBool(r1)
}

// FPImageBitmap Is TFPImageBitmap Class
func (m TIs) FPImageBitmap() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPImageBitmapClass()))
	return GoBool(r1)
}

// FPList Is TFPList Class
func (m TIs) FPList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPListClass()))
	return GoBool(r1)
}

// FPListEnumerator Is TFPListEnumerator Class
func (m TIs) FPListEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPListEnumeratorClass()))
	return GoBool(r1)
}

// FPMemoryImage Is TFPMemoryImage Class
func (m TIs) FPMemoryImage() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPMemoryImageClass()))
	return GoBool(r1)
}

// FPPalette Is TFPPalette Class
func (m TIs) FPPalette() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPPaletteClass()))
	return GoBool(r1)
}

// FPRectRegion Is TFPRectRegion Class
func (m TIs) FPRectRegion() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FPRectRegionClass()))
	return GoBool(r1)
}

// FileDialog Is TFileDialog Class
func (m TIs) FileDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FileDialogClass()))
	return GoBool(r1)
}

// FindDialog Is TFindDialog Class
func (m TIs) FindDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FindDialogClass()))
	return GoBool(r1)
}

// FloatSpinEdit Is TFloatSpinEdit Class
func (m TIs) FloatSpinEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FloatSpinEditClass()))
	return GoBool(r1)
}

// FlowPanel Is TFlowPanel Class
func (m TIs) FlowPanel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FlowPanelClass()))
	return GoBool(r1)
}

// FlowPanelControl Is TFlowPanelControl Class
func (m TIs) FlowPanelControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FlowPanelControlClass()))
	return GoBool(r1)
}

// FlowPanelControlList Is TFlowPanelControlList Class
func (m TIs) FlowPanelControlList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FlowPanelControlListClass()))
	return GoBool(r1)
}

// Font Is TFont Class
func (m TIs) Font() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FontClass()))
	return GoBool(r1)
}

// FontDialog Is TFontDialog Class
func (m TIs) FontDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FontDialogClass()))
	return GoBool(r1)
}

// Form Is TForm Class
func (m TIs) Form() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FormClass()))
	return GoBool(r1)
}

// Frame Is TFrame Class
func (m TIs) Frame() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(FrameClass()))
	return GoBool(r1)
}

// GEEdit Is TGEEdit Class
func (m TIs) GEEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GEEditClass()))
	return GoBool(r1)
}

// GIFImage Is TGIFImage Class
func (m TIs) GIFImage() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GIFImageClass()))
	return GoBool(r1)
}

// Gauge Is TGauge Class
func (m TIs) Gauge() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GaugeClass()))
	return GoBool(r1)
}

// Graphic Is TGraphic Class
func (m TIs) Graphic() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GraphicClass()))
	return GoBool(r1)
}

// GraphicControl Is TGraphicControl Class
func (m TIs) GraphicControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GraphicControlClass()))
	return GoBool(r1)
}

// GraphicsObject Is TGraphicsObject Class
func (m TIs) GraphicsObject() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GraphicsObjectClass()))
	return GoBool(r1)
}

// GridColumn Is TGridColumn Class
func (m TIs) GridColumn() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GridColumnClass()))
	return GoBool(r1)
}

// GridColumnTitle Is TGridColumnTitle Class
func (m TIs) GridColumnTitle() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GridColumnTitleClass()))
	return GoBool(r1)
}

// GridColumns Is TGridColumns Class
func (m TIs) GridColumns() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GridColumnsClass()))
	return GoBool(r1)
}

// GroupBox Is TGroupBox Class
func (m TIs) GroupBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(GroupBoxClass()))
	return GoBool(r1)
}

// HeaderControl Is THeaderControl Class
func (m TIs) HeaderControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(HeaderControlClass()))
	return GoBool(r1)
}

// HeaderSection Is THeaderSection Class
func (m TIs) HeaderSection() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(HeaderSectionClass()))
	return GoBool(r1)
}

// HeaderSections Is THeaderSections Class
func (m TIs) HeaderSections() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(HeaderSectionsClass()))
	return GoBool(r1)
}

// IDesigner Is TIDesigner Class
func (m TIs) IDesigner() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(IDesignerClass()))
	return GoBool(r1)
}

// Icon Is TIcon Class
func (m TIs) Icon() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(IconClass()))
	return GoBool(r1)
}

// IconOptions Is TIconOptions Class
func (m TIs) IconOptions() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(IconOptionsClass()))
	return GoBool(r1)
}

// Image Is TImage Class
func (m TIs) Image() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ImageClass()))
	return GoBool(r1)
}

// ImageButton Is TImageButton Class
func (m TIs) ImageButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ImageButtonClass()))
	return GoBool(r1)
}

// ImageList Is TImageList Class
func (m TIs) ImageList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ImageListClass()))
	return GoBool(r1)
}

// IniFile Is TIniFile Class
func (m TIs) IniFile() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(IniFileClass()))
	return GoBool(r1)
}

// ItemProp Is TItemProp Class
func (m TIs) ItemProp() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ItemPropClass()))
	return GoBool(r1)
}

// JPEGImage Is TJPEGImage Class
func (m TIs) JPEGImage() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(JPEGImageClass()))
	return GoBool(r1)
}

// LCLComponent Is TLCLComponent Class
func (m TIs) LCLComponent() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LCLComponentClass()))
	return GoBool(r1)
}

// LCLReferenceComponent Is TLCLReferenceComponent Class
func (m TIs) LCLReferenceComponent() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LCLReferenceComponentClass()))
	return GoBool(r1)
}

// Label Is TLabel Class
func (m TIs) Label() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LabelClass()))
	return GoBool(r1)
}

// LabeledEdit Is TLabeledEdit Class
func (m TIs) LabeledEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LabeledEditClass()))
	return GoBool(r1)
}

// LazAccessibleObject Is TLazAccessibleObject Class
func (m TIs) LazAccessibleObject() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LazAccessibleObjectClass()))
	return GoBool(r1)
}

// LazAccessibleObjectEnumerator Is TLazAccessibleObjectEnumerator Class
func (m TIs) LazAccessibleObjectEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LazAccessibleObjectEnumeratorClass()))
	return GoBool(r1)
}

// LazDockForm Is TLazDockForm Class
func (m TIs) LazDockForm() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LazDockFormClass()))
	return GoBool(r1)
}

// LazDockPage Is TLazDockPage Class
func (m TIs) LazDockPage() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LazDockPageClass()))
	return GoBool(r1)
}

// LazDockPages Is TLazDockPages Class
func (m TIs) LazDockPages() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LazDockPagesClass()))
	return GoBool(r1)
}

// LazDockSplitter Is TLazDockSplitter Class
func (m TIs) LazDockSplitter() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LazDockSplitterClass()))
	return GoBool(r1)
}

// LazDockTree Is TLazDockTree Class
func (m TIs) LazDockTree() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LazDockTreeClass()))
	return GoBool(r1)
}

// LazDockZone Is TLazDockZone Class
func (m TIs) LazDockZone() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LazDockZoneClass()))
	return GoBool(r1)
}

// LinkLabel Is TLinkLabel Class
func (m TIs) LinkLabel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(LinkLabelClass()))
	return GoBool(r1)
}

// List Is TList Class
func (m TIs) List() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListClass()))
	return GoBool(r1)
}

// ListBox Is TListBox Class
func (m TIs) ListBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListBoxClass()))
	return GoBool(r1)
}

// ListColumn Is TListColumn Class
func (m TIs) ListColumn() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListColumnClass()))
	return GoBool(r1)
}

// ListColumns Is TListColumns Class
func (m TIs) ListColumns() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListColumnsClass()))
	return GoBool(r1)
}

// ListControlItem Is TListControlItem Class
func (m TIs) ListControlItem() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListControlItemClass()))
	return GoBool(r1)
}

// ListControlItems Is TListControlItems Class
func (m TIs) ListControlItems() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListControlItemsClass()))
	return GoBool(r1)
}

// ListEnumerator Is TListEnumerator Class
func (m TIs) ListEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListEnumeratorClass()))
	return GoBool(r1)
}

// ListItem Is TListItem Class
func (m TIs) ListItem() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListItemClass()))
	return GoBool(r1)
}

// ListItems Is TListItems Class
func (m TIs) ListItems() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListItemsClass()))
	return GoBool(r1)
}

// ListItemsEnumerator Is TListItemsEnumerator Class
func (m TIs) ListItemsEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListItemsEnumeratorClass()))
	return GoBool(r1)
}

// ListView Is TListView Class
func (m TIs) ListView() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ListViewClass()))
	return GoBool(r1)
}

// MainMenu Is TMainMenu Class
func (m TIs) MainMenu() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MainMenuClass()))
	return GoBool(r1)
}

// MaskEdit Is TMaskEdit Class
func (m TIs) MaskEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MaskEditClass()))
	return GoBool(r1)
}

// Memo Is TMemo Class
func (m TIs) Memo() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MemoClass()))
	return GoBool(r1)
}

// MemoScrollBar Is TMemoScrollBar Class
func (m TIs) MemoScrollBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MemoScrollBarClass()))
	return GoBool(r1)
}

// MemoryStream Is TMemoryStream Class
func (m TIs) MemoryStream() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MemoryStreamClass()))
	return GoBool(r1)
}

// Menu Is TMenu Class
func (m TIs) Menu() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MenuClass()))
	return GoBool(r1)
}

// MenuItem Is TMenuItem Class
func (m TIs) MenuItem() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MenuItemClass()))
	return GoBool(r1)
}

// MenuItemEnumerator Is TMenuItemEnumerator Class
func (m TIs) MenuItemEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MenuItemEnumeratorClass()))
	return GoBool(r1)
}

// MergedMenuItems Is TMergedMenuItems Class
func (m TIs) MergedMenuItems() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MergedMenuItemsClass()))
	return GoBool(r1)
}

// Monitor Is TMonitor Class
func (m TIs) Monitor() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MonitorClass()))
	return GoBool(r1)
}

// Mouse Is TMouse Class
func (m TIs) Mouse() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(MouseClass()))
	return GoBool(r1)
}

// Object Is TObject Class
func (m TIs) Object() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ObjectClass()))
	return GoBool(r1)
}

// OpenDialog Is TOpenDialog Class
func (m TIs) OpenDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(OpenDialogClass()))
	return GoBool(r1)
}

// OpenPictureDialog Is TOpenPictureDialog Class
func (m TIs) OpenPictureDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(OpenPictureDialogClass()))
	return GoBool(r1)
}

// OwnedCollection Is TOwnedCollection Class
func (m TIs) OwnedCollection() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(OwnedCollectionClass()))
	return GoBool(r1)
}

// PageControl Is TPageControl Class
func (m TIs) PageControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PageControlClass()))
	return GoBool(r1)
}

// PageSetupDialog Is TPageSetupDialog Class
func (m TIs) PageSetupDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PageSetupDialogClass()))
	return GoBool(r1)
}

// PaintBox Is TPaintBox Class
func (m TIs) PaintBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PaintBoxClass()))
	return GoBool(r1)
}

// Panel Is TPanel Class
func (m TIs) Panel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PanelClass()))
	return GoBool(r1)
}

// PaperSize Is TPaperSize Class
func (m TIs) PaperSize() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PaperSizeClass()))
	return GoBool(r1)
}

// ParaAttributes Is TParaAttributes Class
func (m TIs) ParaAttributes() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ParaAttributesClass()))
	return GoBool(r1)
}

// Pen Is TPen Class
func (m TIs) Pen() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PenClass()))
	return GoBool(r1)
}

// Persistent Is TPersistent Class
func (m TIs) Persistent() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PersistentClass()))
	return GoBool(r1)
}

// Picture Is TPicture Class
func (m TIs) Picture() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PictureClass()))
	return GoBool(r1)
}

// Pixmap Is TPixmap Class
func (m TIs) Pixmap() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PixmapClass()))
	return GoBool(r1)
}

// PopupMenu Is TPopupMenu Class
func (m TIs) PopupMenu() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PopupMenuClass()))
	return GoBool(r1)
}

// PortableAnyMapGraphic Is TPortableAnyMapGraphic Class
func (m TIs) PortableAnyMapGraphic() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PortableAnyMapGraphicClass()))
	return GoBool(r1)
}

// PortableNetworkGraphic Is TPortableNetworkGraphic Class
func (m TIs) PortableNetworkGraphic() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PortableNetworkGraphicClass()))
	return GoBool(r1)
}

// PreviewFileControl Is TPreviewFileControl Class
func (m TIs) PreviewFileControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PreviewFileControlClass()))
	return GoBool(r1)
}

// PreviewFileDialog Is TPreviewFileDialog Class
func (m TIs) PreviewFileDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PreviewFileDialogClass()))
	return GoBool(r1)
}

// PrintDialog Is TPrintDialog Class
func (m TIs) PrintDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PrintDialogClass()))
	return GoBool(r1)
}

// Printer Is TPrinter Class
func (m TIs) Printer() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PrinterClass()))
	return GoBool(r1)
}

// PrinterCanvas Is TPrinterCanvas Class
func (m TIs) PrinterCanvas() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PrinterCanvasClass()))
	return GoBool(r1)
}

// PrinterSetupDialog Is TPrinterSetupDialog Class
func (m TIs) PrinterSetupDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(PrinterSetupDialogClass()))
	return GoBool(r1)
}

// ProgressBar Is TProgressBar Class
func (m TIs) ProgressBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ProgressBarClass()))
	return GoBool(r1)
}

// RadioButton Is TRadioButton Class
func (m TIs) RadioButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(RadioButtonClass()))
	return GoBool(r1)
}

// RadioGroup Is TRadioGroup Class
func (m TIs) RadioGroup() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(RadioGroupClass()))
	return GoBool(r1)
}

// RasterImage Is TRasterImage Class
func (m TIs) RasterImage() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(RasterImageClass()))
	return GoBool(r1)
}

// Region Is TRegion Class
func (m TIs) Region() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(RegionClass()))
	return GoBool(r1)
}

// Registry Is TRegistry Class
func (m TIs) Registry() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(RegistryClass()))
	return GoBool(r1)
}

// ReplaceDialog Is TReplaceDialog Class
func (m TIs) ReplaceDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ReplaceDialogClass()))
	return GoBool(r1)
}

// RichEdit Is TRichEdit Class
func (m TIs) RichEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(RichEditClass()))
	return GoBool(r1)
}

// RichMemo Is TRichMemo Class
func (m TIs) RichMemo() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(RichMemoClass()))
	return GoBool(r1)
}

// RichMemoInline Is TRichMemoInline Class
func (m TIs) RichMemoInline() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(RichMemoInlineClass()))
	return GoBool(r1)
}

// SaveDialog Is TSaveDialog Class
func (m TIs) SaveDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(SaveDialogClass()))
	return GoBool(r1)
}

// SavePictureDialog Is TSavePictureDialog Class
func (m TIs) SavePictureDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(SavePictureDialogClass()))
	return GoBool(r1)
}

// Screen Is TScreen Class
func (m TIs) Screen() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ScreenClass()))
	return GoBool(r1)
}

// ScrollBar Is TScrollBar Class
func (m TIs) ScrollBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ScrollBarClass()))
	return GoBool(r1)
}

// ScrollBox Is TScrollBox Class
func (m TIs) ScrollBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ScrollBoxClass()))
	return GoBool(r1)
}

// ScrollingWinControl Is TScrollingWinControl Class
func (m TIs) ScrollingWinControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ScrollingWinControlClass()))
	return GoBool(r1)
}

// SelectDirectoryDialog Is TSelectDirectoryDialog Class
func (m TIs) SelectDirectoryDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(SelectDirectoryDialogClass()))
	return GoBool(r1)
}

// Shape Is TShape Class
func (m TIs) Shape() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ShapeClass()))
	return GoBool(r1)
}

// ShortCutList Is TShortCutList Class
func (m TIs) ShortCutList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ShortCutListClass()))
	return GoBool(r1)
}

// SizeConstraints Is TSizeConstraints Class
func (m TIs) SizeConstraints() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(SizeConstraintsClass()))
	return GoBool(r1)
}

// SpeedButton Is TSpeedButton Class
func (m TIs) SpeedButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(SpeedButtonClass()))
	return GoBool(r1)
}

// SpinEdit Is TSpinEdit Class
func (m TIs) SpinEdit() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(SpinEditClass()))
	return GoBool(r1)
}

// Splitter Is TSplitter Class
func (m TIs) Splitter() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(SplitterClass()))
	return GoBool(r1)
}

// StaticText Is TStaticText Class
func (m TIs) StaticText() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StaticTextClass()))
	return GoBool(r1)
}

// StatusBar Is TStatusBar Class
func (m TIs) StatusBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StatusBarClass()))
	return GoBool(r1)
}

// StatusPanel Is TStatusPanel Class
func (m TIs) StatusPanel() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StatusPanelClass()))
	return GoBool(r1)
}

// StatusPanels Is TStatusPanels Class
func (m TIs) StatusPanels() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StatusPanelsClass()))
	return GoBool(r1)
}

// Stream Is TStream Class
func (m TIs) Stream() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StreamClass()))
	return GoBool(r1)
}

// StringGrid Is TStringGrid Class
func (m TIs) StringGrid() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StringGridClass()))
	return GoBool(r1)
}

// StringList Is TStringList Class
func (m TIs) StringList() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StringListClass()))
	return GoBool(r1)
}

// Strings Is TStrings Class
func (m TIs) Strings() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StringsClass()))
	return GoBool(r1)
}

// StringsEnumerator Is TStringsEnumerator Class
func (m TIs) StringsEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(StringsEnumeratorClass()))
	return GoBool(r1)
}

// TabSheet Is TTabSheet Class
func (m TIs) TabSheet() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TabSheetClass()))
	return GoBool(r1)
}

// TaskDialog Is TTaskDialog Class
func (m TIs) TaskDialog() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TaskDialogClass()))
	return GoBool(r1)
}

// TaskDialogBaseButtonItem Is TTaskDialogBaseButtonItem Class
func (m TIs) TaskDialogBaseButtonItem() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TaskDialogBaseButtonItemClass()))
	return GoBool(r1)
}

// TaskDialogButtonItem Is TTaskDialogButtonItem Class
func (m TIs) TaskDialogButtonItem() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TaskDialogButtonItemClass()))
	return GoBool(r1)
}

// TaskDialogButtons Is TTaskDialogButtons Class
func (m TIs) TaskDialogButtons() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TaskDialogButtonsClass()))
	return GoBool(r1)
}

// TaskDialogButtonsEnumerator Is TTaskDialogButtonsEnumerator Class
func (m TIs) TaskDialogButtonsEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TaskDialogButtonsEnumeratorClass()))
	return GoBool(r1)
}

// TaskDialogRadioButtonItem Is TTaskDialogRadioButtonItem Class
func (m TIs) TaskDialogRadioButtonItem() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TaskDialogRadioButtonItemClass()))
	return GoBool(r1)
}

// TextAttributes Is TTextAttributes Class
func (m TIs) TextAttributes() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TextAttributesClass()))
	return GoBool(r1)
}

// ThemeServices Is TThemeServices Class
func (m TIs) ThemeServices() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ThemeServicesClass()))
	return GoBool(r1)
}

// Timer Is TTimer Class
func (m TIs) Timer() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TimerClass()))
	return GoBool(r1)
}

// ToggleBox Is TToggleBox Class
func (m TIs) ToggleBox() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ToggleBoxClass()))
	return GoBool(r1)
}

// ToolBar Is TToolBar Class
func (m TIs) ToolBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ToolBarClass()))
	return GoBool(r1)
}

// ToolBarEnumerator Is TToolBarEnumerator Class
func (m TIs) ToolBarEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ToolBarEnumeratorClass()))
	return GoBool(r1)
}

// ToolButton Is TToolButton Class
func (m TIs) ToolButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ToolButtonClass()))
	return GoBool(r1)
}

// ToolWindow Is TToolWindow Class
func (m TIs) ToolWindow() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ToolWindowClass()))
	return GoBool(r1)
}

// TrackBar Is TTrackBar Class
func (m TIs) TrackBar() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TrackBarClass()))
	return GoBool(r1)
}

// TrayIcon Is TTrayIcon Class
func (m TIs) TrayIcon() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TrayIconClass()))
	return GoBool(r1)
}

// TreeNode Is TTreeNode Class
func (m TIs) TreeNode() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TreeNodeClass()))
	return GoBool(r1)
}

// TreeNodes Is TTreeNodes Class
func (m TIs) TreeNodes() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TreeNodesClass()))
	return GoBool(r1)
}

// TreeNodesEnumerator Is TTreeNodesEnumerator Class
func (m TIs) TreeNodesEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TreeNodesEnumeratorClass()))
	return GoBool(r1)
}

// TreeView Is TTreeView Class
func (m TIs) TreeView() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(TreeViewClass()))
	return GoBool(r1)
}

// UpDown Is TUpDown Class
func (m TIs) UpDown() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(UpDownClass()))
	return GoBool(r1)
}

// ValueListEditor Is TValueListEditor Class
func (m TIs) ValueListEditor() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ValueListEditorClass()))
	return GoBool(r1)
}

// ValueListStrings Is TValueListStrings Class
func (m TIs) ValueListStrings() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(ValueListStringsClass()))
	return GoBool(r1)
}

// WinControl Is TWinControl Class
func (m TIs) WinControl() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(WinControlClass()))
	return GoBool(r1)
}

// WinControlEnumerator Is TWinControlEnumerator Class
func (m TIs) WinControlEnumerator() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(WinControlEnumeratorClass()))
	return GoBool(r1)
}

// WindowMagnetOptions Is TWindowMagnetOptions Class
func (m TIs) WindowMagnetOptions() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(WindowMagnetOptionsClass()))
	return GoBool(r1)
}

// XButton Is TXButton Class
func (m TIs) XButton() bool {
	r1 := LCL().SysCallN(3724, uintptr(m), uintptr(XButtonClass()))
	return GoBool(r1)
}
