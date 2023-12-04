package barbuilder

// TouchBar describes how to manipulate a Touch Bar
type TouchBar interface {
	// Install will associate the Touch Bar described by `config` to your main window
	Install(config Configuration) error

	// Debug is an alternative to `Install` which block the current thread to run a debug application
	// This is useful if you are developing/testing without a UI framework
	Debug(config Configuration) error

	// Update will update the currently installed Touch Bar using the provided `config`
	Update(config Configuration) error

	// Uninstall will remove
	Uninstall() error
}

// Options specifies the different settings available when creating a new Touch Bar
type Options struct {
	// EventErrorLogger will be called if any internal error arise
	// when handling events coming from interaction with your `Configuration`
	EventErrorLogger func(err error)
}

// Configuration represents the UI setup of your Touch Bar
type Configuration struct {
	// Items are the items to be displayed in your Touch Bar
	Items []Item
	// Escape is an optional setting to have a custom escape button to close you Touch Bar
	Escape *Item
}
