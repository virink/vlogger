package vlogger

import (
	"fmt"
	"os"
	"runtime"
	"sync"
)

type brush func(string) string

func newBrush(color string) brush {
	return func(text string) string {
		return fmt.Sprintf("\033[%sm%s\033[0m", color, text)
	}
}

var colors = map[int]brush{
	LevelSuccess: newBrush("1;32"), // 绿色
	LevelFailed:  newBrush("1;35"), // 紫色
	LevelNormal:  newBrush("0;37"), // 白色
	LevelPanic:   newBrush("1;41"), // 红色底
	LevelError:   newBrush("1;31"), // 红色
	LevelDebug:   newBrush("1;33"), // 黄色
}

type consoleLogger struct {
	sync.Mutex
	Level    int  `json:"level"`
	Colorful bool `json:"color"`
}

func (c *consoleLogger) Init() {
	c.Level = LevelError
}

func (c *consoleLogger) LogWrite(msgText interface{}, level int) error {
	if level > c.Level {
		return nil
	}
	msg, ok := msgText.(string)
	if !ok {
		return nil
	}
	if runtime.GOOS != "windows" {
		msg = colors[level](msg)
	}
	c.printlnConsole(msg)
	return nil
}

func (c *consoleLogger) Destroy() {}

func (c *consoleLogger) SetLevel(level int) {
	c.Level = level
}

func (c *consoleLogger) printlnConsole(msg string) {
	c.Lock()
	defer c.Unlock()
	// os.Stdout.Write(append([]byte(msg), '\n'))
	os.Stdout.WriteString(msg + "\n")
}

// SetConsole -
func (l *Logger) SetConsole() *Logger {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.outputs[AdapterConsole] = &consoleLogger{
		Level: l.level,
	}
	return l
}
