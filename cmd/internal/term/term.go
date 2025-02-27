//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

package term

import (
	"github.com/pterm/pterm"
	"os"
)

const (
	Major = 2
	Minor = 5
	Build = 4
)

var Logger *pterm.Logger
var TermOut = new(Termout)
var Section *pterm.SectionPrinter
var IsWindows10 = true

func init() {
	// logger
	Logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	TermOut = new(Termout)
	Logger.ShowTime = false
	Logger.ShowCaller = false

	// Section
	Section = &pterm.SectionPrinter{
		Style:           &pterm.Style{pterm.Bold, pterm.FgGreen},
		Level:           0,
		IndentCharacter: "",
	}
}

type Termout struct {
}

func (m *Termout) Write(p []byte) (n int, err error) {
	Section.Println(string(p))
	return len(p), nil
}

func GoENERGY() {
	pterm.Println(pterm.LightBlue("      GO\n") + pterm.LightBlue("    ENERGY"))
}

func BoxPrintln(a ...interface{}) {
	pterm.DefaultBox.Println(a...)
}

func Println(a ...interface{}) {
	pterm.Println(a...)
}

func TextInputWith(text, delimiter string) string {
	tiw := pterm.DefaultInteractiveTextInput.WithMultiLine(false).WithOnInterruptFunc(func() {
		os.Exit(1)
	})
	tiw.DefaultText = text
	tiw.Delimiter = delimiter
	result, err := tiw.Show()
	if err != nil {
		return ""
	}
	return result
}

// WithBoolean helps an option setter (WithXXX(b ...bool) to return true, if no boolean is set, but false if it's explicitly set to false.
func WithBoolean(b []bool) bool {
	if len(b) == 0 {
		b = append(b, true)
	}
	return b[0]
}

// NewCancelationSignal for keeping track of a cancelation
func NewCancelationSignal(interruptFunc func()) (func(), func()) {
	canceled := false

	cancel := func() {
		canceled = true
	}

	exit := func() {
		if canceled && interruptFunc != nil {
			interruptFunc()
		}
	}

	return cancel, exit
}
