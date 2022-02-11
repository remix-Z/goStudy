package main

import "fmt"

//关闭通道
func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)

	<-ch1
	<-ch1
	x, ok := <-ch1
	if !ok {
		fmt.Println("error", x, ok) //0 false
		return
	}
	fmt.Println(x, ok)
}
