package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request) {
	str, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
		return
	}
	w.Write([]byte(str))
}

func f2(w http.ResponseWriter, r *http.Request) {
	// 对于get请求，参数都放在url上
	queryParams := r.URL.Query()
	fmt.Println(queryParams.Get("name")) // 自动帮我们识别URL中的query params
	fmt.Println(queryParams.Get("age"))  // 自动帮我们识别URL中的query params
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/posts/Go/socket/", f1)
	http.HandleFunc("/hello/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)
}
