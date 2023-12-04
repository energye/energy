package bar

import (
	"fmt"
	"github.com/energye/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/energye/energy/v2/pkgs/touchbar/barutils"
)

func MakeCatalog(switcher barutils.Switcher, update func()) barbuilder.Item {
	fmt.Println("MakeCatalog")
	// TODO: showcase Escape
	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Catalog",
		CollapsedImage: barbuilder.TBBookmarksTemplate,
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Catalog",
				},
			},
			makeCommonCatalog(switcher),
			makeLabelCatalog(),
			makeButtonCatalog(update),
			barutils.VirtualPopover(barbuilder.Popover{
				CollapsedText:  "Next",
				CollapsedImage: barbuilder.TBBookmarksTemplate,
				Bar: []barbuilder.Item{
					makePopoverCatalog(switcher),
					makeSliderCatalog(switcher, update),
				},
			}, switcher),
		},
	}, switcher)
}
