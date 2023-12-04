package barbuilder

type OtherItemsProxy struct {
}

func (me *OtherItemsProxy) isAnItem() {}

var _ Item = &OtherItemsProxy{}

type SpaceSmall struct {
}

func (me *SpaceSmall) isAnItem() {}

var _ Item = &SpaceSmall{}

type SpaceLarge struct {
}

var _ Item = &SpaceLarge{}

func (me *SpaceLarge) isAnItem() {}

type SpaceFlexible struct {
}

var _ Item = &SpaceFlexible{}

func (me *SpaceFlexible) isAnItem() {}

// TODO: no idea how to use them?
// type CharacterPicker struct {
// }

// var _ Item = &CharacterPicker{}

// func (me *CharacterPicker) isAnItem() {}

// type CandidateList struct {
// }

// var _ Item = &CandidateList{}

// func (me *CandidateList) isAnItem() {}

// type TextFormat struct {
// }

// var _ Item = &TextFormat{}

// func (me *TextFormat) isAnItem() {}

// type TextAlignment struct {
// }

// var _ Item = &TextAlignment{}

// func (me *TextAlignment) isAnItem() {}

// type TextColorPicker struct {
// }

// var _ Item = &TextColorPicker{}

// func (me *TextColorPicker) isAnItem() {}

// type TextList struct {
// }

// var _ Item = &TextList{}

// func (me *TextList) isAnItem() {}

// type TextStyle struct {
// }

// var _ Item = &TextStyle{}

// func (me *TextStyle) isAnItem() {}
