//go:build linux

//Note that you need to have github.com/knightpp/dbus-codegen-go installed from "custom" branch
//go:generate dbus-codegen-go -prefix org.kde -package notifier -output internal/generated/notifier/status_notifier_item.go internal/StatusNotifierItem.xml
//go:generate dbus-codegen-go -prefix com.canonical -package menu -output internal/generated/menu/dbus_menu.go internal/DbusMenu.xml

package systray

import (
	"bytes"
	"fmt"
	"github.com/energye/energy/v3/application/internal/systray/internal/generated/menu"
	"github.com/energye/energy/v3/application/internal/systray/internal/generated/notifier"
	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/godbus/dbus/v5/prop"
	"image"
	_ "image/png" // used only here
	"log"
	"os"
	"sync"
	"time"
)

var (
	currentID = int32(0)
)

const (
	path     = "/StatusNotifierItem"
	menuPath = "/StatusNotifierMenu"
)

// Tray is a basic type that handles the dbus functionality
type Tray struct {
	conn                *dbus.Conn // the DBus connection that we will use
	iconData            []byte     // icon data for the main systray icon
	title, tooltipTitle string     // title and tooltip state
	visible             bool
	lock                sync.Mutex
	props, menuProps    *prop.Properties
	menuVersion         uint32
	menuItems           map[int32]*MenuItem
	menu                *MenuItem
	usni                *UnimplementedStatusNotifierItem
}

func (m *Tray) Menu() *MenuItem {
	if m.menu == nil {
		m.menu = &MenuItem{layout: &menuLayout{
			V0: 0,
			V1: map[string]dbus.Variant{},
			V2: []dbus.Variant{},
		}}
	}
	return m.menu
}

// SetIcon sets the systray icon.
// iconBytes should be the content of .ico for windows and .ico/.jpg/.png
// for other platforms.
func (m *Tray) SetIcon(iconBytes []byte) {
	m.iconData = iconBytes
	if m.props == nil {
		return
	}
	m.props.SetMust("org.kde.StatusNotifierItem", "IconPixmap", []PX{convertToPixels(iconBytes)})
	if m.conn == nil {
		return
	}
	err := notifier.Emit(m.conn, &notifier.StatusNotifierItem_NewIconSignal{
		Path: path,
		Body: &notifier.StatusNotifierItem_NewIconSignalBody{},
	})
	if err != nil {
		log.Printf("systray error: failed to emit new icon signal: %s\n", err)
		return
	}
}

// SetTitle sets the systray title, only available on Mac and Linux.
func (m *Tray) SetTitle(t string) {
	m.title = t

	if m.props == nil {
		return
	}
	dbusErr := m.props.Set("org.kde.StatusNotifierItem", "Title", dbus.MakeVariant(t))
	if dbusErr != nil {
		log.Printf("systray error: failed to set Title prop: %s\n", dbusErr)
		return
	}
	if m.conn == nil {
		return
	}
	err := notifier.Emit(m.conn, &notifier.StatusNotifierItem_NewTitleSignal{
		Path: path,
		Body: &notifier.StatusNotifierItem_NewTitleSignalBody{},
	})
	if err != nil {
		log.Printf("systray error: failed to emit new title signal: %s\n", err)
		return
	}
}

// SetVisible
func (m *Tray) SetVisible(value bool) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.visible = value
	props := m.props
	if props == nil {
		return
	}
	dbusErr := props.Set("org.kde.StatusNotifierItem", "Visible", dbus.MakeVariant(value))
	if dbusErr != nil {
		log.Printf("systray error: failed to set Visible prop: %s\n", dbusErr)
		return
	}
	err := m.conn.Emit(path, "org.kde.StatusNotifierItem.NewVisible", dbus.MakeVariant(value))
	if err != nil {
		log.Printf("systray error: failed to emit new Visible signal: %s\n", err)
		return
	}
}

// SetTooltip sets the systray tooltip to display on mouse hover of the tray icon,
// only available on Mac and Windows.
func (m *Tray) SetTooltip(tooltipTitle string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.tooltipTitle = tooltipTitle
	props := m.props

	if props == nil {
		return
	}
	dbusErr := props.Set("org.kde.StatusNotifierItem", "ToolTip", dbus.MakeVariant(tooltip{V2: tooltipTitle}))
	if dbusErr != nil {
		log.Printf("systray error: failed to set ToolTip prop: %s\n", dbusErr)
		return
	}
}

