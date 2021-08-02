package taillog

import (
	"fmt"
	"github.com/sgm/study/10/log_agent/etcd"
	"time"
)

// tailTask管理者
type taillogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

var tskMgr *taillogMgr

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &taillogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
	}

	for _, logEntry := range logEntryConf {
		// 初始化的时候起了多少个tailtask都要记下来，为了后续判断方便
		tailTask := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailTask
	}

	go tskMgr.run()
}

// 监听自己的newConfChan，有了新的配置过来之后就做对应的处理
func (t *taillogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				if ok {
					// 原来就有，不需要操作
					continue
				} else {
					// 1、配置新增
					t.tskMap[mk] = NewTailTask(conf.Path, conf.Topic)
				}
				// 3、配置变更
			}
			// 2、配置删除
			// 找出原来t.logEntry有，但是newConf中没有的，要删掉
			for _, c1 := range t.logEntry { // 从原来的配置中依次拿出配置项
				isDelete := true
				for _, c2 := range newConf { // 去新的配置中逐一进行比较
					if c2.Path == c1.Path && c2.Topic == c1.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					// 把c1对应的tailObj给停掉
					mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
					//t.tskMap[mk] ==> tailObj
					t.tskMap[mk].cancelFunc()
				}
			}

			fmt.Println("新的配置来了:", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// 向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
