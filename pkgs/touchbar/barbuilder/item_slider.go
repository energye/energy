package barbuilder

type SliderAccessoryWidth string

const (
	SliderAccessoryDefault SliderAccessoryWidth = "default"
	SliderAccessoryWide    SliderAccessoryWidth = "wide"
)

type SliderOnChange func(value float64)

type Slider struct {
	CommonProperties

	Label            string
	StartValue       float64
	MinimumValue     float64
	MaximumValue     float64
	MinimumAccessory Image
	MaximumAccessory Image
	AccessoryWidth   SliderAccessoryWidth
	OnChange         SliderOnChange
}

var _ Item = &Slider{}

func (me *Slider) isAnItem() {}
