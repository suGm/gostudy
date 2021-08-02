package myLogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// Logger对象
type Logger struct {
	Level LogLevel
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOW, err
	}
}

// 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return Logger{
		Level: level,
	}
}

func (l Logger) enable(logLevel LogLevel) bool {
	return l.Level <= logLevel
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime caller() failed")
		return
	}

	funcName = runtime.FuncForPC(pc).Name()

	return funcName, path.Base(file), line
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		now := time.Now()
		funcName, fileName, line := getInfo(2)
		fmt.Printf("[%s] [%s] [funcName:%s] [fileName:%s] [line:%d] %s\n", now.Format("2006-01-02 15:04:05"), "Debug", funcName, fileName, line, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		now := time.Now()
		funcName, fileName, line := getInfo(2)
		fmt.Printf("[%s] [%s] [funcName:%s] [fileName:%s] [line:%d] %s\n", now.Format("2006-01-02 15:04:05"), "Info", funcName, fileName, line, msg)

	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		now := time.Now()
		funcName, fileName, line := getInfo(2)
		fmt.Printf("[%s] [%s] [funcName:%s] [fileName:%s] [line:%d] %s\n", now.Format("2006-01-02 15:04:05"), "Warning", funcName, fileName, line, msg)

	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		now := time.Now()
		funcName, fileName, line := getInfo(2)
		fmt.Printf("[%s] [%s] [funcName:%s] [fileName:%s] [line:%d] %s\n", now.Format("2006-01-02 15:04:05"), "Error", funcName, fileName, line, msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		now := time.Now()
		funcName, fileName, line := getInfo(2)
		fmt.Printf("[%s] [%s] [funcName:%s] [fileName:%s] [line:%d] %s\n", now.Format("2006-01-02 15:04:05"), "Fatal", funcName, fileName, line, msg)
	}
}
