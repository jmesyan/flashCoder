package utils

import (
	"fmt"
	"log"
)

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	InfoLevel
	DebugLevel
)

type LogHandler struct {
	Handler     *log.Logger
	FilePath    string
	FileName    string
	RecordLevel LogLevel
}

func (l *LogHandler) GetRecordLevel(level string) LogLevel {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	case "panic":
		return PanicLevel
	default:
		return InfoLevel
	}
}

func (l *LogHandler) Debug(args ...interface{}) {
	if l.RecordLevel >= DebugLevel {
		l.Handler.SetPrefix("[Debug]")
		l.Handler.Println(args...)
	}
}

func (l *LogHandler) Info(args ...interface{}) {
	if l.RecordLevel >= InfoLevel {
		l.Handler.SetPrefix("[Info]")
		l.Handler.Println(args...)
	}
}

func (l *LogHandler) Error(args ...interface{}) {
	if l.RecordLevel >= ErrorLevel {
		l.Handler.SetPrefix("[Error]")
		l.Handler.Println(args...)
	}
}

func (l *LogHandler) Fatal(args ...interface{}) {
	if l.RecordLevel >= FatalLevel {
		l.Handler.SetPrefix("[Fatal]")
		l.Handler.Fatalln(args...)
	}
}

func (l *LogHandler) Panic(args ...interface{}) {
	if l.RecordLevel >= PanicLevel {
		l.Handler.SetPrefix("[Panic]")
		l.Handler.Panicln(args...)
	}
}
