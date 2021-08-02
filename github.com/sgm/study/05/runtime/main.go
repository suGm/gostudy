package main

import (
	"fmt"
	"path"
	"runtime"
)

func main() {
	getInfo(1)
}

func getInfo(n int) {
	pc, file, line, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime caller() failed")
		return
	}

	funcName := runtime.FuncForPC(pc).Name()

	fmt.Println(funcName)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(path.Base(file))

}
