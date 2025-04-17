//go:build darwin

package barutils

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
)

// Switcher is used to manipulate the internal state of a stackable bar created with `MakeStackableBar`
type Switcher interface {
	Update() error
	Replace(newItems []barbuilder.Item) error

	Pop() error
	Push(newItems []barbuilder.Item) error
}

type switcher struct {
	bar   barbuilder.TouchBar
	stack [][]barbuilder.Item
}

func (me *switcher) Update() error {
	if len(me.stack) < 1 {
		return fmt.Errorf("nothing to update")
	}
	return me.bar.Update(barbuilder.Configuration{Items: me.stack[len(me.stack)-1]})
}

func (me *switcher) Replace(newItems []barbuilder.Item) error {
	if len(me.stack) < 1 {
		return fmt.Errorf("nothing to replace")
	}
	me.stack[len(me.stack)-1] = newItems
	return me.Update()
}

func (me *switcher) Pop() error {
	if len(me.stack) <= 1 {
		return fmt.Errorf("nothing to pop")
	}
	me.stack = me.stack[:len(me.stack)-1]
	return me.Update()
}

func (me *switcher) Push(newItems []barbuilder.Item) error {
	me.stack = append(me.stack, append([]barbuilder.Item{
		&barbuilder.Button{
			Title: "Close",
			OnClick: barbuilder.ButtonOnClick(func() {
				// FIXME: potentially ignores error
				me.Pop()
			}),
		},
	}, newItems...))
	return me.Update()
}

// MakeStackableBar must be used to wrap a bar that will use a `VirtualPopover`
func MakeStackableBar(bar barbuilder.TouchBar, scope func(switcher Switcher) []barbuilder.Item) barbuilder.Configuration {
	switcher := &switcher{
		bar:   bar,
		stack: [][]barbuilder.Item{},
	}
	mainItems := scope(switcher)
	switcher.stack = [][]barbuilder.Item{
		mainItems,
	}
	return barbuilder.Configuration{Items: mainItems}
}

// VirtualPopover is useful if you need to nest more than 1 popover
// Currently the Touch Bar only supports one level of popovers and cannot nest them at all
// Using this function you can easily transform a popover into a other components which allow infinite nesting.
func VirtualPopover(popover barbuilder.Popover, switcher Switcher) barbuilder.Item {
	return &barbuilder.Button{
		Title: popover.CollapsedText,
		Image: popover.CollapsedImage,
		OnClick: func() {
			switcher.Push(popover.Bar)
		},
	}
}
