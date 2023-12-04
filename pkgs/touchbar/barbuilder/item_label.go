package barbuilder

type Label struct {
	CommonProperties

	Content View
}

var _ Item = &Label{}

func (me *Label) isAnItem() {}
