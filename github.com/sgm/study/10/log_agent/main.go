package main

import (
	"fmt"
	"github.com/sgm/study/10/log_agent/conf"
	"github.com/sgm/study/10/log_agent/etcd"
	"github.com/sgm/study/10/log_agent/kafka"
	taillog "github.com/sgm/study/10/log_agent/tail_log"
	"github.com/sgm/study/10/log_agent/utils"
	"gopkg.in/ini.v1"
	"sync"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

// logAgent入口程序
//func run()  {
//	// 1、读取日志
//	for {
//		select {
//			case line := <- taillog.ReadChan():
//				// 2、发送到kafka
//				kafka.SendToKafka(cfg.Topic, line.Text)
//		default:
//			time.Sleep(time.Second)
//		}
//
//	}
//}

func main() {
	// 0、 加载配置文件
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Println("load ini failed, err:", err)
		return
	}

	// 1、初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Println("init kafka failed, err:", err)
		return
	}
	fmt.Println("init kafka success")

	// 2、初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}

	fmt.Println("init etcd success.")

	// 为了实现每个logagent都拉取自己独有的配置，所以要以自己的IP地址作为区分
	ipStr, err := utils.GetOutboundIP()
	if err != nil {
		panic(err)
	}

	etcdConfKey := fmt.Sprintf(cfg.EtcdConf.Key, ipStr)

	// 2.1、从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(etcdConfKey)
	if err != nil {
		fmt.Printf("getConf failed, err:%v\n", err)
		return
	}

	fmt.Println("get conf from etcd success, %v\n", logEntryConf)
	// 2.2、派一个哨兵去监视日志收集项的变化（有变化及时通知我的lkogAgent实时收集目标）
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}

	// 3、收集日志发往kafka
	taillog.Init(logEntryConf) // 因为NewConfChan访问了tskMgr的newConfChan,这个channel是在taillog.Init()中执行的初始化
	// 3.1、循环每一个日志收集项，创建TailObj
	newConfChan := taillog.NewConfChan() // 从taillog包中获取对外暴露的通道
	var wg sync.WaitGroup
	wg.Add(1)
	go etcd.WatchConf(cfg.EtcdConf.Key, newConfChan)
	wg.Wait()

	// 具体业务

	//run()
}

// kafka终端读取数据 (windows)
// kafka-console-consumer.bat --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning
