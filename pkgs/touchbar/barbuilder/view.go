package barbuilder

// View describes a piece of UI that macOS can render,
// usually this is used by elements that can be rendered in different ways (e.g. `Label`).
type View interface {
	isAView()
}

// ContentLabel displays some text with an optional color
type ContentLabel struct {
	Text  string
	Color Color
}

var _ View = &ContentLabel{}

func (me *ContentLabel) isAView() {}

// ContentImage displays an image
type ContentImage struct {
	Image Image
}

var _ View = &ContentImage{}

func (me *ContentImage) isAView() {}
