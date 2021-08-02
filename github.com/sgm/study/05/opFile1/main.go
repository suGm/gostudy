package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	writeFile3()
}

func writeFile() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	// write
	fileObj.Write([]byte("lianxi 111"))
	fileObj.WriteString("xixixi")

	fileObj.Close()
}

func writeFile2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	defer fileObj.Close()

	//创建一个写对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("emmmmmm\n") // 写到缓存中
	wr.Flush()                  // 将缓存中的内容写入到文件
}

func writeFile3() {
	str := "emmmm"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
}
