// Package gop ...
package gop

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Stdout is the default stdout for gop.P .
var Stdout io.Writer = os.Stdout

const indentUnit = "    "

// Theme to color values
type Theme func(t Type) []Style

// ThemeDefault colors for Sprint
var ThemeDefault = func(t Type) []Style {
	switch t {
	case TypeName:
		return []Style{Cyan}
	case Bool, Chan:
		return []Style{Blue}
	case Rune, Byte, String:
		return []Style{Yellow}
	case Number:
		return []Style{Green}
	case Func:
		return []Style{Magenta}
	case Comment:
		return []Style{White}
	case Nil:
		return []Style{Red}
	case Error:
		return []Style{Underline, Red}
	default:
		return []Style{None}
	}
}

// ThemeNone colors for Sprint
var ThemeNone = func(t Type) []Style {
	return []Style{None}
}

// F is a shortcut for Format with color
func F(v interface{}) string {
	return Format(Tokenize(v), ThemeDefault)
}

// P pretty print the values
func P(values ...interface{}) error {
	list := []interface{}{}
	for _, v := range values {
		list = append(list, F(v))
	}

	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc).Name()
	cwd, _ := os.Getwd()
	file, _ = filepath.Rel(cwd, file)
	tpl := Stylize("// %s %s:%d (%s)\n", ThemeDefault(Comment))
	_, _ = fmt.Fprintf(Stdout, tpl, time.Now().Format(time.RFC3339Nano), file, line, fn)

	_, err := fmt.Fprintln(Stdout, list...)
	return err
}

// Plain is a shortcut for Format with plain color
func Plain(v interface{}) string {
	return Format(Tokenize(v), ThemeNone)
}

// Format a list of tokens
func Format(ts []*Token, theme Theme) string {
	out := ""
	depth := 0
	for i, t := range ts {
		if oneOf(t.Type, SliceOpen, MapOpen, StructOpen) {
			depth++
		}
		if i < len(ts)-1 && oneOf(ts[i+1].Type, SliceClose, MapClose, StructClose) {
			depth--
		}

		styles := theme(t.Type)

		switch t.Type {
		case SliceOpen, MapOpen, StructOpen:
			out += Stylize(t.Literal, styles) + "\n"
		case SliceItem, MapKey, StructKey:
			out += strings.Repeat(indentUnit, depth)
		case Colon, InlineComma, Chan:
			out += Stylize(t.Literal, styles) + " "
		case Comma:
			out += Stylize(t.Literal, styles) + "\n"
		case SliceClose, MapClose, StructClose:
			if strings.HasSuffix(out, "{\n") {
				out = out[:len(out)-1]
				out += Stylize(t.Literal, styles)
			} else {
				out += strings.Repeat(indentUnit, depth) + Stylize(t.Literal, styles)
			}
		case String:
			out += Stylize(readableStr(depth, t.Literal), styles)
		default:
			out += Stylize(t.Literal, styles)
		}
	}

	return out
}

func oneOf(t Type, list ...Type) bool {
	for _, el := range list {
		if t == el {
			return true
		}
	}
	return false
}

// To make multi-line string block more human readable.
// Split newline into two strings, convert "\t" into tab.
// Such as foramt string: "line one \n\t line two" into:
//
//	"line one \n" +
//	"	 line two"
func readableStr(depth int, s string) string {
	if ((len(s) > LongStringLen) || strings.Contains(s, "\n") || strings.Contains(s, `"`)) && !strings.Contains(s, "`") {
		return "`" + s + "`"
	}

	s = fmt.Sprintf("%#v", s)
	s, _ = replaceEscaped(s, 't', "	")

	indent := strings.Repeat(indentUnit, depth+1)
	if n, has := replaceEscaped(s, 'n', "\\n\" +\n"+indent+"\""); has {
		return "\"\" +\n" + indent + n
	}

	return s
}

// We use a simple state machine to replace escaped char like "\n"
func replaceEscaped(s string, escaped rune, new string) (string, bool) {
	type State int
	const (
		init State = iota
		prematch
		match
	)

	state := init
	out := ""
	buf := ""
	has := false

	onInit := func(r rune) {
		state = init
		out += buf + string(r)
		buf = ""
	}

	onPrematch := func() {
		state = prematch
		buf = "\\"
	}

	onEscape := func() {
		state = match
		out += new
		buf = ""
		has = true
	}

	for _, r := range s {
		switch state {
		case prematch:
			switch r {
			case escaped:
				onEscape()
			default:
				onInit(r)
			}

		case match:
			switch r {
			case '\\':
				onPrematch()
			default:
				onInit(r)
			}

		default:
			switch r {
			case '\\':
				onPrematch()
			default:
				onInit(r)
			}
		}
	}

	return out, has
}
