package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//resp, err := http.Get("http://127.0.0.1:9090/hello/?name=sgm&age=18")
	//if err != nil {
	//	fmt.Printf("get url failed, err:%v\n", err)
	//	return
	//}

	//data := url.Values{} // url values
	//urlObj, _ := url.Parse("http://127.0.0.1:9090/hello/?name=sgm&age=18")
	//data.Set("name", "sgm")
	//data.Set("age", "9000")
	//queryStr := data.Encode() // URL encode之后的URL
	//fmt.Println(queryStr)
	//urlObj.RawQuery = queryStr
	//req, err := http.NewRequest("GET", urlObj.String(), nil)
	//
	//resp, err := http.DefaultClient.Do(req)
	//
	//if err != nil {
	//	fmt.Println("get url failed, err :%v\n", err)
	//	return
	//}

	// 禁用keepAlive的client
	tr := &http.Transport{
		DisableKeepAlives: true, // 拉取的频率不是很频繁的话，可以禁用长链接
	}

	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("get url failed, err:%v\n", resp)
		return
	}

	defer resp.Body.Close()

	// 从resp中把从服务端中的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("read resp.Body failed, err:%v\n", err)
		return
	}

	fmt.Println(string(b))

}
