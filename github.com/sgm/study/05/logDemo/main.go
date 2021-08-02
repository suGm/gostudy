package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file failed")
		return
	}

	log.SetOutput(fileObj)
	for {
		log.Println("测试日志")
		time.Sleep(time.Second * 2)
	}
}
