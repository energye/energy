package barbuilder

// Color represents a color usable by macOS
type Color interface {
	isAColor()
}

// RGBAColor represents a color in sRGB space given through its individual components
type RGBAColor struct {
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}

var _ Color = &RGBAColor{}

func (me *RGBAColor) isAColor() {}

// HexColor represents a color in sRGB color given as a CSS color string, e.g. `#ff34ad`
// Note: it does not support alpha channel, if you need opacity, use `RGBAColor`
type HexColor string

var _ Color = HexColor("")

func (me HexColor) isAColor() {}

// TODO: add standard colors https://developer.apple.com/documentation/appkit/nscolor/standard_colors?changes=_5&language=objc
// TODO: add UI colors https://developer.apple.com/documentation/appkit/nscolor/ui_element_colors?changes=_5&language=objc
// TODO: add more custom color spaces
