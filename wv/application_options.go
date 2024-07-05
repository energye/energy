package wv

type Options struct {
	Name               string
	DefaultURL         string
	ICON               []byte
	Width              int
	Height             int
	DisableDevTools    bool
	DisableContextMenu bool
	LocalLoad          *LocalLoad
}
