//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !cgo

package gtk3

import (
	"github.com/energye/energy/v3/platform/linux/gtk3/nocgo"
	"github.com/energye/energy/v3/platform/linux/types"
	"unsafe"
)

func Init(args *[]string) {
	nocgo.Init(args)
}

func Main() {
	nocgo.Main()
}

func MainQuit() {
	nocgo.MainQuit()
}

func NewWindow(t types.WindowType) (types.IWindow, error) {
	return nocgo.NewWindow(t)
}

func NewBox(orientation types.Orientation, spacing int) types.IBox {
	return nocgo.NewBox(orientation, spacing)
}

func NewMenuBar() types.IMenuBar {
	return nocgo.NewMenuBar()
}

func NewLayout(hadjustment, vadjustment types.IAdjustment) types.ILayout {
	return nocgo.NewLayout(hadjustment, vadjustment)
}

func NewScrolledWindow(hadjustment, vadjustment types.IAdjustment) types.IScrolledWindow {
	return nocgo.NewScrolledWindow(hadjustment, vadjustment)
}

func NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize float64) types.IAdjustment {
	return nocgo.NewAdjustment(value, lower, upper, stepIncrement, pageIncrement, pageSize)
}

func AsScrolledWindow(ptr unsafe.Pointer) types.IScrolledWindow {
	return nocgo.AsScrolledWindow(ptr)
}

func AsWindow(ptr unsafe.Pointer) types.IWindow {
	return nocgo.AsWindow(ptr)
}

func AsContainer(ptr unsafe.Pointer) types.IContainer {
	return nocgo.AsContainer(ptr)
}

func AsBox(ptr unsafe.Pointer) types.IBox {
	return nocgo.AsBox(ptr)
}

func AsMenuBar(ptr unsafe.Pointer) types.IMenuBar {
	return nocgo.AsMenuBar(ptr)
}

func AsLayout(ptr unsafe.Pointer) types.ILayout {
	return nocgo.AsLayout(ptr)
}

func AsWidget(ptr unsafe.Pointer) types.IWidget {
	return nocgo.AsWidget(ptr)
}

func AsDragContext(ptr unsafe.Pointer) types.IDragContext {
	return nocgo.AsDragContext(ptr)
}

func AsAtom(ptr unsafe.Pointer) types.IAtom {
	return nocgo.AsAtom(ptr)
}

func GdkAtomIntern(atomName string, onlyIfExists bool) types.IAtom {
	return nocgo.GdkAtomIntern(atomName, onlyIfExists)
}

func AsEventButton(ptr unsafe.Pointer) types.IEventButton {
	return nocgo.AsEventButton(ptr)
}

func AsEventCrossing(ptr unsafe.Pointer) types.IEventCrossing {
	return nocgo.AsEventCrossing(ptr)
}

func AsEntry(ptr unsafe.Pointer) types.IEntry {
	return nocgo.AsEntry(ptr)
}

func AsEventKey(p unsafe.Pointer) types.IEventKey {
	return nocgo.AsEventKey(p)
}

func AsContext(ptr unsafe.Pointer) types.IContext {
	return nocgo.AsContext(ptr)
}

func AsSelectionData(ptr unsafe.Pointer) types.ISelectionData {
	return nocgo.AsSelectionData(ptr)
}

func AsEventConfigure(ptr unsafe.Pointer) types.IEventConfigure {
	return nocgo.AsEventConfigure(ptr)
}

func NewEntry() types.IEntry {
	return nocgo.NewEntry()
}

func NewCssProvider() types.ICssProvider {
	return nocgo.NewCssProvider()
}

func AsGdkWindow(ptr unsafe.Pointer) types.IGdkWindow {
	return nocgo.AsGdkWindow(ptr)
}

func SettingsGetDefault() types.ISettings {
	return nocgo.SettingsGetDefault()
}

func AsButton(ptr unsafe.Pointer) types.IButton {
	return nocgo.AsButton(ptr)
}

func NewButton() types.IButton {
	return nocgo.NewButton()
}

func NewButtonWithLabel(label string) types.IButton {
	return nocgo.NewButtonWithLabel(label)
}

