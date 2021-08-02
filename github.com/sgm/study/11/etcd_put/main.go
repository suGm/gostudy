package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

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

	//put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	value := `[{"path":"c:/tmp/nginx.log", "topic":"web_log"},{"path":"d:/xxx/redis/log", "topic":"web_log"}]`
	_, err = cli.Put(ctx, "/logagent/collect_config", value)
	cancel()

	if err != nil {
		fmt.Println("put to etcd failed, err:%b\n", err)
		return
	}
}
