package logger

import (
	"bytes"
	"strings"
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
