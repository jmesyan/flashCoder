package utils

import (
	"log"
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
}

func (l *LogHandler) GetRecordLevel(level string) LogLevel {
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

func (l *LogHandler) GetLogName() string {
	month := time.Now().Format("200601")
	return month + ".log"
}

func (l *LogHandler) SetLogConfig() {
	config := GetGlobalCfg()
	lc := config.Section("logger")
	l.RecordLevel = l.GetRecordLevel(lc.Key("level").String())
	l.FilePath = lc.Key("path").String()
	l.FileName = l.GetLogName()
}

var Loger *LogHandler

func init() {
	Loger = new(LogHandler)
	Loger.SetLogConfig()
	path := Loger.FilePath + "/" + Loger.FileName
	logFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	CheckError(err)
	Loger.Handler = log.New(logFile, "[Info]", log.LstdFlags|log.Llongfile)

}
