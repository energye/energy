//go:build darwin

package darwin

import (
	"fmt"

	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
)

const namespace = "com.energy.touchbar"

const (
	standardOtherItemsProxy = namespace + ".other_items"
	standardSpaceSmall      = namespace + ".small_space"
	standardSpaceLarge      = namespace + ".large_space"
	standardSpaceFlexible   = namespace + ".flexible_space"
	standardCandidateList   = namespace + ".candidates"
	standardCharacterPicker = namespace + ".char_picker"
	standardTextFormat      = namespace + ".text_format"
	standardTextAlignment   = namespace + ".text_align"
	standardTextColorPicker = namespace + ".text_color"
	standardTextList        = namespace + ".text_list"
	standardTextStyle       = namespace + ".text_style"
)

type identifier string

type flatConfig struct {
	Principal       identifier
	Default         []identifier
	Items           map[identifier]item
	OtherItemsProxy bool
	Escape          identifier
}

func makeID(prefix, kind string, i int) identifier {
	return identifier(fmt.Sprintf("%s.%s.%s.%d", namespace, kind, prefix, i))
}

func processItems(prefix string, items []barbuilder.Item, dict map[identifier]item, handlers *handlers) ([]identifier, identifier, error) {
	principal := identifier("")
	list := make([]identifier, 0, len(items))

	for i, item := range items {
		id, result, err := processItem(prefix, i, item, &principal, dict, handlers)
		if err != nil {
			return nil, "", err
		}

		if id == "" {
			return nil, "", fmt.Errorf("id not generated properly for %T (%v)", item, item)
		}

		dict[id] = result
		list = append(list, id)
	}

	return list, principal, nil
}

func processConfig(config *barbuilder.Configuration) (*flatConfig, *handlers, error) {
	dict := make(map[identifier]item, len(config.Items))
	handlers := handlers{
		buttons:      make(map[identifier]barbuilder.ButtonOnClick),
		colorPickers: make(map[identifier]barbuilder.ColorPickerOnSelected),
		customs:      make(map[identifier]barbuilder.CustomOnEvent),
		pickers:      make(map[identifier]barbuilder.PickerOnSelected),
		scrubbers:    make(map[identifier]barbuilder.ScrubberOnChange),
		segments:     make(map[identifier]barbuilder.SegmentedOnClick),
		sliders:      make(map[identifier]barbuilder.SliderOnChange),
		steppers:     make(map[identifier]barbuilder.StepperOnChange),
	}

	list, principal, err := processItems("root", config.Items, dict, &handlers)
	if err != nil {
		return nil, nil, err
	}

	escapeID := identifier("")
	if config.Escape != nil {
		id, result, err := processItem("escape", 0, *config.Escape, &principal, dict, nil)
		if err != nil {
			return nil, nil, err
		}
		dict[id] = result
		escapeID = id
	}

	data := flatConfig{
		Principal: principal,
		Default:   list,
		Items:     dict,
		Escape:    escapeID,
	}
	return &data, &handlers, nil
}
