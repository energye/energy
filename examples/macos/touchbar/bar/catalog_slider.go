//go:build darwin

package bar

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/cef/ipc"

	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barutils"
)

func makeSliderCatalog(switcher barutils.Switcher) barbuilder.Item {
	result := &barbuilder.Label{Content: &barbuilder.ContentLabel{Text: ""}}

	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Slider",
		CollapsedImage: barbuilder.SFSymbol("ruler.fill"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Slider",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Simple",
				PressAndHold:  true,
				Bar: []barbuilder.Item{
					&barbuilder.Slider{
						MinimumValue: 67,
						StartValue:   90,
						MaximumValue: 100,
						OnChange: func(value float64) {
							result.Content = &barbuilder.ContentLabel{Text: fmt.Sprintf("value: %v", value)}
							ipc.Emit("touchbar", 4, "touch bar Slider: "+fmt.Sprintf("%v", value))
							switcher.Update()
						},
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "All options",
				PressAndHold:  true,
				Bar: []barbuilder.Item{
					&barbuilder.Slider{
						Label:            "Normal",
						MinimumValue:     0,
						StartValue:       7,
						MaximumValue:     10,
						MinimumAccessory: barbuilder.TBAlarmTemplate,
						MaximumAccessory: barbuilder.TBAddDetailTemplate,
						AccessoryWidth:   barbuilder.SliderAccessoryWide,
						OnChange: func(value float64) {
							result.Content = &barbuilder.ContentLabel{Text: fmt.Sprintf("value: %v", value)}
							ipc.Emit("touchbar", 4, "touch bar All options: "+fmt.Sprintf("%v", value))
							switcher.Update()
						},
					},
				},
			},
			&barbuilder.SpaceSmall{},
			result,
		},
	}, switcher)
}
