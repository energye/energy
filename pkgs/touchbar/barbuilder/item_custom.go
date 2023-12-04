package barbuilder

type CustomEventKind string

const (
	CustomEventChanged CustomEventKind = "changed"
	CustomEventEnded   CustomEventKind = "ended"
)

type CustomEvent struct {
	Event CustomEventKind
	X     float32
}

type CustomOnEvent func(event CustomEvent)

type Custom struct {
	CommonProperties

	OnEvent CustomOnEvent

	// TODO: how to choose what to render?
	// TODO: add gestures?
}

var _ Item = &Custom{}

func (me *Custom) isAnItem() {}
