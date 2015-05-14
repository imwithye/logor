package logor

import (
	"fmt"
	"log"
	"os"
)

type Logor struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	debugLogger *log.Logger
	errorLogger *log.Logger
}

func New() *Logor {
	logor := new(Logor)
	flag := log.Lshortfile | log.LstdFlags
	logor.infoLogger = log.New(os.Stdout, "[I] ", flag)
	logor.warnLogger = log.New(os.Stdout, "[W] ", flag)
	logor.debugLogger = log.New(os.Stdout, "[D] ", flag)
	logor.errorLogger = log.New(os.Stderr, "[E] ", flag)
	return logor
}

func (l *Logor) Info(v ...interface{}) {
	l.infoLogger.Output(2, fmt.Sprintln(v...))
}

func (l *Logor) Infof(format string, v ...interface{}) {
	l.infoLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logor) Warn(v ...interface{}) {
	l.warnLogger.Output(2, fmt.Sprintln(v...))
}

func (l *Logor) Warnf(format string, v ...interface{}) {
	l.warnLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logor) Debug(v ...interface{}) {
	l.debugLogger.Output(2, fmt.Sprintln(v...))
}

func (l *Logor) Debugf(format string, v ...interface{}) {
	l.debugLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logor) Error(v ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprintln(v...))
}

func (l *Logor) Errorf(format string, v ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logor) Fatal(v ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logor) Fatalf(format string, v ...interface{}) {
	l.errorLogger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
