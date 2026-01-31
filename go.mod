module github.com/energye/energy/v3

go 1.20

require (
	github.com/godbus/dbus/v5 v5.2.2
	golang.org/x/sys v0.30.0
)

replace (
	github.com/energye/cef => ../cef
	github.com/energye/lcl => ../lcl
	github.com/energye/wv => ../wv
)

require (
	github.com/energye/lcl v0.0.0-00010101000000-000000000000
	github.com/energye/wv v0.0.0-00010101000000-000000000000
	github.com/energye/cef v0.0.0-00010101000000-000000000000
)