func AsLabel(ptr unsafe.Pointer) types.ILabel {
	return nocgo.AsLabel(ptr)
}

func NewLabel(str string) types.ILabel {
	return nocgo.NewLabel(str)
}

func AsImage(ptr unsafe.Pointer) types.IImage {
	return nocgo.AsImage(ptr)
}

func NewImage() types.IImage {
	return nocgo.NewImage()
}

func NewImageFromFile(filename string) types.IImage {
	return nocgo.NewImageFromFile(filename)
}

func NewImageFromIconName(iconName string, size types.IconSize) types.IImage {
	return nocgo.NewImageFromIconName(iconName, size)
}

func AsAdjustment(ptr unsafe.Pointer) types.IAdjustment {
	return nocgo.AsAdjustment(ptr)
}

func AsRange(ptr unsafe.Pointer) types.IRange {
	return nocgo.AsRange(ptr)
}

func AsOverlay(ptr unsafe.Pointer) types.IOverlay {
	return nocgo.AsOverlay(ptr)
}

func NewOverlay() types.IOverlay {
	return nocgo.NewOverlay()
}

func NewStack() types.IStack {
	return nocgo.NewStack()
}

func NewStackSwitcher() types.IStackSwitcher {
	return nocgo.NewStackSwitcher()
}

func AsStack(ptr unsafe.Pointer) types.IStack {
	return nocgo.AsStack(ptr)
}

func AsStackSwitcher(ptr unsafe.Pointer) types.IStackSwitcher {
	return nocgo.AsStackSwitcher(ptr)
}

func NewSwitch() types.ISwitch {
	return nocgo.NewSwitch()
}

func AsSwitch(ptr unsafe.Pointer) types.ISwitch {
	return nocgo.AsSwitch(ptr)
}

func NewInfoBar() types.IInfoBar {
	return nocgo.NewInfoBar()
}

func AsInfoBar(ptr unsafe.Pointer) types.IInfoBar {
	return nocgo.AsInfoBar(ptr)
}

func AsMenu(ptr unsafe.Pointer) types.IMenu {
	return nocgo.AsMenu(ptr)
}

func NewMenu() types.IMenu {
	return nocgo.NewMenu()
}

func AsHeaderBar(ptr unsafe.Pointer) types.IHeaderBar {
	return nocgo.AsHeaderBar(ptr)
}

func NewHeaderBar() types.IHeaderBar {
	return nocgo.NewHeaderBar()
}

func AsEventBox(ptr unsafe.Pointer) types.IEventBox {
	return nocgo.AsEventBox(ptr)
}

func NewEventBox() types.IEventBox {
	return nocgo.NewEventBox()
}

func AsFixed(ptr unsafe.Pointer) types.IFixed {
	return nocgo.AsFixed(ptr)
}

func NewFixed() types.IFixed {
	return nocgo.NewFixed()
}

func AsEntryBuffer(ptr unsafe.Pointer) types.IEntryBuffer {
	return nocgo.AsEntryBuffer(ptr)
}

func NewEntryBuffer(initialChars string) types.IEntryBuffer {
	return nocgo.NewEntryBuffer(initialChars)
}

func AsEntryCompletion(ptr unsafe.Pointer) types.IEntryCompletion {
	return nocgo.AsEntryCompletion(ptr)
}

func NewEntryCompletion() types.IEntryCompletion {
	return nocgo.NewEntryCompletion()
}

func AsScrollbar(ptr unsafe.Pointer) types.IScrollbar {
	return nocgo.AsScrollbar(ptr)
}

func NewScrollbar(orientation types.Orientation, adjustment types.IAdjustment) types.IScrollbar {
	return nocgo.NewScrollbar(orientation, adjustment)
}

// Window functions

func WindowGetDefaultIconName() (string, error) {
	return nocgo.WindowGetDefaultIconName()
}

func WindowSetDefaultIconFromFile(file string) error {
	return nocgo.WindowSetDefaultIconFromFile(file)
}

func WindowSetDefaultIconName(s string) {
	nocgo.WindowSetDefaultIconName(s)
}

func WindowSetAutoStartupNotification(setting bool) {
	nocgo.WindowSetAutoStartupNotification(setting)
}

