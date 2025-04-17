//go:build darwin

package darwin

import (
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
)

type item interface{}

type itemButton struct {
	barbuilder.CommonProperties

	Title      string
	Image      barbuilder.Image
	Disabled   bool
	BezelColor barbuilder.Color
}

type itemGroup struct {
	barbuilder.CommonProperties

	Direction          barbuilder.GroupDirection
	Children           []identifier
	PrefersEqualWidth  bool
	PreferredItemWidth float32
}

type itemSlider struct {
	barbuilder.CommonProperties

	Label            string
	StartValue       float64
	MinimumValue     float64
	MaximumValue     float64
	MinimumAccessory barbuilder.Image
	MaximumAccessory barbuilder.Image
	AccessoryWidth   barbuilder.SliderAccessoryWidth
}

type itemPopover struct {
	barbuilder.CommonProperties

	CollapsedText  string
	CollapsedImage barbuilder.Image
	Bar            []identifier
	Principal      identifier
	PressAndHold   bool
}
