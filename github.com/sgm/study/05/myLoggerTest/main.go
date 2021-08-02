package main

import (
	"github.com/sgm/study/05/myLogger"
	"time"
)

func main() {
	log := myLogger.NewLog("debug")
	for true {
		log.Debug("这是一条debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志")
		log.Fatal("这是一条Fatal日志")
		time.Sleep(2 * time.Second)
	}
}
