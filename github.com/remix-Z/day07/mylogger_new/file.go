package mylogger_new

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type FileLogger struct {
	Level       LogLever
	filePath    string //日志文件路径
	fileName    string //日志文件保存文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
	//logChan     chan string 不放字符串类型是因为占用的内存空间很大
	logChan chan *logMsg
}

//定义一个日志的结构体
type logMsg struct {
	line      int      //行号
	timeStamp string   //时间戳
	fileName  string   //文件名
	funcName  string   //函数名
	flag      LogLever //日志级别
	msg       string   //日志内容
}

//构造函数
func NewFileLogger(leverStr, fp, fn string, maxSize int64) *FileLogger {
	LogLever, err := parseLevel(leverStr)
	if err != nil {
		panic(err)
	}

	fl := &FileLogger{
		Level:       LogLever,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *logMsg, 50000),
	}

	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}

	return fl
}

func (f *FileLogger) initFile() error {
	//直接调用path包
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	f.fileObj = fileObj
	f.errFileObj = errFileObj
	//开启多个个后台的goroutine去往文件中
	for i := 0; i < 3; i++ {
		go f.writeLogBackGround()
	}

	return nil
}

//关闭文件
// func (f *FileLogger) fileClose() {
// 	f.fileObj.Close()
// 	f.errFileObj.Close()
// }

//往后台日志中写文件
func (f *FileLogger) writeLogBackGround() {
	for {
		//判断文件大小
		if f.checkFileSize(f.fileObj) {
			newObj := f.newFileSplit(f.fileObj)
			f.fileObj = newObj
		}

		select {
		case logTmp := <-f.logChan:
			//转换
			flag, err := parseLevel02(logTmp.flag)
			if err != nil {
				panic(err)
			}

			if f.Level <= logTmp.flag {
				//行号 时间 文件名 函数名 日志级别 日志内容
				fmt.Fprintf(f.fileObj, "[%d] [%s] [%s] [%s] [%s] %s\n", logTmp.line, logTmp.timeStamp, logTmp.funcName, logTmp.fileName, flag, logTmp.msg)
			}

			if f.checkFileSize(f.errFileObj) {
				newErrObj := f.newFileSplit(f.errFileObj)
				f.errFileObj = newErrObj
			}

			//如果日志级别大于等于error 需要在.err文件中再次记录一遍
			if logTmp.flag >= ERROR {
				//行号 时间 文件名 函数名 日志级别 日志内容
				fmt.Fprintf(f.fileObj, "[%d] [%s] [%s] [%s] [%s] %s\n", logTmp.line, logTmp.timeStamp, logTmp.funcName, logTmp.fileName, flag, logTmp.msg)
			}

		default:
			//防止取不到发生阻塞
			time.Sleep(time.Millisecond * 500)
		}
	}
}

//记录日志的函数	注释内容 放到writeLogBackGround函数中
func (f *FileLogger) myPrint(lever LogLever, format string, a ...interface{}) {
	now := time.Now()
	// //判断文件大小
	// if f.checkFileSize(f.fileObj) {
	// 	newObj := f.newFileSplit(f.fileObj)
	// 	f.fileObj = newObj
	// }

	str := fmt.Sprintf(format, a...)
	funcname, filename, line := f.getInfo(3)
	//if f.Level <= lever {
	//行号 时间 文件名 函数名 日志级别 日志内容
	//fmt.Fprintf(f.fileObj, "[%d] [%s] [%s] [%s] [%s] %s\n", line, now.Format("2006-01-02 15:04:05"), funcname, filename, flag, str)
	//需要先将日志发送到通道中去
	//创建一个logMsg对象
	logTmp := &logMsg{
		line:      line,
		timeStamp: now.Format("2006-01-02 15:04:05"),
		funcName:  funcname,
		fileName:  filename,
		flag:      lever,
		msg:       str,
	}

	//用select 保证通道未满就往通道里写日志，如果满了就抛弃走default分支防止阻塞
	select {
	case f.logChan <- logTmp:
	default:
		//走这个分支，丢弃日志防止阻塞(极端情况)
	}

	//}

	// if f.checkFileSize(f.errFileObj) {
	// 	newErrObj := f.newFileSplit(f.errFileObj)
	// 	f.errFileObj = newErrObj
	// }

	// //如果日志级别大于等于error 需要在.err文件中再次记录一遍
	// if lever >= ERROR {
	// 	//行号 时间 文件名 函数名 日志级别 日志内容
	// 	fmt.Fprintf(f.errFileObj, "[%d] [%s] [%s] [%s] [%s] %s\n", line, time.Now().Format("2006-01-02 15:04:05"), funcname, filename, flag, str)
	// }
}

//Debug日志
func (f *FileLogger) MyDebug02(format string, a ...interface{}) {
	f.myPrint(DEBUG, format, a...)
}

//Info日志
func (f *FileLogger) MyInfo02(format string, a ...interface{}) {
	f.myPrint(INFO, format, a...)
}

//Error日志
func (f *FileLogger) MyError02(format string, a ...interface{}) {
	f.myPrint(ERROR, format, a...)
}

//Warning日志
func (f *FileLogger) MyWarning02(format string, a ...interface{}) {
	f.myPrint(WARNING, format, a...)
}

//Fatal日志
func (f *FileLogger) MyFatal02(format string, a ...interface{}) {
	f.myPrint(FATAL, format, a...)
}

func (f *FileLogger) getInfo(n int) (funcname, file string, line int) {
	//Caller报告当前go线程调用栈所执行的函数的文件和行号信息。
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Println("runtime.Caller failed")
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	//path.Base函数返回路径的最后一个元素。在提取元素前会求掉末尾的斜杠。
	return strings.Split(funcName, ".")[1], path.Base(file), line
}

//判断文件大小
func (f *FileLogger) checkFileSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return false
	}

	//fmt.Println(fileInfo.Size())

	//判断当前文件大小和设置的文件最大值进行比较
	//大于等于的时候返回真，及需要扩容了
	return fileInfo.Size() >= f.maxFileSize
}

//切割文件操作
func (f *FileLogger) newFileSplit(file *os.File) (fileObj *os.File) {
	//切割日志文件
	now := time.Now()
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	logName := path.Join(f.filePath, fileInfo.Name())

	//1.关闭当前日志文件
	file.Close()
	//2.备份 xx.log -> xx.log.bak20220118
	nowStr := now.Format("20060102150405")                 //获取当前时间后缀
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //拼接新的日志文件名加路径
	//重命名
	os.Rename(logName, newLogName)
	//3.打开一个新日志文件
	fileObj, errNew := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if errNew != nil {
		fmt.Printf("open new file failed err:%v\n", errNew)
		return
	}
	//4.将打开的新日志文件复制给f.fileObj
	return fileObj
}
