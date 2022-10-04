package cef

import (
	"io"
	"log"
	"os"
)

type CefLoggerLevel int8

const (
	CefLog_Error CefLoggerLevel = iota
	CefLog_Info
	CefLog_Debug
)
const log_file_name = "cef-lcl.log"

type CefLogger struct {
	logFile *os.File
	logger  *log.Logger
	enable  bool
	level   CefLoggerLevel
}

var Logger = &CefLogger{}

func init() {
	logFile, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	Logger.enable = true
	Logger.logFile = logFile
	Logger.level = CefLog_Error
	Logger.logger = log.New(io.MultiWriter(os.Stdout, logFile), "", log.Ldate|log.Ltime)
}

func SetLogger(logger *CefLogger) {
	if logger != nil {
		Logger = logger
	}
}

func (m *CefLogger) SetLevel(l CefLoggerLevel) {
	m.level = l
}

func (m *CefLogger) SetEnable(enable bool) {
	m.enable = enable
	if !m.enable {
		m.logFile.Close()
		os.Remove(log_file_name)
	}
}

func (m *CefLogger) Error(v ...interface{}) {
	if m.enable && m.level >= CefLog_Error {
		m.logger.SetPrefix("[CEF-LCL-Error] ")
		m.logger.Println(v...)
	}
}

func (m *CefLogger) Errorf(format string, v ...interface{}) {
	if m.enable && m.level >= CefLog_Error {
		m.logger.SetPrefix("[CEF-LCL-Error] ")
		m.logger.Printf(format, v...)
	}
}

func (m *CefLogger) Info(v ...interface{}) {
	if m.enable && m.level >= CefLog_Info {
		m.logger.SetPrefix("[CEF-LCL-Info] ")
		m.logger.Println(v...)
	}
}

func (m *CefLogger) Infof(format string, v ...interface{}) {
	if m.enable && m.level >= CefLog_Info {
		m.logger.SetPrefix("[CEF-LCL-Info] ")
		m.logger.Printf(format, v...)
	}
}

func (m *CefLogger) Debug(v ...interface{}) {
	if m.enable && m.level >= CefLog_Debug {
		m.logger.SetPrefix("[CEF-LCL-Debug] ")
		m.logger.Println(v...)
	}
}

func (m *CefLogger) Debugf(format string, v ...interface{}) {
	if m.enable && m.level >= CefLog_Debug {
		m.logger.SetPrefix("[CEF-LCL-Debug] ")
		m.logger.Printf(format, v...)
	}
}
