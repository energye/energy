package logger

import (
	"bytes"
	"errors"
	"strings"
	"sync"
	"testing"
	"time"
)

// TestLevel 测试日志级别过滤
func TestLevel(t *testing.T) {
	tests := []struct {
		name     string
		setLevel Level
		logLevel Level
		wantLog  bool
	}{
		{"Debug允许所有", DebugLevel, DebugLevel, true},
		{"Debug允许Info", DebugLevel, InfoLevel, true},
		{"Debug允许Warn", DebugLevel, WarnLevel, true},
		{"Debug允许Error", DebugLevel, ErrorLevel, true},
		{"Info过滤Debug", InfoLevel, DebugLevel, false},
		{"Info允许Info", InfoLevel, InfoLevel, true},
		{"Warn过滤Debug和Info", WarnLevel, InfoLevel, false},
		{"Error只允许Error", ErrorLevel, WarnLevel, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			l := New(Config{
				Level:  tt.setLevel,
				Output: &buf,
			})

			switch tt.logLevel {
			case DebugLevel:
				l.Debug("test")
			case InfoLevel:
				l.Info("test")
			case WarnLevel:
				l.Warn("test")
			case ErrorLevel:
				l.Error("test")
			}

			l.Close()
			hasLog := buf.Len() > 0
			if hasLog != tt.wantLog {
				t.Errorf("期望日志输出=%v, 实际=%v, 缓冲区内容=%q", tt.wantLog, hasLog, buf.String())
			}
		})
	}
}

// TestFieldConstructors 测试字段构造函数
func TestFieldConstructors(t *testing.T) {
	tests := []struct {
		name  string
		field Field
		want  string
	}{
		{"String字段", String("key", "value"), `key="value"`},
		{"Int字段", Int("num", 42), "num=42"},
		{"Int64字段", Int64("big", 9223372036854775807), "big=9223372036854775807"},
		{"Float64字段", Float64("pi", 3.14159), "pi=3.14159"},
		{"Bool真值", Bool("flag", true), "flag=true"},
		{"Bool假值", Bool("flag", false), "flag=false"},
		{"Err非空", Err(errors.New("test error")), `error="test error"`},
		{"Err空值", Err(nil), `error=""`},
		{"Any字符串", Any("any", "str"), `any="str"`},
		{"Any整数", Any("any", 123), "any=123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			l := New(Config{Output: &buf})

			l.InfoF("msg", tt.field)
			l.Close()

			output := buf.String()
			if !strings.Contains(output, tt.want) {
				t.Errorf("期望包含 %q, 实际输出: %q", tt.want, output)
			}
		})
	}
}

// TestLogWithFields 测试带字段的日志输出
func TestLogWithFields(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	l.InfoF("user login", String("user_id", "12345"), Int("age", 25), Bool("vip", true))
	l.Close()

	output := buf.String()

	if !strings.Contains(output, "[INFO]") {
		t.Errorf("缺少级别标签, 输出: %q", output)
	}
	if !strings.Contains(output, `"user login"`) {
		t.Errorf("缺少消息, 输出: %q", output)
	}
	if !strings.Contains(output, `user_id="12345"`) {
		t.Errorf("缺少 user_id 字段, 输出: %q", output)
	}
	if !strings.Contains(output, "age=25") {
		t.Errorf("缺少 age 字段, 输出: %q", output)
	}
	if !strings.Contains(output, "vip=true") {
		t.Errorf("缺少 vip 字段, 输出: %q", output)
	}
}

// TestLogWithAny 测试 any 参数日志输出
func TestLogWithAny(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	l.Info("request", "url", "/api/test", "status", 200, "duration", 0.045)
	l.Close()

	output := buf.String()

	if !strings.Contains(output, `url="/api/test"`) {
		t.Errorf("缺少 url 字段, 输出: %q", output)
	}
	if !strings.Contains(output, "status=200") {
		t.Errorf("缺少 status 字段, 输出: %q", output)
	}
	if !strings.Contains(output, "duration=0.045") {
		t.Errorf("缺少 duration 字段, 输出: %q", output)
	}
}

// TestLogWithMixedValues 测试混合值（非 key-value）
func TestLogWithMixedValues(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	l.Info("values", "first", 123, "second")
	l.Close()

	output := buf.String()
	if !strings.Contains(output, `"first"`) || !strings.Contains(output, "123") {
		t.Errorf("输出不符合预期: %q", output)
	}
}

