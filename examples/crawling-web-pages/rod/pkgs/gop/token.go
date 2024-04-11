package gop

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// LongStringLen is the length of that will be treated as long string
var LongStringLen = 16

// LongBytesLen is the length of that will be treated as long bytes
var LongBytesLen = 16

// Type of token
type Type int

const (
	// Nil type
	Nil Type = iota
	// Bool type
	Bool
	// Number type
	Number
	// Float type
	Float
	// Complex type
	Complex
	// String type
	String
	// Byte type
	Byte
	// Rune type
	Rune
	// Chan type
	Chan
	// Func type
	Func
	// Error type
	Error

	// Comment type
	Comment

	// TypeName type
	TypeName

	// ParenOpen type
	ParenOpen
	// ParenClose type
	ParenClose

	// Dot type
	Dot
	// And type
	And

	// SliceOpen type
	SliceOpen
	// SliceItem type
	SliceItem
	// InlineComma type
	InlineComma
	// Comma type
	Comma
	// SliceClose type
	SliceClose

	// MapOpen type
	MapOpen
	// MapKey type
	MapKey
	// Colon type
	Colon
	// MapClose type
	MapClose

	// StructOpen type
	StructOpen
	// StructKey type
	StructKey
	// StructField type
	StructField
	// StructClose type
	StructClose
)

// Token represents a symbol in value layout
type Token struct {
	Type    Type
	Literal string
}

// Tokenize a random Go value
func Tokenize(v interface{}) []*Token {
	return tokenize(newContext(), reflect.ValueOf(v))
}

type path []interface{}

func (p path) tokens() []*Token {
	ts := []*Token{}
	for i, seg := range p {
		ts = append(ts, tokenize(newContext(), reflect.ValueOf(seg))...)
		if i < len(p)-1 {
			ts = append(ts, &Token{InlineComma, ","})
		}
	}
	return ts
}

type context struct {
	global map[uintptr]path
	path   path
}

func newContext() context {
	return context{global: map[uintptr]path{}, path: path{}}
}

func (ctx context) add(p interface{}) context {
	if v := reflect.ValueOf(p); v.Kind() == reflect.Ptr {
		p = v.Pointer()
	}

	return context{
		global: ctx.global,
		path:   append(ctx.path, p),
	}
}

func (ctx context) has(prefix path) bool {
	if len(ctx.path) < len(prefix) {
		return false
	}

	for i := range prefix {
		if !reflect.DeepEqual(prefix[i], ctx.path[i]) {
			return false
		}
	}
	return true
}

func (ctx context) circular(v reflect.Value) []*Token {
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice:
		ptr := v.Pointer()
		if ptr == 0 {
			return nil
		}

		if prev, has := ctx.global[ptr]; has && ctx.has(prev) {
			ts := []*Token{{Func, SymbolCircular}, {ParenOpen, "("}}
			ts = append(ts, prev.tokens()...)
			return append(ts, &Token{ParenClose, ")"}, &Token{Dot, "."},
				&Token{ParenOpen, "("}, typeName(v.Type().String()), &Token{ParenClose, ")"})
		}
		ctx.global[ptr] = ctx.path
	}

	return nil
}

