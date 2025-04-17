//go:build darwin

package bar

import (
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barutils"
)

func makeCommonCatalog(switcher barutils.Switcher) barbuilder.Item {
	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Common",
		CollapsedImage: barbuilder.SFSymbol("figure.stand"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Common",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Principal",
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Before",
						},
					},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Principal: true,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Principal",
						},
					},
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "After",
						},
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Priority",
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityMedium,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Medium",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh,
						},
						Content: &barbuilder.ContentLabel{
							Text: "High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityMedium,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Medium",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow * 2,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Very Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh,
						},
						Content: &barbuilder.ContentLabel{
							Text: "High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh * 2,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Very High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh,
						},
						Content: &barbuilder.ContentLabel{
							Text: "High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityMedium,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Medium",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh * 2,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Very High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow * 2,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Very Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityMedium,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Medium",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh,
						},
						Content: &barbuilder.ContentLabel{
							Text: "VERY LONG TEXT WHICH WILL HIDE THE OTHERS",
						},
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Spaces",
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Small:",
						},
					},
					&barbuilder.SpaceSmall{},
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Large:",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Flexible:",
						},
					},
					&barbuilder.SpaceFlexible{},
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "End",
						},
					},
				},
			},
			// FIXME: doesn't work at the moment
			// &barbuilder.SpaceSmall{},
			// &barbuilder.Popover{
			// 	CollapsedText: "Text",
			// 	Bar: []barbuilder.Item{
			// 		&barbuilder.TextFormat{},
			// 		&barbuilder.TextAlignment{},
			// 		&barbuilder.TextColorPicker{},
			// 		&barbuilder.TextList{},
			// 		&barbuilder.TextStyle{},
			// 	},
			// },
			// &barbuilder.SpaceSmall{},
			// barutils.VirtualPopover(barbuilder.Popover{
			// 	CollapsedText: "Others",
			// 	Bar: []barbuilder.Item{
			// 		&barbuilder.CharacterPicker{},
			// 		&barbuilder.CandidateList{},
			// 	},
			// }, switcher),
		},
	}, switcher)
}
