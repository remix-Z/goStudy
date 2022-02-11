package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//打开文件写内容

func writeDemo01() {
	//fileObj, err := os.OpenFile("./test.txt", os.O_APPEND, 0644)
	//open file failed,err:open ./test.txt: The system cannot find the file specified.

	//os.O_APPEND|os.O_CREATE 位操作
	/*	底层 是16进制表示
		O_RDONLY   = 0x00000
		O_WRONLY   = 0x00001
		O_RDWR     = 0x00002
		O_CREAT    = 0x00040
		O_EXCL     = 0x00080
		O_NOCTTY   = 0x00100
		O_TRUNC    = 0x00200
		O_NONBLOCK = 0x00800
		O_APPEND   = 0x00400
		O_SYNC     = 0x01000
		O_ASYNC    = 0x02000
		O_CLOEXEC  = 0x80000
	*/
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	defer fileObj.Close()

	//Write方法
	//fileObj.Write([]byte("hello world\n"))

	//WriteString方法
	//fileObj.WriteString("hello xiaohei\n")

	//WriteAt 在指定位置写 //os: invalid use of WriteAt on file opened with O_APPEND
	//WriteAt 不能用O_APPEND标识符
	len, err := fileObj.WriteAt([]byte("hello world\n"), 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len)

}

func writeDemo02() {
	fileObj, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	defer fileObj.Close()

	//创建一个写对象
	tmp := bufio.NewWriter(fileObj)
	tmp.WriteString("hello world demo two\n") //写到缓存中 还没有到文件
	tmp.Flush()                               //将缓存中的内容写到文件

}

func writeDemo03() {
	str := "hello world demo three"
	//会truncate再写
	err := ioutil.WriteFile("./test.txt", []byte(str), 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
}

func main() {

	//writeDemo02()

	//学习syscall
	//ptr, err := syscall.CommandLineToArgv("./test.txt", 1)

	writeDemo03()

}
