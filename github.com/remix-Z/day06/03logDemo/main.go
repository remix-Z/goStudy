package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//logDemo
func demo() {

	fileObj, err := os.OpenFile("./logDemo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.SetOutput(fileObj)

	for {
		log.Println("这是一条测试日志")
		time.Sleep(5 * time.Second)
	}
}

//需求分析
/*
	1.支持往不同的位置输出日志
	2.日志分级别 debug trace info warning error fatal ...
	3.日志开关控制 上线之后有一些不需要打印到日志中(可有可无 注释掉就完事了) 
	4.日志要记录时间 行号 文件名 日志级别 日志信息
	5.日志文件要切割(?)
*/
