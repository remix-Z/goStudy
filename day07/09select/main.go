package main

import "fmt"

//select多路复用
//select如果多个case同时满足，select会随机选择一个。对于没有case的select{}会一直等待
func main() {
	ch := make(chan int, 2)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
			fmt.Printf("******[%d]\n", i)
		}
	}
}
