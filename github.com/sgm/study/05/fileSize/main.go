package main

import (
	"fmt"
	"os"
)

// 1、文件对象类型
// 2、获取文件对象的详细信息

func main() {
	fileObj, err := os.Open("./protocol.go")
	if err != nil {
		fmt.Println("open file failed")
		return
	}

	// 获取文件对象类型
	fmt.Printf("%T\n", fileObj)
	// 获取文件对象详细信息
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Println("get file info failed")
		return
	}
	fmt.Println(fileInfo.Size())
}
