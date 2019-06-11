package logtool

import (
	"fmt"
	"log"
	"os"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const (
	// LevelNone 0     不写文件日志, 只输出std
	LevelNone = iota
	// LevelError 1		只写Error级别
	LevelError
	// LevelWarning 2
	LevelWarning
	// LevelInfo 3
	LevelInfo
	// LevelDebug 4		写所有日志
	LevelDebug
)

// Logger ...
type Logger struct {
	level int
	depth int
	std   *log.Logger
	err   *log.Logger
	warn  *log.Logger
	info  *log.Logger
	debug *log.Logger
}

// SetLevel ...
func (l *Logger) SetLevel(level int) {
	l.level = level
}

// NewLogger ...
func NewLogger(logPath string, prefix string, level int, alias string, flag ...int) *Logger {
	logFlag := log.Ldate | log.Lmicroseconds | log.Lshortfile
	if len(flag) > 0 {
		logFlag = flag[0]
	}

	logger := new(Logger)
	logger.depth = 3
	// 默认输出到标准std
	logger.std = log.New(os.Stdout, "", logFlag)
	// 日志级别设置
	logger.SetLevel(level)
	if _, err := os.Stat(logPath); err != nil {
		logger.std.Output(logger.depth, "log_path is not exists!")
		logger.SetLevel(LevelNone)
	}

	// Error 级日志
	logger.err = log.New(&lumberjack.Logger{
		Filename:  fmt.Sprintf("%s/%s.error.log", logPath, alias),
		MaxSize:   50,
		MaxAge:    3,
		LocalTime: true,
		Compress:  false,
	}, prefix+" [E] ", logFlag)

	// Warn 级日志
	logger.warn = log.New(&lumberjack.Logger{
		Filename:  fmt.Sprintf("%s/%s.warn.log", logPath, alias),
		MaxSize:   50,
		MaxAge:    3,
		LocalTime: true,
		Compress:  false,
	}, prefix+" [W] ", logFlag)

	// Info 级日志
	logger.info = log.New(&lumberjack.Logger{
		Filename:  fmt.Sprintf("%s/%s.info.log", logPath, alias),
		MaxSize:   50,
		MaxAge:    3,
		LocalTime: true,
		Compress:  false,
	}, prefix+" [I] ", logFlag)

	// Debug 级日志
	logger.debug = log.New(&lumberjack.Logger{
		Filename:  fmt.Sprintf("%s/%s.debug.log", logPath, alias),
		MaxSize:   50,
		MaxAge:    3,
		LocalTime: true,
		Compress:  false,
	}, prefix+" [D] ", logFlag)

	return logger
}

// Error ...
func (l *Logger) Error(format string, v ...interface{}) {
	l.std.Output(l.depth, fmt.Sprintf("[E] "+format, v...))
	if LevelError > l.level {
		return
	}
	l.err.Output(l.depth, fmt.Sprintf(format, v...))
}

// Warn ...
func (l *Logger) Warn(format string, v ...interface{}) {
	l.std.Output(l.depth, fmt.Sprintf("[W] "+format, v...))
	if LevelWarning > l.level {
		return
	}
	l.warn.Output(l.depth, fmt.Sprintf(format, v...))
}

// Info ...
func (l *Logger) Info(format string, v ...interface{}) {
	l.std.Output(l.depth, fmt.Sprintf("[I] "+format, v...))
	if LevelInfo > l.level {
		return
	}
	l.info.Output(l.depth, fmt.Sprintf(format, v...))
}

// Debug ...
func (l *Logger) Debug(format string, v ...interface{}) {
	l.std.Output(l.depth, fmt.Sprintf("[D] "+format, v...))
	if LevelDebug > l.level {
		return
	}
	l.debug.Output(l.depth, fmt.Sprintf(format, v...))
}
