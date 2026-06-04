package notification

import (
	"github.com/energye/energy/v3/application/pack"
	"github.com/energye/energy/v3/logger"
	"github.com/energye/lcl/api"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	myLogger  *logger.Logger
	loggerOne sync.Once
)

func mustInit() {
	loggerOne.Do(func() {
		getLogFilePath := func() string {
			homeDir, _ := os.UserHomeDir()
			bundleId := pack.Info.Id
			if runtime.GOOS == "darwin" && bundleId != "" {
				logDir := filepath.Join(homeDir, "Library", "Logs", bundleId)
				_ = os.MkdirAll(logDir, 0700)
				return filepath.Join(logDir, "notification.log")
			} else {
				return filepath.Join(homeDir, ".energy", "notification.log")
			}
		}
		file, err := os.OpenFile(getLogFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			myLogger = logger.New(logger.Config{Level: logger.DebugLevel, Output: os.Stdout})
		} else {
			multiWriter := io.MultiWriter(os.Stdout, file)
			myLogger = logger.New(logger.Config{
				Level:  logger.DebugLevel,
				Output: multiWriter,
			})
		}
		api.SetOnReleaseCallback(func() {
			if file != nil {
				_ = file.Close()
			}
			myLogger.Close()
		})
	})
}

func Debug(v ...any) {
	mustInit()
	myLogger.Debug("-", v...)
}

func Info(v ...any) {
	mustInit()
	myLogger.Info("-", v...)
}

func Warn(v ...any) {
	mustInit()
	myLogger.Warn("-", v...)
}

func Error(v ...any) {
	mustInit()
	myLogger.Error("-", v...)
}
