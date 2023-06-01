//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package logger Simple log output
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

const log_file_name = "energy.log"

type CefLogger struct {
	logFile *os.File
	logger  *log.Logger
	enable  bool
	isInit  bool
	level   CefLoggerLevel
}

var logger = &CefLogger{}

func loggerInit() {
	if logger.isInit {
		return
	}
	logger.isInit = true
	logFile, err := os.OpenFile(log_file_name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.enable = false
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
	if enable {
		loggerInit()
	}
}

func Error(v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Error {
		logger.logger.SetPrefix("[ENERGY-Error] ")
		logger.logger.Println(v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Error {
		logger.logger.SetPrefix("[ENERGY-Error] ")
		logger.logger.Printf(format, v...)
	}
}

func Info(v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Info {
		logger.logger.SetPrefix("[ENERGY-Info] ")
		logger.logger.Println(v...)
	}
}

func Infof(format string, v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Info {
		logger.logger.SetPrefix("[ENERGY-Info] ")
		logger.logger.Printf(format, v...)
	}
}

func Debug(v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Debug {
		logger.logger.SetPrefix("[ENERGY-Debug] ")
		logger.logger.Println(v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if logger.enable && logger.level >= CefLog_Debug {
		logger.logger.SetPrefix("[ENERGY-Debug] ")
		logger.logger.Printf(format, v...)
	}
}

func Fatal(v ...interface{}) {
	if logger.enable {
		logger.logger.SetPrefix("[ENERGY-Fatal] ")
		logger.logger.Fatal(v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	if logger.enable {
		logger.logger.SetPrefix("[ENERGY-Fatal] ")
		logger.logger.Fatalf(format, v...)
	}
}