func tokenize(ctx context, v reflect.Value) []*Token {
	if ts, has := tokenizeSpecial(v); has {
		return ts
	}

	if ts := ctx.circular(v); ts != nil {
		return ts
	}

	t := &Token{Nil, ""}

	switch v.Kind() {
	case reflect.Interface:
		return tokenize(ctx, v.Elem())

	case reflect.Bool:
		t.Type = Bool
		if v.Bool() {
			t.Literal = "true"
		} else {
			t.Literal = "false"
		}

	case reflect.String:
		return tokenizeString(v)

	case reflect.Chan:
		if v.Cap() == 0 {
			return []*Token{{Func, "make"}, {ParenOpen, "("},
				{Chan, "chan"}, typeName(v.Type().Elem().String()), {ParenClose, ")"},
				{Comment, wrapComment(formatUintptr(v.Pointer()))}}
		}
		return []*Token{{Func, "make"}, {ParenOpen, "("}, {Chan, "chan"},
			typeName(v.Type().Elem().String()), {InlineComma, ","},
			{Number, strconv.FormatInt(int64(v.Cap()), 10)}, {ParenClose, ")"},
			{Comment, wrapComment(formatUintptr(v.Pointer()))}}

	case reflect.Func:
		return []*Token{{ParenOpen, "("}, {TypeName, v.Type().String()},
			{ParenClose, ")"}, {ParenOpen, "("}, {Nil, "nil"}, {ParenClose, ")"},
			{Comment, wrapComment(formatUintptr(v.Pointer()))}}

	case reflect.Ptr:
		return tokenizePtr(ctx, v)

	case reflect.UnsafePointer:
		return []*Token{typeName("unsafe.Pointer"), {ParenOpen, "("}, typeName("uintptr"),
			{ParenOpen, "("}, {Number, formatUintptr(v.Pointer())}, {ParenClose, ")"}, {ParenClose, ")"}}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Uintptr, reflect.Complex64, reflect.Complex128:
		return tokenizeNumber(v)

	case reflect.Slice, reflect.Array, reflect.Map, reflect.Struct:
		return tokenizeCollection(ctx, v)
	}

	return []*Token{t}
}

func tokenizeSpecial(v reflect.Value) ([]*Token, bool) {
	if v.Kind() == reflect.Invalid {
		return []*Token{{Nil, "nil"}}, true
	} else if r, ok := v.Interface().(rune); ok && unicode.IsGraphic(r) {
		return []*Token{{Rune, strconv.QuoteRune(r)}}, true
	} else if b, ok := v.Interface().(byte); ok {
		return tokenizeByte(&Token{Nil, ""}, b), true
	} else if t, ok := v.Interface().(time.Time); ok {
		return tokenizeTime(t), true
	} else if d, ok := v.Interface().(time.Duration); ok {
		return tokenizeDuration(d), true
	}

	return tokenizeJSON(v)
}

func tokenizeCollection(ctx context, v reflect.Value) []*Token {
	ts := []*Token{}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		if data, ok := v.Interface().([]byte); ok {
			ts = append(ts, tokenizeBytes(data)...)
			break
		} else {
			ts = append(ts, typeName(v.Type().String()))
		}
		if v.Kind() == reflect.Slice && v.Cap() > 0 {
			ts = append(ts, &Token{Comment, formatLenCap(v.Len(), v.Cap())})
		}
		ts = append(ts, &Token{SliceOpen, "{"})
		for i := 0; i < v.Len(); i++ {
			el := v.Index(i)
			ts = append(ts, &Token{SliceItem, ""})
			ts = append(ts, tokenize(ctx.add(i), el)...)
			ts = append(ts, &Token{Comma, ","})
		}
		ts = append(ts, &Token{SliceClose, "}"})

	case reflect.Map:
		ts = append(ts, typeName(v.Type().String()))
		keys := v.MapKeys()
		sort.Slice(keys, func(i, j int) bool {
			return compare(keys[i].Interface(), keys[j].Interface()) < 0
		})
		if len(keys) > 1 {
			ts = append(ts, &Token{Comment, formatLenCap(len(keys), -1)})
		}
		ts = append(ts, &Token{MapOpen, "{"})
		for _, k := range keys {
			ctx := ctx.add(k.Interface())
			ts = append(ts, &Token{MapKey, ""})

			if k.Kind() == reflect.Interface && k.Elem().Kind() == reflect.Ptr {
				ts = append(ts, tokenizeMapKey(ctx, k)...)
			} else {
				ts = append(ts, tokenize(ctx, k)...)
			}

			ts = append(ts, &Token{Colon, ":"})
			ts = append(ts, tokenize(ctx, v.MapIndex(k))...)
			ts = append(ts, &Token{Comma, ","})
		}
		ts = append(ts, &Token{MapClose, "}"})

	case reflect.Struct:
		t := v.Type()

		ts = append(ts, typeName(t.String()))
		ts = append(ts, &Token{StructOpen, "{"})
		for i := 0; i < v.NumField(); i++ {
			name := t.Field(i).Name
			ts = append(ts, &Token{StructKey, ""})
			ts = append(ts, &Token{StructField, name})

			f := v.Field(i)
			if !f.CanInterface() {
				f = GetPrivateField(v, i)
			}
			ts = append(ts, &Token{Colon, ":"})
			ts = append(ts, tokenize(ctx.add(name), f)...)
			ts = append(ts, &Token{Comma, ","})
		}
		ts = append(ts, &Token{StructClose, "}"})
	}

	return ts
}

