package main

import (
	"fmt"
	"time"
)

/*
	三个工人(3个goroutine)干5个任务
*/

//worker pool (goroutine池)
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//开启3个goroutine	一般64或者128 看任务量
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	//输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
