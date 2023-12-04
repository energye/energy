package barbuilder

type ColorPickerKind string

const (
	ColorPickerStandard ColorPickerKind = "standard"
	ColorPickerText     ColorPickerKind = "text"
	ColorPickerStroke   ColorPickerKind = "stroke"
)

type ColorPickerColor struct {
	RGB RGBAColor
}

type ColorPickerOnSelected func(color ColorPickerColor)

type ColorPicker struct {
	CommonProperties

	Kind       ColorPickerKind
	ShowsAlpha bool
	Disabled   bool
	OnSelected ColorPickerOnSelected

	// TODO: custom color list
	// TODO: custom color spaces
}

var _ Item = &ColorPicker{}

func (me *ColorPicker) isAnItem() {}
