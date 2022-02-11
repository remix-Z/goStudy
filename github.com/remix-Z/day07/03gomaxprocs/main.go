package main

import (
	"fmt"
	"runtime"
	"sync"
)

//GOMAXPROCS
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

func main() {
	//默认CPU逻辑核心数，默认跑满整个CPU
	runtime.GOMAXPROCS(256) //WIN不知道为啥不行
	fmt.Println(runtime.NumCPU())	//打印CPU核心数
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
