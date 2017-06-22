package utils

import (
	"flashCoder/app/kernel/log"
	"os"
	"strings"
	"time"
)

type LogLevel uint

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
	Depth       int
}

func (l *LogHandler) GetLogLevel(level string) LogLevel {
	level = strings.ToLower(level)
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
		l.Handler.Println(l.Depth, args...)
	}
}

func (l *LogHandler) Info(args ...interface{}) {
	if l.RecordLevel >= InfoLevel {
		l.Handler.SetPrefix("[Info]")
		l.Handler.Println(l.Depth, args...)
	}
}

func (l *LogHandler) Error(args ...interface{}) {
	if l.RecordLevel >= ErrorLevel {
		l.Handler.SetPrefix("[Error]")
		l.Handler.Println(l.Depth, args...)
	}
}

func (l *LogHandler) Fatal(args ...interface{}) {
	if l.RecordLevel >= FatalLevel {
		l.Handler.SetPrefix("[Fatal]")
		l.Handler.Fatalln(l.Depth, args...)
	}
}

func (l *LogHandler) Panic(args ...interface{}) {
	if l.RecordLevel >= PanicLevel {
		l.Handler.SetPrefix("[Panic]")
		l.Handler.Panicln(l.Depth, args...)
	}
}

func (l *LogHandler) GetLogName() string {
	month := time.Now().Format("200601")
	return month + ".log"
}

func (l *LogHandler) SetLogConfig() {
	config := GetGlobalCfg()
	lc := config.Section("logger")
	l.RecordLevel = l.GetLogLevel(lc.Key("level").String())
	l.FilePath = lc.Key("path").String()
	l.FileName = l.GetLogName()
	l.Depth = 4
}

var Loger *LogHandler

func LogError(level string, err error) {
	if err == nil {
		return
	}

	if Loger == nil {
		Loger = new(LogHandler)
		Loger.SetLogConfig()
		path := Loger.FilePath + "/" + Loger.FileName
		logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		CheckError(err)
		Loger.Handler = log.New(logFile, "[Info]", log.LstdFlags|log.Llongfile)
	}

	loglevel := Loger.GetLogLevel(level)
	switch loglevel {
	case DebugLevel:
		Loger.Debug(err)
	case InfoLevel:
		Loger.Info(err)
	case ErrorLevel:
		Loger.Error(err)
	case FatalLevel:
		Loger.Fatal(err)
	case PanicLevel:
		Loger.Panic(err)
	}

}
