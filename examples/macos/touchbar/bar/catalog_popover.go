//go:build darwin

package bar

import (
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barutils"
)

func makePopoverCatalog(switcher barutils.Switcher) barbuilder.Item {
	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Popover",
		CollapsedImage: barbuilder.SFSymbol("bubble.left.fill"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Popover",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Click Once",
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Releasing doesn't dismiss, click the X when you are done",
						},
					},
					&barbuilder.Button{
						Title: "Click Me",
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedImage: barbuilder.SFSymbol("escape"),
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Releasing doesn't dismiss, click the X when you are done",
						},
					},
					&barbuilder.Button{
						Title: "Click Me",
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText:  "Press & Hold",
				CollapsedImage: barbuilder.SFSymbol("rectangle.compress.vertical"),
				PressAndHold:   true,
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Keep pressing!",
						},
					},
					&barbuilder.Button{
						Title: "Click me by sliding and releasing over me",
					},
				},
			},
		},
	}, switcher)
}