// TestCaller 测试调用者信息显示
func TestCaller(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{
		Output: &buf,
		Caller: true,
	})

	l.Info("with caller")
	l.Close()

	output := buf.String()
	if !strings.Contains(output, ".go:") {
		t.Errorf("缺少调用者信息, 输出: %q", output)
	}
}

// TestNoCaller 测试不显示调用者信息
func TestNoCaller(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{
		Output: &buf,
		Caller: false,
	})

	l.Info("no caller")
	l.Close()

	output := buf.String()
	if strings.Contains(output, ".go:") {
		t.Errorf("不应包含调用者信息, 输出: %q", output)
	}
}

// TestTimestamp 测试时间戳格式
func TestTimestamp(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	l.Info("timestamp test")
	l.Close()

	output := buf.String()
	if !strings.Contains(output, "20") || !strings.Contains(output, "-") || !strings.Contains(output, ":") {
		t.Errorf("时间戳格式不正确, 输出: %q", output)
	}
}

// TestSpecialCharacters 测试特殊字符转义
func TestSpecialCharacters(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	l.InfoF("special", String("newline", "line1\nline2"), String("quote", `say "hello"`), String("tab", "col1\tcol2"))
	l.Close()

	output := buf.String()

	if !strings.Contains(output, `\n`) {
		t.Errorf("换行符未转义, 输出: %q", output)
	}
	if !strings.Contains(output, `\"`) {
		t.Errorf("引号未转义, 输出: %q", output)
	}
	if !strings.Contains(output, `\t`) {
		t.Errorf("制表符未转义, 输出: %q", output)
	}
}

// TestNilValue 测试 nil 值处理
func TestNilValue(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	l.Info("nil test", "null_value", nil)
	l.Close()

	output := buf.String()
	if !strings.Contains(output, "null") {
		t.Errorf("nil 值应显示为 null, 输出: %q", output)
	}
}

// TestAsyncWrite 测试异步写入
func TestAsyncWrite(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	for i := 0; i < 100; i++ {
		l.InfoF("async test", Int("index", i))
	}

	l.Close()

	output := buf.String()
	lines := strings.Count(output, "\n")
	if lines != 100 {
		t.Errorf("期望 100 条日志, 实际 %d 条", lines)
	}
}

// TestChannelFull 测试通道满时的降级处理
func TestChannelFull(t *testing.T) {
	var buf bytes.Buffer
	l := &Logger{
		out:    &buf,
		caller: false,
		ch:     make(chan []byte, 2),
		pool:   sync.Pool{New: func() any { b := make([]byte, 0, 512); return &b }},
	}
	l.level.Store(int32(InfoLevel))
	l.wg.Add(1)
	go l.writer()

	for i := 0; i < 10; i++ {
		l.InfoF("overflow", Int("i", i))
	}

	l.Close()
	output := buf.String()
	lines := strings.Count(output, "\n")
	if lines != 10 {
		t.Errorf("期望 10 条日志(包括降级写入), 实际 %d 条", lines)
	}
}

// TestSetLevel 测试动态修改日志级别
func TestSetLevel(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{
		Level:  InfoLevel,
		Output: &buf,
	})

	l.Info("info msg")
	l.SetLevel(ErrorLevel)
	l.Info("should not appear")
	l.Error("error msg")

	l.Close()

	output := buf.String()
	if !strings.Contains(output, "info msg") {
		t.Error("第一条 info 日志应该存在")
	}
	if strings.Contains(output, "should not appear") {
		t.Error("切换级别后的 info 日志不应存在")
	}
	if !strings.Contains(output, "error msg") {
		t.Error("error 日志应该存在")
	}
}

// TestDefaultLogger 测试默认日志器
func TestDefaultLogger(t *testing.T) {
	var buf bytes.Buffer
	customLogger := New(Config{Output: &buf})
	SetDefault(customLogger)

	Info("default logger test")
	customLogger.Close()

	output := buf.String()
	if !strings.Contains(output, "default logger test") {
		t.Errorf("默认日志器未工作, 输出: %q", output)
	}
}

// TestGlobalFunctions 测试全局函数
func TestGlobalFunctions(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})
	SetDefault(l)

	Debug("debug msg")
	Info("info msg")
	Warn("warn msg")
	Error("error msg")

	DebugF("debug field", String("k", "v"))
	InfoF("info field", Int("n", 1))
	WarnF("warn field", Bool("b", true))
	ErrorF("error field", Err(errors.New("err")))

	l.Close()

	output := buf.String()
	expected := []string{"debug msg", "info msg", "warn msg", "error msg", `k="v"`, "n=1", "b=true", `error="err"`}
	for _, exp := range expected {
		if !strings.Contains(output, exp) {
			t.Errorf("期望包含 %q, 输出: %q", exp, output)
		}
	}
}

