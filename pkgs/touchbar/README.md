# go-touchbar [![Go Reference](https://pkg.go.dev/badge/github.com/LouisBrunner/go-touchbar.svg)](https://pkg.go.dev/github.com/LouisBrunner/go-touchbar)

Go library to integrate the MacBook Touch Bar

## Installation

```bash
go get github.com/LouisBrunner/go-touchbar
```

## Usage

```go
// Setup your window code (including NSApplication/NSWindow on macOS)

tb := touchbar.New(touchbar.Options{})

err := tb.Install(Configuration: touchbar.Configuration{
  // Add your configuration here
})
if err != nil {
  // handle
}

// run your application

// when you want to update the touchbar (even from another routine), call do
err = tb.Update(Configuration: touchbar.Configuration{
  // Add your updated configuration here
})
if err != nil {
  // handle
}

err = tb.Uninstall()
if err != nil {
  // handle
}
```

### Configuration

https://user-images.githubusercontent.com/7120871/197391736-450bef36-4dd6-4c79-8ae7-3bd68a932eec.mp4

See [example application](./examples/tester/main.go) for a real-life example.

Note: most of the widgets are direct translations of the one in Apple's documentation and have similar options.
Please refer to [the official documentation](https://developer.apple.com/documentation/appkit/touch_bar/creating_and_customizing_the_touch_bar?language=objc) for more details.

You configure this library by specifying a list of items from the following options:

- `Button`: a button that can be clicked (has an event handler)
- **NOT IMPLEMENTED** `Candidates`: a list of text options, allows to give custom recommendations for the current text field
- **NOT IMPLEMENTED** `ColorPicker`: used to choose a color (has an event handler)
- **NOT IMPLEMENTED** `Custom`: can render anything you want, feedback any touch with a X position (has an event handler)
- **NOT IMPLEMENTED** `Group`: used to group other items
- `Label`: allows to display text and images
- **NOT IMPLEMENTED** `Picker`: used to pick one or more options from a list of choices, equivalent to radio/checkbox buttons (has an event handler)
- `Popover`: a button which shows more options when clicked, equivalent to a modal or collapse widget
- **NOT IMPLEMENTED** `Scrubber`: allows to select items from a list, which supports scrolling (has an event handler)
- **NOT IMPLEMENTED** `SegmentedControl`: shows multiple buttons grouped together, similar to `Picker` (has an event handler)
- **NOT IMPLEMENTED** `Sharer`: allows to share using the standard macOS sharing system
- `Slider`: a slider used to pick value in a range (has an event handler)
- **NOT IMPLEMENTED** `Stepper`: a control to pick a value in a range through increments (has an event handler)

But also from this list of standard Apple widgets which have no customizable options:

- `OtherItemsProxy`: where to display other Touch Bar closer to the first responder
- `SpaceSmall`: a small space
- `SpaceLarge`: a bigger space
- `SpaceFlexible`: a space that grows as much as possible
<!-- - `CharacterPicker`: opens the macOS character picker (e.g. to pick an emoji)
- `CandidateList`: shows the standard macOS predictive text widget
- `TextFormat`: a group of text formatting options
- `TextAlignment`: allows to pick text alignement
- `TextColorPicker`: allows to pick text color
- `TextList`: allows to pick text listing options
- `TextStyle`: allows to pick text style -->

## Further work

Check TODO/FIXME as well

- (!!!) Finish implementing widgets
- Allow user customization (`customizationLabel`, `templateItems`, etc)
- Layout constraints (e.g. sizing)
- More options for widgets:

  - Color-picker
  - Custom
  - Popover

- Support standard/UI colors
- Support custom images
- Better validation in Go (validator on the structs?)
- A few random crashes: `signal arrived during cgo execution` and `[touchBar itemIdentifiers] was mutated while items array was being built` (in `[WindowController setupTouchBar]`)

## Acknowledgements

This library's API was influenced by [Electron's](https://www.electronjs.org/docs/latest/api/touch-bar).

The [demo application](./examples/tester/demo.go) is a reimplementation of [`electron-touch-bar`](https://github.com/pahund/electron-touch-bar).

The [catalog](./examples/tester/catalog.go) was influenced by [Apple's](https://developer.apple.com/documentation/appkit/touch_bar/creating_and_customizing_the_touch_bar?language=objc), which also helped with the Objective-C part of the implementation.