func tokenizeNumber(v reflect.Value) []*Token {
	t := &Token{Nil, ""}
	ts := []*Token{}
	tname := v.Type().String()

	switch v.Kind() {
	case reflect.Int:
		t.Type = Number
		t.Literal = strconv.FormatInt(v.Int(), 10)
		if tname != "int" {
			ts = append(ts, typeName(tname), &Token{ParenOpen, "("}, t, &Token{ParenClose, ")"})
		} else {
			ts = append(ts, t)
		}

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ts = append(ts, typeName(tname), &Token{ParenOpen, "("})
		t.Type = Number
		t.Literal = strconv.FormatInt(v.Int(), 10)
		ts = append(ts, t, &Token{ParenClose, ")"})

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ts = append(ts, typeName(tname), &Token{ParenOpen, "("})
		t.Type = Number
		t.Literal = strconv.FormatUint(v.Uint(), 10)
		ts = append(ts, t, &Token{ParenClose, ")"})

	case reflect.Float32:
		ts = append(ts, typeName(tname), &Token{ParenOpen, "("})
		t.Type = Number
		t.Literal = strconv.FormatFloat(v.Float(), 'f', -1, 32)
		ts = append(ts, t, &Token{ParenClose, ")"})

	case reflect.Float64:
		t.Type = Number
		t.Literal = strconv.FormatFloat(v.Float(), 'f', -1, 64)
		if !strings.Contains(t.Literal, ".") {
			t.Literal += ".0"
		}
		if tname != "float64" {
			ts = append(ts, typeName(tname), &Token{ParenOpen, "("}, t, &Token{ParenClose, ")"})
		} else {
			ts = append(ts, t)
		}

	case reflect.Complex64:
		ts = append(ts, typeName(tname), &Token{ParenOpen, "("})
		t.Type = Number
		t.Literal = strconv.FormatComplex(v.Complex(), 'f', -1, 64)
		t.Literal = t.Literal[1 : len(t.Literal)-1]
		ts = append(ts, t, &Token{ParenClose, ")"})

	case reflect.Complex128:
		t.Type = Number
		t.Literal = strconv.FormatComplex(v.Complex(), 'f', -1, 128)
		t.Literal = t.Literal[1 : len(t.Literal)-1]
		if tname != "complex128" {
			ts = append(ts, typeName(tname), &Token{ParenOpen, "("}, t, &Token{ParenClose, ")"})
		} else {
			ts = append(ts, t)
		}

	}

	return ts
}

func tokenizeByte(t *Token, b byte) []*Token {
	ts := []*Token{typeName("byte"), {ParenOpen, "("}}
	r := rune(b)
	if unicode.IsGraphic(r) {
		ts = append(ts, &Token{Byte, strconv.QuoteRune(r)})
	} else {
		ts = append(ts, &Token{Byte, "0x" + strconv.FormatUint(uint64(b), 16)})
	}
	return append(ts, &Token{ParenClose, ")"})
}

func tokenizeTime(t time.Time) []*Token {
	ext := GetPrivateFieldByName(reflect.ValueOf(t), "ext").Int()
	ts := []*Token{{Func, SymbolTime}, {ParenOpen, "("}}
	ts = append(ts, &Token{String, t.Format(time.RFC3339Nano)})
	ts = append(ts, &Token{InlineComma, ","}, &Token{Number, strconv.FormatInt(ext, 10)}, &Token{ParenClose, ")"})
	return ts
}

