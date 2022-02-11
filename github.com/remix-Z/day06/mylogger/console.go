package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

//设置日志级别
type LogLever uint16

const (
	UNKNOWN LogLever = iota
	DEBUG
	WARNING
	INFO
	ERROR
	FATAL
)

//在终端上输出日志
type Logger struct {
	Lever LogLever //日志级别开关
}

func parseLevel(s string) (LogLever, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "warning":
		return WARNING, nil
	case "info":
		return INFO, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效日志级别")
		return UNKNOWN, err
	}
}

//NewLog 构造函数
func NewLog(Lever string) Logger {
	//创建一个日志级别
	level, err := parseLevel(Lever)
	if err != nil {
		panic(err)
	}
	return Logger{
		Lever: level,
	}
}

//日志输出
func (l Logger) logprint(lever LogLever, flag string, format string, a ...interface{}) {

	str := fmt.Sprintf(format, a...)
	funcName, fileName, line := getInfo(3)
	//fmt.Fprintln(os.Stdout, str)
	if l.Lever <= lever {
		now := time.Now()
		//行号 时间 文件名 函数名 日志级别 日志内容
		fmt.Printf("[%d] [%s] [%s] [%s] [%-7s] %s\n", line, now.Format("2006-01-02 15:04:05"), fileName, funcName, flag, str)
	}
}

func (l Logger) MyDebug(format string, a ...interface{}) {
	l.logprint(DEBUG, "DEBUG", format, a...)
}

func (l Logger) MyInfo(format string, a ...interface{}) {
	l.logprint(INFO, "INFO", format, a...)
}

func (l Logger) MyError(format string, a ...interface{}) {
	l.logprint(ERROR, "ERROR", format, a...)
}

func (l Logger) MyWarning(format string, a ...interface{}) {
	l.logprint(WARNING, "WARNING", format, a...)
}

func (l Logger) MyFatal(format string, a ...interface{}) {
	l.logprint(FATAL, "FATAL", format, a...)
}

func getInfo(n int) (funcName string, fileName string, line int) {
	//runtime file全路径
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller failed")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	return strings.Split(funcName, ".")[1], path.Base(file), line
}
