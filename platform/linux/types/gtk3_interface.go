//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package types

import (
	"github.com/energye/lcl/types/colors"
	"unsafe"
)

type ISignalHandlerID interface {
	Disconnect()
}

type IObject interface {
	Instance() uintptr
	Ref()
	Unref()
}

type IScreen interface {
	IObject
	GetRGBAVisual() IVisual
	IsComposited() bool
}

type IVisual interface {
	IObject
}

type IGdkWindow interface {
	IObject
	SetDecorations(decorations TGdkWMDecoration)
	WindowGetWidth() int
	WindowGetHeight() int
	GetRootOrigin() (x, y int)
	GetOrigin() (x, y int)
	SetOverrideRedirect(setting bool)
}

type IWidget interface {
	IObject
	GetScreen() IScreen
	SetVisual(visual IVisual)
	SetAppPaintable(paintable bool)
	GetAppPaintable() bool
	GetName() string
	SetName(name string)
	GetAllocation() IRectangle
	SetSizeRequest(width, height int)
	GetSizeRequest() (width, height int)
	GetStyleContext() IStyleContext
	GrabFocus()
	GrabDefault()
	DragGetData(context IDragContext, target IAtom, time uint)
	IsContainer() bool
	Realize()
	Unrealize()
	Map()
	Unmap()
	Show()
	Hide()
	ShowNow()
	ShowAll()
	SetNoShowAll(noShowAll bool)
	GetNoShowAll() bool
	Activate() bool
	IsFocus() bool
	HasFocus() bool
	HasDefault() bool
	HasVisibleFocus() bool
	HasGrab() bool
	IsDrawable() bool
	IsToplevel() bool
	GetSensitive() bool
	IsSensitive() bool
	SetSensitive(sensitive bool)
	GetVisible() bool
	SetVisible(visible bool)
	GetHAlign() Align
	SetHAlign(align Align)
	GetVAlign() Align
	SetVAlign(align Align)
	GetMarginTop() int
	SetMarginTop(margin int)
	GetMarginBottom() int
	SetMarginBottom(margin int)
	GetMarginStart() int
	SetMarginStart(margin int)
	GetMarginEnd() int
	SetMarginEnd(margin int)
	GetHExpand() bool
	SetHExpand(expand bool)
	GetVExpand() bool
	SetVExpand(expand bool)
	GetRealized() bool
	SetRealized(realized bool)
	GetHasWindow() bool
	SetHasWindow(hasWindow bool)
	GetCanFocus() bool
	SetCanFocus(canFocus bool)
	GetCanDefault() bool
	SetCanDefault(canDefault bool)
	GetMapped() bool
	SetMapped(mapped bool)
	GetFocusOnClick() bool
	SetFocusOnClick(focusOnClick bool)
	GetAllocatedWidth() int
	GetAllocatedHeight() int
	SetEvents(events EventMask)
	GetEvents() EventMask
	AddEvents(events EventMask)
	FreezeChildNotify()
	ThawChildNotify()
	SetOpacity(opacity float64)
	GetOpacity() float64
	GetTooltipText() string
	SetTooltipText(text string)
	GetTooltipMarkup() string
	SetTooltipMarkup(markup string)
	QueueDraw()
	SetStateFlags(flags StateFlags, clear bool)
	UnsetStateFlags(flags StateFlags)
	GetStateFlags() StateFlags
	ResetStyle()
	InDestruction() bool
	Destroy()
}

type IContainer interface {
	IWidget
	Add(w IWidget)
	Remove(w IWidget)
	CheckResize()
	GetChildren() IList
	GetFocusChild() IWidget
	SetFocusChild(child IWidget)
	GetBorderWidth() uint
	SetBorderWidth(borderWidth uint)
}

type IBin interface {
	IContainer
}

type IBox interface {
	IContainer
	PackStart(child IWidget, expand, fill bool, padding uint)
	PackEnd(child IWidget, expand, fill bool, padding uint)
	GetHomogeneous() bool
	SetHomogeneous(homogeneous bool)
	GetSpacing() int
	SetSpacing(spacing int)
	ReorderChild(child IWidget, position int)
	SetChildPacking(child IWidget, expand, fill bool, padding uint, packType PackType)
}

type IStyleProvider interface {
	Instance() uintptr
}

type IList interface {
	Instance() uintptr
	Append(data uintptr) IList
	Prepend(data uintptr) IList
	Insert(data uintptr, position int) IList
	Length() uint
	NthDataRaw(n uint) unsafe.Pointer
	Next() IList
	Previous() IList
	First() IList
	Last() IList
	Free()
}

