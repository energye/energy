package barbuilder

type SegmentedOnClick func(clicked []bool)

type Segment struct {
	Label string
	Image Image
}

type SegmentedControl struct {
	CommonProperties

	Segments       []Segment
	SelectMultiple bool
	OnClick        SegmentedOnClick
}

var _ Item = &SegmentedControl{}

func (me *SegmentedControl) isAnItem() {}
