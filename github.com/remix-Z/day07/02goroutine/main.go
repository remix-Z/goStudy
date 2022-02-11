package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//并发
//Go语言层面天生支持并发,可以设置
/*
	并发:同一时间段内执行多个任务	(单个服务器)	为了充分利用处理器的核
	并行:同一时刻执行多个任务	(多台服务器同时处理多个任务)
	goroutine类似于线程

	goroutine什么时候结束?
	函数结束,对应函数创建的goroutine就结束了

*/

//waitGroup
func demo01() {
	//随机数种子
	rand.Seed(time.Now().UnixNano())

	for j := 0; j < 5; j++ {
		r1 := rand.Int()    //int64的随机数
		r2 := rand.Intn(10) //0 - 10
		fmt.Println(r1, -r2)
	}
}

func demo02(i int) {
	defer wg.Done() //计数器-1
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

//程序启动本身就是创建一个goroutine main.goroutine
func teset01() {
	// for i := 0; i < 100; i++ {
	// 	//go hello(i) //go关键字:开启一个单独的goroutine去执行这个任务
	// 	//匿名函数
	// 	go func(i int) {
	// 		fmt.Println(i) //这个i是从函数外部来的,闭包...(所以直接加个参数)
	// 	}(i)
	// }

	demo01()

	fmt.Println("test01")

	//main函数结束后，由main函数启动的goroutine也都结束了(类似linux里的init)
	time.Sleep(time.Second)

}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)    //计数器+1
		go demo02(i) //如何确定这10个goroutine都结束了
	}

	fmt.Println("main")
	//time.Sleep(time.Second) i
	wg.Wait() //等待wg计数器降为0
}