type IStyleContext interface {
	IObject
	AddClass(class_name string)
	RemoveClass(class_name string)
	HasClass(className string) bool
	AddProvider(provider IStyleProvider, prio uint)
	RemoveProvider(provider IStyleProvider)
	Save()
	Restore()
	SetState(flags StateFlags)
}

type IWindow interface {
	IBin
	GetDefaultSize() (width, height int)
	SetDefaultSize(width, height int)
	SetDecorated(setting bool)
	GetDecorated() bool
	SetDeletable(setting bool)
	GetDeletable() bool
	Maximize()
	Unmaximize()
	Fullscreen()
	Unfullscreen()
	Iconify()
	Deiconify()
	Stick()
	Unstick()
	SetTitle(title string)
	GetTitle() string
	SetTitlebar(titlebar IWidget)
	SetResizable(resizable bool)
	GetResizable() bool
	SetModal(modal bool)
	GetModal() bool
	SetGravity(gravity Gravity)
	GetGravity() Gravity
	SetDestroyWithParent(setting bool)
	GetDestroyWithParent() bool
	SetHideTitlebarWhenMaximized(setting bool)
	GetHideTitlebarWhenMaximized() bool
	IsActive() bool
	HasToplevelFocus() bool
	Present()
	PresentWithTime(ts uint32)
	SetKeepAbove(setting bool)
	SetKeepBelow(setting bool)
	SetTypeHint(typeHint WindowTypeHint)
	GetTypeHint() WindowTypeHint
	SetSkipTaskbarHint(setting bool)
	GetSkipTaskbarHint() bool
	SetSkipPagerHint(setting bool)
	GetSkipPagerHint() bool
	SetUrgencyHint(setting bool)
	GetUrgencyHint() bool
	SetAcceptFocus(setting bool)
	GetAcceptFocus() bool
	SetFocusOnMap(setting bool)
	GetFocusOnMap() bool
	SetStartupID(sid string)
	SetRole(s string)
	GetRole() (string, error)
	GetPosition() (int, int)
	GetSize() (width, height int)
	GetIconName() (string, error)
	SetIconName(name string)
	GetWindowType() WindowType
	HasGroup() bool
	Move(x, y int)
	Resize(width, height int)
	GetMnemonicsVisible() bool
	SetMnemonicsVisible(setting bool)
	GetFocusVisible() bool
	SetFocusVisible(setting bool)
	BeginResizeDrag(edge WindowEdge, button ButtonType, rootX, rootY int, timestamp uint32)
	BeginMoveDrag(button ButtonType, rootX, rootY int, timestamp uint32)
	SetOnMap(fn TMapEvent) ISignalHandlerID
	SetOnDraw(fn TDrawEvent) ISignalHandlerID
	SetOnConfigure(fn TConfigureEvent) ISignalHandlerID
	SetOnDestroy(fn TNotifyEvent) ISignalHandlerID
}

type IMenuShell interface {
	IContainer
	Append(child IWidget)
	Prepend(child IWidget)
	Insert(child IWidget, position int)
	Deactivate()
	SelectItem(child IWidget)
	SelectFirst(searchSensitive bool)
	Deselect()
	ActivateItem(child IWidget, forceDeactivate bool)
	Cancel()
	SetTakeFocus(takeFocus bool)
	GetTakeFocus() bool
	GetSelectedItem() (IWidget, error)
	GetParentShell() (IMenuShell, error)
	SetOnDeactivate(fn TNotifyEvent) ISignalHandlerID
	SetOnSelectionDone(fn TNotifyEvent) ISignalHandlerID
}

type IMenuItem interface {
	IWidget
	SetSubmenu(submenu IWidget)
	GetSubmenu() IMenu
	SetOnActivate(fn TNotifyEvent) ISignalHandlerID
}

type IScrolledWindow interface {
	IBin
	SetPolicy(hScrollbarPolicy, vScrollbarPolicy PolicyType)
	SetMinContentWidth(width int)
	GetMinContentWidth() int
	SetMinContentHeight(height int)
	GetMinContentHeight() int
	SetShadowType(type_ ShadowType)
	GetShadowType() ShadowType
}

type IMenuBar interface {
	IMenuShell
}

