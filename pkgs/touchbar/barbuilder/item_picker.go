package barbuilder

type PickerOnSelected func(i int)

type Picker struct {
	CommonProperties

	Items           []View
	SingleSelection bool
	Collapsed       bool
	OnSelected      PickerOnSelected
}

var _ Item = &Picker{}

func (me *Picker) isAnItem() {}
