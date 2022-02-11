/*
	为保证业务代码执行性能将之前写的日志库改写为异步记录日志的方式
	业务代码记录日志先放到通道中，然后启一个后台goroutine专门从通道中取日志，往文件中写
	&Q:
		1.日志库中channel怎么用?
		2.什么时候起后台goroutine去写日志到文件中?
*/

package main

import "github.com/remix-Z/day06/mylogger"

//var wg sync.WaitGroup

func main() {
	// fp := "./"
	// fn := "mydebugDemo.log" //文件名
	// //fatal := "致命错误"
	// log02 := mylogger.NewFileLogger("debug", fp, fn, 1024)

	// log02.MyDebug02("111")

	// str := "这是一条Debug日志"
	// len := len(str)
	// ch := make(chan interface{}, len) //容量呢?一次写入多大
	// //主要涉及的是打开文件的时机 可能会造成堵塞
	// //就是把要写入的内容放到channel中，然后再从channel读到内容写入文件
	// wg.Add(10)
	// for i := 0; i < 10; i++ {
	// 	go test01(ch, str, strconv.Itoa(i))
	// }

	// wg.Add(1)
	// go test02()

	// wg.Wait()

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
}
