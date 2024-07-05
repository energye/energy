package wv

type Options struct {
	Name            string
	DefaultURL      string
	ICON            []byte
	Width           int
	Height          int
	EnabledDevTools bool
	LocalLoad       *LocalLoad
}
