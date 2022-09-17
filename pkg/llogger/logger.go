package llogger

import (
	"fmt"
	"log"
	"os"
)

type LeveledLogger struct {
	l     *log.Logger
	level int
}

func NewLogger(loglevel int) *LeveledLogger {
	return &LeveledLogger{
		l:     log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile),
		level: loglevel,
	}
}

func (l *LeveledLogger) Traceln(msg string) {
	l.Tracef("%s\n", msg)
}

func (l *LeveledLogger) Debugln(msg string) {
	l.Debugf("%s\n", msg)
}

func (l *LeveledLogger) Errorln(msg string) {
	l.Errorf("%s\n", msg)
}

func (l *LeveledLogger) Infoln(msg string) {
	l.Infof("%s\n", msg)
}

func (l *LeveledLogger) Warnln(msg string) {
	l.Warnf("%s\n", msg)
}

func (l *LeveledLogger) Tracef(format string, msg string) {
	l.Printf(TraceLevel, format, msg)
}

func (l *LeveledLogger) Debugf(format, msg string) {
	l.Printf(DebugLevel, format, msg)
}

func (l *LeveledLogger) Errorf(format string, msg string) {
	l.Printf(ErrorLevel, format, msg)
}

func (l *LeveledLogger) Infof(format string, msg string) {
	l.Printf(InfoLevel, format, msg)
}

func (l *LeveledLogger) Warnf(format string, msg string) {
	l.Printf(WarnLevel, format, msg)
}

func (l *LeveledLogger) Printf(level int, format, msg string) {
	if level >= l.level {
		color := getColorByLevel(level)
		levelText := getTextByLevel(level)
		color.setColor()
		l.l.Printf("%s%s", levelText, fmt.Sprintf(format, msg))
		color.resetColor()
	}
}
