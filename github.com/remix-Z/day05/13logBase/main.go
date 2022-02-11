package main

import (
	"fmt"
	"os"
	"strconv"
)

//logger.Trace()
//logger.Debug()
//logger.Warning()
//logger.Info()
//logger.Error()
//日志库
//接口:可以输出到终端01 输出到文件02 输出到kafka03
//可以往不同的输出位置记录日志
//日志分为五种级别
func main() {
	//写日志
	//fmt.Fprintf(os.Stdout, "这是一条日志记录")
	// tmp, err := os.OpenFile("./logDebug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	var logger Log
	//logger.Info01("这是一条日志记录")

	logger.route = "./logInfo.log"
	logger.Info02("这是一条日志记录")
}

type Log struct {
	name  string //日志名
	route string //日志路径
	//日志大小
}

//标准日志
func (log *Log) Info01(str string) {
	fmt.Fprintf(os.Stdout, str)
}

func (log *Log) Info02(str string) {
	tmp, err := os.OpenFile(log.route, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < 10; i++ {
		tmp.WriteString(str + strconv.Itoa(i) + "\n")
	}

}
