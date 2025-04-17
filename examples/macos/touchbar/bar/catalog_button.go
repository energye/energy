//go:build darwin

package bar

import (
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barutils"
)

func makeButtonCatalog(switcher barutils.Switcher) barbuilder.Item {
	result := &barbuilder.Label{Content: &barbuilder.ContentLabel{Text: ""}}

	return &barbuilder.Popover{
		CollapsedText:  "Button",
		CollapsedImage: barbuilder.SFSymbol("hand.point.up.fill"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Button",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Button{
				Title: "Plain",
				OnClick: func() {
					result.Content = &barbuilder.ContentLabel{Text: "Button1 pressed"}
					switcher.Update()
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Button{
				Image: barbuilder.TBAlarmTemplate,
				OnClick: func() {
					result.Content = &barbuilder.ContentLabel{Text: "Button2 pressed"}
					switcher.Update()
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Button{
				Title:      "With Icon & Color",
				Image:      barbuilder.SFSymbol("exclamationmark.triangle.fill"),
				BezelColor: barbuilder.HexColor("#e35412"),
				OnClick: func() {
					result.Content = &barbuilder.ContentLabel{Text: "Button3 pressed"}
					switcher.Update()
				},
			},
			&barbuilder.SpaceSmall{},
			result,
			&barbuilder.SpaceSmall{},
			&barbuilder.Button{
				Title:    "Disabled",
				Image:    barbuilder.SFSymbol("sunrise.fill"),
				Disabled: true,
				OnClick: func() {
					result.Content = &barbuilder.ContentLabel{Text: "Button4 pressed"}
					switcher.Update()
				},
			},
		},
	}
}