// CssProvider

func CssProviderGetNamed(name string, variant string) types.ICssProvider {
	return nocgo.CssProviderGetNamed(name, variant)
}

// New components

func NewButtonWithMnemonic(label string) types.IButton {
	return nocgo.NewButtonWithMnemonic(label)
}

func NewDialog() types.IDialog {
	return nocgo.NewDialog()
}

func NewProgressBar() types.IProgressBar {
	return nocgo.NewProgressBar()
}

func NewNotebook() types.INotebook {
	return nocgo.NewNotebook()
}

func NewStatusbar() types.IStatusbar {
	return nocgo.NewStatusbar()
}

func NewMessageDialog(parent types.IWindow, flags types.DialogFlags, mType types.MessageType, buttons types.ButtonsType, message string) types.IMessageDialog {
	return nocgo.MessageDialogNew(parent, flags, mType, buttons, message)
}

func NewCheckButton() types.ICheckButton {
	return nocgo.NewCheckButton()
}

func NewSpinButton(adjustment types.IAdjustment, climbRate float64, digits uint) types.ISpinButton {
	return nocgo.NewSpinButton(adjustment, climbRate, digits)
}

// Text widgets

func NewTextView() types.ITextView {
	return nocgo.NewTextView()
}

func NewTextViewWithBuffer(buffer types.ITextBuffer) types.ITextView {
	var buf *nocgo.TextBuffer
	if buffer != nil {
		buf = buffer.(*nocgo.TextBuffer)
	}
	return nocgo.NewTextViewWithBuffer(buf)
}

func NewTextBuffer() types.ITextBuffer {
	return nocgo.NewTextBuffer()
}

func NewTextTagTable() types.ITextTagTable {
	return nocgo.NewTextTagTable()
}

// TreeView + ListStore

func NewTreeView() types.ITreeView {
	return nocgo.NewTreeView()
}

func NewListStore(columnTypes ...types.Type) types.IListStore {
	return nocgo.NewListStore(columnTypes...)
}

func NewTreeStore(columnTypes ...types.Type) types.ITreeStore {
	return nocgo.NewTreeStore(columnTypes...)
}

func NewTreeViewColumn() types.ITreeViewColumn {
	return nocgo.NewTreeViewColumn()
}

func NewCellRendererText() types.ICellRenderer {
	return nocgo.NewCellRendererText()
}

// FileChooserDialog + ComboBox

func NewFileChooserDialog(title string, parent types.IWindow, action types.FileChooserAction) types.IFileChooserDialog {
	return nocgo.NewFileChooserDialog(title, parent, action)
}

func NewComboBoxText() types.IComboBoxText {
	return nocgo.NewComboBoxText()
}

// Grid + Separator + RadioButton + AboutDialog

func NewGrid() types.IGrid {
	return nocgo.NewGrid()
}

func NewSeparator(orientation types.Orientation) types.ISeparator {
	return nocgo.NewSeparator(orientation)
}

func NewRadioButtonWithLabelFromWidget(radioGroupMember *nocgo.RadioButton, label string) *nocgo.RadioButton {
	return nocgo.NewRadioButtonWithLabelFromWidget(radioGroupMember, label)
}

func NewAboutDialog() types.IAboutDialog {
	return nocgo.NewAboutDialog()
}

// ColorChooserDialog + FontChooserDialog

func NewColorChooserDialog(title string, parent types.IWindow) types.IColorChooserDialog {
	return nocgo.NewColorChooserDialog(title, parent)
}

func NewFontChooserDialog(title string, parent types.IWindow) types.IFontChooserDialog {
	return nocgo.NewFontChooserDialog(title, parent)
}

// MenuItem

func NewMenuItem() types.IMenuItem {
	return nocgo.NewMenuItem()
}

func MenuItemNewWithLabel(label string) types.IMenuItem {
	return nocgo.MenuItemNewWithLabel(label)
}

func SeparatorMenuItemNew() types.IMenuItem {
	return nocgo.SeparatorMenuItemNew()
}

// FileFilter

func NewFileFilter() types.IFileFilter {
	return nocgo.NewFileFilter()
}
