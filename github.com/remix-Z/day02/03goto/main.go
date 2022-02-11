package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'B' {
				//break //break只能跳出当前循环
				goto LABEL1
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}

LABEL1:
	fmt.Println("over")
}

//同理 break和continue 都可以用标签
// break LABEL或是continue LABEL
