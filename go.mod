module github.com/energye/energy/v3

go 1.20

require (
	golang.org/x/sys v0.30.0
)

replace (
	github.com/energye/lcl => ../lcl
)

require (
	github.com/energye/lcl v0.0.0-beta
)