type ILayout interface {
	IContainer
	Put(w IWidget, x, y int)
	Move(w IWidget, x, y int)
	SetSize(width, height uint)
	GetSize() (width, height uint)
}

type ICssProvider interface {
	IObject
	LoadFromPath(path string) error
	LoadFromData(data string) error
	ToString() string
}

type ISettings interface {
	IObject
	SetOnThemeChanged(fn TThemeChangedEvent) ISignalHandlerID
}

type IRectangle interface {
	GetX() int
	SetX(x int)
	GetY() int
	SetY(y int)
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type IDragContext interface {
	IObject
	ListTargets() IList
	Finish(success bool, del bool, time uint)
	Status(actions DragAction, time uint)
}

type ISelectionData interface {
	Instance() uintptr
	GetLength() int
	GetData() []byte
	SetData(atom TAtom, data []byte)
	GetText() string
	SetText(text string) bool
	SetURIs(uris []string) bool
	GetURIs() []string
	Free()
}

type IAtom interface {
	Name() string
	Atom() TAtom
}

type IEvent interface {
	Instance() uintptr
	Free()
	ScanCode() int
}

type IEventKey interface {
	IEvent
	KeyVal() uint
	HardwareKeyCode() uint16
	Type() EventType
	State() uint
}

type IEventButton interface {
	IEvent
	X() float64
	Y() float64
	XRoot() float64
	YRoot() float64
	Button() ButtonType
	State() uint
	Time() uint32
	Type() EventType
}

type IEventCrossing interface {
	IEvent
	X() float64
	Y() float64
	XRoot() float64
	YRoot() float64
	State() uint
	Time() uint32
	Type() EventType
	Mode() CrossingMode
	Detail() NotifyType
	Focus() bool
}

type IEventConfigure interface {
	IEvent
	X() int
	Y() int
	Width() int
	Height() int
	Type() EventType
}

type IContext interface {
	Instance() uintptr
	Status() Status
	Close()
	Save()
	Restore()
	SetSourceRGB(red, green, blue float64)
	SetSourceRGBA(red, green, blue, alpha float64)
	SetLineWidth(width float64)
	GetLineWidth() float64
	Clip()
	ClipPreserve()
	ResetClip()
	Rectangle(x, y, w, h float64)
	Arc(xc, yc, radius, angle1, angle2 float64)
	LineTo(x, y float64)
	MoveTo(x, y float64)
	Fill()
	FillPreserve()
	ClosePath()
	NewPath()
	Paint()
	PaintWithAlpha(alpha float64)
	Stroke()
	StrokePreserve()
	CopyPage()
	ShowPage()
	PushGroup()
	PopGroupToSource()
}

type IEntry interface {
	IWidget
	SetText(text string)
	GetText() string
	GetTextLength() uint16
	SetVisibility(visible bool)
	GetVisibility() bool
	SetMaxLength(max int)
	GetMaxLength() int
	SetHasFrame(setting bool)
	GetHasFrame() bool
	SetWidthChars(nChars int)
	GetWidthChars() int
	SetActivatesDefault(setting bool)
	GetActivatesDefault() bool
	SetPlaceholderText(text string)
	GetPlaceholderText() string
	ProgressPulse()
	ResetIMContext()
	SetOnChanged(fn TTextChangedEvent) ISignalHandlerID
	SetOnCommit(fn TTextCommitEvent) ISignalHandlerID
	SetOnKeyPress(fn TTextKeyEvent) ISignalHandlerID
	SetOnKeyRelease(fn TTextKeyEvent) ISignalHandlerID
}

type IWebkit2 interface {
	IWidget
	OpenDevTools()
	SetBackgroundColor(color *colors.TARGB)
	SetOnDragDataReceived(fn TDragDataReceivedEvent) ISignalHandlerID
	SetOnDragDrop(fn TDragDropEvent) ISignalHandlerID
	SetOnDragMotion(fn TDragMotionEvent) ISignalHandlerID
	SetOnDragLeave(fn TDragLeaveEvent) ISignalHandlerID
	SetOnDragDataDelete(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID
	SetOnDragBegin(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID
	SetOnDragEnd(fn TDragDataDeleteOrBeginOrEndEvent) ISignalHandlerID
	SetOnFocusIn(fn TFocusInEvent) ISignalHandlerID
	SetOnFocusOut(fn TFocusOutEvent) ISignalHandlerID
	Undo()
	CanUndo() bool
	Redo()
	CanRedo() bool
	Cut()
	CanCut() bool
	Copy()
	CanCopy() bool
	Paste()
	CanPaste() bool
	Delete()
	CanDelete() bool
	SelectAll()
}

// IButton is a representation of GTK's GtkButton.
type IButton interface {
	IBin
	Clicked()
	SetRelief(newStyle ReliefStyle)
	GetRelief() ReliefStyle
	SetLabel(label string)
	GetLabel() (string, error)
	SetUseUnderline(useUnderline bool)
	GetUseUnderline() bool
	SetOnClick(fn TNotifyEvent) ISignalHandlerID
	SetOnEnter(fn TLeaveEnterNotifyEvent) ISignalHandlerID
	SetOnLeave(fn TLeaveEnterNotifyEvent) ISignalHandlerID
}

// ILabel is a representation of GTK's GtkLabel.
type ILabel interface {
	IWidget
	SetText(str string)
	GetText() (string, error)
	SetMarkup(str string)
	SetJustify(jtype Justification)
	GetJustify() Justification
	SetEllipsize(mode EllipsizeMode)
	GetEllipsize() EllipsizeMode
	SetWidthChars(nChars int)
	GetWidthChars() int
	SetLineWrap(wrap bool)
	GetLineWrap() bool
	SetSelectable(setting bool)
	GetSelectable() bool
	SetLabel(str string)
	GetLabel() string
}

// IImage is a representation of GTK's GtkImage.
type IImage interface {
	IWidget
	Clear()
	SetFromFile(filename string)
	SetFromIconName(iconName string, size IconSize)
	SetPixelSize(pixelSize int)
	GetStorageType() ImageType
	GetIconName() (string, IconSize)
	GetPixelSize() int
}

// IAdjustment is a representation of GTK's GtkAdjustment.
type IAdjustment interface {
	IObject
	GetValue() float64
	SetValue(value float64)
	GetLower() float64
	SetLower(lower float64)
	GetUpper() float64
	SetUpper(upper float64)
	GetPageSize() float64
	SetPageSize(pageSize float64)
	GetStepIncrement() float64
	SetStepIncrement(stepIncrement float64)
}

// IRange is a representation of GTK's GtkRange.
type IRange interface {
	IWidget
	GetAdjustment() IAdjustment
	SetAdjustment(adj IAdjustment)
	GetValue() float64
	SetValue(value float64)
	SetIncrements(step, page float64)
	SetRange(min, max float64)
	GetInverted() bool
	SetInverted(setting bool)
	SetOnValueChanged(fn TValueChangedEvent) ISignalHandlerID
}

// IEditable is a representation of GTK's GtkEditable.
type IEditable interface {
	IWidget
	SelectRegion(startPos, endPos int)
	GetSelectionBounds() (start, end int, ok bool)
	SetEditable(isEditable bool)
	GetEditable() bool
}

// IOverlay is a representation of GTK's GtkOverlay.
type IOverlay interface {
	IContainer
	AddOverlay(widget IWidget)
}

// IMenu is a representation of GTK's GtkMenu.
type IMenu interface {
	IMenuShell
}

// IHeaderBar is a representation of GTK's GtkHeaderBar.
type IHeaderBar interface {
	IContainer
	SetTitle(title string)
	GetTitle() string
	SetShowCloseButton(setting bool)
	GetShowCloseButton() bool
	PackStart(child IWidget)
	PackEnd(child IWidget)
}

// IEventBox is a representation of GTK's GtkEventBox.
type IEventBox interface {
	IBin
	SetOnClick(fn TButtonPressEvent) ISignalHandlerID
	SetOnEnter(fn TLeaveEnterNotifyEvent) ISignalHandlerID
	SetOnLeave(fn TLeaveEnterNotifyEvent) ISignalHandlerID
}

// IStack is a representation of GTK's GtkStack.
type IStack interface {
	IContainer
	AddNamed(child IWidget, name string)
	AddTitled(child IWidget, name, title string)
	SetVisibleChild(child IWidget)
	GetVisibleChild() IWidget
	SetVisibleChildName(name string)
	GetVisibleChildName() string
	SetVisibleChildFull(name string, transition StackTransitionType)
	SetHomogeneous(homogeneous bool)
	GetHomogeneous() bool
	SetTransitionDuration(duration uint)
	GetTransitionDuration() uint
	SetTransitionType(transition StackTransitionType)
	GetTransitionType() StackTransitionType
}

// IStackSwitcher is a representation of GTK's GtkStackSwitcher.
type IStackSwitcher interface {
	IBox
	SetStack(stack IStack)
	GetStack() IStack
}

// ISwitch is a representation of GTK's GtkSwitch.
type ISwitch interface {
	IWidget
	GetActive() bool
	SetActive(isActive bool)
	GetState() bool
	SetState(state bool)
	SetOnActiveNotify(fn TNotifyActiveEvent) ISignalHandlerID
}

// IInfoBar is a representation of GTK's GtkInfoBar.
type IInfoBar interface {
	IBox
	AddActionWidget(child IWidget, responseId int)
	AddButton(buttonText string, responseId int)
	SetResponseSensitive(responseId int, setting bool)
	SetDefaultResponse(responseId int)
	SetMessageType(messageType MessageType)
	GetMessageType() MessageType
	GetActionArea() IWidget
	GetContentArea() IBox
	SetShowCloseButton(setting bool)
	GetShowCloseButton() bool
	SetOnResponse(fn TResponseEvent) ISignalHandlerID
}

// IFixed is a representation of GTK's GtkFixed.
type IFixed interface {
	IContainer
	Put(w IWidget, x, y int)
	Move(w IWidget, x, y int)
}

// IEntryBuffer is a representation of GTK's GtkEntryBuffer.
type IEntryBuffer interface {
	IObject
	GetText() (string, error)
	SetText(text string)
	GetBytes() uint
	GetLength() uint
	SetOnDeletedText(fn TDeletedTextEvent) ISignalHandlerID
}

// IEntryCompletion is a representation of GTK's GtkEntryCompletion.
type IEntryCompletion interface {
	IObject
	SetTextColumn(column int)
	GetTextColumn() int
	SetMinimumKeyLength(length int)
	GetMinimumKeyLength() int
	SetOnMatchSelected(fn TMatchSelectedEvent) ISignalHandlerID
	SetOnActionActivated(fn TActionActivatedEvent) ISignalHandlerID
}

// IScrollbar is a representation of GTK's GtkScrollbar.
type IScrollbar interface {
	IRange
}

// IRequisition is a representation of GTK's GtkRequisition.
type IRequisition interface {
	GetWidth() int
	GetHeight() int
	Free()
}

// IDialog is a representation of GTK's GtkDialog.
type IDialog interface {
	IWindow
	Run() int
	Response(responseId int)
	AddButton(buttonText string, responseId int) IButton
	SetDefaultResponse(responseId int)
	GetContentArea() IBox
	SetOnResponse(fn TResponseEvent) ISignalHandlerID
}

// IProgressBar is a representation of GTK's GtkProgressBar.
type IProgressBar interface {
	IWidget
	SetFraction(fraction float64)
	GetFraction() float64
	Pulse()
	SetText(text string)
	SetShowText(showText bool)
	SetPulseStep(fraction float64)
}

// INotebook is a representation of GTK's GtkNotebook.
type INotebook interface {
	IContainer
	AppendPage(child IWidget, tabLabel IWidget) int
	RemovePage(pageNum int)
	GetCurrentPage() int
	SetCurrentPage(pageNum int)
	GetNPages() int
	SetShowTabs(showTabs bool)
	SetShowBorder(showBorder bool)
	SetTabPos(pos PositionType)
	SetScrollable(scrollable bool)
	GetScrollable() bool
	SetOnSwitchPage(fn TSwitchPageEvent) ISignalHandlerID
	SetOnPageAdded(fn TPageEvent) ISignalHandlerID
	SetOnPageRemoved(fn TPageEvent) ISignalHandlerID
}

// IStatusbar is a representation of GTK's GtkStatusbar.
type IStatusbar interface {
	IBox
	GetContextId(contextDescription string) uint
	Push(contextId uint, text string) uint
	Pop(contextId uint)
	RemoveAll(contextId uint)
}

// IMessageDialog is a representation of GTK's GtkMessageDialog.
type IMessageDialog interface {
	IDialog
	FormatSecondaryText(message string)
}

// ISpinButton is a representation of GTK's GtkSpinButton.
type ISpinButton interface {
	IEntry
	GetValue() float64
	SetValue(value float64)
	SetRange(min, max float64)
	SetIncrements(step, page float64)
	SetDigits(digits uint)
	SetOnValueChanged(fn TValueChangedEvent) ISignalHandlerID
}

// ICheckButton is a representation of GTK's GtkCheckButton.
type ICheckButton interface {
	IButton
	GetActive() bool
	SetActive(isActive bool)
	SetOnToggled(fn TNotifyEvent) ISignalHandlerID
}

// ICellRenderer is a representation of GTK's GtkCellRenderer.
type ICellRenderer interface {
	IObject
}

// ITreeModel is a representation of GTK's GtkTreeModel.
type ITreeModel interface {
	IObject
}

// ITreeStore is a representation of GTK's GtkTreeStore (hierarchical tree data).
type ITreeStore interface {
	IObject
	Append(parent ITreeIter) ITreeIter
	SetValue(iter ITreeIter, column int, value string)
	Remove(iter ITreeIter) bool
	Clear()
}

// ITextIter is a representation of GTK's GtkTextIter.
type ITextIter interface {
	GetOffset() int
	GetLine() int
	GetLineOffset() int
	SetOffset(charOffset int)
	SetLine(lineNumber int)
	ForwardChar() bool
	BackwardChar() bool
	ForwardLine() bool
	BackwardLine() bool
	IsEnd() bool
	IsStart() bool
	Equal(other ITextIter) bool
}

// ITextBuffer is a representation of GTK's GtkTextBuffer.
type ITextBuffer interface {
	IObject
	SetText(text string)
	GetText(start, end ITextIter, includeHiddenChars bool) string
	GetBounds() (start, end ITextIter)
	GetCharCount() int
	GetLineCount() int
	Insert(iter ITextIter, text string)
	InsertAtCursor(text string)
	Delete(start, end ITextIter)
	GetStartIter() ITextIter
	GetEndIter() ITextIter
	GetIterAtOffset(charOffset int) ITextIter
	GetIterAtLine(lineNumber int) ITextIter
	GetModified() bool
	SetModified(setting bool)
	PlaceCursor(iter ITextIter)
	GetSelectionBounds() (start, end ITextIter, ok bool)
	DeleteSelection(interactive, defaultEditable bool) bool
	GetInsert() ITextMark
	CreateMark(markName string, where ITextIter, leftGravity bool) ITextMark
	GetMark(name string) ITextMark
	DeleteMark(mark ITextMark)
	GetIterAtMark(mark ITextMark) ITextIter
	ApplyTagByName(name string, start, end ITextIter)
	RemoveTagByName(name string, start, end ITextIter)
	SetOnChanged(fn TNotifyEvent) ISignalHandlerID
}

// ITextMark is a representation of GTK's GtkTextMark.
type ITextMark interface {
	IObject
	SetVisible(setting bool)
	GetVisible() bool
	GetDeleted() bool
	GetName() string
	GetBuffer() ITextBuffer
}

// ITextTag is a representation of GTK's GtkTextTag.
type ITextTag interface {
	IObject
	GetPriority() int
	SetPriority(priority int)
}

// ITextTagTable is a representation of GTK's GtkTextTagTable.
type ITextTagTable interface {
	IObject
	Add(tag ITextTag) bool
	Lookup(name string) ITextTag
	Remove(tag ITextTag)
}

// ITextView is a representation of GTK's GtkTextView.
type ITextView interface {
	IWidget
	GetBuffer() ITextBuffer
	SetBuffer(buffer ITextBuffer)
	SetEditable(editable bool)
	GetEditable() bool
	SetWrapMode(wrapMode WrapMode)
	GetWrapMode() WrapMode
	SetCursorVisible(visible bool)
	GetCursorVisible() bool
	SetOverwrite(overwrite bool)
	GetOverwrite() bool
	SetJustification(justify Justification)
	GetJustification() Justification
	SetAcceptsTab(acceptsTab bool)
	GetAcceptsTab() bool
	SetLeftMargin(margin int)
	GetLeftMargin() int
	SetRightMargin(margin int)
	GetRightMargin() int
	SetIndent(indent int)
	GetIndent() int
	SetPixelsAboveLines(px int)
	GetPixelsAboveLines() int
	SetPixelsBelowLines(px int)
	GetPixelsBelowLines() int
	SetPixelsInsideWrap(px int)
	GetPixelsInsideWrap() int
	ScrollToIter(iter ITextIter, withinMargin float64, useAlign bool, xalign, yalign float64) bool
	PlaceCursorOnscreen() bool
	ResetImContext()
}

// ITreeIter is a representation of GTK's GtkTreeIter.
type ITreeIter interface {
	Instance() uintptr
}

// IListStore is a representation of GTK's GtkListStore.
type IListStore interface {
	IObject
	Append() ITreeIter
	SetValue(iter ITreeIter, column int, value string)
	Remove(iter ITreeIter) bool
	Clear()
}

// ITreeViewColumn is a representation of GTK's GtkTreeViewColumn.
type ITreeViewColumn interface {
	IObject
	SetTitle(title string)
	GetTitle() string
	PackStart(cell ICellRenderer, expand bool)
	PackEnd(cell ICellRenderer, expand bool)
	AddAttribute(renderer ICellRenderer, attribute string, column int)
	SetOnClicked(fn TNotifyEvent) ISignalHandlerID
	SetResizable(resizable bool)
	GetResizable() bool
	SetSizing(sizing TreeViewColumnSizing)
	GetSizing() TreeViewColumnSizing
	SetFixedWidth(fixedWidth int)
	GetFixedWidth() int
	SetMinWidth(minWidth int)
	GetMinWidth() int
	SetMaxWidth(maxWidth int)
	GetMaxWidth() int
	SetExpand(expand bool)
	GetExpand() bool
	SetSortColumnId(sortColumnId int)
	GetSortColumnId() int
	SetSortIndicator(setting bool)
	GetSortIndicator() bool
	SetReorderable(reorderable bool)
	GetReorderable() bool
	SetAlignment(xalign float32)
	GetAlignment() float32
	GetWidth() int
	SetSpacing(spacing int)
	GetSpacing() int
}

// ITreeSelection is a representation of GTK's GtkTreeSelection.
type ITreeSelection interface {
	IObject
	SetMode(mode SelectionMode)
	GetMode() SelectionMode
	SetOnChanged(fn TNotifyEvent) ISignalHandlerID
	GetSelected() (ITreeModel, ITreeIter)
	CountSelectedRows() int
	SelectAll()
	UnselectAll()
}

// ITreeView is a representation of GTK's GtkTreeView.
type ITreeView interface {
	IContainer
	GetModel() IListStore
	SetModel(model IListStore)
	SetTreeModel(model ITreeModel)
	GetTreeModel() ITreeModel
	GetSelection() ITreeSelection
	AppendColumn(column ITreeViewColumn) int
	SetHeadersVisible(show bool)
	GetHeadersVisible() bool
	ExpandAll()
	CollapseAll()
	SetOnRowActivated(fn TRowActivatedEvent) ISignalHandlerID
	SetOnCursorChanged(fn TNotifyEvent) ISignalHandlerID
	SetOnRowExpanded(fn TTreeRowExpandCollapseEvent) ISignalHandlerID
	SetOnRowCollapsed(fn TTreeRowExpandCollapseEvent) ISignalHandlerID
	SetOnTestExpandRow(fn TTestExpandRowEvent) ISignalHandlerID
	SetOnTestCollapseRow(fn TTestExpandRowEvent) ISignalHandlerID
	SetOnColumnsChanged(fn TNotifyEvent) ISignalHandlerID
	SetOnSelectAll(fn TNotifyEvent) ISignalHandlerID
	SetOnUnselectAll(fn TNotifyEvent) ISignalHandlerID
	SetOnToggleCursorRow(fn TNotifyEvent) ISignalHandlerID
	SetOnStartInteractiveSearch(fn TNotifyEvent) ISignalHandlerID
}

// IFileChooserDialog is a representation of GTK's GtkFileChooserDialog.
type IFileChooserDialog interface {
	IDialog
	GetFilename() string
	SetFilename(filename string) bool
	SetCurrentFolder(folder string) bool
	SetCurrentName(name string)
	GetCurrentFolder() string
	SetAction(action FileChooserAction)
	GetAction() FileChooserAction
	SetSelectMultiple(selectMultiple bool)
	GetSelectMultiple() bool
	SetFilter(filter IFileFilter)
	GetFilter() IFileFilter
	AddFilter(filter IFileFilter)
	SetOnSelectionChanged(fn TNotifyEvent) ISignalHandlerID
	SetOnFileActivated(fn TNotifyEvent) ISignalHandlerID
}

// IFileFilter is a representation of GTK's GtkFileFilter.
type IFileFilter interface {
	IObject
	SetName(name string)
	GetName() string
	AddPattern(pattern string)
	AddMimeType(mimeType string)
}

// IComboBoxText is a representation of GTK's GtkComboBoxText.
type IComboBoxText interface {
	IWidget
	Append(id string, text string)
	AppendText(text string)
	GetActiveText() string
	RemoveAll()
	GetActive() int
	SetActive(index int)
	SetOnChanged(fn TNotifyEvent) ISignalHandlerID
}

// IGrid is a representation of GTK's GtkGrid.
type IGrid interface {
	IContainer
	Attach(child IWidget, left, top, width, height int)
	AttachNextTo(child, sibling IWidget, side PositionType, width, height int)
	SetRowSpacing(spacing uint)
	GetRowSpacing() uint
	SetColumnSpacing(spacing uint)
	GetColumnSpacing() uint
	SetRowHomogeneous(homogeneous bool)
	GetRowHomogeneous() bool
	SetColumnHomogeneous(homogeneous bool)
	GetColumnHomogeneous() bool
}

// ISeparator is a representation of GTK's GtkSeparator.
type ISeparator interface {
	IWidget
}

// IRadioButton is a representation of GTK's GtkRadioButton.
type IRadioButton interface {
	ICheckButton
}

// IAboutDialog is a representation of GTK's GtkAboutDialog.
type IAboutDialog interface {
	IDialog
	SetProgramName(name string)
	GetProgramName() string
	SetVersion(version string)
	GetVersion() string
	SetComments(comments string)
	GetComments() string
	SetWebsite(website string)
	GetWebsite() string
	SetWebsiteLabel(label string)
	GetWebsiteLabel() string
	SetLicense(license string)
	GetLicense() string
	SetAuthors(authors []string)
	SetTranslatorCredits(credits string)
}

// IColorChooserDialog is a representation of GTK's GtkColorChooserDialog.
type IColorChooserDialog interface {
	IDialog
	GetUseAlpha() bool
	SetUseAlpha(useAlpha bool)
	GetRGBA() GdkRGBA
}

// IFontChooserDialog is a representation of GTK's GtkFontChooserDialog.
type IFontChooserDialog interface {
	IDialog
	GetFont() string
	SetFont(font string)
	GetPreviewText() string
	SetPreviewText(text string)
}

// IScale is a representation of GTK's GtkScale (HScale/VScale base).
type IScale interface {
	IRange
	SetDigits(digits int)
	GetDigits() int
	SetDrawValue(drawValue bool)
	GetDrawValue() bool
	SetValuePos(pos PositionType)
	GetValuePos() PositionType
}

// ISpinner is a representation of GTK's GtkSpinner.
type ISpinner interface {
	IWidget
	Start()
	Stop()
}

// ILevelBar is a representation of GTK's GtkLevelBar.
type ILevelBar interface {
	IWidget
	SetValue(value float64)
	GetValue() float64
	SetMinValue(value float64)
	GetMinValue() float64
	SetMaxValue(value float64)
	GetMaxValue() float64
	SetMode(mode LevelBarMode)
	GetMode() LevelBarMode
	SetOnOffsetChanged(fn TOffsetChangedEvent) ISignalHandlerID
}

// IPaned is a representation of GTK's GtkPaned.
type IPaned interface {
	IContainer
	Add1(child IWidget)
	Add2(child IWidget)
	SetPosition(position int)
	GetPosition() int
	SetWideHandle(wide bool)
	GetWideHandle() bool
}

// IListBox is a representation of GTK's GtkListBox.
type IListBox interface {
	IContainer
	Prepend(child IWidget)
	Insert(child IWidget, position int)
	SelectRow(row IWidget)
	GetSelectedRow() IWidget
	SetSelectionMode(mode SelectionMode)
	GetSelectionMode() SelectionMode
	SetOnRowSelected(fn TNotifyEvent) ISignalHandlerID
}

// IPopover is a representation of GTK's GtkPopover.
type IPopover interface {
	IBin
	SetRelativeTo(widget IWidget)
	GetRelativeTo() IWidget
	SetPosition(position PositionType)
	GetPosition() PositionType
	SetModal(modal bool)
	GetModal() bool
	Popdown()
	Popup()
	SetOnClosed(fn TNotifyEvent) ISignalHandlerID
}

// ISearchEntry is a representation of GTK's GtkSearchEntry.
type ISearchEntry interface {
	IEntry
}

// IRevealer is a representation of GTK's GtkRevealer.
type IRevealer interface {
	IBin
	SetRevealChild(reveal bool)
	GetRevealChild() bool
	SetTransitionDuration(duration uint)
	GetTransitionDuration() uint
	SetTransitionType(t RevealerTransitionType)
	GetTransitionType() RevealerTransitionType
}
