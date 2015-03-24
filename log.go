package log

import (
	"io"
	L "log"
	"os"
)

type LogLevel byte

type Logger struct {
	log   *L.Logger
	level LogLevel
}

var Out, Err, Default *Logger

const (
	LogQuiet = iota
	LogError
	LogWarn
	LogNormal
	LogInfo
	LogDebug
)

const (
	// Bits or'ed together to control what's printed. There is no control over the
	// order they appear (the order listed here) or the format they present (as
	// described in the comments).  A colon appears after these items:
	//	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
	Ldate         = 1 << iota     // the date: 2009/01/23
	Ltime                         // the time: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func init() {
	Out = New(os.Stdout, "", LstdFlags, LogDebug)
	Err = New(os.Stderr, "", LstdFlags, LogDebug)
	Default = Out
}

func New(out io.Writer, prefix string, flag int, lvl LogLevel) *Logger {
	return &Logger{
		log:   L.New(out, prefix, flag),
		level: lvl,
	}
}

func (l *Logger) lvlLog(lvl LogLevel, fmt string, v []interface{}) {
	if lvl <= l.level {
		l.log.Printf(fmt, v...)
	}
}

func (l *Logger) Error(fmt string, v ...interface{}) {
	l.lvlLog(LogError, fmt, v)
}

func (l *Logger) Debug(fmt string, v ...interface{}) {
	l.lvlLog(LogDebug, fmt, v)
}

func (l *Logger) Warn(fmt string, v ...interface{}) {
	l.lvlLog(LogWarn, fmt, v)
}

func (l *Logger) Info(fmt string, v ...interface{}) {
	l.lvlLog(LogInfo, fmt, v)
}

func (l *Logger) Normal(fmt string, v ...interface{}) {
	l.lvlLog(LogNormal, fmt, v)
}

func (l *Logger) Fatal(fmt string, v ...interface{}) {
	L.Fatalf(fmt, v...)
}

func (l *Logger) Flags() int {
	return l.log.Flags()
}

func (l *Logger) SetFlags(flags int) {
	l.log.SetFlags(flags)
}

func (l *Logger) SetLevel(lvl LogLevel) {
	l.level = lvl
}

func Error(fmt string, v ...interface{}) {
	Default.lvlLog(LogError, fmt, v)
}

func Debug(fmt string, v ...interface{}) {
	Default.lvlLog(LogDebug, fmt, v)
}

func Warn(fmt string, v ...interface{}) {
	Default.lvlLog(LogWarn, fmt, v)
}

func Info(fmt string, v ...interface{}) {
	Default.lvlLog(LogInfo, fmt, v)
}

func Normal(fmt string, v ...interface{}) {
	Default.lvlLog(LogNormal, fmt, v)
}

func Fatal(fmt string, v ...interface{}) {
	L.Fatalf(fmt, v...)
}

func TODO(fmt string, v ...interface{}) {
	Default.lvlLog(LogQuiet, fmt, v)
}

func Flags() int {
	return Default.Flags()
}

func SetFlags(flags int) {
	Default.SetFlags(flags)
}

func SetLevel(lvl LogLevel) {
	Default.level = lvl
}

func SetDefault(l *Logger) {
	Default = l
}
