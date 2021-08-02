package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

// etcd watch
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}

	fmt.Println("connect to etcd success")

	defer cli.Close()

	// watch
	// 哨兵一直监视sgm的变化
	ch := cli.Watch(context.Background(), "sgm")

	for wresp := range ch {
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v Key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
		}
	}
}
