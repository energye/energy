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
	"fmt"
	"github.com/pterm/pterm"
	"sort"

	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

var (
	// DefaultInteractiveMultiselect is the default InteractiveMultiselect printer.
	DefaultInteractiveMultiselect = InteractiveMultiselectPrinter{
		TextStyle:      &pterm.ThemeDefault.PrimaryStyle,
		DefaultText:    "Please select your options",
		Options:        []string{},
		OptionStyle:    &pterm.ThemeDefault.DefaultText,
		DefaultOptions: []string{},
		MaxHeight:      5,
		Selector:       ">",
		SelectorStyle:  &pterm.ThemeDefault.SecondaryStyle,
		Filter:         true,
		KeySelect:      keys.Space,
		KeyConfirm:     keys.Enter,
		Up:             'w',
		Down:           's',
		Left:           'a',
		Right:          'd',
		Checkmark:      &pterm.ThemeDefault.Checkmark,
	}
)

// InteractiveMultiselectPrinter is a printer for interactive multiselect menus.
type InteractiveMultiselectPrinter struct {
	DefaultText     string
	TextStyle       *pterm.Style
	Options         []string
	OptionStyle     *pterm.Style
	DefaultOptions  []string
	MaxHeight       int
	Selector        string
	SelectorStyle   *pterm.Style
	Filter          bool
	Checkmark       *pterm.Checkmark
	OnInterruptFunc func()

	selectedOption        int
	selectedOptions       []int
	text                  string
	fuzzySearchString     string
	fuzzySearchMatches    []string
	displayedOptions      []string
	displayedOptionsStart int
	displayedOptionsEnd   int

	// KeySelect is the select key. It cannot be keys.Space when Filter is enabled.
	KeySelect keys.KeyCode

	// KeyConfirm is the confirm key. It cannot be keys.Space when Filter is enabled.
	KeyConfirm keys.KeyCode

	// 自定义上下左右选择键
	Up    keys.KeyCode
	Down  keys.KeyCode
	Left  keys.KeyCode
	Right keys.KeyCode
}

// WithOptions sets the options.
func (p InteractiveMultiselectPrinter) WithOptions(options []string) *InteractiveMultiselectPrinter {
	p.Options = options
	return &p
}

func (p *InteractiveMultiselectPrinter) CheckmarkANSI() {
	if !IsWindows10 {
		p.Checkmark.Checked = "+"
		p.Checkmark.Unchecked = "-"
	}
}

// WithDefaultOptions sets the default options.
func (p InteractiveMultiselectPrinter) WithDefaultOptions(options []string) *InteractiveMultiselectPrinter {
	p.DefaultOptions = options
	return &p
}

// WithDefaultText sets the default text.
func (p InteractiveMultiselectPrinter) WithDefaultText(text string) *InteractiveMultiselectPrinter {
	p.DefaultText = text
	return &p
}

// WithMaxHeight sets the maximum height of the select menu.
func (p InteractiveMultiselectPrinter) WithMaxHeight(maxHeight int) *InteractiveMultiselectPrinter {
	p.MaxHeight = maxHeight
	return &p
}

// WithFilter sets the Filter option
func (p InteractiveMultiselectPrinter) WithFilter(b ...bool) *InteractiveMultiselectPrinter {
	p.Filter = WithBoolean(b)
	return &p
}

// WithKeySelect sets the confirm key
// It cannot be keys.Space when Filter is enabled.
func (p InteractiveMultiselectPrinter) WithKeySelect(keySelect keys.KeyCode) *InteractiveMultiselectPrinter {
	p.KeySelect = keySelect
	return &p
}

// WithKeyConfirm sets the confirm key
// It cannot be keys.Space when Filter is enabled.
func (p InteractiveMultiselectPrinter) WithKeyConfirm(keyConfirm keys.KeyCode) *InteractiveMultiselectPrinter {
	p.KeyConfirm = keyConfirm
	return &p
}

// WithCheckmark sets the checkmark
func (p InteractiveMultiselectPrinter) WithCheckmark(checkmark *pterm.Checkmark) *InteractiveMultiselectPrinter {
	p.Checkmark = checkmark
	return &p
}

// OnInterrupt sets the function to execute on exit of the input reader
func (p InteractiveMultiselectPrinter) WithOnInterruptFunc(exitFunc func()) *InteractiveMultiselectPrinter {
	p.OnInterruptFunc = exitFunc
	return &p
}