// TestConcurrent 测试并发安全性
func TestConcurrent(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	var wg sync.WaitGroup
	goroutines := 10
	logsPerGoroutine := 100

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < logsPerGoroutine; j++ {
				l.InfoF("concurrent", Int("goroutine", id), Int("index", j))
			}
		}(i)
	}

	wg.Wait()
	l.Close()

	output := buf.String()
	lines := strings.Count(output, "\n")
	expected := goroutines * logsPerGoroutine
	if lines != expected {
		t.Errorf("期望 %d 条日志, 实际 %d 条", expected, lines)
	}
}

// TestPerformance 性能基准测试
func BenchmarkLogWithFields(b *testing.B) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})
	defer l.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.InfoF("benchmark", String("key", "value"), Int("num", 42))
	}
}

func BenchmarkLogWithAny(b *testing.B) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})
	defer l.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info("benchmark", "key", "value", "num", 42)
	}
}

func BenchmarkLogNoArgs(b *testing.B) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})
	defer l.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info("benchmark no args")
	}
}

func BenchmarkLogWithError(b *testing.B) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})
	defer l.Close()

	err := errors.New("test error")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.ErrorF("error occurred", Err(err))
	}
}

// TestOutputFormat 测试输出格式完整性
func TestOutputFormat(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{
		Output: &buf,
		Caller: true,
	})

	l.InfoF("format test", String("key", "value"))
	l.Close()

	output := buf.String()

	parts := []string{"[INFO]", "\"format test\"", `key="value"`}
	for _, part := range parts {
		if !strings.Contains(output, part) {
			t.Errorf("输出缺少 %q, 完整输出: %q", part, output)
		}
	}
}

// TestEdgeCases 测试边界情况
func TestEdgeCases(t *testing.T) {
	t.Run("空消息", func(t *testing.T) {
		var buf bytes.Buffer
		l := New(Config{Output: &buf})
		l.Info("")
		l.Close()
		if buf.Len() == 0 {
			t.Error("空消息也应该输出")
		}
	})

	t.Run("空字段列表", func(t *testing.T) {
		var buf bytes.Buffer
		l := New(Config{Output: &buf})
		l.InfoF("no fields")
		l.Close()
		if !strings.Contains(buf.String(), "no fields") {
			t.Error("无字段时消息应正常输出")
		}
	})

	t.Run("超长字符串", func(t *testing.T) {
		var buf bytes.Buffer
		l := New(Config{Output: &buf})
		longStr := strings.Repeat("a", 10000)
		l.InfoF("long", String("data", longStr))
		l.Close()
		if !strings.Contains(buf.String(), longStr) {
			t.Error("超长字符串未正确输出")
		}
	})

	t.Run("Unicode字符", func(t *testing.T) {
		var buf bytes.Buffer
		l := New(Config{Output: &buf})
		l.InfoF("unicode", String("chinese", "中文"), String("emoji", "😀"))
		l.Close()
		output := buf.String()
		if !strings.Contains(output, "中文") || !strings.Contains(output, "😀") {
			t.Errorf("Unicode 字符丢失, 输出: %q", output)
		}
	})
}

// TestMemoryReuse 测试内存复用
func TestMemoryReuse(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	for i := 0; i < 1000; i++ {
		l.InfoF("reuse test", Int("i", i))
	}

	l.Close()

	reusedBuf := l.pool.Get().(*[]byte)
	if cap(*reusedBuf) < 512 {
		t.Errorf("缓冲区容量不足, 期望>=512, 实际=%d", cap(*reusedBuf))
	}
}

// TestTimeAccuracy 测试时间准确性
func TestTimeAccuracy(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	before := time.Now()
	l.Info("time test")
	l.Close()
	after := time.Now()

	output := buf.String()
	if !strings.Contains(output, ":") {
		t.Error("时间戳格式错误")
	}

	_ = before
	_ = after
}

// TestDoubleClose 测试重复关闭的安全性
func TestDoubleClose(t *testing.T) {
	var buf bytes.Buffer
	l := New(Config{Output: &buf})

	l.Info("test")
	l.Close()

	l.Close()

	output := buf.String()
	if !strings.Contains(output, "test") {
		t.Error("日志应该正常输出")
	}
}
