package barbuilder

// Item represents a UI element that can be rendered in a Touch Bar by macOS
type Item interface {
	isAnItem()
}

// ItemPriority informs macOS on which items to hide first when space becomes constrained
// Lower priority elements are hidden first.
// Note: you can use custom value if you need more granularity than the pre-defined defaults
type ItemPriority float32

const (
	ItemPriorityLow    ItemPriority = -1000
	ItemPriorityMedium ItemPriority = 0
	ItemPriorityHigh   ItemPriority = 1000
)

// CommonProperties describes attributes supported by all `Item`s
type CommonProperties struct {
	// Priority informs macOS on which items to hide first, see `ItemPriority` for more details.
	Priority ItemPriority
	// Principal will position the `Item` at the center of the bar
	Principal bool
}