func tokenizeDuration(d time.Duration) []*Token {
	ts := []*Token{}
	ts = append(ts, typeName(SymbolDuration), &Token{ParenOpen, "("})
	ts = append(ts, &Token{String, d.String()})
	ts = append(ts, &Token{ParenClose, ")"})
	return ts
}

func tokenizeString(v reflect.Value) []*Token {
	s := v.String()
	ts := []*Token{{String, s}}
	if v.Len() >= LongStringLen {
		ts = append(ts, &Token{Comment, formatLenCap(len(s), -1)})
	}
	return ts
}

func tokenizeBytes(data []byte) []*Token {
	ts := []*Token{}

	if utf8.Valid(data) {
		s := string(data)
		ts = append(ts, typeName("[]byte"), &Token{ParenOpen, "("})
		ts = append(ts, &Token{String, s})
		ts = append(ts, &Token{ParenClose, ")"})
	} else {
		ts = append(ts, &Token{Func, SymbolBase64}, &Token{ParenOpen, "("})
		ts = append(ts, &Token{String, base64.StdEncoding.EncodeToString(data)})
		ts = append(ts, &Token{ParenClose, ")"})
	}
	if len(data) >= LongBytesLen {
		ts = append(ts, &Token{Comment, formatLenCap(len(data), -1)})
	}
	return ts
}

func tokenizeMapKey(ctx context, v reflect.Value) []*Token {
	ts := []*Token{}
	ts = append(ts,
		&Token{ParenOpen, "("}, typeName(v.Type().String()), &Token{ParenClose, ")"},
		&Token{ParenOpen, "("}, &Token{Nil, "nil"}, &Token{ParenClose, ")"},
		&Token{Comment, wrapComment(formatUintptr(v.Elem().Pointer()))},
	)
	return ts
}

func tokenizePtr(ctx context, v reflect.Value) []*Token {
	ts := []*Token{}

	if v.Elem().Kind() == reflect.Invalid {
		ts = append(ts,
			&Token{ParenOpen, "("}, typeName(v.Type().String()), &Token{ParenClose, ")"},
			&Token{ParenOpen, "("}, &Token{Nil, "nil"}, &Token{ParenClose, ")"})
		return ts
	}

	needFn := false

	switch v.Elem().Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		if _, ok := v.Elem().Interface().([]byte); ok {
			needFn = true
		}
	default:
		needFn = true
	}

	if needFn {
		ts = append(ts, &Token{Func, SymbolPtr}, &Token{ParenOpen, "("})
		ts = append(ts, tokenize(ctx, v.Elem())...)
		ts = append(ts, &Token{ParenClose, ")"}, &Token{Dot, "."}, &Token{ParenOpen, "("},
			typeName(v.Type().String()), &Token{ParenClose, ")"})
	} else {
		ts = append(ts, &Token{And, "&"})
		ts = append(ts, tokenize(ctx, v.Elem())...)
	}

	return ts
}

func tokenizeJSON(v reflect.Value) ([]*Token, bool) {
	var jv interface{}
	ts := []*Token{}
	s := ""
	if v.Kind() == reflect.String {
		s = v.String()
		err := json.Unmarshal([]byte(s), &jv)
		if err != nil {
			return nil, false
		}
		ts = append(ts, &Token{Func, SymbolJSONStr})
	} else if b, ok := v.Interface().([]byte); ok {
		err := json.Unmarshal(b, &jv)
		if err != nil {
			return nil, false
		}
		s = string(b)
		ts = append(ts, &Token{Func, SymbolJSONBytes})
	}

	_, isObj := jv.(map[string]interface{})
	_, isArr := jv.(map[string]interface{})

	if isObj || isArr {
		ts = append(ts, &Token{ParenOpen, "("})
		ts = append(ts, Tokenize(jv)...)
		ts = append(ts, &Token{InlineComma, ","},
			&Token{String, s}, &Token{ParenClose, ")"})
		return ts, true
	}

	return nil, false
}

func typeName(t string) *Token {
	return &Token{TypeName, t}
}
