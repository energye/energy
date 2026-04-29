package assets

import "embed"

//go:embed  tab
var tab embed.FS

func Tab(file string) []byte {
	data, err := tab.ReadFile("tab/" + file)
	if err != nil {
		return nil
	}
	return data
}
