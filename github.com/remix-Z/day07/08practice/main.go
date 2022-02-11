package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
	使用goroutine和channel实现一个计算int64随机数各个位数和的程序
	1.开启一个goroutine循环生成int64的随机数，发送到jobChan
	2.开启24个goroutine从jobChan中取出随机数计算各个位数的和，将结果发送到resultChan
	3.主goroutine从resultChan中取到结果并打印输出到终端
*/

type resultChan struct {
	value int64
	sum   int64
}

var wg sync.WaitGroup
var ch1 = make(chan int64, 100)
var ch2 = make(chan *resultChan, 100)

//随机生成随机数，发送到channel中
func function01(ch chan<- int64) {
	defer wg.Done() //可有可无,因为在goroutine中就无限造数
	//随机数种子
	rand.Seed(time.Now().UnixNano())

	//循环
	for {
		ch <- rand.Int63()
		//time.Sleep(time.Second) //防止放太快
		time.Sleep(time.Millisecond * 500)
	}

}

//从通道1中取值，计算各位数的和发送到通道2中
func function02(ch1 <-chan int64, ch2 chan<- *resultChan) {
	defer wg.Done()

	var i int64

	//我也得循环去取么? en
	for {
		num := <-ch1
		n := num

		for n > 0 {
			i += n % 10
			n /= 10
		}

		n1 := &resultChan{
			value: num,
			sum:   i,
		}
		ch2 <- n1
	}
}

func test01() {
	//开启一个goroutine循环生成随机数
	wg.Add(1)
	go function01(ch1)

	//开启24个goroutine求各位数和
	wg.Add(24)
	for j := 0; j < 24; j++ {
		go function02(ch1, ch2)
	}

	fmt.Println("遍历ch2")
	for s := range ch2 {
		fmt.Printf("value:%d sum:%d\n", s.value, s.sum)
	}

	wg.Wait()
}

//求各位数和小工具
func toolDemo(n int) {
	var i int
	for n > 0 {
		i += n % 10
		n /= 10
	}
	fmt.Println(i)
}

func main() {
	// var i int
	// fmt.Printf("请输入:")
	// fmt.Scanln(&i)
	// toolDemo(i)

	test01()
}
