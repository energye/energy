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

type Level int8

const (
	LError Level = iota
	LInfo
	LDebug
)

var logFileName = "energy.log"

// Logger logger conifg
type Logger struct {
	logFile *os.File
	logger  *log.Logger
	enable  bool
	isInit  bool
	level   Level
}

var logger = &Logger{}

func loggerInit() {
	if logger.isInit {
		return
	}
	logger.isInit = true
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.enable = false
		return
	}
	logger.enable = true
	logger.logFile = logFile
	logger.level = LError
	logger.logger = log.New(io.MultiWriter(os.Stdout, logFile), "", log.Ldate|log.Ltime)
}

// SetLogFile set log file full path
func SetLogFile(filePath string) {
	logFileName = filePath
}

// SetLevel set log level
func SetLevel(l Level) {
	logger.level = l
}

// SetEnable enable log, default true
func SetEnable(enable bool) {
	logger.enable = enable
	if enable {
		loggerInit()
	}
}

// Enable return log Enable
func Enable() bool {
	return logger.enable
}

// Error level
func Error(v ...interface{}) {
	if logger.enable && logger.level >= LError {
		logger.logger.SetPrefix("[ENERGY-Error] ")
		logger.logger.Println(v...)
	}
}

// Errorf level fmt
func Errorf(format string, v ...interface{}) {
	if logger.enable && logger.level >= LError {
		logger.logger.SetPrefix("[ENERGY-Error] ")
		logger.logger.Printf(format, v...)
	}
}

// Info level
func Info(v ...interface{}) {
	if logger.enable && logger.level >= LInfo {
		logger.logger.SetPrefix("[ENERGY-Info] ")
		logger.logger.Println(v...)
	}
}

// Infof level fmt
func Infof(format string, v ...interface{}) {
	if logger.enable && logger.level >= LInfo {
		logger.logger.SetPrefix("[ENERGY-Info] ")
		logger.logger.Printf(format, v...)
	}
}

// Debug level
func Debug(v ...interface{}) {
	if logger.enable && logger.level >= LDebug {
		logger.logger.SetPrefix("[ENERGY-Debug] ")
		logger.logger.Println(v...)
	}
}

// Debugf level fmt
func Debugf(format string, v ...interface{}) {
	if logger.enable && logger.level >= LDebug {
		logger.logger.SetPrefix("[ENERGY-Debug] ")
		logger.logger.Printf(format, v...)
	}
}

// Fatal level
func Fatal(v ...interface{}) {
	if logger.enable {
		logger.logger.SetPrefix("[ENERGY-Fatal] ")
		logger.logger.Fatal(v...)
	}
}

// Fatalf level fmt
func Fatalf(format string, v ...interface{}) {
	if logger.enable {
		logger.logger.SetPrefix("[ENERGY-Fatal] ")
		logger.logger.Fatalf(format, v...)
	}
}
