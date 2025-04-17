//go:build darwin

package darwin

import (
	"fmt"
	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
)

func processItem(prefix string, i int, item barbuilder.Item, principal *identifier, dict map[identifier]item, handlers *handlers) (identifier, interface{}, error) {
	id := identifier("")
	isPrincipal := false
	var result interface{} = item

	switch widget := item.(type) {
	// customizable
	case *barbuilder.Button:
		id = makeID(prefix, "button", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		result = itemButton{
			CommonProperties: widget.CommonProperties,
			Title:            widget.Title,
			Image:            widget.Image,
			Disabled:         widget.Disabled,
			BezelColor:       widget.BezelColor,
		}
		handlers.buttons[id] = widget.OnClick

	case *barbuilder.Candidates:
		id = makeID(prefix, "candidates", i)
		isPrincipal = widget.Principal

	case *barbuilder.ColorPicker:
		id = makeID(prefix, "colorpicker", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.colorPickers[id] = widget.OnSelected

	case *barbuilder.Custom:
		id = makeID(prefix, "custom", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.customs[id] = widget.OnEvent

	case *barbuilder.Group:
		id = makeID(prefix, "group", i)
		isPrincipal = widget.Principal
		list, principal, err := processItems(fmt.Sprintf("%s.group.%d", prefix, i), widget.Children, dict, handlers)
		if err != nil {
			return "", nil, err
		}
		if principal != "" {
			return "", nil, fmt.Errorf("principal is not supported in sub touch bars")
		}
		result = itemGroup{
			CommonProperties:   widget.CommonProperties,
			Direction:          widget.Direction,
			Children:           list,
			PrefersEqualWidth:  widget.PrefersEqualWidth,
			PreferredItemWidth: widget.PreferredItemWidth,
		}

	case *barbuilder.Label:
		id = makeID(prefix, "label", i)
		isPrincipal = widget.Principal

	case *barbuilder.Picker:
		id = makeID(prefix, "picker", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.pickers[id] = widget.OnSelected

	case *barbuilder.Popover:
		id = makeID(prefix, "popover", i)
		isPrincipal = widget.Principal
		list, principal, err := processItems(fmt.Sprintf("%s.popover.%d", prefix, i), widget.Bar, dict, handlers)
		if err != nil {
			return "", nil, err
		}
		result = itemPopover{
			CommonProperties: widget.CommonProperties,
			CollapsedText:    widget.CollapsedText,
			CollapsedImage:   widget.CollapsedImage,
			Bar:              list,
			Principal:        principal,
			PressAndHold:     widget.PressAndHold,
		}

	case *barbuilder.Scrubber:
		id = makeID(prefix, "scrubber", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.scrubbers[id] = widget.OnChange

	case *barbuilder.SegmentedControl:
		id = makeID(prefix, "segmented", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.segments[id] = widget.OnClick

	case *barbuilder.Sharer:
		id = makeID(prefix, "sharer", i)
		isPrincipal = widget.Principal

	case *barbuilder.Slider:
		id = makeID(prefix, "slider", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.sliders[id] = widget.OnChange
		result = itemSlider{
			CommonProperties: widget.CommonProperties,
			Label:            widget.Label,
			StartValue:       widget.StartValue,
			MinimumValue:     widget.MinimumValue,
			MaximumValue:     widget.MaximumValue,
			MinimumAccessory: widget.MinimumAccessory,
			MaximumAccessory: widget.MaximumAccessory,
			AccessoryWidth:   widget.AccessoryWidth,
		}

	case *barbuilder.Stepper:
		id = makeID(prefix, "stepper", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.steppers[id] = widget.OnChange

		// standards
	case *barbuilder.OtherItemsProxy:
		id = standardOtherItemsProxy
	case *barbuilder.SpaceSmall:
		id = standardSpaceSmall
	case *barbuilder.SpaceLarge:
		id = standardSpaceLarge
	case *barbuilder.SpaceFlexible:
		id = standardSpaceFlexible
	// case *barbuilder.CharacterPicker:
	// 	id = standardCharacterPicker
	// case *barbuilder.CandidateList:
	// 	id = standardCandidateList
	// case *barbuilder.TextFormat:
	// 	id = standardTextFormat
	// case *barbuilder.TextAlignment:
	// 	id = standardTextAlignment
	// case *barbuilder.TextColorPicker:
	// 	id = standardTextColorPicker
	// case *barbuilder.TextList:
	// 	id = standardTextList
	// case *barbuilder.TextStyle:
	// 	id = standardTextStyle

	default:
		return "", nil, fmt.Errorf("unknown item type %T (%v)", item, item)
	}

	if isPrincipal {
		if *principal != "" {
			return "", nil, fmt.Errorf("duplicate principal: %v vs %v", *principal, id)
		}
		*principal = id
	}

	return id, result, nil
}
