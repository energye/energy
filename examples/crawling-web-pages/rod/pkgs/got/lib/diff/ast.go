// Package diff ...
package diff

// TokenLine of tokens
type TokenLine struct {
	Type   Type
	Tokens []*Token
}

// ParseTokenLines of tokens
func ParseTokenLines(ts []*Token) []*TokenLine {
	list := []*TokenLine{}
	var l *TokenLine
	for _, t := range ts {
		switch t.Type {
		case SameSymbol, AddSymbol, DelSymbol:
			l = &TokenLine{}
			list = append(list, l)
			l.Type = t.Type
		}
		l.Tokens = append(l.Tokens, t)
	}
	return list
}

// SpreadTokenLines to tokens
func SpreadTokenLines(lines []*TokenLine) []*Token {
	out := []*Token{}
	for _, l := range lines {
		out = append(out, l.Tokens...)
	}
	return out
}
