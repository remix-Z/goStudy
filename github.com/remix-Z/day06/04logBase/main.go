package main

//日志库
//需求分析
/*
	1.支持往不同的位置输出日志
	2.日志分级别 debug trace info warning error fatal ...
	3.日志开关控制 上线之后有一些不需要打印到日志中(可有可无 注释掉就完事了)
	4.日志要记录时间 行号 文件名 日志级别 日志信息
	5.日志文件要切割(?)
		1.按文件大小切割,每次记录日志之前都要判断当前写的这个文件的文件大小、
		2.按日期切割
*/

import (
	"github.com/remix-Z/day06/mylogger"
)

func main() {
	// log := mylogger.NewLog("debug")
	// err := "ERROR!ERROR!ERROR!ERROR!ERROR!"
	// name := "大山大河"
	// i := 1000000

	// //实现的是打印到终端
	// for {
	// 	log.MyDebug("这是一条Debug日志 name:%s no:%d", name, i)
	// 	log.MyError("这是一条Error日志 err:%v", err)
	// 	log.MyInfo("这是一条Info日志")
	// 	log.MyWarning("这是一条Warning日志")
	// 	log.MyFatal("这是一条Fatal日志")
	// 	time.Sleep(2 * time.Second)
	// }

	//func NewFileLogger(leverStr, fp, fn string, maxSize int64) *FileLogger {
	fp := "./"
	fn := "mydebugDemo.log" //文件名
	fatal := "致命错误"
	log02 := mylogger.NewFileLogger("debug", fp, fn, 1024)
	for i := 0; i < 10; i++ {
		log02.MyDebug02("nihao 这是一条Debug日志")
		log02.MyInfo02("nihao 这是一条Info日志")
		log02.MyWarning02("nihao 这是一条Warning日志")
		log02.MyError02("nihao 这是一条Error日志")
		log02.MyFatal02("nihao 这是一条Fatal日志 %v", fatal)
		//time.Sleep(time.Second)
	}

	// fileObj.Close()
	// errFileObj.Close()

}
