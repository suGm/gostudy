package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"strings"
)

var producer *nsq.Producer

// 初始生产者
func initProducer(str string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(str, config)
	if err != nil {
		fmt.Printf("create producer failed,err:%v\n", err)
		return err
	}
	return nil
}

func main() {
	nsqAddress := "127.0.0.1:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		fmt.Printf("init producer failed,err:%v\n", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("read string from stdin failed, err:%v\n", err)
			return
		}
		data = strings.TrimSpace(data)
		if strings.ToUpper(data) == "Q" {
			break
		}

		err = producer.Publish("topic_demo", []byte(data))

		if err != nil {
			fmt.Printf("publish msg to nsq failed,err:%v\n", err)
			continue
		}
	}
}
