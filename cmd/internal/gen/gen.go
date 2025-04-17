//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package gen

import (
	"encoding/json"
	"errors"
	"github.com/cyber-xxm/energy/v2/cmd/internal/command"
	"github.com/cyber-xxm/energy/v2/cmd/internal/project"
	"github.com/cyber-xxm/energy/v2/cmd/internal/term"
	"io/ioutil"
)

// Gen 生成 windows icon syso
func Gen(c *command.Config) (err error) {
	gen := c.Gen
	if gen.Icon {
		outPath, err := GeneraICON(gen.IconFilePath, gen.OutPath)
		if err != nil {
			return err
		}
		term.Section.Println("ICON save path", outPath)
	}
	if gen.Syso {
		info := project.Info{}
		infoData, err := ioutil.ReadFile(gen.InfoFilePath)
		if err != nil {
			return errors.New("Read infoFile error:" + err.Error())
		}
		err = json.Unmarshal(infoData, &info)
		if err != nil {
			return err
		}
		outPath, err := GeneraSYSO(gen.Name, gen.IconFilePath, gen.ManifestFilePath, gen.OutPath, gen.Arch, info)
		if err != nil {
			return err
		}
		term.Section.Println("SYSO save path", outPath)
	}
	return nil
}
