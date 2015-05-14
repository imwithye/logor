package main

import "github.com/imwithye/logor"

func main() {
	// Create a new Logor instance
	l := logor.New()
	l.Error("error")
	l.Warn("warn")
	l.Info("info")

	// Logor is set to InfoLevel by default
	l.Debug("debug 1")
	l.Trace("trace 1")

	// Set Logor to a higer level
	l.Level = logor.TraceLevel

	l.Debug("debug 2")
	l.Trace("trace 2")

	// Get a shared instance(singleton) of Logor
	l = logor.GetLogor()
	l.Debug("debug 3")
	l.Level = logor.DebugLevel

	// Level will be set to DebugLevel since it is a shared instance
	l = logor.GetLogor()
	l.Debug("debug 4")
}