type UnimplementedStatusNotifierItem struct {
	contextMenu       func(x int32, y int32)
	activate          func(x int32, y int32)
	dActivate         func(x int32, y int32)
	secondaryActivate func(x int32, y int32)
	scroll            func(delta int32, orientation string)
	dActivateTime     int64
}

func (*UnimplementedStatusNotifierItem) iface() string {
	return notifier.InterfaceStatusNotifierItem
}

func (m *UnimplementedStatusNotifierItem) ContextMenu(x int32, y int32) (err *dbus.Error) {
	if m.contextMenu != nil {
		m.contextMenu(x, y)
	} else {
		err = &dbus.ErrMsgUnknownMethod
	}
	return
}

func (m *UnimplementedStatusNotifierItem) Activate(x int32, y int32) (err *dbus.Error) {
	if m.dActivateTime == 0 {
		m.dActivateTime = time.Now().UnixMilli()
	} else {
		nowMilli := time.Now().UnixMilli()
		if nowMilli-m.dActivateTime < dClickTimeMinInterval {
			m.dActivateTime = dClickTimeMinInterval
			if m.dActivate != nil {
				m.dActivate(x, y)
				return
			}
		} else {
			m.dActivateTime = nowMilli
		}
	}

	if m.activate != nil {
		m.activate(x, y)
	} else {
		err = &dbus.ErrMsgUnknownMethod
	}
	return
}

func (m *UnimplementedStatusNotifierItem) SecondaryActivate(x int32, y int32) (err *dbus.Error) {
	if m.secondaryActivate != nil {
		m.secondaryActivate(x, y)
	} else {
		err = &dbus.ErrMsgUnknownMethod
	}
	return
}

func (m *UnimplementedStatusNotifierItem) Scroll(delta int32, orientation string) (err *dbus.Error) {
	if m.scroll != nil {
		m.scroll(delta, orientation)
	} else {
		err = &dbus.ErrMsgUnknownMethod
	}
	return
}

func (m *Tray) SetOnClick(fn func()) {
	m.usni.activate = func(x int32, y int32) {
		fn()
	}
}

func (m *Tray) SetOnDClick(fn func()) {
	m.usni.dActivate = func(x int32, y int32) {
		fn()
	}
}

func (m *Tray) NativeEnd() {
	if m.conn != nil {
		_ = m.conn.Close()
		m.conn = nil
	}
}

func NativeStart() *Tray {
	m := &Tray{ /*menu: &menuLayout{}, */ menuVersion: 1, usni: &UnimplementedStatusNotifierItem{}, menuItems: make(map[int32]*MenuItem)}
	conn, _ := dbus.ConnectSessionBus()
	err := notifier.ExportStatusNotifierItem(conn, path, m.usni)
	if err != nil {
		log.Printf("systray error: failed to export status notifier item: %s\n", err)
	}
	err = menu.ExportDbusmenu(conn, menuPath, m)
	if err != nil {
		log.Printf("systray error: failed to export status notifier item: %s\n", err)
	}

	name := fmt.Sprintf("org.kde.StatusNotifierItem-%d-1", os.Getpid()) // register id 1 for this process
	_, err = conn.RequestName(name, dbus.NameFlagDoNotQueue)
	if err != nil {
		log.Printf("systray error: failed to request name: %s\n", err)
		// it's not critical error: continue
	}
	props, err := prop.Export(conn, path, m.createPropSpec())
	if err != nil {
		log.Printf("systray error: failed to export notifier item properties to bus: %s\n", err)
		return nil
	}
	menuProps, err := prop.Export(conn, menuPath, m.createMenuPropSpec())
	if err != nil {
		log.Printf("systray error: failed to export notifier menu properties to bus: %s\n", err)
		return nil
	}

	node := introspect.Node{
		Name: path,
		Interfaces: []introspect.Interface{
			introspect.IntrospectData,
			prop.IntrospectData,
			notifier.IntrospectDataStatusNotifierItem,
		},
	}
	err = conn.Export(introspect.NewIntrospectable(&node), path, "org.freedesktop.DBus.Introspectable")
	if err != nil {
		log.Printf("systray error: failed to export node introspection: %s\n", err)
		return nil
	}

	menuNode := introspect.Node{
		Name: menuPath,
		Interfaces: []introspect.Interface{
			introspect.IntrospectData,
			prop.IntrospectData,
			menu.IntrospectDataDbusmenu,
		},
	}
	err = conn.Export(introspect.NewIntrospectable(&menuNode), menuPath, "org.freedesktop.DBus.Introspectable")
	if err != nil {
		log.Printf("systray error: failed to export menu node introspection: %s\n", err)
		return nil
	}

	m.lock.Lock()
	m.conn = conn
	m.props = props
	m.menuProps = menuProps
	m.lock.Unlock()

	var (
		obj  dbus.BusObject
		call *dbus.Call
	)
	obj = conn.Object("org.kde.StatusNotifierWatcher", "/StatusNotifierWatcher")
	call = obj.Call("org.kde.StatusNotifierWatcher.RegisterStatusNotifierItem", 0, path)
	if call.Err != nil {
		log.Printf("systray error: failed to register our icon with the notifier watcher (maybe no tray is running?): %s\n", call.Err)
	}
	return m
}

