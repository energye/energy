//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build cgo

package gtk3

import (
	"github.com/energye/energy/v3/platform/linux/gtk3/cgo"
	"github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

func Init(args *[]string) {
	cgo.Init(args)
}

func Main() {
	cgo.Main()
}

func MainQuit() {
	cgo.MainQuit()
}

func NewWindow(t types.WindowType) (types.IWindow, error) {
	return cgo.NewWindow(t)
}

func NewBox(orientation types.Orientation, spacing int) types.IBox {
	return cgo.NewBox(orientation, spacing)
}

func NewMenuBar() types.IMenuBar {
	return cgo.NewMenuBar()
}

func NewLayout(hadjustment, vadjustment types.IAdjustment) types.ILayout {
	var hadj, vadj *cgo.Adjustment
	if hadjustment != nil {
		hadj = hadjustment.(*cgo.Adjustment)
	}
	if vadjustment != nil {
		vadj = vadjustment.(*cgo.Adjustment)
	}
	return cgo.NewLayout(hadj, vadj)
}

func NewScrolledWindow(hadjustment, vadjustment types.IAdjustment) types.IScrolledWindow {
	var hadj, vadj *cgo.Adjustment
	if hadjustment != nil {
		hadj = hadjustment.(*cgo.Adjustment)
	}
	if vadjustment != nil {
		vadj = vadjustment.(*cgo.Adjustment)
	}
	return cgo.NewScrolledWindow(hadj, vadj)
}

func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) types.IAdjustment {
	adjustment, _ := cgo.NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize)
	return adjustment
}

func AsScrolledWindow(ptr unsafe.Pointer) types.IScrolledWindow {
	return cgo.AsScrolledWindow(ptr)
}

func AsWindow(ptr unsafe.Pointer) types.IWindow {
	return cgo.AsWindow(ptr)
}
func AsContainer(ptr unsafe.Pointer) types.IContainer {
	return cgo.AsContainer(ptr)
}

func AsBox(ptr unsafe.Pointer) types.IBox {
	return cgo.AsBox(ptr)
}

func AsMenuBar(ptr unsafe.Pointer) types.IMenuBar {
	return cgo.AsMenuBar(ptr)
}

func AsWidget(ptr unsafe.Pointer) types.IWidget {
	return cgo.AsWidget(ptr)
}

func AsLayout(ptr unsafe.Pointer) types.ILayout {
	return cgo.AsLayout(ptr)
}

func AsDragContext(ptr unsafe.Pointer) types.IDragContext {
	return cgo.AsDragContext(ptr)
}

func AsAtom(ptr unsafe.Pointer) types.IAtom {
	return cgo.AsAtom(ptr)
}

func GdkAtomIntern(atomName string, onlyIfExists bool) types.IAtom {
	return cgo.GdkAtomIntern(atomName, onlyIfExists)
}

func AsEventButton(ptr unsafe.Pointer) types.IEventButton {
	return cgo.AsEventButton(ptr)
}

func AsEventCrossing(ptr unsafe.Pointer) types.IEventCrossing {
	return cgo.AsEventCrossing(ptr)
}

func AsEntry(ptr unsafe.Pointer) types.IEntry {
	return cgo.AsEntry(ptr)
}

func AsEventKey(p unsafe.Pointer) types.IEventKey {
	return cgo.AsEventKey(p)
}

func AsContext(ptr unsafe.Pointer) types.IContext {
	return cgo.AsContext(ptr)
}

func AsSelectionData(ptr unsafe.Pointer) types.ISelectionData {
	return cgo.AsSelectionData(ptr)
}

func AsEventConfigure(ptr unsafe.Pointer) types.IEventConfigure {
	return cgo.AsEventConfigure(ptr)
}

func NewEntry() types.IEntry {
	return cgo.NewEntry()
}

func NewCssProvider() types.ICssProvider {
	return cgo.NewCssProvider()
}

func AsGdkWindow(ptr unsafe.Pointer) types.IGdkWindow {
	return cgo.AsGdkWindow(ptr)
}

func SettingsGetDefault() types.ISettings {
	return cgo.SettingsGetDefault()
}

func AsButton(ptr unsafe.Pointer) types.IButton {
	return cgo.AsButton(ptr)
}

func NewButton() types.IButton {
	return cgo.NewButton()
}

func NewButtonWithLabel(label string) types.IButton {
	return cgo.NewButtonWithLabel(label)
}

