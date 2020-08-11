package vlogger

import (
	"os"
	"runtime"
	"sync"
)

type brush func(string) string

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
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
	LogLevel string
}

func (c *consoleLogger) Init() {
	c.Level = LevelError
	if runtime.GOOS == "windows" {
		c.Colorful = false
	}
	c.LogLevel = LevelMap[c.Level]
}

func (c *consoleLogger) Debug() {
	c.Level = LevelDebug
	c.LogLevel = LevelMap[c.Level]
	return
}

func (c *consoleLogger) LogWrite(msgText interface{}, level int) error {
	if level > c.Level {
		return nil
	}
	msg, ok := msgText.(string)
	if !ok {
		return nil
	}
	if c.Colorful {
		msg = colors[level](msg)
	}
	c.printlnConsole(msg)
	return nil
}

func (c *consoleLogger) Destroy() {}

func (c *consoleLogger) ID() int {
	return AdapterConsole
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
		Level:    LevelError,
		Colorful: runtime.GOOS != "windows",
	}
	return l
}
