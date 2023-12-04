package barbuilder

type Popover struct {
	CommonProperties

	CollapsedText  string
	CollapsedImage Image
	Bar            []Item
	PressAndHold   bool

	// TODO: support for custom button with `collapsedRepresentation`
	// TODO: support custom close button `HideCloseButton`
}

var _ Item = &Popover{}

func (me *Popover) isAnItem() {}
