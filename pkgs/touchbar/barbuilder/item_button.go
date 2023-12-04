package barbuilder

type ButtonOnClick func()

type Button struct {
	CommonProperties

	Title      string
	Image      Image
	Disabled   bool
	BezelColor Color
	OnClick    ButtonOnClick
}

var _ Item = &Button{}

func (me *Button) isAnItem() {}
