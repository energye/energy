//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under GNU General Public License v3.0
//
//----------------------------------------

package logger

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

var logger = &CefLogger{}

func init() {
	logFile, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	logger.enable = true
	logger.logFile = logFile
	logger.level = CefLog_Error
	logger.logger = log.New(io.MultiWriter(os.Stdout, logFile), "", log.Ldate|log.Ltime)
}

func SetLevel(l CefLoggerLevel) {
	logger.level = l
}

func SetEnable(enable bool) {
	logger.enable = enable
	if !logger.enable {
		logger.logFile.Close()
		os.Remove(log_file_name)
	}
}

func Error(v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Error {
		logger.logger.SetPrefix("[CEF-LCL-Error] ")
		logger.logger.Println(v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Error {
		logger.logger.SetPrefix("[CEF-LCL-Error] ")
		logger.logger.Printf(format, v...)
	}
}

func Info(v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Info {
		logger.logger.SetPrefix("[CEF-LCL-Info] ")
		logger.logger.Println(v...)
	}
}

func Infof(format string, v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Info {
		logger.logger.SetPrefix("[CEF-LCL-Info] ")
		logger.logger.Printf(format, v...)
	}
}

func Debug(v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Debug {
		logger.logger.SetPrefix("[CEF-LCL-Debug] ")
		logger.logger.Println(v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Debug {
		logger.logger.SetPrefix("[CEF-LCL-Debug] ")
		logger.logger.Printf(format, v...)
	}
}
