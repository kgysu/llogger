package llogger

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

var (
	std = logger(InfoLevel)
)

type LeveledLogger struct {
	l           *log.Logger
	level       int
	logToStdErr bool
}

func NewLogger(loglevel int) *LeveledLogger {
	std = logger(loglevel)
	return std
}

func logger(loglevel int) *LeveledLogger {
	return &LeveledLogger{
		l:           log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		level:       loglevel,
		logToStdErr: false,
	}
}

func (l *LeveledLogger) SetLogToStdErr(shouldLog bool) {
	l.logToStdErr = shouldLog
}

// Struct Methods

func (l *LeveledLogger) Traceln(v ...any) {
	l.Print(TraceLevel, fmt.Sprintln(v...))
}

func (l *LeveledLogger) Debugln(v ...any) {
	l.Print(DebugLevel, fmt.Sprintln(v...))
}

func (l *LeveledLogger) Errorln(v ...any) {
	l.Print(ErrorLevel, fmt.Sprintln(v...))
}

func (l *LeveledLogger) Infoln(v ...any) {
	l.Print(InfoLevel, fmt.Sprintln(v...))
}

func (l *LeveledLogger) Warnln(v ...any) {
	l.Print(WarnLevel, fmt.Sprintln(v...))
}

func (l *LeveledLogger) Tracef(format string, v ...any) {
	l.Print(TraceLevel, fmt.Sprintf(format, v...))
}

func (l *LeveledLogger) Debugf(format string, v ...any) {
	l.Print(DebugLevel, fmt.Sprintf(format, v...))
}

func (l *LeveledLogger) Errorf(format string, v ...any) {
	l.Print(ErrorLevel, fmt.Sprintf(format, v...))
}

func (l *LeveledLogger) Infof(format string, v ...any) {
	l.Print(InfoLevel, fmt.Sprintf(format, v...))
}

func (l *LeveledLogger) Warnf(format string, v ...any) {
	l.Print(WarnLevel, fmt.Sprintf(format, v...))
}

func (l *LeveledLogger) Print(level int, msg string) {
	if level >= l.level {
		color := getColorByLevel(level)
		levelText := getTextByLevel(level)

		if level == ErrorLevel && l.logToStdErr {
			now := time.Now().Format("2006/01/02 15:04:05.999999")
			file, line := getCaller()
			fmt.Fprintf(os.Stderr, "%s %s:%d: %s %s", now, file, line, levelText, msg)
		}

		color.setColor()
		l.l.Output(3, fmt.Sprintf("%s%s", levelText, msg))
		color.resetColor()
	}
}

func getCaller() (string, int) {
	var ok bool
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	return file, line
}

// Public Methods

func Traceln(v ...any) {
	std.Print(TraceLevel, fmt.Sprintln(v...))
}

func Debugln(v ...any) {
	std.Print(DebugLevel, fmt.Sprintln(v...))
}

func Errorln(v ...any) {
	std.Print(ErrorLevel, fmt.Sprintln(v...))
}

func Infoln(v ...any) {
	std.Print(InfoLevel, fmt.Sprintln(v...))
}

func Warnln(v ...any) {
	std.Print(WarnLevel, fmt.Sprintln(v...))
}

func Tracef(format string, v ...any) {
	std.Print(TraceLevel, fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...any) {
	std.Print(DebugLevel, fmt.Sprintf(format, v...))
}

func Errorf(format string, v ...any) {
	std.Print(ErrorLevel, fmt.Sprintf(format, v...))
}

func Infof(format string, v ...any) {
	std.Print(InfoLevel, fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...any) {
	std.Print(WarnLevel, fmt.Sprintf(format, v...))
}
