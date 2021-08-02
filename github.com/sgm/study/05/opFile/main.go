package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 打开文件

func main() {
	readFromFile3()
}

// ioutil读取文件
func readFromFile3() {
	ret, err := ioutil.ReadFile("./protocol.go")

	if err != nil {
		fmt.Println("open file failed")
		return
	}

	fmt.Println(string(ret))
}

// 利用bufio这个包读取文件
func readFromFile2() {
	fileObj, err := os.Open("./protocol.go")

	if err != nil {
		fmt.Println("open file failed")
		return
	}

	defer fileObj.Close()

	// 创建一个用来从文件中读取内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("read file error")
			return
		}

		fmt.Print(line)
	}
}

func readFromFile1() {
	fileObj, err := os.Open("./protocol.go")

	if err != nil {
		fmt.Println("open file failed")
		return
	}

	// 关闭文件
	defer fileObj.Close()

	// 读文件
	var tmp = make([]byte, 128) // 指定读的长度
	//var tmp = [128]byte

	for {
		n, err := fileObj.Read(tmp[:])

		if err == io.EOF {
			fmt.Println("读完了")
			return
		}

		if err != nil {
			fmt.Println("read from file failed")
			return
		}

		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}
