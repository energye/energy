//go:build darwin

package darwin

import (
	"encoding/json"
	"fmt"

	"github.com/cyber-xxm/energy/v2/pkgs/touchbar/barbuilder"
)

type handlers struct {
	buttons      map[identifier]barbuilder.ButtonOnClick
	colorPickers map[identifier]barbuilder.ColorPickerOnSelected
	customs      map[identifier]barbuilder.CustomOnEvent
	pickers      map[identifier]barbuilder.PickerOnSelected
	scrubbers    map[identifier]barbuilder.ScrubberOnChange
	segments     map[identifier]barbuilder.SegmentedOnClick
	sliders      map[identifier]barbuilder.SliderOnChange
	steppers     map[identifier]barbuilder.StepperOnChange
}

const (
	eventButton      = "button"
	eventColorPicker = "color_picker"
	eventCustom      = "custom"
	eventPicker      = "picker"
	eventScrubber    = "scrubber"
	eventSegment     = "segment"
	eventSlider      = "slider"
	eventStepper     = "stepper"
)

type event struct {
	Kind   string
	Target identifier
	Data   json.RawMessage
}

func (me *touchBar) handleEventLogic(eventJSON string) (func(), error) {
	event := event{}
	err := json.Unmarshal([]byte(eventJSON), &event)
	if err != nil {
		return nil, err
	}

	switch event.Kind {
	case eventButton:
		handler, found := me.handlers.buttons[event.Target]
		if !found {
			return nil, fmt.Errorf("unknown button %v", event.Target)
		}
		return handler, nil

	case eventColorPicker:
		handler, found := me.handlers.colorPickers[event.Target]
		if !found {
			return nil, fmt.Errorf("unknown color picker %v", event.Target)
		}
		if handler == nil {
			return nil, nil
		}
		data := barbuilder.ColorPickerColor{}
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return nil, err
		}
		return func() {
			handler(data)
		}, nil

	case eventCustom:
		handler, found := me.handlers.customs[event.Target]
		if !found {
			return nil, fmt.Errorf("unknown custom %v", event.Target)
		}
		if handler == nil {
			return nil, nil
		}
		data := barbuilder.CustomEvent{}
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return nil, err
		}
		return func() {
			handler(data)
		}, nil

	case eventPicker:
		handler, found := me.handlers.pickers[event.Target]
		if !found {
			return nil, fmt.Errorf("unknown picker %v", event.Target)
		}
		if handler == nil {
			return nil, nil
		}
		data := 0
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return nil, err
		}
		return func() {
			handler(data)
		}, nil

	case eventScrubber:
		handler, found := me.handlers.scrubbers[event.Target]
		if !found {
			return nil, fmt.Errorf("unknown scrubber %v", event.Target)
		}
		if handler == nil {
			return nil, nil
		}
		data := 0
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return nil, err
		}
		return func() {
			handler(data)
		}, nil

	case eventSegment:
		handler, found := me.handlers.segments[event.Target]
		if !found {
			return nil, fmt.Errorf("unknown segment %v", event.Target)
		}
		if handler == nil {
			return nil, nil
		}
		data := []bool{}
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return nil, err
		}
		return func() {
			handler(data)
		}, nil

	case eventSlider:
		handler, found := me.handlers.sliders[event.Target]
		if !found {
			return nil, fmt.Errorf("unknown slider %v", event.Target)
		}
		if handler == nil {
			return nil, nil
		}
		data := float64(0)
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return nil, err
		}
		return func() {
			handler(data)
		}, nil

	case eventStepper:
		handler, found := me.handlers.steppers[event.Target]
		if !found {
			return nil, fmt.Errorf("unknown stepper %v", event.Target)
		}
		if handler == nil {
			return nil, nil
		}
		data := float64(0)
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return nil, err
		}
		return func() {
			handler(data)
		}, nil

	default:
		return nil, fmt.Errorf("unknown kind %v", event.Kind)
	}
}

func (me *touchBar) handleEvent(eventJSON string) {
	me.lock.Lock()
	handler, err := me.handleEventLogic(eventJSON)
	if err != nil && me.options.EventErrorLogger != nil {
		me.options.EventErrorLogger(err)
	}
	me.lock.Unlock()
	if err == nil && handler != nil {
		handler()
	}
}
