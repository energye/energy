package barbuilder

type GroupDirection string

const (
	GroupMatchApp    GroupDirection = "app"
	GroupLeftToRight GroupDirection = "left_to_right"
	GroupRightToLeft GroupDirection = "right_to_left"
)

type Group struct {
	CommonProperties

	Direction          GroupDirection
	Children           []Item
	PrefersEqualWidth  bool
	PreferredItemWidth float32 // only used when PrefersEqualWidth is true
}

var _ Item = &Group{}

func (me *Group) isAnItem() {}
