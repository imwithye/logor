package logor

import (
	"testing"
)

func TestLogorLevelSetting(t *testing.T) {
	l := New()
	l.Error("should be displayed")
	l.Debug("should not be displayed")
	l.Trace("should not be displayed")

	l.Level = TraceLevel
	l.Debug("should be displayed")
	l.Trace("should be displayed")

	l.Level = FatalLevel
	l.Error("should not be displayed")
	l.Trace("should not be displayed")
}

func TestSingtelton(t *testing.T) {
	GetLogor().Level = TraceLevel

	l := GetLogor()
	l.Debug("should be displayed")
	l.Trace("should be displayed")
	l.Level = FatalLevel

	l = GetLogor()
	l.Error("should not be displayed")
}
