package gen

import (
	"encoding/json"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestIcon(t *testing.T) {
	wd, _ := os.Getwd()
	testdata := filepath.Join(wd, "testdata")
	icon := filepath.Join(testdata, "Go-Energy.png")
	outPath := filepath.Join(wd, "out")
	os.MkdirAll(outPath, fs.ModePerm)
	println("wd:", wd)
	outPath, err := GeneraICON(icon, outPath)
	if err != nil {
		panic(err)
	}
}

func TestSYSO(t *testing.T) {
	var infoJSON = `{
    "companyName": "demo",
    "productName": "demo",
    "fileVersion": "1.0.0",
    "productVersion": "1.0.0",
    "copyright": "Copyright.....",
    "comments": "Built using ENERGY (https://github.com/energye/energy)",
    "fileDescription": "Built using ENERGY (https://github.com/energye/energy)"
}`
	wd, _ := os.Getwd()
	iconPath := filepath.Join(wd, "out", "icon.ico")
	outPath := filepath.Join(wd, "out")
	info := project.Info{}
	err := json.Unmarshal([]byte(infoJSON), &info)
	if err != nil {
		panic(err)
	}
	_, err = GeneraSYSO("demo", iconPath, "", outPath, "amd64", info)
	if err != nil {
		panic(err)
	}
	_, err = GeneraSYSO("demo", iconPath, "", outPath, "386", info)
	if err != nil {
		panic(err)
	}
}
