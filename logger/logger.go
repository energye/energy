package logger

import (
	"fmt"
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

type Config struct {
	Level      Level
	Output     io.Writer
	Caller     bool
	BufferSize int // 默认 65536
	BatchSize  int // 默认 256
}

type Logger struct {
	out    io.Writer
	level  atomic.Int32
	caller bool

	ch        chan *[]byte
	batchSize int

	wg   sync.WaitGroup
	pool sync.Pool
	mu   sync.Mutex

	closed atomic.Bool
}

var def atomic.Pointer[Logger]

func New(c Config) *Logger {
	if c.Output == nil {
		c.Output = os.Stdout
	}
	if c.BufferSize <= 0 {
		c.BufferSize = 65536
	}
	if c.BatchSize <= 0 {
		c.BatchSize = 256
	}

	l := &Logger{
		out:       c.Output,
		caller:    c.Caller,
		ch:        make(chan *[]byte, c.BufferSize),
		batchSize: c.BatchSize,
		pool: sync.Pool{
			New: func() any {
				b := make([]byte, 0, 512)
				return &b
			},
		},
	}

	l.level.Store(int32(c.Level))
	l.wg.Add(1)
	go l.writer()
	return l
}

func init() {
	def.Store(New(Config{
		Level:  InfoLevel,
		Output: os.Stdout,
	}))
}

func SetDefault(l *Logger) {
	if l != nil {
		def.Store(l)
	}
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

	for bp := range l.ch {
		l.writeOne(bp)

		// 批量清空积压日志，保持顺序
		for i := 0; i < l.batchSize; i++ {
			select {
			case bp2, ok := <-l.ch:
				if !ok {
					return
				}
				l.writeOne(bp2)
			default:
				goto NEXT
			}
		}
	NEXT:
	}
}

func (l *Logger) writeOne(bp *[]byte) {
	b := *bp

	l.mu.Lock()
	writeFull(l.out, b)
	l.mu.Unlock()

	*bp = b[:0]
	l.pool.Put(bp)
}

func writeFull(w io.Writer, b []byte) {
	for len(b) > 0 {
		n, err := w.Write(b)
		if err != nil {
			return
		}
		b = b[n:]
	}
}

func Debug(msg string, args ...any) { L().log(DebugLevel, msg, args...) }
func Info(msg string, args ...any)  { L().log(InfoLevel, msg, args...) }
func Warn(msg string, args ...any)  { L().log(WarnLevel, msg, args...) }
func Error(msg string, args ...any) { L().log(ErrorLevel, msg, args...) }

func (l *Logger) Debug(msg string, args ...any) { l.log(DebugLevel, msg, args...) }
func (l *Logger) Info(msg string, args ...any)  { l.log(InfoLevel, msg, args...) }
func (l *Logger) Warn(msg string, args ...any)  { l.log(WarnLevel, msg, args...) }
func (l *Logger) Error(msg string, args ...any) { l.log(ErrorLevel, msg, args...) }

func (l *Logger) log(level Level, msg string, args ...any) {
	if l == nil || l.closed.Load() {
		return
	}
	if int32(level) < l.level.Load() {
		return
	}

	bp := l.pool.Get().(*[]byte)
	b := (*bp)[:0]

	b = appendHeader(b, level, l.caller, 3)
	b = append(b, msg...)

	if len(args) > 0 {
		b = append(b, ' ')
		for i, a := range args {
			if i > 0 {
				b = append(b, ' ')
			}
			b = appendValue(b, a)
		}
	}

	b = append(b, '\n')
	*bp = b

	// GUI程序推荐：统一异步队列，顺序稳定
	l.ch <- bp
}

func appendHeader(b []byte, lv Level, caller bool, skip int) []byte {
	b = append(b, '[')
	b = appendLevel(b, lv)
	b = append(b, "] "...)
	b = time.Now().AppendFormat(b, timeFmt)
	b = append(b, ' ')

	if caller {
		_, file, line, ok := runtime.Caller(skip)
		if ok {
			b = appendCaller(b, file, line)
		}
	}
	return b
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

func appendValue(b []byte, v any) []byte {
	switch x := v.(type) {
	case nil:
		return append(b, "null"...)
	case string:
		return append(b, x...)
	case int:
		return strconv.AppendInt(b, int64(x), 10)
	case int8:
		return strconv.AppendInt(b, int64(x), 10)
	case int16:
		return strconv.AppendInt(b, int64(x), 10)
	case int32:
		return strconv.AppendInt(b, int64(x), 10)
	case int64:
		return strconv.AppendInt(b, x, 10)
	case uint:
		return strconv.AppendUint(b, uint64(x), 10)
	case uint8:
		return strconv.AppendUint(b, uint64(x), 10)
	case uint16:
		return strconv.AppendUint(b, uint64(x), 10)
	case uint32:
		return strconv.AppendUint(b, uint64(x), 10)
	case uint64:
		return strconv.AppendUint(b, x, 10)
	case float32:
		return strconv.AppendFloat(b, float64(x), 'f', -1, 32)
	case float64:
		return strconv.AppendFloat(b, x, 'f', -1, 64)
	case bool:
		return strconv.AppendBool(b, x)
	case error:
		return append(b, x.Error()...)
	default:
		return append(b, fmt.Sprintf("%v", x)...)
	}
}