func AsLabel(ptr unsafe.Pointer) types.ILabel {
	return cgo.AsLabel(ptr)
}

func NewLabel(str string) types.ILabel {
	return cgo.NewLabel(str)
}

func AsImage(ptr unsafe.Pointer) types.IImage {
	return cgo.AsImage(ptr)
}

func NewImage() types.IImage {
	return cgo.NewImage()
}

func NewImageFromFile(filename string) types.IImage {
	return cgo.NewImageFromFile(filename)
}

func NewImageFromIconName(iconName string, size types.IconSize) types.IImage {
	return cgo.NewImageFromIconName(iconName, size)
}

func AsAdjustment(ptr unsafe.Pointer) types.IAdjustment {
	return cgo.AsAdjustment(ptr)
}

func AsRange(ptr unsafe.Pointer) types.IRange {
	return cgo.AsRange(ptr)
}

func AsOverlay(ptr unsafe.Pointer) types.IOverlay {
	return cgo.AsOverlay(ptr)
}

func NewOverlay() types.IOverlay {
	return cgo.NewOverlay()
}

func NewStack() types.IStack {
	return cgo.NewStack()
}

func NewStackSwitcher() types.IStackSwitcher {
	return cgo.NewStackSwitcher()
}

func AsStack(ptr unsafe.Pointer) types.IStack {
	return cgo.AsStack(ptr)
}

func AsStackSwitcher(ptr unsafe.Pointer) types.IStackSwitcher {
	return cgo.AsStackSwitcher(ptr)
}

func NewSwitch() types.ISwitch {
	return cgo.NewSwitch()
}

func AsSwitch(ptr unsafe.Pointer) types.ISwitch {
	return cgo.AsSwitch(ptr)
}

func NewInfoBar() types.IInfoBar {
	return cgo.NewInfoBar()
}

func AsInfoBar(ptr unsafe.Pointer) types.IInfoBar {
	return cgo.AsInfoBar(ptr)
}

func NewMenu() types.IMenu {
	m, _ := cgo.NewMenu()
	return m
}

func AsHeaderBar(ptr unsafe.Pointer) types.IHeaderBar {
	return cgo.AsHeaderBar(ptr)
}

func NewHeaderBar() types.IHeaderBar {
	return cgo.NewHeaderBar()
}

func AsEventBox(ptr unsafe.Pointer) types.IEventBox {
	return cgo.AsEventBox(ptr)
}

func NewEventBox() types.IEventBox {
	return cgo.NewEventBox()
}

func AsFixed(ptr unsafe.Pointer) types.IFixed {
	return cgo.AsFixed(ptr)
}

func NewFixed() types.IFixed {
	return cgo.NewFixed()
}

func AsEntryBuffer(ptr unsafe.Pointer) types.IEntryBuffer {
	return cgo.AsEntryBuffer(ptr)
}

func NewEntryBuffer(initialChars string) types.IEntryBuffer {
	return cgo.NewEntryBuffer(initialChars, -1)
}

func AsEntryCompletion(ptr unsafe.Pointer) types.IEntryCompletion {
	return cgo.AsEntryCompletion(ptr)
}

func NewEntryCompletion() types.IEntryCompletion {
	return cgo.NewEntryCompletion()
}

func AsScrollbar(ptr unsafe.Pointer) types.IScrollbar {
	return cgo.AsScrollbar(ptr)
}

func NewScrollbar(orientation types.Orientation, adjustment types.IAdjustment) types.IScrollbar {
	if adjustment == nil {
		return cgo.NewScrollbar(orientation, nil)
	}
	return cgo.NewScrollbar(orientation, adjustment.(*cgo.Adjustment))
}

// Window functions

func WindowGetDefaultIconName() (string, error) {
	return cgo.WindowGetDefaultIconName()
}

func WindowSetDefaultIconFromFile(file string) error {
	return cgo.WindowSetDefaultIconFromFile(file)
}

func WindowSetDefaultIconName(s string) {
	cgo.WindowSetDefaultIconName(s)
}

func WindowSetAutoStartupNotification(setting bool) {
	cgo.WindowSetAutoStartupNotification(setting)
}

// CssProvider

func CssProviderGetNamed(name string, variant string) types.ICssProvider {
	return cgo.CssProviderGetNamed(name, variant)
}

// New components

func NewButtonWithMnemonic(label string) types.IButton {
	return cgo.NewButtonWithMnemonic(label)
}

