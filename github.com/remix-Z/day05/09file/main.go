package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//文件
//字节读
func readFromFile() {
	//可以相对路径 或 绝对路径
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	//需要关闭文件
	defer fileObj.Close()

	//读取文件
	// tmp := make([]byte, 128) //需要指定读取的长度
	// n, err := fileObj.Read(tmp)
	// if err != nil {
	// 	fmt.Printf("read file failed,err:%v\n", err)
	// 	return
	// }

	// fmt.Printf("读了%d个字节\n", n)
	// fmt.Println(string(tmp))

	//如果想读完
	for {
		readTmp := make([]byte, 128)
		n, err := fileObj.Read(readTmp)
		//EOF
		if err == io.EOF {
			fmt.Println("end of file")
			return
		}

		if err != nil {
			fmt.Printf("read file failed,err:%v\n", err)
			return
		}

		//fmt.Printf("本次读了%d个字节\n", n)
		fmt.Println(string(readTmp))
		if n < 128 {
			return
		}
	}
}

//bufio包
func readFromFileBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}

	//需要关闭文件
	defer fileObj.Close()

	//创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	// line, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Printf("read file failed,err:%v", err)
	// 	return
	// }
	// if err == io.EOF {
	// 	fmt.Println("end of file")
	// 	return
	// }

	// fmt.Println(line)

	//读完
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			//fmt.Println("end of file")
			return
		}

		if err != nil {
			fmt.Printf("read file failed,err:%v", err)
			return
		}

		fmt.Print(line)
	}
}

//ioutil包 更简洁
func readFromFileIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file failed,err:%v", err)
		return
	}
	fmt.Println(string(ret))
}

func main() {
	//readFromFileBufio()
	readFromFileIoutil()
}