// Show shows the interactive multiselect menu and returns the selected entry.
func (p *InteractiveMultiselectPrinter) Show(text ...string) ([]string, error) {
	// should be the first defer statement to make sure it is executed last
	// and all the needed cleanup can be done before
	cancel, exit := NewCancelationSignal(p.OnInterruptFunc)
	defer exit()

	if len(text) == 0 || pterm.Sprint(text[0]) == "" {
		text = []string{p.DefaultText}
	}

	p.text = p.TextStyle.Sprint(text[0])
	p.fuzzySearchMatches = append([]string{}, p.Options...)

	if p.MaxHeight == 0 {
		p.MaxHeight = DefaultInteractiveMultiselect.MaxHeight
	}

	maxHeight := p.MaxHeight
	if maxHeight > len(p.fuzzySearchMatches) {
		maxHeight = len(p.fuzzySearchMatches)
	}

	if len(p.Options) == 0 {
		return nil, fmt.Errorf("no options provided")
	}

	p.displayedOptions = append([]string{}, p.fuzzySearchMatches[:maxHeight]...)
	p.displayedOptionsStart = 0
	p.displayedOptionsEnd = maxHeight

	for _, option := range p.DefaultOptions {
		p.selectOption(option)
	}

	area, err := pterm.DefaultArea.Start(p.renderSelectMenu())
	defer area.Stop()
	if err != nil {
		return nil, fmt.Errorf("could not start area: %w", err)
	}

	if p.Filter && (p.KeyConfirm == keys.Space || p.KeySelect == keys.Space) {
		return nil, fmt.Errorf("if filter/search is active, keys.Space can not be used for KeySelect or KeyConfirm")
	}

	area.Update(p.renderSelectMenu())

	cursor.Hide()
	defer cursor.Show()
	err = keyboard.Listen(func(keyInfo keys.Key) (stop bool, err error) {
		key := keyInfo.Code
		if p.MaxHeight > len(p.fuzzySearchMatches) {
			maxHeight = len(p.fuzzySearchMatches)
		} else {
			maxHeight = p.MaxHeight
		}
		var up = func() {
			if p.selectedOption > 0 {
				p.selectedOption--
				if p.selectedOption < p.displayedOptionsStart {
					p.displayedOptionsStart--
					p.displayedOptionsEnd--
					if p.displayedOptionsStart < 0 {
						p.displayedOptionsStart = 0
						p.displayedOptionsEnd = maxHeight
					}
					p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)
				}
			} else {
				p.selectedOption = len(p.fuzzySearchMatches) - 1
				p.displayedOptionsStart = len(p.fuzzySearchMatches) - maxHeight
				p.displayedOptionsEnd = len(p.fuzzySearchMatches)
				p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)
			}
			area.Update(p.renderSelectMenu())
		}
		var down = func() {
			p.displayedOptions = p.fuzzySearchMatches[:maxHeight]
			if p.selectedOption < len(p.fuzzySearchMatches)-1 {
				p.selectedOption++
				if p.selectedOption >= p.displayedOptionsEnd {
					p.displayedOptionsStart++
					p.displayedOptionsEnd++
					p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)
				}
			} else {
				p.selectedOption = 0
				p.displayedOptionsStart = 0
				p.displayedOptionsEnd = maxHeight
				p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)
			}

			area.Update(p.renderSelectMenu())
		}
		var left = func() {
			p.selectedOptions = []int{}
			area.Update(p.renderSelectMenu())
		}
		var right = func() {
			p.selectedOptions = []int{}
			for i := 0; i < len(p.Options); i++ {
				p.selectedOptions = append(p.selectedOptions, i)
			}
			area.Update(p.renderSelectMenu())
		}
		switch string(keyInfo.Runes) {
		case string(rune(p.Up)):
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			up()
		case string(rune(p.Down)):
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			down()
		case string(rune(p.Left)):
			left()
		case string(rune(p.Right)):
			right()
		}

		switch key {
		case p.KeyConfirm:
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			area.Update(p.renderFinishedMenu())
			return true, nil
		case p.KeySelect:
			if len(p.fuzzySearchMatches) > 0 {
				// Select option if not already selected
				p.selectOption(p.fuzzySearchMatches[p.selectedOption])
			}
			area.Update(p.renderSelectMenu())
		case keys.RuneKey:
			if p.Filter {
				// Fuzzy search for options
				// append to fuzzy search string
				p.fuzzySearchString += keyInfo.String()
				p.selectedOption = 0
				p.displayedOptionsStart = 0
				p.displayedOptionsEnd = maxHeight
				p.displayedOptions = append([]string{}, p.fuzzySearchMatches[:maxHeight]...)
			}
			area.Update(p.renderSelectMenu())
		case keys.Space:
			if p.Filter {
				p.fuzzySearchString += " "
				p.selectedOption = 0
				area.Update(p.renderSelectMenu())
			}
		case keys.Backspace:
			// Remove last character from fuzzy search string
			if len(p.fuzzySearchString) > 0 {
				// Handle UTF-8 characters
				p.fuzzySearchString = string([]rune(p.fuzzySearchString)[:len([]rune(p.fuzzySearchString))-1])
			}

			if p.fuzzySearchString == "" {
				p.fuzzySearchMatches = append([]string{}, p.Options...)
			}

			p.renderSelectMenu()

			if len(p.fuzzySearchMatches) > p.MaxHeight {
				maxHeight = p.MaxHeight
			} else {
				maxHeight = len(p.fuzzySearchMatches)
			}

			p.selectedOption = 0
			p.displayedOptionsStart = 0
			p.displayedOptionsEnd = maxHeight
			p.displayedOptions = append([]string{}, p.fuzzySearchMatches[p.displayedOptionsStart:p.displayedOptionsEnd]...)

			area.Update(p.renderSelectMenu())
		case keys.Left:
			// Unselect all options
			left()
		case keys.Right:
			// Select all options
			right()
		case keys.Up:
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			up()
		case keys.Down, p.Down:
			if len(p.fuzzySearchMatches) == 0 {
				return false, nil
			}
			down()
		case keys.CtrlC:
			cancel()
			return true, nil
		}

		return false, nil
	})
	if err != nil {
		pterm.Error.Println(err)
		return nil, fmt.Errorf("failed to start keyboard listener: %w", err)
	}

	var result []string
	for _, selectedOption := range p.selectedOptions {
		result = append(result, p.Options[selectedOption])
	}

	return result, nil
}

