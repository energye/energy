//go:build darwin

package bar

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barutils"
)

func MakeCatalog(switcher barutils.Switcher) barbuilder.Item {
	fmt.Println("MakeCatalog")
	// TODO: showcase Escape
	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "目录",
		CollapsedImage: barbuilder.TBBookmarksTemplate,
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Catalog",
				},
			},
			makeCommonCatalog(switcher),
			makeLabelCatalog(),
			makeButtonCatalog(switcher),
			barutils.VirtualPopover(barbuilder.Popover{
				CollapsedText:  "Next",
				CollapsedImage: barbuilder.TBBookmarksTemplate,
				Bar: []barbuilder.Item{
					makePopoverCatalog(switcher),
					makeSliderCatalog(switcher),
				},
			}, switcher),
		},
	}, switcher)
}
