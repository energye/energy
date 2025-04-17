//go:build darwin

package bar

import (
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
)

func makeLabelCatalog() barbuilder.Item {
	return &barbuilder.Popover{
		CollapsedText:  "Label",
		CollapsedImage: barbuilder.SFSymbol("text.alignleft"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Label",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text:  "Color #f42309",
					Color: barbuilder.HexColor("#f42309"),
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Touch Bar Icon:",
				},
			},
			&barbuilder.Label{
				Content: &barbuilder.ContentImage{
					Image: barbuilder.TBAddTemplate,
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "SF Symbol:",
				},
			},
			&barbuilder.Label{
				Content: &barbuilder.ContentImage{
					Image: barbuilder.SFSymbol("text.bubble"),
				},
			},
		},
	}
}
