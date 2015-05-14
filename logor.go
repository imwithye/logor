package logor

import (
	"fmt"
	"log"
	"os"
)

const (
	Off        = 0
	FatalLevel = 100
	ErrorLevel = 200
	WarnLevel  = 300
	InfoLevel  = 400
	DebugLevel = 500
	TraceLevel = 600
)

var logor *Logor = nil

type Logor struct {
	Level       int
	FatalLogger *log.Logger
	ErrorLogger *log.Logger
	WarnLogger  *log.Logger
	InfoLogger  *log.Logger
	DebugLogger *log.Logger
	TraceLogger *log.Logger
}

func New() *Logor {
	l := new(Logor)
	flag := log.Lshortfile | log.LstdFlags

	l.Level = InfoLevel
	l.FatalLogger = log.New(os.Stderr, "[F] ", flag)
	l.ErrorLogger = log.New(os.Stderr, "[E] ", flag)
	l.WarnLogger = log.New(os.Stdout, "[W] ", flag)
	l.InfoLogger = log.New(os.Stdout, "[I] ", flag)
	l.DebugLogger = log.New(os.Stdout, "[D] ", flag)
	l.TraceLogger = log.New(os.Stdout, "[T] ", flag)

	return l
}

func GetLogor() *Logor {
	if logor == nil {
		logor = New()
	}
	return logor
}

func (l *Logor) Fatal(v ...interface{}) {
	if l.Level >= FatalLevel {
		l.FatalLogger.Output(2, fmt.Sprintln(v...))
		os.Exit(1)
	}
}

func (l *Logor) Fatalf(format string, v ...interface{}) {
	if l.Level >= FatalLevel {
		l.FatalLogger.Output(2, fmt.Sprintf(format, v...))
		os.Exit(1)
	}
}

func (l *Logor) FatalCode(code int, v ...interface{}) {
	if l.Level >= FatalLevel {
		l.FatalLogger.Output(2, fmt.Sprintln(v...))
		os.Exit(code)
	}
}

func (l *Logor) FatalfCode(code int, format string, v ...interface{}) {
	if l.Level >= FatalLevel {
		l.FatalLogger.Output(2, fmt.Sprintf(format, v...))
		os.Exit(code)
	}
}

func (l *Logor) Error(v ...interface{}) {
	if l.Level >= ErrorLevel {
		l.ErrorLogger.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logor) Errorf(format string, v ...interface{}) {
	if l.Level >= ErrorLevel {
		l.ErrorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logor) Warn(v ...interface{}) {
	if l.Level >= WarnLevel {
		l.WarnLogger.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logor) Warnf(format string, v ...interface{}) {
	if l.Level >= WarnLevel {
		l.WarnLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logor) Info(v ...interface{}) {
	if l.Level >= InfoLevel {
		l.InfoLogger.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logor) Infof(format string, v ...interface{}) {
	if l.Level >= InfoLevel {
		l.InfoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logor) Debug(v ...interface{}) {
	if l.Level >= DebugLevel {
		l.DebugLogger.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logor) Debugf(format string, v ...interface{}) {
	if l.Level >= DebugLevel {
		l.DebugLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func (l *Logor) Trace(v ...interface{}) {
	if l.Level >= TraceLevel {
		l.TraceLogger.Output(2, fmt.Sprintln(v...))
	}
}

func (l *Logor) Tracef(format string, v ...interface{}) {
	if l.Level >= TraceLevel {
		l.TraceLogger.Output(2, fmt.Sprintf(format, v...))
	}
}
