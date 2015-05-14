package main

import "github.com/imwithye/logor"

func main() {
	l := logor.New()
	l.Info("log info")
	l.Warn("log warn")
	l.Debug("log debug 1")
	l.Error("log error")

	l.Level = logor.TraceLevel

	l.Debug("log debug 2")
	l.Trace("log trace")
	l.FatalCode(2, "log fatal")
}