func (*Tray) iface() string {
	return notifier.InterfaceStatusNotifierItem
}

func (t *Tray) createPropSpec() map[string]map[string]*prop.Prop {
	t.lock.Lock()
	defer t.lock.Unlock()
	return map[string]map[string]*prop.Prop{
		"org.kde.StatusNotifierItem": {
			"Status": {
				Value:    "Active", // Passive, Active or NeedsAttention
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"Title": {
				Value:    t.title,
				Writable: true,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"Id": {
				Value:    "1",
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"Category": {
				Value:    "ApplicationStatus",
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"IconName": {
				Value:    "",
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"IconPixmap": {
				Value:    []PX{convertToPixels(t.iconData)},
				Writable: true,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"IconThemePath": {
				Value:    "",
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"ItemIsMenu": {
				Value:    false,
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"Menu": {
				Value:    dbus.ObjectPath(menuPath),
				Writable: true,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"ToolTip": {
				Value:    tooltip{V2: t.tooltipTitle},
				Writable: true,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"Visible": {
				Value:    t.visible,
				Writable: true,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
		}}
}

func (m *Tray) Refresh() {
	if m.conn == nil || m.menuProps == nil {
		return
	}
	m.menuVersion++
	dbusErr := m.menuProps.Set("com.canonical.dbusmenu", "Version", dbus.MakeVariant(m.menuVersion))
	if dbusErr != nil {
		log.Printf("systray error: failed to update menu version: %s\n", dbusErr)
		return
	}
	err := menu.Emit(m.conn, &menu.Dbusmenu_LayoutUpdatedSignal{
		Path: menuPath,
		Body: &menu.Dbusmenu_LayoutUpdatedSignalBody{
			Revision: m.menuVersion,
		},
	})
	if err != nil {
		log.Printf("systray error: failed to emit layout updated signal: %s\n", err)
	}
}

func (m *Tray) ResetMenu() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.menu.Clear(m)
	m.menuItems = make(map[int32]*MenuItem)
	m.menuVersion++
	m.Refresh()
}

// copyLayout makes full copy of layout
func copyLayout(in *menuLayout, depth int32) *menuLayout {
	out := menuLayout{
		V0: in.V0,
		V1: make(map[string]dbus.Variant, len(in.V1)),
	}
	for k, v := range in.V1 {
		out.V1[k] = v
	}
	if depth != 0 {
		depth--
		out.V2 = make([]dbus.Variant, len(in.V2))
		for i, v := range in.V2 {
			out.V2[i] = dbus.MakeVariant(copyLayout(v.Value().(*menuLayout), depth))
		}
	} else {
		out.V2 = []dbus.Variant{}
	}
	return &out
}

// GetLayout is com.canonical.dbusmenu.GetLayout method.
func (m *Tray) GetLayout(parentID int32, recursionDepth int32, propertyNames []string) (revision uint32, layout menuLayout, err *dbus.Error) {
	revision = m.menuVersion
	if parentID == 0 {
		layout = *copyLayout(m.Menu().layout, recursionDepth)
	} else if partItem, ok := m.menuItems[parentID]; ok {
		layout = *copyLayout(partItem.layout, recursionDepth)
	}
	return
}

// GetGroupProperties is com.canonical.dbusmenu.GetGroupProperties method.
func (m *Tray) GetGroupProperties(ids []int32, propertyNames []string) (properties []struct {
	V0 int32
	V1 map[string]dbus.Variant
}, err *dbus.Error) {
	for _, id := range ids {
		if item, ok := m.menuItems[id]; ok {
			p := struct {
				V0 int32
				V1 map[string]dbus.Variant
			}{
				V0: item.layout.V0,
				V1: make(map[string]dbus.Variant, len(item.layout.V1)),
			}
			properties = append(properties, p)
		}
	}
	return
}

// GetProperty is com.canonical.dbusmenu.GetProperty method.
func (m *Tray) GetProperty(id int32, name string) (value dbus.Variant, err *dbus.Error) {
	if item, ok := m.menuItems[id]; ok {
		value = item.layout.V1[name]
	}
	return
}

// Event is com.canonical.dbusmenu.Event method.
func (m *Tray) Event(id int32, eventID string, data dbus.Variant, timestamp uint32) (err *dbus.Error) {
	if eventID == "clicked" {
		if item, ok := m.menuItems[id]; ok && item.click != nil {
			item.click()
		}
	} else if eventID == "" {

	}
	return
}

// EventGroup is com.canonical.dbusmenu.EventGroup method.
func (m *Tray) EventGroup(events []struct {
	V0 int32
	V1 string
	V2 dbus.Variant
	V3 uint32
}) (idErrors []int32, err *dbus.Error) {
	for _, event := range events {
		if event.V1 == "clicked" {
			if item, ok := m.menuItems[event.V0]; ok && item.click != nil {
				item.click()
			}
		}
	}
	return
}

// AboutToShow is com.canonical.dbusmenu.AboutToShow method.
func (t *Tray) AboutToShow(id int32) (needUpdate bool, err *dbus.Error) {
	return
}

// AboutToShowGroup is com.canonical.dbusmenu.AboutToShowGroup method.
func (t *Tray) AboutToShowGroup(ids []int32) (updatesNeeded []int32, idErrors []int32, err *dbus.Error) {
	return
}

func (t *Tray) createMenuPropSpec() map[string]map[string]*prop.Prop {
	return map[string]map[string]*prop.Prop{
		"com.canonical.dbusmenu": {
			"Version": {
				Value:    t.menuVersion,
				Writable: true,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"TextDirection": {
				Value:    "ltr",
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"Status": {
				Value:    "normal",
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
			"IconThemePath": {
				Value:    []string{},
				Writable: false,
				Emit:     prop.EmitTrue,
				Callback: nil,
			},
		},
	}
}

// menuLayout is a named struct to map into generated bindings. It represents the layout of a menu item
type menuLayout = struct {
	V0 int32                   // the unique ID of this item
	V1 map[string]dbus.Variant // properties for this menu item layout
	V2 []dbus.Variant          // child menu item layouts
}

// PX is picture pix map structure with width and high
type PX struct {
	W, H int
	Pix  []byte
}

// tooltip is our data for a tooltip property.
// Param names need to match the generated code...
type tooltip = struct {
	V0 string // name
	V1 []PX   // icons
	V2 string // title
	V3 string // description
}

func convertToPixels(data []byte) PX {
	if len(data) == 0 {
		return PX{}
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Printf("Failed to read icon format %v", err)
		return PX{}
	}

	return PX{
		img.Bounds().Dx(), img.Bounds().Dy(),
		argbForImage(img),
	}
}

func argbForImage(img image.Image) []byte {
	w, h := img.Bounds().Dx(), img.Bounds().Dy()
	data := make([]byte, w*h*4)
	i := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			data[i] = byte(a)
			data[i+1] = byte(r)
			data[i+2] = byte(g)
			data[i+3] = byte(b)
			i += 4
		}
	}
	return data
}
