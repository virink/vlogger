package vlogger

import (
	"fmt"
	"testing"
)

// Try each log level in decreasing order of priority.
func testConsoleCalls(bl *Logger) {
	bl.Success("Success text", "ok", "emmm")
	bl.Failed("Failed text")
	bl.Normal("Normal text")
	bl.Error("Error", fmt.Errorf("testtt error text"))
	bl.Debug("Debug text")
	// bl.Panic("Error Panic")
}

func TestConsole(t *testing.T) {
	log := NewLogger().SetLevel(LevelDebug)
	testConsoleCalls(log)
	log.SetLevel(LevelError)
	testConsoleCalls(log)
}
