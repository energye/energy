package cef

import "github.com/energye/golcl/lcl/api/dllimports"

func init() {
	var energyImportDefs = []*dllimports.ImportTable{
		dllimports.NewEnergyImport("CEFApplication_Create", 0),
		dllimports.NewEnergyImport("CEFApplication_Free", 0),
		dllimports.NewEnergyImport("CEFStartMainProcess", 0),
		dllimports.NewEnergyImport("CEFStartSubProcess", 0),
		dllimports.NewEnergyImport("AddCustomCommandLine", 0),
		dllimports.NewEnergyImport("CEFApplication_ExecuteJS", 0),
		dllimports.NewEnergyImport("CEFWindow_UpdateSize", 0),
		dllimports.NewEnergyImport("CEFWindow_OnEnter", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
		dllimports.NewEnergyImport("", 0),
	}
	dllimports.SetEnergyImportDefs(energyImportDefs)
}

const (
	internale_CEFApplication_Create = iota
	internale_CEFApplication_Free
	internale_CEFStartMainProcess
	internale_CEFStartSubProcess
	internale_AddCustomCommandLine
	internale_CEFApplication_ExecuteJS
	internale_CEFWindow_UpdateSize
)