func NewDialog() types.IDialog {
	return cgo.NewDialog()
}

func NewProgressBar() types.IProgressBar {
	return cgo.NewProgressBar()
}

func NewScale(adjustment types.IAdjustment) types.IScale {
	return cgo.NewHScale(adjustment)
}

func NewHScale(adjustment types.IAdjustment) types.IScale {
	return cgo.NewHScale(adjustment)
}

func NewVScale(adjustment types.IAdjustment) types.IScale {
	return cgo.NewVScale(adjustment)
}

func NewSpinner() types.ISpinner {
	return cgo.NewSpinner()
}

func NewLevelBar() types.ILevelBar {
	return cgo.NewLevelBar()
}

func NewNotebook() types.INotebook {
	return cgo.NewNotebook()
}

func NewStatusbar() types.IStatusbar {
	return cgo.NewStatusbar()
}

func NewMessageDialog(parent types.IWindow, flags types.DialogFlags, mType types.MessageType, buttons types.ButtonsType, message string) types.IMessageDialog {
	return cgo.MessageDialogNew(parent, flags, mType, buttons, message)
}

func NewCheckButton() types.ICheckButton {
	return cgo.NewCheckButton()
}

func NewSpinButton(adjustment types.IAdjustment, climbRate float64, digits uint) types.ISpinButton {
	var adj *cgo.Adjustment
	if adjustment != nil {
		adj = adjustment.(*cgo.Adjustment)
	}
	return cgo.NewSpinButton(adj, climbRate, digits)
}

// Text widgets

func NewTextView() types.ITextView {
	return cgo.NewTextView()
}

func NewTextViewWithBuffer(buffer types.ITextBuffer) types.ITextView {
	var buf *cgo.TextBuffer
	if buffer != nil {
		buf = buffer.(*cgo.TextBuffer)
	}
	return cgo.NewTextViewWithBuffer(buf)
}

func NewTextBuffer() types.ITextBuffer {
	return cgo.NewTextBuffer()
}

func NewTextTagTable() types.ITextTagTable {
	return cgo.NewTextTagTable()
}

// TreeView + ListStore

func NewTreeView() types.ITreeView {
	return cgo.NewTreeView()
}

func NewListStore(columnTypes ...types.Type) types.IListStore {
	return cgo.NewListStore(columnTypes...)
}

func NewTreeStore(columnTypes ...types.Type) types.ITreeStore {
	return cgo.NewTreeStore(columnTypes...)
}

func NewTreeViewColumn() types.ITreeViewColumn {
	return cgo.NewTreeViewColumn()
}

func NewCellRendererText() types.ICellRenderer {
	return cgo.NewCellRendererText()
}

// FileChooserDialog + ComboBox

func NewFileChooserDialog(title string, parent types.IWindow, action types.FileChooserAction) types.IFileChooserDialog {
	return cgo.NewFileChooserDialog(title, parent, action)
}

func NewComboBoxText() types.IComboBoxText {
	return cgo.NewComboBoxText()
}

// Grid + Separator + RadioButton + AboutDialog

func NewGrid() types.IGrid {
	return cgo.NewGrid()
}

func NewSeparator(orientation types.Orientation) types.ISeparator {
	return cgo.NewSeparator(orientation)
}

func NewRadioButtonWithLabelFromWidget(radioGroupMember *cgo.RadioButton, label string) *cgo.RadioButton {
	return cgo.NewRadioButtonWithLabelFromWidget(radioGroupMember, label)
}

func NewAboutDialog() types.IAboutDialog {
	return cgo.NewAboutDialog()
}

// ColorChooserDialog + FontChooserDialog

func NewColorChooserDialog(title string, parent types.IWindow) types.IColorChooserDialog {
	return cgo.NewColorChooserDialog(title, parent)
}

func NewFontChooserDialog(title string, parent types.IWindow) types.IFontChooserDialog {
	return cgo.NewFontChooserDialog(title, parent)
}

// MenuItem

func NewMenuItem() types.IMenuItem {
	item, _ := cgo.NewMenuItem()
	return item
}

func MenuItemNewWithLabel(label string) types.IMenuItem {
	item, _ := cgo.MenuItemNewWithLabel(label)
	if item == nil {
		return nil
	}
	return item
}

func SeparatorMenuItemNew() types.IMenuItem {
	return cgo.SeparatorMenuItemNew()
}

// FileFilter

func NewFileFilter() types.IFileFilter {
	return cgo.NewFileFilter()
}
