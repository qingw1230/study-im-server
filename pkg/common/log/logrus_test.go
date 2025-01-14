package log

import "testing"

func TestLog(t *testing.T) {
	globalLogger.Debug("test debug: ", "DEBUG")
	globalLogger.Info("test info: ", "INFO")
	globalLogger.Warn("test warn: ", "WARN")
	globalLogger.Error("test error: ", "ERROR")
}
