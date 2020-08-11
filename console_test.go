package vlogger

import "testing"

// Try each log level in decreasing order of priority.
func testConsoleCalls(bl *Logger) {
	bl.Success("Success")
	bl.Failed("Failed")
	bl.Normal("Normal")
	// bl.Panic("Error")
	bl.Error("Error")
	bl.Debug("Debug")
}

func TestConsole(t *testing.T) {
	log := NewLogger().SetLevel(LevelDebug)
	testConsoleCalls(log)
	log.SetLevel(LevelError)
	testConsoleCalls(log)
}
