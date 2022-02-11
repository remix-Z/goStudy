package main

import (
	"fmt"
	"sync"
)

//优雅的循环从通道中取值
/*
	channel小练习
	1.启动goroutine，生成100个数发送到ch1
	2.启动goroutine，从ch1中取值，计算其平方放到ch2中
	3.在main中从ch2中取值
*/

var wg sync.WaitGroup
var once sync.Once

//死锁
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	//单向通道
	var ch3 <-chan int

	wg.Add(3) //3个goroutine
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)
	wg.Wait()

	for ret := range ch2 {
		fmt.Println(ret)
	}

	fmt.Printf("%T\n", ch3) //ch3类型是<-chan int
}

//单向通道 这个概念 确保这个通道只能接收值
func f1(ch1 chan<- int) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		ch1 <- i
	}

	close(ch1)
}

//ch1只能取值，ch2只能存值，限制函数参数
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()

	for {
		s, ok := <-ch1
		if !ok {
			break
		}

		ch2 <- s * s
	}

	once.Do(func() { close(ch2) }) //确保某个操作只执行一次
}
