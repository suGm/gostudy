package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("revc from chan, value:%v\n", v)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// cpu 调优 go tool pprof cpu.pprof
func main() {
	var isCpuPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCpuPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCpuPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}

		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)

	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err : %v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}
