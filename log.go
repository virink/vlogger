package vlogger

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Level
const (
	LevelSuccess = iota // 成功
	LevelFailed         // 失败
	LevelNormal         // 普通
	LevelError          // 用户错误
	LevelDebug          // 用户级调试
	LevelPanic          // 系统错误
)

// LevelMap 日志等级和描述映射关系
var LevelMap = map[int]string{
	LevelSuccess: "[+]",
	LevelFailed:  "[-]",
	LevelNormal:  "[*]",
	LevelError:   "[!]",
	LevelDebug:   "[#]",
	LevelPanic:   "[~]",
}

const (
	defaultTimeFormat = "15:04:05" // 日志输出默认格式
	fullTimeFormat    = "2006-01-02 15:04:05"
)

// Adapter ID
const (
	AdapterConsole = iota
	AdapterFile
)

type loginfo struct {
	Time    string
	Level   string
	Name    string
	Path    string
	Content string
}

// LoggerAdapter log provider interface
type LoggerAdapter interface {
	Init()
	LogWrite(msg interface{}, level int) error
	Destroy()
	SetLevel(level int)
}

// Logger -
type Logger struct {
	lock       sync.Mutex
	level      int
	timeFormat string
	callDepth  int
	usePath    string
	outputs    map[int]LoggerAdapter
}

// NewLogger Create New Logger
func NewLogger() *Logger {
	l := new(Logger)
	l.outputs = make(map[int]LoggerAdapter)
	l.SetCallDepth(2)
	l.SetLevel(LevelError)
	l.SetTimeFormat(defaultTimeFormat)
	l.SetConsole()
	l.SetTrimPath("src/")
	return l
}

// SetLevel Set Level
func (l *Logger) SetLevel(level int) *Logger {
	l.level = level
	for op := range l.outputs {
		l.outputs[op].SetLevel(level)
	}
	return l
}

// SetTimeFormat Set Time Format
func (l *Logger) SetTimeFormat(timeFormat string) *Logger {
	l.timeFormat = timeFormat
	return l
}

// SetTrimPath Set Trim Path
func (l *Logger) SetTrimPath(trimPath string) *Logger {
	l.usePath = trimPath
	return l
}

// SetCallDepth Set Call Depth
func (l *Logger) SetCallDepth(depth int) *Logger {
	l.callDepth = depth
	return l
}

// Reset Reset
func (l *Logger) Reset() *Logger {
	for _, l := range l.outputs {
		l.Destroy()
	}
	l.outputs = nil
	return l
}

func (l *Logger) writeToLoggers(msg *loginfo, level int) {
	// time level path content
	msgStr := fmt.Sprintf("%s %s%s %s", msg.Time, msg.Level, msg.Path, msg.Content)
	for _, op := range l.outputs {
		op.LogWrite(msgStr, level)
	}
}

func (l *Logger) writeMsg(level int, format interface{}, v ...interface{}) {
	if level > l.level && level < 0 {
		return
	}
	msgSt := new(loginfo)
	msg := formatLog(format, v...)
	src := ""
	if l.level == LevelDebug {
		if _, file, lineno, ok := runtime.Caller(l.callDepth); ok {
			src = strings.Replace(
				fmt.Sprintf(" %s:%d", stringTrim(file, l.usePath), lineno), "%2e", ".", -1)
		}
	}
	msgSt.Time = time.Now().Format(l.timeFormat)
	msgSt.Level = LevelMap[level]
	msgSt.Path = src
	msgSt.Content = msg
	l.writeToLoggers(msgSt, level)
}

// Success Log SUCCESS level message.
func (l *Logger) Success(format interface{}, v ...interface{}) {
	l.writeMsg(LevelSuccess, format, v...)
}

// Failed Log FAILED level message.
func (l *Logger) Failed(format interface{}, v ...interface{}) {
	l.writeMsg(LevelFailed, format, v...)
}

// Normal Log NORMAL level message.
func (l *Logger) Normal(format interface{}, v ...interface{}) {
	l.writeMsg(LevelNormal, format, v...)
}

// Panic -
func (l *Logger) Panic(format interface{}, v ...interface{}) {
	panic(formatLog(format, v))
}

// Error Log ERROR level message.
func (l *Logger) Error(format interface{}, v ...interface{}) {
	l.writeMsg(LevelError, format, v...)
}

// Debug Log DEBUG level message.
func (l *Logger) Debug(format interface{}, v ...interface{}) {
	l.writeMsg(LevelDebug, format, v...)
}
