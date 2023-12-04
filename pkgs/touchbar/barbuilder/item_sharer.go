package barbuilder

type Sharer struct {
	CommonProperties

	ButtonImage Image
	ButtonLabel string
	Disabled    bool

	// TODO: needs delegate or something?
}

var _ Item = &Sharer{}

func (me *Sharer) isAnItem() {}
