package logger

import (
	"io"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const timeFmt = "2006-01-02 15:04:05.000"

type Level int32

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

type Field struct {
	Key  string
	Kind byte
	S    string
	I    int64
	F    float64
	B    bool
	A    any
}

func String(k, v string) Field { return Field{Key: k, Kind: 's', S: v} }

func Int(k string, v int) Field { return Field{Key: k, Kind: 'i', I: int64(v)} }

func Int64(k string, v int64) Field { return Field{Key: k, Kind: 'i', I: v} }

func Float64(k string, v float64) Field { return Field{Key: k, Kind: 'f', F: v} }

func Bool(k string, v bool) Field { return Field{Key: k, Kind: 'b', B: v} }

func Any(k string, v any) Field { return Field{Key: k, Kind: 'a', A: v} }

func Err(err error) Field {
	if err == nil {
		return String("error", "")
	}
	return String("error", err.Error())
}

type Config struct {
	Level  Level
	Output io.Writer
	Caller bool
}

type Logger struct {
	out    io.Writer
	level  atomic.Int32
	caller bool
	ch     chan []byte
	wg     sync.WaitGroup
	pool   sync.Pool
	closed atomic.Bool
}

var def atomic.Pointer[Logger]

func New(c Config) *Logger {
	if c.Output == nil {
		c.Output = os.Stdout
	}
	l := &Logger{out: c.Output, caller: c.Caller, ch: make(chan []byte, 1024), pool: sync.Pool{New: func() any { b := make([]byte, 0, 512); return &b }}}
	l.level.Store(int32(c.Level))
	l.wg.Add(1)
	go l.writer()
	return l
}

func init() {
	def.Store(New(Config{Level: InfoLevel, Output: os.Stdout}))
}

func SetDefault(l *Logger) {
	def.Store(l)
}

func L() *Logger {
	return def.Load()
}
func (l *Logger) SetLevel(v Level) {
	l.level.Store(int32(v))
}
func (l *Logger) Close() {
	if !l.closed.CompareAndSwap(false, true) {
		return
	}
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) writer() {
	defer l.wg.Done()
	for b := range l.ch {
		_, _ = l.out.Write(b)
		p := &b
		*p = (*p)[:0]
		l.pool.Put(p)
	}
}

func Debug(msg string, args ...any) {
	L().logAny(DebugLevel, msg, args...)
}
func Info(msg string, args ...any) {
	L().logAny(InfoLevel, msg, args...)
}
func Warn(msg string, args ...any) {
	L().logAny(WarnLevel, msg, args...)
}
func Error(msg string, args ...any) {
	L().logAny(ErrorLevel, msg, args...)
}
func DebugF(msg string, fields ...Field) {
	L().logFields(DebugLevel, msg, fields...)
}
func InfoF(msg string, fields ...Field) {
	L().logFields(InfoLevel, msg, fields...)
}
func WarnF(msg string, fields ...Field) {
	L().logFields(WarnLevel, msg, fields...)
}
func ErrorF(msg string, fields ...Field) {
	L().logFields(ErrorLevel, msg, fields...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logAny(DebugLevel, msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logAny(InfoLevel, msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logAny(WarnLevel, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.logAny(ErrorLevel, msg, args...)
}

func (l *Logger) DebugF(msg string, fields ...Field) {
	l.logFields(DebugLevel, msg, fields...)
}

func (l *Logger) InfoF(msg string, fields ...Field) {
	l.logFields(InfoLevel, msg, fields...)
}

func (l *Logger) WarnF(msg string, fields ...Field) {
	l.logFields(WarnLevel, msg, fields...)
}

func (l *Logger) ErrorF(msg string, fields ...Field) {
	l.logFields(ErrorLevel, msg, fields...)
}

func (l *Logger) logAny(level Level, msg string, args ...any) {
	if int32(level) < l.level.Load() {
		return
	}
	bp := l.pool.Get().(*[]byte)
	b := (*bp)[:0]
	b = append(b, '[')
	b = appendLevel(b, level)
	b = append(b, "] "...)
	b = time.Now().AppendFormat(b, timeFmt)
	b = append(b, ' ')
	if l.caller {
		_, file, line, _ := runtime.Caller(2)
		b = appendCaller(b, file, line)
	}
	b = appendQuoted(b, msg)
	if len(args) > 0 {
		b = appendArgs(b, args...)
	}

	b = append(b, '\n')
	select {
	case l.ch <- b:
	default:
		_, _ = l.out.Write(b)
		*bp = (*bp)[:0]
		l.pool.Put(bp)
	}
}

func (l *Logger) logFields(level Level, msg string, fields ...Field) {
	if int32(level) < l.level.Load() {
		return
	}
	bp := l.pool.Get().(*[]byte)
	b := (*bp)[:0]
	b = append(b, '[')
	b = appendLevel(b, level)
	b = append(b, "] "...)
	b = time.Now().AppendFormat(b, timeFmt)
	b = append(b, ' ')
	if l.caller {
		_, file, line, _ := runtime.Caller(2)
		b = appendCaller(b, file, line)
	}
	b = appendQuoted(b, msg)
	for _, f := range fields {
		b = append(b, ' ')
		b = append(b, f.Key...)
		b = append(b, '=')
		b = appendField(b, f)
	}
	b = append(b, '\n')
	select {
	case l.ch <- b:
	default:
		_, _ = l.out.Write(b)
		*bp = (*bp)[:0]
		l.pool.Put(bp)
	}
}

func appendLevel(b []byte, lv Level) []byte {
	switch lv {
	case DebugLevel:
		return append(b, "DEBUG"...)
	case InfoLevel:
		return append(b, "INFO"...)
	case WarnLevel:
		return append(b, "WARN"...)
	default:
		return append(b, "ERROR"...)
	}
}

func appendCaller(b []byte, file string, line int) []byte {
	for i := len(file) - 1; i >= 0; i-- {
		if file[i] == '/' || file[i] == '\\' {
			file = file[i+1:]
			break
		}
	}
	b = append(b, file...)
	b = append(b, ':')
	b = strconv.AppendInt(b, int64(line), 10)
	b = append(b, ' ')
	return b
}

func appendArgs(b []byte, args ...any) []byte {
	b = append(b, " |"...)
	if len(args)%2 == 0 && len(args) >= 2 {
		if _, ok := args[0].(string); ok {
			return appendKeyValuePairs(b, args...)
		}
	}
	for i, arg := range args {
		if i > 0 {
			b = append(b, ' ')
		}
		b = appendValue(b, arg)
	}
	return b
}

func appendKeyValuePairs(b []byte, args ...any) []byte {
	for i := 0; i < len(args); i += 2 {
		if i > 0 {
			b = append(b, ' ')
		}

		key, _ := args[i].(string)
		b = append(b, key...)
		b = append(b, '=')
		b = appendValue(b, args[i+1])
	}
	return b
}

func appendValue(b []byte, v any) []byte {
	if v == nil {
		return append(b, "null"...)
	}
	switch val := v.(type) {
	case string:
		return appendQuoted(b, val)
	case int:
		return strconv.AppendInt(b, int64(val), 10)
	case int8:
		return strconv.AppendInt(b, int64(val), 10)
	case int16:
		return strconv.AppendInt(b, int64(val), 10)
	case int32:
		return strconv.AppendInt(b, int64(val), 10)
	case int64:
		return strconv.AppendInt(b, val, 10)
	case uint:
		return strconv.AppendUint(b, uint64(val), 10)
	case uint8:
		return strconv.AppendUint(b, uint64(val), 10)
	case uint16:
		return strconv.AppendUint(b, uint64(val), 10)
	case uint32:
		return strconv.AppendUint(b, uint64(val), 10)
	case uint64:
		return strconv.AppendUint(b, val, 10)
	case float32:
		return strconv.AppendFloat(b, float64(val), 'f', -1, 32)
	case float64:
		return strconv.AppendFloat(b, val, 'f', -1, 64)
	case bool:
		return strconv.AppendBool(b, val)
	case error:
		return appendQuoted(b, val.Error())
	default:
		return appendQuoted(b, simpleToString(val))
	}
}

func appendField(b []byte, f Field) []byte {
	switch f.Kind {
	case 's':
		return appendQuoted(b, f.S)
	case 'i':
		return strconv.AppendInt(b, f.I, 10)
	case 'f':
		return strconv.AppendFloat(b, f.F, 'f', -1, 64)
	case 'b':
		return strconv.AppendBool(b, f.B)
	default:
		return appendQuoted(b, toString(f.A))
	}
}

func appendQuoted(b []byte, s string) []byte {
	b = append(b, '"')
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '"' || c == '\\' {
			b = append(b, '\\')
		}
		if c == '\n' {
			b = append(b, '\\', 'n')
			continue
		}
		if c == '\r' {
			b = append(b, '\\', 'r')
			continue
		}
		if c == '\t' {
			b = append(b, '\\', 't')
			continue
		}
		b = append(b, c)
	}
	return append(b, '"')
}

func toString(v any) string {
	if v == nil {
		return "null"
	}
	if s, ok := v.(string); ok {
		return s
	}
	if e, ok := v.(error); ok {
		return e.Error()
	}
	return "?"
}

func simpleToString(v any) string {
	if v == nil {
		return "null"
	}
	if s, ok := v.(string); ok {
		return s
	}
	if e, ok := v.(error); ok {
		return e.Error()
	}
	return "?"
}
