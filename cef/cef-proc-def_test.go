package cef

import (
	"github.com/energye/golcl/lcl/api/dllimports"
	"testing"
)

func TestProcDef(t *testing.T) {
	for i, impTab := range dllimports.GetEnergyImports() {
		println(i, impTab.Name())
	}
}
