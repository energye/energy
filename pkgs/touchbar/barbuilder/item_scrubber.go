package barbuilder

type ScrubberOnChange func(i int)

type ScrubberOutline string

const (
	ScrubberOutlineNone       ScrubberOutline = "none"
	ScrubberOutlineBackground ScrubberOutline = "background"
	ScrubberOutlineOutline    ScrubberOutline = "outline"
)

type ScrubberMode string

const (
	ScrubberModeFree  ScrubberMode = "free"
	ScrubberModeFixed ScrubberMode = "fixed"
)

type Scrubber struct {
	CommonProperties

	Items            []View
	ShowsArrows      bool
	Mode             ScrubberMode
	SelectionOutline ScrubberOutline
	OnChange         ScrubberOnChange
}

var _ Item = &Scrubber{}

func (me *Scrubber) isAnItem() {}
