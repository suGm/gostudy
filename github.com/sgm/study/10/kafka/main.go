package main

import (
	"fmt"
	"github.com/shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	// tailf包使用
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要 leader 和 follow 都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success_channel返回
	// 构造一个消息
	msg := &sarama.ProducerMessage{
		Topic: "web_log",
		Value: sarama.StringEncoder("this is a test log"),
	}

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed,err:", err)
		return
	}
	fmt.Println("pid:%v offset:%v\n", pid, offset)
}
