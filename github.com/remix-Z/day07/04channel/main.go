package main

import (
	"fmt"
	"sync"
)

//channel
/*
	channel是一种特殊的类型
	channel是goroutine之间的连接
	遵循先入先出
	格式:var 变量 chan 类型
	引用类型 需要初始化开辟空间才能使用(make)

	通道操作
	1.发送: <-	(箭头表示数据的流向)
	2.接收: <-
	3.关闭: close()
*/

var wg sync.WaitGroup

func main() {
	var c1 chan int //需要指定通道中元素类型
	defer close(c1) //关不关都行 最终程序结束会自动回收

	fmt.Printf("%T\n", c1) //chan int

	wg.Add(1)
	c1 = make(chan int, 100) //无缓冲区通道初始化
	go func() {
		defer wg.Done()
		fmt.Printf("后台goroutine从通道b中取到数据:%d\n", recv(c1))
		//fmt.Printf("后台goroutine从通道b中取到数据:%d\n", recv(c1))	需要再次接收才能接收到第二次的20
	}()

	fmt.Println("发送到通道c1中")
	send(c1, 10)

	fmt.Println("发送到通道c1中")
	send(c1, 20)

	//c1 = make(chan int, 100) //需要初始化开辟空间才能使用(make)
	fmt.Println(c1)
	wg.Wait()

	//var c2 chan func(int, int) int
}

//通道操作
//发送
func send(c chan int, i int) chan int {
	c <- i
	return c
}

//接收
func recv(c chan int) int {
	return <-c
}