func (p InteractiveMultiselectPrinter) findOptionByText(text string) int {
	for i, option := range p.Options {
		if option == text {
			return i
		}
	}
	return -1
}

func (p *InteractiveMultiselectPrinter) isSelected(optionText string) bool {
	for _, selectedOption := range p.selectedOptions {
		if p.Options[selectedOption] == optionText {
			return true
		}
	}

	return false
}

func (p *InteractiveMultiselectPrinter) selectOption(optionText string) {
	if p.isSelected(optionText) {
		// Remove from selected options
		for i, selectedOption := range p.selectedOptions {
			if p.Options[selectedOption] == optionText {
				p.selectedOptions = append(p.selectedOptions[:i], p.selectedOptions[i+1:]...)
				break
			}
		}
	} else {
		// Add to selected options
		p.selectedOptions = append(p.selectedOptions, p.findOptionByText(optionText))
	}
}

func (p *InteractiveMultiselectPrinter) renderSelectMenu() string {
	var content string
	content += pterm.Sprintf("%s: %s\n", p.text, p.fuzzySearchString)

	// find options that match fuzzy search string
	rankedResults := fuzzy.RankFindFold(p.fuzzySearchString, p.Options)
	// map rankedResults to fuzzySearchMatches
	p.fuzzySearchMatches = []string{}
	if len(rankedResults) != len(p.Options) {
		sort.Sort(rankedResults)
	}
	for _, result := range rankedResults {
		p.fuzzySearchMatches = append(p.fuzzySearchMatches, result.Target)
	}

	indexMapper := make([]string, len(p.fuzzySearchMatches))
	for i := 0; i < len(p.fuzzySearchMatches); i++ {
		// if in displayed options range
		if i >= p.displayedOptionsStart && i < p.displayedOptionsEnd {
			indexMapper[i] = p.fuzzySearchMatches[i]
		}
	}

	for i, option := range indexMapper {
		if option == "" {
			continue
		}
		var checkmark string
		if p.isSelected(option) {
			checkmark = fmt.Sprintf("[%s]", p.Checkmark.Checked)
		} else {
			checkmark = fmt.Sprintf("[%s]", p.Checkmark.Unchecked)
		}
		if i == p.selectedOption {
			content += pterm.Sprintf("%s %s %s\n", p.renderSelector(), checkmark, option)
		} else {
			content += pterm.Sprintf("  %s %s\n", checkmark, option)
		}
	}

	help := fmt.Sprintf("Operate:\n  %s: %s | %s: %s\n  left: [%s, %s] | right: [%s, %s]\n  up: [%s, %s] | down: [%s, %s]", p.KeySelect, pterm.Bold.Sprint("select"),
		p.KeyConfirm, pterm.Bold.Sprint("confirm"), string(rune(p.Left)), pterm.Bold.Sprint("none"), string(rune(p.Right)), pterm.Bold.Sprint("all"), string(rune(p.Up)), pterm.Bold.Sprint("up"), string(rune(p.Down)), pterm.Bold.Sprint("down"))
	if p.Filter {
		help += fmt.Sprintf("| type to %s", pterm.Bold.Sprint("filter"))
	}
	content += pterm.ThemeDefault.SecondaryStyle.Sprintfln(help)

	return content
}

func (p InteractiveMultiselectPrinter) renderFinishedMenu() string {
	var content string
	content += pterm.Sprintf("%s: %s\n", p.text, p.fuzzySearchString)
	for _, option := range p.selectedOptions {
		content += pterm.Sprintf("  %s %s\n", p.renderSelector(), p.Options[option])
	}

	return content
}

func (p InteractiveMultiselectPrinter) renderSelector() string {
	return p.SelectorStyle.Sprint(p.Selector)
}
