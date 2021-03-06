package kafka

import (
	"fmt"
	"github.com/shopify/sarama"
	"time"
)

// 专门往kafka写日志的模块

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer // 声明一个全局的连接kafka的生产者的client
	logDataChan chan *logData
)

// Init初始化Client
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()

	// tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	// 连接kafka
	client, err = sarama.NewSyncProducer(addrs, config)

	if err != nil {
		fmt.Println("producer closed, err:", err)
	}

	// 初始化logDataChan
	logDataChan = make(chan *logData, maxSize)
	// 开启后台的goroutine从通道中取数据发往kafka
	go SendToKafka()
	return
}

// 真正往kafka发送日志的函数
func SendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{
				Topic: ld.topic,
				Value: sarama.StringEncoder(ld.data),
			}

			// 发送到kafka
			pid, offset, err := client.SendMessage(msg)

			if err != nil {
				fmt.Println("send msg failed, err:", err)
				return
			}

			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// 给外部暴露的一个函数，该函数只把日志数据发送到一个内部的channel中
func SendToChan(topic, data string) {
	msg := &logData{
		topic: topic,
		data:  data,
	}
	logDataChan <- msg
}
